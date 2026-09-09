package geo2day

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var ErrFetchCatalog = catalog.ErrFetchCatalog

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
	FormatOsmPbfMd5     = "osm.pbf.md5"
)

// Provider implements provider.Provider for geo2day.com.
// Field alignment optimized.
type Provider struct {
	Client   *http.Client
	BaseURL  string
	StartURL string
}

// New creates a new Geo2Day scraper provider.
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
				MaxIdleConnsPerHost: concurrencyLimit,
				MaxConnsPerHost:     concurrencyLimit * maxConnsMultiplier,
				IdleConnTimeout:     defaultIdleTimeout,
				ForceAttemptHTTP2:   true,
			},
		},
	}
}

// NewProvider creates a new Geo2Day scraper provider.
//
// Deprecated: Use New instead.
func NewProvider() *Provider {
	return New()
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
		FormatOsmPbfMd5:       {ID: FormatOsmPbfMd5, Loc: ".md5"},
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
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("crawl context canceled: %w", err)
	}

	cat := catalog.New()
	cat.BaseURL = p.BaseURL
	cat.Formats = DefaultFormats()

	crawlCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var (
		visitedMu     sync.Mutex
		visited       = make(map[string]bool)
		workChan      = make(chan string, workChanCapacity)
		activeWorkers sync.WaitGroup
		errOnce       sync.Once
		firstErr      error
	)

	recordErr := func(err error) {
		errOnce.Do(func() {
			firstErr = err

			cancel()
		})
	}

	enqueue := func(targetURL string) {
		if err := crawlCtx.Err(); err != nil {
			recordErr(fmt.Errorf("crawl context canceled: %w", err))

			return
		}

		visitedMu.Lock()
		if visited[targetURL] {
			visitedMu.Unlock()

			return
		}

		visited[targetURL] = true
		visitedMu.Unlock()

		activeWorkers.Add(1)

		go func() {
			select {
			case workChan <- targetURL:
			case <-crawlCtx.Done():
				activeWorkers.Done()
			}
		}()
	}

	crawlState := &crawlContext{
		workers: &activeWorkers,
		onError: recordErr,
		enqueue: enqueue,
		tokens:  make(chan struct{}, concurrencyLimit),
	}

	enqueue(p.StartURL)

	var dispatcherWG sync.WaitGroup

	dispatcherWG.Go(func() {
		defer func() {
			if r := recover(); r != nil {
				crawlState.onError(fmt.Errorf("%w: panic in geo2day dispatcher: %v", ErrFetchCatalog, r))
			}
		}()

		p.runDispatcher(crawlCtx, workChan, cat, crawlState)
	})

	activeWorkers.Wait()
	close(workChan)
	dispatcherWG.Wait()

	if err := ctx.Err(); err != nil && firstErr == nil {
		firstErr = fmt.Errorf("crawl context canceled: %w", err)
	}

	if firstErr != nil {
		return nil, firstErr
	}

	normalizeElementNames(cat)

	return cat, nil
}

func normalizeElementNames(cat *catalog.Catalog) {
	for id, elem := range cat.Elements {
		elem.Name = cmp.Or(elem.Name, id)
		cat.Elements[id] = elem
	}
}

type crawlContext struct {
	workers *sync.WaitGroup
	onError func(error)
	enqueue func(string)
	tokens  chan struct{}
}

func (p *Provider) runDispatcher(ctx context.Context, work <-chan string, cat *catalog.Catalog, state *crawlContext) {
	for target := range work {
		if ctx.Err() != nil {
			state.workers.Done()

			continue
		}

		select {
		case <-ctx.Done():
			state.onError(fmt.Errorf("crawl context canceled: %w", ctx.Err()))
			state.workers.Done()

			continue

		case state.tokens <- struct{}{}:
		}

		go p.processWorker(ctx, target, cat, state)
	}
}

func (p *Provider) processWorker(ctx context.Context, pageURL string, cat *catalog.Catalog, state *crawlContext) {
	defer func() {
		if r := recover(); r != nil {
			state.onError(fmt.Errorf("%w: panic in geo2day worker for %s: %v", ErrFetchCatalog, pageURL, r))
		}

		<-state.tokens
		state.workers.Done()
	}()

	subPages, err := p.fetchAndProcessPage(ctx, pageURL, cat)
	if err != nil {
		state.onError(err)

		return
	}

	for _, page := range subPages {
		state.enqueue(page)
	}
}

