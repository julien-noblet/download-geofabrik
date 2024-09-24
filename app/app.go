package app

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/julien-noblet/download-geofabrik/formats"
)

const (
	defaultConfigFile = "./geofabrik.yml"
	serviceGeofabrik  = "geofabrik"
	serviceOSMFr      = "openstreetmap.fr"
	serviceOSMToday   = "osmtoday"
	serviceBBBike     = "bbbike"
)

// App encapsulates the application state and configuration.
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
	dOutputDir  *string

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
	Clist     *kingpin.CmdClause
	lmd       *bool
	fService  *string
}

// NewApp initializes a new App instance with default settings.
func NewApp() *App {
	myApp := &App{
		App: kingpin.New("download-geofabrik", "A command-line tool for downloading OSM files."),
	}

	myApp.initFlags()
	myApp.initCommands()

	return myApp
}

// initFlags initializes the flags for the application.
func (a *App) initFlags() {
	a.fConfig = a.App.Flag("config", "Set Config file.").Default(defaultConfigFile).Short('c').String()
	a.fNodownload = a.App.Flag("nodownload", "Do not download file (test only)").Short('n').Bool()
	a.fVerbose = a.App.Flag("verbose", "Be verbose").Short('v').Bool()
	a.fQuiet = a.App.Flag("quiet", "Be quiet").Short('q').Bool()
	a.fProgress = a.App.Flag("progress", "Add a progress bar (implies quiet)").Bool()
	a.fService = a.App.Flag("service",
		"Can switch to another service. "+
			"You can use \"geofabrik\", \"openstreetmap.fr\" \"osmtoday\" or \"bbbike\". "+
			"It automatically change config file if -c is unused.").
		Default(serviceGeofabrik).String()
}

// initCommands initializes the commands for the application.
func (a *App) initCommands() {
	a.Cdownload = a.App.Command("download", "Download element")
	a.delement = a.Cdownload.Arg("element", "OSM element").Required().String()
	a.dCheck = a.Cdownload.Flag("check", "Control with checksum (default) Use --no-check to discard control").
		Default("true").Bool()
	a.dOutputDir = a.Cdownload.Flag(
		"output_directory",
		"Set output directory, you can use also OUTPUT_DIR env variable",
	).Short('d').String()

	a.initFormatFlags()

	a.Cgenerate = a.App.Command("generate", "Generate a new config file")
	a.Clist = a.App.Command("list", "Show elements available")
	a.lmd = a.Clist.Flag("markdown", "generate list in Markdown format").Bool()
}

// initFormatFlags initializes the format flags for the download command.
func (a *App) initFormatFlags() {
	a.dosmBz2 = a.Cdownload.Flag(formats.FormatOsmBz2, "Download osm.bz2 if available").Short('B').Bool()
	a.dosmGz = a.Cdownload.Flag(formats.FormatOsmGz, "Download osm.gz if available").Short('G').Bool()
	a.dshpZip = a.Cdownload.Flag(formats.FormatShpZip, "Download shp.zip if available").Short('S').Bool()
	a.dosmPbf = a.Cdownload.Flag(formats.FormatOsmPbf, "Download osm.pbf (default)").Short('P').Bool()
	a.doshPbf = a.Cdownload.Flag(formats.FormatOshPbf, "Download osh.pbf").Short('H').Bool()
	a.dstate = a.Cdownload.Flag(formats.FormatState, "Download state.txt file").Short('s').Bool()
	a.dpoly = a.Cdownload.Flag(formats.FormatPoly, "Download poly file").Short('p').Bool()
	a.dkml = a.Cdownload.Flag(formats.FormatKml, "Download kml file").Short('k').Bool()
	a.dgeojson = a.Cdownload.Flag(formats.FormatGeoJSON, "Download GeoJSON file").Short('g').Bool()
	a.dgarmin = a.Cdownload.Flag("garmin-osm", "Download Garmin OSM file").Short('O').Bool()
	a.dmaps = a.Cdownload.Flag("mapsforge", "Download Mapsforge file").Short('m').Bool()
	a.dmbtiles = a.Cdownload.Flag("mbtiles", "Download MBTiles file").Short('M').Bool()
	a.dcsv = a.Cdownload.Flag("csv", "Download CSV file").Short('C').Bool()
	a.dgarminonroad = a.Cdownload.Flag("garminonroad", "Download Garmin Onroad file").Short('r').Bool()
	a.dgarminontrail = a.Cdownload.Flag("garminontrail", "Download Garmin Ontrail file").Short('t').Bool()
	a.dgarminopentopo = a.Cdownload.Flag("garminopenTopo", "Download Garmin OpenTopo file").Short('o').Bool()
}
