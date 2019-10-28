package openstreetmapfr

import (
	"regexp"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/gocolly/colly"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper"
)

// OpenstreetmapFR Scrapper
type OpenstreetmapFR struct {
	*scrapper.Scrapper
}

const (
	defaultTimeout = time.Second * 30
)

func GetDefault() *OpenstreetmapFR {
	t := defaultTimeout

	return &OpenstreetmapFR{
		Scrapper: &scrapper.Scrapper{
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
			FormatDefinition: formats.FormatDefinitions{
				formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: "-latest.osm.pbf"},
				"osm.pbf.md5":        {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
				formats.FormatPoly:   {ID: formats.FormatPoly, Loc: ".poly", BasePath: "../polygons/"},
				formats.FormatState:  {ID: formats.FormatState, Loc: ".state.txt"},
			},
			Timeout: &t,
		},
	}
}

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
			el := element.Element{
				Parent: gparent,
				Name:   parent,
				ID:     parent,
				Meta:   true,
			}

			if err := o.Config.MergeElement(&el); err != nil {
				log.WithError(err).Errorf("can't merge %s", el.Name)
			}

			if gparent != "" {
				o.makeParents(gparent, gparents[:len(gparents)-1])
			}
		}
	}
}

func (o *OpenstreetmapFR) parseHref(href string) { //nolint:gocyclo // TODO: Fix cyclo complexity
	log.Debugf("Parsing: %s", href)

	if !strings.Contains(href, "?") && !strings.Contains(href, "-latest") && href[0] != '/' {
		parent, parents := openstreetmapFRGetParent(href)
		if !o.Config.Exist(parent) {
			o.makeParents(parent, parents)
		}

		valsplit := strings.Split(parents[len(parents)-1], ".")
		if valsplit[0] != "" {
			log.Debugf("Parsing %s", valsplit[0])

			extension := strings.Join(valsplit[1:], ".")
			if strings.Contains(extension, "state.txt") { // I'm shure it can be refactorized
				extension = formats.FormatState
			}

			log.Debugf("Add %s format", extension)

			el := element.Element{
				Parent: parent,
				Name:   valsplit[0],
				ID:     valsplit[0],
				Meta:   false,
			}
			if !o.Config.Exist(valsplit[0]) {
				el.Formats = append(el.Formats, extension)

				if err := o.Config.MergeElement(&el); err != nil {
					log.WithError(err).Errorf("can't merge %s", el.Name)
				}
			} else {
				log.Debugf("%s already exist, Merging formats", valsplit[0])

				o.Config.AddExtension(valsplit[0], extension)
			}
		}
	}
}

func (o *OpenstreetmapFR) parse(e *colly.HTMLElement, c *colly.Collector) {
	href := e.Request.AbsoluteURL(e.Attr("href"))
	if href[len(href)-1] == '/' {
		log.Debugf("Next: %s", href)

		if err := c.Visit(href); err != nil && err != colly.ErrAlreadyVisited {
			if err != colly.ErrNoURLFiltersMatch {
				log.WithError(err).Error("can't get url")
			} else {
				log.Debugf("URL: %v is not matching URLFilters\n", href)
			}
		}
	} else {
		o.parseHref(href)
	}
}
