package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name    string
	href    string
	ext     Ext
	want    ElementSlice
	parent  string
	parents []string
}{
	// TODO: Add test cases.
	{
		name:    "Void ",
		href:    "http://osm.fr/",
		want:    ElementSlice{},
		ext:     Ext{Elements: make(map[string]Element)},
		parent:  "",
		parents: []string{"http:", "", "osm.fr", ""},
	},
	{
		name:    "1st level subfolder",
		href:    "http://osm.fr/one/",
		want:    ElementSlice{},
		ext:     Ext{Elements: make(map[string]Element)},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "one", ""},
	},
	{
		name:    "2nd level subfolder",
		href:    "http://osm.fr/one/two/",
		want:    ElementSlice{},
		ext:     Ext{Elements: make(map[string]Element)},
		parent:  "two",
		parents: []string{"http:", "", "osm.fr", "one", "two", ""},
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
		ext:     Ext{Elements: make(map[string]Element)},
		parent:  "",
		parents: []string{"http:", "", "osm.fr", "extracts", "object.osm.pbf"},
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
		ext:     Ext{Elements: make(map[string]Element)},
		parent:  "",
		parents: []string{"http:", "", "osm.fr", "polygons", "object.osm.pbf"},
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
		ext:     Ext{Elements: make(map[string]Element)},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "extracts", "one", "object.osm.pbf"},
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
		ext:     Ext{Elements: make(map[string]Element)},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "polygons", "one", "object.osm.pbf"},
	},
	{
		name: "2nd level osm.pbf",
		href: "http://osm.fr/polygons/one/two/object.osm.pbf",
		want: ElementSlice{
			"object": Element{
				ID:      "object",
				Name:    "object",
				Formats: []string{"osm.pbf"},
				Parent:  "two",
			},
		},
		ext:     Ext{Elements: make(map[string]Element)},
		parent:  "two",
		parents: []string{"http:", "", "osm.fr", "polygons", "one", "two", "object.osm.pbf"},
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
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "extracts", "one", "object.state.txt"},
	},
}

func Test_openstreetmapFRParseHref(t *testing.T) {
	for tn := range tests {
		t.Run(tests[tn].name, func(t *testing.T) {
			openstreetmapFRParseHref(tests[tn].href, &tests[tn].ext)
			if !reflect.DeepEqual(tests[tn].ext.Elements, tests[tn].want) {
				t.Errorf("openstreetmapFRParseHref() = %v len:%d, want %v len:%d", tests[tn].ext.Elements, len(tests[tn].ext.Elements), tests[tn].want, len(tests[tn].want))
			}

		})
	}
}

func Test_openstreetmapFRGetParent(t *testing.T) {
	for tn := range tests {
		t.Run(tests[tn].name, func(t *testing.T) {
			p, ps := openstreetmapFRGetParent(tests[tn].href)
			if !reflect.DeepEqual(p, tests[tn].parent) {
				t.Errorf("openstreetmapFRGetParent() = %v want %v ", p, tests[tn].parent)
			}
			if !reflect.DeepEqual(ps, tests[tn].parents) {
				t.Errorf("openstreetmapFRGetParent() = %v want %v ", ps, tests[tn].parents)
			}

		})
	}
}
