package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

const (
	geofabrikPb          = 401
	geofabrikAsync       = true
	geofabrikDomainGlob  = "*"
	geofabrikParallelism = 20
	geofabrikRandomDelay = 5 * time.Second
	geofabrikMaxDepth    = 0 // infinite
)

var (
	geofabrikAllowedDomains = []string{`download.geofabrik.de`}
	geofabrikURLFilters     = []*regexp.Regexp{
		regexp.MustCompile(`https://download\.geofabrik\.de/.+\.html$`),
		regexp.MustCompile(`https://download\.geofabrik\.de/$`),
	}
	geofabrikConfig = &Config{
		BaseURL: `https://download.geofabrik.de`,
		Formats: map[string]format{
			//geofabrik.Formats["osh.pbf"] = format{ID: "osh.pbf", Loc: ".osh.pbf"}
			//geofabrik.Formats["osh.pbf.md5"] = format{ID: "osh.pbf.md5", Loc: ".osh.pbf.md5"}
			"osm.bz2":     {ID: "osm.bz2", Loc: "-latest.osm.bz2"},
			"osm.bz2.md5": {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
			"osm.pbf":     {ID: "osm.pbf", Loc: "-latest.osm.pbf"},
			"osm.pbf.md5": {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
			"poly":        {ID: "poly", Loc: ".poly"},
			"kml":         {ID: "kml", Loc: ".kml"},
			"state":       {ID: "state", Loc: "-updates/state.txt"},
			"shp.zip":     {ID: "shp.zip", Loc: "-latest-free.shp.zip"},
		},
		Elements:      ElementSlice{},
		ElementsMutex: &sync.RWMutex{},
	}
	geofabrikLimit = &colly.LimitRule{
		DomainGlob:  geofabrikDomainGlob,
		Parallelism: geofabrikParallelism,
		RandomDelay: geofabrikRandomDelay,
	}
)

func geofabrikGetColly() *colly.Collector {
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains(geofabrikAllowedDomains...),
		colly.URLFilters(geofabrikURLFilters...),
		colly.Async(geofabrikAsync),
		colly.MaxDepth(geofabrikMaxDepth),
	)
	catch(c.Limit(geofabrikLimit))
	c.OnError(func(r *colly.Response, err error) {
		catch(fmt.Errorf("request URL: %v failed with response: %v\nerror: %v", r.Request.URL, r, err.Error()))
	})
	return c
}

func geofabrikGetParent(url string) (string, string) {
	r := strings.Split(url, "/")
	if len(r) < 5 { // <4 should be impossible
		return "", strings.Join(r[:len(r)-1], "/")
	}
	return r[len(r)-2], strings.Join(r[:len(r)-1], "/")
}

func geofabrikFile(url string) (string, string) {
	urls := strings.Split(url, "/")
	file := urls[len(urls)-1]
	f2 := strings.Split(file, ".")
	return f2[0], strings.Join(f2[1:], ".")
}

func geofabrikMakeParent(e Element, gparent string) *Element {
	if e.hasParent() {
		return &Element{
			ID:     e.Parent,
			Name:   e.Parent,
			Parent: gparent,
			Meta:   true,
		}
	}
	return nil
}

func geofabrikParseSubregion(e *colly.HTMLElement, config *Config, c *colly.Collector) {
	e.ForEach("td.subregion", func(_ int, el *colly.HTMLElement) {
		el.ForEach("a", func(_ int, sub *colly.HTMLElement) {
			href := sub.Request.AbsoluteURL(sub.Attr("href"))
			id, extension := geofabrikFile(href)
			if extension == "html" {
				parent, pp := geofabrikGetParent(href)
				if id == "georgia" {
					switch parent {
					case "us":
						id = "georgia-us"
					case "europe":
						id = "georgia-eu"
					}
				}
				element := Element{
					ID:     id,
					Name:   sub.Text,
					Parent: parent,
					Meta:   true,
				}
				if !config.Exist(parent) && parent != "" {
					gparent, _ := geofabrikGetParent(pp)
					if *fVerbose && !*fQuiet && !*fProgress {
						log.Println("Create Meta", element.Parent, "parent:", gparent, pp)
					}
					if gp := geofabrikMakeParent(element, gparent); gp != nil {
						catch(config.mergeElement(gp))
					}
				}
				catch(config.mergeElement(&element))
				if *fVerbose && !*fQuiet && !*fProgress {
					log.Println("Add:", href)
				}
				if err := c.Visit(href); err != nil && err != colly.ErrAlreadyVisited {
					catch(err)
				}
			}
		})
	})
}

func geofabrikParseFormat(id, format string, config *Config) {
	switch format {
	case "osm.pbf":
		config.AddExtension(id, format)
		config.AddExtension(id, "kml")   // not checked!
		config.AddExtension(id, "state") // not checked!
	case "osm.pbf.md5":
		config.AddExtension(id, format)
	case "osm.bz2":
		config.AddExtension(id, format)
	case "osm.bz2.md5":
		config.AddExtension(id, format)
	case "poly":
		config.AddExtension(id, format)
	case "shp.zip":
		config.AddExtension(id, format)
	}
}

func geofabrikParseLi(e *colly.HTMLElement, config *Config, c *colly.Collector) {
	e.ForEach("a", func(_ int, el *colly.HTMLElement) {
		_, format := geofabrikFile(el.Attr("href"))
		id, _ := geofabrikFile(el.Request.URL.String())
		if id == "georgia" { // Exception
			parent, _ := geofabrikGetParent(el.Request.AbsoluteURL(el.Attr("href")))
			switch parent {
			case "us":
				id = "georgia-us"
			case "europe":
				id = "georgia-eu"
			}
		}
		geofabrikParseFormat(id, format, config)
	})
}
