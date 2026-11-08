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
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	initialRetryDelay   = time.Second
	maxRetryDelay       = 30 * time.Second
)

var (
	// ErrFromURL is returned when downloading an extract fails due to network or filesystem errors.
	ErrFromURL = errors.New("cannot download element")

	// ErrServerStatusCode is returned when the remote HTTP server responds with a non-2xx status code.
	ErrServerStatusCode = errors.New("server return code error")

	// ErrDestinationExists is returned when a download destination already exists and force is disabled.
	ErrDestinationExists = errors.New("destination file already exists")
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

	// Force permits replacing an existing output file.
	Force bool

	// Progress controls whether a progress bar is displayed during download.
	Progress bool

	// Retries is the number of additional attempts for transient download failures.
	Retries int
}

// Downloader executes atomic extract downloads with in-flight hash computation and connection pooling.
type Downloader struct {
	// Catalog provides element metadata and format URL templates.
	Catalog *catalog.Catalog

	// Options contains execution options.
	Options *Options

	client   *http.Client
	wait     func(context.Context, time.Duration) error
	lastHash sync.Map // map[string]string: filePath -> hexMD5 computed in-flight
}

var copyBufPool = sync.Pool{
	New: func() any {
		buf := make([]byte, streamBufferSize)

		return &buf
	},
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
				ResponseHeaderTimeout: defaultTimeout,
				ExpectContinueTimeout: continueTimeout,
				ForceAttemptHTTP2:     true,
				DisableCompression:    true, // Avoid decompression CPU overhead for already-compressed OSM files
				ReadBufferSize:        streamBufferSize,
				WriteBufferSize:       streamBufferSize,
			},
		},
		wait: waitContext,
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
	return d.fromURL(ctx, myURL, fileName, d.opts().Force)
}

func (d *Downloader) fromURL(ctx context.Context, myURL, fileName string, force bool) (err error) {
	slog.Debug("Downloading", "url", myURL, "file", fileName)

	if d.opts().NoDownload {
		return nil
	}

	if checkErr := d.checkDestination(fileName, force); checkErr != nil {
		return checkErr
	}

	for retriesPerformed := 0; ; retriesPerformed++ {
		retryable, retryAfter, downloadErr := d.downloadOnce(ctx, myURL, fileName, force)
		if downloadErr == nil {
			return nil
		}

		if !retryable || errors.Is(downloadErr, context.Canceled) ||
			errors.Is(downloadErr, context.DeadlineExceeded) || retriesPerformed >= d.opts().Retries {
			return downloadErr
		}

		delay := retryDelay(retriesPerformed, retryAfter, time.Now())

		wait := d.wait
		if wait == nil {
			wait = waitContext
		}

		if waitErr := wait(ctx, delay); waitErr != nil {
			return errors.Join(downloadErr, fmt.Errorf("waiting before retry: %w", waitErr))
		}
	}
}

func (d *Downloader) downloadOnce(ctx context.Context, myURL, fileName string, force bool) (retryable bool, retryAfter string, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, myURL, http.NoBody)
	if err != nil {
		return false, "", fmt.Errorf("creating request for %s: %w", myURL, err)
	}

	client := cmp.Or(d.client, http.DefaultClient)

	response, err := client.Do(req)
	if err != nil {
		return isTransientNetworkError(err), "", fmt.Errorf("downloading %s: %w", myURL, err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			closeErr := fmt.Errorf("closing response body for %s: %w", myURL, cerr)
			err = errors.Join(err, closeErr)
		}
	}()

	if response.StatusCode != http.StatusOK {
		statusErr := fmt.Errorf("%w: downloading %s: server status code %d",
			ErrServerStatusCode, myURL, response.StatusCode)

		return isRetryableStatus(response.StatusCode), response.Header.Get("Retry-After"), statusErr
	}

	err = d.saveToFile(fileName, response, force)

	var readErr *responseReadError

	return errors.As(err, &readErr), "", err
}

func (d *Downloader) checkDestination(fileName string, force bool) error {
	if force {
		return nil
	}

	if _, err := os.Lstat(fileName); err == nil {
		return fmt.Errorf("%w: %s", ErrDestinationExists, fileName)
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("checking destination %s: %w", fileName, err)
	}

	return nil
}

