package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app         = kingpin.New("download-geofabrik", "A command-line tool for downloading OSM files.")
	fConfig     = app.Flag("config", "Set Config file.").Default("./geofabrik.yml").Short('c').String()
	fNodownload = app.Flag("nodownload", "Do not download file (test only)").Short('n').Bool()
	fVerbose    = app.Flag("verbose", "Be verbose").Short('v').Bool()
	fQuiet      = app.Flag("quiet", "Be quiet").Short('q').Bool()
	fProxyHTTP  = app.Flag("proxy-http", "Use http proxy, format: proxy_address:port").Default("").String()
	fProxySock5 = app.Flag("proxy-sock5", "Use Sock5 proxy, format: proxy_address:port").Default("").String()
	fProxyUser  = app.Flag("proxy-user", "Proxy user").Default("").String()
	fProxyPass  = app.Flag("proxy-pass", "Proxy password").Default("").String()

	update = app.Command("update", "Update geofabrik.yml from github *** DEPRECATED you should prefer use generate ***")
	fURL   = update.Flag("url", "Url for config source").Default("https://raw.githubusercontent.com/julien-noblet/download-geofabrik/master/geofabrik.yml").String()

	list = app.Command("list", "Show elements available")
	lmd  = list.Flag("markdown", "generate list in Markdown format").Bool()

	download = app.Command("download", "Download element") //TODO : add d as command
	delement = download.Arg("element", "OSM element").Required().String()
	dosmBz2  = download.Flag("osm.bz2", "Download osm.bz2 if available").Short('B').Bool()
	dshpZip  = download.Flag("shp.zip", "Download shp.zip if available").Short('S').Bool()
	dosmPbf  = download.Flag("osm.pbf", "Download osm.pbf (default)").Short('P').Bool()
	doshPbf  = download.Flag("osh.pbf", "Download osh.pbf (default)").Short('H').Bool()
	dstate   = download.Flag("state", "Download state.txt file").Short('s').Bool()
	dpoly    = download.Flag("poly", "Download poly file").Short('p').Bool()

	generate = app.Command("generate", "Generate a new config file")
)

func listAllRegions(c Config, format string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ShortName", "Is in", "Long Name", "formats"})
	if format == "Markdown" {
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
	}
	keys := make(sort.StringSlice, len(c.Elements))
	i := 0
	for k := range c.Elements {
		keys[i] = k
		i++
	}
	keys.Sort()
	for _, item := range keys {
		table.Append([]string{item, c.Elements[c.Elements[item].Parent].Name, c.Elements[item].Name, miniFormats(c.Elements[item].Formats)})
	}
	table.Render()
	fmt.Printf("Total elements: %#v\n", len(c.Elements))
}

// UpdateConfig : simple script to download lastest config from repo
func UpdateConfig(myURL *string, myconfig string) {
	downloadFromURL(myURL, myconfig)
	if !*fQuiet {
		log.Println("*** DEPRECATED you should prefer use generate ***")
	}
	if !*fVerbose {
		log.Println("Congratulation, you have the latest geofabrik.yml")
	}
}

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case list.FullCommand():
		var format = ""
		if *lmd {
			format = "Markdown"
		}
		listAllRegions(*loadConfig(*fConfig), format)
	case update.FullCommand():
		UpdateConfig(fURL, *fConfig)
	case download.FullCommand():
		formatFile := getFormats()
		for _, format := range *formatFile {
			downloadFromURL(elem2URL(loadConfig(*fConfig), findElem(loadConfig(*fConfig), *delement), format), *delement+"."+format)
		}
	case generate.FullCommand():
		Generate(*fConfig)
	}

}
