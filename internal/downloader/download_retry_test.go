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

	var requestCount atomic.Int32

	d := New(catalog.New(), &Options{Retries: 1})

	d.wait = func(_ context.Context, delay time.Duration) error {
		assert.Equal(t, initialRetryDelay, delay)

		return nil
	}
	d.client = &http.Client{
		Transport: roundTripperFunc(func(request *http.Request) (*http.Response, error) {
			if requestCount.Add(1) == 1 {
				return nil, &net.OpError{Op: "dial", Net: "tcp", Err: syscall.ECONNRESET}
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     make(http.Header),
				Body:       io.NopCloser(http.NoBody),
				Request:    request,
			}, nil
		}),
	}

	err := d.FromURL(context.Background(), "http://example.test/extract", filepath.Join(t.TempDir(), "extract.bin"))

	require.NoError(t, err)
	assert.EqualValues(t, 2, requestCount.Load())
}

func TestFromURLRetriesResponseReadErrors(t *testing.T) {
	t.Parallel()

	var requestCount atomic.Int32

	d := New(catalog.New(), &Options{Retries: 1})
	d.wait = func(_ context.Context, delay time.Duration) error {
		assert.Equal(t, initialRetryDelay, delay)

		return nil
	}
	d.client = &http.Client{
		Transport: roundTripperFunc(func(request *http.Request) (*http.Response, error) {
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
		}),
	}

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

			assert.Equal(t, test.want, retryDelay(test.retries, test.retryAfter, now))
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

	outputDir := t.TempDir()
	cfg := catalog.New()
	cfg.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf"}
	cfg.Formats["osm.pbf.md5"] = catalog.Format{ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"}
	cfg.Elements["monaco"] = catalog.Element{
		ID:      "monaco",
		Name:    "Monaco",
		Formats: []string{catalog.FormatOsmPbf, "osm.pbf.md5"},
	}
	cfg.BaseURL = server.URL

	d := New(cfg, &Options{Check: true, Retries: 1, OutputDirectory: outputDir + string(os.PathSeparator)})

	var waits int

	d.wait = func(_ context.Context, delay time.Duration) error {
		waits++

		assert.Equal(t, initialRetryDelay, delay)

		return nil
	}

	targetFile := filepath.Join(outputDir, "monaco.osm.pbf")
	require.NoError(t, os.WriteFile(targetFile, content, 0o600))

	assert.True(t, d.Checksum(context.Background(), "monaco", catalog.FormatOsmPbf))
	assert.EqualValues(t, 2, sidecarRequests.Load())
	assert.Equal(t, 1, waits)
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

	outputDir := t.TempDir()
	cfg := catalog.New()
	cfg.Formats[catalog.FormatOsmPbf] = catalog.Format{ID: catalog.FormatOsmPbf, Loc: "-latest.osm.pbf"}
	cfg.Formats["osm.pbf.md5"] = catalog.Format{ID: "osm.pbf.md5", Loc: "-latest.osm.pbf.md5"}
	cfg.Elements["monaco"] = catalog.Element{
		ID:      "monaco",
		Name:    "Monaco",
		Formats: []string{catalog.FormatOsmPbf, "osm.pbf.md5"},
	}
	cfg.BaseURL = server.URL

	d := New(cfg, &Options{Check: true, Retries: 2, OutputDirectory: outputDir + string(os.PathSeparator)})

	var waits int

	d.wait = func(_ context.Context, _ time.Duration) error {
		waits++

		return nil
	}

	require.NoError(t, os.WriteFile(filepath.Join(outputDir, "monaco.osm.pbf"), content, 0o600))
	assert.False(t, d.Checksum(context.Background(), "monaco", catalog.FormatOsmPbf))
	assert.EqualValues(t, 1, requestCount.Load())
	assert.Zero(t, waits)
}

func TestParseRetryAfterDeltaSeconds(t *testing.T) {
	t.Parallel()

	delay, ok := parseRetryAfter(strconv.Itoa(int(maxRetryDelay/time.Second)+1), time.Now())
	require.True(t, ok)
	assert.Equal(t, maxRetryDelay, delay)

	delay, ok = parseRetryAfter("0", time.Now())
	require.True(t, ok)
	assert.Zero(t, delay)

	_, ok = parseRetryAfter("-1", time.Now())
	assert.False(t, ok)

	_, ok = parseRetryAfter("not-a-date", time.Now())
	assert.False(t, ok)
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

		return copy(buffer, "partial"), syscall.ECONNRESET
	}

	return 0, io.EOF
}

func (*failingReadCloser) Close() error {
	return nil
}
