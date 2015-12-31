package system

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/kerr"
)

// Enforcer is a rule with properties that need to be enforced against data.
type Enforcer interface {
	Enforce(ctx context.Context, data interface{}) (bool, string, error)
}

type CollectionRule interface {
	GetItemsRule() RuleInterface
}
type DefaultRule interface {
	GetDefault() interface{}
}

type DummyRule struct {
	*Object
	*Rule
	Default interface{}
	Items   RuleInterface
}

func (d *DummyRule) GetDefault() interface{} {
	return d.Default
}

func (d *DummyRule) GetItemsRule() RuleInterface {
	return d.Items
}

var _ DefaultRule = (*DummyRule)(nil)
var _ CollectionRule = (*DummyRule)(nil)

func init() {
	json.RegisterInterface(reflect.TypeOf((*RuleInterface)(nil)).Elem(), reflect.TypeOf(&DummyRule{}))
}

type RuleWrapper struct {
	ctx       context.Context
	Interface RuleInterface
	Struct    *Rule
	Parent    *Type
}

func (r *RuleWrapper) GetReflectType() (reflect.Type, error) {

	if r.Struct.Interface {
		typ, ok := r.Parent.Id.GetReflectInterface()
		if !ok {
			return nil, kerr.New("QGUVEUTXAN", nil, "Type interface for %s not found", r.Parent.Id.Value())
		}
		return typ, nil
	}

	switch r.Parent.Native.Value() {
	case "object", "number", "bool", "string":
		typ, ok := r.Parent.Id.GetReflectType()
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
		items, err := WrapRule(r.ctx, itemsRule)
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
func WrapRule(ctx context.Context, r RuleInterface) (*RuleWrapper, error) {

	ob, ok := r.(ObjectInterface)
	if !ok {
		return nil, kerr.New("VKFNPJDNVB", nil, "%T does not implement ObjectInterface", r)
	}

	// ob.GetObject(nil).Type will be a @ rule type, so we change to the type,
	// which will be the parent
	t, ok := ob.GetObject(nil).Type.ChangeToType().GetType(ctx)
	if !ok {
		return nil, kerr.New("KYCTDXKFYR", nil, "GetType: type %v not found", ob.GetObject(nil).Type.ChangeToType().Value())
	}

	return &RuleWrapper{ctx: ctx, Interface: r, Parent: t, Struct: r.GetRule(nil)}, nil
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
	w, err := WrapRule(r.ctx, ir)
	if err != nil {
		return nil, kerr.New("SDSMCXSWOF", err, "WrapRule")
	}
	return w, nil
}

// HoldsDisplayType returns the string to display when communicating to
// the end user what type this rule holds.
func (r *RuleWrapper) HoldsDisplayType(ctx context.Context) (string, error) {
	str, err := r.Parent.Id.ValueContext(ctx)
	if err != nil {
		return "", kerr.New("OPIFCOHGWI", err, "ValueContext")
	}
	if r.Struct.Interface || r.Parent.Interface {
		str += "*"
	}
	return str, nil
}
