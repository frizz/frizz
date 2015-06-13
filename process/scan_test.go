package process

import (
	"io/ioutil"
	"testing"

	"kego.io/assert"
	"kego.io/system"

	"os"

	"path/filepath"
)

func TestScan(t *testing.T) {

	test := `{
	"description": "a",
	"type": "system:type",
	"id": "b",
	"properties": {
		"c": {
			"type": "system:property",
			"item": {
				"type": "system:@string"
			}
		}
	}
}`

	d, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.RemoveAll(d)

	f := filepath.Join(d, "a.json")
	err = ioutil.WriteFile(f, []byte(test), 0644)
	assert.NoError(t, err)

	err = Scan(d, false, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	ty, ok := system.GetType("d.e/f:b")
	assert.True(t, ok)
	assert.Equal(t, "a", ty.Description)

}

func TestScan_rule(t *testing.T) {

	test := `{
	"description": "a",
	"type": "system:type",
	"id": "b",
	"properties": {
		"c": {
			"type": "system:property",
			"item": {
				"type": "system:@string"
			}
		}
	},
	"rule": {
		"description": "d",
		"type": "system:type",
		"id": "@b",
		"is": ["system:rule"],
		"properties": {
			"e": {
				"description": "f",
				"type": "system:property",
				"item": {
					"type": "system:@string"
				}
			}
		}
	}
}`

	d, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	defer os.Remove(d)

	f := filepath.Join(d, "a.json")
	err = ioutil.WriteFile(f, []byte(test), 0644)
	assert.NoError(t, err)
	defer os.Remove(f)

	err = Scan(d, false, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	ty, ok := system.GetType("d.e/f:b")
	assert.True(t, ok)
	assert.Equal(t, "@b", ty.Rule.Id)

}

func TestScan_errors(t *testing.T) {

	err := Scan("/this-folder-doesnt-exist", false, "", map[string]string{})
	assert.IsError(t, err, "XHHQSAVCKK")

	d, err := ioutil.TempDir("", "")
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

func TestProcessScannedFile(t *testing.T) {
	err := processScannedFile("/this-file-doesnt-exist.json", false, "", map[string]string{})
	assert.IsError(t, err, "NMWROTKPLJ")
}
