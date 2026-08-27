package openstreetmapfr_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type errTransport struct{}

func (errTransport) RoundTrip(_ *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(iotest.ErrReader(errors.New("read error"))),
	}, nil
}

// mock pages exercising all crawler branches:
// - root extracts files (parent = extracts / polygons)
// - subdirectories with and without trailing slashes
// - full URLs (https://)
// - duplicate directory links
// - files with no extension, hidden files (.hidden)
// - deep 4-level nesting (europe -> france -> ile_de_france -> paris).
func buildMockOSMFRHTML(baseURL string) map[string]string {
	return map[string]string{
		"/extracts/": fmt.Sprintf(`<!DOCTYPE html>
<html><body><table>
  <tr><td><a class="nav" title="help">no href tag</a></td></tr>
  <tr><td><a href="europe/">europe/</a></td></tr>
  <tr><td><a href="asia/">asia/</a></td></tr>
  <tr><td><a href="%[1]s/extracts/europe/">duplicate europe</a></td></tr>
  <tr><td><a href="direct.osm.pbf">direct.osm.pbf</a></td></tr>
  <tr><td><a href="noextension">noextension</a></td></tr>
  <tr><td><a href=".hidden">.hidden</a></td></tr>
  <tr><td><a href="../">parent</a></td></tr>
  <tr><td><a href="?C=N;O=D">sort</a></td></tr>
  <tr><td><a href="france-latest.osm.pbf">france-latest.osm.pbf</a></td></tr>
</table></body></html>`, baseURL),

		"/extracts/europe/": fmt.Sprintf(`<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="france.osm.pbf">france.osm.pbf</a></td></tr>
  <tr><td><a href="france.poly">france.poly</a></td></tr>
  <tr><td><a href="france.state.txt">france.state.txt</a></td></tr>
  <tr><td><a href="monaco.osm.pbf">monaco.osm.pbf</a></td></tr>
  <tr><td><a href="france/">france/</a></td></tr>
  <tr><td><a href="%[1]s/extracts/europe/france/">absolute france</a></td></tr>
</table></body></html>`, baseURL),

		"/extracts/europe/france/": `<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="ile_de_france/">ile_de_france/</a></td></tr>
  <tr><td><a href="south.osm.pbf">south.osm.pbf</a></td></tr>
</table></body></html>`,

		"/extracts/europe/france/ile_de_france/": `<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="paris.osm.pbf">paris.osm.pbf</a></td></tr>
  <tr><td><a href="paris.poly">paris.poly</a></td></tr>
</table></body></html>`,

		"/extracts/asia/": `<!DOCTYPE html>
<html><body><table>
  <tr><td><a href="japan.osm.pbf">japan.osm.pbf</a></td></tr>
</table></body></html>`,
	}
}

func newMockOSMFRServer() *httptest.Server {
	var pages map[string]string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		html, found := pages[r.URL.Path]
		if found {
			_, _ = w.Write([]byte(html))

			return
		}

		_, _ = w.Write([]byte(`<!DOCTYPE html><html><body><table></table></body></html>`))
	}))

	pages = buildMockOSMFRHTML(ts.URL)

	return ts
}

func TestOSMFR_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := newMockOSMFRServer()
	defer ts.Close()

	p := openstreetmapfr.NewProvider()
	p.StartURL = ts.URL + "/extracts/"
	p.BaseURL = ts.URL + "/extracts"
	p.Client = ts.Client()

	assert.Equal(t, "openstreetmap.fr", p.Name())
	assert.Equal(t, "openstreetmap.fr.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("france"))
	assert.True(t, cat.Exist("monaco"))
	assert.True(t, cat.Exist("japan"))
	assert.True(t, cat.Exist("paris"))
	assert.True(t, cat.Exist("direct"))

	fr, exists := cat.Get("france")
	assert.True(t, exists)
	assert.True(t, fr.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, fr.ContainsFormat(catalog.FormatPoly))
	assert.True(t, fr.ContainsFormat(catalog.FormatState))

	// Exception case: south under france should become france_south
	assert.True(t, cat.Exist("france_south"))

	// Deep nesting: paris should have parent ile_de_france
	paris, exists := cat.Get("paris")
	assert.True(t, exists)
	assert.Equal(t, "ile_de_france", paris.Parent)
}

func TestOSMFR_FetchCatalog_HTMLParseError(t *testing.T) {
	t.Parallel()

	p := openstreetmapfr.NewProvider()
	p.StartURL = "http://example.com"
	p.Client = &http.Client{Transport: errTransport{}}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "error parsing HTML")
}

func TestOSMFR_FetchCatalog_Errors(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "server error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	p := openstreetmapfr.NewProvider()
	p.StartURL = ts.URL + "/extracts/"
	p.Client = ts.Client()

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, openstreetmapfr.ErrFetchCatalog)
}

func TestOSMFR_FetchCatalog_InvalidURL(t *testing.T) {
	t.Parallel()

	p := openstreetmapfr.NewProvider()
	p.StartURL = "http://invalid-host-that-does-not-exist.test:9999/extracts/"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestOSMFR_FetchCatalog_InvalidURLSyntax(t *testing.T) {
	t.Parallel()

	p := openstreetmapfr.NewProvider()
	p.StartURL = "http://[::1]:namedport/extracts/"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestOSMFR_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := newMockOSMFRServer()
	defer ts.Close()

	p := openstreetmapfr.NewProvider()
	p.StartURL = ts.URL + "/extracts/"
	p.Client = ts.Client()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestOSMFR_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	p := openstreetmapfr.NewProvider()
	p.StartURL = "http://invalid-host-that-does-not-exist.test:9999/extracts/"
	p.Client = nil

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestOSMFR_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := openstreetmapfr.DefaultFormats()
	assert.Contains(t, formats, catalog.FormatOsmPbf)
	assert.Contains(t, formats, catalog.FormatPoly)
	assert.Contains(t, formats, catalog.FormatState)
	assert.Contains(t, formats, "osm.pbf.md5")
}

func Benchmark_OSMFR_FetchCatalog_Mock(b *testing.B) {
	ts := newMockOSMFRServer()
	defer ts.Close()

	p := openstreetmapfr.NewProvider()
	p.StartURL = ts.URL + "/extracts/"
	p.BaseURL = ts.URL + "/extracts"
	p.Client = ts.Client()

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}

func TestOSMFR_FetchCatalog_ContextCancelledDuringCrawl(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var baseURL string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		if r.URL.Path == "/extracts/" {
			var links strings.Builder

			links.WriteString("<!DOCTYPE html><html><body><table>")

			for i := range 60 {
				fmt.Fprintf(&links, "<tr><td><a href=\"%s/extracts/sub%d/\">Sub %d</a></td></tr>", baseURL, i, i)
			}

			links.WriteString("</table></body></html>")
			_, _ = w.Write([]byte(links.String()))

			return
		}

		// Cancel context on first subpage request
		cancel()
		time.Sleep(10 * time.Millisecond)

		_, _ = w.Write([]byte(`<!DOCTYPE html><html><body><table></table></body></html>`))
	}))
	defer ts.Close()

	baseURL = ts.URL

	p := openstreetmapfr.NewProvider()
	p.StartURL = ts.URL + "/extracts/"
	p.BaseURL = ts.URL + "/extracts"
	p.Client = ts.Client()

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}
