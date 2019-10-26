package main

import "testing"

func Test_downloadFromURL(t *testing.T) {
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
		*fNodownload = tt.fNodownload
		*fQuiet = tt.fQuiet
		*fProgress = tt.fProgress
		t.Run(tt.name, func(t *testing.T) {
			if err := downloadFromURL(tt.args.myURL, tt.args.fileName); err != nil != tt.wantErr {
				t.Errorf("downloadFromURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
