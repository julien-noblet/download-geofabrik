package mcp_test

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/mcp"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	mcpSDK "github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testVersion    = "1.0.0-test"
	serviceDefault = "geofabrik"
	elementFrance  = "france"
	elementMonaco  = "monaco"
	regionEurope   = "europe"
)

func createTestCatalog(t *testing.T) string {
	t.Helper()

	tempDir := t.TempDir()
	catPath := filepath.Join(tempDir, "test-catalog.yml")

	cat := catalog.New()
	cat.BaseURL = "https://example.com/osm"
	cat.Formats = catalog.FormatDefinitions{
		formats.FormatOsmPbf: {
			ID:   formats.FormatOsmPbf,
			Loc:  ".osm.pbf",
			Type: "OpenStreetMap binary format",
		},
		formats.FormatPoly: {
			ID:   formats.FormatPoly,
			Loc:  ".poly",
			Type: "Polygon filter file",
		},
		formats.FormatOsmPbf + ".md5": {
			ID:   formats.FormatOsmPbf + ".md5",
			Loc:  ".osm.pbf.md5",
			Type: "MD5 checksum",
		},
	}

	cat.Elements[regionEurope] = catalog.Element{
		ID:      regionEurope,
		Name:    "Europe",
		Formats: catalog.Formats{},
		Meta:    true,
	}

	cat.Elements[elementFrance] = catalog.Element{
		ID:      elementFrance,
		Name:    "France",
		Parent:  regionEurope,
		Formats: catalog.Formats{formats.FormatOsmPbf, formats.FormatPoly},
	}

	cat.Elements[elementMonaco] = catalog.Element{
		ID:      elementMonaco,
		Name:    "Monaco",
		Parent:  regionEurope,
		Formats: catalog.Formats{formats.FormatOsmPbf},
	}

	require.NoError(t, cat.SaveFile(catPath))

	return catPath
}

func TestNewServer(t *testing.T) {
	t.Parallel()

	srv := mcp.NewServer(testVersion)
	require.NotNil(t, srv)
	assert.NotNil(t, srv.MCPServer())

	tools := srv.MCPServer().ListTools()
	assert.Contains(t, tools, "list_services")
	assert.Contains(t, tools, "regenerate_catalog")
	assert.Contains(t, tools, "list_elements")
	assert.Contains(t, tools, "get_element")
	assert.Contains(t, tools, "list_formats")
	assert.Contains(t, tools, "download_element")
}

func TestListServicesTool(t *testing.T) {
	t.Parallel()

	srv := mcp.NewServer(testVersion)
	ctx := context.Background()

	tool := srv.MCPServer().GetTool("list_services")
	require.NotNil(t, tool)

	callReq := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name: "list_services",
			Arguments: map[string]any{
				"check_local": true,
			},
		},
	}

	res, err := tool.Handler(ctx, callReq)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.False(t, res.IsError)

	textContent, ok := res.Content[0].(mcpSDK.TextContent)
	require.True(t, ok)

	var services []mcp.ServiceInfo

	err = json.Unmarshal([]byte(textContent.Text), &services)
	require.NoError(t, err)
	assert.NotEmpty(t, services)

	foundGeofabrik := false

	for _, s := range services {
		if s.Name == serviceDefault {
			foundGeofabrik = true

			assert.Equal(t, "geofabrik.yml", s.DefaultConfigFile)

			break
		}
	}

	assert.True(t, foundGeofabrik, "geofabrik service should be in the list")
}

func TestListFormatsTool(t *testing.T) {
	t.Parallel()

	srv := mcp.NewServer(testVersion)
	ctx := context.Background()

	tool := srv.MCPServer().GetTool("list_formats")
	require.NotNil(t, tool)

	callReq := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name:      "list_formats",
			Arguments: map[string]any{},
		},
	}

	res, err := tool.Handler(ctx, callReq)
	require.NoError(t, err)
	require.False(t, res.IsError)

	textContent, ok := res.Content[0].(mcpSDK.TextContent)
	require.True(t, ok)

	var specs []mcp.FormatSpecification

	err = json.Unmarshal([]byte(textContent.Text), &specs)
	require.NoError(t, err)
	assert.NotEmpty(t, specs)

	catFile := createTestCatalog(t)
	callReqWithCatalog := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name: "list_formats",
			Arguments: map[string]any{
				"service":     serviceDefault,
				"config_file": catFile,
			},
		},
	}

	resWithCat, err := tool.Handler(ctx, callReqWithCatalog)
	require.NoError(t, err)
	require.False(t, resWithCat.IsError)
}

