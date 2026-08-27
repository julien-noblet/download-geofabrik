package catalog_test

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const geofabrikYml = "../../geofabrik.yml"

func TestCatalog_LoadFile(t *testing.T) {
	t.Parallel()

	cat, err := catalog.LoadFile(geofabrikYml)
	require.NoError(t, err)
	assert.NotNil(t, cat)
	assert.True(t, cat.Exist("france"))
	assert.True(t, cat.Exist("europe"))
	assert.False(t, cat.Exist("non_existent_region"))
}

func TestCatalog_LoadFile_Errors(t *testing.T) {
	t.Parallel()

	_, err := catalog.LoadFile("non_existent_file.yml")
	require.Error(t, err)
}

func TestCatalog_SaveAndLoad(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.BaseURL = "https://download.geofabrik.de"
	cat.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: ".osm.pbf"}
	cat.AddElement(&catalog.Element{
		ID:      "monaco",
		Name:    "Monaco",
		Formats: catalog.Formats{catalog.FormatOsmPbf},
	})

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "sub", "catalog.yml")

	err := cat.SaveFile(filePath)
	require.NoError(t, err)

	loaded, err := catalog.LoadFile(filePath)
	require.NoError(t, err)
	assert.True(t, loaded.Exist("monaco"))
	elem, exists := loaded.Get("monaco")
	assert.True(t, exists)
	assert.Equal(t, "Monaco", elem.Name)
	assert.True(t, elem.ContainsFormat(catalog.FormatOsmPbf))
}

func TestCatalog_ResolveURL(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.BaseURL = "https://download.geofabrik.de"
	cat.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf"}
	cat.Formats[catalog.FormatPoly] = catalog.Format{ID: catalog.FormatPoly, Loc: ".poly"}

	cat.AddElement(&catalog.Element{ID: "europe", Name: "Europe", Meta: true})
	cat.AddElement(&catalog.Element{
		ID:      "france",
		Name:    "France",
		Parent:  "europe",
		Formats: catalog.Formats{catalog.FormatOsmPbf, catalog.FormatPoly},
	})

	fr, err := cat.Find("france")
	require.NoError(t, err)

	url, err := cat.ResolveURL(fr, catalog.FormatOsmPbf)
	require.NoError(t, err)
	assert.Equal(t, "https://download.geofabrik.de/europe/france-latest.osm.pbf", url)

	// Format not supported
	_, err = cat.ResolveURL(fr, catalog.FormatShpZip)
	require.Error(t, err)
	require.ErrorIs(t, err, catalog.ErrFormatNotFound)

	// Nil element
	_, err = cat.ResolveURL(nil, catalog.FormatOsmPbf)
	require.Error(t, err)
}

func TestCatalog_ResolveURL_CycleDetection(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.BaseURL = "https://download.geofabrik.de"
	cat.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: ".osm.pbf"}

	// Create cyclical parent reference A -> B -> A
	cat.AddElement(&catalog.Element{ID: "nodeA", Parent: "nodeB", Formats: catalog.Formats{catalog.FormatOsmPbf}})
	cat.AddElement(&catalog.Element{ID: "nodeB", Parent: "nodeA", Formats: catalog.Formats{catalog.FormatOsmPbf}})

	elemA, err := cat.Find("nodeA")
	require.NoError(t, err)

	_, err = cat.ResolveURL(elemA, catalog.FormatOsmPbf)
	require.Error(t, err)
	assert.ErrorIs(t, err, catalog.ErrMaxHierarchyDepth)
}

func TestCatalog_MergeElement(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.AddElement(&catalog.Element{
		ID:      "paris",
		Name:    "Paris",
		Parent:  "france",
		Formats: catalog.Formats{catalog.FormatOsmPbf},
	})

	// Merge with new format
	err := cat.MergeElement(&catalog.Element{
		ID:      "paris",
		Parent:  "france",
		Formats: catalog.Formats{catalog.FormatPoly},
	})
	require.NoError(t, err)

	paris, exists := cat.Get("paris")
	assert.True(t, exists)
	assert.True(t, paris.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, paris.ContainsFormat(catalog.FormatPoly))

	// Conflicting parent error
	err = cat.MergeElement(&catalog.Element{
		ID:     "paris",
		Parent: "germany",
	})
	require.Error(t, err)
	assert.ErrorIs(t, err, catalog.ErrParentMismatch)
}

