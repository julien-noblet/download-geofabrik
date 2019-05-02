package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	yaml "gopkg.in/yaml.v2"
)

//Sample data:
func getHTML(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

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
				element: &Element{ID: "us", Parent: "north-america", Formats: []string{"osm.pbf"}},
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
		t.Run(tt.name, func(t *testing.T) {
			e := Ext{
				Elements: sampleElementValidPtr,
			}
			if err := e.mergeElement(tt.args.element); err != nil != tt.wantErr {
				t.Errorf("Ext.mergeElement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func createHTMLElement(t *testing.T, in string) *colly.HTMLElement {
	sel := "*"
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer([]byte(in)))
	if err != nil {
		t.Fatal(err)
	}
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})
	return elements[0]
}

func Test_contains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "Contain", args: args{s: []string{"a"}, e: "a"}, want: true},
		{name: "Contain", args: args{s: []string{"a", "b"}, e: "a"}, want: true},
		{name: "Not Contain", args: args{s: []string{"a", "b"}, e: "c"}, want: false},
		{name: "Void s", args: args{s: []string{}, e: "c"}, want: false},
		{name: "Void e", args: args{s: []string{"a", "b"}, e: ""}, want: false},
		{name: "Void s,e", args: args{s: []string{}, e: ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
