package geo2day

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var ErrFetchCatalog = errors.New("failed to fetch catalog")

const (
	ProviderName       = "geo2day"
	DefaultConfigFile  = "geo2day.yml"
	BaseURL            = "https://geo2day.com"
	StartURL           = "https://geo2day.com/"
	defaultTimeout     = 30 * time.Second
	defaultKeepAlive   = 30 * time.Second
	defaultIdleTimeout = 90 * time.Second
	minExtParts        = 2
	minGeo2DayParts    = 4
)

// Provider implements provider.Provider for geo2day.com.
// Field alignment optimized.
type Provider struct {
	Client   *http.Client
	BaseURL  string
	StartURL string
}

// NewProvider creates a new Geo2Day scraper provider.
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
				IdleConnTimeout: defaultIdleTimeout,
			},
		},
	}
}

// Name returns the provider name.
func (p *Provider) Name() string {
	return ProviderName
}

// Description returns a description of the provider.
func (p *Provider) Description() string {
	return "Geo2Day OSM extracts downloads service"
}

// DefaultConfigFile returns the default configuration file name.
func (p *Provider) DefaultConfigFile() string {
	return DefaultConfigFile
}

// DefaultFormats returns format definitions for Geo2Day.
func DefaultFormats() catalog.FormatDefinitions {
	return catalog.FormatDefinitions{
		"osm.pbf.md5":         {ID: "osm.pbf.md5", Loc: ".md5"},
		catalog.FormatGeoJSON: {ID: catalog.FormatGeoJSON, Loc: ".geojson"},
		catalog.FormatOsmPbf:  {ID: catalog.FormatOsmPbf, Loc: ".pbf"},
		catalog.FormatPoly:    {ID: catalog.FormatPoly, Loc: ".poly"},
	}
}

var exceptionList = []struct {
	ID     string
	Parent string
}{
	{"la_rioja", "argentina"},
	{"la_rioja", "spain"},
	{"guyane", "france"},
	{"guyane", "south-america"},
	{"sevastopol", "ukraine"},
	{"sevastopol", "russia"},
	{"limburg", "netherlands"},
	{"limburg", "flanders"},
	{"cordoba", "argentina"},
	{"cordoba", "andalucia"},
	{"georgia", "asia"},
	{"georgia", "us"},
}

// FetchCatalog scrapes the Geo2Day index and builds a catalog.Catalog.
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
		return nil, fmt.Errorf("%w: unexpected HTTP %d", ErrFetchCatalog, resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot parse HTML: %w", err)
	}

	cat := catalog.New()
	cat.BaseURL = p.BaseURL
	cat.Formats = DefaultFormats()

	doc.Find("table td a").Each(func(_ int, selection *goquery.Selection) {
		p.processLink(cat, selection)
	})

	return cat, nil
}

func (p *Provider) processLink(cat *catalog.Catalog, selection *goquery.Selection) {
	href, exists := selection.Attr("href")
	if !exists {
		return
	}

	rawID, ext := splitFileExt(href)
	if rawID == "" {
		return
	}

	parent, _ := splitParent(href)
	elem := catalog.Element{
		ID:     rawID,
		Name:   selection.Text(),
		Parent: parent,
		Meta:   ext == "html" || ext == "",
	}

	applyExceptions(&elem)
	_ = cat.MergeElement(&elem)

	if ext != "" && ext != "html" {
		if ext == "pbf" {
			ext = catalog.FormatOsmPbf
		}

		cat.AddExtension(elem.ID, ext)

		if ext == catalog.FormatOsmPbf {
			cat.AddExtension(elem.ID, "osm.pbf.md5")
		}
	}
}

func applyExceptions(elem *catalog.Element) {
	for _, exc := range exceptionList {
		if elem.ID == exc.ID && elem.Parent == exc.Parent {
			elem.ID = fmt.Sprintf("%s-%s", elem.ID, elem.Parent)

			return
		}
	}
}

func splitFileExt(urlStr string) (filename, extension string) {
	parts := strings.Split(urlStr, "/")
	last := parts[len(parts)-1]

	dotParts := strings.Split(last, ".")
	if len(dotParts) < minExtParts {
		return dotParts[0], ""
	}

	return dotParts[0], strings.Join(dotParts[1:], ".")
}

func splitParent(urlStr string) (parent, path string) {
	parts := strings.Split(urlStr, "/")
	if len(parts) < minGeo2DayParts {
		return "", strings.Join(parts[:len(parts)-1], "/")
	}

	return parts[len(parts)-2], strings.Join(parts[:len(parts)-1], "/")
}
