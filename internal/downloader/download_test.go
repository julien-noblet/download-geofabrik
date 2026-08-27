package download_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	download "github.com/julien-noblet/download-geofabrik/internal/downloader"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
)

func Test_DownloadFromURL(t *testing.T) {
	t.Parallel()

	type args struct {
		myURL string
	}

	tests := []struct {
		name        string
		args        args
		fNodownload bool
		fQuiet      bool
		fProgress   bool
		wantErr     bool
	}{
		{
			name: "try fNodownload=true",
			args: args{
				myURL: "https://download.geofabrik.de/this_url_should_not_exist",
			},
			fNodownload: true,
			wantErr:     false,
		},
		{
			name:        "404 error from geofabrik",
			fNodownload: false,
			args: args{
				myURL: "https://download.geofabrik.de/this_url_should_not_exist",
			},
			wantErr: true,
		},
		{
			name:        "OK download from geofabrik",
			fNodownload: false,
			fQuiet:      false,
			fProgress:   true,
			args: args{
				myURL: "https://download.geofabrik.de/europe/andorra.poly",
			},
			wantErr: false,
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			tmpDir := t.TempDir()
			targetFile := filepath.Join(tmpDir, "download.test")

			opts := &config.Options{
				NoDownload:      thisTest.fNodownload,
				Quiet:           thisTest.fQuiet,
				Progress:        thisTest.fProgress,
				OutputDirectory: tmpDir + "/",
				FormatFlags:     make(map[string]bool),
			}
			cfg := &config.Config{} // Empty config for FromURL

			d := download.NewDownloader(cfg, opts)

			if err := d.FromURL(context.Background(), thisTest.args.myURL, targetFile); (err != nil) != thisTest.wantErr {
				t.Errorf("Downloader.FromURL() error = %v, wantErr %v", err, thisTest.wantErr)
			}
		})
	}
}

