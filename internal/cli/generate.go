package cli

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/julien-noblet/download-geofabrik/internal/generator"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var generateProgress bool

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate configuration file",
	RunE:  runGenerate,
}

// RegisterGenerateCmd registers the generate command and its flags to rootCmd.
func RegisterGenerateCmd() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().BoolVarP(&generateProgress, "progress", "p", true, "Show progress bar")
}

func runGenerate(cmd *cobra.Command, _ []string) error {
	cfgFile := viper.ConfigFileUsed()
	if cfgFile == "" {
		if service != "" {
			cfgFile = service + ".yml"
		} else {
			cfgFile = catalog.DefaultConfigFile
		}
	}

	slog.Info("Generating config", "service", service, "file", cfgFile)

	if err := generator.Generate(cmd.Context(), service, cfgFile); err != nil {
		return fmt.Errorf("generation failed: %w", err)
	}

	return nil
}
