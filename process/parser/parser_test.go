package parser

import (
	"testing"

	"io/ioutil"
	"path/filepath"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/context/sysctx"
	"kego.io/process/generate"
	"kego.io/system"
	"kego.io/tests"
)

func TestRuleId(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	pathA, dirA := cb.TempPackage("a", map[string]string{
		"a.json": `{
			"type": "system:type",
			"id": "a",
			"rule": {
				"type": "system:type",
				"id": "foo"
			}
		}`,
	})
	cb.Path(pathA).Dir(dirA).Cmd().Sempty().Jsystem()
	_, err := Parse(cb.Ctx(), pathA)
	assert.IsError(t, err, "VFUNPHUFHD")
	assert.HasError(t, err, "JKARKEDTIW")

	cb.TempFile("a.json", `{
			"type": "system:type",
			"id": "a",
			"rule": {
				"type": "system:type"
			}
		}`)
	_, err = Parse(cb.Ctx(), pathA)
	assert.IsError(t, err, "VFUNPHUFHD")
	assert.HasError(t, err, "LMALEMKFDI")
}

func TestTypesUnknownType(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	pathA, dirA := cb.TempPackage("a", map[string]string{
		"a.json": `{
			"type": "system:type",
			"id": "a",
			"rules": [
				{
					"type": "system:foo"
				}
			]
		}`,
	})
	cb.Path(pathA).Dir(dirA).Cmd().Sempty().Jsystem()
	_, err := Parse(cb.Ctx(), pathA)
	// unknown types are tolerated when scanning types
	require.NoError(t, err)

	cb.TempFile("b.json", `{
			"type": "system:package",
			"rules": [
				{
					"type": "foo:bar"
				}
			]
		}`)
	_, err = Parse(cb.Ctx(), pathA)
	// b.json is not a type so unknown packages and types are permitted
	require.NoError(t, err)

	cb.TempFile("c.json", `{
			"type": "system:type",
			"id": "c",
			"rules": [
				{
					"type": "foo:bar"
				}
			]
		}`)
	_, err = Parse(cb.Ctx(), pathA)
	// unknown packages are not tolerated when scanning types
	assert.HasError(t, err, "DKKFLKDKYI")

}

func TestNoIdError(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	pathA, dirA := cb.TempPackage("a", map[string]string{
		"a.json": `{
			"type": "system:type"
		}`,
	})
	cb.Path(pathA).Dir(dirA).Cmd().Sempty().Jsystem()
	_, err := Parse(cb.Ctx(), pathA)
	assert.IsError(t, err, "VFUNPHUFHD")
	assert.HasError(t, err, "DLLMKTDYFW")

}

func TestNoTypeError(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	pathA, dirA := cb.TempPackage("a", map[string]string{
		"a.json": `{
			"id": "a"
		}`,
	})
	cb.Path(pathA).Dir(dirA).Cmd().Sempty().Jsystem()
	_, err := Parse(cb.Ctx(), pathA)
	assert.HasError(t, err, "NUKWIHYFMQ")

}

func TestGoGet(t *testing.T) {
	cb := tests.New().TempGopath(true).Cmd()
	defer cb.Cleanup()
	pathA, _ := cb.TempPackage("a", map[string]string{
		"a.json": `{
			"type": "system:type",
			"id": "a"
		}`,
	})
	err := GoGet(cb.Ctx(), pathA)
	require.NoError(t, err)

	cb.CmdUpdate(true)
	err = GoGet(cb.Ctx(), pathA)
	assert.IsError(t, err, "NIKCKQAKUI")

}

func TestImport(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	pathA, _ := cb.TempPackage("a", map[string]string{
		"a.json": `{
			"type": "system:type",
			"id": "a"
		}`,
	})
	pathB, dirB := cb.TempPackage("b", map[string]string{
		"package.json": `{
			"type": "system:package",
			"aliases": {
				"a": "` + pathA + `"
			}
		}`,
		"b.json": `{
			"type": "a:a",
			"id": "a"
		}`,
		"c.json": `{
			"type": "system:foo",
			"id": "c"
		}`,
	})
	cb.Path(pathB).Dir(dirB).Cmd().Sempty().Jsystem()

	_, err := Parse(cb.Ctx(), pathB)
	require.NoError(t, err)

}

