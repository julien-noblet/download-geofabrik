package openstreetmapfr_test

import (
	"net/url"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/julien-noblet/download-geofabrik/scrapper/openstreetmapfr"
)

//nolint:lll // this func contain html extracts
func TestOpenstreetmapFR_parse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		html    string
		url     string
		want    element.MapElement
		element []element.Element
	}{
		{
			name: "sample",
			html: `<a href="fiji.osm.pbf">fiji.osm.pbf</a>`,
			url:  `https://download.openstreetmap.fr/extracts/`,
			want: element.MapElement{"fiji": element.Element{ID: "fiji", File: "fiji", Name: "fiji", Formats: element.Formats{formats.FormatOsmPbf}}},
		},
		{
			name: "Merge",
			html: `
				<table>
					<tbody><tr><th valign="top"><img src="/icons/blank.gif" alt="[ICO]"></th><th><a href="?C=N;O=D">Name</a></th><th><a href="?C=M;O=A">Last modified</a></th><th><a href="?C=S;O=A">Size</a></th><th><a href="?C=D;O=A">Description</a></th></tr>
					<tr><th colspan="5"><hr></th></tr>
					<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="fiji-latest.osm.pbf">fiji-latest.osm.pbf</a></td><td align="right">2019-05-09 08:14  </td><td align="right">7.6M</td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="fiji.osm.pbf">fiji.osm.pbf</a></td><td align="right">2019-05-09 08:14  </td><td align="right">7.6M</td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="fiji.state.txt">fiji.state.txt</a></td><td align="right">2019-05-09 08:13  </td><td align="right"> 87 </td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_metro_dom_com_nc-latest.osm.pbf">france_metro_dom_com_nc-latest.osm.pbf</a></td><td align="right">2019-05-09 08:37  </td><td align="right">3.9G</td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_metro_dom_com_nc.osm.pbf">france_metro_dom_com_nc.osm.pbf</a></td><td align="right">2019-05-09 08:37  </td><td align="right">3.9G</td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="france_metro_dom_com_nc.state.txt">france_metro_dom_com_nc.state.txt</a></td><td align="right">2019-05-09 08:15  </td><td align="right"> 87 </td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_taaf-latest.osm.pbf">france_taaf-latest.osm.pbf</a></td><td align="right">2019-05-09 08:15  </td><td align="right">1.8M</td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="france_taaf.osm.pbf">france_taaf.osm.pbf</a></td><td align="right">2019-05-09 08:15  </td><td align="right">1.8M</td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="france_taaf.state.txt">france_taaf.state.txt</a></td><td align="right">2019-05-09 08:15  </td><td align="right"> 87 </td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="kiribati-latest.osm.pbf">kiribati-latest.osm.pbf</a></td><td align="right">2019-05-09 08:14  </td><td align="right">1.2M</td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/unknown.gif" alt="[   ]"></td><td><a href="kiribati.osm.pbf">kiribati.osm.pbf</a></td><td align="right">2019-05-09 08:14  </td><td align="right">1.2M</td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/text.gif" alt="[TXT]"></td><td><a href="kiribati.state.txt">kiribati.state.txt</a></td><td align="right">2019-05-09 08:14  </td><td align="right"> 87 </td><td>&nbsp;</td></tr>
					<tr><th colspan="5"><hr></th></tr>
				</tbody></table>`,
			url: `https://download.openstreetmap.fr/extracts/merge/`,
			want: element.MapElement{
				"merge":                   element.Element{ID: "merge", Name: "merge", Meta: true},
				"fiji":                    element.Element{ID: "fiji", File: "fiji", Name: "fiji", Formats: element.Formats{formats.FormatOsmPbf, formats.FormatState}, Parent: "merge"},
				"france_metro_dom_com_nc": element.Element{ID: "france_metro_dom_com_nc", File: "france_metro_dom_com_nc", Name: "france_metro_dom_com_nc", Formats: element.Formats{formats.FormatOsmPbf, formats.FormatState}, Parent: "merge"},
				"merge_france_taaf":       element.Element{ID: "merge_france_taaf", File: "france_taaf", Name: "merge_france_taaf", Formats: element.Formats{formats.FormatOsmPbf, formats.FormatState}, Parent: "merge"},
				"kiribati":                element.Element{ID: "kiribati", File: "kiribati", Name: "kiribati", Formats: element.Formats{formats.FormatOsmPbf, formats.FormatState}, Parent: "merge"},
			},
		},
		{
			name: "cgi-bin and replication",
			html: `
				<tbody>
					<tr><th valign="top"><img src="/icons/blank.gif" alt="[ICO]"></th><th><a href="?C=N;O=D">Name</a></th><th><a href="?C=M;O=A">Last modified</a></th><th><a href="?C=S;O=A">Size</a></th><th><a href="?C=D;O=A">Description</a></th></tr>
					<tr><th colspan="5"><hr></th></tr>
					<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="cgi-bin/">cgi-bin/</a></td><td align="right">2014-07-26 18:40  </td><td align="right">  - </td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="replication/">replication/</a></td><td align="right">2014-01-17 20:12  </td><td align="right">  - </td><td>&nbsp;</td></tr>
					<tr><th colspan="5"><hr></th></tr>
				</tbody>`,
			url:  `https://download.openstreetmap.fr/`,
			want: element.MapElement{},
		},
		// TODO: Add test cases.
	}
	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(thisTest.html))
			myURL, _ := url.Parse(thisTest.url)
			myElement := &colly.HTMLElement{
				DOM: dom.Selection,
				Response: &colly.Response{
					Request: &colly.Request{URL: myURL},
				},
			}
			myOpenstreetmapFR := openstreetmapfr.GetDefault()
			myCollector := myOpenstreetmapFR.Collector() // Need a Collector to visit

			if len(thisTest.element) > 0 {
				for _, el := range thisTest.element {
					if err := myOpenstreetmapFR.Config.MergeElement(&el); err != nil {
						t.Errorf("Bad tests o.Config.mergeElement() can't merge %#v - %v", el, err)
					}
				}
			}

			myElement.ForEach("a", func(_ int, elmt *colly.HTMLElement) {
				myOpenstreetmapFR.Parse(elmt, myCollector)
			})

			if !reflect.DeepEqual(myOpenstreetmapFR.Config.Elements, thisTest.want) {
				t.Errorf("parse() fail, got %#v,\n want %#v", myOpenstreetmapFR.Config.Elements, thisTest.want)
			}
		})
	}
}

