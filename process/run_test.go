package process

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"kego.io/process/parser"

	"fmt"

	"github.com/davelondon/kerr"
	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/context/envctx"
	"kego.io/tests"
)

func TestRun(t *testing.T) {

	cb := tests.New().RealGopath()
	defer cb.Cleanup()

	path, dir := cb.TempPackage("d", map[string]string{
		"a.yaml": `
			type: system:type
			id: a
			fields:
				b:
					type: system:@string
					max-length: 5`,
		"d.go": `package d`,
	})

	cb.Path(path).Dir(dir).Jauto().Wg().Sauto(parser.Parse)

	env := envctx.FromContext(cb.Ctx())

	err := Generate(cb.Ctx(), env)
	require.NoError(t, err)

	b, err := ioutil.ReadFile(filepath.Join(dir, "generated.go"))
	require.NoError(t, err)
	assert.Contains(t, string(b), `pkg.Init("a", func() interface{} { return new(A) }, func() interface{} { return new(ARule) }, func() reflect.Type { return reflect.TypeOf((*AInterface)(nil)).Elem() })`)
	assert.Contains(t, string(b), fmt.Sprintf("%v", env.Hash))

	err = RunValidateCommand(cb.Ctx())
	require.NoError(t, err)

	file1, err := os.Stat(filepath.Join(dir, ".localke", "validate"))
	require.NoError(t, err)
	time1 := file1.ModTime()

	err = RunValidateCommand(cb.Ctx())
	require.NoError(t, err)

	cb.TempFile("c.yaml", `
		type: a
		id: c
		b: foo`)

	err = RunValidateCommand(cb.Ctx())
	require.NoError(t, err)

	// should not rebuild validate command
	file2, err := os.Stat(filepath.Join(dir, ".localke", "validate"))
	require.NoError(t, err)
	time2 := file2.ModTime()
	assert.Equal(t, time1, time2)

	cb.TempFile("e.yaml", `
		type: a
		id: e
		b: tooolong`)

	err = RunValidateCommand(cb.Ctx())
	assert.IsError(t, err, "KFNIOHWCBT")
	assert.HasError(t, err, "ETWHPXTUVB")

	cb.TempFile("f.yaml", `
		type: system:type
		id: f
		fields:
			a:
				type: system:@string`)

	// This loads the new system.Type into the system cache
	cb.Sauto(parser.Parse)
	// This generates a new generated.go
	err = Generate(cb.Ctx(), env)
	require.NoError(t, err)

	// This will re-run the build, but still return the validation error
	err = RunValidateCommand(cb.Ctx())
	assert.IsError(t, err, "KFNIOHWCBT")
	assert.HasError(t, err, "ETWHPXTUVB")

	cb.RemoveTempFile("e.yaml")

	cb.TempFile("h.yaml", `
		type: system:type
		id: h
		fields:
			a:
				type: system:@string`)

	// This loads the new system.Type into the system cache
	cb.Sauto(parser.Parse)
	// This generates a new generated.go
	err = Generate(cb.Ctx(), env)
	require.NoError(t, err)

	// This will re-run the build, but not return the validation error
	err = RunValidateCommand(cb.Ctx())
	require.NoError(t, err)

	// should rebuild validate command
	file3, err := os.Stat(filepath.Join(dir, ".localke", "validate"))
	require.NoError(t, err)
	time3 := file3.ModTime()
	assert.NotEqual(t, time1, time3)

	cb.TempFile("g.yaml", `
		type: system:type
		id: g
		fields:
			a:
				type: system:@string`)

	// We add a new type, but we haven't generated the struct, so it will fail with hash changed
	err = runValidateCommand(cb.Ctx(), false, false)
	assert.IsError(t, err, "DTTHRRJSSF")

	err = os.Remove(filepath.Join(dir, ".localke", "validate"))
	require.NoError(t, err)

	_, err = os.Stat(filepath.Join(dir, ".localke", "validate"))
	assert.True(t, os.IsNotExist(err))

	err = runValidateCommand(cb.Ctx(), false, false)
	assert.IsError(t, err, "DTTHRRJSSF")
	_, ok := kerr.Source(err).(*os.PathError)
	assert.True(t, ok)

}
