package main

import (
	"fmt"
	"os"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
)

var version = "dev"

func main() {
	cli.Version = version

	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
