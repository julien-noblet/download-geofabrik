package download_test

import (
	"os"
	"sync"
	"testing"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/download"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
)

//nolint:paralleltest // Can't be parallelized
func Test_DownloadFromURL(t *testing.T) {
	type args struct {
		myURL    string
		fileName string
	}

	tests := []struct {
		name        string
		args        args
		fNodownload bool
		fQuiet      bool
		fProgress   bool
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name: "try fNodownload=true",
			args: args{
				myURL:    "https://download.geofabrik.de/this_url_should_not_exist",
				fileName: "/tmp/download-geofabrik.test",
			},
			fNodownload: true,
			wantErr:     false,
		},
		{
			name:        "404 error from geofabrik",
			fNodownload: false,
			args: args{
				myURL:    "https://download.geofabrik.de/this_url_should_not_exist",
				fileName: "/tmp/download-geofabrik.test",
			},
			wantErr: true,
		},
		{
			name:        "OK download from geofabrik",
			fNodownload: false,
			fQuiet:      false,
			fProgress:   true,
			args: args{
				myURL:    "https://download.geofabrik.de/europe/andorra.poly",
				fileName: "/tmp/download-geofabrik.test",
			},
			wantErr: false,
		},
		{
			name:        "OK download from osmfr",
			fNodownload: false,
			fQuiet:      false,
			fProgress:   true,
			args: args{
				myURL:    "http://download.openstreetmap.fr/extracts/africa/spain/ceuta.osm.pbf",
				fileName: "/tmp/download-geofabrik.test",
			},
			wantErr: false,
		},
	}

	for _, thisTest := range tests {
		viper.Set("noDownload", thisTest.fNodownload)
		viper.Set("quiet", thisTest.fQuiet)
		viper.Set("progress", thisTest.fProgress)
		t.Run(thisTest.name, func(t *testing.T) {
			if err := download.FromURL(thisTest.args.myURL, thisTest.args.fileName); err != nil != thisTest.wantErr {
				t.Errorf("download.FromURL() error = %v, wantErr %v", err, thisTest.wantErr)
			}
		})
	}
}

func TestFile(t *testing.T) {
	type args struct {
		configPtr *config.Config
		element   string
		format    string
		output    string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestFile",
			args: args{
				configPtr: &config.Config{
					Formats: formats.FormatDefinitions{
						formats.FormatPoly: {ID: formats.FormatPoly, Loc: ".poly", ToLoc: "", BasePath: "polygons/", BaseURL: ""},
					},
					Elements: element.MapElement{
						"africa": element.Element{ID: "africa", Name: "Africa", Formats: []string{formats.FormatPoly}},
					},
					ElementsMutex: &sync.RWMutex{},
					BaseURL:       `https://download.openstreetmap.fr/`,
				},

				element: "africa",
				format:  formats.FormatPoly,
				output:  "/tmp/download-geofabrik.test",
			},
			wantErr: false,
		},
		{
			name: "TestFile not found",
			args: args{
				configPtr: &config.Config{
					Formats: formats.FormatDefinitions{
						formats.FormatPoly: {ID: formats.FormatPoly, Loc: ".poly", ToLoc: "", BasePath: "polygons/", BaseURL: ""},
					},
					Elements: element.MapElement{
						"africa": element.Element{ID: "africa", Name: "Africa", Formats: []string{formats.FormatPoly}},
					},
					ElementsMutex: &sync.RWMutex{},
					BaseURL:       `https://download.openstreetmap.fr/`,
				},
				element: "europe",
				format:  formats.FormatPoly,
				output:  "/tmp/download-geofabrik.test",
			},
			wantErr: true,
		},
		{
			name: "TestFile url error",
			args: args{
				configPtr: &config.Config{
					Formats: formats.FormatDefinitions{
						formats.FormatPoly: {ID: formats.FormatPoly, Loc: ".poly", ToLoc: "", BasePath: "polygons/", BaseURL: ""},
					},
					Elements: element.MapElement{
						"africa": element.Element{ID: "africa", Name: "Africa", Formats: []string{formats.FormatPoly}},
					},
					ElementsMutex: &sync.RWMutex{},
					BaseURL:       `https://xx.fr/`,
				},
				element: "africa",
				format:  formats.FormatPoly,
				output:  "/tmp/download-geofabrik.test",
			},
			wantErr: true,
		},
		{
			name: "TestFile format error",
			args: args{
				configPtr: &config.Config{
					Formats: formats.FormatDefinitions{
						formats.FormatPoly: {ID: formats.FormatPoly, Loc: ".poly", ToLoc: "", BasePath: "polygons/", BaseURL: ""},
					},
					Elements: element.MapElement{
						"africa": element.Element{ID: "africa", Name: "Africa", Formats: []string{formats.FormatPoly}},
					},
					ElementsMutex: &sync.RWMutex{},
					BaseURL:       `https://download.openstreetmap.fr/`,
				},
				element: "africa",
				format:  "unknown",
				output:  "/tmp/download-geofabrik.test",
			},
			wantErr: true,
		},
	}

	viper.Set(config.ViperVerbose, true)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := download.File(tt.args.configPtr, tt.args.element, tt.args.format, tt.args.output)
			if (got != nil) != tt.wantErr {
				t.Errorf("File() error = %v, wantErr %v", got, tt.wantErr)
			}
		})
	}
}

