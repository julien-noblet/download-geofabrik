package main

import (
	"reflect"
	"testing"
)

var sampleFormatValidPtr = map[string]format{
	// Blank
	"": {
		ID:       "",
		Loc:      "",
		BasePath: "",
	}, formatOsmPbf: {
		ID:  formatOsmPbf,
		Loc: ".osm.pbf",
		//BasePath: "/",
	}, formatState: {
		ID:       formatState,
		Loc:      "-updates/state.txt",
		BasePath: "../state/",
	}, formatPoly: {
		ID:      formatPoly,
		Loc:     ".poly",
		BaseURL: "http://my.new.url/folder",
	}, formatOsmBz2: {
		ID:       formatOsmBz2,
		Loc:      ".osm.bz2",
		BasePath: "../osmbz2/",
		BaseURL:  "http://my.new.url/folder",
	}, formatOsmGz: {
		ID:       formatOsmGz,
		Loc:      ".osm.gz",
		BasePath: "../osmgz/",
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
		{name: "No Formats", args: args{s: []string(nil)}, want: ""},
		{name: "state only", args: args{s: []string{formatState}}, want: "s"},
		{name: "osm.pbf only", args: args{s: []string{formatOsmPbf}}, want: "P"},
		{name: "osm.bz2 only", args: args{s: []string{formatOsmBz2}}, want: "B"},
		{name: "osm.gz only", args: args{s: []string{formatOsmGz}}, want: "G"},
		{name: "osh.pbf only", args: args{s: []string{formatOshPbf}}, want: "H"},
		{name: "poly only", args: args{s: []string{formatPoly}}, want: "p"},
		{name: "shp.zip only", args: args{s: []string{formatShpZip}}, want: "S"},
		{name: "kml only", args: args{s: []string{formatKml}}, want: "k"},
		{name: "state and osm.pbf", args: args{s: []string{formatOsmPbf, formatState}}, want: "sP"},
		{name: "state and osm.bz2", args: args{s: []string{formatState, formatOsmBz2}}, want: "sB"},
		{name: "state and osh.pbf", args: args{s: []string{formatOshPbf, formatState}}, want: "sH"},
		{name: "state and osm.bz2", args: args{s: []string{formatState, formatOsmBz2}}, want: "sB"},
		{name: "state and osm.bz2", args: args{s: []string{formatState, formatOsmBz2}}, want: "sB"},
		{name: "state and poly", args: args{s: []string{formatState, formatPoly}}, want: "sp"},
		{name: "state and shp.zip", args: args{s: []string{formatState, formatShpZip}}, want: "sS"},
		{name: "state and kml", args: args{s: []string{formatState, formatKml}}, want: "sk"},
		// Not testing all combinaisons!
		{name: "osm.pbf and shp.zip", args: args{s: []string{formatOsmPbf, formatShpZip}}, want: "PS"},
		// With all
		{name: "All formats", args: args{s: []string{formatState, formatOsmBz2, formatOsmPbf, "osh.pbh", formatPoly, formatKml, formatShpZip}}, want: "sPBpSk"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := miniFormats(tt.args.s); got != tt.want {
				t.Errorf("miniFormats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHashable(t *testing.T) {
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
		{name: "Test is osm.pbf is hashable", args: args{format: formatOsmPbf, file: "./geofabrik.yml"}, want: true, want1: "osm.pbf.md5", want2: "md5"},
		{name: "Test is kml is hashable", args: args{format: formatKml, file: "./geofabrik.yml"}, want: false, want1: "", want2: ""},
	}

	for _, tt := range tests {
		tt := tt

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
		dosmGz  bool
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
				dosmGz:  false,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{formatOsmPbf},
		}, {
			name: "dosmPbf",
			flags: dflags{
				dosmPbf: true,
				doshPbf: false,
				dosmBz2: false,
				dosmGz:  false,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{formatOsmPbf},
		}, {
			name: "doshPbf",
			flags: dflags{
				dosmPbf: false,
				doshPbf: true,
				dosmBz2: false,
				dosmGz:  false,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{formatOshPbf},
		}, {
			name: "dosmPbf doshPbf",
			flags: dflags{
				dosmPbf: true,
				doshPbf: true,
				dosmBz2: false,
				dosmGz:  false,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{formatOsmPbf, formatOshPbf},
		}, {
			name: "dosmBz2 dshpZip",
			flags: dflags{
				dosmPbf: false,
				doshPbf: false,
				dosmBz2: true,
				dosmGz:  false,
				dshpZip: true,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{formatOsmBz2, formatShpZip},
		}, {
			name: "dstate dpoly",
			flags: dflags{
				dosmPbf: false,
				doshPbf: false,
				dosmBz2: false,
				dosmGz:  false,
				dshpZip: false,
				dstate:  true,
				dpoly:   true,
				dkml:    false,
			},
			want: []string{formatState, formatPoly},
		}, {
			name: "dkml",
			flags: dflags{
				dosmPbf: false,
				doshPbf: false,
				dosmBz2: false,
				dosmGz:  false,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    true,
			},
			want: []string{formatKml},
		},
		{
			name: "dosmGz",
			flags: dflags{
				dosmPbf: false,
				doshPbf: false,
				dosmBz2: false,
				dosmGz:  true,
				dshpZip: false,
				dstate:  false,
				dpoly:   false,
				dkml:    false,
			},
			want: []string{formatOsmGz},
		},
	}

	for _, tt := range tests {
		tt := tt
		*dosmPbf = tt.flags.dosmPbf
		*doshPbf = tt.flags.doshPbf
		*dosmBz2 = tt.flags.dosmBz2
		*dosmGz = tt.flags.dosmGz
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
