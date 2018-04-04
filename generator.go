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

// Ext simple struct for managing ElementSlice and crawler
type Ext struct {
	*gocrawl.DefaultExtender
	Elements ElementSlice
}

// addHash find if a hash is available and append it to e
func (e *Element) addHash(myel *goquery.Selection) {
	a := myel.Find("a")
	if a.Length() == 2 { // If only 1 a there is no hash
		validHash := []string{"md5"}
		val, exist := a.Eq(1).Attr("href")
		if exist {
			splitted := strings.Split(val, ".")
			hash := splitted[len(splitted)-1]
			if stringInSlice(&hash, &validHash) {
				hashfile := strings.Join(splitted[1:], ".")
				//fmt.Println(hashfile)
				e.Formats = append(e.Formats, hashfile)
			}
		}
	}
}

func (e *Ext) parseGeofabrik(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	var thisElement Element
	downloadMain := doc.Find("div#download-main")
	parent, haveParent := doc.Find("p a").Attr("href")                        // I'm not shure it is parent
	if haveParent && !strings.Contains(parent, "https://www.geofabrik.de/") { // Removing https?
		// or try with slicer and == should be quicker?
		parent = parent[0 : len(parent)-5] // remove ".html"
		if parent == "index" {             // first level
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
		us.Formats = []string{}
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

func (e *Ext) mergeElement(element *Element) error {
	if cE, ok := e.Elements[element.ID]; ok {
		if cE.Parent != element.Parent {
			return fmt.Errorf("Cant merge : Parent mismatch")
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
	return nil
}

func (e *Ext) parseOSMfr(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	parents := strings.Split(doc.Url.Path, "/")
	parent := parents[len(parents)-2]          // Get x in this kind of url http(s)://1/2/.../x/
	if strings.EqualFold(parent, "extracts") { // should I try == or a switch?
		parent = ""
	} else if strings.EqualFold(parent, "polygons") {
		parent = ""
	}
	list := doc.Find("table tr")
	for line := range list.Nodes {
		singleElement := list.Eq(line)
		link := singleElement.Find("a")
		//index := 0
		for aa := range link.Nodes {
			a := link.Eq(aa)
			vallink, link := a.Attr("href") // get first link
			if link {
				// Filtering
				if !strings.Contains(vallink, "?") && !strings.Contains(vallink, "-latest") && vallink[0] != '/' && !strings.EqualFold(vallink, "cgi-bin/") && vallink[len(vallink)-1] != '/' {
					element := *new(Element)
					element.Parent = parent
					valsplit := strings.Split(vallink, ".")
					name := valsplit[0]
					//log.Println("name", name)
					ext := strings.Join(valsplit[1:], ".")
					if strings.Contains(ext, "state.txt") { // I'm shure it can be refactorized
						ext = "state"
					}
					element.ID = name
					element.Name = name
					if *fVerbose && !*fQuiet && !*fProgress {
						log.Println("parsing", vallink)
					}
					if !strings.EqualFold(e.Elements[name].ID, name) {
						element.Formats = append(element.Formats, ext)
						err := e.mergeElement(&element)
						if err != nil {
							log.Panicln("Can't merge element,", err)
						}
					} else {
						if *fVerbose && !*fQuiet && !*fProgress {
							log.Println(name, "already exist")
							log.Println("Merging formats")
						}
						et := e.Elements[name]
						if len(et.Formats) == 0 {
							et.Meta = true
						} else {
							et.Meta = false
						}
						et.Formats = append(et.Formats, ext)
						e.Elements[name] = et
					}
				}
			}
		}
	}
	return nil, true
}

func (e *Ext) parseGisLab(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	list := doc.Find("table tr")
	for line := range list.Nodes {
		tds := list.Eq(line).Find("td")
		if tds.Length() == 6 {
			element := *(new(Element))
			element.ID = tds.Eq(0).Text()
			element.Name = tds.Eq(1).Text()
			element.Formats = append(element.Formats, "osm.pbf") // Not checked elements
			element.Formats = append(element.Formats, "osm.bz2") // Pray for non changing data structure...
			element.Formats = append(element.Formats, "poly")    // Not checked but seems to be used for generating osm.pbf/osm.bz2
			if *fVerbose && !*fQuiet {
				log.Println("Adding", element.Name)
			}
			err := e.mergeElement(&element)
			if err != nil {
				log.Panicln("Can't merge element,", err)
			}
		}
	}
	return nil, true
}

// Visit launch right crawler
func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	if *fVerbose && !*fQuiet && !*fProgress {
		fmt.Printf("Visit: %s\n", ctx.URL())
	}
	if *fProgress {
		bar.Increment()
	}
	switch ctx.URL().Host {
	case "download.geofabrik.de":
		return e.parseGeofabrik(ctx, res, doc)
	case "download.openstreetmap.fr":
		return e.parseOSMfr(ctx, res, doc)
	case "be.gis-lab.info":
		return e.parseGisLab(ctx, res, doc)
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
	} else if strings.Contains(ctx.URL().Path, "?") {
		return false
	} else if ctx.URL().Path[len(ctx.URL().Path)-1:] == "/" {
		return true
	} else if strings.Contains(ctx.URL().Path, ".html") {
		return true
	} else if strings.Contains(ctx.URL().Path, ".php") {
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
	if *fProgress {
		maxPb := 400 // default value is a realy magicaly set :)
		switch url { // Todo: found a better way!
		case "https://download.geofabrik.de/":
			maxPb = 392 // Magical!
		case "https://download.openstreetmap.fr/":
			maxPb = 51 // Magical
		}
		bar = pb.New(maxPb)
		bar.Start()
	}
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
	switch *fService {
	case "geofabrik":
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

	case "openstreetmap.fr":
		var myConfig Config
		myConfig.BaseURL = "https://download.openstreetmap.fr/extracts"
		myConfig.Formats = make(map[string]format)
		myConfig.Formats["osm.pbf"] = format{ID: "osm.pbf", Loc: "-latest.osm.pbf"}
		myConfig.Formats["poly"] = format{ID: "poly", Loc: ".poly", BasePath: "../polygons/"}
		myConfig.Formats["state"] = format{ID: "state", Loc: ".state.txt"}
		GenerateCrawler("https://download.openstreetmap.fr/", configfile, &myConfig)
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
		GenerateCrawler("http://be.gis-lab.info/project/osm_dump/iframe.php", configfile, &myConfig)
		if !*fQuiet {
			log.Println(configfile, " generated.")
		}
	default:
		log.Println("Service not reconized, please use one of geofabrik, openstreetmap.fr or gislab")
	}

}
