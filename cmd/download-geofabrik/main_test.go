package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
)

func TestMainBuild(t *testing.T) {
	// Verify that the command builds successfully
	tmpDir := t.TempDir()

	binPath := filepath.Join(tmpDir, "download-geofabrik")

	cmd := exec.CommandContext(t.Context(), "go", "build", "-o", binPath, ".")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "Build failed: %s", output)

	// Verify basic run (help)
	cmd = exec.CommandContext(t.Context(), binPath, "--help")
	output, err = cmd.CombinedOutput()
	require.NoError(t, err)
	assert.Contains(t, string(output), "download-geofabrik")
}

func TestRun(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var stderr bytes.Buffer

		code := run(&stderr, func() error { return nil })

		assert.Equal(t, 0, code)
		assert.Empty(t, stderr.String())
		assert.Equal(t, version, cli.Version)
	})

	t.Run("error", func(t *testing.T) {
		var stderr bytes.Buffer

		code := run(&stderr, func() error { return errors.New("boom") })

		assert.Equal(t, 1, code)
		assert.Contains(t, stderr.String(), "Error: boom")
	})
}

func TestMain_Help(t *testing.T) {
	origArgs := os.Args

	t.Cleanup(func() {
		os.Args = origArgs
	})

	os.Args = []string{"download-geofabrik", "--help"}

	main()
}
