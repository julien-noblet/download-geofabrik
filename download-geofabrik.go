package main

import (
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/julien-noblet/download-geofabrik/app"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/generator"
	"github.com/julien-noblet/download-geofabrik/lists"
	"github.com/spf13/viper"
)

var version = "devel"

func main() {
	myApp := app.NewApp()
	myApp.App.Version(version) // Add version flag
	commands := kingpin.MustParse(myApp.App.Parse(os.Args[1:]))

	myApp.ConfigureViper()
	app.ConfigureLogging()
	app.SetOutputDir()

	app.CheckService()

	switch commands {
	case myApp.Clist.FullCommand():
		lists.ListCommand()
	case myApp.Cdownload.FullCommand():
		app.DownloadCommand()
	case myApp.Cgenerate.FullCommand():
		generator.Generate(viper.GetString(config.ViperConfig))
	}
}
