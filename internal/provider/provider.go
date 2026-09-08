package provider

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"slices"
	"sync"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var (
	// ErrProviderNotFound is returned when a requested provider name is not in the registry.
	ErrProviderNotFound = errors.New("provider not found")

	// ErrFetchCatalog is returned when a provider fails to fetch or parse remote catalog data.
	ErrFetchCatalog = errors.New("failed to fetch catalog")
)

// Provider defines the standard interface for an OSM data catalog provider.
type Provider interface {
	// Name returns the unique service identifier for the provider (e.g. "geofabrik").
	Name() string

	// Description returns a human-readable summary of the provider and its coverage.
	Description() string

	// DefaultConfigFile returns the default YAML filename used to cache the provider catalog.
	DefaultConfigFile() string

	// FetchCatalog scrapes or queries the provider API to build and return a complete *catalog.Catalog.
	// It must respect context cancellation and timeouts.
	FetchCatalog(ctx context.Context) (*catalog.Catalog, error)
}

var (
	registryMu sync.RWMutex
	registry   = make(map[string]Provider)
)

// Register registers a provider in the global registry.
func Register(prov Provider) {
	if prov == nil || prov.Name() == "" {
		return
	}

	registryMu.Lock()
	defer registryMu.Unlock()

	registry[prov.Name()] = prov
}

// Get retrieves a registered provider by name.
//
//nolint:ireturn // Factory function returning interface
func Get(name string) (Provider, error) {
	registryMu.RLock()
	defer registryMu.RUnlock()

	prov, exists := registry[name]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrProviderNotFound, name)
	}

	return prov, nil
}

// List returns all registered provider names sorted alphabetically.
func List() []string {
	registryMu.RLock()
	defer registryMu.RUnlock()

	return slices.Sorted(maps.Keys(registry))
}
