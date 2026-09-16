package catalog

import (
	"slices"
)

// Supported format identifiers.
const (
	FormatOsmPbf         = "osm.pbf"
	FormatPbf            = "pbf"
	FormatOsmBz2         = "osm.bz2"
	FormatOsmGz          = "osm.gz"
	FormatOshPbf         = "osh.pbf"
	FormatShpZip         = "shp.zip"
	FormatPoly           = "poly"
	FormatKml            = "kml"
	FormatState          = "state.txt"
	FormatGeoJSON        = "geojson"
	FormatGarminOSM      = "garmin-osm.zip"
	FormatMapsforge      = "mapsforge-osm.zip"
	FormatMBTiles        = "mbtiles.zip"
	FormatCSV            = "csv.xz"
	FormatGarminOnroad   = "garmin-onroad.zip"
	FormatGarminOntrail  = "garmin-ontrail.zip"
	FormatGarminOpenTopo = "garmin-opentopo.zip"
	FormatOBF            = "obf"
	FormatGPKG           = "gpkg"
	FormatO5m            = "o5m"
	FormatO5mZst         = "o5m.zst"
)

// Format flag keys used in CLI options and configurations.
const (
	KeyOsmPbf         = "dosmPbf"
	KeyPbf            = "dPbf"
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
	KeyOBF            = "dobf"
	KeyGPKG           = "dgpkg"
	KeyO5m            = "do5m"
	KeyO5mZst         = "do5mZst"
)

// Format represents the metadata and URL resolution rules for a file format.
type Format struct {
	// ID is the file extension identifier (e.g., ".osm.pbf").
	ID string `json:"id" yaml:"ext"`

	// Loc is the relative URL pattern or suffix template for the format file.
	Loc string `json:"loc" yaml:"loc"`

	// BasePath is an optional intermediate path prefix within the provider URL hierarchy.
	BasePath string `json:"basepath,omitempty" yaml:"basepath,omitempty"`

	// BaseURL is an optional override for the root download base URL for this format.
	BaseURL string `json:"baseurl,omitempty" yaml:"baseurl,omitempty"`

	// ToLoc is an optional transformation target or link pattern.
	ToLoc string `json:"toloc,omitempty" yaml:"toloc,omitempty"`

	// Type specifies the format classification (e.g. data or checksum).
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// FormatDefinitions maps format IDs to Format specifications.
type FormatDefinitions map[string]Format

// MiniFormat maps full format names to single-letter display abbreviations.
type MiniFormat struct {
	// FullName is the format identifier string (e.g., "osm.pbf").
	FullName string

	// ShortName is the single-letter CLI table display abbreviation (e.g., "P").
	ShortName string
}

var miniFormatBytes = map[string]byte{
	FormatState:   's',
	FormatOsmBz2:  'B',
	FormatOsmGz:   'G',
	FormatOshPbf:  'H',
	FormatOsmPbf:  'P',
	FormatPbf:     'P',
	FormatPoly:    'p',
	FormatKml:     'k',
	FormatShpZip:  'S',
	FormatGeoJSON: 'g',
	FormatOBF:     'o',
	FormatGPKG:    'K',
	FormatO5m:     '5',
	FormatO5mZst:  'Z',
}

// GetMiniFormats returns a compact string representation of the given format slice.
func GetMiniFormats(formatList []string) string {
	if len(formatList) == 0 {
		return ""
	}

	var buf [16]byte

	idx := 0

	for _, fullName := range formatList {
		if ch, exists := miniFormatBytes[fullName]; exists && idx < len(buf) {
			buf[idx] = ch
			idx++
		}
	}

	return string(buf[:idx])
}

var defaultFlagToFormat = map[string]string{
	KeyOsmPbf:         FormatOsmPbf,
	KeyPbf:            FormatPbf,
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
	KeyOBF:            FormatOBF,
	KeyGPKG:           FormatGPKG,
	KeyO5m:            FormatO5m,
	KeyO5mZst:         FormatO5mZst,
}

// GetFormats converts a map of enabled boolean flags into a sorted slice of format IDs.
func GetFormats(flagMap map[string]bool) []string {
	var formats []string

	for key, format := range defaultFlagToFormat {
		if enabled, ok := flagMap[key]; ok && enabled {
			formats = append(formats, format)
		}
	}

	if len(formats) == 0 {
		return []string{FormatOsmPbf}
	}

	slices.Sort(formats)

	return formats
}