func TestCatalog_IsHashable(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: ".osm.pbf"}
	cat.Formats["osm.pbf.md5"] = catalog.Format{ID: "osm.pbf.md5", Loc: ".osm.pbf.md5"}

	ok, hashExt, hashType := cat.IsHashable(catalog.FormatOsmPbf)
	assert.True(t, ok)
	assert.Equal(t, "osm.pbf.md5", hashExt)
	assert.Equal(t, "md5", hashType)

	ok, _, _ = cat.IsHashable(catalog.FormatPoly)
	assert.False(t, ok)
}

func TestCatalog_SortedKeys_And_All(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.AddElement(&catalog.Element{ID: "c"})
	cat.AddElement(&catalog.Element{ID: "a"})
	cat.AddElement(&catalog.Element{ID: "b"})

	keys := cat.SortedKeys()
	assert.Equal(t, []string{"a", "b", "c"}, keys)

	var iteratedKeys []string
	for k := range cat.All() {
		iteratedKeys = append(iteratedKeys, k)
	}

	assert.Len(t, iteratedKeys, 3)
}

func TestCatalog_AddExtension(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.AddElement(&catalog.Element{ID: "monaco"})
	cat.AddExtension("monaco", catalog.FormatOsmPbf)

	elem, exists := cat.Get("monaco")
	assert.True(t, exists)
	assert.True(t, elem.ContainsFormat(catalog.FormatOsmPbf))

	// Non-existent element should not panic
	cat.AddExtension("unknown", catalog.FormatOsmPbf)
}

func TestCatalog_SaveStream(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.AddElement(&catalog.Element{ID: "monaco"})

	var buf bytes.Buffer

	err := cat.Save(&buf)
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "monaco")
}

func TestElement_Helpers(t *testing.T) {
	t.Parallel()

	elem := &catalog.Element{
		ID:     "monaco",
		Name:   "Monaco",
		Parent: "europe",
	}

	assert.True(t, elem.HasParent())
	assert.Equal(t, "monaco", elem.Filename())

	elem.File = "monaco-extract"
	assert.Equal(t, "monaco-extract", elem.Filename())

	elem.AddFormat(catalog.FormatOsmPbf)
	assert.True(t, elem.ContainsFormat(catalog.FormatOsmPbf))
	assert.False(t, elem.ContainsFormat(catalog.FormatPoly))

	parent := elem.CreateParentElement("world")
	require.NotNil(t, parent)
	assert.Equal(t, "europe", parent.ID)
	assert.Equal(t, "world", parent.Parent)
	assert.True(t, parent.Meta)

	var nilElem *catalog.Element
	assert.False(t, nilElem.HasParent())
	assert.Empty(t, nilElem.Filename())
	assert.False(t, nilElem.ContainsFormat("any"))
	assert.Nil(t, nilElem.CreateParentElement(""))
}

func TestFormat_GetFormats(t *testing.T) {
	t.Parallel()

	flags := map[string]bool{
		catalog.KeyOsmPbf: true,
		catalog.KeyPoly:   true,
	}

	formats := catalog.GetFormats(flags)
	assert.Equal(t, []string{catalog.FormatOsmPbf, catalog.FormatPoly}, formats)

	emptyFlags := map[string]bool{}
	defaultFormats := catalog.GetFormats(emptyFlags)
	assert.Equal(t, []string{catalog.FormatOsmPbf}, defaultFormats)
}

func TestFormat_GetMiniFormats(t *testing.T) {
	t.Parallel()

	mini := catalog.GetMiniFormats([]string{catalog.FormatOsmPbf, catalog.FormatState, catalog.FormatPoly})
	assert.Equal(t, "Psp", mini)

	assert.Empty(t, catalog.GetMiniFormats(nil))
}

func Benchmark_Catalog_Exist(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		_ = cat.Exist("france")
	}
}

func Benchmark_Catalog_Get(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		_, _ = cat.Get("france")
	}
}

func Benchmark_Catalog_Find(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		_, _ = cat.Find("france")
	}
}

func Benchmark_Catalog_ResolveURL(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	fr, _ := cat.Find("france")

	for range b.N {
		_, _ = cat.ResolveURL(fr, catalog.FormatOsmPbf)
	}
}

func Benchmark_Catalog_SortedKeys(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		_ = cat.SortedKeys()
	}
}
