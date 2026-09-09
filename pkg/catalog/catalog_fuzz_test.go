package catalog_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func FuzzCatalogUnmarshalYAML(f *testing.F) {
	geofabrikData, err := os.ReadFile("../../geofabrik.yml")
	if err == nil {
		f.Add(geofabrikData)
	}

	f.Add([]byte(""))
	f.Add([]byte("baseURL: https://example.com\nformats:\n  osm.pbf:\n    id: osm.pbf\n    loc: .osm.pbf\n"))
	f.Add([]byte("elements:\n  france:\n    id: france\n    name: France\n    formats:\n      - osm.pbf\n"))
	f.Add([]byte("invalid: yaml: [}"))

	f.Fuzz(func(_ *testing.T, data []byte) {
		cat := catalog.New()
		if err := yaml.Unmarshal(data, cat); err != nil {
			return
		}

		_ = cat.SortedKeys()
	})
}

func FuzzCatalogFind(f *testing.F) {
	cat, err := catalog.LoadFile("../../geofabrik.yml")
	require.NoError(f, err)

	f.Add("france")
	f.Add("rhone-alpes")
	f.Add("rhone_alpes")
	f.Add("non_existent")
	f.Add("")
	f.Add("../../../etc/passwd")
	f.Add("\x00\xff")

	f.Fuzz(func(_ *testing.T, query string) {
		_ = cat.Exist(query)
		_, _ = cat.Get(query)
		_, _ = cat.Find(query)
	})
}

func FuzzCatalogResolveURL(f *testing.F) {
	cat, err := catalog.LoadFile("../../geofabrik.yml")
	require.NoError(f, err)

	elem, err := cat.Find("france")
	require.NoError(f, err)

	f.Add("osm.pbf")
	f.Add("osm.pbf.md5")
	f.Add("poly")
	f.Add("unknown_format")
	f.Add("")
	f.Add("../../evil")

	f.Fuzz(func(_ *testing.T, format string) {
		_, _ = cat.ResolveURL(elem, format)
	})
}
