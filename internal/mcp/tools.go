package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	downloader "github.com/julien-noblet/download-geofabrik/internal/downloader"
	"github.com/julien-noblet/download-geofabrik/internal/generator"
	"github.com/julien-noblet/download-geofabrik/internal/provider"
	"github.com/julien-noblet/download-geofabrik/internal/provider/bbbike"
	"github.com/julien-noblet/download-geofabrik/internal/provider/geo2day"
	"github.com/julien-noblet/download-geofabrik/internal/provider/geofabrik"
	"github.com/julien-noblet/download-geofabrik/internal/provider/movisda"
	"github.com/julien-noblet/download-geofabrik/internal/provider/openstreetmapfr"
	"github.com/julien-noblet/download-geofabrik/internal/provider/osmch"
	"github.com/julien-noblet/download-geofabrik/internal/provider/osmfitvutbr"
	"github.com/julien-noblet/download-geofabrik/internal/provider/osmit"
	"github.com/julien-noblet/download-geofabrik/internal/provider/osmkewllu"
	"github.com/julien-noblet/download-geofabrik/internal/provider/osmtw"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	mcpSDK "github.com/mark3labs/mcp-go/mcp"
)

var (
	errInvalidService = errors.New("invalid or unknown service")
	errEmptyElementID = errors.New("element_id cannot be empty")
	errNilOrEmpty     = errors.New("nil element or empty format")
	errFormatMissing  = errors.New("format definition missing from catalog")
)

const (
	defaultListLimit = 50
	maxListLimit     = 500
	fieldStatus      = "status"
	statusSuccess    = "success"
	statusError      = "error"
)

// ServiceInfo describes an available OSM data provider.
// Field alignment optimized.
type ServiceInfo struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	DefaultConfigFile string `json:"defaultConfigFile"`
	HasLocalConfig    bool   `json:"hasLocalConfig"`
}

// ElementSummary summarizes an element for listing.
// Field alignment optimized.
type ElementSummary struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Parent  string   `json:"parent,omitempty"`
	Formats []string `json:"formats"`
	Meta    bool     `json:"meta,omitempty"`
}

// ListElementsResult contains paginated elements and metadata.
// Field alignment optimized.
type ListElementsResult struct {
	Service       string           `json:"service"`
	Elements      []ElementSummary `json:"elements"`
	TotalElements int              `json:"totalElements"`
	TotalMatched  int              `json:"totalMatched"`
	Limit         int              `json:"limit"`
	Offset        int              `json:"offset"`
}

// FormatURLDetail provides detailed URL and integrity information for a single format.
// Field alignment optimized.
type FormatURLDetail struct {
	FormatID          string `json:"formatId"`
	DownloadURL       string `json:"downloadUrl,omitempty"`
	ChecksumFormatID  string `json:"checksumFormatId,omitempty"`
	ChecksumURL       string `json:"checksumUrl,omitempty"`
	Error             string `json:"error,omitempty"`
	ChecksumAvailable bool   `json:"checksumAvailable"`
}

// ElementDetail represents the complete metadata and resolved URLs for an element.
// Field alignment optimized.
type ElementDetail struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Parent  string            `json:"parent,omitempty"`
	File    string            `json:"file,omitempty"`
	Service string            `json:"service"`
	Formats []FormatURLDetail `json:"formats"`
	Meta    bool              `json:"meta,omitempty"`
}

// DownloadFileResult reports the status of a single downloaded format.
// Field alignment optimized.
type DownloadFileResult struct {
	Format           string `json:"format"`
	TargetFile       string `json:"targetFile"`
	DownloadURL      string `json:"downloadUrl,omitempty"`
	ChecksumURL      string `json:"checksumUrl,omitempty"`
	Status           string `json:"status"`
	Error            string `json:"error,omitempty"`
	FileSizeBytes    int64  `json:"fileSizeBytes"`
	ChecksumVerified bool   `json:"checksumVerified"`
}

// DownloadResult reports the outcome of a download_element call.
// Field alignment optimized.
type DownloadResult struct {
	Service   string               `json:"service"`
	ElementID string               `json:"elementId"`
	OutputDir string               `json:"outputDir"`
	Files     []DownloadFileResult `json:"files"`
	DryRun    bool                 `json:"dryRun"`
}

