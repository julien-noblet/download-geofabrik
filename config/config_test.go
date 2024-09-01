package config_test

import (
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
)

const (
	geofabrikYml       = "../geofabrik.yml"
	openstreetmapFRYml = "../openstreetmap.fr.yml"
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
			Ext:      "",
			Loc:      "",
			BasePath: "",
		}, formats.FormatOsmPbf: {
			Ext: formats.FormatOsmPbf,
			Loc: ".osm.pbf",
			// BasePath: "/",
		}, formats.FormatState: {
			Ext:      formats.FormatState,
			Loc:      "-updates/state.txt",
			BasePath: "../state/",
		}, formats.FormatPoly: {
			Ext:     formats.FormatPoly,
			Loc:     ".poly",
			BaseURL: "http://my.new.url/folder",
		}, formats.FormatOsmBz2: {
			Ext:      formats.FormatOsmBz2,
			Loc:      ".osm.bz2",
			BasePath: "../osmbz2/",
			BaseURL:  "http://my.new.url/folder",
		}, formats.FormatOsmGz: {
			Ext:      formats.FormatOsmGz,
			Loc:      ".osm.gz",
			BasePath: "../osmgz/",
			BaseURL:  "http://my.new.url/folder",
		},
	}
}

func SampleConfigValidPtr() *config.Config {
	return &config.Config{
		BaseURL:  "https://my.base.url",
		Formats:  sampleFormatValidPtr(),
		Elements: sampleElementValidPtr(),
	}
}

func sampleFakeGeorgiaPtr() *element.Element {
	return &element.Element{
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
}

func sampleFakeGeorgia2Ptr() *element.Element {
	return &element.Element{
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
}

func sampleFakeGeorgia3Ptr() *element.Element {
	return &element.Element{
		ID:   "georgia-us3",
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
		Parent: "notus2",
	}
}

func sampleNotUS2Ptr() *element.Element {
	return &element.Element{
		ID:     "notus2",
		Name:   "notus2",
		Meta:   true,
		Parent: "north", // bad parent not exist!
	}
}

func Benchmark_Exist_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := config.LoadConfig(geofabrikYml)
	e := "france"

	for n := 0; n < b.N; n++ {
		c.Exist(e)
	}
}

func Benchmark_loadConfig_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		if _, err := config.LoadConfig(geofabrikYml); err != nil {
			b.Error(err.Error())
		}
	}
}

func Benchmark_loadConfig_osmfr_yml(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		if _, err := config.LoadConfig(openstreetmapFRYml); err != nil {
			b.Error(err.Error())
		}
	}
}

func Test_loadConfig(t *testing.T) {
	t.Parallel()

	type args struct {
		configFile string
	}

	tests := []struct {
		args    args
		want    *config.Config
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Check config file not exist",
			args:    args{configFile: "./this_file_not_exists"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Check config not yaml",
			args:    args{configFile: "../LICENSE"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check config valid geofabrik.yml",
			args: args{configFile: geofabrikYml},
			want: &config.Config{
				BaseURL:       "https://download.geofabrik.de",
				Formats:       sampleFormatValidPtr(),
				Elements:      sampleElementValidPtr(),
				ElementsMutex: &sync.RWMutex{},
			},
			wantErr: false,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := config.LoadConfig(thisTest.args.configFile)
			if err != nil != thisTest.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, thisTest.wantErr)

				return
			}
			if got != nil {
				if reflect.TypeOf(got) != reflect.TypeOf(thisTest.want) {
					t.Errorf("loadConfig() is a %v, want %v", reflect.TypeOf(got), reflect.TypeOf(thisTest.want))
				}
				// Check BaseURL
				if got.BaseURL != thisTest.want.BaseURL {
					t.Errorf("loadConfig().BaseURL = %v, want %v", got.BaseURL, thisTest.want.BaseURL)
				}
				// Must have formats but ne need to check if they are valids!
				if reflect.TypeOf(got.Formats) != reflect.TypeOf(thisTest.want.Formats) {
					t.Errorf("loadConfig().Formats is a %v, want %v", reflect.TypeOf(got.Formats), reflect.TypeOf(thisTest.want.Formats))
				}
				// Must have Elements but no check if they are valids
				if reflect.TypeOf(got.Elements) != reflect.TypeOf(thisTest.want.Elements) {
					t.Errorf("loadConfig().Elements is a %v, want %v", reflect.TypeOf(got.Elements), reflect.TypeOf(thisTest.want.Elements))
				}
			}
		})
	}
}

