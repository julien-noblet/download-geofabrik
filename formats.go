package main

import "strings"

type format struct {
	ID       string `yaml:"ext"`
	Loc      string `yaml:"loc"`
	BasePath string `yaml:"basepath,omitempty"`
}

func miniFormats(s []string) string {
	res := make([]string, 7)
	for _, item := range s {
		switch item {
		case "state":
			res[0] = "s"
		case "osm.pbf":
			res[1] = "P"
		case "osm.bz2":
			res[2] = "B"
		case "osh.pbf":
			res[3] = "H"
		case "poly":
			res[4] = "p"
		case "shp.zip":
			res[5] = "S"
		case "kml":
			res[6] = "k"
		}
	}
	return strings.Join(res, "")
}

func getFormats() *[]string {
	var formatFile []string
	if *dosmPbf {
		formatFile = append(formatFile, "osm.pbf")
	}
	if *doshPbf {
		formatFile = append(formatFile, "osh.pbf")
	}
	if *dosmBz2 {
		formatFile = append(formatFile, "osm.bz2")
	}
	if *dshpZip {
		formatFile = append(formatFile, "shp.zip")
	}
	if *dstate {
		formatFile = append(formatFile, "state")
	}
	if *dpoly {
		formatFile = append(formatFile, "poly")
	}
	if *dkml {
		formatFile = append(formatFile, "kml")
	}
	if len(formatFile) == 0 {
		formatFile = append(formatFile, "osm.pbf")
	}
	return &formatFile
}
