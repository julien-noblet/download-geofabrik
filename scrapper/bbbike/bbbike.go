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

// Constants for magic numbers.
const (
	progressBarCount = 237
	parallelism      = 20
)

// Bbbike Scrapper.
type Bbbike struct {
	*scrapper.Scrapper
}

// GetDefault returns the default configuration for Bbbike scrapper.
func GetDefault() *Bbbike {
	return &Bbbike{
		Scrapper: &scrapper.Scrapper{
			PB:             progressBarCount,
			Async:          true,
			Parallelism:    parallelism,
			MaxDepth:       0,
			AllowedDomains: []string{`download.bbbike.org`},
			BaseURL:        `https://download.bbbike.org/osm/bbbike`,
			StartURL:       `https://download.bbbike.org/osm/bbbike/`,
			URLFilters: []*regexp.Regexp{
				regexp.MustCompile(`https://download\.bbbike\.org/osm/bbbike/[A-Z].+$`),
				regexp.MustCompile(`https://download\.bbbike\.org/osm/bbbike/$`),
			},
			FormatDefinition: formats.FormatDefinitions{
				formats.FormatCSV:            {ID: formats.FormatCSV, Loc: ".osm.csv.xz", ToLoc: ".osm.csv.xz"},
				formats.FormatGarminOSM:      {ID: formats.FormatGarminOSM, Loc: ".osm.garmin-osm.zip"},
				formats.FormatGarminOnroad:   {ID: formats.FormatGarminOnroad, Loc: ".osm.garmin-onroad-latin1.zip"},
				formats.FormatGarminOntrail:  {ID: formats.FormatGarminOntrail, Loc: ".osm.garmin-ontrail-latin1.zip"},
				formats.FormatGarminOpenTopo: {ID: formats.FormatGarminOpenTopo, Loc: ".osm.garmin-opentopo-latin1.zip"},
				formats.FormatGeoJSON:        {ID: formats.FormatGeoJSON, Loc: ".osm.geojson.xz", ToLoc: ".geojson.xz"},
				formats.FormatMBTiles:        {ID: formats.FormatMBTiles, Loc: ".osm.mbtiles-openmaptiles.zip", ToLoc: "osm.mbtiles-openmaptiles.zip"},
				formats.FormatMapsforge:      {ID: formats.FormatMapsforge, Loc: ".osm.mapsforge-osm.zip"},
				formats.FormatOsmGz:          {ID: formats.FormatOsmGz, Loc: ".osm.gz"},
				formats.FormatOsmPbf:         {ID: formats.FormatOsmPbf, Loc: ".osm.pbf"},
				formats.FormatPoly:           {ID: formats.FormatPoly, Loc: ".poly"},
				formats.FormatShpZip:         {ID: formats.FormatShpZip, Loc: ".osm.shp.zip"},
			},
		},
	}
}

// Collector represents Bbbike's scrapper.
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

// ParseList parses the list of elements from the HTML.
func (b *Bbbike) ParseList(e *colly.HTMLElement, c *colly.Collector) {
	e.ForEach("a", func(_ int, el *colly.HTMLElement) {
		href := el.Request.AbsoluteURL(el.Attr("href"))
		log.Debugf("Parse: %s", href)

		if err := c.Visit(href); err != nil && !errors.Is(err, colly.ErrNoURLFiltersMatch) {
			log.WithError(err).Error("can't get url")
		}
	})
}

// GetName extracts the name from the given string.
func GetName(h3 string) string {
	const prefixLength = 17

	return h3[prefixLength:] // remove "OSM extracts for "
}

// ParseSidebar parses the sidebar information from the HTML.
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
