// 2015-2019 copyright Julien Noblet

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"sync"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

//Sample data:
func getHTML(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func TestElementSlice_Generate(t *testing.T) {
	tests := []struct {
		name    string
		e       ElementSlice
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Marshalling OK, no error",
			e:       sampleElementValidPtr,
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
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
		t.Run(tt.name, func(t *testing.T) {
			Generate(tt.args.configfile)
		})
	}
}

func TestConfig_mergeElement(t *testing.T) {
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

// createHTMLElement is not used --- removing dead code!
// func createHTMLElement(t *testing.T, in string) *colly.HTMLElement {
// 	sel := "*"
// 	ctx := &colly.Context{}
// 	resp := &colly.Response{
// 		Request: &colly.Request{
// 			Ctx: ctx,
// 		},
// 		Ctx: ctx,
// 	}
// 	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer([]byte(in)))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	elements := []*colly.HTMLElement{}
// 	i := 0
// 	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
// 		for _, n := range s.Nodes {
// 			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
// 			i++
// 		}
// 	})
// 	return elements[0]
// }
