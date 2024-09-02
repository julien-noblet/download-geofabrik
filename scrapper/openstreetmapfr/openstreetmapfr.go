package openstreetmapfr

import (
	"errors"
	"fmt"
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
	exeptionList   = "central" +
		"central-east" +
		"central-north" +
		"central-south" +
		"central-west" +
		"central_east" +
		"central_north" +
		"central_south" +
		"central_west" +
		"coastral" +
		"east" +
		"east_central" +
		"east-central" +
		"eastern" +
		"lake" +
		"north" +
		"north_central" +
		"north-central" +
		"north-east" +
		"north-eastern" +
		"north-west" +
		"north-western" +
		"north_east" +
		"north_eastern" +
		"north_west" +
		"north_western" +
		"northeast" +
		"northern" +
		"northwest" +
		"south" +
		"south_central" +
		"south-central" +
		"south-east" +
		"south-south" +
		"south-west" +
		"south_central" +
		"south_east" +
		"south_south" +
		"south_west" +
		"southeast" +
		"southern" +
		"southwest" +
		"west" +
		"west_central" +
		"west-central" +
		"western" +
		"france_taaf" +
		"sevastopol" +
		"la_rioja" +
		"jura"
	passList = "HEADER"
)

func GetDefault() *OpenstreetmapFR {
	timeout := defaultTimeout

	return &OpenstreetmapFR{
		Scrapper: &scrapper.Scrapper{ //nolint:exhaustruct // I'm lazy
			PB:             1195, //nolint:gomnd // There is 1195 element right now
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
				formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: "-latest.osm.pbf", BasePath: "", BaseURL: ""},
				"osm.pbf.md5":        {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5", BasePath: "", BaseURL: ""},
				formats.FormatPoly:   {ID: formats.FormatPoly, Loc: ".poly", BasePath: "../polygons/", BaseURL: ""},
				formats.FormatState:  {ID: formats.FormatState, Loc: ".state.txt", BasePath: "", BaseURL: ""},
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

func exeptions(name, parent string) string {
	if strings.Contains(exeptionList, name) {
		return fmt.Sprintf("%v_%v", parent, name)
	}

	return name
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
			if strings.Contains(passList, valsplit[0]) { // There is a list of element that should be passed
				return
			}

			name := valsplit[0]
			file := name
			name = exeptions(name, parent)
			log.Debugf("Parsing %s", name)

			extension := strings.Join(valsplit[1:], ".")
			if strings.Contains(extension, "state.txt") { // I'm shure it can be refactorized
				extension = formats.FormatState
			}

			log.Debugf("Add %s format", extension)

			myElement := element.Element{ //nolint:exhaustruct // I'm lazy
				Parent: parent,
				Name:   name,
				File:   file,
				ID:     name,
				Meta:   false,
			}
			if !o.Config.Exist(name) {
				myElement.Formats = append(myElement.Formats, extension)

				if err := o.Config.MergeElement(&myElement); err != nil {
					log.WithError(err).Errorf("can't merge %s", myElement.Name)
				}
			} else {
				log.Debugf("%s already exist, Merging formats", name)

				o.Config.AddExtension(name, extension)
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
