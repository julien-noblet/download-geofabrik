package cli_test

import (
	"testing"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
	"github.com/stretchr/testify/assert"
)

func TestRegisterMCPCmd(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		cli.RegisterMCPCmd()
	})
}
