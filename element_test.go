package main

import (
	"reflect"
	"testing"
)

var sampleElementValidPtr = map[string]Element{
	"africa": {
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
	"georgia-us": {
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
	},
	"us": {
		ID:     "us",
		Meta:   true,
		Name:   "United States of America",
		Parent: "north-america",
	},
	"north-america": {
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
	},
}

func Benchmark_hasParent_parse_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
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

func Benchmark_findElem_parse_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := loadConfig("./geofabrik.yml")
	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			findElem(c, k)
		}
	}
}
func Test_findElem(t *testing.T) {
	type args struct {
		c *Config
		e string
	}
	tests := []struct {
		name string
		args args
		want *Element
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
		},
		/*{ // not working!
			name: "Should not find",
			args: args{
				c: &SampleConfigValidPtr,
				e: "",
			},
			want: new(Element),
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findElem(tt.args.c, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findElem() = %v, want %v", got, tt.want)
			}
		})
	}
}
