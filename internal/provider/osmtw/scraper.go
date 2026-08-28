package osmtw

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
	ProviderName               = "osm.kcwu.csie.org"
	DefaultConfigFile          = "osm.kcwu.csie.org.yml"
	BaseURL                    = "https://osm.kcwu.csie.org/download/tw-extract"
	StartURL                   = "https://osm.kcwu.csie.org/download/tw-extract/"
	defaultTimeout             = 30 * time.Second
	defaultKeepAlive           = 30 * time.Second
	defaultIdleTimeout         = 90 * time.Second
	defaultMaxIdleConns        = 20
	defaultMaxIdleConnsPerHost = 10
)

// Provider implements provider.Provider for osm.kcwu.csie.org (Taiwan OSM extracts).
// Field alignment optimized.
type Provider struct {
	Client   *http.Client
	BaseURL  string
	StartURL string
}

// NewProvider creates a new Taiwan OSM scraper provider.
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
	return "Taiwan OSM extract downloads service (osm.kcwu.csie.org)"
}

// DefaultConfigFile returns the default configuration filename.
func (p *Provider) DefaultConfigFile() string {
	return DefaultConfigFile
}

// DefaultFormats returns format definitions supported by osm.kcwu.csie.org.
func DefaultFormats() catalog.FormatDefinitions {
	return catalog.FormatDefinitions{
		catalog.FormatO5m:    {ID: catalog.FormatO5m, Loc: "-latest.o5m", BasePath: "recent/"},
		catalog.FormatO5mZst: {ID: catalog.FormatO5mZst, Loc: "-latest.o5m.zst", BasePath: "recent/"},
	}
}

// FetchCatalog scrapes the index of osm.kcwu.csie.org and generates a catalog.Catalog.
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

	if err := parseTaiwanHTML(resp.Body, cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func parseTaiwanHTML(reader io.Reader, cat *catalog.Catalog) error {
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
			parseLink(href, cat)

			return
		}

		if !more {
			return
		}
	}
}

func parseLink(href string, cat *catalog.Catalog) {
	if shouldSkipHref(href) {
		return
	}

	// Direct links or subdirectories
	dir := strings.Trim(href, "/")
	if dir == "recent" || dir == "basemap" {
		elem := catalog.Element{
			ID:      "taiwan",
			Name:    "Taiwan",
			File:    "taiwan",
			Formats: catalog.Formats{catalog.FormatO5m, catalog.FormatO5mZst},
		}
		_ = cat.MergeElement(&elem)
	}
}

func shouldSkipHref(href string) bool {
	return href == "" ||
		strings.Contains(href, "?") ||
		strings.Contains(href, "http:") ||
		strings.Contains(href, "https:") ||
		strings.HasPrefix(href, "/") ||
		strings.HasPrefix(href, "..") ||
		strings.HasPrefix(href, ".") ||
		strings.HasSuffix(href, ".txt") ||
		strings.Trim(href, "/") == "diff"
}
