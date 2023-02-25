package element_test

import (
	"reflect"
	"testing"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
)

const (
	geofabrikYml = "../geofabrik.yml"
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

func Benchmark_HasParent_parse_geofabrik_yml(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)

	for n := 0; n < b.N; n++ {
		for _, v := range c.Elements {
			v.HasParent()
		}
	}
}

func TestElement_HasParent(t *testing.T) {
	tests := []struct {
		name   string
		fields element.Element
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "us Have parent",
			fields: sampleElementValidPtr["us"],
			want:   true,
		},
		{
			name:   "Africa Haven't parent",
			fields: sampleElementValidPtr["Africa"],
			want:   false,
		},
		{
			name:   "Haven't parent 2",
			fields: element.Element{ID: "", File: "", Meta: true, Name: "", Formats: []string(nil), Parent: ""},
			want:   false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			e := &element.Element{
				ID:      tt.fields.ID,
				File:    tt.fields.File,
				Meta:    tt.fields.Meta,
				Name:    tt.fields.Name,
				Formats: tt.fields.Formats,
				Parent:  tt.fields.Parent,
			}
			if got := e.HasParent(); got != tt.want {
				t.Errorf("Element.HasParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_StringInSlice_parse_geofabrik_yml(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)
	sliceE := element.Formats{}

	for k := range c.Elements {
		sliceE = append(sliceE, k)
	}

	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			k := k
			element.StringInSlice(&k, &sliceE)
		}
	}
}

func Benchmark_contains_parse_geofabrik_yml(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)
	sliceE := element.Formats{}

	for k := range c.Elements {
		sliceE = append(sliceE, k)
	}

	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			sliceE.Contains(k)
		}
	}
}

func Benchmark_StringInSlice_parse_geofabrik_yml_France_formats_osm_pbf(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)
	myformats := c.Elements["france"].Formats
	format := formats.FormatOsmPbf

	for n := 0; n < b.N; n++ {
		element.StringInSlice(&format, &myformats)
	}
}

func Benchmark_contain_parse_geofabrik_yml_France_formats_osm_pbf(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)
	myformats := c.Elements["france"].Formats
	format := formats.FormatOsmPbf

	for n := 0; n < b.N; n++ {
		myformats.Contains(format)
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		e string
		s element.Formats
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "Contain", args: args{s: element.Formats{"a"}, e: "a"}, want: true},
		{name: "Contain", args: args{s: element.Formats{"a", "b"}, e: "a"}, want: true},
		{name: "Not Contain", args: args{s: element.Formats{"a", "b"}, e: "c"}, want: false},
		{name: "Void s", args: args{s: element.Formats{}, e: "c"}, want: false},
		{name: "Void e", args: args{s: element.Formats{"a", "b"}, e: ""}, want: false},
		{name: "Void s,e", args: args{s: element.Formats{}, e: ""}, want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Contains(tt.args.e); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
			if got := element.StringInSlice(&tt.args.e, &tt.args.s); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MakeParent(t *testing.T) {
	type args struct {
		gparent string
		e       element.Element
	}

	tests := []struct {
		want *element.Element
		name string
		args args
	}{
		{name: "No Parents", args: args{e: element.Element{ID: "a", Name: "a"}, gparent: ""}, want: nil},
		{name: "Have Parent with no gparent", args: args{e: element.Element{ID: "a", Name: "a", Parent: "p"}, gparent: ""}, want: &element.Element{ID: "p", Name: "p", Meta: true}},
		{name: "Have Parent with gparent", args: args{e: element.Element{ID: "a", Name: "a", Parent: "p"}, gparent: "gp"}, want: &element.Element{ID: "p", Name: "p", Meta: true, Parent: "gp"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := element.MakeParent(&tt.args.e, tt.args.gparent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("element.MakeParent() = %v, want %v", got, tt.want)
			}
		})
	}
}
