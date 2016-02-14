package command

import (
	"testing"

	"kego.io/context/sysctx"
	"kego.io/kerr/assert"
	"kego.io/process/parser"
	"kego.io/process/tests"
	_ "kego.io/system"
)

func TestComparePackageHash(t *testing.T) {
	cb := tests.New().TempGopath(true)
	path, _ := cb.TempPackage("a", map[string]string{
		"a.yaml": "type: system:package",
	})
	cb.Path(path).Jempty()

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

}
