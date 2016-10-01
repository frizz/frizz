package system_test

import (
	"testing"

	"context"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/process"
	"kego.io/system"
	"kego.io/tests/repacker"
	"kego.io/tests/unpacker"
)

var systemContext context.Context

func initialise() context.Context {
	if systemContext == nil {
		ctx, _, err := process.Initialise(context.Background(), process.Options{Path: "kego.io/system"})
		if err != nil {
			panic(err)
		}
		systemContext = ctx
	}
	return systemContext
}

func TestBoolExt(t *testing.T) {
	testBool(t, unpacker.Unmarshal)
	testBool(t, unpacker.Unpack)
	testBool(t, unpacker.Decode)
	testBool(t, repacker.Repack)
}
func testBool(t *testing.T, up unpacker.Interface) {

	ctx := initialise()

	data := `{
		"description": "This is the native json bool data type",
		"type": "type",
		"id": "bool",
		"native": "bool",
		"alias": {
			"type": "json:@bool"
		},
		"rule": {
			"description": "Restriction rules for bools",
			"type": "type",
			"embed": ["rule"],
			"fields": {
				"default": {
					"description": "Default value of this is missing or null",
					"type": "@bool",
					"optional": true
				}
			}
		}
	}`

	var i interface{}
	err := up.Process(ctx, []byte(data), &i)
	assert.NoError(t, err)
	f, ok := i.(*system.Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "This is the native json bool data type", f.Description)
	assert.NotNil(t, f.Native)
	assert.Equal(t, "bool", f.Native.Value())
}

func TestTypeExt(t *testing.T) {
	testType(t, unpacker.Unmarshal)
	testType(t, unpacker.Unpack)
	testType(t, unpacker.Decode)
	testType(t, repacker.Repack)
}
func testType(t *testing.T, up unpacker.Interface) {

	ctx := initialise()

	data := `{
		"description": "This is the most basic type.",
		"type": "type",
		"id": "type",
		"fields": {
			"native": {
				"description": "This is the native json type that represents this type. If omitted, default is object.",
				"type": "@string",
				"enum": ["string", "number", "bool", "array", "object", "map"],
				"default": "object",
				"optional": true
			},
			"interface": {
				"description": "Is this type an interface?",
				"type": "@bool",
				"default": false,
				"optional": true
			},
			"fields": {
				"description": "Each field is listed with it's type",
				"type": "@map",
				"items": {
					"type": "@rule"
				},
				"optional": true
			},
			"rule": {
				"description": "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.",
				"type": "@type",
				"optional": true
			}
		},
		"rule": {
			"description": "Restriction rules for types",
			"type": "type",
			"embed": ["rule"]
		}
	}`

	var i interface{}
	err := up.Process(ctx, []byte(data), &i)
	require.NoError(t, err)
	f, ok := i.(*system.Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "This is the most basic type.", f.Description)
	assert.NotNil(t, f.Native)
	assert.Equal(t, "object", f.Native.Value())
	assert.Equal(t, "Is this type an interface?", f.Fields["interface"].(system.ObjectInterface).GetObject(nil).Description)
	assert.Equal(t, true, f.Fields["interface"].GetRule(nil).Optional)
	r, ok := f.Fields["interface"].(*system.BoolRule)
	assert.True(t, ok, "Wrong type %T\n", f.Fields["interface"])
	assert.NotNil(t, r.Default)
	assert.Equal(t, false, r.Default.Value())

}
