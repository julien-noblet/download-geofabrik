package downloader

import (
	"cmp"
	"context"
	"crypto/md5" //nolint:gosec // MD5 is used to control with md5sum files
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	pb "github.com/cheggaaa/pb/v3"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
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
	streamBufferSize    = 128 * 1024 // 128KB buffer for optimal socket and disk throughput
)

var (
	// ErrFromURL is returned when downloading an extract fails due to network or filesystem errors.
	ErrFromURL = errors.New("can't download element")

	// ErrServerStatusCode is returned when the remote HTTP server responds with a non-2xx status code.
	ErrServerStatusCode = errors.New("server return code error")
)

// Options holds runtime options configuring downloader execution.
type Options struct {
	// FormatFlags maps format CLI keys to their enabled state.
	FormatFlags map[string]bool

	// ConfigFile specifies the path to the YAML catalog file.
	ConfigFile string

	// Service specifies the extract provider service identifier.
	Service string

	// OutputDirectory specifies the target directory for downloaded files.
	OutputDirectory string

	// Check controls whether checksum verification is performed after download.
	Check bool

	// Verbose enables detailed debug logging.
	Verbose bool

	// Quiet suppresses informational log output.
	Quiet bool

	// NoDownload enables dry-run mode, printing what would be downloaded without network transfer.
	NoDownload bool

	// Progress controls whether a progress bar is displayed during download.
	Progress bool
}

// Downloader executes atomic extract downloads with in-flight hash computation and connection pooling.
type Downloader struct {
	// Catalog provides element metadata and format URL templates.
	Catalog *catalog.Catalog

	// Options contains execution options.
	Options *Options

	client   *http.Client
	lastHash sync.Map // map[string]string: filePath -> hexMD5 computed in-flight
}

// New creates a new Downloader with connection pooling and high-throughput buffers.
func New(cat *catalog.Catalog, opts *Options) *Downloader {
	return &Downloader{
		Catalog: cat,
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
				DisableCompression:    true, // Avoid decompression CPU overhead for already-compressed OSM files
				ReadBufferSize:        streamBufferSize,
				WriteBufferSize:       streamBufferSize,
			},
		},
	}
}

// NewDownloader creates a new Downloader with connection pooling and high-throughput buffers.
//
// Deprecated: Use New instead.
func NewDownloader(cat *catalog.Catalog, opts *Options) *Downloader {
	return New(cat, opts)
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

	client := cmp.Or(d.client, http.DefaultClient)

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

// saveToFile saves the response body to a file atomically with single-pass MD5 computation and progress bar support.
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

	hasher := md5.New() //nolint:gosec // MD5 is used for checksum control with provider md5 files
	writer := io.MultiWriter(file, hasher)

	currentProgress, err := d.copyBody(writer, response)
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

	var digest [md5.Size]byte
	hasher.Sum(digest[:0])
	d.lastHash.Store(fileName, hex.EncodeToString(digest[:]))

	slog.Info("Downloaded", "file", fileName)
	slog.Debug("Bytes downloaded", "bytes", currentProgress)

	return nil
}

// FileExists checks if a file exists at the given path.
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil
}

// FileExist checks if a file exists at the given path.
//
// Deprecated: Use FileExists instead.
func FileExist(filePath string) bool {
	return FileExists(filePath)
}

// DownloadFile downloads a file based on the catalog and element.
func (d *Downloader) DownloadFile(ctx context.Context, elementID, formatName, outputPath string) error {
	formatDef, ok := d.Catalog.Formats[formatName]
	if !ok {
		slog.Error("Format not found in config", "format", formatName)

		return fmt.Errorf("%w: %s", catalog.ErrFormatNotFound, formatName)
	}

	format := formatDef.ID

	myElem, err := d.Catalog.Find(elementID)
	if err != nil {
		slog.Error("Element not found", "element", elementID, "error", err)

		return fmt.Errorf("%w: %s", catalog.ErrElementNotFound, elementID)
	}

	myURL, err := d.Catalog.ResolveURL(myElem, format)
	if err != nil {
		slog.Error("URL generation failed", "error", err)

		return fmt.Errorf("can't find url: %w", err)
	}

	err = d.FromURL(ctx, myURL, outputPath)
	if err != nil {
		slog.Error("Download failed", "error", err)

		return fmt.Errorf("%w: %w", ErrFromURL, err)
	}

	return nil
}

func (d *Downloader) verifyChecksum(targetFile, hashFile string) bool {
	cachedVal, ok := d.lastHash.LoadAndDelete(targetFile)
	if !ok {
		return VerifyFileChecksum(targetFile, hashFile)
	}

	cachedDigest, isStr := cachedVal.(string)
	if !isStr || cachedDigest == "" {
		return VerifyFileChecksum(targetFile, hashFile)
	}

	slog.Debug("Using in-flight computed MD5", "file", targetFile, "hash", cachedDigest)

	ret, err := CheckFileHash(hashFile, cachedDigest)
	if err != nil {
		slog.Error("Checksum error", "error", err)
	}

	if ret {
		slog.Info("Checksum OK", "file", targetFile)
	} else {
		slog.Warn("Checksum MISMATCH", "file", targetFile)
	}

	return ret
}

// Checksum downloads and verifies the checksum of a file, using in-flight computed MD5 when available.
func (d *Downloader) Checksum(ctx context.Context, elementID, formatName string) bool {
	if !d.Options.Check {
		return false
	}

	ok, _, _ := d.Catalog.IsHashable(formatName)
	if !ok {
		slog.Warn("No checksum provided", "file", d.Options.OutputDirectory+elementID+"."+formatName)

		return false
	}

	hashType := "md5"
	fhash := formatName + "." + hashType

	myElem, err := d.Catalog.Find(elementID)
	if err != nil {
		slog.Error("Element not found", "element", elementID, "error", err)

		return false
	}

	if !myElem.Formats.Contains(fhash) {
		slog.Warn("No checksum provided", "file", d.Options.OutputDirectory+elementID+"."+d.Catalog.Formats[formatName].ID)

		return false
	}

	myURL, err := d.Catalog.ResolveURL(myElem, fhash)
	if err != nil {
		slog.Error("URL generation failed", "error", err)

		return false
	}

	outputPath := d.Options.OutputDirectory + elementID
	targetFile := outputPath + "." + d.Catalog.Formats[formatName].ID
	hashFile := outputPath + "." + fhash

	if e := d.FromURL(ctx, myURL, hashFile); e != nil {
		slog.Error("Checksum download failed", "error", e)

		return false
	}

	return d.verifyChecksum(targetFile, hashFile)
}
