package node // import "kego.io/system/node"

// xke: {"package": {"complete": true}}

import (
	"fmt"

	"reflect"

	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/ke"
	"kego.io/system"
)

type Node struct {
	Parent *Node
	Array  []*Node
	Map    map[string]*Node

	Key         string            // in an object or a map, this is the key
	Index       int               // in an array, this is the index
	Origin      *system.Reference // in an object, this is the type that the field originated from - e.g. perhaps an embedded type
	ValueString string
	ValueNumber float64
	ValueBool   bool
	Value       interface{} // unmarshalled value
	Val         reflect.Value
	Null        bool // null is true if the json is null or the field is missing
	Missing     bool // missing is only true if the field is missing
	Rule        *system.RuleWrapper
	Type        *system.Type
	JsonType    json.Type
}

func Unmarshal(ctx context.Context, data []byte) (*Node, error) {
	n := NewNode()
	if err := ke.UnmarshalUntyped(ctx, data, n); err != nil {
		return nil, kerr.Wrap("QDWFKJOJPQ", err)
	}
	return n, nil
}

func NewNode() *Node {
	n := &Node{Index: -1}
	return n
}

// Unpack unpacks a node from an unpackable
func (n *Node) Unpack(ctx context.Context, in json.Packed) error {
	n.InitialiseRoot()
	if err := n.SetValueUnpack(ctx, in); err != nil {
		return kerr.Wrap("WVFVMGWPQJ", err)
	}
	return nil
}

var _ json.Unpacker = (*Node)(nil)

func (n *Node) SetValueString(ctx context.Context, value string) error {
	if err := n.SetValueUnpack(ctx, json.Pack(value)); err != nil {
		return kerr.Wrap("GAMJNECRUW", err)
	}
	return nil
}

func (n *Node) SetValueNumber(ctx context.Context, value float64) error {
	if err := n.SetValueUnpack(ctx, json.Pack(value)); err != nil {
		return kerr.Wrap("SOJGUGHXSX", err)
	}
	return nil
}

func (n *Node) SetValueBool(ctx context.Context, value bool) error {
	if err := n.SetValueUnpack(ctx, json.Pack(value)); err != nil {
		return kerr.Wrap("AWRMEACQWR", err)
	}
	return nil
}

func (n *Node) DeleteObjectChild(ctx context.Context, field string) error {
	if n.JsonType != json.J_OBJECT {
		return kerr.New("BMUSITINTC", "Must be J_OBJECT")
	}
	child := n.Map[field]
	if err := child.setValueEmpty(ctx); err != nil {
		return kerr.Wrap("HPXJKGQYLL", err)
	}
	return nil
}

func (n *Node) DeleteMapChild(key string) error {
	if n.JsonType != json.J_MAP {
		return kerr.New("ACRGPCPPFK", "Must be J_MAP")
	}
	delete(n.Map, key)
	deleteFromMap(n.Val, key)
	n.propagateValChange()
	return nil
}

func (n *Node) DeleteArrayChild(index int) error {
	if n.JsonType != json.J_ARRAY {
		return kerr.New("NFVEWWCSMV", "Must be J_ARRAY")
	}
	n.Array = append(n.Array[:index], n.Array[index+1:]...)
	for i, n := range n.Array {
		n.Index = i
	}
	deleteFromSlice(n.Val, index)
	n.propagateValChange()
	return nil
}

func (n *Node) ReorderArrayChild(from, to int) error {

	if n.JsonType != json.J_ARRAY {
		return kerr.New("MHEXGBUQOL", "Must be J_ARRAY")
	}

	a := n.Array

	// remove the item we're moving
	item := a[from]

	a = append(
		a[:from],
		a[from+1:]...)

	// insert it back in the correct place
	n.Array = append(
		a[:to],
		append([]*Node{item}, a[to:]...)...)

	// correct the indexes
	for i, c := range n.Array {
		c.Index = i
	}

	val := n.Val.Index(from).Interface()
	deleteFromSlice(n.Val, from)
	insertIntoSlice(n.Val, to, reflect.ValueOf(val))
	n.propagateValChange()
	return nil
}

func deleteFromMap(v reflect.Value, key string) {
	v.SetMapIndex(reflect.ValueOf(key), reflect.Value{})
}

func deleteFromSlice(v reflect.Value, i int) {
	v.Set(reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len())))
}

func appendToSlice(v reflect.Value, val reflect.Value) {
	v.Set(reflect.Append(v, val))
}

func insertIntoSlice(v reflect.Value, i int, val reflect.Value) {
	v.Set(reflect.AppendSlice(v.Slice(0, i+1), v.Slice(i, v.Len())))
	v.Index(i).Set(val)
}

