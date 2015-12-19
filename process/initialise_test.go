package process

import (
	"os"
	"strings"
	"testing"

	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestInitialise(t *testing.T) {

	current, err := os.Getwd()
	assert.NoError(t, err)
	defer os.Chdir(current)

	namespace, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	pathA, dirA, _, err := tests.CreateTemporaryPackage(namespace, "a", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
		"a.go":   "package a",
	})
	pathB, dirB, _, err := tests.CreateTemporaryPackage(namespace, "b", map[string]string{
		"b.json": `{"type": "system:type", "id": "b"}`,
		"b.go":   "package b",
	})

	err = os.Chdir(dirA)
	assert.NoError(t, err)

	ctx, _, err := Initialise(nil)
	assert.NoError(t, err)
	env, ok := envctx.FromContext(ctx)
	assert.True(t, ok)
	cmd, ok := cmdctx.FromContext(ctx)
	assert.True(t, ok)
	assert.Equal(t, dirA, cmd.Dir)
	assert.Equal(t, pathA, env.Path)

	err = os.Chdir("/")
	assert.NoError(t, err)

	ctx, _, err = Initialise(&FromDefaults{
		Path: pathB,
	})
	env, ok = envctx.FromContext(ctx)
	assert.True(t, ok)
	cmd, ok = cmdctx.FromContext(ctx)
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, dirB, cmd.Dir)
	assert.Equal(t, pathB, env.Path)

	_, _, err = Initialise(&FromDefaults{
		Path: "",
	})
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

	namespace, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	pathA, dirA, _, err := tests.CreateTemporaryPackage(namespace, "a", nil)

	dir, err := getPackageDir(pathA, os.Getenv("GOPATH"))
	assert.NoError(t, err)
	assert.Equal(t, dirA, dir)

	dir, err = getPackageDir("a.b/c", os.Getenv("GOPATH"))
	assert.IsError(t, err, "SUTCWEVRXS")

}
