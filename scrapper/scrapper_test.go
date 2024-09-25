package scrapper_test

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper"
)

func TestGetParent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		url   string
		want  string
		want2 string
	}{
		// TODO: Add test cases.
		{
			name:  "No Parent",
			url:   "https://download.geofabrik.de/test.html",
			want:  "",
			want2: "https://download.geofabrik.de",
		},
		{
			name:  "1 parent",
			url:   "https://download.geofabrik.de/parent1/test.html",
			want:  "parent1",
			want2: "https://download.geofabrik.de/parent1",
		},
		{
			name:  "2 parents",
			url:   "https://download.geofabrik.de/parent1/parent2/test.html",
			want:  "parent2",
			want2: "https://download.geofabrik.de/parent1/parent2",
		},
		{
			name:  "grand parents",
			url:   "https://download.geofabrik.de/parent1/parent2",
			want:  "parent1",
			want2: "https://download.geofabrik.de/parent1",
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, got2 := scrapper.GetParent(thisTest.url)
			if got != thisTest.want {
				t.Errorf("GetParent() = %v, want %v", got, thisTest.want)
			}

			if got2 != thisTest.want2 {
				t.Errorf("GetParent() = %v, want %v", got2, thisTest.want2)
			}
		})
	}
}

func Test_FileExt(t *testing.T) {
	t.Parallel()

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
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, ext := scrapper.FileExt(thisTest.url)
			if got != thisTest.want {
				t.Errorf("FileExt() = %v, want %v", got, thisTest.want)
			}

			if ext != thisTest.want2 {
				t.Errorf("FileExt() = %v, want %v", ext, thisTest.want2)
			}
		})
	}
}

