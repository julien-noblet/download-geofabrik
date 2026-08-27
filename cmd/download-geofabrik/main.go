package main

import (
	"os"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
)

var version = "dev"

func main() {
	cli.Version = version

	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
