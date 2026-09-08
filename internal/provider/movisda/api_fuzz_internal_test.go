package movisda

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func FuzzDecodeGeoJSON(f *testing.F) {
	f.Add([]byte(`{"features":[{"properties":{"prefix":"france","name":"France","name_en":"France"}}]}`))
	f.Add([]byte(`{"features":[{"properties":{"prefix":"--region--","name":"Region"}}]}`))
	f.Add([]byte(`{"features":[{"properties":{"prefix":"","name":""}}]}`))
	f.Add([]byte(`{}`))
	f.Add([]byte(`{"features":[]}`))
	f.Add([]byte(`invalid json`))

	f.Fuzz(func(_ *testing.T, jsonData []byte) {
		var geo adminGeoJSON
		if err := json.NewDecoder(bytes.NewReader(jsonData)).Decode(&geo); err != nil {
			return
		}

		cat := catalog.New()
		cat.BaseURL = MovisdaBaseURL
		cat.Formats = DefaultFormats()

		for i := range geo.Features {
			elem := buildElement(&geo.Features[i])
			if elem.ID == "" {
				continue
			}

			_ = cat.MergeElement(&elem)
		}
	})
}

func FuzzBuildElement(f *testing.F) {
	f.Add("france", "France", "France EN")
	f.Add("-alpes-", "Alpes", "")
	f.Add("", "", "")
	f.Add("---", "", "")
	f.Add("special/chars", "Name", "Name EN")

	f.Fuzz(func(_ *testing.T, prefix, name, nameEN string) {
		feat := geoJSONFeature{}
		feat.Properties.Prefix = prefix
		feat.Properties.Name = name
		feat.Properties.NameEN = nameEN

		_ = buildElement(&feat)
	})
}
