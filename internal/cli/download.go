package cli

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	config "github.com/julien-noblet/download-geofabrik/internal/config"
	downloader "github.com/julien-noblet/download-geofabrik/internal/downloader"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	RunE:  runDownload,
}

func RegisterDownloadCmd() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&outputDir, "output-dir", "d", "", "Set output directory")
	downloadCmd.Flags().BoolVar(&check, "check", true, "Control with checksum (default). Use --no-check to discard control")
	downloadCmd.Flags().BoolVarP(&noDownload, "nodownload", "n", false, "Do not download file (test only)")
	downloadCmd.Flags().BoolVar(&downloadProgress, "progress", true, "Show progress bar")

	// Add format flags
	// These mimic the original kingpin flags
	addFormatFlag(formats.KeyOsmPbf, "P", "Download osm.pbf (default)")
	addFormatFlag(formats.KeyOshPbf, "H", "Download osh.pbf")
	addFormatFlag(formats.KeyOsmGz, "G", "Download osm.gz")
	addFormatFlag(formats.KeyOsmBz2, "B", "Download osm.bz2")
	addFormatFlag(formats.KeyShpZip, "S", "Download shp.zip")
	addFormatFlag(formats.KeyState, "s", "Download state.txt")
	addFormatFlag(formats.KeyPoly, "p", "Download poly")
	addFormatFlag(formats.KeyKml, "k", "Download kml")
	addFormatFlag(formats.KeyGeoJSON, "g", "Download GeoJSON")
	addFormatFlag(formats.KeyGarminOSM, "O", "Download Garmin OSM")

	// Others...
	addFormatFlag(formats.KeyMapsforge, "m", "Download Mapsforge")
	addFormatFlag(formats.KeyMBTiles, "M", "Download MBTiles")
	addFormatFlag(formats.KeyCSV, "C", "Download CSV")
	addFormatFlag(formats.KeyGarminOnroad, "r", "Download Garmin Onroad")
	addFormatFlag(formats.KeyGarminOntrail, "t", "Download Garmin Ontrail")
	addFormatFlag(formats.KeyGarminOpenTopo, "o", "Download Garmin OpenTopo")
}

func addFormatFlag(key, shorthand, usage string) {
	val := false
	formatFlags[key] = &val
	downloadCmd.Flags().BoolVarP(&val, key, shorthand, false, usage)
}

func runDownload(_ *cobra.Command, args []string) error {
	elementID := args[0]

	// Prepare Options
	// Note: rootCmd flags (config file, verbose) should be parsed already.
	// I need to access them. They are bound to Viper in root.go (I need to ensure that).

	cfgFile := viper.GetString("config")
	if cfgFile == "" {
		cfgFile = config.DefaultConfigFile
	}

	opts := &config.Options{
		ConfigFile:      cfgFile,
		OutputDirectory: outputDir,
		Check:           check,
		Verbose:         viper.GetBool("verbose"),
		Quiet:           viper.GetBool("quiet"),
		NoDownload:      noDownload,
		Progress:        downloadProgress,
		FormatFlags:     make(map[string]bool),
	}

	// Fill format flags
	for k, v := range formatFlags {
		opts.FormatFlags[k] = *v
	}

	// Ensure output dir has separator?
	if opts.OutputDirectory == "" {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}

		opts.OutputDirectory = wd + string(os.PathSeparator)
	} else if opts.OutputDirectory[len(opts.OutputDirectory)-1] != os.PathSeparator {
		opts.OutputDirectory += string(os.PathSeparator)
	}

	// Load Config
	cfg, err := config.LoadConfig(opts.ConfigFile)
	if err != nil {
		slog.Error("Failed to load config", "file", opts.ConfigFile, "error", err)

		return fmt.Errorf("failed to load config: %w", err)
	}

	// Determine active formats
	activeFormats := formats.GetFormats(opts.FormatFlags)

	downloaderInstance := downloader.NewDownloader(cfg, opts)

	ctx := context.Background()

	for _, format := range activeFormats {
		// Filename calculation?
		// original used `GetFilename`.
		// `filename := GetFilename(viper.GetString(viperOutputDirectoryKey), viper.GetString(viperElementKey))`

		// I should reconstruct the file path.
		// `Downloader.DownloadFile` takes outputPath (e.g. dir/elementName.ext) or just dir/elementName?
		// My implementation of `DownloadFile` takes `outputPath`.
		// And checks `format` ID from config.

		// Construct the base filename (without extension).
		// Original: `r.FindStringSubmatch(outputDir + element)[0]` -> basically basename of element?
		// No, if element is path?

		// Construct the base filename (without extension).
		outFile := opts.OutputDirectory + elementID
		// Get format details for extension
		formatDef := cfg.Formats[format]
		targetFile := outFile + "." + formatDef.ID

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
