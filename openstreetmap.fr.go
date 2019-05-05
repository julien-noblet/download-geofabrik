package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

const (
	openstreetmapFRPb = 144
)

func openstreetmapFRGetParent(href string) (string, []string) {
	parents := strings.Split(href, "/")
	var parent string
	if len(parents) > 4 {
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

func openstreetmapFRParseHref(href string, config *Config) {
	if *fVerbose && !*fQuiet && !*fProgress {
		log.Println("Parsing:", href)
	}
	//	if !strings.Contains(href, "?") && !strings.Contains(href, "-latest") && href[0] != '/' && !strings.EqualFold(href, "cgi-bin/") {
	if !strings.Contains(href, "?") && !strings.Contains(href, "-latest") && href[0] != '/' {
		parent, parents := openstreetmapFRGetParent(href)
		valsplit := strings.Split(parents[len(parents)-1], ".")
		if valsplit[0] != "" {
			if *fVerbose && !*fQuiet && !*fProgress {
				log.Println("Parsing", valsplit[0])
				// Output:
				// parsing valsplit[0]
			}
			extention := strings.Join(valsplit[1:], ".")
			if strings.Contains(extention, "state.txt") { // I'm shure it can be refactorized
				extention = "state"
			}
			if *fVerbose && !*fQuiet && !*fProgress {
				log.Println("Add", extention, "format")
			}
			element := Element{
				Parent: parent,
				Name:   valsplit[0],
				ID:     valsplit[0],
				Meta:   false,
			}
			if !config.Exist(valsplit[0]) {
				element.Formats = append(element.Formats, extention)
				err := config.mergeElement(&element)
				if err != nil {
					catch(fmt.Errorf("can't merge element, %v", err))
					// Panic
				}
			} else {
				e, _ := config.GetElement(valsplit[0])
				fmt.Println(valsplit[0], "Exists, ID: ", e)
				if *fVerbose && !*fQuiet && !*fProgress {
					log.Println(valsplit[0], "already exist")
					log.Println("Merging formats")
				}
				config.AddExtension(valsplit[0], extention)
			}
		}
	}
}

func openstreetmapFRParse(e *colly.HTMLElement, config *Config, c *colly.Collector) {
	href := e.Request.AbsoluteURL(e.Attr("href"))
	if href[len(href)-1] == '/' {
		if *fVerbose && !*fQuiet && !*fProgress {
			log.Println("Next:", href)
		}
		if err := c.Visit(href); err != nil && err != colly.ErrAlreadyVisited {
			catch(err)
		}
	} else {
		openstreetmapFRParseHref(href, config)
	}

}
