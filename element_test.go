package main

import (
	"reflect"
	"testing"
)

var sampleAfricaElementPtr = Element{
	ID:   "africa",
	Name: "Africa",
	Formats: []string{
		"osm.pbf",
		"osm.pbf.md5",
		"osm.bz2",
		"osm.bz2.md5",
		"osh.pbf",
		"osh.pbf.md5",
		"poly",
		"kml",
		"state",
	},
}
var sampleGeorgiaUsElementPtr = Element{
	ID:   "georgia-us",
	File: "georgia",
	Name: "Georgia (US State)",
	Formats: []string{
		"osm.pbf",
		"osm.pbf.md5",
		"shp.zip",
		"osm.bz2",
		"osm.bz2.md5",
		"osh.pbf",
		"osh.pbf.md5",
		"poly",
		"kml",
		"state",
	},
	Parent: "us",
}
var sampleUsElementPtr = Element{
	ID:     "us",
	Meta:   true,
	Name:   "United States of America",
	Parent: "north-america",
}
var sampleNorthAmericaElementPtr = Element{
	ID:   "north-america",
	Name: "North America",
	Formats: []string{
		"osm.pbf",
		"osm.pbf.md5",
		"osm.bz2",
		"osm.bz2.md5",
		"osh.pbf",
		"osh.pbf.md5",
		"poly",
		"kml",
		"state",
	},
}
var sampleElementValidPtr = map[string]Element{
	"africa":        sampleAfricaElementPtr,
	"georgia-us":    sampleGeorgiaUsElementPtr,
	"us":            sampleUsElementPtr,
	"north-america": sampleNorthAmericaElementPtr,
}

//Creating some fake samples
var sampleFakeGeorgiaPtr = Element{
	ID:   "georgia-usf",
	File: "georgia-fake",
	Name: "Georgia (US State) - fake test",
	Formats: []string{
		"osm.pbf",
		"osm.pbf.md5",
		"shp.zip",
		"osm.bz2",
		"osm.bz2.md5",
		"osh.pbf",
		"osh.pbf.md5",
		"poly",
		"kml",
		"state",
	},
	Parent: "us", // keep good parent!
}
var sampleFakeGeorgia2Ptr = Element{
	ID:   "georgia-us2",
	File: "georgia",
	Name: "Georgia (US State)",
	Formats: []string{
		"osm.pbf",
		"osm.pbf.md5",
		"shp.zip",
		"osm.bz2",
		"osm.bz2.md5",
		"osh.pbf",
		"osh.pbf.md5",
		"poly",
		"kml",
		"state",
	},
	Parent: "notus", // bad parent not exist!
}
var sampleFakeGeorgia3Ptr = Element{
	ID:   "georgia-us3",
	File: "georgia",
	Name: "Georgia (US State)",
	Formats: []string{
		"osm.pbf",
		"osm.pbf.md5",
		"shp.zip",
		"osm.bz2",
		"osm.bz2.md5",
		"osh.pbf",
		"osh.pbf.md5",
		"poly",
		"kml",
		"state",
	},
	Parent: "notus2",
}
var sampleNotUS2Ptr = Element{
	ID:     "notus2",
	Name:   "notus2",
	Meta:   true,
	Parent: "north", // bad parent not exist!
}

func Benchmark_hasParent_parse_geofabrik_yml(b *testing.B) {
	c, _ := loadConfig("./geofabrik.yml")
	for n := 0; n < b.N; n++ {
		for _, v := range c.Elements {
			v.hasParent()
		}
	}
}

