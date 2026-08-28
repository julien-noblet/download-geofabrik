package generator //nolint:testpackage // testing internal functions

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/internal/provider"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
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

type mockGeneratorProvider struct {
	catalogData *catalog.Catalog
	name        string
	shouldErr   bool
}

func (m *mockGeneratorProvider) Name() string              { return m.name }
func (m *mockGeneratorProvider) Description() string       { return "Mock provider for generator tests" }
func (m *mockGeneratorProvider) DefaultConfigFile() string { return m.name + ".yml" }
func (m *mockGeneratorProvider) FetchCatalog(_ context.Context) (*catalog.Catalog, error) {
	if m.shouldErr {
		return nil, errors.New("mock fetch failed")
	}

	return m.catalogData, nil
}

func TestPerformGenerate_Success(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.Elements["test-elem"] = catalog.Element{
		ID:      "test-elem",
		Name:    "Test Element",
		Formats: []string{formats.FormatPoly, formats.FormatOsmPbf},
	}

	prov := &mockGeneratorProvider{
		name:        "mock-gen-success",
		catalogData: cat,
	}
	provider.Register(prov)

	tmpDir := t.TempDir()
	outFile := filepath.Join(tmpDir, "out.yml")

	err := PerformGenerate("mock-gen-success", false, outFile)
	require.NoError(t, err)

	loaded, err := catalog.LoadFile(outFile)
	require.NoError(t, err)
	assert.Contains(t, loaded.Elements, "test-elem")
	// Verify formats were sorted
	assert.Equal(t, catalog.Formats{formats.FormatOsmPbf, formats.FormatPoly}, loaded.Elements["test-elem"].Formats)
}

func TestPerformGenerate_FetchError(t *testing.T) {
	t.Parallel()

	prov := &mockGeneratorProvider{
		name:      "mock-gen-err",
		shouldErr: true,
	}
	provider.Register(prov)

	tmpDir := t.TempDir()
	outFile := filepath.Join(tmpDir, "out.yml")

	err := PerformGenerate("mock-gen-err", false, outFile)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mock fetch failed")
}

func TestPerformGenerate_SaveError(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	prov := &mockGeneratorProvider{
		name:        "mock-gen-save-err",
		catalogData: cat,
	}
	provider.Register(prov)

	// An impossible directory path
	outFile := "/dev/null/impossible/path.yml"

	err := PerformGenerate("mock-gen-save-err", false, outFile)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to write config")
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
