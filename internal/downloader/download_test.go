package download_test

import (
	"context"
	"crypto/md5" //nolint:gosec // MD5 used by Geofabrik checksum format
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	download "github.com/julien-noblet/download-geofabrik/internal/downloader"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func Test_DownloadFromURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		path        string
		fNodownload bool
		fQuiet      bool
		fProgress   bool
		wantErr     bool
	}{
		{
			name:        "try fNodownload=true",
			path:        "/this_url_should_not_exist",
			fNodownload: true,
			wantErr:     false,
		},
		{
			name:        "404 error from server",
			path:        "/this_url_should_not_exist",
			fNodownload: false,
			wantErr:     true,
		},
		{
			name:        "OK download from mock server",
			path:        "/europe/andorra.poly",
			fNodownload: false,
			fQuiet:      false,
			fProgress:   true,
			wantErr:     false,
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/europe/andorra.poly":
					_, _ = w.Write([]byte("mock poly file content"))
				default:
					http.NotFound(w, r)
				}
			}))
			defer ts.Close()

			tmpDir := t.TempDir()
			targetFile := filepath.Join(tmpDir, "download.test")

			opts := &download.Options{
				NoDownload:      thisTest.fNodownload,
				Quiet:           thisTest.fQuiet,
				Progress:        thisTest.fProgress,
				OutputDirectory: tmpDir + "/",
				FormatFlags:     make(map[string]bool),
			}
			cfg := catalog.New()

			d := download.NewDownloader(cfg, opts)

			if err := d.FromURL(context.Background(), ts.URL+thisTest.path, targetFile); (err != nil) != thisTest.wantErr {
				t.Errorf("Downloader.FromURL() error = %v, wantErr %v", err, thisTest.wantErr)
			}
		})
	}
}

func TestFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		element string
		format  string
		wantErr bool
	}{
		{
			name:    "TestFile",
			element: "africa",
			format:  catalog.FormatPoly,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/polygons/africa.poly", "//polygons/africa.poly":
					_, _ = w.Write([]byte("africa poly data"))
				default:
					http.NotFound(w, r)
				}
			}))
			defer ts.Close()

			cfg := catalog.New()
			cfg.Formats[catalog.FormatPoly] = catalog.Format{ID: catalog.FormatPoly, Loc: ".poly", BasePath: "polygons/"}
			cfg.Elements["africa"] = catalog.Element{ID: "africa", Name: "Africa", Formats: []string{catalog.FormatPoly}}
			cfg.BaseURL = ts.URL + "/"

			tmpDir := t.TempDir()
			outputFile := filepath.Join(tmpDir, "download.test")

			opts := &download.Options{
				Verbose:         true,
				OutputDirectory: tmpDir + "/",
				FormatFlags:     make(map[string]bool),
			}

			d := download.NewDownloader(cfg, opts)

			err := d.DownloadFile(context.Background(), tt.element, tt.format, outputFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChecksum(t *testing.T) {
	t.Parallel()

	testData := []byte("dummy pbf content for checksum test")
	hash := md5.Sum(testData) //nolint:gosec // MD5 used by Geofabrik checksum format
	md5Content := fmt.Sprintf("%x  monaco-latest.osm.pbf\n", hash)

	makeServer := func() *httptest.Server {
		return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/europe/monaco-latest.osm.pbf":
				_, _ = w.Write(testData)
			case "/europe/monaco-latest.osm.pbf.md5":
				_, _ = w.Write([]byte(md5Content))
			default:
				http.NotFound(w, r)
			}
		}))
	}

	tests := []struct {
		name   string
		format string
		check  bool
		want   bool
	}{
		{"No Check Poly", catalog.FormatPoly, false, false},
		{"Check Poly (no MD5 def)", catalog.FormatPoly, true, false},
		{"Check PBF", catalog.FormatOsmPbf, true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ts := makeServer()
			defer ts.Close()

			tmpDir := t.TempDir()
			cfg := catalog.New()
			cfg.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf"}
			cfg.Formats["osm.pbf.md5"] = catalog.Format{ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"}
			cfg.Formats[catalog.FormatPoly] = catalog.Format{ID: catalog.FormatPoly, Loc: ".poly"}
			cfg.Elements["monaco"] = catalog.Element{ID: "monaco", Name: "Monaco", Formats: []string{catalog.FormatOsmPbf, "osm.pbf.md5", catalog.FormatPoly}}
			cfg.BaseURL = ts.URL + "/europe"

			subOpts := &download.Options{
				Check:           tt.check,
				OutputDirectory: tmpDir + "/",
				FormatFlags:     make(map[string]bool),
			}
			subDownloader := download.NewDownloader(cfg, subOpts)
			targetFile := filepath.Join(tmpDir, "monaco.osm.pbf")

			err := subDownloader.DownloadFile(context.Background(), "monaco", catalog.FormatOsmPbf, targetFile)
			if err != nil {
				t.Fatalf("Failed setup download: %v", err)
			}

			got := subDownloader.Checksum(context.Background(), "monaco", tt.format)
			if got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("Checksum with non-existent element", func(t *testing.T) {
		t.Parallel()

		ts := makeServer()
		defer ts.Close()

		tmpDir := t.TempDir()
		cfg := catalog.New()
		cfg.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf"}
		cfg.BaseURL = ts.URL + "/europe"

		subOpts := &download.Options{Check: true, OutputDirectory: tmpDir + "/"}
		subDownloader := download.NewDownloader(cfg, subOpts)

		got := subDownloader.Checksum(context.Background(), "non_existent", catalog.FormatOsmPbf)
		if got {
			t.Errorf("Checksum() = true, want false for non-existent element")
		}
	})

	t.Run("Checksum with check=false", func(t *testing.T) {
		t.Parallel()

		ts := makeServer()
		defer ts.Close()

		tmpDir := t.TempDir()
		cfg := catalog.New()
		cfg.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf"}
		cfg.Elements["monaco"] = catalog.Element{ID: "monaco", Name: "Monaco", Formats: []string{catalog.FormatOsmPbf}}
		cfg.BaseURL = ts.URL + "/europe"

		subOpts := &download.Options{Check: false, OutputDirectory: tmpDir + "/"}
		subDownloader := download.NewDownloader(cfg, subOpts)

		got := subDownloader.Checksum(context.Background(), "monaco", catalog.FormatOsmPbf)
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

func TestChecksum_AdditionalScenarios(t *testing.T) {
	t.Parallel()

	testData := []byte("hello world osm data")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/europe/monaco-latest.osm.pbf":
			_, _ = w.Write(testData)
		case "/europe/monaco-latest.osm.pbf.md5":
			// Return a mismatched checksum
			_, _ = w.Write([]byte("00000000000000000000000000000000  monaco-latest.osm.pbf\n"))
		case "/europe/notfound.osm.pbf.md5":
			http.NotFound(w, r)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(ts.Close)

	cfg := catalog.New()
	cfg.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf"}
	cfg.Formats["osm.pbf.md5"] = catalog.Format{ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"}
	cfg.Elements["monaco"] = catalog.Element{ID: "monaco", Name: "Monaco", Formats: []string{catalog.FormatOsmPbf, "osm.pbf.md5"}}
	cfg.Elements["missing"] = catalog.Element{ID: "missing", Name: "Missing", Formats: []string{catalog.FormatOsmPbf, "osm.pbf.md5"}}
	cfg.Elements["no_fhash"] = catalog.Element{ID: "no_fhash", Name: "NoFHash", Formats: []string{catalog.FormatOsmPbf}}
	cfg.BaseURL = ts.URL + "/europe"

	t.Run("Checksum mismatch", func(t *testing.T) {
		t.Parallel()

		subTmp := t.TempDir()
		subOpts := &download.Options{Check: true, OutputDirectory: subTmp + "/"}
		subD := download.NewDownloader(cfg, subOpts)
		targetFile := filepath.Join(subTmp, "monaco.osm.pbf")

		err := subD.DownloadFile(context.Background(), "monaco", catalog.FormatOsmPbf, targetFile)
		if err != nil {
			t.Fatalf("DownloadFile failed: %v", err)
		}

		got := subD.Checksum(context.Background(), "monaco", catalog.FormatOsmPbf)
		if got {
			t.Errorf("Checksum() = true, want false when MD5 does not match")
		}
	})

	t.Run("Checksum download 404", func(t *testing.T) {
		t.Parallel()

		subTmp := t.TempDir()
		subOpts := &download.Options{Check: true, OutputDirectory: subTmp + "/"}
		subD := download.NewDownloader(cfg, subOpts)

		got := subD.Checksum(context.Background(), "missing", catalog.FormatOsmPbf)
		if got {
			t.Errorf("Checksum() = true, want false when md5 404s")
		}
	})

	t.Run("Checksum Elem2URL failure", func(t *testing.T) {
		t.Parallel()

		subTmp := t.TempDir()
		subOpts := &download.Options{Check: true, OutputDirectory: subTmp + "/"}
		subD := download.NewDownloader(cfg, subOpts)

		got := subD.Checksum(context.Background(), "no_fhash", catalog.FormatOsmPbf)
		if got {
			t.Errorf("Checksum() = true, want false when format hash not defined")
		}
	})
}

func TestDownload_ProgressBar(t *testing.T) {
	t.Parallel()

	// Large payload > 512KB to trigger progress bar
	largePayload := make([]byte, 600*1024)
	for i := range largePayload {
		largePayload[i] = byte(i % 256)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Length", strconv.Itoa(len(largePayload)))
		_, _ = w.Write(largePayload)
	}))
	defer ts.Close()

	tmpDir := t.TempDir()
	targetFile := filepath.Join(tmpDir, "large.bin")

	opts := &download.Options{
		Progress:        true,
		Quiet:           false,
		OutputDirectory: tmpDir + "/",
		FormatFlags:     make(map[string]bool),
	}
	d := download.NewDownloader(catalog.New(), opts)

	err := d.FromURL(context.Background(), ts.URL+"/large.bin", targetFile)
	if err != nil {
		t.Fatalf("FromURL with progress bar failed: %v", err)
	}
}

func TestDownload_ContextCanceled(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(100 * time.Millisecond)

		_, _ = w.Write([]byte("data"))
	}))
	defer ts.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	tmpDir := t.TempDir()
	d := download.NewDownloader(catalog.New(), &download.Options{})

	err := d.FromURL(ctx, ts.URL+"/canceled.bin", filepath.Join(tmpDir, "canceled.bin"))
	if err == nil {
		t.Errorf("expected error on canceled context, got nil")
	}
}

func TestDownloadFile_Errors(t *testing.T) {
	t.Parallel()

	cfg := catalog.New()
	cfg.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: ".osm.pbf"}
	cfg.Elements["valid_elem"] = catalog.Element{ID: "valid_elem", Formats: []string{catalog.FormatOsmPbf}}
	cfg.BaseURL = "http://invalid-host-that-does-not-exist.example.com"

	opts := &download.Options{
		OutputDirectory: t.TempDir() + "/",
		FormatFlags:     make(map[string]bool),
	}
	d := download.NewDownloader(cfg, opts)

	t.Run("Element not found", func(t *testing.T) {
		t.Parallel()

		err := d.DownloadFile(context.Background(), "unknown_id", catalog.FormatOsmPbf, filepath.Join(t.TempDir(), "out.bin"))
		if err == nil {
			t.Errorf("expected error for unknown element, got nil")
		}
	})

	t.Run("Format not found in config", func(t *testing.T) {
		t.Parallel()

		err := d.DownloadFile(context.Background(), "valid_elem", "nonexistent_format", filepath.Join(t.TempDir(), "out.bin"))
		if err == nil {
			t.Errorf("expected error for nonexistent format in config, got nil")
		}
	})

	t.Run("Format not available on element", func(t *testing.T) {
		t.Parallel()

		err := d.DownloadFile(context.Background(), "valid_elem", catalog.FormatPoly, filepath.Join(t.TempDir(), "out.bin"))
		if err == nil {
			t.Errorf("expected error for unsupported format, got nil")
		}
	})

	t.Run("Network connection error", func(t *testing.T) {
		t.Parallel()

		err := d.DownloadFile(context.Background(), "valid_elem", catalog.FormatOsmPbf, filepath.Join(t.TempDir(), "out.bin"))
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
	cfg := catalog.New()
	opts := &download.Options{}

	for range b.N {
		_ = download.NewDownloader(cfg, opts)
	}
}
