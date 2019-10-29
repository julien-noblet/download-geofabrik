package config_test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
)

const (
	geofabrikYml       = "../geofabrik.yml"
	openstreetmapFRYml = "../openstreetmap.fr.yml"
)

var sampleAfricaElementPtr = element.Element{
	ID:   "africa",
	Name: "Africa",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
}
var sampleGeorgiaUsElementPtr = element.Element{
	ID:   "georgia-us",
	File: "georgia",
	Name: "Georgia (US State)",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatShpZip,
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
	Parent: "us",
}
var sampleUsElementPtr = element.Element{
	ID:     "us",
	Meta:   true,
	Name:   "United States of America",
	Parent: "north-america",
}
var sampleNorthAmericaElementPtr = element.Element{
	ID:   "north-america",
	Name: "North America",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
}
var sampleElementValidPtr = map[string]element.Element{
	"africa":        sampleAfricaElementPtr,
	"georgia-us":    sampleGeorgiaUsElementPtr,
	"us":            sampleUsElementPtr,
	"north-america": sampleNorthAmericaElementPtr,
}
var sampleFormatValidPtr = map[string]formats.Format{
	// Blank
	"": {
		ID:       "",
		Loc:      "",
		BasePath: "",
	}, formats.FormatOsmPbf: {
		ID:  formats.FormatOsmPbf,
		Loc: ".osm.pbf",
		//BasePath: "/",
	}, formats.FormatState: {
		ID:       formats.FormatState,
		Loc:      "-updates/state.txt",
		BasePath: "../state/",
	}, formats.FormatPoly: {
		ID:      formats.FormatPoly,
		Loc:     ".poly",
		BaseURL: "http://my.new.url/folder",
	}, formats.FormatOsmBz2: {
		ID:       formats.FormatOsmBz2,
		Loc:      ".osm.bz2",
		BasePath: "../osmbz2/",
		BaseURL:  "http://my.new.url/folder",
	}, formats.FormatOsmGz: {
		ID:       formats.FormatOsmGz,
		Loc:      ".osm.gz",
		BasePath: "../osmgz/",
		BaseURL:  "http://my.new.url/folder",
	},
}

var SampleConfigValidPtr = config.Config{
	BaseURL:  "https://my.base.url",
	Formats:  sampleFormatValidPtr,
	Elements: sampleElementValidPtr,
}
var sampleFakeGeorgiaPtr = element.Element{
	ID:   "georgia-usf",
	File: "georgia-fake",
	Name: "Georgia (US State) - fake test",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatShpZip,
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
	Parent: "us", // keep good parent!
}
var sampleFakeGeorgia2Ptr = element.Element{
	ID:   "georgia-us2",
	File: "georgia",
	Name: "Georgia (US State)",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatShpZip,
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
	Parent: "notus", // bad parent not exist!
}
var sampleFakeGeorgia3Ptr = element.Element{
	ID:   "georgia-us3",
	File: "georgia",
	Name: "Georgia (US State)",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatShpZip,
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
	Parent: "notus2",
}
var sampleNotUS2Ptr = element.Element{
	ID:     "notus2",
	Name:   "notus2",
	Meta:   true,
	Parent: "north", // bad parent not exist!
}

func Benchmark_Exist_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := config.LoadConfig(geofabrikYml)
	e := "france"

	for n := 0; n < b.N; n++ {
		c.Exist(e)
	}
}

func Benchmark_loadConfig_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		if _, err := config.LoadConfig(geofabrikYml); err != nil {
			b.Error(err.Error())
		}
	}
}

func Benchmark_loadConfig_osmfr_yml(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		if _, err := config.LoadConfig(openstreetmapFRYml); err != nil {
			b.Error(err.Error())
		}
	}
}

