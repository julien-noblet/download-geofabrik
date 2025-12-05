package main

import (
	"fmt"
	"os"

	"github.com/julien-noblet/download-geofabrik/app"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/generator"
	"github.com/julien-noblet/download-geofabrik/lists"
	"github.com/spf13/viper"
)

var version = "devel"

func main() {
	if err := Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

// Run executes the application logic with the provided arguments.
func Run(args []string) error {
	myApp := app.NewApp()
	myApp.App.Version(version) // Add version flag

	commands, err := myApp.App.Parse(args)
	if err != nil {
		return fmt.Errorf("unable to parse arguments: %w", err)
	}

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

	return nil
}
