package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

const geofabrikPb = 401

func geofabrikGetParent(url string) (string, string) {
	r := strings.Split(url, "/")
	if len(r) < 5 {
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

func geofabrikParseSubregion(e *colly.HTMLElement, ext *Ext, c *colly.Collector) {
	e.ForEach("td.subregion", func(_ int, el *colly.HTMLElement) {
		el.ForEach("a", func(_ int, sub *colly.HTMLElement) {
			href := sub.Request.AbsoluteURL(sub.Attr("href"))
			//name := sub.Text
			//log.Println(name, ":", href)
			id, extension := geofabrikFile(href)
			if extension == "html" {
				parent, pp := geofabrikGetParent(href)
				element := Element{
					ID:     id,
					Name:   sub.Text,
					Parent: parent,
					Meta:   true,
				}
				if !ext.Exist(parent) && parent != "" {
					gparent, _ := geofabrikGetParent(pp)
					if *fVerbose && !*fQuiet && !*fProgress {
						log.Println("Create Meta", element.Parent, "parent:", gparent, pp)
					}
					gp := geofabrikMakeParent(element, gparent)
					if gp != nil {
						ext.mergeElement(gp)
					}
				}
				ext.mergeElement(&element)
				if *fVerbose && !*fQuiet && !*fProgress {
					log.Println("Add:", href)
				}
				c.Visit(href)
			}
		})
	})
}

func geofabrikAddExtension(id, format string, ext *Ext) {
	ext.ElementsMutex.RLock()
	element := ext.Elements[id]
	ext.ElementsMutex.RUnlock()
	if !contains(element.Formats, format) {
		if *fVerbose && !*fQuiet && !*fProgress {
			log.Println("Add", format, "to", element.ID)
		}
		element.Formats = append(element.Formats, format)
		ext.mergeElement(&element)
	}
}

func geofabrikParseLi(e *colly.HTMLElement, ext *Ext, c *colly.Collector) {
	e.ForEach("a", func(_ int, el *colly.HTMLElement) {
		_, format := geofabrikFile(el.Attr("href"))
		id, _ := geofabrikFile(el.Request.URL.String())
		switch format {
		case "osm.pbf":
			geofabrikAddExtension(id, format, ext)
			geofabrikAddExtension(id, "kml", ext)   // not checked!
			geofabrikAddExtension(id, "state", ext) // not checked!
		case "osm.pbf.md5":
			geofabrikAddExtension(id, format, ext)
		case "osm.bz2":
			geofabrikAddExtension(id, format, ext)
		case "osm.bz2.md5":
			geofabrikAddExtension(id, format, ext)
		case "poly":
			geofabrikAddExtension(id, format, ext)
		case "shp.zip":
			geofabrikAddExtension(id, format, ext)
		}

	})
}
