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

// Geofabrik Scrapper
type Geofabrik struct {
	*scrapper.Scrapper
}

func GetDefault() *Geofabrik {
	return &Geofabrik{
		Scrapper: &scrapper.Scrapper{
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

// Collector represent geofabrik's scrapper
func (g *Geofabrik) Collector(options ...interface{}) *colly.Collector {
	c := g.Scrapper.Collector(options)
	c.OnHTML("#subregions", func(e *colly.HTMLElement) {
		g.parseSubregion(e, c)
	})
	c.OnHTML("#specialsubregions", func(e *colly.HTMLElement) {
		g.parseSubregion(e, c)
	})
	c.OnHTML("li", func(e *colly.HTMLElement) {
		g.parseLi(e, c)
	})

	return c
}

func (g *Geofabrik) parseSubregion(e *colly.HTMLElement, c *colly.Collector) {
	e.ForEach("td.subregion", func(_ int, el *colly.HTMLElement) {
		el.ForEach("a", func(_ int, sub *colly.HTMLElement) {
			href := sub.Request.AbsoluteURL(sub.Attr("href"))
			id, extension := scrapper.FileExt(href)
			var file string
			if extension == "html" {
				parent, pp := scrapper.GetParent(href)
				if id == "georgia" { //nolint:goconst // Georgia is in Europe & US
					switch parent {
					case "us":
						id = "georgia-us"
						file = "georgia"
					case "europe":
						id = "georgia-eu"
						file = "georgia"
					}
				}
				if id == "guatemala" && parent == "south-america" { //nolint:goconst // guatemala is also in central-america
					id = "guatemala-south-america"
					file = "guatemala"
				}
				el := element.Element{
					ID:     id,
					Name:   sub.Text,
					Parent: parent,
					Meta:   true,
				}
				if file != "" {
					el.File = file
				}
				if !g.Config.Exist(parent) && parent != "" { // Case of parent should exist not already in Slice
					gparent, _ := scrapper.GetParent(pp)
					log.Debugf("Create Meta %s parent: %s %v", el.Parent, gparent, pp)

					if gp := element.MakeParent(&el, gparent); gp != nil {
						if err := g.Config.MergeElement(gp); err != nil {
							log.WithError(err).Errorf("can't merge %s", el.Name)
						}
					}
				}
				if err := g.Config.MergeElement(&el); err != nil {
					log.WithError(err).Errorf("can't merge %s", el.Name)
				}
				log.Debugf("Add: %s", href)

				if err := c.Visit(href); err != nil && !errors.Is(err, colly.ErrAlreadyVisited) {
					log.WithError(err).Error("can't get url")
				}
			}
		})
	})
}

// ParseFormat Add Extension to ID
// In this case, we add a kml and a state for all .osm.pbf
func (g *Geofabrik) ParseFormat(id, format string) {
	g.Scrapper.ParseFormat(id, format)

	if format == formats.FormatOsmPbf {
		g.Config.AddExtension(id, formats.FormatKml)   // not checked!
		g.Config.AddExtension(id, formats.FormatState) // not checked!
	}
}

func (g *Geofabrik) parseLi(e *colly.HTMLElement, c *colly.Collector) { //nolint:unparam,lll // *colly.Collector is passed as param but unused in this case
	e.ForEach("a", func(_ int, el *colly.HTMLElement) {
		_, format := scrapper.FileExt(el.Attr("href"))
		id, _ := scrapper.FileExt(el.Request.URL.String()) // id can't be extracted from href
		if id == "georgia" {                               // Exception
			parent, _ := scrapper.GetParent(el.Request.AbsoluteURL(el.Attr("href")))
			switch parent {
			case "us":
				id = "georgia-us"
			case "europe":
				id = "georgia-eu"
			}
		}
		if id == "guatemala" {
			parent, _ := scrapper.GetParent(el.Request.AbsoluteURL(el.Attr("href")))
			if parent == "south-america" {
				id = "guatemala-south-america"
			}
		}
		g.ParseFormat(id, format)
	})
}
