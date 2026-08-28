package provider_test

import (
	"context"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/provider"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockProvider struct {
	name string
}

func (m *mockProvider) Name() string              { return m.name }
func (m *mockProvider) Description() string       { return "Mock provider" }
func (m *mockProvider) DefaultConfigFile() string { return "mock.yml" }
func (m *mockProvider) FetchCatalog(_ context.Context) (*catalog.Catalog, error) {
	return catalog.New(), nil
}

func TestProvider_Registry(t *testing.T) {
	t.Parallel()

	mock := &mockProvider{name: "mock_test_provider"}
	provider.Register(mock)

	p, err := provider.Get("mock_test_provider")
	require.NoError(t, err)
	assert.Equal(t, "mock_test_provider", p.Name())
	assert.Equal(t, "mock.yml", p.DefaultConfigFile())
	assert.Equal(t, "Mock provider", p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, cat)

	_, err = provider.Get("non_existent_provider_xyz")
	require.Error(t, err)
	require.ErrorIs(t, err, provider.ErrProviderNotFound)

	list := provider.List()
	assert.Contains(t, list, "mock_test_provider")

	// Nil / empty provider register
	provider.Register(nil)
	provider.Register(&mockProvider{name: ""})
}

func TestRegisterDefaultProviders(t *testing.T) {
	t.Parallel()

	provider.RegisterDefaultProviders()

	list := provider.List()
	assert.Contains(t, list, "geofabrik")
	assert.Contains(t, list, "openstreetmap.fr")
	assert.Contains(t, list, "bbbike")
	assert.Contains(t, list, "geo2day")
	assert.Contains(t, list, "movisda")
	assert.Contains(t, list, "planet.osm.ch")
	assert.Contains(t, list, "osm.kewl.lu")
	assert.Contains(t, list, "osm.fit.vutbr.cz")
	assert.Contains(t, list, "osmit-estratti")
	assert.Contains(t, list, "osm.kcwu.csie.org")
}
