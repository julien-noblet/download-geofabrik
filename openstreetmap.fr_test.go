package main

import (
	"reflect"
	"testing"
)

func Test_openstreetmapFRParseHref(t *testing.T) {
	tests := []struct {
		name string
		href string
		ext  Ext
		want ElementSlice
	}{
		// TODO: Add test cases.
		{
			name: "Void ",
			href: "/",
			want: ElementSlice{},
			ext:  Ext{Elements: make(map[string]Element)},
		},
		{
			name: "1st level subfolder",
			href: "http://osm.fr/one/",
			want: ElementSlice{},
			ext:  Ext{Elements: make(map[string]Element)},
		},
		{
			name: "2nd level subfolder",
			href: "http://osm.fr/one/two/",
			want: ElementSlice{},
			ext:  Ext{Elements: make(map[string]Element)},
		},
		{
			name: "top level extracts osm.pbf",
			href: "http://osm.fr/extracts/object.osm.pbf",
			want: ElementSlice{
				"object": Element{
					ID:      "object",
					Name:    "object",
					Formats: []string{"osm.pbf"},
				},
			},
			ext: Ext{Elements: make(map[string]Element)},
		},
		{
			name: "top level polygons osm.pbf",
			href: "http://osm.fr/polygons/object.osm.pbf",
			want: ElementSlice{
				"object": Element{
					ID:      "object",
					Name:    "object",
					Formats: []string{"osm.pbf"},
				},
			},
			ext: Ext{Elements: make(map[string]Element)},
		},
		{
			name: "sub level extracts osm.pbf",
			href: "http://osm.fr/extracts/one/object.osm.pbf",
			want: ElementSlice{
				"object": Element{
					ID:      "object",
					Name:    "object",
					Formats: []string{"osm.pbf"},
					Parent:  "one",
				},
			},
			ext: Ext{Elements: make(map[string]Element)},
		},
		{
			name: "sub level polygons osm.pbf",
			href: "http://osm.fr/polygons/one/object.osm.pbf",
			want: ElementSlice{
				"object": Element{
					ID:      "object",
					Name:    "object",
					Formats: []string{"osm.pbf"},
					Parent:  "one",
				},
			},
			ext: Ext{Elements: make(map[string]Element)},
		},
		{
			name: "sub level extracts osm.pbf + state",
			href: "http://osm.fr/extracts/one/object.state.txt",
			want: ElementSlice{
				"object": Element{
					ID:      "object",
					Name:    "object",
					Formats: []string{"osm.pbf", "state"},
					Parent:  "one",
				},
			},
			ext: Ext{
				Elements: ElementSlice{
					"object": Element{
						ID:      "object",
						Name:    "object",
						Formats: []string{"osm.pbf"},
						Parent:  "one",
					},
				},
			},
		},
	}
	for tn := range tests {
		t.Run(tests[tn].name, func(t *testing.T) {
			openstreetmapFRParseHref(tests[tn].href, &tests[tn].ext)
			if !reflect.DeepEqual(tests[tn].ext.Elements, tests[tn].want) {
				t.Errorf("openstreetmapFRParseHref() = %v len:%d, want %v len:%d", tests[tn].ext.Elements, len(tests[tn].ext.Elements), tests[tn].want, len(tests[tn].want))
			}

		})
	}
}
