package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"

	"golang.org/x/net/proxy"
)

const progressMinimal = 512 * 1024 // Don't display progress bar if size < 512kb

func downloadFromURL(myURL string, fileName string) error {
	if *fVerbose && !*fQuiet {
		log.Println("Downloading", myURL, "to", fileName)
	}

	if !*fNodownload {

		transport := &http.Transport{}
		if *fProxyHTTP != "" {
			u, _ := url.Parse(myURL)
			//log.Println(u.Scheme +"://"+ *fProxyHTTP)
			proxyURL, err := url.Parse(u.Scheme + "://" + *fProxyHTTP)
			if *fProxyUser != "" && *fProxyPass != "" {
				proxyURL, err = url.Parse(u.Scheme + "://" + *fProxyUser + ":" + *fProxyPass + *fProxyHTTP)
			}
			if err != nil {
				return fmt.Errorf("Wrong proxy url, please use format proxy_address:port")
			}
			transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		}
		client := &http.Client{Transport: transport}
		if *fProxySock5 != "" {
			auth := proxy.Auth{User: *fProxyUser, Password: *fProxyPass}
			dialer, err := proxy.SOCKS5("tcp", *fProxySock5, &auth, proxy.Direct)
			if err != nil {
				return fmt.Errorf("Can't connect to the proxy: %v", err)
			}
			transport.Dial = dialer.Dial
		}
		response, err := client.Get(myURL)
		if err != nil {
			return fmt.Errorf("Error while downloading %s - %v", myURL, err)
		}
		if response.StatusCode != 200 {
			if response.StatusCode == 404 {
				return fmt.Errorf("Error while downloading %v, server return code %d\nPlease use 'download-geofabrik generate' to re-create your yml file", myURL, response.StatusCode)
			}
			return fmt.Errorf("Error while downloading %v, server return code %d", myURL, response.StatusCode)
		}
		defer response.Body.Close()

		// If no error, create file
		// TODO: check file existence first with io.IsExist
		// and use a new flag (like f) to force overwrite
		flags := os.O_CREATE | os.O_WRONLY
		var f *os.File
		f, err = os.OpenFile(fileName, flags, 0666)
		if err != nil {
			return fmt.Errorf("Error while creating %s - %v", fileName, err)
		}
		defer f.Close()
		var output io.Writer
		output = f
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
