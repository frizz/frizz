package node

import (
	"testing"

	"reflect"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/process/parser"
	"kego.io/process/tests"
	"kego.io/system"
)

type labelled struct {
	v string
}

func (l *labelled) Label(ctx context.Context) string {
	return l.v
}

type embedsObject struct {
	*system.Object
}

func TestNode_Label(t *testing.T) {
	var n *Node
	cb := tests.New()

	// Nil node
	assert.Equal(t, "(nil)", n.Label(cb.Ctx()))

	// Root node (parent == nil)
	n = &Node{Key: "", Index: -1}
	assert.Equal(t, "root", n.Label(cb.Ctx()))

	// Node in a map (key != "")
	n = &Node{Key: "a", Index: -1, Parent: &Node{}}
	assert.Equal(t, "a", n.Label(cb.Ctx()))

	// Node in an array (index > -1)
	n = &Node{Key: "", Index: 2, Parent: &Node{}}
	assert.Equal(t, "2", n.Label(cb.Ctx()))

	// This is an invalid Node, but the key should override the index
	n = &Node{Key: "a", Index: 2, Parent: &Node{}}
	assert.Equal(t, "a", n.Label(cb.Ctx()))

	// The label for a node in a map will always be the map key even if the value implements
	// system.Labelled.
	n = &Node{Key: "a", Index: -1, Parent: &Node{}, Value: &labelled{v: "c"}}
	assert.Equal(t, "a", n.Label(cb.Ctx()))

	// A node in an array will use the value returned by the system.Labelled interface if it
	// implements it.
	n = &Node{Key: "", Index: 2, Parent: &Node{}, Value: &labelled{v: "c"}}
	assert.Equal(t, "c", n.Label(cb.Ctx()))

	// This is an invalid node
	n = &Node{Key: "", Index: -1, Parent: &Node{}}
	assert.Equal(t, "(?)", n.Label(cb.Ctx()))
}

func TestNode_Path(t *testing.T) {
	var n *Node
	cb := tests.New()
	assert.Nil(t, n.Root())
	r := &Node{
		Key:   "",
		Index: -1,
	}
	n = &Node{
		Key:   "a",
		Index: -1,
		Parent: &Node{
			Key:   "b",
			Index: -1,
			Parent: &Node{
				Index:  2,
				Parent: r,
			},
		}}
	assert.Equal(t, "root/2/b/a", n.Path(cb.Ctx()))
	assert.Equal(t, r, n.Root())
}

func TestExtractType(t *testing.T) {
	//extractType(ctx context.Context, in json.Packed, rule *system.RuleWrapper) (*system.Type, error)

	cb := tests.Context("a.b/c")
	r := &system.RuleWrapper{
		Interface: &system.StringRule{},
		Struct:    &system.Rule{Interface: true},
		Parent:    &system.Type{Interface: true},
	}
	_, err := extractType(cb.Ctx(), json.Pack(true), r)
	assert.IsError(t, err, "TDXTPGVFAK")

	typ := &system.Type{}
	r = &system.RuleWrapper{
		Interface: &system.StringRule{},
		Struct:    &system.Rule{},
		Parent:    typ,
	}
	ty, err := extractType(cb.Ctx(), json.Pack(true), r)
	assert.NoError(t, err)
	assert.Equal(t, typ, ty)

	ty, err = extractType(cb.Ctx(), nil, nil)
	assert.NoError(t, err)
	assert.Nil(t, ty)

	ty, err = extractType(cb.Ctx(), json.Pack(1.0), nil)
	assert.IsError(t, err, "DLSQRFLINL")

	r = &system.RuleWrapper{
		Struct: &system.Rule{Interface: true},
		Parent: typ,
	}
	ty, err = extractType(cb.Ctx(), nil, r)
	assert.NoError(t, err)
	assert.Nil(t, ty)

	ty, err = extractType(cb.Ctx(), json.Pack("a"), r)
	assert.NoError(t, err)
	assert.Equal(t, typ, ty)

	ty, err = extractType(cb.Ctx(), json.Pack(1.0), r)
	assert.NoError(t, err)
	assert.Equal(t, typ, ty)

	ty, err = extractType(cb.Ctx(), json.Pack(true), r)
	assert.NoError(t, err)
	assert.Equal(t, typ, ty)

	ty, err = extractType(cb.Ctx(), json.Pack([]interface{}{""}), r)
	assert.IsError(t, err, "SNYLGBJYTM")

	ty, err = extractType(cb.Ctx(), json.Pack(map[string]interface{}{}), nil)
	assert.IsError(t, err, "HBJVDKAKBJ")

	ty, err = extractType(cb.Ctx(), json.Pack(map[string]interface{}{"a": "b"}), r)
	assert.IsError(t, err, "HBJVDKAKBJ")

	cb.Sempty()

	ty, err = extractType(cb.Ctx(), json.Pack(map[string]interface{}{"type": "a"}), r)
	assert.IsError(t, err, "IJFMJJWVCA")

	typa := &system.Type{}
	cb.Stype("a", typa)
	ty, err = extractType(cb.Ctx(), json.Pack(map[string]interface{}{"type": "a"}), r)
	assert.NoError(t, err)
	assert.Equal(t, typa, ty)

}

