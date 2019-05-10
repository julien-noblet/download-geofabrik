package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"
)

const progressMinimal = 512 * 1024 // Don't display progress bar if size < 512kb

func downloadFromURL(myURL string, fileName string) error {
	if *fVerbose && !*fQuiet {
		log.Println("Downloading", myURL, "to", fileName)
	}

	if !*fNodownload {

		client := &http.Client{Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          0,
			IdleConnTimeout:       5 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
		}}

		response, err := client.Get(myURL)
		if err != nil {
			return fmt.Errorf("Error while downloading %s - %v", myURL, err)
		}
		if response.StatusCode != http.StatusOK {
			if response.StatusCode == http.StatusNotFound {
				return fmt.Errorf("Error while downloading %v, server return code %d\nPlease use 'download-geofabrik generate' to re-create your yml file", myURL, response.StatusCode)
			}
			return fmt.Errorf("Error while downloading %v, server return code %d", myURL, response.StatusCode)
		}
		defer func() {
			catch(response.Body.Close())
		}()

		// If no error, create file
		// TODO: check file existence first with io.IsExist
		// and use a new cmd flag (like f) to force overwrite
		flags := os.O_CREATE | os.O_WRONLY
		var f *os.File
		f, err = os.OpenFile(fileName, flags, 0666)
		if err != nil {
			return fmt.Errorf("Error while creating %s - %v", fileName, err)
		}
		defer func() {
			catch(f.Close())
		}()
		var output io.Writer = f
		var n int64
		var progressBar *pb.ProgressBar
		if !*fQuiet && *fProgress && response.ContentLength > progressMinimal {

			progressBar = pb.New64(response.ContentLength)
			progressBar.SetUnits(pb.U_BYTES)
			progressBar.ShowTimeLeft = true
			progressBar.ShowSpeed = true
			progressBar.RefreshRate = time.Millisecond * 100 // reduce cpu usage, 100 seems to be a good value
			progressBar.Start()
			defer progressBar.Finish()
			output = io.MultiWriter(output, progressBar)
		}
		n, err = io.Copy(output, response.Body)
		if err != nil {
			return fmt.Errorf("Error while writing %s - %v", fileName, err)
		}
		if !*fQuiet {
			if progressBar != nil {
				progressBar.Finish() // Force finish
			}
			log.Println(fileName, "downloaded.")
			if *fVerbose {
				log.Println(n, "bytes downloaded.")
			}
		}
	}
	return nil // Everything is ok
}
