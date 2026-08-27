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

	// Short path <= minParentListLength (4)
	p, parts := getParent("http://example.com/file.pbf")
	assert.Empty(t, p)
	assert.Len(t, parts, 4)

	// Path with extracts as parent
	p, _ = getParent("https://download.openstreetmap.fr/extracts/europe.pbf")
	assert.Empty(t, p)

	// Path with polygons as parent
	p, _ = getParent("https://download.openstreetmap.fr/polygons/europe.pbf")
	assert.Empty(t, p)

	// Normal path
	p, _ = getParent("https://download.openstreetmap.fr/extracts/europe/france.osm.pbf")
	assert.Equal(t, "europe", p)
}

func TestInternal_GetGparent(t *testing.T) {
	t.Parallel()

	// Short path < minParentListLength
	assert.Empty(t, getGparent([]string{"a", "b"}))

	// Blocklisted gparent
	assert.Empty(t, getGparent([]string{"https:", "", "download.openstreetmap.fr", "europe", "france.osm.pbf"}))
	assert.Empty(t, getGparent([]string{"https:", "", "extracts", "europe", "france.osm.pbf"}))
	assert.Empty(t, getGparent([]string{"https:", "", "polygons", "europe", "france.osm.pbf"}))

	// Valid grandparent
	assert.Equal(t, "europe", getGparent([]string{"https:", "", "download.openstreetmap.fr", "extracts", "europe", "france", "paris.osm.pbf"}))
}