func TestListElementsTool(t *testing.T) {
	t.Parallel()

	srv := mcp.NewServer(testVersion)
	ctx := context.Background()
	catFile := createTestCatalog(t)

	tool := srv.MCPServer().GetTool("list_elements")
	require.NotNil(t, tool)

	callReq := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name: "list_elements",
			Arguments: map[string]any{
				"service":     serviceDefault,
				"config_file": catFile,
				"limit":       10,
				"offset":      0,
			},
		},
	}

	res, err := tool.Handler(ctx, callReq)
	require.NoError(t, err)
	require.False(t, res.IsError)

	textContent, ok := res.Content[0].(mcpSDK.TextContent)
	require.True(t, ok)

	var listRes mcp.ListElementsResult

	err = json.Unmarshal([]byte(textContent.Text), &listRes)
	require.NoError(t, err)
	assert.Equal(t, 3, listRes.TotalElements)
	assert.Equal(t, 3, listRes.TotalMatched)
	assert.Len(t, listRes.Elements, 3)

	callReqSearch := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name: "list_elements",
			Arguments: map[string]any{
				"service":     serviceDefault,
				"config_file": catFile,
				"search":      elementMonaco,
			},
		},
	}

	resSearch, err := tool.Handler(ctx, callReqSearch)
	require.NoError(t, err)
	require.False(t, resSearch.IsError)

	textContentSearch, ok := resSearch.Content[0].(mcpSDK.TextContent)
	require.True(t, ok)

	var searchRes mcp.ListElementsResult

	err = json.Unmarshal([]byte(textContentSearch.Text), &searchRes)
	require.NoError(t, err)
	assert.Equal(t, 1, searchRes.TotalMatched)
	require.Len(t, searchRes.Elements, 1)
	assert.Equal(t, elementMonaco, searchRes.Elements[0].ID)

	callReqParent := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name: "list_elements",
			Arguments: map[string]any{
				"service":     serviceDefault,
				"config_file": catFile,
				"parent":      regionEurope,
			},
		},
	}

	resParent, err := tool.Handler(ctx, callReqParent)
	require.NoError(t, err)
	require.False(t, resParent.IsError)

	textContentParent, ok := resParent.Content[0].(mcpSDK.TextContent)
	require.True(t, ok)

	var parentRes mcp.ListElementsResult

	err = json.Unmarshal([]byte(textContentParent.Text), &parentRes)
	require.NoError(t, err)
	assert.Equal(t, 2, parentRes.TotalMatched)
}

func TestGetElementTool(t *testing.T) {
	t.Parallel()

	srv := mcp.NewServer(testVersion)
	ctx := context.Background()
	catFile := createTestCatalog(t)

	tool := srv.MCPServer().GetTool("get_element")
	require.NotNil(t, tool)

	callReq := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name: "get_element",
			Arguments: map[string]any{
				"service":     serviceDefault,
				"element_id":  elementFrance,
				"config_file": catFile,
			},
		},
	}

	res, err := tool.Handler(ctx, callReq)
	require.NoError(t, err)
	require.False(t, res.IsError)

	textContent, ok := res.Content[0].(mcpSDK.TextContent)
	require.True(t, ok)

	var detail mcp.ElementDetail

	err = json.Unmarshal([]byte(textContent.Text), &detail)
	require.NoError(t, err)
	assert.Equal(t, elementFrance, detail.ID)
	assert.Equal(t, "France", detail.Name)
	assert.Equal(t, regionEurope, detail.Parent)
	assert.NotEmpty(t, detail.Formats)

	var pbfDetail *mcp.FormatURLDetail

	for _, f := range detail.Formats {
		if f.FormatID == formats.FormatOsmPbf {
			pbfDetail = &f

			break
		}
	}

	require.NotNil(t, pbfDetail)
	assert.Equal(t, "https://example.com/osm/europe/france.osm.pbf", pbfDetail.DownloadURL)
	assert.True(t, pbfDetail.ChecksumAvailable)
	assert.Equal(t, "https://example.com/osm/europe/france.osm.pbf.md5", pbfDetail.ChecksumURL)

	callReqNotFound := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name: "get_element",
			Arguments: map[string]any{
				"service":     serviceDefault,
				"element_id":  "non-existent",
				"config_file": catFile,
			},
		},
	}

	resNotFound, err := tool.Handler(ctx, callReqNotFound)
	require.NoError(t, err)
	assert.True(t, resNotFound.IsError)

	callReqEmpty := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name: "get_element",
			Arguments: map[string]any{
				"element_id": "",
			},
		},
	}

	resEmpty, err := tool.Handler(ctx, callReqEmpty)
	require.NoError(t, err)
	assert.True(t, resEmpty.IsError)
}

