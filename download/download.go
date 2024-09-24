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
)

func FromURL(myURL, fileName string) error { //nolint:cyclop // TODO: Refactor!
	log.Debugf("Downloading %s to %s", myURL, fileName)

	if !viper.GetBool("noDownload") { //nolint:nestif // TODO : Refactor?
		client := &http.Client{Transport: &http.Transport{ //nolint:exhaustruct // I'm lazy
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{ //nolint:exhaustruct // I'm lazy
				Timeout:   60 * time.Second, //nolint:gomnd // 60 seconds
				KeepAlive: 30 * time.Second, //nolint:gomnd // 30 seconds
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          0,
			IdleConnTimeout:       5 * time.Second,  //nolint:gomnd //  5 seconds
			TLSHandshakeTimeout:   10 * time.Second, //nolint:gomnd // 10 seconds
			ExpectContinueTimeout: 5 * time.Second,  //nolint:gomnd //  5 seconds
		}}

		response, err := client.Get(myURL)
		if err != nil {
			return fmt.Errorf("error while downloading %s - %w", myURL, err)
		}

		if response.StatusCode != http.StatusOK {
			if response.StatusCode == http.StatusNotFound {
				return fmt.Errorf("error while downloading %v, server return code %d\n"+
					"Please use '"+os.Args[0]+" generate' to re-create your yml file",
					myURL, response.StatusCode)
			}

			return fmt.Errorf("error while downloading %v, server return code %d", myURL, response.StatusCode)
		}

		defer func() {
			if e := response.Body.Close(); e != nil {
				log.WithError(e).Fatal("can't close HTTP connection")
			}
		}()

		// If no error, create file
		// TODO: check file existence first with io.IsExist
		// and use a new cmd flag (like f) to force overwrite
		flags := os.O_CREATE | os.O_WRONLY

		file, err := os.OpenFile(fileName, flags, 0o644) //nolint:gomnd // 0o644 is the default mode
		if err != nil {
			return fmt.Errorf("error while creating %s - %w", fileName, err)
		}

		defer func() {
			if e := file.Close(); e != nil {
				log.WithError(e).Fatal("can't close file")
			}
		}()

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

		if progressBar != nil {
			progressBar.Finish() // Force finish
		}

		log.Infof("%s downloaded.", fileName)
		log.Debugf("%v bytes downloaded.", currentProgress)
	}

	return nil // Everything is ok
}

func FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}

	return false
}

func DownloadFile(configPtr *config.Config, element, format, output string) {
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

func GetOutputFileName(configPtr *config.Config, element string, myFormat *formats.Format) string {
	myElem, err := config.FindElem(configPtr, element)
	if err != nil {
		log.WithError(err).Fatalf(config.ErrFindElem.Error(), element)
	}

	var extension string

	if myFormat.ToLoc != "" {
		extension = myFormat.ToLoc
	} else {
		extension = "." + myFormat.ID
	}

	return myElem.ID + extension
}

func DownloadChecksum(format string) bool {
	ret := false

	if viper.GetBool(config.ViperCheck) {
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

			if e := FromURL(myURL, viper.GetString("output_directory")+viper.GetString(config.ViperElement)+"."+fhash); e != nil {
				log.WithError(e).Fatal(ErrFromURL)
			}

			log.Infof("Hashing %s", viper.GetString(config.ViperOutputDirectory)+viper.GetString(config.ViperElement)+"."+format)

			hashed, err := HashFileMD5(viper.GetString(config.ViperOutputDirectory) + viper.GetString(config.ViperElement) + "." + format)
			if err != nil {
				log.WithError(err).Fatal("can't hash file")
			}

			log.Debugf("MD5 : %s", hashed)

			ret, err = ControlHash(viper.GetString(config.ViperOutputDirectory)+viper.GetString(config.ViperElement)+"."+fhash, hashed)
			if err != nil {
				log.WithError(err).Error("checksum error")
			}

			if ret {
				log.Infof("Checksum OK for %s", viper.GetString(config.ViperOutputDirectory)+viper.GetString(config.ViperElement)+"."+format)
			} else {
				log.Infof("Checksum MISMATCH for %s", viper.GetString(config.ViperOutputDirectory)+viper.GetString(config.ViperElement)+"."+format)
			}

			return ret
		}

		log.Warnf("No checksum provided for %s", viper.GetString(config.ViperOutputDirectory)+viper.GetString(config.ViperElement)+"."+format)
	}

	return ret
}