func TestImportSystem(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	pathA, dirA := cb.TempPackage("a", map[string]string{
		"package.json": `{
			"type": "system:package",
			"aliases": {
				"s": "kego.io/system"
			}
		}`,
	})
	cb.Path(pathA).Dir(dirA).Cmd().Sempty().Jsystem()
	_, err := Parse(cb.Ctx(), pathA)
	assert.IsError(t, err, "EWMLNJDXKC")

	pathB, dirB := cb.TempPackage("b", map[string]string{
		"package.json": `{
			"type": "system:package",
			"aliases": {
				"system": "a.b/c"
			}
		}`,
	})
	cb.Path(pathB).Dir(dirB)
	_, err = Parse(cb.Ctx(), pathB)
	assert.IsError(t, err, "EWMLNJDXKC")

}

func TestCircularImport(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	pathA, dirA := cb.TempPackage("a", map[string]string{})
	pathB, dirB := cb.TempPackage("b", map[string]string{})
	ioutil.WriteFile(filepath.Join(dirA, "package.json"), []byte(`
			{
				"type": "system:package",
				"aliases": {
					"b": "`+pathB+`"
				}
			}
		`), 0777)
	ioutil.WriteFile(filepath.Join(dirB, "package.json"), []byte(`
			{
				"type": "system:package",
				"aliases": {
					"a": "`+pathA+`"
				}
			}
		`), 0777)

	cb.Path(pathA).Dir(dirA).Cmd().Sempty().Jsystem()

	_, err := Parse(cb.Ctx(), pathA)
	assert.IsError(t, err, "NOVMGYKHHI")
	assert.HasError(t, err, "SCSCFJPPHD")

}

func TestScan_errors(t *testing.T) {

	cb := tests.Context("").Cmd().Sempty().Jsystem().TempGopath(true)
	defer cb.Cleanup()

	_, err := Parse(cb.Ctx(), "does-not-exist")
	require.IsError(t, err, "SBALWXUPKN") // Error from ScanForEnv
	assert.HasError(t, err, "NIKCKQAKUI") // GoGet: exit status 1: package does-not-exist

	path, _ := cb.TempPackage("a", map[string]string{
		"a.json": `{
			"type": "system:package",
			"recursive": false
		}`,
		"b.json": "foo",
	})

	cb.Path(path)

	_, err = Parse(cb.Ctx(), path)
	assert.IsError(t, err, "VFUNPHUFHD")  // Error from scanForTypes
	assert.HasError(t, err, "HCYGNBDFFA") // Error trying to unmarshal a type

	path, _ = cb.TempPackage("a", map[string]string{
		"a.json": "foo",
	})

	cb.Path(path)

	_, err = Parse(cb.Ctx(), path)
	assert.IsError(t, err, "GJRHNGGWFD")  // Error from ScanForEnv
	assert.HasError(t, err, "MTDCXBYBEJ") // Error trying to scan for packages

}

func TestParseRule(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
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

	cb.Path(path).Dir(dir).Cmd().Sempty().Jsystem()

	_, err := Parse(cb.Ctx(), path)
	require.NoError(t, err)

	scache := sysctx.FromContext(cb.Ctx())
	i, ok := scache.GetType(path, "b")
	assert.True(t, ok)
	ty, ok := i.(*system.Type)
	assert.True(t, ok)
	assert.Equal(t, "@b", ty.Rule.Id.Name)

}

func TestParse1(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	path, dir := cb.TempPackage("a", map[string]string{
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

	cb.Path(path).Dir(dir).Cmd().Sempty().Jsystem()

	_, err := Parse(cb.Ctx(), path)
	require.NoError(t, err)

	scache := sysctx.FromContext(cb.Ctx())
	i, ok := scache.GetType(path, "b")
	assert.True(t, ok)
	ty, ok := i.(*system.Type)
	assert.True(t, ok)
	assert.Equal(t, "a", ty.Description)

}

func TestParse(t *testing.T) {

	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

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
	path, dir := cb.TempPackage("a", files)

	path = "kego.io/system"

	cb.Path(path).Dir(dir).Cmd().Sempty().Jsystem()

	pi, err := Parse(cb.Ctx(), path)
	require.NoError(t, err)

	_, err = generate.Structs(cb.Ctx(), pi.Env)
	require.NoError(t, err)

}
