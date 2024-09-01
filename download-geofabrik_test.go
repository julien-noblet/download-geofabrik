package main

import (
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/download"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

//nolint:paralleltest // Can't be parallelized
func Test_checkService(t *testing.T) {
	tests := []struct {
		name       string
		service    string
		config     string
		wantConfig string
		want       bool
	}{
		// TODO: Add test cases.
		{
			name:    "checkService(), fService = geofabrik",
			service: "geofabrik",
			want:    true,
		},
		{
			name:       "checkService(), fService = openstreetmap.fr",
			service:    "openstreetmap.fr",
			config:     "./geofabrik.yml",
			want:       true,
			wantConfig: "./openstreetmap.fr.yml",
		},
		{
			name:       "checkService(), fService = bbbike",
			service:    "bbbike",
			config:     "./geofabrik.yml",
			want:       true,
			wantConfig: "./bbbike.yml",
		},
		{
			name:    "checkService(), fService = anothermap",
			service: "anothermap",
			want:    false,
		},
		{
			name:    "checkService(), fService = \"\"",
			service: "",
			want:    false,
		},
	}
	for _, thisTest := range tests {
		thisTest := thisTest
		*fService = thisTest.service
		t.Run(thisTest.name, func(t *testing.T) {
			if thisTest.config != "" {
				*fConfig = thisTest.config
			}

			if got := checkService(); got != thisTest.want {
				t.Errorf("checkService() = %v, want %v", got, thisTest.want)
			}

			if thisTest.wantConfig != "" && *fConfig != thisTest.wantConfig {
				t.Errorf("checkService() haven't change fConfig, want %v have %v", thisTest.wantConfig, *fConfig)
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

func Test_hashFileMD5(t *testing.T) {
	t.Parallel()

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
		{
			name:    "Check with LICENSE file",
			args:    args{filePath: "./LICENSE"},
			want:    "65d26fcc2f35ea6a181ac777e42db1ea",
			wantErr: false,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := hashFileMD5(thisTest.args.filePath)
			if err != nil != thisTest.wantErr {
				t.Errorf("hashFileMD5(%v) error = %v, wantErr %v", thisTest.args.filePath, err, thisTest.wantErr)

				return
			}

			if got != thisTest.want {
				t.Errorf("hashFileMD5() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

func Benchmark_hashFileMD5_LICENSE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if _, err := hashFileMD5("./LICENSE"); err != nil {
			b.Error(err.Error())
		}
	}
}

func Benchmark_controlHash_LICENSE(b *testing.B) {
	hash, _ := hashFileMD5("./LICENSE")
	hashfile := "/tmp/download-geofabrik-test.hash"

	if err := os.WriteFile(hashfile, []byte(hash), 0o600); err != nil {
		b.Errorf("Can't write file %s err: %v", hashfile, err)
	}

	for n := 0; n < b.N; n++ {
		if _, err := controlHash(hashfile, hash); err != nil {
			b.Error(err.Error())
		}
	}
}

func Test_controlHash(t *testing.T) {
	t.Parallel()

	type args struct {
		hashfile string
		hash     string
	}

	tests := []struct {
		args       args
		name       string
		fileToHash string
		want       bool
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name:       "Check with LICENSE file",
			fileToHash: "./LICENSE",
			args:       args{hashfile: "/tmp/download-geofabrik-test.hash", hash: "65d26fcc2f35ea6a181ac777e42db1ea"},
			want:       true,
			wantErr:    false,
		},
		{
			name:       "Check with LICENSE file wrong hash",
			fileToHash: "./LICENSE",
			args:       args{hashfile: "/tmp/download-geofabrik-test.hash", hash: "65d26fcc2f35ea6a181ac777e42db1eb"},
			want:       false,
			wantErr:    false,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		hash, _ := hashFileMD5(thisTest.fileToHash)

		if err := os.WriteFile(thisTest.args.hashfile, []byte(hash), 0o600); err != nil {
			t.Errorf("can't write file %s err: %v", thisTest.args.hashfile, err)
		}

		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := controlHash(thisTest.args.hashfile, thisTest.args.hash)
			if err != nil != thisTest.wantErr {
				t.Errorf("controlHash() error = %v, wantErr %v", err, thisTest.wantErr)

				return
			}

			if got != thisTest.want {
				t.Errorf("controlHash() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

// Test_downloadChecksum I don't know why sometimes controlHash fail :'(
// seems geofabrik have a limit download I reach sometimes :/.
//
//nolint:paralleltest // Can't be parallelized
func Test_downloadChecksum(t *testing.T) {
	type args struct {
		format string
	}

	*fQuiet = true // be silent!
	tests := []struct {
		name     string
		fConfig  string
		delement string
		args     args
		dCheck   bool
		want     bool
	}{
		// TODO: Add test cases.
		{
			name:     "dCheck = false monaco.osm.pbf from geofabrik",
			dCheck:   false,
			fConfig:  "./geofabrik.yml",
			delement: "monaco",
			args:     args{format: formats.FormatOsmPbf},
			want:     false,
		},
		{
			name:     "dCheck = true monaco.osm.pbf from geofabrik",
			fConfig:  "./geofabrik.yml",
			dCheck:   true,
			delement: "monaco",
			args:     args{format: formats.FormatOsmPbf},
			want:     true,
		},
		{
			name:    "dCheck = true monaco.poly from geofabrik",
			fConfig: "./geofabrik.yml",
			dCheck:  true, delement: "monaco",
			args: args{format: formats.FormatPoly},
			want: false,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		*dCheck = thisTest.dCheck
		*fConfig = thisTest.fConfig
		*delement = thisTest.delement
		t.Run(thisTest.name, func(t *testing.T) {
			if *dCheck { // If I want to compare checksum, Download file
				configPtr, err := config.LoadConfig(*fConfig)
				if err != nil {
					t.Error(err)
				}

				myElem, err := config.FindElem(configPtr, *delement)
				if err != nil {
					t.Error(err)
				}

				myURL, err := config.Elem2URL(configPtr, myElem, thisTest.args.format)
				if err != nil {
					t.Error(err)
				}

				err = download.FromURL(myURL, *delement+"."+thisTest.args.format)
				if err != nil {
					t.Error(err)
				}
			}
			// now real test
			if got := downloadChecksum(thisTest.args.format); got != thisTest.want {
				t.Errorf("downloadChecksum() = %v, want %v", got, thisTest.want)
			}

			os.Remove("monaco.osm.pbf")     // clean
			os.Remove("monaco.osm.pbf.md5") // clean
		})
	}
}

//nolint:paralleltest // Can't be parallelized
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
	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			*lmd = thisTest.lmd
			fakelistAllRegions := func(configuration *config.Config, format string) { //nolint:revive // I need to test format
				assert.Equal(t, thisTest.want, format)
			}
			patch := monkey.Patch(listAllRegions, fakelistAllRegions)

			defer patch.Unpatch()

			listCommand()
		})
	}
}

//nolint:paralleltest // Can't be parallelized
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
		delement      string
		wantURL       string
		wantOutput    string
		formatsFlags  fFlags
		dCheck        bool
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
	for _, thisTest := range tests {
		thisTest := thisTest
		*dosmPbf = thisTest.formatsFlags.dosmPbf || false
		*doshPbf = thisTest.formatsFlags.doshPbf || false
		*dosmBz2 = thisTest.formatsFlags.dosmBz2 || false
		*dshpZip = thisTest.formatsFlags.dshpZip || false
		*dstate = thisTest.formatsFlags.dstate || false
		*dpoly = thisTest.formatsFlags.dpoly || false
		*dkml = thisTest.formatsFlags.dkml || false
		*fConfig = thisTest.fConfig
		*dCheck = thisTest.dCheck || false
		*delement = thisTest.delement
		t.Run(thisTest.name, func(t *testing.T) {
			fakedownloadFromURL := func(myURL string, output string) error {
				assert.Equal(t, thisTest.wantURL, myURL)
				assert.Equal(t, thisTest.wantOutput, output)

				return nil
			}
			fakedownloadChecksum := func(format string) bool { //nolint:revive // I need to test format
				return thisTest.checksumValid
			}
			fakefileExist := func(filePath string) bool { //nolint:revive // I need to test filePath
				return thisTest.fakefileExist
			}
			patch := monkey.Patch(download.FromURL, fakedownloadFromURL)

			defer patch.Unpatch()

			patch2 := monkey.Patch(downloadChecksum, fakedownloadChecksum)

			defer patch2.Unpatch()

			patch3 := monkey.Patch(fileExist, fakefileExist)

			defer patch3.Unpatch()
			downloadCommand()
		})
	}
}

func Test_configureBool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		config string
		flag   bool
	}{
		// TODO: Add test cases.
		{name: "true", flag: true, config: "test1"},
		{name: "false", flag: false, config: "test2"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			configureBool(&tt.flag, tt.config)
			assert.Equal(t, true, viper.IsSet(tt.config), "Is set")
			assert.Equal(t, tt.flag, viper.GetBool(tt.config), "ok")
		})
	}
}
