package process

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"kego.io/kerr/assert"
)

func Validate_NeedsTypes(t *testing.T) {

	d, err := ioutil.TempDir("", "temporary")
	assert.NoError(t, err)
	defer os.Remove(d)

	aFile := filepath.Join(d, "a.json")
	aJson := `{
			"description": "a",
			"type": "system:type",
			"id": "a",
			"fields": {
				"a": {
					"type": "system:@string"
				}
			}
		}`
	err = ioutil.WriteFile(aFile, []byte(aJson), 0644)
	assert.NoError(t, err)
	defer os.Remove(aFile)

	err = Validate(settings{dir: d, path: "d.e/f"})
	assert.NoError(t, err)
}
func TestValidate_error1(t *testing.T) {

	d, err := ioutil.TempDir("", "temporary")
	assert.NoError(t, err)
	defer os.Remove(d)

	bFile := filepath.Join(d, "b.json")
	bJson := `{
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
	}`
	err = ioutil.WriteFile(bFile, []byte(bJson), 0644)
	assert.NoError(t, err)
	defer os.Remove(bFile)

	err = Validate(settings{dir: d, path: "d.e/f"})
	// @string is invalid because minLength > maxLength
	assert.HasError(t, err, "DCIARXKRXN")

}
