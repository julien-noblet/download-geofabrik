package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/julien-noblet/download-geofabrik/internal/provider"
	mcpSDK "github.com/mark3labs/mcp-go/mcp"
)

const (
	uriSchemeServices = "geofabrik://services"
	uriSchemeFormats  = "geofabrik://formats"
	mimeTypeJSON      = "application/json"
)

func (s *Server) registerResources() {
	s.mcpServer.AddResource(
		mcpSDK.NewResource(
			uriSchemeServices,
			"OSM Data Services",
			mcpSDK.WithResourceDescription("List of available OpenStreetMap data provider services with descriptions and default config files"),
			mcpSDK.WithMIMEType(mimeTypeJSON),
		),
		s.handleResourceServices,
	)

	s.mcpServer.AddResource(
		mcpSDK.NewResource(
			uriSchemeFormats,
			"OSM File Formats",
			mcpSDK.WithResourceDescription("List of supported OpenStreetMap file formats and extensions"),
			mcpSDK.WithMIMEType(mimeTypeJSON),
		),
		s.handleResourceFormats,
	)

	s.mcpServer.AddResourceTemplate(
		mcpSDK.NewResourceTemplate(
			"geofabrik://catalog/{service}",
			"Service Catalog",
			mcpSDK.WithTemplateDescription("Complete catalog containing elements and formats for a given OSM provider service"),
			mcpSDK.WithTemplateMIMEType(mimeTypeJSON),
		),
		s.handleResourceCatalogTemplate,
	)
}

func (s *Server) handleResourceServices(_ context.Context, request mcpSDK.ReadResourceRequest) ([]mcpSDK.ResourceContents, error) {
	names := provider.List()
	services := make([]ServiceInfo, 0, len(names))

	for _, name := range names {
		prov, err := provider.Get(name)
		if err != nil {
			continue
		}

		defaultCfg := prov.DefaultConfigFile()
		hasLocal := false

		if defaultCfg != "" {
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

	data, err := json.MarshalIndent(services, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal services: %w", err)
	}

	return []mcpSDK.ResourceContents{
		mcpSDK.TextResourceContents{
			URI:      request.Params.URI,
			MIMEType: mimeTypeJSON,
			Text:     string(data),
		},
	}, nil
}

func (s *Server) handleResourceFormats(_ context.Context, request mcpSDK.ReadResourceRequest) ([]mcpSDK.ResourceContents, error) {
	formatSpecs := collectGlobalFormats()

	data, err := json.MarshalIndent(formatSpecs, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal formats: %w", err)
	}

	return []mcpSDK.ResourceContents{
		mcpSDK.TextResourceContents{
			URI:      request.Params.URI,
			MIMEType: mimeTypeJSON,
			Text:     string(data),
		},
	}, nil
}

func (s *Server) handleResourceCatalogTemplate(ctx context.Context, request mcpSDK.ReadResourceRequest) ([]mcpSDK.ResourceContents, error) {
	parsedURL, err := url.Parse(request.Params.URI)
	if err != nil {
		return nil, fmt.Errorf("invalid resource URI %s: %w", request.Params.URI, err)
	}

	serviceName := strings.TrimPrefix(parsedURL.Path, "/")
	if serviceName == "" {
		serviceName = parsedURL.Host
	}

	cat, err := s.loadOrFetchCatalog(ctx, serviceName, "")
	if err != nil {
		return nil, fmt.Errorf("failed to load catalog for %s: %w", serviceName, err)
	}

	data, err := json.MarshalIndent(cat, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal catalog for %s: %w", serviceName, err)
	}

	return []mcpSDK.ResourceContents{
		mcpSDK.TextResourceContents{
			URI:      request.Params.URI,
			MIMEType: mimeTypeJSON,
			Text:     string(data),
		},
	}, nil
}
