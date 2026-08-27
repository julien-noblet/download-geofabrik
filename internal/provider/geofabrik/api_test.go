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
    },
    {
      "properties": {
        "id": "germany",
        "name": "Germany",
        "parent": "europe",
        "urls": {
          "pbf": "https://download.geofabrik.de/europe/germany-latest.osm.pbf",
          "bz2": "https://download.geofabrik.de/europe/germany-latest.osm.bz2",
          "history": "https://download.geofabrik.de/europe/germany.osh.pbf"
        }
      }
    },
    {
      "properties": {
        "id": "antarctica",
        "name": "Antarctica"
      }
    }
  ]
}`

const mockGeofabrikConflictJSON = `{
  "features": [
    {
      "properties": {
        "id": "france",
        "name": "France",
        "parent": "europe",
        "urls": {"pbf": "https://download.geofabrik.de/europe/france-latest.osm.pbf"}
      }
    },
    {
      "properties": {
        "id": "france",
        "name": "France Conflicting",
        "parent": "asia",
        "urls": {"pbf": "https://download.geofabrik.de/asia/france-latest.osm.pbf"}
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
	assert.True(t, cat.Exist("germany"))
	assert.True(t, cat.Exist("antarctica"))

	fr, exists := cat.Get("france")
	assert.True(t, exists)
	assert.Equal(t, "europe", fr.Parent)
	assert.True(t, fr.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, fr.ContainsFormat("osm.pbf.md5"))
	assert.True(t, fr.ContainsFormat(catalog.FormatOsmBz2))
	assert.True(t, fr.ContainsFormat("osm.bz2.md5"))
	assert.True(t, fr.ContainsFormat(catalog.FormatShpZip))
	assert.True(t, fr.ContainsFormat(catalog.FormatPoly))
	assert.True(t, fr.ContainsFormat(catalog.FormatKml))
	assert.True(t, fr.ContainsFormat(catalog.FormatState))

	de, exists := cat.Get("germany")
	assert.True(t, exists)
	assert.True(t, de.ContainsFormat(catalog.FormatOshPbf))
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
	assert.ErrorIs(t, err, geofabrik.ErrFetchCatalog)
}

func TestGeofabrik_FetchCatalog_InvalidURL(t *testing.T) {
	t.Parallel()

	p := geofabrik.NewProvider()
	p.IndexURL = "http://invalid-host-that-does-not-exist.test:9999/index.json"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestGeofabrik_FetchCatalog_InvalidURLSyntax(t *testing.T) {
	t.Parallel()

	p := geofabrik.NewProvider()
	p.IndexURL = "http://[::1]:namedport/index.json"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestGeofabrik_FetchCatalog_InvalidJSON(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{not valid json`))
	}))
	defer ts.Close()

	p := geofabrik.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestGeofabrik_FetchCatalog_MergeConflict(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockGeofabrikConflictJSON))
	}))
	defer ts.Close()

	p := geofabrik.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot merge element")
}

func TestGeofabrik_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockGeofabrikIndexJSON))
	}))
	defer ts.Close()

	p := geofabrik.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestGeofabrik_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	p := geofabrik.NewProvider()
	p.IndexURL = "http://invalid-host-that-does-not-exist.test:9999/index.json"
	p.Client = nil

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestGeofabrik_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := geofabrik.DefaultFormats()
	assert.Contains(t, formats, catalog.FormatOsmPbf)
	assert.Contains(t, formats, catalog.FormatOsmBz2)
	assert.Contains(t, formats, catalog.FormatShpZip)
	assert.Contains(t, formats, catalog.FormatPoly)
	assert.Contains(t, formats, catalog.FormatKml)
	assert.Contains(t, formats, catalog.FormatState)
	assert.Contains(t, formats, "osm.pbf.md5")
	assert.Contains(t, formats, "osm.bz2.md5")
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
