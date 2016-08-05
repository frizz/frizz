package server

import (
	"testing"

	"github.com/davelondon/ktest/assert"
)

func TestIsInside(t *testing.T) {
	assert.NoError(t, checkFilename("/a", "/b"))
	assert.IsError(t, checkFilename("/a", "../b"), "CNMFTLYQFG")
	assert.NoError(t, checkFilename("/a/b", "../b/c"))
	assert.NoError(t, checkFilename("/a/", "/b"))
	assert.IsError(t, checkFilename("/a/", "../b"), "CNMFTLYQFG")
	assert.NoError(t, checkFilename("/a/b/", "../b/c"))
	assert.NoError(t, checkFilename("/a/b/", "../../a/b/c"))
	assert.IsError(t, checkFilename("/a/b/", "../../a/c/d"), "CNMFTLYQFG")
}
