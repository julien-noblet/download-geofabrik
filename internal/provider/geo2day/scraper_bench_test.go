package geo2day_test

import (
	"testing"
)

func BenchmarkFetchCatalogMock(b *testing.B) {
	ts := newMockGeo2DayServer()
	defer ts.Close()

	p := newTestProvider(ts)

	ctx := b.Context()

	for b.Loop() {
		_, _ = p.FetchCatalog(ctx)
	}
}
