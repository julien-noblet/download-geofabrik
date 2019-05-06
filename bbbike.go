package main

import (
	"log"
	"time"

	"github.com/gocolly/colly"
)

const (
	bbbikePb          = 236
	bbbikeAsync       = true
	bbbikeDomainGlob  = "*"
	bbbikeParallelism = 20
	bbbikeRandomDelay = 5 * time.Second
	bbbikeMaxDepth    = 0 // infinite
)

func bbbikeParseList(e *colly.HTMLElement, config *Config, c *colly.Collector) {
	e.ForEach("a", func(_ int, el *colly.HTMLElement) {
		href := el.Request.AbsoluteURL(el.Attr("href"))
		if *fVerbose && !*fQuiet && !*fProgress {
			log.Println("Parse:", href)
		}
		catch(c.Visit(href))
	})
}

func bbbikeGetName(h3 string) string {
	ret := h3[17:] // remove "OSM extracts for "
	return ret
}

func bbbikeParseSidebar(e *colly.HTMLElement, config *Config, c *colly.Collector) {
	name := bbbikeGetName(e.ChildText("h3"))
	element := Element{
		ID:   name,
		Name: name,
		File: name + "/" + name,
		Formats: elementFormats{
			"osm.pbf",
			"osm.gz",
			"shp.zip",
		},
	}
	if *fVerbose && !*fQuiet && !*fProgress {
		log.Println("Add", name)
	}
	if err := config.mergeElement(&element); err != nil {
		panic(err)
	}
}
