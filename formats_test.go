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
		ID:  "state",
		Loc: "-updates/state.txt",
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
