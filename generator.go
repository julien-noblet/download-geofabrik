package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/gocolly/colly"
	pb "gopkg.in/cheggaaa/pb.v1"
	yaml "gopkg.in/yaml.v2"
)

var bar *pb.ProgressBar

// ElementSlice contain all Elements
// TODO: It's not a slice but a MAP!!!!
type ElementSlice map[string]Element

// Generate make the slice which contain all Elements
func (e ElementSlice) Generate(myConfig *Config) ([]byte, error) {
	myConfig.Elements = e
	return yaml.Marshal(myConfig)
}

// Generate make the slice which contain all Elements
func (e *Ext) Exist(id string) bool {
	e.ElementsMutex.RLock()
	r := reflect.DeepEqual(e.Elements[id], Element{})
	e.ElementsMutex.RUnlock()
	return r
}

// Ext simple struct for managing ElementSlice and crawler
type Ext struct {
	//*gocrawl.DefaultExtender
	Elements      ElementSlice
	ElementsMutex sync.RWMutex
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}

func (e *Ext) mergeElement(element *Element) error {
	e.ElementsMutex.RLock()
	cE, ok := e.Elements[element.ID]
	e.ElementsMutex.RUnlock()
	if ok {
		if cE.Parent != element.Parent {
			return fmt.Errorf("Cant merge : Parent mismatch")
		}
		e.ElementsMutex.Lock()
		for _, f := range element.Formats {
			if !contains(cE.Formats, f) {
				cE.Formats = append(cE.Formats, f)
			}
		}
		e.ElementsMutex.Unlock()
		if len(cE.Formats) == 0 {
			cE.Meta = true
		} else {
			cE.Meta = false
		}
		e.ElementsMutex.Lock()
		e.Elements[element.ID] = cE
		e.ElementsMutex.Unlock()
	} else {
		e.ElementsMutex.Lock()
		e.Elements[element.ID] = *element
		e.ElementsMutex.Unlock()
	}
	return nil
}