func (d *Downloader) copyBody(dst io.Writer, response *http.Response) (int64, error) {
	bufPtr, ok := copyBufPool.Get().(*[]byte)
	if !ok {
		buf := make([]byte, streamBufferSize)
		bufPtr = &buf
	}

	defer copyBufPool.Put(bufPtr)

	opts := d.opts()

	source := &trackingReader{reader: response.Body}
	if opts.Progress && !opts.Quiet && response.ContentLength > progressMinimal {
		progressBar := pb.Full.Start64(response.ContentLength)
		barReader := progressBar.NewProxyReader(source)

		defer progressBar.Finish()

		written, err := io.CopyBuffer(dst, barReader, *bufPtr)
		if err != nil {
			if source.err != nil {
				return written, &responseReadError{err: source.err}
			}

			return written, fmt.Errorf("copying response with progress: %w", err)
		}

		return written, nil
	}

	written, err := io.CopyBuffer(dst, source, *bufPtr)
	if err != nil {
		if source.err != nil {
			return written, &responseReadError{err: source.err}
		}

		return written, fmt.Errorf("copying response: %w", err)
	}

	return written, nil
}

type trackingReader struct {
	reader io.Reader
	err    error
}

func (r *trackingReader) Read(buffer []byte) (int, error) {
	bytesRead, err := r.reader.Read(buffer)
	if err != nil && !errors.Is(err, io.EOF) {
		r.err = err
	}

	if err != nil && !errors.Is(err, io.EOF) {
		return bytesRead, fmt.Errorf("reading response body: %w", err)
	}

	// io.EOF must remain identifiable as the end-of-stream sentinel.
	//nolint:wrapcheck // io.Reader.Read must return io.EOF unchanged.
	return bytesRead, err
}

type responseReadError struct {
	err error
}

func (e *responseReadError) Error() string {
	return fmt.Sprintf("reading response body: %v", e.err)
}

func (e *responseReadError) Unwrap() error {
	return e.err
}

func isRetryableStatus(statusCode int) bool {
	return statusCode == http.StatusTooManyRequests ||
		statusCode >= http.StatusInternalServerError && statusCode <= 599
}

func isTransientNetworkError(err error) bool {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}

	var requestErr *url.Error
	if !errors.As(err, &requestErr) {
		return false
	}

	if _, ok := errors.AsType[*net.OpError](requestErr.Err); ok {
		return true
	}

	networkErr, ok := errors.AsType[net.Error](requestErr.Err)

	return ok && networkErr.Timeout()
}

func retryDelay(retriesPerformed int, retryAfter string, now time.Time) time.Duration {
	if delay, ok := parseRetryAfter(retryAfter, now); ok {
		return min(delay, maxRetryDelay)
	}

	delay := initialRetryDelay
	for range retriesPerformed {
		if delay >= maxRetryDelay/2 {
			return maxRetryDelay
		}

		delay *= 2
	}

	return delay
}

func parseRetryAfter(value string, now time.Time) (time.Duration, bool) {
	value = strings.TrimSpace(value)

	if seconds, err := strconv.ParseInt(value, 10, 64); err == nil && seconds >= 0 {
		if seconds >= int64(maxRetryDelay/time.Second) {
			return maxRetryDelay, true
		}

		return time.Duration(seconds) * time.Second, true
	}

	retryTime, err := http.ParseTime(value)
	if err != nil {
		return 0, false
	}

	return max(retryTime.Sub(now), 0), true
}

func waitContext(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return fmt.Errorf("waiting for retry: %w", ctx.Err())
	case <-timer.C:
		return nil
	}
}

func (d *Downloader) saveToFile(fileName string, response *http.Response, force bool) (err error) {
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

	if err := publishFile(tmpFileName, fileName, force); err != nil {
		return err
	}

	var digest [md5.Size]byte

	sum := hasher.Sum(digest[:0])
	d.lastHash.Store(fileName, hex.EncodeToString(sum))

	slog.Info("Downloaded", "file", fileName)
	slog.Debug("Bytes downloaded", "bytes", currentProgress)

	return nil
}