// FormatSpecification represents a supported file format definition.
// Field alignment optimized.
type FormatSpecification struct {
	ID          string `json:"id"`
	Extension   string `json:"extension"`
	Description string `json:"description,omitempty"`
	BaseURL     string `json:"baseUrl,omitempty"`
	BasePath    string `json:"basePath,omitempty"`
}

type dryRunParams struct {
	cat     *catalog.Catalog
	elem    *catalog.Element
	service string
	outDir  string
	formats []string
}

type downloadExecutionParams struct {
	cat       *catalog.Catalog
	elem      *catalog.Element
	cfgFile   string
	service   string
	outDir    string
	formats   []string
	checkHash bool
}

func (s *Server) registerTools() {
	s.mcpServer.AddTool(s.toolListServices(), s.handleListServices)
	s.mcpServer.AddTool(s.toolRegenerateCatalog(), s.handleRegenerateCatalog)
	s.mcpServer.AddTool(s.toolListElements(), s.handleListElements)
	s.mcpServer.AddTool(s.toolGetElement(), s.handleGetElement)
	s.mcpServer.AddTool(s.toolListFormats(), s.handleListFormats)
	s.mcpServer.AddTool(s.toolDownloadElement(), s.handleDownloadElement)
}

func (s *Server) toolListServices() mcpSDK.Tool {
	desc := "List all available OpenStreetMap data provider services (e.g. geofabrik, openstreetmap.fr, bbbike, geo2day, movisda, etc.) " +
		"with metadata and local config status."

	return mcpSDK.NewTool("list_services",
		mcpSDK.WithDescription(desc),
		mcpSDK.WithBoolean("check_local",
			mcpSDK.Description("Check whether local YAML configuration files exist on disk (default: true)."),
		),
	)
}

func (s *Server) toolRegenerateCatalog() mcpSDK.Tool {
	return mcpSDK.NewTool("regenerate_catalog",
		mcpSDK.WithDescription("Fetch the latest OSM catalog from the specified service provider and save or update the local YAML configuration file."),
		mcpSDK.WithString("service",
			mcpSDK.Required(),
			mcpSDK.Description("Name of the service to regenerate (e.g. 'geofabrik', 'bbbike', 'openstreetmap.fr', or 'all')."),
		),
		mcpSDK.WithString("config_file",
			mcpSDK.Description("Optional custom destination path for the YAML configuration file (default is <service>.yml)."),
		),
	)
}

func (s *Server) toolListElements() mcpSDK.Tool {
	desc := "List and search elements (regions, countries, cities) in a service catalog to discover available file formats " +
		"(osm.pbf, shp.zip, poly, etc.) and metadata."

	return mcpSDK.NewTool("list_elements",
		mcpSDK.WithDescription(desc),
		mcpSDK.WithString("service",
			mcpSDK.Description("Service name to query (default: 'geofabrik')."),
		),
		mcpSDK.WithString("config_file",
			mcpSDK.Description("Optional custom configuration file path."),
		),
		mcpSDK.WithString("search",
			mcpSDK.Description("Case-insensitive search filter on element ID or Name."),
		),
		mcpSDK.WithString("parent",
			mcpSDK.Description("Filter by parent region identifier (e.g. 'europe', 'france', 'asia')."),
		),
		mcpSDK.WithInteger("limit",
			mcpSDK.Description("Maximum number of elements to return (default: 50, 0 for all, max: 500)."),
		),
		mcpSDK.WithInteger("offset",
			mcpSDK.Description("Pagination offset (default: 0)."),
		),
	)
}

func (s *Server) toolGetElement() mcpSDK.Tool {
	desc := "Get detailed information about a specific element in a service, including resolved download URLs for each format " +
		"and checksum verification availability."

	return mcpSDK.NewTool("get_element",
		mcpSDK.WithDescription(desc),
		mcpSDK.WithString("element_id",
			mcpSDK.Required(),
			mcpSDK.Description("Element identifier (e.g. 'france', 'berlin', 'rhone-alpes')."),
		),
		mcpSDK.WithString("service",
			mcpSDK.Description("Service name (default: 'geofabrik')."),
		),
		mcpSDK.WithString("config_file",
			mcpSDK.Description("Optional custom configuration file path."),
		),
	)
}

func (s *Server) toolListFormats() mcpSDK.Tool {
	return mcpSDK.NewTool("list_formats",
		mcpSDK.WithDescription("List file format definitions supported across services or defined in a specific service catalog."),
		mcpSDK.WithString("service",
			mcpSDK.Description("Optional service name to get format definitions specific to that provider."),
		),
		mcpSDK.WithString("config_file",
			mcpSDK.Description("Optional custom configuration file path."),
		),
	)
}

