package download_test

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/download"
	"github.com/spf13/viper"
)

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

	for _, tt := range tests {
		tt := tt
		viper.Set("noDownload", tt.fNodownload)
		viper.Set("quiet", tt.fQuiet)
		viper.Set("progress", tt.fProgress)
		t.Run(tt.name, func(t *testing.T) {
			if err := download.FromURL(tt.args.myURL, tt.args.fileName); err != nil != tt.wantErr {
				t.Errorf("download.FromURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
