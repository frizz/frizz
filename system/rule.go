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
type DefaultRule interface {
	GetDefault() interface{}
}

type dummyRule struct {
	*Object
	*Rule
	Default interface{}
}

func (d *dummyRule) GetDefault() interface{} {
	return d.Default
}

var _ DefaultRule = (*dummyRule)(nil)

func init() {
	json.RegisterInterface(reflect.TypeOf((*RuleInterface)(nil)).Elem(), reflect.TypeOf(&dummyRule{}))
}

type RuleWrapper struct {
	Interface RuleInterface
	Struct    *Rule
	Parent    *Type
}

func (r *RuleWrapper) GetReflectType() (reflect.Type, error) {

	if r.Struct.Interface {
		typ, _, ok := json.GetTypeInterface(r.Parent.Id.Package, r.Parent.Id.Name)
		if !ok {
			return nil, kerr.New("QGUVEUTXAN", nil, "Type interface for %s not found", r.Parent.Id.Value())
		}
		return typ, nil
	}

	switch r.Parent.Native.Value() {
	case "object", "number", "bool", "string":
		typ, _, ok := json.GetType(r.Parent.Id.Package, r.Parent.Id.Name)
		if !ok {
			return nil, kerr.New("DLAJJPJDPL", nil, "Type %s not found", r.Parent.Id.Value())
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
		if r.Parent.Native.Value() == "map" {
			return reflect.MapOf(reflect.TypeOf(""), itemsType), nil
		}
		return reflect.SliceOf(itemsType), nil
	default:
		return nil, kerr.New("VDEORSSUWA", nil, "Unknown native %s", r.Parent.Native.Value())
	}
}

func WrapEmptyRule(t *Type) *RuleWrapper {
	return &RuleWrapper{Interface: nil, Parent: t}
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

	return &RuleWrapper{Interface: r, Parent: t, Struct: r.GetRule()}, nil
}

// ItemsRule returns Items rule for a collection Rule.
func (r *RuleWrapper) ItemsRule() (*RuleWrapper, error) {
	if !r.Parent.IsNativeCollection() {
		return nil, kerr.New("VPAGXSTQHM", nil, "%s is not a collection", r.Parent.Id.Value())
	}
	c, ok := r.Interface.(CollectionRule)
	if !ok {
		return nil, kerr.New("TNRVQVJIFH", nil, "%T is not a CollectionRule", r.Interface)
	}
	ir := c.GetItemsRule()
	if ir == nil {
		return nil, kerr.New("SUJLYBXPYS", nil, "%s has nil items rule", r.Parent.Id.Value())
	}
	w, err := WrapRule(ir)
	if err != nil {
		return nil, kerr.New("SDSMCXSWOF", err, "WrapRule")
	}
	return w, nil
}

// HoldsDisplayType returns the string to display when communicating to
// the end user what type this rule holds.
func (r *RuleWrapper) HoldsDisplayType(path string, aliases map[string]string) (string, error) {
	str, err := r.Parent.Id.ValueContext(path, aliases)
	if err != nil {
		return "", kerr.New("OPIFCOHGWI", err, "ValueContext")
	}
	if r.Struct.Interface || r.Parent.Interface {
		str += "*"
	}
	return str, nil
}
