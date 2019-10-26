// 2015-2019 copyright Julien Noblet

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/gocolly/colly"
)

func TestGetParent(t *testing.T) {
	tests := []struct {
		name  string
		url   string
		want  string
		want2 string
	}{
		// TODO: Add test cases.
		{name: "No Parent", url: "https://download.geofabrik.de/test.html", want: "", want2: "https://download.geofabrik.de"},
		{name: "1 parent", url: "https://download.geofabrik.de/parent1/test.html", want: "parent1", want2: "https://download.geofabrik.de/parent1"},
		{name: "2 parents", url: "https://download.geofabrik.de/parent1/parent2/test.html", want: "parent2", want2: "https://download.geofabrik.de/parent1/parent2"},
		{name: "grand parents", url: "https://download.geofabrik.de/parent1/parent2", want: "parent1", want2: "https://download.geofabrik.de/parent1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := GetParent(tt.url)
			if got != tt.want {
				t.Errorf("GetParent() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("GetParent() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_FileExt(t *testing.T) {
	tests := []struct {
		name  string
		url   string
		want  string
		want2 string
	}{
		// TODO: Add test cases.
		{name: "No Parent", url: "https://download.geofabrik.de/test.html", want: "test", want2: "html"},
		{name: "No Parent long ext", url: "https://download.geofabrik.de/test.ext.html", want: "test", want2: "ext.html"},
		{name: "1 Parent", url: "https://download.geofabrik.de/parent/test.html", want: "test", want2: "html"},
		{name: "1 Parent long ext", url: "https://download.geofabrik.de/parent/test.ext.html", want: "test", want2: "ext.html"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ext := FileExt(tt.url)
			if got != tt.want {
				t.Errorf("FileExt() = %v, want %v", got, tt.want)
			}
			if ext != tt.want2 {
				t.Errorf("FileExt() = %v, want %v", ext, tt.want2)
			}
		})
	}
}

func Test_ParseFormat(t *testing.T) {
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
		{name: "Add osm.pbf but already in",
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
					Formats: []string{"osm.pbf"},
					Meta:    false,
				},
			},
		},
		{name: "Add osm.pbf",
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
					Formats: []string{"osm.pbf"},
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
		{name: "Add shp.zip",
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
		{name: "Add unk format",
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
		{name: "Add osm.pbf on meta",
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
					Formats: []string{"osm.pbf"},
					Meta:    false,
				},
			},
		},
	}
	for tn := range tests {
		t.Run(tests[tn].name, func(t *testing.T) {
			s := Scrapper{
				Config: &tests[tn].args.c,
			}
			s.Config.Formats = formatDefinitions{
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

func TestScrapper_PB(t *testing.T) {
	for i := 0; i < 10; i++ {
		want := rand.Int()
		s := Scrapper{
			PB: want,
		}
		out := s.GetPB()
		if out != want {
			t.Errorf("GetPB() got %d, want %d", out, want)
		}
	}
}

func TestScrapper_GetStartURL(t *testing.T) {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "\\/?."

	for i := 0; i < 1024; i++ {
		var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
		want := stringWithCharset(seededRand, i, charset)
		s := Scrapper{
			StartURL: want,
		}
		out := s.GetStartURL()
		if out != want {
			t.Errorf("GetStartURL() got %v, want %v", out, want)
		}
	}
}

func stringWithCharset(seededRand *rand.Rand, length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func TestScrapper_Limit(t *testing.T) {

	tests := []struct {
		name   string
		fields Scrapper
		want   *colly.LimitRule
	}{
		// TODO: Add test cases.
		{name: "Default",
			fields: Scrapper{DomainGlob: "*", Parallelism: 1, RandomDelay: 5 * time.Second},
			want:   &colly.LimitRule{DomainGlob: "*", Parallelism: 1, RandomDelay: 5 * time.Second},
		},
		{name: "Void",
			fields: Scrapper{},
			want:   &colly.LimitRule{DomainGlob: "*", Parallelism: 1, RandomDelay: 5 * time.Second},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scrapper{
				BaseURL:          tt.fields.BaseURL,
				StartURL:         tt.fields.StartURL,
				Config:           tt.fields.Config,
				PB:               tt.fields.PB,
				Async:            tt.fields.Async,
				DomainGlob:       tt.fields.DomainGlob,
				Parallelism:      tt.fields.Parallelism,
				RandomDelay:      tt.fields.RandomDelay,
				MaxDepth:         tt.fields.MaxDepth,
				AllowedDomains:   tt.fields.AllowedDomains,
				URLFilters:       tt.fields.URLFilters,
				FormatDefinition: tt.fields.FormatDefinition,
			}
			if got := s.Limit(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scrapper.Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScrapper_GetConfig(t *testing.T) {
	tests := []struct {
		name   string
		fields Scrapper
		want   *Config
	}{
		// TODO: Add test cases.
		{name: "Void",
			want: &Config{Elements: ElementSlice{}, ElementsMutex: &sync.RWMutex{}},
		},
		{name: "Void + BaseURL",
			fields: Scrapper{BaseURL: "http://my.url"},
			want:   &Config{Elements: ElementSlice{}, ElementsMutex: &sync.RWMutex{}, BaseURL: "http://my.url"},
		},
		{name: "Void + FormatDefinition",
			fields: Scrapper{FormatDefinition: formatDefinitions{"ext": format{ID: "ext"}}},
			want:   &Config{Elements: ElementSlice{}, ElementsMutex: &sync.RWMutex{}, Formats: formatDefinitions{"ext": format{ID: "ext"}}},
		},
		{name: "Void + BaseURL",
			fields: Scrapper{
				BaseURL:          "http://my.url",
				FormatDefinition: formatDefinitions{"ext": format{ID: "ext"}},
			},
			want: &Config{
				Elements:      ElementSlice{},
				ElementsMutex: &sync.RWMutex{},
				BaseURL:       "http://my.url",
				Formats:       formatDefinitions{"ext": format{ID: "ext"}},
			},
		},
		{name: "Config Exist",
			fields: Scrapper{
				Config: &Config{
					Elements: ElementSlice{
						"a": Element{ID: "a"},
					},
					ElementsMutex: &sync.RWMutex{},
					BaseURL:       "http://my.url",
					Formats:       formatDefinitions{"ext": format{ID: "ext"}},
				},
			},
			want: &Config{
				Elements: ElementSlice{
					"a": Element{ID: "a"},
				},
				ElementsMutex: &sync.RWMutex{},
				BaseURL:       "http://my.url",
				Formats:       formatDefinitions{"ext": format{ID: "ext"}},
			},
		},
		{name: "Config Exist+ Base URL",
			fields: Scrapper{
				BaseURL: "http://my.url",
				Config: &Config{
					Elements: ElementSlice{
						"a": Element{ID: "a"},
					},
					ElementsMutex: &sync.RWMutex{},
					BaseURL:       "http://old.url",
					Formats:       formatDefinitions{"ext": format{ID: "ext"}},
				},
			},
			want: &Config{
				Elements: ElementSlice{
					"a": Element{ID: "a"},
				},
				ElementsMutex: &sync.RWMutex{},
				BaseURL:       "http://my.url",
				Formats:       formatDefinitions{"ext": format{ID: "ext"}},
			},
		},
		{name: "Config Exist+ Base URL + Format Definition",
			fields: Scrapper{
				BaseURL: "http://my.url",
				Config: &Config{
					Elements: ElementSlice{
						"a": Element{ID: "a"},
					},
					ElementsMutex: &sync.RWMutex{},
					BaseURL:       "http://old.url",
					Formats:       formatDefinitions{"old": format{ID: "old"}},
				},
				FormatDefinition: formatDefinitions{"ext": format{ID: "ext"}},
			},
			want: &Config{
				Elements: ElementSlice{
					"a": Element{ID: "a"},
				},
				ElementsMutex: &sync.RWMutex{},
				BaseURL:       "http://my.url",
				Formats:       formatDefinitions{"ext": format{ID: "ext"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scrapper{
				BaseURL:          tt.fields.BaseURL,
				Config:           tt.fields.Config,
				FormatDefinition: tt.fields.FormatDefinition,
			}
			if got := s.GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scrapper.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
