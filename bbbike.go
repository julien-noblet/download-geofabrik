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
		c.Visit(href)
	})
}

func bbbikeGetName(h3 string) string {
	// remove "OSM extracts for "
	ret := h3[17:]
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
			"shp.zip",
		},
	}
	/*e.ForEach("a.download_link", func(_ int, el *colly.HTMLElement) {
		switch el.Text {
		case " Protocolbuffer (PBF) ":
			element.Formats = append(element.Formats, "osm.pbf")
		case " OSM XML gzip'd ":
			element.Formats = append(element.Formats, "osm.gz")
		case " Shapefile (Esri) ":
			element.Formats = append(element.Formats, "shp.zip")
		}

	})
	*/
	if *fVerbose && !*fQuiet && !*fProgress {
		log.Println("Add", name)
	}
	err := ext.mergeElement(&element)
	if err != nil {
		panic(err)
	}
}
