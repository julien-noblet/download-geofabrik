package cli

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/julien-noblet/download-geofabrik/internal/downloader"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var (
	// ErrMissingElementArg is returned when the download command is called without an element argument.
	ErrMissingElementArg = errors.New("missing required element argument")

	// Flags for download command.
	outputDir        string
	check            bool
	noDownload       bool
	downloadProgress bool
)

type formatFlag struct {
	key       string
	shorthand string
	usage     string
}

var formatFlagList = []formatFlag{
	{key: catalog.KeyOsmPbf, shorthand: "P", usage: "Download osm.pbf (default)"},
	{key: catalog.KeyOshPbf, shorthand: "H", usage: "Download osh.pbf"},
	{key: catalog.KeyOsmGz, shorthand: "G", usage: "Download osm.gz"},
	{key: catalog.KeyOsmBz2, shorthand: "B", usage: "Download osm.bz2"},
	{key: catalog.KeyShpZip, shorthand: "S", usage: "Download shp.zip"},
	{key: catalog.KeyState, shorthand: "", usage: "Download state.txt"},
	{key: catalog.KeyPoly, shorthand: "p", usage: "Download poly"},
	{key: catalog.KeyKml, shorthand: "k", usage: "Download kml"},
	{key: catalog.KeyGeoJSON, shorthand: "g", usage: "Download GeoJSON"},
	{key: catalog.KeyGarminOSM, shorthand: "O", usage: "Download Garmin OSM"},
	{key: catalog.KeyMapsforge, shorthand: "m", usage: "Download Mapsforge"},
	{key: catalog.KeyMBTiles, shorthand: "M", usage: "Download MBTiles"},
	{key: catalog.KeyCSV, shorthand: "C", usage: "Download CSV"},
	{key: catalog.KeyGarminOnroad, shorthand: "r", usage: "Download Garmin Onroad"},
	{key: catalog.KeyGarminOntrail, shorthand: "t", usage: "Download Garmin Ontrail"},
	{key: catalog.KeyGarminOpenTopo, shorthand: "o", usage: "Download Garmin OpenTopo"},
	{key: catalog.KeyOBF, shorthand: "", usage: "Download OBF"},
	{key: catalog.KeyGPKG, shorthand: "K", usage: "Download GeoPackage"},
	{key: catalog.KeyO5m, shorthand: "5", usage: "Download o5m"},
	{key: catalog.KeyO5mZst, shorthand: "Z", usage: "Download o5m.zst"},
	{key: catalog.KeyPbf, shorthand: "", usage: "Download pbf"},
}

var downloadCmd = &cobra.Command{
	Use:   "download [element]",
	Short: "Download element",
	Args:  cobra.ExactArgs(1),
	ValidArgsFunction: func(_ *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		cfgFile := viper.ConfigFileUsed()
		if cfgFile == "" {
			if service != "" {
				cfgFile = service + ".yml"
			} else {
				cfgFile = catalog.DefaultConfigFile
			}
		}

		cat, err := catalog.LoadFile(cfgFile)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		var matches []string

		for _, key := range cat.SortedKeys() {
			if strings.HasPrefix(key, toComplete) {
				matches = append(matches, key)
			}
		}

		return matches, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: runDownload,
}

// RegisterDownloadCmd registers the download command and its flags to rootCmd.
func RegisterDownloadCmd() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&outputDir, "output-dir", "d", "", "Set output directory")
	downloadCmd.Flags().BoolVar(&check, "check", true, "Control with checksum (default). Use --no-check to discard control")
	downloadCmd.Flags().BoolVarP(&noDownload, "nodownload", "n", false, "Do not download file (test only)")
	downloadCmd.Flags().BoolVar(&downloadProgress, "progress", true, "Show progress bar")

	for _, flagDef := range formatFlagList {
		downloadCmd.Flags().BoolP(flagDef.key, flagDef.shorthand, false, flagDef.usage)
	}
}

func buildDownloadOptions(cmd *cobra.Command) (*downloader.Options, error) {
	cfgFile := viper.ConfigFileUsed()
	if cfgFile == "" {
		if service != "" {
			cfgFile = service + ".yml"
		} else {
			cfgFile = catalog.DefaultConfigFile
		}
	}

	outDir, err := resolveOutputDir(outputDir)
	if err != nil {
		return nil, err
	}

	opts := &downloader.Options{
		ConfigFile:      cfgFile,
		OutputDirectory: outDir,
		Check:           check,
		Verbose:         viper.GetBool("verbose"),
		Quiet:           viper.GetBool("quiet"),
		NoDownload:      noDownload,
		Progress:        downloadProgress,
		FormatFlags:     make(map[string]bool, len(formatFlagList)),
	}

	flags := cmd.Flags()
	for _, flagDef := range formatFlagList {
		if enabled, err := flags.GetBool(flagDef.key); err == nil && enabled {
			opts.FormatFlags[flagDef.key] = true
		}
	}

	return opts, nil
}

