package app

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/text"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/spf13/viper"
)

func ConfigureLogging() {
	log.SetLevel(log.InfoLevel)
	log.SetHandler(cli.Default)

	if viper.GetBool(config.ViperQuiet) || viper.GetBool(config.ViperProgress) {
		log.SetLevel(log.ErrorLevel)
	}

	if viper.GetBool(config.ViperVerbose) {
		log.SetLevel(log.DebugLevel)
		log.SetHandler(text.Default)
	}
}
