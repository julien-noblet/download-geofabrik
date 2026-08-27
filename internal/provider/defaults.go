package provider

import (
	"github.com/julien-noblet/download-geofabrik/internal/provider/bbbike"
	"github.com/julien-noblet/download-geofabrik/internal/provider/geo2day"
	"github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik"
	"github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr"
)

// RegisterDefaultProviders registers all built-in catalog providers.
func RegisterDefaultProviders() {
	Register(geofabrik.NewProvider())
	Register(openstreetmapfr.NewProvider())
	Register(bbbike.NewProvider())
	Register(geo2day.NewProvider())
}
