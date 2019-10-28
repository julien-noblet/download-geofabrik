package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	pb "github.com/cheggaaa/pb/v3"
	"github.com/gocolly/colly"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/scrapper"
	"github.com/julien-noblet/download-geofabrik/scrapper/bbbike"
	"github.com/julien-noblet/download-geofabrik/scrapper/geofabrik"
	"github.com/julien-noblet/download-geofabrik/scrapper/openstreetmapfr"
	"github.com/spf13/viper"
)

func write(c *config.Config, filename string) {
	out, _ := c.Generate()
	filename, _ = filepath.Abs(filename)

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Panicln(fmt.Errorf(" File error: %v ", err))
	}

	if !viper.GetBool("quiet") {
		log.Println(filename, " generated.")
	}
}

// Generate main function
func Generate(configfile string) {
	var (
		bar *pb.ProgressBar
		s   scrapper.IScrapper
	)

	switch viper.GetString("service") {
	case "geofabrik":
		s = geofabrik.GetDefault()
	case "openstreetmap.fr":
		s = openstreetmapfr.GetDefault()
	case "bbbike":
		s = bbbike.GetDefault()

	default:
		log.Panicln(fmt.Errorf("service not recognized, please use one of geofabrik, openstreetmap.fr or bbbike"))
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
		log.Panicln(err)
	}

	c.Wait()
	write(s.GetConfig(), configfile)
}
