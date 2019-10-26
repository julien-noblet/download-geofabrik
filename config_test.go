package main

import (
	"reflect"
	"sync"
	"testing"
)

var SampleConfigValidPtr = Config{
	BaseURL:  "https://my.base.url",
	Formats:  sampleFormatValidPtr,
	Elements: sampleElementValidPtr,
}

func Benchmark_Exist_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := loadConfig("./geofabrik.yml")
	e := "france"

	for n := 0; n < b.N; n++ {
		c.Exist(e)
	}
}

func Benchmark_loadConfig_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		if _, err := loadConfig("./geofabrik.yml"); err != nil {
			b.Error(err.Error())
		}
	}
}

func Benchmark_loadConfig_osmfr_yml(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		if _, err := loadConfig("./openstreetmap.fr.yml"); err != nil {
			b.Error(err.Error())
		}
	}
}

func Test_loadConfig(t *testing.T) {
	type args struct {
		configFile string
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Check config file not exist",
			args:    args{configFile: "./this_file_not_exists"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Check config not yaml",
			args:    args{configFile: "./LICENSE"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check config valid geofabrik.yml",
			args: args{configFile: "./geofabrik.yml"},
			want: &Config{
				BaseURL:       "https://download.geofabrik.de",
				Formats:       sampleFormatValidPtr,
				Elements:      sampleElementValidPtr,
				ElementsMutex: &sync.RWMutex{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadConfig(tt.args.configFile)
			if err != nil != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
					t.Errorf("loadConfig() is a %v, want %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
				}
				// Check BaseURL
				if got.BaseURL != tt.want.BaseURL {
					t.Errorf("loadConfig().BaseURL = %v, want %v", got.BaseURL, tt.want.BaseURL)
				}
				// Must have formats but ne need to check if they are valids!
				if reflect.TypeOf(got.Formats) != reflect.TypeOf(tt.want.Formats) {
					t.Errorf("loadConfig().Formats is a %v, want %v", reflect.TypeOf(got.Formats), reflect.TypeOf(tt.want.Formats))
				}
				// Must have Elements but no check if they are valids
				if reflect.TypeOf(got.Elements) != reflect.TypeOf(tt.want.Elements) {
					t.Errorf("loadConfig().Elements is a %v, want %v", reflect.TypeOf(got.Elements), reflect.TypeOf(tt.want.Elements))
				}
			}
		})
	}
}

func Test_AddExtension(t *testing.T) {
	type args struct {
		id     string
		format string
	}

	tests := []struct {
		name string
		c    Config
		args args
		want ElementSlice
	}{
		{
			name: "Add osm.pbf but already in",
			c: Config{
				Elements: ElementSlice{
					"a": Element{
						ID:      "a",
						Name:    "a",
						Formats: []string{formatOsmPbf},
						Meta:    false,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
			args: args{
				id:     "a",
				format: formatOsmPbf,
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf",
			c: Config{
				Elements: ElementSlice{
					"a": Element{
						ID:      "a",
						Name:    "a",
						Formats: []string{},
						Meta:    false,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
			args: args{
				id:     "a",
				format: formatOsmPbf,
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf on meta",
			c: Config{
				Elements: ElementSlice{
					"a": Element{
						ID:      "a",
						Name:    "a",
						Formats: []string{},
						Meta:    true,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
			args: args{
				id:     "a",
				format: formatOsmPbf,
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formatOsmPbf},
					Meta:    false,
				},
			},
		},
	}

	for tn := range tests {
		tn := tn
		t.Run(tests[tn].name, func(t *testing.T) {
			tests[tn].c.AddExtension(tests[tn].args.id, tests[tn].args.format)
			if !reflect.DeepEqual(tests[tn].c.Elements, tests[tn].want) {
				t.Errorf("AddExtension() got %v, want %v", tests[tn].c.Elements, tests[tn].want)
			}
		})
	}
}

func Benchmark_GetElement_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := loadConfig("./geofabrik.yml")
	e := "france"

	for n := 0; n < b.N; n++ {
		if _, err := c.GetElement(e); err != nil {
			panic(err)
		}
	}
}

func TestConfig_GetElement(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		id      string
		want    *Element
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "France from geofabrik",
			file:    "./geofabrik.yml",
			id:      "france",
			want:    &Element{ID: "france", Name: "France", Parent: "europe", Formats: elementFormats{formatOsmPbf, formatKml, formatState, "osm.pbf.md5", formatOsmBz2, "osm.bz2.md5", formatPoly}},
			wantErr: false,
		},
		{
			name:    "Franc from geofabrik",
			file:    "./geofabrik.yml",
			id:      "franc",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c, _ := loadConfig(tt.file)
			got, err := c.GetElement(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.GetElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.GetElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
