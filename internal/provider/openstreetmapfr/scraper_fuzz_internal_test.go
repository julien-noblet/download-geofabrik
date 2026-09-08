package openstreetmapfr

import (
	"bytes"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func FuzzParseHTMLStream(f *testing.F) {
	prov := NewProvider()

	f.Add([]byte(`<a href="europe/">europe/</a><a href="france-latest.osm.pbf">france</a>`), "https://download.openstreetmap.fr/extracts/")
	f.Add([]byte(`<a href="russia/central-latest.osm.pbf">central</a>`), "https://download.openstreetmap.fr/extracts/russia/")
	f.Add([]byte(`<a href="../">Parent</a><a href="?C=N;O=D">Sort</a>`), "https://download.openstreetmap.fr/extracts/")
	f.Add([]byte(""), "https://download.openstreetmap.fr/extracts/")
	f.Add([]byte("<a href=\"/root\">Root</a>"), "https://download.openstreetmap.fr/extracts/")
	f.Add([]byte("<invalid html <> >>>"), "")

	f.Fuzz(func(_ *testing.T, htmlData []byte, currentURL string) {
		cat := catalog.New()
		cat.BaseURL = prov.BaseURL
		cat.Formats = DefaultFormats()

		_, _ = prov.parseHTMLStream(bytes.NewReader(htmlData), currentURL, cat)
	})
}

func FuzzGetParent(f *testing.F) {
	f.Add("https://download.openstreetmap.fr/extracts/europe/france/alpes_de_haute_provence.osm.pbf")
	f.Add("https://download.openstreetmap.fr/extracts/europe/france.osm.pbf")
	f.Add("http://example.com/file.pbf")
	f.Add("")
	f.Add("/")
	f.Add("////")
	f.Add("a/b/c/d/e/f")
	f.Add("../../../etc/passwd")

	f.Fuzz(func(_ *testing.T, href string) {
		_, _ = getParent(href)
	})
}

func FuzzResolveURL(f *testing.F) {
	f.Add("https://download.openstreetmap.fr/extracts", "europe/")
	f.Add("https://download.openstreetmap.fr/extracts/", "europe/")
	f.Add("https://download.openstreetmap.fr/extracts/", "https://other.com/file.pbf")
	f.Add("http://example.com", "http://other.com")
	f.Add("", "")
	f.Add("/", "/")

	f.Fuzz(func(_ *testing.T, base, ref string) {
		_ = resolveURL(base, ref)
	})
}

func FuzzExceptions(f *testing.F) {
	f.Add("central", "russia")
	f.Add("france_taaf", "france")
	f.Add("normal_name", "europe")
	f.Add("", "")
	f.Add("east", "")

	f.Fuzz(func(_ *testing.T, name, parent string) {
		_ = exceptions(name, parent)
	})
}
