package pkghelp

import (
	"testing"

	"github.com/davelondon/ktest/assert"
)

func TestDirEqual(t *testing.T) {
	assert.True(t, DirEqual("/a/b", "/a/b/"))
	assert.False(t, DirEqual("/a/b", "/a/b/c"))
	assert.False(t, DirEqual("/a/b/", "/a/b/c"))
}
