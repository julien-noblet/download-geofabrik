package generator_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/generator"
	yaml "gopkg.in/yaml.v3"
)

func sampleAfricaElementPtr() *element.Element {
	return &element.Element{
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
}

func sampleGeorgiaUsElementPtr() *element.Element {
	return &element.Element{
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
}

func sampleUsElementPtr() *element.Element {
	return &element.Element{
		ID:     "us",
		Meta:   true,
		Name:   "United States of America",
		Parent: "north-america",
	}
}

func sampleNorthAmericaElementPtr() *element.Element {
	return &element.Element{
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
}

func sampleElementValidPtr() map[string]element.Element {
	return map[string]element.Element{
		"africa":        *sampleAfricaElementPtr(),
		"georgia-us":    *sampleGeorgiaUsElementPtr(),
		"us":            *sampleUsElementPtr(),
		"north-america": *sampleNorthAmericaElementPtr(),
	}
}

func sampleFormatValidPtr() map[string]formats.Format {
	return map[string]formats.Format{
		// Blank
		"": {
			ID:       "",
			Loc:      "",
			BasePath: "",
		}, formats.FormatOsmPbf: {
			ID:  formats.FormatOsmPbf,
			Loc: ".osm.pbf",
			// BasePath: "/",
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
}

func SampleConfigValidPtr() config.Config {
	return config.Config{
		BaseURL:  "https://my.base.url",
		Formats:  sampleFormatValidPtr(),
		Elements: sampleElementValidPtr(),
	}
}

func TestSlice_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		e       element.Slice
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Marshaling OK, no error",
			e:       sampleElementValidPtr(),
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, thisTest := range tests {
		thisTest := thisTest
		myConfig := SampleConfigValidPtr()
		myConfig.Elements = map[string]element.Element{} // void Elements
		myConfig.Elements = thisTest.e
		thisTest.want, _ = yaml.Marshal(myConfig)
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := myConfig.Generate()
			if err != nil != thisTest.wantErr {
				t.Errorf("Slice.Generate() error = %v, wantErr %v", err, thisTest.wantErr)

				return
			}

			if !reflect.DeepEqual(got, thisTest.want) {
				t.Errorf("Slice.Generate() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	t.Parallel()

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
			t.Parallel()
			generator.Generate(tt.args.configfile)
		})
	}
}

func Test_write(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		output string
	}{
		// TODO: Add test cases.
		{name: "geofabrik", input: "../geofabrik.yml", output: "/tmp/test.yml"},
	}
	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			c, _ := config.LoadConfig(thisTest.input)
			generator.Write(c, thisTest.output)

			input, err := os.ReadFile(thisTest.input)
			if err != nil {
				t.Errorf("read() error = %v", err)
			}

			output, err := os.ReadFile(thisTest.output)
			if err != nil {
				t.Errorf("read() error = %v", err)
			}

			reflect.DeepEqual(input, output)
		})
	}
}

func TestCleanup(t *testing.T) {
	t.Parallel()

	type args struct {
		c *config.Config
	}

	tests := []struct {
		name string
		args args
		want element.Formats
	}{
		// TODO: Add test cases.
		{
			name: "example 1",
			args: args{
				c: &config.Config{
					BaseURL: "https://my.base.url",
					Formats: formats.FormatDefinitions{
						formats.FormatOsmPbf: {
							ID:  formats.FormatOsmPbf,
							Loc: ".osm.pbf",
							// BasePath: "/",
						}, formats.FormatState: {
							ID:       formats.FormatState,
							Loc:      "-updates/state.txt",
							BasePath: "../state/",
						}, formats.FormatPoly: {
							ID: formats.FormatPoly,

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
					},
					Elements: element.Slice{
						"africa": {
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
						},
					},
				},
			},
			want: element.Formats{
				formats.FormatKml,
				formats.FormatOshPbf,
				"osh.pbf.md5",
				formats.FormatOsmBz2,
				"osm.bz2.md5",
				formats.FormatOsmPbf,
				"osm.pbf.md5",
				formats.FormatPoly,
				formats.FormatState,
			},
		},
		{
			name: "example 2",
			args: args{
				c: &config.Config{
					BaseURL: "https://my.base.url",
					Formats: formats.FormatDefinitions{},
					Elements: element.Slice{
						"africa": {
							ID:   "africa",
							Name: "Africa",
							Formats: []string{
								formats.FormatOsmPbf,
								formats.FormatGeoJSON,
								formats.FormatPoly,
								formats.FormatState,
							},
						},
					},
				},
			},
			want: element.Formats{
				formats.FormatGeoJSON,
				formats.FormatOsmPbf,
				formats.FormatPoly,
				formats.FormatState,
			},
		},
	}

	for _, tt := range tests {
		myTest := tt

		t.Run(myTest.name, func(t *testing.T) {
			t.Parallel()

			af := myTest.args.c.Elements["africa"]
			generator.Cleanup(myTest.args.c)
			// compare af.Formats != tt.want
			if !reflect.DeepEqual(af.Formats, myTest.want) {
				t.Errorf("Cleanup() = %v, want %v", af.Formats, myTest.want)
			}
		})
	}
}
