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
	"maps"
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
	ErrFromURL = errors.New("cannot download element")

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
	optsCopy := &Options{}
	if opts != nil {
		*optsCopy = *opts
		optsCopy.FormatFlags = maps.Clone(opts.FormatFlags)
	}

	return &Downloader{
		Catalog: cat,
		Options: optsCopy,
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

func (d *Downloader) opts() *Options {
	if d == nil || d.Options == nil {
		return &Options{}
	}

	return d.Options
}

// FromURL downloads a file from a URL to a specified file path.
func (d *Downloader) FromURL(ctx context.Context, myURL, fileName string) (err error) {
	slog.Debug("Downloading", "url", myURL, "file", fileName)

	if d.opts().NoDownload {
		return nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, myURL, http.NoBody)
	if err != nil {
		return fmt.Errorf("creating request for %s: %w", myURL, err)
	}

	client := cmp.Or(d.client, http.DefaultClient)

	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("downloading %s: %w", myURL, err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			closeErr := fmt.Errorf("closing response body for %s: %w", myURL, cerr)
			err = errors.Join(err, closeErr)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: downloading %s: server status code %d",
			ErrServerStatusCode, myURL, response.StatusCode)
	}

	return d.saveToFile(fileName, response)
}

func (d *Downloader) copyBody(dst io.Writer, response *http.Response) (int64, error) {
	opts := d.opts()
	if opts.Progress && !opts.Quiet && response.ContentLength > progressMinimal {
		progressBar := pb.Full.Start64(response.ContentLength)
		barReader := progressBar.NewProxyReader(response.Body)

		defer progressBar.Finish()

		written, err := io.Copy(dst, barReader)
		if err != nil {
			return written, fmt.Errorf("copying response with progress: %w", err)
		}

		return written, nil
	}

	written, err := io.Copy(dst, response.Body)
	if err != nil {
		return written, fmt.Errorf("copying response: %w", err)
	}

	return written, nil
}

func (d *Downloader) saveToFile(fileName string, response *http.Response) (err error) {
	tmpFileName := fileName + ".tmp"

	if err = os.MkdirAll(filepath.Dir(fileName), dirMode); err != nil {
		return fmt.Errorf("creating directory %s: %w", filepath.Dir(fileName), err)
	}

	file, err := os.OpenFile(tmpFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, fileMode)
	if err != nil {
		return fmt.Errorf("creating %s: %w", tmpFileName, err)
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

	hasher := md5.New() //nolint:gosec // MD5 is used to control with md5sum files
	multiWriter := io.MultiWriter(file, hasher)

	currentProgress, err := d.copyBody(multiWriter, response)
	if err != nil {
		return err
	}

	fileClosed = true

	if cerr := file.Close(); cerr != nil {
		return fmt.Errorf("closing %s: %w", tmpFileName, cerr)
	}

	if err := os.Rename(tmpFileName, fileName); err != nil {
		return fmt.Errorf("renaming %s to %s: %w", tmpFileName, fileName, err)
	}

	var digest [md5.Size]byte

	sum := hasher.Sum(digest[:0])
	d.lastHash.Store(fileName, hex.EncodeToString(sum))

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
	if d == nil || d.Catalog == nil {
		return catalog.ErrNilCatalog
	}

	formatDef, ok := d.Catalog.GetFormat(formatName)
	if !ok {
		return fmt.Errorf("%w: %s", catalog.ErrFormatNotFound, formatName)
	}

	format := formatDef.ID

	myElem, err := d.Catalog.Find(elementID)
	if err != nil {
		return fmt.Errorf("finding element %s: %w", elementID, err)
	}

	myURL, err := d.Catalog.ResolveURL(myElem, format)
	if err != nil {
		return fmt.Errorf("resolving url for %s: %w", elementID, err)
	}

	if err := d.FromURL(ctx, myURL, outputPath); err != nil {
		return fmt.Errorf("downloading from url: %w", err)
	}

	return nil
}

func (d *Downloader) verifyChecksum(targetFile, hashFile string) bool {
	var ret bool

	if inFlightHash, ok := d.lastHash.Load(targetFile); ok {
		cachedDigest, isStr := inFlightHash.(string)
		if !isStr || cachedDigest == "" {
			slog.Debug("Cached hash invalid, reading file", "file", targetFile)

			ret = VerifyFileChecksum(targetFile, hashFile)
		} else {
			slog.Debug("Using in-flight hash", "file", targetFile, "hash", cachedDigest)

			ret, _ = CheckFileHash(hashFile, cachedDigest)
		}
	} else {
		ret = VerifyFileChecksum(targetFile, hashFile)
	}

	if ret {
		slog.Info("Checksum OK", "file", targetFile)
	} else {
		slog.Error("Checksum MISMATCH", "file", targetFile)
	}

	return ret
}

// Checksum downloads and verifies the checksum of a file, using in-flight computed MD5 when available.
func (d *Downloader) Checksum(ctx context.Context, elementID, formatName string) bool {
	if d == nil || d.Catalog == nil || !d.opts().Check {
		return false
	}

	isHashable, _, _ := d.Catalog.IsHashable(formatName)
	if !isHashable {
		slog.Warn("No checksum provided",
			"element", elementID,
			"format", formatName,
			"file", filepath.Join(d.opts().OutputDirectory, elementID+"."+formatName),
		)

		return false
	}

	hashType := "md5"
	fhash := formatName + "." + hashType

	myElem, err := d.Catalog.Find(elementID)
	if err != nil {
		slog.Error("Element not found", "element", elementID, "error", err)

		return false
	}

	formatDef, hasFormat := d.Catalog.GetFormat(formatName)
	if !hasFormat {
		return false
	}

	if !myElem.Formats.Contains(fhash) {
		slog.Warn("No checksum provided",
			"element", elementID,
			"format", formatDef.ID,
			"file", filepath.Join(d.opts().OutputDirectory, elementID+"."+formatDef.ID),
		)

		return false
	}

	myURL, err := d.Catalog.ResolveURL(myElem, fhash)
	if err != nil {
		slog.Error("URL generation failed", "element", elementID, "hash_format", fhash, "error", err)

		return false
	}

	outputPath := d.opts().OutputDirectory + elementID
	targetFile := outputPath + "." + formatDef.ID
	hashFile := outputPath + "." + fhash

	if err := d.FromURL(ctx, myURL, hashFile); err != nil {
		slog.Error("Checksum download failed", "element", elementID, "url", myURL, "file", hashFile, "error", err)

		return false
	}
	defer d.lastHash.Delete(hashFile)

	return d.verifyChecksum(targetFile, hashFile)
}