func (s *Server) toolDownloadElement() mcpSDK.Tool {
	return mcpSDK.NewTool("download_element",
		mcpSDK.WithDescription("Download an OSM extract for a given element and format(s) from a provider service."),
		mcpSDK.WithString("element_id",
			mcpSDK.Required(),
			mcpSDK.Description("Element identifier to download (e.g. 'monaco', 'france', 'berlin')."),
		),
		mcpSDK.WithString("service",
			mcpSDK.Description("Service to download from (default: 'geofabrik')."),
		),
		mcpSDK.WithArray("formats",
			mcpSDK.Description("Formats to download (e.g. ['osm.pbf'], ['shp.zip', 'poly']). Defaults to ['osm.pbf'] if omitted."),
			mcpSDK.WithStringItems(),
		),
		mcpSDK.WithString("output_dir",
			mcpSDK.Description("Directory where files will be saved (default: current working directory)."),
		),
		mcpSDK.WithBoolean("check_checksum",
			mcpSDK.Description("Verify MD5 checksum if available (default: true)."),
		),
		mcpSDK.WithBoolean("dry_run",
			mcpSDK.Description("Simulate the download and resolve URLs without downloading actual data (default: false)."),
		),
		mcpSDK.WithString("config_file",
			mcpSDK.Description("Optional custom configuration file path."),
		),
	)
}

//nolint:gocritic // parameter signature defined by mcpSDK.ToolHandlerFunc interface
func (s *Server) handleListServices(_ context.Context, request mcpSDK.CallToolRequest) (*mcpSDK.CallToolResult, error) {
	checkLocal := request.GetBool("check_local", true)
	names := provider.List()
	services := make([]ServiceInfo, 0, len(names))

	for _, name := range names {
		prov, err := provider.Get(name)
		if err != nil {
			continue
		}

		defaultCfg := prov.DefaultConfigFile()
		hasLocal := false

		if checkLocal && defaultCfg != "" {
			if _, statErr := os.Stat(defaultCfg); statErr == nil {
				hasLocal = true
			}
		}

		services = append(services, ServiceInfo{
			Name:              prov.Name(),
			Description:       prov.Description(),
			DefaultConfigFile: defaultCfg,
			HasLocalConfig:    hasLocal,
		})
	}

	return jsonToolResult(services)
}

//nolint:gocritic // parameter signature defined by mcpSDK.ToolHandlerFunc interface
func (s *Server) handleRegenerateCatalog(ctx context.Context, request mcpSDK.CallToolRequest) (*mcpSDK.CallToolResult, error) {
	serviceName := strings.TrimSpace(request.GetString("service", ""))
	if serviceName == "" {
		return mcpSDK.NewToolResultError("service parameter is required"), nil
	}

	customCfg := request.GetString("config_file", "")

	if serviceName == "all" {
		return s.regenerateAllCatalogs(ctx)
	}

	prov, _ := provider.Get(serviceName)
	if prov == nil {
		return mcpSDK.NewToolResultError(fmt.Sprintf("%v: %s", errInvalidService, serviceName)), nil
	}

	targetFile := customCfg
	if targetFile == "" {
		targetFile = prov.DefaultConfigFile()
	}

	if genErr := generator.PerformGenerateContext(ctx, serviceName, false, targetFile); genErr != nil {
		return mcpSDK.NewToolResultError(fmt.Sprintf("failed to regenerate catalog for %s: %v", serviceName, genErr)), nil
	}

	cat, loadErr := catalog.LoadFile(targetFile)
	elementCount := 0

	if loadErr == nil && cat != nil {
		elementCount = len(cat.Elements)
	}

	response := map[string]any{
		fieldStatus:     statusSuccess,
		"service":       serviceName,
		"configFile":    targetFile,
		"totalElements": elementCount,
		"message":       fmt.Sprintf("Successfully regenerated catalog for %s (%d elements saved to %s)", serviceName, elementCount, targetFile),
	}

	return jsonToolResult(response)
}

