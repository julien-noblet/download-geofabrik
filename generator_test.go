package main

import (
	"net/http"
	"net/url"
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

func TestExt_parseOSMfr(t *testing.T) {
	osmfrValid1 := `
	<h1>Index of /extracts</h1>
	<table><tbody><tr><th><img src="/icons/blank.gif" alt="[ICO]"></th><th><a href="?C=N;O=D">Name</a></th><th><a href="?C=M;O=A">Last modified</a></th><th><a href="?C=S;O=A">Size</a></th><th><a href="?C=D;O=A">Description</a></th></tr><tr><th colspan="5"><hr></th></tr>
	<tr><td valign="top"><img src="/icons/back.gif" alt="[DIR]"></td><td><a href="/">Parent Directory</a></td><td>&nbsp;</td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="africa-latest.osm.pbf">africa-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:43  </td><td align="right">2.3G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="africa.osm.pbf">africa.osm.pbf</a></td><td align="right">22-Mar-2018 02:43  </td><td align="right">2.3G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="africa.state.txt">africa.state.txt</a></td><td align="right">22-Mar-2018 02:25  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="africa/">africa/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="asia-latest.osm.pbf">asia-latest.osm.pbf</a></td><td align="right">22-Mar-2018 01:39  </td><td align="right">6.0G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="asia.osm.pbf">asia.osm.pbf</a></td><td align="right">22-Mar-2018 01:39  </td><td align="right">6.0G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="asia.state.txt">asia.state.txt</a></td><td align="right">22-Mar-2018 00:50  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="asia/">asia/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="central-america-latest.osm.pbf">central-america-latest.osm.pbf</a></td><td align="right">22-Mar-2018 00:52  </td><td align="right">336M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="central-america.osm.pbf">central-america.osm.pbf</a></td><td align="right">22-Mar-2018 00:52  </td><td align="right">336M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="central-america.state.txt">central-america.state.txt</a></td><td align="right">22-Mar-2018 00:49  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="central-america/">central-america/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="europe-latest.osm.pbf">europe-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:31  </td><td align="right"> 20G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="europe.osm.pbf">europe.osm.pbf</a></td><td align="right">22-Mar-2018 02:31  </td><td align="right"> 20G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="europe.state.txt">europe.state.txt</a></td><td align="right">22-Mar-2018 00:09  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="europe/">europe/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="merge/">merge/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="north-america/">north-america/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="oceania-latest.osm.pbf">oceania-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:42  </td><td align="right">570M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="oceania.osm.pbf">oceania.osm.pbf</a></td><td align="right">22-Mar-2018 02:42  </td><td align="right">570M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="oceania.state.txt">oceania.state.txt</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="oceania/">oceania/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="south-america-latest.osm.pbf">south-america-latest.osm.pbf</a></td><td align="right">22-Mar-2018 01:00  </td><td align="right">1.3G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="south-america.osm.pbf">south-america.osm.pbf</a></td><td align="right">22-Mar-2018 01:00  </td><td align="right">1.3G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="south-america.state.txt">south-america.state.txt</a></td><td align="right">22-Mar-2018 00:49  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="south-america/">south-america/</a></td><td align="right">22-Mar-2017 21:10  </td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><th colspan="5"><hr></th></tr>
	</tbody></table>
	<address>Apache/2.2.22 (Debian) Server at download.openstreetmap.fr Port 80</address>
	`

	osmfrValid2 := `
	<h1>Index of /extracts/merge</h1>
	<table><tbody><tr><th><img src="/icons/blank.gif" alt="[ICO]"></th><th><a href="?C=N;O=D">Name</a></th><th><a href="?C=M;O=A">Last modified</a></th><th><a href="?C=S;O=A">Size</a></th><th><a href="?C=D;O=A">Description</a></th></tr><tr><th colspan="5"><hr></th></tr>
	<tr><td valign="top"><img src="/icons/back.gif" alt="[DIR]"></td><td><a href="/extracts/">Parent Directory</a></td><td>&nbsp;</td><td align="right">  - </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="fiji-latest.osm.pbf">fiji-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">7.2M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="fiji.osm.pbf">fiji.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">7.2M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="fiji.state.txt">fiji.state.txt</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_metro_dom_com_nc-latest.osm.pbf">france_metro_dom_com_nc-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:56  </td><td align="right">3.6G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_metro_dom_com_nc.osm.pbf">france_metro_dom_com_nc.osm.pbf</a></td><td align="right">22-Mar-2018 02:56  </td><td align="right">3.6G</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="france_metro_dom_com_nc.state.txt">france_metro_dom_com_nc.state.txt</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_taaf-latest.osm.pbf">france_taaf-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">1.0M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_taaf.osm.pbf">france_taaf.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">1.0M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="france_taaf.state.txt">france_taaf.state.txt</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="israel_and_palestine-latest.osm.pbf">israel_and_palestine-latest.osm.pbf</a></td><td align="right">21-Mar-2017 02:37  </td><td align="right"> 60M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="israel_and_palestine.osm.pbf">israel_and_palestine.osm.pbf</a></td><td align="right">21-Mar-2017 02:37  </td><td align="right"> 60M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="israel_and_palestine.state.txt">israel_and_palestine.state.txt</a></td><td align="right">11-Sep-2017 13:52  </td><td align="right">190 </td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="kiribati-latest.osm.pbf">kiribati-latest.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">1.1M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="kiribati.osm.pbf">kiribati.osm.pbf</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right">1.1M</td><td>&nbsp;</td></tr>
	<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="kiribati.state.txt">kiribati.state.txt</a></td><td align="right">22-Mar-2018 02:38  </td><td align="right"> 86 </td><td>&nbsp;</td></tr>
	<tr><th colspan="5"><hr></th></tr>
	</tbody></table>
	<address>Apache/2.2.22 (Debian) Server at download.openstreetmap.fr Port 80</address>
	`

	var f func(s string, myUrl string) *goquery.Document
	f = func(s string, myUrl string) *goquery.Document {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
		doc.Url, err = url.Parse(myUrl)
		if err != nil {
			t.Error(err)
		}
		return doc
	}
	type fields struct {
		DefaultExtender *gocrawl.DefaultExtender // not used
		Elements        ElementSlice
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
			args:   args{doc: f(osmfrValid1, "http://download.openstreetmap.fr/extracts/")},
			want1:  true,
			want: ElementSlice{
				"africa":          {ID: "africa", Meta: false, Name: "africa", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
				"asia":            {ID: "asia", Meta: false, Name: "asia", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
				"central-america": {ID: "central-america", Meta: false, Name: "central-america", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
				"europe":          {ID: "europe", Meta: false, Name: "europe", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
				"oceania":         {ID: "oceania", Meta: false, Name: "oceania", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
				"south-america":   {ID: "south-america", Meta: false, Name: "south-america", Formats: []string{"osm.pbf", "state"}, Parent: "", File: ""},
			},
		},
		{
			name:   "osm valid2",
			fields: Ext{Elements: ElementSlice{}},
			args:   args{doc: f(osmfrValid2, "http://download.openstreetmap.fr/extracts/merge/")},
			want: ElementSlice{
				"france_metro_dom_com_nc": {ID: "france_metro_dom_com_nc", File: "", Meta: false, Name: "france_metro_dom_com_nc", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"france_taaf":             {ID: "france_taaf", File: "", Meta: false, Name: "france_taaf", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"israel_and_palestine":    {ID: "israel_and_palestine", File: "", Meta: false, Name: "israel_and_palestine", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"kiribati":                {ID: "kiribati", File: "", Meta: false, Name: "kiribati", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
				"fiji":                    {ID: "fiji", File: "", Meta: false, Name: "fiji", Formats: []string{"osm.pbf", "state"}, Parent: "merge"},
			},
			want1: true,
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
