package main

import (
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func Test_gislabAddExt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Element
	}{
		// TODO: Add test cases.
		{name: "Unknown ext", in: `<td><a href="http://test.org/test.xyz">Test.xyz</a></td>`, want: Element{}},
		{name: "osm.pbf ext", in: `<td><a href="http://test.org/test.osm.pbf">test.osm.pbf</a></td>`, want: Element{Formats: []string{"osm.pbf"}}},
		{name: "osm.bz2 ext", in: `<td><a href="http://test.org/test.osm.bz2">test.osm.bz2</a></td>`, want: Element{Formats: []string{"osm.bz2"}}},
	}
	for _, tt := range tests {
		element := createHTMLElement(t, tt.in)
		e := Element{}
		t.Run(tt.name, func(t *testing.T) {
			gislabAddExt(element, &e)
			if !reflect.DeepEqual(e, tt.want) {
				t.Errorf("gislabAddExt() = %v, want %v", e, tt.want)
			}
		})
	}
}

func Test_gislabParse(t *testing.T) {
	tests := []struct {
		name string
		html string
		want ElementSlice
	}{
		// TODO: Add test cases.
		{name: "Void 1 tr 1 td", html: `<table>
			<tbody>
				<tr>
				<td></td>
				</tr>
			</tbody>
			</table>`,
			want: ElementSlice{}},
		{name: "Void 1 tr 1 td",
			html: `<table>
			<tbody>
				<tr>
				<th></th>
				</tr>
			</tbody>
			</table>`,
			want: ElementSlice{}},
		{name: "Void header only",
			html: `<table>
			<tbody>
				<tr>
					<th>ISO&nbsp;3166</th>
					<th>Страна/регион</th>
					<th>osm.pbf</th>
					<th>osm.bz2</th>
					<th colspan="2">&nbsp;</th>
				</tr>
			</tbody>
			</table>`,
			want: ElementSlice{}},
		{name: "1 elmt with Header",
			html: `<table>
			<tbody>
				<tr>
					<th>ISO&nbsp;3166</th>
					<th>Страна/регион</th>
					<th>osm.pbf</th>
					<th>osm.bz2</th>
					<th colspan="2">&nbsp;</th>
				</tr>
				<tr>
					<td>AM</td>
					<td>Армения</td>
					<td style="white-space:nowrap;"><a href="/AM.osm.pbf">1 Jan (1 MB)</a></td>
					<td style="white-space:nowrap;"><a href="AM.osm.bz2">1 Jan (1 MB)</a></td>
					<td><a href="/AM/" target="_blank">архив</a></td>
					<td><a href="/diff/AM/" target="_blank">обновления</a></td>
				</tr>
			</tbody>
			</table>`,
			want: ElementSlice{"AM": Element{ID: "AM", Name: "Армения", Formats: []string{"osm.pbf", "osm.bz2", "poly"}}}},
	}
	for _, tt := range tests {

		g := Gislab{
			Scrapper: &Scrapper{
				PB:             1,
				Async:          true,
				Parallelism:    1,
				MaxDepth:       0,
				AllowedDomains: []string{`be.gis-lab.info`, `data.gis-lab.info`, `raw.githubusercontent.com`},
				BaseURL:        `http://be.gis-lab.info/project/osm_dump`,
				StartURL:       `http://be.gis-lab.info/project/osm_dump/iframe.php`,
				URLFilters: []*regexp.Regexp{
					regexp.MustCompile(`http[s]?://be\.gis-lab\.info/project/osm_dump/iframe\.php$`), // Only 1 page!
				},
				FormatDefinition: formatDefinitions{
					"osm.pbf": format{ID: "osm.pbf", BaseURL: "http://data.gis-lab.info/osm_dump/dump", BasePath: "latest/", Loc: ".osm.pbf"},
					"osm.bz2": format{ID: "osm.bz2", BaseURL: "http://data.gis-lab.info/osm_dump/dump", BasePath: "latest/", Loc: ".osm.bz2"},
					"poly":    format{ID: "poly", BaseURL: "https://raw.githubusercontent.com/nextgis/osmdump_poly/master", Loc: ".poly"},
				},
			},
		}
		dom, _ := goquery.NewDocumentFromReader(strings.NewReader(tt.html))
		url, _ := url.Parse(g.StartURL)
		e := &colly.HTMLElement{
			DOM: dom.Selection,
			Response: &colly.Response{
				Request: &colly.Request{URL: url},
			},
			Request: &colly.Request{URL: url},
		}
		c := g.Collector()
		t.Run(tt.name, func(t *testing.T) {
			if err := g.parse(e, c); err != nil {
				t.Errorf("gislabParse crash : %v", err)
			}
			if !reflect.DeepEqual(g.Config.Elements, tt.want) {
				t.Errorf("gislabParse() = %#v len:%d, want %#v len:%d", g.Config.Elements, len(g.Config.Elements), tt.want, len(tt.want))
			}
		})
	}
}
