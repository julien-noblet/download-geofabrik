package config_test

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestElem2preURL_WithParent(t *testing.T) {
	cfg := &config.Config{
		BaseURL: "https://example.com",
		Elements: element.MapElement{
			"parent": {
				ID:   "parent",
				Name: "Parent Element",
			},
			"child": {
				ID:     "child",
				Name:   "Child Element",
				Parent: "parent",
			},
		},
		Formats:       formats.FormatDefinitions{},
		ElementsMutex: &sync.RWMutex{},
	}

	childElem, err := cfg.GetElement("child")
	require.NoError(t, err)

	url, err := config.Elem2preURL(cfg, childElem)
	require.NoError(t, err)

	expectedURL := "https://example.com/parent/child"
	assert.Equal(t, expectedURL, url)
}

func TestElem2preURL_NoParent_WithBaseURLArgs(t *testing.T) {
	cfg := &config.Config{
		BaseURL: "https://example.com",
		Elements: element.MapElement{
			"item": {ID: "item"},
		},
		ElementsMutex: &sync.RWMutex{},
	}
	elem, err := cfg.GetElement("item")
	require.NoError(t, err)

	url, err := config.Elem2preURL(cfg, elem, "custom")
	require.NoError(t, err)
	assert.Equal(t, "https://example.com/custom/item", url)

	url, err = config.Elem2preURL(cfg, elem, "http://other.com", "path")
	require.NoError(t, err)
	assert.Equal(t, "http://other.com/path/item", url)

	url, err = config.Elem2preURL(cfg, elem)
	require.NoError(t, err)
	assert.Equal(t, "https://example.com/item", url)
}

func TestGenerate(t *testing.T) {
	cfg := &config.Config{
		BaseURL: "https://example.com",
		Elements: element.MapElement{
			"one": {ID: "one"},
		},
	}
	data, err := cfg.Generate()
	require.NoError(t, err)
	assert.Contains(t, string(data), "https://example.com")
	assert.Contains(t, string(data), "one")
}

func TestMergeElement(t *testing.T) {
	cfg := &config.Config{
		Elements:      make(element.MapElement),
		ElementsMutex: &sync.RWMutex{},
	}

	el1 := &element.Element{ID: "e1", Parent: "p1", Formats: []string{"osm"}}
	err := cfg.MergeElement(el1)
	require.NoError(t, err)
	assert.True(t, cfg.Exist("e1"))

	el2 := &element.Element{ID: "e1", Parent: "p1", Formats: []string{"pbf"}}
	err = cfg.MergeElement(el2)
	require.NoError(t, err)

	e, err := cfg.GetElement("e1")
	require.NoError(t, err)
	assert.Contains(t, e.Formats, "osm")
	assert.Contains(t, e.Formats, "pbf")

	elBad := &element.Element{ID: "e1", Parent: "p2"}
	err = cfg.MergeElement(elBad)
	assert.ErrorIs(t, err, config.ErrParentMismatch)
}

func TestAddExtension(t *testing.T) {
	cfg := &config.Config{
		Elements: element.MapElement{
			"e1": {ID: "e1", Formats: []string{"osm"}},
		},
		ElementsMutex: &sync.RWMutex{},
	}

	cfg.AddExtension("e1", "pbf")
	e, _ := cfg.GetElement("e1")
	assert.Contains(t, e.Formats, "pbf")
}

func TestElem2URL(t *testing.T) {
	cfg := &config.Config{
		BaseURL: "https://example.com",
		Formats: formats.FormatDefinitions{
			"osm.pbf": {Loc: "-latest.osm.pbf"},
		},
		Elements: element.MapElement{
			"e1": {ID: "e1", Formats: []string{"osm.pbf"}},
		},
		ElementsMutex: &sync.RWMutex{},
	}

	e, _ := cfg.GetElement("e1")
	url, err := config.Elem2URL(cfg, e, "osm.pbf")
	require.NoError(t, err)
	assert.Equal(t, "https://example.com/e1-latest.osm.pbf", url)

	_, err = config.Elem2URL(cfg, e, "missing")
	assert.ErrorIs(t, err, config.ErrFormatNotExist)
}

func TestLoadConfig(t *testing.T) {
	content := `
baseURL: https://test.com
elements:
  e1:
    id: e1
`
	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "config.yml")
	err := os.WriteFile(f, []byte(content), 0o600)
	require.NoError(t, err)

	cfg, err := config.LoadConfig(f)
	require.NoError(t, err)
	assert.Equal(t, "https://test.com", cfg.BaseURL)
	assert.True(t, cfg.Exist("e1"))
}

