package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"

	"gopkg.in/yaml.v2"
)

type config struct {
	BaseURL  string             `yaml:"baseURL"`
	Formats  map[string]format  `yaml:"formats"`
	Elements map[string]Element `yaml:"elements"`
}

type format struct {
	ID  string `yaml:"ext"`
	Loc string `yaml:"loc"`
}

type Element struct {
	ID     string   `yaml:"id"`
	File   string   `yaml:"file,omitempty"`
	Meta   bool     `yaml:"meta,omitempty"`
	Name   string   `yaml:"name"`
	Files  []string `yaml:"files,omitempty"`
	Parent string   `yaml:"parent,omitempty"`
}

type ElementSlice map[string]Element

func (e ElementSlice) Generate() ([]byte, error) {
	var myConfig config
	myConfig.BaseURL = "http://download.geofabrik.de"
	myConfig.Formats = make(map[string]format)
	myConfig.Formats["osm.bz2"] = format{ID: "osm.bz2", Loc: "-latest.osm.bz2"}
	myConfig.Formats["osm.pbf"] = format{ID: "osm.pbf", Loc: "-latest.osm.pbf"}
	myConfig.Formats["poly"] = format{ID: "poly", Loc: ".poly"}
	myConfig.Formats["state"] = format{ID: "state", Loc: ".state"}
	myConfig.Formats["shp.zip"] = format{ID: "shp.zip", Loc: "-latest.shp.zip"}
	myConfig.Elements = e
	return yaml.Marshal(myConfig)
}

type Ext struct {
	*gocrawl.DefaultExtender
	Elements ElementSlice
}

func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	fmt.Printf("Visit: %s\n", ctx.URL())
	var thisElement Element
	downloadMain := doc.Find("div#download-main")
	parent, haveParent := doc.Find("p a").Attr("href")
	if haveParent && strings.Index(parent, "http://www.geofabrik.de/") == -1 {
		parent = parent[0 : len(parent)-5]
		if parent == "index" {
			parent = ""
		} else {
			temp := strings.Split(parent, "/")
			parent = temp[len(temp)-1]
		}
		thisElement.Parent = parent
		for element := range downloadMain.Nodes {
			singleElement := downloadMain.Eq(element)
			name := singleElement.Find("h2").Text()
			thisElement.Name = name
			index := 0
			li := singleElement.Find("li")
			for el := range li.Nodes {
				myel := li.Eq(el)
				linkval, link := myel.Find("a").Attr("href")
				if link {
					switch index {
					case 0: // osm.pbf
						thisElement.Files = append(thisElement.Files, "osm.pbf")
					case 1: // shp.zip
						thisElement.Files = append(thisElement.Files, "shp.zip")
					case 2: // osm.bz2
						thisElement.Files = append(thisElement.Files, "osm.bz2")
					case 3: // poly
						thisElement.Files = append(thisElement.Files, "poly")
						thisElement.ID = linkval[0 : len(linkval)-5]

					case 4: //-updates
						thisElement.Files = append(thisElement.Files, "state")
					}
				}
				index++
			}
		}
		if len(thisElement.Files) == 0 {
			thisElement.Meta = true
		}
		//Execptions!
		// Only Georgia (EU and US)
		if thisElement.ID == "georgia" {
			thisElement.File = "georgia"
			if thisElement.Parent == "europe" {
				thisElement.Name = "Georgia (Europe country)"
				thisElement.ID = "georgia-eu"
			} else {
				thisElement.Name = "Georgia (US State)"
				thisElement.ID = "georgia-us"
			}
		}
		e.Elements[thisElement.ID] = thisElement
	}
	fmt.Println("OK")
	fmt.Printf("Elements: %d\n", len(e.Elements))

	return nil, true
}

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	if isVisited {
		return false
	}
	if strings.Index(ctx.URL().Path, ".html") != -1 {
		return true
		//	} else if strings.Index(ctx.URL().Path, ".pbf") != -1 {
		//		return false
		//	} else if strings.Index(ctx.URL().Path, ".poly") != -1 {
		//		return false
		//	} else if strings.Index(ctx.URL().Path, ".bz2") != -1 {
		//		return false
		//	} else if strings.Index(ctx.URL().Path, ".zip") != -1 {
		//		return false
	} else if ctx.URL().Path[len(ctx.URL().Path)-1:] == "/" {
		return true
		//	} else if ctx.URL().Path[len(ctx.URL().Path)-8:] == "-updates" {
		//		return false
		//	} else {
		//		return false
	}
	return false
}

func main() {
	ext := &Ext{&gocrawl.DefaultExtender{}, make(map[string]Element)}
	// Set custom options
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = time.Duration(1000) * time.Microsecond
	opts.LogFlags = gocrawl.LogError
	opts.SameHostOnly = true //false
	//opts.MaxVisits = 5

	c := gocrawl.NewCrawlerWithOptions(opts)
	c.Run("http://download.geofabrik.de/")
	// Done so wee need to generate yml :)
	out, _ := ext.Elements.Generate()
	//fmt.Printf("%q\n", out)

	filename, _ := filepath.Abs("geofabrik.yml")
	err := ioutil.WriteFile(filename, out, 0644)
	if err != nil {
		log.Fatalln(" File error: %v ", err)
		os.Exit(1)
	}

}