func (n *Node) InitialiseRoot() {
	n.resetAllValues()
}

func (n *Node) InitialiseArrayChild(ctx context.Context, parent *Node, index int, exists bool) error {
	n.resetAllValues()
	n.Parent = parent
	n.Index = index
	n.Missing = false

	rule, err := parent.Rule.ItemsRule()
	if err != nil {
		return kerr.Wrap("SBJVMGUOOA", err)
	}
	n.Rule = rule

	t, err := extractType(ctx, json.Pack(nil), rule)
	if err != nil {
		return kerr.Wrap("EQNRHQWXFJ", err)
	}
	n.Type = t

	if index > len(parent.Array) {
		return kerr.New("GHJIDXABLL", "Index out of bounds")
	} else if index == len(parent.Array) {
		parent.Array = append(parent.Array, n)
	} else {
		parent.Array[index] = n
	}

	if !exists {
		if index > parent.Val.Len() {
			return kerr.New("YFKMXFUPHY", "Index out of bounds")
		} else if index == parent.Val.Len() {
			appendToSlice(parent.Val, reflect.Zero(parent.Val.Type().Elem()))
		} else {
			parent.Val.Index(index).Set(reflect.Zero(parent.Val.Type().Elem()))
		}
		n.propagateValChange()
	}
	n.initialiseValFromParent()

	return nil
}

func (n *Node) InitialiseMapChild(ctx context.Context, parent *Node, key string, exists bool) error {
	n.resetAllValues()
	n.Parent = parent
	n.Key = key
	n.Missing = false

	rule, err := parent.Rule.ItemsRule()
	if err != nil {
		return kerr.Wrap("MMUUKGFEAA", err)
	}
	n.Rule = rule

	t, err := extractType(ctx, json.Pack(nil), rule)
	if err != nil {
		return kerr.Wrap("YYHFMEVMRO", err)
	}
	n.Type = t

	parent.Map[key] = n
	if !exists {
		parent.Val.SetMapIndex(reflect.ValueOf(key), reflect.Zero(parent.Val.Type().Elem()))
		n.propagateValChange()
	}
	n.initialiseValFromParent()
	return nil
}

func (n *Node) initialiseNewField(ctx context.Context, parent *Node, rule *system.RuleWrapper, key string, origin *system.Reference, exists bool) error {

	n.resetAllValues()
	n.Parent = parent
	n.Rule = rule
	n.Key = key
	n.Origin = origin

	t, err := extractType(ctx, json.Pack(nil), rule)
	if err != nil {
		return kerr.Wrap("RBDBRRUVMM", err)
	}
	n.Type = t

	parent.Map[key] = n
	rt, err := rule.GetReflectType()
	if err != nil {
		return kerr.Wrap("QMGGBWEMPT", err)
	}
	p := parent.Val
	for p.Kind() == reflect.Interface || p.Kind() == reflect.Ptr {
		p = p.Elem()
	}
	f := p.FieldByName(system.GoName(key))
	if !f.IsValid() {
		// sometimes the parent object doesn't have the field e.g. DummyRule
		return nil
	}

	if !exists {
		f.Set(reflect.Zero(rt))
		n.propagateValChange()
	}

	n.initialiseValFromParent()

	return nil
}

//func (n *Node) InitialiseExistingField(ctx context.Context) {
//	n.initialiseValFromParent()
//}

func (n *Node) resetAllValues() {
	n.Parent = nil
	n.Array = []*Node{}
	n.Map = map[string]*Node{}
	n.Key = ""
	n.Index = -1
	n.Origin = nil
	n.ValueString = ""
	n.ValueNumber = 0.0
	n.ValueBool = false
	n.Value = nil
	n.Val = reflect.Value{}
	n.Null = true
	n.Missing = true
	n.Rule = nil
	n.Type = nil
	n.JsonType = json.J_NULL
}

func (n *Node) initialiseValFromParent() {
	switch n.Parent.JsonType {
	case json.J_OBJECT:
		v := n.Parent.Val
		for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		n.Val = v.FieldByName(system.GoName(n.Key))
	case json.J_MAP:
		n.Val = n.Parent.Val.MapIndex(reflect.ValueOf(n.Key))
	case json.J_ARRAY:
		n.Val = n.Parent.Val.Index(n.Index)
	}
	n.Value = n.Val.Interface()
}

