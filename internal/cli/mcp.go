package cli

import (
	"fmt"

	"github.com/julien-noblet/download-geofabrik/internal/mcp"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start Model Context Protocol (MCP) server for LLM integration",
	Long: `Start the Model Context Protocol (MCP) server communicating over standard I/O (stdio).
This allows LLMs (Claude Desktop, Antigravity, Cursor, etc.) to query OSM catalogs,
inspect formats, trigger catalog regeneration, and download extracts.`,
	RunE: runMCP,
}

// RegisterMCPCmd adds the mcp command to rootCmd.
func RegisterMCPCmd() {
	rootCmd.AddCommand(mcpCmd)
}

func runMCP(cmd *cobra.Command, _ []string) error {
	srv := mcp.NewServer(Version)

	if err := srv.ServeStdio(cmd.Context()); err != nil {
		return fmt.Errorf("failed to run MCP server: %w", err)
	}

	return nil
}
