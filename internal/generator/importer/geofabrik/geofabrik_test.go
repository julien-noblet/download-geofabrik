package geofabrik_test

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/generator/importer/geofabrik"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetIndex(t *testing.T) {
	t.Parallel()
	viper.Set("log", true)

	tests := []struct {
		name    string
		myURL   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test",
			myURL:   geofabrik.GeofabrikIndexURL,
			wantErr: false,
		},
		{
			name:    "Test 404",
			myURL:   "https://google.com/404",
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
				assert.NotNil(t, index)

				if len(index.Features) < 10 {
					t.Errorf("GetIndex() error I should have more features!!!")
				}

				converted, err := geofabrik.Convert(index)
				if converted == nil || err != nil {
					t.Errorf("GetIndex() error cant convert !!!\n%v", err)
				}

				if e, err := converted.GetElement("france"); err != nil || e == nil {
					t.Errorf("GetIndex() error cant find element !!!\nconfig=%v\nerr=%v", converted, err)
				}
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