func TestExtractFields(t *testing.T) {
	cb := tests.Context("a.b/c").Ssystem(parser.Parse)
	f := map[string]*system.Field{"a": nil}
	ty := &system.Type{Fields: map[string]system.RuleInterface{"a": nil}}
	err := extractFields(cb.Ctx(), f, ty)
	assert.IsError(t, err, "BARXPFXQNB")

	ty = &system.Type{Basic: true, Embed: []*system.Reference{system.NewReference("a", "b")}}
	err = extractFields(cb.Ctx(), f, ty)
	assert.HasError(t, err, "VEKXQDJFGD")

}

func TestInitialiseFields(t *testing.T) {
	cb := tests.Context("a.b/c")
	n := NewNode()
	err := n.InitialiseFields(cb.Ctx(), json.Pack(true))
	// Should be J_MAP
	assert.IsError(t, err, "CVCRNWMDYF")

	cb.Ssystem(parser.Parse)
	n.Type = &system.Type{}
	err = n.InitialiseFields(cb.Ctx(), json.Pack(map[string]interface{}{"foo": "bar"}))
	assert.IsError(t, err, "SRANLETJRS")

}

func TestSetValueBool(t *testing.T) {
	cb := tests.Context("a.b/c")

	n := NewNode()
	n.Rule = &system.RuleWrapper{}
	err := n.SetValueBool(cb.Ctx(), true)
	assert.IsError(t, err, "AWRMEACQWR")
	assert.HasError(t, err, "XOOUKLGORQ")

	n.Null = true
	n.Missing = true
	n.JsonType = json.J_NULL

	cb.Ssystem(parser.Parse)
	n.Rule, _ = system.WrapRule(cb.Ctx(), &system.BoolRule{Object: &system.Object{Type: system.NewReference("kego.io/system", "@bool")}, Rule: &system.Rule{}})
	err = n.SetValueBool(cb.Ctx(), true)
	assert.NoError(t, err)

	assert.False(t, n.Null)
	assert.False(t, n.Missing)
	assert.Equal(t, json.J_BOOL, n.JsonType)
	assert.True(t, n.ValueBool)
	v, ok := n.Value.(*system.Bool)
	assert.True(t, ok)
	assert.True(t, v.Value())

}

