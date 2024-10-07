package app

import (
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/spf13/viper"
)

const (
	geofabrikParse = "geofabrik-parse"
)

// serviceConfigMap maps services to their respective configuration files.
var serviceConfigMap = map[string]string{
	serviceGeofabrik: "./geofabrik.yml",
	geofabrikParse:   "./geofabrik.yml",
	serviceOSMFr:     "./openstreetmap.fr.yml",
	serviceOSMToday:  "./osmtoday.yml",
	serviceBBBike:    "./bbbike.yml",
}

// CheckService checks the service and sets the appropriate configuration file.
func CheckService() bool {
	service := viper.GetString(config.ViperService)

	if configFile, exists := serviceConfigMap[service]; exists {
		SetConfigFile(configFile)

		return true
	}

	return false
}

// SetConfigFile sets the configuration file if the current config file is the default.
func SetConfigFile(configFile string) {
	if viper.GetString(config.ViperConfig) == "" && configFile == "" {
		viper.Set(config.ViperConfig, defaultConfigFile)

		return
	}

	if configFile != "" {
		viper.Set(config.ViperConfig, configFile)
	}
}
