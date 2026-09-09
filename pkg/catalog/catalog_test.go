package catalog_test

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

const geofabrikYml = "../../geofabrik.yml"

func TestCatalog_LoadFile(t *testing.T) {
	t.Parallel()

	cat, err := catalog.LoadFile(geofabrikYml)
	require.NoError(t, err)
	assert.NotNil(t, cat)
	assert.True(t, cat.Exist("france"))
	assert.True(t, cat.Exist("europe"))
	assert.True(t, cat.Exist("rhone-alpes"))
	assert.True(t, cat.Exist("rhone_alpes"))
	r1, found1 := cat.Get("rhone-alpes")
	assert.True(t, found1)

	r2, found2 := cat.Get("rhone_alpes")
	assert.True(t, found2)
	assert.Equal(t, r1.ID, r2.ID)

	f1, err := cat.Find("rhone_alpes")
	require.NoError(t, err)
	assert.Equal(t, "rhone-alpes", f1.ID)
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

	_, err = cat.ResolveURL(fr, catalog.FormatShpZip)
	require.Error(t, err)
	require.ErrorIs(t, err, catalog.ErrFormatNotFound)

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

func TestCatalog_DateResolution(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.BaseURL = "https://test.com"
	cat.Formats[catalog.FormatOsmPbf] = catalog.Format{
		ID:       catalog.FormatOsmPbf,
		Loc:      ".osm.pbf",
		BasePath: "czech_republic/",
	}
	cat.AddElement(&catalog.Element{
		ID:   "czech_republic",
		Name: "Czech Republic",
		File: "czech-republic",
	})
	cat.AddElement(&catalog.Element{
		ID:   "latest",
		Name: "Czech Republic (latest)",
		File: "czech_republic-2026-09-06",
	})

	assert.True(t, cat.Exist("2026-09-06"))
	elem, exists := cat.Get("2026-09-06")
	assert.True(t, exists)
	assert.Equal(t, "2026-09-06", elem.ID)
	assert.Equal(t, "czech_republic-2026-09-06", elem.File)

	elemPtr, err := cat.Find("2026-09-06")
	require.NoError(t, err)
	assert.Equal(t, "2026-09-06", elemPtr.ID)

	url, err := cat.ResolveURL(elemPtr, catalog.FormatOsmPbf)
	require.NoError(t, err)
	assert.Equal(t, "https://test.com/czech_republic/czech_republic-2026-09-06.osm.pbf", url)

	assert.False(t, cat.Exist("invalid-date"))

	_, exists = cat.Get("invalid-date")
	assert.False(t, exists)

	_, err = cat.Find("invalid-date")
	require.ErrorIs(t, err, catalog.ErrElementNotFound)
}

func TestCatalog_ConcurrentAccess(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.BaseURL = "https://download.example.com"
	cat.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: ".osm.pbf"}

	const (
		numGoroutines = 20
		iterations    = 50
	)

	var waitGroup sync.WaitGroup
	waitGroup.Add(numGoroutines * 3)

	// Writer goroutines: AddElement
	for goroutineIdx := range numGoroutines {
		go func(gID int) {
			defer waitGroup.Done()

			for iterIdx := range iterations {
				elemID := fmt.Sprintf("elem_%d_%d", gID, iterIdx)
				cat.AddElement(&catalog.Element{
					ID:      elemID,
					Name:    elemID,
					Formats: catalog.Formats{catalog.FormatOsmPbf},
				})
			}
		}(goroutineIdx)
	}

	// Reader goroutines: Len, Get, Find, Exists, GetFormat
	for goroutineIdx := range numGoroutines {
		go func(gID int) {
			defer waitGroup.Done()

			for iterIdx := range iterations {
				elemID := fmt.Sprintf("elem_%d_%d", gID, iterIdx)
				_ = cat.Len()
				_, _ = cat.Get(elemID)
				_, _ = cat.Find(elemID)
				_ = cat.Exists(elemID)
				_, _ = cat.GetFormat(catalog.FormatOsmPbf)
			}
		}(goroutineIdx)
	}

	// Modifier goroutines: AddExtension
	for goroutineIdx := range numGoroutines {
		go func(gID int) {
			defer waitGroup.Done()

			for iterIdx := range iterations {
				elemID := fmt.Sprintf("elem_%d_%d", gID, iterIdx)
				cat.AddExtension(elemID, "custom_format")
			}
		}(goroutineIdx)
	}

	waitGroup.Wait()
	assert.Positive(t, cat.Len())
}

func TestCatalog_SliceAliasing(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.AddElement(&catalog.Element{
		ID:      "test-element",
		Formats: catalog.Formats{"format-a", "format-b"},
	})

	elem1, ok := cat.Get("test-element")
	require.True(t, ok)

	elem2, err := cat.Find("test-element")
	require.NoError(t, err)

	// Mutating the returned formats slice must NOT mutate the catalog's internal state
	elem1.Formats[0] = "mutated-format"
	elem2.Formats[1] = "mutated-format-2"

	original, ok := cat.Get("test-element")
	require.True(t, ok)
	assert.Equal(t, catalog.Formats{"format-a", "format-b"}, original.Formats)
}

func Benchmark_Catalog_Exist(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		_ = cat.Exist("france")
	}
}

func TestCatalog_NilReceiver(t *testing.T) {
	t.Parallel()

	var cat *catalog.Catalog

	assert.Equal(t, 0, cat.Len())
	assert.Nil(t, cat.SortedKeys())
	assert.False(t, cat.Exists("france"))
	assert.False(t, cat.Exist("france"))

	_, exists := cat.Get("france")
	assert.False(t, exists)

	_, hasFormat := cat.GetFormat(catalog.FormatOsmPbf)
	assert.False(t, hasFormat)

	isHashable, hashExt, hashType := cat.IsHashable(catalog.FormatOsmPbf)
	assert.False(t, isHashable)
	assert.Empty(t, hashExt)
	assert.Empty(t, hashType)

	_, err := cat.Find("france")
	require.ErrorIs(t, err, catalog.ErrNilCatalog)

	err = cat.SaveFile("somefile.yml")
	require.ErrorIs(t, err, catalog.ErrNilCatalog)

	err = cat.Save(io.Discard)
	require.ErrorIs(t, err, catalog.ErrNilCatalog)

	err = cat.MergeElement(&catalog.Element{ID: "test"})
	require.ErrorIs(t, err, catalog.ErrNilCatalog)

	_, err = cat.ResolveURL(&catalog.Element{ID: "test"}, catalog.FormatOsmPbf)
	require.ErrorIs(t, err, catalog.ErrNilCatalog)

	_, err = cat.ResolvePreURL(&catalog.Element{ID: "test"})
	require.ErrorIs(t, err, catalog.ErrNilCatalog)
}

func Benchmark_Catalog_Get(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		_, _ = cat.Get("france")
	}
}

func Benchmark_Catalog_Find(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		_, _ = cat.Find("france")
	}
}

func Benchmark_Catalog_ResolveURL(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	fr, _ := cat.Find("france")

	for b.Loop() {
		_, _ = cat.ResolveURL(fr, catalog.FormatOsmPbf)
	}
}

func Benchmark_Catalog_SortedKeys(b *testing.B) {
	cat, err := catalog.LoadFile(geofabrikYml)
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		_ = cat.SortedKeys()
	}
}
