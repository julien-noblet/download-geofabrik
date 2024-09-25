package download_test

import (
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
		thisTest := thisTest
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
	t.Parallel()

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
					Elements: element.Slice{
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
					Elements: element.Slice{
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
					Elements: element.Slice{
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
					Elements: element.Slice{
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
		t.Parallel()

		t.Run(tt.name, func(t *testing.T) {
			got := download.File(tt.args.configPtr, tt.args.element, tt.args.format, tt.args.output)
			if (got != nil) != tt.wantErr {
				t.Errorf("File() error = %v, wantErr %v", got, tt.wantErr)
			}
		})
	}
}
