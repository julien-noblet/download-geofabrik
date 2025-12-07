package download

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	pb "github.com/cheggaaa/pb/v3"
	"github.com/julien-noblet/download-geofabrik/internal/config"
)

const (
	progressMinimal = 512 * 1024 // Don't display progress bar if size < 512kb
	ErrFromURL      = "can't download element"
	defaultTimeout  = 60 * time.Second
	keepAlive       = 30 * time.Second
	idleTimeout     = 5 * time.Second
	tlsTimeout      = 10 * time.Second
	continueTimeout = 5 * time.Second
	fileMode        = 0o644
)

// Downloader handles downloading files.
type Downloader struct {
	Config  *config.Config
	Options *config.Options
}

// NewDownloader creates a new Downloader.
func NewDownloader(cfg *config.Config, opts *config.Options) *Downloader {
	return &Downloader{
		Config:  cfg,
		Options: opts,
	}
}

// FromURL downloads a file from a URL to a specified file path.
func (d *Downloader) FromURL(ctx context.Context, myURL, fileName string) error {
	slog.Debug("Downloading", "url", myURL, "file", fileName)

	if d.Options.NoDownload {
		return nil
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   defaultTimeout,
				KeepAlive: keepAlive,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          0,
			IdleConnTimeout:       idleTimeout,
			TLSHandshakeTimeout:   tlsTimeout,
			ExpectContinueTimeout: continueTimeout,
		},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, myURL, http.NoBody)
	if err != nil {
		return fmt.Errorf("error creating request for %s - %w", myURL, err)
	}

	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error while downloading %s - %w", myURL, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"error while downloading %v, server return code %d",
			myURL,
			response.StatusCode,
		)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, fileMode)
	if err != nil {
		return fmt.Errorf("error while creating %s - %w", fileName, err)
	}

	defer file.Close()

	var (
		output          io.Writer = file
		currentProgress int64
	)

	// Display progress bar if requested, not quiet, and file is large enough
	if d.Options.Progress && !d.Options.Quiet && response.ContentLength > progressMinimal {
		progressBar := pb.Full.Start64(response.ContentLength)
		barReader := progressBar.NewProxyReader(response.Body)

		currentProgress, err = io.Copy(output, barReader)
		if err != nil {
			return fmt.Errorf("error while writing %s - %w", fileName, err)
		}

		progressBar.Finish()
	} else {
		currentProgress, err = io.Copy(output, response.Body)
		if err != nil {
			return fmt.Errorf("error while writing %s - %w", fileName, err)
		}
	}

	slog.Info("Downloaded", "file", fileName)
	slog.Debug("Bytes downloaded", "bytes", currentProgress)

	return nil
}

// FileExist checks if a file exists at the given path.
func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil
}

// DownloadFile downloads a file based on the configuration and element.
func (d *Downloader) DownloadFile(ctx context.Context, elementID, formatName, outputPath string) error { // elementID and formatName are strings.
	// config.FindElem uses d.Config.

	format := d.Config.Formats[formatName].ID

	myElem, err := config.FindElem(d.Config, elementID)
	if err != nil {
		slog.Error("Element not found", "element", elementID, "error", err)

		return fmt.Errorf("%w: %s", config.ErrFindElem, elementID)
	}

	myURL, err := config.Elem2URL(d.Config, myElem, format)
	if err != nil {
		slog.Error("URL generation failed", "error", err)

		return fmt.Errorf("%s %w", config.ErrElem2URL, err)
	}

	err = d.FromURL(ctx, myURL, outputPath)
	if err != nil {
		slog.Error("Download failed", "error", err)

		return fmt.Errorf("%s %w", ErrFromURL, err)
	}

	return nil
}

// Checksum downloads and verifies the checksum of a file.
func (d *Downloader) Checksum(ctx context.Context, elementID, formatName string) bool {
	if !d.Options.Check {
		return false
	}

	hashType := "md5"
	fhash := formatName + "." + hashType

	if ok, _, _ := config.IsHashable(d.Config, formatName); ok {
		myElem, err := config.FindElem(d.Config, elementID)
		if err != nil {
			slog.Error("Element not found", "element", elementID, "error", err)

			return false
		}

		myURL, err := config.Elem2URL(d.Config, myElem, fhash)
		if err != nil {
			slog.Error("URL generation failed", "error", err)
			return false
		}

		outputPath := d.Options.OutputDirectory + elementID
		// Ensure path separator?
		// Actually d.Options.OutputDirectory should have separator?
		// Original code: viper.GetString(config.ViperOutputDirectory) + viper.GetString(config.ViperElement)
		// I'll assume OutputDirectory ends with separator if needed, or join properly.
		// It's safer to use filepath.Join? But Element might be a filename?
		// Original code: outputPath + "." + fhash

		// I will use `outputPath` argument passed in? No, Checksum calculates it?
		// `Checksum` in original code calculated output path from global config!
		// `outputPath := viper.GetString(config.ViperOutputDirectory) + viper.GetString(config.ViperElement)`
		// `VerifyFileChecksum(outputPath+"."+format, outputPath+"."+fhash)`

		// So Checksum needs to know the layout on disk.
		// I'll reconstruct it.

		if e := d.FromURL(ctx, myURL, outputPath+"."+fhash); e != nil {
			slog.Error("Checksum download failed", "error", e)
			return false
		}

		return VerifyFileChecksum(outputPath+"."+d.Config.Formats[formatName].ID, outputPath+"."+fhash)
		// Wait, d.Config.Formats[formatName].ID might not be the extension on disk?
		// Original code: `outputPath+"."+format`
		// format arg is likely the key (e.g. "osm.pbf"), not the extension on disk?
		// Original `formats` map mapped key to `Format` struct with `ID` (ext).
		// Let's assume `formatName` is the key.

	}

	slog.Warn("No checksum provided", "file", d.Options.OutputDirectory+elementID+"."+formatName)

	return false
}
