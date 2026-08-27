package geo2day_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/provider/geo2day"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// buildMockHTML creates mock HTML pages with hrefs that use the test server's
// BaseURL, which is required because splitParent needs full-URL paths to
// correctly extract the parent component.
func buildMockHTML(baseURL string) map[string]string {
	return map[string]string{
		// Root page: lists continents as .html sub-pages and continent-level files.
		"/": fmt.Sprintf(`<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="%[1]s/europe.html">Europe</a></td></tr>
  <tr><td><a href="%[1]s/africa.html">Africa</a></td></tr>
  <tr><td><a href="%[1]s/europe.pbf">Europe PBF</a></td></tr>
  <tr><td><a href="%[1]s/europe.poly">Europe Poly</a></td></tr>
  <tr><td><a href="%[1]s/africa.pbf">Africa PBF</a></td></tr>
</table></body></html>`, baseURL),

		// europe.html: lists countries + a sub-page for france
		"/europe.html": fmt.Sprintf(`<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="%[1]s/europe/france.html">France</a></td></tr>
  <tr><td><a href="%[1]s/europe/france.pbf">France PBF</a></td></tr>
  <tr><td><a href="%[1]s/europe/france.poly">France Poly</a></td></tr>
  <tr><td><a href="%[1]s/europe/germany.pbf">Germany PBF</a></td></tr>
  <tr><td><a href="%[1]s/europe/germany.poly">Germany Poly</a></td></tr>
  <tr><td><a href="%[1]s/europe/spain.pbf">Spain PBF</a></td></tr>
</table></body></html>`, baseURL),

		// africa.html: lists countries
		"/africa.html": fmt.Sprintf(`<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="%[1]s/africa/morocco.pbf">Morocco PBF</a></td></tr>
  <tr><td><a href="%[1]s/africa/morocco.poly">Morocco Poly</a></td></tr>
</table></body></html>`, baseURL),

		// europe/france.html: lists regions, including an exception case
		"/europe/france.html": fmt.Sprintf(`<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="%[1]s/europe/france/ile-de-france.pbf">Ile-de-France PBF</a></td></tr>
  <tr><td><a href="%[1]s/europe/france/ile-de-france.poly">Ile-de-France Poly</a></td></tr>
  <tr><td><a href="%[1]s/europe/france/bretagne.pbf">Bretagne PBF</a></td></tr>
  <tr><td><a href="%[1]s/europe/france/guyane.pbf">Guyane PBF</a></td></tr>
</table></body></html>`, baseURL),
	}
}

// newMockGeo2DayServer creates a httptest server that routes requests
// to the appropriate mock HTML page based on URL path.
// It uses a two-phase init because the mock HTML needs the server URL
// for constructing absolute hrefs.
func newMockGeo2DayServer() *httptest.Server {
	var pages map[string]string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		html, found := pages[r.URL.Path]
		if found {
			_, _ = w.Write([]byte(html))

			return
		}

		// Return empty table for unknown sub-pages (stops crawling)
		_, _ = w.Write([]byte(`<!DOCTYPE html><html><body><table></table></body></html>`))
	}))

	pages = buildMockHTML(ts.URL)

	return ts
}

func newTestProvider(ts *httptest.Server) *geo2day.Provider {
	p := geo2day.NewProvider()
	p.StartURL = ts.URL + "/"
	p.BaseURL = ts.URL
	p.Client = ts.Client()

	return p
}

func TestGeo2Day_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := newMockGeo2DayServer()
	defer ts.Close()

	p := newTestProvider(ts)

	assert.Equal(t, "geo2day", p.Name())
	assert.Equal(t, "geo2day.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	// Root-level continent files
	assert.True(t, cat.Exist("europe"), "europe should exist from root .pbf")
	assert.True(t, cat.Exist("africa"), "africa should exist from root .pbf")

	europeElem, exists := cat.Get("europe")
	assert.True(t, exists)
	assert.True(t, europeElem.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, europeElem.ContainsFormat(catalog.FormatPoly))

	// Sub-page: europe.html -> countries
	assert.True(t, cat.Exist("france"), "france should exist from europe sub-page")
	assert.True(t, cat.Exist("germany"), "germany should exist from europe sub-page")
	assert.True(t, cat.Exist("spain"), "spain should exist from europe sub-page")

	fr, exists := cat.Get("france")
	assert.True(t, exists)
	assert.True(t, fr.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, fr.ContainsFormat(catalog.FormatPoly))

	germany, exists := cat.Get("germany")
	assert.True(t, exists)
	assert.True(t, germany.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, germany.ContainsFormat(catalog.FormatPoly))

	// Sub-page: africa.html -> countries
	assert.True(t, cat.Exist("morocco"), "morocco should exist from africa sub-page")

	morocco, exists := cat.Get("morocco")
	assert.True(t, exists)
	assert.True(t, morocco.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, morocco.ContainsFormat(catalog.FormatPoly))

	// Third-level page: europe/france.html -> regions
	assert.True(t, cat.Exist("ile-de-france"), "ile-de-france should exist from france sub-page")
	assert.True(t, cat.Exist("bretagne"), "bretagne should exist from france sub-page")

	idf, exists := cat.Get("ile-de-france")
	assert.True(t, exists)
	assert.True(t, idf.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, idf.ContainsFormat(catalog.FormatPoly))

	// Exception mapping: guyane under france should become guyane-france
	assert.True(t, cat.Exist("guyane-france"), "guyane under france should be renamed to guyane-france")
}

