package cli_test

import (
	"io"
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
)

func BenchmarkCLIExecuteHelp(b *testing.B) {
	cli.RootCmd.SetArgs([]string{"--help"})
	cli.RootCmd.SetOut(io.Discard)
	cli.RootCmd.SetErr(io.Discard)
	b.Cleanup(func() {
		cli.RootCmd.SetOut(nil)
		cli.RootCmd.SetErr(nil)
	})

	for b.Loop() {
		_ = cli.Execute()
	}
}
