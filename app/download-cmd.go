package app

import (
	"regexp"

	"github.com/apex/log"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/download"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
)

const (
	viperConfigKey          = config.ViperConfig
	viperOutputDirectoryKey = config.ViperOutputDirectory
	viperElementKey         = config.ViperElement
	viperCheckKey           = config.ViperCheck
)

// DownloadCommand handles the download command logic.
func DownloadCommand() {
	formatFile := formats.GetFormats()

	configPtr, err := config.LoadConfig(viper.GetString(viperConfigKey))
	if err != nil {
		log.WithError(err).Fatal(config.ErrLoadConfig)
	}

	filename := GetFilename(viper.GetString(viperOutputDirectoryKey), viper.GetString(viperElementKey))

	for _, format := range formatFile {
		myFormat := configPtr.Formats[format]
		outputFilePath := viper.GetString(viperOutputDirectoryKey) + filename + "." + myFormat.ID

		if viper.GetBool(viperCheckKey) && IsHashable(configPtr, myFormat.ID) {
			HandleHashableFormat(configPtr, myFormat.ID, outputFilePath)
		} else {
			DownloadFile(configPtr, myFormat.ID, outputFilePath)
		}
	}
}

// GetFilename extracts the filename from the given directory and element.
func GetFilename(outputDir, element string) string {
	r := regexp.MustCompile(`.*[\\/]?([A-Za-z_-]*)$`)

	return r.FindStringSubmatch(outputDir + element)[0]
}

// IsHashable checks if the format is hashable.
func IsHashable(configPtr *config.Config, formatID string) bool {
	ok, _, _ := config.IsHashable(configPtr, formatID)

	return ok
}

// HandleHashableFormat handles the download logic for hashable formats.
func HandleHashableFormat(configPtr *config.Config, formatID, outputFilePath string) {
	if download.FileExist(outputFilePath) {
		if !download.Checksum(formatID) {
			log.Infof("Checksum mismatch, re-downloading %v", outputFilePath)
			DownloadFile(configPtr, formatID, outputFilePath)
			download.Checksum(formatID)
		} else {
			log.Info("Checksum match, no download!")
		}
	} else {
		DownloadFile(configPtr, formatID, outputFilePath)

		if !download.Checksum(formatID) {
			log.Warnf("Checksum mismatch, please re-download %s", outputFilePath)
		}
	}
}

// DownloadFile handles the file download logic.
func DownloadFile(configPtr *config.Config, formatID, outputFilePath string) {
	err := download.File(
		configPtr,
		viper.GetString(viperElementKey),
		formatID,
		outputFilePath,
	)
	if err != nil {
		log.WithError(err).Fatal("Download error")
	}
}
