package formats_test

import (
	"reflect"
	"testing"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
)

func Benchmark_miniFormats_parse_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := config.LoadConfig("../geofabrik.yml")

	for n := 0; n < b.N; n++ {
		for _, v := range c.Elements {
			formats.MiniFormats(v.Formats)
		}
	}
}
func Test_miniFormats(t *testing.T) {
	type args struct {
		s []string
	}

	tests := []struct {
		name string
		want string
		args args
	}{
		// TODO: Add test cases.
		{name: "No Formats", args: args{s: []string(nil)}, want: ""},
		{name: "state only", args: args{s: []string{formats.FormatState}}, want: "s"},
		{name: "osm.pbf only", args: args{s: []string{formats.FormatOsmPbf}}, want: "P"},
		{name: "osm.bz2 only", args: args{s: []string{formats.FormatOsmBz2}}, want: "B"},
		{name: "osm.gz only", args: args{s: []string{formats.FormatOsmGz}}, want: "G"},
		{name: "osh.pbf only", args: args{s: []string{formats.FormatOshPbf}}, want: "H"},
		{name: "poly only", args: args{s: []string{formats.FormatPoly}}, want: "p"},
		{name: "shp.zip only", args: args{s: []string{formats.FormatShpZip}}, want: "S"},
		{name: "kml only", args: args{s: []string{formats.FormatKml}}, want: "k"},
		{name: "state and osm.pbf", args: args{s: []string{formats.FormatOsmPbf, formats.FormatState}}, want: "sP"},
		{name: "state and osm.bz2", args: args{s: []string{formats.FormatState, formats.FormatOsmBz2}}, want: "sB"},
		{name: "state and osh.pbf", args: args{s: []string{formats.FormatOshPbf, formats.FormatState}}, want: "sH"},
		{name: "state and osm.bz2", args: args{s: []string{formats.FormatState, formats.FormatOsmBz2}}, want: "sB"},
		{name: "state and osm.bz2", args: args{s: []string{formats.FormatState, formats.FormatOsmBz2}}, want: "sB"},
		{name: "state and poly", args: args{s: []string{formats.FormatState, formats.FormatPoly}}, want: "sp"},
		{name: "state and shp.zip", args: args{s: []string{formats.FormatState, formats.FormatShpZip}}, want: "sS"},
		{name: "state and kml", args: args{s: []string{formats.FormatState, formats.FormatKml}}, want: "sk"},
		// Not testing all combinaisons!
		{name: "osm.pbf and shp.zip", args: args{s: []string{formats.FormatOsmPbf, formats.FormatShpZip}}, want: "PS"},
		// With all
		{name: "All formats", args: args{s: []string{formats.FormatState, formats.FormatOsmBz2, formats.FormatOsmPbf, "osh.pbh", formats.FormatPoly, formats.FormatKml, formats.FormatShpZip}}, want: "sPBpSk"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := formats.MiniFormats(tt.args.s); got != tt.want {
				t.Errorf("formats.MiniFormats() = %v, want %v", got, tt.want)
			}
		})
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
			want: []string{formats.FormatOsmPbf},
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
			want: []string{formats.FormatOsmPbf},
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
			want: []string{formats.FormatOshPbf},
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
			want: []string{formats.FormatOsmPbf, formats.FormatOshPbf},
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
			want: []string{formats.FormatOsmBz2, formats.FormatShpZip},
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
			want: []string{formats.FormatState, formats.FormatPoly},
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
			want: []string{formats.FormatKml},
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
			want: []string{formats.FormatOsmGz},
		},
	}

	for _, tt := range tests {
		tt := tt
		viper.Set("dosmPbf", tt.flags.dosmPbf)
		viper.Set("doshPbf", tt.flags.doshPbf)
		viper.Set("dosmBz2", tt.flags.dosmBz2)
		viper.Set("dosmGz", tt.flags.dosmGz)
		viper.Set("dshpZip", tt.flags.dshpZip)
		viper.Set("dstate", tt.flags.dstate)
		viper.Set("dpoly", tt.flags.dpoly)
		viper.Set("dkml", tt.flags.dkml)
		t.Run(tt.name, func(t *testing.T) {
			if got := formats.GetFormats(); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("formats.GetFormats() = %v, want %v", *got, tt.want)
			}
		})
	}
}
