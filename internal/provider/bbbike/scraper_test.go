package bbbike_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/iotest"

	"github.com/julien-noblet/download-geofabrik/internal/provider/bbbike"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockBBBikeHTML = `<!DOCTYPE html>
<html>
<head><title>BBBike extracts</title></head>
<body>
<div class="list">
  <table>
    <tbody>
      <tr><td><a name="no_href">no href link</a></td></tr>
      <tr><td><a href="Aachen/">Aachen</a></td></tr>
      <tr><td><a href="Berlin/">Berlin</a></td></tr>
      <tr><td><a href="Paris/">Paris</a></td></tr>
      <tr><td><a href="Tokyo/">Tokyo</a></td></tr>
      <tr><td><a href="NewYork/">NewYork</a></td></tr>
      <tr><td><a href="../">Parent</a></td></tr>
      <tr><td><a href="?C=M;O=A">Sort</a></td></tr>
      <tr><td><a href="http://example.com">External</a></td></tr>
      <tr><td><a href="some/deep/path">Deep</a></td></tr>
      <tr><td><a href="lowercase/">lowercase</a></td></tr>
    </tbody>
  </table>
</div>
</body>
</html>`

// newMockServer creates a test server that serves the given HTML body.
func newMockServer(body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(body))
	}))
}

// newProviderWithServer creates a Provider pointing at the given test server.
func newProviderWithServer(ts *httptest.Server) *bbbike.Provider {
	p := bbbike.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = ts.Client()

	return p
}

type errTransport struct{}

func (errTransport) RoundTrip(_ *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(iotest.ErrReader(errors.New("read error"))),
	}, nil
}

func TestBBBike_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockBBBikeHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	assert.Equal(t, "bbbike", p.Name())
	assert.Equal(t, "bbbike.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	// Valid cities must be present.
	validCities := []string{"Aachen", "Berlin", "Paris", "Tokyo", "NewYork"}
	for _, city := range validCities {
		assert.True(t, cat.Exist(city), "expected city %s to exist", city)
	}

	// Invalid links must be skipped.
	invalidEntries := []string{"..", "?C=M;O=A", "http://example.com", "some/deep/path", "lowercase"}
	for _, entry := range invalidEntries {
		assert.False(t, cat.Exist(entry), "expected entry %q to be skipped", entry)
	}

	// Every valid city must have all standard BBBike formats.
	expectedFormats := []string{
		catalog.FormatOsmPbf,
		catalog.FormatGarminOSM,
		catalog.FormatPoly,
		catalog.FormatShpZip,
		catalog.FormatOsmGz,
		catalog.FormatGeoJSON,
		catalog.FormatMapsforge,
		catalog.FormatMBTiles,
		catalog.FormatCSV,
		catalog.FormatGarminOnroad,
		catalog.FormatGarminOntrail,
		catalog.FormatGarminOpenTopo,
	}

	for _, city := range validCities {
		elem, exists := cat.Get(city)
		require.True(t, exists, "city %s must exist", city)

		for _, fmtID := range expectedFormats {
			assert.True(t, elem.ContainsFormat(fmtID), "city %s missing format %s", city, fmtID)
		}
	}
}

func TestBBBike_FetchCatalog_HTMLParseError(t *testing.T) {
	t.Parallel()

	p := bbbike.NewProvider()
	p.StartURL = "http://example.com"
	p.Client = &http.Client{Transport: errTransport{}}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot parse HTML")
}

func TestBBBike_FetchCatalog_HTTPError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "bad gateway", http.StatusBadGateway)
	}))
	defer ts.Close()

	p := newProviderWithServer(ts)

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, bbbike.ErrFetchCatalog)
}

func TestBBBike_FetchCatalog_InvalidURL(t *testing.T) {
	t.Parallel()

	p := bbbike.NewProvider()
	p.StartURL = "http://127.0.0.1:1" // unreachable port
	p.Client = &http.Client{}

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, bbbike.ErrFetchCatalog)
}

func TestBBBike_FetchCatalog_InvalidURLSyntax(t *testing.T) {
	t.Parallel()

	p := bbbike.NewProvider()
	p.StartURL = "http://[::1]:namedport/"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestBBBike_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockBBBikeHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestBBBike_FetchCatalog_EmptyPage(t *testing.T) {
	t.Parallel()

	emptyHTML := `<!DOCTYPE html><html><body></body></html>`

	ts := newMockServer(emptyHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.Empty(t, cat.Elements)
}

func TestBBBike_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockBBBikeHTML)
	defer ts.Close()

	p := bbbike.NewProvider()
	p.StartURL = ts.URL
	p.BaseURL = ts.URL
	p.Client = nil // nil client should fall back to http.DefaultClient

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.True(t, cat.Exist("Berlin"))
}

func TestBBBike_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := bbbike.DefaultFormats()
	require.NotNil(t, formats)

	expectedKeys := []string{
		catalog.FormatOsmPbf,
		catalog.FormatGarminOSM,
		catalog.FormatPoly,
		catalog.FormatShpZip,
		catalog.FormatOsmGz,
		catalog.FormatGeoJSON,
		catalog.FormatMapsforge,
		catalog.FormatMBTiles,
		catalog.FormatCSV,
		catalog.FormatGarminOnroad,
		catalog.FormatGarminOntrail,
		catalog.FormatGarminOpenTopo,
	}

	for _, key := range expectedKeys {
		fmtDef, exists := formats[key]
		assert.True(t, exists, "expected format key %s", key)
		assert.Equal(t, key, fmtDef.ID, "format ID must match key for %s", key)
		assert.NotEmpty(t, fmtDef.Loc, "format %s must have a Loc", key)
	}
}

func TestBBBike_FetchCatalog_CityFileField(t *testing.T) {
	t.Parallel()

	ts := newMockServer(mockBBBikeHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	// The File field should be "CityName/CityName" for every valid city.
	cities := []string{"Aachen", "Berlin", "Paris", "Tokyo", "NewYork"}
	for _, city := range cities {
		elem, exists := cat.Get(city)
		require.True(t, exists, "city %s must exist", city)
		assert.Equal(t, city+"/"+city, elem.File, "File field for %s", city)
	}
}

func Benchmark_BBBike_FetchCatalog_Mock(b *testing.B) {
	ts := newMockServer(mockBBBikeHTML)
	defer ts.Close()

	p := newProviderWithServer(ts)

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}
