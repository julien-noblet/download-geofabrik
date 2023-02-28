package openstreetmapfr

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper"
)

// OpenstreetmapFR Scrapper.
type OpenstreetmapFR struct {
	*scrapper.Scrapper
}

const (
	defaultTimeout = time.Second * 30
)

func GetDefault() *OpenstreetmapFR {
	timeout := defaultTimeout

	return &OpenstreetmapFR{
		Scrapper: &scrapper.Scrapper{ //nolint:exhaustruct // I'm lazy
			PB:             149, //nolint:gomnd // There is 149 element right now
			Async:          true,
			Parallelism:    20, //nolint:gomnd // 20 threads for scrapping
			MaxDepth:       0,
			AllowedDomains: []string{`download.openstreetmap.fr`},
			BaseURL:        `https://download.openstreetmap.fr/extracts`,
			StartURL:       `https://download.openstreetmap.fr/`,
			URLFilters: []*regexp.Regexp{
				regexp.MustCompile(`https://download\.openstreetmap\.fr/$`),
				regexp.MustCompile(`https://download\.openstreetmap\.fr/extracts/(\w.+|)$`), //nolint:gocritic,lll // I don't know why gocrtic is complaining about this
				regexp.MustCompile(`https://download\.openstreetmap\.fr/polygons/(\w.+|)$`), //nolint:gocritic,lll // I don't know why gocrtic is complaining about this
				regexp.MustCompile(`https://download.openstreetmap.fr/cgi-bin/^(.*)$`),      //nolint:gocritic // ^ is intentional
				regexp.MustCompile(`https://download.openstreetmap.fr/replication/^(.*|)$`), //nolint:gocritic // ^ is intentional
			},
			FormatDefinition: formats.FormatDefinitions{
				formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: "-latest.osm.pbf"},
				"osm.pbf.md5":        {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
				formats.FormatPoly:   {ID: formats.FormatPoly, Loc: ".poly", BasePath: "../polygons/"},
				formats.FormatState:  {ID: formats.FormatState, Loc: ".state.txt"},
			},
			Timeout: &timeout,
		},
	}
}

// Collector represent geofabrik's scrapper.
func (o *OpenstreetmapFR) Collector() *colly.Collector {
	c := o.Scrapper.Collector()
	c.OnHTML("a", func(e *colly.HTMLElement) {
		o.Parse(e, c)
	})

	return c
}

func GetParent(href string) (parent string, parents []string) { //nolint:nonamedreturns // I prefer to name those ones
	parents = strings.Split(href, "/")
	if len(parents) > 4 { //nolint:gomnd // need at least 4 elments to have a parent :
		// http://1/2/.../x
		// 1    2 3 4 ...
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

//nolint:cyclop // TODO: Refactor?
func (o *OpenstreetmapFR) MakeParents(parent string, gparents []string) {
	if parent != "" { //nolint:nestif // TODO : Refactor?
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
			myElement := element.Element{ //nolint:exhaustruct // I'm lazy
				Parent: gparent,
				Name:   parent,
				ID:     parent,
				Meta:   true,
			}

			if err := o.Config.MergeElement(&myElement); err != nil {
				log.WithError(err).Errorf("can't merge %s", myElement.Name)
			}

			if gparent != "" {
				o.MakeParents(gparent, gparents[:len(gparents)-1])
			}
		}
	}
}

func (o *OpenstreetmapFR) ParseHref(href string) {
	log.Debugf("Parsing: %s", href)

	if !strings.Contains(href, "?") && !strings.Contains(href, "-latest") && href[0] != '/' { //nolint:nestif // TODO : Refactor?
		parent, parents := GetParent(href)
		if !o.Config.Exist(parent) {
			o.MakeParents(parent, parents)
		}

		valsplit := strings.Split(parents[len(parents)-1], ".")
		if valsplit[0] != "" {
			log.Debugf("Parsing %s", valsplit[0])

			extension := strings.Join(valsplit[1:], ".")
			if strings.Contains(extension, "state.txt") { // I'm shure it can be refactorized
				extension = formats.FormatState
			}

			log.Debugf("Add %s format", extension)

			myElement := element.Element{ //nolint:exhaustruct // I'm lazy
				Parent: parent,
				Name:   valsplit[0],
				ID:     valsplit[0],
				Meta:   false,
			}
			if !o.Config.Exist(valsplit[0]) {
				myElement.Formats = append(myElement.Formats, extension)

				if err := o.Config.MergeElement(&myElement); err != nil {
					log.WithError(err).Errorf("can't merge %s", myElement.Name)
				}
			} else {
				log.Debugf("%s already exist, Merging formats", valsplit[0])

				o.Config.AddExtension(valsplit[0], extension)
			}
		}
	}
}

func (o *OpenstreetmapFR) Parse(e *colly.HTMLElement, c *colly.Collector) {
	href := e.Request.AbsoluteURL(e.Attr("href"))
	if href[len(href)-1] == '/' { //nolint:nestif // TODO : Refactor?
		log.Debugf("Next: %s", href)

		if err := c.Visit(href); err != nil && !errors.Is(err, colly.ErrAlreadyVisited) {
			if !errors.Is(err, colly.ErrNoURLFiltersMatch) {
				log.WithError(err).Error("can't get url")
			} else {
				log.Debugf("URL: %v is not matching URLFilters\n", href)
			}
		}
	} else {
		o.ParseHref(href)
	}
}
