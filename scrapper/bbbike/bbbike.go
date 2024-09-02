package bbbike

import (
	"errors"
	"regexp"

	"github.com/apex/log"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper"
)

// Bbbike Scrapper.
type Bbbike struct {
	*scrapper.Scrapper
}

func GetDefault() *Bbbike {
	return &Bbbike{
		Scrapper: &scrapper.Scrapper{ //nolint:exhaustruct // I'm lazy
			PB:             237, //nolint:gomnd // there is 237 element at this time
			Async:          true,
			Parallelism:    20, //nolint:gomnd // Use 20 threads for scraping
			MaxDepth:       0,
			AllowedDomains: []string{`download.bbbike.org`},
			BaseURL:        `https://download.bbbike.org/osm/bbbike`,
			StartURL:       `https://download.bbbike.org/osm/bbbike/`,
			URLFilters: []*regexp.Regexp{
				regexp.MustCompile(`https://download\.bbbike\.org/osm/bbbike/[A-Z].+$`),
				regexp.MustCompile(`https://download\.bbbike\.org/osm/bbbike/$`),
			},
			FormatDefinition: formats.FormatDefinitions{
				formats.FormatCSV:            {ID: formats.FormatCSV, Loc: ".osm.csv.xz", ToLoc: ".osm.csv.xz", BasePath: "", BaseURL: ""},
				formats.FormatGarminOSM:      {ID: formats.FormatGarminOSM, Loc: ".osm.garmin-osm.zip", ToLoc: "", BasePath: "", BaseURL: ""},
				formats.FormatGarminOnroad:   {ID: formats.FormatGarminOnroad, Loc: ".osm.garmin-onroad-latin1.zip", ToLoc: "", BasePath: "", BaseURL: ""},
				formats.FormatGarminOntrail:  {ID: formats.FormatGarminOntrail, Loc: ".osm.garmin-ontrail-latin1.zip", ToLoc: "", BasePath: "", BaseURL: ""},
				formats.FormatGarminOpenTopo: {ID: formats.FormatGarminOpenTopo, Loc: ".osm.garmin-opentopo-latin1.zip", ToLoc: "", BasePath: "", BaseURL: ""},
				formats.FormatGeoJSON:        {ID: formats.FormatGeoJSON, Loc: ".osm.geojson.xz", ToLoc: ".geojson.xz", BasePath: "", BaseURL: ""},
				formats.FormatMBTiles:        {ID: formats.FormatMBTiles, Loc: ".osm.mbtiles-openmaptiles.zip", ToLoc: "osm.mbtiles-openmaptiles.zip", BasePath: "", BaseURL: ""}, //nolint:lll // I don't want to split this line
				formats.FormatMapsforge:      {ID: formats.FormatMapsforge, Loc: ".osm.mapsforge-osm.zip", ToLoc: "", BasePath: "", BaseURL: ""},
				formats.FormatOsmGz:          {ID: formats.FormatOsmGz, Loc: ".osm.gz", ToLoc: "", BasePath: "", BaseURL: ""},
				formats.FormatOsmPbf:         {ID: formats.FormatOsmPbf, Loc: ".osm.pbf", ToLoc: "", BasePath: "", BaseURL: ""},
				formats.FormatPoly:           {ID: formats.FormatPoly, Loc: ".poly", ToLoc: "", BasePath: "", BaseURL: ""},
				formats.FormatShpZip:         {ID: formats.FormatShpZip, Loc: ".osm.shp.zip", ToLoc: "", BasePath: "", BaseURL: ""},
			},
		},
	}
}

// Collector represent geofabrik's scrapper.
func (b *Bbbike) Collector() *colly.Collector {
	myCollector := b.Scrapper.Collector()
	myCollector.OnHTML("div.list tbody", func(e *colly.HTMLElement) {
		b.ParseList(e, myCollector)
	})
	myCollector.OnHTML("#sidebar", func(e *colly.HTMLElement) {
		b.ParseSidebar(e, myCollector)
	})

	return myCollector
}

func (b *Bbbike) ParseList(e *colly.HTMLElement, c *colly.Collector) {
	e.ForEach("a", func(_ int, el *colly.HTMLElement) {
		href := el.Request.AbsoluteURL(el.Attr("href"))
		log.Debugf("Parse: %s", href)

		if err := c.Visit(href); err != nil && !errors.Is(err, colly.ErrNoURLFiltersMatch) { // Not matching
			log.WithError(err).Error("can't get url")
		}
	})
}

func GetName(h3 string) string {
	ret := h3[17:] // remove "OSM extracts for "

	return ret
}

func (b *Bbbike) ParseSidebar(e *colly.HTMLElement, _ *colly.Collector) {
	name := GetName(e.ChildText("h3"))
	myElement := element.Element{
		ID:     name,
		Name:   name,
		File:   name + "/" + name,
		Parent: "",
		Formats: element.Formats{
			formats.FormatCSV,
			formats.FormatGarminOSM,
			formats.FormatGarminOnroad,
			formats.FormatGarminOntrail,
			formats.FormatGarminOpenTopo,
			formats.FormatGeoJSON,
			formats.FormatMBTiles,
			formats.FormatMapsforge,
			formats.FormatOsmGz,
			formats.FormatOsmPbf,
			formats.FormatPoly,
			formats.FormatShpZip,
		},
		Meta: false,
	}

	log.Debugf("Add %s", name)

	if err := b.Config.MergeElement(&myElement); err != nil {
		log.WithError(err).Errorf("can't merge %s", myElement.Name)
	}
}
