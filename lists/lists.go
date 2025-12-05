package lists

import (
	"fmt"
	"os"
	"sort"

	"github.com/apex/log"
	"github.com/julien-noblet/download-geofabrik/config"
	"github.com/julien-noblet/download-geofabrik/formats"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
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
		err := table.Append(
			item,
			configuration.Elements[configuration.Elements[item].Parent].Name,
			configuration.Elements[item].Name,
			formats.GetMiniFormats(configuration.Elements[item].Formats),
		)
		if err != nil {
			return err
		}
	}

	if err := table.Render(); err != nil {
		return err
	}
	fmt.Printf("Total elements: %#v\n", len(configuration.Elements)) //nolint:forbidigo // I want to print the number of elements

	return nil
}

// CreateTable creates a table with the specified format.
func CreateTable(format string) *tablewriter.Table {
	// Options
	opts := []tablewriter.Option{
		tablewriter.WithHeader([]string{"ShortName", "Is in", "Long Name", "formats"}),
		tablewriter.WithAlignment(tw.MakeAlign(4, tw.AlignLeft)),
	}

	if format == markdownFormat {
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
