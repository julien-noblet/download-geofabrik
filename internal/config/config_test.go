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
	elem, _ := cfg.GetElement("item")

	// Case 1: 1 arg
	url, err := config.Elem2preURL(cfg, elem, "custom")
	require.NoError(t, err)
	assert.Equal(t, "https://example.com/custom/item", url)

	// Case 2: 2 args
	url, err = config.Elem2preURL(cfg, elem, "http://other.com", "path")
	require.NoError(t, err)
	assert.Equal(t, "http://other.com/path/item", url)

	// Case default: 0 args
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

	// Merge update checks
	el2 := &element.Element{ID: "e1", Parent: "p1", Formats: []string{"pbf"}}
	err = cfg.MergeElement(el2)
	require.NoError(t, err)

	e, err := cfg.GetElement("e1")
	require.NoError(t, err)
	assert.Contains(t, e.Formats, "osm")
	assert.Contains(t, e.Formats, "pbf")

	// Parent mismatch
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
