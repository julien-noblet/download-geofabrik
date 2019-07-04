package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

//OpenstreetmapFR Scrapper
type OpenstreetmapFR struct {
	*Scrapper
}

var openstreetmapFR = OpenstreetmapFR{
	Scrapper: &Scrapper{
		PB:             144,
		Async:          true,
		Parallelism:    20,
		MaxDepth:       0,
		AllowedDomains: []string{`download.openstreetmap.fr`},
		BaseURL:        `https://download.openstreetmap.fr/extracts`,
		StartURL:       `https://download.openstreetmap.fr/`,
		URLFilters: []*regexp.Regexp{
			//regexp.MustCompile(`https://download\.openstreetmap\.fr/([extracts|polygons]\w.+|)`),
			regexp.MustCompile(`https://download\.openstreetmap\.fr/$`),
			regexp.MustCompile(`https://download\.openstreetmap\.fr/extracts/(\w.+|)$`),
			regexp.MustCompile(`https://download\.openstreetmap\.fr/polygons/(\w.+|)$`),
		},
		FormatDefinition: formatDefinitions{
			"osm.pbf": {ID: "osm.pbf", Loc: "-latest.osm.pbf"},
			"poly":    {ID: "poly", Loc: ".poly", BasePath: "../polygons/"},
			"state":   {ID: "state", Loc: ".state.txt"},
		},
	},
}

//Collector represent geofabrik's scrapper
func (o *OpenstreetmapFR) Collector(options ...interface{}) *colly.Collector {
	c := o.Scrapper.Collector(options)
	c.OnHTML("a", func(e *colly.HTMLElement) {
		o.parse(e, c)
	})
	return c
}

func openstreetmapFRGetParent(href string) (string, []string) {
	parents := strings.Split(href, "/")
	var parent string
	if len(parents) > 4 {
		parent = parents[len(parents)-2] // Get x in this kind of url http(s)://1/2/.../x/
	} else {
		parent = ""
	}
	if strings.EqualFold(parent, "extracts") { // should I try == or a switch?
		parent = ""
	} else if strings.EqualFold(parent, "polygons") {
		parent = ""
	}
	return parent, parents
}

func (o *OpenstreetmapFR) makeParents(parent string, gparents []string) {
	if parent != "" {
		var gparent string
		if gparents == nil || len(gparents) < 3 {
			gparent = ""
		} else {
			gparent = gparents[len(gparents)-3] // Remove 2 last
			if gparent == "http:" || gparent == "osm.fr" || gparent == "extracts" || gparent == "polygons" {
				gparent = ""
			}
		}
		if !o.Config.Exist(parent) {
			element := Element{
				Parent: gparent,
				Name:   parent,
				ID:     parent,
				Meta:   true,
			}
			err := o.Config.mergeElement(&element)
			if err != nil {
				catch(fmt.Errorf("can't merge element, %v", err))
				// Panic
			}
			if gparent != "" {
				o.makeParents(gparent, gparents[:len(gparents)-1])
			}
		}
	}
}

func (o *OpenstreetmapFR) parseHref(href string) {
	if *fVerbose && !*fQuiet && !*fProgress {
		log.Println("Parsing:", href)
	}
	//	if !strings.Contains(href, "?") && !strings.Contains(href, "-latest") && href[0] != '/' && !strings.EqualFold(href, "cgi-bin/") {
	if !strings.Contains(href, "?") && !strings.Contains(href, "-latest") && href[0] != '/' {
		parent, parents := openstreetmapFRGetParent(href)
		if !o.Config.Exist(parent) {
			o.makeParents(parent, parents)
		}
		valsplit := strings.Split(parents[len(parents)-1], ".")
		if valsplit[0] != "" {
			if *fVerbose && !*fQuiet && !*fProgress {
				log.Println("Parsing", valsplit[0])
				// Output:
				// parsing valsplit[0]
			}
			extention := strings.Join(valsplit[1:], ".")
			if strings.Contains(extention, "state.txt") { // I'm shure it can be refactorized
				extention = "state"
			}
			if *fVerbose && !*fQuiet && !*fProgress {
				log.Println("Add", extention, "format")
			}
			element := Element{
				Parent: parent,
				Name:   valsplit[0],
				ID:     valsplit[0],
				Meta:   false,
			}
			if !o.Config.Exist(valsplit[0]) {
				element.Formats = append(element.Formats, extention)
				err := o.Config.mergeElement(&element)
				if err != nil {
					catch(fmt.Errorf("can't merge element, %v", err))
					// Panic
				}
			} else {
				if *fVerbose && !*fQuiet && !*fProgress {
					log.Println(valsplit[0], "already exist")
					log.Println("Merging formats")
				}
				o.Config.AddExtension(valsplit[0], extention)
			}
		}
	}
}

func (o *OpenstreetmapFR) parse(e *colly.HTMLElement, c *colly.Collector) {
	href := e.Request.AbsoluteURL(e.Attr("href"))
	if href[len(href)-1] == '/' {
		if *fVerbose && !*fQuiet && !*fProgress {
			log.Println("Next:", href)
		}
		if err := c.Visit(href); err != nil && err != colly.ErrAlreadyVisited {
			if err != colly.ErrNoURLFiltersMatch {
				catch(err)
			} else {
				if *fVerbose && !*fProgress && !*fQuiet {
					log.Printf("URL: %v is not matching URLFilters\n", href)
				}
			}
		}
	} else {
		o.parseHref(href)
	}

}
