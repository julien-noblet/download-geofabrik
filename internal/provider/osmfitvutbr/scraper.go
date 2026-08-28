package osmfitvutbr

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
	ProviderName               = "osm.fit.vutbr.cz"
	DefaultConfigFile          = "osm.fit.vutbr.cz.yml"
	BaseURL                    = "https://osm.fit.vutbr.cz/extracts"
	StartURL                   = "https://osm.fit.vutbr.cz/extracts/"
	defaultTimeout             = 30 * time.Second
	defaultKeepAlive           = 30 * time.Second
	defaultIdleTimeout         = 90 * time.Second
	defaultMaxIdleConns        = 20
	defaultMaxIdleConnsPerHost = 10
)

// Provider implements provider.Provider for osm.fit.vutbr.cz (Czech Republic extracts).
// Field alignment optimized.
type Provider struct {
	Client   *http.Client
	BaseURL  string
	StartURL string
}

// NewProvider creates a new FIT VUTBR Czechia scraper provider.
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
	return "Czech Republic OSM extract downloads service (FIT VUTBR)"
}

// DefaultConfigFile returns the default configuration filename.
func (p *Provider) DefaultConfigFile() string {
	return DefaultConfigFile
}

// DefaultFormats returns format definitions supported by osm.fit.vutbr.cz.
func DefaultFormats() catalog.FormatDefinitions {
	return catalog.FormatDefinitions{
		catalog.FormatOsmPbf: {ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf", BasePath: "czech_republic/"},
		catalog.FormatOsmBz2: {ID: catalog.FormatOsmBz2, Loc: "-latest.osm.bz2", BasePath: "czech_republic/"},
		catalog.FormatPoly:   {ID: catalog.FormatPoly, Loc: ".poly"},
	}
}

// FetchCatalog scrapes the index of osm.fit.vutbr.cz and generates a catalog.Catalog.
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

	if err := parseFitVutbrHTML(resp.Body, cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func parseFitVutbrHTML(reader io.Reader, cat *catalog.Catalog) error {
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

	// Subdirectory representing an extract region (e.g. czech_republic/)
	if strings.HasSuffix(href, "/") {
		dir := strings.Trim(href, "/")
		elem := catalog.Element{
			ID:      dir,
			Name:    formatName(dir),
			Formats: catalog.Formats{catalog.FormatOsmPbf, catalog.FormatOsmBz2},
		}
		_ = cat.MergeElement(&elem)

		return
	}

	// Direct polygon file (e.g. czech-republic.poly)
	if strings.HasSuffix(href, ".poly") {
		raw := strings.TrimSuffix(href, ".poly")
		elemID := strings.ReplaceAll(raw, "-", "_")
		elem := catalog.Element{
			ID:   elemID,
			Name: formatName(elemID),
			File: raw,
		}
		_ = cat.MergeElement(&elem)
		cat.AddExtension(elemID, catalog.FormatPoly)
	}
}

func formatName(name string) string {
	words := strings.Split(name, "_")
	for i, w := range words {
		if w != "" {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}

	return strings.Join(words, " ")
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
		strings.HasPrefix(href, "v6-planet")
}
