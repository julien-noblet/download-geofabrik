package lists

import (
	"fmt"
	"os"
	"sort"

	"github.com/apex/log"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

const (
	markdownFormat = "Markdown"
)

// ListAllRegions lists all regions in the specified format.
func ListAllRegions(configuration *config.Config, format string) error {
	table := CreateTable(format)
	keys := GetSortedKeys(configuration)

	for _, item := range keys {
		table.Append([]string{
			item,
			configuration.Elements[configuration.Elements[item].Parent].Name,
			configuration.Elements[item].Name,
			formats.GetMiniFormats(configuration.Elements[item].Formats),
		})
	}

	table.Render()
	fmt.Printf("Total elements: %#v\n", len(configuration.Elements)) //nolint:forbidigo // I want to print the number of elements

	return nil
}

// CreateTable creates a table with the specified format.
func CreateTable(format string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ShortName", "Is in", "Long Name", "formats"})

	if format == markdownFormat {
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
	}

	return table
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

// ListCommand executes the list command.
func ListCommand() {
	format := ""

	configPtr, err := config.LoadConfig(viper.GetString(config.ViperConfig))
	if err != nil {
		log.WithError(err).Fatal(config.ErrLoadConfig)
	}

	if viper.GetBool(config.ViperListFormatMarkdown) {
		format = markdownFormat
	}

	if err := ListAllRegions(configPtr, format); err != nil {
		log.WithError(err).Fatal("Failed to list all regions")
	}
}
