package main

import (
	"crypto/md5" //nolint:gosec // I only can get MD5sum files
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/text"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/download"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/generator"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

var version = "devel"

var ( // TODO: move from kingpin to cobra
	app      = kingpin.New("download-geofabrik", "A command-line tool for downloading OSM files.") //nolint:gochecknoglobals // global
	fService = app.Flag("service",                                                                 //nolint:gochecknoglobals // global
		"Can switch to another service. "+
			"You can use \"geofabrik\", \"openstreetmap.fr\" \"osmtoday\" or \"bbbike\". "+
			"It automatically change config file if -c is unused.").
		Default("geofabrik").String()
	fConfig     = app.Flag("config", "Set Config file.").Default("./geofabrik.yml").Short('c').String() //nolint:gochecknoglobals // global
	fNodownload = app.Flag("nodownload", "Do not download file (test only)").Short('n').Bool()          //nolint:gochecknoglobals // global
	fVerbose    = app.Flag("verbose", "Be verbose").Short('v').Bool()                                   //nolint:gochecknoglobals // global
	fQuiet      = app.Flag("quiet", "Be quiet").Short('q').Bool()                                       //nolint:gochecknoglobals // global
	fProgress   = app.Flag("progress", "Add a progress bar (implie quiet)").Bool()                      //nolint:gochecknoglobals // global

	list = app.Command("list", "Show elements available")                   //nolint:gochecknoglobals // global
	lmd  = list.Flag("markdown", "generate list in Markdown format").Bool() //nolint:gochecknoglobals // global
	// TODO : add dDownload as command.
	dDownload = app.Command("download", "Download element")                                                   //nolint:gochecknoglobals // global
	delement  = dDownload.Arg("element", "OSM element").Required().String()                                   //nolint:gochecknoglobals // global
	dosmBz2   = dDownload.Flag(formats.FormatOsmBz2, "Download osm.bz2 if available").Short('B').Bool()       //nolint:gochecknoglobals // global
	dosmGz    = dDownload.Flag(formats.FormatOsmGz, "Download osm.gz if available").Short('G').Bool()         //nolint:gochecknoglobals // global
	dshpZip   = dDownload.Flag(formats.FormatShpZip, "Download shp.zip if available").Short('S').Bool()       //nolint:gochecknoglobals // global
	dosmPbf   = dDownload.Flag(formats.FormatOsmPbf, "Download osm.pbf (default)").Short('P').Bool()          //nolint:gochecknoglobals // global
	doshPbf   = dDownload.Flag(formats.FormatOshPbf, "Download osh.pbf").Short('H').Bool()                    //nolint:gochecknoglobals // global
	dstate    = dDownload.Flag(formats.FormatState, "Download state.txt file").Short('s').Bool()              //nolint:gochecknoglobals // global
	dpoly     = dDownload.Flag(formats.FormatPoly, "Download poly file").Short('p').Bool()                    //nolint:gochecknoglobals // global
	dkml      = dDownload.Flag(formats.FormatKml, "Download kml file").Short('k').Bool()                      //nolint:gochecknoglobals // global
	dgeojson  = dDownload.Flag(formats.FormatGeoJSON, "Download geojson file").Short('g').Bool()              //nolint:gochecknoglobals // global
	dCheck    = dDownload.Flag("check", "Control with checksum (default) Use --no-check to discard control"). //nolint:gochecknoglobals // global
			Default("true").Bool()
	dOutputDir = dDownload.Flag( //nolint:gochecknoglobals // global
		"output_directory",
		"Set output directory, you can use also OUTPUT_DIR env variable",
	).
		Short('d').String()

	generate = app.Command("generate", "Generate a new config file") //nolint:gochecknoglobals // global
)

func listAllRegions(configuration *config.Config, format string) {
	table := tablewriter.NewWriter(os.Stdout)
	keys := make(sort.StringSlice, len(configuration.Elements))
	i := 0

	for k := range configuration.Elements {
		keys[i] = k
		i++
	}

	keys.Sort()
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ShortName", "Is in", "Long Name", "formats"})

	if format == "Markdown" {
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
	}

	for _, item := range keys {
		table.Append([]string{
			item,
			configuration.Elements[configuration.Elements[item].Parent].Name,
			configuration.Elements[item].Name,
			formats.MiniFormats(configuration.Elements[item].Formats),
		})
	}

	table.Render()
	fmt.Printf("Total elements: %#v\n", len(configuration.Elements))
}

func checkService() bool {
	switch *fService {
	case "geofabrik":
		return true
	case "geofabrik-parse":
		return true
	case "openstreetmap.fr":
		if strings.EqualFold(*fConfig, "./geofabrik.yml") {
			*fConfig = "./openstreetmap.fr.yml"
		}

		return true
	case "osmtoday":
		if strings.EqualFold(*fConfig, "./geofabrik.yml") {
			*fConfig = "./osmtoday.yml"
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

func listCommand() {
	format := ""

	configPtr, err := config.LoadConfig(*fConfig)
	if err != nil {
		log.WithError(err).Fatal(config.ErrLoadConfig)
	}

	if *lmd {
		format = "Markdown"
	}

	listAllRegions(configPtr, format)
}

func downloadFile(configPtr *config.Config, element, format, output string) {
	format = configPtr.Formats[format].Ext
	myElem, err := config.FindElem(configPtr, element)
	if err != nil {
		log.WithError(err).Fatalf(config.ErrFindElem, element)
	}

	myURL, err := config.Elem2URL(configPtr, myElem, format)
	if err != nil {
		log.WithError(err).Fatal(config.ErrElem2URL)
	}

	err = download.FromURL(myURL, output)
	if err != nil {
		log.WithError(err).Fatal(download.ErrFromURL)
	}
}

func downloadCommand() {
	formatFile := formats.GetFormats()

	configPtr, err := config.LoadConfig(*fConfig)
	if err != nil {
		log.WithError(err).Fatal(config.ErrLoadConfig)
	}

	r := regexp.MustCompile(`.*[\\/]?([A-Za-z_-]*)$`) // Trick for handle / in name
	filename := r.FindStringSubmatch(*dOutputDir + *delement)[0]

	for _, format := range *formatFile {
		myFormat := configPtr.Formats[format]
		if ok, _, _ := config.IsHashable(configPtr, myFormat.Ext); *dCheck && ok { //nolint:nestif // TODO : Refactor?
			if fileExist(*dOutputDir + *delement + "." + myFormat.Ext) {
				if !downloadChecksum(myFormat.Ext) {
					log.Infof("Checksum mismatch, re-downloading %v", *dOutputDir+filename+"."+myFormat.Ext)
					downloadFile(configPtr, *delement, myFormat.Ext, *dOutputDir+filename+"."+myFormat.Ext)
					downloadChecksum(myFormat.Ext)
				} else {
					log.Info("Checksum match, no download!")
				}
			} else {
				downloadFile(configPtr, *delement, myFormat.Ext, *dOutputDir+filename+"."+myFormat.Ext)
				if !downloadChecksum(myFormat.Ext) {
					log.Warnf("Checksum mismatch, please re-download %s", *dOutputDir+filename+"."+myFormat.Ext)
				}
			}
		} else {
			downloadFile(configPtr, *delement, myFormat.Ext, *dOutputDir+filename+"."+myFormat.Ext)
		}
	}
}

func configureBool(flag *bool, name string) {
	viper.Set(name, false)

	if *flag {
		viper.Set(name, true)
	}
}

func main() {
	app.Version(version) // Add version flag
	commands := kingpin.MustParse(app.Parse(os.Args[1:]))

	log.SetLevel(log.InfoLevel)
	log.SetHandler(cli.Default)

	if *fQuiet || *fProgress {
		log.SetLevel(log.ErrorLevel)
	}

	if *fVerbose {
		log.SetLevel(log.DebugLevel)
		log.SetHandler(text.Default)
	}

	configureBool(fNodownload, "noDownload")
	configureBool(fProgress, "progress")

	configureBool(doshPbf, "doshPbf")
	configureBool(dosmPbf, "dosmPbf")
	configureBool(dosmGz, "dosmGz")
	configureBool(dosmBz2, "dosmBz2")
	configureBool(dshpZip, "dshpZip")
	configureBool(dstate, "dstate")
	configureBool(dpoly, "dpoly")
	configureBool(dkml, "dkml")

	viper.Set("service", fService)

	if *dOutputDir == "" {
		if *dOutputDir = os.Getenv("OUTPUT_DIR"); *dOutputDir == "" {
			var err error

			*dOutputDir, err = os.Getwd()
			if err != nil {
				panic(err)
			}
		}
	}

	*dOutputDir += string(os.PathSeparator)

	checkService()

	switch commands {
	case list.FullCommand():
		listCommand()
	case dDownload.FullCommand():
		downloadCommand()
	case generate.FullCommand():
		generator.Generate(*fConfig)
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
		hash := md5.New() //nolint:gosec // I use md5 to control with md5sum files

		file, err := os.Open(filePath)
		if err != nil {
			return returnMD5String, fmt.Errorf("can't open %s : %w", filePath, err)
		}

		defer func() {
			err := file.Close()
			if err != nil {
				log.WithError(err).Fatal("can't save file")
			}
		}()

		if _, err := io.Copy(hash, file); err != nil {
			return returnMD5String, fmt.Errorf("can't copy %s : %w", filePath, err)
		}

		hashInBytes := hash.Sum(nil)[:16]
		returnMD5String = hex.EncodeToString(hashInBytes)

		return returnMD5String, nil
	}

	return returnMD5String, nil
}

func controlHash(hashfile, hash string) (bool, error) {
	if fileExist(hashfile) {
		file, err := os.ReadFile(hashfile)
		if err != nil {
			return false, fmt.Errorf("can't read %s : %w", hashfile, err)
		}

		filehash := strings.Split(string(file), " ")[0]

		log.Debugf("Hash from file :%s", filehash)

		return strings.EqualFold(hash, filehash), nil
	}

	return false, nil
}

func downloadChecksum(format string) bool {
	ret := false

	if *dCheck { //nolint:nestif // TODO : Refactor?
		hash := "md5"
		fhash := format + "." + hash

		configPtr, err := config.LoadConfig(*fConfig)
		if err != nil {
			log.WithError(err).Fatal(config.ErrLoadConfig)
		}

		if ok, _, _ := config.IsHashable(configPtr, format); ok {
			myElem, err := config.FindElem(configPtr, *delement)
			if err != nil {
				log.WithError(err).Fatalf(config.ErrFindElem, *delement)
			}

			myURL, err := config.Elem2URL(configPtr, myElem, fhash)
			if err != nil {
				log.WithError(err).Fatal(config.ErrElem2URL)
			}

			if e := download.FromURL(myURL, *dOutputDir+*delement+"."+fhash); e != nil {
				log.WithError(e).Fatal(download.ErrFromURL)
			}

			log.Infof("Hashing %s", *dOutputDir+*delement+"."+format)

			hashed, err := hashFileMD5(*dOutputDir + *delement + "." + format)
			if err != nil {
				log.WithError(err).Fatal("can't hash file")
			}

			log.Debugf("MD5 : %s", hashed)

			ret, err = controlHash(*dOutputDir+*delement+"."+fhash, hashed)
			if err != nil {
				log.WithError(err).Error("checksum error")
			}

			if ret {
				log.Infof("Checksum OK for %s", *dOutputDir+*delement+"."+format)
			} else {
				log.Infof("Checksum MISMATCH for %s", *dOutputDir+*delement+"."+format)
			}

			return ret
		}

		log.Warnf("No checksum provided for %s", *dOutputDir+*delement+"."+format)
	}

	return ret
}
