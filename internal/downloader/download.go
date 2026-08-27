package download

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	pb "github.com/cheggaaa/pb/v3"
	"github.com/julien-noblet/download-geofabrik/internal/config"
)

const (
	progressMinimal     = 512 * 1024 // Don't display progress bar if size < 512kb
	defaultTimeout      = 60 * time.Second
	keepAlive           = 30 * time.Second
	idleTimeout         = 90 * time.Second
	tlsTimeout          = 10 * time.Second
	continueTimeout     = 5 * time.Second
	fileMode            = 0o644
	dirMode             = 0o755
	maxIdleConns        = 100
	maxIdleConnsPerHost = 20
)

var (
	ErrFromURL          = errors.New("can't download element")
	ErrServerStatusCode = errors.New("server return code error")
)

// Downloader handles downloading files.
type Downloader struct {
	Config  *config.Config
	Options *config.Options
	client  *http.Client
}

// NewDownloader creates a new Downloader with connection pooling.
func NewDownloader(cfg *config.Config, opts *config.Options) *Downloader {
	return &Downloader{
		Config:  cfg,
		Options: opts,
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   defaultTimeout,
					KeepAlive: keepAlive,
				}).DialContext,
				MaxIdleConns:          maxIdleConns,
				MaxIdleConnsPerHost:   maxIdleConnsPerHost,
				IdleConnTimeout:       idleTimeout,
				TLSHandshakeTimeout:   tlsTimeout,
				ExpectContinueTimeout: continueTimeout,
			},
		},
	}
}

// FromURL downloads a file from a URL to a specified file path.
func (d *Downloader) FromURL(ctx context.Context, myURL, fileName string) (err error) {
	slog.Debug("Downloading", "url", myURL, "file", fileName)

	if d.Options.NoDownload {
		return nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, myURL, http.NoBody)
	if err != nil {
		return fmt.Errorf("error creating request for %s - %w", myURL, err)
	}

	client := d.client
	if client == nil {
		client = http.DefaultClient
	}

	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error while downloading %s - %w", myURL, err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("error while closing response body for %s - %w", myURL, cerr)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: error while downloading %v, server return code %d",
			ErrServerStatusCode, myURL, response.StatusCode)
	}

	return d.saveToFile(fileName, response)
}

func (d *Downloader) copyBody(dst io.Writer, response *http.Response) (int64, error) {
	if d.Options.Progress && !d.Options.Quiet && response.ContentLength > progressMinimal {
		progressBar := pb.Full.Start64(response.ContentLength)
		barReader := progressBar.NewProxyReader(response.Body)

		defer progressBar.Finish()

		written, err := io.Copy(dst, barReader)
		if err != nil {
			return written, fmt.Errorf("error copying response with progress: %w", err)
		}

		return written, nil
	}

	written, err := io.Copy(dst, response.Body)
	if err != nil {
		return written, fmt.Errorf("error copying response body: %w", err)
	}

	return written, nil
}

// saveToFile saves the response body to a file atomically with progress bar support.
func (d *Downloader) saveToFile(fileName string, response *http.Response) (err error) {
	if dir := filepath.Dir(fileName); dir != "" {
		if merr := os.MkdirAll(dir, dirMode); merr != nil {
			return fmt.Errorf("error creating directory %s - %w", dir, merr)
		}
	}

	tmpFileName := fileName + ".tmp"

	file, err := os.OpenFile(tmpFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, fileMode)
	if err != nil {
		return fmt.Errorf("error while creating %s - %w", tmpFileName, err)
	}

	fileClosed := false
	defer func() {
		if !fileClosed {
			_ = file.Close()
		}

		if err != nil {
			_ = os.Remove(tmpFileName)
		}
	}()

	currentProgress, err := d.copyBody(file, response)
	if err != nil {
		return fmt.Errorf("error while writing %s - %w", tmpFileName, err)
	}

	fileClosed = true

	if cerr := file.Close(); cerr != nil {
		return fmt.Errorf("error while closing %s - %w", tmpFileName, cerr)
	}

	if err := os.Rename(tmpFileName, fileName); err != nil {
		return fmt.Errorf("error while renaming %s to %s - %w", tmpFileName, fileName, err)
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
func (d *Downloader) DownloadFile(ctx context.Context, elementID, formatName, outputPath string) error {
	// elementID and formatName are strings.
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

		return fmt.Errorf("%w: %w", config.ErrElem2URL, err)
	}

	err = d.FromURL(ctx, myURL, outputPath)
	if err != nil {
		slog.Error("Download failed", "error", err)

		return fmt.Errorf("%w: %w", ErrFromURL, err)
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

		if e := d.FromURL(ctx, myURL, outputPath+"."+fhash); e != nil {
			slog.Error("Checksum download failed", "error", e)

			return false
		}

		return VerifyFileChecksum(outputPath+"."+d.Config.Formats[formatName].ID, outputPath+"."+fhash)
	}

	slog.Warn("No checksum provided", "file", d.Options.OutputDirectory+elementID+"."+formatName)

	return false
}