func TestDownloadElementToolDryRun(t *testing.T) {
	t.Parallel()

	srv := mcp.NewServer(testVersion)
	ctx := context.Background()
	catFile := createTestCatalog(t)
	outDir := t.TempDir()

	tool := srv.MCPServer().GetTool("download_element")
	require.NotNil(t, tool)

	callReq := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name: "download_element",
			Arguments: map[string]any{
				"service":     serviceDefault,
				"element_id":  elementFrance,
				"formats":     []string{formats.FormatOsmPbf, formats.FormatPoly},
				"config_file": catFile,
				"output_dir":  outDir,
				"dry_run":     true,
			},
		},
	}

	res, err := tool.Handler(ctx, callReq)
	require.NoError(t, err)
	require.False(t, res.IsError)

	textContent, ok := res.Content[0].(mcpSDK.TextContent)
	require.True(t, ok)

	var dlResult mcp.DownloadResult

	err = json.Unmarshal([]byte(textContent.Text), &dlResult)
	require.NoError(t, err)

	assert.True(t, dlResult.DryRun)
	assert.Equal(t, elementFrance, dlResult.ElementID)
	assert.Len(t, dlResult.Files, 2)
	assert.Equal(t, "dry_run", dlResult.Files[0].Status)
	assert.Equal(t, "https://example.com/osm/europe/france.osm.pbf", dlResult.Files[0].DownloadURL)
	assert.Equal(t, "https://example.com/osm/europe/france.osm.pbf.md5", dlResult.Files[0].ChecksumURL)
}

func TestRegenerateCatalogToolErrors(t *testing.T) {
	t.Parallel()

	srv := mcp.NewServer(testVersion)
	ctx := context.Background()

	tool := srv.MCPServer().GetTool("regenerate_catalog")
	require.NotNil(t, tool)

	callReqMissing := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name:      "regenerate_catalog",
			Arguments: map[string]any{},
		},
	}

	resMissing, err := tool.Handler(ctx, callReqMissing)
	require.NoError(t, err)
	assert.True(t, resMissing.IsError)

	callReqUnknown := mcpSDK.CallToolRequest{
		Params: mcpSDK.CallToolParams{
			Name: "regenerate_catalog",
			Arguments: map[string]any{
				"service": "unknown-service-xyz",
			},
		},
	}

	resUnknown, err := tool.Handler(ctx, callReqUnknown)
	require.NoError(t, err)
	assert.True(t, resUnknown.IsError)
}

func TestResources(t *testing.T) {
	t.Parallel()

	srv := mcp.NewServer(testVersion)
	ctx := context.Background()

	resources := srv.MCPServer().ListResources()
	require.NotEmpty(t, resources)

	srvResource, exists := resources["geofabrik://services"]
	require.True(t, exists)

	reqServices := mcpSDK.ReadResourceRequest{
		Params: mcpSDK.ReadResourceParams{
			URI: "geofabrik://services",
		},
	}

	resServices, err := srvResource.Handler(ctx, reqServices)
	require.NoError(t, err)
	require.NotEmpty(t, resServices)

	fmtResource, exists := resources["geofabrik://formats"]
	require.True(t, exists)

	reqFormats := mcpSDK.ReadResourceRequest{
		Params: mcpSDK.ReadResourceParams{
			URI: "geofabrik://formats",
		},
	}

	resFormats, err := fmtResource.Handler(ctx, reqFormats)
	require.NoError(t, err)
	require.NotEmpty(t, resFormats)
}
