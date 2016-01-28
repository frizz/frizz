package node_test

import (
	"testing"

	"golang.org/x/net/context"
	"kego.io/kerr/assert"
	"kego.io/process"
	_ "kego.io/system"
	"kego.io/system/node"
)

func TestFoo(t *testing.T) {

	ctx, _, err := process.Initialise(context.Background(), process.FromDefaults{Path: "kego.io/system"})
	assert.NoError(t, err)

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
	_, err = node.Unmarshal(ctx, []byte(s))
	assert.NoError(t, err)
}
