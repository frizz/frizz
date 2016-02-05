package packages_test

import (
	"os"
	"strings"
	"testing"

	"kego.io/context/vosctx"
	"kego.io/kerr/assert"
	. "kego.io/process/packages"
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

	cb := tests.NewContextBuilder().TempGopath(false)
	defer cb.Cleanup()

	pathA, dirA, _ := cb.TempPackage("a", nil)

	vos := vosctx.FromContext(cb.Ctx())

	dir, err := GetDirFromEmptyPackage(pathA, vos.Getenv("GOPATH"))
	assert.NoError(t, err)
	assert.Equal(t, dirA, dir)

	dir, err = GetDirFromEmptyPackage("a.b/c", vos.Getenv("GOPATH"))
	assert.IsError(t, err, "SUTCWEVRXS")

}
