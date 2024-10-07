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
			mutex.Unlock()

			mutex.Lock()
			app.SetConfigFile(tt.configFile)
			mutex.Unlock()

			mutex.RLock()
			if viper.GetString(config.ViperConfig) != tt.want {
				t.Errorf("SetConfigFile() = %v, want %v", viper.GetString(config.ViperConfig), tt.want)
			}
			mutex.RUnlock()
		})
	}
}

func TestCheckService(t *testing.T) {
	mutex := sync.RWMutex{}

	tests := []struct {
		name        string
		serviceName string
		wantbool    bool
		want        string
	}{
		{
			name:        "Test CheckService with geofabrik",
			serviceName: "geofabrik",
			wantbool:    true,
			want:        "./geofabrik.yml",
		},
		{
			name:        "Test CheckService with geofabrik-parse",
			serviceName: "geofabrik-parse",
			wantbool:    true,
			want:        "./geofabrik.yml",
		},
		{
			name:        "Test CheckService with openstreetmap.fr",
			serviceName: "openstreetmap.fr",
			wantbool:    true,
			want:        "./openstreetmap.fr.yml",
		},
		{
			name:        "Test CheckService with osmtoday",
			serviceName: "osmtoday",
			wantbool:    true,
			want:        "./osmtoday.yml",
		},
		{
			name:        "Test CheckService with bbbike",
			serviceName: "bbbike",
			wantbool:    true,
			want:        "./bbbike.yml",
		},
		{
			name:        "Test CheckService with unknown service",
			serviceName: "unknown",
			wantbool:    false,
			want:        "",
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