func TestSetValueNumber(t *testing.T) {
	cb := tests.Context("a.b/c")

	n := NewNode()
	n.Rule = &system.RuleWrapper{}
	err := n.SetValueNumber(cb.Ctx(), 2)
	assert.IsError(t, err, "SOJGUGHXSX")
	assert.HasError(t, err, "XOOUKLGORQ")

	n.Null = true
	n.Missing = true
	n.JsonType = json.J_NULL

	cb.Ssystem(parser.Parse)
	n.Rule, _ = system.WrapRule(cb.Ctx(), &system.NumberRule{Object: &system.Object{Type: system.NewReference("kego.io/system", "@number")}, Rule: &system.Rule{}})
	err = n.SetValueNumber(cb.Ctx(), 2)
	assert.NoError(t, err)

	assert.False(t, n.Null)
	assert.False(t, n.Missing)
	assert.Equal(t, json.J_NUMBER, n.JsonType)
	assert.Equal(t, 2.0, n.ValueNumber)
	v, ok := n.Value.(*system.Number)
	assert.True(t, ok)
	assert.Equal(t, 2.0, v.Value())
}

func TestSetValueString(t *testing.T) {
	cb := tests.Context("a.b/c")

	n := NewNode()
	n.Rule = &system.RuleWrapper{}
	err := n.SetValueString(cb.Ctx(), "foo")
	assert.IsError(t, err, "GAMJNECRUW")
	assert.HasError(t, err, "XOOUKLGORQ")

	n.Null = true
	n.Missing = true
	n.JsonType = json.J_NULL

	cb.Ssystem(parser.Parse)
	n.Rule, _ = system.WrapRule(cb.Ctx(), &system.StringRule{Object: &system.Object{Type: system.NewReference("kego.io/system", "@string")}, Rule: &system.Rule{}})
	err = n.SetValueString(cb.Ctx(), "foo")
	assert.NoError(t, err)

	assert.False(t, n.Null)
	assert.False(t, n.Missing)
	assert.Equal(t, json.J_STRING, n.JsonType)
	assert.Equal(t, "foo", n.ValueString)
	v, ok := n.Value.(*system.String)
	assert.True(t, ok)
	assert.Equal(t, "foo", v.Value())
}

