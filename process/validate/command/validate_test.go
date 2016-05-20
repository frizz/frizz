package command

import (
	"os"
	"testing"

	"github.com/davelondon/ktest/assert"
	"kego.io/context/sysctx"
	"kego.io/process/parser"
	"kego.io/process/tests"
	_ "kego.io/process/validate/tests"
	_ "kego.io/system"
)

func TestValidateMain(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	dir := cb.SetTemp("kego.io/system")
	cb.Path("kego.io/system").Dir(dir).Jempty().Sempty()

	cancelled := false
	cancel := func() {
		cancelled = true
	}

	logged := ""
	log := func(message string) {
		logged = message
	}

	interrupt := make(chan os.Signal, 1)

	exitStatus := validateMain(cb.Ctx(), cancel, log, interrupt)
	assert.Equal(t, 1, exitStatus)
	assert.False(t, cancelled)
	assert.Contains(t, logged, "NHXWLPHCHL")
	logged = ""

	cb.Jauto().Sauto(parser.Parse)

	exitStatus = validateMain(cb.Ctx(), cancel, log, interrupt)
	assert.Equal(t, 0, exitStatus)
	assert.False(t, cancelled)
	assert.Equal(t, logged, "")

	cb.TempFile("a.yaml", `
		type: type
		id: a`)
	cb.Sauto(parser.Parse)

	exitStatus = validateMain(cb.Ctx(), cancel, log, interrupt)
	assert.Equal(t, 3, exitStatus)
	assert.False(t, cancelled)
	assert.Equal(t, logged, "")

	cb.RemoveTempFile("a.yaml").
		TempFile("a.yaml", `
			type: "@string"
			id: a
			maxLength: 1
			minLength: 2`).
		Sauto(parser.Parse)

	exitStatus = validateMain(cb.Ctx(), cancel, log, interrupt)
	assert.Equal(t, 4, exitStatus)
	assert.False(t, cancelled)
	assert.Equal(t, logged, "MaxLength 1 must not be less than MinLength 2")
	logged = ""

	cb.RemoveTempFile("a.yaml").Sauto(parser.Parse).TempFile("b.yaml", `%`)

	exitStatus = validateMain(cb.Ctx(), cancel, log, interrupt)
	assert.Equal(t, 1, exitStatus)
	assert.False(t, cancelled)
	assert.Contains(t, logged, "IHSVWAUAYW")

}

func TestComparePackageHash(t *testing.T) {
	cb := tests.New().TempGopath(true)
	path, _ := cb.TempPackage("a", map[string]string{
		"a.yaml": "type: system:package",
	})
	cb.Path(path).Jempty().Spkg(path)

	// "a.b/c" not found in scache.
	changes, err := comparePackageHash(cb.Ctx(), "a.b/c")
	assert.IsError(t, err, "NHXWLPHCHL")
	assert.False(t, changes)

	// path not found in jcache
	changes, err = comparePackageHash(cb.Ctx(), path)
	assert.NoError(t, err)
	assert.True(t, changes)

	cb.Jsystem().Jpkg(path, 999).Sauto(parser.Parse)

	// hash changed
	changes, err = comparePackageHash(cb.Ctx(), path)
	assert.NoError(t, err)
	assert.True(t, changes)

	scache := sysctx.FromContext(cb.Ctx())
	pi, _ := scache.Get(path)
	cb.Jpkg(path, pi.Hash)

	// hash correct
	changes, err = comparePackageHash(cb.Ctx(), path)
	assert.NoError(t, err)
	assert.False(t, changes)

	pi.Aliases["c"] = "a.b/c"

	changes, err = comparePackageHash(cb.Ctx(), path)
	assert.IsError(t, err, "DGJTLHQOCQ")
	assert.False(t, changes)

	pi1 := scache.Set("a.b/c")

	changes, err = comparePackageHash(cb.Ctx(), path)
	assert.NoError(t, err)
	assert.True(t, changes)

	cb.Jpkg("a.b/c", 1)
	pi1.Hash = 2

	changes, err = comparePackageHash(cb.Ctx(), path)
	assert.NoError(t, err)
	assert.True(t, changes)

	pi1.Hash = 1

	changes, err = comparePackageHash(cb.Ctx(), path)
	assert.NoError(t, err)
	assert.False(t, changes)

}
