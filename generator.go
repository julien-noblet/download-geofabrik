package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"sync"
	"time"

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
	var myConfig *Config
	var scrapper IScrapper
	switch *fService {
	case "geofabrik":
		scrapper = &geofabrik
	case "openstreetmap.fr":
		myConfig = &Config{
			BaseURL:       "https://download.openstreetmap.fr/extracts",
			Formats:       make(map[string]format),
			Elements:      ElementSlice{},
			ElementsMutex: &sync.RWMutex{},
		}
		myConfig.Formats["osm.pbf"] = format{ID: "osm.pbf", Loc: "-latest.osm.pbf"}
		myConfig.Formats["poly"] = format{ID: "poly", Loc: ".poly", BasePath: "../polygons/"}
		myConfig.Formats["state"] = format{ID: "state", Loc: ".state.txt"}
		if *fProgress {
			bar = pb.New(openstreetmapFRPb)
			bar.Start()
		}
		//GenerateCrawler("https://download.openstreetmap.fr/", configfile, &myConfig)
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			colly.AllowedDomains("download.openstreetmap.fr"),
			colly.URLFilters(
				regexp.MustCompile(`https://download.openstreetmap.fr/([^cgi\-bin][^replication]\w.+|)`),
				//				regexp.MustCompile(`https://download.openstreetmap.fr/[a-z][^cgi\-bin^replication].+`),
				//regexp.MustCompile(`https://download.openstreetmap.fr/extracts/[a-z].+`),
				//regexp.MustCompile(`https://download.openstreetmap.fr/polygons/[a-z].+`),
			),
			colly.Async(true),
			colly.MaxDepth(0),
		)
		catch(c.Limit(&colly.LimitRule{
			DomainGlob:  "*",
			Parallelism: 20,
			RandomDelay: 5 * time.Second,
		}))

		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})
		c.OnHTML("a", func(e *colly.HTMLElement) {
			openstreetmapFRParse(e, myConfig, c)
		})
		c.OnScraped(func(*colly.Response) {
			if *fProgress {
				bar.Increment()
			}
		})
		catch(c.Visit("https://download.openstreetmap.fr/"))
		c.Wait()

	case "gislab":
		myConfig.BaseURL = "http://be.gis-lab.info/project/osm_dump"
		myConfig.Formats = make(map[string]format)
		myConfig.Formats["osm.pbf"] = format{ID: "osm.pbf", BaseURL: "http://data.gis-lab.info/osm_dump/dump", BasePath: "latest/", Loc: ".osm.pbf"}
		myConfig.Formats["osm.bz2"] = format{ID: "osm.bz2", BaseURL: "http://data.gis-lab.info/osm_dump/dump", BasePath: "latest/", Loc: ".osm.bz2"}
		myConfig.Formats["poly"] = format{ID: "poly", BaseURL: "https://raw.githubusercontent.com/nextgis/osmdump_poly/master", Loc: ".poly"}
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			colly.AllowedDomains("be.gis-lab.info"),
		)
		c.OnHTML("table", func(e *colly.HTMLElement) {
			catch(gislabParse(e, myConfig))
		})
		catch(c.Visit("http://be.gis-lab.info/project/osm_dump/iframe.php"))

		c.Wait()
		//GenerateCrawler("http://be.gis-lab.info/project/osm_dump/iframe.php", configfile, &myConfig)

	case "bbbike":
		scrapper = &bbbike

	default:
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