func (s *Server) regenerateAllCatalogs(ctx context.Context) (*mcpSDK.CallToolResult, error) {
	names := provider.List()
	results := make(map[string]any, len(names))

	for _, name := range names {
		prov, err := provider.Get(name)
		if err != nil {
			results[name] = map[string]any{fieldStatus: statusError, "error": err.Error()}

			continue
		}

		targetFile := prov.DefaultConfigFile()

		if genErr := generator.PerformGenerateContext(ctx, name, false, targetFile); genErr != nil {
			results[name] = map[string]any{fieldStatus: statusError, "error": genErr.Error()}

			continue
		}

		cat, _ := catalog.LoadFile(targetFile)
		count := 0

		if cat != nil {
			count = len(cat.Elements)
		}

		results[name] = map[string]any{
			fieldStatus:     statusSuccess,
			"configFile":    targetFile,
			"totalElements": count,
		}
	}

	return jsonToolResult(map[string]any{
		fieldStatus: statusSuccess,
		"services":  results,
	})
}

//nolint:gocritic // parameter signature defined by mcpSDK.ToolHandlerFunc interface
func (s *Server) handleListElements(ctx context.Context, request mcpSDK.CallToolRequest) (*mcpSDK.CallToolResult, error) {
	serviceName := request.GetString("service", config.DefaultService)
	customCfg := request.GetString("config_file", "")
	searchQuery := strings.ToLower(strings.TrimSpace(request.GetString("search", "")))
	parentFilter := strings.TrimSpace(request.GetString("parent", ""))
	limit := request.GetInt("limit", defaultListLimit)
	offset := request.GetInt("offset", 0)

	if limit < 0 {
		limit = defaultListLimit
	} else if limit > maxListLimit {
		limit = maxListLimit
	}

	cat, err := s.loadOrFetchCatalog(ctx, serviceName, customCfg)
	if err != nil {
		return mcpSDK.NewToolResultError(fmt.Sprintf("failed to load catalog for %s: %v", serviceName, err)), nil
	}

	filtered := filterElements(cat, searchQuery, parentFilter)
	totalMatched := len(filtered)
	pagedElements := applyPagination(filtered, offset, limit)

	result := ListElementsResult{
		Service:       serviceName,
		TotalElements: len(cat.Elements),
		TotalMatched:  totalMatched,
		Limit:         limit,
		Offset:        offset,
		Elements:      pagedElements,
	}

	return jsonToolResult(result)
}

func filterElements(cat *catalog.Catalog, searchQuery, parentFilter string) []ElementSummary {
	keys := cat.SortedKeys()
	filtered := make([]ElementSummary, 0, len(keys))

	for _, key := range keys {
		elem, exists := cat.Get(key)
		if !exists {
			continue
		}

		if parentFilter != "" && !strings.EqualFold(elem.Parent, parentFilter) {
			continue
		}

		if searchQuery != "" {
			idMatch := strings.Contains(strings.ToLower(elem.ID), searchQuery)
			nameMatch := strings.Contains(strings.ToLower(elem.Name), searchQuery)

			if !idMatch && !nameMatch {
				continue
			}
		}

		formatsCopy := make([]string, len(elem.Formats))
		copy(formatsCopy, elem.Formats)

		filtered = append(filtered, ElementSummary{
			ID:      elem.ID,
			Name:    elem.Name,
			Parent:  elem.Parent,
			Formats: formatsCopy,
			Meta:    elem.Meta,
		})
	}

	return filtered
}

func applyPagination(items []ElementSummary, offset, limit int) []ElementSummary {
	total := len(items)
	if offset >= total {
		return []ElementSummary{}
	}

	end := total
	if limit > 0 && offset+limit < total {
		end = offset + limit
	}

	return items[offset:end]
}

