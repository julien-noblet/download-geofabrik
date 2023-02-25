package bbbike

import (
	"errors"
	"regexp"

	"github.com/apex/log"
	"github.com/gocolly/colly"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper"
)

// Bbbike Scrapper
type Bbbike struct {
	*scrapper.Scrapper
}

func GetDefault() *Bbbike {
	return &Bbbike{
		Scrapper: &scrapper.Scrapper{
			PB:             236, //nolint:gomnd // there is 236 element at this time
			Async:          true,
			Parallelism:    20, //nolint:gomnd // Use 20 threads for scraping
			MaxDepth:       0,
			AllowedDomains: []string{`download.bbbike.org`},
			BaseURL:        `https://download.bbbike.org/osm/bbbike`,
			StartURL:       `https://download.bbbike.org/osm/bbbike/`,
			URLFilters: []*regexp.Regexp{
				regexp.MustCompile(`https://download\.bbbike\.org/osm/bbbike/[A-Z].+$`),
				regexp.MustCompile(`https://download\.bbbike\.org/osm/bbbike/$`),
			},
			FormatDefinition: formats.FormatDefinitions{
				formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: ".osm.pbf"},
				formats.FormatShpZip: {ID: formats.FormatShpZip, Loc: ".osm.shp.zip"},
				formats.FormatOsmGz:  {ID: formats.FormatOsmGz, Loc: ".osm.gz"},
				formats.FormatPoly:   {ID: formats.FormatPoly, Loc: ".poly"},
			},
		},
	}
}

// Collector represent geofabrik's scrapper
func (b *Bbbike) Collector(options ...interface{}) *colly.Collector {
	c := b.Scrapper.Collector(options)
	c.OnHTML("div.list tbody", func(e *colly.HTMLElement) {
		b.parseList(e, c)
	})
	c.OnHTML("#sidebar", func(e *colly.HTMLElement) {
		b.parseSidebar(e, c)
	})

	return c
}

func (b *Bbbike) parseList(e *colly.HTMLElement, c *colly.Collector) {
	e.ForEach("a", func(_ int, el *colly.HTMLElement) {
		href := el.Request.AbsoluteURL(el.Attr("href"))
		log.Debugf("Parse: %s", href)

		if err := c.Visit(href); err != nil && !errors.Is(err, colly.ErrNoURLFiltersMatch) { // Not matching
			log.WithError(err).Error("can't get url")
		}
	})
}

func bbbikeGetName(h3 string) string {
	ret := h3[17:] // remove "OSM extracts for "

	return ret
}

func (b *Bbbike) parseSidebar(e *colly.HTMLElement, c *colly.Collector) { //nolint:unparam,lll // *colly.Collector is passed as param but unused in this case
	name := bbbikeGetName(e.ChildText("h3"))
	el := element.Element{
		ID:     name,
		Name:   name,
		File:   name + "/" + name,
		Parent: "",
		Formats: element.Formats{
			formats.FormatOsmPbf,
			formats.FormatOsmGz,
			formats.FormatShpZip,
		},
		Meta: false,
	}

	log.Debugf("Add %s", name)

	if err := b.Config.MergeElement(&el); err != nil {
		log.WithError(err).Errorf("can't merge %s", el.Name)
	}
}
