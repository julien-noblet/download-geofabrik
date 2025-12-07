package cli

import (
	"log/slog"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/lists"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	markdown bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show elements available",
	RunE:  runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&markdown, "markdown", false, "Generate list in Markdown format")
}

func runList(cmd *cobra.Command, args []string) error {
	cfgFile := viper.GetString("config")
	if cfgFile == "" {
		cfgFile = config.DefaultConfigFile
	}

	opts := &config.Options{
		ConfigFile: cfgFile,
	}

	cfg, err := config.LoadConfig(opts.ConfigFile)
	if err != nil {
		slog.Error("Failed to load config", "file", opts.ConfigFile, "error", err)
		return err
	}

	format := ""
	if markdown {
		format = lists.MarkdownFormat
	}

	if err := lists.ListAllRegions(cfg, format); err != nil {
		slog.Error("Failed to list all regions", "error", err)
		return err
	}

	return nil
}