//nolint:gocritic // parameter signature defined by mcpSDK.ToolHandlerFunc interface
func (s *Server) handleGetElement(ctx context.Context, request mcpSDK.CallToolRequest) (*mcpSDK.CallToolResult, error) {
	elementID := strings.TrimSpace(request.GetString("element_id", ""))
	if elementID == "" {
		return mcpSDK.NewToolResultError(errEmptyElementID.Error()), nil
	}

	serviceName := request.GetString("service", config.DefaultService)
	customCfg := request.GetString("config_file", "")

	cat, err := s.loadOrFetchCatalog(ctx, serviceName, customCfg)
	if err != nil {
		return mcpSDK.NewToolResultError(fmt.Sprintf("failed to load catalog for %s: %v", serviceName, err)), nil
	}

	elem, err := cat.Find(elementID)
	if err != nil {
		return mcpSDK.NewToolResultError(fmt.Sprintf("element '%s' not found in service '%s': %v", elementID, serviceName, err)), nil
	}

	formatDetails := make([]FormatURLDetail, 0, len(elem.Formats))

	for _, formatID := range elem.Formats {
		detail := FormatURLDetail{
			FormatID: formatID,
		}

		resolvedURL, urlErr := cat.ResolveURL(elem, formatID)
		if urlErr != nil {
			detail.Error = urlErr.Error()
		} else {
			detail.DownloadURL = resolvedURL
		}

		hashable, hashExt, _ := cat.IsHashable(formatID)
		detail.ChecksumAvailable = hashable

		if hashable {
			detail.ChecksumFormatID = hashExt

			if hashURL, hashErr := resolveFormatOrHashURL(cat, elem, hashExt); hashErr == nil {
				detail.ChecksumURL = hashURL
			}
		}

		formatDetails = append(formatDetails, detail)
	}

	result := ElementDetail{
		ID:      elem.ID,
		Name:    elem.Name,
		Parent:  elem.Parent,
		File:    elem.File,
		Service: serviceName,
		Formats: formatDetails,
		Meta:    elem.Meta,
	}

	return jsonToolResult(result)
}

//nolint:gocritic // parameter signature defined by mcpSDK.ToolHandlerFunc interface
func (s *Server) handleListFormats(ctx context.Context, request mcpSDK.CallToolRequest) (*mcpSDK.CallToolResult, error) {
	serviceName := request.GetString("service", "")
	customCfg := request.GetString("config_file", "")

	if serviceName != "" {
		cat, err := s.loadOrFetchCatalog(ctx, serviceName, customCfg)
		if err != nil {
			return mcpSDK.NewToolResultError(fmt.Sprintf("failed to load catalog for %s: %v", serviceName, err)), nil
		}

		formatSpecs := make([]FormatSpecification, 0, len(cat.Formats))

		for formatID, formatDef := range cat.Formats {
			formatSpecs = append(formatSpecs, FormatSpecification{
				ID:          formatID,
				Extension:   formatDef.Loc,
				Description: formatDef.Type,
				BaseURL:     formatDef.BaseURL,
				BasePath:    formatDef.BasePath,
			})
		}

		slices.SortFunc(formatSpecs, func(a, b FormatSpecification) int {
			return strings.Compare(a.ID, b.ID)
		})

		return jsonToolResult(formatSpecs)
	}

	return jsonToolResult(collectGlobalFormats())
}

func collectGlobalFormats() []FormatSpecification {
	allDefs := make(map[string]catalog.Format)

	for _, formatMap := range []catalog.FormatDefinitions{
		geofabrik.DefaultFormats(),
		bbbike.DefaultFormats(),
		openstreetmapfr.DefaultFormats(),
		geo2day.DefaultFormats(),
		movisda.DefaultFormats(),
		osmch.DefaultFormats(),
		osmkewllu.DefaultFormats(),
		osmfitvutbr.DefaultFormats(),
		osmit.DefaultFormats(),
		osmtw.DefaultFormats(),
	} {
		for formatID, formatDef := range formatMap {
			if _, exists := allDefs[formatID]; !exists {
				allDefs[formatID] = formatDef
			}
		}
	}

	specs := make([]FormatSpecification, 0, len(allDefs))

	for formatID, formatDef := range allDefs {
		specs = append(specs, FormatSpecification{
			ID:          formatID,
			Extension:   formatDef.Loc,
			Description: formatDef.Type,
			BaseURL:     formatDef.BaseURL,
			BasePath:    formatDef.BasePath,
		})
	}

	slices.SortFunc(specs, func(a, b FormatSpecification) int {
		return strings.Compare(a.ID, b.ID)
	})

	return specs
}