func TestFile(t *testing.T) {
	t.Parallel()

	type args struct {
		configPtr *config.Config
		element   string
		format    string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestFile",
			args: args{
				configPtr: &config.Config{
					Formats: formats.FormatDefinitions{
						formats.FormatPoly: {ID: formats.FormatPoly, Loc: ".poly", ToLoc: "", BasePath: "polygons/", BaseURL: ""},
					},
					Elements: element.MapElement{
						"africa": element.Element{ID: "africa", Name: "Africa", Formats: []string{formats.FormatPoly}},
					},
					BaseURL: `https://download.openstreetmap.fr/`,
				},
				element: "africa",
				format:  formats.FormatPoly,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmpDir := t.TempDir()
			outputFile := filepath.Join(tmpDir, "download.test")

			opts := &config.Options{
				Verbose:         true,
				OutputDirectory: tmpDir + "/",
				FormatFlags:     make(map[string]bool),
			}

			d := download.NewDownloader(tt.args.configPtr, opts)

			err := d.DownloadFile(context.Background(), tt.args.element, tt.args.format, outputFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChecksum(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	cfg := &config.Config{
		Formats: formats.FormatDefinitions{
			formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: "-latest.osm.pbf"},
			"osm.pbf.md5":        {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
			formats.FormatPoly:   {ID: formats.FormatPoly, Loc: ".poly"},
		},
		Elements: element.MapElement{
			"monaco": element.Element{ID: "monaco", Name: "Monaco", Formats: []string{formats.FormatOsmPbf, "osm.pbf.md5", formats.FormatPoly}},
		},
		BaseURL: "https://download.geofabrik.de/europe",
	}

	opts := &config.Options{
		Check:           true,
		OutputDirectory: tmpDir + "/",
		FormatFlags:     make(map[string]bool),
	}

	d := download.NewDownloader(cfg, opts)

	// Download monaco first into tmpDir
	targetFile := filepath.Join(tmpDir, "monaco.osm.pbf")

	err := d.DownloadFile(context.Background(), "monaco", formats.FormatOsmPbf, targetFile)
	if err != nil {
		t.Fatalf("Failed setup download: %v", err)
	}

	// Test Checksum
	tests := []struct {
		name   string
		format string
		check  bool
		want   bool
	}{
		{"No Check Poly", formats.FormatPoly, false, false},
		{"Check Poly (no MD5 def)", formats.FormatPoly, true, false},
		{"Check PBF", formats.FormatOsmPbf, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Isolate options per parallel subtest to prevent data races
			subOpts := &config.Options{
				Check:           tt.check,
				OutputDirectory: tmpDir + "/",
				FormatFlags:     make(map[string]bool),
			}
			subDownloader := download.NewDownloader(cfg, subOpts)

			got := subDownloader.Checksum(context.Background(), "monaco", tt.format)
			if got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("Checksum with non-existent element", func(t *testing.T) {
		t.Parallel()

		subOpts := &config.Options{Check: true, OutputDirectory: tmpDir + "/"}
		subDownloader := download.NewDownloader(cfg, subOpts)

		got := subDownloader.Checksum(context.Background(), "non_existent", formats.FormatOsmPbf)
		if got {
			t.Errorf("Checksum() = true, want false for non-existent element")
		}
	})

	t.Run("Checksum with check=false", func(t *testing.T) {
		t.Parallel()

		subOpts := &config.Options{Check: false, OutputDirectory: tmpDir + "/"}
		subDownloader := download.NewDownloader(cfg, subOpts)

		got := subDownloader.Checksum(context.Background(), "monaco", formats.FormatOsmPbf)
		if got {
			t.Errorf("Checksum() = true, want false when Check is false")
		}
	})
}

func TestFileExist(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	existingFile := filepath.Join(tmpDir, "exist.txt")

	if err := os.WriteFile(existingFile, []byte("ok"), 0o600); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	tests := []struct {
		name string
		path string
		want bool
	}{
		{"Existing file", existingFile, true},
		{"Existing directory", tmpDir, true},
		{"Non-existing file", filepath.Join(tmpDir, "missing.txt"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := download.FileExist(tt.path)
			if got != tt.want {
				t.Errorf("FileExist(%s) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func TestDownloadFile_Errors(t *testing.T) {
	t.Parallel()

	cfg := &config.Config{
		Formats: formats.FormatDefinitions{
			formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: ".osm.pbf"},
		},
		Elements: element.MapElement{
			"valid_elem": element.Element{ID: "valid_elem", Formats: []string{formats.FormatOsmPbf}},
		},
		BaseURL: "http://invalid-host-that-does-not-exist.example.com",
	}

	opts := &config.Options{
		OutputDirectory: t.TempDir() + "/",
		FormatFlags:     make(map[string]bool),
	}
	d := download.NewDownloader(cfg, opts)

	t.Run("Element not found", func(t *testing.T) {
		t.Parallel()

		err := d.DownloadFile(context.Background(), "unknown_id", formats.FormatOsmPbf, filepath.Join(t.TempDir(), "out.bin"))
		if err == nil {
			t.Errorf("expected error for unknown element, got nil")
		}
	})

	t.Run("Format not available on element", func(t *testing.T) {
		t.Parallel()

		err := d.DownloadFile(context.Background(), "valid_elem", formats.FormatPoly, filepath.Join(t.TempDir(), "out.bin"))
		if err == nil {
			t.Errorf("expected error for unsupported format, got nil")
		}
	})

	t.Run("Network connection error", func(t *testing.T) {
		t.Parallel()

		err := d.DownloadFile(context.Background(), "valid_elem", formats.FormatOsmPbf, filepath.Join(t.TempDir(), "out.bin"))
		if err == nil {
			t.Errorf("expected download error, got nil")
		}
	})
}

func Benchmark_FileExist(b *testing.B) {
	tmpDir := b.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(f, []byte("data"), 0o600)

	for range b.N {
		_ = download.FileExist(f)
	}
}

func Benchmark_NewDownloader(b *testing.B) {
	cfg := &config.Config{}
	opts := &config.Options{}

	for range b.N {
		_ = download.NewDownloader(cfg, opts)
	}
}
