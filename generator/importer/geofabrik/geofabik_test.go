package geofabrik_test

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/generator/importer/geofabrik"
	"github.com/spf13/viper"
)

func TestGetIndex(t *testing.T) {
	t.Parallel()
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
	}
	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			index, err := geofabrik.GetIndex(thisTest.myURL)
			if (err != nil) != thisTest.wantErr {
				t.Errorf("GetIndex() error = %v, wantErr %v", err, thisTest.wantErr)
			}
			if len(index.Features) < 10 {
				t.Errorf("GetIndex() error I should have more features!!!")
			}
			c, err := geofabrik.Convert(index)
			if c == nil || err != nil {
				t.Errorf("GetIndex() error cant convert !!!\n%v", err)
			}
			if e, err := c.GetElement("france"); err != nil || e == nil {
				t.Errorf("GetIndex() error cant find element !!!\nconfig=%v\nerr=%v", c, err)
			}
		})
	}
}
