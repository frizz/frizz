package system

import (
	"reflect"

	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/json"
)

// Enforcer is a rule with properties that need to be enforced against data.
type Enforcer interface {
	Enforce(ctx context.Context, data interface{}) (success bool, message string, error error)
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
	jsonctx.InitDummy(reflect.TypeOf((*RuleInterface)(nil)).Elem(), reflect.TypeOf(&DummyRule{}))
}

type RuleWrapper struct {
	Ctx       context.Context
	Interface RuleInterface
	Struct    *Rule
	Parent    *Type
}

func (r *RuleWrapper) PermittedTypes() []*Type {
	if !r.Parent.Interface && !r.Struct.Interface {
		return []*Type{r.Parent}
	}
	return GetAllTypesThatImplementInterface(r.Ctx, r.Parent)
}

func (r *RuleWrapper) ZeroValue() (reflect.Value, error) {
	rt, err := r.GetReflectType()
	if err != nil {
		return reflect.Value{}, kerr.Wrap("DWWGAWUNCN", err)
	}
	return reflect.Zero(rt), nil
}

func (r *RuleWrapper) GetReflectType() (reflect.Type, error) {

	if r.Struct != nil && r.Struct.Interface {
		typ, ok := r.Parent.Id.GetReflectInterface(r.Ctx)
		if !ok {
			return nil, kerr.New("QGUVEUTXAN", "Type interface for %s not found", r.Parent.Id.Value())
		}
		return typ, nil
	}

	switch r.Parent.Native.Value() {
	case "object", "number", "bool", "string":
		typ, ok := r.Parent.Id.GetReflectType(r.Ctx)
		if !ok {
			return nil, kerr.New("DLAJJPJDPL", "Type %s not found", r.Parent.Id.Value())
		}
		return typ, nil
	case "array", "map":
		c, ok := r.Interface.(CollectionRule)
		if !ok {
			return nil, kerr.New("GSYSHQOWNH", "Collection types must have rule that implements CollectionRule")
		}
		itemsRule := c.GetItemsRule()
		items, err := WrapRule(r.Ctx, itemsRule)
		if err != nil {
			return nil, kerr.Wrap("EDEMPPVUNW", err)
		}
		itemsType, err := items.GetReflectType()
		if err != nil {
			return nil, kerr.Wrap("LMKEHHWHKL", err)
		}
		if r.Parent.NativeJsonType() == json.J_MAP {
			return reflect.MapOf(reflect.TypeOf(""), itemsType), nil
		}
		return reflect.SliceOf(itemsType), nil
	default:
		return nil, kerr.New("VDEORSSUWA", "Unknown native %s", r.Parent.Native.Value())
	}
}

func WrapEmptyRule(ctx context.Context, t *Type) *RuleWrapper {
	return &RuleWrapper{Ctx: ctx, Interface: nil, Parent: t}
}
func WrapRule(ctx context.Context, r RuleInterface) (*RuleWrapper, error) {

	ob, ok := r.(ObjectInterface)
	if !ok {
		return nil, kerr.New("VKFNPJDNVB", "%T does not implement ObjectInterface", r)
	}

	// ob.GetObject(nil).Type will be a @ rule type, so we change to the type,
	// which will be the parent
	t, ok := ob.GetObject(nil).Type.ChangeToType().GetType(ctx)
	if !ok {
		return nil, kerr.New("KYCTDXKFYR", "GetType: type %v not found", ob.GetObject(nil).Type.ChangeToType().Value())
	}

	return &RuleWrapper{Ctx: ctx, Interface: r, Parent: t, Struct: r.GetRule(nil)}, nil
}

// ItemsRule returns Items rule for a collection Rule.
func (r *RuleWrapper) ItemsRule() (*RuleWrapper, error) {
	if !r.Parent.IsNativeCollection() {
		return nil, kerr.New("VPAGXSTQHM", "%s is not a collection", r.Parent.Id.Value())
	}
	c, ok := r.Interface.(CollectionRule)
	if !ok {
		return nil, kerr.New("TNRVQVJIFH", "%T is not a CollectionRule", r.Interface)
	}
	ir := c.GetItemsRule()
	if ir == nil {
		return nil, kerr.New("SUJLYBXPYS", "%s has nil items rule", r.Parent.Id.Value())
	}
	w, err := WrapRule(r.Ctx, ir)
	if err != nil {
		return nil, kerr.Wrap("SDSMCXSWOF", err)
	}
	return w, nil
}

// HoldsDisplayType returns the string to display when communicating to
// the end user what type this rule holds.
func (r *RuleWrapper) DisplayType() (str string, err error) {
	if r.Parent.IsNativeCollection() {
		items, err := r.ItemsRule()
		if err != nil {
			return "", kerr.Wrap("CIRIPVYUBF", err)
		}
		inner, err := items.DisplayType()
		if err != nil {
			return "", kerr.Wrap("AGRCYBLJDJ", err)
		}
		if r.Parent.IsNativeArray() {
			return "[]" + inner, nil
		}
		return "map[]" + inner, nil
	}
	str, err = r.Parent.Id.ValueContext(r.Ctx)
	if err != nil {
		return "", kerr.Wrap("OPIFCOHGWI", err)
	}
	if r.Struct.Interface || r.Parent.Interface {
		str += "*"
	}
	return str, nil
}
