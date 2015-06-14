package process

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"kego.io/assert"
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

	err = Validate(d, "d.e/f", map[string]string{})
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

	err = Validate(d, "d.e/f", map[string]string{})
	// @string is invalid because minLength > maxLength
	assert.HasError(t, err, "DCIARXKRXN")

}

func TestValidate_error2(t *testing.T) {

	err := Scan("/this-folder-doesnt-exist", false, "", map[string]string{})
	assert.IsError(t, err, "XHHQSAVCKK")

	d, err := ioutil.TempDir("", "temporary")
	assert.NoError(t, err)
	defer os.Remove(d)

	f := filepath.Join(d, "a.json")
	err = ioutil.WriteFile(f, []byte("foo"), 0644)
	assert.NoError(t, err)
	defer os.Remove(f)

	err = Scan(d, false, "d.e/f", map[string]string{})
	assert.IsError(t, err, "XHHQSAVCKK")
	assert.HasError(t, err, "DHTURNTIXE")

}

func TestValidateReader(t *testing.T) {
	err := processScannedFile("/this-file-doesnt-exist.json", false, "", map[string]string{})
	assert.IsError(t, err, "NMWROTKPLJ")
}