func TestGeo2Day_FetchCatalog_Exceptions(t *testing.T) {
	t.Parallel()

	// Serve a page with guyane under south-america to test exception renaming
	var baseURL string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = fmt.Fprintf(w, `<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="%s/south-america/guyane.pbf">Guyane PBF</a></td></tr>
  <tr><td><a href="%s/south-america/guyane.poly">Guyane Poly</a></td></tr>
</table></body></html>`, baseURL, baseURL)
	}))
	defer ts.Close()

	baseURL = ts.URL

	p := newTestProvider(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	// guyane under south-america should become guyane-south-america
	assert.True(t, cat.Exist("guyane-south-america"), "guyane under south-america should be renamed")
}

func TestGeo2Day_FetchCatalog_HTTPError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "server error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	p := newTestProvider(ts)

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, geo2day.ErrFetchCatalog)
}

func TestGeo2Day_FetchCatalog_InvalidURL(t *testing.T) {
	t.Parallel()

	p := geo2day.NewProvider()
	p.StartURL = "http://invalid-host-that-does-not-exist.test:9999/"
	p.BaseURL = "http://invalid-host-that-does-not-exist.test:9999"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestGeo2Day_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := newMockGeo2DayServer()
	defer ts.Close()

	p := newTestProvider(ts)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestGeo2Day_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := geo2day.DefaultFormats()
	assert.Contains(t, formats, catalog.FormatOsmPbf)
	assert.Contains(t, formats, catalog.FormatPoly)
	assert.Contains(t, formats, catalog.FormatGeoJSON)
	assert.Contains(t, formats, "osm.pbf.md5")
}

func TestGeo2Day_FetchCatalog_MD5Extension(t *testing.T) {
	t.Parallel()

	var baseURL string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = fmt.Fprintf(w, `<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="%s/europe/france.pbf">France PBF</a></td></tr>
</table></body></html>`, baseURL)
	}))
	defer ts.Close()

	baseURL = ts.URL

	p := newTestProvider(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)

	fr, exists := cat.Get("france")
	assert.True(t, exists)
	assert.True(t, fr.ContainsFormat(catalog.FormatOsmPbf), "should have osm.pbf format")
	assert.True(t, fr.ContainsFormat("osm.pbf.md5"), "should automatically add .md5 for .pbf files")
}

func TestGeo2Day_FetchCatalog_EmptyTable(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<!DOCTYPE html><html><body><table></table></body></html>`))
	}))
	defer ts.Close()

	p := newTestProvider(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.Empty(t, cat.Elements)
}

func TestGeo2Day_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	p := geo2day.NewProvider()
	p.StartURL = "http://invalid-host-that-does-not-exist.test:9999/"
	p.BaseURL = "http://invalid-host-that-does-not-exist.test:9999"
	p.Client = nil

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestGeo2Day_FetchCatalog_ExternalLinksIgnored(t *testing.T) {
	t.Parallel()

	var baseURL string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = fmt.Fprintf(w, `<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="https://external-site.com/europe.html">External</a></td></tr>
  <tr><td><a href="%s/europe/france.pbf">France PBF</a></td></tr>
</table></body></html>`, baseURL)
	}))
	defer ts.Close()

	baseURL = ts.URL

	p := newTestProvider(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	// External .html links should not be crawled, but france.pbf should be present
	assert.True(t, cat.Exist("france"), "france should exist")
}

func Benchmark_Geo2Day_FetchCatalog_Mock(b *testing.B) {
	ts := newMockGeo2DayServer()
	defer ts.Close()

	p := newTestProvider(ts)

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}
