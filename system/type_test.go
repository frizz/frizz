package system

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
			"properties": {
				"default": {
					"description": "Default value of this is missing or null",
					"type": "property",
					"optional": true,
					"item": {
						"type": "@bool"
					}
				}
			}
		}
	}`

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Description.Exists)
	assert.Equal(t, f.Description.Value, "This is the native json bool data type")
	assert.True(t, f.Native.Exists)
	assert.Equal(t, f.Native.Value, "bool")
	assert.Equal(t, f.Rule.Id.Value, "@bool")

}

func TestType(t *testing.T) {
	data := `{
		"description": "This is the most basic type.",
		"type": "type",
		"id": "type",
		"properties": {
			"extends": {
				"description": "Type which this should extend",
				"type": "property",
				"optional": true,
				"item": {
					"type": "@reference",
					"default": "kego.io/system:object"
				}
			},
			"is": {
				"description": "Array of interface types that this type should support",
				"type": "property",
				"optional": true,
				"item": {
					"type": "@array",
					"items": {
						"type": "@reference"
					}
				}
			},
			"native": {
				"description": "This is the native json type that represents this type. If omitted, default is object.",
				"type": "property",
				"optional": true,
				"item": {
					"type": "@string",
					"enum": ["string", "number", "bool", "array", "object", "map"],
					"default": "object"
				}
			},
			"interface": {
				"description": "Is this type an interface?",
				"type": "property",
				"optional": true,
				"item": {
					"type": "@bool",
					"default": false
				}
			},
			"properties": {
				"description": "Each field is listed with it's type",
				"type": "property",
				"optional": true,
				"item": {
					"type": "@map",
					"items": {
						"type": "@property"
					}
				}
			},
			"rule": {
				"description": "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.",
				"optional": true,
				"type": "property",
				"item": {
					"type": "@type"
				}
			}
		},
		"rule": {
			"description": "Restriction rules for types",
			"type": "type",
			"id": "@type",
			"is": ["rule"]
		}
	}`

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.True(t, f.Description.Exists)
	assert.Equal(t, f.Description.Value, "This is the most basic type.")
	assert.True(t, f.Native.Exists)
	assert.Equal(t, f.Native.Value, "object")
	assert.Equal(t, f.Rule.Id.Value, "@type")
	assert.True(t, f.Properties["interface"].Description.Exists)
	assert.Equal(t, f.Properties["interface"].Description.Value, "Is this type an interface?")
	assert.Equal(t, f.Properties["interface"].Optional.Value, true)
	r, ok := f.Properties["interface"].Item.(*Bool_rule)
	assert.True(t, ok, "Wrong type %T\n", f.Properties["interface"].Item)
	assert.True(t, r.Default.Exists)
	assert.Equal(t, r.Default.Value, false)

}

func unmarshalDiagram(t *testing.T) {
	diagram := `{
		"description": "This is a type of image, which just contains the url of the image",
		"type": "system:type",
		"id": "diagram",
		"properties": {
			"url": {
				"type": "system:property",
				"item": {
					"type": "system:@string"
				}
			}
		},
		"rule": {
			"description": "Restriction rules for diagram",
			"type": "system:type",
			"id": "@diagram",
			"is": ["system:rule"],
			"properties": {
				"default": {
					"description": "Default value",
					"type": "system:property",
					"optional": true,
					"defaulter": true,
					"item": {
						"type": "@diagram"
					}
				}
			}
		}
	}`

	var i interface{}
	err := json.UnmarshalTyped([]byte(diagram), &i, "kego.io/gallery", map[string]string{"system": "kego.io/system"})
	assert.NoError(t, err)
	d, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, d)

	fullname := fmt.Sprintf("%s:%s", "kego.io/gallery", d.Id.Value)
	RegisterType(fullname, d)
	if d.Rule != nil {
		rulename := fmt.Sprintf("%s:%s", "kego.io/gallery", d.Rule.Id.Value)
		RegisterType(rulename, d.Rule)
	}

}

func TestUnknownRule(t *testing.T) {

	unmarshalDiagram(t)

	data := `{
		"description": "This represents a gallery - it's just a list of images",
		"type": "system:type",
		"id": "gallery",
		"properties": {
			"image": {
				"type": "system:property",
				"optional": true,
				"item": {
					"type": "@diagram",
					"default": {
						"type": "diagram",
						"url": "def"
					}
				}
			},
			"foo": {
				"type": "system:property",
				"optional": true,
				"item": {
					"type": "system:@bool",
					"default": true
				}
			},
			"ref": {
				"type": "system:property",
				"optional": true,
				"item": {
					"type": "system:@reference",
					"default": "image"
				}
			}
		}
	}`

	var i interface{}
	err := json.UnmarshalTyped([]byte(data), &i, "kego.io/gallery", map[string]string{"system": "kego.io/system"})
	assert.NoError(t, err)
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)

	s, err := f.Properties["image"].GoTypeDescriptor(map[string]string{"system": "kego.io/system"}, "kego.io/gallery")
	assert.NoError(t, err)
	assert.Equal(t, s, "*Diagram `kego:\"{\\\"default\\\": {\\\"type\\\":\\\"diagram\\\",\\\"url\\\":\\\"def\\\"}}\"`")

	b, err := f.Properties["foo"].GoTypeDescriptor(map[string]string{"system": "kego.io/system"}, "kego.io/gallery")
	assert.NoError(t, err)
	assert.Equal(t, b, "system.Bool `kego:\"{\\\"default\\\": true}\"`")

	r, err := f.Properties["ref"].GoTypeDescriptor(map[string]string{"system": "kego.io/system"}, "kego.io/gallery")
	assert.NoError(t, err)
	assert.Equal(t, r, "system.Reference `kego:\"{\\\"default\\\": \\\"kego.io/gallery:image\\\"}\"`")
}
