package app

import (
	"os"

	"github.com/apex/log"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
)

const (
	defaultOutputDirEnv  = "OUTPUT_DIR"
	defaultPathSeparator = string(os.PathSeparator)
)

// ConfigureViper sets up the Viper configuration based on the application flags.
func (a *App) ConfigureViper() {
	viper.Set(config.ViperConfig, *a.fConfig)
	viper.Set(config.ViperService, *a.fService)
	setViperBool(config.ViperNoDL, a.fNodownload)
	setViperBool(config.ViperProgress, a.fProgress)
	setViperBool(config.ViperVerbose, a.fVerbose)
	setViperBool(config.ViperQuiet, a.fQuiet)
	viper.Set(config.ViperElement, *a.delement)

	// Add other configurations here
	viper.Set(config.ViperOutputDirectory, *a.dOutputDir)
	setViperBool(config.ViperCheck, a.dCheck)
	setViperBool(config.ViperListFormatMarkdown, a.lmd)

	// Formats:
	setViperBool(formats.FormatOsmBz2, a.dosmBz2)
	setViperBool(formats.FormatOsmGz, a.dosmGz)
	setViperBool(formats.FormatShpZip, a.dshpZip)
	setViperBool(formats.FormatOsmPbf, a.dosmPbf)
	setViperBool(formats.FormatOshPbf, a.doshPbf)
	setViperBool(formats.FormatState, a.dstate)
	setViperBool(formats.FormatPoly, a.dpoly)
	setViperBool(formats.FormatKml, a.dkml)
	setViperBool(formats.FormatGeoJSON, a.dgeojson)
	setViperBool(formats.FormatGarminOSM, a.dgarmin)
	setViperBool(formats.FormatMapsforge, a.dmaps)
	setViperBool(formats.FormatMBTiles, a.dmbtiles)
	setViperBool(formats.FormatCSV, a.dcsv)
	setViperBool(formats.FormatGarminOnroad, a.dgarminonroad)
	setViperBool(formats.FormatGarminOntrail, a.dgarminontrail)
	setViperBool(formats.FormatGarminOpenTopo, a.dgarminopentopo)
}

// setViperBool sets a boolean flag in Viper configuration.
func setViperBool(name string, flag *bool) {
	viper.Set(name, false)

	if *flag {
		viper.Set(name, true)
	}
}

// SetOutputDir sets the output directory in Viper configuration.
func SetOutputDir() {
	outputDir := viper.GetString(config.ViperOutputDirectory)
	if outputDir == "" {
		if outputDir = os.Getenv(defaultOutputDirEnv); outputDir == "" {
			var err error

			outputDir, err = os.Getwd()
			if err != nil {
				log.WithError(err).Fatal("Failed to get current working directory")
			}
		}
	}

	outputDir += defaultPathSeparator
	viper.Set(config.ViperOutputDirectory, outputDir)
}
