package formats

import (
	"slices"
	"strings"

	"github.com/spf13/viper"
)

type Format struct {
	ID       string `yaml:"ext"`
	Loc      string `yaml:"loc"`
	BasePath string `yaml:"basepath,omitempty"`
	BaseURL  string `yaml:"baseurl,omitempty"`
}

type (
	FormatDefinitions map[string]Format
	MiniFormat        struct {
		ShortName string
		FullName  string
	}
)

const (
	FormatState   = "state"
	FormatOsmPbf  = "osm.pbf"
	FormatOsmGz   = "osm.gz"
	FormatOsmBz2  = "osm.bz2"
	FormatOshPbf  = "osh.pbf"
	FormatPoly    = "poly"
	FormatShpZip  = "shp.zip"
	FormatKml     = "kml"
	FormatGeoJSON = "geojson"
)

// MiniFormats get formats of an Element
//
//	and return a string
//	according to download-geofabrik short flags.
func MiniFormats(miniFormat []string) string {
	miniFormatList := []MiniFormat{
		{ShortName: "s", FullName: FormatState},
		{ShortName: "P", FullName: FormatOsmPbf},
		{ShortName: "G", FullName: FormatOsmGz},
		{ShortName: "B", FullName: FormatOsmBz2},
		{ShortName: "H", FullName: FormatOshPbf},
		{ShortName: "p", FullName: FormatPoly},
		{ShortName: "S", FullName: FormatShpZip},
		{ShortName: "k", FullName: FormatKml},
		{ShortName: "g", FullName: FormatGeoJSON},
	}

	res := make([]string, len(miniFormatList))

	for _, item := range miniFormat {
		for i, format := range miniFormatList {
			if item == format.FullName {
				res[i] = format.ShortName
			}
		}
	}

	return strings.Join(res, "")
}

// GetFormats return a pointer to a slice with formats.
func GetFormats() *[]string {
	var formatFile []string

	options := map[string]string{
		"dosmPbf":  FormatOsmPbf,
		"doshPbf":  FormatOshPbf,
		"dosmGz":   FormatOsmGz,
		"dosmBz2":  FormatOsmBz2,
		"dshpZip":  FormatShpZip,
		"dstate":   FormatState,
		"dpoly":    FormatPoly,
		"dkml":     FormatKml,
		"dgeojson": FormatGeoJSON,
	}

	for k, v := range options {
		if viper.GetBool(k) {
			formatFile = append(formatFile, v)
		}
	}

	if len(formatFile) == 0 {
		formatFile = append(formatFile, FormatOsmPbf)
	}

	slices.Sort(formatFile)

	return &formatFile
}
