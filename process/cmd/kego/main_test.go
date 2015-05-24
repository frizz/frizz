package main

import (
	"os"
	"strings"
	"testing"

	"kego.io/assert"
)

func TestGetPackage(t *testing.T) {
	gopath := "/Users/dave/go"
	dir := "/Users/dave/go/src/github.com/foo/bar"
	pkg, err := getPackage(dir, gopath)
	assert.NoError(t, err)
	assert.Equal(t, "github.com/foo/bar", pkg)

	gopath = strings.Join([]string{"/Users/another/path", "/Users/dave/go", "/one/more"}, string(os.PathListSeparator))
	pkg, err = getPackage(dir, gopath)
	assert.NoError(t, err)
	assert.Equal(t, "github.com/foo/bar", pkg)
}
