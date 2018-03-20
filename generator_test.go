package main

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestElementSlice_Generate(t *testing.T) {
	myConfig := &SampleConfigValidPtr
	myYaml, _ := yaml.Marshal(*myConfig)
	myConfig.Elements = *new(map[string]Element) // void Elements

	type args struct {
		myConfig *Config
	}
	tests := []struct {
		name    string
		e       ElementSlice
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Marshalling OK, no error",
			e:       sampleElementValidPtr,
			args:    args{myConfig: myConfig},
			want:    myYaml,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.Generate(tt.args.myConfig)
			if (err != nil) != tt.wantErr {
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
		t.Run(tt.name, func(t *testing.T) {
			Generate(tt.args.configfile)
		})
	}
}
