package app_test

import (
	"sync"
	"testing"

	"github.com/julien-noblet/download-geofabrik/app"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/spf13/viper"
)

func TestSetConfigFile(t *testing.T) {
	mutex := sync.RWMutex{}

	tests := []struct {
		name       string
		configFile string
		want       string
	}{
		{
			name:       "Test SetConfigFile with empty string",
			configFile: "",
			want:       "./geofabrik.yml",
		},
		{
			name:       "Test SetConfigFile with geofabrik.yml",
			configFile: "./geofabrik.yml",
			want:       "./geofabrik.yml",
		},
		{
			name:       "Test SetConfigFile with openstreetmap.fr.yml",
			configFile: "./openstreetmap.fr.yml",
			want:       "./openstreetmap.fr.yml",
		},
		{
			name:       "Test SetConfigFile with osmtoday.yml",
			configFile: "./osmtoday.yml",
			want:       "./osmtoday.yml",
		},
		{
			name:       "Test SetConfigFile with bbbike.yml",
			configFile: "./bbbike.yml",
			want:       "./bbbike.yml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutex.Lock()
			viper.Set(config.ViperConfig, "")

			app.SetConfigFile(tt.configFile)

			if viper.GetString(config.ViperConfig) != tt.want {
				t.Errorf("SetConfigFile() = %v, want %v", viper.GetString(config.ViperConfig), tt.want)
			}
			mutex.Unlock()
		})
	}
}

func TestCheckService(t *testing.T) {
	mutex := sync.RWMutex{}

	tests := []struct {
		name        string
		serviceName string
		want        string
		wantbool    bool
	}{
		{
			name:        "Test CheckService with geofabrik",
			serviceName: "geofabrik",
			want:        "./geofabrik.yml",
			wantbool:    true,
		},
		{
			name:        "Test CheckService with geofabrik-parse",
			serviceName: "geofabrik-parse",
			want:        "./geofabrik.yml",
			wantbool:    true,
		},
		{
			name:        "Test CheckService with openstreetmap.fr",
			serviceName: "openstreetmap.fr",
			want:        "./openstreetmap.fr.yml",
			wantbool:    true,
		},
		{
			name:        "Test CheckService with osmtoday",
			serviceName: "osmtoday",
			want:        "./osmtoday.yml",
			wantbool:    true,
		},
		{
			name:        "Test CheckService with bbbike",
			serviceName: "bbbike",
			want:        "./bbbike.yml",
			wantbool:    true,
		},
		{
			name:        "Test CheckService with unknown service",
			serviceName: "unknown",
			want:        "",
			wantbool:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutex.Lock()
			defer mutex.Unlock()

			viper.Set(config.ViperService, tt.serviceName)
			defer viper.Set(config.ViperService, "")

			if got := app.CheckService(); got != tt.wantbool {
				t.Errorf("CheckService() = %v, want %v", got, tt.wantbool)
			}

			if tt.wantbool && (viper.GetString(config.ViperConfig) != tt.want) {
				t.Errorf("CheckService() = %v, want %v", viper.GetString(config.ViperConfig), tt.want)
			}
		})
	}
}
