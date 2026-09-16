package geofabrik_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik"
)

func BenchmarkFetchCatalogMock(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockGeofabrikIndexJSON))
	}))
	defer ts.Close()

	p := geofabrik.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	ctx := b.Context()

	for b.Loop() {
		_, _ = p.FetchCatalog(ctx)
	}
}
