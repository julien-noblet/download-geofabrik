package cli_test

import (
	"os"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecuteHelp(t *testing.T) {
	// Execute() now returns validation errors or nil.
	// Since we can't easily capture stdout/stderr without more complex refactoring or redirection,
	// checking that it returns nil for cases where it should just print help/version is a basic start.

	// However, cobra's Execute might need args.
	// `rootCmd` is package level variable. We can modify it for tests but should be careful.

	// A safe way is to set args on rootCmd before calling Execute,
	// but Execute() in root.go calls initCLI() which might reset things or bind flags again.
	// Let's rely on the fact that Execute() returns error.

	// If we run without args, it prints help and returns nil (RunE calls cmd.Help()).
	// We need to inject args if we want to test other things.
	// But Execute() does not take args. It uses os.Args by default unless we change it.

	// So calling Execute() directly in test will try to parse actual test binary flags which might fail or be weird.
	// We should probably allow passing args to Execute or just test the commands directly?
	// But the goal is to test Execute().

	// Changing Execute to take args would be a bigger refactor.
	// Simpler: Just make successful call to Execute?
	// Without args, it just runs help.

	// IMPORTANT: cobra uses os.Args[1:] if we don't set args.
	// When running "go test", os.Args contains test flags.
	// This will likely cause unknown flag errors in cobra.

	// We should probably refactor Execute to allow passing args, or set os.Args mock?
	// Setting os.Args is risky.

	// Better: Refactor `Execute` to accept args?
	// Or just test `rootCmd` directly in other tests and acknowledge `Execute` is hard to test as is?

	// Let's try to verify `rootCmd` configuration instead.
	assert.NotNil(t, cli.RootCmd)
	assert.Equal(t, "download-geofabrik", cli.RootCmd.Use)
	cli.RootCmd.SetArgs([]string{"--help"})

	err := cli.Execute()
	assert.NoError(t, err)
}

func TestDefaultConfigLoading(t *testing.T) {
	// Create temp dir
	tmpDir := t.TempDir()

	// Create default config file
	configFile := tmpDir + "/geofabrik.yml"

	err := os.WriteFile(configFile, []byte(testConfigContent), 0o600)
	require.NoError(t, err)

	// Change cwd
	t.Chdir(tmpDir)

	// Reset globs
	cli.ResetGlobs()
	viper.Reset()

	// Run list command without --config
	// It should pick up geofabrik.yml in cwd
	cli.RootCmd.SetArgs([]string{"list"})

	err = cli.Execute()
	assert.NoError(t, err)
}
