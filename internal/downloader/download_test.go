package download_test

import (
	"context"
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
}
