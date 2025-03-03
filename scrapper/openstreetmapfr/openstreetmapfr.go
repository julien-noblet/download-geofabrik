package openstreetmapfr

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper"
)

// OpenstreetmapFR Scrapper.
type OpenstreetmapFR struct {
	*scrapper.Scrapper
}

const (
	defaultTimeout      = time.Second * 30
	passList            = "HEADER"
	nbElmt              = 1196            // Number of elements in openstreetmap.fr
	parallelism         = 20              // use 20 routines to scrape openstreetmap.fr
	randomDelay         = time.Second * 5 // Random delay between 0 and 5 seconds
	minParentListLength = 4
)

var exceptionList = map[string]struct{}{
	"central":       {},
	"central-east":  {},
	"central-north": {},
	"central-south": {},
	"central-west":  {},
	"central_east":  {},
	"central_north": {},
	"central_south": {},
	"central_west":  {},
	"coastral":      {},
	"east":          {},
	"east_central":  {},
	"east-central":  {},
	"eastern":       {},
	"lake":          {},
	"north":         {},
	"north_central": {},
	"north-central": {},
	"north-east":    {},
	"north-eastern": {},
	"north-west":    {},
	"north-western": {},
	"north_east":    {},
	"north_eastern": {},
	"north_west":    {},
	"north_western": {},
	"northeast":     {},
	"northern":      {},
	"northwest":     {},
	"south":         {},
	"south_central": {},
	"south-central": {},
	"south-east":    {},
	"south-south":   {},
	"south-west":    {},
	"south_east":    {},
	"south_south":   {},
	"south_west":    {},
	"southeast":     {},
	"southern":      {},
	"southwest":     {},
	"west":          {},
	"west_central":  {},
	"west-central":  {},
	"western":       {},
	"france_taaf":   {},
	"sevastopol":    {},
	"la_rioja":      {},
	"jura":          {},
	"santa_cruz":    {},
}

// GetDefault returns a default instance of OpenstreetmapFR.
func GetDefault() *OpenstreetmapFR {
	timeout := defaultTimeout

	return &OpenstreetmapFR{
		Scrapper: &scrapper.Scrapper{
			PB:             nbElmt,
			Async:          true,
			Parallelism:    parallelism,
			MaxDepth:       0,
			AllowedDomains: []string{`download.openstreetmap.fr`},
			BaseURL:        `https://download.openstreetmap.fr/extracts`,
			StartURL:       `https://download.openstreetmap.fr/`,
			URLFilters: []*regexp.Regexp{
				regexp.MustCompile(`https://download\.openstreetmap\.fr/$`),
				regexp.MustCompile(`https://download\.openstreetmap\.fr/extracts/(\w.+|)$`), //nolint:gocritic // This is a valid regexp
				regexp.MustCompile(`https://download\.openstreetmap\.fr/polygons/(\w.+|)$`), //nolint:gocritic // This is a valid regexp
				regexp.MustCompile(`https://download.openstreetmap.fr/cgi-bin/^(.*)$`),      //nolint:gocritic // This is a valid regexp
				regexp.MustCompile(`https://download.openstreetmap.fr/replication/^(.*|)$`), //nolint:gocritic // This is a valid regexp
			},
			FormatDefinition: formats.FormatDefinitions{
				"osm.pbf.md5":        {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5", ToLoc: "", BasePath: "", BaseURL: ""},
				formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: "-latest.osm.pbf", ToLoc: "", BasePath: "", BaseURL: ""},
				formats.FormatPoly:   {ID: formats.FormatPoly, Loc: ".poly", ToLoc: "", BasePath: "../polygons/", BaseURL: ""},
				formats.FormatState:  {ID: formats.FormatState, Loc: ".state.txt", ToLoc: "", BasePath: "", BaseURL: ""},
			},
			Timeout: timeout,
			Config: &config.Config{
				Formats:       formats.FormatDefinitions{},
				Elements:      element.MapElement{},
				ElementsMutex: &sync.RWMutex{},
				BaseURL:       `https://download.openstreetmap.fr/`,
			},
			DomainGlob:  "*",
			RandomDelay: randomDelay,
		},
	}
}

// Collector returns a Colly collector for OpenstreetmapFR.
func (o *OpenstreetmapFR) Collector() *colly.Collector {
	c := o.Scrapper.Collector()
	c.OnHTML("a", func(e *colly.HTMLElement) {
		o.Parse(e, c)
	})

	return c
}

