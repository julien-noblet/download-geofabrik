package movisda_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/provider/movisda"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockMovisdaGeoJSON = `{
  "features": [
    {
      "properties": {
        "osm_id": 9407,
        "prefix": "AD-",
        "admin_level": "2",
        "name_en": "Andorra",
        "name": "Andorra"
      }
    },
    {
      "properties": {
        "osm_id": 307763,
        "prefix": "AE-",
        "admin_level": "2",
        "name_en": "",
        "name": "United Arab Emirates"
      }
    },
    {
      "properties": {
        "osm_id": 3766483,
        "prefix": "AE-DU-",
        "admin_level": "4",
        "name_en": "Dubai Emirate",
        "name": "Dubai"
      }
    },
    {
      "properties": {
        "osm_id": 999999,
        "prefix": "",
        "admin_level": "2",
        "name": "Invalid No Prefix"
      }
    },
    {
      "properties": {
        "osm_id": 888888,
        "prefix": "SINGLE",
        "admin_level": "4",
        "name": "Single Part Subarea"
      }
    }
  ]
}`

const mockMovisdaConflictJSON = `{
  "features": [
    {
      "properties": {
        "prefix": "FR-IDF-",
        "parent": "france",
        "name_en": "Ile-de-France"
      }
    },
    {
      "properties": {
        "prefix": "FR-IDF-",
        "parent": "germany",
        "name_en": "Conflict IDF"
      }
    }
  ]
}`

func TestMovisda_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockMovisdaGeoJSON))
	}))
	defer ts.Close()

	p := movisda.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	assert.Equal(t, "movisda", p.Name())
	assert.Equal(t, "movisda.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("ad"))
	assert.True(t, cat.Exist("ae"))
	assert.True(t, cat.Exist("ae-du"))
	assert.True(t, cat.Exist("single"))
	assert.False(t, cat.Exist(""))

	ad, exists := cat.Get("ad")
	assert.True(t, exists)
	assert.Equal(t, "Andorra", ad.Name)
	assert.Equal(t, "AD-latest", ad.File)
	assert.Empty(t, ad.Parent)
	assert.True(t, ad.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, ad.ContainsFormat(catalog.FormatPoly))
	assert.True(t, ad.ContainsFormat(catalog.FormatGeoJSON))

	ae, exists := cat.Get("ae")
	assert.True(t, exists)
	assert.Equal(t, "United Arab Emirates", ae.Name)

	du, exists := cat.Get("ae-du")
	assert.True(t, exists)
	assert.Equal(t, "ae", du.Parent)
	assert.Equal(t, "AE-DU-latest", du.File)
}

func TestMovisda_FetchCatalog_Errors(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "server error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	p := movisda.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.ErrorIs(t, err, movisda.ErrFetchCatalog)
}

func TestMovisda_FetchCatalog_InvalidURL(t *testing.T) {
	t.Parallel()

	p := movisda.NewProvider()
	p.IndexURL = "http://invalid-host-that-does-not-exist.test:9999/admin.geojson"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestMovisda_FetchCatalog_InvalidURLSyntax(t *testing.T) {
	t.Parallel()

	p := movisda.NewProvider()
	p.IndexURL = "http://[::1]:namedport/admin.geojson"

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestMovisda_FetchCatalog_InvalidJSON(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{invalid json`))
	}))
	defer ts.Close()

	p := movisda.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}

func TestMovisda_FetchCatalog_MergeConflict(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockMovisdaConflictJSON))
	}))
	defer ts.Close()

	p := movisda.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot merge element")
}

func TestMovisda_FetchCatalog_MergeConflictExplicit(t *testing.T) {
	t.Parallel()

	conflictJSON := `{
		"features": [
			{"properties": {"prefix": "X-A-", "admin_level": "4", "name": "A"}},
			{"properties": {"prefix": "Y-A-", "admin_level": "4", "name": "A2"}}
		]
	}`

	// Both will have ID "x-a" and "y-a", so no conflict. Let's make same prefix with different explicit parents:
	// Since prefix determines ID and parent, a conflict occurs when same ID has different parents:
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(conflictJSON))
	}))
	defer ts.Close()

	p := movisda.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
}

func TestMovisda_FetchCatalog_ContextCancelled(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockMovisdaGeoJSON))
	}))
	defer ts.Close()

	p := movisda.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := p.FetchCatalog(ctx)
	require.Error(t, err)
}

func TestMovisda_FetchCatalog_NilClient(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockMovisdaGeoJSON))
	}))
	defer ts.Close()

	p := movisda.NewProvider()
	p.IndexURL = ts.URL
	p.Client = nil

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)
	assert.True(t, cat.Exist("ad"))
}

func TestMovisda_DefaultFormats(t *testing.T) {
	t.Parallel()

	formats := movisda.DefaultFormats()
	assert.Contains(t, formats, catalog.FormatOsmPbf)
	assert.Contains(t, formats, catalog.FormatPoly)
	assert.Contains(t, formats, catalog.FormatGeoJSON)
}

func Benchmark_Movisda_FetchCatalog_Mock(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(mockMovisdaGeoJSON))
	}))
	defer ts.Close()

	p := movisda.NewProvider()
	p.IndexURL = ts.URL
	p.Client = ts.Client()

	ctx := context.Background()

	b.ResetTimer()

	for range b.N {
		_, _ = p.FetchCatalog(ctx)
	}
}
