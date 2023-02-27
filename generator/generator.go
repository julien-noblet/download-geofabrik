package generator

import (
	"os"
	"path/filepath"

	"github.com/apex/log"
	pb "github.com/cheggaaa/pb/v3"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/generator/importer/geofabrik"
	"github.com/julien-noblet/download-geofabrik/scrapper"
	"github.com/julien-noblet/download-geofabrik/scrapper/bbbike"
	geofabrikScrapper "github.com/julien-noblet/download-geofabrik/scrapper/geofabrik"
	"github.com/julien-noblet/download-geofabrik/scrapper/openstreetmapfr"
	"github.com/spf13/viper"
)

func Write(c *config.Config, filename string) {
	out, _ := c.Generate()
	filename, _ = filepath.Abs(filename)

	if err := os.WriteFile(filename, out, 0o600); err != nil { //nolint:gomnd // 0o600 is the default mode for new files
		log.WithError(err).Fatal("can't write file")
	}

	log.Infof("%s generated.", filename)
}

// Generate main function.
func Generate(configfile string) { //nolint:cyclop // TODO : Refactor
	var (
		bar        *pb.ProgressBar
		myScrapper scrapper.IScrapper
	)

	switch viper.GetString("service") {
	case "geofabrik":
		index, err := geofabrik.GetIndex(geofabrik.GeofabrikIndexURL)
		if err != nil {
			log.WithError(err)
		}

		c, err := geofabrik.Convert(index)
		if err != nil {
			log.WithError(err)
		}

		Write(c, configfile)

		return // Exit function!
	case "geofabrik-parse":
		myScrapper = geofabrikScrapper.GetDefault()
	case "openstreetmap.fr":
		myScrapper = openstreetmapfr.GetDefault()
	case "bbbike":
		myScrapper = bbbike.GetDefault()

	default:
		log.Error("service not recognized, please use one of geofabrik, openstreetmap.fr or bbbike")
	}

	if viper.GetBool("progress") {
		bar = pb.New(myScrapper.GetPB())
		bar.Start()
	}

	collector := myScrapper.Collector()
	collector.OnScraped(func(*colly.Response) {
		if viper.GetBool("progress") {
			bar.Increment()
		}
	})

	if err := collector.Visit(myScrapper.GetStartURL()); err != nil {
		log.WithError(err).Error("can't get url")
	}

	collector.Wait()
	Write(myScrapper.GetConfig(), configfile)
}
