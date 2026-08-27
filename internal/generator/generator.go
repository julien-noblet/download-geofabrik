package generator

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/provider"
)

const (
	filePermission         = 0o600
	ServiceGeofabrik       = "geofabrik"
	ServiceGeofabrikParse  = "geofabrik-parse"
	ServiceOpenStreetMapFR = "openstreetmap.fr"
	ServiceGeo2Day         = "geo2day"
	ServiceBBBike          = "bbbike"
)

var ErrUnknownService = errors.New("unknown service")

// Write writes the generated configuration to a file.
func Write(c *config.Config, filename string) error {
	out, err := c.Generate()
	if err != nil {
		return fmt.Errorf("failed to generate config: %w", err)
	}

	absFilename, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for filename: %w", err)
	}

	if err := os.WriteFile(absFilename, out, filePermission); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	slog.Info("Generated config file", "file", absFilename)

	return nil
}

// Generate generates the configuration based on the specified service.
func Generate(service string, progress bool, configfile string) error {
	return GenerateContext(context.Background(), service, progress, configfile)
}

// GenerateContext generates the configuration based on the specified service with context.
func GenerateContext(ctx context.Context, service string, progress bool, configfile string) error {
	return PerformGenerateContext(ctx, service, progress, configfile)
}

// PerformGenerate handles the generation logic using registered providers.
func PerformGenerate(service string, progress bool, configfile string) error {
	return PerformGenerateContext(context.Background(), service, progress, configfile)
}

// PerformGenerateContext handles the generation logic using registered providers with context.
func PerformGenerateContext(ctx context.Context, service string, _ bool, configfile string) error {
	provider.RegisterDefaultProviders()

	lookupService := service
	if service == ServiceGeofabrikParse {
		lookupService = ServiceGeofabrik
	}

	prov, err := provider.Get(lookupService)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrUnknownService, service)
	}

	slog.Info("Fetching catalog from provider", "service", service, "description", prov.Description())

	cat, err := prov.FetchCatalog(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch catalog from %s: %w", service, err)
	}

	// Sort formats within each element for deterministic YAML output
	for k, elem := range cat.Elements {
		slices.Sort(elem.Formats)
		cat.Elements[k] = elem
	}

	if err := cat.SaveFile(configfile); err != nil {
		return fmt.Errorf("failed to write config to %s: %w", configfile, err)
	}

	slog.Info("Generated config file", "file", configfile, "elements", len(cat.Elements))

	return nil
}

// Cleanup sorts the formats in the configuration elements.
func Cleanup(c *config.Config) {
	for _, elem := range c.Elements {
		slices.Sort(elem.Formats)
	}
}
