package process

import (
	"os"
	"testing"

	"kego.io/context/envctx"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestInitialise(t *testing.T) {

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

	cb := tests.NewContextBuilder().OsWd(dirA)

	ctx, _, err := Initialise(cb.Ctx(), nil)
	assert.NoError(t, err)
	env := envctx.FromContext(ctx)
	assert.Equal(t, dirA, env.Dir)
	assert.Equal(t, pathA, env.Path)

	cb.OsWd("/")

	ctx, _, err = Initialise(ctx, &FromDefaults{
		Path: pathB,
	})
	env = envctx.FromContext(ctx)
	assert.NoError(t, err)
	assert.Equal(t, dirB, env.Dir)
	assert.Equal(t, pathB, env.Path)

	_, _, err = Initialise(ctx, &FromDefaults{
		Path: "",
	})
	assert.IsError(t, err, "ADNJKTLAWY")
	assert.HasError(t, err, "CXOETFPTGM")
}
