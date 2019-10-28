package generator_test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/generator"
	yaml "gopkg.in/yaml.v2"
)

var sampleAfricaElementPtr = element.Element{
	ID:   "africa",
	Name: "Africa",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
}
var sampleGeorgiaUsElementPtr = element.Element{
	ID:   "georgia-us",
	File: "georgia",
	Name: "Georgia (US State)",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatShpZip,
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
	Parent: "us",
}
var sampleUsElementPtr = element.Element{
	ID:     "us",
	Meta:   true,
	Name:   "United States of America",
	Parent: "north-america",
}
var sampleNorthAmericaElementPtr = element.Element{
	ID:   "north-america",
	Name: "North America",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
}
var sampleElementValidPtr = map[string]element.Element{
	"africa":        sampleAfricaElementPtr,
	"georgia-us":    sampleGeorgiaUsElementPtr,
	"us":            sampleUsElementPtr,
	"north-america": sampleNorthAmericaElementPtr,
}
var sampleFormatValidPtr = map[string]formats.Format{
	// Blank
	"": {
		ID:       "",
		Loc:      "",
		BasePath: "",
	}, formats.FormatOsmPbf: {
		ID:  formats.FormatOsmPbf,
		Loc: ".osm.pbf",
		//BasePath: "/",
	}, formats.FormatState: {
		ID:       formats.FormatState,
		Loc:      "-updates/state.txt",
		BasePath: "../state/",
	}, formats.FormatPoly: {
		ID:      formats.FormatPoly,
		Loc:     ".poly",
		BaseURL: "http://my.new.url/folder",
	}, formats.FormatOsmBz2: {
		ID:       formats.FormatOsmBz2,
		Loc:      ".osm.bz2",
		BasePath: "../osmbz2/",
		BaseURL:  "http://my.new.url/folder",
	}, formats.FormatOsmGz: {
		ID:       formats.FormatOsmGz,
		Loc:      ".osm.gz",
		BasePath: "../osmgz/",
		BaseURL:  "http://my.new.url/folder",
	},
}

var SampleConfigValidPtr = config.Config{
	BaseURL:  "https://my.base.url",
	Formats:  sampleFormatValidPtr,
	Elements: sampleElementValidPtr,
}
var sampleFakeGeorgiaPtr = element.Element{
	ID:   "georgia-usf",
	File: "georgia-fake",
	Name: "Georgia (US State) - fake test",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatShpZip,
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
	Parent: "us", // keep good parent!
}
var sampleFakeGeorgia2Ptr = element.Element{
	ID:   "georgia-us2",
	File: "georgia",
	Name: "Georgia (US State)",
	Formats: []string{
		formats.FormatOsmPbf,
		"osm.pbf.md5",
		formats.FormatShpZip,
		formats.FormatOsmBz2,
		"osm.bz2.md5",
		formats.FormatOshPbf,
		"osh.pbf.md5",
		formats.FormatPoly,
		formats.FormatKml,
		formats.FormatState,
	},
	Parent: "notus", // bad parent not exist!
}

func TestSlice_Generate(t *testing.T) {
	tests := []struct {
		name    string
		e       element.Slice
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Marshaling OK, no error",
			e:       sampleElementValidPtr,
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		myConfig := &SampleConfigValidPtr
		myConfig.Elements = map[string]element.Element{} // void Elements
		myConfig.Elements = tt.e
		tt.want, _ = yaml.Marshal(*myConfig)
		t.Run(tt.name, func(t *testing.T) {
			got, err := myConfig.Generate()
			if err != nil != tt.wantErr {
				t.Errorf("Slice.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		configfile string
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			generator.Generate(tt.args.configfile)
		})
	}
}

func TestConfig_mergeElement(t *testing.T) {
	type args struct {
		element *element.Element
	}

	myFakeGeorgia := sampleFakeGeorgia2Ptr
	myFakeGeorgia.ID = "georgia-us"
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "adding a new element",
			args: args{
				element: &sampleFakeGeorgiaPtr,
			},
			wantErr: false,
		},
		{
			name: "adding a same element",
			args: args{
				element: &sampleAfricaElementPtr,
			},
			wantErr: false,
		},
		{
			name: "adding a meta element",
			args: args{
				element: &sampleUsElementPtr,
			},
			wantErr: false,
		},
		{
			name: "adding a format to a meta element",
			args: args{
				element: &element.Element{ID: "us", Parent: "north-america", Formats: []string{formats.FormatOsmPbf}},
			},
			wantErr: false,
		},
		{
			name: "adding a same element parent differ",
			args: args{
				element: &myFakeGeorgia,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := config.Config{
				Elements:      sampleElementValidPtr,
				ElementsMutex: &sync.RWMutex{},
			}
			if err := c.MergeElement(tt.args.element); err != nil != tt.wantErr {
				t.Errorf("Ext.mergeElement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
