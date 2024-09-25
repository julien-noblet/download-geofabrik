package app_test

import (
	"testing"

	"github.com/apex/log"
	"github.com/julien-noblet/download-geofabrik/app"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/spf13/viper"
)

func TestConfigureLogging(t *testing.T) {
	type args struct {
		quiet    bool
		verbose  bool
		progress bool
	}

	tests := []struct {
		name string
		args args
		want log.Level
	}{
		{
			name: "Test ConfigureLogging with quiet mode",
			args: args{
				quiet: true,
			},
			want: log.ErrorLevel,
		},
		{
			name: "Test ConfigureLogging with verbose mode",
			args: args{
				verbose: true,
			},
			want: log.DebugLevel,
		},
		{
			name: "Test ConfigureLogging with progress mode",
			args: args{
				progress: true,
			},
			want: log.ErrorLevel,
		},
		{
			name: "Test ConfigureLogging with quiet and verbose mode",
			args: args{
				quiet:   true,
				verbose: true,
			},
			want: log.DebugLevel,
		},
		{
			name: "Test ConfigureLogging with quiet and progress mode",
			args: args{
				quiet:    true,
				progress: true,
			},
			want: log.ErrorLevel,
		},
		{
			name: "Test ConfigureLogging with verbose and progress mode",
			args: args{
				verbose:  true,
				progress: true,
			},
			want: log.DebugLevel,
		},
		{
			name: "Test ConfigureLogging with quiet, verbose and progress mode",
			args: args{
				quiet:    true,
				verbose:  true,
				progress: true,
			},
			want: log.DebugLevel,
		},
		{
			name: "Test ConfigureLogging with no mode",
			args: args{},
			want: log.InfoLevel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			viper.Set(config.ViperQuiet, tt.args.quiet)
			viper.Set(config.ViperVerbose, tt.args.verbose)
			viper.Set(config.ViperProgress, tt.args.progress)
			app.ConfigureLogging()
		})
	}
}
