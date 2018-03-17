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
	yaml "gopkg.in/yaml.v2"
)

// ElementSlice contain all Elements
type ElementSlice map[string]Element

// Generate make the slice which contain all Elements
func (e ElementSlice) Generate(myConfig *Config) ([]byte, error) {
	myConfig.Elements = e
	return yaml.Marshal(myConfig)
}

// Ext simple struct for managing ElementSlice and crawler
type Ext struct {
	*gocrawl.DefaultExtender
	Elements ElementSlice
}

func (e *Element) addHash(myel *goquery.Selection) {
	validHash := []string{"md5"}
	if myel.Find("a").Length() == 2 {
		l2v, l2z := myel.Find("a").Eq(1).Attr("href")
		if l2z {
			hash := strings.Split(l2v, ".")[len(strings.Split(l2v, "."))-1]
			if stringInSlice(&hash, &validHash) {
				hashfile := strings.Join(strings.Split(l2v, ".")[1:], ".")
				//fmt.Println(hashfile)
				e.Formats = append(e.Formats, hashfile)
			}
		}
	}
}

func (e *Ext) parseGeofabrik(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	var thisElement Element
	downloadMain := doc.Find("div#download-main")
	parent, haveParent := doc.Find("p a").Attr("href")
	if haveParent && !strings.Contains(parent, "https://www.geofabrik.de/") {
		parent = parent[0 : len(parent)-5] // remove ".html"
		if parent == "index" {             //first level
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
						thisElement.ID = linkval[0 : len(linkval)-15]
						thisElement.Formats = append(thisElement.Formats, "osm.pbf")
						thisElement.addHash(myel)
					case 1: // shp.zip
						thisElement.Formats = append(thisElement.Formats, "shp.zip")
						thisElement.addHash(myel)
					case 2: // osm.bz2
						thisElement.Formats = append(thisElement.Formats, "osm.bz2")
						thisElement.addHash(myel)
					case 3: // osh.pbf
						thisElement.Formats = append(thisElement.Formats, "osh.pbf")
						thisElement.addHash(myel)
					case 4: // poly & kml
						thisElement.Formats = append(thisElement.Formats, "poly")
						thisElement.Formats = append(thisElement.Formats, "kml")
						thisElement.addHash(myel)
					case 5: //-updates
						thisElement.Formats = append(thisElement.Formats, "state")
						thisElement.addHash(myel)
					}
				}
				index++
			}
		}
		if len(thisElement.Formats) == 0 {
			thisElement.Meta = true
		}
		// Workaround to fix #10
		var us Element
		us.Meta = true
		us.ID = "us"
		us.Name = "United States of America"
		us.Parent = "north-america"
		e.Elements[us.ID] = us

		//Exceptions!
		// Only Georgia (EU and US)
		if thisElement.ID == "georgia" {
			thisElement.File = "georgia"
			if thisElement.Parent == "europe" {
				thisElement.Name = "Georgia (Europe country)"
				thisElement.ID = "georgia-eu"
			} else {
				thisElement.Name = "Georgia (US State)"
				thisElement.ID = "georgia-us"
				thisElement.Parent = "us"
			}
		}
		// List of US to fix #10
		usList := map[string]bool{
			"alabama":              true,
			"alaska":               true,
			"arizona":              true,
			"arkansas":             true,
			"california":           true,
			"colorado":             true,
			"connecticut":          true,
			"delaware":             true,
			"district-of-columbia": true,
			"florida":              true,
			"georgia":              false, // Since there is also georgia in europe....
			"hawaii":               true,
			"idaho":                true,
			"illinois":             true,
			"indiana":              true,
			"iowa":                 true,
			"kansas":               true,
			"kentucky":             true,
			"louisiana":            true,
			"maine":                true,
			"maryland":             true,
			"massachusetts":        true,
			"michigan":             true,
			"minnesota":            true,
			"mississippi":          true,
			"missouri":             true,
			"montana":              true,
			"nebraska":             true,
			"nevada":               true,
			"new-hampshire":        true,
			"new-jersey":           true,
			"new-mexico":           true,
			"new-york":             true,
			"north-carolina":       true,
			"north-dakota":         true,
			"ohio":                 true,
			"oklahoma":             true,
			"oregon":               true,
			"pennsylvania":         true,
			"puerto-rico":          true,
			"rhode-island":         true,
			"south-carolina":       true,
			"south-dakota":         true,
			"tennessee":            true,
			"texas":                true,
			"utah":                 true,
			"vermont":              true,
			"virginia":             true,
			"washington":           true,
			"west-virginia":        true,
			"wisconsin":            true,
			"wyoming":              true}

		if usList[thisElement.ID] {
			thisElement.Parent = "us"
		}
		if thisElement.Name != "OpenStreetMap Data Extracts" {
			e.Elements[thisElement.ID] = thisElement
		}
	}
	return nil, true
}

func (e *Ext) mergeElement(element *Element) {
	if cE, ok := e.Elements[element.ID]; ok {
		//cE := &e.Elements[element.ID]

		if cE.Parent != element.Parent {
			panic(fmt.Sprintln("Error! : Parent mismatch!"))
		}
		cE.Formats = append(cE.Formats, element.Formats...)
		if len(cE.Formats) == 0 {
			cE.Meta = true
		} else {
			cE.Meta = false
		}
		e.Elements[element.ID] = cE
	} else {
		e.Elements[element.ID] = *element
	}
}

