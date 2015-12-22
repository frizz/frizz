package scan

import (
	"testing"

	"kego.io/kerr/assert"
	"kego.io/system"

	"os"

	"kego.io/process/tests"
)

func TestScan(t *testing.T) {

	n, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	path, dir, _, err := tests.CreateTemporaryPackage(n, "a", map[string]string{
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

	err = ScanForTypes(tests.AllCtx(tests.Ctx{Dir: dir, Path: path}), false)
	assert.NoError(t, err)
	defer system.Unregister(path, "b")

	h, ok := system.GetGlobal(path, "b")
	assert.True(t, ok)
	ty, ok := h.Object.(*system.Type)
	assert.True(t, ok)
	assert.Equal(t, "a", ty.Description)

}

func TestScanRule(t *testing.T) {

	n, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	path, dir, _, err := tests.CreateTemporaryPackage(n, "a", map[string]string{
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
				"embed": ["system:rule"],
				"fields": {
					"e": {
						"description": "f",
						"type": "system:@string"
					}
				}
			}
		}`,
	})

	err = ScanForTypes(tests.AllCtx(tests.Ctx{Dir: dir, Path: path}), false)
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

	err := ScanForTypes(tests.AllCtx(tests.Ctx{Dir: "/this-folder-doesnt-exist", Recursive: true}), false)
	assert.IsError(t, err, "XHHQSAVCKK")

	err = ScanForTypes(tests.AllCtx(tests.Ctx{Dir: "/this-folder-doesnt-exist", Recursive: false}), false)
	assert.IsError(t, err, "CDYLDBLHKT")

	n, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	path, dir, _, err := tests.CreateTemporaryPackage(n, "a", map[string]string{
		"a.json": "foo",
	})

	err = ScanForTypes(tests.AllCtx(tests.Ctx{Dir: dir, Path: path, Recursive: true}), false)
	assert.IsError(t, err, "XHHQSAVCKK")
	assert.HasError(t, err, "DHTURNTIXE")

	err = ScanForTypes(tests.AllCtx(tests.Ctx{Dir: dir, Path: path, Recursive: false}), false)
	assert.IsError(t, err, "IAPRUHFTAD")
	assert.HasError(t, err, "DHTURNTIXE")

}
