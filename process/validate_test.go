package process

import (
	"os"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/pkgtest"
	"kego.io/process/settings"
	"kego.io/system"
)

func Validate_NeedsTypes(t *testing.T) {

	n, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	files := map[string]string{
		"a.json": `{
			"description": "a",
			"type": "system:type",
			"id": "a",
			"fields": {
				"a": {
					"type": "system:@string"
				}
			}
		}`,
	}
	path, dir, _, err := pkgtest.CreateTemporaryPackage(n, "a", files)
	assert.NoError(t, err)

	// this is a type, so we need to register it with a hash to stop validate erroring.
	hash, err := getHash("a.json", path, map[string]string{}, []byte(files["a.json"]))
	assert.NoError(t, err)
	system.Register(path, "a", &system.Type{}, hash)
	defer system.Unregister(path, "a")

	err = Validate(&settings.Settings{Dir: dir, Path: path})
	assert.NoError(t, err)

}
func TestValidate_error1(t *testing.T) {

	n, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	files := map[string]string{
		"b.json": `{
			"description": "b",
			"type": "system:type",
			"id": "b",
			"fields": {
				"b": {
					"type": "system:@string",
					"minLength": 10,
					"maxLength": 5
				}
			}
		}`,
	}
	path, dir, _, err := pkgtest.CreateTemporaryPackage(n, "b", files)
	assert.NoError(t, err)

	// this is a type, so we need to register it with a hash to stop validate erroring.
	hash, err := getHash("b.json", path, map[string]string{}, []byte(files["b.json"]))
	assert.NoError(t, err)
	system.Register(path, "b", &system.Type{}, hash)
	defer system.Unregister(path, "b")

	err = Validate(&settings.Settings{Dir: dir, Path: path})
	// @string is invalid because minLength > maxLength
	assert.HasError(t, err, "YLONAMFUAG")

}
