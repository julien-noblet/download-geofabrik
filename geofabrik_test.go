package main

import (
	"testing"
)

func Test_geofabrikGetParent(t *testing.T) {

	tests := []struct {
		name  string
		url   string
		want  string
		want2 string
	}{
		// TODO: Add test cases.
		{name: "No Parent", url: "https://download.geofabrik.de/test.html", want: "", want2: "https://download.geofabrik.de"},
		{name: "1 parent", url: "https://download.geofabrik.de/parent1/test.html", want: "parent1", want2: "https://download.geofabrik.de/parent1"},
		{name: "2 parents", url: "https://download.geofabrik.de/parent1/parent2/test.html", want: "parent2", want2: "https://download.geofabrik.de/parent1/parent2"},
		{name: "grand parents", url: "https://download.geofabrik.de/parent1/parent2", want: "parent1", want2: "https://download.geofabrik.de/parent1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := geofabrikGetParent(tt.url)
			if got != tt.want {
				t.Errorf("geofabrikGetParent() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("geofabrikGetParent() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_geofabrikFileWOExt(t *testing.T) {
	tests := []struct {
		name  string
		url   string
		want  string
		want2 string
	}{
		// TODO: Add test cases.
		{name: "No Parent", url: "https://download.geofabrik.de/test.html", want: "test", want2: "html"},
		{name: "No Parent long ext", url: "https://download.geofabrik.de/test.ext.html", want: "test", want2: "ext.html"},
		{name: "1 Parent", url: "https://download.geofabrik.de/parent/test.html", want: "test", want2: "html"},
		{name: "1 Parent long ext", url: "https://download.geofabrik.de/parent/test.ext.html", want: "test", want2: "ext.html"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ext := geofabrikFile(tt.url)
			if got != tt.want {
				t.Errorf("geofabrikFileWOExt() = %v, want %v", got, tt.want)
			}
			if ext != tt.want2 {
				t.Errorf("geofabrikFileWOExt() = %v, want %v", ext, tt.want2)
			}
		})
	}
}
