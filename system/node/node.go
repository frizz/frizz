package node // import "kego.io/system/node"

// ke: {"package": {"complete": true}}

import (
	"fmt"

	"reflect"

	"sort"

	"context"

	"github.com/davelondon/kerr"
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
	hash        uint64
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

func (n *Node) SetValue(ctx context.Context, value interface{}) error {
	switch n.JsonType {
	case json.J_STRING:
		val := value.(string)
		if n.ValueString == val {
			// ignore the change if there's no change to the value
			return nil
		}
		if err := n.SetValueString(ctx, val); err != nil {
			return kerr.Wrap("NCIMXDORED", err)
		}
	case json.J_BOOL:
		val := value.(bool)
		if n.ValueBool == val {
			// ignore the change if there's no change to the value
			return nil
		}
		if err := n.SetValueBool(ctx, val); err != nil {
			return kerr.Wrap("HKFEEMFRHR", err)
		}
	case json.J_NUMBER:
		val := value.(float64)
		if n.ValueNumber == val {
			// ignore the change if there's no change to the value
			return nil
		}
		if err := n.SetValueNumber(ctx, val); err != nil {
			return kerr.Wrap("LBEBBNFJVG", err)
		}
	}
	return nil
}

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
	if err := child.setZero(ctx, true, true); err != nil {
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
	n.Value = n.Val.Interface()
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

func (n *Node) InitialiseArrayItem(ctx context.Context, parent *Node, index int) error {
	if err := n.initialiseCollectionItem(ctx, parent, "", index); err != nil {
		return kerr.Wrap("UHIEPCAQAL", err)
	}
	return nil
}
func (n *Node) InitialiseMapItem(ctx context.Context, parent *Node, key string) error {
	if err := n.initialiseCollectionItem(ctx, parent, key, -1); err != nil {
		return kerr.Wrap("SCHXGKLKOV", err)
	}
	return nil
}
func (n *Node) initialiseCollectionItem(ctx context.Context, parent *Node, key string, index int) error {
	n.resetAllValues()
	n.Parent = parent
	n.Index = index
	n.Key = key
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

	return nil
}

func (n *Node) AddToArray(ctx context.Context, parent *Node, index int, updateParentVal bool) error {

	if index > len(parent.Array) {
		return kerr.New("GHJIDXABLL", "Index out of bounds")
	} else if index == len(parent.Array) {
		parent.Array = append(parent.Array, n)
	} else {
		na := []*Node{}
		for i := 0; i < len(parent.Array)+1; i++ {
			if i < index {
				na = append(na, parent.Array[i])
			} else if i == index {
				na = append(na, n)
			} else {
				na = append(na, parent.Array[i-1])
			}
		}
		parent.Array = na
	}

	if updateParentVal {
		val := n.Val
		if val == (reflect.Value{}) {
			val = reflect.Zero(parent.Val.Type().Elem())
		}
		if index > parent.Val.Len() {
			return kerr.New("YFKMXFUPHY", "Index out of bounds")
		} else if index == parent.Val.Len() {
			appendToSlice(parent.Val, val)
		} else {
			insertIntoSlice(parent.Val, index, val)
		}
		parent.Value = parent.Val.Interface()
	}

	n.initialiseValFromParent()

	return nil
}

func (n *Node) AddToMap(ctx context.Context, parent *Node, key string, updateParentVal bool) error {

	parent.Map[key] = n

	if updateParentVal {
		val := n.Val
		if val == (reflect.Value{}) {
			val = reflect.Zero(parent.Val.Type().Elem())
		}
		parent.Val.SetMapIndex(reflect.ValueOf(key), val)
	}

	n.initialiseValFromParent()

	return nil
}

func (n *Node) initialiseObjectField(ctx context.Context, parent *Node, rule *system.RuleWrapper, key string, origin *system.Reference) error {

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

	return nil

}

func (n *Node) AddToObject(ctx context.Context, parent *Node, rule *system.RuleWrapper, key string, updateParentVal bool) error {

	parent.Map[key] = n

	if updateParentVal {
		rt, err := rule.GetReflectType()
		if err != nil {
			return kerr.Wrap("QMGGBWEMPT", err)
		}
		val := n.Val
		if val == (reflect.Value{}) {
			val = reflect.Zero(rt)
		}
		p := parent.Val
		for p.Kind() == reflect.Interface || p.Kind() == reflect.Ptr {
			p = p.Elem()
		}
		f := p.FieldByName(system.GoName(key))
		f.Set(val)
	}

	n.initialiseValFromParent()

	return nil
}

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

func (n *Node) SetValueUnpack(ctx context.Context, in json.Packed) error {
	if err := n.setValue(ctx, in, true); err != nil {
		return kerr.Wrap("FAVEHOUYHB", err)
	}
	return nil
}

func (n *Node) setValue(ctx context.Context, in json.Packed, unpack bool) error {

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

	if unpack {
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
		n.setVal(reflect.ValueOf(n.Value))
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
			if err := childNode.InitialiseArrayItem(ctx, n, i); err != nil {
				return kerr.Wrap("XHQKQTNRJV", err)
			}
			if err := childNode.AddToArray(ctx, n, i, false); err != nil {
				return kerr.Wrap("VWWYPDIJKP", err)
			}
			if err := childNode.setValue(ctx, child, false); err != nil {
				return kerr.Wrap("KUCBPFFJNT", err)
			}
		}
	case json.J_MAP:
		n.Map = map[string]*Node{}
		children := in.Map()
		for name, child := range children {
			childNode := NewNode()
			if err := childNode.InitialiseMapItem(ctx, n, name); err != nil {
				return kerr.Wrap("TBNWBMJDIE", err)
			}
			if err := childNode.AddToMap(ctx, n, name, false); err != nil {
				return kerr.Wrap("HTOPDOKPRE", err)
			}
			if err := childNode.setValue(ctx, child, false); err != nil {
				return kerr.Wrap("LWCSAHSBDF", err)
			}
		}
	case json.J_OBJECT:
		if err := n.initialiseFields(ctx, in, false); err != nil {
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
	if err := n.setZero(ctx, null, false); err != nil {
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

func (n *Node) setZero(ctx context.Context, null bool, missing bool) error {

	if missing && !null {
		return kerr.New("NYQULBBBHO", "If missing, must also be null")
	}

	if missing && (n.Parent == nil || n.Parent.JsonType != json.J_OBJECT) {
		return kerr.New("XRYLQWRNPH", "Parent must be J_OBJECT")
	}

	if n.Type == nil {
		return kerr.New("ABXFQOYCBA", "Can't set value without a type")
	}

	if n.Type.IsNativeCollection() && n.Rule == nil {
		return kerr.New("VGKTIRMDTJ", "Can't create collection zero value without a rule")
	}

	n.Missing = missing
	n.Null = null
	n.ValueString = ""
	n.ValueNumber = 0.0
	n.ValueBool = false
	n.Array = []*Node{}
	n.Map = map[string]*Node{}

	if null {
		n.JsonType = json.J_NULL
	} else {
		// if this node was previously null, we must reset the json type
		n.JsonType = n.Type.NativeJsonType()
	}

	var rv reflect.Value
	if n.Type.IsNativeCollection() {
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

	if !null && n.Type.IsNativeObject() {
		if err := n.initialiseFields(ctx, nil, true); err != nil {
			return kerr.Wrap("VSAXCHGCOG", err)
		}
		if err := n.setCorrectTypeField(ctx); err != nil {
			return kerr.Wrap("CUCJDNBBSU", err)
		}
	}

	return nil
}

func (n *Node) setCorrectTypeField(ctx context.Context) error {
	typeField, ok := n.Map["type"]
	if !ok {
		return kerr.New("DQKGYKFQKJ", "type field not found")
	}
	typeString, err := n.Type.Id.ValueContext(ctx)
	if err != nil {
		return kerr.Wrap("MRHEBUUXBB", err)
	}
	if err := typeField.SetValueString(ctx, typeString); err != nil {
		return kerr.Wrap("CURDKCQLGS", err)
	}
	return nil
}

func (n *Node) SetIdField(ctx context.Context, id string) error {
	idField, ok := n.Map["id"]
	if !ok {
		return kerr.New("QVLAEAEWEB", "id field not found")
	}
	if err := idField.SetValueString(ctx, id); err != nil {
		return kerr.Wrap("GYMSSRFPMO", err)
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

func (n *Node) initialiseFields(ctx context.Context, in json.Packed, updateVal bool) error {

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
		if err := childNode.initialiseObjectField(ctx, n, rule, name, typeField.Origin); err != nil {
			return kerr.Wrap("ILIHBXGROP", err)
		}
		if err := childNode.AddToObject(ctx, n, rule, name, updateVal); err != nil {
			return kerr.Wrap("LJUGPMWNPD", err)
		}
		if valueExists {
			if err := childNode.setValue(ctx, valueField, false); err != nil {
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
		// If we have a rule with the parent, and it's not an interface, then
		// we just return the parent type of the rule.
		return rule.Parent, nil
	}

	// if the rule is nil (e.g. unpacking into an unknown type) or the type is
	// an interface, we ensure the input is a map
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
			// if the input value is a native value, we will be unpacking into
			// the parent type of the rule
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
	if !t.Basic && !t.Interface {
		// All types apart from Basic types embed system:object
		ob, ok := system.GetTypeFromCache(ctx, "kego.io/system", "object")
		if !ok {
			return kerr.New("YRFWOTIGFT", "Type system:object not found in sys ctx")
		}
		if err := extractFields(ctx, fields, ob); err != nil {
			return kerr.Wrap("DTQEFALIMM", err)
		}
	}
	for _, embedRef := range t.Embed {
		embed, ok := system.GetTypeFromCache(ctx, embedRef.Package, embedRef.Name)
		if !ok {
			return kerr.New("SLIRILCARQ", "Type %s not found in sys ctx", embedRef)
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
		if o, ok := n.Value.(system.ObjectInterface); ok {
			ob := o.GetObject(ctx)
			if ob != nil && ob.Id != nil {
				return ob.Id.Name
			}
		}
		return "root"
	}
	if n.Key != "" {
		return n.Key
	}
	if l, ok := n.Value.(system.Labelled); ok && ctx != nil {
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

func (n *Node) Hash() uint64 {
	return n.hash
}

func (n *Node) RecomputeHash(ctx context.Context, children bool) error {
	h := NodeHasher{}
	h.String = n.ValueString
	h.Number = n.ValueNumber
	h.Bool = n.ValueBool
	h.Missing = n.Missing
	h.Null = n.Null
	for _, c := range n.Array {
		if children {
			if err := c.RecomputeHash(ctx, true); err != nil {
				return kerr.Wrap("WCYWBSKEJH", err)
			}
		}
		h.Array = append(h.Array, c.hash)
	}

	// Map must be in order, so order keys
	var keys []string
	for k, _ := range n.Map {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		c := n.Map[k]
		if children {
			if err := c.RecomputeHash(ctx, true); err != nil {
				return kerr.Wrap("QDVRLABPYN", err)
			}
		}
		h.Map = append(h.Map, MapItem{Key: k, Hash: c.hash})
	}

	hash, err := h.Hash(ctx)
	if err != nil {
		return kerr.Wrap("GJNHTFFNAM", err)
	}
	n.hash = hash
	return nil
}

func (n *Node) Backup() *Node {
	return &Node{
		Parent:      n.Parent,
		Array:       n.Array,
		Map:         n.Map,
		Key:         n.Key,
		Index:       n.Index,
		Origin:      n.Origin,
		ValueString: n.ValueString,
		ValueNumber: n.ValueNumber,
		ValueBool:   n.ValueBool,
		Value:       n.Value,
		Val:         reflect.ValueOf(n.Value),
		Null:        n.Null,
		Missing:     n.Missing,
		Rule:        n.Rule,
		Type:        n.Type,
		JsonType:    n.JsonType,
	}
}

func (n *Node) Restore(ctx context.Context, b *Node) {
	n.Parent = b.Parent
	n.Array = b.Array
	n.Map = b.Map
	n.Key = b.Key
	n.Index = b.Index
	n.Origin = b.Origin
	n.ValueString = b.ValueString
	n.ValueNumber = b.ValueNumber
	n.ValueBool = b.ValueBool
	n.Value = b.Value
	n.Val = reflect.ValueOf(n.Value)
	n.Null = b.Null
	n.Missing = b.Missing
	n.Rule = b.Rule
	n.Type = b.Type
	n.JsonType = b.JsonType
}

func (n *Node) NativeValue() interface{} {
	switch n.JsonType {
	case json.J_STRING:
		return n.ValueString
	case json.J_NUMBER:
		return n.ValueNumber
	case json.J_BOOL:
		return n.ValueBool
	}
	return nil
}

func (n *Node) Print(ctx context.Context) string {
	b, err := ke.MarshalContext(ctx, n.Value)
	if err != nil {
		// ke: {"block": {"notest": true}}
		return err.Error()
	}
	return string(b)
}
