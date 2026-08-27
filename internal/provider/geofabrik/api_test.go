package geofabrik_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockGeofabrikIndexJSON = `{
  "features": [
    {
      "properties": {
        "id": "europe",
        "name": "Europe"
      }
    },
    {
      "properties": {
        "id": "france",
        "name": "France",
        "parent": "europe",
        "urls": {
          "pbf": "https://download.geofabrik.de/europe/france-latest.osm.pbf",
          "bz2": "https://download.geofabrik.de/europe/france-latest.osm.bz2",
          "shp": "https://download.geofabrik.de/europe/france-latest-free.shp.zip"
        }
      }
    }
  ]
}`

func TestGeofabrik_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockGeofabrikIndexJSON))
	}))
	defer ts.Close()

	p := geofabrik.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	assert.Equal(t, "geofabrik", p.Name())
	assert.Equal(t, "geofabrik.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("europe"))
	assert.True(t, cat.Exist("france"))

	fr, exists := cat.Get("france")
	assert.True(t, exists)
	assert.Equal(t, "europe", fr.Parent)
	assert.True(t, fr.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, fr.ContainsFormat(catalog.FormatOsmBz2))
	assert.True(t, fr.ContainsFormat(catalog.FormatShpZip))
	assert.True(t, fr.ContainsFormat(catalog.FormatPoly))
}

func TestGeofabrik_FetchCatalog_Errors(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "server error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	p := geofabrik.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)

	p.IndexURL = "http://invalid-host-that-does-not-exist.test:9999/index.json"
	_, err = p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func Benchmark_Geofabrik_FetchCatalog_Mock(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockGeofabrikIndexJSON))
	}))
	defer ts.Close()

	p := geofabrik.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}
