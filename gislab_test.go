package main

import (
	"reflect"
	"testing"
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
		in   string
		want ElementSlice
	}{
		// TODO: Add test cases.
		{name: "Void 1 tr 1 td", in: `<table>
		<tbody>
			<tr>
			<td></td>
			</tr>
		</tbody>
		</table>`, want: ElementSlice{}},
		{name: "Void 1 tr 1 td", in: `<table>
		<tbody>
			<tr>
			<th></th>
			</tr>
		</tbody>
		</table>`, want: ElementSlice{}},
		{name: "Void header only", in: `<table>
		<tbody>
			<tr>
				<th>ISO&nbsp;3166</th>
				<th>Страна/регион</th>
				<th>osm.pbf</th>
				<th>osm.bz2</th>
				<th colspan="2">&nbsp;</th>
			</tr>
		</tbody>
		</table>`, want: ElementSlice{}},
		{name: "1 elmt with Header", in: `<table>
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
		</table>`, want: ElementSlice{"AM": Element{ID: "AM", Name: "Армения", Formats: []string{"osm.pbf", "osm.bz2", "poly"}}}},
	}
	for _, tt := range tests {
		element := createHTMLElement(t, tt.in)
		e := Ext{Elements: make(map[string]Element)}
		t.Run(tt.name, func(t *testing.T) {
			gislabParse(element, &e)
			if !reflect.DeepEqual(e.Elements, tt.want) {
				t.Errorf("gislabParse() = %v len:%d, want %v len:%d", e.Elements, len(e.Elements), tt.want, len(tt.want))
			}
		})
	}
}