func TestInitialiseWithConcreteType(t *testing.T) {
	type num float64
	type str string
	type bol bool
	type ob struct{ *system.Object }

	cb := tests.New().Jauto().Sauto(parser.Parse)
	cb.Path("a.b/c")

	ty := &system.Type{
		Object: &system.Object{Id: system.NewReference("a.b/c", "d")},
	}

	n := NewNode()
	n.Missing = true
	n.Null = true
	n.JsonType = json.J_NULL
	n.ValueNumber = 1

	ty.Native = system.NewString("number")
	cb.Jtype("d", reflect.TypeOf(num(0)))

	err := n.InitialiseWithConcreteType(cb.Ctx(), ty)
	assert.NoError(t, err)
	assert.Equal(t, n.Type, ty)
	assert.False(t, n.Missing)
	assert.False(t, n.Null)
	assert.Equal(t, 0.0, n.ValueNumber)
	assert.Equal(t, n.JsonType, json.J_NUMBER)
	f, ok := n.Value.(num)
	assert.True(t, ok)
	assert.Equal(t, num(0.0), f)
	f, ok = n.Val.Interface().(num)
	assert.True(t, ok)
	assert.Equal(t, num(0.0), f)

	ty.Native = system.NewString("string")
	cb.Jtype("d", reflect.TypeOf(str("")))
	err = n.InitialiseWithConcreteType(cb.Ctx(), ty)
	assert.NoError(t, err)
	assert.Equal(t, "", n.ValueString)
	assert.Equal(t, n.JsonType, json.J_STRING)
	s, ok := n.Value.(str)
	assert.True(t, ok)
	assert.Equal(t, str(""), s)
	s, ok = n.Val.Interface().(str)
	assert.True(t, ok)
	assert.Equal(t, str(""), s)

	ty.Native = system.NewString("bool")
	cb.Jtype("d", reflect.TypeOf(bol(false)))
	err = n.InitialiseWithConcreteType(cb.Ctx(), ty)
	assert.NoError(t, err)
	assert.Equal(t, false, n.ValueBool)
	assert.Equal(t, n.JsonType, json.J_BOOL)
	b, ok := n.Value.(bol)
	assert.True(t, ok)
	assert.Equal(t, bol(false), b)
	b, ok = n.Val.Interface().(bol)
	assert.True(t, ok)
	assert.Equal(t, bol(false), b)

	ty.Native = system.NewString("array")
	err = n.InitialiseWithConcreteType(cb.Ctx(), ty)
	// Can't create collection zero value without a rule
	assert.IsError(t, err, "VGKTIRMDTJ")

	r, err := system.WrapRule(cb.Ctx(), &system.ArrayRule{
		Object: &system.Object{Type: system.NewReference("kego.io/system", "@array")},
		Rule:   &system.Rule{},
		Items: &system.StringRule{
			Object: &system.Object{Type: system.NewReference("kego.io/system", "@string")},
			Rule:   &system.Rule{},
		},
	})
	require.NoError(t, err)
	n.Rule = r
	err = n.InitialiseWithConcreteType(cb.Ctx(), r.Parent)
	require.NoError(t, err)
	require.IsType(t, []*system.String{}, n.Value)
	require.IsType(t, []*system.String{}, n.Val.Interface())

	r, err = system.WrapRule(cb.Ctx(), &system.MapRule{
		Object: &system.Object{Type: system.NewReference("kego.io/system", "@map")},
		Rule:   &system.Rule{},
		Items: &system.StringRule{
			Object: &system.Object{Type: system.NewReference("kego.io/system", "@string")},
			Rule:   &system.Rule{},
		},
	})
	require.NoError(t, err)
	n.Rule = r
	err = n.InitialiseWithConcreteType(cb.Ctx(), r.Parent)
	require.NoError(t, err)
	require.IsType(t, map[string]*system.String{}, n.Value)
	require.IsType(t, map[string]*system.String{}, n.Val.Interface())

	n = NewNode()
	n.Missing = true
	n.Null = true
	n.JsonType = json.J_NULL

	cb.Jtype("d", reflect.TypeOf(&ob{}))
	ty.Native = system.NewString("object")
	err = n.InitialiseWithConcreteType(cb.Ctx(), ty)
	assert.NoError(t, err)
	assert.Equal(t, "d", n.Map["type"].ValueString)
	assert.IsType(t, &ob{}, n.Value)
	assert.IsType(t, &ob{}, n.Val.Interface())

	// InitialiseFields will always create type
	assert.SkipError("DQKGYKFQKJ")

}

func TestNodeUnpack(t *testing.T) {
	n := NewNode()
	err := n.Unpack(context.Background(), json.Pack("a"))
	assert.IsError(t, err, "FUYLKYTQYD")

	cb := tests.Context("kego.io/system").Ssystem(parser.Parse)
	s := map[string]interface{}{"type": "type"}
	err = n.Unpack(cb.Ctx(), json.Pack(s))
	assert.NoError(t, err)
}

