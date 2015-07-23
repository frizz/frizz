package process

import (
	"os"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/internal/pkgtest"
	"kego.io/system"
)

func Validate_NeedsTypes(t *testing.T) {

	n, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.Remove(n)

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
	path, dir, err := pkgtest.CreateTemporaryPackage(n, "a", files)
	assert.NoError(t, err)

	// this is a type, so we need to register it with a hash to stop validate erroring.
	hash, err := getHash("a.json", path, map[string]string{}, []byte(files["a.json"]))
	assert.NoError(t, err)
	system.RegisterType(path, "a", &system.Type{}, hash)
	defer system.UnregisterType(path, "a")

	err = Validate(settings{dir: dir, path: path})
	assert.NoError(t, err)

}
func TestValidate_error1(t *testing.T) {

	n, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.Remove(n)

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
	path, dir, err := pkgtest.CreateTemporaryPackage(n, "b", files)
	assert.NoError(t, err)

	// this is a type, so we need to register it with a hash to stop validate erroring.
	hash, err := getHash("b.json", path, map[string]string{}, []byte(files["b.json"]))
	assert.NoError(t, err)
	system.RegisterType(path, "b", &system.Type{}, hash)
	defer system.UnregisterType(path, "b")

	err = Validate(settings{dir: dir, path: path})
	// @string is invalid because minLength > maxLength
	assert.HasError(t, err, "YLONAMFUAG")

}