func Test_AddExtension(t *testing.T) {
	t.Parallel()

	type args struct {
		id     string
		format string
	}

	tests := []struct {
		args args
		c    config.Config
		want element.Slice
		name string
	}{
		{
			name: "Add osm.pbf but already in",
			c: config.Config{
				Elements: element.Slice{
					"a": element.Element{
						ID:      "a",
						Name:    "a",
						Formats: []string{formats.FormatOsmPbf},
						Meta:    false,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
			args: args{
				id:     "a",
				format: formats.FormatOsmPbf,
			},
			want: element.Slice{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf",
			c: config.Config{
				Elements: element.Slice{
					"a": element.Element{
						ID:      "a",
						Name:    "a",
						Formats: []string{},
						Meta:    false,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
			args: args{
				id:     "a",
				format: formats.FormatOsmPbf,
			},
			want: element.Slice{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf},
					Meta:    false,
				},
			},
		},
		{
			name: "Add osm.pbf on meta",
			c: config.Config{
				Elements: element.Slice{
					"a": element.Element{
						ID:      "a",
						Name:    "a",
						Formats: []string{},
						Meta:    true,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
			args: args{
				id:     "a",
				format: formats.FormatOsmPbf,
			},
			want: element.Slice{
				"a": element.Element{
					ID:      "a",
					Name:    "a",
					Formats: []string{formats.FormatOsmPbf},
					Meta:    false,
				},
			},
		},
	}

	for tn := range tests {
		tn := tn
		t.Run(tests[tn].name, func(t *testing.T) {
			t.Parallel()
			tests[tn].c.AddExtension(tests[tn].args.id, tests[tn].args.format)
			if !reflect.DeepEqual(tests[tn].c.Elements, tests[tn].want) {
				t.Errorf("AddExtension() got %v, want %v", tests[tn].c.Elements, tests[tn].want)
			}
		})
	}
}

func Benchmark_GetElement_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := config.LoadConfig(geofabrikYml)
	e := "france"

	for n := 0; n < b.N; n++ {
		if _, err := c.GetElement(e); err != nil {
			panic(err)
		}
	}
}

func Test_config_GetElement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		want    *element.Element
		name    string
		file    string
		id      string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "France from geofabrik",
			file: geofabrikYml,
			id:   "france",
			want: &element.Element{
				ID:   "france",
				Name: "France", Parent: "europe",
				Formats: element.Formats{
					formats.FormatOshPbf,
					formats.FormatOsmPbf,
					"osm.pbf.md5",
					formats.FormatOsmBz2,
					"osm.bz2.md5",
					formats.FormatPoly,
					formats.FormatKml,
					formats.FormatState,
				},
			},
			wantErr: false,
		},
		{
			name:    "Franc from geofabrik",
			file:    geofabrikYml,
			id:      "franc",
			want:    nil,
			wantErr: true,
		},
	}
	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			c, _ := config.LoadConfig(thisTest.file)
			got, err := c.GetElement(thisTest.id)
			if (err != nil) != thisTest.wantErr {
				t.Errorf("Config.GetElement() error = %v, wantErr %v", err, thisTest.wantErr)

				return
			}
			if thisTest.want != nil {
				tfrance := &element.Element{
					ID:     thisTest.want.ID,
					Name:   thisTest.want.Name,
					File:   thisTest.want.File,
					Parent: thisTest.want.Parent,
					Meta:   thisTest.want.Meta,
				}
				tgot := &element.Element{
					ID:     got.ID,
					Name:   got.Name,
					File:   got.File,
					Parent: got.Parent,
					Meta:   got.Meta,
				}
				if !reflect.DeepEqual(tgot, tfrance) {
					t.Errorf("Config.GetElement() = %v, want %v", got, thisTest.want)
				}

				for _, k := range thisTest.want.Formats {
					if !(got.Formats.Contains(k)) {
						t.Errorf("%v format should exist, got=%v", k, got)
					}
				}

				for _, k := range got.Formats {
					if !(thisTest.want.Formats.Contains(k)) {
						t.Errorf("%v format should not exist, got=%v", k, got)
					}
				}
			} else if !reflect.DeepEqual(got, thisTest.want) {
				t.Errorf("Config.GetElement() = %v, want %v", got, thisTest.want)
			}
			if !thisTest.wantErr && (err == nil) {
				for _, k := range thisTest.want.Formats {
					if !(got.Formats.Contains(k)) {
						t.Errorf("%v format should exist, got=%v", k, got)
					}
				}

				for _, k := range got.Formats {
					if !(thisTest.want.Formats.Contains(k)) {
						t.Errorf("%v format should not exist, got=%v", k, got)
					}
				}
			}
		})
	}
}

func Test_IsHashable(t *testing.T) {
	t.Parallel()

	type args struct {
		format string
		file   string
	}

	tests := []struct {
		args  args
		name  string
		want1 string
		want2 string
		want  bool
	}{
		// TODO: Add test cases.
		{name: "Test is osm.pbf is hashable", args: args{format: formats.FormatOsmPbf, file: geofabrikYml}, want: true, want1: "osm.pbf.md5", want2: "md5"},
		{name: "Test is kml is hashable", args: args{format: formats.FormatKml, file: geofabrikYml}, want: false, want1: "", want2: ""},
	}

	for _, thisTest := range tests {
		thisTest := thisTest

		myConfig, err := config.LoadConfig(thisTest.args.file)
		if err != nil {
			t.Error(err)
		}

		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, got1, got2 := config.IsHashable(myConfig, thisTest.args.format)
			if got != thisTest.want {
				t.Errorf("config.IsHashable() got = %v, want %v", got, thisTest.want)
			}
			if got1 != thisTest.want1 {
				t.Errorf("config.IsHashable() got1 = %v, want %v", got1, thisTest.want1)
			}
			if got2 != thisTest.want2 {
				t.Errorf("config.IsHashable() got2 = %v, want %v", got2, thisTest.want2)
			}
		})
	}
}

