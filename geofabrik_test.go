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
				format: "osm.pbf",
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf"},
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
					Formats: []string{"osm.pbf", "kml", "state"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf",
			args: args{
				id:     "a",
				format: "osm.pbf",
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
					Formats: []string{"osm.pbf", "kml", "state"},
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
							Formats: []string{"osm.pbf", "kml", "state"},
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
					Formats: []string{"osm.pbf", "kml", "state", "osm.pbf.md5"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.bz2",
			args: args{
				id:     "a",
				format: "osm.bz2",
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf", "kml", "state"},
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
					Formats: []string{"osm.pbf", "kml", "state", "osm.bz2"},
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
							Formats: []string{"osm.pbf", "kml", "state"},
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
					Formats: []string{"osm.pbf", "kml", "state", "osm.bz2.md5"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add poly",
			args: args{
				id:     "a",
				format: "poly",
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf", "kml", "state"},
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
					Formats: []string{"osm.pbf", "kml", "state", "poly"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add shp.zip",
			args: args{
				id:     "a",
				format: "shp.zip",
				c: Config{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf", "kml", "state"},
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
					Formats: []string{"osm.pbf", "kml", "state", "shp.zip"},
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
							Formats: []string{"osm.pbf", "kml", "state"},
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
					Formats: []string{"osm.pbf", "kml", "state"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf on meta",
			args: args{
				id:     "a",
				format: "osm.pbf",
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
					Formats: []string{"osm.pbf", "kml", "state"},
					Meta:    false,
				},
			},
		},
	}
	for tn := range tests {
		t.Run(tests[tn].name, func(t *testing.T) {
			s := Geofabrik{Scrapper: &Scrapper{}}
			s.Scrapper.Config = &tests[tn].args.c
			s.Scrapper.Config.Formats = formatDefinitions{
				"osh.pbf":     {ID: "osh.pbf", Loc: ".osh.pbf"},
				"osh.pbf.md5": format{ID: "osh.pbf.md5", Loc: ".osh.pbf.md5"},
				"osm.bz2":     {ID: "osm.bz2", Loc: "-latest.osm.bz2"},
				"osm.bz2.md5": {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
				"osm.pbf":     {ID: "osm.pbf", Loc: "-latest.osm.pbf"},
				"osm.pbf.md5": {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
				"poly":        {ID: "poly", Loc: ".poly"},
				"kml":         {ID: "kml", Loc: ".kml"},
				"state":       {ID: "state", Loc: "-updates/state.txt"},
				"shp.zip":     {ID: "shp.zip", Loc: "-latest-free.shp.zip"},
			}

			s.ParseFormat(tests[tn].args.id, tests[tn].args.format)
			if !reflect.DeepEqual(tests[tn].args.c.Elements, tests[tn].want) {
				t.Errorf("ParseFormat() got %v, want %v", tests[tn].args.c.Elements, tests[tn].want)
			}
		})
	}
}

func TestGeofabrik_parseLi(t *testing.T) {
	type fields struct {
		Scrapper *Scrapper
	}

	tests := []struct {
		name    string
		fields  fields
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
			url:     `http://download.geofabrik.de/toto.html`,
			element: &Element{ID: "toto"},
			want:    ElementSlice{"toto": Element{ID: "toto", Formats: elementFormats{"osm.pbf", "kml", "state"}}},
		},
		{name: "georgia-us",
			html: `
			<li>
				<p>some text</p>
				<a href="georgia.osm.pbf">anotherText</a>
			</li>`,
			url:     `http://download.geofabrik.de/north-america/us/georgia.html`,
			element: &Element{ID: "georgia-us"},
			want:    ElementSlice{"georgia-us": Element{ID: "georgia-us", Formats: elementFormats{"osm.pbf", "kml", "state"}}},
		},
		{name: "georgia-eu",
			html: `
			<li>
				<p>some text</p>
				<a href="georgia.osm.pbf">anotherText</a>
			</li>`,
			url:     `http://download.geofabrik.de/europe/georgia.html`,
			element: &Element{ID: "georgia-eu"},
			want:    ElementSlice{"georgia-eu": Element{ID: "georgia-eu", Formats: elementFormats{"osm.pbf", "kml", "state"}}},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(tt.html))
			url, _ := url.Parse(tt.url)
			e := &colly.HTMLElement{
				DOM: dom.Selection,
				Response: &colly.Response{
					Request: &colly.Request{URL: url},
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
						//geofabrik.Formats["osh.pbf"] = format{ID: "osh.pbf", Loc: ".osh.pbf"}
						//geofabrik.Formats["osh.pbf.md5"] = format{ID: "osh.pbf.md5", Loc: ".osh.pbf.md5"}
						"osm.bz2":     {ID: "osm.bz2", Loc: "-latest.osm.bz2"},
						"osm.bz2.md5": {ID: "osm.bz2.md5", Loc: "-latest.osm.bz2.md5"},
						"osm.pbf":     {ID: "osm.pbf", Loc: "-latest.osm.pbf"},
						"osm.pbf.md5": {ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"},
						"poly":        {ID: "poly", Loc: ".poly"},
						"kml":         {ID: "kml", Loc: ".kml"},
						"state":       {ID: "state", Loc: "-updates/state.txt"},
						"shp.zip":     {ID: "shp.zip", Loc: "-latest-free.shp.zip"},
					},
				},
			}
			g.GetConfig()
			g.Config.mergeElement(tt.element)
			g.parseLi(e, nil)
			if !reflect.DeepEqual(g.Config.Elements, tt.want) {
				t.Errorf("parseLi() fail, got %v, want %v", g.Config.Elements, tt.want)
			}

		})
	}
}
