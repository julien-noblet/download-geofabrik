package geo2day

import (
	"bytes"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func FuzzParseHTMLStream(f *testing.F) {
	prov := NewProvider()

	f.Add([]byte(`<a href="africa.html">Africa</a><a href="algeria.osm.pbf">Algeria</a>`))
	f.Add([]byte(`<a href="morocco.poly">Morocco Poly</a><a href="algeria.osm.pbf.md5">MD5</a>`))
	f.Add([]byte(`<a href="#top">Top</a><a href="mailto:info@geo2day.com">Mail</a>`))
	f.Add([]byte(""))
	f.Add([]byte("<invalid html <> >>>"))

	f.Fuzz(func(_ *testing.T, htmlData []byte) {
		cat := catalog.New()
		cat.BaseURL = prov.BaseURL
		cat.Formats = DefaultFormats()

		_, _ = prov.parseHTMLStream(bytes.NewReader(htmlData), cat)
	})
}

func FuzzSplitFileExt(f *testing.F) {
	f.Add("https://geo2day.com/algeria.osm.pbf")
	f.Add("algeria.osm.pbf?query=1#hash")
	f.Add("file.poly")
	f.Add("no_ext")
	f.Add("")
	f.Add("/")
	f.Add("a/b/c/d.e.f.g")

	f.Fuzz(func(_ *testing.T, urlStr string) {
		_, _ = splitFileExt(urlStr)
	})
}

func FuzzSplitParent(f *testing.F) {
	f.Add("https://geo2day.com/africa/algeria.osm.pbf")
	f.Add("https://geo2day.com/algeria.osm.pbf")
	f.Add("africa/algeria.osm.pbf")
	f.Add("https://geo2day.com/")
	f.Add("://")
	f.Add("")
	f.Add("/")
	f.Add("a/b/c/d/e")

	f.Fuzz(func(_ *testing.T, urlStr string) {
		_, _ = splitParent(urlStr)
	})
}

func FuzzSanitizeFileName(f *testing.F) {
	f.Add("[DIR]")
	f.Add("d41d8cd98f00b204e9800998ecf8427e")
	f.Add("valid-name")
	f.Add("")
	f.Add("   ")

	f.Fuzz(func(_ *testing.T, name string) {
		_ = sanitizeFileName(name)
	})
}
