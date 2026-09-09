package cli

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	downloader "github.com/julien-noblet/download-geofabrik/internal/downloader"
	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

var (
	// Flags for download command.
	outputDir        string
	check            bool
	noDownload       bool
	downloadProgress bool
	// Format flags.
	formatFlags = make(map[string]*bool)
)

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

	// Add format flags
	addFormatFlag(catalog.KeyOsmPbf, "P", "Download osm.pbf (default)")
	addFormatFlag(catalog.KeyOshPbf, "H", "Download osh.pbf")
	addFormatFlag(catalog.KeyOsmGz, "G", "Download osm.gz")
	addFormatFlag(catalog.KeyOsmBz2, "B", "Download osm.bz2")
	addFormatFlag(catalog.KeyShpZip, "S", "Download shp.zip")
	addFormatFlag(catalog.KeyState, "", "Download state.txt")
	addFormatFlag(catalog.KeyPoly, "p", "Download poly")
	addFormatFlag(catalog.KeyKml, "k", "Download kml")
	addFormatFlag(catalog.KeyGeoJSON, "g", "Download GeoJSON")
	addFormatFlag(catalog.KeyGarminOSM, "O", "Download Garmin OSM")

	// Others...
	addFormatFlag(catalog.KeyMapsforge, "m", "Download Mapsforge")
	addFormatFlag(catalog.KeyMBTiles, "M", "Download MBTiles")
	addFormatFlag(catalog.KeyCSV, "C", "Download CSV")
	addFormatFlag(catalog.KeyGarminOnroad, "r", "Download Garmin Onroad")
	addFormatFlag(catalog.KeyGarminOntrail, "t", "Download Garmin Ontrail")
	addFormatFlag(catalog.KeyGarminOpenTopo, "o", "Download Garmin OpenTopo")
	addFormatFlag(catalog.KeyOBF, "", "Download OBF")
	addFormatFlag(catalog.KeyGPKG, "K", "Download GeoPackage")
	addFormatFlag(catalog.KeyO5m, "5", "Download o5m")
	addFormatFlag(catalog.KeyO5mZst, "Z", "Download o5m.zst")
	addFormatFlag(catalog.KeyPbf, "", "Download pbf")
}

func addFormatFlag(key, shorthand, usage string) {
	val := false
	formatFlags[key] = &val
	downloadCmd.Flags().BoolVarP(&val, key, shorthand, false, usage)
}

func buildDownloadOptions() (*downloader.Options, error) {
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
		FormatFlags:     make(map[string]bool, len(formatFlags)),
	}

	for k, v := range formatFlags {
		opts.FormatFlags[k] = *v
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

	if dir[len(dir)-1] != os.PathSeparator {
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
	if elem == nil {
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

func resolveFormat(cat *catalog.Catalog, elem *catalog.Element, format string) (string, bool) {
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
	elementID := args[0]

	opts, err := buildDownloadOptions()
	if err != nil {
		return err
	}

	cat, err := catalog.LoadFile(opts.ConfigFile)
	if err != nil {
		slog.Error("Failed to load config", "file", opts.ConfigFile, "error", err)

		return fmt.Errorf("failed to load config: %w", err)
	}

	myElem, err := cat.Find(elementID)
	if err != nil {
		slog.Error("Element not found", "element", elementID, "error", err)

		return fmt.Errorf("%w: %s", catalog.ErrElementNotFound, elementID)
	}

	var activeFormats []string
	if hasExplicitFormat(opts.FormatFlags) {
		activeFormats = catalog.GetFormats(opts.FormatFlags)
	} else {
		activeFormats = []string{selectDefaultFormat(cat, myElem)}
	}

	downloaderInstance := downloader.NewDownloader(cat, opts)
	ctx := cmd.Context()

	for _, rawFormat := range activeFormats {
		format, ok := resolveFormat(cat, myElem, rawFormat)
		if !ok {
			slog.Error("Format not available for element", "format", rawFormat, "element", elementID)

			return fmt.Errorf("%w: %s for %s", catalog.ErrFormatNotFound, rawFormat, elementID)
		}

		formatDef := cat.Formats[format]
		targetFile := opts.OutputDirectory + elementID + "." + formatDef.ID

		slog.Info("Processing", "element", elementID, "format", format)

		if err := processDownload(ctx, downloaderInstance, opts.Check, elementID, format, targetFile); err != nil {
			return err
		}
	}

	return nil
}

func processDownload(ctx context.Context, downloaderInstance *downloader.Downloader, check bool, elementID, format, targetFile string) error {
	if !check {
		if err := downloaderInstance.DownloadFile(ctx, elementID, format, targetFile); err != nil {
			return fmt.Errorf("download failed: %w", err)
		}

		return nil
	}

	shouldDownload := true

	if downloader.FileExist(targetFile) {
		if downloaderInstance.Checksum(ctx, elementID, format) {
			slog.Info("File already exists and checksum matches", "file", targetFile)

			shouldDownload = false
		} else {
			slog.Warn("Checksum mismatch or verification failed, re-downloading", "file", targetFile)
		}
	}

	if shouldDownload {
		if err := downloaderInstance.DownloadFile(ctx, elementID, format, targetFile); err != nil {
			return fmt.Errorf("download failed: %w", err)
		}
		// Verify again
		downloaderInstance.Checksum(ctx, elementID, format)
	}

	return nil
}
