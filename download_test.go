package main

import "testing"

func Test_downloadFromURL(t *testing.T) {
	type args struct {
		myURL    string
		fileName string
	}
	tests := []struct {
		name        string
		fNodownload bool
		fQuiet      bool
		fProgress   bool
		args        args
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
				myURL:    "https://download.geofabrik.de/europe/andorra.osh.pbf.md5",
				fileName: "/tmp/download-geofabrik.test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		*fNodownload = tt.fNodownload
		*fQuiet = tt.fQuiet
		*fProgress = tt.fProgress
		t.Run(tt.name, func(t *testing.T) {
			if err := downloadFromURL(tt.args.myURL, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("downloadFromURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
