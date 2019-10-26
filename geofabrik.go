package main

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

// Geofabrik Scrapper
type Geofabrik struct {
	*Scrapper
}

var geofabrik = Geofabrik{
	Scrapper: &Scrapper{
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
		FormatDefinition: formatDefinitions{
			formatOsmBz2:  {ID: formatOsmBz2, Loc: "-latest.osm.bz2"},
			"osm.bz2.md5": {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
			formatOsmPbf:  {ID: formatOsmPbf, Loc: "-latest.osm.pbf"},
			"osm.pbf.md5": {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
			formatPoly:    {ID: formatPoly, Loc: ".poly"},
			formatKml:     {ID: formatKml, Loc: ".kml"},
			formatState:   {ID: formatState, Loc: "-updates/state.txt"},
			formatShpZip:  {ID: formatShpZip, Loc: "-latest-free.shp.zip"},
		},
	},
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
			id, extension := FileExt(href)
			if extension == "html" {
				parent, pp := GetParent(href)
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
				element := Element{
					ID:     id,
					Name:   sub.Text,
					Parent: parent,
					Meta:   true,
				}
				if !g.Config.Exist(parent) && parent != "" { // Case of parent should exist not already in ElementSlice
					gparent, _ := GetParent(pp)
					if *fVerbose && !*fQuiet && !*fProgress {
						log.Println("Create Meta", element.Parent, "parent:", gparent, pp)
					}
					if gp := MakeParent(&element, gparent); gp != nil {
						catch(g.Config.mergeElement(gp))
					}
				}
				catch(g.Config.mergeElement(&element))
				if *fVerbose && !*fQuiet && !*fProgress {
					log.Println("Add:", href)
				}
				if err := c.Visit(href); err != nil && err != colly.ErrAlreadyVisited {
					catch(err)
				}
			}
		})
	})
}

// ParseFormat Add Extension to ID
// In this case, we add a kml and a state for all .osm.pbf
func (g *Geofabrik) ParseFormat(id, format string) {
	g.Scrapper.ParseFormat(id, format)

	if format == formatOsmPbf {
		g.Config.AddExtension(id, formatKml)   // not checked!
		g.Config.AddExtension(id, formatState) // not checked!
	}
}

func (g *Geofabrik) parseLi(e *colly.HTMLElement, c *colly.Collector) { //nolint:unparam
	e.ForEach("a", func(_ int, el *colly.HTMLElement) {
		_, format := FileExt(el.Attr("href"))
		id, _ := FileExt(el.Request.URL.String()) // id can't be extracted from href
		if id == "georgia" {                      // Exception
			parent, _ := GetParent(el.Request.AbsoluteURL(el.Attr("href")))
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
