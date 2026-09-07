package geofabrik_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/generator/importer/geofabrik"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockGeofabrikIndexJSON = `{
  "features": [
    {"properties": {"id": "europe", "name": "Europe"}},
    {
      "properties": {
        "id": "france",
        "name": "France",
        "parent": "europe",
        "urls": {
          "pbf": "https://download.geofabrik.de/europe/france-latest.osm.pbf",
          "bz2": "https://download.geofabrik.de/europe/france-latest.osm.bz2",
          "shp": "https://download.geofabrik.de/europe/france-latest-free.shp.zip",
          "history": "https://download.geofabrik.de/europe/france.osh.pbf"
        }
      }
    },
    {
      "properties": {
        "id": "germany",
        "name": "Germany",
        "parent": "europe",
        "urls": {
          "pbf": "https://download.geofabrik.de/europe/germany-latest.osm.pbf"
        }
      }
    },
    {"properties": {"id": "italy", "name": "Italy", "parent": "europe"}},
    {"properties": {"id": "spain", "name": "Spain", "parent": "europe"}},
    {"properties": {"id": "united_kingdom", "name": "United Kingdom", "parent": "europe"}},
    {"properties": {"id": "asia", "name": "Asia"}},
    {"properties": {"id": "japan", "name": "Japan", "parent": "asia"}},
    {"properties": {"id": "africa", "name": "Africa"}},
    {"properties": {"id": "antarctica", "name": "Antarctica"}}
  ]
}`

func TestGetIndex(t *testing.T) {
	t.Parallel()
	viper.Set("log", true)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/404":
			http.NotFound(w, r)
		case "/500":
			http.Error(w, "internal server error", http.StatusInternalServerError)
		case "/invalid-json":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{not-json`))

		default:
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockGeofabrikIndexJSON))
		}
	}))
	t.Cleanup(ts.Close)

	tests := []struct {
		name    string
		myURL   string
		wantErr bool
	}{
		{
			name:    "Test Success",
			myURL:   ts.URL,
			wantErr: false,
		},
		{
			name:    "Test 404",
			myURL:   ts.URL + "/404",
			wantErr: true,
		},
		{
			name:    "Test 500",
			myURL:   ts.URL + "/500",
			wantErr: true,
		},
		{
			name:    "Test Invalid JSON",
			myURL:   ts.URL + "/invalid-json",
			wantErr: true,
		},
		{
			name:    "Test Invalid URL Syntax",
			myURL:   "http://[::1]:namedport/index.json",
			wantErr: true,
		},
		{
			name:    "Test Connection Error",
			myURL:   "http://127.0.0.1:0",
			wantErr: true,
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			index, err := geofabrik.GetIndex(thisTest.myURL)
			if thisTest.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, index)

				if len(index.Features) < 10 {
					t.Errorf("GetIndex() error I should have more features!!!")
				}

				converted, err := geofabrik.Convert(index)
				require.NoError(t, err)
				require.NotNil(t, converted)

				e, err := converted.GetElement("france")
				require.NoError(t, err)
				require.NotNil(t, e)
			}
		})
	}
}

func TestFormatDefinition(t *testing.T) {
	t.Parallel()

	defs := geofabrik.FormatDefinition()
	assert.NotEmpty(t, defs)
	assert.Contains(t, defs, "osm.pbf")
	assert.Contains(t, defs, "osm.pbf.md5")
	assert.Contains(t, defs, "osm.bz2")
	assert.Contains(t, defs, "osm.bz2.md5")
}

func TestConvert_MockData(t *testing.T) {
	t.Parallel()

	mockIndex := &geofabrik.Index{
		Features: []geofabrik.IndexElement{
			{
				ElementProperties: geofabrik.IndexElementProperties{
					ID:     "europe",
					Name:   "Europe",
					Parent: "",
					Urls: map[string]string{
						"pbf":     "https://example.com/europe.pbf",
						"bz2":     "https://example.com/europe.bz2",
						"shp":     "https://example.com/europe.shp",
						"history": "https://example.com/europe.history",
					},
				},
			},
			{
				ElementProperties: geofabrik.IndexElementProperties{
					ID:     "france",
					Name:   "France",
					Parent: "europe",
					Urls: map[string]string{
						"pbf": "https://example.com/france.pbf",
					},
				},
			},
		},
	}

	cfg, err := geofabrik.Convert(mockIndex)
	require.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.True(t, cfg.Exist("europe"))
	assert.True(t, cfg.Exist("france"))

	fr, err := cfg.GetElement("france")
	require.NoError(t, err)
	assert.Equal(t, "europe", fr.Parent)
	assert.Contains(t, fr.Formats, "osm.pbf")
	assert.Contains(t, fr.Formats, "osm.pbf.md5")
}

func TestConvert_ConflictError(t *testing.T) {
	t.Parallel()

	conflictingIndex := &geofabrik.Index{
		Features: []geofabrik.IndexElement{
			{
				ElementProperties: geofabrik.IndexElementProperties{
					ID:     "france",
					Name:   "France",
					Parent: "europe",
					Urls:   map[string]string{"pbf": "https://example.com/france.pbf"},
				},
			},
			{
				ElementProperties: geofabrik.IndexElementProperties{
					ID:     "france",
					Name:   "France Conflicting",
					Parent: "asia",
					Urls:   map[string]string{"pbf": "https://example.com/france.pbf"},
				},
			},
		},
	}

	cfg, err := geofabrik.Convert(conflictingIndex)
	require.Error(t, err)
	assert.Nil(t, cfg)
}

func Benchmark_FormatDefinition(b *testing.B) {
	for range b.N {
		_ = geofabrik.FormatDefinition()
	}
}

func Benchmark_Convert(b *testing.B) {
	mockIndex := &geofabrik.Index{
		Features: []geofabrik.IndexElement{
			{
				ElementProperties: geofabrik.IndexElementProperties{
					ID:   "france",
					Name: "France",
					Urls: map[string]string{"pbf": "url", "bz2": "url", "shp": "url"},
				},
			},
		},
	}

	for range b.N {
		_, _ = geofabrik.Convert(mockIndex)
	}
}
