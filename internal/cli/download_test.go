package cli_test

import (
	"crypto/md5" //nolint:gosec // MD5 is used by provider checksum files.
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"sync/atomic"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
)

const (
	downloadHelperProcessEnv = "DOWNLOAD_GEOFABRIK_TEST_HELPER"
	downloadHelperArgsEnv    = "DOWNLOAD_GEOFABRIK_TEST_ARGS"
	downloadHelperErrorEnv   = "DOWNLOAD_GEOFABRIK_TEST_ERROR"

	testDownloadElement = "test-elem"
	testGeoJSONPath     = "/test-elem/success.geojson"
	testOsmPbfPath      = "/test-elem/failure.osm.pbf"
	testGeoJSONContent  = `{"type":"FeatureCollection","features":[]}`
	testOsmPbfContent   = "osm.pbf test content"
)

func TestDownloadCLIHelperProcess(t *testing.T) {
	if os.Getenv(downloadHelperProcessEnv) != "1" {
		return
	}

	var args []string
	if err := json.Unmarshal([]byte(os.Getenv(downloadHelperArgsEnv)), &args); err != nil {
		t.Fatalf("decode CLI arguments: %v", err)
	}

	cli.RootCmd.SetArgs(args)

	err := cli.Execute()

	expectedError := os.Getenv(downloadHelperErrorEnv)
	if expectedError == "" {
		require.NoError(t, err)

		return
	}

	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)
}

func TestDownloadMultiFormatCLIOutcomes(t *testing.T) {
	tests := []struct {
		name        string
		expectedErr string
		failOsmPbf  bool
	}{
		{
			name: "all formats download successfully",
		},
		{
			name:        "returns an error after one format succeeds",
			failOsmPbf:  true,
			expectedErr: "server status code 503",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case testGeoJSONPath:
					_, _ = w.Write([]byte(testGeoJSONContent))

				case testOsmPbfPath:
					if test.failOsmPbf {
						http.Error(w, "unavailable", http.StatusServiceUnavailable)

						return
					}

					_, _ = w.Write([]byte(testOsmPbfContent))

				default:
					http.NotFound(w, r)
				}
			}))
			t.Cleanup(server.Close)

			configFile := filepath.Join(t.TempDir(), "catalog.yml")
			configContent := fmt.Sprintf(`
elements:
  %s:
    id: %s
    files: ["geojson", "osm.pbf"]
formats:
  geojson:
    ext: "geojson"
    loc: "/success.geojson"
    baseurl: %q
  osm.pbf:
    ext: "osm.pbf"
    loc: "/failure.osm.pbf"
    baseurl: %q
`, testDownloadElement, testDownloadElement, server.URL, server.URL)
			require.NoError(t, os.WriteFile(configFile, []byte(configContent), 0o600))

			outputDir := t.TempDir()
			runDownloadCLIHelper(t, []string{
				"download",
				testDownloadElement,
				"--config", configFile,
				"--output-dir", outputDir,
				"--check=false",
				"--progress=false",
				"-g",
				"-P",
			}, test.expectedErr)

			geoJSON, err := os.ReadFile(filepath.Join(outputDir, testDownloadElement+".geojson"))
			require.NoError(t, err)
			assert.JSONEq(t, testGeoJSONContent, string(geoJSON))

			osmPbf, err := os.ReadFile(filepath.Join(outputDir, testDownloadElement+".osm.pbf"))
			if test.failOsmPbf {
				require.ErrorIs(t, err, os.ErrNotExist)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, []byte(testOsmPbfContent), osmPbf)
		})
	}
}

