// 2015-2019 copyright Julien Noblet

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	pb "github.com/cheggaaa/pb/v3"
	"github.com/gocolly/colly"
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

	default:
		catch(fmt.Errorf("Service not reconized, please use one of geofabrik, openstreetmap.fr or bbbike"))
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
