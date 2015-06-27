package main

import (
	"flag"
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

	"gopkg.in/yaml.v2"
)

type config struct {
	BaseURL  string             `yaml:"baseURL"`
	Elements map[string]element `yaml:"elements,flow"`
}

type element struct {
	ID     string   `yaml:"id"`
	Name   string   `yaml:"name"`
	Files  []string `yaml:"files"`
	Parent string   `yaml:"parent"`
}

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

func downloadFromURL(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

func findElem(c *config, e string) element {
	res := c.Elements[e]
	return res
}

func elem2preURL(c *config, e element) string {
	var res string
	if e.hasParent() {
		res = elem2preURL(c, findElem(c, e.Parent)) + "/"
		if e.ID == "georgia-eu" || e.ID == "georgia-us" {
			res = res + "georgia"
		} else {
			res = res + e.ID
		}
	} else {
		res = c.BaseURL + "/" + e.ID
	}
	return res
}

func elem2URL(c *config, e element, ext string) string {
	res := elem2preURL(c, e)

	res += "-latest." + ext
	if !stringInSlice(ext, e.Files) {
		fmt.Println("Error!!!\n" + res + " not exist")
	}
	return res
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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

func main() {
	configFile := flag.String("config", "./geofabrik.yml", "Config for downloading OSMFiles")
	nodownload := flag.Bool("n", false, "Download")
	osmBz2 := flag.Bool("osm.bz2", false, "Download osm.bz2 if available")
	shpZip := flag.Bool("shp.zip", false, "Download shp.zip if available")
	osmPbf := flag.Bool("osm.pbf", false, "Download osm.pbf (default)")
	list := flag.Bool("list", false, "list all elements")
	update := flag.Bool("update", false, "Update geofabrik.yml from github")

	flag.Parse()

	if *update {
		downloadFromURL("https://raw.githubusercontent.com/julien-noblet/download-geofabrik/stable/geofabrik.yml")
		log.Fatalln("\nCongratulation, you have the latest geofabrik.yml\n")
	}

	if (flag.NArg() < 1) && !*list && !*update {
		log.Fatalln("\nThis program needs one argument or more\nMore info at: https://github.com/julien-noblet/download-geofabrik\n")
	}

	filename, _ := filepath.Abs(*configFile)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var myConfig config
	err = yaml.Unmarshal(file, &myConfig)
	if err != nil {
		panic(err)
	}

	if *list {
		listAllRegions(myConfig)
	}

	var formatFile []string
	if *osmPbf {
		formatFile = append(formatFile, "osm.pbf")
	}
	if *osmBz2 {
		formatFile = append(formatFile, "osm.bz2")
	}
	if *shpZip {
		formatFile = append(formatFile, "shp.zip")
	}
	if len(formatFile) == 0 {
		formatFile = append(formatFile, "osm.pbf")
	}

	if !*nodownload {
		for _, elname := range flag.Args() {
			for _, format := range formatFile {
				downloadFromURL(elem2URL(&myConfig, findElem(&myConfig, elname), format))
			}
		}
	} else {
		for _, elname := range flag.Args() {
			for _, format := range formatFile {
				fmt.Printf("(not) Downloading : %#v", elem2URL(&myConfig, findElem(&myConfig, elname), format))
			}
		}
	}

}
