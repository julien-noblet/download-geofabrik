package formats

import (
	"strings"

	"github.com/spf13/viper"
)

type Format struct {
	ID       string `yaml:"ext"`
	Loc      string `yaml:"loc"`
	BasePath string `yaml:"basepath,omitempty"`
	BaseURL  string `yaml:"baseurl,omitempty"`
}

type FormatDefinitions map[string]Format

const (
	FormatState  = "state"
	FormatOsmPbf = "osm.pbf"
	FormatOsmGz  = "osm.gz"
	FormatOsmBz2 = "osm.bz2"
	FormatOshPbf = "osh.pbf"
	FormatPoly   = "poly"
	FormatShpZip = "shp.zip"
	FormatKml    = "kml"
)

// MiniFormats get formats of an Element
//  and return a string
//  according to download-geofabrik short flags.
func MiniFormats(s []string) string {
	res := make([]string, 7) //nolint:gomnd // 7 is the number of formats

	for _, item := range s {
		switch item {
		case FormatState:
			res[0] = "s"
		case FormatOsmPbf:
			res[1] = "P"
		case FormatOsmGz:
			res[2] = "G"
		case FormatOsmBz2:
			res[2] = "B"
		case FormatOshPbf:
			res[3] = "H"
		case FormatPoly:
			res[4] = "p"
		case FormatShpZip:
			res[5] = "S"
		case FormatKml:
			res[6] = "k"
		}
	}

	return strings.Join(res, "")
}

// GetFormats return a pointer to a slice with formats
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

	if len(formatFile) == 0 {
		formatFile = append(formatFile, FormatOsmPbf)
	}

	return &formatFile
}
