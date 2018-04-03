package main

import (
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
	}, "poly": {
		ID:      "poly",
		Loc:     ".poly",
		BaseURL: "http://my.new.url/folder",
	}, "osm.bz2": {
		ID:       "osm.bz2",
		Loc:      ".osm.bz2",
		BasePath: "../osmbz2/",
		BaseURL:  "http://my.new.url/folder",
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
