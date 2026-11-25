//nolint:testpackage // Tests set the private wait hook to avoid real retry delays.
package downloader

import (
	"context"
	"crypto/md5" //nolint:gosec // MD5 is used by the provider checksum format
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

func TestFromURLRetriesTransientStatuses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		retryAfter string
		wantError  string
		statuses   []int
		wantWaits  []time.Duration
		retries    int
		wantCount  int32
	}{
		{
			name:      "retries server errors with exponential backoff",
			statuses:  []int{http.StatusServiceUnavailable, http.StatusBadGateway, http.StatusOK},
			retries:   2,
			wantCount: 3,
			wantWaits: []time.Duration{time.Second, 2 * time.Second},
		},
		{
			name:       "retries too many requests using Retry-After",
			statuses:   []int{http.StatusTooManyRequests, http.StatusOK},
			retries:    1,
			retryAfter: "4",
			wantCount:  2,
			wantWaits:  []time.Duration{4 * time.Second},
		},
		{
			name:      "does not retry other client errors",
			statuses:  []int{http.StatusBadRequest, http.StatusOK},
			retries:   2,
			wantError: "server status code 400",
			wantCount: 1,
		},
		{
			name:      "zero retries allows only the initial attempt",
			statuses:  []int{http.StatusServiceUnavailable, http.StatusOK},
			retries:   0,
			wantError: "server status code 503",
			wantCount: 1,
		},
		{
			name:      "returns the final failure after retries are exhausted",
			statuses:  []int{http.StatusServiceUnavailable, http.StatusBadGateway, http.StatusOK},
			retries:   1,
			wantError: "server status code 502",
			wantCount: 2,
			wantWaits: []time.Duration{time.Second},
		},
		{
			name:       "caps Retry-After",
			statuses:   []int{http.StatusServiceUnavailable, http.StatusOK},
			retries:    1,
			retryAfter: "60",
			wantCount:  2,
			wantWaits:  []time.Duration{maxRetryDelay},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var requestCount atomic.Int32

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				attempt := int(requestCount.Add(1)) - 1
				status := test.statuses[min(attempt, len(test.statuses)-1)]

				if test.retryAfter != "" {
					w.Header().Set("Retry-After", test.retryAfter)
				}

				if status != http.StatusOK {
					w.WriteHeader(status)

					return
				}

				_, _ = w.Write([]byte("downloaded"))
			}))
			t.Cleanup(server.Close)

			d := New(catalog.New(), &Options{Retries: test.retries})

			var waits []time.Duration

			d.wait = func(_ context.Context, delay time.Duration) error {
				waits = append(waits, delay)

				return nil
			}

			outputPath := filepath.Join(t.TempDir(), "extract.bin")
			err := d.FromURL(context.Background(), server.URL, outputPath)

			if test.wantError != "" {
				require.ErrorIs(t, err, ErrServerStatusCode)
				require.ErrorContains(t, err, test.wantError)
			} else {
				require.NoError(t, err)

				content, readErr := os.ReadFile(outputPath)
				require.NoError(t, readErr)
				assert.Equal(t, "downloaded", string(content))
			}

			assert.Equal(t, test.wantCount, requestCount.Load())
			assert.Equal(t, test.wantWaits, waits)
		})
	}
}

func TestFromURLRetriesNetworkErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		networkErr error
		name       string
	}{
		{name: "net.OpError", networkErr: &net.OpError{Op: "dial", Net: "tcp", Err: syscall.ECONNRESET}},
		{name: "io.EOF", networkErr: io.EOF},
		{name: "io.ErrUnexpectedEOF", networkErr: io.ErrUnexpectedEOF},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var requestCount atomic.Int32

			d := newRetryTestDownloader(t, 1, func(request *http.Request) (*http.Response, error) {
				if requestCount.Add(1) == 1 {
					return nil, test.networkErr
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     make(http.Header),
					Body:       io.NopCloser(http.NoBody),
					Request:    request,
				}, nil
			})

			err := d.FromURL(context.Background(), "http://example.test/extract", filepath.Join(t.TempDir(), "extract.bin"))

			require.NoError(t, err)
			assert.EqualValues(t, 2, requestCount.Load())
		})
	}
}

func TestFromURLRetriesResponseReadErrors(t *testing.T) {
	t.Parallel()

	var requestCount atomic.Int32

	d := newRetryTestDownloader(t, 1, func(request *http.Request) (*http.Response, error) {
		if requestCount.Add(1) == 1 {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       &failingReadCloser{},
				Request:    request,
			}, nil
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("complete")),
			Request:    request,
		}, nil
	})

	outputPath := filepath.Join(t.TempDir(), "extract.bin")
	require.NoError(t, d.FromURL(context.Background(), "http://example.test/extract", outputPath))

	content, err := os.ReadFile(outputPath)
	require.NoError(t, err)
	assert.Equal(t, "complete", string(content))
	assert.EqualValues(t, 2, requestCount.Load())
}

