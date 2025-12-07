package lists

import (
	"fmt"
	"os"
	"sort"

	"github.com/julien-noblet/download-geofabrik/internal/config"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

const (
	MarkdownFormat = "Markdown"
)

// ListAllRegions lists all regions in the specified format.
func ListAllRegions(configuration *config.Config, format string) error {
	table := CreateTable(format)
	keys := GetSortedKeys(configuration)

	for _, item := range keys {
		err := table.Append(
			item,
			configuration.Elements[configuration.Elements[item].Parent].Name,
			configuration.Elements[item].Name,
			formats.GetMiniFormats(configuration.Elements[item].Formats),
		)
		if err != nil {
			return fmt.Errorf("unable to append: %w", err)
		}
	}

	if err := table.Render(); err != nil {
		return fmt.Errorf("unable to render table: %w", err)
	}
	fmt.Printf("Total elements: %#v\n", len(configuration.Elements)) //nolint:forbidigo // I want to print the number of elements

	return nil
}

// CreateTable creates a table with the specified format.
func CreateTable(format string) *tablewriter.Table {
	// Options
	opts := []tablewriter.Option{
		tablewriter.WithHeader([]string{"ShortName", "Is in", "Long Name", "formats"}),
		tablewriter.WithAlignment(tw.MakeAlign(4, tw.AlignLeft)), //nolint:mnd // 4 columns to align left
	}

	if format == MarkdownFormat {
		opts = append(opts, tablewriter.WithRendition(tw.Rendition{
			Symbols: tw.NewSymbols(tw.StyleMarkdown),
			Borders: tw.Border{
				Left:   tw.On,
				Top:    tw.Off,
				Right:  tw.On,
				Bottom: tw.Off,
			},
		}))
	}

	return tablewriter.NewTable(os.Stdout, opts...)
}

// GetSortedKeys returns the sorted keys of the configuration elements.
func GetSortedKeys(configuration *config.Config) []string {
	keys := make(sort.StringSlice, len(configuration.Elements))
	i := 0

	for k := range configuration.Elements {
		keys[i] = k
		i++
	}

	keys.Sort()

	return keys
}
