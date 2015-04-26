package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPackage(t *testing.T) {
	gopath := "/Users/dave/go"
	dir := "/Users/dave/go/src/github.com/foo/bar"
	pkg, err := getPackage(dir, gopath)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "github.com/foo/bar", pkg)
}

func TestGetPackageMultiplePath(t *testing.T) {
	gopath := strings.Join([]string{"/Users/another/path", "/Users/dave/go", "/one/more"}, string(os.PathListSeparator))
	dir := "/Users/dave/go/src/github.com/foo/bar"
	pkg, err := getPackage(dir, gopath)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "github.com/foo/bar", pkg)
}
