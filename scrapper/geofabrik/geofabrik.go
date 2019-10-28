package geofabrik

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper"
	"github.com/spf13/viper"
)

// Geofabrik Scrapper
type Geofabrik struct {
	*scrapper.Scrapper
}

func GetDefault() *Geofabrik {
	return &Geofabrik{
		Scrapper: &scrapper.Scrapper{
			PB:             402,
			Async:          true,
			Parallelism:    20,
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
			if extension == "html" {
				parent, pp := scrapper.GetParent(href)
				if id == "georgia" {
					switch parent {
					case "us":
						id = "georgia-us" //nolint:goconst
					case "europe":
						id = "georgia-eu"
					}
				}
				if id == "guatemala" {
					if parent == "south-america" {
						id = "guatemala-south-america"
					}
				}
				el := element.Element{
					ID:     id,
					Name:   sub.Text,
					Parent: parent,
					Meta:   true,
				}
				if !g.Config.Exist(parent) && parent != "" { // Case of parent should exist not already in Slice
					gparent, _ := scrapper.GetParent(pp)
					if viper.GetBool("verbose") && !viper.GetBool("quiet") && !viper.GetBool("progress") {
						log.Println("Create Meta", el.Parent, "parent:", gparent, pp)
					}
					if gp := element.MakeParent(&el, gparent); gp != nil {
						if err := g.Config.MergeElement(gp); err != nil {
							log.Panicln(err)
						}
					}
				}
				if err := g.Config.MergeElement(&el); err != nil {
					log.Panicln(err)
				}
				if viper.GetBool("verbose") && !viper.GetBool("quiet") && !viper.GetBool("progress") {
					log.Println("Add:", href)
				}
				if err := c.Visit(href); err != nil && err != colly.ErrAlreadyVisited {
					log.Panicln(err)
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

func (g *Geofabrik) parseLi(e *colly.HTMLElement, c *colly.Collector) { //nolint:unparam
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
		g.ParseFormat(id, format)
	})
}
