package app

import (
	"strings"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/spf13/viper"
)

const (
	geofabrikService = "geofabrik"
	geofabrikParse   = "geofabrik-parse"
	openstreetmapFr  = "openstreetmap.fr"
	osmtodayService  = "osmtoday"
	bbbikeService    = "bbbike"
)

// serviceConfigMap maps services to their respective configuration files.
var serviceConfigMap = map[string]string{
	openstreetmapFr: "./openstreetmap.fr.yml",
	osmtodayService: "./osmtoday.yml",
	bbbikeService:   "./bbbike.yml",
}

// CheckService checks the service and sets the appropriate configuration file.
func CheckService() bool {
	service := viper.GetString(config.ViperService)
	if service == geofabrikService || service == geofabrikParse {
		return true
	}

	if configFile, exists := serviceConfigMap[service]; exists {
		SetConfigFile(configFile)

		return true
	}

	return false
}

// SetConfigFile sets the configuration file if the current config file is the default.
func SetConfigFile(configFile string) {
	if strings.EqualFold(viper.GetString(config.ViperConfig), defaultConfigFile) {
		viper.Set(config.ViperConfig, configFile)
	}
}
