package generator //nolint:testpackage // testing internal functions

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/julien-noblet/download-geofabrik/internal/provider"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

type mockGeneratorProvider struct {
	catalogData *catalog.Catalog
	name        string
	shouldErr   bool
}

func (m *mockGeneratorProvider) Name() string              { return m.name }
func (m *mockGeneratorProvider) Description() string       { return "Mock provider for generator tests" }
func (m *mockGeneratorProvider) DefaultConfigFile() string { return m.name + ".yml" }
func (m *mockGeneratorProvider) FetchCatalog(_ context.Context) (*catalog.Catalog, error) {
	if m.shouldErr {
		return nil, errors.New("mock fetch failed")
	}

	return m.catalogData, nil
}

func TestGenerate_Unknown(t *testing.T) {
	t.Parallel()

	err := Generate(t.Context(), "unknown_service", "/tmp/dummy.yml")
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrUnknownService)
}

func TestGenerate_Success(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	cat.Elements["test-elem"] = catalog.Element{
		ID:      "test-elem",
		Name:    "Test Element",
		Formats: []string{catalog.FormatPoly, catalog.FormatOsmPbf},
	}

	prov := &mockGeneratorProvider{
		name:        "mock-gen-success",
		catalogData: cat,
	}
	provider.Register(prov)

	tmpDir := t.TempDir()
	outFile := filepath.Join(tmpDir, "out.yml")

	err := Generate(t.Context(), "mock-gen-success", outFile)
	require.NoError(t, err)

	loaded, err := catalog.LoadFile(outFile)
	require.NoError(t, err)
	assert.Contains(t, loaded.Elements, "test-elem")
	assert.Equal(t, catalog.Formats{catalog.FormatOsmPbf, catalog.FormatPoly}, loaded.Elements["test-elem"].Formats)
}

func TestGenerate_FetchError(t *testing.T) {
	t.Parallel()

	prov := &mockGeneratorProvider{
		name:      "mock-gen-err",
		shouldErr: true,
	}
	provider.Register(prov)

	tmpDir := t.TempDir()
	outFile := filepath.Join(tmpDir, "out.yml")

	err := Generate(t.Context(), "mock-gen-err", outFile)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mock fetch failed")
}

func TestGenerate_SaveError(t *testing.T) {
	t.Parallel()

	cat := catalog.New()
	prov := &mockGeneratorProvider{
		name:        "mock-gen-save-err",
		catalogData: cat,
	}
	provider.Register(prov)

	// An impossible directory path
	outFile := "/dev/null/impossible/path.yml"

	err := Generate(t.Context(), "mock-gen-save-err", outFile)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to write config")
}
