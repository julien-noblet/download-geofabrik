package ui_test

import (
	"io"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/ui"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func BenchmarkPrintTableStandard(b *testing.B) {
	cat, err := catalog.LoadFile("../../geofabrik.yml")
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		if err := ui.PrintTable(cat, false, io.Discard); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPrintTableMarkdown(b *testing.B) {
	cat, err := catalog.LoadFile("../../geofabrik.yml")
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		if err := ui.PrintTable(cat, true, io.Discard); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPrintJSON(b *testing.B) {
	cat, err := catalog.LoadFile("../../geofabrik.yml")
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		if err := ui.PrintJSON(cat, io.Discard); err != nil {
			b.Fatal(err)
		}
	}
}
