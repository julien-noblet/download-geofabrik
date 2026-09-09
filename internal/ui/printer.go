package ui

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

const (
	tabPadding = 4
)

// PrintTable formats and outputs the catalog as an ASCII or Markdown table.
func PrintTable(cat *catalog.Catalog, isMarkdown bool, writer io.Writer) error {
	if cat == nil {
		return nil
	}

	if isMarkdown {
		return printMarkdownTable(cat, writer)
	}

	return printStandardTable(cat, writer)
}

func printMarkdownTable(cat *catalog.Catalog, writer io.Writer) error {
	tabWriter := tabwriter.NewWriter(writer, 0, 0, 1, ' ', 0)
	fmt.Fprintln(tabWriter, "| ShortName | Is in | Long Name | formats |")
	fmt.Fprintln(tabWriter, "| --- | --- | --- | --- |")

	for _, elementID := range cat.SortedKeys() {
		elem, _ := cat.Get(elementID)
		parentName := ""

		if elem.Parent != "" {
			if parentElem, exists := cat.Get(elem.Parent); exists {
				parentName = parentElem.Name
			}
		}

		fmt.Fprintf(tabWriter, "| %s | %s | %s | %s |\n",
			elementID,
			parentName,
			elem.Name,
			catalog.GetMiniFormats(elem.Formats),
		)
	}

	if err := tabWriter.Flush(); err != nil {
		return fmt.Errorf("unable to flush markdown table: %w", err)
	}

	return nil
}

func printStandardTable(cat *catalog.Catalog, writer io.Writer) error {
	tabWriter := tabwriter.NewWriter(writer, 0, 0, tabPadding, ' ', 0)
	fmt.Fprintln(tabWriter, "SHORTNAME\tIS IN\tLONG NAME\tFORMATS")

	for _, elementID := range cat.SortedKeys() {
		elem, _ := cat.Get(elementID)
		parentName := ""

		if elem.Parent != "" {
			if parentElem, exists := cat.Get(elem.Parent); exists {
				parentName = parentElem.Name
			}
		}

		fmt.Fprintf(tabWriter, "%s\t%s\t%s\t%s\n",
			elementID,
			parentName,
			elem.Name,
			catalog.GetMiniFormats(elem.Formats),
		)
	}

	if err := tabWriter.Flush(); err != nil {
		return fmt.Errorf("unable to flush table: %w", err)
	}

	fmt.Fprintf(writer, "Total elements: %d\n", len(cat.Elements))

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
		return fmt.Errorf("unable to encode json: %w", err)
	}

	return nil
}
