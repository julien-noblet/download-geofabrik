package bbbike_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

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
      <tr><td><a href="Aachen/">Aachen</a></td></tr>
      <tr><td><a href="Berlin/">Berlin</a></td></tr>
      <tr><td><a href="Paris/">Paris</a></td></tr>
      <tr><td><a href="../">Parent</a></td></tr>
      <tr><td><a href="?C=M;O=A">Sort</a></td></tr>
    </tbody>
  </table>
</div>
</body>
</html>`

func TestBBBike_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(mockBBBikeHTML))
	}))
	defer ts.Close()

	p := bbbike.NewProvider()
	p.StartURL = ts.URL
	p.Client = ts.Client()

	assert.Equal(t, "bbbike", p.Name())
	assert.Equal(t, "bbbike.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("Aachen"))
	assert.True(t, cat.Exist("Berlin"))
	assert.True(t, cat.Exist("Paris"))
	assert.False(t, cat.Exist("Parent"))

	paris, exists := cat.Get("Paris")
	assert.True(t, exists)
	assert.True(t, paris.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, paris.ContainsFormat(catalog.FormatGarminOSM))
}

func TestBBBike_FetchCatalog_Errors(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "error", http.StatusBadGateway)
	}))
	defer ts.Close()

	p := bbbike.NewProvider()
	p.StartURL = ts.URL
	p.Client = ts.Client()

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func Benchmark_BBBike_FetchCatalog_Mock(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(mockBBBikeHTML))
	}))
	defer ts.Close()

	p := bbbike.NewProvider()
	p.StartURL = ts.URL
	p.Client = ts.Client()

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}
