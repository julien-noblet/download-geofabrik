package main

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
)

func TestElementSlice_Generate(t *testing.T) {
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

func TestExt_parseGeofabrik(t *testing.T) {
	type fields struct {
		DefaultExtender *gocrawl.DefaultExtender
		Elements        ElementSlice
	}
	type args struct {
		ctx *gocrawl.URLContext
		res *http.Response
		doc *goquery.Document
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		//
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: tt.fields.DefaultExtender,
				Elements:        tt.fields.Elements,
			}
			got, got1 := e.parseGeofabrik(tt.args.ctx, tt.args.res, tt.args.doc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ext.parseGeofabrik() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Ext.parseGeofabrik() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExt_mergeElement(t *testing.T) {
	type fields struct {
		DefaultExtender *gocrawl.DefaultExtender
		Elements        ElementSlice
	}
	type args struct {
		element *Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: tt.fields.DefaultExtender,
				Elements:        tt.fields.Elements,
			}
			e.mergeElement(tt.args.element)
		})
	}
}

func TestExt_Visit(t *testing.T) {
	type fields struct {
		DefaultExtender *gocrawl.DefaultExtender
		Elements        ElementSlice
	}
	type args struct {
		ctx *gocrawl.URLContext
		res *http.Response
		doc *goquery.Document
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: tt.fields.DefaultExtender,
				Elements:        tt.fields.Elements,
			}
			got, got1 := e.Visit(tt.args.ctx, tt.args.res, tt.args.doc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ext.Visit() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Ext.Visit() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExt_Filter(t *testing.T) {
	type fields struct {
		DefaultExtender *gocrawl.DefaultExtender
		Elements        ElementSlice
	}
	type args struct {
		ctx       *gocrawl.URLContext
		isVisited bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: tt.fields.DefaultExtender,
				Elements:        tt.fields.Elements,
			}
			if got := e.Filter(tt.args.ctx, tt.args.isVisited); got != tt.want {
				t.Errorf("Ext.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateCrawler(t *testing.T) {
	type args struct {
		url      string
		fname    string
		myConfig *Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// No test because need to check a new file...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateCrawler(tt.args.url, tt.args.fname, tt.args.myConfig)
		})
	}
}
