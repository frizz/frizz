package process

import (
	"io/ioutil"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/system"

	"os"

	"path/filepath"
)

func TestScan(t *testing.T) {

	test := `{
	"description": "a",
	"type": "system:type",
	"id": "b",
	"fields": {
		"c": {
			"type": "system:@string"
		}
	}
}`

	d, err := ioutil.TempDir("", "temporary")
	assert.NoError(t, err)
	defer os.RemoveAll(d)

	f := filepath.Join(d, "a.json")
	err = ioutil.WriteFile(f, []byte(test), 0644)
	assert.NoError(t, err)

	err = ScanForTypes(d, false, false, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	ty, ok := system.GetType("d.e/f", "b")
	assert.True(t, ok)
	assert.Equal(t, "a", ty.Description)

}

func TestScan_rule(t *testing.T) {

	test := `{
	"description": "a",
	"type": "system:type",
	"id": "b",
	"fields": {
		"c": {
			"type": "system:@string"
		}
	},
	"rule": {
		"description": "d",
		"type": "system:type",
		"id": "@b",
		"is": ["system:rule"],
		"embed": ["system:ruleBase"],
		"fields": {
			"e": {
				"description": "f",
				"type": "system:@string"
			}
		}
	}
}`

	d, err := ioutil.TempDir("", "temporary")
	assert.NoError(t, err)
	defer os.Remove(d)

	f := filepath.Join(d, "a.json")
	err = ioutil.WriteFile(f, []byte(test), 0644)
	assert.NoError(t, err)
	defer os.Remove(f)

	err = ScanForTypes(d, false, false, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	ty, ok := system.GetType("d.e/f", "b")
	assert.True(t, ok)
	assert.Equal(t, "@b", ty.Rule.Id.Name)

}

func TestScan_errors(t *testing.T) {

	err := ScanForTypes("/this-folder-doesnt-exist", false, true, "", map[string]string{})
	assert.IsError(t, err, "XHHQSAVCKK")

	err = ScanForTypes("/this-folder-doesnt-exist", false, false, "", map[string]string{})
	assert.IsError(t, err, "CDYLDBLHKT")

	d, err := ioutil.TempDir("", "temporary")
	assert.NoError(t, err)
	defer os.Remove(d)

	f := filepath.Join(d, "a.json")
	err = ioutil.WriteFile(f, []byte("foo"), 0644)
	assert.NoError(t, err)
	defer os.Remove(f)

	err = ScanForTypes(d, false, true, "d.e/f", map[string]string{})
	assert.IsError(t, err, "XHHQSAVCKK")
	assert.HasError(t, err, "DHTURNTIXE")

	err = ScanForTypes(d, false, false, "d.e/f", map[string]string{})
	assert.IsError(t, err, "IAPRUHFTAD")
	assert.HasError(t, err, "DHTURNTIXE")

}

//func TestProcessScannedFile(t *testing.T) {
//	err := processScannedFile("/this-file-doesnt-exist.json", false, "", map[string]string{})
//	assert.IsError(t, err, "NMWROTKPLJ")
//}
