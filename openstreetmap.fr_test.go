package main

import (
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func TestOpenstreetmapFR_parse(t *testing.T) {
	tests := []struct {
		name    string
		html    string
		url     string
		want    ElementSlice
		element []Element
	}{
		{name: "sample",
			html: `<a href="fiji.osm.pbf">fiji.osm.pbf</a>`,
			url:  `https://download.openstreetmap.fr/extracts/`,
			want: ElementSlice{"fiji": Element{ID: "fiji", Name: "fiji", Formats: elementFormats{"osm.pbf"}}},
		},
		{name: "Merge",
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
			want: ElementSlice{
				"merge":                   Element{ID: "merge", Name: "merge", Meta: true},
				"fiji":                    Element{ID: "fiji", Name: "fiji", Formats: elementFormats{"osm.pbf", "state"}, Parent: "merge"},
				"france_metro_dom_com_nc": Element{ID: "france_metro_dom_com_nc", Name: "france_metro_dom_com_nc", Formats: elementFormats{"osm.pbf", "state"}, Parent: "merge"},
				"france_taaf":             Element{ID: "france_taaf", Name: "france_taaf", Formats: elementFormats{"osm.pbf", "state"}, Parent: "merge"},
				"kiribati":                Element{ID: "kiribati", Name: "kiribati", Formats: elementFormats{"osm.pbf", "state"}, Parent: "merge"}},
		},
		{name: "cgi-bin and replication",
			html: `
				<tbody>
					<tr><th valign="top"><img src="/icons/blank.gif" alt="[ICO]"></th><th><a href="?C=N;O=D">Name</a></th><th><a href="?C=M;O=A">Last modified</a></th><th><a href="?C=S;O=A">Size</a></th><th><a href="?C=D;O=A">Description</a></th></tr>
					<tr><th colspan="5"><hr></th></tr>
					<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="cgi-bin/">cgi-bin/</a></td><td align="right">2014-07-26 18:40  </td><td align="right">  - </td><td>&nbsp;</td></tr>
					<tr><td valign="top"><img src="/icons/folder.gif" alt="[DIR]"></td><td><a href="replication/">replication/</a></td><td align="right">2014-01-17 20:12  </td><td align="right">  - </td><td>&nbsp;</td></tr>
					<tr><th colspan="5"><hr></th></tr>
				</tbody>`,
			url:  `https://download.openstreetmap.fr/`,
			want: ElementSlice{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(tt.html))
			url, _ := url.Parse(tt.url)
			e := &colly.HTMLElement{
				DOM: dom.Selection,
				Response: &colly.Response{
					Request: &colly.Request{URL: url},
				},
			}
			o := OpenstreetmapFR{
				Scrapper: &Scrapper{
					PB:             144,
					Async:          true,
					Parallelism:    20,
					MaxDepth:       0,
					AllowedDomains: []string{`download.openstreetmap.fr`},
					BaseURL:        `https://download.openstreetmap.fr/extracts`,
					StartURL:       `https://download.openstreetmap.fr/`,
					// URLFilters: []*regexp.Regexp{
					// regexp.MustCompile(`https://download.openstreetmap.fr/([^cgi\-bin][^replication]\w.+|)`),
					// },
					URLFilters: []*regexp.Regexp{
						//regexp.MustCompile(`https://download\.openstreetmap\.fr/([extracts|polygons]\w.+|)`),
						regexp.MustCompile(`https://download\.openstreetmap\.fr/$`),
						regexp.MustCompile(`https://download\.openstreetmap\.fr/extracts/(\w.+|)$`),
						regexp.MustCompile(`https://download\.openstreetmap\.fr/polygons/(\w.+|)$`),
						regexp.MustCompile(`https://download.openstreetmap.fr/cgi-bin/^(.*)$`),
						regexp.MustCompile(`https://download.openstreetmap.fr/replication/^(.*|)$`),
					},
					FormatDefinition: formatDefinitions{
						"osm.pbf": {ID: "osm.pbf", Loc: "-latest.osm.pbf"},
						"poly":    {ID: "poly", Loc: ".poly", BasePath: "../polygons/"},
						"state":   {ID: "state", Loc: ".state.txt"},
					},
				},
			}
			c := o.Collector() // Need a Collector to visit

			if len(tt.element) > 0 {
				for _, el := range tt.element {
					err := o.Config.mergeElement(&el)
					if err != nil {
						t.Errorf("Bad tests o.Config.mergeElement() can't merge %#v - %v", el, err)
					}
				}
			}
			//fmt.Printf("%#v", e)
			e.ForEach("a", func(_ int, elmt *colly.HTMLElement) {
				o.parse(elmt, c)
			})

			if !reflect.DeepEqual(o.Config.Elements, tt.want) {
				t.Errorf("parse() fail, got %#v,\n want %#v", o.Config.Elements, tt.want)
			}

		})
	}
}

