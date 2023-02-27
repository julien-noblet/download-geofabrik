package geofabrik

import (
	"errors"
	"regexp"

	"github.com/apex/log"
	"github.com/gocolly/colly"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper"
)

// Geofabrik Scrapper.
type Geofabrik struct {
	*scrapper.Scrapper
}

func GetDefault() *Geofabrik {
	return &Geofabrik{
		Scrapper: &scrapper.Scrapper{ //nolint:exhaustruct // I'm lazy
			PB:             412, //nolint:gomnd // there is 412 items
			Async:          true,
			Parallelism:    20, //nolint:gomnd // use 20 threads for scrapping
			MaxDepth:       0,
			AllowedDomains: []string{`download.geofabrik.de`},
			BaseURL:        `https://download.geofabrik.de`,
			StartURL:       `https://download.geofabrik.de/`,
			URLFilters: []*regexp.Regexp{
				regexp.MustCompile(`https://download\.geofabrik\.de/.+\.html$`),
				regexp.MustCompile(`https://download\.geofabrik\.de/$`),
			},
			FormatDefinition: formats.FormatDefinitions{
				formats.FormatOsmBz2: {ID: formats.FormatOsmBz2, Loc: "-latest.osm.bz2"},
				"osm.bz2.md5":        {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
				formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: "-latest.osm.pbf"},
				"osm.pbf.md5":        {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
				formats.FormatPoly:   {ID: formats.FormatPoly, Loc: ".poly"},
				formats.FormatKml:    {ID: formats.FormatKml, Loc: ".kml"},
				formats.FormatState:  {ID: formats.FormatState, Loc: "-updates/state.txt"},
				formats.FormatShpZip: {ID: formats.FormatShpZip, Loc: "-latest-free.shp.zip"},
			},
		},
	}
}

// Collector represent geofabrik's scrapper.
func (g *Geofabrik) Collector() *colly.Collector {
	c := g.Scrapper.Collector()
	c.OnHTML("#subregions", func(e *colly.HTMLElement) {
		g.ParseSubregion(e, c)
	})
	c.OnHTML("#specialsubregions", func(e *colly.HTMLElement) {
		g.ParseSubregion(e, c)
	})
	c.OnHTML("li", func(e *colly.HTMLElement) {
		g.ParseLi(e, c)
	})

	return c
}

//nolint:cyclop // TODO : Refactoring?
func (g *Geofabrik) ParseSubregion(e *colly.HTMLElement, myCollector *colly.Collector) {
	e.ForEach("td.subregion", func(_ int, el *colly.HTMLElement) {
		el.ForEach("a", func(_ int, sub *colly.HTMLElement) {
			href := sub.Request.AbsoluteURL(sub.Attr("href"))
			myID, extension := scrapper.FileExt(href)
			var file string
			if extension == "html" { //nolint:nestif // TODO : Refactor?
				parent, parentPath := scrapper.GetParent(href)
				if myID == "georgia" { //nolint:goconst // Georgia is in Europe & US
					switch parent {
					case "us":
						myID = "georgia-us"
						file = "georgia"
					case "europe":
						myID = "georgia-eu"
						file = "georgia"
					}
				}
				if myID == "guatemala" && parent == "south-america" { //nolint:goconst // guatemala is also in central-america
					myID = "guatemala-south-america"
					file = "guatemala"
				}
				myElement := element.Element{ //nolint:exhaustruct // I'm lazy
					ID:     myID,
					Name:   sub.Text,
					Parent: parent,
					Meta:   true,
				}
				if file != "" {
					myElement.File = file
				}
				if !g.Config.Exist(parent) && parent != "" { // Case of parent should exist not already in Slice
					gparent, _ := scrapper.GetParent(parentPath)
					log.Debugf("Create Meta %s parent: %s %v", myElement.Parent, gparent, parentPath)

					if gp := element.MakeParent(&myElement, gparent); gp != nil {
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
		})
	})
}

// ParseFormat Add Extension to ID
// In this case, we add a kml and a state for all .osm.pbf .
func (g *Geofabrik) ParseFormat(id, format string) {
	g.Scrapper.ParseFormat(id, format)

	if format == formats.FormatOsmPbf {
		g.Config.AddExtension(id, formats.FormatKml)   // not checked!
		g.Config.AddExtension(id, formats.FormatState) // not checked!
	}
}

func (g *Geofabrik) ParseLi(e *colly.HTMLElement, c *colly.Collector) {
	e.ForEach("a", func(_ int, element *colly.HTMLElement) {
		_, format := scrapper.FileExt(element.Attr("href"))
		myID, _ := scrapper.FileExt(element.Request.URL.String()) // id can't be extracted from href
		if myID == "georgia" {                                    // Exception
			parent, _ := scrapper.GetParent(element.Request.AbsoluteURL(element.Attr("href")))
			switch parent {
			case "us":
				myID = "georgia-us"
			case "europe":
				myID = "georgia-eu"
			}
		}
		if myID == "guatemala" {
			parent, _ := scrapper.GetParent(element.Request.AbsoluteURL(element.Attr("href")))
			if parent == "south-america" {
				myID = "guatemala-south-america"
			}
		}
		g.ParseFormat(myID, format)
	})
}
