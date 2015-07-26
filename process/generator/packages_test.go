package generator

import (
	"os"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/internal/pkgtest"
)

func TestGuessPackageName(t *testing.T) {
	a := guessPackageName("a")
	assert.Equal(t, "a", a)
	a = guessPackageName("a/b")
	assert.Equal(t, "b", a)
	a = guessPackageName("a/b/c")
	assert.Equal(t, "c", a)
	a = guessPackageName("a/b/c-d")
	assert.Equal(t, "d", a)
	a = guessPackageName("a/b/c-d/")
	assert.Equal(t, "d", a)
	a = guessPackageName("a.b")
	assert.Equal(t, "a", a)
	a = guessPackageName("a/b.c")
	assert.Equal(t, "b", a)
	a = guessPackageName("a/b-c.d")
	assert.Equal(t, "c", a)
}

func TestGetPackageName(t *testing.T) {

	dir, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(dir)

	path, _, _, err := pkgtest.CreateTemporaryPackage(dir, "a", map[string]string{})
	assert.NoError(t, err)
	name := getPackageName(path)
	assert.Equal(t, "a", name)

	path, _, _, err = pkgtest.CreateTemporaryPackage(dir, "b", map[string]string{
		"foo.go": "package c",
	})
	assert.NoError(t, err)
	name = getPackageName(path)
	assert.Equal(t, "c", name)

}
