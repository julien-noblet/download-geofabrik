package main

import (
	"log"

	"github.com/gocolly/colly"
)

const bbbikePb = 236

func bbbikeParseList(e *colly.HTMLElement, ext *Ext, c *colly.Collector) {
	e.ForEach("a", func(_ int, el *colly.HTMLElement) {
		href := el.Request.AbsoluteURL(el.Attr("href"))
		if *fVerbose && !*fQuiet && !*fProgress {
			log.Println("Parse:", href)
		}
		err := c.Visit(href)
		if err != nil {
			log.Panicln(err)
		}
	})
}

func bbbikeGetName(h3 string) string {
	ret := h3[17:] // remove "OSM extracts for "
	return ret
}

func bbbikeParseSidebar(e *colly.HTMLElement, ext *Ext, c *colly.Collector) {
	name := bbbikeGetName(e.ChildText("h3"))
	element := Element{
		ID:   name,
		Name: name,
		File: name + "/" + name,
		Formats: []string{
			"osm.pbf",
			"osm.gz",
			"shp.zip",
		},
	}
	if *fVerbose && !*fQuiet && !*fProgress {
		log.Println("Add", name)
	}
	err := ext.mergeElement(&element)
	if err != nil {
		panic(err)
	}
}
