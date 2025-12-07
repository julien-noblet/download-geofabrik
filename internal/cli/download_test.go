package cli_test

import (
	"os"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownloadCmd_NoDownload(t *testing.T) {
	// Setup
	// Reset viper to avoid pollution
	viper.Reset()

	// We need to call initCLI to set up flags on rootCmd?
	// Or just RegisterDownloadCmd?
	// downloadCmd is a global variable in the package, so we need to be careful.
	// Ideally we should reset flags.

	// Let's create a fresh environment if possible, but the code relies on globals.
	// We will trust valid usage for now.

	// Register command if not already (Execute does this, but we might verify individual command)
	// RegisterDownloadCmd() has side effects (adding to rootCmd).
	// If we run multiple tests, we might get duplicates if we keep calling it.
	// But `cobra` handles that gracefully? No, it might panic or duplicate.
	// Ideally, we test `downloadCmd` directly via its `RunE` or via `rootCmd` if properly set.

	// RegisterDownloadCmd() has side effects (adding to rootCmd).
	// We rely on Execute() to do this.

	// Let's try to run `downloadCmd` directly with `SetArgs`.
	// We need to initialize flags first.

	// We need to initialize viper bindings similar to main
	// initCLI also binds persistent flags.
	// Let's manually set up what we need.

	// Create a temporary directory for output
	tmpDir := t.TempDir()

	// Set args
	// We use "param" as element.
	// We use --nodownload to avoid network.
	// We need a config file, or defaults.
	// If no config file is found, it looks for "geofabrik.yml".
	// We should probably provide a dummy config or rely on defaults failing gracefully if we mocked network?
	// But --nodownload expects to read config to find the element URL even if it doesn't download?
	// Let's check `runDownload`.

	// `downloaderInstance := downloader.NewDownloader(cfg, opts)`
	// `processDownload` -> `downloaderInstance.DownloadFile`
	// `DownloadFile` -> `config.FindElem` -> `config.Elem2URL` -> `FromURL`
	// `FromURL` checks `d.Options.NoDownload`.

	// So we need a valid config that contains the element we are asking for,
	// OR we can mock the config loading?
	// Config loading happens in `runDownload`: `cfg, err := config.LoadConfig(opts.ConfigFile)`

	// If `config.LoadConfig` fails, `runDownload` returns error.
	// We need a valid config file.

	// Let's create a minimal config file.
	configFile := tmpDir + "/geofabrik.yml"

	err := os.WriteFile(configFile, []byte(testConfigContent), 0o600)
	require.NoError(t, err)

	// Set flags
	cli.RootCmd.SetArgs([]string{"download", "test-elem", "--config", configFile, "--nodownload", "--output-dir", tmpDir})

	// To properly test, we should run `rootCmd.Execute()` because `downloadCmd` is a subcommand.
	// `Execute` in `root.go` calls `initCLI` then `rootCmd.Execute()`.
	// But `Execute` also calls `Register...` which updates `rootCmd`.
	// If we call `Execute()`, it might double register if we are not careful?
	// `RegisterDownloadCmd` adds `downloadCmd` to `rootCmd`.
	// Cobra's `AddCommand` is idempotent-ish (adds to commands slice).

	// Let's rely on `rootCmd` being package global.
	// We can call `Execute()` but we need to prevent `initCLI` and `Register...` from messing up if called repeatedly?
	// The current implementation of `Execute` is:
	/*
	   func Execute() error {
	       initCLI()
	       RegisterDownloadCmd()
	       ...
	       return rootCmd.Execute()
	   }
	*/
	// If we test `Execute`, we are stuck with that logic.
	// Better to test `rootCmd.Execute()` directly after manual setup?

	// Setup for test
	// We need to make sure flags are registered.
	// Resetting globals is hard.

	// Workaround: Use a separate test harness that doesn't rely on `Execute` but sets up `rootCmd`.
	// But `initCLI` is private.
	// `RegisterDownloadCmd` is public.

	// Let's try calling `Execute` once per test?
	// Or just call `Execute()` and accept it.

	// Wait, `Execute` takes no args. It uses `os.Args`.
	// We can set `os.Args`?
	oldArgs := os.Args

	defer func() { os.Args = oldArgs }()

	os.Args = []string{"download-geofabrik", "download", "test-elem", "--config", configFile, "--nodownload", "--output-dir", tmpDir}

	// Run Execute
	err = cli.Execute()

	// Assert
	assert.NoError(t, err)
}

func TestDownloadCmd_InvalidArgs(t *testing.T) {
	oldArgs := os.Args

	defer func() { os.Args = oldArgs }()

	// Missing element arg
	os.Args = []string{"download-geofabrik", "download"}

	// Execute
	// Note: cobra might print to stderr.
	err := cli.Execute()

	require.Error(t, err)

	// Test invalid arguments (missing element)
	// Need to set up a dummy config file for this test case as well
	tmpDir := t.TempDir()

	configFile := tmpDir + "/geofabrik.yml"

	err = os.WriteFile(configFile, []byte(testConfigContent), 0o600)
	require.NoError(t, err)

	cli.RootCmd.SetArgs([]string{"download", "--config", configFile})

	err = cli.Execute()
	require.Error(t, err)
}

func TestDownloadCmd_NoCheck(t *testing.T) {
	tmpDir := t.TempDir()

	configFile := tmpDir + "/geofabrik.yml"

	err := os.WriteFile(configFile, []byte(testConfigContent), 0o600)
	require.NoError(t, err)

	cli.RootCmd.SetArgs([]string{"download", "test-elem", "--config", configFile, "--nodownload", "--output-dir", tmpDir, "--check=false"})

	err = cli.Execute()
	require.NoError(t, err)
}

func TestDownloadCmd_FileExists(t *testing.T) {
	// Test file exists path
	tmpDir := t.TempDir()

	configFile := tmpDir + "/geofabrik.yml"

	err := os.WriteFile(configFile, []byte(testConfigContent), 0o600)
	require.NoError(t, err)

	// Create dummy file
	targetFile := tmpDir + "/test-elem.osm.pbf"
	err = os.WriteFile(targetFile, []byte("dummy content"), 0o600)
	require.NoError(t, err)

	cli.RootCmd.SetArgs([]string{"download", "test-elem", "--config", configFile, "--nodownload", "--output-dir", tmpDir})

	err = cli.Execute()
	require.NoError(t, err)
}

func TestDownloadCmd_DefaultOutputDir(t *testing.T) {
	// Test default output dir (current directory)
	tmpDir := t.TempDir()

	configFile := tmpDir + "/geofabrik.yml"

	err := os.WriteFile(configFile, []byte(testConfigContent), 0o600)
	require.NoError(t, err)

	t.Chdir(tmpDir)

	// Reset globs
	// We can't easily reset private globs from cli_test without more exports or helpers.
	// But since we are running in a separate process for "go test" usually? No, same process.
	// `cli` package state persists.
	// Exporting `ResetGlobs`?

	// Let's create `ResetGlobs` in export_test.go as well.
	cli.ResetGlobs()

	viper.Reset()

	// Explicitly UNSET outputDir flag variable?
	// It's a string variable in download.go
	// We might need to reset that too via ResetGlobs.

	cli.RootCmd.SetArgs([]string{"download", "test-elem", "--config", configFile, "--nodownload"})

	err = cli.Execute()
	require.NoError(t, err)
}
