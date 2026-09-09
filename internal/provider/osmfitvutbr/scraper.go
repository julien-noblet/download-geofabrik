package osmfitvutbr

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

// New creates a new FIT VUTBR Czechia scraper provider.
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

// NewProvider creates a new FIT VUTBR Czechia scraper provider.
//
// Deprecated: Use New instead.
func NewProvider() *Provider {
	return New()
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
		catalog.FormatOsmPbf: {ID: catalog.FormatOsmPbf, Loc: ".osm.pbf", BasePath: "czech_republic/"},
		catalog.FormatOsmBz2: {ID: catalog.FormatOsmBz2, Loc: ".osm.bz2", BasePath: "czech_republic/"},
		catalog.FormatPoly:   {ID: catalog.FormatPoly, Loc: ".poly"},
	}
}

func (p *Provider) fetchHTML(ctx context.Context, targetURL string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := cmp.Or(p.Client, http.DefaultClient)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFetchCatalog, err)
	}

	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()

		return nil, fmt.Errorf("%w: unexpected http status %d", ErrFetchCatalog, resp.StatusCode)
	}

	return resp.Body, nil
}

// FetchCatalog scrapes the index of osm.fit.vutbr.cz and generates a catalog.Catalog.
func (p *Provider) FetchCatalog(ctx context.Context) (*catalog.Catalog, error) {
	body, err := p.fetchHTML(ctx, p.StartURL)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	cat := catalog.New()
	cat.BaseURL = p.BaseURL
	cat.Formats = DefaultFormats()

	subdirs, err := parseFitVutbrRootHTML(body, cat)
	if err != nil {
		return nil, err
	}

	for _, dir := range subdirs {
		subURL := strings.TrimSuffix(p.StartURL, "/") + "/" + dir + "/"

		subBody, err := p.fetchHTML(ctx, subURL)
		if err != nil {
			return nil, err
		}

		err = parseFitVutbrSubdirHTML(subBody, cat, dir)
		_ = subBody.Close()

		if err != nil {
			return nil, err
		}
	}

	return cat, nil
}

func parseFitVutbrRootHTML(reader io.Reader, cat *catalog.Catalog) ([]string, error) {
	tokenizer := html.NewTokenizer(reader)

	var subdirs []string

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			if errors.Is(tokenizer.Err(), io.EOF) {
				return subdirs, nil
			}

			return nil, fmt.Errorf("cannot parse html: %w", tokenizer.Err())

		case html.StartTagToken, html.SelfClosingTagToken:
			if href := extractHref(tokenizer); href != "" {
				parseRootLink(href, cat, &subdirs)
			}

		case html.TextToken, html.EndTagToken, html.CommentToken, html.DoctypeToken:
			// Non-anchor tokens
		}
	}
}

func parseFitVutbrSubdirHTML(reader io.Reader, cat *catalog.Catalog, dir string) error {
	tokenizer := html.NewTokenizer(reader)

	var (
		latestPbfDate string
		latestPbfBase string
		latestBz2Date string
		latestBz2Base string
	)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			if errors.Is(tokenizer.Err(), io.EOF) {
				addLatestElement(cat, dir, latestPbfBase, latestBz2Base)

				return nil
			}

			return fmt.Errorf("cannot parse html: %w", tokenizer.Err())

		case html.StartTagToken, html.SelfClosingTagToken:
			if href := extractHref(tokenizer); href != "" {
				parseSubdirLink(href, dir, &latestPbfDate, &latestPbfBase, &latestBz2Date, &latestBz2Base)
			}

		case html.TextToken, html.EndTagToken, html.CommentToken, html.DoctypeToken:
			// Non-anchor tokens
		}
	}
}

func extractHref(tokenizer *html.Tokenizer) string {
	tagName, hasAttr := tokenizer.TagName()
	if string(tagName) != "a" || !hasAttr {
		return ""
	}

	for {
		key, val, more := tokenizer.TagAttr()
		if string(key) == "href" {
			return string(val)
		}

		if !more {
			return ""
		}
	}
}

func parseRootLink(href string, cat *catalog.Catalog, subdirs *[]string) {
	if shouldSkipHref(href) {
		return
	}

	// Subdirectory representing an extract region (e.g. czech_republic/)
	if strings.HasSuffix(href, "/") {
		dir := strings.Trim(href, "/")
		*subdirs = append(*subdirs, dir)

		return
	}

	// Direct polygon file (e.g. czech-republic.poly)
	if raw, ok := strings.CutSuffix(href, ".poly"); ok {
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

func parseSubdirLink(href, dir string, latestPbfDate, latestPbfBase, latestBz2Date, latestBz2Base *string) {
	if shouldSkipHref(href) {
		return
	}

	var (
		formatID string
		base     string
	)

	switch {
	case strings.HasSuffix(href, ".osm.pbf"):
		formatID = catalog.FormatOsmPbf
		base, _ = strings.CutSuffix(href, ".osm.pbf")

	case strings.HasSuffix(href, ".osm.bz2"):
		formatID = catalog.FormatOsmBz2
		base, _ = strings.CutSuffix(href, ".osm.bz2")

	default:
		return
	}

	date, ok := strings.CutPrefix(base, dir+"-")
	if !ok || date == "" {
		return
	}

	updateLatest(formatID, date, base, latestPbfDate, latestPbfBase, latestBz2Date, latestBz2Base)
}

func updateLatest(formatID, date, base string, latestPbfDate, latestPbfBase, latestBz2Date, latestBz2Base *string) {
	switch formatID {
	case catalog.FormatOsmPbf:
		if *latestPbfDate == "" || date > *latestPbfDate {
			*latestPbfDate = date
			*latestPbfBase = base
		}

	case catalog.FormatOsmBz2:
		if *latestBz2Date == "" || date > *latestBz2Date {
			*latestBz2Date = date
			*latestBz2Base = base
		}
	}
}

func addLatestElement(cat *catalog.Catalog, dir, latestPbfBase, latestBz2Base string) {
	if latestPbfBase != "" {
		latestElem := catalog.Element{
			ID:   "latest",
			Name: formatName(dir) + " (latest)",
			File: latestPbfBase,
		}
		_ = cat.MergeElement(&latestElem)
		cat.AddExtension("latest", catalog.FormatOsmPbf)
	} else if latestBz2Base != "" {
		latestElem := catalog.Element{
			ID:   "latest",
			Name: formatName(dir) + " (latest)",
			File: latestBz2Base,
		}
		_ = cat.MergeElement(&latestElem)
		cat.AddExtension("latest", catalog.FormatOsmBz2)
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
