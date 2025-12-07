package geo2day

import (
	"errors"
	"fmt"
	"log/slog"
	"regexp"
	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/internal/scrapper"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
)

// Constants for magic numbers and URLs.
const (
	progressBarCount = 1003 // number of elements
	parallelism      = 20   // number of parallel downloads
	baseURL          = "https://geo2day.com"
	startURL         = baseURL + "/"
)

// Geo2day Scrapper.
type Geo2day struct {
	*scrapper.Scrapper
}

// GetDefault returns the default configuration for Osmtoday scrapper.
func GetDefault() *Geo2day {
	urlFilters := []*regexp.Regexp{
		regexp.MustCompile(`https://geo2day\.com/.+\.html$`),
		regexp.MustCompile(`https://geo2day\.com/$`),
	}

	formatDefinition := formats.FormatDefinitions{
		"osm.pbf.md5":         {ID: "osm.pbf.md5", Loc: ".md5"},
		formats.FormatGeoJSON: {ID: formats.FormatGeoJSON, Loc: ".geojson"},
		formats.FormatOsmPbf:  {ID: formats.FormatOsmPbf, Loc: ".pbf"},
		formats.FormatPoly:    {ID: formats.FormatPoly, Loc: ".poly"},
	}

	return &Geo2day{
		Scrapper: &scrapper.Scrapper{
			PB:               progressBarCount,
			Async:            true,
			Parallelism:      parallelism,
			MaxDepth:         0,
			AllowedDomains:   []string{`geo2day.com`},
			BaseURL:          baseURL,
			StartURL:         startURL,
			URLFilters:       urlFilters,
			FormatDefinition: formatDefinition,
			Config: &config.Config{
				Formats:       formats.FormatDefinitions{},
				Elements:      element.MapElement{},
				ElementsMutex: &sync.RWMutex{},
				BaseURL:       "",
			},
		},
	}
}

// Collector represents Osmtoday's scrapper.
func (g *Geo2day) Collector() *colly.Collector {
	myCollector := g.Scrapper.Collector()
	myCollector.OnHTML(".row", func(e *colly.HTMLElement) {
		g.ParseLi(e, myCollector)
	})
	myCollector.OnHTML("table", func(e *colly.HTMLElement) {
		g.ParseSubregion(e, myCollector)
	})

	return myCollector
}

// Exceptions handles special cases for certain IDs.
func (g *Geo2day) Exceptions(myElement *element.Element) *element.Element {
	exceptions := []struct {
		ID     string
		Parent string
	}{
		{"la_rioja", "argentina"},
		{"la_rioja", "spain"},
		{"guyane", "france"},
		{"guyane", "south-america"},
		{"sevastopol", "ukraine"},
		{"sevastopol", "russia"},
		{"limburg", "netherlands"},
		{"limburg", "flanders"},
		{"cordoba", "argentina"},
		{"cordoba", "andalucia"},
		{"georgia", "asia"},
		{"georgia", "us"},
	}

	for _, exception := range exceptions {
		if myElement.ID == exception.ID && myElement.Parent == exception.Parent {
			myElement.ID = fmt.Sprintf("%s-%s", myElement.ID, myElement.Parent)
		}
	}

	return myElement
}

// ParseSubregion parses the subregion information from the HTML.
func (g *Geo2day) ParseSubregion(e *colly.HTMLElement, myCollector *colly.Collector) {
	e.ForEach("td", func(_ int, el *colly.HTMLElement) {
		el.ForEach("a", func(_ int, sub *colly.HTMLElement) {
			href := sub.Request.AbsoluteURL(sub.Attr("href"))

			myID, extension := scrapper.FileExt(href)
			if myID == "" {
				slog.Debug("myID is empty", "href", href)

				return
			}

			if extension == "html" {
				g.handleHTMLExtension(sub, href, myID, myCollector)
			} else {
				parent, _ := scrapper.GetParent(href)

				myElement := element.Element{
					ID:     myID,
					Name:   sub.Text,
					Parent: parent,
					Meta:   true,
				}
				myElement = *g.Exceptions(&myElement)

				if err := g.Config.MergeElement(&myElement); err != nil {
					slog.Error("Can't merge", "name", myElement.Name, "error", err)
				}

				g.ParseFormat(myElement.ID, extension)
			}
		})
	})
}

// handleHTMLExtension handles the HTML extension case.
func (g *Geo2day) handleHTMLExtension(sub *colly.HTMLElement, href, myID string, myCollector *colly.Collector) {
	parent, parentPath := scrapper.GetParent(href)

	myElement := element.Element{
		ID:     myID,
		Name:   sub.Text,
		Parent: parent,
		Meta:   true,
	}

	myElement = *g.Exceptions(&myElement)

	if !g.Config.Exist(parent) && parent != "" {
		gparent, _ := scrapper.GetParent(parentPath)
		slog.Debug("Create Meta", "parent", myElement.Parent, "gparent", gparent, "path", parentPath)

		if gp := element.CreateParentElement(&myElement, gparent); gp != nil {
			if err := g.Config.MergeElement(gp); err != nil {
				slog.Error("Can't merge", "name", myElement.Name, "error", err)
			}
		}
	}

	if err := g.Config.MergeElement(&myElement); err != nil {
		slog.Error("Can't merge", "name", myElement.Name, "error", err)
	}

	slog.Debug("Add", "href", href)

	if err := myCollector.Visit(href); err != nil && !errors.Is(err, &colly.AlreadyVisitedError{}) {
		slog.Error("Can't get url", "error", err)
	}
}

// ParseFormat adds extensions to the ID.
func (g *Geo2day) ParseFormat(id, format string) {
	g.ParseFormatService(id, format, &g.FormatDefinition)

	if format == formats.FormatOsmPbf {
		g.Config.AddExtension(id, "osm.pbf.md5")
	}
}

// ParseLi parses the list items from the HTML.
func (g *Geo2day) ParseLi(e *colly.HTMLElement, _ *colly.Collector) {
	e.ForEach("a", func(_ int, element *colly.HTMLElement) {
		_, format := scrapper.FileExt(element.Attr("href"))
		myID, _ := scrapper.FileExt(element.Request.URL.String())

		g.ParseFormat(myID, format)
	})
}
