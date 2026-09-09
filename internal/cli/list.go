package cli

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/julien-noblet/download-geofabrik/internal/ui"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var (
	markdown   bool
	jsonOutput bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show elements available",
	RunE:  runList,
}

// RegisterListCmd registers the list command and its flags to rootCmd.
func RegisterListCmd() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&markdown, "markdown", false, "Generate list in Markdown format")
	listCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output list in JSON format")
}

func runList(cmd *cobra.Command, _ []string) error {
	cfgFile := viper.ConfigFileUsed()
	if cfgFile == "" {
		if service != "" {
			cfgFile = service + ".yml"
		} else {
			cfgFile = catalog.DefaultConfigFile
		}
	}

	cat, err := catalog.LoadFile(cfgFile)
	if err != nil {
		slog.Error("Failed to load catalog", "file", cfgFile, "error", err)

		return fmt.Errorf("failed to load catalog: %w", err)
	}

	out := cmd.OutOrStdout()

	if jsonOutput {
		if err := ui.PrintJSON(cat, out); err != nil {
			return fmt.Errorf("failed to output JSON: %w", err)
		}

		return nil
	}

	if err := ui.PrintTable(cat, markdown, out); err != nil {
		slog.Error("Failed to render table", "error", err)

		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
