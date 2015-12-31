package pkgutils

import (
	"os"
	"strings"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestGetPackagePath(t *testing.T) {

	gopath := "/Users/dave/go"
	dir := "/Users/dave/go/src/github.com/foo/bar"
	pkg, err := getPackageFromDir(dir, gopath)
	assert.NoError(t, err)
	assert.Equal(t, "github.com/foo/bar", pkg)

	gopath = strings.Join([]string{"/Users/another/path", "/Users/dave/go", "/one/more"}, string(os.PathListSeparator))
	pkg, err = getPackageFromDir(dir, gopath)
	assert.NoError(t, err)
	assert.Equal(t, "github.com/foo/bar", pkg)

}

func TestGetPackageDir(t *testing.T) {

	current, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(current)

	namespace, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	pathA, dirA, _, err := tests.CreateTemporaryPackage(namespace, "a", nil)

	dir, err := getDirFromEmptyPackage(pathA, os.Getenv("GOPATH"))
	assert.NoError(t, err)
	assert.Equal(t, dirA, dir)

	dir, err = getDirFromEmptyPackage("a.b/c", os.Getenv("GOPATH"))
	assert.IsError(t, err, "SUTCWEVRXS")

}
