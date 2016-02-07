package process

import (
	"io/ioutil"
	"testing"

	"kego.io/process/parser"

	"path/filepath"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestGetInfo(t *testing.T) {
	cb := tests.Context("a.b/c").Wg().Sempty().Jsystem().TempGopath(false)
	defer cb.Cleanup()
	_, dirA := cb.TempPackage("a", map[string]string{
		"generated.go": "",
	})
	info, found, err := getInfo(cb.Ctx(), dirA)
	assert.NoError(t, err)
	assert.Nil(t, info)
	assert.False(t, found)

}

func TestGenerateAll(t *testing.T) {

	cb := tests.Context("a.b/c").Wg().Sempty().Jsystem().TempGopath(false)
	defer cb.Cleanup()

	pathA, dirA := cb.TempPackage("a", map[string]string{
		"a.yml": `
id: a
type: system:type`,
	})

	err := GenerateAll(cb.Ctx(), pathA, map[string]bool{})
	assert.IsError(t, err, "XMVXECGDOX")

	cb.Path(pathA).Sauto(parser.Parse)

	done := map[string]bool{pathA: true}
	err = GenerateAll(cb.Ctx(), pathA, done)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(done))

	err = GenerateAll(cb.Ctx(), pathA, map[string]bool{})
	assert.NoError(t, err)

	b, err := ioutil.ReadFile(filepath.Join(dirA, "generated.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(b), "type A struct")

	pathB, _ := cb.TempPackage("b", map[string]string{
		"pkg.yml": `
type: system:package
aliases: {"` + pathA + `": "a"}`,
		"b.yml": `
id: b
type: system:type`,
	})

	cb.Path(pathB).Sauto(parser.Parse)
	err = GenerateAll(cb.Ctx(), pathB, map[string]bool{})
	assert.NoError(t, err)

}

func TestGenerate(t *testing.T) {

	cb := tests.NewContextBuilder().TempGopath(false)
	defer cb.Cleanup()

	path, dir := cb.TempPackage("d", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
		"b.json": `{"type": "c", "id": "b"}`,
		"d.go":   `package d`,
	})

	cb.Path(path).Dir(dir).Cmd().Wg().Jsystem().Sauto(parser.Parse)

	err := Generate(cb.Ctx(), cb.Env())
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "jsonctx.InitPackage")

}

func TestGenerate_path(t *testing.T) {

	cb := tests.NewContextBuilder().TempGopath(false)
	defer cb.Cleanup()

	path, dir := cb.TempPackage("z", map[string]string{
		"a.json": `{"type": "system:type", "id": "a"}`,
	})

	cb.Path(path).Dir(dir).Cmd().Wg().Jsystem().Sauto(parser.Parse)

	err := Generate(cb.Ctx(), cb.Env())
	assert.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated.go"))
	assert.NoError(t, err)
	assert.Contains(t, string(genBytes), "package z\n")

}
