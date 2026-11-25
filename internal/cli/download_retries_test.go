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

	tests := []struct {
		name      string
		flagValue string
		wantErr   string
		want      int
	}{
		{name: "zero disables retries", flagValue: "0", want: 0},
		{name: "positive value", flagValue: "1", want: 1},
		{name: "negative value returns error", flagValue: "-1", wantErr: "--retries must be non-negative"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.NoError(t, downloadCmd.Flags().Set("retries", test.flagValue))

			options, err := buildDownloadOptions(downloadCmd)
			if test.wantErr != "" {
				require.ErrorContains(t, err, test.wantErr)

				return
			}

			require.NoError(t, err)
			require.Equal(t, test.want, options.Retries)
		})
	}
}
