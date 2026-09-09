package provider

import (
	"github.com/julien-noblet/download-geofabrik/internal/provider/bbbike"
	"github.com/julien-noblet/download-geofabrik/internal/provider/geo2day"
	"github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik"
	"github.com/julien-noblet/download-geofabrik/internal/provider/movisda"
	"github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr"
	"github.com/julien-noblet/download-geofabrik/internal/provider/osmch"
	"github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr"
	"github.com/julien-noblet/download-geofabrik/internal/provider/osmit"
	"github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu"
	"github.com/julien-noblet/download-geofabrik/internal/provider/osmtw"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

// RegisterDefaultProviders registers all built-in catalog providers.
func RegisterDefaultProviders() {
	Register(geofabrik.New())
	Register(openstreetmapfr.New())
	Register(bbbike.New())
	Register(geo2day.New())
	Register(movisda.New())
	Register(osmch.New())
	Register(osmkewllu.New())
	Register(osmfitvutbr.New())
	Register(osmit.New())
	Register(osmtw.New())
}

// AllDefaultFormats returns the default format definitions for all built-in providers.
func AllDefaultFormats() []catalog.FormatDefinitions {
	return []catalog.FormatDefinitions{
		geofabrik.DefaultFormats(),
		openstreetmapfr.DefaultFormats(),
		bbbike.DefaultFormats(),
		geo2day.DefaultFormats(),
		movisda.DefaultFormats(),
		osmch.DefaultFormats(),
		osmkewllu.DefaultFormats(),
		osmfitvutbr.DefaultFormats(),
		osmit.DefaultFormats(),
		osmtw.DefaultFormats(),
	}
}
