package app

import (
	"regexp"

	"github.com/apex/log"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/download"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
)

func DownloadCommand() {
	formatFile := formats.GetFormats()

	configPtr, err := config.LoadConfig(viper.GetString(config.ViperConfig))
	if err != nil {
		log.WithError(err).Fatal(config.ErrLoadConfig)
	}

	r := regexp.MustCompile(`.*[\\/]?([A-Za-z_-]*)$`) // Trick for handle / in name
	filename := r.FindStringSubmatch(viper.GetString(config.ViperOutputDirectory) + viper.GetString(config.ViperElement))[0]

	for _, format := range *formatFile {
		myFormat := configPtr.Formats[format]
		if ok, _, _ := config.IsHashable(configPtr, myFormat.ID); viper.GetBool(config.ViperCheck) && ok { //nolint:nestif // TODO : Refactor?
			if download.FileExist(viper.GetString(config.ViperOutputDirectory) + viper.GetString(config.ViperElement) + "." + myFormat.ID) {
				if !download.DownloadChecksum(myFormat.ID) {
					log.Infof("Checksum mismatch, re-downloading %v", viper.GetString(config.ViperOutputDirectory)+filename+"."+myFormat.ID)
					download.DownloadFile(configPtr, viper.GetString(config.ViperElement), myFormat.ID, viper.GetString(config.ViperOutputDirectory)+filename+"."+myFormat.ID)
					download.DownloadChecksum(myFormat.ID)
				} else {
					log.Info("Checksum match, no download!")
				}
			} else {
				download.DownloadFile(
					configPtr,
					viper.GetString(config.ViperElement),
					myFormat.ID,
					viper.GetString(config.ViperOutputDirectory)+download.GetOutputFileName(configPtr, viper.GetString(config.ViperElement), &myFormat),
				)

				if !download.DownloadChecksum(myFormat.ID) {
					log.Warnf(
						"Checksum mismatch, please re-download %s",
						viper.GetString(config.ViperOutputDirectory)+download.GetOutputFileName(configPtr, viper.GetString(config.ViperElement), &myFormat),
					)
				}
			}
		} else {
			download.DownloadFile(
				configPtr,
				viper.GetString(config.ViperElement),
				myFormat.ID,
				viper.GetString(config.ViperOutputDirectory)+download.GetOutputFileName(configPtr, viper.GetString(config.ViperElement), &myFormat),
			)
		}
	}
}