func publishFile(tmpFileName, fileName string, force bool) error {
	if force {
		if err := os.Rename(tmpFileName, fileName); err != nil {
			return fmt.Errorf("renaming %s to %s: %w", tmpFileName, fileName, err)
		}

		return nil
	}

	if err := os.Link(tmpFileName, fileName); err != nil {
		if errors.Is(err, os.ErrExist) {
			return fmt.Errorf("%w: %s", ErrDestinationExists, fileName)
		}

		return fmt.Errorf("publishing %s as %s: %w", tmpFileName, fileName, err)
	}

	if err := os.Remove(tmpFileName); err != nil {
		return fmt.Errorf("removing temporary file %s: %w", tmpFileName, err)
	}

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

// DownloadFile downloads an extract file by element ID and format name.
func (d *Downloader) DownloadFile(ctx context.Context, elementID, formatName, outputPath string) error {
	myElem, err := d.Catalog.Find(elementID)
	if err != nil {
		return fmt.Errorf("finding element %s: %w", elementID, err)
	}

	myURL, err := d.Catalog.ResolveURL(myElem, formatName)
	if err != nil {
		return fmt.Errorf("generating URL for %s (%s): %w", elementID, formatName, err)
	}

	slog.Info("Starting download", "element", elementID, "format", formatName, "url", myURL)

	return d.FromURL(ctx, myURL, outputPath)
}

func (d *Downloader) verifyChecksum(targetFile, hashFile string) (bool, error) {
	var (
		ret bool
		err error
	)

	if inFlightHash, ok := d.lastHash.Load(targetFile); ok {
		cachedDigest, isStr := inFlightHash.(string)
		if !isStr || cachedDigest == "" {
			slog.Debug("Cached hash invalid, reading file", "file", targetFile)

			ret, err = VerifyFileChecksum(targetFile, hashFile)
		} else {
			slog.Debug("Using in-flight hash", "file", targetFile, "hash", cachedDigest)

			ret, err = CheckFileHash(hashFile, cachedDigest)
		}
	} else {
		ret, err = VerifyFileChecksum(targetFile, hashFile)
	}

	if err != nil {
		return false, err
	}

	if ret {
		slog.Info("Checksum OK", "file", targetFile)
	}

	return ret, nil
}

type checksumInfo struct {
	url        string
	formatID   string
	hashFormat string
}

func (d *Downloader) checksumInfo(elementID, formatName string) (checksumInfo, bool, error) {
	isHashable, hashFormat, _ := d.Catalog.IsHashable(formatName)
	if !isHashable {
		slog.Warn("No checksum provided",
			"element", elementID,
			"format", formatName,
			"file", filepath.Join(d.opts().OutputDirectory, elementID+"."+formatName),
		)

		return checksumInfo{}, false, nil
	}

	myElem, err := d.Catalog.Find(elementID)
	if err != nil {
		return checksumInfo{}, false, fmt.Errorf("finding element %s for checksum: %w", elementID, err)
	}

	formatDef, hasFormat := d.Catalog.GetFormat(formatName)
	if !hasFormat {
		return checksumInfo{}, false, fmt.Errorf("%w: %s", catalog.ErrFormatNotFound, formatName)
	}

	if !myElem.Formats.Contains(hashFormat) {
		slog.Warn("No checksum provided",
			"element", elementID,
			"format", formatDef.ID,
			"file", filepath.Join(d.opts().OutputDirectory, elementID+"."+formatDef.ID),
		)

		return checksumInfo{}, false, nil
	}

	myURL, err := d.Catalog.ResolveURL(myElem, hashFormat)
	if err != nil {
		slog.Warn("Checksum unavailable",
			"element", elementID,
			"hash_format", hashFormat,
			"error", err,
		)

		return checksumInfo{}, false, nil
	}

	return checksumInfo{
		url:        myURL,
		formatID:   formatDef.ID,
		hashFormat: hashFormat,
	}, true, nil
}

// Checksum downloads and verifies the checksum of a file, using in-flight computed MD5 when available.
func (d *Downloader) Checksum(ctx context.Context, elementID, formatName string) (bool, error) {
	if d == nil || d.Catalog == nil || !d.opts().Check {
		return false, nil
	}

	checksum, available, err := d.checksumInfo(elementID, formatName)
	if err != nil {
		return false, err
	}

	if !available {
		return false, nil
	}

	outputPath := filepath.Join(d.opts().OutputDirectory, elementID)
	targetFile := outputPath + "." + checksum.formatID
	hashFile := outputPath + "." + checksum.hashFormat

	if downloadErr := d.FromURL(ctx, checksum.url, hashFile); downloadErr != nil {
		slog.Warn("Checksum unavailable", "element", elementID, "url", checksum.url, "file", hashFile, "error", downloadErr)

		return false, nil
	}
	defer d.lastHash.Delete(hashFile)
	defer d.lastHash.Delete(targetFile)

	verified, err := d.verifyChecksum(targetFile, hashFile)
	if err != nil {
		return false, fmt.Errorf("verifying checksum for %s: %w", targetFile, err)
	}

	return verified, nil
}
