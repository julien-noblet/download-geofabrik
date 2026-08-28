package catalog

import (
	"slices"
	"strings"
)

// Supported format identifiers.
const (
	FormatOsmPbf         = "osm.pbf"
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
	KeyOshPbf         = "doshPbf"
	KeyOsmGz          = "dosmGz"
	KeyOsmBz2         = "dosmBz2"
	KeyShpZip         = "dshpZip"
	KeyState          = "dstate"
	KeyPoly           = "dpoly"
	KeyKml            = "dkml"
	KeyGeoJSON        = "dgeojson"
	KeyGarminOSM      = "dgarminOSM"
	KeyMapsforge      = "dmapsforge"
	KeyMBTiles        = "dmbtiles"
	KeyCSV            = "dcsv"
	KeyGarminOnroad   = "dgarminOnroad"
	KeyGarminOntrail  = "dgarminOntrail"
	KeyGarminOpenTopo = "dgarminOpenTopo"
	KeyOBF            = "dobf"
	KeyGPKG           = "dgpkg"
	KeyO5m            = "do5m"
	KeyO5mZst         = "do5mZst"
)

// Format represents the metadata and URL resolution rules for a file format.
type Format struct {
	ID       string `json:"id"                 yaml:"ext"`
	Loc      string `json:"loc"                yaml:"loc"`
	BasePath string `json:"basepath,omitempty" yaml:"basepath,omitempty"`
	BaseURL  string `json:"baseurl,omitempty"  yaml:"baseurl,omitempty"`
	ToLoc    string `json:"toloc,omitempty"    yaml:"toloc,omitempty"`
	Type     string `json:"type,omitempty"     yaml:"type,omitempty"`
}

// FormatDefinitions maps format IDs to Format specifications.
type FormatDefinitions map[string]Format

// MiniFormat maps full format names to single-letter display abbreviations.
type MiniFormat struct {
	FullName  string
	ShortName string
}

var defaultMiniFormats = []MiniFormat{
	{FullName: FormatState, ShortName: "s"},
	{FullName: FormatOsmBz2, ShortName: "B"},
	{FullName: FormatOsmGz, ShortName: "G"},
	{FullName: FormatOshPbf, ShortName: "H"},
	{FullName: FormatOsmPbf, ShortName: "P"},
	{FullName: FormatPoly, ShortName: "p"},
	{FullName: FormatKml, ShortName: "k"},
	{FullName: FormatShpZip, ShortName: "S"},
	{FullName: FormatGeoJSON, ShortName: "g"},
	{FullName: FormatOBF, ShortName: "o"},
	{FullName: FormatGPKG, ShortName: "K"},
	{FullName: FormatO5m, ShortName: "5"},
	{FullName: FormatO5mZst, ShortName: "Z"},
}

// GetMiniFormats returns a compact string representation of the given format slice.
func GetMiniFormats(formatList []string) string {
	if len(formatList) == 0 {
		return ""
	}

	var shortNames []string

	for _, fullName := range formatList {
		for _, mini := range defaultMiniFormats {
			if fullName == mini.FullName {
				shortNames = append(shortNames, mini.ShortName)

				break
			}
		}
	}

	return strings.Join(shortNames, "")
}

// GetFormats converts a map of enabled boolean flags into a sorted slice of format IDs.
func GetFormats(flagMap map[string]bool) []string {
	flagToFormat := map[string]string{
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
		KeyOBF:            FormatOBF,
		KeyGPKG:           FormatGPKG,
		KeyO5m:            FormatO5m,
		KeyO5mZst:         FormatO5mZst,
	}

	var formats []string

	for key, format := range flagToFormat {
		if enabled, ok := flagMap[key]; ok && enabled {
			formats = append(formats, format)
		}
	}

	if len(formats) == 0 {
		formats = append(formats, FormatOsmPbf)
	}

	slices.Sort(formats)

	return formats
}
