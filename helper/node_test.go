package helper

import (
	"testing"

	"fmt"

	"kego.io/kerr/assert"
	_ "kego.io/system"
	_ "kego.io/system/types"
)

func TestFoo(t *testing.T) {
	s := `{
		"description": "Restriction rules for bools",
		"type": "type",
		"is": ["rule"],
		"fields": {
			"default": {
				"description": "Default value if this is missing or null",
				"type": "@bool",
				"optional": true
			}
		}
	}`
	n, err := Unpack([]byte(s), "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	fmt.Printf("%#v\n", n)
}
