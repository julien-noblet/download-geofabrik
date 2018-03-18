package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/proxy"
)

func downloadFromURL(myURL *string, fileName string) {
	if *fVerbose {
		log.Println("Downloading", *myURL, "to", fileName)
	}

	if !*fNodownload {
		// TODO: check file existence first with io.IsExist
		output, err := os.Create(fileName)
		if err != nil {
			log.Fatalln(fmt.Errorf("Error while creating %s - %v", fileName, err))
			return
		}
		defer output.Close()
		transport := &http.Transport{}
		if *fProxyHTTP != "" {
			u, _ := url.Parse(*myURL)
			//log.Println(u.Scheme +"://"+ *fProxyHTTP)
			proxyURL, err := url.Parse(u.Scheme + "://" + *fProxyHTTP)
			if *fProxyUser != "" && *fProxyPass != "" {
				proxyURL, err = url.Parse(u.Scheme + "://" + *fProxyUser + ":" + *fProxyPass + *fProxyHTTP)
			}
			if err != nil {
				log.Fatalln(fmt.Errorf(" Wrong proxy url, please use format proxy_address:port"))
				return
			}
			transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		}
		client := &http.Client{Transport: transport}
		if *fProxySock5 != "" {
			auth := proxy.Auth{User: *fProxyUser, Password: *fProxyPass}
			dialer, err := proxy.SOCKS5("tcp", *fProxySock5, &auth, proxy.Direct)
			if err != nil {
				log.Fatalln(fmt.Errorf(" Can't connect to the proxy: %v", err))
				return
			}
			transport.Dial = dialer.Dial
		}
		response, err := client.Get(*myURL)
		if err != nil {
			log.Fatalln(fmt.Errorf("Error while downloading %s - %v", *myURL, err))
			return
		}
		defer response.Body.Close()

		n, err := io.Copy(output, response.Body)
		if err != nil {
			log.Fatalln(fmt.Errorf("Error while downloading %s - %v", *myURL, err))
			return
		}
		if !*fQuiet {
			log.Println(fileName, " downloaded.")
		}
		if *fVerbose {
			log.Println(n, " bytes downloaded.")
		}
	}
}
