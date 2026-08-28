package osmch_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/iotest"

	"github.com/julien-noblet/download-geofabrik/internal/provider/osmch"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockOSMCHHTML = `<!DOCTYPE html>
<html>
<head><title>Index of /</title></head>
<body>
<table>
  <tr><td><a name="no_href">No href</a></td></tr>
  <tr><td><a href="?C=N;O=D">Name sort</a></td></tr>
  <tr><td><a href="http://osm.org">External</a></td></tr>
  <tr><td><a href="https://example.com">External HTTPS</a></td></tr>
  <tr><td><a href="/root/path">Root</a></td></tr>
  <tr><td><a href="../">Parent</a></td></tr>
  <tr><td><a href="replication/">Directory</a></td></tr>
  <tr><td><a href=".hidden">Hidden</a></td></tr>
  <tr><td><a href="unknown.txt">Unknown format</a></td></tr>
  <tr><td><a href="switzerland-exact.osm.pbf">switzerland-exact.osm.pbf</a></td></tr>
  <tr><td><a href="switzerland-exact.poly">switzerland-exact.poly</a></td></tr>
  <tr><td><a href="switzerland-padded.osm.pbf">switzerland-padded.osm.pbf</a></td></tr>
  <tr><td><a href="switzerland-padded.poly">switzerland-padded.poly</a></td></tr>
  <tr><td><a href="switzerland.obf">switzerland.obf</a></td></tr>
  <tr><td><a href="switzerland.pbf">switzerland.pbf</a></td></tr>
  <tr><td><a href="switzerland-garmin.zip">switzerland-garmin.zip</a></td></tr>
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

func newProviderWithServer(ts *httptest.Server) *osmch.Provider {
	p := osmch.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = ts.Client()

	return p
}

func TestOSMCH_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockOSMCHHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	assert.Equal(t, "planet.osm.ch", p.Name())
	assert.Equal(t, "planet.osm.ch.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("switzerland-exact"))
	assert.True(t, cat.Exist("switzerland-padded"))
	assert.True(t, cat.Exist("switzerland"))
	assert.False(t, cat.Exist("unknown"))

	exact, exists := cat.Get("switzerland-exact")
	assert.True(t, exists)
	assert.True(t, exact.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, exact.ContainsFormat(catalog.FormatPoly))

	padded, exists := cat.Get("switzerland-padded")
	assert.True(t, exists)
	assert.True(t, padded.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, padded.ContainsFormat(catalog.FormatPoly))

	ch, exists := cat.Get("switzerland")
	assert.True(t, exists)
	assert.True(t, ch.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, ch.ContainsFormat(catalog.FormatOBF))
	assert.True(t, ch.ContainsFormat(catalog.FormatGarminOSM))
}

func TestOSMCH_FetchCatalog_HTMLParseError(t *testing.T) {
	t.Parallel()

	p := osmch.NewProvider()
	p.StartURL = "http://example.com"
	p.Client = &http.Client{Transport: errTransport{}}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot parse HTML")
}

func TestOSMCH_FetchCatalog_HTTPError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "bad gateway", http.StatusBadGateway)
	}))
	defer ts.Close()

	p := newProviderWithServer(ts)

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmch.ErrFetchCatalog)
}

func TestOSMCH_FetchCatalog_InvalidURL(t *testing.T) {
	t.Parallel()

	p := osmch.NewProvider()
	p.StartURL = "http://127.0.0.1:1"
	p.Client = &http.Client{}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmch.ErrFetchCatalog)
}

func TestOSMCH_FetchCatalog_InvalidURLSyntax(t *testing.T) {
	t.Parallel()

	p := osmch.NewProvider()
	p.StartURL = "http://[::1]:namedport/"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestOSMCH_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockOSMCHHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestOSMCH_FetchCatalog_EmptyPage(t *testing.T) {
	t.Parallel()

	ts := newMockServer(`<!DOCTYPE html><html><body></body></html>`)
	defer ts.Close()

	p := newProviderWithServer(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.Empty(t, cat.Elements)
}

func TestOSMCH_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockOSMCHHTML)
	defer ts.Close()

	p := osmch.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = nil

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.True(t, cat.Exist("switzerland"))
}

func TestOSMCH_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := osmch.DefaultFormats()
	assert.Contains(t, formats, catalog.FormatOsmPbf)
	assert.Contains(t, formats, catalog.FormatPoly)
	assert.Contains(t, formats, catalog.FormatOBF)
	assert.Contains(t, formats, catalog.FormatGarminOSM)
}

func Benchmark_OSMCH_FetchCatalog_Mock(b *testing.B) {
	ts := newMockServer(mockOSMCHHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}
