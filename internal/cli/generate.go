package cli

import (
	"log/slog"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/internal/generator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	service          string
	generateProgress bool
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate configuration file",
	RunE:  runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(
		&service,
		"service",
		"s",
		"geofabrik",
		"Service to use (geofabrik, geofabrik-parse, openstreetmap.fr, osmtoday, bbbike)",
	)
	generateCmd.Flags().BoolVarP(&generateProgress, "progress", "p", true, "Show progress bar")

	// Bind to viper if needed, or just pass flags manually.
	// Logic in generator.go was using Viper.
	// I refactored generator.go to use args.
}

func runGenerate(_ *cobra.Command, _ []string) error {
	cfgFile := viper.GetString("config")
	if cfgFile == "" {
		cfgFile = config.DefaultConfigFile
	}

	slog.Info("Generating config", "service", service, "file", cfgFile)

	if err := generator.Generate(service, generateProgress, cfgFile); err != nil {
		slog.Error("Generation failed", "error", err)
		return err
	}

	return nil
}
