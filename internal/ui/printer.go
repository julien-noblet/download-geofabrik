package ui

import (
	"bufio"
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
	bufWriter := bufio.NewWriter(writer)
	defer func() {
		_ = bufWriter.Flush()
	}()

	if _, err := bufWriter.WriteString("| ShortName | Is in | Long Name | formats |\n| --- | --- | --- | --- |\n"); err != nil {
		return fmt.Errorf("writing markdown table header: %w", err)
	}

	parentNameCache := make(map[string]string)

	for _, elementID := range cat.SortedKeys() {
		elem, _ := cat.Get(elementID)
		parentName := ""

		if elem.Parent != "" {
			cached, ok := parentNameCache[elem.Parent]
			if ok {
				parentName = cached
			} else if parentElem, exists := cat.Get(elem.Parent); exists {
				parentName = parentElem.Name
				parentNameCache[elem.Parent] = parentName
			}
		}

		fmt.Fprintf(bufWriter, "| %s | %s | %s | %s |\n",
			elementID,
			parentName,
			elem.Name,
			catalog.GetMiniFormats(elem.Formats),
		)
	}

	return nil
}

func printStandardTable(cat *catalog.Catalog, writer io.Writer) error {
	tabWriter := tabwriter.NewWriter(writer, 0, 0, tabPadding, ' ', 0)
	fmt.Fprintln(tabWriter, "SHORTNAME\tIS IN\tLONG NAME\tFORMATS")

	keys := cat.SortedKeys()
	parentNameCache := make(map[string]string)

	for _, elementID := range keys {
		elem, _ := cat.Get(elementID)
		parentName := ""

		if elem.Parent != "" {
			cached, ok := parentNameCache[elem.Parent]
			if ok {
				parentName = cached
			} else if parentElem, exists := cat.Get(elem.Parent); exists {
				parentName = parentElem.Name
				parentNameCache[elem.Parent] = parentName
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

	fmt.Fprintf(writer, "Total elements: %d\n", len(keys))

	return nil
}

// PrintJSON formats and outputs the entire catalog as formatted JSON.
func PrintJSON(cat *catalog.Catalog, writer io.Writer) error {
	if cat == nil {
		return nil
	}

	bufWriter := bufio.NewWriter(writer)
	defer func() {
		_ = bufWriter.Flush()
	}()

	encoder := json.NewEncoder(bufWriter)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(cat); err != nil {
		return fmt.Errorf("unable to encode json: %w", err)
	}

	return nil
}
