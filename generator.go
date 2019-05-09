package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gocolly/colly"
	pb "gopkg.in/cheggaaa/pb.v1"
)

func write(config *Config, filename string) {
	out, _ := config.Generate()
	filename, _ = filepath.Abs(filename)
	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		catch(fmt.Errorf(" File error: %v ", err))
	}
	if !*fQuiet {
		log.Println(filename, " generated.")
	}
}

//Generate main function
func Generate(configfile string) {
	var bar *pb.ProgressBar
	var scrapper IScrapper
	switch *fService {
	case "geofabrik":
		scrapper = &geofabrik
	case "openstreetmap.fr":

		scrapper = &openstreetmapFR

	case "bbbike":
		scrapper = &bbbike

		catch(fmt.Errorf("Service not reconized, please use one of geofabrik, openstreetmap.fr or gislab"))
	}
	if *fProgress {
		bar = pb.New(scrapper.GetPB())
		bar.Start()
	}
	c := scrapper.Collector()
	/*c.WithTransport(&http.Transport{
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
	})*/

	c.OnScraped(func(*colly.Response) {
		if *fProgress {
			bar.Increment()
		}
	})
	catch(c.Visit(scrapper.GetStartURL()))
	c.Wait()
	write(scrapper.GetConfig(), configfile)
}
