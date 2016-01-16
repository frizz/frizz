package parse_test

import (
	"os"
	"testing"

	"kego.io/parse"

	"kego.io/context/sysctx"
	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/tests"
	"kego.io/system"
)

func TestScan_errors(t *testing.T) {

	cb := tests.Context("").Cmd().Sempty().Jsystem()

	_, err := parse.Parse(cb.Ctx(), "does-not-exist")
	assert.IsError(t, err, "GJRHNGGWFD")
	assert.HasError(t, err, "SUTCWEVRXS")

	n, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	path, _, _, err := tests.CreateTemporaryPackage(n, "a", map[string]string{
		"a.json": `{
			"type": "system:package",
			"recursive": false
		}`,
		"b.json": "foo",
	})
	assert.NoError(t, err)

	cb.Path(path)

	_, err = parse.Parse(cb.Ctx(), path)
	assert.IsError(t, err, "VFUNPHUFHD")
	assert.HasError(t, err, "HCYGNBDFFA")

	path, _, _, err = tests.CreateTemporaryPackage(n, "a", map[string]string{
		"a.json": "foo",
	})
	assert.NoError(t, err)

	cb.Path(path)

	_, err = parse.Parse(cb.Ctx(), path)
	assert.IsError(t, err, "GJRHNGGWFD")
	assert.HasError(t, err, "XTEQCAYQJP")

}

func TestParseRule(t *testing.T) {

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
	assert.NoError(t, err)

	cb := tests.Context(path).Dir(dir).Cmd().Sempty().Jsystem()

	_, err = parse.Parse(cb.Ctx(), path)
	assert.NoError(t, err)

	scache := sysctx.FromContext(cb.Ctx())
	i, ok := scache.GetType(path, "b")
	assert.True(t, ok)
	ty, ok := i.(*system.Type)
	assert.True(t, ok)
	assert.Equal(t, "@b", ty.Rule.Id.Name)

}

func TestParse1(t *testing.T) {

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
	assert.NoError(t, err)

	cb := tests.Context(path).Dir(dir).Cmd().Sempty().Jsystem()

	_, err = parse.Parse(cb.Ctx(), path)
	assert.NoError(t, err)

	scache := sysctx.FromContext(cb.Ctx())
	i, ok := scache.GetType(path, "b")
	assert.True(t, ok)
	ty, ok := i.(*system.Type)
	assert.True(t, ok)
	assert.Equal(t, "a", ty.Description)

}

func TestParse(t *testing.T) {

	n, err := tests.CreateTemporaryNamespace()
	assert.NoError(t, err)
	defer os.RemoveAll(n)

	files := map[string]string{
		"a.json": `{
			"type": "system:type",
			"id": "a",
			"rule": {
				"type": "system:type",
				"embed": ["system:rule"],
				"fields": {
					"a": {
						"type": "system:@bool"
					}
				}
			}
		}`,
		"b.json": `{
			"type": "system:type",
			"id": "b",
			"fields": {
				"c": {
					"type": "@a",
					"a": true,
					"optional": true
				},
				"d": {
					"type": "@b"
				}
			}
		}`,
	}
	path, dir, _, err := tests.CreateTemporaryPackage(n, "a", files)
	assert.NoError(t, err)

	path = "kego.io/system"

	cb := tests.Context(path).Dir(dir).Cmd().Sempty().Jsystem()

	pi, err := parse.Parse(cb.Ctx(), path)
	assert.NoError(t, err)

	_, err = process.Structs(cb.Ctx(), pi.Environment)
	assert.NoError(t, err)
}