//nolint:gocritic // parameter signature defined by mcpSDK.ToolHandlerFunc interface
func (s *Server) handleDownloadElement(ctx context.Context, request mcpSDK.CallToolRequest) (*mcpSDK.CallToolResult, error) {
	elementID := strings.TrimSpace(request.GetString("element_id", ""))
	if elementID == "" {
		return mcpSDK.NewToolResultError(errEmptyElementID.Error()), nil
	}

	serviceName := request.GetString("service", config.DefaultService)
	customCfg := request.GetString("config_file", "")
	outDir := request.GetString("output_dir", "")
	checkChecksum := request.GetBool("check_checksum", true)
	dryRun := request.GetBool("dry_run", false)

	requestedFormats := request.GetStringSlice("formats", []string{formats.FormatOsmPbf})
	if len(requestedFormats) == 0 {
		requestedFormats = []string{formats.FormatOsmPbf}
	}

	resolvedDir, err := resolveOutputDirectory(outDir)
	if err != nil {
		return mcpSDK.NewToolResultError(fmt.Sprintf("invalid output directory: %v", err)), nil
	}

	cat, cfgFile, err := s.ensureCatalogFile(ctx, serviceName, customCfg)
	if err != nil {
		return mcpSDK.NewToolResultError(fmt.Sprintf("failed to prepare catalog for %s: %v", serviceName, err)), nil
	}

	elem, err := cat.Find(elementID)
	if err != nil {
		return mcpSDK.NewToolResultError(fmt.Sprintf("element '%s' not found in %s: %v", elementID, serviceName, err)), nil
	}

	if dryRun {
		return s.performDryRunDownload(&dryRunParams{
			cat:     cat,
			elem:    elem,
			service: serviceName,
			outDir:  resolvedDir,
			formats: requestedFormats,
		})
	}

	return s.performActualDownload(ctx, &downloadExecutionParams{
		cat:       cat,
		elem:      elem,
		cfgFile:   cfgFile,
		service:   serviceName,
		outDir:    resolvedDir,
		formats:   requestedFormats,
		checkHash: checkChecksum,
	})
}

func resolveOutputDirectory(dir string) (string, error) {
	if dir == "" {
		workingDir, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get working directory: %w", err)
		}

		return workingDir + string(os.PathSeparator), nil
	}

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("invalid path %s: %w", dir, err)
	}

	if !strings.HasSuffix(absDir, string(os.PathSeparator)) {
		absDir += string(os.PathSeparator)
	}

	return absDir, nil
}

func (s *Server) performDryRunDownload(params *dryRunParams) (*mcpSDK.CallToolResult, error) {
	results := make([]DownloadFileResult, 0, len(params.formats))

	for _, formatID := range params.formats {
		targetPath := params.outDir + params.elem.ID + "." + formatID
		downloadURL, urlErr := params.cat.ResolveURL(params.elem, formatID)

		fileRes := DownloadFileResult{
			Format:     formatID,
			TargetFile: targetPath,
			Status:     "dry_run",
		}

		if urlErr != nil {
			fileRes.Error = urlErr.Error()
		} else {
			fileRes.DownloadURL = downloadURL
		}

		if hashable, hashExt, _ := params.cat.IsHashable(formatID); hashable {
			if hashURL, hashErr := resolveFormatOrHashURL(params.cat, params.elem, hashExt); hashErr == nil {
				fileRes.ChecksumURL = hashURL
			}
		}

		if statInfo, statErr := os.Stat(targetPath); statErr == nil {
			fileRes.FileSizeBytes = statInfo.Size()
			fileRes.Status = "dry_run_exists"
		}

		results = append(results, fileRes)
	}

	return jsonToolResult(DownloadResult{
		Service:   params.service,
		ElementID: params.elem.ID,
		OutputDir: params.outDir,
		Files:     results,
		DryRun:    true,
	})
}

