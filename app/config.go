package app

import (
	"os"

	"github.com/apex/log"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
)

func (a *App) ConfigureViper() {
	viper.Set(config.ViperConfig, *a.fConfig)
	viper.Set(config.ViperService, *a.fService)
	ConfigureBool(config.ViperNoDL, a.fNodownload)
	ConfigureBool(config.ViperProgress, a.fProgress)
	ConfigureBool(config.ViperVerbose, a.fVerbose)
	ConfigureBool(config.ViperQuiet, a.fQuiet)
	viper.Set(config.ViperElement, *a.delement)

	// Add other configurations here
	viper.Set(config.ViperOutputDirectory, *a.dOutputDir)
	ConfigureBool(config.ViperCheck, a.dCheck)
	ConfigureBool(config.ViperListFormatMarkdown, a.lmd)

	// Formats:
	ConfigureBool(formats.FormatOsmBz2, a.dosmBz2)
	ConfigureBool(formats.FormatOsmGz, a.dosmGz)
	ConfigureBool(formats.FormatShpZip, a.dshpZip)
	ConfigureBool(formats.FormatOsmPbf, a.dosmPbf)
	ConfigureBool(formats.FormatOshPbf, a.doshPbf)
	ConfigureBool(formats.FormatState, a.dstate)
	ConfigureBool(formats.FormatPoly, a.dpoly)
	ConfigureBool(formats.FormatKml, a.dkml)
	ConfigureBool(formats.FormatGeoJSON, a.dgeojson)
	ConfigureBool(formats.FormatGarminOSM, a.dgarmin)
	ConfigureBool(formats.FormatMapsforge, a.dmaps)
	ConfigureBool(formats.FormatMBTiles, a.dmbtiles)
	ConfigureBool(formats.FormatCSV, a.dcsv)
	ConfigureBool(formats.FormatGarminOnroad, a.dgarminonroad)
	ConfigureBool(formats.FormatGarminOntrail, a.dgarminontrail)
	ConfigureBool(formats.FormatGarminOpenTopo, a.dgarminopentopo)

}

func ConfigureBool(name string, flag *bool) {
	viper.Set(name, false)

	if *flag {
		viper.Set(name, true)
	}
}

func SetOutputDir() {
	outputDir := viper.GetString(config.ViperOutputDirectory)
	if outputDir == "" {
		if outputDir = os.Getenv("OUTPUT_DIR"); outputDir == "" {
			var err error

			outputDir, err = os.Getwd()
			if err != nil {
				log.WithError(err).Fatal("Failed to get current working directory")
			}
		}
	}

	outputDir += string(os.PathSeparator)
	viper.Set(config.ViperOutputDirectory, outputDir)
}
