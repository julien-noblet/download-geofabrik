package element_test

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
)

const (
	geofabrikYml = "../geofabrik.yml"
)

func sampleAfricaElementPtr() *element.Element {
	return &element.Element{
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
}

func sampleGeorgiaUsElementPtr() *element.Element {
	return &element.Element{
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
}

func sampleUsElementPtr() *element.Element {
	return &element.Element{
		ID:     "us",
		Meta:   true,
		Name:   "United States of America",
		Parent: "north-america",
	}
}

func sampleNorthAmericaElementPtr() *element.Element {
	return &element.Element{
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
}

func sampleElementValidPtr() map[string]element.Element {
	return map[string]element.Element{
		"africa":        *sampleAfricaElementPtr(),
		"georgia-us":    *sampleGeorgiaUsElementPtr(),
		"us":            *sampleUsElementPtr(),
		"north-america": *sampleNorthAmericaElementPtr(),
	}
}

func Benchmark_HasParent_parse_geofabrik_yml(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)

	for range make([]struct{}, b.N) {
		for _, v := range c.Elements {
			v.HasParent()
		}
	}
}

func TestElement_HasParent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		fields element.Element
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "us Have parent",
			fields: sampleElementValidPtr()["us"],
			want:   true,
		},
		{
			name:   "Africa Haven't parent",
			fields: sampleElementValidPtr()["Africa"],
			want:   false,
		},
		{
			name:   "Haven't parent 2",
			fields: element.Element{ID: "", File: "", Meta: true, Name: "", Formats: []string(nil), Parent: ""},
			want:   false,
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			myElement := &element.Element{
				ID:      thisTest.fields.ID,
				File:    thisTest.fields.File,
				Meta:    thisTest.fields.Meta,
				Name:    thisTest.fields.Name,
				Formats: thisTest.fields.Formats,
				Parent:  thisTest.fields.Parent,
			}
			if got := myElement.HasParent(); got != thisTest.want {
				t.Errorf("Element.HasParent() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

func Benchmark_contains_parse_geofabrik_yml(b *testing.B) {
	myConfig, _ := config.LoadConfig(geofabrikYml)
	sliceE := element.Formats{}

	for key := range myConfig.Elements {
		sliceE = append(sliceE, key)
	}

	for range make([]struct{}, b.N) {
		for k := range myConfig.Elements {
			sliceE.Contains(k)
		}
	}
}

func Benchmark_contain_parse_geofabrik_yml_France_formats_osm_pbf(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)
	myformats := c.Elements["france"].Formats
	format := formats.FormatOsmPbf

	for range make([]struct{}, b.N) {
		myformats.Contains(format)
	}
}

func Test_MakeParent(t *testing.T) {
	t.Parallel()

	type args struct {
		gparent string
		e       element.Element
	}

	tests := []struct {
		want *element.Element
		name string
		args args
	}{
		{
			name: "No Parents",
			args: args{e: element.Element{ID: "a", Name: "a", Parent: ""}, gparent: ""},
			want: nil,
		},
		{
			name: "Have Parent with no gparent",
			args: args{e: element.Element{ID: "a", Name: "a", Parent: "p"}, gparent: ""},
			want: &element.Element{ID: "p", Name: "p", Meta: true},
		},
		{
			name: "Have Parent with gparent",
			args: args{e: element.Element{ID: "a", Name: "a", Parent: "p"}, gparent: "gp"},
			want: &element.Element{ID: "p", Name: "p", Meta: true, Parent: "gp"},
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got := element.CreateParentElement(&thisTest.args.e, thisTest.args.gparent)
			if got == nil && thisTest.want == nil {
				return
			}

			if got.ID != thisTest.want.ID ||
				got.Name != thisTest.want.Name ||
				got.Meta != thisTest.want.Meta ||
				got.Parent != thisTest.want.Parent {
				t.Errorf("element.MakeParent() = %+v,gparent = %+v, want %+v", got, thisTest.args.gparent, thisTest.want)
			}
		})
	}
}
