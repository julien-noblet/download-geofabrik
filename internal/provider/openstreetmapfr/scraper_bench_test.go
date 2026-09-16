package openstreetmapfr_test

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr"
)

func BenchmarkFetchCatalogMock(b *testing.B) {
	ts := newMockOSMFRServer()
	defer ts.Close()

	p := openstreetmapfr.NewProvider()
	p.StartURL = ts.URL + "/extracts/"
	p.BaseURL = ts.URL + "/extracts"
	p.Client = ts.Client()

	ctx := b.Context()

	for b.Loop() {
		_, _ = p.FetchCatalog(ctx)
	}
}
