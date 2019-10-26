package main

import (
	"strings"
)

type format struct {
	ID       string `yaml:"ext"`
	Loc      string `yaml:"loc"`
	BasePath string `yaml:"basepath,omitempty"`
	BaseURL  string `yaml:"baseurl,omitempty"`
}

type formatDefinitions map[string]format

const (
	formatState  = "state"
	formatOsmPbf = "osm.pbf"
	formatOsmGz  = "osm.gz"
	formatOsmBz2 = "osm.bz2"
	formatOshPbf = "osh.pbf"
	formatPoly   = "poly"
	formatShpZip = "shp.zip"
	formatKml    = "kml"
)

// miniFormats get formats of an Element
//  and return a string
//  according to download-geofabrik short flags.
func miniFormats(s []string) string {
	res := make([]string, 7)

	for _, item := range s {
		switch item {
		case formatState:
			res[0] = "s"
		case formatOsmPbf:
			res[1] = "P"
		case formatOsmGz:
			res[2] = "G"
		case formatOsmBz2:
			res[2] = "B"
		case formatOshPbf:
			res[3] = "H"
		case formatPoly:
			res[4] = "p"
		case formatShpZip:
			res[5] = "S"
		case formatKml:
			res[6] = "k"
		}
	}

	return strings.Join(res, "")
}

func isHashable(c *Config, format string) (hashable bool, hash, hashType string) {
	hashs := []string{"md5"} // had to be globalized?

	if _, ok := c.Formats[format]; ok {
		for _, h := range hashs {
			hash := format + "." + h
			if _, ok := c.Formats[hash]; ok {
				return true, hash, h
			}
		}
	}

	return false, "", ""
}

// getFormats return a pointer to a slice with formats
func getFormats() *[]string {
	var formatFile []string
	if *dosmPbf {
		formatFile = append(formatFile, formatOsmPbf)
	}

	if *doshPbf {
		formatFile = append(formatFile, formatOshPbf)
	}

	if *dosmGz {
		formatFile = append(formatFile, formatOsmGz)
	}

	if *dosmBz2 {
		formatFile = append(formatFile, formatOsmBz2)
	}

	if *dshpZip {
		formatFile = append(formatFile, formatShpZip)
	}

	if *dstate {
		formatFile = append(formatFile, formatState)
	}

	if *dpoly {
		formatFile = append(formatFile, formatPoly)
	}

	if *dkml {
		formatFile = append(formatFile, formatKml)
	}

	if len(formatFile) == 0 {
		formatFile = append(formatFile, formatOsmPbf)
	}

	return &formatFile
}
