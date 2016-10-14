package system

import (
	"context"

	"github.com/davelondon/kerr"
	"kego.io/packer"
)

/*
func init() {
	pkg := jsonctx.InitPackage("kego.io/system", 0)
	pkg.InitNew("array", nil, New_ArrayRule)
	pkg.InitNew("map", nil, New_MapRule)
	pkg.InitNew("object", New_Object, New_ObjectRule)
	pkg.InitNew("rule", New_Rule, New_RuleRule)
	pkg.InitNew("string", New_String, New_StringRule)
	pkg.InitNew("int", New_Int, New_IntRule)
	pkg.InitNew("bool", New_Bool, New_BoolRule)
	pkg.InitNew("number", New_Number, New_NumberRule)
	pkg.InitNew("reference", New_Reference, New_ReferenceRule)
	pkg.InitNew("package", New_Package, New_PackageRule)
	pkg.InitNew("tags", New_Tags, New_TagsRule)
	pkg.InitNew("type", New_Type, New_TypeRule)
}*/

/*
func New_String(ctx context.Context) interface{} {
	v := new(String)
	return v
}

func New_Bool(ctx context.Context) interface{} {
	v := new(Bool)
	return v
}

func New_Number(ctx context.Context) interface{} {
	v := new(Number)
	return v
}

func New_Reference(ctx context.Context) interface{} {
	v := new(Reference)
	return v
}

func New_Package(ctx context.Context) interface{} {
	v := new(Package)
	v.Object = new(Object)
	return v
}

func New_Tags(ctx context.Context) interface{} {
	v := new(Tags)
	return v
}

func New_Type(ctx context.Context) interface{} {
	v := new(Type)
	v.Object = new(Object)
	return v
}

func New_Object(ctx context.Context) interface{} {
	v := new(Object)
	return v
}

func New_Int(ctx context.Context) interface{} {
	v := new(Int)
	return v
}

func New_Rule(ctx context.Context) interface{} {
	v := new(Rule)
	return v
}
*/

/*
func UnpackRuleInterface(ctx context.Context, in packer.Packed) (RuleInterface, error) {
	i, err := UnpackUnknownType(ctx, in, true)
	if err != nil {
		return nil, kerr.Wrap("LTNSFRGNRT", err)
	}
	ob, ok := i.(RuleInterface)
	if !ok {
		return nil, kerr.New("GCAKXDFTKO", "%T does not implement system.RuleInterface", i)
	}
	return ob, nil
}

func UnpackObjectInterface(ctx context.Context, in packer.Packed) (ObjectInterface, error) {
	switch in.Type() {
	case packer.J_MAP:
		i, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, kerr.Wrap("WNOJQCIETR", err)
		}
		ob, ok := i.(ObjectInterface)
		if !ok {
			return nil, kerr.New("AFQCJDSUOV", "%T does not implement system.ObjectInterface", i)
		}
		return ob, nil
	default:
		return nil, kerr.New("DKCJPKINBA", "Unpacking into an ObjectInterface, so input must be a map. Found %s", in.Type())
	}
}

func UnpackStringInterface(ctx context.Context, in packer.Packed) (StringInterface, error) {
	// String is a native type, so we accept a string or a typed map
	switch in.Type() {
	case packer.J_STRING:
		ob := new(String)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, kerr.Wrap("SCHSENXXQD", err)
		}
		return ob, nil
	case packer.J_MAP:
		ob, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, kerr.Wrap("NHTCXPAVQD", err)
		}
		i, ok := ob.(StringInterface)
		if !ok {
			return nil, kerr.New("OSCYJRJGTT", "%T does not implement system.StringInterface", ob)
		}
		return i, nil
	default:
		return nil, kerr.New("DKCJPKINBA", "Unpacking into a StringInterface, so input must be a string or a map. Found %s", in.Type())
	}
}

func UnpackTagsInterface(ctx context.Context, in packer.Packed) (TagsInterface, error) {
	// Tags is an array type, so we accept an array or a typed map
	switch in.Type() {
	case packer.J_ARRAY:
		ob := new(Tags)
		if err := ob.Unpack(ctx, in, false); err != nil {
			return nil, kerr.Wrap("INSRJQTFEP", err)
		}
		return ob, nil
	case packer.J_MAP:
		ob, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, kerr.Wrap("NEJGMXXRAK", err)
		}
		i, ok := ob.(TagsInterface)
		if !ok {
			return nil, kerr.New("WXWPSQJIVA", "%T does not implement system.TagsInterface", ob)
		}
		return i, nil
	default:
		return nil, kerr.New("FYGBHLNPBY", "Unpacking into a TagsInterface, so input must be an array or a map. Found %s", in.Type())
	}
}
*/

type Mappy map[string]string

type MappyInterface interface {
	GetMappy(context.Context) error
}

func UnpackMappyInterface(ctx context.Context, in packer.Packed) (MappyInterface, error) {
	// Mappy is an map type, so we only accept a typed map
	switch in.Type() {
	case packer.J_MAP:
		ob, err := UnpackUnknownType(ctx, in, true)
		if err != nil {
			return nil, kerr.Wrap("QSTEIBUNWO", err)
		}
		i, ok := ob.(MappyInterface)
		if !ok {
			return nil, kerr.New("SMJXFKMKSP", "%T does not implement system.MappyInterface", ob)
		}
		return i, nil
	default:
		return nil, kerr.New("VYHBOJFSIM", "Unpacking into a MappyInterface, so input must be a map. Found %s", in.Type())
	}
}

func (v *Mappy) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if iface {
		if in.Type() != packer.J_MAP {
			return kerr.New("DTLQPHOJDT", "Mappy.Unpack: %s must by J_MAP", in.Type())
		}
		in = in.Map()["value"]
	}

	if in.Type() != packer.J_MAP {
		return kerr.New("KYPSXYBLNC", "Mappy.Unpack: %s must by J_MAP", in.Type())
	}

	for key, value := range in.Map() {
		ob, err := UnpackString(ctx, value)
		if err != nil {
			return kerr.Wrap("HOLNALWYBA", err)
		}
		(*v)[key] = ob
	}

	return nil
}

/*
func New_ArrayRule(ctx context.Context) interface{} {
	v := new(ArrayRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_MapRule(ctx context.Context) interface{} {
	v := new(MapRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_StringRule(ctx context.Context) interface{} {
	v := new(StringRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_IntRule(ctx context.Context) interface{} {
	v := new(IntRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_BoolRule(ctx context.Context) interface{} {
	v := new(BoolRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_NumberRule(ctx context.Context) interface{} {
	v := new(NumberRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_ReferenceRule(ctx context.Context) interface{} {
	v := new(ReferenceRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_RuleRule(ctx context.Context) interface{} {
	v := new(RuleRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_ObjectRule(ctx context.Context) interface{} {
	v := new(ObjectRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_PackageRule(ctx context.Context) interface{} {
	v := new(PackageRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_TagsRule(ctx context.Context) interface{} {
	v := new(TagsRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}

func New_TypeRule(ctx context.Context) interface{} {
	v := new(TypeRule)
	v.Object = new(Object)
	v.Rule = new(Rule)
	return v
}
*/