func (n *Node) setValueEmpty(ctx context.Context) error {

	if n.Parent == nil || n.Parent.JsonType != json.J_OBJECT {
		return kerr.New("XRYLQWRNPH", "Parent must be J_OBJECT")
	}

	n.Array = []*Node{}
	n.Map = map[string]*Node{}

	n.ValueString = ""
	n.ValueNumber = 0.0
	n.ValueBool = false

	n.Val.Set(reflect.Zero(n.Val.Type()))
	n.Value = n.Val.Interface()
	n.propagateValChange()

	n.Null = true
	n.Missing = true
	n.JsonType = json.J_NULL
	return nil
}

func (n *Node) SetValueUnpack(ctx context.Context, in json.Packed) error {
	if err := n.setValue(ctx, in, false); err != nil {
		return kerr.Wrap("FAVEHOUYHB", err)
	}
	return nil
}

func (n *Node) setValue(ctx context.Context, in json.Packed, exists bool) error {

	n.Missing = false

	objectType, err := extractType(ctx, in, n.Rule)
	if err != nil {
		return kerr.Wrap("MKMNOOYQJY", err)
	}
	n.Type = objectType

	if n.Rule == nil && objectType != nil {
		n.Rule = system.WrapEmptyRule(ctx, objectType)
	}

	if in.Type() == json.J_MAP && n.Type.IsNativeObject() {
		// for objects and maps, Type() from the json.Packed is always J_MAP, so we correct it
		// for object types here.
		n.JsonType = json.J_OBJECT
	} else {
		n.JsonType = in.Type()
	}
	n.Null = in.Type() == json.J_NULL

	// validate json type
	if !n.Null && n.Type.NativeJsonType() != n.JsonType {
		return kerr.New("VEPLUIJXSN", "json type is %s but object type is %s", n.JsonType, n.Type.NativeJsonType())
	}

	if !exists {
		var rv reflect.Value
		if n.Rule.Struct == nil {
			if err := json.Unpack(ctx, in, &n.Value); err != nil {
				return kerr.Wrap("CQMWGPLYIJ", err)
			}
		} else {
			t, err := n.Rule.GetReflectType()
			if err != nil {
				return kerr.Wrap("DQJDYPIANO", err)
			}
			if err := json.UnpackFragment(ctx, in, &n.Value, t); err != nil {
				return kerr.Wrap("PEVKGFFHLL", err)
			}
		}
		rv = reflect.ValueOf(n.Value)

		n.setVal(rv)
		n.propagateValChange()
	}

	switch n.Type.NativeJsonType() {
	case json.J_STRING:
		n.ValueString = in.String()
	case json.J_NUMBER:
		n.ValueNumber = in.Number()
	case json.J_BOOL:
		n.ValueBool = in.Bool()
	case json.J_ARRAY:
		children := in.Array()
		for i, child := range children {
			childNode := NewNode()
			if err := childNode.InitialiseArrayChild(ctx, n, i, exists); err != nil {
				return kerr.Wrap("VWWYPDIJKP", err)
			}
			if err := childNode.setValue(ctx, child, true); err != nil {
				return kerr.Wrap("KUCBPFFJNT", err)
			}
		}
	case json.J_MAP:
		n.Map = map[string]*Node{}
		children := in.Map()
		for name, child := range children {
			childNode := NewNode()
			if err := childNode.InitialiseMapChild(ctx, n, name, exists); err != nil {
				return kerr.Wrap("HTOPDOKPRE", err)
			}
			if err := childNode.setValue(ctx, child, true); err != nil {
				return kerr.Wrap("LWCSAHSBDF", err)
			}
		}
	case json.J_OBJECT:
		if err := n.initialiseFields(ctx, in, true); err != nil {
			return kerr.Wrap("XCRYJWKPKP", err)
		}
	}

	return nil
}

func (n *Node) SetValueZero(ctx context.Context, null bool, t *system.Type) error {
	if t != nil {
		if err := n.setType(t); err != nil {
			return kerr.Wrap("TGPRRSUSMQ", err)
		}
	}
	if err := n.setZero(ctx, null); err != nil {
		return kerr.Wrap("JWCKCAJXUB", err)
	}
	return nil
}

func (n *Node) setType(t *system.Type) error {
	if t.Interface {
		return kerr.New("VHOSYBMDQL", "Can't set type to an interface - must be concrete type.")
	}
	n.Type = t
	n.JsonType = t.NativeJsonType()
	return nil
}

