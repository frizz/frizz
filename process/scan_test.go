package process

import (
	"testing"

	"kego.io/kerr/assert"
	"kego.io/system"

	"os"

	"kego.io/process/pkgtest"
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

	err = ScanForTypes(false, settings{dir: dir, path: path})
	assert.NoError(t, err)
	defer system.UnregisterType(path, "b")

	ty, _, ok := system.GetType(path, "b")
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
		}`,
	})

	err = ScanForTypes(false, settings{dir: dir, path: path})
	assert.NoError(t, err)
	defer system.UnregisterType(path, "b")
	defer system.UnregisterType(path, "@b")

	ty, _, ok := system.GetType(path, "b")
	assert.True(t, ok)
	assert.Equal(t, "@b", ty.Rule.Id.Name)

}

func TestScan_errors(t *testing.T) {

	err := ScanForTypes(false, settings{dir: "/this-folder-doesnt-exist", recursive: true})
	assert.IsError(t, err, "XHHQSAVCKK")

	err = ScanForTypes(false, settings{dir: "/this-folder-doesnt-exist", recursive: false})
	assert.IsError(t, err, "CDYLDBLHKT")

	n, err := pkgtest.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	path, dir, _, err := pkgtest.CreateTemporaryPackage(n, "a", map[string]string{
		"a.json": "foo",
	})

	err = ScanForTypes(false, settings{dir: dir, path: path, recursive: true})
	assert.IsError(t, err, "XHHQSAVCKK")
	assert.HasError(t, err, "DHTURNTIXE")

	err = ScanForTypes(false, settings{dir: dir, path: path, recursive: false})
	assert.IsError(t, err, "IAPRUHFTAD")
	assert.HasError(t, err, "DHTURNTIXE")

}
