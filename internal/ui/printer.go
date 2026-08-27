package ui

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

const (
	tableColumnCount = 4
)

// PrintTable formats and outputs the catalog as an ASCII or Markdown table.
func PrintTable(cat *catalog.Catalog, isMarkdown bool, writer io.Writer) error {
	if cat == nil {
		return nil
	}

	opts := []tablewriter.Option{
		tablewriter.WithHeader([]string{"ShortName", "Is in", "Long Name", "formats"}),
		tablewriter.WithAlignment(tw.MakeAlign(tableColumnCount, tw.AlignLeft)),
	}

	if isMarkdown {
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

	table := tablewriter.NewTable(writer, opts...)
	keys := cat.SortedKeys()

	for _, elementID := range keys {
		elem, _ := cat.Get(elementID)
		parentName := ""

		if elem.Parent != "" {
			if parentElem, exists := cat.Get(elem.Parent); exists {
				parentName = parentElem.Name
			}
		}

		err := table.Append(
			elementID,
			parentName,
			elem.Name,
			catalog.GetMiniFormats(elem.Formats),
		)
		if err != nil {
			return fmt.Errorf("unable to append row: %w", err)
		}
	}

	if err := table.Render(); err != nil {
		return fmt.Errorf("unable to render table: %w", err)
	}

	if !isMarkdown {
		fmt.Fprintf(writer, "Total elements: %d\n", len(cat.Elements))
	}

	return nil
}

// PrintJSON formats and outputs the entire catalog as formatted JSON.
func PrintJSON(cat *catalog.Catalog, writer io.Writer) error {
	if cat == nil {
		return nil
	}

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(cat); err != nil {
		return fmt.Errorf("unable to encode JSON: %w", err)
	}

	return nil
}
