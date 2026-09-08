package osmfitvutbr

import (
	"bytes"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func FuzzParseFitVutbrRootHTML(f *testing.F) {
	f.Add([]byte(`<a href="czech_republic/">Czech</a><a href="czech-republic.poly">Poly</a>`))
	f.Add([]byte(`<a href="../">Parent</a><a href="?C=N;O=D">Sort</a>`))
	f.Add([]byte(``))
	f.Add([]byte(`<invalid html <> >>>`))

	f.Fuzz(func(_ *testing.T, htmlData []byte) {
		cat := catalog.New()
		cat.BaseURL = BaseURL
		cat.Formats = DefaultFormats()

		_, _ = parseFitVutbrRootHTML(bytes.NewReader(htmlData), cat)
	})
}

func FuzzParseFitVutbrSubdirHTML(f *testing.F) {
	f.Add([]byte(`<a href="czech_republic-2023-01-01.osm.pbf">pbf</a><a href="czech_republic-2023-01-01.osm.bz2">bz2</a>`), "czech_republic")
	f.Add([]byte(`<a href="../">Parent</a>`), "czech_republic")
	f.Add([]byte(``), "unknown")
	f.Add([]byte(`<invalid html <> >>>`), "")

	f.Fuzz(func(_ *testing.T, htmlData []byte, dir string) {
		cat := catalog.New()
		cat.BaseURL = BaseURL
		cat.Formats = DefaultFormats()

		_ = parseFitVutbrSubdirHTML(bytes.NewReader(htmlData), cat, dir)
	})
}

func FuzzParseSubdirLink(f *testing.F) {
	f.Add("czech_republic-2023-01-01.osm.pbf", "czech_republic")
	f.Add("czech_republic-2023-01-01.osm.bz2", "czech_republic")
	f.Add("other-2023.osm.pbf", "czech_republic")
	f.Add("../", "czech_republic")
	f.Add("", "")
	f.Add("czech_republic-", "czech_republic")

	f.Fuzz(func(_ *testing.T, href, dir string) {
		var (
			latestPbfDate string
			latestPbfBase string
			latestBz2Date string
			latestBz2Base string
		)

		parseSubdirLink(href, dir, &latestPbfDate, &latestPbfBase, &latestBz2Date, &latestBz2Base)
	})
}

func FuzzFormatName(f *testing.F) {
	f.Add("czech_republic")
	f.Add("prague")
	f.Add("")
	f.Add("a_b_c")
	f.Add("---")

	f.Fuzz(func(_ *testing.T, id string) {
		_ = formatName(id)
	})
}
