package geo2day_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/provider/geo2day"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockGeo2DayHTML = `<!DOCTYPE html>
<html>
<body>
<table>
  <tr><td><a href="https://geo2day.com/europe/france.pbf">France</a></td></tr>
  <tr><td><a href="https://geo2day.com/europe/france.poly">France Poly</a></td></tr>
  <tr><td><a href="https://geo2day.com/south-america/guyane.pbf">Guyane</a></td></tr>
</table>
</body>
</html>`

func TestGeo2Day_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(mockGeo2DayHTML))
	}))
	defer ts.Close()

	p := geo2day.NewProvider()
	p.StartURL = ts.URL
	p.Client = ts.Client()

	assert.Equal(t, "geo2day", p.Name())
	assert.Equal(t, "geo2day.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("france"))
	fr, exists := cat.Get("france")
	assert.True(t, exists)
	assert.True(t, fr.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, fr.ContainsFormat(catalog.FormatPoly))

	// Check exception mapping: guyane under south-america should become guyane-south-america
	assert.True(t, cat.Exist("guyane-south-america"))
}

func TestGeo2Day_FetchCatalog_Errors(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "server error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	p := geo2day.NewProvider()
	p.StartURL = ts.URL
	p.Client = ts.Client()

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}
