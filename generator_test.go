package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
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

var osmfrValidHTML1 = getHTML("http://download.openstreetmap.fr/extracts/")

var osmfrValidHTML2 = getHTML("http://download.openstreetmap.fr/extracts/merge/")

var osmfrPolygonJapanValidHTML = getHTML("http://download.openstreetmap.fr/polygons/asia/japan/")

var geofabrikSouthAmericaHTML = getHTML("https://download.geofabrik.de/south-america.html")

var geofabrikDistrictOfColumbiaHTML = getHTML("https://download.geofabrik.de/north-america/us/district-of-columbia.html")

var geofabrikShikokuHTML = getHTML("https://download.geofabrik.de/asia/japan/shikoku.html")

var gislabSampleHTML = getHTML("http://be.gis-lab.info/project/osm_dump/iframe.php")

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
				DefaultExtender: new(gocrawl.DefaultExtender),
				Elements:        sampleElementValidPtr,
			}
			if err := e.mergeElement(tt.args.element); err != nil != tt.wantErr {
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

func TestExt_parseOSMfr(t *testing.T) {
	var f func(s string, myUrl string) *goquery.Document
	f = func(s string, myUrl string) *goquery.Document {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
		doc.Url, err = url.Parse(myUrl)
		if err != nil {
			t.Error(err)
		}
		return doc
	}
	type args struct {
		ctx *gocrawl.URLContext // not used
		res *http.Response      // not used
		doc *goquery.Document
	}
	tests := []struct {
		name   string
		fields Ext
		args   args
		want   ElementSlice
		want1  bool
	}{
		// TODO: Add test cases.
		{
			name:   "osm valid1",
			fields: Ext{Elements: ElementSlice{}},
			args:   args{doc: f(osmfrValidHTML1, "http://download.openstreetmap.fr/extracts/")},
			want1:  true,
			want: ElementSlice{
				"africa":          {ID: "africa", Meta: false, Name: "africa", Formats: []string{"osm.pbf", "osm.pbf.md5", "state"}, Parent: "", File: ""},
				"asia":            {ID: "asia", Meta: false, Name: "asia", Formats: []string{"osm.pbf", "osm.pbf.md5", "state"}, Parent: "", File: ""},
				"central-america": {ID: "central-america", Meta: false, Name: "central-america", Formats: []string{"osm.pbf", "osm.pbf.md5", "state"}, Parent: "", File: ""},
				"europe":          {ID: "europe", Meta: false, Name: "europe", Formats: []string{"osm.pbf", "osm.pbf.md5", "state"}, Parent: "", File: ""},
				"north-america":   {ID: "north-america", Meta: false, Name: "north-america", Formats: []string{"osm.pbf", "osm.pbf.md5", "state"}, Parent: "", File: ""},
				"oceania":         {ID: "oceania", Meta: false, Name: "oceania", Formats: []string{"osm.pbf", "osm.pbf.md5", "state"}, Parent: "", File: ""},
				"russia":          {ID: "russia", Meta: false, Name: "russia", Formats: []string{"osm.pbf", "osm.pbf.md5", "state"}, Parent: "", File: ""},
				"south-america":   {ID: "south-america", Meta: false, Name: "south-america", Formats: []string{"osm.pbf", "osm.pbf.md5", "state"}, Parent: "", File: ""},
			},
		},
		{
			name:   "osm valid2",
			fields: Ext{Elements: ElementSlice{}},
			args:   args{doc: f(osmfrValidHTML2, "http://download.openstreetmap.fr/extracts/merge/")},
			want: ElementSlice{
				"france_metro_dom_com_nc": {ID: "france_metro_dom_com_nc", File: "", Meta: false, Name: "france_metro_dom_com_nc", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"france_taaf":             {ID: "france_taaf", File: "", Meta: false, Name: "france_taaf", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				//"israel_and_palestine":    {ID: "israel_and_palestine", File: "", Meta: false, Name: "israel_and_palestine", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"kiribati": {ID: "kiribati", File: "", Meta: false, Name: "kiribati", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"fiji":     {ID: "fiji", File: "", Meta: false, Name: "fiji", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
			},
			want1: true,
		}, {
			name:   "osm Poly valid",
			fields: Ext{Elements: ElementSlice{}},
			args:   args{doc: f(osmfrPolygonJapanValidHTML, "http://download.openstreetmap.fr/polygons/asia/japan/")},
			want1:  true,
			want: ElementSlice{
				"chubu":    {ID: "chubu", Meta: false, Name: "chubu", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"chugoku":  {ID: "chugoku", Meta: false, Name: "chugoku", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"hokkaido": {ID: "hokkaido", Meta: false, Name: "hokkaido", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"kansai":   {ID: "kansai", Meta: false, Name: "kansai", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"kanto":    {ID: "kanto", Meta: false, Name: "kanto", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"kyushu":   {ID: "kyushu", Meta: false, Name: "kyushu", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"shikoku":  {ID: "shikoku", Meta: false, Name: "shikoku", Formats: []string{"poly"}, Parent: "japan", File: ""},
				"tohoku":   {ID: "tohoku", Meta: false, Name: "tohoku", Formats: []string{"poly"}, Parent: "japan", File: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: tt.fields.DefaultExtender,
				Elements:        tt.fields.Elements,
			}
			_, got1 := e.parseOSMfr(tt.args.ctx, tt.args.res, tt.args.doc)
			if !reflect.DeepEqual(e.Elements, tt.want) {
				t.Errorf("Ext.parseOSMfr() \nExt.Elements= %+v, want %+v", e.Elements, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Ext.parseOSMfr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExt_parseGisLab(t *testing.T) {
	var f func(s string, myUrl string) *goquery.Document
	f = func(s string, myUrl string) *goquery.Document {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
		doc.Url, err = url.Parse(myUrl)
		if err != nil {
			t.Error(err)
		}
		return doc
	}
	type args struct {
		ctx *gocrawl.URLContext // not used
		res *http.Response      // not used
		doc *goquery.Document
	}
	tests := []struct {
		name   string
		fields Ext
		args   args
		want   ElementSlice
		want1  bool
	}{
		// TODO: Add test cases.
		{
			name:   "osm valid1",
			fields: Ext{Elements: ElementSlice{}},
			args:   args{doc: f(gislabSampleHTML, "http://be.gis-lab.info/project/osm_dump/iframe.php")},
			want1:  true,
			want: ElementSlice{
				"local": {ID: "local", Name: "локальное покрытие", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"AM":    {ID: "AM", Name: "Армения", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"AZ":    {ID: "AZ", Name: "Азербайджан", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"BY":    {ID: "BY", Name: "Беларусь", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"EE":    {ID: "EE", Name: "Эстония", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"TJ":    {ID: "TJ", Name: "Таджикистан", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"TM":    {ID: "TM", Name: "Туркмения", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"UA":    {ID: "UA", Name: "Украина", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"UZ":    {ID: "UZ", Name: "Узбекистан", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"RU-AD": {ID: "RU-AD", Name: "Адыгея", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"RU-AL": {ID: "RU-AL", Name: "Алтай", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
				"RU":    {ID: "RU", Name: "Российская Федерация", Formats: []string{"osm.pbf", "osm.bz2", "poly"}, File: "", Meta: false, Parent: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: tt.fields.DefaultExtender,
				Elements:        tt.fields.Elements,
			}
			_, got1 := e.parseGisLab(tt.args.ctx, tt.args.res, tt.args.doc)
			if !reflect.DeepEqual(e.Elements, tt.want) {
				// check one by one, if gislab have add some files don't panic!
				for i, k := range tt.want {
					if !reflect.DeepEqual(e.Elements[i], k) {
						t.Errorf("Ext.parseGisLab() \nExt.Elements= %+v, want %+v", e.Elements[i], k)
					}
				}
				//t.Errorf("Ext.parseGisLab() \nExt.Elements= %+v, want %+v", e.Elements, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Ext.parseGisLab() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Benchmark_Parser_osmfrValidHTML1(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(osmfrValidHTML1))
	doc.Url, err = url.Parse("http://download.openstreetmap.fr/extracts/")
	if err != nil {
		b.Error(err)
	}

	ext := Ext{Elements: ElementSlice{}}

	for n := 0; n < b.N; n++ {
		ext.parseOSMfr(nil, nil, doc)
	}
}
func Benchmark_Parser_osmfrValidHTML2(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(osmfrValidHTML2))
	doc.Url, err = url.Parse("http://download.openstreetmap.fr/extracts/merge/")
	if err != nil {
		b.Error(err)
	}

	ext := Ext{Elements: ElementSlice{}}

	for n := 0; n < b.N; n++ {
		ext.parseOSMfr(nil, nil, doc)
	}
}

func Benchmark_Parser_osmfrPolygonJapanValidHTML(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(osmfrPolygonJapanValidHTML))
	doc.Url, err = url.Parse("http://download.openstreetmap.fr/polygons/asia/japan/")
	if err != nil {
		b.Error(err)
	}

	ext := Ext{Elements: ElementSlice{}}

	for n := 0; n < b.N; n++ {
		ext.parseOSMfr(nil, nil, doc)
	}
}

func TestExt_parseGeofabrik(t *testing.T) {
	var f func(s string) *goquery.Document //, myUrl string) *goquery.Document
	f = func(s string) *goquery.Document { //, myUrl string) *goquery.Document {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
		//	doc.Url, err = url.Parse(myUrl) // not needed
		if err != nil {
			t.Error(err)
		}
		return doc
	}
	type args struct {
		ctx *gocrawl.URLContext // not used
		res *http.Response      // not used
		doc *goquery.Document
	}
	tests := []struct {
		name  string
		args  args         // Contain tests
		want  ElementSlice // to compare with e.Elements
		want1 bool         // always true
	}{
		// TODO: Add test cases.
		{
			name:  "Parse Geofabrik SouthAmerica",
			args:  args{doc: f(geofabrikSouthAmericaHTML)},
			want1: true,
			want: ElementSlice{
				"us":            {ID: "us", File: "", Meta: true, Name: "United States of America", Formats: []string{}, Parent: "north-america"},
				"south-america": {ID: "south-america", File: "", Meta: false, Name: "South America", Formats: []string{"osm.pbf", "osm.pbf.md5", "osm.bz2", "osm.bz2.md5", "poly", "kml", "state"}, Parent: ""},
			},
		},
		{
			name:  "Parse Geofabrik District of Columbia",
			args:  args{doc: f(geofabrikDistrictOfColumbiaHTML)},
			want1: true,
			want: ElementSlice{
				"us":                   {ID: "us", File: "", Meta: true, Name: "United States of America", Formats: []string{}, Parent: "north-america"},
				"district-of-columbia": {ID: "district-of-columbia", File: "", Meta: false, Name: "District of Columbia", Formats: []string{"osm.pbf", "osm.pbf.md5", "shp.zip", "osm.bz2", "osm.bz2.md5", "poly", "kml", "state"}, Parent: "us"},
			},
		},
		{
			name:  "Parse Geofabrik Shikoku",
			args:  args{doc: f(geofabrikShikokuHTML)},
			want1: true,
			want: ElementSlice{
				"us":      {ID: "us", File: "", Meta: true, Name: "United States of America", Formats: []string{}, Parent: "north-america"},
				"shikoku": {ID: "shikoku", File: "", Meta: false, Name: "Shikoku", Formats: []string{"osm.pbf", "osm.pbf.md5", "shp.zip", "osm.bz2", "osm.bz2.md5", "poly", "kml", "state"}, Parent: "japan"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Ext{
				DefaultExtender: new(gocrawl.DefaultExtender),
				Elements:        ElementSlice{},
			}
			e.parseGeofabrik(tt.args.ctx, tt.args.res, tt.args.doc)

			if !reflect.DeepEqual(e.Elements, tt.want) {
				t.Errorf("Ext.parseGeofabrik() got = %+v, want %+v", e.Elements, tt.want)
			}
			/* // Since it's always true, not need to check it!
			if got1 != tt.want1 {
				t.Errorf("Ext.parseGeofabrik() got1 = %v, want %v", got1, tt.want1)
			}
			*/
		})
	}
}

func Benchmark_Parser_geofabrikSouthAmericaHTML(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(geofabrikSouthAmericaHTML))
	if err != nil {
		b.Error(err)
	}
	ext := Ext{Elements: ElementSlice{}}
	for n := 0; n < b.N; n++ {
		ext.parseGeofabrik(nil, nil, doc)
	}
}

func Benchmark_Parser_geofabrikDistrictOfColumbiaHTML(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(geofabrikDistrictOfColumbiaHTML))
	if err != nil {
		b.Error(err)
	}
	ext := Ext{Elements: ElementSlice{}}
	for n := 0; n < b.N; n++ {
		ext.parseGeofabrik(nil, nil, doc)
	}
}

func Benchmark_Parser_geofabrikShikokuHTML(b *testing.B) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(geofabrikShikokuHTML))
	if err != nil {
		b.Error(err)
	}
	ext := Ext{Elements: ElementSlice{}}
	for n := 0; n < b.N; n++ {
		ext.parseGeofabrik(nil, nil, doc)
	}
}
