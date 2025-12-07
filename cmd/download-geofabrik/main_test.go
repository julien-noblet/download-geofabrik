package main

import (
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