func (s *Server) performActualDownload(ctx context.Context, params *downloadExecutionParams) (*mcpSDK.CallToolResult, error) {
	cfg, err := config.LoadConfig(params.cfgFile)
	if err != nil {
		return mcpSDK.NewToolResultError(fmt.Sprintf("failed to load configuration %s: %v", params.cfgFile, err)), nil
	}

	opts := &config.Options{
		ConfigFile:      params.cfgFile,
		OutputDirectory: params.outDir,
		Check:           params.checkHash,
		Progress:        false,
		Quiet:           true,
	}

	downloaderInstance := downloader.NewDownloader(cfg, opts)
	results := make([]DownloadFileResult, 0, len(params.formats))

	for _, formatID := range params.formats {
		formatDef, exists := cfg.Formats[formatID]
		if !exists {
			results = append(results, DownloadFileResult{
				Format: formatID,
				Error:  "unknown format: " + formatID,
				Status: "failed",
			})

			continue
		}

		targetFile := params.outDir + params.elem.ID + "." + formatDef.ID
		fileRes := DownloadFileResult{
			Format:     formatID,
			TargetFile: targetFile,
		}

		downloadURL, urlErr := params.cat.ResolveURL(params.elem, formatID)
		if urlErr == nil {
			fileRes.DownloadURL = downloadURL
		}

		if hashable, hashExt, _ := params.cat.IsHashable(formatID); hashable {
			if hashURL, hashErr := resolveFormatOrHashURL(params.cat, params.elem, hashExt); hashErr == nil {
				fileRes.ChecksumURL = hashURL
			}
		}

		slog.Debug("MCP downloading element", "element", params.elem.ID, "format", formatID, "target", targetFile)

		downloadErr := downloaderInstance.DownloadFile(ctx, params.elem.ID, formatID, targetFile)
		if downloadErr != nil {
			fileRes.Error = downloadErr.Error()
			fileRes.Status = "failed"
			results = append(results, fileRes)

			continue
		}

		if params.checkHash {
			checksumOK := downloaderInstance.Checksum(ctx, params.elem.ID, formatID)
			fileRes.ChecksumVerified = checksumOK
		}

		if statInfo, statErr := os.Stat(targetFile); statErr == nil {
			fileRes.FileSizeBytes = statInfo.Size()
		}

		fileRes.Status = "downloaded"
		results = append(results, fileRes)
	}

	return jsonToolResult(DownloadResult{
		Service:   params.service,
		ElementID: params.elem.ID,
		OutputDir: params.outDir,
		Files:     results,
		DryRun:    false,
	})
}

func (s *Server) ensureCatalogFile(ctx context.Context, serviceName, customCfg string) (*catalog.Catalog, string, error) {
	targetFile := customCfg
	if targetFile == "" {
		prov, err := provider.Get(serviceName)
		if err != nil {
			return nil, "", fmt.Errorf("%w: %s", errInvalidService, serviceName)
		}

		targetFile = prov.DefaultConfigFile()
	}

	if _, statErr := os.Stat(targetFile); os.IsNotExist(statErr) {
		slog.Info("Catalog file not found, generating on the fly", "service", serviceName, "file", targetFile)

		if genErr := generator.PerformGenerateContext(ctx, serviceName, false, targetFile); genErr != nil {
			return nil, "", fmt.Errorf("failed to generate catalog for %s: %w", serviceName, genErr)
		}
	}

	cat, err := catalog.LoadFile(targetFile)
	if err != nil {
		return nil, "", fmt.Errorf("failed to load catalog %s: %w", targetFile, err)
	}

	return cat, targetFile, nil
}

func (s *Server) loadOrFetchCatalog(ctx context.Context, serviceName, customCfg string) (*catalog.Catalog, error) {
	if customCfg != "" {
		cat, err := catalog.LoadFile(customCfg)
		if err != nil {
			return nil, fmt.Errorf("failed to load custom catalog %s: %w", customCfg, err)
		}

		return cat, nil
	}

	prov, err := provider.Get(serviceName)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errInvalidService, serviceName)
	}

	defaultFile := prov.DefaultConfigFile()
	if _, statErr := os.Stat(defaultFile); statErr == nil {
		cat, loadErr := catalog.LoadFile(defaultFile)
		if loadErr == nil {
			return cat, nil
		}
	}

	slog.Info("Fetching catalog live from provider", "service", serviceName)

	cat, fetchErr := prov.FetchCatalog(ctx)
	if fetchErr != nil {
		return nil, fmt.Errorf("failed to fetch catalog from %s: %w", serviceName, fetchErr)
	}

	return cat, nil
}

func jsonToolResult(v any) (*mcpSDK.CallToolResult, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return mcpSDK.NewToolResultError(fmt.Sprintf("failed to serialize response: %v", err)), nil
	}

	return mcpSDK.NewToolResultText(string(data)), nil
}

func resolveFormatOrHashURL(cat *catalog.Catalog, elem *catalog.Element, formatID string) (string, error) {
	if elem == nil || formatID == "" {
		return "", errNilOrEmpty
	}

	formatDef, exists := cat.Formats[formatID]
	if !exists {
		return "", fmt.Errorf("%w: %s", errFormatMissing, formatID)
	}

	baseURL := formatDef.BaseURL
	if baseURL == "" {
		baseURL = cat.BaseURL
	}

	preURL, err := cat.ResolvePreURL(elem, baseURL, formatDef.BasePath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve pre URL: %w", err)
	}

	return preURL + formatDef.Loc, nil
}
