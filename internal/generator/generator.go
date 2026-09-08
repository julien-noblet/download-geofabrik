package generator

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"

	"github.com/julien-noblet/download-geofabrik/internal/provider"
)

const (
	ServiceGeofabrik       = "geofabrik"
	ServiceGeofabrikParse  = "geofabrik-parse"
	ServiceOpenStreetMapFR = "openstreetmap.fr"
	ServiceGeo2Day         = "geo2day"
	ServiceBBBike          = "bbbike"
	ServiceMovisda         = "movisda"
	ServiceOSMCH           = "planet.osm.ch"
	ServiceOSMKewlLu       = "osm.kewl.lu"
	ServiceOSMFitVutbr     = "osm.fit.vutbr.cz"
	ServiceOSMIt           = "osmit-estratti"
	ServiceOSMTW           = "osm.kcwu.csie.org"
)

var ErrUnknownService = errors.New("unknown service")

// Generate generates the configuration catalog file for the specified service.
func Generate(ctx context.Context, service, configfile string) error {
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
