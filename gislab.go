package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

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

func gislabParse(e *colly.HTMLElement, ext *Ext) error {
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
			err := ext.mergeElement(&element)
			if err != nil {
				panic(err)
			}
		}
	})
	return nil
}
