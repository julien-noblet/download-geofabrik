package openstreetmapfr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternal_ResolveURL(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "https://example.com/file.osm.pbf", resolveURL("https://example.com", "https://example.com/file.osm.pbf"))
	assert.Equal(t, "http://example.com/file.osm.pbf", resolveURL("https://example.com", "http://example.com/file.osm.pbf"))
	assert.Equal(t, "https://example.com/base/file.osm.pbf", resolveURL("https://example.com/base", "file.osm.pbf"))
	assert.Equal(t, "https://example.com/base/file.osm.pbf", resolveURL("https://example.com/base/", "file.osm.pbf"))
}

func TestInternal_GetParent(t *testing.T) {
	t.Parallel()

	t.Run("short path", func(t *testing.T) {
		t.Parallel()

		parent, parts := getParent("http://example.com/file.pbf")
		assert.Empty(t, parent)
		assert.Len(t, parts, 4)
	})

	t.Run("extracts as parent", func(t *testing.T) {
		t.Parallel()

		parent, _ := getParent("https://download.openstreetmap.fr/extracts/europe.pbf")
		assert.Empty(t, parent)
	})

	t.Run("polygons as parent", func(t *testing.T) {
		t.Parallel()

		parent, _ := getParent("https://download.openstreetmap.fr/polygons/europe.pbf")
		assert.Empty(t, parent)
	})

	t.Run("normal path", func(t *testing.T) {
		t.Parallel()

		parent, _ := getParent("https://download.openstreetmap.fr/extracts/europe/france.osm.pbf")
		assert.Equal(t, "europe", parent)
	})
}

func TestInternal_GetGparent(t *testing.T) {
	t.Parallel()

	t.Run("short path", func(t *testing.T) {
		t.Parallel()

		assert.Empty(t, getGparent([]string{"a", "b"}))
	})

	t.Run("blocklisted gparent", func(t *testing.T) {
		t.Parallel()

		assert.Empty(t, getGparent([]string{"https:", "", "download.openstreetmap.fr", "europe", "france.osm.pbf"}))
		assert.Empty(t, getGparent([]string{"https:", "", "extracts", "europe", "france.osm.pbf"}))
		assert.Empty(t, getGparent([]string{"https:", "", "polygons", "europe", "france.osm.pbf"}))
	})

	t.Run("valid grandparent", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "europe", getGparent([]string{"https:", "", "download.openstreetmap.fr", "extracts", "europe", "france", "paris.osm.pbf"}))
	})
}
