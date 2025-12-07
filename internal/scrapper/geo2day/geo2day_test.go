package geo2day_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/internal/scrapper/geo2day"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeo2day_Exceptions(t *testing.T) {
	t.Parallel()

	type args struct {
		e element.Element
	}

	tests := []struct {
		name string
		args args
		want element.Element
	}{
		// TODO: Add test cases.
		{
			name: "la_rioja in argentina",
			args: args{
				e: element.Element{
					ID:     "la_rioja",
					Parent: "argentina",
					File:   "la_rioja",
				},
			},
			want: element.Element{
				ID:     "la_rioja-argentina",
				Parent: "argentina",
				File:   "la_rioja",
			},
		},
		{
			name: "la_rioja in spain",
			args: args{
				e: element.Element{
					ID:     "la_rioja",
					Parent: "spain",
					File:   "la_rioja",
				},
			},
			want: element.Element{
				ID:     "la_rioja-spain",
				Parent: "spain",
				File:   "la_rioja",
			},
		},
		{
			name: "guyane in france",
			args: args{
				e: element.Element{
					ID:     "guyane",
					Parent: "france",
					File:   "guyane",
				},
			},
			want: element.Element{
				ID:     "guyane-france",
				Parent: "france",
				File:   "guyane",
			},
		},
		{
			name: "guyane in south-america",
			args: args{
				e: element.Element{
					ID:     "guyane",
					Parent: "south-america",
					File:   "guyane",
				},
			},
			want: element.Element{
				ID:     "guyane-south-america",
				Parent: "south-america",
				File:   "guyane",
			},
		},
		{
			name: "sevastopol in ukraine",
			args: args{
				e: element.Element{
					ID:     "sevastopol",
					Parent: "ukraine",
					File:   "sevastopol",
				},
			},
			want: element.Element{
				ID:     "sevastopol-ukraine",
				Parent: "ukraine",
				File:   "sevastopol",
			},
		},
		{
			name: "sevastopol in russia",
			args: args{
				e: element.Element{
					ID:     "sevastopol",
					Parent: "russia",
					File:   "sevastopol",
				},
			},
			want: element.Element{
				ID:     "sevastopol-russia",
				Parent: "russia",
				File:   "sevastopol",
			},
		},
		{
			name: "limburg in netherlands",
			args: args{
				e: element.Element{
					ID:     "limburg",
					Parent: "netherlands",
					File:   "limburg",
				},
			},
			want: element.Element{
				ID:     "limburg-netherlands",
				Parent: "netherlands",
				File:   "limburg",
			},
		},
		{
			name: "limburg in flanders",
			args: args{
				e: element.Element{
					ID:     "limburg",
					Parent: "flanders",
					File:   "limburg",
				},
			},
			want: element.Element{
				ID:     "limburg-flanders",
				Parent: "flanders",
				File:   "limburg",
			},
		},
		{
			name: "cordoba in argentina",
			args: args{
				e: element.Element{
					ID:     "cordoba",
					Parent: "argentina",
					File:   "cordoba",
				},
			},
			want: element.Element{
				ID:     "cordoba-argentina",
				Parent: "argentina",
				File:   "cordoba",
			},
		},
		{
			name: "cordoba in andalucia",
			args: args{
				e: element.Element{
					ID:     "cordoba",
					Parent: "andalucia",
					File:   "cordoba",
				},
			},
			want: element.Element{
				ID:     "cordoba-andalucia",
				Parent: "andalucia",
				File:   "cordoba",
			},
		},
		{
			name: "georgia in usa",
			args: args{
				e: element.Element{
					ID:     "georgia",
					Parent: "us",
					File:   "georgia",
				},
			},
			want: element.Element{
				ID:     "georgia-us",
				Parent: "us",
				File:   "georgia",
			},
		},
		{
			name: "georgia in asia",
			args: args{
				e: element.Element{
					ID:     "georgia",
					Parent: "asia",
					File:   "georgia",
				},
			},
			want: element.Element{
				ID:     "georgia-asia",
				Parent: "asia",
				File:   "georgia",
			},
		},
		{
			name: "france is not in the list",
			args: args{
				e: element.Element{
					ID:     "france",
					Parent: "europe",
					File:   "france",
				},
			},
			want: element.Element{
				ID:     "france",
				Parent: "europe",
				File:   "france",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			g := geo2day.GetDefault()

			if got := g.Exceptions(&tt.args.e); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("Geo2day.Exceptions() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestGeo2day_Collector(t *testing.T) {
	// Mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			fmt.Fprintln(w, `
				<html>
				<body>
					<table>
						<td>
							<a href="/europe.html">Europe</a>
							<a href="/planet.osm.pbf">Planet PBF</a>
						</td>
					</table>
					<div class="row">
						<a href="/planet.osm.pbf">Planet PBF</a>
					</div>
				</body>
				</html>
			`)

		case "/europe.html":
			fmt.Fprintln(w, `
				<html>
				<body>
					<table>
						<td>
                            <a href="/europe/france.html">France</a>
                            <a href="/europe/germany.osm.pbf">Germany</a>
						</td>
					</table>
				</body>
				</html>
			`)

		case "/europe/france.html":
			fmt.Fprintln(w, `
				<html>
				<body>
					<table>
						<td>
                            <a href="/europe/france/ile-de-france.osm.pbf">Ile-de-France</a>
						</td>
					</table>
				</body>
				</html>
			`)

		default:
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	g := geo2day.GetDefault()
	g.BaseURL = ts.URL
	g.StartURL = ts.URL + "/"
	g.AllowedDomains = nil // Allow all domains (specifically the test server)
	g.URLFilters = nil     // Disable URL filters to allow test server URLs
	g.Parallelism = 1      // Process sequentially to avoid race in test logic if any (though map access needs sync)

	c := g.Collector()

	// Visit the start URL
	err := c.Visit(g.StartURL)
	require.NoError(t, err)
	c.Wait()

	// Helper to get element by key safely
	getElement := func(key string) (element.Element, bool) {
		g.Config.ElementsMutex.RLock()
		defer g.Config.ElementsMutex.RUnlock() // Defer unlock

		el, ok := g.Config.Elements[key]

		return el, ok
	}

	// 1. Check Planet PBF (Attached to root element "")
	rootEl, found := getElement("")
	assert.True(t, found, "root element should exist")

	if found {
		assert.Contains(t, rootEl.Formats, formats.FormatOsmPbf)
	}

	// 2. Check Europe (HTML)
	europe, found := getElement("europe")
	assert.True(t, found, "europe element should exist")
	assert.Equal(t, "europe", europe.ID)
	// 3. Check traversal to France
	france, found := getElement("france")
	assert.True(t, found, "france element should exist")

	if found {
		assert.Equal(t, "france", france.ID)
		assert.Equal(t, "europe", france.Parent)
	}

	// 4. Check contents within France (Ile-de-france)
	// ID "ile-de-france"
	idf, found := getElement("ile-de-france")
	assert.True(t, found, "ile-de-france element should exist")

	if found {
		// Parent should be "france"
		assert.Equal(t, "france", idf.Parent)
	}
}
