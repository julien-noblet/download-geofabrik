package main

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

// Bbbike Scrapper
type Bbbike struct {
	*Scrapper
}

var bbbike = Bbbike{
	Scrapper: &Scrapper{
		PB:             236,
		Async:          true,
		Parallelism:    20,
		MaxDepth:       0,
		AllowedDomains: []string{`download.bbbike.org`},
		BaseURL:        `https://download.bbbike.org/osm/bbbike`,
		StartURL:       `https://download.bbbike.org/osm/bbbike/`,
		URLFilters: []*regexp.Regexp{
			regexp.MustCompile(`https://download\.bbbike\.org/osm/bbbike/[A-Z].+$`),
			regexp.MustCompile(`https://download\.bbbike\.org/osm/bbbike/$`),
		},
		FormatDefinition: formatDefinitions{
			formatOsmPbf: {ID: formatOsmPbf, Loc: ".osm.pbf"},
			formatShpZip: {ID: formatShpZip, Loc: ".osm.shp.zip"},
			formatOsmGz:  {ID: formatOsmGz, Loc: ".osm.gz"},
			formatPoly:   {ID: formatPoly, Loc: ".poly"},
		},
	},
}

// Collector represent geofabrik's scrapper
func (b *Bbbike) Collector(options ...interface{}) *colly.Collector {
	c := b.Scrapper.Collector(options)
	c.OnHTML("div.list tbody", func(e *colly.HTMLElement) {
		b.parseList(e, c)
	})
	c.OnHTML("#sidebar", func(e *colly.HTMLElement) {
		b.parseSidebar(e, c)
	})

	return c
}

func (b *Bbbike) parseList(e *colly.HTMLElement, c *colly.Collector) {
	e.ForEach("a", func(_ int, el *colly.HTMLElement) {
		href := el.Request.AbsoluteURL(el.Attr("href"))
		if *fVerbose && !*fQuiet && !*fProgress {
			log.Println("Parse:", href)
		}
		if err := c.Visit(href); err != nil && err != colly.ErrNoURLFiltersMatch { // Not matching
			catch(err)
		}
	})
}

func bbbikeGetName(h3 string) string {
	ret := h3[17:] // remove "OSM extracts for "
	return ret
}

func (b *Bbbike) parseSidebar(e *colly.HTMLElement, c *colly.Collector) { //nolint:unparam
	name := bbbikeGetName(e.ChildText("h3"))
	element := Element{
		ID:   name,
		Name: name,
		File: name + "/" + name,
		Formats: elementFormats{
			formatOsmPbf,
			formatOsmGz,
			formatShpZip,
		},
	}

	if *fVerbose && !*fQuiet && !*fProgress {
		log.Println("Add", name)
	}

	if err := b.Config.mergeElement(&element); err != nil {
		panic(err)
	}
}
