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
	"github.com/spf13/viper"
)

const (
	progressMinimal = 512 * 1024 // Don't display progress bar if size < 512kb
	ErrFromURL      = "can't download element"
)

func FromURL(myURL, fileName string) error {
	log.Debugf("Downloading %s to %s", myURL, fileName)

	if !viper.GetBool("noDownload") {
		client := &http.Client{Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
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

		f, err := os.OpenFile(fileName, flags, 0o644) //nolint:gomnd // 0o644 is the default mode
		if err != nil {
			return fmt.Errorf("error while creating %s - %w", fileName, err)
		}

		defer func() {
			if e := f.Close(); e != nil {
				log.WithError(e).Fatal("can't close file")
			}
		}()

		var (
			output      io.Writer = f
			n           int64
			progressBar *pb.ProgressBar
		)

		if viper.GetBool("progress") && response.ContentLength > progressMinimal {
			progressBar = pb.Full.Start64(response.ContentLength)
			barReader := progressBar.NewProxyReader(response.Body)

			n, err = io.Copy(output, barReader)
			if err != nil {
				return fmt.Errorf("error while writing %s - %w", fileName, err)
			}

			defer progressBar.Finish()
		} else {
			n, err = io.Copy(output, response.Body)
			if err != nil {
				return fmt.Errorf("error while writing %s - %w", fileName, err)
			}
		}

		if progressBar != nil {
			progressBar.Finish() // Force finish
		}

		log.Infof("%s downloaded.", fileName)
		log.Debugf("%v bytes downloaded.", n)
	}

	return nil // Everything is ok
}
