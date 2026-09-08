package osmch

import (
	"bytes"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func FuzzParseOSMCHHTML(f *testing.F) {
	f.Add([]byte(`<a href="switzerland.osm.pbf">switzerland</a><a href="switzerland.poly">poly</a>`))
	f.Add([]byte(`<a href="switzerland-garmin.zip">garmin</a><a href="switzerland.obf">obf</a>`))
	f.Add([]byte(`<a href="../">Parent</a><a href="?C=N;O=D">Sort</a>`))
	f.Add([]byte(``))
	f.Add([]byte(`<invalid html <> >>>`))

	f.Fuzz(func(_ *testing.T, htmlData []byte) {
		cat := catalog.New()
		cat.BaseURL = BaseURL
		cat.Formats = DefaultFormats()

		_ = parseOSMCHHTML(bytes.NewReader(htmlData), cat)
	})
}

func FuzzParseHrefToIDAndFormat(f *testing.F) {
	f.Add("switzerland.osm.pbf")
	f.Add("switzerland.poly")
	f.Add("switzerland.obf")
	f.Add("switzerland-garmin.zip")
	f.Add("switzerland.pbf")
	f.Add("switzerland.unknown")
	f.Add("")

	f.Fuzz(func(_ *testing.T, href string) {
		_, _ = parseHrefToIDAndFormat(href)
	})
}
