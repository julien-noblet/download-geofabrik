package generator //nolint:testpackage // testing internal functions

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v3"
)

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

func sampleElementValidPtr() map[string]element.Element {
	return map[string]element.Element{
		"africa": *sampleAfricaElementPtr(),
	}
}

func sampleFormatValidPtr() map[string]formats.Format {
	return map[string]formats.Format{
		formats.FormatOsmPbf: {
			ID:  formats.FormatOsmPbf,
			Loc: ".osm.pbf",
		},
		formats.FormatState: {
			ID:       formats.FormatState,
			Loc:      "-updates/state.txt",
			BasePath: "../state/",
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

	myConfig := SampleConfigValidPtr()
	want, err := yaml.Marshal(myConfig)
	require.NoError(t, err)

	got, err := myConfig.Generate()
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestGenerate_Unknown(t *testing.T) {
	t.Parallel()

	err := Generate("unknown_service", false, "/tmp/dummy.yml")
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrUnknownService)
}

func Test_write(t *testing.T) {
	t.Parallel()

	input := "../../geofabrik.yml"
	tmpDir := t.TempDir()
	output := filepath.Join(tmpDir, "test.yml")

	c, err := config.LoadConfig(input)
	require.NoError(t, err)

	err = Write(c, output)
	require.NoError(t, err)

	_, err = os.ReadFile(output)
	require.NoError(t, err)
}

func TestCleanup(t *testing.T) {
	t.Parallel()

	cfg := &config.Config{
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
	}

	Cleanup(cfg)
	af := cfg.Elements["africa"]
	expected := element.Formats{
		formats.FormatGeoJSON,
		formats.FormatOsmPbf,
		formats.FormatPoly,
		formats.FormatState,
	}

	if !reflect.DeepEqual(af.Formats, expected) {
		t.Errorf("Cleanup() = %v, want %v", af.Formats, expected)
	}
}

func Benchmark_Cleanup(b *testing.B) {
	for range b.N {
		cfg := SampleConfigValidPtr()
		Cleanup(&cfg)
	}
}

func Benchmark_Slice_Generate(b *testing.B) {
	cfg := SampleConfigValidPtr()

	for range b.N {
		_, _ = cfg.Generate()
	}
}

func Benchmark_Write(b *testing.B) {
	cfg := SampleConfigValidPtr()
	tmpDir := b.TempDir()
	outFile := filepath.Join(tmpDir, "out.yml")

	for range b.N {
		_ = Write(&cfg, outFile)
	}
}
