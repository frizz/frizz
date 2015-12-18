package node_test

import (
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
	"kego.io/process/tests"
	_ "kego.io/system"
	"kego.io/system/node"
	_ "kego.io/system/types"
)

func TestFoo(t *testing.T) {
	s := `{
		"description": "Restriction rules for bools",
		"type": "type",
		"embed": ["rule"],
		"fields": {
			"default": {
				"description": "Default value if this is missing or null",
				"type": "@bool",
				"optional": true
			}
		}
	}`
	n := &node.Node{}
	err := json.UnmarshalUntyped(tests.PathCtx("kego.io/system"), []byte(s), n)
	assert.NoError(t, err)
}
