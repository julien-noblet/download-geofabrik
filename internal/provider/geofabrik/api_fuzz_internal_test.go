package geofabrik

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func FuzzDecodeIndexJSON(f *testing.F) {
	f.Add([]byte(`{"features":[{"properties":{"id":"france","name":"France",` +
		`"urls":{"pbf":"https://download.geofabrik.de/europe/france-latest.osm.pbf"}}}]}`))
	f.Add([]byte(`{"features":[{"properties":{"id":"monaco","name":"Monaco","parent":"europe",` +
		`"urls":{"shp":"https://download.geofabrik.de/europe/monaco.shp.zip",` +
		`"history":"https://download.geofabrik.de/europe/monaco.osh.pbf"}}}]}`))
	f.Add([]byte(`{}`))
	f.Add([]byte(`{"features":[]}`))
	f.Add([]byte(`invalid json`))

	f.Fuzz(func(_ *testing.T, jsonData []byte) {
		var index indexJSON
		if err := json.NewDecoder(bytes.NewReader(jsonData)).Decode(&index); err != nil {
			return
		}

		cat := catalog.New()
		cat.BaseURL = GeofabrikBaseURL
		cat.Formats = DefaultFormats()

		for _, feat := range index.Features {
			elem := catalog.Element{
				ID:      feat.Properties.ID,
				Name:    feat.Properties.Name,
				Parent:  feat.Properties.Parent,
				Formats: extractFormats(feat.Properties.Urls),
			}

			_ = cat.MergeElement(&elem)
		}
	})
}

func FuzzExtractFormats(f *testing.F) {
	f.Add("pbf")
	f.Add("bz2")
	f.Add("shp")
	f.Add("history")
	f.Add("unknown")
	f.Add("")

	f.Fuzz(func(_ *testing.T, key string) {
		urls := map[string]string{key: "https://example.com/file"}
		_ = extractFormats(urls)
	})
}
