package geo2day

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
)

func TestSplitParent_EdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("scheme without path", func(t *testing.T) {
		t.Parallel()

		parent, path := splitParent("https://geo2day.com")
		assert.Empty(t, parent)
		assert.Empty(t, path)
	})

	t.Run("slash only", func(t *testing.T) {
		t.Parallel()

		parent, path := splitParent("/")
		assert.Empty(t, parent)
		assert.Empty(t, path)
	})
}

func TestProcessFileLink_EmptyRawID(t *testing.T) {
	t.Parallel()

	p := NewProvider()
	cat := catalog.New()
	p.processFileLink(cat, "/foo.pbf", "", "pbf", "test")
	assert.Empty(t, cat.Elements)
}
