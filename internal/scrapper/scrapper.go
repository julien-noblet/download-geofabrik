package scrapper

import (
	"errors"
	"log/slog"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
)

// IScrapper represents a colly Scrapper.
type IScrapper interface {
	GetConfig() *config.Config
	Collector() *colly.Collector
	Limit() *colly.LimitRule
	GetPB() int
	GetStartURL() string
	ParseFormat(id, format string)
}

// Scrapper defines a default scrapper.
type Scrapper struct {
	Config           *config.Config
	FormatDefinition formats.FormatDefinitions
	BaseURL          string
	DomainGlob       string
	StartURL         string
	URLFilters       []*regexp.Regexp
	AllowedDomains   []string
	Timeout          time.Duration
	RandomDelay      time.Duration
	MaxDepth         int
	Parallelism      int
	PB               int
	mu               sync.RWMutex
	Async            bool
}

const (
	defaultRandomDelay           = 5 * time.Second
	defaultTimeout               = 60 * time.Second
	defaultKeepAlive             = 30 * time.Second
	defaultIdleConnTimeout       = 5 * time.Second
	defaultTLSHandshakeTimeout   = 10 * time.Second
	defaultExpectContinueTimeout = 5 * time.Second
	minParentListLength          = 5
)

// NewScrapper creates a new Scrapper instance with default values.
func NewScrapper(baseURL, startURL string, allowedDomains []string) *Scrapper {
	return &Scrapper{
		RandomDelay:    defaultRandomDelay,
		Timeout:        defaultTimeout,
		Parallelism:    1,
		BaseURL:        baseURL,
		StartURL:       startURL,
		AllowedDomains: allowedDomains,
	}
}

// GetConfig initializes a *config.Config from fields.
func (s *Scrapper) GetConfig() *config.Config {
	s.mu.RLock()

	if s.Config != nil {
		defer s.mu.RUnlock()

		return s.Config
	}

	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	s.Config = s.initializeConfig()

	return s.Config
}

// initializeConfig initializes the configuration with default values.
func (s *Scrapper) initializeConfig() *config.Config {
	return &config.Config{
		Elements:      element.MapElement{},
		ElementsMutex: &sync.RWMutex{},
		Formats:       s.FormatDefinition,
		BaseURL:       s.BaseURL,
	}
}

// Limit defines LimitRules.
func (s *Scrapper) Limit() *colly.LimitRule {
	if s.DomainGlob == "" {
		s.DomainGlob = "*"
	}

	if s.Parallelism <= 1 {
		s.Parallelism = 1
	}

	return &colly.LimitRule{
		DomainGlob:  s.DomainGlob,
		Parallelism: s.Parallelism,
		RandomDelay: s.RandomDelay,
	}
}

// Collector initializes a *colly.Collector.
func (s *Scrapper) Collector(_ ...interface{}) *colly.Collector {
	myCollector := colly.NewCollector(
		colly.AllowedDomains(s.AllowedDomains...),
		colly.URLFilters(s.URLFilters...),
		colly.Async(s.Async),
		colly.MaxDepth(s.MaxDepth),
	)

	if s.Timeout != 0 {
		myCollector.SetRequestTimeout(s.Timeout)
	}

	myCollector.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   defaultTimeout,
			KeepAlive: defaultKeepAlive,
		}).DialContext,
		IdleConnTimeout:       defaultIdleConnTimeout,
		TLSHandshakeTimeout:   defaultTLSHandshakeTimeout,
		ExpectContinueTimeout: defaultExpectContinueTimeout,
	})

	s.Config = s.GetConfig()
	if err := myCollector.Limit(s.Limit()); err != nil {
		slog.Error("Can't update limit", "error", err)
	}

	myCollector.OnError(func(r *colly.Response, err error) {
		if !errors.Is(err, colly.ErrForbiddenURL) && !errors.Is(err, colly.ErrForbiddenDomain) && err.Error() != "Forbidden" {
			slog.Debug("Request failed", "url", r.Request.URL, "error", err)
		} else {
			slog.Debug("Forbidden URL", "url", r.Request.URL)
		}
	})

	return myCollector
}

// GetStartURL returns StartURL.
func (s *Scrapper) GetStartURL() string {
	return s.StartURL
}

// GetPB returns PB.
func (s *Scrapper) GetPB() int {
	return s.PB
}

// ParseFormat adds Extensions to ID.
func (s *Scrapper) ParseFormat(id, format string) {
	s.AddExtension(id, format, &s.Config.Formats)
}

// ParseFormatService adds Extensions to ID.
func (s *Scrapper) ParseFormatService(id, format string, def *formats.FormatDefinitions) {
	s.AddExtension(id, format, def)
}

// AddExtension adds an extension to the configuration.
func (s *Scrapper) AddExtension(id, format string, def *formats.FormatDefinitions) {
	for f, i := range *def {
		if format == i.ID {
			s.Config.AddExtension(id, f)
		}
	}
}

// FileExt returns filename and extension.
func FileExt(url string) (filename, extension string) {
	urls := strings.Split(url, "/")
	f := strings.Split(urls[len(urls)-1], ".")

	return f[0], strings.Join(f[1:], ".")
}

// GetParent returns filename and path.
func GetParent(url string) (filename, path string) {
	r := strings.Split(url, "/")
	if len(r) < minParentListLength {
		return "", strings.Join(r[:len(r)-1], "/")
	}

	return r[len(r)-2], strings.Join(r[:len(r)-1], "/")
}
