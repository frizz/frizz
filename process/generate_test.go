package process_test

import (
	"io/ioutil"
	"testing"

	"kego.io/process/parser"

	"path/filepath"

	"fmt"

	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/tests"
)

func TestGenerateAll(t *testing.T) {

	cb := tests.Context("a.b/c").TempGopath(true)
	defer cb.Cleanup()

	path, dir, tests := cb.TempPackage("foo", map[string]string{
		"foo.go": "package foo",
	})
	fmt.Println(path, dir, tests)

}

func TestGenerate(t *testing.T) {

	cb := tests.NewContextBuilder().TempGopath(false)
	defer cb.Cleanup()

	path, dir, _ := cb.TempPackage("d", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
		"b.json": `{"type": "c", "id": "b"}`,
		"d.go":   `package d`,
	})

	cb.Path(path).Dir(dir).Cmd().Wg().Jsystem().Sauto(parser.Parse)

	err := process.Generate(cb.Ctx(), cb.Env())
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "jsonctx.InitPackage")

}

func TestGenerate_path(t *testing.T) {

	cb := tests.NewContextBuilder().TempGopath(false)
	defer cb.Cleanup()

	path, dir, _ := cb.TempPackage("z", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
	})

	cb.Path(path).Dir(dir).Cmd().Wg().Jsystem().Sauto(parser.Parse)

	err := process.Generate(cb.Ctx(), cb.Env())
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "package z\n")

}
