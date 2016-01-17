package process

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"kego.io/process/parser"

	"fmt"

	"kego.io/context/envctx"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestRun(t *testing.T) {

	namespace, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(namespace)

	path, dir, _, err := tests.CreateTemporaryPackage(namespace, "d", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
		"d.go":   `package d`,
	})

	cb := tests.Context(path).Dir(dir).Jauto().Wg().Sauto(parser.Parse)

	env := envctx.FromContext(cb.Ctx())

	err = Generate(cb.Ctx(), env)
	assert.NoError(t, err)

	b, err := ioutil.ReadFile(filepath.Join(dir, "generated.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(b), "pkg.InitType(\"a\", reflect.TypeOf((*A)(nil)), reflect.TypeOf((*ARule)(nil)), reflect.TypeOf((*AInterface)(nil)).Elem())")
	assert.Contains(t, string(b), fmt.Sprintf("%v", env.Hash))

	err = BuildAndRunLocalCommand(cb.Ctx())
	assert.NoError(t, err)

	_, err = os.Stat(filepath.Join(dir, ".localke", "validate"))
	assert.NoError(t, err)

}
