package cli_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
)

func TestGenerateCmd_UnknownService(t *testing.T) {
	// Reset globs not needed as we don't rely on it?
	// But Execute uses it.
	cli.ResetGlobs()

	cli.RootCmd.SetArgs([]string{"generate", "--service", "unknown-service"})

	err := cli.Execute()
	assert.Error(t, err)
	// Check if error message is relevant?
	// "generation failed: ..."
}