func Test_ParseFormat(t *testing.T) {
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
					Formats: []string{formats.FormatOsmPbf},
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
					Formats: []string{formats.FormatOsmPbf},
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
					Formats: []string{formats.FormatOsmPbf},
					Meta:    false,
				},
			},
		},
	}
	for thisTest := range tests {
		t.Run(tests[thisTest].name, func(t *testing.T) {
			t.Parallel()

			myScrapper := scrapper.Scrapper{
				Config: &tests[thisTest].args.c,
			}
			myScrapper.Config.Formats = formats.FormatDefinitions{
				formats.FormatOshPbf: {ID: formats.FormatOshPbf, Loc: ".osh.pbf"},
				"osh.pbf.md5":        formats.Format{ID: "osh.pbf.md5", Loc: ".osh.pbf.md5"},
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

func TestScrapper_PB(t *testing.T) {
	t.Parallel()

	for range [10]int{} {
		want, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			t.Fatalf("Failed to generate random number: %v", err)
		}

		wantInt := int(want.Int64())
		myScrapper := scrapper.Scrapper{
			PB: wantInt,
		}

		t.Run(fmt.Sprintf("PB: %d", want), func(t *testing.T) {
			t.Parallel()

			out := myScrapper.GetPB()
			if out != wantInt {
				t.Errorf("GetPB() got %d, want %d", out, want)
			}
		})
	}
}

func generateRandomString(n int, charset string) (string, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}

	return string(b), nil
}

func TestScrapper_GetStartURL(t *testing.T) {
	t.Parallel()

	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "\\/?."

	for i := range [1024]int{} {
		want, err := generateRandomString(i, charset)
		if err != nil {
			t.Fatalf("Failed to generate random string: %v", err)
		}
		myScrapper := scrapper.Scrapper{
			StartURL: want,
		}

		t.Run(fmt.Sprintf("StartURL: %s", want), func(t *testing.T) {
			t.Parallel()

			out := myScrapper.GetStartURL()
			if out != want {
				t.Errorf("GetStartURL() got %v, want %v", out, want)
			}
		})
	}
}

func TestScrapper_Limit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		want   *colly.LimitRule
		fields scrapper.Scrapper
	}{
		// TODO: Add test cases.
		{
			name:   "Default",
			fields: scrapper.Scrapper{DomainGlob: "*", Parallelism: 1, RandomDelay: 5 * time.Second},
			want:   &colly.LimitRule{DomainGlob: "*", Parallelism: 1, RandomDelay: 5 * time.Second},
		},
		{
			name:   "Void",
			fields: scrapper.Scrapper{},
			want:   &colly.LimitRule{DomainGlob: "*", Parallelism: 1, RandomDelay: 5 * time.Second},
		},
	}
	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			myScrapper := &scrapper.Scrapper{
				BaseURL:          thisTest.fields.BaseURL,
				StartURL:         thisTest.fields.StartURL,
				Config:           thisTest.fields.Config,
				PB:               thisTest.fields.PB,
				Async:            thisTest.fields.Async,
				DomainGlob:       thisTest.fields.DomainGlob,
				Parallelism:      thisTest.fields.Parallelism,
				RandomDelay:      thisTest.fields.RandomDelay,
				MaxDepth:         thisTest.fields.MaxDepth,
				AllowedDomains:   thisTest.fields.AllowedDomains,
				URLFilters:       thisTest.fields.URLFilters,
				FormatDefinition: thisTest.fields.FormatDefinition,
			}
			if got := myScrapper.Limit(); !reflect.DeepEqual(got, thisTest.want) {
				t.Errorf("Scrapper.Limit() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

func TestScrapper_GetConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		want   *config.Config
		fields scrapper.Scrapper
	}{
		// TODO: Add test cases.
		{
			name: "Void",
			want: &config.Config{Elements: element.MapElement{}, ElementsMutex: &sync.RWMutex{}},
		},
		{
			name:   "Void + BaseURL",
			fields: scrapper.Scrapper{BaseURL: "http://my.url"},
			want:   &config.Config{Elements: element.MapElement{}, ElementsMutex: &sync.RWMutex{}, BaseURL: "http://my.url"},
		},
		{
			name:   "Void + FormatDefinition",
			fields: scrapper.Scrapper{FormatDefinition: formats.FormatDefinitions{"ext": formats.Format{ID: "ext"}}},
			want: &config.Config{
				Elements:      element.MapElement{},
				ElementsMutex: &sync.RWMutex{},
				Formats:       formats.FormatDefinitions{"ext": formats.Format{ID: "ext"}},
			},
		},
		{
			name: "Void + BaseURL",
			fields: scrapper.Scrapper{
				BaseURL:          "http://my.url",
				FormatDefinition: formats.FormatDefinitions{"ext": formats.Format{ID: "ext"}},
			},
			want: &config.Config{
				Elements:      element.MapElement{},
				ElementsMutex: &sync.RWMutex{},
				BaseURL:       "http://my.url",
				Formats:       formats.FormatDefinitions{"ext": formats.Format{ID: "ext"}},
			},
		},
		{
			name: "Config Exist",
			fields: scrapper.Scrapper{
				Config: &config.Config{
					Elements: element.MapElement{
						"a": element.Element{ID: "a"},
					},
					ElementsMutex: &sync.RWMutex{},
					BaseURL:       "http://my.url",
					Formats:       formats.FormatDefinitions{"ext": formats.Format{ID: "ext"}},
				},
			},
			want: &config.Config{
				Elements: element.MapElement{
					"a": element.Element{ID: "a"},
				},
				ElementsMutex: &sync.RWMutex{},
				BaseURL:       "http://my.url",
				Formats:       formats.FormatDefinitions{"ext": formats.Format{ID: "ext"}},
			},
		},
		{
			name: "Config Exist+ Base URL",
			fields: scrapper.Scrapper{
				BaseURL: "http://my.url",
				Config: &config.Config{
					Elements: element.MapElement{
						"a": element.Element{ID: "a"},
					},
					ElementsMutex: &sync.RWMutex{},
					BaseURL:       "http://old.url",
					Formats:       formats.FormatDefinitions{"ext": formats.Format{ID: "ext"}},
				},
			},
			want: &config.Config{
				Elements: element.MapElement{
					"a": element.Element{ID: "a"},
				},
				ElementsMutex: &sync.RWMutex{},
				BaseURL:       "http://my.url",
				Formats:       formats.FormatDefinitions{"ext": formats.Format{ID: "ext"}},
			},
		},
		{
			name: "Config Exist+ Base URL + Format Definition",
			fields: scrapper.Scrapper{
				BaseURL: "http://my.url",
				Config: &config.Config{
					Elements: element.MapElement{
						"a": element.Element{ID: "a"},
					},
					ElementsMutex: &sync.RWMutex{},
					BaseURL:       "http://old.url",
					Formats:       formats.FormatDefinitions{"old": formats.Format{ID: "old"}},
				},
				FormatDefinition: formats.FormatDefinitions{"ext": formats.Format{ID: "ext"}},
			},
			want: &config.Config{
				Elements: element.MapElement{
					"a": element.Element{ID: "a"},
				},
				ElementsMutex: &sync.RWMutex{},
				BaseURL:       "http://my.url",
				Formats:       formats.FormatDefinitions{"ext": formats.Format{ID: "ext"}},
			},
		},
	}
	for _, thisTest := range tests {
		thisTest := thisTest

		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			s := &scrapper.Scrapper{
				BaseURL:          thisTest.fields.BaseURL,
				Config:           thisTest.fields.Config,
				FormatDefinition: thisTest.fields.FormatDefinition,
			}
			if got := s.GetConfig(); !reflect.DeepEqual(got, thisTest.want) {
				t.Errorf("Scrapper.GetConfig() = %v, want %v", got, thisTest.want)
			}
		})
	}
}
