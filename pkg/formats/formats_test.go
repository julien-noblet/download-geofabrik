package formats_test

import (
	"reflect"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
)

func Benchmark_miniFormats_parse_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := config.LoadConfig("../geofabrik.yml")

	for range make([]struct{}, b.N) {
		for _, v := range c.Elements {
			formats.GetMiniFormats(v.Formats)
		}
	}
}

func Test_GetMiniFormats(t *testing.T) {
	t.Parallel()

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
		{name: "state and osm.pbf", args: args{s: []string{formats.FormatOsmPbf, formats.FormatState}}, want: "Ps"},
		{name: "state and osm.bz2", args: args{s: []string{formats.FormatState, formats.FormatOsmBz2}}, want: "sB"},
		{name: "state and osh.pbf", args: args{s: []string{formats.FormatOshPbf, formats.FormatState}}, want: "Hs"},
		{name: "state and osm.bz2", args: args{s: []string{formats.FormatState, formats.FormatOsmBz2}}, want: "sB"},
		{name: "state and osm.bz2", args: args{s: []string{formats.FormatState, formats.FormatOsmBz2}}, want: "sB"},
		{name: "state and poly", args: args{s: []string{formats.FormatState, formats.FormatPoly}}, want: "sp"},
		{name: "state and shp.zip", args: args{s: []string{formats.FormatState, formats.FormatShpZip}}, want: "sS"},
		{name: "state and kml", args: args{s: []string{formats.FormatState, formats.FormatKml}}, want: "sk"},
		{name: "state and geojson", args: args{s: []string{formats.FormatState, formats.FormatGeoJSON}}, want: "sg"},
		{name: "osm.pbf and geojson", args: args{s: []string{formats.FormatOsmPbf, formats.FormatGeoJSON}}, want: "Pg"},
		// Not testing all combinaisons!
		{name: "osm.pbf and shp.zip", args: args{s: []string{formats.FormatOsmPbf, formats.FormatShpZip}}, want: "PS"},
		// With all
		{
			name: "All formats",
			args: args{s: []string{
				formats.FormatState,
				formats.FormatOsmBz2,
				formats.FormatOsmPbf,
				"osh.pbh",
				formats.FormatPoly,
				formats.FormatKml,
				formats.FormatShpZip,
			}},
			want: "sBPpkS",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := formats.GetMiniFormats(tt.args.s); got != tt.want {
				t.Errorf("formats.MiniFormats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFormats(t *testing.T) {
	t.Parallel()

	tests := []struct {
		flags map[string]bool
		name  string
		want  []string
	}{
		{
			name:  "none",
			flags: map[string]bool{},
			want:  []string{formats.FormatOsmPbf},
		},
		{
			name:  "dosmPbf",
			flags: map[string]bool{formats.KeyOsmPbf: true},
			want:  []string{formats.FormatOsmPbf},
		},
		{
			name:  "doshPbf",
			flags: map[string]bool{formats.KeyOshPbf: true},
			want:  []string{formats.FormatOshPbf},
		},
		{
			name: "dosmPbf doshPbf",
			flags: map[string]bool{
				formats.KeyOsmPbf: true,
				formats.KeyOshPbf: true,
			},
			want: []string{formats.FormatOshPbf, formats.FormatOsmPbf},
		},
		{
			name: "dosmBz2 dshpZip",
			flags: map[string]bool{
				formats.KeyOsmBz2: true,
				formats.KeyShpZip: true,
			},
			want: []string{formats.FormatOsmBz2, formats.FormatShpZip},
		},
		{
			name: "dstate dpoly",
			flags: map[string]bool{
				formats.KeyState: true,
				formats.KeyPoly:  true,
			},
			want: []string{formats.FormatPoly, formats.FormatState},
		},
		{
			name:  "dkml",
			flags: map[string]bool{formats.KeyKml: true},
			want:  []string{formats.FormatKml},
		},
		{
			name:  "dosmGz",
			flags: map[string]bool{formats.KeyOsmGz: true},
			want:  []string{formats.FormatOsmGz},
		},
		{
			name:  "dgeojson",
			flags: map[string]bool{formats.KeyGeoJSON: true},
			want:  []string{formats.FormatGeoJSON},
		},
	}

	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			if got := formats.GetFormats(thisTest.flags); !reflect.DeepEqual(got, thisTest.want) {
				t.Errorf("formats.GetFormats() = %v, want %v", got, thisTest.want)
			}
		})
	}
}