var openstreetmapFRtests = []struct { //nolint:gochecknoglobals // global
	name    string
	href    string
	c       config.Config
	want    element.MapElement
	parent  string
	parents []string
}{
	// TODO: Add test cases.
	{
		name:    "Void ",
		href:    "http://osm.fr/",
		want:    element.MapElement{},
		c:       config.Config{Elements: make(map[string]element.Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "",
		parents: []string{"http:", "", "osm.fr", ""},
	},
	{
		name: "1st level subfolder",
		href: "http://osm.fr/one/",
		want: element.MapElement{
			"one": element.Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
		},
		c:       config.Config{Elements: make(map[string]element.Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "one", ""},
	},
	{
		name: "2nd level subfolder",
		href: "http://osm.fr/one/two/",
		want: element.MapElement{
			"one": element.Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
			"two": element.Element{
				ID:     "two",
				Meta:   true,
				Name:   "two",
				Parent: "one",
			},
		},
		c:       config.Config{Elements: make(map[string]element.Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "two",
		parents: []string{"http:", "", "osm.fr", "one", "two", ""},
	},
	{
		name: "top level extracts osm.pbf",
		href: "http://osm.fr/extracts/object.osm.pbf",
		want: element.MapElement{
			"object": element.Element{
				ID:      "object",
				File:    "object",
				Name:    "object",
				Formats: []string{formats.FormatOsmPbf},
			},
		},
		c:       config.Config{Elements: make(map[string]element.Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "",
		parents: []string{"http:", "", "osm.fr", "extracts", "object.osm.pbf"},
	},
	{
		name: "top level polygons osm.pbf",
		href: "http://osm.fr/polygons/object.osm.pbf",
		want: element.MapElement{
			"object": element.Element{
				ID:      "object",
				File:    "object",
				Name:    "object",
				Formats: []string{formats.FormatOsmPbf},
			},
		},
		c:       config.Config{Elements: make(map[string]element.Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "",
		parents: []string{"http:", "", "osm.fr", "polygons", "object.osm.pbf"},
	},
	{
		name: "sub level extracts osm.pbf",
		href: "http://osm.fr/extracts/one/object.osm.pbf",
		want: element.MapElement{
			"object": element.Element{
				ID:      "object",
				File:    "object",
				Name:    "object",
				Formats: []string{formats.FormatOsmPbf},
				Parent:  "one",
			},
			"one": element.Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
		},
		c:       config.Config{Elements: make(map[string]element.Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "extracts", "one", "object.osm.pbf"},
	},
	{
		name: "sub level polygons osm.pbf",
		href: "http://osm.fr/polygons/one/object.osm.pbf",
		want: element.MapElement{
			"object": element.Element{
				ID:      "object",
				File:    "object",
				Name:    "object",
				Meta:    false,
				Formats: []string{formats.FormatOsmPbf},
				Parent:  "one",
			},
			"one": element.Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
		},
		c:       config.Config{Elements: make(map[string]element.Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "polygons", "one", "object.osm.pbf"},
	},
	{
		name: "2nd level osm.pbf",
		href: "http://osm.fr/polygons/one/two/object.osm.pbf",
		want: element.MapElement{
			"object": element.Element{
				ID:      "object",
				File:    "object",
				Name:    "object",
				Formats: []string{formats.FormatOsmPbf},
				Parent:  "two",
			},
			"one": element.Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
			"two": element.Element{
				ID:     "two",
				Meta:   true,
				Name:   "two",
				Parent: "one",
			},
		},
		c:       config.Config{Elements: make(map[string]element.Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "two",
		parents: []string{"http:", "", "osm.fr", "polygons", "one", "two", "object.osm.pbf"},
	},
	{
		name: "sub level extracts osm.pbf + state",
		href: "http://osm.fr/extracts/one/object.state.txt",
		want: element.MapElement{
			"object": element.Element{
				ID:      "object",
				File:    "object",
				Name:    "object",
				Formats: []string{formats.FormatOsmPbf, formats.FormatState},
				Parent:  "one",
			},
			"one": element.Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
		},
		c: config.Config{
			Elements: element.MapElement{
				"object": element.Element{
					ID:      "object",
					File:    "object",
					Name:    "object",
					Formats: []string{formats.FormatOsmPbf},
					Parent:  "one",
				},
			},
			ElementsMutex: &sync.RWMutex{},
		},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "extracts", "one", "object.state.txt"},
	},
}

func TestOpenstreetmapFR_parseHref(t *testing.T) {
	t.Parallel()

	for thisTest := range openstreetmapFRtests {
		thisTest := thisTest
		myOpenstreetmapFR := openstreetmapfr.GetDefault()
		myOpenstreetmapFR.Config = &openstreetmapFRtests[thisTest].c
		t.Run(openstreetmapFRtests[thisTest].name, func(t *testing.T) {
			t.Parallel()

			myOpenstreetmapFR.ParseHref(openstreetmapFRtests[thisTest].href)

			if !reflect.DeepEqual(openstreetmapFRtests[thisTest].c.Elements, openstreetmapFRtests[thisTest].want) {
				t.Errorf(
					"parseHref() = \n%#v len:%d, want \n%#v len:%d\n",
					openstreetmapFRtests[thisTest].c.Elements,
					len(openstreetmapFRtests[thisTest].c.Elements),
					openstreetmapFRtests[thisTest].want,
					len(openstreetmapFRtests[thisTest].want),
				)
			}
		})
	}
}

func Test_openstreetmapFRGetParent(t *testing.T) {
	t.Parallel()

	for thisTest := range openstreetmapFRtests {
		ThisTest := thisTest
		t.Run(openstreetmapFRtests[ThisTest].name, func(t *testing.T) {
			t.Parallel()

			p, grandParent := openstreetmapfr.GetParent(openstreetmapFRtests[ThisTest].href)
			if !reflect.DeepEqual(p, openstreetmapFRtests[ThisTest].parent) {
				t.Errorf("openstreetmapFRGetParent() = %v want %v ", p, openstreetmapFRtests[ThisTest].parent)
			}

			if !reflect.DeepEqual(grandParent, openstreetmapFRtests[ThisTest].parents) {
				t.Errorf("openstreetmapFRGetParent() = %v want %v ", grandParent, openstreetmapFRtests[ThisTest].parents)
			}
		})
	}
}

func TestOpenstreetmapFR_makeParents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		elements element.MapElement
		want     element.MapElement
		name     string
		parent   string
		gparents []string
	}{
		// TODO: Add test cases.
		{
			name: "1 fake", // should not append!!!
			// TODO: Raise an error?
			parent:   "toto",
			gparents: []string{"toto", "some.osm.pbf"},
			want:     element.MapElement{"toto": element.Element{ID: "toto", Name: "toto", Meta: true}},
		},
		{
			name:     "1 parent",
			parent:   "toto",
			gparents: []string{"http:", "", "osm.fr", "extracts", "toto", "some.osm.pbf"},
			want:     element.MapElement{"toto": element.Element{ID: "toto", Name: "toto", Meta: true}},
		},
		{
			name:     "2 parents",
			parent:   "toto",
			gparents: []string{"http:", "", "osm.fr", "extracts", "tata", "toto", "some.osm.pbf"},
			want: element.MapElement{
				"toto": element.Element{ID: "toto", Name: "toto", Meta: true, Parent: "tata"},
				"tata": element.Element{ID: "tata", Name: "tata", Meta: true},
			},
		},
		{
			name:     "3 parents",
			parent:   "toto",
			gparents: []string{"http:", "", "osm.fr", "extracts", "tata", "titi", "toto", "some.osm.pbf"},
			want: element.MapElement{
				"toto": element.Element{ID: "toto", Name: "toto", Meta: true, Parent: "titi"},
				"titi": element.Element{ID: "titi", Name: "titi", Meta: true, Parent: "tata"},
				"tata": element.Element{ID: "tata", Name: "tata", Meta: true},
			},
		},
		{
			name:     "one parent",
			parent:   "one",
			gparents: []string{"http:", "", "osm.fr", "extracts", "one", "object.osm.pbf"},
			want: element.MapElement{
				"one": element.Element{ID: "one", Name: "one", Meta: true},
			},
		},
	}
	for _, thisTest := range tests {
		thisTest := thisTest
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			myOpenstreetmapFR := openstreetmapfr.GetDefault()
			myOpenstreetmapFR.GetConfig()

			if len(thisTest.elements) > 0 {
				for _, e := range thisTest.elements {
					if err := myOpenstreetmapFR.Config.MergeElement(&e); err != nil {
						t.Error(err)
					}
				}
			}

			myOpenstreetmapFR.MakeParents(thisTest.parent, thisTest.gparents)

			if !reflect.DeepEqual(myOpenstreetmapFR.Config.Elements, thisTest.want) {
				t.Errorf("makeParent() fail: got\n %#v,\n want \n %#v.\n", myOpenstreetmapFR.Config.Elements, thisTest.want)
			}
		})
	}
}

func BenchmarkExeptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		openstreetmapfr.Exceptions("north-east", "parent")
		openstreetmapfr.Exceptions("central", "parent")
		openstreetmapfr.Exceptions("unknown", "parent")
	}
}