func TestDownloadCLI_ChecksumOutcomes(t *testing.T) {
	tests := []struct {
		name             string
		checksumContent  string
		expectedError    string
		existingContent  string
		includeChecksum  bool
		checksumNotFound bool
		force            bool
	}{
		{
			name:            "checksum mismatch fails the command",
			includeChecksum: true,
			checksumContent: "00000000000000000000000000000000  failure.osm.pbf\n",
			expectedError:   "checksum mismatch",
		},
		{
			name:            "force redownloads a stale file before mismatch fails",
			checksumContent: "00000000000000000000000000000000  failure.osm.pbf\n",
			expectedError:   "checksum mismatch",
			existingContent: "stale content",
			includeChecksum: true,
			force:           true,
		},
		{
			name: "missing checksum allows the command to succeed",
		},
		{
			name:             "unavailable advertised checksum allows the command to succeed",
			includeChecksum:  true,
			checksumNotFound: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case testOsmPbfPath:
					_, _ = w.Write([]byte(testOsmPbfContent))

				case testOsmPbfPath + ".md5":
					if test.checksumNotFound {
						http.NotFound(w, r)

						return
					}

					_, _ = fmt.Fprint(w, test.checksumContent)

				default:
					http.NotFound(w, r)
				}
			}))
			t.Cleanup(server.Close)

			elementFiles := `["osm.pbf"]`
			if test.includeChecksum {
				elementFiles = `["osm.pbf", "osm.pbf.md5"]`
			}

			configContent := fmt.Sprintf(`
elements:
  %s:
    id: %s
    files: %s
formats:
  osm.pbf:
    ext: "osm.pbf"
    loc: "/failure.osm.pbf"
    baseurl: %q
  osm.pbf.md5:
    ext: "osm.pbf.md5"
    loc: "/failure.osm.pbf.md5"
    baseurl: %q
`,
				testDownloadElement, testDownloadElement, elementFiles,
				server.URL,
				server.URL,
			)
			configFile := filepath.Join(t.TempDir(), "catalog.yml")
			require.NoError(t, os.WriteFile(configFile, []byte(configContent), 0o600))

			outputDir := t.TempDir()

			targetFile := filepath.Join(outputDir, testDownloadElement+".osm.pbf")
			if test.existingContent != "" {
				require.NoError(t, os.WriteFile(targetFile, []byte(test.existingContent), 0o600))
			}

			args := []string{
				"download",
				testDownloadElement,
				"--config", configFile,
				"--output-dir", outputDir,
				"--progress=false",
			}
			if test.force {
				args = append(args, "--force")
			}

			runDownloadCLIHelper(t, args, test.expectedError)

			downloadedFile, err := os.ReadFile(targetFile)
			require.NoError(t, err)
			assert.Equal(t, testOsmPbfContent, string(downloadedFile))
		})
	}
}

func TestDownloadCLIExistingDestinationRequiresForce(t *testing.T) {
	tests := []struct {
		name        string
		expectedErr string
		existing    bool
		force       bool
	}{
		{
			name:        "rejects an existing file with a matching checksum by default",
			expectedErr: "destination file already exists",
			existing:    true,
		},
		{
			name:     "force replaces an existing file with a matching checksum",
			existing: true,
			force:    true,
		},
		{
			name:  "force downloads when the destination is absent",
			force: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var dataRequestCount atomic.Int32

			var checksumRequestCount atomic.Int32

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.URL.Path {
				case "/test-elem/test-elem.osm.pbf":
					dataRequestCount.Add(1)

					_, _ = w.Write([]byte(testOsmPbfContent))

				case "/test-elem/test-elem.osm.pbf.md5":
					checksumRequestCount.Add(1)

					hash := md5.Sum([]byte(testOsmPbfContent)) //nolint:gosec // MD5 is used by provider checksum files.
					_, _ = fmt.Fprintf(w, "%x  test-elem.osm.pbf\n", hash)

				default:
					http.NotFound(w, r)
				}
			}))
			t.Cleanup(server.Close)

			configFile := filepath.Join(t.TempDir(), "catalog.yml")
			configContent := fmt.Sprintf(`
elements:
  %s:
    id: %s
    files: ["osm.pbf", "osm.pbf.md5"]
formats:
  osm.pbf:
    ext: "osm.pbf"
    loc: "/test-elem.osm.pbf"
    baseurl: %q
  osm.pbf.md5:
    ext: "osm.pbf.md5"
    loc: "/test-elem.osm.pbf.md5"
    baseurl: %q
`, testDownloadElement, testDownloadElement, server.URL, server.URL)
			require.NoError(t, os.WriteFile(configFile, []byte(configContent), 0o600))

			outputDir := t.TempDir()

			targetFile := filepath.Join(outputDir, testDownloadElement+".osm.pbf")
			if test.existing {
				require.NoError(t, os.WriteFile(targetFile, []byte(testOsmPbfContent), 0o600))
			}

			args := []string{
				"download",
				testDownloadElement,
				"--config", configFile,
				"--output-dir", outputDir,
				"--progress=false",
			}
			if test.force {
				args = append(args, "--force")
			}

			runDownloadCLIHelper(t, args, test.expectedErr)

			content, err := os.ReadFile(targetFile)
			require.NoError(t, err)
			assert.Equal(t, testOsmPbfContent, string(content))
			assert.NoFileExists(t, targetFile+".tmp")

			if test.expectedErr != "" {
				assert.Equal(t, int32(0), dataRequestCount.Load())
				assert.Equal(t, int32(0), checksumRequestCount.Load())
			} else {
				assert.Equal(t, int32(1), dataRequestCount.Load())
				assert.Equal(t, int32(1), checksumRequestCount.Load())
			}
		})
	}
}

