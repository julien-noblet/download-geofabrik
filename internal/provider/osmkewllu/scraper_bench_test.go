package osmkewllu_test

import (
	"testing"
)

func BenchmarkFetchCatalogMock(b *testing.B) {
	ts := newMockServer(mockKewlLuHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx := b.Context()

	for b.Loop() {
		_, _ = p.FetchCatalog(ctx)
	}
}
