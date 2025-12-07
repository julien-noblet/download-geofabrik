package bbbike_test

import (
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/julien-noblet/download-geofabrik/internal/element"
	"github.com/julien-noblet/download-geofabrik/internal/scrapper/bbbike"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
)

func Test_bbbike_getName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		h3   string
		want string
	}{
		// TODO: Add test cases.
		{name: "Valid", h3: "OSM extracts for Test", want: "Test"},
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			if got := bbbike.GetName(thisTest.h3); got != thisTest.want {
				t.Errorf("bbbikeGetName() = %v, want %v", got, thisTest.want)
			}
		})
	}
}

//nolint:lll // this func contain html extracts
func Test_bbbike_parseList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		elements *element.MapElement
		want     element.MapElement
		name     string
		html     string
		url      string
	}{
		{
			name: "sample",
			html: `
			<tbody>
				<tr class="d"><td class="n"><a href="../">..</a>/</td><td class="m">&nbsp;</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Aachen/">Aachen</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Aarhus/">Aarhus</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Adelaide/">Adelaide</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Albuquerque/">Albuquerque</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Alexandria/">Alexandria</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Amsterdam/">Amsterdam</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Antwerpen/">Antwerpen</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Arnhem/">Arnhem</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Auckland/">Auckland</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Augsburg/">Augsburg</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Austin/">Austin</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Baghdad/">Baghdad</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Baku/">Baku</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Balaton/">Balaton</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bamberg/">Bamberg</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bangkok/">Bangkok</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Barcelona/">Barcelona</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Basel/">Basel</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Beijing/">Beijing</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Beirut/">Beirut</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Berkeley/">Berkeley</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Berlin/">Berlin</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bern/">Bern</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bielefeld/">Bielefeld</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Birmingham/">Birmingham</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bochum/">Bochum</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bogota/">Bogota</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bombay/">Bombay</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bonn/">Bonn</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bordeaux/">Bordeaux</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Boulder/">Boulder</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="BrandenburgHavel/">BrandenburgHavel</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Braunschweig/">Braunschweig</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bremen/">Bremen</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bremerhaven/">Bremerhaven</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Brisbane/">Brisbane</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bristol/">Bristol</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Brno/">Brno</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bruegge/">Bruegge</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Bruessel/">Bruessel</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Budapest/">Budapest</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="BuenosAires/">BuenosAires</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Cairo/">Cairo</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Calgary/">Calgary</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Cambridge/">Cambridge</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="CambridgeMa/">CambridgeMa</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Canberra/">Canberra</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="CapeTown/">CapeTown</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Chemnitz/">Chemnitz</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Chicago/">Chicago</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="ClermontFerrand/">ClermontFerrand</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Colmar/">Colmar</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Copenhagen/">Copenhagen</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Cork/">Cork</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Corsica/">Corsica</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Corvallis/">Corvallis</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Cottbus/">Cottbus</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Cracow/">Cracow</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="CraterLake/">CraterLake</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Curitiba/">Curitiba</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Cusco/">Cusco</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Dallas/">Dallas</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Darmstadt/">Darmstadt</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Davis/">Davis</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="DenHaag/">DenHaag</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Denver/">Denver</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Dessau/">Dessau</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Dortmund/">Dortmund</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Dresden/">Dresden</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Dublin/">Dublin</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Duesseldorf/">Duesseldorf</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Duisburg/">Duisburg</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Edinburgh/">Edinburgh</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Eindhoven/">Eindhoven</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Emden/">Emden</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Erfurt/">Erfurt</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Erlangen/">Erlangen</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Eugene/">Eugene</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Flensburg/">Flensburg</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="FortCollins/">FortCollins</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Frankfurt/">Frankfurt</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="FrankfurtOder/">FrankfurtOder</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Freiburg/">Freiburg</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Gdansk/">Gdansk</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Genf/">Genf</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Gent/">Gent</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Gera/">Gera</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Glasgow/">Glasgow</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Gliwice/">Gliwice</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Goerlitz/">Goerlitz</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Goeteborg/">Goeteborg</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Goettingen/">Goettingen</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Graz/">Graz</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Groningen/">Groningen</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Halifax/">Halifax</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Halle/">Halle</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Hamburg/">Hamburg</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Hamm/">Hamm</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Hannover/">Hannover</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Heilbronn/">Heilbronn</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Helsinki/">Helsinki</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Hertogenbosch/">Hertogenbosch</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Huntsville/">Huntsville</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Innsbruck/">Innsbruck</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Istanbul/">Istanbul</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Jena/">Jena</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Jerusalem/">Jerusalem</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Johannesburg/">Johannesburg</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Kaiserslautern/">Kaiserslautern</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Karlsruhe/">Karlsruhe</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Kassel/">Kassel</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Katowice/">Katowice</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Kaunas/">Kaunas</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Kiel/">Kiel</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Kiew/">Kiew</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Koblenz/">Koblenz</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Koeln/">Koeln</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Konstanz/">Konstanz</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="LakeGarda/">LakeGarda</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="LaPaz/">LaPaz</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="LaPlata/">LaPlata</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Lausanne/">Lausanne</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Leeds/">Leeds</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Leipzig/">Leipzig</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Lima/">Lima</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Linz/">Linz</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Lisbon/">Lisbon</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Liverpool/">Liverpool</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Ljubljana/">Ljubljana</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Lodz/">Lodz</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="London/">London</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Luebeck/">Luebeck</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Luxemburg/">Luxemburg</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Lyon/">Lyon</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Maastricht/">Maastricht</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Madison/">Madison</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Madrid/">Madrid</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Magdeburg/">Magdeburg</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Mainz/">Mainz</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Malmoe/">Malmoe</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Manchester/">Manchester</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Mannheim/">Mannheim</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Marseille/">Marseille</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Melbourne/">Melbourne</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Memphis/">Memphis</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="MexicoCity/">MexicoCity</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Miami/">Miami</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Moenchengladbach/">Moenchengladbach</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Montevideo/">Montevideo</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Montpellier/">Montpellier</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Montreal/">Montreal</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Moscow/">Moscow</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Muenchen/">Muenchen</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Muenster/">Muenster</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="NewDelhi/">NewDelhi</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="NewOrleans/">NewOrleans</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="NewYork/">NewYork</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Nuernberg/">Nuernberg</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Oldenburg/">Oldenburg</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Oranienburg/">Oranienburg</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Orlando/">Orlando</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Oslo/">Oslo</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Osnabrueck/">Osnabrueck</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Ostrava/">Ostrava</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Ottawa/">Ottawa</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Paderborn/">Paderborn</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Palma/">Palma</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="PaloAlto/">PaloAlto</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Paris/">Paris</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Perth/">Perth</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Philadelphia/">Philadelphia</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="PhnomPenh/">PhnomPenh</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Portland/">Portland</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="PortlandME/">PortlandME</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Porto/">Porto</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="PortoAlegre/">PortoAlegre</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Potsdam/">Potsdam</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Poznan/">Poznan</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Prag/">Prag</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Providence/">Providence</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Regensburg/">Regensburg</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Riga/">Riga</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="RiodeJaneiro/">RiodeJaneiro</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Rostock/">Rostock</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Rotterdam/">Rotterdam</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Ruegen/">Ruegen</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Saarbruecken/">Saarbruecken</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Sacramento/">Sacramento</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Saigon/">Saigon</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Salzburg/">Salzburg</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="SanFrancisco/">SanFrancisco</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="SanJose/">SanJose</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="SanktPetersburg/">SanktPetersburg</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="SantaBarbara/">SantaBarbara</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="SantaCruz/">SantaCruz</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Santiago/">Santiago</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Sarajewo/">Sarajewo</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Schwerin/">Schwerin</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Seattle/">Seattle</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Seoul/">Seoul</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Sheffield/">Sheffield</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Singapore/">Singapore</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Sofia/">Sofia</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Stockholm/">Stockholm</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Stockton/">Stockton</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Strassburg/">Strassburg</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Stuttgart/">Stuttgart</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Sucre/">Sucre</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Sydney/">Sydney</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Szczecin/">Szczecin</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Tallinn/">Tallinn</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Tehran/">Tehran</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Tilburg/">Tilburg</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Tokyo/">Tokyo</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Toronto/">Toronto</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Toulouse/">Toulouse</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Trondheim/">Trondheim</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Tucson/">Tucson</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Turin/">Turin</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="UlanBator/">UlanBator</a>/</td><td class="m">2019-May-03 05:14:30</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Ulm/">Ulm</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Usedom/">Usedom</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Utrecht/">Utrecht</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Vancouver/">Vancouver</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Victoria/">Victoria</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="WarenMueritz/">WarenMueritz</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Warsaw/">Warsaw</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="WashingtonDC/">WashingtonDC</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Waterloo/">Waterloo</a>/</td><td class="m">2019-May-03 05:14:29</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Wien/">Wien</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Wroclaw/">Wroclaw</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Wuerzburg/">Wuerzburg</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Wuppertal/">Wuppertal</a>/</td><td class="m">2019-May-03 05:14:27</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Zagreb/">Zagreb</a>/</td><td class="m">2019-May-03 05:14:28</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
				<tr class="d"><td class="n"><a href="Zuerich/">Zuerich</a>/</td><td class="m">2019-May-03 05:14:26</td><td class="s">- &nbsp;</td><td class="t">Directory</td></tr>
			</tbody>`,
			url:      `https://download.bbbike.org/osm/bbbike/`,
			elements: &element.MapElement{},
			want:     element.MapElement{},
		},

		// TODO: Add test cases.
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(thisTest.html))
			myURL, _ := url.Parse(thisTest.url)
			myElement := &colly.HTMLElement{
				DOM: dom.Selection,
				Response: &colly.Response{
					Request: &colly.Request{URL: myURL},
				},
				Request: &colly.Request{URL: myURL},
			}

			defaultbbbike := bbbike.GetDefault()
			myCollector := defaultbbbike.Collector() // Need a Collector to visit

			for _, elemem := range *thisTest.elements {
				if err := defaultbbbike.Config.MergeElement(&elemem); err != nil {
					t.Errorf("Bad tests g.Config.mergeElement() can't merge %v - %v", elemem, err)
				}
			}

			defaultbbbike.ParseList(myElement, myCollector)

			if !reflect.DeepEqual(defaultbbbike.Config.Elements, thisTest.want) {
				t.Errorf("parseSubregion() fail, got \n%v, want \n%v", defaultbbbike.Config.Elements, thisTest.want)
			}
		})
	}
}

