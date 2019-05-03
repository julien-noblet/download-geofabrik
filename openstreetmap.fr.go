package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	pb "gopkg.in/cheggaaa/pb.v1"
)

const (
	openstreetmapFRPb = 144
)

// Example:
func openstreetmapFRGetNext(href string, callback func(interface{})) {
	if *fVerbose && !*fQuiet && !*fProgress {
		log.Println("Next:", href)
	}
	callback(href)
}

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

func openstreetmapFRParseHref(href string, ext *Ext) {
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
			}
			ext.ElementsMutex.RLock()
			eq := strings.EqualFold(ext.Elements[valsplit[0]].ID, valsplit[0])
			ext.ElementsMutex.RUnlock()
			if !eq {
				element.Formats = append(element.Formats, extention)
				err := ext.mergeElement(&element)
				if err != nil {
					catch(fmt.Errorf("can't merge element, %v", err))
					// Panic
				}
			} else {
				if *fVerbose && !*fQuiet && !*fProgress {
					log.Println(valsplit[0], "already exist")
					log.Println("Merging formats")
				}
				ext.ElementsMutex.RLock()
				et := ext.Elements[valsplit[0]]
				ext.ElementsMutex.RUnlock()
				if len(et.Formats) == 0 {
					et.Meta = true
				} else {
					et.Meta = false
				}
				et.Formats = append(et.Formats, extention)
				ext.ElementsMutex.Lock()
				ext.Elements[valsplit[0]] = et
				ext.ElementsMutex.Unlock()
			}
		}
	}
}

func openstreetmapFRParse(e *colly.HTMLElement, ext *Ext, bar *pb.ProgressBar, callback func(interface{})) {
	href := e.Request.AbsoluteURL(e.Attr("href"))
	if href[len(href)-1] == '/' {
		openstreetmapFRGetNext(href, callback)
	} else {
		openstreetmapFRParseHref(href, ext)
	}

}
