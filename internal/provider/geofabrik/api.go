package geofabrik

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var ErrFetchCatalog = errors.New("failed to fetch catalog")

const (
	ProviderName               = "geofabrik"
	DefaultConfigFile          = "geofabrik.yml"
	GeofabrikIndexURL          = "https://download.geofabrik.de/index-v1-nogeom.json"
	GeofabrikBaseURL           = "https://download.geofabrik.de"
	defaultTimeout             = 60 * time.Second
	defaultKeepAlive           = 30 * time.Second
	defaultIdleTimeout         = 90 * time.Second
	defaultMaxIdleConns        = 20
	defaultMaxIdleConnsPerHost = 10
	alwaysPresentFormatsCount  = 3
	formatPerURLMultiplier     = 2
)

// Provider implements provider.Provider for the Geofabrik service using JSON API.
// Field alignment optimized.
type Provider struct {
	Client   *http.Client
	IndexURL string
	BaseURL  string
}

// NewProvider creates a new Geofabrik API provider with tuned transport.
func NewProvider() *Provider {
	return &Provider{
		IndexURL: GeofabrikIndexURL,
		BaseURL:  GeofabrikBaseURL,
		Client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   defaultTimeout,
					KeepAlive: defaultKeepAlive,
				}).DialContext,
				MaxIdleConns:        defaultMaxIdleConns,
				MaxIdleConnsPerHost: defaultMaxIdleConnsPerHost,
				IdleConnTimeout:     defaultIdleTimeout,
				ForceAttemptHTTP2:   true,
			},
		},
	}
}

// Name returns the provider's unique service name.
func (p *Provider) Name() string {
	return ProviderName
}

// Description returns a human-readable description.
func (p *Provider) Description() string {
	return "Geofabrik OSM extract downloads service (JSON API)"
}

// DefaultConfigFile returns the default YAML filename.
func (p *Provider) DefaultConfigFile() string {
	return DefaultConfigFile
}

// DefaultFormats returns format definitions used by Geofabrik.
func DefaultFormats() catalog.FormatDefinitions {
	return catalog.FormatDefinitions{
		"osm.bz2.md5":         {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
		"osm.pbf.md5":         {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
		catalog.FormatKml:     {ID: catalog.FormatKml, Loc: ".kml"},
		catalog.FormatMBTiles: {ID: catalog.FormatMBTiles, Loc: "-latest-free.mbtiles.zip", ToLoc: "latest-free.mbtiles.zip"},
		catalog.FormatOsmBz2:  {ID: catalog.FormatOsmBz2, Loc: "-latest.osm.bz2"},
		catalog.FormatOsmPbf:  {ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf"},
		catalog.FormatPoly:    {ID: catalog.FormatPoly, Loc: ".poly"},
		catalog.FormatShpZip:  {ID: catalog.FormatShpZip, Loc: "-shortbread-1.0.mbtiles"},
		catalog.FormatState:   {ID: catalog.FormatState, Loc: "-updates/state.txt"},
	}
}

type indexJSON struct {
	Features []struct {
		Properties struct {
			Urls   map[string]string `json:"urls"`
			ID     string            `json:"id"`
			Name   string            `json:"name"`
			Parent string            `json:"parent,omitempty"`
		} `json:"properties"`
	} `json:"features"`
}

// FetchCatalog downloads the index JSON and builds a catalog.Catalog.
func (p *Provider) FetchCatalog(ctx context.Context) (*catalog.Catalog, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.IndexURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("cannot create request for %s: %w", p.IndexURL, err)
	}

	client := p.Client
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w from %s: %w", ErrFetchCatalog, p.IndexURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: unexpected HTTP %d from %s", ErrFetchCatalog, resp.StatusCode, p.IndexURL)
	}

	var index indexJSON
	if err := json.NewDecoder(resp.Body).Decode(&index); err != nil {
		return nil, fmt.Errorf("cannot decode index JSON: %w", err)
	}

	cat := catalog.New()
	cat.BaseURL = p.BaseURL
	cat.Formats = DefaultFormats()

	for _, feat := range index.Features {
		elem := catalog.Element{
			ID:      feat.Properties.ID,
			Name:    feat.Properties.Name,
			Parent:  feat.Properties.Parent,
			Formats: extractFormats(feat.Properties.Urls),
		}

		if err := cat.MergeElement(&elem); err != nil {
			return nil, fmt.Errorf("cannot merge element %s: %w", elem.ID, err)
		}
	}

	return cat, nil
}

func extractFormats(urls map[string]string) catalog.Formats {
	formatList := make(catalog.Formats, 0, len(urls)*formatPerURLMultiplier+alwaysPresentFormatsCount)

	for k := range urls {
		switch k {
		case "pbf":
			formatList = append(formatList, catalog.FormatOsmPbf, "osm.pbf.md5")
		case "bz2":
			formatList = append(formatList, catalog.FormatOsmBz2, "osm.bz2.md5")
		case "shp":
			formatList = append(formatList, catalog.FormatShpZip)
		case "history":
			formatList = append(formatList, catalog.FormatOshPbf)
		}
	}

	formatList = append(formatList, catalog.FormatPoly, catalog.FormatKml, catalog.FormatState)

	return formatList
}
