package packages_test

import (
	"testing"

	"path/filepath"

	"kego.io/context/vosctx"
	"kego.io/kerr/assert"
	. "kego.io/process/packages"
	"kego.io/process/tests"
)

func TestGetCurrentGopath(t *testing.T) {
	cb := tests.NewContextBuilder()
	abc := filepath.Join("a", "b", "c")
	def := filepath.Join("d", "e", "f")
	cb.OsVar("GOPATH", abc+string(filepath.ListSeparator)+def)
	gop := GetCurrentGopath(cb.Ctx())
	assert.Equal(t, abc, gop)
	cb.OsWd(filepath.Join("d", "e", "f", "g", "h"))
	gop = GetCurrentGopath(cb.Ctx())
	assert.Equal(t, def, gop)
}

func TestGetPackageFromDir(t *testing.T) {
	cb := tests.NewContextBuilder().TempGopath(false)
	defer cb.Cleanup()
	packagePath, packageDir := cb.TempPackage("a", map[string]string{})

	calculatedPath, err := GetPackageFromDir(cb.Ctx(), packageDir)
	assert.NoError(t, err)
	assert.Equal(t, packagePath, calculatedPath)

	vos := vosctx.FromContext(cb.Ctx())
	cb.OsVar("GOPATH", "/fdskljsfdash/"+string(filepath.ListSeparator)+vos.Getenv("GOPATH"))

	calculatedPath, err = GetPackageFromDir(cb.Ctx(), packageDir)
	assert.NoError(t, err)
	assert.Equal(t, packagePath, calculatedPath)

	cb.OsVar("GOPATH", "/fdskljsfdash/")
	_, err = GetPackageFromDir(cb.Ctx(), packageDir)
	assert.IsError(t, err, "CXOETFPTGM")
}

func TestGetDirFromEmptyPackage(t *testing.T) {
	cb := tests.NewContextBuilder().TempGopath(false)
	defer cb.Cleanup()
	packagePath, packageDir := cb.TempPackage("a", map[string]string{})

	_, err := GetDirFromEmptyPackage(cb.Ctx(), "a.b/c")
	assert.IsError(t, err, "SUTCWEVRXS")

	calculatedDir, err := GetDirFromEmptyPackage(cb.Ctx(), packagePath)
	assert.NoError(t, err)
	assert.Equal(t, packageDir, calculatedDir)

	vos := vosctx.FromContext(cb.Ctx())
	cb.OsVar("GOPATH", "/fdskljsfdash/"+string(filepath.ListSeparator)+vos.Getenv("GOPATH"))

	// This will now need two loops around to get the package
	calculatedDir, err = GetDirFromEmptyPackage(cb.Ctx(), packagePath)
	assert.NoError(t, err)
	assert.Equal(t, packageDir, calculatedDir)

}
func TestGetDirFromPackage(t *testing.T) {
	cb := tests.NewContextBuilder().TempGopath(false)
	defer cb.Cleanup()
	packagePath, packageDir := cb.TempPackage("a", map[string]string{})
	calculatedDir, err := GetDirFromPackage(cb.Ctx(), packagePath)
	assert.NoError(t, err)
	assert.Equal(t, packageDir, calculatedDir)

	cb.TempFile("a.go", "package a")

	calculatedDir, err = GetDirFromPackage(cb.Ctx(), packagePath)
	assert.NoError(t, err)
	assert.Equal(t, packageDir, calculatedDir)

}
