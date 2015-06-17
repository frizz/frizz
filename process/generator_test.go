package process

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"path/filepath"

	"kego.io/kerr/assert"
)

func TestGenerateFiles_path(t *testing.T) {

	currentDir, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(currentDir)

	dir, err := ioutil.TempDir(currentDir, "temporary")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	currentGopath := os.Getenv("GOPATH")
	defer os.Setenv("GOPATH", currentGopath)
	os.Setenv("GOPATH", dir)

	// in this case, the package name is not the same as the
	// folder, so we must use the path flag
	pkgDir := filepath.Join(dir, "src", "x.y", "funnyfolder")
	os.MkdirAll(pkgDir, 0777)

	os.Chdir(pkgDir)

	data := `{"type": "system:type", "id": "a"}`
	err = ioutil.WriteFile(filepath.Join(pkgDir, "a.json"), []byte(data), 0777)
	assert.NoError(t, err)

	// we correct the path by setting the path flag
	defer func() { *generatorPathFlag = "" }()
	*generatorPathFlag = "x.y/z"

	err = GenerateFiles(F_MAIN)
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(pkgDir, "generated.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "package z\n")

}

func TestGenerateAndRunCmd(t *testing.T) {

	currentDir, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(currentDir)

	dir, err := ioutil.TempDir(currentDir, "temporary")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	currentGopath := os.Getenv("GOPATH")
	defer os.Setenv("GOPATH", currentGopath)
	// with this test, we will be compiling and executing
	// the temporary command, so we still need access to the
	// real gopath
	newGopath := strings.Join([]string{currentGopath, dir}, string(os.PathListSeparator))
	os.Setenv("GOPATH", newGopath)

	pkgDir := filepath.Join(dir, "src", "b.c", "d")
	os.MkdirAll(pkgDir, 0777)

	os.Chdir(pkgDir)

	data := `{"type": "system:type", "id": "a"}`
	err = ioutil.WriteFile(filepath.Join(pkgDir, "a.json"), []byte(data), 0777)
	assert.NoError(t, err)

	err = ioutil.WriteFile(filepath.Join(pkgDir, "d.go"), []byte("package d"), 0777)
	assert.NoError(t, err)

	err = GenerateAndRunCmd(F_CMD_TYPES)
	assert.NoError(t, err)

	bytes, err := ioutil.ReadFile(filepath.Join(pkgDir, "types", "generated.go"))
	assert.NoError(t, err)
	source := string(bytes)
	assert.Contains(t, source, "system.RegisterType")

}

func TestGenerateFiles(t *testing.T) {

	currentDir, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(currentDir)

	dir, err := ioutil.TempDir(currentDir, "temporary")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	currentGopath := os.Getenv("GOPATH")
	defer os.Setenv("GOPATH", currentGopath)
	os.Setenv("GOPATH", dir)

	pkgDir := filepath.Join(dir, "src", "e.f", "g")
	os.MkdirAll(pkgDir, 0777)

	os.Chdir("/")

	err = GenerateFiles(F_MAIN)
	assert.IsError(t, err, "MHLOQTSAOX")

	os.Chdir(pkgDir)

	err = GenerateFiles(F_MAIN)
	assert.IsError(t, err, "XFNESBLBTQ")
	assert.HasError(t, err, "HQLAEMCHBM")

	data := `{"type": "system:type", "id": "a"}`
	err = ioutil.WriteFile(filepath.Join(pkgDir, "a.json"), []byte(data), 0777)
	assert.NoError(t, err)

	data = `{"type": "c", "id": "b"}`
	err = ioutil.WriteFile(filepath.Join(pkgDir, "b.json"), []byte(data), 0777)
	assert.NoError(t, err)

	err = GenerateFiles(F_MAIN)
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(pkgDir, "generated.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "json.RegisterType")

	// This will error because of unknown types in b.json
	err = GenerateFiles(F_TYPES)
	assert.IsError(t, err, "XYIUHERDHE")
	assert.HasError(t, err, "KWNPDUJNYP")

	os.Remove(filepath.Join(pkgDir, "b.json"))

	err = GenerateFiles(F_TYPES)
	assert.NoError(t, err)

	bytes, err := ioutil.ReadFile(filepath.Join(pkgDir, "types", "generated.go"))
	assert.NoError(t, err)
	source := string(bytes)
	assert.Contains(t, source, "system.RegisterType")

}

func Test_parseOptions(t *testing.T) {

	currentDir, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(currentDir)

	dir, err := ioutil.TempDir(currentDir, "temporary")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	currentGopath := os.Getenv("GOPATH")
	defer os.Setenv("GOPATH", currentGopath)
	os.Setenv("GOPATH", dir)

	pkgDir := filepath.Join(dir, "src", "x.y", "z")
	os.MkdirAll(pkgDir, 0777)

	os.Chdir(pkgDir)

	test, dir, path, imports, err := parseOptions()
	assert.NoError(t, err)
	assert.Equal(t, pkgDir, dir)
	assert.Equal(t, false, test)
	assert.Equal(t, "x.y/z", path)
	assert.Equal(t, map[string]string{}, imports)

	defer func() { *generatorTestFlag = false }()
	*generatorTestFlag = true
	test, _, _, _, err = parseOptions()
	assert.NoError(t, err)
	assert.Equal(t, true, test)

	defer func() { *generatorPathFlag = "" }()
	*generatorPathFlag = "a.b/c"
	_, _, path, _, err = parseOptions()
	assert.NoError(t, err)
	assert.Equal(t, "a.b/c", path)

	*generatorPathFlag = ""
	os.Chdir("/")
	defer os.Chdir(currentDir)
	_, _, _, _, err = parseOptions()
	assert.IsError(t, err, "PSRAWHQCPV")
	assert.HasError(t, err, "CXOETFPTGM")
}

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
