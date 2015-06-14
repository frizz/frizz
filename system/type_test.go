package system

import (
	"fmt"
	"testing"

	"kego.io/assert"
	"kego.io/json"
)

func TestBool(t *testing.T) {
	data := `{
		"description": "This is the native json bool data type",
		"type": "type",
		"id": "bool",
		"native": "bool",
		"rule": {
			"description": "Restriction rules for bools",
			"type": "type",
			"id": "@bool",
			"is": ["rule"],
			"embed": ["ruleBase"],
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
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "This is the native json bool data type", f.Description)
	assert.True(t, f.Native.Exists)
	assert.Equal(t, "bool", f.Native.Value)
	assert.Equal(t, "@bool", f.Rule.Id)

}

func TestType(t *testing.T) {
	data := `{
		"description": "This is the most basic type.",
		"type": "type",
		"id": "type",
		"fields": {
			"extends": {
				"description": "Type which this should extend",
				"type": "@reference",
				"default": "kego.io/system:object",
				"optional": true
			},
			"is": {
				"description": "Array of interface types that this type should support",
				"type": "@array",
				"items": {
					"type": "@reference"
				},
				"optional": true
			},
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
			"id": "@type",
			"is": ["rule"],
			"embed": ["ruleBase"]
		}
	}`

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "This is the most basic type.", f.Description)
	assert.True(t, f.Native.Exists)
	assert.Equal(t, "object", f.Native.Value)
	assert.Equal(t, "@type", f.Rule.Id)
	assert.Equal(t, "Is this type an interface?", f.Fields["interface"].(Object).GetBase().Description)
	assert.Equal(t, true, f.Fields["interface"].GetRuleBase().Optional)
	r, ok := f.Fields["interface"].(*Bool_rule)
	assert.True(t, ok, "Wrong type %T\n", f.Fields["interface"])
	assert.True(t, r.Default.Exists)
	assert.Equal(t, false, r.Default.Value)

}

func unmarshalDiagram(t *testing.T) {
	diagram := `{
		"description": "This is a type of image, which just contains the url of the image",
		"type": "system:type",
		"id": "diagram",
		"fields": {
			"url": {
				"type": "system:@string"
			}
		},
		"rule": {
			"description": "Restriction rules for diagram",
			"type": "system:type",
			"id": "@diagram",
			"is": ["system:rule"],
			"embed": ["system:ruleBase"],
			"fields": {
				"default": {
					"description": "Default value",
					"type": "@diagram",
					"optional": true
				}
			}
		}
	}`

	var i interface{}
	unknown, err := json.Unmarshal([]byte(diagram), &i, "kego.io/gallery", map[string]string{})
	assert.False(t, unknown)
	assert.NoError(t, err)
	d, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, d)

	fullname := fmt.Sprintf("%s:%s", "kego.io/gallery", d.Id)
	RegisterType(fullname, d)
	if d.Rule != nil {
		rulename := fmt.Sprintf("%s:%s", "kego.io/gallery", d.Rule.Id)
		RegisterType(rulename, d.Rule)
	}

}

func TestUnknownRule(t *testing.T) {

	unmarshalDiagram(t)

	data := `{
		"description": "This represents a gallery - it's just a list of images",
		"type": "system:type",
		"id": "gallery",
		"fields": {
			"image": {
				"type": "@diagram",
				"default": {
					"type": "diagram",
					"url": "def"
				},
				"optional": true
			},
			"foo": {
				"type": "system:@bool",
				"default": true,
				"optional": true
			},
			"ref": {
				"type": "system:@reference",
				"default": "image",
				"optional": true
			}
		}
	}`

	var i interface{}
	unknown, err := json.Unmarshal([]byte(data), &i, "kego.io/gallery", map[string]string{})
	assert.True(t, unknown)
	assert.NoError(t, err)
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)

	s, err := GoTypeDescriptor(f.Fields["image"], "kego.io/gallery", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "*Diagram `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/gallery:diagram\\\",\\\"value\\\":{\\\"type\\\":\\\"diagram\\\",\\\"url\\\":\\\"def\\\"},\\\"path\\\":\\\"kego.io/gallery\\\"}}\"`", s)

	b, err := GoTypeDescriptor(f.Fields["foo"], "kego.io/gallery", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "system.Bool `kego:\"{\\\"default\\\":{\\\"value\\\":true}}\"`", b)

	r, err := GoTypeDescriptor(f.Fields["ref"], "kego.io/gallery", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "system.Reference `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/system:reference\\\",\\\"value\\\":\\\"kego.io/gallery:image\\\",\\\"path\\\":\\\"kego.io/gallery\\\"}}\"`", r)

}

func TestUnregisterType(t *testing.T) {
	UnregisterType("")
}

func TestTypeHasExtends(t *testing.T) {
	y := &Type{}
	h := y.HasExtends()
	assert.False(t, h)

	y = &Type{Extends: NewReference("a.b/c", "d")}
	h = y.HasExtends()
	assert.True(t, h)
}

func TestNativeGoType(t *testing.T) {
	n, err := nativeGoType("string")
	assert.NoError(t, err)
	assert.Equal(t, "string", n)

	n, err = nativeGoType("number")
	assert.NoError(t, err)
	assert.Equal(t, "float64", n)

	n, err = nativeGoType("bool")
	assert.NoError(t, err)
	assert.Equal(t, "bool", n)

	_, err = nativeGoType("a")
	assert.IsError(t, err, "TXQIDRBJRH")
}

func TestTypeIsNativeType(t *testing.T) {
	y := &Type{Native: NewString("bool")}
	n := y.IsNativeObject()
	assert.False(t, n)

	y = &Type{Native: NewString("array")}
	n = y.IsNativeObject()
	assert.False(t, n)

	y = &Type{Native: NewString("object")}
	n = y.IsNativeObject()
	assert.True(t, n)

	// Will also return true for any other value
	y = &Type{Native: NewString("a")}
	n = y.IsNativeObject()
	assert.True(t, n)
}

func TestTypeNativeValueGolangType(t *testing.T) {
	y := &Type{Native: NewString("bool")}
	n, err := y.NativeValueGolangType()
	assert.NoError(t, err)
	assert.Equal(t, "bool", n)
}

func TestTypeGoName(t *testing.T) {
	y := &Type{Base: &Base{Id: "aa"}}
	n := y.GoName()
	assert.Equal(t, "Aa", n)
}
