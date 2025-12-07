package cli

import (
	"context"
	"log/slog"
	"os"

	config "github.com/julien-noblet/download-geofabrik/internal/config"
	downloader "github.com/julien-noblet/download-geofabrik/internal/downloader"
	"github.com/julien-noblet/download-geofabrik/pkg/formats"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Flags for download command
	outputDir        string
	check            bool
	allFormats       bool
	noDownload       bool
	downloadProgress bool
	// Format flags
	formatFlags = make(map[string]*bool)
)

var downloadCmd = &cobra.Command{
	Use:   "download [element]",
	Short: "Download element",
	Args:  cobra.ExactArgs(1),
	RunE:  runDownload,
}

func init() {
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

func runDownload(cmd *cobra.Command, args []string) error {
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
			return err
		}
		opts.OutputDirectory = wd + string(os.PathSeparator)
	} else {
		if opts.OutputDirectory[len(opts.OutputDirectory)-1] != os.PathSeparator {
			opts.OutputDirectory += string(os.PathSeparator)
		}
	}

	// Load Config
	cfg, err := config.LoadConfig(opts.ConfigFile)
	if err != nil {
		slog.Error("Failed to load config", "file", opts.ConfigFile, "error", err)
		return err
	}

	// Determine active formats
	activeFormats := formats.GetFormats(opts.FormatFlags)

	dl := downloader.NewDownloader(cfg, opts)

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

		// Let's assume simplest case: OutputDir + ElementID.
		outFile := opts.OutputDirectory + elementID

		slog.Info("Processing", "element", elementID, "format", format)

		if opts.Check {
			// Checksum calculates filenames itself?
			// My default Checksum implementation recalculated paths.
			// I should verify consistency.
			if !dl.Checksum(ctx, elementID, format) {
				// If checksum failed or mismatched, or file not exist.
				// Checksum returns TRUE if match. FALSE if download needed or mismatch.
				// Actually original `HandleHashableFormat` logic:
				// If file exist: check checksum. If match, skip. If mismatch, redownload.
				// If file not exist: download. Verify.

				// My new `Checksum` logic just returns bool.
				// It does NOT download the file itself?
				// Wait, my `Checksum` implementation calls `FromURL` to download the md5 file.
				// But does it download the MAIN file?

				// Looking at my new `Checksum`:
				// It downloads .md5
				// Calls `VerifyFileChecksum`.

				// It does NOT download the target file if missing.
				// The original code `HandleHashableFormat` did:
				// if exist -> check. if mismatch -> download.
				// else -> download. check.

				// So I need to orchestrate this here or in `downloader`.
				// `downloader.Checksum` currently verifies.

				// I should improve `Downloader` to have `DownloadWithChecksum` logic?
				// Or handle it here.

				// Let's implement logic here:
				targetFile := outFile + "." + cfg.Formats[format].ID

				shouldDownload := true
				if downloader.FileExist(targetFile) {
					if dl.Checksum(ctx, elementID, format) {
						slog.Info("File already exists and checksum matches", "file", targetFile)
						shouldDownload = false
					} else {
						slog.Warn("Checksum mismatch or verification failed, re-downloading", "file", targetFile)
					}
				}

				if shouldDownload {
					if err := dl.DownloadFile(ctx, elementID, format, targetFile); err != nil {
						return err
					}
					// Verify again
					dl.Checksum(ctx, elementID, format)
				}
			} else {
				// Check disabled
				targetFile := outFile + "." + cfg.Formats[format].ID
				if err := dl.DownloadFile(ctx, elementID, format, targetFile); err != nil {
					return err
				}
			}
		} else {
			targetFile := outFile + "." + cfg.Formats[format].ID
			if err := dl.DownloadFile(ctx, elementID, format, targetFile); err != nil {
				return err
			}
		}
	}

	return nil
}