// GetParent returns the parent and the list of parents from a given href.
func GetParent(href string) (parent string, parentList []string) {
	// Remove the last / from the href to avoid empty string in the parent list
	href = strings.TrimSuffix(href, "/")

	parentList = strings.Split(href, "/")
	if len(parentList) > minParentListLength {
		parent = parentList[len(parentList)-2]
	} else {
		parent = ""
	}

	if strings.EqualFold(parent, "extracts") || strings.EqualFold(parent, "polygons") {
		parent = ""
	}

	return parent, parentList
}

// MakeParents creates parent elements recursively.
func (o *OpenstreetmapFR) MakeParents(parent string, gparents []string) {
	if parent == "" {
		return
	}

	gparent := getGparent(gparents)

	if !o.Config.Exist(parent) {
		o.createAndMergeElement(parent, gparent)

		if gparent != "" {
			o.MakeParents(gparent, gparents[:len(gparents)-1])
		}
	}
}

// getGparent returns the grandparent from a list of parents.
func getGparent(gparents []string) string {
	if len(gparents) < minParentListLength {
		return ""
	}

	gparent := gparents[len(gparents)-3]
	if gparent == "http:" || gparent == "openstreetmap.fr" || gparent == "extracts" || gparent == "polygons" {
		return ""
	}

	return gparent
}

// createAndMergeElement creates and merges an element into the configuration.
func (o *OpenstreetmapFR) createAndMergeElement(parent, gparent string) {
	myElement := element.Element{
		Parent:  gparent,
		Name:    parent,
		ID:      parent,
		Formats: []string{},
		File:    "",
		Meta:    true,
	}

	if err := o.Config.MergeElement(&myElement); err != nil {
		log.WithError(err).Errorf("can't merge %s", myElement.Name)
	}
}

// Exceptions returns the exception name if it exists in the exception list.
func Exceptions(name, parent string) string {
	if _, exists := exceptionList[name]; exists {
		return fmt.Sprintf("%v_%v", parent, name)
	}

	return name
}

// ParseHref parses the href and updates the configuration.
func (o *OpenstreetmapFR) ParseHref(href string) {
	log.Debugf("Parsing: %s", href)

	if strings.Contains(href, "?") || strings.Contains(href, "-latest") || href[0] == '/' {
		return
	}

	parent, parents := GetParent(href)
	if !o.Config.Exist(parent) {
		o.MakeParents(parent, parents)
	}

	valsplit := strings.Split(parents[len(parents)-1], ".")
	if valsplit[0] == "" || len(strings.Split(href, "/")) <= minParentListLength {
		return
	}

	if strings.Contains(passList, valsplit[0]) {
		return
	}

	name := Exceptions(valsplit[0], parent)
	log.Debugf("Parsing %s", name)

	extension := strings.Join(valsplit[1:], ".")
	if strings.Contains(extension, "state.txt") {
		extension = formats.FormatState
	}

	log.Debugf("Add %s format", extension)

	file := ""
	if extension != "" {
		file = name
	}

	o.addOrUpdateElement(parent, name, file, extension)
}

// addOrUpdateElement adds or updates an element in the configuration.
func (o *OpenstreetmapFR) addOrUpdateElement(parent, name, file, extension string) {
	myElement := element.Element{
		ID:      name,
		File:    file,
		Name:    name,
		Parent:  parent,
		Formats: []string{},
		Meta:    false,
	}

	if extension == "" {
		myElement.File = ""
		myElement.Meta = true
	}

	if !o.Config.Exist(name) {
		if extension != "" {
			myElement.Formats = append(myElement.Formats, extension)
		}

		if err := o.Config.MergeElement(&myElement); err != nil {
			log.WithError(err).Errorf("can't merge %s", myElement.Name)
		}
	} else {
		log.Debugf("%s already exist, Merging formats", name)

		if extension != "" {
			o.Config.AddExtension(name, extension)
		}
	}
}

// Parse parses the HTML element and visits the URL if it's a directory.
func (o *OpenstreetmapFR) Parse(e *colly.HTMLElement, c *colly.Collector) {
	href := e.Request.AbsoluteURL(e.Attr("href"))
	if isDirectory(href) {
		log.Debugf("Next: %s", href)
		visitURL(c, href)
	} else {
		o.ParseHref(href)
	}
}

// isDirectory checks if the URL is a directory.
func isDirectory(href string) bool {
	return href[len(href)-1] == '/'
}

// visitURL visits the URL and handles errors.
func visitURL(c *colly.Collector, href string) {
	if err := c.Visit(href); err != nil && !errors.Is(err, colly.ErrAlreadyVisited) {
		if !errors.Is(err, colly.ErrNoURLFiltersMatch) {
			log.WithError(err).Error("can't get url")
		} else {
			log.Debugf("URL: %v is not matching URLFilters\n", href)
		}
	}
}
