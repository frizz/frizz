package process_test

import (
	"io/ioutil"
	"os"
	"testing"

	"path/filepath"

	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/settings"
	"kego.io/process/tests"
	_ "kego.io/system/types"
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

	set := &settings.Settings{Dir: dir, Path: path}
	err = process.Generate(process.S_STRUCTS, set)
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated-structs.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "json.Register")

	// This will error because of unknown types in b.json
	err = process.Generate(process.S_TYPES, set)
	assert.IsError(t, err, "XYIUHERDHE")
	assert.HasError(t, err, "FKCPTUWJWW")

	os.Remove(filepath.Join(dir, "b.json"))

	err = process.Generate(process.S_TYPES, set)
	assert.NoError(t, err)

	bytes, err := ioutil.ReadFile(filepath.Join(dir, "types", "generated-types.go"))
	assert.NoError(t, err)
	source := string(bytes)
	assert.Contains(t, source, "system.Register")

}

func TestGenerate_path(t *testing.T) {

	namespace, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	path, dir, _, err := tests.CreateTemporaryPackage(namespace, "z", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
	})

	err = process.Generate(process.S_STRUCTS, &settings.Settings{Dir: dir, Path: path})
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated-structs.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "package z\n")

}
