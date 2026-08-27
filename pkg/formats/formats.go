package formats

import (
	"slices"
)

// Format represents a file format with various attributes.
type Format struct {
	ID       string `yaml:"ext"`
	Loc      string `yaml:"loc"`
	BasePath string `yaml:"basepath,omitempty"`
	BaseURL  string `yaml:"baseurl,omitempty"`
	ToLoc    string `yaml:"toloc,omitempty"`
	Type     string `yaml:"type,omitempty"`
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

// Configuration keys.
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

var miniFormatMap = map[string]byte{
	FormatState:          's',
	FormatOsmPbf:         'P',
	FormatOsmGz:          'G',
	FormatOsmBz2:         'B',
	FormatOshPbf:         'H',
	FormatPoly:           'p',
	FormatShpZip:         'S',
	FormatKml:            'k',
	FormatGeoJSON:        'g',
	FormatGarminOntrail:  't',
	FormatGarminOnroad:   'r',
	FormatGarminOpenTopo: 'o',
	FormatGarminOSM:      'O',
	FormatMapsforge:      'm',
	FormatMBTiles:        'M',
	FormatCSV:            'C',
}

type keyFormatPair struct {
	key    string
	format string
}

var keyFormatPairs = []keyFormatPair{
	{KeyOsmPbf, FormatOsmPbf},
	{KeyOshPbf, FormatOshPbf},
	{KeyOsmGz, FormatOsmGz},
	{KeyOsmBz2, FormatOsmBz2},
	{KeyShpZip, FormatShpZip},
	{KeyState, FormatState},
	{KeyPoly, FormatPoly},
	{KeyKml, FormatKml},
	{KeyGeoJSON, FormatGeoJSON},
	{KeyGarminOSM, FormatGarminOSM},
	{KeyMapsforge, FormatMapsforge},
	{KeyMBTiles, FormatMBTiles},
	{KeyCSV, FormatCSV},
	{KeyGarminOnroad, FormatGarminOnroad},
	{KeyGarminOntrail, FormatGarminOntrail},
	{KeyGarminOpenTopo, FormatGarminOpenTopo},
}

// GetMiniFormats returns a string of short format names based on the provided full format names.
func GetMiniFormats(fullFormatNames []string) string {
	if len(fullFormatNames) == 0 {
		return ""
	}

	buf := make([]byte, 0, len(fullFormatNames))

	for _, fullName := range fullFormatNames {
		if b, ok := miniFormatMap[fullName]; ok {
			buf = append(buf, b)
		}
	}

	return string(buf)
}

// GetFormats returns a slice of format strings based on the configuration map.
// The config map should contain keys like KeyOsmPbf with boolean true/false.
func GetFormats(config map[string]bool) []string {
	formatList := make([]string, 0, len(config))

	for _, pair := range keyFormatPairs {
		if config[pair.key] {
			formatList = append(formatList, pair.format)
		}
	}

	if len(formatList) == 0 {
		formatList = append(formatList, FormatOsmPbf)
	}

	slices.Sort(formatList)

	return formatList
}
