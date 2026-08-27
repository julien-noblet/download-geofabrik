package geo2day

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"golang.org/x/net/html"
)

var ErrFetchCatalog = errors.New("failed to fetch catalog")

const (
	ProviderName        = "geo2day"
	DefaultConfigFile   = "geo2day.yml"
	BaseURL             = "https://geo2day.com"
	StartURL            = "https://geo2day.com/"
	defaultTimeout      = 15 * time.Second
	defaultKeepAlive    = 30 * time.Second
	defaultIdleTimeout  = 90 * time.Second
	concurrencyLimit    = 25
	maxConnsMultiplier  = 2
	defaultMaxIdleConns = 100
	workChanCapacity    = 1000
	minExtParts         = 2
	minGeo2DayParts     = 4
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
				MaxIdleConns:        defaultMaxIdleConns,
				MaxIdleConnsPerHost: concurrencyLimit,
				MaxConnsPerHost:     concurrencyLimit * maxConnsMultiplier,
				IdleConnTimeout:     defaultIdleTimeout,
				ForceAttemptHTTP2:   true,
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

// FetchCatalog scrapes the Geo2Day index concurrently and builds a catalog.Catalog.
func (p *Provider) FetchCatalog(ctx context.Context) (*catalog.Catalog, error) {
	cat := catalog.New()
	cat.BaseURL = p.BaseURL
	cat.Formats = DefaultFormats()

	var (
		visitedMu     sync.Mutex
		visited       = make(map[string]bool)
		workChan      = make(chan string, workChanCapacity)
		activeWorkers sync.WaitGroup
		errOnce       sync.Once
		firstErr      error
	)

	sem := make(chan struct{}, concurrencyLimit)

	recordErr := func(err error) {
		errOnce.Do(func() {
			firstErr = err
		})
	}

	enqueue := func(targetURL string) {
		visitedMu.Lock()
		if visited[targetURL] {
			visitedMu.Unlock()

			return
		}

		visited[targetURL] = true
		visitedMu.Unlock()

		activeWorkers.Add(1)

		go func() {
			workChan <- targetURL
		}()
	}

	enqueue(p.StartURL)

	go func() {
		for target := range workChan {
			select {
			case <-ctx.Done():
				recordErr(fmt.Errorf("crawl context canceled: %w", ctx.Err()))
				activeWorkers.Done()

				continue

			case sem <- struct{}{}:
			}

			go func(pageURL string) {
				defer func() {
					<-sem
					activeWorkers.Done()
				}()

				subPages, err := p.fetchAndProcessPage(ctx, pageURL, cat)
				if err != nil {
					recordErr(err)

					return
				}

				for _, page := range subPages {
					enqueue(page)
				}
			}(target)
		}
	}()

	activeWorkers.Wait()
	close(workChan)

	if firstErr != nil {
		return nil, firstErr
	}

	return cat, nil
}

func (p *Provider) fetchAndProcessPage(ctx context.Context, currentURL string, cat *catalog.Catalog) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, currentURL, http.NoBody)
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

	return p.parseHTMLStream(resp.Body, cat)
}

func (p *Provider) parseHTMLStream(reader io.Reader, cat *catalog.Catalog) ([]string, error) {
	var subPages []string

	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			if errors.Is(tokenizer.Err(), io.EOF) {
				return subPages, nil
			}

			return nil, fmt.Errorf("cannot parse HTML: %w", tokenizer.Err())

		case html.StartTagToken, html.SelfClosingTagToken:
			if subPage := p.handleAnchor(tokenizer, tokenType, cat); subPage != "" {
				subPages = append(subPages, subPage)
			}

		case html.TextToken, html.EndTagToken, html.CommentToken, html.DoctypeToken:
			// Non-anchor tokens
		}
	}
}

func (p *Provider) handleAnchor(tokenizer *html.Tokenizer, tokenType html.TokenType, cat *catalog.Catalog) string {
	tagName, hasAttr := tokenizer.TagName()
	if string(tagName) != "a" || !hasAttr {
		return ""
	}

	href := extractHref(tokenizer)
	if href == "" {
		return ""
	}

	var text string

	if tokenType == html.StartTagToken {
		if nextType := tokenizer.Next(); nextType == html.TextToken {
			text = string(tokenizer.Text())
		}
	}

	fullURL := p.resolveURL(href)
	rawID, ext := splitFileExt(href)

	if ext == "html" {
		var subPage string
		if strings.HasPrefix(fullURL, p.BaseURL) || strings.HasPrefix(fullURL, p.StartURL) {
			subPage = fullURL
		}

		p.processHTMLLink(cat, fullURL, rawID, text)

		return subPage
	}

	p.processFileLink(cat, href, rawID, ext, text)

	return ""
}

func extractHref(tokenizer *html.Tokenizer) string {
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

func (p *Provider) resolveURL(ref string) string {
	if strings.HasPrefix(ref, "http://") || strings.HasPrefix(ref, "https://") {
		return ref
	}

	trimmedBase := strings.TrimRight(p.BaseURL, "/")

	if !strings.HasPrefix(ref, "/") {
		ref = "/" + ref
	}

	return trimmedBase + ref
}

func (p *Provider) processHTMLLink(cat *catalog.Catalog, fullURL, rawID, name string) {
	parent, parentPath := splitParent(fullURL)
	elem := catalog.Element{
		ID:     rawID,
		Name:   name,
		Parent: parent,
		Meta:   true,
	}

	applyExceptions(&elem)

	if !cat.Exist(parent) && parent != "" {
		gparent, _ := splitParent(parentPath)
		metaParent := catalog.Element{
			ID:     parent,
			Name:   parent,
			Parent: gparent,
			Meta:   true,
		}
		_ = cat.MergeElement(&metaParent)
	}

	_ = cat.MergeElement(&elem)
}

func (p *Provider) processFileLink(cat *catalog.Catalog, href, rawID, ext, name string) {
	if rawID == "" {
		return
	}

	parent, _ := splitParent(href)
	elem := catalog.Element{
		ID:     rawID,
		Name:   name,
		Parent: parent,
		Meta:   false,
	}

	applyExceptions(&elem)
	_ = cat.MergeElement(&elem)

	if ext != "" {
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