func (n *Node) setZero(ctx context.Context, null bool) error {

	if n.Type == nil {
		return kerr.New("ABXFQOYCBA", "Can't set value without a type")
	}
	n.Missing = false
	n.Null = null
	n.ValueString = ""
	n.ValueNumber = 0.0
	n.ValueBool = false
	n.Array = []*Node{}
	n.Map = map[string]*Node{}
	if !null {
		// if this node was previously null, we must reset the json type
		n.JsonType = n.Type.NativeJsonType()
	}

	var rv reflect.Value
	if n.Type.IsNativeCollection() {
		if n.Rule == nil {
			return kerr.New("VGKTIRMDTJ", "Can't create collection zero value without a rule")
		}
		var err error
		if rv, err = n.Rule.ZeroValue(null); err != nil {
			return kerr.Wrap("WNQLTRJRBD", err)
		}
	} else {
		// this is for both objects and native values
		var err error
		if rv, err = n.Type.ZeroValue(ctx, null); err != nil {
			return kerr.Wrap("UDBVTIDRIK", err)
		}
	}
	n.Value = rv.Interface()

	n.setVal(rv)
	n.propagateValChange()

	if !null && n.Type.IsNativeObject() {
		if err := n.initialiseFields(ctx, nil, false); err != nil {
			return kerr.Wrap("VSAXCHGCOG", err)
		}
		if err := n.setCorrectTypeField(ctx); err != nil {
			return kerr.Wrap("CUCJDNBBSU", err)
		}
	}

	return nil
}

func (n *Node) setCorrectTypeField(ctx context.Context) error {
	typeString, err := n.Type.Id.ValueContext(ctx)
	if err != nil {
		return kerr.Wrap("MRHEBUUXBB", err)
	}
	typeField, ok := n.Map["type"]
	if !ok {
		return kerr.New("DQKGYKFQKJ", "type field not found")
	}
	if err := typeField.SetValueString(ctx, typeString); err != nil {
		return kerr.Wrap("CURDKCQLGS", err)
	}
	return nil
}

func (n *Node) setVal(rv reflect.Value) {
	if n.Parent == nil {
		n.Val = rv
	} else if n.Parent.Type.IsNativeMap() {
		n.Parent.Val.SetMapIndex(reflect.ValueOf(n.Key), rv)
		n.Val = n.Parent.Val.MapIndex(reflect.ValueOf(n.Key))
		for n.Val.Kind() == reflect.Interface {
			n.Val = n.Val.Elem()
		}
	} else if n.Parent.Type.IsNativeArray() {
		n.Val = n.Parent.Val.Index(n.Index)
		n.Val.Set(rv)
		for n.Val.Kind() == reflect.Interface {
			n.Val = n.Val.Elem()
		}
	} else {
		n.Val.Set(rv)
	}
}

func (n *Node) propagateValChange() {
	// Maps are not addressable, so we must check all ancestors and set accordingly.

	/*
		c := n
		for c.Parent != nil {
			if c.Parent.JsonType == json.J_MAP {
				c.Parent.Val.SetMapIndex(reflect.ValueOf(c.Key), c.Val)
			}
			c = c.Parent
		}
	*/

}

func (n *Node) initialiseFields(ctx context.Context, in json.Packed, exists bool) error {

	valueFields := map[string]json.Packed{}
	if in != nil && in.Type() != json.J_NULL {
		valueFields = in.Map()
	}

	typeFields := map[string]*system.Field{}
	if err := extractFields(ctx, typeFields, n.Type); err != nil {
		return kerr.Wrap("LPWTOSATQE", err)
	}

	for name, typeField := range typeFields {
		rule, err := system.WrapRule(ctx, typeField.Rule)
		if err != nil {
			return kerr.Wrap("YWFSOLOBXH", err)
		}
		valueField, valueExists := valueFields[name]
		childNode := NewNode()
		if err := childNode.initialiseNewField(ctx, n, rule, name, typeField.Origin, exists); err != nil {
			return kerr.Wrap("LJUGPMWNPD", err)
		}
		if valueExists {
			if err := childNode.setValue(ctx, valueField, true); err != nil {
				return kerr.Wrap("UWOTRJJVNK", err)
			}
		}
	}

	for name, _ := range valueFields {
		_, ok := typeFields[name]
		if !ok {
			return kerr.New("SRANLETJRS", "Extra field %s", name)
		}
	}

	return nil
}