func TestNodeExtract(t *testing.T) {

	cb := tests.Context("kego.io/system").Ssystem(parser.Parse)

	s := map[string]interface{}{
		"description": "Restriction rules for bools",
		"type":        "type",
		"embed":       []interface{}{"rule"},
		"fields": map[string]interface{}{
			"default": map[string]interface{}{
				"description": "Default value if this is missing or null",
				"type":        "@bool",
				"optional":    true,
			},
		},
	}
	n := NewNode()
	err := n.extract(cb.Ctx(), nil, "", -1, &system.Reference{}, 0, json.Pack(s), true, nil)
	assert.NoError(t, err)
	f, ok := n.Map["fields"]
	assert.True(t, ok)
	d, ok := f.Map["default"]
	assert.True(t, ok)
	o, ok := d.Map["optional"]
	assert.True(t, ok)
	assert.True(t, o.ValueBool)

	// Can't extract the object type from the rule (can't have interface rule and parent)
	r := &system.RuleWrapper{
		Interface: &system.StringRule{},
		Struct:    &system.Rule{Interface: true},
		Parent:    &system.Type{Interface: true},
	}
	err = n.extract(cb.Ctx(), nil, "", -1, &system.Reference{}, 0, json.Pack(s), true, r)
	assert.IsError(t, err, "RBDBRRUVMM")
	assert.HasError(t, err, "TDXTPGVFAK")

	// Rule specifies a string type, but we're unpacking a number
	r = &system.RuleWrapper{
		Interface: &system.StringRule{},
		Struct:    &system.Rule{},
		Parent:    &system.Type{Native: system.NewString("string")},
	}
	err = n.extract(cb.Ctx(), nil, "", -1, &system.Reference{}, 0, json.Pack(1.0), true, r)
	assert.IsError(t, err, "VEPLUIJXSN")

	// Id is the wrong type
	s = map[string]interface{}{"type": "package", "id": 1.0}
	err = n.extract(cb.Ctx(), nil, "", -1, &system.Reference{}, 0, json.Pack(s), true, nil)
	assert.IsError(t, err, "RUHTPJFDOG")

	s = map[string]interface{}{"type": "package"}
	err = n.extract(cb.Ctx(), nil, "", -1, &system.Reference{}, 0, json.Pack(s), true, nil)
	assert.NoError(t, err)

	r.Parent = &system.Type{Object: &system.Object{Id: system.NewReference("kego.io/system", "string")}, Native: system.NewString("string")}
	r.Ctx = cb.Ctx()
	err = n.extract(cb.Ctx(), nil, "", -1, &system.Reference{}, 0, nil, true, r)
	assert.NoError(t, err)

	err = n.extract(cb.Ctx(), nil, "", -1, &system.Reference{}, 0, json.Pack("a"), true, r)
	assert.NoError(t, err)
	assert.Equal(t, "a", n.ValueString)

	r.Parent = &system.Type{Object: &system.Object{Id: system.NewReference("kego.io/system", "number")}, Native: system.NewString("number")}
	err = n.extract(cb.Ctx(), nil, "", -1, &system.Reference{}, 0, json.Pack(2.0), true, r)
	assert.NoError(t, err)
	assert.Equal(t, 2.0, n.ValueNumber)

	r.Parent = &system.Type{Object: &system.Object{Id: system.NewReference("kego.io/system", "bool")}, Native: system.NewString("bool")}
	err = n.extract(cb.Ctx(), nil, "", -1, &system.Reference{}, 0, json.Pack(true), true, r)
	assert.NoError(t, err)
	assert.Equal(t, true, n.ValueBool)

	// This is checked in Node.UpdateValue -> Rule.GetReflectType
	assert.SkipError("IUTONSPQOL")
	assert.SkipError("RTQUNQEKUY")

}

func TestUnmarshalError(t *testing.T) {
	_, err := Unmarshal(context.Background(), []byte("a"))
	assert.IsError(t, err, "QDWFKJOJPQ")
}

func TestUnmarshal(t *testing.T) {

	cb := tests.Context("kego.io/system").Ssystem(parser.Parse)

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
	n, err := Unmarshal(cb.Ctx(), []byte(s))
	assert.NoError(t, err)
	f, ok := n.Map["fields"]
	assert.True(t, ok)
	d, ok := f.Map["default"]
	assert.True(t, ok)
	o, ok := d.Map["optional"]
	assert.True(t, ok)
	assert.True(t, o.ValueBool)
}

