package main

import (
	"reflect"
	"sync"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestElementSlice_Generate(t *testing.T) {
	tests := []struct {
		name    string
		e       ElementSlice
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
		myConfig.Elements = map[string]Element{} // void Elements
		myConfig.Elements = tt.e
		tt.want, _ = yaml.Marshal(*myConfig)
		t.Run(tt.name, func(t *testing.T) {
			got, err := myConfig.Generate()
			if err != nil != tt.wantErr {
				t.Errorf("ElementSlice.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElementSlice.Generate() = %v, want %v", got, tt.want)
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
			Generate(tt.args.configfile)
		})
	}
}

func TestConfig_mergeElement(t *testing.T) {
	type args struct {
		element *Element
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
				element: &Element{ID: "us", Parent: "north-america", Formats: []string{formatOsmPbf}},
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
			c := Config{
				Elements:      sampleElementValidPtr,
				ElementsMutex: &sync.RWMutex{},
			}
			if err := c.mergeElement(tt.args.element); err != nil != tt.wantErr {
				t.Errorf("Ext.mergeElement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
