package osmtw_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/iotest"

	"github.com/julien-noblet/download-geofabrik/internal/provider/osmtw"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockTaiwanHTML = `<!DOCTYPE html>
<html>
<head><title>Index of /download/tw-extract/</title></head>
<body>
<table>
  <tr><td><a name="no_href">No href</a></td></tr>
  <tr><td><a href="?C=N;O=D">Name sort</a></td></tr>
  <tr><td><a href="http://external.com">External</a></td></tr>
  <tr><td><a href="https://external.com">External HTTPS</a></td></tr>
  <tr><td><a href="/root">Root</a></td></tr>
  <tr><td><a href="../">Parent</a></td></tr>
  <tr><td><a href=".hidden">Hidden</a></td></tr>
  <tr><td><a href="README.txt">README.txt</a></td></tr>
  <tr><td><a href="diff/">diff/</a></td></tr>
  <tr><td><a href="basemap/">basemap/</a></td></tr>
  <tr><td><a href="recent/">recent/</a></td></tr>
  <tr><td><a href="other_dir/">other_dir/</a></td></tr>
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

func newProviderWithServer(ts *httptest.Server) *osmtw.Provider {
	p := osmtw.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = ts.Client()

	return p
}

func TestOSMTW_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockTaiwanHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	assert.Equal(t, "osm.kcwu.csie.org", p.Name())
	assert.Equal(t, "osm.kcwu.csie.org.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("taiwan"))
	assert.False(t, cat.Exist("diff"))
	assert.False(t, cat.Exist("other_dir"))

	tw, exists := cat.Get("taiwan")
	assert.True(t, exists)
	assert.Equal(t, "Taiwan", tw.Name)
	assert.True(t, tw.ContainsFormat(catalog.FormatO5m))
	assert.True(t, tw.ContainsFormat(catalog.FormatO5mZst))
}

func TestOSMTW_FetchCatalog_HTMLParseError(t *testing.T) {
	t.Parallel()

	p := osmtw.NewProvider()
	p.StartURL = "http://example.com"
	p.Client = &http.Client{Transport: errTransport{}}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot parse HTML")
}

func TestOSMTW_FetchCatalog_HTTPError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "bad gateway", http.StatusBadGateway)
	}))
	defer ts.Close()

	p := newProviderWithServer(ts)

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmtw.ErrFetchCatalog)
}

func TestOSMTW_FetchCatalog_InvalidURL(t *testing.T) {
	t.Parallel()

	p := osmtw.NewProvider()
	p.StartURL = "http://127.0.0.1:1"
	p.Client = &http.Client{}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmtw.ErrFetchCatalog)
}

func TestOSMTW_FetchCatalog_InvalidURLSyntax(t *testing.T) {
	t.Parallel()

	p := osmtw.NewProvider()
	p.StartURL = "http://[::1]:namedport/"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestOSMTW_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockTaiwanHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestOSMTW_FetchCatalog_EmptyPage(t *testing.T) {
	t.Parallel()

	ts := newMockServer(`<!DOCTYPE html><html><body></body></html>`)
	defer ts.Close()

	p := newProviderWithServer(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.Empty(t, cat.Elements)
}

func TestOSMTW_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockTaiwanHTML)
	defer ts.Close()

	p := osmtw.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = nil

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.True(t, cat.Exist("taiwan"))
}

func TestOSMTW_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := osmtw.DefaultFormats()
	assert.Contains(t, formats, catalog.FormatO5m)
	assert.Contains(t, formats, catalog.FormatO5mZst)
}

func Benchmark_OSMTW_FetchCatalog_Mock(b *testing.B) {
	ts := newMockServer(mockTaiwanHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}
