package process

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"kego.io/uerr"
)

func TestValidate(t *testing.T) {

	d, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.Remove(d)

	aFile := filepath.Join(d, "a.json")
	aJson := `{
			"description": "a",
			"type": "system:type",
			"id": "a",
			"properties": {
				"a": {
					"type": "system:property",
					"item": {
						"type": "system:@string"
					}
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

	d, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.Remove(d)

	bFile := filepath.Join(d, "b.json")
	bJson := `{
		"description": "b",
		"type": "system:type",
		"id": "b",
		"properties": {
			"b": {
				"type": "system:property",
				"item": {
					"type": "system:@string",
					"minLength": 10,
					"maxLength": 5
				}
			}
		}
	}`
	err = ioutil.WriteFile(bFile, []byte(bJson), 0644)
	assert.NoError(t, err)
	defer os.Remove(bFile)

	err = Validate(d, "d.e/f", map[string]string{})
	// @string is invalid because minLength > maxLength
	uerr.Stack(t, err, "DCIARXKRXN")

}

func TestValidate_error2(t *testing.T) {

	err := Scan("/this-folder-doesnt-exist", "", map[string]string{})
	uerr.Assert(t, err, "XHHQSAVCKK")

	d, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.Remove(d)

	f := filepath.Join(d, "a.json")
	err = ioutil.WriteFile(f, []byte("foo"), 0644)
	assert.NoError(t, err)
	defer os.Remove(f)

	err = Scan(d, "d.e/f", map[string]string{})
	uerr.Assert(t, err, "XHHQSAVCKK")
	uerr.Stack(t, err, "DHTURNTIXE")

}

func TestValidateReader(t *testing.T) {
	err := processScannedFile("/this-file-doesnt-exist.json", "", map[string]string{})
	uerr.Assert(t, err, "NMWROTKPLJ")
}