func TestElement_hasParent(t *testing.T) {

	tests := []struct {
		name   string
		fields Element
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "us Have parent",
			fields: sampleElementValidPtr["us"],
			want:   true,
		}, {
			name:   "Africa Haven't parent",
			fields: sampleElementValidPtr["Africa"],
			want:   false,
		},
		{
			name:   "Haven't parent 2",
			fields: Element{ID: "", File: "", Meta: true, Name: "", Formats: *new([]string), Parent: ""},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Element{
				ID:      tt.fields.ID,
				File:    tt.fields.File,
				Meta:    tt.fields.Meta,
				Name:    tt.fields.Name,
				Formats: tt.fields.Formats,
				Parent:  tt.fields.Parent,
			}
			if got := e.hasParent(); got != tt.want {
				t.Errorf("Element.hasParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findElem_parse_all_geofabrik_yml(b *testing.B) {
	c, _ := loadConfig("./geofabrik.yml")
	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			findElem(c, k)
		}
	}
}

func Benchmark_GetElement_parse_all_geofabrik_yml(b *testing.B) {
	c, _ := loadConfig("./geofabrik.yml")
	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			if _, err := c.GetElement(k); err != nil {
				panic(err)
			}
		}
	}
}

func Benchmark_findElem_parse_France_geofabrik_yml(b *testing.B) {
	c, _ := loadConfig("./geofabrik.yml")
	for n := 0; n < b.N; n++ {
		findElem(c, "france")
	}
}

func Test_findElem(t *testing.T) {
	type args struct {
		c *Config
		e string
	}
	tests := []struct {
		name    string
		args    args
		want    *Element
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Find",
			args: args{
				c: &SampleConfigValidPtr,
				e: "africa",
			},
			want: &Element{
				ID:   "africa",
				Name: "Africa",
				Formats: []string{
					"osm.pbf",
					"osm.pbf.md5",
					"osm.bz2",
					"osm.bz2.md5",
					"osh.pbf",
					"osh.pbf.md5",
					"poly",
					"kml",
					"state",
				},
			},
			wantErr: false,
		},
		{
			name: "Cant find notInList",
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := findElem(tt.args.c, tt.args.e)
			if err != nil != tt.wantErr {
				t.Errorf("findElem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_stringInSlice_parse_geofabrik_yml(b *testing.B) {
	c, _ := loadConfig("./geofabrik.yml")
	sliceE := elementFormats{}
	for k := range c.Elements {
		sliceE = append(sliceE, k)
	}
	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			stringInSlice(&k, &sliceE)
		}
	}
}
func Benchmark_contains_parse_geofabrik_yml(b *testing.B) {
	c, _ := loadConfig("./geofabrik.yml")
	sliceE := elementFormats{}
	for k := range c.Elements {
		sliceE = append(sliceE, k)
	}
	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			sliceE.contains(k)
		}
	}
}
func Benchmark_stringInSlice_parse_geofabrik_yml_France_formats_osm_pbf(b *testing.B) {
	c, _ := loadConfig("./geofabrik.yml")
	formats := c.Elements["france"].Formats
	format := "osm.pbf"
	for n := 0; n < b.N; n++ {
		stringInSlice(&format, &formats)
	}
}

