package generator

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"

	pb "github.com/cheggaaa/pb/v3"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/generator/importer/geofabrik"
	"github.com/julien-noblet/download-geofabrik/internal/scrapper"
	"github.com/julien-noblet/download-geofabrik/internal/scrapper/bbbike"
	geofabrikScrapper "github.com/julien-noblet/download-geofabrik/internal/scrapper/geofabrik"
	"github.com/julien-noblet/download-geofabrik/internal/scrapper/openstreetmapfr"
	"github.com/julien-noblet/download-geofabrik/internal/scrapper/osmtoday"
)

const (
	filePermission         = 0o600
	ServiceGeofabrik       = "geofabrik"
	ServiceGeofabrikParse  = "geofabrik-parse"
	ServiceOpenStreetMapFR = "openstreetmap.fr"
	ServiceOSMToday        = "osmtoday"
	ServiceBBBike          = "bbbike"
)

var ErrUnknownService = errors.New("unknown service")

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

	slog.Info("Generated config file", "file", filename)

	return nil
}

// Generate generates the configuration based on the specified service.
func Generate(service string, progress bool, configfile string) error {
	// The original `Generate` function used a map of handlers.
	// With the refactoring, `PerformGenerate` now encapsulates the logic
	// for all services, including the distinction between Geofabrik and
	// the scrapper-based services.
	// Therefore, `Generate` can directly call `PerformGenerate`.
	return PerformGenerate(service, progress, configfile)
}

// PerformGenerate handles the generation logic for all services.
func PerformGenerate(service string, progress bool, configfile string) error {
	var myScrapper scrapper.IScrapper

	switch service {
	case ServiceGeofabrik:
		return handleGeofabrik(configfile, progress)
	case ServiceGeofabrikParse:
		myScrapper = geofabrikScrapper.GetDefault()
	case ServiceOpenStreetMapFR:
		myScrapper = openstreetmapfr.GetDefault()
	case ServiceOSMToday:
		myScrapper = osmtoday.GetDefault()
	case ServiceBBBike:
		myScrapper = bbbike.GetDefault()
	default:
		return fmt.Errorf("%w: %s", ErrUnknownService, service)
	}

	if progress {
		handleProgress(myScrapper)
	} else {
		collector := myScrapper.Collector()
		visitAndWait(collector, myScrapper.GetStartURL())
	}

	myconfig := myScrapper.GetConfig()
	Cleanup(myconfig)

	if err := Write(myconfig, configfile); err != nil {
		slog.Error("Failed to write config", "error", err)

		return err
	}

	return nil
}

// handleGeofabrik handles the Geofabrik service.
func handleGeofabrik(configfile string, _ bool) error {
	index, err := geofabrik.GetIndex(geofabrik.GeofabrikIndexURL)
	if err != nil {
		slog.Error("Failed to get geofabrik index", "error", err)

		return fmt.Errorf("failed to get index: %w", err)
	}

	myConfig, err := geofabrik.Convert(index)
	if err != nil {
		slog.Error("Failed to convert geofabrik index", "error", err)

		return fmt.Errorf("failed to convert index: %w", err)
	}

	Cleanup(myConfig)

	if err := Write(myConfig, configfile); err != nil {
		slog.Error("Failed to write config", "error", err)

		return err
	}

	return nil
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
		slog.Error("Can't get url", "error", err)

		return
	}

	collector.Wait()
}

// Cleanup sorts the formats in the configuration elements.
func Cleanup(c *config.Config) {
	for _, elem := range c.Elements {
		slices.Sort(elem.Formats)
	}
}
