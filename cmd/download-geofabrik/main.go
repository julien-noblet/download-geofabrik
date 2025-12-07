package main

import (
	"os"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
