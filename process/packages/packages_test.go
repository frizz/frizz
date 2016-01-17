package packages

import (
	"os"
	"strings"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestGetPackagePath(t *testing.T) {

	cb := tests.Context("").OsVar("GOPATH", "/Users/dave/go")

	dir := "/Users/dave/go/src/github.com/foo/bar"
	pkg, err := GetPackageFromDir(cb.Ctx(), dir)
	assert.NoError(t, err)
	assert.Equal(t, "github.com/foo/bar", pkg)

	cb.OsVar("GOPATH", strings.Join([]string{"/Users/another/path", "/Users/dave/go", "/one/more"}, string(os.PathListSeparator)))
	pkg, err = GetPackageFromDir(cb.Ctx(), dir)
	assert.NoError(t, err)
	assert.Equal(t, "github.com/foo/bar", pkg)

}

func TestGetPackageDir(t *testing.T) {

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
