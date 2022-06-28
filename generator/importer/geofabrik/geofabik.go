package geofabrik

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
)

const (
	GeofabrikIndexURL = `https://download.geofabrik.de/index-v1-nogeom.json`
	GeofabrikBaseURL  = `https://download.geofabrik.de`
)

var GeofabrikFormatDefinition = formats.FormatDefinitions{
	formats.FormatOsmBz2: {ID: formats.FormatOsmBz2, Loc: "-latest.osm.bz2"},
	"osm.bz2.md5":        {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
	formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: "-latest.osm.pbf"},
	"osm.pbf.md5":        {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
	formats.FormatPoly:   {ID: formats.FormatPoly, Loc: ".poly"},
	formats.FormatKml:    {ID: formats.FormatKml, Loc: ".kml"},
	formats.FormatState:  {ID: formats.FormatState, Loc: "-updates/state.txt"},
	formats.FormatShpZip: {ID: formats.FormatShpZip, Loc: "-latest-free.shp.zip"},
}

type Index struct {
	Features []IndexElement `json:"features"`
}

type IndexElement struct {
	ElementProperties IndexElementProperties `json:"properties"`
}

type IndexElementProperties struct {
	Urls      map[string]string `json:"urls"`
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Parent    string            `json:"parent,omitempty"`
	Iso3166_1 []string          `json:"iso3166-1:alpha2,omitempty"`
	Iso3166_2 []string          `json:"iso3166-2,omitempty"`
}

// GetIndex download Index and Unmarshall the json
func GetIndex(myURL string) (*Index, error) {
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
		return nil, fmt.Errorf("error while downloading %s - %w", myURL, err)
	}

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("error while downloading %v, server return code %d\n"+
				"Please use '"+os.Args[0]+" generate' to re-create your yml file",
				myURL, response.StatusCode)
		}

		return nil, fmt.Errorf("error while downloading %v, server return code %d", myURL, response.StatusCode)
	}

	defer func() {
		if e := response.Body.Close(); e != nil {
			log.WithError(e).Fatal("can't close HTTP connection")
		}
	}()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body %w", err)
	}

	var mygeofabrikIndex Index

	err = json.Unmarshal(bodyBytes, &mygeofabrikIndex)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling response body %w", err)
	}

	return &mygeofabrikIndex, nil
}

func Convert(g *Index) (*config.Config, error) {
	c := &config.Config{
		Formats:       GeofabrikFormatDefinition,
		BaseURL:       GeofabrikBaseURL,
		Elements:      element.Slice{},
		ElementsMutex: &sync.RWMutex{},
	}

	for _, ge := range g.Features {
		var e element.Element

		if viper.GetBool("log") {
			log.Debugf("ID:%v", ge.ElementProperties.ID)
		}

		e.ID = ge.ElementProperties.ID
		e.Parent = ge.ElementProperties.Parent
		e.Name = ge.ElementProperties.Name

		for k := range ge.ElementProperties.Urls {
			switch k {
			case "pbf":
				e.Formats = append(e.Formats, formats.FormatOsmPbf, "osm.pbf.md5")
			case "bz2":
				e.Formats = append(e.Formats, formats.FormatOsmBz2, "osm.bz2.md5")
			case "shp":
				e.Formats = append(e.Formats, formats.FormatShpZip)
			case "history":
				e.Formats = append(e.Formats, formats.FormatOshPbf)
			}
		}

		e.Formats = append(e.Formats, formats.FormatPoly, formats.FormatKml, formats.FormatState)

		err := c.MergeElement(&e)
		if err != nil {
			return nil, fmt.Errorf("error while merging element %v - %w", e, err)
		}
	}

	return c, nil
}
