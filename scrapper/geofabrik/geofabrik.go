package geofabrik

import (
	"errors"
	"regexp"

	"github.com/apex/log"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper"
)

// Constants for magic numbers and URLs.
const (
	progressBarCount = 509 // number of elements
	parallelism      = 20  // number of parallel downloads
	baseURL          = "https://download.geofabrik.de"
	startURL         = baseURL + "/"
)

// Geofabrik Scrapper.
type Geofabrik struct {
	*scrapper.Scrapper
}

// GetDefault returns the default configuration for Geofabrik scrapper.
func GetDefault() *Geofabrik {
	urlFilters := []*regexp.Regexp{
		regexp.MustCompile(`https://download\.geofabrik\.de/.+\.html$`),
		regexp.MustCompile(`https://download\.geofabrik\.de/$`),
	}

	formatDefinition := formats.FormatDefinitions{
		"osm.bz2.md5":         {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
		"osm.pbf.md5":         {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
		formats.FormatKml:     {ID: formats.FormatKml, Loc: ".kml"},
		formats.FormatMBTiles: {ID: formats.FormatMBTiles, Loc: "-latest-free.mbtiles.zip", ToLoc: "latest-free.mbtiles.zip"},
		formats.FormatOsmBz2:  {ID: formats.FormatOsmBz2, Loc: "-latest.osm.bz2"},
		formats.FormatOsmPbf:  {ID: formats.FormatOsmPbf, Loc: "-latest.osm.pbf"},
		formats.FormatPoly:    {ID: formats.FormatPoly, Loc: ".poly"},
		formats.FormatShpZip:  {ID: formats.FormatShpZip, Loc: "-shortbread-1.0.mbtiles"},
		formats.FormatState:   {ID: formats.FormatState, Loc: "-updates/state.txt"},
	}

	return &Geofabrik{
		Scrapper: &scrapper.Scrapper{
			PB:               progressBarCount,
			Async:            true,
			Parallelism:      parallelism,
			MaxDepth:         0,
			AllowedDomains:   []string{`download.geofabrik.de`},
			BaseURL:          baseURL,
			StartURL:         startURL,
			URLFilters:       urlFilters,
			FormatDefinition: formatDefinition,
		},
	}
}

// Collector represents Geofabrik's scrapper.
func (g *Geofabrik) Collector() *colly.Collector {
	myCollector := g.Scrapper.Collector()
	myCollector.OnHTML("#subregions", func(e *colly.HTMLElement) {
		g.ParseSubregion(e, myCollector)
	})
	myCollector.OnHTML("#specialsubregions", func(e *colly.HTMLElement) {
		g.ParseSubregion(e, myCollector)
	})
	myCollector.OnHTML("li", func(e *colly.HTMLElement) {
		g.ParseLi(e, myCollector)
	})

	return myCollector
}

// ParseSubregion parses the subregion information from the HTML.
func (g *Geofabrik) ParseSubregion(e *colly.HTMLElement, myCollector *colly.Collector) {
	e.ForEach("td.subregion", func(_ int, el *colly.HTMLElement) {
		el.ForEach("a", func(_ int, sub *colly.HTMLElement) {
			href := sub.Request.AbsoluteURL(sub.Attr("href"))
			myID, extension := scrapper.FileExt(href)

			if extension == "html" {
				g.handleHTMLExtension(sub, href, myID, myCollector)
			}
		})
	})
}

// handleHTMLExtension handles the HTML extension case.
func (g *Geofabrik) handleHTMLExtension(sub *colly.HTMLElement, href, myID string, myCollector *colly.Collector) {
	parent, parentPath := scrapper.GetParent(href)
	myID, file := g.handleSpecialCases(myID, parent)

	myElement := element.Element{
		ID:     myID,
		Name:   sub.Text,
		Parent: parent,
		Meta:   true,
		File:   file,
	}

	if !g.Config.Exist(parent) && parent != "" {
		gparent, _ := scrapper.GetParent(parentPath)
		log.Debugf("Create Meta %s parent: %s %v", myElement.Parent, gparent, parentPath)

		if gp := element.CreateParentElement(&myElement, gparent); gp != nil {
			if err := g.Config.MergeElement(gp); err != nil {
				log.WithError(err).Errorf("can't merge %s", myElement.Name)
			}
		}
	}

	if err := g.Config.MergeElement(&myElement); err != nil {
		log.WithError(err).Errorf("can't merge %s", myElement.Name)
	}

	log.Debugf("Add: %s", href)

	if err := myCollector.Visit(href); err != nil && !errors.Is(err, colly.ErrAlreadyVisited) {
		log.WithError(err).Error("can't get url")
	}
}

// handleSpecialCases handles special cases for certain IDs.
func (g *Geofabrik) handleSpecialCases(myID, parent string) (newID, file string) {
	const georgia = "georgia"

	switch myID {
	case georgia:
		switch parent {
		case "us":
			myID = georgia + "-us"
			file = georgia

		case "europe":
			myID = georgia + "-eu"
			file = georgia
		}

	case "guatemala":
		if parent == "south-america" {
			myID = "guatemala-south-america"
			file = "guatemala"
		}
	}

	return myID, file
}

// ParseFormat adds extensions to the ID.
func (g *Geofabrik) ParseFormat(id, format string) {
	g.Scrapper.ParseFormat(id, format)

	if format == formats.FormatOsmPbf {
		g.Config.AddExtension(id, formats.FormatKml)
		g.Config.AddExtension(id, formats.FormatState)
	}
}

// ParseLi parses the list items from the HTML.
func (g *Geofabrik) ParseLi(e *colly.HTMLElement, _ *colly.Collector) {
	e.ForEach("a", func(_ int, element *colly.HTMLElement) {
		_, format := scrapper.FileExt(element.Attr("href"))

		myID, _ := scrapper.FileExt(element.Request.URL.String())
		grandParent, _ := scrapper.GetParent(element.Request.AbsoluteURL(element.Attr("href")))
		myID, _ = g.handleSpecialCases(myID, grandParent)

		g.ParseFormat(myID, format)
	})
}
