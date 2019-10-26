package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// OpenstreetmapFR Scrapper
type OpenstreetmapFR struct {
	*Scrapper
}

var (
	t               = time.Second * 30
	openstreetmapFR = OpenstreetmapFR{
		Scrapper: &Scrapper{
			PB:             149,
			Async:          true,
			Parallelism:    20,
			MaxDepth:       0,
			AllowedDomains: []string{`download.openstreetmap.fr`},
			BaseURL:        `https://download.openstreetmap.fr/extracts`,
			StartURL:       `https://download.openstreetmap.fr/`,
			URLFilters: []*regexp.Regexp{
				regexp.MustCompile(`https://download\.openstreetmap\.fr/$`),
				regexp.MustCompile(`https://download\.openstreetmap\.fr/extracts/(\w.+|)$`),
				regexp.MustCompile(`https://download\.openstreetmap\.fr/polygons/(\w.+|)$`),
				regexp.MustCompile(`https://download.openstreetmap.fr/cgi-bin/^(.*)$`),
				regexp.MustCompile(`https://download.openstreetmap.fr/replication/^(.*|)$`),
			},
			FormatDefinition: formatDefinitions{
				formatOsmPbf:  {ID: formatOsmPbf, Loc: "-latest.osm.pbf"},
				"osm.pbf.md5": {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
				formatPoly:    {ID: formatPoly, Loc: ".poly", BasePath: "../polygons/"},
				formatState:   {ID: formatState, Loc: ".state.txt"},
			},
			Timeout: &t,
		},
	}
)

// Collector represent geofabrik's scrapper
func (o *OpenstreetmapFR) Collector(options ...interface{}) *colly.Collector {
	c := o.Scrapper.Collector(options)
	c.OnHTML("a", func(e *colly.HTMLElement) {
		o.parse(e, c)
	})

	return c
}

func openstreetmapFRGetParent(href string) (parent string, parents []string) {
	var p string

	pp := strings.Split(href, "/")
	if len(pp) > 4 {
		p = pp[len(pp)-2] // Get x in this kind of url http(s)://1/2/.../x/
	} else {
		p = ""
	}

	if strings.EqualFold(p, "extracts") { // should I try == or a switch?
		p = ""
	} else if strings.EqualFold(p, "polygons") {
		p = ""
	}

	return p, pp
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

			if err := o.Config.mergeElement(&element); err != nil {
				catch(fmt.Errorf("can't merge element, %v", err))
			}

			if gparent != "" {
				o.makeParents(gparent, gparents[:len(gparents)-1])
			}
		}
	}
}

func (o *OpenstreetmapFR) parseHref(href string) { //nolint:gocyclo // TODO: Fix cyclo complexity
	if *fVerbose && !*fQuiet && !*fProgress {
		log.Println("Parsing:", href)
	}

	if !strings.Contains(href, "?") && !strings.Contains(href, "-latest") && href[0] != '/' {
		parent, parents := openstreetmapFRGetParent(href)
		if !o.Config.Exist(parent) {
			o.makeParents(parent, parents)
		}

		valsplit := strings.Split(parents[len(parents)-1], ".")
		if valsplit[0] != "" {
			if *fVerbose && !*fQuiet && !*fProgress {
				log.Println("Parsing", valsplit[0])
			}

			extension := strings.Join(valsplit[1:], ".")
			if strings.Contains(extension, "state.txt") { // I'm shure it can be refactorized
				extension = formatState
			}

			if *fVerbose && !*fQuiet && !*fProgress {
				log.Println("Add", extension, "format")
			}

			element := Element{
				Parent: parent,
				Name:   valsplit[0],
				ID:     valsplit[0],
				Meta:   false,
			}
			if !o.Config.Exist(valsplit[0]) {
				element.Formats = append(element.Formats, extension)

				if err := o.Config.mergeElement(&element); err != nil {
					catch(fmt.Errorf("can't merge element, %v", err))
				}
			} else {
				if *fVerbose && !*fQuiet && !*fProgress {
					log.Println(valsplit[0], "already exist")
					log.Println("Merging formats")
				}
				o.Config.AddExtension(valsplit[0], extension)
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
			} else if *fVerbose && !*fProgress && !*fQuiet {
				log.Printf("URL: %v is not matching URLFilters\n", href)
			}
		}
	} else {
		o.parseHref(href)
	}
}
