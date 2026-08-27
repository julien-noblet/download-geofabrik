package ui_test

import (
	"bytes"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/ui"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestCatalog() *catalog.Catalog {
	cat := catalog.New()
	cat.BaseURL = "https://download.geofabrik.de"
	cat.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: ".osm.pbf"}
	cat.AddElement(&catalog.Element{ID: "europe", Name: "Europe", Meta: true})
	cat.AddElement(&catalog.Element{
		ID:      "france",
		Name:    "France",
		Parent:  "europe",
		Formats: catalog.Formats{catalog.FormatOsmPbf},
	})

	return cat
}

func TestPrintTable_Standard(t *testing.T) {
	t.Parallel()

	cat := createTestCatalog()

	var buf bytes.Buffer

	err := ui.PrintTable(cat, false, &buf)
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "france")
	assert.Contains(t, buf.String(), "France")
	assert.Contains(t, buf.String(), "Total elements: 2")
}

func TestPrintTable_Markdown(t *testing.T) {
	t.Parallel()

	cat := createTestCatalog()

	var buf bytes.Buffer

	err := ui.PrintTable(cat, true, &buf)
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "|")
	assert.Contains(t, buf.String(), "france")
}

func TestPrintJSON(t *testing.T) {
	t.Parallel()

	cat := createTestCatalog()

	var buf bytes.Buffer

	err := ui.PrintJSON(cat, &buf)
	require.NoError(t, err)
	assert.Contains(t, buf.String(), `"id": "france"`)
	assert.Contains(t, buf.String(), `"baseURL": "https://download.geofabrik.de"`)
}

func TestPrint_Nil(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	assert.NoError(t, ui.PrintTable(nil, false, &buf))
	assert.NoError(t, ui.PrintJSON(nil, &buf))
}