func TestFindElem_SeparatorNormalization(t *testing.T) {
	content := `
baseURL: https://test.com
elements:
  rhone_alpes:
    id: rhone_alpes
    name: Rhone-Alpes
  ile-de-france:
    id: ile-de-france
    name: Ile-de-France
`
	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "config.yml")
	err := os.WriteFile(f, []byte(content), 0o600)
	require.NoError(t, err)

	cfg, err := config.LoadConfig(f)
	require.NoError(t, err)

	// Direct lookup
	e1, err := config.FindElem(cfg, "rhone_alpes")
	require.NoError(t, err)
	assert.Equal(t, "rhone_alpes", e1.ID)

	// Hyphen fallback to underscore
	e2, err := config.FindElem(cfg, "rhone-alpes")
	require.NoError(t, err)
	assert.Equal(t, "rhone_alpes", e2.ID)
	assert.True(t, cfg.Exist("rhone-alpes"))

	// Underscore fallback to hyphen
	e3, err := config.FindElem(cfg, "ile_de_france")
	require.NoError(t, err)
	assert.Equal(t, "ile-de-france", e3.ID)
	assert.True(t, cfg.Exist("ile_de_france"))

	// Not found
	_, err = config.FindElem(cfg, "nonexistent")
	require.ErrorIs(t, err, config.ErrFindElem)
	assert.False(t, cfg.Exist("nonexistent"))

	// Nil config
	_, err = config.FindElem(nil, "rhone_alpes")
	require.ErrorIs(t, err, config.ErrFindElem)
}

func TestFindElem_DateResolution(t *testing.T) {
	content := `
baseURL: https://test.com
formats:
  osm.pbf:
    ext: osm.pbf
    loc: .osm.pbf
    basepath: czech_republic/
elements:
  czech_republic:
    id: czech_republic
    name: Czech Republic
    file: czech-republic
    files:
      - poly
  latest:
    id: latest
    file: czech_republic-2026-09-06
    name: Czech Republic (latest)
    files:
      - osm.pbf
`
	tmpDir := t.TempDir()
	f := filepath.Join(tmpDir, "config.yml")
	err := os.WriteFile(f, []byte(content), 0o600)
	require.NoError(t, err)

	cfg, err := config.LoadConfig(f)
	require.NoError(t, err)

	// Valid date resolution
	e, err := config.FindElem(cfg, "2026-09-06")
	require.NoError(t, err)
	assert.Equal(t, "2026-09-06", e.ID)
	assert.Equal(t, "czech_republic-2026-09-06", e.File)
	assert.Equal(t, "Czech Republic 2026-09-06", e.Name)
	assert.True(t, cfg.Exist("2026-09-06"))

	// Historical date resolution
	eHist, err := config.FindElem(cfg, "2006-04-03")
	require.NoError(t, err)
	assert.Equal(t, "2006-04-03", eHist.ID)
	assert.Equal(t, "czech_republic-2006-04-03", eHist.File)

	// URL generation for resolved date
	url, err := config.Elem2URL(cfg, e, "osm.pbf")
	require.NoError(t, err)
	assert.Equal(t, "https://test.com/czech_republic/czech_republic-2026-09-06.osm.pbf", url)

	// Invalid date
	_, err = config.FindElem(cfg, "2026-99-99")
	require.ErrorIs(t, err, config.ErrFindElem)
	assert.False(t, cfg.Exist("2026-99-99"))
}

func TestIsHashable(t *testing.T) {
	cfg := &config.Config{
		Formats: formats.FormatDefinitions{
			"osm.pbf":     {},
			"osm.pbf.md5": {},
		},
	}
	isH, hash, ext := config.IsHashable(cfg, "osm.pbf")
	assert.True(t, isH)
	assert.Equal(t, "osm.pbf.md5", hash)
	assert.Equal(t, "md5", ext)

	isH, _, _ = config.IsHashable(cfg, "other")
	assert.False(t, isH)
}

const geofabrikYml = "../../geofabrik.yml"

func Benchmark_Exist_geofabrik_yml(b *testing.B) {
	c, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		c.Exist("france")
	}
}

func Benchmark_loadConfig_geofabrik_yml(b *testing.B) {
	for range b.N {
		_, _ = config.LoadConfig(geofabrikYml)
	}
}

func Benchmark_GetElement_geofabrik_yml(b *testing.B) {
	c, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		_, _ = c.GetElement("france")
	}
}

func Benchmark_IsHashable_geofabrik_yml(b *testing.B) {
	c, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		for f := range c.Formats {
			config.IsHashable(c, f)
		}
	}
}

func Benchmark_findElem_parse_all_geofabrik_yml(b *testing.B) {
	c, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		for k := range c.Elements {
			_, _ = config.FindElem(c, k)
		}
	}
}

func Benchmark_GetElement_parse_all_geofabrik_yml(b *testing.B) {
	c, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		for k := range c.Elements {
			_, _ = c.GetElement(k)
		}
	}
}

func Benchmark_FindElem_parse_France_geofabrik_yml(b *testing.B) {
	c, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		_, _ = config.FindElem(c, "france")
	}
}

func Benchmark_Elem2preURL_parse_France_geofabrik_yml(b *testing.B) {
	c, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	france, err := config.FindElem(c, "france")
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		_, _ = config.Elem2preURL(c, france)
	}
}

func Benchmark_Elem2URL_parse_France_geofabrik_yml(b *testing.B) {
	c, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	france, err := config.FindElem(c, "france")
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		_, _ = config.Elem2URL(c, france, formats.FormatState)
	}
}
