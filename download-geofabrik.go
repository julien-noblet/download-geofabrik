package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"

	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

type config struct {
	BaseURL  string             `yaml:"baseURL"`
	Elements map[string]element `yaml:"elements,flow"`
}

type element struct {
	ID     string   `yaml:"id"`
	File   string   `yaml:"file"`
	Name   string   `yaml:"name"`
	Files  []string `yaml:"files"`
	Parent string   `yaml:"parent"`
}

var (
	app        = kingpin.New("download-geofabrik", "A command-line tool for downloading OSM files.")
	Fconfig    = app.Flag("config", "Set Config file.").Default("./geofabrik.yml").Short('c').String()
	nodownload = app.Flag("nodownload", "Do not download file (test only)").Short('n').Bool()
	verbose    = app.Flag("verbose", "Be verbose").Short('v').Bool()

	update = app.Command("update", "Update geofabrik.yml from github")
	url    = update.Flag("url", "Url for config source").Default("https://raw.githubusercontent.com/julien-noblet/download-geofabrik/stable/geofabrik.yml").String()

	list = app.Command("list", "Show elements available")

	download = app.Command("download", "Download element") //TODO : add d as command
	delement = download.Arg("element", "OSM element").Required().String()
	dosmBz2  = download.Flag("osm.bz2", "Download osm.bz2 if available").Short('B').Bool()
	dshpZip  = download.Flag("shp.zip", "Download shp.zip if available").Short('S').Bool()
	dosmPbf  = download.Flag("osm.pbf", "Download osm.pbf (default)").Short('P').Bool()
	dstate   = download.Flag("state", "Download state.txt file").Short('s').Bool()
	dpoly    = download.Flag("poly", "Download poly file").Short('p').Bool()
)

func (e *element) hasParent() bool {
	return len(e.Parent) != 0
}

func miniFormats(s []string) string {
	res := make([]string, 3)
	for _, item := range s {
		if item == "osm.pbf" {
			res[0] = "p"
		}
		if item == "osm.bz2" {
			res[1] = "b"
		}
		if item == "shp.zip" {
			res[2] = "s"
		}
	}

	return strings.Join(res, "")
}

func downloadFromURL(url string, fileName string) {
	if *verbose == true {
		log.Println(" Downloading", url, "to", fileName)
	}

	if *nodownload == false {
		// TODO: check file existence first with io.IsExist
		output, err := os.Create(fileName)
		if err != nil {
			log.Fatalln(" Error while creating", fileName, "-", err)
			return
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			log.Fatalln(" Error while downloading", url, "-", err)
			return
		}
		defer response.Body.Close()

		n, err := io.Copy(output, response.Body)
		if err != nil {
			log.Fatalln(" Error while downloading", url, "-", err)
			return
		}

		if *verbose == true {
			log.Println(" ", n, "bytes downloaded.")
		}
	}
}
func elem2preURL(c config, e element) string {
	var res string
	if e.hasParent() {
		res = elem2preURL(c, findElem(c, e.Parent)) + "/"
		if e.File != "" { //TODO use file in config???
			res = res + e.File
		} else {
			res = res + e.ID
		}
	} else {
		res = c.BaseURL + "/" + e.ID
	}
	return res
}

func elem2URL(c config, e element, ext string) string {
	res := elem2preURL(c, e)
	switch ext {
	case "state":
		res += "-updates/state.txt"
	case "poly":
		res += "." + ext
	default:
		res += "-latest." + ext
		if !stringInSlice(ext, e.Files) {
			log.Fatalln(" Error!!!" + res + "not exist")
		}
	}

	return res
}

func findElem(c config, e string) element {
	res := c.Elements[e]
	if res.ID == "" {
		log.Fatalln(" " + e + " is not in config! Please use \"list\" command!")
	}
	return res
}
func getFormats() []string {
	var formatFile []string
	if *dosmPbf {
		formatFile = append(formatFile, "osm.pbf")
	}
	if *dosmBz2 {
		formatFile = append(formatFile, "osm.bz2")
	}
	if *dshpZip {
		formatFile = append(formatFile, "shp.zip")
	}
	if *dstate {
		formatFile = append(formatFile, "state")
	}
	if *dpoly {
		formatFile = append(formatFile, "poly")
	}
	if len(formatFile) == 0 {
		formatFile = append(formatFile, "osm.pbf")
	}
	return formatFile
}

func listAllRegions(c config) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ShortName", "Is in", "Long Name", "formats"})
	keys := make(sort.StringSlice, len(c.Elements))
	i := 0
	for k := range c.Elements {
		keys[i] = k
		i++
	}
	keys.Sort()
	for _, item := range keys {
		table.Append([]string{item, c.Elements[c.Elements[item].Parent].Name, c.Elements[item].Name, miniFormats(c.Elements[item].Files)})
	}
	table.Render()
	fmt.Printf("Total elements: %#v\n", len(c.Elements))
}

func loadConfig(configFile string) config {
	filename, _ := filepath.Abs(configFile)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(" File error: %v", err)
		os.Exit(1)
	}
	var myConfig config
	err = yaml.Unmarshal(file, &myConfig)
	if err != nil {
		panic(err)
	}
	return myConfig

}
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func UpdateConfig(url string, myconfig string) {
	downloadFromURL(url, myconfig)
	fmt.Println("Congratulation, you have the latest geofabrik.yml\n")
}

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case list.FullCommand():
		listAllRegions(loadConfig(*Fconfig))
	case update.FullCommand():
		UpdateConfig(*url, *Fconfig)
	case download.FullCommand():
		formatFile := getFormats()
		for _, format := range formatFile {
			downloadFromURL(elem2URL(loadConfig(*Fconfig), findElem(loadConfig(*Fconfig), *delement), format), *delement+"."+format)
		}
	}
}
