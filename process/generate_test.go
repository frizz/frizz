package process

import (
	"io/ioutil"
	"testing"

	"kego.io/process/parser"

	"path/filepath"

	"os"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/tests"
)

func TestSave(t *testing.T) {
	d, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(d)
	err = save(d, []byte{}, "b", false)
	require.NoError(t, err)
	if _, err := os.Stat(filepath.Join(d, "b")); err == nil {
		assert.Fail(t, "File should not exist")
	}

	// backup is on, but the file doesn't exist, so nothing to backup
	err = save(d, []byte("a"), "c", true)
	require.NoError(t, err)
	val, err := ioutil.ReadFile(filepath.Join(d, "c"))
	require.NoError(t, err)
	assert.Equal(t, "a", string(val))
	if _, err := os.Stat(filepath.Join(d, "c.backup")); err == nil {
		assert.Fail(t, "Backup file should not exist")
	}

	// backup is off, so the file is overwritten
	err = save(d, []byte("b"), "c", false)
	require.NoError(t, err)
	val, err = ioutil.ReadFile(filepath.Join(d, "c"))
	require.NoError(t, err)
	assert.Equal(t, "b", string(val))
	_, err = os.Stat(filepath.Join(d, "c.backup"))
	assert.Error(t, err)

	// backup is on and the file exists, so we should backup
	err = save(d, []byte("c"), "c", true)
	require.NoError(t, err)
	val, err = ioutil.ReadFile(filepath.Join(d, "c"))
	require.NoError(t, err)
	assert.Equal(t, "c", string(val))
	back, err := ioutil.ReadFile(filepath.Join(d, "c.backup"))
	require.NoError(t, err)
	assert.Equal(t, "b", string(back))

	// backup is on and the file exists, so we should backup and update the backup
	err = save(d, []byte("d"), "c", true)
	require.NoError(t, err)
	val, err = ioutil.ReadFile(filepath.Join(d, "c"))
	require.NoError(t, err)
	assert.Equal(t, "d", string(val))
	back, err = ioutil.ReadFile(filepath.Join(d, "c.backup"))
	require.NoError(t, err)
	assert.Equal(t, "c", string(back))

	// dir doesn't exist
	err = save(filepath.Join(d, "e"), []byte("f"), "g", true)
	require.NoError(t, err)
	val, err = ioutil.ReadFile(filepath.Join(d, "e", "g"))
	require.NoError(t, err)
	assert.Equal(t, "f", string(val))
}

func TestGetInfo(t *testing.T) {
	cb := tests.Context("a.b/c").Wg().Sempty().Jsystem().TempGopath(false)
	defer cb.Cleanup()
	_, dirA := cb.TempPackage("a", map[string]string{
		"generated.go": "",
	})
	info, found, err := getInfo(cb.Ctx(), dirA)
	require.NoError(t, err)
	assert.Nil(t, info)
	assert.False(t, found)

}

func TestGenerateAll(t *testing.T) {

	cb := tests.Context("a.b/c").Wg().Sempty().Jsystem().TempGopath(true)
	defer cb.Cleanup()

	pathA, dirA := cb.TempPackage("a", map[string]string{
		"a.yml": `
id: a
type: system:type
fields:
	a:
		type: system:@string`,
	})

	err := GenerateAll(cb.Ctx(), pathA, map[string]bool{})
	assert.IsError(t, err, "XMVXECGDOX")

	cb.Path(pathA).Sauto(parser.Parse)

	done := map[string]bool{pathA: true}
	err = GenerateAll(cb.Ctx(), pathA, done)
	require.NoError(t, err)
	assert.Equal(t, 1, len(done))

	err = GenerateAll(cb.Ctx(), pathA, map[string]bool{})
	require.NoError(t, err)

	b, err := ioutil.ReadFile(filepath.Join(dirA, "generated.go"))
	require.NoError(t, err)
	assert.Contains(t, string(b), "type A struct")

	pathB, _ := cb.TempPackage("b", map[string]string{
		"pkg.yml": `
type: system:package
aliases: {"a": "` + pathA + `"}`,
		"b.yml": `
id: b
type: system:type
fields:
	a:
		type: system:@string`,
	})

	cb.Path(pathB).Sauto(parser.Parse)
	err = GenerateAll(cb.Ctx(), pathB, map[string]bool{})
	require.NoError(t, err)

}

func TestGenerate(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	path, dir := cb.TempPackage("d", map[string]string{
		"a.json": `{"type": "system:type", "id": "a", "fields": {"a": {"type": "system:@string"}}}`,
		"b.json": `{"type": "c", "id": "b"}`,
		"d.go":   `package d`,
	})

	cb.Path(path).Dir(dir).Cmd().Wg().Jsystem().Sauto(parser.Parse)

	err := Generate(cb.Ctx(), cb.Env())
	require.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated.go"))
	require.NoError(t, err)
	assert.Contains(t, string(genBytes), "jsonctx.InitPackage")

}

func TestGenerate_path(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	path, dir := cb.TempPackage("z", map[string]string{
		"a.json": `{"type": "system:type", "id": "a", "fields": {"a": {"type": "system:@string"}}}`,
	})

	cb.Path(path).Dir(dir).Cmd().Wg().Jsystem().Sauto(parser.Parse)

	err := Generate(cb.Ctx(), cb.Env())
	require.NoError(t, err)

	genBytes, err := ioutil.ReadFile(filepath.Join(dir, "generated.go"))
	require.NoError(t, err)
	assert.Contains(t, string(genBytes), "package z\n")

}
