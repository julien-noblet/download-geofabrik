package osmtw

import (
	"bytes"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func FuzzParseTaiwanHTML(f *testing.F) {
	f.Add([]byte(`<a href="recent/">recent/</a><a href="basemap/">basemap/</a>`))
	f.Add([]byte(`<a href="diff/">diff/</a><a href="changes.txt">changes</a>`))
	f.Add([]byte(`<a href="../">Parent</a><a href="?C=N;O=D">Sort</a>`))
	f.Add([]byte(``))
	f.Add([]byte(`<invalid html <> >>>`))

	f.Fuzz(func(_ *testing.T, htmlData []byte) {
		cat := catalog.New()
		cat.BaseURL = BaseURL
		cat.Formats = DefaultFormats()

		_ = parseTaiwanHTML(bytes.NewReader(htmlData), cat)
	})
}

func FuzzParseLink(f *testing.F) {
	f.Add("recent")
	f.Add("recent/")
	f.Add("basemap")
	f.Add("diff")
	f.Add("other")
	f.Add("../")
	f.Add("")

	f.Fuzz(func(_ *testing.T, href string) {
		cat := catalog.New()
		cat.BaseURL = BaseURL
		cat.Formats = DefaultFormats()

		parseLink(href, cat)
	})
}
