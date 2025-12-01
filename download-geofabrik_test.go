package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun_List(t *testing.T) {
	// This test assumes geofabrik.yml exists and is valid.
	// It also assumes running "list" command doesn't fail.
	args := []string{"list"}
	err := Run(args)
	assert.NoError(t, err)
}

func TestRun_InvalidArg(t *testing.T) {
	args := []string{"invalid"}
	err := Run(args)
	assert.Error(t, err)
}
