package process

import (
	"os"
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
	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)
	assert.Equal(t, dirA, cmd.Dir)
	assert.Equal(t, pathA, env.Path)

	err = os.Chdir("/")
	assert.NoError(t, err)

	ctx, _, err = Initialise(&FromDefaults{
		Path: pathB,
	})
	env = envctx.FromContext(ctx)
	cmd = cmdctx.FromContext(ctx)
	assert.NoError(t, err)
	assert.Equal(t, dirB, cmd.Dir)
	assert.Equal(t, pathB, env.Path)

	_, _, err = Initialise(&FromDefaults{
		Path: "",
	})
	assert.IsError(t, err, "PSRAWHQCPV")
	assert.HasError(t, err, "CXOETFPTGM")
}
