package main

import (
	"reflect"
	"sync"
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

func Test_geofabrikMakeParent(t *testing.T) {
	type args struct {
		e       Element
		gparent string
	}
	tests := []struct {
		name string
		args args
		want *Element
	}{
		{name: "No Parents", args: args{e: Element{ID: "a", Name: "a"}, gparent: ""}, want: nil},
		{name: "Have Parent with no gparent", args: args{e: Element{ID: "a", Name: "a", Parent: "p"}, gparent: ""}, want: &Element{ID: "p", Name: "p", Meta: true}},
		{name: "Have Parent with gparent", args: args{e: Element{ID: "a", Name: "a", Parent: "p"}, gparent: "gp"}, want: &Element{ID: "p", Name: "p", Meta: true, Parent: "gp"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := geofabrikMakeParent(tt.args.e, tt.args.gparent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("geofabrikMakeParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_geofabrikAddExtension(t *testing.T) {
	type args struct {
		id     string
		format string
		ext    Ext
	}
	tests := []struct {
		name string
		args args
		want ElementSlice
	}{
		{
			name: "Add osm.pbf but already in",
			args: args{
				id:     "a",
				format: "osm.pbf",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf"},
							Meta:    false,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf",
			args: args{
				id:     "a",
				format: "osm.pbf",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{},
							Meta:    false,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf on meta",
			args: args{
				id:     "a",
				format: "osm.pbf",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{},
							Meta:    true,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf"},
					Meta:    false,
				},
			},
		},
	}
	for tn := range tests {
		t.Run(tests[tn].name, func(t *testing.T) {
			geofabrikAddExtension(tests[tn].args.id, tests[tn].args.format, &tests[tn].args.ext)
			if !reflect.DeepEqual(tests[tn].args.ext.Elements, tests[tn].want) {
				t.Errorf("geofabrikAddExtension() got %v, want %v", tests[tn].args.ext.Elements, tests[tn].want)
			}
		})
	}
}

func Test_geofabrikParseFormat(t *testing.T) {
	type args struct {
		id     string
		format string
		ext    Ext
	}
	tests := []struct {
		name string
		args args
		want ElementSlice
	}{
		{
			name: "Add osm.pbf but already in",
			args: args{
				id:     "a",
				format: "osm.pbf",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf"},
							Meta:    false,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf", "kml", "state"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf",
			args: args{
				id:     "a",
				format: "osm.pbf",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{},
							Meta:    false,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf", "kml", "state"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf.md5",
			args: args{
				id:     "a",
				format: "osm.pbf.md5",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf", "kml", "state"},
							Meta:    false,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf", "kml", "state", "osm.pbf.md5"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.bz2",
			args: args{
				id:     "a",
				format: "osm.bz2",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf", "kml", "state"},
							Meta:    false,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf", "kml", "state", "osm.bz2"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.bz2.md5",
			args: args{
				id:     "a",
				format: "osm.bz2.md5",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf", "kml", "state"},
							Meta:    false,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf", "kml", "state", "osm.bz2.md5"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add poly",
			args: args{
				id:     "a",
				format: "poly",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf", "kml", "state"},
							Meta:    false,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf", "kml", "state", "poly"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add shp.zip",
			args: args{
				id:     "a",
				format: "shp.zip",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf", "kml", "state"},
							Meta:    false,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf", "kml", "state", "shp.zip"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add unk format",
			args: args{
				id:     "a",
				format: "unk",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{"osm.pbf", "kml", "state"},
							Meta:    false,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf", "kml", "state"},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf on meta",
			args: args{
				id:     "a",
				format: "osm.pbf",
				ext: Ext{
					Elements: ElementSlice{
						"a": Element{
							ID:      "a",
							Name:    "a",
							Formats: []string{},
							Meta:    true,
						},
					},
					ElementsMutex: sync.RWMutex{},
				},
			},
			want: ElementSlice{
				"a": Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{"osm.pbf", "kml", "state"},
					Meta:    false,
				},
			},
		},
	}
	for tn := range tests {
		t.Run(tests[tn].name, func(t *testing.T) {
			geofabrikParseFormat(tests[tn].args.id, tests[tn].args.format, &tests[tn].args.ext)
			if !reflect.DeepEqual(tests[tn].args.ext.Elements, tests[tn].want) {
				t.Errorf("geofabrikParseFormat() got %v, want %v", tests[tn].args.ext.Elements, tests[tn].want)
			}
		})
	}
}
