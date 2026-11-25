package scraperutil_test

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/julien-noblet/download-geofabrik/internal/provider/scraperutil"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var errTestSentinel = errors.New("test error")

func TestCrawl(t *testing.T) {
	t.Parallel()

	t.Run("visits unique pages and deduplicates", func(t *testing.T) {
		t.Parallel()

		var (
			mu      sync.Mutex
			visited []string
		)

		cfg := scraperutil.CrawlConfig{
			StartURL:         "https://example.com/",
			ProviderName:     "test",
			PanicErr:         errTestSentinel,
			WorkChanCapacity: 10,
			ConcurrencyLimit: 2,
		}

		err := scraperutil.Crawl(t.Context(), cfg, func(_ context.Context, pageURL string) ([]string, error) {
			mu.Lock()

			visited = append(visited, pageURL)

			mu.Unlock()

			if pageURL == "https://example.com/" {
				return []string{"https://example.com/a/", "https://example.com/a/", "https://example.com/b/"}, nil
			}

			return nil, nil
		})

		require.NoError(t, err)
		assert.ElementsMatch(t, []string{"https://example.com/", "https://example.com/a/", "https://example.com/b/"}, visited)
	})

	t.Run("propagates worker error and cancels", func(t *testing.T) {
		t.Parallel()

		cfg := scraperutil.CrawlConfig{
			StartURL:         "https://example.com/",
			ProviderName:     "test",
			PanicErr:         errTestSentinel,
			WorkChanCapacity: 10,
			ConcurrencyLimit: 2,
		}

		err := scraperutil.Crawl(t.Context(), cfg, func(_ context.Context, _ string) ([]string, error) {
			return nil, errTestSentinel
		})

		require.ErrorIs(t, err, errTestSentinel)
	})

	t.Run("recovers worker panic", func(t *testing.T) {
		t.Parallel()

		cfg := scraperutil.CrawlConfig{
			StartURL:         "https://example.com/",
			ProviderName:     "test",
			PanicErr:         errTestSentinel,
			WorkChanCapacity: 10,
			ConcurrencyLimit: 2,
		}

		err := scraperutil.Crawl(t.Context(), cfg, func(_ context.Context, _ string) ([]string, error) {
			panic("boom")
		})

		require.ErrorIs(t, err, errTestSentinel)
		assert.Contains(t, err.Error(), "panic in test worker")
	})

	t.Run("respects canceled context", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(t.Context())
		cancel()

		cfg := scraperutil.CrawlConfig{
			StartURL:         "https://example.com/",
			ProviderName:     "test",
			PanicErr:         errTestSentinel,
			WorkChanCapacity: 10,
			ConcurrencyLimit: 2,
		}

		err := scraperutil.Crawl(ctx, cfg, func(_ context.Context, _ string) ([]string, error) {
			return nil, nil
		})

		require.ErrorIs(t, err, context.Canceled)
	})
}

func TestParseFlatDirectory(t *testing.T) {
	t.Parallel()

	htmlInput := `
		<html><body>
			<a href="../">Parent</a>
			<a name="anchor">NoHref</a>
			<a href="zurich.osm.pbf">zurich.osm.pbf</a>
			<a href="zurich.poly"/>
			<a href="ignored.txt">ignored.txt</a>
		</body></html>
	`

	cat := catalog.New()
	err := scraperutil.ParseFlatDirectory(
		strings.NewReader(htmlInput),
		cat,
		func(href string) bool { return strings.HasPrefix(href, "..") },
		func(href string) (string, string) {
			if id, ok := strings.CutSuffix(href, ".osm.pbf"); ok {
				return id, catalog.FormatOsmPbf
			}

			if id, ok := strings.CutSuffix(href, ".poly"); ok {
				return id, catalog.FormatPoly
			}

			return "", ""
		},
	)

	require.NoError(t, err)
	assert.Contains(t, cat.Elements, "zurich")
	assert.ElementsMatch(t, []string{catalog.FormatOsmPbf, catalog.FormatPoly}, cat.Elements["zurich"].Formats)
}
