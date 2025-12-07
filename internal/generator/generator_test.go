package generator //nolint:testpackage // testing internal functions

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"testing"

	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/internal/generator/importer/geofabrik"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v3"
)

// MockScrapper implements scrapper.IScrapper.
type MockScrapper struct {
	BaseURL  string
	StartURL string
	PB       int
}

func (m *MockScrapper) GetConfig() *config.Config {
	return &config.Config{
		Elements:      element.MapElement{},
		ElementsMutex: &sync.RWMutex{},
	}
}

func (m *MockScrapper) Collector() *colly.Collector {
	// Return a collector that visits mock URL?
	// We need it to NOT hit internet.
	return colly.NewCollector()
}

func (m *MockScrapper) Limit() *colly.LimitRule {
	return &colly.LimitRule{}
}

func (m *MockScrapper) GetPB() int {
	return m.PB
}

func (m *MockScrapper) GetStartURL() string {
	return m.StartURL
}
func (m *MockScrapper) ParseFormat(_, _ string) {}

func sampleAfricaElementPtr() *element.Element {
	return &element.Element{
		ID:   "africa",
		Name: "Africa",
		Formats: []string{
			formats.FormatOsmPbf,
			"osm.pbf.md5",
			formats.FormatOsmBz2,
			"osm.bz2.md5",
			formats.FormatOshPbf,
			"osh.pbf.md5",
			formats.FormatPoly,
			formats.FormatKml,
			formats.FormatState,
		},
	}
}

func sampleGeorgiaUsElementPtr() *element.Element {
	return &element.Element{
		ID:   "georgia-us",
		File: "georgia",
		Name: "Georgia (US State)",
		Formats: []string{
			formats.FormatOsmPbf,
			"osm.pbf.md5",
			formats.FormatShpZip,
			formats.FormatOsmBz2,
			"osm.bz2.md5",
			formats.FormatOshPbf,
			"osh.pbf.md5",
			formats.FormatPoly,
			formats.FormatKml,
			formats.FormatState,
		},
		Parent: "us",
	}
}

func sampleUsElementPtr() *element.Element {
	return &element.Element{
		ID:     "us",
		Meta:   true,
		Name:   "United States of America",
		Parent: "north-america",
	}
}

func sampleNorthAmericaElementPtr() *element.Element {
	return &element.Element{
		ID:   "north-america",
		Name: "North America",
		Formats: []string{
			formats.FormatOsmPbf,
			"osm.pbf.md5",
			formats.FormatOsmBz2,
			"osm.bz2.md5",
			formats.FormatOshPbf,
			"osh.pbf.md5",
			formats.FormatPoly,
			formats.FormatKml,
			formats.FormatState,
		},
	}
}

func sampleElementValidPtr() map[string]element.Element {
	return map[string]element.Element{
		"africa":        *sampleAfricaElementPtr(),
		"georgia-us":    *sampleGeorgiaUsElementPtr(),
		"us":            *sampleUsElementPtr(),
		"north-america": *sampleNorthAmericaElementPtr(),
	}
}

func sampleFormatValidPtr() map[string]formats.Format {
	return map[string]formats.Format{
		// Blank
		"": {
			ID:       "",
			Loc:      "",
			BasePath: "",
		}, formats.FormatOsmPbf: {
			ID:  formats.FormatOsmPbf,
			Loc: ".osm.pbf",
			// BasePath: "/",
		}, formats.FormatState: {
			ID:       formats.FormatState,
			Loc:      "-updates/state.txt",
			BasePath: "../state/",
		}, formats.FormatPoly: {
			ID:      formats.FormatPoly,
			Loc:     ".poly",
			BaseURL: "http://my.new.url/folder",
		}, formats.FormatOsmBz2: {
			ID:       formats.FormatOsmBz2,
			Loc:      ".osm.bz2",
			BasePath: "../osmbz2/",
			BaseURL:  "http://my.new.url/folder",
		}, formats.FormatOsmGz: {
			ID:       formats.FormatOsmGz,
			Loc:      ".osm.gz",
			BasePath: "../osmgz/",
			BaseURL:  "http://my.new.url/folder",
		},
	}
}

func SampleConfigValidPtr() config.Config {
	return config.Config{
		BaseURL:  "https://my.base.url",
		Formats:  sampleFormatValidPtr(),
		Elements: sampleElementValidPtr(),
	}
}

