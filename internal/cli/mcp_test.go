package cli_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/julien-noblet/download-geofabrik/internal/cli"
)

func TestRegisterMCPCmd(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		cli.RegisterMCPCmd()
	})
}
