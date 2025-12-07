package geofabrik_test

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/generator/importer/geofabrik"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetIndex(t *testing.T) {
	viper.Set("log", true)

	tests := []struct {
		name    string
		myURL   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test",
			myURL:   geofabrik.GeofabrikIndexURL,
			wantErr: false,
		},
		{
			name:    "Test 404",
			myURL:   "https://google.com/404",
			wantErr: true,
		},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			index, err := geofabrik.GetIndex(thisTest.myURL)
			if thisTest.wantErr {
				require.Error(t, err)
			} else {
				assert.NotNil(t, index)

				if len(index.Features) < 10 {
					t.Errorf("GetIndex() error I should have more features!!!")
				}

				converted, err := geofabrik.Convert(index)
				if converted == nil || err != nil {
					t.Errorf("GetIndex() error cant convert !!!\n%v", err)
				}

				if e, err := converted.GetElement("france"); err != nil || e == nil {
					t.Errorf("GetIndex() error cant find element !!!\nconfig=%v\nerr=%v", converted, err)
				}
			}
		})
	}
}
