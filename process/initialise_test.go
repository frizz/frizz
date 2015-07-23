package process

import (
	"os"
	"strings"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/internal/pkgtest"
)

func TestInitialise(t *testing.T) {

	current, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(current)

	namespace, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	pathA, dirA, err := pkgtest.CreateTemporaryPackage(namespace, "a", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
		"a.go":   "package a",
	})
	pathB, dirB, err := pkgtest.CreateTemporaryPackage(namespace, "b", map[string]string{
		"b.json": `{"type": "system:type", "id": "b"}`,
		"b.go":   "package b",
	})

	err = os.Chdir(dirA)
	assert.NoError(t, err)

	set, err := InitialiseAutomatic()
	assert.NoError(t, err)
	assert.Equal(t, dirA, set.dir)
	assert.Equal(t, pathA, set.path)

	err = os.Chdir("/")
	assert.NoError(t, err)

	set, err = InitialiseManually(false, false, false, false, false, pathB)
	assert.NoError(t, err)
	assert.Equal(t, dirB, set.dir)
	assert.Equal(t, pathB, set.path)

	_, err = InitialiseManually(false, false, false, false, false, "")
	assert.IsError(t, err, "PSRAWHQCPV")
	assert.HasError(t, err, "CXOETFPTGM")
}

func TestGetPackagePath(t *testing.T) {

	gopath := "/Users/dave/go"
	dir := "/Users/dave/go/src/github.com/foo/bar"
	pkg, err := getPackagePath(dir, gopath)
	assert.NoError(t, err)
	assert.Equal(t, "github.com/foo/bar", pkg)

	gopath = strings.Join([]string{"/Users/another/path", "/Users/dave/go", "/one/more"}, string(os.PathListSeparator))
	pkg, err = getPackagePath(dir, gopath)
	assert.NoError(t, err)
	assert.Equal(t, "github.com/foo/bar", pkg)

}

func TestGetPackageDir(t *testing.T) {

	current, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(current)

	namespace, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	pathA, dirA, err := pkgtest.CreateTemporaryPackage(namespace, "a", nil)

	dir, err := getPackageDir(pathA, os.Getenv("GOPATH"))
	assert.NoError(t, err)
	assert.Equal(t, dirA, dir)

	dir, err = getPackageDir("a.b/c", os.Getenv("GOPATH"))
	assert.IsError(t, err, "SUTCWEVRXS")

}