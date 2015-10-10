package process

import (
	"testing"

	"kego.io/kerr/assert"
	"kego.io/system"

	"os"

	"kego.io/process/pkgtest"
	"kego.io/process/settings"
)

func TestScan(t *testing.T) {

	n, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	path, dir, _, err := pkgtest.CreateTemporaryPackage(n, "a", map[string]string{
		"a.json": `{
			"description": "a",
			"type": "system:type",
			"id": "b",
			"fields": {
				"c": {
					"type": "system:@string"
				}
			}
		}`,
	})

	err = ScanForTypes(false, &settings.Settings{Dir: dir, Path: path})
	assert.NoError(t, err)
	defer system.Unregister(path, "b")

	h, ok := system.GetGlobal(path, "b")
	assert.True(t, ok)
	ty, ok := h.Object.(*system.Type)
	assert.True(t, ok)
	assert.Equal(t, "a", ty.Description)

}

func TestScan_rule(t *testing.T) {

	n, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	path, dir, _, err := pkgtest.CreateTemporaryPackage(n, "a", map[string]string{
		"a.json": `{
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
				"is": ["system:rule"],
				"fields": {
					"e": {
						"description": "f",
						"type": "system:@string"
					}
				}
			}
		}`,
	})

	err = ScanForTypes(false, &settings.Settings{Dir: dir, Path: path})
	assert.NoError(t, err)
	defer system.Unregister(path, "b")
	defer system.Unregister(path, "@b")

	h, ok := system.GetGlobal(path, "b")
	assert.True(t, ok)
	ty, ok := h.Object.(*system.Type)
	assert.True(t, ok)
	assert.Equal(t, "@b", ty.Rule.Id.Name)

}

func TestScan_errors(t *testing.T) {

	err := ScanForTypes(false, &settings.Settings{Dir: "/this-folder-doesnt-exist", Recursive: true})
	assert.IsError(t, err, "XHHQSAVCKK")

	err = ScanForTypes(false, &settings.Settings{Dir: "/this-folder-doesnt-exist", Recursive: false})
	assert.IsError(t, err, "CDYLDBLHKT")

	n, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	path, dir, _, err := pkgtest.CreateTemporaryPackage(n, "a", map[string]string{
		"a.json": "foo",
	})

	err = ScanForTypes(false, &settings.Settings{Dir: dir, Path: path, Recursive: true})
	assert.IsError(t, err, "XHHQSAVCKK")
	assert.HasError(t, err, "DHTURNTIXE")

	err = ScanForTypes(false, &settings.Settings{Dir: dir, Path: path, Recursive: false})
	assert.IsError(t, err, "IAPRUHFTAD")
	assert.HasError(t, err, "DHTURNTIXE")

}
