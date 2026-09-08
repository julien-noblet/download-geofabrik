package bbbike

import (
	"bytes"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func FuzzParseBBBikeHTML(f *testing.F) {
	f.Add([]byte(`<html><body><a href="Berlin/">Berlin</a><a href="Paris/">Paris</a></body></html>`))
	f.Add([]byte(`<a href="../">Parent</a><a href="?C=N;O=D">Sort</a><a href="Aachen">Aachen</a>`))
	f.Add([]byte(""))
	f.Add([]byte("<invalid html <> >>>"))
	f.Add([]byte("<a href=\"/root\">Root</a><a href=\"123num\">123</a>"))

	f.Fuzz(func(_ *testing.T, htmlData []byte) {
		cat := catalog.New()
		cat.BaseURL = BaseURL
		cat.Formats = DefaultFormats()

		_ = parseBBBikeHTML(bytes.NewReader(htmlData), cat)
	})
}

func FuzzCleanCityName(f *testing.F) {
	f.Add("Berlin")
	f.Add("Berlin/")
	f.Add("/Berlin")
	f.Add("berlin")
	f.Add("../etc/passwd")
	f.Add("?C=N;O=D")
	f.Add("http://example.com")
	f.Add("")
	f.Add("\x00")

	f.Fuzz(func(_ *testing.T, href string) {
		_ = cleanCityName(href)
	})
}