func (e *Ext) parseOSMfr(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	//var thisElement Element
	parent := doc.Find("h1").Text()
	fmt.Println(parent)

	list := doc.Find("table tr")
	for line := range list.Nodes {
		singleElement := list.Eq(line)
		lien := singleElement.Find("a")
		//index := 0
		for aa := range lien.Nodes {
			var element Element
			element.Parent = parent
			a := lien.Eq(aa)
			vallink, link := a.Attr("href")
			if link {
				if strings.Contains(vallink, "/") {
					// /!\ meta! or levelup
					if vallink[0:1] != "/" {
						//meta
						// looking if already in e.Elements
						element.Meta = true
						t := strings.Split(vallink, "/")
						element.ID = t[len(t)-1]
					}
				}
			}
			e.mergeElement(&element)
		}
	}
	return nil, true
}

// Visit launch right crawler
func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	if *fVerbose {
		fmt.Printf("Visit: %s\n", ctx.URL())
	}
	switch ctx.URL().Host {
	case "download.geofabrik.de":
		return e.parseGeofabrik(ctx, res, doc)
	case "download.openstreetmap.fr":
		return e.parseOSMfr(ctx, res, doc)
	default:
		panic(fmt.Sprintln("Panic! " + ctx.URL().Host + " is not supported!"))
	}

}

// Filter remove non needed urls.
func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	if isVisited {
		return false
	}
	if len(ctx.URL().RawQuery) != 0 {
		return false
		// TODO: refactorize? Use config file?
	} else if strings.Contains(ctx.URL().Path, "newshapes.html") {
		return false
	} else if strings.Contains(ctx.URL().Path, "technical.html") {
		return false
	} else if strings.Contains(ctx.URL().Path, "robots.txt") {
		return false
	} else if strings.Contains(ctx.URL().Path, "replication") {
		return false
	} else if strings.Contains(ctx.URL().Path, "cgi-bin") {
		return false
	} else if strings.Contains(ctx.URL().Path, ".pdf") {
		return false
	} else if strings.Contains(ctx.URL().Path, ".pbf") {
		return false
	} else if strings.Contains(ctx.URL().Path, ".poly") {
		return false
	} else if strings.Contains(ctx.URL().Path, ".kml") {
		return false
	} else if strings.Contains(ctx.URL().Path, ".bz2") {
		return false
	} else if strings.Contains(ctx.URL().Path, ".zip") {
		return false
	} else if ctx.URL().Path[len(ctx.URL().Path)-1:] == "/" {
		return true
	} else if strings.Contains(ctx.URL().Path, ".html") {
		return true
		//	} else if ctx.URL().Path[len(ctx.URL().Path)-8:] == "-updates" {
		//		return false
		//	} else {
		//		return false
	}
	return false
}

// GenerateCrawler creating a gocrawl to parse the website.
func GenerateCrawler(url string, fname string, myConfig *Config) {
	ext := &Ext{&gocrawl.DefaultExtender{}, make(map[string]Element)}
	// Set custom options
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 100 * time.Millisecond
	opts.LogFlags = gocrawl.LogError
	//	opts.LogFlags = gocrawl.LogAll
	opts.SameHostOnly = true //false
	opts.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36"
	opts.MaxVisits = 15000

	file := gocrawl.NewCrawlerWithOptions(opts)
	file.Run(url)
	out, _ := ext.Elements.Generate(myConfig)
	filename, _ := filepath.Abs(fname)
	err := ioutil.WriteFile(filename, out, 0644)
	if err != nil {
		log.Fatalln(fmt.Errorf(" File error: %v ", err))
		os.Exit(1)
	}
}

//Generate main function
func Generate(configfile string) {
	//Generate geofabrik.yml
	var geofabrik Config
	geofabrik.BaseURL = "https://download.geofabrik.de"
	geofabrik.Formats = make(map[string]format)
	//TODO: make a function for adding formats
	geofabrik.Formats["osh.pbf"] = format{ID: "osh.pbf", Loc: ".osh.pbf"}
	geofabrik.Formats["osh.pbf.md5"] = format{ID: "osh.pbf.md5", Loc: ".osh.pbf.md5"}
	geofabrik.Formats["osm.bz2"] = format{ID: "osm.bz2", Loc: "-latest.osm.bz2"}
	geofabrik.Formats["osm.bz2.md5"] = format{ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"}
	geofabrik.Formats["osm.pbf"] = format{ID: "osm.pbf", Loc: "-latest.osm.pbf"}
	geofabrik.Formats["osm.pbf.md5"] = format{ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"}
	geofabrik.Formats["poly"] = format{ID: "poly", Loc: ".poly"}
	geofabrik.Formats["kml"] = format{ID: "kml", Loc: ".kml"}
	geofabrik.Formats["state"] = format{ID: "state", Loc: "-updates/state.txt"}
	geofabrik.Formats["shp.zip"] = format{ID: "shp.zip", Loc: "-latest-free.shp.zip"}
	GenerateCrawler("https://download.geofabrik.de/", configfile, &geofabrik)
	if !*fQuiet {
		log.Println(configfile, " generated.")
	}
	/*
		var myConfig config
		myConfig.BaseURL = "https://download.openstreetmap.fr/extracts"
		myConfig.Formats = make(map[string]format)
		myConfig.Formats["osm.pbf"] = format{ID: "osm.pbf", Loc: ".osm.pbf"}
		myConfig.Formats["poly"] = format{ID: "poly", Loc: ".poly", BasePath: "../polygons/"}
		myConfig.Formats["state"] = format{ID: "state", Loc: ".state.txt"}
		generate("https://download.openstreetmap.fr/", "openstreetmap_fr.yml", &myConfig) // TODO: Not Working!
	*/
}