func TestSlice_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		e       element.MapElement
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Marshaling OK, no error",
			e:       sampleElementValidPtr(),
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, thisTest := range tests {
		myConfig := SampleConfigValidPtr()
		myConfig.Elements = map[string]element.Element{} // void Elements
		myConfig.Elements = thisTest.e
		thisTest.want, _ = yaml.Marshal(myConfig)
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := myConfig.Generate()
			if err != nil != thisTest.wantErr {
				t.Errorf("Slice.Generate() error = %v, wantErr %v", err, thisTest.wantErr)

				return
			}

			if !reflect.DeepEqual(got, thisTest.want) {
				t.Errorf("Slice.Generate() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	type args struct {
		service    string
		configfile string
		progress   bool
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "run",
			args: args{
				service:    ServiceGeofabrik,
				progress:   false,
				configfile: "/tmp/gen_test.yml",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := Generate(tt.args.service, tt.args.progress, tt.args.configfile); err != nil {
				t.Errorf("Generate() error = %v", err)
			}
		})
	}
}

func Test_write(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		output string
	}{
		// TODO: Add test cases.
		{name: "geofabrik", input: "../../geofabrik.yml", output: "/tmp/test.yml"},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			c, _ := config.LoadConfig(thisTest.input)

			err := Write(c, thisTest.output)
			if err != nil {
				t.Errorf("write() error = %v", err)
			}

			input, err := os.ReadFile(thisTest.input)
			if err != nil {
				t.Errorf("read() error = %v", err)
			}

			output, err := os.ReadFile(thisTest.output)
			if err != nil {
				t.Errorf("read() error = %v", err)
			}

			reflect.DeepEqual(input, output)
		})
	}
}

func TestCleanup(t *testing.T) {
	t.Parallel()

	type args struct {
		c *config.Config
	}

	tests := []struct {
		name string
		args args
		want element.Formats
	}{
		// TODO: Add test cases.
		{
			name: "example 1",
			args: args{
				c: &config.Config{
					BaseURL: "https://my.base.url",
					Formats: formats.FormatDefinitions{
						formats.FormatOsmPbf: {
							ID:  formats.FormatOsmPbf,
							Loc: ".osm.pbf",
							// BasePath: "/",
						}, formats.FormatState: {
							ID:       formats.FormatState,
							Loc:      "-updates/state.txt",
							BasePath: "../state/",
						}, formats.FormatPoly: {
							ID:      formats.FormatPoly,
							Loc:     ".poly",
							BaseURL: "http://my.new.url/folder",
						}, formats.FormatOsmBz2: {
							ID:       formats.FormatOsmBz2,
							Loc:      ".osm.bz2",
							BasePath: "../osmbz2/",
							BaseURL:  "http://my.new.url/folder",
						}, formats.FormatOsmGz: {
							ID:       formats.FormatOsmGz,
							Loc:      ".osm.gz",
							BasePath: "../osmgz/",
							BaseURL:  "http://my.new.url/folder",
						},
					},
					Elements: element.MapElement{
						"africa": {
							ID:   "africa",
							Name: "Africa",
							Formats: []string{
								formats.FormatOsmPbf,
								"osm.pbf.md5",
								formats.FormatOsmBz2,
								"osm.bz2.md5",
								formats.FormatOshPbf,
								"osh.pbf.md5",
								formats.FormatPoly,
								formats.FormatKml,
								formats.FormatState,
							},
						},
					},
				},
			},
			want: element.Formats{
				formats.FormatKml,
				formats.FormatOshPbf,
				"osh.pbf.md5",
				formats.FormatOsmBz2,
				"osm.bz2.md5",
				formats.FormatOsmPbf,
				"osm.pbf.md5",
				formats.FormatPoly,
				formats.FormatState,
			},
		},
		{
			name: "example 2",
			args: args{
				c: &config.Config{
					BaseURL: "https://my.base.url",
					Formats: formats.FormatDefinitions{},
					Elements: element.MapElement{
						"africa": {
							ID:   "africa",
							Name: "Africa",
							Formats: []string{
								formats.FormatOsmPbf,
								formats.FormatGeoJSON,
								formats.FormatPoly,
								formats.FormatState,
							},
						},
					},
				},
			},
			want: element.Formats{
				formats.FormatGeoJSON,
				formats.FormatOsmPbf,
				formats.FormatPoly,
				formats.FormatState,
			},
		},
	}

	for _, tt := range tests {
		myTest := tt

		t.Run(myTest.name, func(t *testing.T) {
			t.Parallel()

			af := myTest.args.c.Elements["africa"]
			Cleanup(myTest.args.c)
			// compare af.Formats != tt.want
			if !reflect.DeepEqual(af.Formats, myTest.want) {
				t.Errorf("Cleanup() = %v, want %v", af.Formats, myTest.want)
			}
		})
	}
}

