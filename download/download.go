package download

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/apex/log"
	pb "github.com/cheggaaa/pb/v3"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
)

const (
	progressMinimal = 512 * 1024 // Don't display progress bar if size < 512kb
	ErrFromURL      = "can't download element"
	timeout         = 60 * time.Second
	keepAlive       = 30 * time.Second
	idleTimeout     = 5 * time.Second
	tlsTimeout      = 10 * time.Second
	continueTimeout = 5 * time.Second
	fileMode        = 0o644
)

// FromURL downloads a file from a URL to a specified file path.
func FromURL(myURL, fileName string) error {
	log.Debugf("Downloading %s to %s", myURL, fileName)

	if viper.GetBool("noDownload") {
		return nil
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: keepAlive,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          0,
			IdleConnTimeout:       idleTimeout,
			TLSHandshakeTimeout:   tlsTimeout,
			ExpectContinueTimeout: continueTimeout,
		},
	}

	response, err := client.Get(myURL)
	if err != nil {
		return fmt.Errorf("error while downloading %s - %w", myURL, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error while downloading %v, server return code %d\n Please use '%s generate' to re-create your yml file", myURL, response.StatusCode, os.Args[0])
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, fileMode)
	if err != nil {
		return fmt.Errorf("error while creating %s - %w", fileName, err)
	}
	defer file.Close()

	var (
		output          io.Writer = file
		currentProgress int64
		progressBar     *pb.ProgressBar
	)

	if viper.GetBool("progress") && response.ContentLength > progressMinimal {
		progressBar = pb.Full.Start64(response.ContentLength)
		barReader := progressBar.NewProxyReader(response.Body)

		currentProgress, err = io.Copy(output, barReader)
		if err != nil {
			return fmt.Errorf("error while writing %s - %w", fileName, err)
		}

		defer progressBar.Finish()
	} else {
		currentProgress, err = io.Copy(output, response.Body)
		if err != nil {
			return fmt.Errorf("error while writing %s - %w", fileName, err)
		}
	}

	log.Infof("%s downloaded.", fileName)
	log.Debugf("%v bytes downloaded.", currentProgress)

	return nil
}

// FileExist checks if a file exists at the given path.
func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil
}

// File downloads a file based on the configuration and element.
func File(configPtr *config.Config, element, format, output string) {
	format = configPtr.Formats[format].ID

	myElem, err := config.FindElem(configPtr, element)
	if err != nil {
		log.WithError(err).Fatalf(config.ErrFindElem.Error(), element)
	}

	myURL, err := config.Elem2URL(configPtr, myElem, format)
	if err != nil {
		log.WithError(err).Fatal(config.ErrElem2URL)
	}

	err = FromURL(myURL, output)
	if err != nil {
		log.WithError(err).Fatal(ErrFromURL)
	}
}

// GetOutputFileName generates the output file name based on the element and format.
func GetOutputFileName(configPtr *config.Config, element string, myFormat *formats.Format) string {
	myElem, err := config.FindElem(configPtr, element)
	if err != nil {
		log.WithError(err).Fatalf(config.ErrFindElem.Error(), element)
	}

	extension := myFormat.ToLoc
	if extension == "" {
		extension = "." + myFormat.ID
	}

	return myElem.ID + extension
}

// Checksum downloads and verifies the checksum of a file.
func Checksum(format string) bool {
	if !viper.GetBool(config.ViperCheck) {
		return false
	}

	hash := "md5"
	fhash := format + "." + hash

	configPtr, err := config.LoadConfig(viper.GetString(config.ViperConfig))
	if err != nil {
		log.WithError(err).Fatal(config.ErrLoadConfig)
	}

	if ok, _, _ := config.IsHashable(configPtr, format); ok {
		myElem, err := config.FindElem(configPtr, viper.GetString(config.ViperElement))
		if err != nil {
			log.WithError(err).Fatalf(config.ErrFindElem.Error(), viper.GetString(config.ViperElement))
		}

		myURL, err := config.Elem2URL(configPtr, myElem, fhash)
		if err != nil {
			log.WithError(err).Fatal(config.ErrElem2URL)
		}

		outputPath := viper.GetString("output_directory") + viper.GetString(config.ViperElement) + "." + fhash
		if e := FromURL(myURL, outputPath); e != nil {
			log.WithError(e).Fatal(ErrFromURL)
		}

		return VerifyChecksum(outputPath, format)
	}

	log.Warnf("No checksum provided for %s", viper.GetString(config.ViperOutputDirectory)+viper.GetString(config.ViperElement)+"."+format)
	return false
}
