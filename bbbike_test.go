package main

import "testing"

func Test_bbbikeGetName(t *testing.T) {

	tests := []struct {
		name string
		h3   string
		want string
	}{
		// TODO: Add test cases.
		{name: "Valid", h3: "OSM extracts for Test", want: "Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bbbikeGetName(tt.h3); got != tt.want {
				t.Errorf("bbbikeGetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
