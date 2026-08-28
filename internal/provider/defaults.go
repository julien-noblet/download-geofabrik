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
)

// RegisterDefaultProviders registers all built-in catalog providers.
func RegisterDefaultProviders() {
	Register(geofabrik.NewProvider())
	Register(openstreetmapfr.NewProvider())
	Register(bbbike.NewProvider())
	Register(geo2day.NewProvider())
	Register(movisda.NewProvider())
	Register(osmch.NewProvider())
	Register(osmkewllu.NewProvider())
	Register(osmfitvutbr.NewProvider())
	Register(osmit.NewProvider())
	Register(osmtw.NewProvider())
}