//nolint:lll // this func contain html extracts
func Test_bbbike_parseSidebar(t *testing.T) {
	t.Parallel()

	tests := []struct {
		elements *element.MapElement
		want     element.MapElement
		name     string
		html     string
		url      string
	}{
		{
			name: "Toulouse",
			html: `<div id="sidebar">
				<div id="routes"><h3 style="text-align:center">OSM extracts for Toulouse</h3>
				<p>
				Welcome to BBBike's free download server!
				This server has data extracts from the OpenStreetMap project
				for the area Toulouse in
				a variety of 
				<a href="https://extract.bbbike.org/extract-screenshots.html">formats and styles</a>
				for you to use:
				</p>

				<!-- <table> -->

				<a class="download_link" href="./Toulouse.osm.pbf" title="Thu May  2 09:45:34 2019"> Protocolbuffer (PBF) <span class="size">38M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.gz" title="Thu May  2 18:42:55 2019"> OSM XML gzip'd <span class="size">70M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.shp.zip" title="Thu May  2 19:23:31 2019"> Shapefile (Esri) <span class="size">68M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.garmin-onroad.zip" title="Thu May  2 20:06:17 2019"> Garmin Onroad (UTF-8) <span class="size">1.9M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.garmin-opentopo.zip" title="Thu May  2 20:08:33 2019"> Garmin OpenTopoMap (UTF-8) <span class="size">11M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.garmin-osm.zip" title="Thu May  2 20:07:36 2019"> Garmin OSM (UTF-8) <span class="size">11M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.geojson.xz" title="Thu May  2 18:57:15 2019"> Osmium GeoJSON 7z (xz) <span class="size">65M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.svg-osm.zip" title="Fri May  3 03:13:17 2019"> SVG mapnik <span class="size">35M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.mapsforge-osm.zip" title="Fri May  3 01:42:08 2019"> Mapsforge OSM <span class="size">18M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.navit.zip" title="Thu May  2 19:36:10 2019"> Navit <span class="size">26M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.garmin-onroad-latin1.zip" title="Thu May  2 20:06:48 2019"> Garmin Onroad (latin1) <span class="size">1.8M</span>
				</a>
				<a class="download_link" href="./Toulouse.osm.csv.xz" title="Thu May  2 18:52:30 2019"> csv 7z (xz) <span class="size">5.2M</span>
				</a>
				<br><a class="small" href="./Toulouse.poly">Poly</a>
				<a class="small" href="./CHECKSUM.txt" title="Fri May  3 05:14:26 2019">CHECKSUM.txt</a>
				<!-- </table> -->

				<p>
				Didn't find the area you want?<br>
				<a href="https://extract.bbbike.org/">Select your custom region</a>
				- a rectangle or polygon up to 6000x4000km large, or 512MB file size.
				</p>

				<span id="big_donate_image">
				<center>
				<a href="https://www.bbbike.org/community.html"><img class="logo" alt="donate" height="47" width="126" src="/images/btn_donateCC_LG.gif"></a>
				</center>
				</span>

				<br>
				<span style="font-size:small">
				<a href="https://extract.bbbike.org/extract.html" target="_new">help</a> |
				<a href="https://extract.bbbike.org/extract-screenshots.html" target="_new">screenshots</a> |
				<a href="https://extract.bbbike.org/" target="_new">extracts</a> |
				<a href="https://extract.bbbike.org/extract.html#extract-pro">commercial support</a>
				</span>
				<hr>

				<span class="city">
				Start bicycle routing for <a style="font-size:x-large" href="https://www.bbbike.org/Toulouse/">Toulouse</a>
				</span>
				</div>
				</div>`,
			url:      `https://download.bbbike.org/osm/bbbike/Toulouse/`,
			elements: &element.MapElement{},
			want: element.MapElement{
				"Toulouse": element.Element{
					ID:     "Toulouse",
					File:   "Toulouse/Toulouse",
					Name:   "Toulouse",
					Parent: "",
					Formats: element.Formats{
						formats.FormatCSV,
						formats.FormatGarminOSM,
						formats.FormatGarminOnroad,
						formats.FormatGarminOntrail,
						formats.FormatGarminOpenTopo,
						formats.FormatGeoJSON,
						formats.FormatMBTiles,
						formats.FormatMapsforge,
						formats.FormatOsmGz,
						formats.FormatOsmPbf,
						formats.FormatPoly,
						formats.FormatShpZip,
					},
				},
			},
		},

		// TODO: Add test cases.
	}
	for _, thisTest := range tests {
		t.Run(thisTest.name, func(t *testing.T) {
			t.Parallel()

			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(thisTest.html))
			myURL, _ := url.Parse(thisTest.url)
			myElement := &colly.HTMLElement{
				DOM: dom.Selection,
				Response: &colly.Response{
					Request: &colly.Request{URL: myURL},
				},
				Request: &colly.Request{URL: myURL},
			}

			defaultBbbike := bbbike.GetDefault()
			myCollector := defaultBbbike.Collector() // Need a Collector to visit

			for _, elemem := range *thisTest.elements {
				if err := defaultBbbike.Config.MergeElement(&elemem); err != nil {
					t.Errorf("Bad tests g.Config.mergeElement() can't merge %v - %v", elemem, err)
				}
			}

			defaultBbbike.ParseSidebar(myElement, myCollector)

			if !reflect.DeepEqual(defaultBbbike.Config.Elements, thisTest.want) {
				t.Errorf("parseSubregion() fail, got \n%v, want \n%v", defaultBbbike.Config.Elements, thisTest.want)
			}
		})
	}
}
