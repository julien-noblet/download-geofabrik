package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
					_, err := fmt.Println("Recovered in f", r)
					if err != nil {
						t.Errorf(err.Error())
					}
				}
			}()
			catch(tt.args.err)
			t.Errorf("The code did not panic")
		})
	}
}

func Test_hashFileMD5(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Check with LICENSE file", args: args{filePath: "./LICENSE"}, want: "65d26fcc2f35ea6a181ac777e42db1ea", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashFileMD5(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("hashFileMD5(%v) error = %v, wantErr %v", tt.args.filePath, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hashFileMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_hashFileMD5_LICENSE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		hashFileMD5("./LICENSE")
	}
}
func Benchmark_controlHash_LICENSE(b *testing.B) {
	hash, _ := hashFileMD5("./LICENSE")
	hashfile := "/tmp/download-geofabrik-test.hash"
	ioutil.WriteFile(hashfile, []byte(hash), 0644)
	for n := 0; n < b.N; n++ {
		controlHash(hashfile, hash)
	}
}

func Test_controlHash(t *testing.T) {
	type args struct {
		hashfile string
		hash     string
	}
	tests := []struct {
		name       string
		args       args
		want       bool
		wantErr    bool
		fileToHash string
	}{
		// TODO: Add test cases.
		{name: "Check with LICENSE file", fileToHash: "./LICENSE", args: args{hashfile: "/tmp/download-geofabrik-test.hash", hash: "65d26fcc2f35ea6a181ac777e42db1ea"}, want: true, wantErr: false},
		{name: "Check with LICENSE file wrong hash", fileToHash: "./LICENSE", args: args{hashfile: "/tmp/download-geofabrik-test.hash", hash: "65d26fcc2f35ea6a181ac777e42db1eb"}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		hash, _ := hashFileMD5(tt.fileToHash)
		ioutil.WriteFile(tt.args.hashfile, []byte(hash), 0644)
		t.Run(tt.name, func(t *testing.T) {
			got, err := controlHash(tt.args.hashfile, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("controlHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("controlHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_downloadChecksum(t *testing.T) {
	*fQuiet = true // be silent!
	type args struct {
		format string
	}
	tests := []struct {
		name     string
		fConfig  string
		dCheck   bool
		delement string
		args     args
		want     bool
	}{
		// TODO: Add test cases.
		{name: "dCheck = false", dCheck: false, want: false},
		{name: "dCheck = true andorra.osm.pbf from geofabrik", fConfig: "./geofabrik.yml", dCheck: true, delement: "andorra", args: args{format: "osm.pbf"}, want: false},
	}
	for _, tt := range tests {
		*dCheck = tt.dCheck
		*fConfig = tt.fConfig
		*delement = tt.delement
		t.Run(tt.name, func(t *testing.T) {
			if got := downloadChecksum(tt.args.format); got != tt.want {
				t.Errorf("downloadChecksum() = %v, want %v", got, tt.want)
			}
			os.Remove("andorra.osm.pbf.md5") // clean
		})
	}
}