func extractType(ctx context.Context, in json.Packed, rule *system.RuleWrapper) (*system.Type, error) {

	parentInterface := rule != nil && rule.Parent != nil && rule.Parent.Interface
	ruleInterface := rule != nil && rule.Struct != nil && rule.Struct.Interface

	if parentInterface && ruleInterface {
		return nil, kerr.New("TDXTPGVFAK", "Can't have interface type and interface rule at the same time")
	}

	if rule != nil && rule.Parent != nil && !parentInterface && !ruleInterface {
		// If we have a rule with the parent, and it's not an interface, then we just return the
		// parent type of the rule.
		return rule.Parent, nil
	}

	// if the rule is nil (e.g. unpacking into an unknown type) or the type is an interface, we
	// ensure the input is a map
	if rule == nil || parentInterface {
		if in == nil {
			return nil, nil
		}
		switch in.Type() {
		case json.J_NULL:
			// item is nil, so we don't know the concrete type yet.
			return nil, nil
		case json.J_MAP:
			break
		default:
			return nil, kerr.New("DLSQRFLINL", "Input %s should be J_MAP if rule is nil or an interface type", in.Type())
		}
	}

	// if the rule is an interface rule, we ensure the input is a map or a native value
	if ruleInterface {
		if in == nil {
			return nil, nil
		}
		switch in.Type() {
		case json.J_NULL:
			// item is nil, so we don't know the concrete type yet.
			return nil, nil
		case json.J_MAP:
			break
		case json.J_STRING, json.J_NUMBER, json.J_BOOL:
			// if the input value is a native value, we will be unpacking into the parent
			// type of the rule
			return rule.Parent, nil
		default:
			return nil, kerr.New("SNYLGBJYTM", "Input %s should be J_MAP, J_STRING, J_NUMBER or J_BOOL if rule is interface rule", in.Type())
		}
	}

	ob := in.Map()
	typeField, ok := ob["type"]
	if !ok {
		return nil, kerr.New("HBJVDKAKBJ", "Input must have type field if rule is nil or an interface type")
	}
	var r system.Reference
	if err := r.Unpack(ctx, typeField); err != nil {
		return nil, kerr.Wrap("YXHGIBXCOC", err)
	}
	t, ok := r.GetType(ctx)
	if !ok {
		return nil, kerr.New("IJFMJJWVCA", "Could not find type %s", r.Value())
	}

	return t, nil
}

func extractFields(ctx context.Context, fields map[string]*system.Field, t *system.Type) error {
	getType := func(r *system.Reference) (*system.Type, error) {
		t, ok := system.GetTypeFromCache(ctx, r.Package, r.Name)
		if !ok {
			return nil, kerr.New("VEKXQDJFGD", "Type %s not found in sys ctx", r.String())
		}
		return t, nil
	}
	if !t.Basic && !t.Interface {
		// All types apart from Basic types embed system:object
		ob, err := getType(system.NewReference("kego.io/system", "object"))
		if err != nil {
			return kerr.Wrap("YRFWOTIGFT", err)
		}
		if err := extractFields(ctx, fields, ob); err != nil {
			return kerr.Wrap("DTQEFALIMM", err)
		}
	}
	for _, embedRef := range t.Embed {
		embed, err := getType(embedRef)
		if err != nil {
			return kerr.Wrap("SLIRILCARQ", err)
		}
		if err := extractFields(ctx, fields, embed); err != nil {
			return kerr.Wrap("JWAPCVIYBJ", err)
		}
	}
	for name, rule := range t.Fields {
		if _, ok := fields[name]; ok {
			return kerr.New("BARXPFXQNB", "Duplicate field %s", name)
		}
		fields[name] = &system.Field{Name: name, Rule: rule, Origin: t.Id}
	}
	return nil
}

func (n *Node) Label(ctx context.Context) string {
	if n == nil {
		return "(nil)"
	}
	if n.Parent == nil {
		return "root"
	}
	if n.Key != "" {
		return n.Key
	}
	if l, ok := n.Value.(system.Labelled); ok {
		if s := l.Label(ctx); s != "" {
			return s
		}
	}
	if n.Index > -1 {
		return fmt.Sprintf("%d", n.Index)
	}
	return "(?)"
}

func (n *Node) Path() (path string) {
	for n != nil {
		if path != "" {
			path = n.Label(nil) + "/" + path
		} else {
			path = n.Label(nil)
		}
		n = n.Parent
	}
	return path
}

func (n *Node) Root() *Node {
	for n != nil {
		if n.Parent == nil {
			return n
		}
		n = n.Parent
	}
	return nil
}

func (n *Node) DisplayType(ctx context.Context) (string, error) {
	if n.Missing || n.Null {
		return "null", nil
	}
	if n.Type.IsNativeCollection() {
		str, err := n.Rule.DisplayType()
		if err != nil {
			return "", kerr.Wrap("AKTVTUJXWS", err)
		}
		return str, nil
	}
	str, err := n.Type.Id.ValueContext(ctx)
	if err != nil {
		return "", kerr.Wrap("BGWYLIFSNN", err)
	}
	return str, nil
}
