package main

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
)

type IScrapper interface {
	GetConfig() *Config
	Collector(options ...interface{}) *colly.Collector
	Limit() *colly.LimitRule
	GetPB() int
	GetStartURL() string
	ParseFormat(id, format string)
}

//Scrapper define a default scrapper
type Scrapper struct {
	BaseURL          string
	StartURL         string
	Config           *Config       // ptr to Config Element
	PB               int           // For ProgressBar
	Async            bool          // true by default
	DomainGlob       string        // "*" by default
	Parallelism      int           // >1
	RandomDelay      time.Duration // 5 * time.Second by default
	MaxDepth         int           // 0 to infinite
	AllowedDomains   []string
	URLFilters       []*regexp.Regexp
	FormatDefinition formatDefinitions
}

//GetConfig init a *Config from fields
func (s *Scrapper) GetConfig() *Config {
	if s.Config == nil {
		s.Config = &Config{
			Elements:      ElementSlice{},  // should be void
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

//Limit define LimitRules
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

//Collector *colly.Collector // Init Collector
func (s *Scrapper) Collector(options ...interface{}) *colly.Collector {
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains(s.AllowedDomains...),
		colly.URLFilters(s.URLFilters...),
		colly.Async(s.Async),
		colly.MaxDepth(s.MaxDepth),
	)
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
	catch(c.Limit(s.Limit()))
	c.OnError(func(r *colly.Response, err error) {
		if err != colly.ErrForbiddenURL && err != colly.ErrForbiddenDomain && err.Error() != "Forbidden" {
			catch(fmt.Errorf("request URL: %v failed with response: %v\nerror: %v", r.Request.URL, r, err.Error()))
		} else {
			if *fVerbose == true && *fProgress == false && *fQuiet == false {
				log.Printf("URL: %v is forbidden\n", r.Request.URL)
			}
		}
	})
	return c
}

func (s *Scrapper) GetStartURL() string {
	return s.StartURL
}

//GetPB return PB
func (s *Scrapper) GetPB() int {
	return s.PB
}

//ParseFormat add Extentions to ID
func (s *Scrapper) ParseFormat(id, format string) {
	for i := range s.Config.Formats {
		if format == i {
			s.Config.AddExtension(id, format)
		}
	}
}

//FileExt return filename, ext
func FileExt(url string) (string, string) {
	urls := strings.Split(url, "/") // Todo: Try with regexp?
	f := strings.Split(urls[len(urls)-1], ".")
	return f[0], strings.Join(f[1:], ".")
}

//GetParent return filename, path
func GetParent(url string) (string, string) {
	r := strings.Split(url, "/")
	if len(r) < 5 { // <4 should be impossible
		return "", strings.Join(r[:len(r)-1], "/")
	}
	return r[len(r)-2], strings.Join(r[:len(r)-1], "/")
}
