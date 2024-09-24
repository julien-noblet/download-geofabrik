package app

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/julien-noblet/download-geofabrik/formats"
)

const (
	defaultConfigFile = "./geofabrik.yml"
)

type App struct {
	App         *kingpin.Application
	fConfig     *string
	fNodownload *bool
	fVerbose    *bool
	fQuiet      *bool
	fProgress   *bool
	Cdownload   *kingpin.CmdClause
	delement    *string
	dCheck      *bool

	// Output directory
	dOutputDir *string

	// Formats
	dosmBz2         *bool
	dosmGz          *bool
	dshpZip         *bool
	dosmPbf         *bool
	doshPbf         *bool
	dstate          *bool
	dpoly           *bool
	dkml            *bool
	dgeojson        *bool
	dgarmin         *bool
	dmaps           *bool
	dmbtiles        *bool
	dcsv            *bool
	dgarminonroad   *bool
	dgarminontrail  *bool
	dgarminopentopo *bool

	Cgenerate *kingpin.CmdClause

	Clist *kingpin.CmdClause
	lmd   *bool

	fService *string
}

func NewApp() *App {
	myApp := &App{
		App: kingpin.New("download-geofabrik", "A command-line tool for downloading OSM files."),
	}

	myApp.fConfig = myApp.App.Flag("config", "Set Config file.").Default(defaultConfigFile).Short('c').String()
	myApp.fNodownload = myApp.App.Flag("nodownload", "Do not download file (test only)").Short('n').Bool()
	myApp.fVerbose = myApp.App.Flag("verbose", "Be verbose").Short('v').Bool()
	myApp.fQuiet = myApp.App.Flag("quiet", "Be quiet").Short('q').Bool()
	myApp.fProgress = myApp.App.Flag("progress", "Add a progress bar (implies quiet)").Bool()

	myApp.Cdownload = myApp.App.Command("download", "Download element")
	myApp.delement = myApp.Cdownload.Arg("element", "OSM element").Required().String()
	myApp.dosmBz2 = myApp.Cdownload.Flag(formats.FormatOsmBz2, "Download osm.bz2 if available").Short('B').Bool()
	myApp.dosmGz = myApp.Cdownload.Flag(formats.FormatOsmGz, "Download osm.gz if available").Short('G').Bool()
	myApp.dshpZip = myApp.Cdownload.Flag(formats.FormatShpZip, "Download shp.zip if available").Short('S').Bool()
	myApp.dosmPbf = myApp.Cdownload.Flag(formats.FormatOsmPbf, "Download osm.pbf (default)").Short('P').Bool()
	myApp.doshPbf = myApp.Cdownload.Flag(formats.FormatOshPbf, "Download osh.pbf").Short('H').Bool()
	myApp.dstate = myApp.Cdownload.Flag(formats.FormatState, "Download state.txt file").Short('s').Bool()
	myApp.dpoly = myApp.Cdownload.Flag(formats.FormatPoly, "Download poly file").Short('p').Bool()
	myApp.dkml = myApp.Cdownload.Flag(formats.FormatKml, "Download kml file").Short('k').Bool()
	myApp.dgeojson = myApp.Cdownload.Flag(formats.FormatGeoJSON, "Download GeoJSON file").Short('g').Bool()
	myApp.dgarmin = myApp.Cdownload.Flag("garmin-osm", "Download Garmin OSM file").Short('O').Bool()
	myApp.dmaps = myApp.Cdownload.Flag("mapsforge", "Download Mapsforge file").Short('m').Bool()
	myApp.dmbtiles = myApp.Cdownload.Flag("mbtiles", "Download MBTiles file").Short('M').Bool()
	myApp.dcsv = myApp.Cdownload.Flag("csv", "Download CSV file").Short('C').Bool()
	myApp.dgarminonroad = myApp.Cdownload.Flag("garminonroad", "Download Garmin Onroad file").Short('r').Bool()
	myApp.dgarminontrail = myApp.Cdownload.Flag("garminontrail", "Download Garmin Ontrail file").Short('t').Bool()
	myApp.dgarminopentopo = myApp.Cdownload.Flag("garminopenTopo", "Download Garmin OpenTopo file").Short('o').Bool()
	myApp.dCheck = myApp.Cdownload.Flag("check", "Control with checksum (default) Use --no-check to discard control").
		Default("true").Bool()
	myApp.dOutputDir = myApp.Cdownload.Flag(
		"output_directory",
		"Set output directory, you can use also OUTPUT_DIR env variable",
	).
		Short('d').String()

	myApp.Cgenerate = myApp.App.Command("generate", "Generate a new config file")

	myApp.Clist = myApp.App.Command("list", "Show elements available")
	myApp.lmd = myApp.Clist.Flag("markdown", "generate list in Markdown format").Bool()

	myApp.fService = myApp.App.Flag("service",
		"Can switch to another service. "+
			"You can use \"geofabrik\", \"openstreetmap.fr\" \"osmtoday\" or \"bbbike\". "+
			"It automatically change config file if -c is unused.").
		Default("geofabrik").String()

	return myApp
}
