package builder

import (
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
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

	cb := tests.NewContextBuilder().RealGopath()
	defer cb.Cleanup()

	path, _ := cb.TempPackage("a", map[string]string{})
	name := getPackageName(path)
	assert.Equal(t, "a", name)

	path, _ = cb.TempPackage("b", map[string]string{
		"foo.go": "package c",
	})
	name = getPackageName(path)
	assert.Equal(t, "c", name)

}
