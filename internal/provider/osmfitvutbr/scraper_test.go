package osmfitvutbr_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

const mockFitVutbrRootHTML = `<!DOCTYPE html>
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
</table>
</body>
</html>`

const mockFitVutbrSubdirHTML = `<!DOCTYPE html>
<html>
<head><title>Index of /extracts/czech_republic</title></head>
<body>
<table>
  <tr><td><a name="no_href">No href</a></td></tr>
  <tr><td><a href="/extracts/">Parent Directory</a></td></tr>
  <tr><td><a href="?C=N;O=D">Name sort</a></td></tr>
  <tr><td><a href="README.txt">README.txt</a></td></tr>
  <tr><td><a href="czech_republic-2006-04-03.osm.bz2">2006-04-03</a></td></tr>
  <tr><td><a href="czech_republic-2006-05-01.osm.bz2">2006-05-01</a></td></tr>
  <tr><td><a href="czech_republic-2026-09-05.osm.pbf">2026-09-05</a></td></tr>
  <tr><td><a href="czech_republic-2026-09-06.osm.pbf">2026-09-06</a></td></tr>
  <tr><td><a href="czech_republic-.osm.pbf">Empty date</a></td></tr>
  <tr><td><a href="unrelated_file.bin">Unrelated</a></td></tr>
  <tr><td><a href="czech_republic.osm.pbf">czech_republic.osm.pbf</a></td></tr>
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

func newMockServer(rootHTML, subdirHTML string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		if strings.HasSuffix(r.URL.Path, "/czech_republic/") {
			_, _ = w.Write([]byte(subdirHTML))

			return
		}

		_, _ = w.Write([]byte(rootHTML))
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

	ts := newMockServer(mockFitVutbrRootHTML, mockFitVutbrSubdirHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	assert.Equal(t, "osm.fit.vutbr.cz", p.Name())
	assert.Equal(t, "osm.fit.vutbr.cz.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("czech_republic"))
	assert.False(t, cat.Exist("README"))

	cz, exists := cat.Get("czech_republic")
	assert.True(t, exists)
	assert.Equal(t, "Czech Republic", cz.Name)
	assert.True(t, cz.ContainsFormat(catalog.FormatPoly))
	assert.False(t, cz.ContainsFormat(catalog.FormatOsmPbf))

	dateElem, exists := cat.Get("2026-09-06")
	assert.True(t, exists)
	assert.Equal(t, "Czech Republic 2026-09-06", dateElem.Name)
	assert.Equal(t, "czech_republic-2026-09-06", dateElem.File)
	assert.True(t, dateElem.ContainsFormat(catalog.FormatOsmPbf))

	bz2Elem, exists := cat.Get("2006-04-03")
	assert.True(t, exists)
	assert.Equal(t, "Czech Republic 2006-04-03", bz2Elem.Name)
	assert.Equal(t, "czech_republic-2006-04-03", bz2Elem.File)
	assert.True(t, bz2Elem.ContainsFormat(catalog.FormatOsmBz2))

	latestElem, exists := cat.Get("latest")
	assert.True(t, exists)
	assert.Equal(t, "Czech Republic (latest)", latestElem.Name)
	assert.Equal(t, "czech_republic-2026-09-06", latestElem.File)
	assert.True(t, latestElem.ContainsFormat(catalog.FormatOsmPbf))
}

func TestOSMFitVutbr_FetchCatalog_OnlyBz2(t *testing.T) {
	t.Parallel()

	onlyBz2HTML := `<!DOCTYPE html><html><body><table>
		<tr><td><a href="czech_republic-2006-04-03.osm.bz2">2006-04-03</a></td></tr>
	</table></body></html>`

	ts := newMockServer(mockFitVutbrRootHTML, onlyBz2HTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)

	latestElem, exists := cat.Get("latest")
	assert.True(t, exists)
	assert.Equal(t, "Czech Republic (latest)", latestElem.Name)
	assert.Equal(t, "czech_republic-2006-04-03", latestElem.File)
	assert.True(t, latestElem.ContainsFormat(catalog.FormatOsmBz2))
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

func TestOSMFitVutbr_FetchCatalog_SubdirHTMLParseError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		if strings.HasSuffix(r.URL.Path, "/czech_republic/") {
			w.Header().Set("Content-Length", "100")
			_, _ = w.Write([]byte("<unclosed"))

			return
		}

		_, _ = w.Write([]byte(mockFitVutbrRootHTML))
	}))
	defer ts.Close()

	p := newProviderWithServer(ts)
	client := ts.Client()
	originalTransport := client.Transport
	client.Transport = roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		if strings.HasSuffix(req.URL.Path, "/czech_republic/") {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(iotest.ErrReader(errors.New("subdir read error"))),
			}, nil
		}

		return originalTransport.RoundTrip(req)
	})
	p.Client = client

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

func TestOSMFitVutbr_FetchCatalog_SubdirHTTPError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/czech_republic/") {
			http.Error(w, "not found", http.StatusNotFound)

			return
		}

		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(mockFitVutbrRootHTML))
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

func TestOSMFitVutbr_FetchCatalog_SubdirInvalidURLSyntax(t *testing.T) {
	t.Parallel()

	ts := newMockServer(`<!DOCTYPE html><html><body><table><tr><td><a href="subdir/">subdir/</a></td></tr></table></body></html>`, "")
	defer ts.Close()

	p := newProviderWithServer(ts)
	p.StartURL = "http://[::1]:namedport"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestOSMFitVutbr_FetchCatalog_SubdirNetworkError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(mockFitVutbrRootHTML))
	}))
	defer ts.Close()

	p := newProviderWithServer(ts)
	client := ts.Client()
	originalTransport := client.Transport
	client.Transport = roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		if strings.HasSuffix(req.URL.Path, "/czech_republic/") {
			return nil, errors.New("connection reset")
		}

		return originalTransport.RoundTrip(req)
	})
	p.Client = client

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, osmfitvutbr.ErrFetchCatalog)
}

func TestOSMFitVutbr_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockFitVutbrRootHTML, mockFitVutbrSubdirHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestOSMFitVutbr_FetchCatalog_EmptyPage(t *testing.T) {
	t.Parallel()

	ts := newMockServer(`<!DOCTYPE html><html><body></body></html>`, "")
	defer ts.Close()

	p := newProviderWithServer(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.Empty(t, cat.Elements)
}

func TestOSMFitVutbr_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockFitVutbrRootHTML, mockFitVutbrSubdirHTML)
	defer ts.Close()

	p := osmfitvutbr.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = nil

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.True(t, cat.Exist("czech_republic"))
	assert.True(t, cat.Exist("2026-09-06"))
}

func TestOSMFitVutbr_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := osmfitvutbr.DefaultFormats()
	assert.Contains(t, formats, catalog.FormatOsmPbf)
	assert.Contains(t, formats, catalog.FormatOsmBz2)
	assert.Contains(t, formats, catalog.FormatPoly)
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func Benchmark_OSMFitVutbr_FetchCatalog_Mock(b *testing.B) {
	ts := newMockServer(mockFitVutbrRootHTML, mockFitVutbrSubdirHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx := b.Context()

	for b.Loop() {
		_, _ = p.FetchCatalog(ctx)
	}
}
