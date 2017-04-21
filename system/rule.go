package system

import (
	"reflect"

	"context"

	"fmt"

	"frizz.io/context/jsonctx"
	"github.com/dave/kerr"
)

// Enforcer is a rule with properties that need to be enforced against data.
type Enforcer interface {
	Enforce(ctx context.Context, data interface{}) (fail bool, messages []string, error error)
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

func (v *DummyRule) Unpack(ctx context.Context, in Packed, iface bool) error {
	if in == nil || in.Type() == J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	if field, ok := in.Map()["items"]; ok && field.Type() != J_NULL {
		ob, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Items = ob
	}
	if field, ok := in.Map()["default"]; ok && field.Type() != J_NULL {
		ob, err := UnpackInterface(ctx, field)
		if err != nil {
			return err
		}
		v.Default = ob
	}
	return nil
}

func UnpackInterface(ctx context.Context, in Packed) (interface{}, error) {
	switch in.Type() {
	case J_MAP:
		i, err := UnpackUnknownType(ctx, in, true, "", "")
		if err != nil {
			return nil, err
		}
		return i, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into interface{}.", in.Type())
	}
}

func (d *DummyRule) GetDefault() interface{} {
	if d.Default == nil {
		return nil
	}
	return d.Default
}

func (d *DummyRule) GetItemsRule() RuleInterface {
	return d.Items
}

var _ DefaultRule = (*DummyRule)(nil)
var _ Unpacker = (*DummyRule)(nil)
var _ CollectionRule = (*DummyRule)(nil)

func init() {
	jsonctx.InitDummy(reflect.TypeOf((*RuleInterface)(nil)).Elem(), reflect.TypeOf(&DummyRule{}))

	pkg := jsonctx.InitPackage("frizz.io/system")
	pkg.Dummy("rule", func() interface{} { return new(DummyRule) })
}

type RuleWrapper struct {
	Ctx       context.Context
	Interface RuleInterface
	Struct    *Rule
	Parent    *Type
}

func (r *RuleWrapper) InnerType(ctx context.Context) *Type {
	inner := r
	for {
		cr, ok := inner.Interface.(CollectionRule)
		if !ok {
			return inner.Parent
		}
		ir := cr.GetItemsRule()
		if ir == nil {
			return inner.Parent
		}
		inner = WrapRule(ctx, ir)
	}
}

func (r *RuleWrapper) PassedAsPointer(ctx context.Context) bool {
	kind, alias := r.Kind(ctx)
	switch kind {
	case KindStruct:
		return true
	case KindValue:
		if alias {
			return true
		}
	}
	return false
}

func (r *RuleWrapper) PassedAsPointerString(ctx context.Context) string {
	if r.PassedAsPointer(ctx) {
		return "*"
	}
	return ""
}
func (r *RuleWrapper) Kind(ctx context.Context) (kind Kind, alias bool) {

	if cr, ok := r.Interface.(CollectionRule); ok {
		// DummyRule always implements CollectionRule but sometimes the actual
		// type doesn't support it.
		ir := cr.GetItemsRule()
		if ir != nil {
			return Kind(r.Parent.CustomKind.Value()), false
		}
	}

	if r.Struct != nil && r.Struct.Interface {
		return KindInterface, false
	}

	return r.Parent.Kind(ctx)
}

func (r *RuleWrapper) PermittedTypes() []*Type {
	typeInterface := r.Parent != nil && r.Parent.Interface
	ruleInteface := r.Struct != nil && r.Struct.Interface
	if !typeInterface && !ruleInteface {
		return []*Type{r.Parent}
	}
	return GetAllTypesThatImplementInterface(r.Ctx, r.Parent)
}

func (r *RuleWrapper) ZeroValue(null bool) (reflect.Value, error) {
	rt, err := r.GetReflectType()
	if err != nil {
		return reflect.Value{}, kerr.Wrap("DWWGAWUNCN", err)
	}
	return zeroValue(rt, null), nil
}

func (r *RuleWrapper) GetReflectType() (reflect.Type, error) {

	if r.Struct != nil && r.Struct.Interface {
		typ, ok := r.Parent.GetReflectInterface(r.Ctx)
		if !ok {
			return nil, kerr.New("QGUVEUTXAN", "Type interface for %s not found", r.Parent.Id.Value())
		}
		return typ, nil
	}

	if c, ok := r.Interface.(CollectionRule); ok {
		itemsRule := c.GetItemsRule()
		if itemsRule != nil {
			items := WrapRule(r.Ctx, itemsRule)
			itemsType, err := items.GetReflectType()
			if err != nil {
				return nil, kerr.Wrap("LMKEHHWHKL", err)
			}
			if r.Parent.NativeJsonType(r.Ctx) == J_MAP {
				return reflect.MapOf(reflect.TypeOf(""), itemsType), nil
			}
			return reflect.SliceOf(itemsType), nil
		}
	}

	typ, ok := r.Parent.GetReflectType(r.Ctx)
	if !ok {
		return nil, kerr.New("DLAJJPJDPL", "Type %s not found", r.Parent.Id.Value())
	}

	return typ, nil

}

func WrapEmptyRule(ctx context.Context, t *Type) *RuleWrapper {
	return &RuleWrapper{Ctx: ctx, Interface: nil, Parent: t}
}
func WrapRule(ctx context.Context, r RuleInterface) *RuleWrapper {

	ob, ok := r.(ObjectInterface)
	if !ok {
		panic(kerr.New("VKFNPJDNVB", "%T does not implement ObjectInterface", r).Error())
	}

	// ob.GetObject(nil).Type will be a @ rule type, so we change to the type,
	// which will be the parent
	t, ok := ob.GetObject(nil).Type.ChangeToType().GetType(ctx)
	if !ok {
		panic(kerr.New("KYCTDXKFYR", "GetType: type %v not found", ob.GetObject(nil).Type.ChangeToType().Value()).Error())
	}

	return &RuleWrapper{Ctx: ctx, Interface: r, Parent: t, Struct: r.GetRule(nil)}
}

func (r *RuleWrapper) IsCollection() bool {
	return r.Parent.IsNativeCollection()
}

// ItemsRule returns Items rule for a collection Rule.
func (r *RuleWrapper) ItemsRule() (*RuleWrapper, error) {
	if !r.IsCollection() {
		return nil, kerr.New("VPAGXSTQHM", "%s is not a collection", r.Parent.Id.Value())
	}
	// I don't think this should be here:
	/*
		if r.Parent.Alias != nil {
			aw, err := WrapRule(r.Ctx, r.Parent.Alias)
			if err != nil {
				return nil, kerr.Wrap("PVCNTDVGWA", err)
			}
			if aw.IsCollection() {
				ir, err := aw.ItemsRule()
				if err != nil {
					return nil, kerr.Wrap("UIGQFXJLJE", err)
				}
				return ir, nil
			}
		}
	*/
	c, ok := r.Interface.(CollectionRule)
	if !ok {
		return nil, kerr.New("TNRVQVJIFH", "%T is not a CollectionRule", r.Interface)
	}
	ir := c.GetItemsRule()
	if ir == nil {
		return nil, kerr.New("SUJLYBXPYS", "%s has nil items rule", r.Parent.Id.Value())
	}
	w := WrapRule(r.Ctx, ir)
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
