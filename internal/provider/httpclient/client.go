package httpclient

import (
	"net"
	"net/http"
	"time"
)

const (
	defaultTimeout             = 30 * time.Second
	defaultKeepAlive           = 30 * time.Second
	defaultIdleTimeout         = 90 * time.Second
	defaultMaxIdleConns        = 20
	defaultMaxIdleConnsPerHost = 10
	crawlerMaxIdleConns        = 100
	crawlerConnsMultiplier     = 2
)

// TransportConfig holds parameters for configuring the provider HTTP transport.
type TransportConfig struct {
	Timeout             time.Duration
	KeepAlive           time.Duration
	IdleTimeout         time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	MaxConnsPerHost     int
}

// DefaultTransportConfig returns the standard configuration used by default scrapers.
func DefaultTransportConfig() TransportConfig {
	return TransportConfig{
		Timeout:             defaultTimeout,
		KeepAlive:           defaultKeepAlive,
		IdleTimeout:         defaultIdleTimeout,
		MaxIdleConns:        defaultMaxIdleConns,
		MaxIdleConnsPerHost: defaultMaxIdleConnsPerHost,
		MaxConnsPerHost:     0,
	}
}

// New creates a pre-configured *http.Client optimized for scraping provider catalogs with default 30s timeout.
func New() *http.Client {
	return NewWithConfig(DefaultTransportConfig())
}

// NewWithTimeout creates a pre-configured *http.Client with a custom dialer timeout.
func NewWithTimeout(timeout time.Duration) *http.Client {
	cfg := DefaultTransportConfig()
	cfg.Timeout = timeout

	return NewWithConfig(cfg)
}

// NewCrawlerClient creates a pre-configured *http.Client tuned for high-concurrency crawlers with default 30s timeout.
func NewCrawlerClient(concurrencyLimit int) *http.Client {
	return NewCrawlerClientWithTimeout(defaultTimeout, concurrencyLimit)
}

// NewCrawlerClientWithTimeout creates a pre-configured *http.Client tuned for high-concurrency crawlers with custom timeout.
func NewCrawlerClientWithTimeout(timeout time.Duration, concurrencyLimit int) *http.Client {
	return NewWithConfig(TransportConfig{
		Timeout:             timeout,
		KeepAlive:           defaultKeepAlive,
		IdleTimeout:         defaultIdleTimeout,
		MaxIdleConns:        crawlerMaxIdleConns,
		MaxIdleConnsPerHost: concurrencyLimit,
		MaxConnsPerHost:     concurrencyLimit * crawlerConnsMultiplier,
	})
}

// NewWithConfig creates a pre-configured *http.Client with the given transport configuration.
func NewWithConfig(cfg TransportConfig) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   cfg.Timeout,
				KeepAlive: cfg.KeepAlive,
			}).DialContext,
			MaxIdleConns:        cfg.MaxIdleConns,
			MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
			MaxConnsPerHost:     cfg.MaxConnsPerHost,
			IdleConnTimeout:     cfg.IdleTimeout,
			ForceAttemptHTTP2:   true,
		},
	}
}
