package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

//Gislab Scrapper
type Gislab struct {
	*Scrapper
}

var gislab = Gislab{
	Scrapper: &Scrapper{
		PB:             1,
		Async:          true,
		Parallelism:    1,
		MaxDepth:       0,
		AllowedDomains: []string{`be.gis-lab.info`, `data.gis-lab.info`, `raw.githubusercontent.com`},
		BaseURL:        `http://be.gis-lab.info/project/osm_dump`,
		StartURL:       `http://be.gis-lab.info/project/osm_dump/iframe.php`,
		URLFilters: []*regexp.Regexp{
			regexp.MustCompile(`http[s]?://be\.gis-lab\.info/project/osm_dump/iframe\.php$`), // Only 1 page!
		},
		FormatDefinition: formatDefinitions{
			"osm.pbf": format{ID: "osm.pbf", BaseURL: "http://data.gis-lab.info/osm_dump/dump", BasePath: "latest/", Loc: ".osm.pbf"},
			"osm.bz2": format{ID: "osm.bz2", BaseURL: "http://data.gis-lab.info/osm_dump/dump", BasePath: "latest/", Loc: ".osm.bz2"},
			"poly":    format{ID: "poly", BaseURL: "https://raw.githubusercontent.com/nextgis/osmdump_poly/master", Loc: ".poly"},
		},
	},
}

//Collector represent Gislab's scrapper
func (g *Gislab) Collector(options ...interface{}) *colly.Collector {
	c := g.Scrapper.Collector(options)
	c.OnHTML("table", func(e *colly.HTMLElement) {
		catch(g.parse(e, c))
	})
	return c
}

func gislabAddExt(td *colly.HTMLElement, element *Element) {
	if strings.Contains(td.ChildAttr("a", "href"), "osm.pbf") {
		if *fVerbose && !*fQuiet && !*fProgress {
			log.Println("Adding", td.ChildAttr("a", "href"))
		}
		element.Formats = append(element.Formats, "osm.pbf") // Not checked elements
	}
	if strings.Contains(td.ChildAttr("a", "href"), "osm.bz2") {
		if *fVerbose && !*fQuiet && !*fProgress {
			log.Println("Adding", td.ChildAttr("a", "href"))
		}
		element.Formats = append(element.Formats, "osm.bz2") // Not checked elements
	}
}

func (g *Gislab) parse(e *colly.HTMLElement, c *colly.Collector) error {
	e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
		if el.ChildText("a:nth-child(1)") != "" {
			element := Element{
				ID:   el.ChildText("td:nth-child(1)"),
				Name: el.ChildText("td:nth-child(2)"),
			}
			el.ForEach("td", func(_ int, td *colly.HTMLElement) {
				gislabAddExt(td, &element)
			})
			element.Formats = append(element.Formats, "poly") // Not checked but seems to be used for generating osm.pbf/osm.bz2
			if *fVerbose && !*fQuiet && !*fProgress {
				log.Println("Adding", element.Name)
			}
			catch(g.Config.mergeElement(&element))
		}
	})
	return nil
}
