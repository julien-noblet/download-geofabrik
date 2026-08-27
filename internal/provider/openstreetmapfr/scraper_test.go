package openstreetmapfr_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockOSMFRIndexHTML = `<!DOCTYPE html>
<html>
<body>
<table>
  <tr><td><a href="europe/">europe/</a></td></tr>
  <tr><td><a href="asia/">asia/</a></td></tr>
  <tr><td><a href="../">parent</a></td></tr>
  <tr><td><a href="?C=N;O=D">sort</a></td></tr>
  <tr><td><a href="france-latest.osm.pbf">france-latest.osm.pbf</a></td></tr>
</table>
</body>
</html>`

const mockOSMFREuropeHTML = `<!DOCTYPE html>
<html>
<body>
<table>
  <tr><td><a href="france.osm.pbf">france.osm.pbf</a></td></tr>
  <tr><td><a href="france.poly">france.poly</a></td></tr>
  <tr><td><a href="france.state.txt">france.state.txt</a></td></tr>
  <tr><td><a href="monaco.osm.pbf">monaco.osm.pbf</a></td></tr>
  <tr><td><a href="france/">france/</a></td></tr>
</table>
</body>
</html>`

const mockOSMFRFranceHTML = `<!DOCTYPE html>
<html>
<body>
<table>
  <tr><td><a href="ile_de_france.osm.pbf">ile_de_france.osm.pbf</a></td></tr>
  <tr><td><a href="ile_de_france.poly">ile_de_france.poly</a></td></tr>
  <tr><td><a href="south.osm.pbf">south.osm.pbf</a></td></tr>
</table>
</body>
</html>`

const mockOSMFRAsiaHTML = `<!DOCTYPE html>
<html>
<body>
<table>
  <tr><td><a href="japan.osm.pbf">japan.osm.pbf</a></td></tr>
</table>
</body>
</html>`

func newMockOSMFRServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		switch r.URL.Path {
		case "/extracts/", "/extracts":
			_, _ = w.Write([]byte(mockOSMFRIndexHTML))
		case "/extracts/europe/", "/extracts/europe":
			_, _ = w.Write([]byte(mockOSMFREuropeHTML))
		case "/extracts/europe/france/", "/extracts/europe/france":
			_, _ = w.Write([]byte(mockOSMFRFranceHTML))
		case "/extracts/asia/", "/extracts/asia":
			_, _ = w.Write([]byte(mockOSMFRAsiaHTML))
		default:
			_, _ = w.Write([]byte(`<!DOCTYPE html><html><body><table></table></body></html>`))
		}
	}))
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
	assert.True(t, cat.Exist("ile_de_france"))

	fr, exists := cat.Get("france")
	assert.True(t, exists)
	assert.True(t, fr.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, fr.ContainsFormat(catalog.FormatPoly))
	assert.True(t, fr.ContainsFormat(catalog.FormatState))

	// Exception case: south under france should become france_south
	assert.True(t, cat.Exist("france_south"))
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
