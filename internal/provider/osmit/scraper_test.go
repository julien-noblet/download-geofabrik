package osmit_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/provider/osmit"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockOSMItHTML = `<!DOCTYPE html><html><head><title>Estratti OSM Italia</title></head><body>OK</body></html>`

func newMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(mockOSMItHTML))
	}))
}

func newProviderWithServer(ts *httptest.Server) *osmit.Provider {
	p := osmit.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = ts.Client()

	return p
}

func TestOSMIt_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := newMockServer()
	defer ts.Close()

	p := newProviderWithServer(ts)

	assert.Equal(t, "osmit-estratti", p.Name())
	assert.Equal(t, "osmit-estratti.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("piemonte"))
	assert.True(t, cat.Exist("lazio"))
	assert.True(t, cat.Exist("sicilia"))
	assert.True(t, cat.Exist("torino"))
	assert.True(t, cat.Exist("roma"))
	assert.True(t, cat.Exist("milano"))

	piemonte, exists := cat.Get("piemonte")
	assert.True(t, exists)
	assert.Equal(t, "Piemonte", piemonte.Name)
	assert.True(t, piemonte.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, piemonte.ContainsFormat(catalog.FormatGPKG))
	assert.True(t, piemonte.ContainsFormat(catalog.FormatOBF))
	assert.True(t, piemonte.ContainsFormat(catalog.FormatGarminOSM))

	torino, exists := cat.Get("torino")
	assert.True(t, exists)
	assert.Equal(t, "Torino", torino.Name)
	assert.True(t, torino.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, torino.ContainsFormat(catalog.FormatGPKG))
}

func TestOSMIt_FetchCatalog_HTTPError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "bad gateway", http.StatusBadGateway)
	}))
	defer ts.Close()

	p := newProviderWithServer(ts)

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmit.ErrFetchCatalog)
}

func TestOSMIt_FetchCatalog_InvalidURL(t *testing.T) {
	t.Parallel()

	p := osmit.NewProvider()
	p.StartURL = "http://127.0.0.1:1"
	p.Client = &http.Client{}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmit.ErrFetchCatalog)
}

func TestOSMIt_FetchCatalog_InvalidURLSyntax(t *testing.T) {
	t.Parallel()

	p := osmit.NewProvider()
	p.StartURL = "http://[::1]:namedport/"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestOSMIt_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := newMockServer()
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestOSMIt_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	ts := newMockServer()
	defer ts.Close()

	p := osmit.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = nil

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.True(t, cat.Exist("roma"))
}

func TestOSMIt_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := osmit.DefaultFormats()
	assert.Contains(t, formats, catalog.FormatOsmPbf)
	assert.Contains(t, formats, catalog.FormatGPKG)
	assert.Contains(t, formats, catalog.FormatOBF)
	assert.Contains(t, formats, catalog.FormatGarminOSM)
}

func Benchmark_OSMIt_FetchCatalog_Mock(b *testing.B) {
	ts := newMockServer()
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}
