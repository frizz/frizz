package process_test

import (
	"io/ioutil"
	"os"
	"testing"

	"path/filepath"

	"kego.io/kerr/assert"
	"kego.io/parse"
	"kego.io/process"
	"kego.io/process/tests"
)

func TestGenerate(t *testing.T) {

	namespace, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	path, dir, _, err := tests.CreateTemporaryPackage(namespace, "d", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
		"b.json": `{"type": "c", "id": "b"}`,
		"d.go":   `package d`,
	})
	assert.NoError(t, err)

	cb := tests.Context(path).Dir(dir).Cmd().Wg().Jsystem().Sauto(parse.Parse)

	err = process.Generate(cb.Ctx(), cb.Env())
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "jsonctx.InitPackage")

}

func TestGenerate_path(t *testing.T) {

	namespace, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	path, dir, _, err := tests.CreateTemporaryPackage(namespace, "z", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
	})

	cb := tests.Context(path).Dir(dir).Cmd().Wg().Jsystem().Sauto(parse.Parse)

	err = process.Generate(cb.Ctx(), cb.Env())
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "package z\n")

}