func TestPerformGenerate_WithMock(t *testing.T) {
	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, `{
    "features": [
        {
            "properties": {
                "id": "test-region",
                "name": "Test Region",
                "urls": {
                    "pbf": "https://example.com/test-region.osm.pbf"
                }
            }
        }
    ]
}`)
	}))
	defer ts.Close()

	// Override URL
	oldURL := geofabrik.GeofabrikIndexURL
	geofabrik.GeofabrikIndexURL = ts.URL

	defer func() { geofabrik.GeofabrikIndexURL = oldURL }()

	// Temp config file
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yml")

	// Run Generate
	err := PerformGenerate(ServiceGeofabrik, false, configFile)
	require.NoError(t, err)

	// Verify file exists
	_, err = os.Stat(configFile)
	require.NoError(t, err)

	// Verify content (basic check)
	content, err := os.ReadFile(configFile)
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-region")
	assert.Contains(t, string(content), "osm.pbf")
}

func TestHandleProgress(_ *testing.T) {
	// Mock scrapper with mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "OK")
	}))
	defer ts.Close()

	ms := &MockScrapper{
		StartURL: ts.URL,
		PB:       10,
	}

	// Capture stdout? No need, just ensure no panic and coverage.
	handleProgress(ms)
}

func TestVisitAndWait(_ *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "OK")
	}))
	defer ts.Close()

	c := colly.NewCollector()
	visitAndWait(c, ts.URL)
}

func TestPerformGenerate_Unknown(t *testing.T) {
	err := PerformGenerate("unknown", false, "")
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrUnknownService)
}

func TestPerformGenerate_Geofabrik_Error(t *testing.T) {
	// Mock server returns 500
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	oldURL := geofabrik.GeofabrikIndexURL
	geofabrik.GeofabrikIndexURL = ts.URL

	defer func() { geofabrik.GeofabrikIndexURL = oldURL }()

	err := PerformGenerate(ServiceGeofabrik, false, "dummy")
	require.Error(t, err)
	// Check error message contains "failed to get index"
	assert.Contains(t, err.Error(), "failed to get index")
}

func TestPerformGenerate_Geofabrik_InvalidJSON(t *testing.T) {
	// Mock server returns invalid JSON
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, `invalid json`)
	}))
	defer ts.Close()

	oldURL := geofabrik.GeofabrikIndexURL
	geofabrik.GeofabrikIndexURL = ts.URL

	defer func() { geofabrik.GeofabrikIndexURL = oldURL }()

	err := PerformGenerate(ServiceGeofabrik, false, "dummy")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get index") // GetIndex fails at unmarshal -> "error while unmarshalling" -> wrapped?
	// In geofabrik.go: return nil, fmt.Errorf(ErrUnmarshallingBody, err)
	// In generator.go: return fmt.Errorf("failed to get index: %w", err) (wait, GetIndex returns error)
	// generator.go: err := geofabrik.GetIndex... if err != nil ...
}

func TestPerformGenerate_Geofabrik_WriteError(t *testing.T) {
	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, `{
    "features": []
}`)
	}))
	defer ts.Close()

	oldURL := geofabrik.GeofabrikIndexURL
	geofabrik.GeofabrikIndexURL = ts.URL

	defer func() { geofabrik.GeofabrikIndexURL = oldURL }()

	// Invalid config file path (directory)
	tmpDir := t.TempDir()

	err := PerformGenerate(ServiceGeofabrik, false, tmpDir)
	require.Error(t, err)
	// Write fails because tmpDir is a directory or use /dev/null/fail?
	// os.WriteFile on directory fails with "is a directory" on Linux.
	// Or use a non-existent directory structure /non/existent/path.
}

func TestVisitAndWait_Error(_ *testing.T) {
	c := colly.NewCollector()
	// Invalid URL causes Visit to return error immediately
	// e.g. schemes must be present
	visitAndWait(c, ":/invalid-url")
}
