package downloader_test

import (
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/downloader"
)

func BenchmarkComputeMD5Hash(b *testing.B) {
	for b.Loop() {
		if _, err := downloader.ComputeMD5Hash("../../LICENSE"); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkCheckFileHash(b *testing.B) {
	oldLogger := slog.Default()

	slog.SetDefault(slog.New(slog.DiscardHandler))
	b.Cleanup(func() {
		slog.SetDefault(oldLogger)
	})

	hash, err := downloader.ComputeMD5Hash("../../LICENSE")
	if err != nil {
		b.Fatal(err)
	}

	tmpDir := b.TempDir()
	hashfile := filepath.Join(tmpDir, "test.hash")

	if err := os.WriteFile(hashfile, []byte(hash), 0o600); err != nil {
		b.Fatalf("Can't write file %s err: %v", hashfile, err)
	}

	for b.Loop() {
		if _, err := downloader.CheckFileHash(hashfile, hash); err != nil {
			b.Error(err.Error())
		}
	}
}
