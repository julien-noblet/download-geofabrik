package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	yaml "gopkg.in/yaml.v2"
)

func TestElementSlice_Generate(t *testing.T) {
	myConfig := &SampleConfigValidPtr
	myYaml, _ := yaml.Marshal(*myConfig)
	myConfig.Elements = map[string]Element{} // void Elements

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

func TestExt_mergeElement(t *testing.T) {
	myFakeGeorgia := sampleFakeGeorgia2Ptr
	myFakeGeorgia.ID = "georgia-us"
	type args struct {
		element *Element
	}
	tests := []struct {
		name    string
		fields  Ext
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "adding a new element",
			fields: Ext{
				DefaultExtender: new(gocrawl.DefaultExtender),
				Elements:        sampleElementValidPtr,
			},
			args: args{
				element: &sampleFakeGeorgiaPtr,
			},
			wantErr: false,
		},
		{
			name: "adding a same element",
			fields: Ext{
				DefaultExtender: new(gocrawl.DefaultExtender),
				Elements:        sampleElementValidPtr,
			},
			args: args{
				element: &sampleAfricaElementPtr,
			},
			wantErr: false,
		},
		{
			name: "adding a meta element",
			fields: Ext{
				DefaultExtender: new(gocrawl.DefaultExtender),
				Elements:        sampleElementValidPtr,
			},
			args: args{
				element: &sampleUsElementPtr,
			},
			wantErr: false,
		},
		{
			name: "adding a same element parent differ",
			fields: Ext{
				DefaultExtender: new(gocrawl.DefaultExtender),
				Elements:        sampleElementValidPtr,
			},
			args: args{
				element: &myFakeGeorgia,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: tt.fields.DefaultExtender,
				Elements:        tt.fields.Elements,
			}
			if err := e.mergeElement(tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("Ext.mergeElement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Benchmark_Element_addHash_noHash(b *testing.B) {
	sampleElement := Element{
		ID:      "test",
		Formats: []string{"osm.pbf"},
	}
	sampleHTML := `<p><a href="example.osm.pbf" >example.osm.pbf</a></p>`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTML))
	if err != nil {
		b.Error(err)
	}
	sampleSelection := doc.Find("p")
	for n := 0; n < b.N; n++ {
		sampleElement.addHash(sampleSelection)
	}
}
func Benchmark_Element_addHash_Hash(b *testing.B) {
	sampleElement := Element{
		ID:      "test",
		Formats: []string{"osm.pbf"},
	}
	sampleHTML := `<p><a href="example.osm.pbf" >example.osm.pbf</a><a href="example.osm.pbf.md5" >example.osm.pbf.md5</a></p>`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTML))
	if err != nil {
		b.Error(err)
	}
	sampleSelection := doc.Find("p")
	for n := 0; n < b.N; n++ {
		sampleElement.addHash(sampleSelection)
	}
}

func TestElement_addHash(t *testing.T) {
	sampleElement := Element{
		ID:      "test",
		Formats: []string{"osm.pbf"},
	}
	sampleElementWithHash := sampleElement
	sampleElementWithHash.Formats = append(sampleElementWithHash.Formats, "osm.pbf.md5")
	sampleHTMLWithHash := "<p><a href=\"example.osm.pbf\">example.osm.pbf</a> and <a href=\"example.osm.pbf.md5\">example.osm.pbf.md5</a></p>"
	docWithHash, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTMLWithHash))
	if err != nil {
		t.Error(err)
	}
	sampleSelectionWithHash := docWithHash.Find("p")
	sampleHTMLWithoutHash := `<p><a href="example.osm.pbf" >example.osm.pbf</a></p>`
	docWithoutHash, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTMLWithoutHash))
	if err != nil {
		t.Error(err)
	}
	sampleSelectionWithoutHash := docWithoutHash.Find("p")
	sampleHTMLWithWrongHash := "<p><a href=\"example.osm.pbf\">example.osm.pbf</a> and <a href=\"example.osm.pbf.sha\">example.osm.pbf.sha</a></p>"
	docWithWrongHash, err := goquery.NewDocumentFromReader(strings.NewReader(sampleHTMLWithWrongHash))
	if err != nil {
		t.Error(err)
	}
	sampleSelectionWithWrongHash := docWithWrongHash.Find("p")

	type args struct {
		myel *goquery.Selection
	}
	tests := []struct {
		name        string
		fields      Element
		args        args
		wantElement Element
	}{
		// TODO: Add test cases.
		{
			name:        "Hash found",
			fields:      sampleElement,
			args:        args{myel: sampleSelectionWithHash},
			wantElement: sampleElementWithHash,
		}, {
			name:        "NoHash found",
			fields:      sampleElement,
			args:        args{myel: sampleSelectionWithoutHash},
			wantElement: sampleElement,
		}, {
			name:        "Wrong Hash found",
			fields:      sampleElement,
			args:        args{myel: sampleSelectionWithWrongHash},
			wantElement: sampleElement,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Element{
				ID:      tt.fields.ID,
				File:    tt.fields.File,
				Meta:    tt.fields.Meta,
				Name:    tt.fields.Name,
				Formats: tt.fields.Formats,
				Parent:  tt.fields.Parent,
			}

			e.addHash(tt.args.myel)
			if !reflect.DeepEqual(*e, tt.wantElement) {
				t.Errorf("addHash(%v) e=%v, want %v", tt.args.myel, *e, tt.wantElement)
			}
		})
	}
}
