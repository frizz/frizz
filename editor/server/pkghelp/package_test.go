package pkghelp

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
)

func TestDirEqual(t *testing.T) {
	assert.True(t, DirEqual("/a/b", "/a/b/"))
	assert.False(t, DirEqual("/a/b", "/a/b/c"))
	assert.False(t, DirEqual("/a/b/", "/a/b/c"))
}

func TestCheck(t *testing.T) {
	var full string
	var err error

	full, err = Check("/a", "/b", false)
	require.NoError(t, err)
	assert.Equal(t, "/a/b", full)

	full, err = Check("/a", "../b", false)
	require.IsError(t, err, "CNMFTLYQFG")

	full, err = Check("/a/b", "../b/c", false)
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c", full)

	full, err = Check("/a/", "/b", false)
	require.NoError(t, err)
	assert.Equal(t, "/a/b", full)

	full, err = Check("/a/", "../b", false)
	require.IsError(t, err, "CNMFTLYQFG")

	full, err = Check("/a/b/", "../b/c", false)
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c", full)

	full, err = Check("/a/b/", "../../a/b/c", false)
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c", full)

	full, err = Check("/a/b/", "../../a/c/d", false)
	require.IsError(t, err, "CNMFTLYQFG")

	full, err = Check("/a/b/", "c/d", false)
	require.IsError(t, err, "OABDMRQSJO")

	full, err = Check("/a/b/", "c/d", true)
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c/d", full)
}
