package system

import (
	"reflect"

	"kego.io/json"
	"kego.io/kerr"
)

// Enforcer is a rule with properties that need to be enforced against data.
type Enforcer interface {
	Enforce(data interface{}, path string, aliases map[string]string) (bool, string, error)
}

type CollectionRule interface {
	GetItemsRule() RuleInterface
}

func init() {
	type dummyRule struct {
		*Object
		*Rule
		Default interface{}
	}
	json.RegisterInterface(reflect.TypeOf((*RuleInterface)(nil)).Elem(), reflect.TypeOf(&dummyRule{}))
}

type RuleWrapper struct {
	Interface RuleInterface
	Type      *Type
}

func (r *RuleWrapper) GetReflectType() (reflect.Type, error) {

	if r.Interface.GetRule().Interface {
		typ, _, ok := json.GetTypeInterface(r.Type.Id.Package, r.Type.Id.Name)
		if !ok {
			return nil, kerr.New("QGUVEUTXAN", nil, "Type interface for %s not found", r.Type.Id.Value())
		}
		return typ, nil
	}

	switch r.Type.Native.Value {
	case "object", "number", "bool", "string":
		typ, _, ok := json.GetType(r.Type.Id.Package, r.Type.Id.Name)
		if !ok {
			return nil, kerr.New("DLAJJPJDPL", nil, "Type %s not found", r.Type.Id.Value())
		}
		if r.Type.Native.Value != "object" && typ.Kind() == reflect.Ptr {
			return typ.Elem(), nil
		}
		return typ, nil
	case "array", "map":
		c, ok := r.Interface.(CollectionRule)
		if !ok {
			return nil, kerr.New("GSYSHQOWNH", nil, "Map types must have rule that implements CollectionRule")
		}
		itemsRule := c.GetItemsRule()
		items, err := WrapRule(itemsRule)
		if err != nil {
			return nil, kerr.New("EDEMPPVUNW", err, "NewRuleHolder")
		}
		itemsType, err := items.GetReflectType()
		if err != nil {
			return nil, kerr.New("LMKEHHWHKL", err, "GetReflectType")
		}
		if r.Type.Native.Value == "map" {
			return reflect.MapOf(reflect.TypeOf(""), itemsType), nil
		}
		return reflect.SliceOf(itemsType), nil
	default:
		return nil, kerr.New("VDEORSSUWA", nil, "Unknown native %s", r.Type.Native.Value)
	}
}

func WrapEmptyRule(t *Type) *RuleWrapper {
	return &RuleWrapper{Interface: nil, Type: t}
}
func WrapRule(r RuleInterface) (*RuleWrapper, error) {

	ob, ok := r.(ObjectInterface)
	if !ok {
		return nil, kerr.New("VKFNPJDNVB", nil, "%T does not implement ObjectInterface", r)
	}

	// ob.GetObject().Type will be a @ rule type, so we change to the type,
	// which will be the parent
	t, ok := ob.GetObject().Type.ChangeToType().GetType()
	if !ok {
		return nil, kerr.New("KYCTDXKFYR", nil, "GetType: type %v not found", ob.GetObject().Type.ChangeToType().Value())
	}

	return &RuleWrapper{Interface: r, Type: t}, nil
}

// ItemsRule returns Items rule for a collection Rule.
func (r *RuleWrapper) ItemsRule() (*RuleWrapper, error) {
	if !r.Type.IsNativeCollection() {
		return nil, kerr.New("VPAGXSTQHM", nil, "%s is not a collection", r.Type.Id.Value())
	}
	c, ok := r.Interface.(CollectionRule)
	if !ok {
		return nil, kerr.New("TNRVQVJIFH", nil, "%T is not a CollectionRule", r.Interface)
	}
	ir := c.GetItemsRule()
	if ir == nil {
		return nil, kerr.New("SUJLYBXPYS", nil, "%s has nil items rule", r.Type.Id.Value())
	}
	w, err := WrapRule(ir)
	if err != nil {
		return nil, kerr.New("SDSMCXSWOF", err, "WrapRule")
	}
	return w, nil
}

func RuleFieldByReflection(object interface{}, name string) (value interface{}, pointer interface{}, ok bool, err error) {
	v := reflect.ValueOf(object)
	if v.Kind() != reflect.Ptr {
		return nil, nil, false, kerr.New("QOYMWPXWUO", nil, "val.Kind (%s) is not a Ptr: %v", name, v.Kind())
	}
	if v.Elem().Kind() != reflect.Struct {
		return nil, nil, false, kerr.New("IGOUOBGXAN", nil, "val.Elem().Kind (%s) is not a Struct: %v", name, v.Elem().Kind())
	}
	value, pointer, _, found, zero, err := GetObjectField(v.Elem(), name)

	// zero => !ok
	return value, pointer, found && !zero, err
}

func GetMapMember(v reflect.Value, name string) (object interface{}, pointer interface{}, value reflect.Value, found bool, zero bool, err error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	member := v.MapIndex(reflect.ValueOf(name))
	return returnValue(member)
}
func GetArrayMember(v reflect.Value, index int) (object interface{}, pointer interface{}, value reflect.Value, found bool, zero bool, err error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	member := v.Index(index)
	return returnValue(member)
}
func GetObjectField(v reflect.Value, name string) (object interface{}, pointer interface{}, value reflect.Value, found bool, zero bool, err error) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	field := v.FieldByName(name)
	return returnValue(field)
}
func returnValue(v reflect.Value) (object interface{}, pointer interface{}, value reflect.Value, found bool, zero bool, err error) {
	value = v
	empty := reflect.Value{}
	if v == empty {
		// Field does not exist
		return
	}
	if v.Kind() == reflect.Ptr {
		// If it's a pointer we should only return not found if
		// it's nil:
		if v.IsNil() {
			zero = true
			return
		}
	} else if v.Kind() == reflect.Map || v.Kind() == reflect.Slice {
		if v.Len() == 0 {
			zero = true
		}
	} else {
		// If it's not a pointer, we return not found if it's an
		// zero value
		nilValue := reflect.Zero(v.Type())
		if v.Interface() == nilValue.Interface() {
			zero = true
		}
	}
	found = true
	object = v.Interface()
	// This prevents **Foo being returned for pointer when value is already *Foo
	if v.Kind() == reflect.Ptr {
		pointer = v.Interface()
	} else if v.CanAddr() {
		pointer = v.Addr().Interface()
	}
	return
}

// HoldsDisplayType returns the string to display when communicating to
// the end user what type this rule holds.
func (r *RuleWrapper) HoldsDisplayType(path string, aliases map[string]string) (string, error) {
	str, err := r.Type.Id.ValueContext(path, aliases)
	if err != nil {
		return "", kerr.New("OPIFCOHGWI", err, "ValueContext")
	}
	if r.Interface.GetRule().Interface || r.Type.Interface {
		str += "*"
	}
	return str, nil
}
