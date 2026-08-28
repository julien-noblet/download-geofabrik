package osmfitvutbr_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/iotest"

	"github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockFitVutbrHTML = `<!DOCTYPE html>
<html>
<head><title>Index of /extracts</title></head>
<body>
<table>
  <tr><td><a name="no_href">No href</a></td></tr>
  <tr><td><a href="?C=N;O=D">Name sort</a></td></tr>
  <tr><td><a href="http://external.com">External</a></td></tr>
  <tr><td><a href="https://external.com">External HTTPS</a></td></tr>
  <tr><td><a href="/extracts/">Parent Directory</a></td></tr>
  <tr><td><a href="../">Parent</a></td></tr>
  <tr><td><a href=".hidden">Hidden</a></td></tr>
  <tr><td><a href="README.txt">README.txt</a></td></tr>
  <tr><td><a href="v6-planet-060403.osm.bz2">Legacy planet</a></td></tr>
  <tr><td><a href="czech-republic.poly">czech-republic.poly</a></td></tr>
  <tr><td><a href="czech_republic/">czech_republic/</a></td></tr>
  <tr><td><a href="slovakia/">slovakia/</a></td></tr>
  <tr><td><a href="custom__area/">custom__area/</a></td></tr>
</table>
</body>
</html>`

type errTransport struct{}

func (errTransport) RoundTrip(_ *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(iotest.ErrReader(errors.New("read error"))),
	}, nil
}

func newMockServer(body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(body))
	}))
}

func newProviderWithServer(ts *httptest.Server) *osmfitvutbr.Provider {
	p := osmfitvutbr.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = ts.Client()

	return p
}

func TestOSMFitVutbr_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockFitVutbrHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	assert.Equal(t, "osm.fit.vutbr.cz", p.Name())
	assert.Equal(t, "osm.fit.vutbr.cz.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("czech_republic"))
	assert.True(t, cat.Exist("slovakia"))
	assert.False(t, cat.Exist("README"))

	cz, exists := cat.Get("czech_republic")
	assert.True(t, exists)
	assert.Equal(t, "Czech Republic", cz.Name)
	assert.True(t, cz.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, cz.ContainsFormat(catalog.FormatOsmBz2))
	assert.True(t, cz.ContainsFormat(catalog.FormatPoly))
}

func TestOSMFitVutbr_FetchCatalog_HTMLParseError(t *testing.T) {
	t.Parallel()

	p := osmfitvutbr.NewProvider()
	p.StartURL = "http://example.com"
	p.Client = &http.Client{Transport: errTransport{}}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot parse HTML")
}

func TestOSMFitVutbr_FetchCatalog_HTTPError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "bad gateway", http.StatusBadGateway)
	}))
	defer ts.Close()

	p := newProviderWithServer(ts)

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmfitvutbr.ErrFetchCatalog)
}

func TestOSMFitVutbr_FetchCatalog_InvalidURL(t *testing.T) {
	t.Parallel()

	p := osmfitvutbr.NewProvider()
	p.StartURL = "http://127.0.0.1:1"
	p.Client = &http.Client{}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmfitvutbr.ErrFetchCatalog)
}

func TestOSMFitVutbr_FetchCatalog_InvalidURLSyntax(t *testing.T) {
	t.Parallel()

	p := osmfitvutbr.NewProvider()
	p.StartURL = "http://[::1]:namedport/"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestOSMFitVutbr_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockFitVutbrHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestOSMFitVutbr_FetchCatalog_EmptyPage(t *testing.T) {
	t.Parallel()

	ts := newMockServer(`<!DOCTYPE html><html><body></body></html>`)
	defer ts.Close()

	p := newProviderWithServer(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.Empty(t, cat.Elements)
}

func TestOSMFitVutbr_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockFitVutbrHTML)
	defer ts.Close()

	p := osmfitvutbr.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = nil

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.True(t, cat.Exist("czech_republic"))
}

func TestOSMFitVutbr_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := osmfitvutbr.DefaultFormats()
	assert.Contains(t, formats, catalog.FormatOsmPbf)
	assert.Contains(t, formats, catalog.FormatOsmBz2)
	assert.Contains(t, formats, catalog.FormatPoly)
}

func Benchmark_OSMFitVutbr_FetchCatalog_Mock(b *testing.B) {
	ts := newMockServer(mockFitVutbrHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}
