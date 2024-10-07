package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/apex/log"
	pb "github.com/cheggaaa/pb/v3"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/generator/importer/geofabrik"
	"github.com/julien-noblet/download-geofabrik/scrapper"
	"github.com/julien-noblet/download-geofabrik/scrapper/bbbike"
	geofabrikScrapper "github.com/julien-noblet/download-geofabrik/scrapper/geofabrik"
	"github.com/julien-noblet/download-geofabrik/scrapper/openstreetmapfr"
	"github.com/julien-noblet/download-geofabrik/scrapper/osmtoday"
	"github.com/spf13/viper"
)

const (
	filePermission         = 0o600
	serviceGeofabrik       = "geofabrik"
	serviceGeofabrikParse  = "geofabrik-parse"
	serviceOpenStreetMapFR = "openstreetmap.fr"
	serviceOSMToday        = "osmtoday"
	serviceBBBike          = "bbbike"
)

// Write writes the generated configuration to a file.
func Write(c *config.Config, filename string) error {
	out, err := c.Generate()
	if err != nil {
		return fmt.Errorf("failed to generate config: %w", err)
	}

	filename, err = filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for filename: %w", err)
	}

	if err := os.WriteFile(filename, out, filePermission); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	log.Infof("%s generated.", filename)

	return nil
}

// Generate generates the configuration based on the specified service.
func Generate(configfile string) {
	service := viper.GetString("service")
	serviceHandlers := map[string]func(string){
		serviceGeofabrik:       handleGeofabrik,
		serviceGeofabrikParse:  handleScrapperService,
		serviceOpenStreetMapFR: handleScrapperService,
		serviceOSMToday:        handleScrapperService,
		serviceBBBike:          handleScrapperService,
	}

	if handler, exists := serviceHandlers[service]; exists {
		handler(configfile)
	} else {
		log.Errorf("service not recognized: %s, please use one of geofabrik, openstreetmap.fr, osmtoday or bbbike", service)
	}
}

// handleScrapperService handles the scrapper service based on the specified service.
func handleScrapperService(configfile string) {
	myScrapper := getScrapper(viper.GetString("service"))

	if viper.GetBool("progress") {
		handleProgress(myScrapper)
	} else {
		collector := myScrapper.Collector()
		visitAndWait(collector, myScrapper.GetStartURL())
	}

	myconfig := myScrapper.GetConfig()
	Cleanup(myconfig)

	if err := Write(myconfig, configfile); err != nil {
		log.WithError(err).Error("failed to write config")
	}
}

// getScrapper returns the appropriate scrapper based on the service.
func getScrapper(service string) scrapper.IScrapper { //nolint:ireturn // have to return the interface
	switch service {
	case serviceGeofabrikParse:
		return geofabrikScrapper.GetDefault()
	case serviceOpenStreetMapFR:
		return openstreetmapfr.GetDefault()
	case serviceOSMToday:
		return osmtoday.GetDefault()
	case serviceBBBike:
		return bbbike.GetDefault()
	default:
		log.Fatalf("unknown service: %s", service)

		return nil
	}
}

// handleProgress handles the progress bar for the scrapper.
func handleProgress(myScrapper scrapper.IScrapper) {
	bar := pb.New(myScrapper.GetPB())
	bar.Start()
	defer bar.Finish()

	collector := myScrapper.Collector()
	collector.OnScraped(func(*colly.Response) {
		bar.Increment()
	})
	visitAndWait(collector, myScrapper.GetStartURL())
}

// visitAndWait visits the URL and waits for the collector to finish.
func visitAndWait(collector *colly.Collector, url string) {
	if err := collector.Visit(url); err != nil {
		log.WithError(err).Error("can't get url")

		return
	}

	collector.Wait()
}

// handleGeofabrik handles the Geofabrik service.
func handleGeofabrik(configfile string) {
	index, err := geofabrik.GetIndex(geofabrik.GeofabrikIndexURL)
	if err != nil {
		log.WithError(err).Error("failed to get geofabrik index")

		return
	}

	myConfig, err := geofabrik.Convert(index)
	if err != nil {
		log.WithError(err).Error("failed to convert geofabrik index")

		return
	}

	Cleanup(myConfig)

	if err := Write(myConfig, configfile); err != nil {
		log.WithError(err).Error("failed to write config")
	}
}

// Cleanup sorts the formats in the configuration elements.
func Cleanup(c *config.Config) {
	for _, elem := range c.Elements {
		slices.Sort(elem.Formats)
	}
}