func (p *Provider) fetchAndProcessPage(ctx context.Context, currentURL string, cat *catalog.Catalog) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, currentURL, http.NoBody)
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
		return nil, fmt.Errorf("%w: unexpected http %d", ErrFetchCatalog, resp.StatusCode)
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

			return nil, fmt.Errorf("cannot parse html: %w", tokenizer.Err())

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
	if href == "" || shouldSkipHref(href, p.BaseURL, p.StartURL) {
		return ""
	}

	text := extractAnchorText(tokenizer, tokenType)
	fullURL := p.resolveURL(href)

	rawID, ext := splitFileExt(href)
	if shouldSkipRawID(rawID) {
		return ""
	}

	if ext == "html" {
		return p.handleHTMLLink(cat, fullURL, rawID, text)
	}

	if isSupportedFormat(ext) {
		p.processFileLink(cat, href, rawID, ext, text)
	}

	return ""
}

func shouldSkipRawID(rawID string) bool {
	return rawID == "" || strings.HasPrefix(rawID, "#")
}

func extractAnchorText(tokenizer *html.Tokenizer, tokenType html.TokenType) string {
	if tokenType == html.StartTagToken {
		if nextType := tokenizer.Next(); nextType == html.TextToken {
			return strings.TrimSpace(string(tokenizer.Text()))
		}
	}

	return ""
}

func (p *Provider) handleHTMLLink(cat *catalog.Catalog, fullURL, rawID, text string) string {
	if rawID == "index" {
		return ""
	}

	var subPage string
	if strings.HasPrefix(fullURL, p.BaseURL) || strings.HasPrefix(fullURL, p.StartURL) {
		subPage = fullURL
	}

	p.processHTMLLink(cat, fullURL, rawID, text)

	return subPage
}

func shouldSkipHref(href, baseURL, startURL string) bool {
	if strings.HasPrefix(href, "#") ||
		strings.HasPrefix(href, "mailto:") ||
		strings.HasPrefix(href, "javascript:") ||
		strings.HasPrefix(href, "data:") ||
		strings.HasPrefix(href, "vbscript:") ||
		strings.HasPrefix(href, "tel:") {
		return true
	}

	if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
		return !strings.HasPrefix(href, baseURL) && !strings.HasPrefix(href, startURL)
	}

	return false
}

func isSupportedFormat(ext string) bool {
	switch ext {
	case "pbf", "osm.pbf", "poly", "geojson", "md5":
		return true
	default:
		return false
	}
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

const md5HashLength = 32

func isMD5Hash(s string) bool {
	if len(s) != md5HashLength {
		return false
	}

	for _, c := range s {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
			return false
		}
	}

	return true
}

func sanitizeFileName(name string) string {
	name = strings.TrimSpace(name)
	if (strings.HasPrefix(name, "[") && strings.HasSuffix(name, "]")) || isMD5Hash(name) {
		return ""
	}

	return name
}

func (p *Provider) processFileLink(cat *catalog.Catalog, href, rawID, ext, name string) {
	if rawID == "" {
		return
	}

	parent, _ := splitParent(href)
	cleanName := sanitizeFileName(name)

	elem := catalog.Element{
		ID:     rawID,
		Name:   cleanName,
		Parent: parent,
		Meta:   false,
	}

	applyExceptions(&elem)
	_ = cat.MergeElement(&elem)

	formatID := ext

	switch ext {
	case "pbf":
		formatID = catalog.FormatOsmPbf
	case "md5":
		formatID = FormatOsmPbfMd5
	}

	cat.AddExtension(elem.ID, formatID)

	if formatID == catalog.FormatOsmPbf {
		cat.AddExtension(elem.ID, FormatOsmPbfMd5)
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
	if idx := strings.IndexAny(urlStr, "?#"); idx != -1 {
		urlStr = urlStr[:idx]
	}

	parts := strings.Split(urlStr, "/")
	last := parts[len(parts)-1]

	dotParts := strings.Split(last, ".")
	if len(dotParts) < minExtParts {
		return dotParts[0], ""
	}

	return dotParts[0], strings.Join(dotParts[1:], ".")
}

func splitParent(urlStr string) (parent, path string) {
	if idx := strings.IndexAny(urlStr, "?#"); idx != -1 {
		urlStr = urlStr[:idx]
	}

	if idx := strings.Index(urlStr, "://"); idx != -1 {
		urlStr = urlStr[idx+3:]
		if slashIdx := strings.Index(urlStr, "/"); slashIdx != -1 {
			urlStr = urlStr[slashIdx:]
		} else {
			return "", ""
		}
	}

	trimmed := strings.Trim(urlStr, "/")
	if trimmed == "" {
		return "", ""
	}

	parts := strings.Split(trimmed, "/")
	if len(parts) <= 1 {
		return "", ""
	}

	parent = parts[len(parts)-2]
	path = "/" + strings.Join(parts[:len(parts)-1], "/")

	return parent, path
}
