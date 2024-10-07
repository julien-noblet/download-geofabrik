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
	FConfig     *string
	FNodownload *bool
	FVerbose    *bool
	FQuiet      *bool
	FProgress   *bool
	Cdownload   *kingpin.CmdClause
	Delement    *string
	DCheck      *bool
	DOutputDir  *string

	// Formats
	DosmBz2         *bool
	DosmGz          *bool
	DshpZip         *bool
	DosmPbf         *bool
	DoshPbf         *bool
	Dstate          *bool
	Dpoly           *bool
	Dkml            *bool
	Dgeojson        *bool
	Dgarmin         *bool
	Dmaps           *bool
	Dmbtiles        *bool
	Dcsv            *bool
	Dgarminonroad   *bool
	Dgarminontrail  *bool
	Dgarminopentopo *bool

	Cgenerate *kingpin.CmdClause
	Clist     *kingpin.CmdClause
	Lmd       *bool
	FService  *string
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
	a.FConfig = a.App.Flag("config", "Set Config file.").Default(defaultConfigFile).Short('c').String()
	a.FNodownload = a.App.Flag("nodownload", "Do not download file (test only)").Short('n').Bool()
	a.FVerbose = a.App.Flag("verbose", "Be verbose").Short('v').Bool()
	a.FQuiet = a.App.Flag("quiet", "Be quiet").Short('q').Bool()
	a.FProgress = a.App.Flag("progress", "Add a progress bar (implies quiet)").Bool()
	a.FService = a.App.Flag("service",
		"Can switch to another service. "+
			"You can use \"geofabrik\", \"openstreetmap.fr\" \"osmtoday\" or \"bbbike\". "+
			"It automatically change config file if -c is unused.").
		Default(serviceGeofabrik).String()
}

// initCommands initializes the commands for the application.
func (a *App) initCommands() {
	a.Cdownload = a.App.Command("download", "Download element")
	a.Delement = a.Cdownload.Arg("element", "OSM element").Required().String()
	a.DCheck = a.Cdownload.Flag("check", "Control with checksum (default) Use --no-check to discard control").
		Default("true").Bool()
	a.DOutputDir = a.Cdownload.Flag(
		"output_directory",
		"Set output directory, you can use also OUTPUT_DIR env variable",
	).Short('d').String()

	a.initFormatFlags()

	a.Cgenerate = a.App.Command("generate", "Generate a new config file")
	a.Clist = a.App.Command("list", "Show elements available")
	a.Lmd = a.Clist.Flag("markdown", "generate list in Markdown format").Bool()
}

// initFormatFlags initializes the format flags for the download command.
func (a *App) initFormatFlags() {
	a.DosmBz2 = a.Cdownload.Flag(formats.FormatOsmBz2, "Download osm.bz2 if available").Short('B').Bool()
	a.DosmGz = a.Cdownload.Flag(formats.FormatOsmGz, "Download osm.gz if available").Short('G').Bool()
	a.DshpZip = a.Cdownload.Flag(formats.FormatShpZip, "Download shp.zip if available").Short('S').Bool()
	a.DosmPbf = a.Cdownload.Flag(formats.FormatOsmPbf, "Download osm.pbf (default)").Short('P').Bool()
	a.DoshPbf = a.Cdownload.Flag(formats.FormatOshPbf, "Download osh.pbf").Short('H').Bool()
	a.Dstate = a.Cdownload.Flag(formats.FormatState, "Download state.txt file").Short('s').Bool()
	a.Dpoly = a.Cdownload.Flag(formats.FormatPoly, "Download poly file").Short('p').Bool()
	a.Dkml = a.Cdownload.Flag(formats.FormatKml, "Download kml file").Short('k').Bool()
	a.Dgeojson = a.Cdownload.Flag(formats.FormatGeoJSON, "Download GeoJSON file").Short('g').Bool()
	a.Dgarmin = a.Cdownload.Flag("garmin-osm", "Download Garmin OSM file").Short('O').Bool()
	a.Dmaps = a.Cdownload.Flag("mapsforge", "Download Mapsforge file").Short('m').Bool()
	a.Dmbtiles = a.Cdownload.Flag("mbtiles", "Download MBTiles file").Short('M').Bool()
	a.Dcsv = a.Cdownload.Flag("csv", "Download CSV file").Short('C').Bool()
	a.Dgarminonroad = a.Cdownload.Flag("garminonroad", "Download Garmin Onroad file").Short('r').Bool()
	a.Dgarminontrail = a.Cdownload.Flag("garminontrail", "Download Garmin Ontrail file").Short('t').Bool()
	a.Dgarminopentopo = a.Cdownload.Flag("garminopenTopo", "Download Garmin OpenTopo file").Short('o').Bool()
}
