package main

import (
	"reflect"
	"testing"
)

var sampleFormatValidPtr = map[string]format{
	//Blank
	"": {
		ID:       "",
		Loc:      "",
		BasePath: "",
	}, "osm.pbf": {
		ID:  "osm.pbf",
		Loc: ".osm.pbf",
		//BasePath: "/",
	}, "state": {
		ID:       "state",
		Loc:      "-updates/state.txt",
		BasePath: "../state/",
	},
}

func Benchmark_miniFormats_parse_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := loadConfig("./geofabrik.yml")
	for n := 0; n < b.N; n++ {
		for _, v := range c.Elements {
			miniFormats(v.Formats)
		}
	}
}
func Test_miniFormats(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "No Formats", args: args{s: *new([]string)}, want: ""},
		{name: "state only", args: args{s: []string{"state"}}, want: "s"},
		{name: "osm.pbf only", args: args{s: []string{"osm.pbf"}}, want: "P"},
		{name: "osm.bz2 only", args: args{s: []string{"osm.bz2"}}, want: "B"},
		{name: "osh.pbf only", args: args{s: []string{"osh.pbf"}}, want: "H"},
		{name: "poly only", args: args{s: []string{"poly"}}, want: "p"},
		{name: "shp.zip only", args: args{s: []string{"shp.zip"}}, want: "S"},
		{name: "kml only", args: args{s: []string{"kml"}}, want: "k"},
		{name: "state and osm.pbf", args: args{s: []string{"osm.pbf", "state"}}, want: "sP"},
		{name: "state and osm.bz2", args: args{s: []string{"state", "osm.bz2"}}, want: "sB"},
		{name: "state and osh.pbf", args: args{s: []string{"osh.pbf", "state"}}, want: "sH"},
		{name: "state and osm.bz2", args: args{s: []string{"state", "osm.bz2"}}, want: "sB"},
		{name: "state and osm.bz2", args: args{s: []string{"state", "osm.bz2"}}, want: "sB"},
		{name: "state and poly", args: args{s: []string{"state", "poly"}}, want: "sp"},
		{name: "state and shp.zip", args: args{s: []string{"state", "shp.zip"}}, want: "sS"},
		{name: "state and kml", args: args{s: []string{"state", "kml"}}, want: "sk"},
		// Not testing all combinaisons!
		{name: "osm.pbf and shp.zip", args: args{s: []string{"osm.pbf", "shp.zip"}}, want: "PS"},
		// With all
		{name: "All formats", args: args{s: []string{"state", "osm.bz2", "osm.pbf", "osh.pbh", "poly", "kml", "shp.zip"}}, want: "sPBpSk"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := miniFormats(tt.args.s); got != tt.want {
				t.Errorf("miniFormats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHashable(t *testing.T) {
	type args struct {
		c      *Config
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
		{name: "Test is osm.pbf is hashable", args: args{format: "osm.pbf", file: "./geofabrik.yml"}, want: true, want1: "osm.pbf.md5", want2: "md5"},
		{name: "Test is kml is hashable", args: args{format: "kml", file: "./geofabrik.yml"}, want: false, want1: "", want2: ""},
	}
	for _, tt := range tests {
		c, err := loadConfig(tt.args.file)
		if err != nil {
			t.Error(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := isHashable(c, tt.args.format)
			if got != tt.want {
				t.Errorf("isHashable() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isHashable() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("isHashable() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
func Benchmark_isHashable_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := loadConfig("./geofabrik.yml")
	for n := 0; n < b.N; n++ {
		for f := range c.Formats {
			isHashable(c, f)
		}
	}
}

func Test_getFormats(t *testing.T) {
	type dflags struct {
		dosmPbf bool
		doshPbf bool
		dosmBz2 bool
		dshpZip bool
		dstate  bool
		dpoly   bool
		dkml    bool
	}
	tests := []struct {
		name  string
		want  []string
		flags dflags
	}{
		// TODO: Add test cases.
		{
			name: "none",
			flags: dflags{
				dosmPbf: false,
				doshPbf: false,
				dosmBz2: false,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{"osm.pbf"},
		}, {
			name: "dosmPbf",
			flags: dflags{
				dosmPbf: true,
				doshPbf: false,
				dosmBz2: false,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{"osm.pbf"},
		}, {
			name: "doshPbf",
			flags: dflags{
				dosmPbf: false,
				doshPbf: true,
				dosmBz2: false,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{"osh.pbf"},
		}, {
			name: "dosmPbf doshPbf",
			flags: dflags{
				dosmPbf: true,
				doshPbf: true,
				dosmBz2: false,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{"osm.pbf", "osh.pbf"},
		}, {
			name: "dosmBz2 dshpZip",
			flags: dflags{
				dosmPbf: false,
				doshPbf: false,
				dosmBz2: true,
				dshpZip: true,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{"osm.bz2", "shp.zip"},
		}, {
			name: "dstate dpoly",
			flags: dflags{
				dosmPbf: false,
				doshPbf: false,
				dosmBz2: false,
				dshpZip: false,
				dstate:  true,
				dpoly:   true,
				dkml:    false,
			},
			want: []string{"state", "poly"},
		}, {
			name: "dkml",
			flags: dflags{
				dosmPbf: false,
				doshPbf: false,
				dosmBz2: false,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    true,
			},
			want: []string{"kml"},
		},
	}
	for _, tt := range tests {
		*dosmPbf = tt.flags.dosmPbf
		*doshPbf = tt.flags.doshPbf
		*dosmBz2 = tt.flags.dosmBz2
		*dshpZip = tt.flags.dshpZip
		*dstate = tt.flags.dstate
		*dpoly = tt.flags.dpoly
		*dkml = tt.flags.dkml
		t.Run(tt.name, func(t *testing.T) {
			if got := getFormats(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("getFormats() = %v, want %v", *got, tt.want)
			}
		})
	}
}
