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

	c.OnScraped(func(*colly.Response) {
		if *fProgress {
			bar.Increment()
		}
	})
	catch(c.Visit(scrapper.GetStartURL()))
	c.Wait()
	write(scrapper.GetConfig(), configfile)
}