func runDownloadCLIHelper(t *testing.T, args []string, expectedError string) {
	t.Helper()

	encodedArgs, err := json.Marshal(args)
	require.NoError(t, err)

	cmd := exec.CommandContext(t.Context(), os.Args[0], "-test.run=^TestDownloadCLIHelperProcess$") //nolint:gosec // os.Args[0] is this test binary.

	cmd.Env = append(os.Environ(),
		downloadHelperProcessEnv+"=1",
		downloadHelperArgsEnv+"="+string(encodedArgs),
		downloadHelperErrorEnv+"="+expectedError,
	)

	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "CLI helper failed: %s", output)
}

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
	cli.ResetGlobs()
	cli.RootCmd.SetArgs([]string{"download"})

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

func TestDownloadCmd_DefaultFormatNonPbf(t *testing.T) {
	tmpDir := t.TempDir()

	cli.ResetGlobs()
	viper.Reset()

	configFile := tmpDir + "/non_pbf.yml"
	content := `
elements:
  custom-elem:
    id: "custom-elem"
    files: ["o5m"]
formats:
  o5m:
    ext: "o5m"
    loc: "custom-elem.o5m"
`
	err := os.WriteFile(configFile, []byte(content), 0o600)
	require.NoError(t, err)

	cli.RootCmd.SetArgs([]string{"download", "custom-elem", "--config", configFile, "--nodownload", "--output-dir", tmpDir})

	err = cli.Execute()
	require.NoError(t, err)
}

func TestDownloadCmd_FormatNotAvailable(t *testing.T) {
	tmpDir := t.TempDir()

	cli.ResetGlobs()
	viper.Reset()

	configFile := tmpDir + "/non_pbf.yml"
	content := `
elements:
  custom-elem:
    id: "custom-elem"
    files: ["o5m"]
formats:
  o5m:
    ext: "o5m"
    loc: "custom-elem.o5m"
`
	err := os.WriteFile(configFile, []byte(content), 0o600)
	require.NoError(t, err)

	cli.RootCmd.SetArgs([]string{"download", "custom-elem", "-P", "--config", configFile, "--nodownload", "--output-dir", tmpDir})

	err = cli.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "format not")
}

func TestDownloadCmd_ElementNotFound(t *testing.T) {
	tmpDir := t.TempDir()

	cli.ResetGlobs()
	viper.Reset()

	configFile := tmpDir + "/geofabrik.yml"
	err := os.WriteFile(configFile, []byte(testConfigContent), 0o600)
	require.NoError(t, err)

	cli.RootCmd.SetArgs([]string{"download", "nonexistent-elem", "--config", configFile, "--nodownload", "--output-dir", tmpDir})

	err = cli.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "element not found")
}

func TestDownloadCmd_OSMTW(t *testing.T) {
	cli.ResetGlobs()
	viper.Reset()

	tmpDir := t.TempDir()

	cli.RootCmd.SetArgs([]string{
		"download", "taiwan",
		"--service", "osm.kcwu.csie.org",
		"--config", "../../osm.kcwu.csie.org.yml",
		"--nodownload",
		"--output-dir", tmpDir,
	})

	err := cli.Execute()
	require.NoError(t, err)
}

func TestDownloadCmd_PlanetOSMCH(t *testing.T) {
	cli.ResetGlobs()
	viper.Reset()

	tmpDir := t.TempDir()

	cli.RootCmd.SetArgs([]string{
		"download", "switzerland",
		"--service", "planet.osm.ch",
		"--config", "../../planet.osm.ch.yml",
		"--nodownload",
		"--output-dir", tmpDir,
	})

	err := cli.Execute()
	require.NoError(t, err)
}

func TestDownloadCmd_PlanetOSMCH_ExplicitP(t *testing.T) {
	cli.ResetGlobs()
	viper.Reset()

	tmpDir := t.TempDir()

	cli.RootCmd.SetArgs([]string{
		"download", "switzerland",
		"-P",
		"--service", "planet.osm.ch",
		"--config", "../../planet.osm.ch.yml",
		"--nodownload",
		"--output-dir", tmpDir,
	})

	err := cli.Execute()
	require.NoError(t, err)
}

func TestDownloadCmd_ValidArgsFunction(t *testing.T) {
	tmpDir := t.TempDir()

	cli.ResetGlobs()
	viper.Reset()

	configFile := tmpDir + "/geofabrik.yml"

	err := os.WriteFile(configFile, []byte(testConfigContent), 0o600)
	require.NoError(t, err)

	viper.SetConfigFile(configFile)

	// Test element completion with prefix
	completions, directive := cli.DownloadCmd.ValidArgsFunction(cli.DownloadCmd, []string{}, "test")
	assert.Equal(t, cobra.ShellCompDirectiveNoFileComp, directive)
	assert.Contains(t, completions, "test-elem")

	// Test with already specified argument (no further completion)
	completions, directive = cli.DownloadCmd.ValidArgsFunction(cli.DownloadCmd, []string{"test-elem"}, "")
	assert.Equal(t, cobra.ShellCompDirectiveNoFileComp, directive)
	assert.Empty(t, completions)
}
