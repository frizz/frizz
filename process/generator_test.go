package process

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"path/filepath"

	"kego.io/kerr/assert"
	"kego.io/process/internal/pkgtest"
)

func TestGenerateFiles_path(t *testing.T) {

	namespace, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	path, dir, err := pkgtest.CreateTemporaryPackage(namespace, "z", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
	})

	err = GenerateFiles(F_MAIN, settings{dir: dir, path: path})
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated-structs.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "package z\n")

}

func TestGenerateAndRunCmd(t *testing.T) {

	namespace, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	path, dir, err := pkgtest.CreateTemporaryPackage(namespace, "d", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
		"d.go":   `package d`,
	})

	err = GenerateAndRunCmd(F_CMD_TYPES, settings{dir: dir, path: path})
	assert.NoError(t, err)

	bytes, err := ioutil.ReadFile(filepath.Join(dir, "types", "generated-types.go"))
	assert.NoError(t, err)
	source := string(bytes)
	assert.Contains(t, source, "system.RegisterType")
}

func TestGenerateFiles(t *testing.T) {

	namespace, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	path, dir, err := pkgtest.CreateTemporaryPackage(namespace, "d", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
		"b.json": `{"type": "c", "id": "b"}`,
		"d.go":   `package d`,
	})

	set := settings{dir: dir, path: path}
	err = GenerateFiles(F_MAIN, set)
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated-structs.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "json.RegisterType")

	// This will error because of unknown types in b.json
	err = GenerateFiles(F_TYPES, set)
	assert.IsError(t, err, "XYIUHERDHE")
	assert.HasError(t, err, "FKCPTUWJWW")

	os.Remove(filepath.Join(dir, "b.json"))

	err = GenerateFiles(F_TYPES, set)
	assert.NoError(t, err)

	bytes, err := ioutil.ReadFile(filepath.Join(dir, "types", "generated-types.go"))
	assert.NoError(t, err)
	source := string(bytes)
	assert.Contains(t, source, "system.RegisterType")

}

func Test_parseOptions(t *testing.T) {

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

	set, err = InitialiseManually(false, false, false, pathB)
	assert.NoError(t, err)
	assert.Equal(t, dirB, set.dir)
	assert.Equal(t, pathB, set.path)

	_, err = InitialiseManually(false, false, false, "")
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

func TestGetDir(t *testing.T) {
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