//Generate main function
func Generate(configfile string) {
	switch *fService {
	case "geofabrik":
		//Generate geofabrik.yml
		var geofabrik Config
		geofabrik.BaseURL = "https://download.geofabrik.de"
		geofabrik.Formats = make(map[string]format)
		//TODO: make a function for adding formats
		//geofabrik.Formats["osh.pbf"] = format{ID: "osh.pbf", Loc: ".osh.pbf"}
		//geofabrik.Formats["osh.pbf.md5"] = format{ID: "osh.pbf.md5", Loc: ".osh.pbf.md5"}
		geofabrik.Formats["osm.bz2"] = format{ID: "osm.bz2", Loc: "-latest.osm.bz2"}
		geofabrik.Formats["osm.bz2.md5"] = format{ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"}
		geofabrik.Formats["osm.pbf"] = format{ID: "osm.pbf", Loc: "-latest.osm.pbf"}
		geofabrik.Formats["osm.pbf.md5"] = format{ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"}
		geofabrik.Formats["poly"] = format{ID: "poly", Loc: ".poly"}
		geofabrik.Formats["kml"] = format{ID: "kml", Loc: ".kml"}
		geofabrik.Formats["state"] = format{ID: "state", Loc: "-updates/state.txt"}
		geofabrik.Formats["shp.zip"] = format{ID: "shp.zip", Loc: "-latest-free.shp.zip"}
		ext := Ext{Elements: make(map[string]Element), ElementsMutex: sync.RWMutex{}}
		if *fProgress {
			bar = pb.New(geofabrikPb)
			bar.Start()
		}
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			colly.AllowedDomains("download.geofabrik.de"),
			colly.URLFilters(
				regexp.MustCompile("https://download.geofabrik.de/.+.html"),
				regexp.MustCompile("https://download.geofabrik.de/$"),
			),
			colly.Async(true),
		)
		c.Limit(&colly.LimitRule{
			Parallelism: 2,
			RandomDelay: 5 * time.Second,
			//Delay: 5 * time.Second,
		})
		/*c.WithTransport(&http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          0,
			IdleConnTimeout:       5 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
		})*/

		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err.Error())
		})

		c.OnHTML("#subregions", func(e *colly.HTMLElement) {
			geofabrikParseSubregion(e, &ext, c)
		})
		c.OnHTML("li", func(e *colly.HTMLElement) {
			geofabrikParseLi(e, &ext, c)
		})
		c.OnScraped(func(*colly.Response) {
			if *fProgress {
				bar.Increment()
			}
		})
		c.Visit("https://download.geofabrik.de/")
		c.Wait()
		out, _ := ext.Elements.Generate(&geofabrik)
		filename, _ := filepath.Abs(configfile)
		err := ioutil.WriteFile(filename, out, 0644)
		if err != nil {
			log.Panicln(fmt.Errorf(" File error: %v ", err))
		}
		if !*fQuiet {
			log.Println(configfile, " generated.")
		}

	case "openstreetmap.fr":
		var myConfig Config
		myConfig.BaseURL = "https://download.openstreetmap.fr/extracts"
		myConfig.Formats = make(map[string]format)
		myConfig.Formats["osm.pbf"] = format{ID: "osm.pbf", Loc: "-latest.osm.pbf"}
		myConfig.Formats["poly"] = format{ID: "poly", Loc: ".poly", BasePath: "../polygons/"}
		myConfig.Formats["state"] = format{ID: "state", Loc: ".state.txt"}
		ext := Ext{Elements: make(map[string]Element), ElementsMutex: sync.RWMutex{}}
		if *fProgress {
			bar = pb.New(openstreetmapFRPb)
			bar.Start()
		}
		//GenerateCrawler("https://download.openstreetmap.fr/", configfile, &myConfig)
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			colly.AllowedDomains("download.openstreetmap.fr"),
			colly.URLFilters(
				regexp.MustCompile("https://download.openstreetmap.fr/$"),
				regexp.MustCompile("https://download.openstreetmap.fr/extracts/$"),
				regexp.MustCompile("https://download.openstreetmap.fr/polygons/$"),
				regexp.MustCompile("https://download.openstreetmap.fr/extracts/[a-zA-Z].+"),
				regexp.MustCompile("https://download.openstreetmap.fr/polygons/[a-zA-Z].+"),
			),
			colly.Async(true),
		)
		c.Limit(&colly.LimitRule{
			Parallelism: 2,
			RandomDelay: 5 * time.Second,
		})

		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})
		c.OnHTML("a", func(e *colly.HTMLElement) {
			openstreetmapFRParse(e, &ext, bar, func(arg interface{}) {
				c.Visit(arg.(string))
			})
		})
		c.OnScraped(func(*colly.Response) {
			if *fProgress {
				bar.Increment()
			}
		})
		c.Visit("https://download.openstreetmap.fr/")
		c.Wait()
		out, _ := ext.Elements.Generate(&myConfig)
		filename, _ := filepath.Abs(configfile)
		err := ioutil.WriteFile(filename, out, 0644)
		if err != nil {
			log.Panicln(fmt.Errorf(" File error: %v ", err))
		}
		if !*fQuiet {
			log.Println(configfile, " generated.")
		}
	case "gislab":
		var myConfig Config
		myConfig.BaseURL = "http://be.gis-lab.info/project/osm_dump"
		myConfig.Formats = make(map[string]format)
		myConfig.Formats["osm.pbf"] = format{ID: "osm.pbf", BaseURL: "http://data.gis-lab.info/osm_dump/dump", BasePath: "latest/", Loc: ".osm.pbf"}
		myConfig.Formats["osm.bz2"] = format{ID: "osm.bz2", BaseURL: "http://data.gis-lab.info/osm_dump/dump", BasePath: "latest/", Loc: ".osm.bz2"}
		myConfig.Formats["poly"] = format{ID: "poly", BaseURL: "https://raw.githubusercontent.com/nextgis/osmdump_poly/master", Loc: ".poly"}
		ext := Ext{Elements: make(map[string]Element)}
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			colly.AllowedDomains("be.gis-lab.info"),
		)
		c.OnHTML("table", func(e *colly.HTMLElement) {
			gislabParse(e, &ext)
		})
		c.Visit("http://be.gis-lab.info/project/osm_dump/iframe.php")
		c.Wait()
		//GenerateCrawler("http://be.gis-lab.info/project/osm_dump/iframe.php", configfile, &myConfig)
		out, _ := ext.Elements.Generate(&myConfig)
		filename, _ := filepath.Abs(configfile)
		err := ioutil.WriteFile(filename, out, 0644)
		if err != nil {
			log.Panicln(fmt.Errorf(" File error: %v ", err))
		}
		if !*fQuiet {
			log.Println(configfile, " generated.")
		}
	case "bbbike":
		var myConfig Config
		myConfig.BaseURL = "https://download.bbbike.org/osm/bbbike/"
		myConfig.Formats = make(map[string]format)
		myConfig.Formats["osm.pbf"] = format{ID: "osm.pbf", Loc: ".osm.pbf"}
		myConfig.Formats["shp.zip"] = format{ID: "shp.zip", Loc: ".osm.shp.zip"}
		ext := Ext{Elements: make(map[string]Element), ElementsMutex: sync.RWMutex{}}
		if *fProgress {
			bar = pb.New(bbbikePb)
			bar.Start()
		}
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			colly.AllowedDomains("download.bbbike.org"),
			colly.URLFilters(
				regexp.MustCompile("https://download.bbbike.org/osm/bbbike/[A-Z].+"),
				regexp.MustCompile("https://download.bbbike.org/osm/bbbike/"),
			),
			//colly.Async(true),
		)
		c.Limit(&colly.LimitRule{
			Parallelism: 1,
			RandomDelay: 5 * time.Second,
			//Delay: 5 * time.Second,
		})
		/*c.WithTransport(&http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          0,
			IdleConnTimeout:       5 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
		})*/

		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err.Error())
		})

		c.OnHTML("div.list", func(e *colly.HTMLElement) {
			bbbikeParseList(e, &ext, c)
		})
		c.OnHTML("#sidebar", func(e *colly.HTMLElement) {
			bbbikeParseSidebar(e, &ext, c)
		})

		c.OnScraped(func(*colly.Response) {
			if *fProgress {
				bar.Increment()
			}
		})
		c.Visit("https://download.bbbike.org/osm/bbbike/")
		c.Wait()
		out, _ := ext.Elements.Generate(&myConfig)
		filename, _ := filepath.Abs(configfile)
		err := ioutil.WriteFile(filename, out, 0644)
		if err != nil {
			log.Panicln(fmt.Errorf(" File error: %v ", err))
		}
		if !*fQuiet {
			log.Println(configfile, " generated.")
		}
	default:
		log.Println("Service not reconized, please use one of geofabrik, openstreetmap.fr or gislab")
	}
}
