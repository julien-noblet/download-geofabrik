package app

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/text"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/spf13/viper"
)

const (
	defaultLogLevel = log.InfoLevel
	quietLogLevel   = log.ErrorLevel
	verboseLogLevel = log.DebugLevel
)

// ConfigureLogging sets up the logging configuration based on the application flags.
func ConfigureLogging() {
	SetDefaultLogging()

	if IsQuietMode() {
		SetQuietLogging()
	}

	if IsVerboseMode() {
		SetVerboseLogging()
	}
}

// SetDefaultLogging sets the default logging configuration.
func SetDefaultLogging() {
	log.SetLevel(defaultLogLevel)
	log.SetHandler(cli.Default)
}

// SetQuietLogging sets the logging configuration for quiet mode.
func SetQuietLogging() {
	log.SetLevel(quietLogLevel)
}

// SetVerboseLogging sets the logging configuration for verbose mode.
func SetVerboseLogging() {
	log.SetLevel(verboseLogLevel)
	log.SetHandler(text.Default)
}

// IsQuietMode checks if the application is in quiet mode.
func IsQuietMode() bool {
	return viper.GetBool(config.ViperQuiet) || viper.GetBool(config.ViperProgress)
}

// IsVerboseMode checks if the application is in verbose mode.
func IsVerboseMode() bool {
	return viper.GetBool(config.ViperVerbose)
}
