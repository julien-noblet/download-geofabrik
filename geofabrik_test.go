package main

import (
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func Test_geofabrikParseFormat(t *testing.T) {
	type args struct {
		id     string
		format string
		c      Config
	}

	tests := []struct {
		name string
		args args
		want ElementSlice
	}{
		{
			name: "Add osm.pbf but already in",
			args: args{
				id:     "a",
				format: formatOsmPbf,
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formatOsmPbf},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf, formatKml, formatState},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf",
			args: args{
				id:     "a",
				format: formatOsmPbf,
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf, formatKml, formatState},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf.md5",
			args: args{
				id:     "a",
				format: "osm.pbf.md5",
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formatOsmPbf, formatKml, formatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf, formatKml, formatState, "osm.pbf.md5"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.bz2",
			args: args{
				id:     "a",
				format: formatOsmBz2,
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formatOsmPbf, formatKml, formatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf, formatKml, formatState, formatOsmBz2},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.bz2.md5",
			args: args{
				id:     "a",
				format: "osm.bz2.md5",
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formatOsmPbf, formatKml, formatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf, formatKml, formatState, "osm.bz2.md5"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add poly",
			args: args{
				id:     "a",
				format: formatPoly,
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formatOsmPbf, formatKml, formatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf, formatKml, formatState, formatPoly},
					Meta:    false,
				},
			},
		},
		{
			name: "Add shp.zip",
			args: args{
				id:     "a",
				format: formatShpZip,
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formatOsmPbf, formatKml, formatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf, formatKml, formatState, formatShpZip},
					Meta:    false,
				},
			},
		},
		{
			name: "Add unk format",
			args: args{
				id:     "a",
				format: "unk",
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{formatOsmPbf, formatKml, formatState},
							Meta:    false,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf, formatKml, formatState},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf on meta",
			args: args{
				id:     "a",
				format: formatOsmPbf,
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{},
							Meta:    true,
						},
					},
					ElementsMutex: &sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf, formatKml, formatState},
					Meta:    false,
				},
			},
		},
	}

	for tn := range tests {
		tn := tn
		t.Run(tests[tn].name, func(t *testing.T) {
			s := Geofabrik{Scrapper: &Scrapper{}}
			s.Scrapper.Config = &tests[tn].args.c
			s.Scrapper.Config.Formats = formatDefinitions{
				formatOshPbf:  {ID: formatOshPbf, Loc: ".osh.pbf"},
				"osh.pbf.md5": format{ID: "osh.pbf.md5", Loc: ".osh.pbf.md5"},
				formatOsmBz2:  {ID: formatOsmBz2, Loc: "-latest.osm.bz2"},
				"osm.bz2.md5": {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
				formatOsmPbf:  {ID: formatOsmPbf, Loc: "-latest.osm.pbf"},
				"osm.pbf.md5": {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
				formatPoly:    {ID: formatPoly, Loc: ".poly"},
				formatKml:     {ID: formatKml, Loc: ".kml"},
				formatState:   {ID: formatState, Loc: "-updates/state.txt"},
				formatShpZip:  {ID: formatShpZip, Loc: "-latest-free.shp.zip"},
			}

			s.ParseFormat(tests[tn].args.id, tests[tn].args.format)
			if !reflect.DeepEqual(tests[tn].args.c.Elements, tests[tn].want) {
				t.Errorf("ParseFormat() got %v, want %v", tests[tn].args.c.Elements, tests[tn].want)
			}
		})
	}
}

func TestGeofabrik_parseLi(t *testing.T) {
	tests := []struct {
		name    string
		html    string
		url     string
		want    ElementSlice
		element *Element
	}{
		{name: "sample",
			html: `
			<li>
				<p>some text</p>
				<a href="toto.osm.pbf">anotherText</a>
			</li>`,
			url:     `https://download.geofabrik.de/toto.html`,
			element: &Element{ID: "toto"},
			want:    ElementSlice{"toto": Element{ID: "toto", Formats: elementFormats{formatOsmPbf, formatKml, formatState}}},
		},
		{name: "georgia-us",
			html: `
			<li>
				<p>some text</p>
				<a href="georgia.osm.pbf">anotherText</a>
			</li>`,
			url:     `https://download.geofabrik.de/north-america/us/georgia.html`,
			element: &Element{ID: "georgia-us"},
			want:    ElementSlice{"georgia-us": Element{ID: "georgia-us", Formats: elementFormats{formatOsmPbf, formatKml, formatState}}},
		},
		{name: "georgia-eu",
			html: `
			<li>
				<p>some text</p>
				<a href="georgia.osm.pbf">anotherText</a>
			</li>`,
			url:     `https://download.geofabrik.de/europe/georgia.html`,
			element: &Element{ID: "georgia-eu"},
			want:    ElementSlice{"georgia-eu": Element{ID: "georgia-eu", Formats: elementFormats{formatOsmPbf, formatKml, formatState}}},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(tt.html))
			u, _ := url.Parse(tt.url)
			e := &colly.HTMLElement{
				DOM: dom.Selection,
				Response: &colly.Response{
					Request: &colly.Request{URL: u},
				},
			}
			g := Geofabrik{
				Scrapper: &Scrapper{
					PB:             401,
					Async:          true,
					Parallelism:    20,
					MaxDepth:       0,
					AllowedDomains: []string{`download.geofabrik.de`},
					BaseURL:        `https://download.geofabrik.de`,
					StartURL:       `https://download.geofabrik.de/`,
					URLFilters: []*regexp.Regexp{
						regexp.MustCompile(`https://download\.geofabrik\.de/.+\.html$`),
						regexp.MustCompile(`https://download\.geofabrik\.de/$`),
					},
					FormatDefinition: formatDefinitions{
						formatOsmBz2:  {ID: formatOsmBz2, Loc: "-latest.osm.bz2"},
						"osm.bz2.md5": {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
						formatOsmPbf:  {ID: formatOsmPbf, Loc: "-latest.osm.pbf"},
						"osm.pbf.md5": {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
						formatPoly:    {ID: formatPoly, Loc: ".poly"},
						formatKml:     {ID: formatKml, Loc: ".kml"},
						formatState:   {ID: formatState, Loc: "-updates/state.txt"},
						formatShpZip:  {ID: formatShpZip, Loc: "-latest-free.shp.zip"},
					},
				},
			}
			g.GetConfig()
			if err := g.Config.mergeElement(tt.element); err != nil {
				t.Errorf("Bad tests g.Config.mergeElement() can't merge %v - %v", tt.element, err)
			}
			g.parseLi(e, nil)
			if !reflect.DeepEqual(g.Config.Elements, tt.want) {
				t.Errorf("parseLi() fail, got %v, want %v", g.Config.Elements, tt.want)
			}
		})
	}
}

func TestGeofabrik_parseSubregion(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		url      string
		want     ElementSlice
		elements *ElementSlice
	}{
		{name: "sample1",
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
			elements: &ElementSlice{},
			want: ElementSlice{
				"africa":            Element{ID: "africa", Name: "Africa", Meta: true},
				"antarctica":        Element{ID: "antarctica", Name: "Antarctica", Meta: true},
				"asia":              Element{ID: "asia", Name: "Asia", Meta: true},
				"australia-oceania": Element{ID: "australia-oceania", Name: "Australia and Oceania", Meta: true},
				"central-america":   Element{ID: "central-america", Name: "Central America", Meta: true},
				"europe":            Element{ID: "europe", Name: "Europe", Meta: true},
				"north-america":     Element{ID: "north-america", Name: "North America", Meta: true},
				"south-america":     Element{ID: "south-america", Name: "South America", Meta: true},
			},
		},
		{name: "sample2 - Europe extract",
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
			elements: &ElementSlice{
				"europe": Element{ID: "europe", Name: "Europe", Meta: true},
			},
			want: ElementSlice{
				"europe":     Element{ID: "europe", Name: "Europe", Meta: true},
				"albania":    Element{ID: "albania", Name: "Albania", Meta: true, Parent: "europe"},
				"andorra":    Element{ID: "andorra", Name: "Andorra", Meta: true, Parent: "europe"},
				"france":     Element{ID: "france", Name: "France", Meta: true, Parent: "europe"},
				"georgia-eu": Element{ID: "georgia-eu", Name: "Georgia (Eastern Europe)", Meta: true, Parent: "europe"},
				"germany":    Element{ID: "germany", Name: "Germany", Meta: true, Parent: "europe"},
			},
		},
		{name: "sample3 - North America extract",
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
			elements: &ElementSlice{
				"north-america": Element{ID: "north-america", Name: "North America", Meta: true},
			},
			want: ElementSlice{
				"north-america": Element{ID: "north-america", Name: "North America", Meta: true},
				"us":            Element{ID: "us", Name: "us", Meta: true, Parent: "north-america"},
				"canada":        Element{ID: "canada", Name: "Canada", Meta: true, Parent: "north-america"},
				"florida":       Element{ID: "florida", Name: "Florida", Meta: true, Parent: "us"},
				"georgia-us":    Element{ID: "georgia-us", Name: "Georgia (US State)", Meta: true, Parent: "us"},
				"hawaii":        Element{ID: "hawaii", Name: "Hawaii", Meta: true, Parent: "us"},
			},
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(tt.html))
			u, _ := url.Parse(tt.url)
			e := &colly.HTMLElement{
				DOM: dom.Selection,
				Response: &colly.Response{
					Request: &colly.Request{URL: u},
				},
				Request: &colly.Request{URL: u},
			}

			g := Geofabrik{
				Scrapper: &Scrapper{
					PB:             401,
					Async:          true,
					Parallelism:    20,
					MaxDepth:       1, // Force a limit!
					AllowedDomains: []string{`download.geofabrik.de`},
					BaseURL:        `https://download.geofabrik.de`,
					StartURL:       `https://download.geofabrik.de/`,
					URLFilters: []*regexp.Regexp{
						regexp.MustCompile(`https://download\.geofabrik\.de/.+\.html$`),
						regexp.MustCompile(`https://download\.geofabrik\.de/$`),
					},
					FormatDefinition: formatDefinitions{
						formatOsmBz2:  {ID: formatOsmBz2, Loc: "-latest.osm.bz2"},
						"osm.bz2.md5": {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
						formatOsmPbf:  {ID: formatOsmPbf, Loc: "-latest.osm.pbf"},
						"osm.pbf.md5": {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
						formatPoly:    {ID: formatPoly, Loc: ".poly"},
						formatKml:     {ID: formatKml, Loc: ".kml"},
						formatState:   {ID: formatState, Loc: "-updates/state.txt"},
						formatShpZip:  {ID: formatShpZip, Loc: "-latest-free.shp.zip"},
					},
				},
			}
			c := g.Collector() // Need a Collector to visit
			for _, elemem := range *tt.elements {
				elemem := elemem
				if err := g.Config.mergeElement(&elemem); err != nil {
					t.Errorf("Bad tests g.Config.mergeElement() can't merge %v - %v", elemem, err)
				}
			}
			g.parseSubregion(e, c)
			if !reflect.DeepEqual(g.Config.Elements, tt.want) {
				t.Errorf("parseSubregion() fail, got \n%v, want \n%v", g.Config.Elements, tt.want)
			}
		})
	}
}
