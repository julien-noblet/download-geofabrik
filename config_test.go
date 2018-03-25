package main

import (
	"reflect"
	"testing"
)

var SampleConfigValidPtr = Config{
	BaseURL:  "https://my.base.url",
	Formats:  sampleFormatValidPtr,
	Elements: sampleElementValidPtr,
}

func Benchmark_loadConfig_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		loadConfig("./geofabrik.yml")
	}
}

func Benchmark_loadConfig_osmfr_yml(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		loadConfig("./openstreetmap.fr.yml")
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
				BaseURL:  "https://download.geofabrik.de",
				Formats:  sampleFormatValidPtr,
				Elements: sampleElementValidPtr,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadConfig(tt.args.configFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			/*if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadConfig() = %v, want %v", got, tt.want)
			}*/
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
