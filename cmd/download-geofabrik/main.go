package main

import (
	"fmt"
	"io"
	"os"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
)

var version = "dev"

func run(stderr io.Writer, execute func() error) int {
	cli.Version = version

	if err := execute(); err != nil {
		fmt.Fprintf(stderr, "Error: %v\n", err)

		return 1
	}

	return 0
}

func main() {
	if code := run(os.Stderr, cli.Execute); code != 0 {
		os.Exit(code)
	}
}
