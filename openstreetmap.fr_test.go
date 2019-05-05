package main

import (
	"reflect"
	"sync"
	"testing"
)

var openstreetmapFRtests = []struct {
	name    string
	href    string
	c       Config
	want    ElementSlice
	parent  string
	parents []string
}{
	// TODO: Add test cases.
	{
		name:    "Void ",
		href:    "http://osm.fr/",
		want:    ElementSlice{},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "",
		parents: []string{"http:", "", "osm.fr", ""},
	},
	{
		name:    "1st level subfolder",
		href:    "http://osm.fr/one/",
		want:    ElementSlice{},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "one", ""},
	},
	{
		name:    "2nd level subfolder",
		href:    "http://osm.fr/one/two/",
		want:    ElementSlice{},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
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
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
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
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
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
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
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
				Meta:    false,
				Formats: []string{"osm.pbf"},
				Parent:  "one",
			},
		},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
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
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
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
		c: Config{
			Elements: ElementSlice{
				"object": Element{
					ID:      "object",
					Name:    "object",
					Formats: []string{"osm.pbf"},
					Parent:  "one",
				},
			},
			ElementsMutex: &sync.RWMutex{},
		},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "extracts", "one", "object.state.txt"},
	},
}

func Test_openstreetmapFRParseHref(t *testing.T) {
	for tn := range openstreetmapFRtests {
		t.Run(openstreetmapFRtests[tn].name, func(t *testing.T) {
			openstreetmapFRParseHref(openstreetmapFRtests[tn].href, &openstreetmapFRtests[tn].c)
			if !reflect.DeepEqual(openstreetmapFRtests[tn].c.Elements, openstreetmapFRtests[tn].want) {
				t.Errorf("openstreetmapFRParseHref() = %v len:%d, want %v len:%d", openstreetmapFRtests[tn].c.Elements, len(openstreetmapFRtests[tn].c.Elements), openstreetmapFRtests[tn].want, len(openstreetmapFRtests[tn].want))
			}
		})
	}
}

func Test_openstreetmapFRGetParent(t *testing.T) {
	for tn := range openstreetmapFRtests {
		t.Run(openstreetmapFRtests[tn].name, func(t *testing.T) {
			p, ps := openstreetmapFRGetParent(openstreetmapFRtests[tn].href)
			if !reflect.DeepEqual(p, openstreetmapFRtests[tn].parent) {
				t.Errorf("openstreetmapFRGetParent() = %v want %v ", p, openstreetmapFRtests[tn].parent)
			}
			if !reflect.DeepEqual(ps, openstreetmapFRtests[tn].parents) {
				t.Errorf("openstreetmapFRGetParent() = %v want %v ", ps, openstreetmapFRtests[tn].parents)
			}

		})
	}
}
