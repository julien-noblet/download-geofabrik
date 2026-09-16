package catalog_test

import (
	"path/filepath"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func BenchmarkCatalogExist(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		_ = cat.Exist("france")
	}
}

func BenchmarkCatalogGet(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		_, _ = cat.Get("france")
	}
}

func BenchmarkCatalogFind(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		_, _ = cat.Find("france")
	}
}

func BenchmarkCatalogResolveURL(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	fr, _ := cat.Find("france")

	for b.Loop() {
		_, _ = cat.ResolveURL(fr, catalog.FormatOsmPbf)
	}
}

func BenchmarkCatalogSortedKeys(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		_ = cat.SortedKeys()
	}
}

func BenchmarkCatalogLoadFile(b *testing.B) {
	for b.Loop() {
		cat, err := catalog.LoadFile(geofabrikYml)
		if err != nil {
			b.Fatal(err)
		}

		_ = cat
	}
}

func BenchmarkCatalogSaveFile(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	tmpDir := b.TempDir()
	outPath := filepath.Join(tmpDir, "out.yml")

	for b.Loop() {
		if err := cat.SaveFile(outPath); err != nil {
			b.Fatal(err)
		}
	}
}
