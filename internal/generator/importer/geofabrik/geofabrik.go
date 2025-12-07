package geofabrik

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
)

var (
	GeofabrikIndexURL = `https://download.geofabrik.de/index-v1-nogeom.json`
	GeofabrikBaseURL  = `https://download.geofabrik.de`

	// ErrDownload          = "error while downloading %v, server returned code %d\nPlease use '%s generate' to re-create your yml file %w"
	// ErrCreatingRequest   = "error while creating request for %s: %w"
	// ErrDownloading       = "error while downloading %s: %w"
	// ErrReadingResponse   = "error while reading response body: %w"
	// ErrUnmarshallingBody = "error while unmarshalling response body: %w"
	// ErrMergingElement    = "error while merging element %v: %w".

	TimeoutDuration       = 60 * time.Second
	KeepAliveDuration     = 30 * time.Second
	IdleConnTimeout       = 5 * time.Second
	TLSHandshakeTimeout   = 10 * time.Second
	ExpectContinueTimeout = 5 * time.Second
)

// HTTPClient is a reusable HTTP client.
var HTTPClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   TimeoutDuration,
			KeepAlive: KeepAliveDuration,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          0,
		IdleConnTimeout:       IdleConnTimeout,
		TLSHandshakeTimeout:   TLSHandshakeTimeout,
		ExpectContinueTimeout: ExpectContinueTimeout,
	},
}

// FormatDefinition returns a map of format definitions.
func FormatDefinition() formats.FormatDefinitions {
	return formats.FormatDefinitions{
		"osm.bz2.md5":         {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
		"osm.pbf.md5":         {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
		formats.FormatKml:     {ID: formats.FormatKml, Loc: ".kml"},
		formats.FormatMBTiles: {ID: formats.FormatMBTiles, Loc: "-latest-free.mbtiles.zip", ToLoc: "latest-free.mbtiles.zip"},
		formats.FormatOsmBz2:  {ID: formats.FormatOsmBz2, Loc: "-latest.osm.bz2"},
		formats.FormatOsmPbf:  {ID: formats.FormatOsmPbf, Loc: "-latest.osm.pbf"},
		formats.FormatPoly:    {ID: formats.FormatPoly, Loc: ".poly"},
		formats.FormatShpZip:  {ID: formats.FormatShpZip, Loc: "-shortbread-1.0.mbtiles"},
		formats.FormatState:   {ID: formats.FormatState, Loc: "-updates/state.txt"},
	}
}

// Index represents the structure of the Geofabrik index.
type Index struct {
	Features []IndexElement `json:"features"`
}

// IndexElement represents an element in the Geofabrik index.
type IndexElement struct {
	ElementProperties IndexElementProperties `json:"properties"`
}

// IndexElementProperties represents the properties of an index element.
type IndexElementProperties struct {
	Urls      map[string]string `json:"urls"`
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Parent    string            `json:"parent,omitempty"`
	Iso3166_1 []string          `json:"iso3166-1:alpha2,omitempty"` //nolint:tagliatelle // That's geofabrik's field name
	Iso3166_2 []string          `json:"iso3166-2,omitempty"`        //nolint:tagliatelle // That's geofabrik's field name
}

// GetIndex downloads the Geofabrik index and unmarshals the JSON response.
func GetIndex(url string) (*Index, error) {
	ctx, cancel := context.WithTimeout(context.Background(), TimeoutDuration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("error while creating request for %s: %w", url, err)
	}

	response, err := HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while downloading %s: %w", url, err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close response: %w", cerr)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, handleHTTPError(response, url)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %w", err)
	}

	var geofabrikIndex Index
	if err := json.Unmarshal(bodyBytes, &geofabrikIndex); err != nil {
		return nil, fmt.Errorf("error while unmarshalling response body: %w", err)
	}

	return &geofabrikIndex, nil
}

// handleHTTPError handles HTTP errors based on the status code.
func handleHTTPError(response *http.Response, url string) error {
	switch response.StatusCode {
	case http.StatusNotFound:
		return fmt.Errorf("error while downloading %s, server returned code %d\n"+
			"Please use '%s generate' to re-create your yml file %w", url, response.StatusCode, os.Args[0], http.ErrNoLocation)

	default:
		return fmt.Errorf("error while downloading %s, server returned code %d\n"+
			"Please use '%s generate' to re-create your yml file %w", url, response.StatusCode, os.Args[0], http.ErrNotSupported)
	}
}

// Convert converts the Geofabrik index to a config.Config object.
func Convert(index *Index) (*config.Config, error) {
	cfg := &config.Config{
		Formats:       FormatDefinition(),
		BaseURL:       GeofabrikBaseURL,
		Elements:      element.MapElement{},
		ElementsMutex: &sync.RWMutex{},
	}

	for _, feature := range index.Features {
		if err := processFeature(cfg, &feature); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

// processFeature processes a single feature from the Geofabrik index.
func processFeature(cfg *config.Config, feature *IndexElement) error {
	var elem element.Element

	slog.Debug("Processing feature", "ID", feature.ElementProperties.ID)

	elem.ID = feature.ElementProperties.ID
	elem.Parent = feature.ElementProperties.Parent
	elem.Name = feature.ElementProperties.Name

	elem.Formats = append(elem.Formats, getFormats(feature.ElementProperties.Urls)...)

	if err := cfg.MergeElement(&elem); err != nil {
		return fmt.Errorf("error while merging element %v: %w", elem, err)
	}

	return nil
}

// getFormats returns the formats based on the URLs.
func getFormats(urls map[string]string) []string {
	myFormats := []string{}

	for k := range urls {
		switch k {
		case "pbf":
			myFormats = append(myFormats, formats.FormatOsmPbf, "osm.pbf.md5")
		case "bz2":
			myFormats = append(myFormats, formats.FormatOsmBz2, "osm.bz2.md5")
		case "shp":
			myFormats = append(myFormats, formats.FormatShpZip)
		case "history":
			myFormats = append(myFormats, formats.FormatOshPbf)
		}
	}

	myFormats = append(myFormats, formats.FormatPoly, formats.FormatKml, formats.FormatState)

	return myFormats
}
