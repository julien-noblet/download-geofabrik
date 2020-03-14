package generator

import (
	"io/ioutil"
	"path/filepath"

	"github.com/apex/log"
	pb "github.com/cheggaaa/pb/v3"
	"github.com/gocolly/colly"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/generator/importer/geofabrik"
	"github.com/julien-noblet/download-geofabrik/scrapper"
	"github.com/julien-noblet/download-geofabrik/scrapper/bbbike"
	geofabrikScrapper "github.com/julien-noblet/download-geofabrik/scrapper/geofabrik"
	"github.com/julien-noblet/download-geofabrik/scrapper/openstreetmapfr"
	"github.com/spf13/viper"
)

func write(c *config.Config, filename string) {
	out, _ := c.Generate()
	filename, _ = filepath.Abs(filename)

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.WithError(err).Fatal("can't write file")
	}

	log.Infof("%s generated.", filename)
}

// Generate main function
func Generate(configfile string) {
	var (
		bar *pb.ProgressBar
		s   scrapper.IScrapper
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

		write(c, configfile)

		return // Exit function!
	case "geofabrik-parse":
		s = geofabrikScrapper.GetDefault()
	case "openstreetmap.fr":
		s = openstreetmapfr.GetDefault()
	case "bbbike":
		s = bbbike.GetDefault()

	default:
		log.Error("service not recognized, please use one of geofabrik, openstreetmap.fr or bbbike")
	}

	if viper.GetBool("progress") {
		bar = pb.New(s.GetPB())
		bar.Start()
	}

	c := s.Collector()
	c.OnScraped(func(*colly.Response) {
		if viper.GetBool("progress") {
			bar.Increment()
		}
	})

	if err := c.Visit(s.GetStartURL()); err != nil {
		log.WithError(err).Error("can't get url")
	}

	c.Wait()
	write(s.GetConfig(), configfile)
}
