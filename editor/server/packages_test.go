package server

import (
	"testing"

	"io/ioutil"
	"path/filepath"

	"github.com/davelondon/ktest/assert"
	"kego.io/tests"
)

func TestCreateTemporaryPackage(t *testing.T) {
	cb := tests.New().TempGopath(false)
	defer cb.Cleanup()

	path, packageDir, err := CreateTemporaryPackage(cb.Ctx(), "a", "b", map[string]string{"c": "d", "e": "f"})

	assert.Regexp(t, `a/js\d{9}/b`, path)
	c, err := ioutil.ReadFile(filepath.Join(packageDir, "c"))
	assert.Nil(t, err)
	assert.Equal(t, []byte("d"), c)
	e, err := ioutil.ReadFile(filepath.Join(packageDir, "e"))
	assert.Nil(t, err)
	assert.Equal(t, []byte("f"), e)

}
