package osmtoday_test

import (
	"reflect"
	"testing"

	"github.com/julien-noblet/download-geofabrik/element"
	"github.com/julien-noblet/download-geofabrik/scrapper/osmtoday"
)

func TestOsmtoday_Exceptions(t *testing.T) {
	type args struct {
		e element.Element
	}

	tests := []struct {
		name string
		args args
		want element.Element
	}{
		// TODO: Add test cases.
		{
			name: "la_rioja in argentina",
			args: args{
				e: element.Element{
					ID:     "la_rioja",
					Parent: "argentina",
					File:   "la_rioja",
				},
			},
			want: element.Element{
				ID:     "la_rioja-argentina",
				Parent: "argentina",
				File:   "la_rioja",
			},
		},
		{
			name: "la_rioja in spain",
			args: args{
				e: element.Element{
					ID:     "la_rioja",
					Parent: "spain",
					File:   "la_rioja",
				},
			},
			want: element.Element{
				ID:     "la_rioja-spain",
				Parent: "spain",
				File:   "la_rioja",
			},
		},
		{
			name: "guyane in france",
			args: args{
				e: element.Element{
					ID:     "guyane",
					Parent: "france",
					File:   "guyane",
				},
			},
			want: element.Element{
				ID:     "guyane-france",
				Parent: "france",
				File:   "guyane",
			},
		},
		{
			name: "guyane in south-america",
			args: args{
				e: element.Element{
					ID:     "guyane",
					Parent: "south-america",
					File:   "guyane",
				},
			},
			want: element.Element{
				ID:     "guyane-south-america",
				Parent: "south-america",
				File:   "guyane",
			},
		},
		{
			name: "sevastopol in ukraine",
			args: args{
				e: element.Element{
					ID:     "sevastopol",
					Parent: "ukraine",
					File:   "sevastopol",
				},
			},
			want: element.Element{
				ID:     "sevastopol-ukraine",
				Parent: "ukraine",
				File:   "sevastopol",
			},
		},
		{
			name: "sevastopol in russia",
			args: args{
				e: element.Element{
					ID:     "sevastopol",
					Parent: "russia",
					File:   "sevastopol",
				},
			},
			want: element.Element{
				ID:     "sevastopol-russia",
				Parent: "russia",
				File:   "sevastopol",
			},
		},
		{
			name: "limburg in netherlands",
			args: args{
				e: element.Element{
					ID:     "limburg",
					Parent: "netherlands",
					File:   "limburg",
				},
			},
			want: element.Element{
				ID:     "limburg-netherlands",
				Parent: "netherlands",
				File:   "limburg",
			},
		},
		{
			name: "limburg in flanders",
			args: args{
				e: element.Element{
					ID:     "limburg",
					Parent: "flanders",
					File:   "limburg",
				},
			},
			want: element.Element{
				ID:     "limburg-flanders",
				Parent: "flanders",
				File:   "limburg",
			},
		},
		{
			name: "cordoba in argentina",
			args: args{
				e: element.Element{
					ID:     "cordoba",
					Parent: "argentina",
					File:   "cordoba",
				},
			},
			want: element.Element{
				ID:     "cordoba-argentina",
				Parent: "argentina",
				File:   "cordoba",
			},
		},
		{
			name: "cordoba in andalucia",
			args: args{
				e: element.Element{
					ID:     "cordoba",
					Parent: "andalucia",
					File:   "cordoba",
				},
			},
			want: element.Element{
				ID:     "cordoba-andalucia",
				Parent: "andalucia",
				File:   "cordoba",
			},
		},
		{
			name: "georgia in usa",
			args: args{
				e: element.Element{
					ID:     "georgia",
					Parent: "us",
					File:   "georgia",
				},
			},
			want: element.Element{
				ID:     "georgia-us",
				Parent: "us",
				File:   "georgia",
			},
		},
		{
			name: "georgia in asia",
			args: args{
				e: element.Element{
					ID:     "georgia",
					Parent: "asia",
					File:   "georgia",
				},
			},
			want: element.Element{
				ID:     "georgia-asia",
				Parent: "asia",
				File:   "georgia",
			},
		},
		{
			name: "france is not in the list",
			args: args{
				e: element.Element{
					ID:     "france",
					Parent: "europe",
					File:   "france",
				},
			},
			want: element.Element{
				ID:     "france",
				Parent: "europe",
				File:   "france",
			},
		},
	}

	for _, tt := range tests {
		test := tt // Reinitialize tt inside the range statement
		t.Run(test.name, func(t *testing.T) {
			g := osmtoday.GetDefault()

			if got := g.Exceptions(&test.args.e); !reflect.DeepEqual(got, &test.want) {
				t.Errorf("Osmtoday.Exceptions() = %v, want %v", *got, test.want)
			}
		})
	}
}

/*
func gWithOsmPbf() *osmtoday.Osmtoday {
	gWithOsmPbf := osmtoday.GetDefault()
	gWithOsmPbf.Config = &config.Config{
		Formats: formats.FormatDefinitions{
			formats.FormatOsmPbf: {ID: formats.FormatOsmPbf, Loc: ".pbf", BasePath: "", BaseURL: ""},
		},
	}

	return gWithOsmPbf
}

func gWithPoly() *osmtoday.Osmtoday {
	gWithPoly := osmtoday.GetDefault()
	gWithPoly.Config.Formats = formats.FormatDefinitions{
		formats.FormatPoly: {ID: formats.FormatPoly, Loc: ".poly", BasePath: "", BaseURL: ""},
	}

	return gWithPoly
}
*/
