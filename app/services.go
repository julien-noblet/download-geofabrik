package app

import (
	"strings"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/spf13/viper"
)

func CheckService() bool {
	switch viper.GetString(config.ViperService) {
	case "geofabrik":
		return true
	case "geofabrik-parse":
		return true
	case "openstreetmap.fr":
		if strings.EqualFold(viper.GetString(config.ViperConfig), "./geofabrik.yml") {
			viper.Set(config.ViperConfig, "./openstreetmap.fr.yml")
		}

		return true
	case "osmtoday":
		if strings.EqualFold(viper.GetString(config.ViperConfig), "./geofabrik.yml") {
			viper.Set(config.ViperConfig, "./osmtoday.yml")
		}

		return true
	case "bbbike":
		if strings.EqualFold(viper.GetString(config.ViperConfig), "./geofabrik.yml") {
			viper.Set(config.ViperConfig, "./bbbike.yml")
		}

		return true
	}

	return false
}
