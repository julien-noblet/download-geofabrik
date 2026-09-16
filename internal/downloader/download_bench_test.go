package downloader_test

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/downloader"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func BenchmarkFileExists(b *testing.B) {
	tmpDir := b.TempDir()
	f := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(f, []byte("data"), 0o600)

	for b.Loop() {
		_ = downloader.FileExists(f)
	}
}

func BenchmarkNewDownloader(b *testing.B) {
	cfg := catalog.New()
	opts := &downloader.Options{}

	for b.Loop() {
		_ = downloader.New(cfg, opts)
	}
}

func BenchmarkDownloadFileStreamMD5(b *testing.B) {
	payload := bytes.Repeat([]byte("0123456789abcdef"), 4096) // 64 KB

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Length", "65536")
		_, _ = w.Write(payload)
	}))
	defer server.Close()

	tmpDir := b.TempDir()
	oldLogger := slog.Default()

	slog.SetDefault(slog.New(slog.DiscardHandler))
	b.Cleanup(func() {
		slog.SetDefault(oldLogger)
	})

	d := downloader.New(catalog.New(), &downloader.Options{
		Progress: false,
		Quiet:    true,
	})
	ctx := context.Background()

	b.SetBytes(int64(len(payload)))

	for b.Loop() {
		targetFile := filepath.Join(tmpDir, "download.bin")

		if err := d.FromURL(ctx, server.URL+"/bench.bin", targetFile); err != nil {
			b.Fatal(err)
		}

		_ = os.Remove(targetFile)
	}
}