func TestChecksum(t *testing.T) {
	// load geofabrik.yml config
	viper.Set(config.ViperConfig, "../geofabrik.yml")
	viper.Set(config.ViperElement, "monaco")

	tests := []struct {
		name   string
		format string
		check  bool
		want   bool
	}{
		{
			name:   "TestChecksum " + formats.FormatPoly + " no check",
			format: formats.FormatPoly,
			check:  false,
			want:   false,
		},
		{
			name:   "TestChecksum " + formats.FormatOsmPbf + " no check",
			format: formats.FormatOsmPbf,
			check:  false,
			want:   false,
		},
		{
			name:   "TestChecksum " + formats.FormatPoly + " check",
			format: formats.FormatPoly,
			check:  true,
			want:   false,
		},
		{
			name:   "TestChecksum " + formats.FormatOsmPbf + " check",
			format: formats.FormatOsmPbf,
			check:  true,
			want:   true,
		},
	}

	// preparation: download monaco.osm.pbf
	configPtr, err := config.LoadConfig("../geofabrik.yml")
	if err != nil {
		t.Error(err)
	}

	err = download.File(configPtr, "monaco", formats.FormatOsmPbf, "monaco.osm.pbf")
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {
		viper.Set(config.ViperCheck, tt.check)

		t.Run(tt.name, func(t *testing.T) {
			if got := download.Checksum(tt.format); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}

	os.Remove("monaco.osm.pbf")     // clean
	os.Remove("monaco.osm.pbf.md5") // clean
	os.Remove("monaco.poly")        // clean just in case
}

// Test_downloadChecksum I don't know why sometimes controlHash fail :'(
// seems geofabrik have a limit download I reach sometimes :/.
//
//nolint:paralleltest // Can't be parallelized
func Test_downloadChecksum(t *testing.T) {
	viper.Set(config.ViperVerbose, true)

	mutex := sync.RWMutex{}

	tests := []struct {
		name     string
		fConfig  string
		delement string
		format   string
		dCheck   bool
		want     bool
	}{
		// TODO: Add test cases.
		{
			name:     "dCheck = false monaco.osm.pbf from geofabrik",
			dCheck:   false,
			fConfig:  "../geofabrik.yml",
			delement: "monaco",
			format:   formats.FormatOsmPbf,
			want:     false,
		},
		{
			name:     "dCheck = true monaco.osm.pbf from geofabrik",
			fConfig:  "../geofabrik.yml",
			dCheck:   true,
			delement: "monaco",
			format:   formats.FormatOsmPbf,
			want:     true,
		},
		{
			name:    "dCheck = true monaco.poly from geofabrik",
			fConfig: "../geofabrik.yml",
			dCheck:  true, delement: "monaco",
			format: formats.FormatPoly,
			want:   false,
		},
	}

	for _, thisTest := range tests {
		mutex.Lock()
		realTestDownloadChecksum(t, thisTest)
		mutex.Unlock()
	}
}

func realTestDownloadChecksum( //nolint:thelper // prevent mutex lock
	t *testing.T,
	thisTest struct {
		name     string
		fConfig  string
		delement string
		format   string
		dCheck   bool
		want     bool
	},
) {
	// load geofabrik.yml config
	viper.Set(config.ViperConfig, thisTest.fConfig)
	viper.Set(config.ViperService, "geofabrik")
	viper.Set(config.ViperElement, thisTest.delement)
	viper.Set(config.ViperCheck, thisTest.dCheck)

	t.Run(thisTest.name, func(t *testing.T) {
		if viper.GetBool(config.ViperCheck) { // If I want to compare checksum, Download file
			configPtr, err := config.LoadConfig(thisTest.fConfig)
			if err != nil {
				t.Error(err)
			}

			myElem, err := config.FindElem(configPtr, thisTest.delement)
			if err != nil {
				t.Error(err)
			}

			myURL, err := config.Elem2URL(configPtr, myElem, thisTest.format)
			if err != nil {
				t.Error(err)
			}

			err = download.FromURL(myURL, thisTest.delement+"."+thisTest.format)
			if err != nil {
				t.Error(err)
			}
		}
		// now real test
		if got := download.Checksum(thisTest.format); got != thisTest.want {
			t.Errorf("download.Checksum() = %v, want %v", got, thisTest.want)
		}

		os.Remove("monaco.osm.pbf")     // clean
		os.Remove("monaco.osm.pbf.md5") // clean
		os.Remove("monaco.poly")        // clean
	})
}
