package bbbike

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"golang.org/x/net/html"
)

var ErrFetchCatalog = errors.New("failed to fetch catalog")

const (
	ProviderName               = "bbbike"
	DefaultConfigFile          = "bbbike.yml"
	BaseURL                    = "https://download.bbbike.org/osm/bbbike"
	StartURL                   = "https://download.bbbike.org/osm/bbbike/"
	defaultTimeout             = 30 * time.Second
	defaultKeepAlive           = 30 * time.Second
	defaultIdleTimeout         = 90 * time.Second
	defaultMaxIdleConns        = 50
	defaultMaxIdleConnsPerHost = 20
)

// Provider implements provider.Provider for download.bbbike.org.
// Field alignment optimized.
type Provider struct {
	Client   *http.Client
	BaseURL  string
	StartURL string
}

// NewProvider creates a new BBBike scraper provider.
func NewProvider() *Provider {
	return &Provider{
		BaseURL:  BaseURL,
		StartURL: StartURL,
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

// Name returns the provider's name.
func (p *Provider) Name() string {
	return ProviderName
}

// Description returns a description of the provider.
func (p *Provider) Description() string {
	return "BBBike OSM city extracts downloads service"
}

// DefaultConfigFile returns the default configuration filename.
func (p *Provider) DefaultConfigFile() string {
	return DefaultConfigFile
}

// DefaultFormats returns format definitions supported by BBBike.
func DefaultFormats() catalog.FormatDefinitions {
	return catalog.FormatDefinitions{
		catalog.FormatCSV:            {ID: catalog.FormatCSV, Loc: ".osm.csv.xz", ToLoc: ".osm.csv.xz"},
		catalog.FormatGarminOSM:      {ID: catalog.FormatGarminOSM, Loc: ".osm.garmin-osm.zip"},
		catalog.FormatGarminOnroad:   {ID: catalog.FormatGarminOnroad, Loc: ".osm.garmin-onroad-latin1.zip"},
		catalog.FormatGarminOntrail:  {ID: catalog.FormatGarminOntrail, Loc: ".osm.garmin-ontrail-latin1.zip"},
		catalog.FormatGarminOpenTopo: {ID: catalog.FormatGarminOpenTopo, Loc: ".osm.garmin-opentopo-latin1.zip"},
		catalog.FormatGeoJSON:        {ID: catalog.FormatGeoJSON, Loc: ".osm.geojson.xz", ToLoc: ".geojson.xz"},
		catalog.FormatMBTiles:        {ID: catalog.FormatMBTiles, Loc: ".osm.mbtiles-openmaptiles.zip", ToLoc: "osm.mbtiles-openmaptiles.zip"},
		catalog.FormatMapsforge:      {ID: catalog.FormatMapsforge, Loc: ".osm.mapsforge-osm.zip"},
		catalog.FormatOsmGz:          {ID: catalog.FormatOsmGz, Loc: ".osm.gz"},
		catalog.FormatOsmPbf:         {ID: catalog.FormatOsmPbf, Loc: ".osm.pbf"},
		catalog.FormatPoly:           {ID: catalog.FormatPoly, Loc: ".poly"},
		catalog.FormatShpZip:         {ID: catalog.FormatShpZip, Loc: ".osm.shp.zip"},
	}
}

var standardBBBikeFormats = catalog.Formats{
	catalog.FormatCSV,
	catalog.FormatGarminOSM,
	catalog.FormatGarminOnroad,
	catalog.FormatGarminOntrail,
	catalog.FormatGarminOpenTopo,
	catalog.FormatGeoJSON,
	catalog.FormatMBTiles,
	catalog.FormatMapsforge,
	catalog.FormatOsmGz,
	catalog.FormatOsmPbf,
	catalog.FormatPoly,
	catalog.FormatShpZip,
}

// FetchCatalog scrapes the BBBike cities index and builds a catalog.Catalog.
func (p *Provider) FetchCatalog(ctx context.Context) (*catalog.Catalog, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.StartURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := p.Client
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFetchCatalog, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: unexpected HTTP status %d", ErrFetchCatalog, resp.StatusCode)
	}

	cat := catalog.New()
	cat.BaseURL = p.BaseURL
	cat.Formats = DefaultFormats()

	if err := parseBBBikeHTML(resp.Body, cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func parseBBBikeHTML(reader io.Reader, cat *catalog.Catalog) error {
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			if errors.Is(tokenizer.Err(), io.EOF) {
				return nil
			}

			return fmt.Errorf("cannot parse HTML: %w", tokenizer.Err())

		case html.StartTagToken, html.SelfClosingTagToken:
			processAnchorTag(tokenizer, cat)

		case html.TextToken, html.EndTagToken, html.CommentToken, html.DoctypeToken:
			// Non-anchor tokens
		}
	}
}

func processAnchorTag(tokenizer *html.Tokenizer, cat *catalog.Catalog) {
	tagName, hasAttr := tokenizer.TagName()
	if string(tagName) != "a" || !hasAttr {
		return
	}

	for {
		key, val, more := tokenizer.TagAttr()
		if string(key) == "href" {
			city := cleanCityName(string(val))
			if city != "" {
				elem := catalog.Element{
					ID:      city,
					Name:    city,
					File:    city + "/" + city,
					Formats: append(catalog.Formats(nil), standardBBBikeFormats...),
				}
				_ = cat.MergeElement(&elem)
			}

			return
		}

		if !more {
			return
		}
	}
}

func cleanCityName(href string) string {
	href = strings.Trim(href, "/")
	if href == "" || strings.Contains(href, "?") || strings.Contains(href, ":") || strings.Contains(href, "..") {
		return ""
	}

	// Must start with uppercase letter (BBBike convention for cities: Aachen, Berlin, Paris, etc.)
	if href[0] >= 'A' && href[0] <= 'Z' && !strings.Contains(href, "/") {
		return href
	}

	return ""
}
