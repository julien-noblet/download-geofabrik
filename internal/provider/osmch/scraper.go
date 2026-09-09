package osmch

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var ErrFetchCatalog = catalog.ErrFetchCatalog

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

// New creates a new Swiss OSM (planet.osm.ch) scraper provider.
func New() *Provider {
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

// NewProvider creates a new Swiss OSM (planet.osm.ch) scraper provider.
//
// Deprecated: Use New instead.
func NewProvider() *Provider {
	return New()
}

// Name returns the unique service name.
func (p *Provider) Name() string {
	if p == nil {
		return ""
	}

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
		catalog.FormatPbf:       {ID: catalog.FormatPbf, Loc: ".pbf"},
		catalog.FormatPoly:      {ID: catalog.FormatPoly, Loc: ".poly"},
		catalog.FormatOBF:       {ID: catalog.FormatOBF, Loc: ".obf"},
		catalog.FormatGarminOSM: {ID: catalog.FormatGarminOSM, Loc: "-garmin.zip"},
	}
}

// FetchCatalog scrapes the index of planet.osm.ch and generates a catalog.Catalog.
func (p *Provider) FetchCatalog(ctx context.Context) (*catalog.Catalog, error) {
	if p == nil {
		return nil, catalog.ErrProviderNil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.StartURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	client := cmp.Or(p.Client, http.DefaultClient)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFetchCatalog, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: unexpected http status %d", ErrFetchCatalog, resp.StatusCode)
	}

	cat := catalog.New()
	cat.BaseURL = p.BaseURL
	cat.Formats = DefaultFormats()

	if err := parseOSMCHHTML(resp.Body, cat); err != nil {
		return nil, fmt.Errorf("parsing catalog html: %w", err)
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

			return fmt.Errorf("cannot parse html: %w", tokenizer.Err())

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
	if trimmed, ok := strings.CutSuffix(href, ".osm.pbf"); ok {
		return trimmed, catalog.FormatOsmPbf
	}

	if trimmed, ok := strings.CutSuffix(href, ".poly"); ok {
		return trimmed, catalog.FormatPoly
	}

	if trimmed, ok := strings.CutSuffix(href, ".obf"); ok {
		return trimmed, catalog.FormatOBF
	}

	if trimmed, ok := strings.CutSuffix(href, "-garmin.zip"); ok {
		return trimmed, catalog.FormatGarminOSM
	}

	if trimmed, ok := strings.CutSuffix(href, ".pbf"); ok {
		return trimmed, catalog.FormatPbf
	}

	return "", ""
}
