package download_test

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/download"
	"github.com/spf13/viper"
)

func Test_DownloadFromURL(t *testing.T) {
	t.Parallel()

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
		{name: "try fNodownload=true", fNodownload: true, wantErr: false},
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
				myURL:    "http://download.openstreetmap.fr/extracts/europe/spain/canarias.osm.pbf",
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
			t.Parallel()
			if err := download.FromURL(thisTest.args.myURL, thisTest.args.fileName); err != nil != thisTest.wantErr {
				t.Errorf("download.FromURL() error = %v, wantErr %v", err, thisTest.wantErr)
			}
		})
	}
}