func TestNode_DisplayType(t *testing.T) {

	n := &Node{Missing: true}
	v, err := n.DisplayType(nil)
	assert.NoError(t, err)
	assert.Equal(t, "null", v)

	n = &Node{Null: true}
	v, err = n.DisplayType(nil)
	assert.NoError(t, err)
	assert.Equal(t, "null", v)

	cb := tests.Context("kego.io/system").Ssystem(parser.Parse)

	s := `{
	"type": "type",
	"embed": ["rule"],
	"fields": {}
}`
	n, err = Unmarshal(cb.Ctx(), []byte(s))
	assert.NoError(t, err)

	v, err = n.Map["fields"].DisplayType(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "map[]rule*", v)

	v, err = n.Map["embed"].DisplayType(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "[]reference", v)

	v, err = n.Map["type"].DisplayType(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, "reference", v)

}

func TestSetValue(t *testing.T) {

	cb := tests.Context("kego.io/system").Ssystem(parser.Parse)

	s := `{
	"description": "a",
	"type": "type",
	"embed": ["c"],
	"fields": {
		"foo": {
			"description": "e",
			"type": "@int",
			"minimum": 2,
			"optional": false
		}
	},
	"rules": [
		{
			"description": "g",
			"type": "@int",
			"minimum": 4,
			"optional": false
		}
	]
}`
	n, err := Unmarshal(cb.Ctx(), []byte(s))
	require.NoError(t, err)

	n.Map["description"].SetValueString(cb.Ctx(), "b")
	assert.Equal(t, "b", n.Value.(*system.Type).Description)
	assert.Equal(t, "b", n.Val.Interface().(*system.Type).Description)

	n.Map["embed"].Array[0].SetValueString(cb.Ctx(), "d")
	assert.Equal(t, "d", n.Value.(*system.Type).Embed[0].Name)
	assert.Equal(t, "d", n.Val.Interface().(*system.Type).Embed[0].Name)

	n.Map["fields"].Map["foo"].Map["description"].SetValueString(cb.Ctx(), "f")
	assert.Equal(t, "f", n.Value.(*system.Type).Fields["foo"].(*system.IntRule).Description)
	assert.Equal(t, "f", n.Val.Interface().(*system.Type).Fields["foo"].(*system.IntRule).Description)

	n.Map["fields"].Map["foo"].Map["minimum"].SetValueNumber(cb.Ctx(), 3.0)
	assert.Equal(t, 3, n.Value.(*system.Type).Fields["foo"].(*system.IntRule).Minimum.Value())
	assert.Equal(t, 3, n.Val.Interface().(*system.Type).Fields["foo"].(*system.IntRule).Minimum.Value())

	n.Map["fields"].Map["foo"].Map["optional"].SetValueBool(cb.Ctx(), true)
	assert.Equal(t, true, n.Value.(*system.Type).Fields["foo"].(*system.IntRule).Optional)
	assert.Equal(t, true, n.Val.Interface().(*system.Type).Fields["foo"].(*system.IntRule).Optional)

	n.Map["rules"].Array[0].Map["description"].SetValueString(cb.Ctx(), "h")
	assert.Equal(t, "h", n.Value.(*system.Type).Rules[0].(*system.IntRule).Description)
	assert.Equal(t, "h", n.Val.Interface().(*system.Type).Rules[0].(*system.IntRule).Description)

	n.Map["rules"].Array[0].Map["minimum"].SetValueNumber(cb.Ctx(), 5.0)
	assert.Equal(t, 5, n.Value.(*system.Type).Rules[0].(*system.IntRule).Minimum.Value())
	assert.Equal(t, 5, n.Val.Interface().(*system.Type).Rules[0].(*system.IntRule).Minimum.Value())

	n.Map["rules"].Array[0].Map["optional"].SetValueBool(cb.Ctx(), true)
	assert.Equal(t, true, n.Value.(*system.Type).Rules[0].(*system.IntRule).Optional)
	assert.Equal(t, true, n.Val.Interface().(*system.Type).Rules[0].(*system.IntRule).Optional)
}
