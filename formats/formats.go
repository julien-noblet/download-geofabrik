package formats

import (
	"strings"

	"github.com/spf13/viper"
)

type Format struct {
	Ext      string `yaml:"ext"`
	Loc      string `yaml:"loc"`
	BasePath string `yaml:"basepath,omitempty"`
	BaseURL  string `yaml:"baseurl,omitempty"`
}

type FormatDefinitions map[string]Format
type MiniFormat struct {
	ShortName string
	FullName  string
}

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
	if viper.GetBool("dosmPbf") {
		formatFile = append(formatFile, FormatOsmPbf)
	}

	if viper.GetBool("doshPbf") {
		formatFile = append(formatFile, FormatOshPbf)
	}

	if viper.GetBool("dosmGz") {
		formatFile = append(formatFile, FormatOsmGz)
	}

	if viper.GetBool("dosmBz2") {
		formatFile = append(formatFile, FormatOsmBz2)
	}

	if viper.GetBool("dshpZip") {
		formatFile = append(formatFile, FormatShpZip)
	}

	if viper.GetBool("dstate") {
		formatFile = append(formatFile, FormatState)
	}

	if viper.GetBool("dpoly") {
		formatFile = append(formatFile, FormatPoly)
	}

	if viper.GetBool("dkml") {
		formatFile = append(formatFile, FormatKml)
	}
	if viper.GetBool("dgeojson") {
		formatFile = append(formatFile, FormatGeoJSON)
	}
	if len(formatFile) == 0 {
		formatFile = append(formatFile, FormatOsmPbf)
	}

	return &formatFile
}