func Benchmark_IsHashable_geofabrik_yml(b *testing.B) {
	// run the Fib function b.N times
	c, _ := config.LoadConfig(geofabrikYml)

	for n := 0; n < b.N; n++ {
		for f := range c.Formats {
			config.IsHashable(c, f)
		}
	}
}

func Benchmark_findElem_parse_all_geofabrik_yml(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)

	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			if _, err := config.FindElem(c, k); err != nil {
				b.Error(err.Error())
			}
		}
	}
}

func Benchmark_GetElement_parse_all_geofabrik_yml(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)

	for n := 0; n < b.N; n++ {
		for k := range c.Elements {
			if _, err := c.GetElement(k); err != nil {
				panic(err)
			}
		}
	}
}

func Benchmark_FindElem_parse_France_geofabrik_yml(b *testing.B) {
	c, _ := config.LoadConfig(geofabrikYml)

	for n := 0; n < b.N; n++ {
		if _, err := config.FindElem(c, "france"); err != nil {
			b.Errorf("%v", err.Error())
		}
	}
}

func Test_FindElem(t *testing.T) {
	t.Parallel()

	type args struct {
		c *config.Config
		e string
	}

	tests := []struct {
		args    args
		want    *element.Element
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Find",
			args: args{
				c: SampleConfigValidPtr(),
				e: "africa",
			},
			want: &element.Element{
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
			wantErr: false,
		},
		{
			name: "Can't find notInList",
			args: args{
				c: SampleConfigValidPtr(),
				e: "notInList",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should not find",
			args: args{
				c: SampleConfigValidPtr(),
				e: "",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := config.FindElem(thisTest.args.c, thisTest.args.e)
			if err != nil != thisTest.wantErr {
				t.Errorf("config.FindElem() error = %v, wantErr %v", err, thisTest.wantErr)

				return
			}
			if !reflect.DeepEqual(got, thisTest.want) {
				t.Errorf("config.FindElem() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

func Test_elem2URL(t *testing.T) {
	t.Parallel()

	type args struct {
		c   *config.Config
		e   *element.Element
		ext string
	}

	localSampleConfigValidPtr := SampleConfigValidPtr()
	localSampleConfigValidPtr.Elements["georgia-us2"] = *sampleFakeGeorgiaPtr() // add it into config
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "top level test config",
			args:    args{c: localSampleConfigValidPtr, e: sampleAfricaElementPtr(), ext: formats.FormatOsmPbf},
			want:    "https://my.base.url/africa.osm.pbf",
			wantErr: false,
		},
		{
			name:    "sub level test config",
			args:    args{c: localSampleConfigValidPtr, e: sampleGeorgiaUsElementPtr(), ext: formats.FormatState},
			want:    "https://my.base.url/../state/north-america/us/georgia-updates/state.txt",
			wantErr: false,
		},
		{
			name:    "BaseUrl test config",
			args:    args{c: localSampleConfigValidPtr, e: sampleGeorgiaUsElementPtr(), ext: formats.FormatPoly},
			want:    "http://my.new.url/folder/north-america/us/georgia.poly",
			wantErr: false,
		},
		{
			name:    "BaseUrl + BasePath test config",
			args:    args{c: localSampleConfigValidPtr, e: sampleGeorgiaUsElementPtr(), ext: formats.FormatOsmBz2},
			want:    "http://my.new.url/folder/../osmbz2/north-america/us/georgia.osm.bz2",
			wantErr: false,
		},
		{
			name:    "sub level test config wrong format",
			args:    args{c: localSampleConfigValidPtr, e: sampleGeorgiaUsElementPtr(), ext: "wrongFmt"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "sub level test config not exists in config",
			args:    args{c: localSampleConfigValidPtr, e: sampleFakeGeorgiaPtr(), ext: formats.FormatState},
			want:    "",
			wantErr: true,
		},
		{
			name:    "sub level test config not exists parent",
			args:    args{c: localSampleConfigValidPtr, e: sampleFakeGeorgia2Ptr(), ext: formats.FormatState},
			want:    "",
			wantErr: true,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := config.Elem2URL(thisTest.args.c, thisTest.args.e, thisTest.args.ext)
			if err != nil != thisTest.wantErr {
				t.Errorf("elem2URL() error = %v, wantErr %v", err, thisTest.wantErr)

				return
			}
			if got != thisTest.want {
				t.Errorf("elem2URL() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

func Test_elem2preURL(t *testing.T) {
	t.Parallel()

	type args struct {
		c *config.Config
		e *element.Element
		b []string
	}

	localSampleConfigValidPtr := SampleConfigValidPtr()
	localSampleConfigValidPtr.Elements["georgia-us2"] = *sampleFakeGeorgia2Ptr() // add it into config
	localSampleConfigValidPtr.Elements["georgia-us3"] = *sampleFakeGeorgia3Ptr() // add it into config
	localSampleConfigValidPtr.Elements["notus2"] = *sampleNotUS2Ptr()            // add it into config
	tests := []struct {
		name    string
		want    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "top level test config",
			args:    args{c: localSampleConfigValidPtr, e: sampleAfricaElementPtr()},
			want:    "https://my.base.url/africa",
			wantErr: false,
		},
		{
			name:    "top level test config with basePath base/",
			args:    args{c: localSampleConfigValidPtr, e: sampleAfricaElementPtr(), b: []string{"base/"}},
			want:    "https://my.base.url/base/africa",
			wantErr: false,
		},
		{
			name:    "sub level test config",
			args:    args{c: localSampleConfigValidPtr, e: sampleGeorgiaUsElementPtr()},
			want:    "https://my.base.url/north-america/us/georgia",
			wantErr: false,
		},
		{
			name:    "sub level test config not exists in config",
			args:    args{c: localSampleConfigValidPtr, e: sampleFakeGeorgiaPtr()},
			want:    "",
			wantErr: true,
		},
		{
			name:    "sub level test config not exists parent",
			args:    args{c: localSampleConfigValidPtr, e: sampleFakeGeorgia2Ptr()},
			want:    "",
			wantErr: true,
		},
		{
			name:    "sub level test config parent exist but not grandparent",
			args:    args{c: localSampleConfigValidPtr, e: sampleFakeGeorgia3Ptr()},
			want:    "",
			wantErr: true,
		},
	}

	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			got, err := config.Elem2preURL(thisTest.args.c, thisTest.args.e, thisTest.args.b...)
			if err != nil != thisTest.wantErr {
				t.Errorf("config.Elem2preURL() =%v error = %v, wantErr %v", got, err, thisTest.wantErr)

				return
			}
			if got != thisTest.want {
				t.Errorf("config.Elem2preURL() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

func Benchmark_Elem2preURL_parse_France_geofabrik_yml(b *testing.B) {
	myConfig, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Error(err)
	}

	france, err := config.FindElem(myConfig, "france")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		_, err = config.Elem2preURL(myConfig, france)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_Elem2URL_parse_France_geofabrik_yml(b *testing.B) {
	myConfig, err := config.LoadConfig(geofabrikYml)
	if err != nil {
		b.Error(err)
	}

	france, err := config.FindElem(myConfig, "france")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		_, err = config.Elem2URL(myConfig, france, formats.FormatState)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_Elem2URL_parse_France_openstreetmap_fr_yml(b *testing.B) {
	myConfig, err := config.LoadConfig(openstreetmapFRYml)
	if err != nil {
		b.Error(err)
	}

	france, err := config.FindElem(myConfig, "france")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		_, err = config.Elem2URL(myConfig, france, formats.FormatPoly)
		if err != nil {
			b.Error(err)
		}
	}
}

func TestConfig_MergeElement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Config     *config.Config
		wantConfig *config.Config
		el         *element.Element
		name       string
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name:   "Add element on void Config",
			Config: &config.Config{ElementsMutex: &sync.RWMutex{}, Elements: element.Slice{}},
			el: &element.Element{
				ID:      "test",
				Name:    "Test",
				Formats: element.Formats{"format1", "format2"},
			},
			wantErr: false,
			wantConfig: &config.Config{
				Elements: element.Slice{
					"test": {
						ID:      "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
		},
		{
			name:   "Add meta element on void Config",
			Config: &config.Config{ElementsMutex: &sync.RWMutex{}, Elements: element.Slice{}},
			el: &element.Element{
				ID:   "test",
				Name: "Test",
				Meta: true,
			},
			wantErr: false,
			wantConfig: &config.Config{
				Elements: element.Slice{
					"test": {
						ID:   "test",
						Name: "Test",
						Meta: true,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
		},
		{
			name: "Add new element on non void Config",
			Config: &config.Config{
				ElementsMutex: &sync.RWMutex{},
				Elements: element.Slice{
					"test2": {
						ID:      "test2",
						Name:    "Test2",
						Formats: element.Formats{"format1", "format2"},
					},
				},
			},
			el: &element.Element{
				ID:      "test",
				Name:    "Test",
				Formats: element.Formats{"format1", "format2"},
			},
			wantErr: false,
			wantConfig: &config.Config{
				Elements: element.Slice{
					"test": {
						ID:      "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
					},
					"test2": {
						ID:      "test2",
						Name:    "Test2",
						Formats: element.Formats{"format1", "format2"},
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
		},
		{
			name: "Add same element on non void Config",
			Config: &config.Config{
				ElementsMutex: &sync.RWMutex{},
				Elements: element.Slice{
					"test": {
						ID:      "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
					},
				},
			},
			el: &element.Element{
				ID:      "test",
				Name:    "Test",
				Formats: element.Formats{"format1", "format2"},
			},
			wantErr: false,
			wantConfig: &config.Config{
				Elements: element.Slice{
					"test": {
						ID:      "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
		},
		{
			name: "Add same element on non void Config and el was meta",
			Config: &config.Config{
				ElementsMutex: &sync.RWMutex{},
				Elements: element.Slice{
					"test": {
						ID:   "test",
						Name: "Test",
						Meta: true,
					},
				},
			},
			el: &element.Element{
				ID:      "test",
				Name:    "Test",
				Formats: element.Formats{"format1", "format2"},
			},
			wantErr: false,
			wantConfig: &config.Config{
				Elements: element.Slice{
					"test": {
						ID:      "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
		},
		{
			name: "Add same void element on non void Config and el was meta",
			Config: &config.Config{
				ElementsMutex: &sync.RWMutex{},
				Elements: element.Slice{
					"test": {
						ID:   "test",
						Name: "Test",
						Meta: true,
					},
				},
			},
			el: &element.Element{
				ID:   "test",
				Name: "Test",
			},
			wantErr: false,
			wantConfig: &config.Config{
				Elements: element.Slice{
					"test": {
						ID:   "test",
						Name: "Test",
						Meta: true,
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
		},
		{
			name: "Add same element on non void Config with wrong parent",
			Config: &config.Config{
				ElementsMutex: &sync.RWMutex{},
				Elements: element.Slice{
					"test": {
						ID:      "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
						Parent:  "parent1",
					},
				},
			},
			el: &element.Element{
				ID:      "test",
				Name:    "Test",
				Formats: element.Formats{"format1", "format2"},
				Parent:  "parent2",
			},
			wantErr: true,
			wantConfig: &config.Config{
				Elements: element.Slice{
					"test": {
						ID:      "test",
						Name:    "Test",
						Formats: element.Formats{"format1", "format2"},
						Parent:  "parent1",
					},
				},
				ElementsMutex: &sync.RWMutex{},
			},
		},
	}
	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			myConfig := thisTest.Config
			err := myConfig.MergeElement(thisTest.el)
			if (err != nil) != thisTest.wantErr {
				t.Errorf("Config.MergeElement() error = %v, wantErr %v", err, thisTest.wantErr)
			}
			if !reflect.DeepEqual(myConfig.Elements, thisTest.wantConfig.Elements) {
				t.Errorf("Config.MergeElement() config.Elements = %v, wantConfig.Elements %v", myConfig.Elements, thisTest.wantConfig.Elements)
			}
		})
	}
}

func TestConfig_Generate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		file    string
		wantErr bool
	}{
		{name: "geofabrik", file: geofabrikYml, wantErr: false},
		// TODO: Add test cases.
	}
	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			c, _ := config.LoadConfig(thisTest.file)
			want, _ := os.ReadFile(geofabrikYml)
			got, err := c.Generate()
			if (err != nil) != thisTest.wantErr {
				t.Errorf("Config.Generate() error = %v, wantErr %v", err, thisTest.wantErr)

				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Config.Generate() = %v, want %v", got, want)
			}
		})
	}
}
