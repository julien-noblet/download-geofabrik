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

var version = "devel"

var (
	app         = kingpin.New("download-geofabrik", "A command-line tool for downloading OSM files.")
	fService    = app.Flag("service", "Can switch to another service. You can use \"geofabrik\", \"openstreetmap.fr\" or \"bbbike\". It automatically change config file if -c is unused.").Default("geofabrik").String()
	fConfig     = app.Flag("config", "Set Config file.").Default("./geofabrik.yml").Short('c').String()
	fNodownload = app.Flag("nodownload", "Do not download file (test only)").Short('n').Bool()
	fVerbose    = app.Flag("verbose", "Be verbose").Short('v').Bool()
	fQuiet      = app.Flag("quiet", "Be quiet").Short('q').Bool()
	fProgress   = app.Flag("progress", "Add a progress bar").Bool()

	list = app.Command("list", "Show elements available")
	lmd  = list.Flag("markdown", "generate list in Markdown format").Bool()

	download = app.Command("download", "Download element") //TODO : add d as command
	delement = download.Arg("element", "OSM element").Required().String()
	dosmBz2  = download.Flag("osm.bz2", "Download osm.bz2 if available").Short('B').Bool()
	dosmGz   = download.Flag("osm.gz", "Download osm.gz if available").Short('G').Bool()
	dshpZip  = download.Flag("shp.zip", "Download shp.zip if available").Short('S').Bool()
	dosmPbf  = download.Flag("osm.pbf", "Download osm.pbf (default)").Short('P').Bool()
	doshPbf  = download.Flag("osh.pbf", "Download osh.pbf").Short('H').Bool()
	dstate   = download.Flag("state", "Download state.txt file").Short('s').Bool()
	dpoly    = download.Flag("poly", "Download poly file").Short('p').Bool()
	dkml     = download.Flag("kml", "Download kml file").Short('k').Bool()
	dCheck   = download.Flag("check", "Control with checksum (default) Use --no-check to discard control").Default("true").Bool()

	generate = app.Command("generate", "Generate a new config file")
)

func listAllRegions(c *Config, format string) {
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

func checkService() bool {
	switch *fService {
	case "geofabrik":
		return true
	case "openstreetmap.fr":
		if strings.EqualFold(*fConfig, "./geofabrik.yml") {
			*fConfig = "./openstreetmap.fr.yml"
		}
		return true
	case "bbbike":
		if strings.EqualFold(*fConfig, "./geofabrik.yml") {
			*fConfig = "./bbbike.yml"
		}
		return true
	}
	return false
}

func catch(err error) {
	if err != nil {
		log.Fatalln(err.Error()) // Fatalln is better than Panic or Println
		// Println only log but dont do exit(1),
		// Panic add a lot of verbose detail for debug but it's too aggressive!
	}
}

func listCommand() {
	var format = ""
	if *lmd {
		format = "Markdown"
	}
	configPtr, err := loadConfig(*fConfig)
	catch(err)
	listAllRegions(configPtr, format)
}

func downloadCommand() {
	configPtr, err := loadConfig(*fConfig)
	catch(err)
	formatFile := getFormats()
	for _, format := range *formatFile {
		if ok, _, _ := isHashable(configPtr, format); *dCheck && ok {
			if fileExist(*delement + "." + format) {
				if !downloadChecksum(format) {
					if !*fQuiet {
						log.Println("Checksum mismatch, re-downloading", *delement+"."+format)
					}
					myElem, err := findElem(configPtr, *delement)
					catch(err)
					myURL, err := elem2URL(configPtr, myElem, format)
					catch(err)
					err = downloadFromURL(myURL, *delement+"."+format)
					catch(err)
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
				err = downloadFromURL(myURL, *delement+"."+format)
				catch(err)
				if !downloadChecksum(format) && !*fQuiet {
					log.Println("Checksum mismatch, please re-download", *delement+"."+format)
				}
			}
		} else {
			myElem, err := findElem(configPtr, *delement)
			catch(err)
			myURL, err := elem2URL(configPtr, myElem, format)
			catch(err)
			err = downloadFromURL(myURL, *delement+"."+format)
			catch(err)
		}
	}
}

func main() {

	app.Version(version) // Add version flag
	commands := kingpin.MustParse(app.Parse(os.Args[1:]))
	checkService()
	switch commands {
	case list.FullCommand():
		listCommand()
	case download.FullCommand():
		downloadCommand()
	case generate.FullCommand():
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
		defer func() {
			err := file.Close()
			catch(err)
		}()
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
		return strings.EqualFold(hash, filehash), nil
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
		if ok, _, _ := isHashable(configPtr, format); ok {
			myElem, err := findElem(configPtr, *delement)
			catch(err)
			myURL, err := elem2URL(configPtr, myElem, fhash)
			catch(err)
			err = downloadFromURL(myURL, *delement+"."+fhash)
			catch(err)
			if *fVerbose && !*fQuiet {
				log.Println("Hashing", *delement+"."+format)
			}
			hashed, err := hashFileMD5(*delement + "." + format)
			catch(err)
			if *fVerbose && !*fQuiet {
				log.Println("MD5 :", hashed)
			}
			ret, err := controlHash(*delement+"."+fhash, hashed)
			catch(err)
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
			log.Println("No checksum provided for", *delement+"."+format)
		}
	}
	return ret
}
