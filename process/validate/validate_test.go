package validate

import (
	"os"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/parse"
	"kego.io/process/tests"
)

func TestValidate_NeedsTypes(t *testing.T) {

	n, err := tests.CreateTemporaryNamespace()
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
	path, dir, _, err := tests.CreateTemporaryPackage(n, "a", files)
	assert.NoError(t, err)

	cb := tests.Context(path).Dir(dir).Jsystem().Sauto(parse.Parse)

	err = ValidatePackage(cb.Ctx())
	assert.NoError(t, err)

}

func TestValidate_error1(t *testing.T) {

	n, err := tests.CreateTemporaryNamespace()
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
	path, dir, _, err := tests.CreateTemporaryPackage(n, "b", files)
	assert.NoError(t, err)

	cb := tests.Context(path).Dir(dir).Jsystem().Sauto(parse.Parse)

	err = ValidatePackage(cb.Ctx())
	// @string is invalid because minLength > maxLength
	assert.HasError(t, err, "YLONAMFUAG")

}
