package osmtoday

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/apex/log"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper"
)

// Osmtoday Scrapper.
type Osmtoday struct {
	*scrapper.Scrapper
}

func GetDefault() *Osmtoday {
	return &Osmtoday{
		Scrapper: &scrapper.Scrapper{ //nolint:exhaustruct // I'm lazy
			PB:             412, //nolint:gomnd // there is 412 items
			Async:          true,
			Parallelism:    20, //nolint:gomnd // use 20 threads for scrapping
			MaxDepth:       0,
			AllowedDomains: []string{`osmtoday.com`},
			BaseURL:        `https://osmtoday.com`,
			StartURL:       `https://osmtoday.com/`,
			URLFilters: []*regexp.Regexp{
				regexp.MustCompile(`https://osmtoday\.com/.+\.html$`),
				regexp.MustCompile(`https://osmtoday\.com/$`),
			},
			FormatDefinition: formats.FormatDefinitions{
				formats.FormatOsmPbf:  {ID: formats.FormatOsmPbf, Loc: ".pbf", BasePath: "", BaseURL: ""},
				"md5":                 {ID: "md5", Loc: ".md5", BasePath: "", BaseURL: ""},
				formats.FormatPoly:    {ID: formats.FormatPoly, Loc: ".poly", BasePath: "", BaseURL: ""},
				formats.FormatGeoJSON: {ID: formats.FormatKml, Loc: ".geojson", BasePath: "", BaseURL: ""},
			},
		},
	}
}

// Collector represent osmtoday's scrapper.
func (g *Osmtoday) Collector() *colly.Collector {
	myCollector := g.Scrapper.Collector()
	myCollector.OnHTML(".row", func(e *colly.HTMLElement) {
		g.ParseLi(e, myCollector)
	})
	myCollector.OnHTML("table", func(e *colly.HTMLElement) {
		g.ParseSubregion(e, myCollector)
	})

	return myCollector
}

func (g *Osmtoday) Exceptions(e element.Element) element.Element {
	// Exception
	// la_rioja is in argentina and spain
	if e.ID == "la_rioja" {
		switch e.Parent {
		case "argentina":
			e.ID = "la_rioja-argentina"
			e.File = "la_rioja"
		case "spain":
			e.ID = "la_rioja-spain"
			e.File = "la_rioja"
		}
	}
	// Guyan is in france
	if e.ID == "guyane" {
		switch e.Parent {
		case "france":
			e.ID = "guyane-france"
			e.File = "guyane"
		case "south-america":
			e.ID = "guyane-south-america"
			e.File = "guyane"
		}
	}
	//    ⨯ can't merge Sevastopol    error=can't merge : Parent mismatch southern_federal_district != ukraine (sevastopol)
	if e.ID == "sevastopol" {
		switch e.Parent {
		case "ukraine":
			e.ID = "sevastopol-ukraine"
			e.File = "sevastopol"
		case "russia":
			e.ID = "sevastopol-russia"
			e.File = "sevastopol"
		}
	}
	//    ⨯ can't merge Limburg       error=can't merge : Parent mismatch netherlands != flanders (limburg)
	if e.ID == "limburg" {
		switch e.Parent {
		case "netherlands":
			e.ID = "limburg-netherlands"
			e.File = "limburg"
		case "flanders":
			e.ID = "limburg-flanders"
			e.File = "limburg"
		}
	}
	//    ⨯ can't merge Cordoba       error=can't merge : Parent mismatch argentina != andalucia (cordoba)
	if e.ID == "cordoba" {
		switch e.Parent {
		case "argentina":
			e.ID = "cordoba-argentina"
			e.File = "cordoba"
		case "andalucia":
			e.ID = "cordoba-andalucia"
			e.File = "cordoba"
		}
	}
	//    ⨯ can't merge Georgia       error=can't merge : Parent mismatch asia != us (georgia)
	if e.ID == "georgia" {
		switch e.Parent {
		case "asia":
			e.ID = "georgia-asia"
			e.File = "georgia"
		case "us":
			e.ID = "georgia-us"
			e.File = "georgia"
		}
	}
	return e
}

//nolint:cyclop // TODO : Refactoring?
func (g *Osmtoday) ParseSubregion(e *colly.HTMLElement, myCollector *colly.Collector) {
	e.ForEach("td", func(_ int, el *colly.HTMLElement) {
		el.ForEach("a", func(_ int, sub *colly.HTMLElement) {
			href := sub.Request.AbsoluteURL(sub.Attr("href"))
			myID, extension := scrapper.FileExt(href)
			if myID == "" {
				// TODO : Move to Debug
				fmt.Println("myID is empty, href:", href)

				return
			}

			var file string

			if extension == "html" { //nolint:nestif // TODO : Refactor?
				parent, parentPath := scrapper.GetParent(href)

				myElement := element.Element{ //nolint:exhaustruct // I'm lazy
					ID:     myID,
					Name:   sub.Text,
					Parent: parent,
					Meta:   true,
				}
				myElement = g.Exceptions(myElement)
				if file != "" {
					myElement.File = file
				}
				if !g.Config.Exist(parent) && parent != "" { // Case of parent should exist not already in Slice
					gparent, _ := scrapper.GetParent(parentPath)
					log.Debugf("Create Meta %s parent: %s %v", myElement.Parent, gparent, parentPath)

					if gp := element.MakeParent(&myElement, gparent); gp != nil {
						if err := g.Config.MergeElement(gp); err != nil {
							log.WithError(err).Errorf("can't merge %s", myElement.Name)
						}
					}
				}
				if err := g.Config.MergeElement(&myElement); err != nil {
					log.WithError(err).Errorf("can't merge %s", myElement.Name)
				}
				log.Debugf("Add: %s", href)

				if err := myCollector.Visit(href); err != nil && !errors.Is(err, colly.ErrAlreadyVisited) {
					log.WithError(err).Error("can't get url")
				}
			} else {
				parent, _ := scrapper.GetParent(href)

				myElement := element.Element{ //nolint:exhaustruct // I'm lazy
					ID:     myID,
					Name:   sub.Text,
					Parent: parent,
					Meta:   true,
				}
				myElement = g.Exceptions(myElement)
				if file != "" {
					myElement.File = file
				}
				g.ParseFormat(myElement.ID, extension)
			}

		})
	})
}

// ParseFormat Add Extension to ID
func (g *Osmtoday) ParseFormat(id, format string) {
	g.Scrapper.ParseFormatService(id, format, &g.Scrapper.FormatDefinition)

	if format == "pbf" {
		g.Config.AddExtension(id, "md5") // not checked!
	}
}

func (g *Osmtoday) ParseLi(e *colly.HTMLElement, _ *colly.Collector) {
	e.ForEach("a", func(_ int, element *colly.HTMLElement) {
		_, format := scrapper.FileExt(element.Attr("href"))
		myID, _ := scrapper.FileExt(element.Request.URL.String()) // id can't be extracted from href

		g.ParseFormat(myID, format)
	})
}
