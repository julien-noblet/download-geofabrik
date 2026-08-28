package osmkewllu_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/iotest"

	"github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockKewlLuHTML = `<!DOCTYPE html>
<html>
<head><title>Index of /luxembourg.osm</title></head>
<body>
<table>
  <tr><td><a name="no_href">No href</a></td></tr>
  <tr><td><a href="?C=N;O=D">Sort</a></td></tr>
  <tr><td><a href="http://external.com">External</a></td></tr>
  <tr><td><a href="https://external.com">External HTTPS</a></td></tr>
  <tr><td><a href="/root">Root</a></td></tr>
  <tr><td><a href="../">Parent</a></td></tr>
  <tr><td><a href="rdiff/">Directory</a></td></tr>
  <tr><td><a href=".hidden">Hidden</a></td></tr>
  <tr><td><a href="archive.sh">Script</a></td></tr>
  <tr><td><a href="luxembourg-diff-20260828.osc.bz2">Diff</a></td></tr>
  <tr><td><a href="luxembourg-20260828.osm.bz2">Archive</a></td></tr>
  <tr><td><a href="luxembourg-20260828.osm.pbf">Timestamped PBF</a></td></tr>
  <tr><td><a href="unknown.txt">Unknown format</a></td></tr>
  <tr><td><a href="luxembourg.osm.pbf">luxembourg.osm.pbf</a></td></tr>
  <tr><td><a href="luxembourg.osm.bz2">luxembourg.osm.bz2</a></td></tr>
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

func newProviderWithServer(ts *httptest.Server) *osmkewllu.Provider {
	p := osmkewllu.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = ts.Client()

	return p
}

func TestOSMKewlLu_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockKewlLuHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	assert.Equal(t, "osm.kewl.lu", p.Name())
	assert.Equal(t, "osm.kewl.lu.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("luxembourg"))
	assert.False(t, cat.Exist("luxembourg-20260828"))
	assert.False(t, cat.Exist("archive"))

	lux, exists := cat.Get("luxembourg")
	assert.True(t, exists)
	assert.True(t, lux.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, lux.ContainsFormat(catalog.FormatOsmBz2))
}

func TestOSMKewlLu_FetchCatalog_HTMLParseError(t *testing.T) {
	t.Parallel()

	p := osmkewllu.NewProvider()
	p.StartURL = "http://example.com"
	p.Client = &http.Client{Transport: errTransport{}}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot parse HTML")
}

func TestOSMKewlLu_FetchCatalog_HTTPError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "bad gateway", http.StatusBadGateway)
	}))
	defer ts.Close()

	p := newProviderWithServer(ts)

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmkewllu.ErrFetchCatalog)
}

func TestOSMKewlLu_FetchCatalog_InvalidURL(t *testing.T) {
	t.Parallel()

	p := osmkewllu.NewProvider()
	p.StartURL = "http://127.0.0.1:1"
	p.Client = &http.Client{}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmkewllu.ErrFetchCatalog)
}

func TestOSMKewlLu_FetchCatalog_InvalidURLSyntax(t *testing.T) {
	t.Parallel()

	p := osmkewllu.NewProvider()
	p.StartURL = "http://[::1]:namedport/"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestOSMKewlLu_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockKewlLuHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestOSMKewlLu_FetchCatalog_EmptyPage(t *testing.T) {
	t.Parallel()

	ts := newMockServer(`<!DOCTYPE html><html><body></body></html>`)
	defer ts.Close()

	p := newProviderWithServer(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.Empty(t, cat.Elements)
}

func TestOSMKewlLu_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockKewlLuHTML)
	defer ts.Close()

	p := osmkewllu.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = nil

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.True(t, cat.Exist("luxembourg"))
}

func TestOSMKewlLu_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := osmkewllu.DefaultFormats()
	assert.Contains(t, formats, catalog.FormatOsmPbf)
	assert.Contains(t, formats, catalog.FormatOsmBz2)
}

func Benchmark_OSMKewlLu_FetchCatalog_Mock(b *testing.B) {
	ts := newMockServer(mockKewlLuHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}
