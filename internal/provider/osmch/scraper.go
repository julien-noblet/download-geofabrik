package osmch

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
	ProviderName               = "planet.osm.ch"
	DefaultConfigFile          = "planet.osm.ch.yml"
	BaseURL                    = "https://planet.osm.ch"
	StartURL                   = "https://planet.osm.ch/"
	defaultTimeout             = 30 * time.Second
	defaultKeepAlive           = 30 * time.Second
	defaultIdleTimeout         = 90 * time.Second
	defaultMaxIdleConns        = 20
	defaultMaxIdleConnsPerHost = 10
)

// Provider implements provider.Provider for planet.osm.ch.
// Field alignment optimized.
type Provider struct {
	Client   *http.Client
	BaseURL  string
	StartURL string
}

// NewProvider creates a new Swiss OSM (planet.osm.ch) scraper provider.
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

// Name returns the unique service name.
func (p *Provider) Name() string {
	return ProviderName
}

// Description returns a human-readable description.
func (p *Provider) Description() string {
	return "Swiss OSM extracts downloads service (planet.osm.ch)"
}

// DefaultConfigFile returns the default configuration filename.
func (p *Provider) DefaultConfigFile() string {
	return DefaultConfigFile
}

// DefaultFormats returns format definitions supported by planet.osm.ch.
func DefaultFormats() catalog.FormatDefinitions {
	return catalog.FormatDefinitions{
		catalog.FormatOsmPbf:    {ID: catalog.FormatOsmPbf, Loc: ".osm.pbf"},
		catalog.FormatPoly:      {ID: catalog.FormatPoly, Loc: ".poly"},
		catalog.FormatOBF:       {ID: catalog.FormatOBF, Loc: ".obf"},
		catalog.FormatGarminOSM: {ID: catalog.FormatGarminOSM, Loc: "-garmin.zip"},
	}
}

// FetchCatalog scrapes the index of planet.osm.ch and generates a catalog.Catalog.
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

	if err := parseOSMCHHTML(resp.Body, cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func parseOSMCHHTML(reader io.Reader, cat *catalog.Catalog) error {
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
			href := string(val)
			parseFileLink(href, cat)

			return
		}

		if !more {
			return
		}
	}
}

func parseFileLink(href string, cat *catalog.Catalog) {
	if shouldSkipHref(href) {
		return
	}

	elementID, formatID := parseHrefToIDAndFormat(href)
	if elementID == "" || formatID == "" {
		return
	}

	elem := catalog.Element{
		ID:   elementID,
		Name: elementID,
	}

	_ = cat.MergeElement(&elem)
	cat.AddExtension(elementID, formatID)
}

func shouldSkipHref(href string) bool {
	return href == "" ||
		strings.Contains(href, "?") ||
		strings.Contains(href, "http:") ||
		strings.Contains(href, "https:") ||
		strings.HasPrefix(href, "/") ||
		strings.HasPrefix(href, "..") ||
		strings.HasSuffix(href, "/") ||
		strings.HasPrefix(href, ".")
}

func parseHrefToIDAndFormat(href string) (elemID, format string) {
	switch {
	case strings.HasSuffix(href, ".osm.pbf"):
		return strings.TrimSuffix(href, ".osm.pbf"), catalog.FormatOsmPbf

	case strings.HasSuffix(href, ".poly"):
		return strings.TrimSuffix(href, ".poly"), catalog.FormatPoly

	case strings.HasSuffix(href, ".obf"):
		return strings.TrimSuffix(href, ".obf"), catalog.FormatOBF

	case strings.HasSuffix(href, "-garmin.zip"):
		return strings.TrimSuffix(href, "-garmin.zip"), catalog.FormatGarminOSM

	case strings.HasSuffix(href, ".pbf"):
		return strings.TrimSuffix(href, ".pbf"), catalog.FormatOsmPbf
	}

	return "", ""
}
