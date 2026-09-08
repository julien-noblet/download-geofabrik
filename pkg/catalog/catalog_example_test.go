package catalog_test

import (
	"fmt"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

// ExampleCatalog_Find demonstrates loading an element from a Catalog
// ExampleCatalog_Find demonstrates finding an element in a Catalog
// with error handling.
func ExampleCatalog_Find() {
	cat := catalog.New()
	cat.BaseURL = "https://download.geofabrik.de"

	cat.AddElement(&catalog.Element{
		ID:      "europe-france",
		Name:    "France",
		Parent:  "europe",
		Formats: catalog.Formats{"osm.pbf"},
	})

	// Exact match
	elem, err := cat.Find("europe-france")
	if err == nil {
		fmt.Printf("Exact: found=true, name=%s\n", elem.Name)
	}

	// Normalizes hyphens and underscores interchangeably
	elem, err = cat.Find("europe_france")
	if err == nil {
		fmt.Printf("Normalized: found=true, name=%s\n", elem.Name)
	}

	// Output:
	// Exact: found=true, name=France
	// Normalized: found=true, name=France
}

// ExampleCatalog_ResolveURL demonstrates constructing absolute download URLs
// for an element based on catalog format definitions.
func ExampleCatalog_ResolveURL() {
	cat := catalog.New()
	cat.BaseURL = "https://download.geofabrik.de"
	cat.Formats[catalog.FormatOsmPbf] = catalog.Format{
		ID:  catalog.FormatOsmPbf,
		Loc: "-latest.osm.pbf",
	}

	elem := &catalog.Element{
		ID:      "antarctica",
		Name:    "Antarctica",
		Formats: catalog.Formats{catalog.FormatOsmPbf},
	}

	cat.AddElement(elem)

	url, err := cat.ResolveURL(elem, catalog.FormatOsmPbf)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Download URL:", url)

	// Output:
	// Download URL: https://download.geofabrik.de/antarctica-latest.osm.pbf
}

// ExampleCatalog_MergeElement demonstrates adding new elements or merging
// format mappings into existing elements.
func ExampleCatalog_MergeElement() {
	cat := catalog.New()

	elem1 := &catalog.Element{
		ID:      "berlin",
		Name:    "Berlin",
		Formats: catalog.Formats{"osm.pbf"},
	}
	_ = cat.MergeElement(elem1)

	elem2 := &catalog.Element{
		ID:      "berlin",
		Name:    "Berlin",
		Formats: catalog.Formats{"osm.bz2"},
	}
	_ = cat.MergeElement(elem2)

	merged, _ := cat.Get("berlin")
	fmt.Printf("Formats count: %d\n", len(merged.Formats))

	// Output:
	// Formats count: 2
}

// ExampleGetMiniFormats demonstrates obtaining single-letter shorthand format
// abbreviations for CLI table output.
func ExampleGetMiniFormats() {
	formats := []string{"osm.pbf", "osm.bz2", "osm.gz"}
	shorthand := catalog.GetMiniFormats(formats)

	fmt.Println("Short format flags:", shorthand)

	// Output:
	// Short format flags: PBG
}
