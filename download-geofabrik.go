package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app         = kingpin.New("download-geofabrik", "A command-line tool for downloading OSM files.")
	fService    = app.Flag("service", "Can switch to another service. You can use \"geofabrik\" or \"openstreetmap.fr\". It automatically change config file if -c is unused.").Default("geofabrik").String()
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
	doshPbf  = download.Flag("osh.pbf", "Download osh.pbf").Short('H').Bool()
	dstate   = download.Flag("state", "Download state.txt file").Short('s').Bool()
	dpoly    = download.Flag("poly", "Download poly file").Short('p').Bool()
	dkml     = download.Flag("kml", "Download kml file").Short('k').Bool()
	dCheck   = download.Flag("check", "control with checksum").Bool()

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
func UpdateConfig(myURL string, myconfig string) error {
	if !*fQuiet {
		log.Print("*** DEPRECATED you should prefer use generate ***")
	}
	err := downloadFromURL(myURL, myconfig)
	if err != nil {
		if *fVerbose {
			log.Println(err)
		}
		return (fmt.Errorf("Can't updating %v please use generate", myconfig))
	}
	if *fVerbose && !*fQuiet {
		log.Println("Congratulation, you have the latest geofabrik.yml")
	}
	return nil
}

func checkService() bool {
	switch *fService {
	case "geofabrik":
		return true
	case "openstreetmap.fr":
		if strings.EqualFold(*fConfig, "./geofabrik.yml") {
			*fConfig = "./openstreetmap.fr.yml"
		}
		return true
	}
	return false
}

func catch(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case list.FullCommand():
		checkService()
		var format = ""
		if *lmd {
			format = "Markdown"
		}
		configPtr, err := loadConfig(*fConfig)
		catch(err)
		listAllRegions(*configPtr, format)
	case update.FullCommand():
		checkService()
		UpdateConfig(*fURL, *fConfig)
	case download.FullCommand():
		checkService()
		configPtr, err := loadConfig(*fConfig)
		catch(err)
		formatFile := getFormats()
		for _, format := range *formatFile {
			if *dCheck && fileExist(*delement+"."+format) {
				if !(downloadChecksum(format)) {
					if !*fQuiet {
						log.Println("Checksum mismatch, re-downloading", *delement+"."+format)
					}
					myElem, err := findElem(configPtr, *delement)
					catch(err)
					myURL, err := elem2URL(configPtr, myElem, format)
					catch(err)
					downloadFromURL(myURL, *delement+"."+format)
					downloadChecksum(format)

				} else {
					if !*fQuiet {
						log.Printf("Checksum match, no download!")
					}
				}
			} else {
				myElem, err := findElem(configPtr, *delement)
				catch(err)
				myURL, err := elem2URL(configPtr, myElem, format)
				catch(err)
				downloadFromURL(myURL, *delement+"."+format)
				downloadChecksum(format)
			}
		}
	case generate.FullCommand():
		checkService()
		Generate(*fConfig)
	}

}

func fileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func hashFileMD5(filePath string) (string, error) {
	var returnMD5String string
	if fileExist(filePath) {
		file, err := os.Open(filePath)
		if err != nil {
			return returnMD5String, err
		}
		defer file.Close()
		hash := md5.New()

		if _, err := io.Copy(hash, file); err != nil {
			return returnMD5String, err
		}
		hashInBytes := hash.Sum(nil)[:16]
		returnMD5String = hex.EncodeToString(hashInBytes)
		return returnMD5String, nil
	}
	return returnMD5String, nil
}

func controlHash(hashfile string, hash string) (bool, error) {
	if fileExist(hashfile) {
		file, err := ioutil.ReadFile(hashfile)
		if err != nil {
			return false, err
		}
		filehash := strings.Split(string(file), " ")[0]
		if *fVerbose && !*fQuiet {
			log.Println("Hash from file :", filehash)
		}
		if strings.EqualFold(hash, filehash) {
			return true, nil
		}
	}
	return false, nil
}

func downloadChecksum(format string) bool {
	ret := false
	if *dCheck {
		hash := "md5"
		fhash := format + "." + hash
		configPtr, err := loadConfig(*fConfig)
		catch(err)
		myElem, err := findElem(configPtr, *delement)
		catch(err)
		if stringInSlice(&fhash, &myElem.Formats) {
			myURL, err := elem2URL(configPtr, myElem, fhash)
			catch(err)
			downloadFromURL(myURL, *delement+"."+fhash)
			if *fVerbose && !*fQuiet {
				log.Println("Hashing", *delement+"."+format)
			}
			hashed, err := hashFileMD5(*delement + "." + format)
			if err != nil {
				log.Panic(fmt.Errorf(err.Error()))
			}
			if *fVerbose && !*fQuiet {
				log.Println("MD5 :", hashed)
			}
			ret, err := controlHash(*delement+"."+fhash, hashed)
			if err != nil {
				log.Panic(fmt.Errorf(err.Error()))
			}
			if !*fQuiet {
				if ret {
					log.Println("Checksum OK for", *delement+"."+format)
				} else {
					log.Println("Checksum MISMATCH for", *delement+"."+format)
				}
			}
			return ret
		}
		if !*fQuiet {
			log.Printf("No checksum provided for" + *delement + "." + format)
		}
	}
	return ret
}
