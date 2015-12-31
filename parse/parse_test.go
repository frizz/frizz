package parse

import (
	"os"
	"testing"

	"fmt"

	"golang.org/x/net/context"
	"kego.io/context/cachectx"
	"kego.io/context/envctx"
	"kego.io/kerr/assert"
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
	env, err := Parse(ctx, path, []string{})
	assert.NoError(t, err)

	ctx = envctx.NewContext(ctx, env)

	s, err := Structs(ctx)
	assert.NoError(t, err)
	fmt.Println(string(s))
}
