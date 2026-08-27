package openstreetmapfr_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mockOSMFRIndexHTML = `<!DOCTYPE html>
<html>
<body>
<table>
  <tr><td><a href="europe/">europe/</a></td></tr>
</table>
</body>
</html>`

const mockOSMFREuropeHTML = `<!DOCTYPE html>
<html>
<body>
<table>
  <tr><td><a href="france.osm.pbf">france.osm.pbf</a></td></tr>
  <tr><td><a href="france.poly">france.poly</a></td></tr>
  <tr><td><a href="monaco.osm.pbf">monaco.osm.pbf</a></td></tr>
</table>
</body>
</html>`

func TestOSMFR_FetchCatalog(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		if r.URL.Path == "/extracts/" || r.URL.Path == "/" {
			_, _ = w.Write([]byte(mockOSMFRIndexHTML))
		} else {
			_, _ = w.Write([]byte(mockOSMFREuropeHTML))
		}
	}))
	defer ts.Close()

	p := openstreetmapfr.NewProvider()
	p.StartURL = ts.URL + "/extracts/"
	p.Client = ts.Client()

	assert.Equal(t, "openstreetmap.fr", p.Name())
	assert.Equal(t, "openstreetmap.fr.yml", p.DefaultConfigFile())
	assert.NotEmpty(t, p.Description())

	cat, err := p.FetchCatalog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cat)

	assert.True(t, cat.Exist("france"))
	assert.True(t, cat.Exist("monaco"))

	fr, exists := cat.Get("france")
	assert.True(t, exists)
	assert.True(t, fr.ContainsFormat(catalog.FormatOsmPbf))
	assert.True(t, fr.ContainsFormat(catalog.FormatPoly))
}

func TestOSMFR_FetchCatalog_Errors(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "server error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	p := openstreetmapfr.NewProvider()
	p.StartURL = ts.URL + "/extracts/"
	p.Client = ts.Client()

	_, err := p.FetchCatalog(context.Background())
	require.Error(t, err)
}
