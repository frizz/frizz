package parse_test

import (
	"os"
	"testing"

	"kego.io/parse"

	"golang.org/x/net/context"
	"kego.io/context/cachectx"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/tests"
)

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
	path, _, _, err := tests.CreateTemporaryPackage(n, "a", files)
	assert.NoError(t, err)

	path = "kego.io/system"

	ctx := context.Background()
	ctx = envctx.NewContext(ctx, &envctx.Env{Path: path})
	ctx = cachectx.NewContext(ctx)
	ctx = cmdctx.NewContext(ctx, &cmdctx.Cmd{})
	ctx = jsonctx.NewContext(ctx)
	pi, err := parse.Parse(ctx, path, []string{})
	assert.NoError(t, err)

	ctx = envctx.NewContext(ctx, pi.Environment)

	_, err = process.Structs(ctx, pi.Environment)
	assert.NoError(t, err)
}
