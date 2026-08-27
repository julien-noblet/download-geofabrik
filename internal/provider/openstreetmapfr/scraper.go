package openstreetmapfr

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
	ProviderName        = "openstreetmap.fr"
	DefaultConfigFile   = "openstreetmap.fr.yml"
	BaseURL             = "https://download.openstreetmap.fr/extracts"
	StartURL            = "https://download.openstreetmap.fr/extracts/"
	defaultTimeout      = 30 * time.Second
	defaultKeepAlive    = 30 * time.Second
	defaultIdleTimeout  = 90 * time.Second
	defaultMaxIdleConns = 100
	concurrencyLimit    = 25
	maxConnsMultiplier  = 2
	workChanCapacity    = 1000
	minParentListLength = 4
)

var exceptionList = map[string]struct{}{
	"central":       {},
	"central-east":  {},
	"central-north": {},
	"central-south": {},
	"central-west":  {},
	"central_east":  {},
	"central_north": {},
	"central_south": {},
	"central_west":  {},
	"coastral":      {},
	"east":          {},
	"east_central":  {},
	"east-central":  {},
	"eastern":       {},
	"lake":          {},
	"north":         {},
	"north_central": {},
	"north-central": {},
	"north-east":    {},
	"north-eastern": {},
	"north-west":    {},
	"north-western": {},
	"north_east":    {},
	"north_eastern": {},
	"north_west":    {},
	"north_western": {},
	"northeast":     {},
	"northern":      {},
	"northwest":     {},
	"south":         {},
	"south_central": {},
	"south-central": {},
	"south-east":    {},
	"south-south":   {},
	"south-west":    {},
	"south_east":    {},
	"south_south":   {},
	"south_west":    {},
	"southeast":     {},
	"southern":      {},
	"southwest":     {},
	"west":          {},
	"west_central":  {},
	"west-central":  {},
	"western":       {},
	"france_taaf":   {},
	"sevastopol":    {},
	"la_rioja":      {},
	"jura":          {},
	"santa_cruz":    {},
}

// Provider implements provider.Provider for download.openstreetmap.fr.
// Field alignment optimized.
type Provider struct {
	Client   *http.Client
	BaseURL  string
	StartURL string
}

// NewProvider creates a new OpenStreetMap.fr scraper provider.
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

// Name returns the unique service name.
func (p *Provider) Name() string {
	return ProviderName
}

// Description returns a human-readable description.
func (p *Provider) Description() string {
	return "OpenStreetMap.fr extract downloads service"
}

// DefaultConfigFile returns the default configuration file name.
func (p *Provider) DefaultConfigFile() string {
	return DefaultConfigFile
}

// DefaultFormats returns format definitions for OpenStreetMap.fr.
func DefaultFormats() catalog.FormatDefinitions {
	return catalog.FormatDefinitions{
		"osm.pbf.md5":        {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
		catalog.FormatOsmPbf: {ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf"},
		catalog.FormatPoly:   {ID: catalog.FormatPoly, Loc: ".poly", BasePath: "../polygons/"},
		catalog.FormatState:  {ID: catalog.FormatState, Loc: ".state.txt"},
	}
}

// FetchCatalog crawls the directory index of OpenStreetMap.fr concurrently and generates a Catalog.
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
				recordErr(fmt.Errorf("context canceled during crawl: %w", ctx.Err()))
				activeWorkers.Done()

				continue

			case sem <- struct{}{}:
			}

			go func(currentURL string) {
				defer func() {
					<-sem
					activeWorkers.Done()
				}()

				links, err := p.fetchAndProcessPage(ctx, currentURL, cat)
				if err != nil {
					recordErr(err)

					return
				}

				for _, link := range links {
					if strings.HasPrefix(link, p.StartURL) {
						enqueue(link)
					}
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
		return nil, fmt.Errorf("error fetching %s: %w", currentURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: unexpected HTTP %d from %s", ErrFetchCatalog, resp.StatusCode, currentURL)
	}

	return p.parseHTMLStream(resp.Body, currentURL, cat)
}

func (p *Provider) parseHTMLStream(reader io.Reader, currentURL string, cat *catalog.Catalog) ([]string, error) {
	var subDirs []string

	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			if errors.Is(tokenizer.Err(), io.EOF) {
				return subDirs, nil
			}

			return nil, fmt.Errorf("error parsing HTML from %s: %w", currentURL, tokenizer.Err())

		case html.StartTagToken, html.SelfClosingTagToken:
			if subDir := p.handleAnchor(tokenizer, currentURL, cat); subDir != "" {
				subDirs = append(subDirs, subDir)
			}

		case html.TextToken, html.EndTagToken, html.CommentToken, html.DoctypeToken:
			// Non-anchor tokens
		}
	}
}

