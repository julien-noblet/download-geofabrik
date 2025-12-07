package geofabrik_test

import (
	"net/url"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/internal/scrapper"
	"github.com/julien-noblet/download-geofabrik/internal/scrapper/geofabrik"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
)

func Test_geofabrikParseFormat(t *testing.T) {
	t.Parallel()

	type args struct {
		id     string
		format string
		c      config.Config
	}

	tests := []struct {
		args args
		want element.MapElement
		name string
	}{
		{
			name: "Add osm.pbf but already in",
			args: args{
				id:     "a",
				format: formats.FormatOsmPbf,
				c: config.Config{
					Elements: element.MapElement{
						"a": element.Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formats.FormatOsmPbf},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: element.MapElement{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf",
			args: args{
				id:     "a",
				format: formats.FormatOsmPbf,
				c: config.Config{
					Elements: element.MapElement{
						"a": element.Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: element.MapElement{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf.md5",
			args: args{
				id:     "a",
				format: "osm.pbf.md5",
				c: config.Config{
					Elements: element.MapElement{
						"a": element.Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: element.MapElement{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState, "osm.pbf.md5"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.bz2",
			args: args{
				id:     "a",
				format: formats.FormatOsmBz2,
				c: config.Config{
					Elements: element.MapElement{
						"a": element.Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: element.MapElement{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState, formats.FormatOsmBz2},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.bz2.md5",
			args: args{
				id:     "a",
				format: "osm.bz2.md5",
				c: config.Config{
					Elements: element.MapElement{
						"a": element.Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: element.MapElement{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState, "osm.bz2.md5"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add poly",
			args: args{
				id:     "a",
				format: formats.FormatPoly,
				c: config.Config{
					Elements: element.MapElement{
						"a": element.Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: element.MapElement{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState, formats.FormatPoly},
					Meta:    false,
				},
			},
		},
		{
			name: "Add shp.zip",
			args: args{
				id:     "a",
				format: formats.FormatShpZip,
				c: config.Config{
					Elements: element.MapElement{
						"a": element.Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: element.MapElement{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState, formats.FormatShpZip},
					Meta:    false,
				},
			},
		},
		{
			name: "Add unk format",
			args: args{
				id:     "a",
				format: "unk",
				c: config.Config{
					Elements: element.MapElement{
						"a": element.Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: element.MapElement{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf on meta",
			args: args{
				id:     "a",
				format: formats.FormatOsmPbf,
				c: config.Config{
					Elements: element.MapElement{
						"a": element.Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{},
							Meta:    true,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: element.MapElement{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
					Meta:    false,
				},
			},
		},
	}

	for thisTest := range tests {
		t.Run(tests[thisTest].name, func(t *testing.T) {
			t.Parallel()

			myScrapper := geofabrik.Geofabrik{Scrapper: &scrapper.Scrapper{}}
			myScrapper.Config = &tests[thisTest].args.c
			myScrapper.Config.Formats = formats.FormatDefinitions{
				formats.FormatOshPbf: {ID: formats.FormatOshPbf, Loc: ".osh.pbf"},
				"osh.pbf.md5":        {ID: "osh.pbf.md5", Loc: ".osh.pbf.md5"},
				formats.FormatOsmBz2: {ID: formats.FormatOsmBz2, Loc: "-latest.osm.bz2"},
				"osm.bz2.md5":        {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
				formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: "-latest.osm.pbf"},
				"osm.pbf.md5":        {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
				formats.FormatPoly:   {ID: formats.FormatPoly, Loc: ".poly"},
				formats.FormatKml:    {ID: formats.FormatKml, Loc: ".kml"},
				formats.FormatState:  {ID: formats.FormatState, Loc: "-updates/state.txt"},
				formats.FormatShpZip: {ID: formats.FormatShpZip, Loc: "-latest-free.shp.zip"},
			}

			myScrapper.ParseFormat(tests[thisTest].args.id, tests[thisTest].args.format)

			if !reflect.DeepEqual(tests[thisTest].args.c.Elements, tests[thisTest].want) {
				t.Errorf("ParseFormat() got %v, want %v", tests[thisTest].args.c.Elements, tests[thisTest].want)
			}
		})
	}
}

func TestGeofabrik_parseLi(t *testing.T) {
	t.Parallel()

	tests := []struct {
		want    element.MapElement
		element *element.Element
		name    string
		html    string
		url     string
	}{
		{
			name: "sample",
			html: `
			<li>
				<p>some text</p>
				<a href="toto.osm.pbf">anotherText</a>
			</li>`,
			url:     `https://download.geofabrik.de/toto.html`,
			element: &element.Element{ID: "toto"},
			want: element.MapElement{"toto": element.Element{
				ID:      "toto",
				Formats: element.Formats{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
			}},
		},
		{
			name: "georgia-us",
			html: `
			<li>
				<p>some text</p>
				<a href="georgia.osm.pbf">anotherText</a>
			</li>`,
			url:     `https://download.geofabrik.de/north-america/us/georgia.html`,
			element: &element.Element{ID: "georgia-us"},
			want: element.MapElement{"georgia-us": element.Element{
				ID:      "georgia-us",
				Formats: element.Formats{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
			}},
		},
		{
			name: "georgia-eu",
			html: `
			<li>
				<p>some text</p>
				<a href="georgia.osm.pbf">anotherText</a>
			</li>`,
			url:     `https://download.geofabrik.de/europe/georgia.html`,
			element: &element.Element{ID: "georgia-eu"},
			want: element.MapElement{"georgia-eu": element.Element{
				ID:      "georgia-eu",
				Formats: element.Formats{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
			}},
		},
		{
			name: "guatemala-south-america",
			html: `
			<li>
				<p>some text</p>
				<a href="guatemala.osm.pbf">anotherText</a>
			</li>`,
			url:     `https://download.geofabrik.de/south-america/guatemala.html`,
			element: &element.Element{ID: "guatemala-south-america"},
			want: element.MapElement{"guatemala-south-america": element.Element{
				ID:      "guatemala-south-america",
				Formats: element.Formats{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState},
			}},
		},
		// TODO: Add test cases.
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(thisTest.html))
			myURL, _ := url.Parse(thisTest.url)
			myElement := &colly.HTMLElement{
				DOM: dom.Selection,
				Response: &colly.Response{
					Request: &colly.Request{URL: myURL},
				},
			}
			defaultGeofabrik := geofabrik.GetDefault()

			defaultGeofabrik.GetConfig()

			if err := defaultGeofabrik.Config.MergeElement(thisTest.element); err != nil {
				t.Errorf("Bad tests g.Config.mergeElement() can't merge %v - %v", thisTest.element, err)
			}

			defaultGeofabrik.ParseLi(myElement, nil)

			if !reflect.DeepEqual(defaultGeofabrik.Config.Elements, thisTest.want) {
				t.Errorf("parseLi() fail, got %v, want %v", defaultGeofabrik.Config.Elements, thisTest.want)
			}
		})
	}
}

//nolint:lll // this func contain html extracts
func TestGeofabrik_parseSubregion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		want     element.MapElement
		elements *element.MapElement
		name     string
		html     string
		url      string
	}{
		{
			name: "sample1",
			html: `
				<table id="subregions">
				<tbody><tr>
				<th class="subregion">Sub Region</th>
				<th colspan="4">Quick Links</th>
				</tr>
				<tr>
				<th>&nbsp;</th>
				<th colspan="2">.osm.pbf</th>
				<th>.shp.zip</th>
				<th>.osm.bz2</th>
				</tr>
				<tr onmouseover="loadkml('africa.kml')"><td class="subregion"><a href="africa.html">Africa</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="africa-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(2.9&nbsp;GB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="africa-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('antarctica.kml')"><td class="subregion"><a href="antarctica.html">Antarctica</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="antarctica-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(29.0&nbsp;MB)</td><td> <a href="antarctica-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="antarctica-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('asia.kml')"><td class="subregion"><a href="asia.html">Asia</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="asia-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(6.8&nbsp;GB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="asia-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('australia-oceania.kml')"><td class="subregion"><a href="australia-oceania.html">Australia and Oceania</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="australia-oceania-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(644&nbsp;MB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="australia-oceania-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('central-america.kml')"><td class="subregion"><a href="central-america.html">Central America</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="central-america-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(345&nbsp;MB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="central-america-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('europe.kml')"><td class="subregion"><a href="europe.html">Europe</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="europe-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(19.5&nbsp;GB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="europe-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('north-america.kml')"><td class="subregion"><a href="north-america.html">North America</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="north-america-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(8.3&nbsp;GB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="north-america-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('south-america.kml')"><td class="subregion"><a href="south-america.html">South America</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(1.5&nbsp;GB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="south-america-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				</tbody></table>`,
			url:      `https://download.geofabrik.de/`,
			elements: &element.MapElement{},
			want: element.MapElement{
				"africa":            element.Element{ID: "africa", Name: "Africa", Meta: true},
				"antarctica":        element.Element{ID: "antarctica", Name: "Antarctica", Meta: true},
				"asia":              element.Element{ID: "asia", Name: "Asia", Meta: true},
				"australia-oceania": element.Element{ID: "australia-oceania", Name: "Australia and Oceania", Meta: true},
				"central-america":   element.Element{ID: "central-america", Name: "Central America", Meta: true},
				"europe":            element.Element{ID: "europe", Name: "Europe", Meta: true},
				"north-america":     element.Element{ID: "north-america", Name: "North America", Meta: true},
				"south-america":     element.Element{ID: "south-america", Name: "South America", Meta: true},
			},
		},
		{
			name: "sample2 - Europe extract",
			html: `
				<table id="subregions">
				<tbody><tr>
				<th class="subregion">Sub Region</th>
				<th colspan="4">Quick Links</th>
				</tr>
				<tr>
				<th>&nbsp;</th>
				<th colspan="2">.osm.pbf</th>
				<th>.shp.zip</th>
				<th>.osm.bz2</th>
				</tr>
				<tr onmouseover="loadkml('europe/albania.kml')"><td class="subregion"><a href="europe/albania.html">Albania</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="europe/albania-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(31.2&nbsp;MB)</td><td> <a href="europe/albania-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="europe/albania-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('europe/andorra.kml')"><td class="subregion"><a href="europe/andorra.html">Andorra</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="europe/andorra-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(1.4&nbsp;MB)</td><td> <a href="europe/andorra-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="europe/andorra-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('europe/france.kml')"><td class="subregion"><a href="europe/france.html">France</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="europe/france-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(3.4&nbsp;GB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="europe/france-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('europe/georgia.kml')"><td class="subregion"><a href="europe/georgia.html">Georgia (Eastern Europe)</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="europe/georgia-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(45.2&nbsp;MB)</td><td> <a href="europe/georgia-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="europe/georgia-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				<tr onmouseover="loadkml('europe/germany.kml')"><td class="subregion"><a href="europe/germany.html">Germany</a></td>
				<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="europe/germany-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(2.9&nbsp;GB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="europe/germany-latest.osm.bz2">[.osm.bz2]</a></td></tr>
				</tbody></table>`,
			url: `https://download.geofabrik.de/europe.html`,
			elements: &element.MapElement{
				"europe": element.Element{ID: "europe", Name: "Europe", Meta: true},
			},
			want: element.MapElement{
				"europe":     element.Element{ID: "europe", Name: "Europe", Meta: true},
				"albania":    element.Element{ID: "albania", Name: "Albania", Meta: true, Parent: "europe"},
				"andorra":    element.Element{ID: "andorra", Name: "Andorra", Meta: true, Parent: "europe"},
				"france":     element.Element{ID: "france", Name: "France", Meta: true, Parent: "europe"},
				"georgia-eu": element.Element{ID: "georgia-eu", File: "georgia", Name: "Georgia (Eastern Europe)", Meta: true, Parent: "europe"},
				"germany":    element.Element{ID: "germany", Name: "Germany", Meta: true, Parent: "europe"},
			},
		},
		{
			name: "sample3 - North America extract",
			html: `
			<table id="subregions">
			<tbody><tr>
			   <th class="subregion">Sub Region</th>
			   <th colspan="4">Quick Links</th>
			</tr>
			<tr>
			   <th>&nbsp;</th>
			   <th colspan="2">.osm.pbf</th>
			   <th>.shp.zip</th>
			   <th>.osm.bz2</th>
			</tr>
			<tr onmouseover="loadkml('north-america/canada.kml')"><td class="subregion"><a href="north-america/canada.html">Canada</a></td>
			<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="north-america/canada-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(2.3&nbsp;GB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="north-america/canada-latest.osm.bz2">[.osm.bz2]</a></td></tr>
			<tr onmouseover="loadkml('north-america/us/florida.kml')"><td class="subregion"><a href="north-america/us/florida.html">Florida</a></td>
			<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="north-america/us/florida-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(211&nbsp;MB)</td><td> <a href="north-america/us/florida-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="north-america/us/florida-latest.osm.bz2">[.osm.bz2]</a></td></tr>
			<tr onmouseover="loadkml('north-america/us/georgia.kml')"><td class="subregion"><a href="north-america/us/georgia.html">Georgia (US State)</a></td>
			<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="north-america/us/georgia-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(170&nbsp;MB)</td><td> <a href="north-america/us/georgia-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="north-america/us/georgia-latest.osm.bz2">[.osm.bz2]</a></td></tr>
			<tr onmouseover="loadkml('north-america/us/hawaii.kml')"><td class="subregion"><a href="north-america/us/hawaii.html">Hawaii</a></td>
			<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="north-america/us/hawaii-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(11.6&nbsp;MB)</td><td> <a href="north-america/us/hawaii-latest-free.shp.zip">[.shp.zip]</a></td><td> <a href="north-america/us/hawaii-latest.osm.bz2">[.osm.bz2]</a></td></tr>
			</tbody></table>`,
			url: `https://download.geofabrik.de/europe.html`,
			elements: &element.MapElement{
				"north-america": element.Element{ID: "north-america", Name: "North America", Meta: true},
			},
			want: element.MapElement{
				"north-america": element.Element{ID: "north-america", Name: "North America", Meta: true},
				"us":            element.Element{ID: "us", File: "", Name: "us", Parent: "north-america", Formats: element.Formats{}, Meta: true},
				"canada":        element.Element{ID: "canada", Name: "Canada", Meta: true, Parent: "north-america"},
				"florida":       element.Element{ID: "florida", Name: "Florida", Meta: true, Parent: "us"},
				"georgia-us":    element.Element{ID: "georgia-us", File: "georgia", Name: "Georgia (US State)", Meta: true, Parent: "us"},
				"hawaii":        element.Element{ID: "hawaii", Name: "Hawaii", Meta: true, Parent: "us"},
			},
		},
		{
			name: "sample4 - South America extract",
			html: `
			<table id="subregions">
			<tbody><tr>
			   <th class="subregion">Sub Region</th>
			   <th colspan="4">Quick Links</th>
			</tr>
			<tr>
			   <th>&nbsp;</th>
			   <th colspan="2">.osm.pbf</th>
			   <th>.shp.zip</th>
			   <th>.osm.bz2</th>
			</tr>
			<tr onmouseover="loadkml('south-america/ecuador.kml')"><td class="subregion"><a href="south-america/ecuador.html">Ecuador</a></td>
			<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/ecuador-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(2.3&nbsp;GB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="south-america/ecuador-latest.osm.bz2">[.osm.bz2]</a></td></tr>
			<tr onmouseover="loadkml('south-america/guatemala.kml')"><td class="subregion"><a href="south-america/guatemala.html">Guatemala</a></td>
			<td style="border-right: 0px; margin-right: 0px; padding-right: 0px;"><a href="south-america/guatemala-latest.osm.pbf">[.osm.pbf]</a></td><td style="border-left: 0px; margin-left: 0px; padding-left: 0px;">(2.3&nbsp;GB)</td><td> <img alt="not available" src="/img/cross.png"></td><td> <a href="south-america/guatemala-latest.osm.bz2">[.osm.bz2]</a></td></tr>
			</tbody></table>`,
			url: `https://download.geofabrik.de/europe.html`,
			elements: &element.MapElement{
				"south-america": element.Element{ID: "south-america", Name: "North America", Meta: true},
			},
			want: element.MapElement{
				"south-america":           element.Element{ID: "south-america", Name: "North America", Meta: true},
				"ecuador":                 element.Element{ID: "ecuador", Name: "Ecuador", Meta: true, Parent: "south-america"},
				"guatemala-south-america": element.Element{ID: "guatemala-south-america", File: "guatemala", Name: "Guatemala", Meta: true, Parent: "south-america"},
			},
		},

		// TODO: Add test cases.
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(thisTest.html))
			myURL, _ := url.Parse(thisTest.url)
			myElement := &colly.HTMLElement{
				DOM: dom.Selection,
				Response: &colly.Response{
					Request: &colly.Request{URL: myURL},
				},
				Request: &colly.Request{URL: myURL},
			}

			defaultGeofabrik := geofabrik.GetDefault()

			myCollector := defaultGeofabrik.Collector() // Need a Collector to visit

			for _, elemem := range *thisTest.elements {
				if err := defaultGeofabrik.Config.MergeElement(&elemem); err != nil {
					t.Errorf("Bad tests g.Config.mergeElement() can't merge %v - %v", elemem, err)
				}
			}

			defaultGeofabrik.ParseSubregion(myElement, myCollector)

			for k := range defaultGeofabrik.Config.Elements {
				if !reflect.DeepEqual(defaultGeofabrik.Config.Elements[k], thisTest.want[k]) {
					t.Errorf("parseSubregion() fail, got \n%+v, want \n%+v", defaultGeofabrik.Config.Elements[k], thisTest.want[k])
				}
			}
		})
	}
}
