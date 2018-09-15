package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
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
		{name: "checkService(), fService = gislab", service: "gislab", config: "./geofabrik.yml", want: true, wantConfig: "./gislab.yml"},
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
		want string
	}{
		// TODO: Add test cases.
		{name: "should display error", args: args{err: fmt.Errorf("test")}, want: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For testing log, need to use Monkey patching
			fakeLogFatal := func(msg ...interface{}) {
				assert.Equal(t, tt.want, msg[0])
			}
			patch := monkey.Patch(log.Fatalln, fakeLogFatal)
			defer patch.Unpatch()
			catch(tt.args.err)
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
			if err != nil != tt.wantErr {
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
			if err != nil != tt.wantErr {
				t.Errorf("controlHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("controlHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_downloadChecksum I don't know why sometimes controlHash fail :'(
// seems geofabrik have a limit download I reach sometimes :/
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
		{name: "dCheck = false monaco.osm.pbf from geofabrik", dCheck: false, fConfig: "./geofabrik.yml", delement: "monaco", args: args{format: "osm.pbf"}, want: false},
		{name: "dCheck = true monaco.osm.pbf from geofabrik", fConfig: "./geofabrik.yml", dCheck: true, delement: "monaco", args: args{format: "osm.pbf"}, want: true},
		{name: "dCheck = true monaco.poly from geofabrik", fConfig: "./geofabrik.yml", dCheck: true, delement: "monaco", args: args{format: "poly"}, want: false},
	}
	for _, tt := range tests {
		*dCheck = tt.dCheck
		*fConfig = tt.fConfig
		*delement = tt.delement
		t.Run(tt.name, func(t *testing.T) {
			if *dCheck { // If I want to compare checksum, Download file
				configPtr, err := loadConfig(*fConfig)
				if err != nil {
					t.Error(err)
				}
				myElem, err := findElem(configPtr, *delement)
				if err != nil {
					t.Error(err)
				}
				myURL, err := elem2URL(configPtr, myElem, tt.args.format)
				if err != nil {
					t.Error(err)
				}
				downloadFromURL(myURL, *delement+"."+tt.args.format)
			}
			// now real test
			if got := downloadChecksum(tt.args.format); got != tt.want {
				t.Errorf("downloadChecksum() = %v, want %v", got, tt.want)
			}
			os.Remove("monaco.osm.pbf")     // clean
			os.Remove("monaco.osm.pbf.md5") // clean
		})
	}
}

func Test_listCommand(t *testing.T) {
	tests := []struct {
		name string
		want string
		lmd  bool
	}{
		// TODO: Add test cases.
		{name: "List normal", want: ""},
		{name: "List Markdown", lmd: true, want: "Markdown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			*lmd = tt.lmd
			fakelistAllRegions := func(configPtr Config, format string) {
				assert.Equal(t, tt.want, format)
			}
			patch := monkey.Patch(listAllRegions, fakelistAllRegions)
			defer patch.Unpatch()
			listCommand()
		})
	}
}

func Test_downloadCommand(t *testing.T) {
	type fFlags struct {
		dosmPbf bool
		doshPbf bool
		dosmBz2 bool
		dshpZip bool
		dstate  bool
		dpoly   bool
		dkml    bool
	}

	tests := []struct {
		name          string
		fConfig       string
		dCheck        bool
		delement      string
		formatsFlags  fFlags
		wantURL       string
		wantOutput    string
		checksumValid bool
		fakefileExist bool
	}{
		// TODO: Add test cases.
		{
			name:     "monaco.osm.pbf from geofabrik.yml no Check",
			fConfig:  "geofabrik.yml",
			dCheck:   false,
			delement: "monaco",
			formatsFlags: fFlags{
				dosmPbf: true,
			},
			wantURL:    "https://download.geofabrik.de/europe/monaco-latest.osm.pbf",
			wantOutput: "monaco.osm.pbf",
		},
		{
			name:     "monaco.osm.pbf from geofabrik.yml with Check",
			fConfig:  "geofabrik.yml",
			dCheck:   true,
			delement: "monaco",
			formatsFlags: fFlags{
				dosmPbf: true,
			},
			wantURL:       "https://download.geofabrik.de/europe/monaco-latest.osm.pbf",
			wantOutput:    "monaco.osm.pbf",
			fakefileExist: false,
			checksumValid: true,
		},
		{
			name:     "monaco.osm.pbf from geofabrik.yml with Check file exist",
			fConfig:  "geofabrik.yml",
			dCheck:   true,
			delement: "monaco",
			formatsFlags: fFlags{
				dosmPbf: true,
			},
			wantURL:       "https://download.geofabrik.de/europe/monaco-latest.osm.pbf",
			wantOutput:    "monaco.osm.pbf",
			fakefileExist: true,
			checksumValid: true,
		},
		{
			name:     "monaco.osm.pbf from geofabrik.yml with Check checksum mismatch",
			fConfig:  "geofabrik.yml",
			dCheck:   true,
			delement: "monaco",
			formatsFlags: fFlags{
				dosmPbf: true,
			},
			wantURL:       "https://download.geofabrik.de/europe/monaco-latest.osm.pbf",
			wantOutput:    "monaco.osm.pbf",
			fakefileExist: false,
			checksumValid: false,
		},
		{
			name:     "monaco.osm.pbf from geofabrik.yml with Check file exist checksum mismatch",
			fConfig:  "geofabrik.yml",
			dCheck:   true,
			delement: "monaco",
			formatsFlags: fFlags{
				dosmPbf: true,
			},
			wantURL:       "https://download.geofabrik.de/europe/monaco-latest.osm.pbf",
			wantOutput:    "monaco.osm.pbf",
			fakefileExist: true,
			checksumValid: false,
		},
	}
	for _, tt := range tests {
		*dosmPbf = tt.formatsFlags.dosmPbf || false
		*doshPbf = tt.formatsFlags.doshPbf || false
		*dosmBz2 = tt.formatsFlags.dosmBz2 || false
		*dshpZip = tt.formatsFlags.dshpZip || false
		*dstate = tt.formatsFlags.dstate || false
		*dpoly = tt.formatsFlags.dpoly || false
		*dkml = tt.formatsFlags.dkml || false
		*fConfig = tt.fConfig
		*dCheck = tt.dCheck || false
		*delement = tt.delement
		t.Run(tt.name, func(t *testing.T) {
			fakedownloadFromURL := func(myURL string, output string) error {
				assert.Equal(t, tt.wantURL, myURL)
				assert.Equal(t, tt.wantOutput, output)
				return nil
			}
			fakedownloadChecksum := func(f string) bool {
				return tt.checksumValid
			}
			fakefileExist := func(f string) bool {
				return tt.fakefileExist
			}
			patch := monkey.Patch(downloadFromURL, fakedownloadFromURL)
			patch2 := monkey.Patch(downloadChecksum, fakedownloadChecksum)
			patch3 := monkey.Patch(fileExist, fakefileExist)
			defer patch.Unpatch()
			defer patch2.Unpatch()
			defer patch3.Unpatch()
			downloadCommand()
			//			izefnof // real error should not compile
		})
	}
}
