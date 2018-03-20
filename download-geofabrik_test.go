package main

import (
	"fmt"
	"testing"
)

func Test_checkService(t *testing.T) {
	tests := []struct {
		name       string
		service    string
		config     string
		want       bool
		wantConfig string
	}{
		// TODO: Add test cases.
		{name: "checkService(), fService = geofabrik", service: "geofabrik", want: true},
		{name: "checkService(), fService = openstreetmap.fr", service: "openstreetmap.fr", config: "./geofabrik.yml", want: true, wantConfig: "./openstreetmap.fr.yml"},
		{name: "checkService(), fService = anothermap", service: "anothermap", want: false},
		{name: "checkService(), fService = \"\"", service: "", want: false},
	}
	for _, tt := range tests {
		*fService = tt.service
		t.Run(tt.name, func(t *testing.T) {
			if tt.config != "" {
				*fConfig = tt.config
			}
			if got := checkService(); got != tt.want {
				t.Errorf("checkService() = %v, want %v", got, tt.want)
			}
			if tt.wantConfig != "" && *fConfig != tt.wantConfig {
				t.Errorf("checkService() haven't change fConfig, want %v have %v", tt.wantConfig, *fConfig)
			}
		})
	}
}

/*func Benchmark_listAllRegions_parse_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := loadConfig("./geofabrik.yml")
	for n := 0; n < b.N; n++ {
		listAllRegions(*c, "")
	}
}*/
/*
func Benchmark_listAllRegions_parse_geofabrik_yml_md(b *testing.B) {
	// run the Fib function b.N times
	c, _ := loadConfig("./geofabrik.yml")
	for n := 0; n < b.N; n++ {
		listAllRegions(*c, "Markdown")
	}
}
*/
func Test_catch(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "should panic", args: args{err: fmt.Errorf("test")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f", r)
				}
			}()
			catch(tt.args.err)
			t.Errorf("The code did not panic")
		})
	}
}