func resolveOutputDir(dir string) (string, error) {
	if dir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get working directory: %w", err)
		}

		return wd + string(os.PathSeparator), nil
	}

	if !strings.HasSuffix(dir, string(os.PathSeparator)) {
		return dir + string(os.PathSeparator), nil
	}

	return dir, nil
}

var preferredDefaultFormats = []string{
	catalog.FormatOsmPbf,
	catalog.FormatPbf,
	catalog.FormatO5m,
	catalog.FormatOsmBz2,
	catalog.FormatOsmGz,
	catalog.FormatGeoJSON,
	catalog.FormatGPKG,
	catalog.FormatShpZip,
	catalog.FormatO5mZst,
}

func hasExplicitFormat(flags map[string]bool) bool {
	for _, enabled := range flags {
		if enabled {
			return true
		}
	}

	return false
}

func selectDefaultFormat(cat *catalog.Catalog, elem *catalog.Element) string {
	if elem == nil || cat == nil || cat.Formats == nil {
		return catalog.FormatOsmPbf
	}

	for _, pref := range preferredDefaultFormats {
		if elem.Formats.Contains(pref) {
			if _, ok := cat.Formats[pref]; ok {
				return pref
			}
		}
	}

	for _, format := range elem.Formats {
		if strings.HasSuffix(format, ".md5") {
			continue
		}

		if _, ok := cat.Formats[format]; ok {
			return format
		}
	}

	return catalog.FormatOsmPbf
}

func resolveFormat(cat *catalog.Catalog, elem *catalog.Element, format string) (string, bool) { //nolint:cyclop // nil guards +2 branches
	if cat == nil || elem == nil || cat.Formats == nil {
		return format, false
	}

	if elem.Formats.Contains(format) {
		if _, exists := cat.Formats[format]; exists {
			return format, true
		}
	}

	if format == catalog.FormatOsmPbf && elem.Formats.Contains(catalog.FormatPbf) {
		if _, exists := cat.Formats[catalog.FormatPbf]; exists {
			return catalog.FormatPbf, true
		}
	}

	if format == catalog.FormatPbf && elem.Formats.Contains(catalog.FormatOsmPbf) {
		if _, exists := cat.Formats[catalog.FormatOsmPbf]; exists {
			return catalog.FormatOsmPbf, true
		}
	}

	return format, false
}

func runDownload(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return ErrMissingElementArg
	}

	elementID := args[0]

	opts, err := buildDownloadOptions(cmd)
	if err != nil {
		return err
	}

	cat, err := catalog.LoadFile(opts.ConfigFile)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	myElem, err := cat.Find(elementID)
	if err != nil {
		return fmt.Errorf("%w: %s", catalog.ErrElementNotFound, elementID)
	}

	var activeFormats []string
	if hasExplicitFormat(opts.FormatFlags) {
		activeFormats = catalog.GetFormats(opts.FormatFlags)
	} else {
		activeFormats = []string{selectDefaultFormat(cat, myElem)}
	}

	client := downloader.New(cat, opts)
	ctx := cmd.Context()

	for _, rawFormat := range activeFormats {
		format, ok := resolveFormat(cat, myElem, rawFormat)
		if !ok {
			return fmt.Errorf("%w: %s for %s", catalog.ErrFormatNotFound, rawFormat, elementID)
		}

		formatDef := cat.Formats[format]
		targetFile := opts.OutputDirectory + elementID + "." + formatDef.ID

		slog.Info("Processing", "element", elementID, "format", format)

		if err := processDownload(ctx, client, opts.Check, elementID, format, targetFile); err != nil {
			return err
		}
	}

	return nil
}

func processDownload(ctx context.Context, client *downloader.Downloader, check bool, elementID, format, targetFile string) error {
	if !check {
		if err := client.DownloadFile(ctx, elementID, format, targetFile); err != nil {
			return fmt.Errorf("download failed: %w", err)
		}

		return nil
	}

	shouldDownload := true

	if downloader.FileExists(targetFile) {
		if client.Checksum(ctx, elementID, format) {
			slog.Info("File already exists and checksum matches", "file", targetFile)

			shouldDownload = false
		} else {
			slog.Warn("Checksum mismatch or verification failed, re-downloading", "file", targetFile)
		}
	}

	if shouldDownload {
		if err := client.DownloadFile(ctx, elementID, format, targetFile); err != nil {
			return fmt.Errorf("download failed: %w", err)
		}
		// Verify again
		client.Checksum(ctx, elementID, format)
	}

	return nil
}
