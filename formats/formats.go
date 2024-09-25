package formats

import (
	"sort"
	"strings"

	"github.com/spf13/viper"
)

// Format represents a file format with various attributes.
type Format struct {
	ID       string `yaml:"ext"`
	Loc      string `yaml:"loc"`
	BasePath string `yaml:"basepath,omitempty"`
	BaseURL  string `yaml:"baseurl,omitempty"`
	ToLoc    string `yaml:"toloc,omitempty"`
}

// FormatDefinitions is a map of format definitions.
type FormatDefinitions map[string]Format

// MiniFormat represents a short and full name pair for a format.
type MiniFormat struct {
	ShortName string
	FullName  string
}

const (
	FormatState          = "state"
	FormatOsmPbf         = "osm.pbf"
	FormatOsmGz          = "osm.gz"
	FormatOsmBz2         = "osm.bz2"
	FormatOshPbf         = "osh.pbf"
	FormatPoly           = "poly"
	FormatShpZip         = "shp.zip"
	FormatKml            = "kml"
	FormatGeoJSON        = "geojson"                        // BBBike & OSM Today only
	FormatGarminOntrail  = "osm.garmin-ontrail-latin1.zip"  // BBBike only
	FormatGarminOnroad   = "osm.garmin-onroad-latin1.zip"   // BBBike only
	FormatGarminOpenTopo = "osm.garmin-opentopo-latin1.zip" // BBBike only
	FormatGarminOSM      = "osm.garmin-osm.zip"             // BBBike only
	FormatMapsforge      = "osm.mapsforge-osm.zip"          // BBBike only
	FormatMBTiles        = "mbtiles"
	FormatCSV            = "csv" // BBBike only
)

// Configuration keys
const (
	KeyOsmPbf         = "dosmPbf"
	KeyOshPbf         = "doshPbf"
	KeyOsmGz          = "dosmGz"
	KeyOsmBz2         = "dosmBz2"
	KeyShpZip         = "dshpZip"
	KeyState          = "dstate"
	KeyPoly           = "dpoly"
	KeyKml            = "dkml"
	KeyGeoJSON        = "dgeojson"
	KeyGarminOSM      = "dgarmin"
	KeyMapsforge      = "dmaps"
	KeyMBTiles        = "dmbtiles"
	KeyCSV            = "dcsv"
	KeyGarminOnroad   = "dgarminonroad"
	KeyGarminOntrail  = "dgarminontrail"
	KeyGarminOpenTopo = "dgarminopentopo"
)

// GetMiniFormats returns a string of short format names based on the provided full format names.
func GetMiniFormats(fullFormatNames []string) string {
	miniFormats := []MiniFormat{
		{ShortName: "s", FullName: FormatState},
		{ShortName: "P", FullName: FormatOsmPbf},
		{ShortName: "G", FullName: FormatOsmGz},
		{ShortName: "B", FullName: FormatOsmBz2},
		{ShortName: "H", FullName: FormatOshPbf},
		{ShortName: "p", FullName: FormatPoly},
		{ShortName: "S", FullName: FormatShpZip},
		{ShortName: "k", FullName: FormatKml},
		{ShortName: "g", FullName: FormatGeoJSON},
		{ShortName: "t", FullName: FormatGarminOntrail},
		{ShortName: "r", FullName: FormatGarminOnroad},
		{ShortName: "o", FullName: FormatGarminOpenTopo},
		{ShortName: "O", FullName: FormatGarminOSM},
		{ShortName: "m", FullName: FormatMapsforge},
		{ShortName: "M", FullName: FormatMBTiles},
		{ShortName: "C", FullName: FormatCSV},
	}

	shortNames := make([]string, 0, len(fullFormatNames))

	for _, fullName := range fullFormatNames {
		for _, format := range miniFormats {
			if fullName == format.FullName {
				shortNames = append(shortNames, format.ShortName)
				break
			}
		}
	}

	return strings.Join(shortNames, "")
}

// GetFormats returns a slice of format strings based on the configuration.
func GetFormats() []string {
	options := map[string]string{
		KeyOsmPbf:         FormatOsmPbf,
		KeyOshPbf:         FormatOshPbf,
		KeyOsmGz:          FormatOsmGz,
		KeyOsmBz2:         FormatOsmBz2,
		KeyShpZip:         FormatShpZip,
		KeyState:          FormatState,
		KeyPoly:           FormatPoly,
		KeyKml:            FormatKml,
		KeyGeoJSON:        FormatGeoJSON,
		KeyGarminOSM:      FormatGarminOSM,
		KeyMapsforge:      FormatMapsforge,
		KeyMBTiles:        FormatMBTiles,
		KeyCSV:            FormatCSV,
		KeyGarminOnroad:   FormatGarminOnroad,
		KeyGarminOntrail:  FormatGarminOntrail,
		KeyGarminOpenTopo: FormatGarminOpenTopo,
	}

	var formatList []string

	for key, format := range options {
		if viper.GetBool(key) {
			formatList = append(formatList, format)
		}
	}

	if len(formatList) == 0 {
		formatList = append(formatList, FormatOsmPbf)
	}

	sort.Strings(formatList)

	return formatList
}
