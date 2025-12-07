package cli_test

import (
	"os"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListCmd(t *testing.T) {
	// Setup mock config
	tmpDir := t.TempDir()

	configFile := tmpDir + "/geofabrik.yml"

	err := os.WriteFile(configFile, []byte(testConfigContent), 0o600)
	require.NoError(t, err)

	// Save original args
	oldArgs := os.Args

	defer func() { os.Args = oldArgs }()

	// Reset global state
	cli.ResetGlobs()
	viper.Reset()
	// viper.Reset clears everything. initCLI is not run again (sync.Once).
	// This might break flag bindings if viper needs them.
	// But initConfig runs and binds config.
	// We need to re-bind flags manually?
	// The flags are bound in initCLI: viper.BindPFlag...
	// If we reset viper, we lose those bindings.
	// So viper.GetBool("verbose") might fail.
	// But list command doesn't use verbose flag logic heavily in the test?
	// runList uses viper.ConfigFileUsed().

	// If we validly set cfgFile, initConfig sets viper config file.

	// Test regular list
	cli.RootCmd.SetArgs([]string{"list", "--config", configFile})

	err = cli.Execute()
	require.NoError(t, err)

	// Test markdown list
	cli.RootCmd.SetArgs([]string{"list", "--config", configFile, "--markdown"})

	err = cli.Execute()
	assert.NoError(t, err)
}
