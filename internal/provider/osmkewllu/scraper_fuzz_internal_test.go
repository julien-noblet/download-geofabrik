package osmkewllu

import (
	"bytes"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func FuzzParseKewlLuHTML(f *testing.F) {
	f.Add([]byte(`<a href="luxembourg.osm.pbf">pbf</a><a href="luxembourg.osm.bz2">bz2</a>`))
	f.Add([]byte(`<a href="luxembourg-diff-001.osm.pbf">diff</a><a href="script.sh">script</a>`))
	f.Add([]byte(`<a href="../">Parent</a><a href="?C=N;O=D">Sort</a>`))
	f.Add([]byte(``))
	f.Add([]byte(`<invalid html <> >>>`))

	f.Fuzz(func(_ *testing.T, htmlData []byte) {
		cat := catalog.New()
		cat.BaseURL = BaseURL
		cat.Formats = DefaultFormats()

		_ = parseKewlLuHTML(bytes.NewReader(htmlData), cat)
	})
}

func FuzzParseHrefToIDAndFormat(f *testing.F) {
	f.Add("luxembourg.osm.pbf")
	f.Add("luxembourg.osm.bz2")
	f.Add("luxembourg-diff.osm.pbf")
	f.Add("other.unknown")
	f.Add("")

	f.Fuzz(func(_ *testing.T, href string) {
		_, _ = parseHrefToIDAndFormat(href)
	})
}
