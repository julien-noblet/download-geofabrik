package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "download-geofabrik",
	Short: "A command-line tool for downloading OSM files",
	Long:  `download-geofabrik is a CLI tool that downloads OpenStreetMap data from Geofabrik.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	RegisterDownloadCmd()
	RegisterGenerateCmd()
	RegisterListCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
