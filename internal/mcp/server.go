package mcp

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/julien-noblet/download-geofabrik/internal/provider"
	"github.com/mark3labs/mcp-go/server"
)

// Server encapsulates the Model Context Protocol (MCP) server for download-geofabrik.
type Server struct {
	mcpServer *server.MCPServer
	version   string
}

// NewServer initializes a new download-geofabrik MCP server with all tools and resources.
func NewServer(version string) *Server {
	provider.RegisterDefaultProviders()

	mcpServer := server.NewMCPServer(
		"download-geofabrik",
		version,
		server.WithToolCapabilities(false),
		server.WithResourceCapabilities(false, false),
		server.WithInstructions("MCP server for download-geofabrik. "+
			"Allows querying OpenStreetMap catalogs across multiple providers, "+
			"listing available data formats, regenerating catalogs, and downloading extracts."),
	)

	serverInstance := &Server{
		mcpServer: mcpServer,
		version:   version,
	}

	serverInstance.registerTools()
	serverInstance.registerResources()

	return serverInstance
}

// ServeStdio starts serving the MCP protocol over standard I/O (stdin/stdout).
func (s *Server) ServeStdio(_ context.Context) error {
	slog.Info("Starting download-geofabrik MCP server on stdio", "version", s.version)

	if err := server.ServeStdio(s.mcpServer); err != nil {
		return fmt.Errorf("mcp stdio server failure: %w", err)
	}

	return nil
}

// MCPServer returns the underlying mark3labs MCPServer instance (useful for testing).
func (s *Server) MCPServer() *server.MCPServer {
	return s.mcpServer
}
