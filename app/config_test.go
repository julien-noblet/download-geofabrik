package app_test

import (
	"os"
	"sync"
	"testing"

	"github.com/julien-noblet/download-geofabrik/app"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/spf13/viper"
)

func Test_setViperBool(t *testing.T) {
	mutex := sync.RWMutex{}

	type args struct {
		name string
		flag bool
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test setViperBool with true",
			args: args{
				name: "test",
				flag: true,
			},
			want: true,
		},
		{
			name: "Test setViperBool with false",
			args: args{
				name: "test",
				flag: false,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutex.Lock()
			app.SetViperBool(tt.args.name, &tt.args.flag)

			if got := viper.Get(tt.args.name); got != tt.want {
				t.Errorf("setViperBool() = %v, want %v", got, tt.want)
			}

			mutex.Unlock()
		})
	}
}

func TestSetOutputDir(t *testing.T) {
	t.Parallel()

	currdir, _ := os.Getwd()

	tests := []struct {
		name            string
		outputDirectory string
		env             string
		want            string
	}{
		{
			name:            "Test SetOutputDir with empty output directory",
			outputDirectory: "",
			env:             "",
			want:            currdir + "/",
		},
		{
			name:            "Test SetOutputDir with output directory",
			outputDirectory: "/tmp",
			env:             "",
			want:            "/tmp/",
		},
		{
			name:            "Test SetOutputDir with empty output directory but env variable",
			env:             "/tmp",
			want:            "/tmp/",
			outputDirectory: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.env != "" {
				t.Setenv("OUTPUT_DIR", tt.env)
			}

			viper.Set("output_directory", tt.outputDirectory)

			app.SetOutputDir()

			if got := viper.GetString("output_directory"); got != tt.want {
				t.Errorf("SetOutputDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_ConfigureViper(t *testing.T) {
	type fields struct {
		fConfig         string
		fNodownload     bool
		fVerbose        bool
		fQuiet          bool
		fProgress       bool
		delement        string
		dCheck          bool
		dOutputDir      string
		dosmBz2         bool
		dosmGz          bool
		dshpZip         bool
		dosmPbf         bool
		doshPbf         bool
		dstate          bool
		dpoly           bool
		dkml            bool
		dgeojson        bool
		dgarmin         bool
		dmaps           bool
		dmbtiles        bool
		dcsv            bool
		dgarminonroad   bool
		dgarminontrail  bool
		dgarminopentopo bool
		lmd             bool
		fService        string
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test ConfigureViper with all flags set to true",
			fields: fields{
				fConfig:         "test",
				fNodownload:     true,
				fVerbose:        true,
				fQuiet:          true,
				fProgress:       true,
				delement:        "test",
				dCheck:          true,
				dOutputDir:      "test",
				dosmBz2:         true,
				dosmGz:          true,
				dshpZip:         true,
				dosmPbf:         true,
				doshPbf:         true,
				dstate:          true,
				dpoly:           true,
				dkml:            true,
				dgeojson:        true,
				dgarmin:         true,
				dmaps:           true,
				dmbtiles:        true,
				dcsv:            true,
				dgarminonroad:   true,
				dgarminontrail:  true,
				dgarminopentopo: true,
				lmd:             true,
				fService:        "test",
			},
		},
		{
			name: "Test ConfigureViper with all flags set to false",
			fields: fields{
				fConfig:         "test",
				fNodownload:     false,
				fVerbose:        false,
				fQuiet:          false,
				fProgress:       false,
				delement:        "test",
				dCheck:          false,
				dOutputDir:      "test",
				dosmBz2:         false,
				dosmGz:          false,
				dshpZip:         false,
				dosmPbf:         false,
				doshPbf:         false,
				dstate:          false,
				dpoly:           false,
				dkml:            false,
				dgeojson:        false,
				dgarmin:         false,
				dmaps:           false,
				dmbtiles:        false,
				dcsv:            false,
				dgarminonroad:   false,
				dgarminontrail:  false,
				dgarminopentopo: false,
				lmd:             false,
				fService:        "test",
			},
		},
		{
			name: "Test ConfigureViper with all flags set to false except fNodownload",
			fields: fields{
				fConfig:         "test",
				fNodownload:     true,
				fVerbose:        false,
				fQuiet:          false,
				fProgress:       false,
				delement:        "test",
				dCheck:          false,
				dOutputDir:      "test",
				dosmBz2:         false,
				dosmGz:          false,
				dshpZip:         false,
				dosmPbf:         false,
				doshPbf:         false,
				dstate:          false,
				dpoly:           false,
				dkml:            false,
				dgeojson:        false,
				dgarmin:         false,
				dmaps:           false,
				dmbtiles:        false,
				dcsv:            false,
				dgarminonroad:   false,
				dgarminontrail:  false,
				dgarminopentopo: false,
				lmd:             false,
				fService:        "test",
			},
		},
		{
			name: "Test ConfigureViper with different settings for string flags",
			fields: fields{
				fConfig:         "fConfig",
				fNodownload:     true,
				fVerbose:        false,
				fQuiet:          false,
				fProgress:       false,
				delement:        "delement",
				dCheck:          false,
				dOutputDir:      "dOutputDir",
				dosmBz2:         false,
				dosmGz:          false,
				dshpZip:         false,
				dosmPbf:         false,
				doshPbf:         false,
				dstate:          false,
				dpoly:           false,
				dkml:            false,
				dgeojson:        false,
				dgarmin:         false,
				dmaps:           false,
				dmbtiles:        false,
				dcsv:            false,
				dgarminonroad:   false,
				dgarminontrail:  false,
				dgarminopentopo: false,
				lmd:             false,
				fService:        "fService",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &app.App{
				FConfig:         &tt.fields.fConfig,
				FNodownload:     &tt.fields.fNodownload,
				FVerbose:        &tt.fields.fVerbose,
				FQuiet:          &tt.fields.fQuiet,
				FProgress:       &tt.fields.fProgress,
				Delement:        &tt.fields.delement,
				DCheck:          &tt.fields.dCheck,
				DOutputDir:      &tt.fields.dOutputDir,
				DosmBz2:         &tt.fields.dosmBz2,
				DosmGz:          &tt.fields.dosmGz,
				DshpZip:         &tt.fields.dshpZip,
				DosmPbf:         &tt.fields.dosmPbf,
				DoshPbf:         &tt.fields.doshPbf,
				Dstate:          &tt.fields.dstate,
				Dpoly:           &tt.fields.dpoly,
				Dkml:            &tt.fields.dkml,
				Dgeojson:        &tt.fields.dgeojson,
				Dgarmin:         &tt.fields.dgarmin,
				Dmaps:           &tt.fields.dmaps,
				Dmbtiles:        &tt.fields.dmbtiles,
				Dcsv:            &tt.fields.dcsv,
				Dgarminonroad:   &tt.fields.dgarminonroad,
				Dgarminontrail:  &tt.fields.dgarminontrail,
				Dgarminopentopo: &tt.fields.dgarminopentopo,
				Lmd:             &tt.fields.lmd,
				FService:        &tt.fields.fService,
			}

			a.ConfigureViper()

			checks := []struct {
				key      string
				expected interface{}
				actual   interface{}
			}{
				{config.ViperConfig, tt.fields.fConfig, viper.GetString(config.ViperConfig)},
				{config.ViperNoDL, tt.fields.fNodownload, viper.GetBool(config.ViperNoDL)},
				{config.ViperVerbose, tt.fields.fVerbose, viper.GetBool(config.ViperVerbose)},
				{config.ViperQuiet, tt.fields.fQuiet, viper.GetBool(config.ViperQuiet)},
				{config.ViperProgress, tt.fields.fProgress, viper.GetBool(config.ViperProgress)},
				{config.ViperElement, tt.fields.delement, viper.GetString(config.ViperElement)},
				{config.ViperCheck, tt.fields.dCheck, viper.GetBool(config.ViperCheck)},
				{config.ViperOutputDirectory, tt.fields.dOutputDir, viper.GetString(config.ViperOutputDirectory)},
				{config.ViperListFormatMarkdown, tt.fields.lmd, viper.GetBool(config.ViperListFormatMarkdown)},
				{formats.FormatOsmBz2, tt.fields.dosmBz2, viper.GetBool(formats.FormatOsmBz2)},
				{formats.FormatOsmGz, tt.fields.dosmGz, viper.GetBool(formats.FormatOsmGz)},
				{formats.FormatShpZip, tt.fields.dshpZip, viper.GetBool(formats.FormatShpZip)},
				{formats.FormatOsmPbf, tt.fields.dosmPbf, viper.GetBool(formats.FormatOsmPbf)},
				{formats.FormatOshPbf, tt.fields.doshPbf, viper.GetBool(formats.FormatOshPbf)},
				{formats.FormatState, tt.fields.dstate, viper.GetBool(formats.FormatState)},
				{formats.FormatPoly, tt.fields.dpoly, viper.GetBool(formats.FormatPoly)},
				{formats.FormatKml, tt.fields.dkml, viper.GetBool(formats.FormatKml)},
				{formats.FormatGeoJSON, tt.fields.dgeojson, viper.GetBool(formats.FormatGeoJSON)},
				{formats.FormatGarminOSM, tt.fields.dgarmin, viper.GetBool(formats.FormatGarminOSM)},
				{formats.FormatMapsforge, tt.fields.dmaps, viper.GetBool(formats.FormatMapsforge)},
				{formats.FormatMBTiles, tt.fields.dmbtiles, viper.GetBool(formats.FormatMBTiles)},
				{formats.FormatCSV, tt.fields.dcsv, viper.GetBool(formats.FormatCSV)},
				{formats.FormatGarminOnroad, tt.fields.dgarminonroad, viper.GetBool(formats.FormatGarminOnroad)},
				{formats.FormatGarminOntrail, tt.fields.dgarminontrail, viper.GetBool(formats.FormatGarminOntrail)},
				{formats.FormatGarminOpenTopo, tt.fields.dgarminopentopo, viper.GetBool(formats.FormatGarminOpenTopo)},
				{config.ViperService, tt.fields.fService, viper.GetString(config.ViperService)},
			}

			for _, check := range checks {
				if check.actual != check.expected {
					t.Errorf("App.ConfigureViper() %s = %v, want %v", check.key, check.actual, check.expected)
				}
			}
		})
	}
}
