package osmfitvutbr_test

import (
	"testing"
)

func BenchmarkFetchCatalogMock(b *testing.B) {
	ts := newMockServer(mockFitVutbrRootHTML, mockFitVutbrSubdirHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx := b.Context()

	for b.Loop() {
		_, _ = p.FetchCatalog(ctx)
	}
}