func Benchmark_contain_parse_geofabrik_yml_France_formats_osm_pbf(b *testing.B) {
	c, _ := loadConfig("./geofabrik.yml")
	formats := c.Elements["france"].Formats
	format := "osm.pbf"
	for n := 0; n < b.N; n++ {
		formats.contains(format)
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		s elementFormats
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "Contain", args: args{s: elementFormats{"a"}, e: "a"}, want: true},
		{name: "Contain", args: args{s: elementFormats{"a", "b"}, e: "a"}, want: true},
		{name: "Not Contain", args: args{s: elementFormats{"a", "b"}, e: "c"}, want: false},
		{name: "Void s", args: args{s: elementFormats{}, e: "c"}, want: false},
		{name: "Void e", args: args{s: elementFormats{"a", "b"}, e: ""}, want: false},
		{name: "Void s,e", args: args{s: elementFormats{}, e: ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.contains(tt.args.e); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
			if got := stringInSlice(&tt.args.e, &tt.args.s); got != tt.want {
				t.Errorf("stringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_elem2URL(t *testing.T) {
	localSampleConfigValidPtr := SampleConfigValidPtr
	localSampleConfigValidPtr.Elements["georgia-us2"] = sampleFakeGeorgiaPtr // add it into config
	type args struct {
		c   *Config
		e   *Element
		ext string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "top level test config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleAfricaElementPtr, ext: "osm.pbf"},
			want:    "https://my.base.url/africa.osm.pbf",
			wantErr: false,
		}, {
			name:    "sub level test config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleGeorgiaUsElementPtr, ext: "state"},
			want:    "https://my.base.url/../state/north-america/us/georgia-updates/state.txt",
			wantErr: false,
		}, {
			name:    "BaseUrl test config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleGeorgiaUsElementPtr, ext: "poly"},
			want:    "http://my.new.url/folder/north-america/us/georgia.poly",
			wantErr: false,
		}, {
			name:    "BaseUrl + BasePath test config",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleGeorgiaUsElementPtr, ext: "osm.bz2"},
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
			args:    args{c: &localSampleConfigValidPtr, e: &sampleFakeGeorgiaPtr, ext: "state"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "sub level test config not exists parent",
			args:    args{c: &localSampleConfigValidPtr, e: &sampleFakeGeorgia2Ptr, ext: "state"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := elem2URL(tt.args.c, tt.args.e, tt.args.ext)
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
	localSampleConfigValidPtr := SampleConfigValidPtr
	localSampleConfigValidPtr.Elements["georgia-us2"] = sampleFakeGeorgia2Ptr // add it into config
	localSampleConfigValidPtr.Elements["georgia-us3"] = sampleFakeGeorgia3Ptr // add it into config
	localSampleConfigValidPtr.Elements["notus2"] = sampleNotUS2Ptr            // add it into config
	type args struct {
		c *Config
		e *Element
		b []string
	}
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := elem2preURL(tt.args.c, tt.args.e, tt.args.b...)
			if err != nil != tt.wantErr {
				t.Errorf("elem2preURL() =%v error = %v, wantErr %v", got, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("elem2preURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Benchmark_elem2preURL_parse_France_geofabrik_yml(b *testing.B) {
	c, err := loadConfig("./geofabrik.yml")
	if err != nil {
		b.Errorf(err.Error())
	}
	france, err := findElem(c, "france")
	if err != nil {
		b.Errorf(err.Error())
	}
	for n := 0; n < b.N; n++ {
		elem2preURL(c, france)
	}
}
func Benchmark_elem2URL_parse_France_geofabrik_yml(b *testing.B) {
	c, err := loadConfig("./geofabrik.yml")
	if err != nil {
		b.Errorf(err.Error())
	}
	france, err := findElem(c, "france")
	if err != nil {
		b.Errorf(err.Error())
	}
	for n := 0; n < b.N; n++ {
		elem2URL(c, france, "state")
	}
}

func Benchmark_elem2URL_parse_France_openstreetmap_fr_yml(b *testing.B) {
	c, err := loadConfig("./openstreetmap.fr.yml")
	if err != nil {
		b.Errorf(err.Error())
	}
	france, err := findElem(c, "france")
	if err != nil {
		b.Errorf(err.Error())
	}
	for n := 0; n < b.N; n++ {
		elem2URL(c, france, "poly")
	}
}
func Test_MakeParent(t *testing.T) {
	type args struct {
		e       Element
		gparent string
	}
	tests := []struct {
		name string
		args args
		want *Element
	}{
		{name: "No Parents", args: args{e: Element{ID: "a", Name: "a"}, gparent: ""}, want: nil},
		{name: "Have Parent with no gparent", args: args{e: Element{ID: "a", Name: "a", Parent: "p"}, gparent: ""}, want: &Element{ID: "p", Name: "p", Meta: true}},
		{name: "Have Parent with gparent", args: args{e: Element{ID: "a", Name: "a", Parent: "p"}, gparent: "gp"}, want: &Element{ID: "p", Name: "p", Meta: true, Parent: "gp"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeParent(tt.args.e, tt.args.gparent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeParent() = %v, want %v", got, tt.want)
			}
		})
	}
}
