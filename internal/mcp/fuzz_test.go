package mcp_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/mcp"
	mcpSDK "github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/require"
)

func FuzzMCPHandleMessage(f *testing.F) {
	srv := mcp.NewServer(testVersion)
	ctx := context.Background()

	f.Add([]byte(`{"jsonrpc":"2.0","id":1,"method":"tools/list"}`))
	f.Add([]byte(`{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"list_services","arguments":{}}}`))
	f.Add([]byte(`{"jsonrpc":"2.0","id":3,"method":"resources/list"}`))
	f.Add([]byte(`{}`))
	f.Add([]byte(`{"jsonrpc":"2.0"`))
	f.Add([]byte("\x00\xff\xfe"))

	f.Fuzz(func(_ *testing.T, rawMsg []byte) {
		_ = srv.MCPServer().HandleMessage(ctx, json.RawMessage(rawMsg))
	})
}

func FuzzListElementsTool(f *testing.F) {
	srv := mcp.NewServer(testVersion)
	ctx := context.Background()

	catFile := createTestCatalog(f)

	tool := srv.MCPServer().GetTool("list_elements")
	require.NotNil(f, tool)

	f.Add("france", "europe", 10, 0)
	f.Add("", "", 50, 0)
	f.Add("non-existent", "unknown", -5, -1)
	f.Add("../../../", "../../", 1000, 500)
	f.Add("\x00\xff", "\x00", 0, 0)

	f.Fuzz(func(_ *testing.T, search string, parent string, limit int, offset int) {
		callReq := mcpSDK.CallToolRequest{
			Params: mcpSDK.CallToolParams{
				Name: "list_elements",
				Arguments: map[string]any{
					"service":     serviceDefault,
					"config_file": catFile,
					"search":      search,
					"parent":      parent,
					"limit":       limit,
					"offset":      offset,
				},
			},
		}

		_, _ = tool.Handler(ctx, callReq)
	})
}

func FuzzGetElementTool(f *testing.F) {
	srv := mcp.NewServer(testVersion)
	ctx := context.Background()

	catFile := createTestCatalog(f)

	tool := srv.MCPServer().GetTool("get_element")
	require.NotNil(f, tool)

	f.Add("france", serviceDefault)
	f.Add("monaco", serviceDefault)
	f.Add("", "")
	f.Add("../../../etc/passwd", "unknown-service")
	f.Add("\x00", "bbbike")

	f.Fuzz(func(_ *testing.T, elementID string, service string) {
		callReq := mcpSDK.CallToolRequest{
			Params: mcpSDK.CallToolParams{
				Name: "get_element",
				Arguments: map[string]any{
					"service":     service,
					"element_id":  elementID,
					"config_file": catFile,
				},
			},
		}

		_, _ = tool.Handler(ctx, callReq)
	})
}

func FuzzDownloadElementTool(f *testing.F) {
	srv := mcp.NewServer(testVersion)
	ctx := context.Background()

	catFile := createTestCatalog(f)

	tool := srv.MCPServer().GetTool("download_element")
	require.NotNil(f, tool)

	f.Add("france", "osm.pbf", "/tmp")
	f.Add("monaco", "poly", "")
	f.Add("", "", "")
	f.Add("../../../etc/passwd", "unknown", "../../evil")
	f.Add("\x00", "\xff", "/invalid\x00path")

	f.Fuzz(func(_ *testing.T, elementID string, format string, outputDir string) {
		callReq := mcpSDK.CallToolRequest{
			Params: mcpSDK.CallToolParams{
				Name: "download_element",
				Arguments: map[string]any{
					"service":     serviceDefault,
					"element_id":  elementID,
					"config_file": catFile,
					"formats":     []any{format},
					"output_dir":  outputDir,
					"dry_run":     true,
				},
			},
		}

		_, _ = tool.Handler(ctx, callReq)
	})
}

func FuzzListFormatsTool(f *testing.F) {
	srv := mcp.NewServer(testVersion)
	ctx := context.Background()

	catFile := createTestCatalog(f)

	tool := srv.MCPServer().GetTool("list_formats")
	require.NotNil(f, tool)

	f.Add(serviceDefault)
	f.Add("bbbike")
	f.Add("")
	f.Add("non-existent")
	f.Add("../../../")

	f.Fuzz(func(_ *testing.T, service string) {
		callReq := mcpSDK.CallToolRequest{
			Params: mcpSDK.CallToolParams{
				Name: "list_formats",
				Arguments: map[string]any{
					"service":     service,
					"config_file": catFile,
				},
			},
		}

		_, _ = tool.Handler(ctx, callReq)
	})
}
