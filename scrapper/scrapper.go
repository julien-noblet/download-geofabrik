package scrapper

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
)

// IScrapper represent a colly Scrapper
type IScrapper interface {
	GetConfig() *config.Config
	Collector(options ...interface{}) *colly.Collector
	Limit() *colly.LimitRule
	GetPB() int
	GetStartURL() string
	ParseFormat(id, format string)
}

// Scrapper define a default scrapper
type Scrapper struct {
	BaseURL          string
	StartURL         string
	Config           *config.Config // ptr to Config Element
	PB               int            // For ProgressBar
	Async            bool           // true by default
	DomainGlob       string         // "*" by default
	Parallelism      int            // >1
	RandomDelay      time.Duration  // 5 * time.Second by default
	MaxDepth         int            // 0 to infinite
	AllowedDomains   []string
	URLFilters       []*regexp.Regexp
	FormatDefinition formats.FormatDefinitions
	Timeout          *time.Duration
}

// GetConfig init a *config.Config from fields
func (s *Scrapper) GetConfig() *config.Config {
	if s.Config == nil {
		s.Config = &config.Config{
			Elements:      element.Slice{}, // should be void
			ElementsMutex: &sync.RWMutex{}, // initialize a new Mutex
		}
	}

	if s.BaseURL != "" {
		s.Config.BaseURL = s.BaseURL
	}

	if s.FormatDefinition != nil {
		s.Config.Formats = s.FormatDefinition
	}

	return s.Config // in case of BaseURL AND/OR FormatDefinition isn't set
}

// Limit define LimitRules
func (s *Scrapper) Limit() *colly.LimitRule {
	if s.DomainGlob == "" {
		s.DomainGlob = "*"
	}

	if s.Parallelism <= 1 { // not //
		s.Parallelism = 1
	}

	if s.RandomDelay == 0 {
		s.RandomDelay = 5 * time.Second
	}

	return &colly.LimitRule{
		DomainGlob:  s.DomainGlob,
		Parallelism: s.Parallelism,
		RandomDelay: s.RandomDelay,
	}
}

// Collector *colly.Collector Init Collector
func (s *Scrapper) Collector(options ...interface{}) *colly.Collector { //nolint:unparam
	c := colly.NewCollector(
		colly.AllowedDomains(s.AllowedDomains...),
		colly.URLFilters(s.URLFilters...),
		colly.Async(s.Async),
		colly.MaxDepth(s.MaxDepth),
	)

	if s.Timeout != nil {
		c.SetRequestTimeout(*s.Timeout)
	}

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          0,
		IdleConnTimeout:       5 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 5 * time.Second,
	})

	s.Config = s.GetConfig() // ensure initialisation
	if err := c.Limit(s.Limit()); err != nil {
		log.Panicln(err)
	}

	c.OnError(func(r *colly.Response, err error) {
		if err != colly.ErrForbiddenURL && err != colly.ErrForbiddenDomain && err.Error() != "Forbidden" {
			log.Panicln(fmt.Errorf("request URL: %v failed with response: %v\nerror: %v", r.Request.URL, r, err.Error()))
		} else if viper.GetBool("verbose") && !viper.GetBool("quiet") && !viper.GetBool("progress") {
			log.Printf("URL: %v is forbidden\n", r.Request.URL)
		}
	})

	return c
}

// GetStartURL return StartURL
func (s *Scrapper) GetStartURL() string {
	return s.StartURL
}

// GetPB return PB
func (s *Scrapper) GetPB() int {
	return s.PB
}

// ParseFormat add Extensions to ID
func (s *Scrapper) ParseFormat(id, format string) {
	for i := range s.Config.Formats {
		if format == i {
			s.Config.AddExtension(id, format)
		}
	}
}

// FileExt return filename, ext
func FileExt(url string) (filename, extension string) {
	urls := strings.Split(url, "/") // Todo: Try with regexp?
	f := strings.Split(urls[len(urls)-1], ".")

	return f[0], strings.Join(f[1:], ".")
}

// GetParent return filename, path
func GetParent(url string) (filename, path string) {
	r := strings.Split(url, "/")
	if len(r) < 5 { // <4 should be impossible
		return "", strings.Join(r[:len(r)-1], "/")
	}

	return r[len(r)-2], strings.Join(r[:len(r)-1], "/")
}