func Test_loadConfig(t *testing.T) {
	type args struct {
		configFile string
	}

	tests := []struct {
		name    string
		args    args
		want    *config.Config
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Check config file not exist",
			args:    args{configFile: "./this_file_not_exists"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Check config not yaml",
			args:    args{configFile: "./LICENSE"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check config valid geofabrik.yml",
			args: args{configFile: geofabrikYml},
			want: &config.Config{
				BaseURL:       "https://download.geofabrik.de",
				Formats:       sampleFormatValidPtr,
				Elements:      sampleElementValidPtr,
				ElementsMutex: &sync.RWMutex{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.LoadConfig(tt.args.configFile)
			if err != nil != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
					t.Errorf("loadConfig() is a %v, want %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
				}
				// Check BaseURL
				if got.BaseURL != tt.want.BaseURL {
					t.Errorf("loadConfig().BaseURL = %v, want %v", got.BaseURL, tt.want.BaseURL)
				}
				// Must have formats but ne need to check if they are valids!
				if reflect.TypeOf(got.Formats) != reflect.TypeOf(tt.want.Formats) {
					t.Errorf("loadConfig().Formats is a %v, want %v", reflect.TypeOf(got.Formats), reflect.TypeOf(tt.want.Formats))
				}
				// Must have Elements but no check if they are valids
				if reflect.TypeOf(got.Elements) != reflect.TypeOf(tt.want.Elements) {
					t.Errorf("loadConfig().Elements is a %v, want %v", reflect.TypeOf(got.Elements), reflect.TypeOf(tt.want.Elements))
				}
			}
		})
	}
}

func Test_AddExtension(t *testing.T) {
	type args struct {
		id     string
		format string
	}

	tests := []struct {
		name string
		c    config.Config
		args args
		want element.Slice
	}{
		{
			name: "Add osm.pbf but already in",
			c: config.Config{
				Elements: element.Slice{
					"a": element.Element{
						ID:      "a",
						Name:    "a",
						Formats: []string{formats.FormatOsmPbf},
						Meta:    false,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
			args: args{
				id:     "a",
				format: formats.FormatOsmPbf,
			},
			want: element.Slice{
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
			c: config.Config{
				Elements: element.Slice{
					"a": element.Element{
						ID:      "a",
						Name:    "a",
						Formats: []string{},
						Meta:    false,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
			args: args{
				id:     "a",
				format: formats.FormatOsmPbf,
			},
			want: element.Slice{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf on meta",
			c: config.Config{
				Elements: element.Slice{
					"a": element.Element{
						ID:      "a",
						Name:    "a",
						Formats: []string{},
						Meta:    true,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
			args: args{
				id:     "a",
				format: formats.FormatOsmPbf,
			},
			want: element.Slice{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf},
					Meta:    false,
				},
			},
		},
	}

	for tn := range tests {
		tn := tn
		t.Run(tests[tn].name, func(t *testing.T) {
			tests[tn].c.AddExtension(tests[tn].args.id, tests[tn].args.format)
			if !reflect.DeepEqual(tests[tn].c.Elements, tests[tn].want) {
				t.Errorf("AddExtension() got %v, want %v", tests[tn].c.Elements, tests[tn].want)
			}
		})
	}
}

func Benchmark_GetElement_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := config.LoadConfig(geofabrikYml)
	e := "france"

	for n := 0; n < b.N; n++ {
		if _, err := c.GetElement(e); err != nil {
			panic(err)
		}
	}
}

func Test_config_GetElement(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		id      string
		want    *element.Element
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "France from geofabrik",
			file:    geofabrikYml,
			id:      "france",
			want:    &element.Element{ID: "france", Name: "France", Parent: "europe", Formats: element.Formats{formats.FormatOsmPbf, formats.FormatKml, formats.FormatState, "osm.pbf.md5", formats.FormatOsmBz2, "osm.bz2.md5", formats.FormatPoly}},
			wantErr: false,
		},
		{
			name:    "Franc from geofabrik",
			file:    geofabrikYml,
			id:      "franc",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c, _ := config.LoadConfig(tt.file)
			got, err := c.GetElement(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.GetElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.GetElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IsHashable(t *testing.T) {
	type args struct {
		format string
		file   string
	}

	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
		want2 string
	}{
		// TODO: Add test cases.
		{name: "Test is osm.pbf is hashable", args: args{format: formats.FormatOsmPbf, file: geofabrikYml}, want: true, want1: "osm.pbf.md5", want2: "md5"},
		{name: "Test is kml is hashable", args: args{format: formats.FormatKml, file: geofabrikYml}, want: false, want1: "", want2: ""},
	}

	for _, tt := range tests {
		tt := tt

		c, err := config.LoadConfig(tt.args.file)
		if err != nil {
			t.Error(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := config.IsHashable(c, tt.args.format)
			if got != tt.want {
				t.Errorf("config.IsHashable() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("config.IsHashable() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("config.IsHashable() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
func Benchmark_IsHashable_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := config.LoadConfig(geofabrikYml)

	for n := 0; n < b.N; n++ {
		for f := range c.Formats {
			config.IsHashable(c, f)
		}
	}
}
func Benchmark_findElem_parse_all_geofabrik_yml(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)

	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			if _, err := config.FindElem(c, k); err != nil {
				b.Error(err.Error())
			}
		}
	}
}

func Benchmark_GetElement_parse_all_geofabrik_yml(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)

	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			if _, err := c.GetElement(k); err != nil {
				panic(err)
			}
		}
	}
}

func Benchmark_FindElem_parse_France_geofabrik_yml(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)

	for n := 0; n < b.N; n++ {
		if _, err := config.FindElem(c, "france"); err != nil {
			b.Errorf(err.Error())
		}
	}
}

func Test_FindElem(t *testing.T) {
	type args struct {
		c *config.Config
		e string
	}

	tests := []struct {
		name    string
		args    args
		want    *element.Element
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Find",
			args: args{
				c: &SampleConfigValidPtr,
				e: "africa",
			},
			want: &element.Element{
				ID:   "africa",
				Name: "Africa",
				Formats: []string{
					formats.FormatOsmPbf,
					"osm.pbf.md5",
					formats.FormatOsmBz2,
					"osm.bz2.md5",
					formats.FormatOshPbf,
					"osh.pbf.md5",
					formats.FormatPoly,
					formats.FormatKml,
					formats.FormatState,
				},
			},
			wantErr: false,
		},
		{
			name: "Can't find notInList",
			args: args{
				c: &SampleConfigValidPtr,
				e: "notInList",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should not find",
			args: args{
				c: &SampleConfigValidPtr,
				e: "",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.FindElem(tt.args.c, tt.args.e)
			if err != nil != tt.wantErr {
				t.Errorf("config.FindElem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("config.FindElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_elem2URL(t *testing.T) {
	type args struct {
		c   *config.Config
		e   *element.Element
		ext string
	}

	localSampleConfigValidPtr := SampleConfigValidPtr
	localSampleConfigValidPtr.Elements["georgia-us2"] = sampleFakeGeorgiaPtr // add it into config
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "top level test config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleAfricaElementPtr, ext: formats.FormatOsmPbf},
			want:    "https://my.base.url/africa.osm.pbf",
			wantErr: false,
		}, {
			name:    "sub level test config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleGeorgiaUsElementPtr, ext: formats.FormatState},
			want:    "https://my.base.url/../state/north-america/us/georgia-updates/state.txt",
			wantErr: false,
		}, {
			name:    "BaseUrl test config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleGeorgiaUsElementPtr, ext: formats.FormatPoly},
			want:    "http://my.new.url/folder/north-america/us/georgia.poly",
			wantErr: false,
		}, {
			name:    "BaseUrl + BasePath test config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleGeorgiaUsElementPtr, ext: formats.FormatOsmBz2},
			want:    "http://my.new.url/folder/../osmbz2/north-america/us/georgia.osm.bz2",
			wantErr: false,
		},
		{
			name:    "sub level test config wrong format",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleGeorgiaUsElementPtr, ext: "wrongFmt"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "sub level test config not exists in config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleFakeGeorgiaPtr, ext: formats.FormatState},
			want:    "",
			wantErr: true,
		},
		{
			name:    "sub level test config not exists parent",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleFakeGeorgia2Ptr, ext: formats.FormatState},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.Elem2URL(tt.args.c, tt.args.e, tt.args.ext)
			if err != nil != tt.wantErr {
				t.Errorf("elem2URL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("elem2URL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_elem2preURL(t *testing.T) {
	type args struct {
		c *config.Config
		e *element.Element
		b []string
	}

	localSampleConfigValidPtr := SampleConfigValidPtr
	localSampleConfigValidPtr.Elements["georgia-us2"] = sampleFakeGeorgia2Ptr // add it into config
	localSampleConfigValidPtr.Elements["georgia-us3"] = sampleFakeGeorgia3Ptr // add it into config
	localSampleConfigValidPtr.Elements["notus2"] = sampleNotUS2Ptr            // add it into config
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "top level test config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleAfricaElementPtr},
			want:    "https://my.base.url/africa",
			wantErr: false,
		}, {
			name:    "top level test config with basePath base/",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleAfricaElementPtr, b: []string{"base/"}},
			want:    "https://my.base.url/base/africa",
			wantErr: false,
		}, {
			name:    "sub level test config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleGeorgiaUsElementPtr},
			want:    "https://my.base.url/north-america/us/georgia",
			wantErr: false,
		}, {
			name:    "sub level test config not exists in config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleFakeGeorgiaPtr},
			want:    "",
			wantErr: true,
		},
		{
			name:    "sub level test config not exists parent",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleFakeGeorgia2Ptr},
			want:    "",
			wantErr: true,
		},
		{
			name:    "sub level test config parent exist but not grandparent",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleFakeGeorgia3Ptr},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.Elem2preURL(tt.args.c, tt.args.e, tt.args.b...)
			if err != nil != tt.wantErr {
				t.Errorf("config.Elem2preURL() =%v error = %v, wantErr %v", got, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("config.Elem2preURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Elem2preURL_parse_France_geofabrik_yml(b *testing.B) {
	c, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Errorf(err.Error())
	}

	france, err := config.FindElem(c, "france")
	if err != nil {
		b.Errorf(err.Error())
	}

	for n := 0; n < b.N; n++ {
		_, err = config.Elem2preURL(c, france)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}
func Benchmark_Elem2URL_parse_France_geofabrik_yml(b *testing.B) {
	c, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Errorf(err.Error())
	}

	france, err := config.FindElem(c, "france")
	if err != nil {
		b.Errorf(err.Error())
	}

	for n := 0; n < b.N; n++ {
		_, err = config.Elem2URL(c, france, formats.FormatState)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}

func Benchmark_Elem2URL_parse_France_openstreetmap_fr_yml(b *testing.B) {
	c, err := config.LoadConfig(openstreetmapFRYml)
	if err != nil {
		b.Errorf(err.Error())
	}

	france, err := config.FindElem(c, "france")
	if err != nil {
		b.Errorf(err.Error())
	}

	for n := 0; n < b.N; n++ {
		_, err = config.Elem2URL(c, france, formats.FormatPoly)
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}

func TestConfig_MergeElement(t *testing.T) {
	tests := []struct {
		name       string
		Config     *config.Config
		el         *element.Element
		wantErr    bool
		wantConfig *config.Config
	}{
		// TODO: Add test cases.
		{name: "Add element on void Config",
			Config: &config.Config{ElementsMutex: &sync.RWMutex{}, Elements: element.Slice{}},
			el: &element.Element{ID: "test",
				Name:    "Test",
				Formats: element.Formats{"format1", "format2"},
			},
			wantErr: false,
			wantConfig: &config.Config{
				Elements: element.Slice{
					"test": {ID: "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
					}},
				ElementsMutex: &sync.RWMutex{},
			},
		},
		{name: "Add new element on non void Config",
			Config: &config.Config{
				ElementsMutex: &sync.RWMutex{},
				Elements: element.Slice{
					"test2": {ID: "test2",
						Name:    "Test2",
						Formats: element.Formats{"format1", "format2"},
					}},
			},
			el: &element.Element{ID: "test",
				Name:    "Test",
				Formats: element.Formats{"format1", "format2"},
			},
			wantErr: false,
			wantConfig: &config.Config{
				Elements: element.Slice{
					"test": {ID: "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
					},
					"test2": {ID: "test2",
						Name:    "Test2",
						Formats: element.Formats{"format1", "format2"},
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
		},
		{name: "Add same element on non void Config",
			Config: &config.Config{
				ElementsMutex: &sync.RWMutex{},
				Elements: element.Slice{
					"test": {ID: "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
					},
				},
			},
			el: &element.Element{ID: "test",
				Name:    "Test",
				Formats: element.Formats{"format1", "format2"},
			},
			wantErr: false,
			wantConfig: &config.Config{
				Elements: element.Slice{
					"test": {ID: "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
					}},
				ElementsMutex: &sync.RWMutex{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := tt.Config
			err := c.MergeElement(tt.el)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.MergeElement() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(c.Elements, tt.wantConfig.Elements) {
				t.Errorf("Config.MergeElement() config.Elements = %v, wantConfig.Elements %v", c.Elements, tt.wantConfig.Elements)
			}
		})
	}
}