var openstreetmapFRtests = []struct {
	name    string
	href    string
	c       Config
	want    ElementSlice
	parent  string
	parents []string
}{
	// TODO: Add test cases.
	{name: "Void ",
		href:    "http://osm.fr/",
		want:    ElementSlice{},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "",
		parents: []string{"http:", "", "osm.fr", ""},
	},
	{name: "1st level subfolder",
		href: "http://osm.fr/one/",
		want: ElementSlice{
			"one": Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
		},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "one", ""},
	},
	{name: "2nd level subfolder",
		href: "http://osm.fr/one/two/",
		want: ElementSlice{
			"one": Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
			"two": Element{
				ID:     "two",
				Meta:   true,
				Name:   "two",
				Parent: "one",
			},
		},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "two",
		parents: []string{"http:", "", "osm.fr", "one", "two", ""},
	},
	{name: "top level extracts osm.pbf",
		href: "http://osm.fr/extracts/object.osm.pbf",
		want: ElementSlice{
			"object": Element{
				ID:      "object",
				Name:    "object",
				Formats: []string{"osm.pbf"},
			},
		},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "",
		parents: []string{"http:", "", "osm.fr", "extracts", "object.osm.pbf"},
	},
	{name: "top level polygons osm.pbf",
		href: "http://osm.fr/polygons/object.osm.pbf",
		want: ElementSlice{
			"object": Element{
				ID:      "object",
				Name:    "object",
				Formats: []string{"osm.pbf"},
			},
		},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "",
		parents: []string{"http:", "", "osm.fr", "polygons", "object.osm.pbf"},
	},
	{name: "sub level extracts osm.pbf",
		href: "http://osm.fr/extracts/one/object.osm.pbf",
		want: ElementSlice{
			"object": Element{
				ID:      "object",
				Name:    "object",
				Formats: []string{"osm.pbf"},
				Parent:  "one",
			},
			"one": Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
		},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "extracts", "one", "object.osm.pbf"},
	},
	{name: "sub level polygons osm.pbf",
		href: "http://osm.fr/polygons/one/object.osm.pbf",
		want: ElementSlice{
			"object": Element{
				ID:      "object",
				Name:    "object",
				Meta:    false,
				Formats: []string{"osm.pbf"},
				Parent:  "one",
			},
			"one": Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
		},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "one",
		parents: []string{"http:", "", "osm.fr", "polygons", "one", "object.osm.pbf"},
	},
	{name: "2nd level osm.pbf",
		href: "http://osm.fr/polygons/one/two/object.osm.pbf",
		want: ElementSlice{
			"object": Element{
				ID:      "object",
				Name:    "object",
				Formats: []string{"osm.pbf"},
				Parent:  "two",
			},
			"one": Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
			"two": Element{
				ID:     "two",
				Meta:   true,
				Name:   "two",
				Parent: "one",
			},
		},
		c:       Config{Elements: make(map[string]Element), ElementsMutex: &sync.RWMutex{}},
		parent:  "two",
		parents: []string{"http:", "", "osm.fr", "polygons", "one", "two", "object.osm.pbf"},
	},
	{name: "sub level extracts osm.pbf + state",
		href: "http://osm.fr/extracts/one/object.state.txt",
		want: ElementSlice{
			"object": Element{
				ID:      "object",
				Name:    "object",
				Formats: []string{"osm.pbf", "state"},
				Parent:  "one",
			},
			"one": Element{
				ID:   "one",
				Meta: true,
				Name: "one",
			},
		},
		c: Config{
			Elements: ElementSlice{
				"object": Element{
					ID:      "object",
					Name:    "object",
					Formats: []string{"osm.pbf"},
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
	for tn := range openstreetmapFRtests {
		o := OpenstreetmapFR{
			Scrapper: &Scrapper{
				PB:             144,
				Async:          true,
				Parallelism:    20,
				MaxDepth:       0,
				AllowedDomains: []string{`download.openstreetmap.fr`},
				BaseURL:        `https://download.openstreetmap.fr/extracts`,
				StartURL:       `https://download.openstreetmap.fr/`,
				// URLFilters: []*regexp.Regexp{
				// 	regexp.MustCompile(`https://download.openstreetmap.fr/([^cgi\-bin][^replication]\w.+|)`),
				// },
				URLFilters: []*regexp.Regexp{
					//regexp.MustCompile(`https://download\.openstreetmap\.fr/([extracts|polygons]\w.+|)`),
					regexp.MustCompile(`https://download\.openstreetmap\.fr/$`),
					regexp.MustCompile(`https://download\.openstreetmap\.fr/extracts/(\w.+|)$`),
					regexp.MustCompile(`https://download\.openstreetmap\.fr/polygons/(\w.+|)$`),
					regexp.MustCompile(`https://download.openstreetmap.fr/cgi-bin/^(.*)$`),
					regexp.MustCompile(`https://download.openstreetmap.fr/replication/^(.*|)$`),
				},
				FormatDefinition: formatDefinitions{
					"osm.pbf": {ID: "osm.pbf", Loc: "-latest.osm.pbf"},
					"poly":    {ID: "poly", Loc: ".poly", BasePath: "../polygons/"},
					"state":   {ID: "state", Loc: ".state.txt"},
				},
			},
		}
		o.Config = &openstreetmapFRtests[tn].c
		t.Run(openstreetmapFRtests[tn].name, func(t *testing.T) {
			o.parseHref(openstreetmapFRtests[tn].href)
			if !reflect.DeepEqual(openstreetmapFRtests[tn].c.Elements, openstreetmapFRtests[tn].want) {
				t.Errorf("parseHref() = \n%#v len:%d, want \n%#v len:%d\n", openstreetmapFRtests[tn].c.Elements, len(openstreetmapFRtests[tn].c.Elements), openstreetmapFRtests[tn].want, len(openstreetmapFRtests[tn].want))
			}
		})
	}
}

func Test_openstreetmapFRGetParent(t *testing.T) {
	for tn := range openstreetmapFRtests {
		t.Run(openstreetmapFRtests[tn].name, func(t *testing.T) {
			p, ps := openstreetmapFRGetParent(openstreetmapFRtests[tn].href)
			if !reflect.DeepEqual(p, openstreetmapFRtests[tn].parent) {
				t.Errorf("openstreetmapFRGetParent() = %v want %v ", p, openstreetmapFRtests[tn].parent)
			}
			if !reflect.DeepEqual(ps, openstreetmapFRtests[tn].parents) {
				t.Errorf("openstreetmapFRGetParent() = %v want %v ", ps, openstreetmapFRtests[tn].parents)
			}

		})
	}
}

func TestOpenstreetmapFR_makeParents(t *testing.T) {

	tests := []struct {
		name     string
		elements ElementSlice
		parent   string
		gparents []string
		want     ElementSlice
	}{
		// TODO: Add test cases.
		{name: "1 fake", // should not append!!!
			// TODO: Raise an error?
			parent:   "toto",
			gparents: []string{"toto", "some.osm.pbf"},
			want:     ElementSlice{"toto": Element{ID: "toto", Name: "toto", Meta: true}},
		},
		{name: "1 parent",
			parent:   "toto",
			gparents: []string{"http:", "", "osm.fr", "extracts", "toto", "some.osm.pbf"},
			want:     ElementSlice{"toto": Element{ID: "toto", Name: "toto", Meta: true}},
		},
		{name: "2 parents",
			parent:   "toto",
			gparents: []string{"http:", "", "osm.fr", "extracts", "tata", "toto", "some.osm.pbf"},
			want: ElementSlice{
				"toto": Element{ID: "toto", Name: "toto", Meta: true, Parent: "tata"},
				"tata": Element{ID: "tata", Name: "tata", Meta: true},
			},
		},
		{name: "3 parents",
			parent:   "toto",
			gparents: []string{"http:", "", "osm.fr", "extracts", "tata", "titi", "toto", "some.osm.pbf"},
			want: ElementSlice{
				"toto": Element{ID: "toto", Name: "toto", Meta: true, Parent: "titi"},
				"titi": Element{ID: "titi", Name: "titi", Meta: true, Parent: "tata"},
				"tata": Element{ID: "tata", Name: "tata", Meta: true},
			},
		},
		{name: "one parent",
			parent:   "one",
			gparents: []string{"http:", "", "osm.fr", "extracts", "one", "object.osm.pbf"},
			want: ElementSlice{
				"one": Element{ID: "one", Name: "one", Meta: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OpenstreetmapFR{
				Scrapper: &Scrapper{
					PB:             144,
					Async:          true,
					Parallelism:    20,
					MaxDepth:       0,
					AllowedDomains: []string{`download.openstreetmap.fr`},
					BaseURL:        `https://download.openstreetmap.fr/extracts`,
					StartURL:       `https://download.openstreetmap.fr/`,
					// URLFilters: []*regexp.Regexp{
					// 	regexp.MustCompile(`https://download.openstreetmap.fr/([^cgi\-bin][^replication]\w.+|)`),
					// },
					URLFilters: []*regexp.Regexp{
						//regexp.MustCompile(`https://download\.openstreetmap\.fr/([extracts|polygons]\w.+|)`),
						regexp.MustCompile(`https://download\.openstreetmap\.fr/$`),
						regexp.MustCompile(`https://download\.openstreetmap\.fr/extracts/(\w.+|)$`),
						regexp.MustCompile(`https://download\.openstreetmap\.fr/polygons/(\w.+|)$`),
						regexp.MustCompile(`https://download.openstreetmap.fr/cgi-bin/^(.*)$`),
						regexp.MustCompile(`https://download.openstreetmap.fr/replication/^(.*|)$`),
					},
					FormatDefinition: formatDefinitions{
						"osm.pbf": {ID: "osm.pbf", Loc: "-latest.osm.pbf"},
						"poly":    {ID: "poly", Loc: ".poly", BasePath: "../polygons/"},
						"state":   {ID: "state", Loc: ".state.txt"},
					},
				},
			}
			o.GetConfig()
			if len(tt.elements) > 0 {
				for _, e := range tt.elements {
					err := o.Config.mergeElement(&e)
					if err != nil {
						t.Error(err)
					}
				}
			}
			o.makeParents(tt.parent, tt.gparents)
			if !reflect.DeepEqual(o.Config.Elements, tt.want) {
				t.Errorf("makeParent() fail: got\n %#v,\n want \n %#v.\n", o.Config.Elements, tt.want)
			}
		})
	}
}
