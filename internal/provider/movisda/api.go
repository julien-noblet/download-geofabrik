package movisda

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var ErrFetchCatalog = errors.New("failed to fetch catalog")

const (
	ProviderName               = "movisda"
	DefaultConfigFile          = "movisda.yml"
	MovisdaIndexURL            = "https://osm.download.movisda.io/admin/Admin-latest.geojson"
	MovisdaBaseURL             = "https://osm.download.movisda.io/admin"
	defaultTimeout             = 60 * time.Second
	defaultKeepAlive           = 30 * time.Second
	defaultIdleTimeout         = 90 * time.Second
	defaultMaxIdleConns        = 20
	defaultMaxIdleConnsPerHost = 10
)

// Provider implements provider.Provider for the Movisda administrative extracts service.
// Field alignment optimized.
type Provider struct {
	Client   *http.Client
	IndexURL string
	BaseURL  string
}

// New creates a new Movisda provider with tuned transport.
func New() *Provider {
	return &Provider{
		IndexURL: MovisdaIndexURL,
		BaseURL:  MovisdaBaseURL,
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

// NewProvider creates a new Movisda provider with tuned transport.
//
// Deprecated: Use New instead.
func NewProvider() *Provider {
	return New()
}

// Name returns the provider's unique service name.
func (p *Provider) Name() string {
	return ProviderName
}

// Description returns a human-readable description.
func (p *Provider) Description() string {
	return "Movisda OSM administrative areas extracts service"
}

// DefaultConfigFile returns the default YAML filename.
func (p *Provider) DefaultConfigFile() string {
	return DefaultConfigFile
}

// DefaultFormats returns format definitions supported by Movisda.
func DefaultFormats() catalog.FormatDefinitions {
	return catalog.FormatDefinitions{
		catalog.FormatOsmPbf:  {ID: catalog.FormatOsmPbf, Loc: ".osm.pbf"},
		catalog.FormatPoly:    {ID: catalog.FormatPoly, Loc: ".poly"},
		catalog.FormatGeoJSON: {ID: catalog.FormatGeoJSON, Loc: ".geojson"},
	}
}

var standardMovisdaFormats = catalog.Formats{
	catalog.FormatOsmPbf,
	catalog.FormatPoly,
	catalog.FormatGeoJSON,
}

//nolint:tagliatelle // GeoJSON schema from upstream
type geoJSONFeature struct {
	Properties struct {
		Name       string `json:"name"`
		NameEN     string `json:"name_en"`
		Prefix     string `json:"prefix"`
		AdminLevel string `json:"admin_level"`
		OSMID      int64  `json:"osm_id"`
	} `json:"properties"`
}

type adminGeoJSON struct {
	Features []geoJSONFeature `json:"features"`
}

// FetchCatalog downloads the Admin-latest.geojson index and builds a catalog.Catalog.
func (p *Provider) FetchCatalog(ctx context.Context) (*catalog.Catalog, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.IndexURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("cannot create request for %s: %w", p.IndexURL, err)
	}

	client := cmp.Or(p.Client, http.DefaultClient)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w from %s: %w", ErrFetchCatalog, p.IndexURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: unexpected http %d from %s", ErrFetchCatalog, resp.StatusCode, p.IndexURL)
	}

	var geo adminGeoJSON
	if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		return nil, fmt.Errorf("cannot decode geojson index: %w", err)
	}

	cat := catalog.New()
	cat.BaseURL = p.BaseURL
	cat.Formats = DefaultFormats()

	for i := range geo.Features {
		elem := buildElement(&geo.Features[i])
		if elem.ID == "" {
			continue
		}

		_ = cat.MergeElement(&elem)
	}

	return cat, nil
}

func buildElement(feat *geoJSONFeature) catalog.Element {
	prefix := feat.Properties.Prefix

	rawID := strings.ToLower(strings.Trim(prefix, "-"))
	if rawID == "" {
		return catalog.Element{}
	}

	name := cmp.Or(feat.Properties.NameEN, feat.Properties.Name)

	return catalog.Element{
		ID:      rawID,
		Name:    name,
		File:    prefix + "latest",
		Formats: slices.Clone(standardMovisdaFormats),
	}
}