func (p *Provider) handleAnchor(tokenizer *html.Tokenizer, currentURL string, cat *catalog.Catalog) string {
	tagName, hasAttr := tokenizer.TagName()
	if string(tagName) != "a" || !hasAttr {
		return ""
	}

	for {
		key, val, more := tokenizer.TagAttr()
		if string(key) == "href" {
			href := string(val)
			if !shouldSkipHref(href) {
				fullURL := resolveURL(currentURL, href)
				if strings.HasSuffix(href, "/") {
					return fullURL
				}

				p.parseFileLink(cat, fullURL)
			}

			return ""
		}

		if !more {
			return ""
		}
	}
}

func shouldSkipHref(href string) bool {
	return href == "" || strings.Contains(href, "?") || strings.Contains(href, "-latest") || strings.HasPrefix(href, "/") || href == "../"
}

func resolveURL(base, ref string) string {
	if strings.HasPrefix(ref, "http://") || strings.HasPrefix(ref, "https://") {
		return ref
	}

	if !strings.HasSuffix(base, "/") {
		base += "/"
	}

	return base + ref
}

func (p *Provider) parseFileLink(cat *catalog.Catalog, href string) {
	parent, parents := getParent(href)
	if parent != "" && !cat.Exist(parent) {
		p.makeParents(cat, parent, parents)
	}

	parts := strings.Split(parents[len(parents)-1], ".")
	if parts[0] == "" || len(parts) < 2 {
		return
	}

	name := exceptions(parts[0], parent)

	ext := strings.Join(parts[1:], ".")
	if strings.Contains(ext, "state.txt") {
		ext = catalog.FormatState
	}

	elem := catalog.Element{
		ID:     name,
		File:   parts[0],
		Name:   name,
		Parent: parent,
	}

	_ = cat.MergeElement(&elem)
	cat.AddExtension(name, ext)
}

func (p *Provider) makeParents(cat *catalog.Catalog, parent string, gparents []string) {
	if parent == "" || cat.Exist(parent) {
		return
	}

	gparent := getGparent(gparents)
	metaElem := catalog.Element{
		ID:     parent,
		Name:   parent,
		Parent: gparent,
		Meta:   true,
	}

	_ = cat.MergeElement(&metaElem)

	if gparent != "" && len(gparents) > 1 {
		p.makeParents(cat, gparent, gparents[:len(gparents)-1])
	}
}

func getParent(href string) (parent string, parts []string) {
	href = strings.TrimSuffix(href, "/")
	parts = strings.Split(href, "/")

	if len(parts) > minParentListLength {
		p := parts[len(parts)-2]
		if strings.EqualFold(p, "extracts") || strings.EqualFold(p, "polygons") {
			return "", parts
		}

		return p, parts
	}

	return "", parts
}

func getGparent(gparents []string) string {
	if len(gparents) < minParentListLength {
		return ""
	}

	gp := gparents[len(gparents)-3]
	if gp == "http:" || gp == "https:" || gp == "download.openstreetmap.fr" || gp == "extracts" || gp == "polygons" {
		return ""
	}

	return gp
}

func exceptions(name, parent string) string {
	if _, exists := exceptionList[name]; exists && parent != "" {
		return fmt.Sprintf("%s_%s", parent, name)
	}

	return name
}
