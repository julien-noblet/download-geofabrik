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
	viper.Set(config.ViperConfig, *a.FConfig)
	viper.Set(config.ViperService, *a.FService)
	SetViperBool(config.ViperNoDL, a.FNodownload)
	SetViperBool(config.ViperProgress, a.FProgress)
	SetViperBool(config.ViperVerbose, a.FVerbose)
	SetViperBool(config.ViperQuiet, a.FQuiet)
	viper.Set(config.ViperElement, *a.Delement)

	// Add other configurations here
	viper.Set(config.ViperOutputDirectory, *a.DOutputDir)
	SetViperBool(config.ViperCheck, a.DCheck)
	SetViperBool(config.ViperListFormatMarkdown, a.Lmd)

	// Formats:
	SetViperBool(formats.FormatOsmBz2, a.DosmBz2)
	SetViperBool(formats.FormatOsmGz, a.DosmGz)
	SetViperBool(formats.FormatShpZip, a.DshpZip)
	SetViperBool(formats.FormatOsmPbf, a.DosmPbf)
	SetViperBool(formats.FormatOshPbf, a.DoshPbf)
	SetViperBool(formats.FormatState, a.Dstate)
	SetViperBool(formats.FormatPoly, a.Dpoly)
	SetViperBool(formats.FormatKml, a.Dkml)
	SetViperBool(formats.FormatGeoJSON, a.Dgeojson)
	SetViperBool(formats.FormatGarminOSM, a.Dgarmin)
	SetViperBool(formats.FormatMapsforge, a.Dmaps)
	SetViperBool(formats.FormatMBTiles, a.Dmbtiles)
	SetViperBool(formats.FormatCSV, a.Dcsv)
	SetViperBool(formats.FormatGarminOnroad, a.Dgarminonroad)
	SetViperBool(formats.FormatGarminOntrail, a.Dgarminontrail)
	SetViperBool(formats.FormatGarminOpenTopo, a.Dgarminopentopo)
}

// SetViperBool sets a boolean flag in Viper configuration.
func SetViperBool(name string, flag *bool) {
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