func TestRetryDelay(t *testing.T) {
	t.Parallel()

	now := time.Date(2026, time.January, 2, 3, 4, 5, 0, time.UTC)

	tests := []struct {
		name       string
		retryAfter string
		retries    int
		want       time.Duration
	}{
		{name: "initial exponential delay", retries: 0, want: time.Second},
		{name: "doubles exponential delay", retries: 2, want: 4 * time.Second},
		{name: "caps exponential delay", retries: 10, want: maxRetryDelay},
		{name: "parses Retry-After date", retries: 0, retryAfter: now.Add(5 * time.Second).Format(http.TimeFormat), want: 5 * time.Second},
		{name: "caps Retry-After date", retries: 0, retryAfter: now.Add(time.Minute).Format(http.TimeFormat), want: maxRetryDelay},
		{name: "invalid Retry-After falls back to backoff", retries: 1, retryAfter: "invalid", want: 2 * time.Second},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			delay, hasDelay := parseRetryAfter(test.retryAfter, now)
			decision := retryDecision{delay: delay, hasDelay: hasDelay}

			assert.Equal(t, test.want, retryDelay(test.retries, decision))
		})
	}
}

func TestChecksumRetriesSidecarFetch(t *testing.T) {
	t.Parallel()

	content := []byte("extract content")
	hash := md5.Sum(content) //nolint:gosec // MD5 is used by the provider checksum format
	hashContent := fmt.Sprintf("%x  monaco-latest.osm.pbf\n", hash)

	var sidecarRequests atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if sidecarRequests.Add(1) == 1 {
			http.Error(w, "temporarily unavailable", http.StatusServiceUnavailable)

			return
		}

		_, _ = w.Write([]byte(hashContent))
	}))
	t.Cleanup(server.Close)

	d, waits := newChecksumTestDownloader(t, server.URL, 1, content)

	checksumOK, err := d.Checksum(context.Background(), "monaco", catalog.FormatOsmPbf)
	require.NoError(t, err)
	assert.True(t, checksumOK)
	assert.EqualValues(t, 2, sidecarRequests.Load())
	assert.Equal(t, 1, *waits)
}

func TestChecksumMismatchDoesNotRetry(t *testing.T) {
	t.Parallel()

	content := []byte("extract content")

	var requestCount atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestCount.Add(1)

		_, _ = w.Write([]byte("00000000000000000000000000000000  monaco-latest.osm.pbf\n"))
	}))
	t.Cleanup(server.Close)

	d, waits := newChecksumTestDownloader(t, server.URL, 2, content)

	checksumOK, err := d.Checksum(context.Background(), "monaco", catalog.FormatOsmPbf)
	assert.False(t, checksumOK)
	require.ErrorContains(t, err, "checksum mismatch")
	assert.EqualValues(t, 1, requestCount.Load())
	assert.Zero(t, *waits)
}

func TestParseRetryAfterDeltaSeconds(t *testing.T) {
	t.Parallel()

	now := time.Now()

	tests := []struct {
		name   string
		value  string
		want   time.Duration
		wantOK bool
	}{
		{name: "caps above maxRetryDelay", value: strconv.Itoa(int(maxRetryDelay/time.Second) + 1), want: maxRetryDelay, wantOK: true},
		{name: "zero seconds", value: "0", want: 0, wantOK: true},
		{name: "negative seconds rejected", value: "-1", wantOK: false},
		{name: "invalid string rejected", value: "not-a-date", wantOK: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			delay, ok := parseRetryAfter(test.value, now)
			assert.Equal(t, test.wantOK, ok)

			if test.wantOK {
				assert.Equal(t, test.want, delay)
			}
		})
	}
}

func newRetryTestDownloader(t *testing.T, retries int, transport roundTripperFunc) *Downloader {
	t.Helper()

	d := New(catalog.New(), &Options{Retries: retries})
	d.wait = func(_ context.Context, delay time.Duration) error {
		assert.Equal(t, initialRetryDelay, delay)

		return nil
	}
	d.client = &http.Client{Transport: transport}

	return d
}

func newChecksumTestDownloader(t *testing.T, serverURL string, retries int, content []byte) (downloader *Downloader, waitCount *int) {
	t.Helper()

	outputDir := t.TempDir()
	cfg := catalog.New()
	cfg.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf"}
	cfg.Formats["osm.pbf.md5"] = catalog.Format{ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"}
	cfg.Elements["monaco"] = catalog.Element{
		ID:      "monaco",
		Name:    "Monaco",
		Formats: []string{catalog.FormatOsmPbf, "osm.pbf.md5"},
	}
	cfg.BaseURL = serverURL

	d := New(cfg, &Options{Check: true, Retries: retries, OutputDirectory: outputDir + string(os.PathSeparator)})

	var waits int

	d.wait = func(_ context.Context, delay time.Duration) error {
		waits++

		assert.Equal(t, initialRetryDelay, delay)

		return nil
	}

	require.NoError(t, os.WriteFile(filepath.Join(outputDir, "monaco.osm.pbf"), content, 0o600))

	return d, &waits
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

type failingReadCloser struct {
	readOnce bool
}

func (r *failingReadCloser) Read(buffer []byte) (int, error) {
	if !r.readOnce {
		r.readOnce = true

		return copy(buffer, "partial"), &net.OpError{Op: "read", Net: "tcp", Err: syscall.ECONNRESET}
	}

	return 0, io.EOF
}

func (*failingReadCloser) Close() error {
	return nil
}
