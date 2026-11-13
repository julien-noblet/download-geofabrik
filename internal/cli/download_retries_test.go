//nolint:testpackage // Tests verify private option wiring without exposing a production accessor.
package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadRetriesFlag(t *testing.T) {
	setupCLI()

	defer func() {
		_ = downloadCmd.Flags().Set("retries", "2")
	}()

	retriesFlag := downloadCmd.Flags().Lookup("retries")
	require.NotNil(t, retriesFlag)
	require.Equal(t, "2", retriesFlag.DefValue)

	require.NoError(t, downloadCmd.Flags().Set("retries", "0"))
	options, err := buildDownloadOptions(downloadCmd)
	require.NoError(t, err)
	require.Zero(t, options.Retries)

	require.NoError(t, downloadCmd.Flags().Set("retries", "1"))
	options, err = buildDownloadOptions(downloadCmd)
	require.NoError(t, err)
	require.Equal(t, 1, options.Retries)

	require.NoError(t, downloadCmd.Flags().Set("retries", "-1"))
	_, err = buildDownloadOptions(downloadCmd)
	require.ErrorContains(t, err, "--retries must be non-negative")
}
