package mcp

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/mark3labs/mcp-go/server"

	"github.com/julien-noblet/download-geofabrik/internal/provider"
)

// Server encapsulates the Model Context Protocol (MCP) server for download-geofabrik.
type Server struct {
	mcpServer *server.MCPServer
	version   string
	genMu     sync.Mutex
}

// ErrServerNotInitialized is returned when ServeStdio is called on a nil or uninitialized Server.
var ErrServerNotInitialized = errors.New("mcp server not initialized: use mcp.NewServer")

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
	if s == nil || s.mcpServer == nil {
		return ErrServerNotInitialized
	}

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
