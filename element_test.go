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

func Benchmark_stringInSlice_parse_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := loadConfig("./geofabrik.yml")
	sliceE := []string{}
	for k := range c.Elements {
		sliceE = append(sliceE, k)
	}
	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			stringInSlice(&k, &sliceE)
		}
	}
}
func Test_stringInSlice(t *testing.T) {
	testString := "test"
	bString := "b"
	type args struct {
		a    *string
		list *[]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "is in slice",
			args: args{a: &testString, list: &[]string{"this", "is", "a", "test", "slice"}},
			want: true,
		},
		{
			name: "not in slice",
			args: args{a: &bString, list: &[]string{"this", "is", "a", "test", "slice"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("stringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_elem2preURL(t *testing.T) {
	africa := sampleElementValidPtr["africa"]
	georgia := sampleElementValidPtr["georgia-us"]
	type args struct {
		c *Config
		e *Element
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "top level test config",
			args: args{c: &SampleConfigValidPtr, e: &africa},
			want: "https://my.base.url/africa",
		}, {
			name: "sub level test config",
			args: args{c: &SampleConfigValidPtr, e: &georgia},
			want: "https://my.base.url/north-america/us/georgia",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := elem2preURL(tt.args.c, tt.args.e); got != tt.want {
				t.Errorf("elem2preURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_elem2URL(t *testing.T) {
	africa := sampleElementValidPtr["africa"]
	georgia := sampleElementValidPtr["georgia-us"]
	type args struct {
		c   *Config
		e   *Element
		ext string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "top level test config",
			args: args{c: &SampleConfigValidPtr, e: &africa, ext: "osm.pbf"},
			want: "https://my.base.url/africa.osm.pbf",
		}, {
			name: "sub level test config",
			args: args{c: &SampleConfigValidPtr, e: &georgia, ext: "state"},
			want: "https://my.base.url/north-america/us/georgia-updates/state.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := elem2URL(tt.args.c, tt.args.e, tt.args.ext); got != tt.want {
				t.Errorf("elem2URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
