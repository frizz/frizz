package system

import (
	"context"

	"github.com/davelondon/kerr"
	"kego.io/context/jsonctx"
	"kego.io/packer"
)

func UnpackBool(ctx context.Context, in packer.Packed) (bool, error) {
	if in == nil || in.Type() == packer.J_NULL {
		return false, nil
	}
	if in.Type() != packer.J_BOOL {
		return false, kerr.New("ITLWMRUKKD", "UnpackBool: %s must by J_BOOL", in.Type())
	}
	return in.Bool(), nil
}

func UnpackString(ctx context.Context, in packer.Packed) (string, error) {
	if in == nil || in.Type() == packer.J_NULL {
		return "", nil
	}
	if in.Type() != packer.J_STRING {
		return "", kerr.New("CIECONONEF", "UnpackString: %s must by J_STRING", in.Type())
	}
	return in.String(), nil
}

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
}

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
	v.Object = New_Object(ctx).(*Object)
	return v
}

func New_Tags(ctx context.Context) interface{} {
	v := new(Tags)
	return v
}

func New_Type(ctx context.Context) interface{} {
	v := new(Type)
	v.Object = New_Object(ctx).(*Object)
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
	i, err := UnpackUnknownType(ctx, in, true)
	if err != nil {
		return nil, kerr.Wrap("WNOJQCIETR", err)
	}
	ob, ok := i.(ObjectInterface)
	if !ok {
		return nil, kerr.New("AFQCJDSUOV", "%T does not implement system.ObjectInterface", i)
	}
	return ob, nil
}

func UnpackStringInterface(ctx context.Context, in packer.Packed) (StringInterface, error) {
	// String is a native type, so we accept a string or a typed map
	switch in.Type() {
	case packer.J_STRING:
		ob := New_String(ctx).(*String)
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
		ob := New_Tags(ctx).(*Tags)
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

func (v *ArrayRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("MSXPOREBLD", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("UYMXFFSCLT", err)
	}

	if field, ok := in.Map()["items"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return kerr.Wrap("TWSEEFJMCO", err)
		}
		v.Items = ob
	}

	if field, ok := in.Map()["max-items"]; ok && field.Type() != packer.J_NULL {
		if v.MaxItems == nil {
			v.MaxItems = New_Int(ctx).(*Int)
		}
		if err := v.MaxItems.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("TQJPOSFUIH", err)
		}
	}

	if field, ok := in.Map()["min-items"]; ok && field.Type() != packer.J_NULL {
		if v.MinItems == nil {
			v.MinItems = New_Int(ctx).(*Int)
		}
		if err := v.MinItems.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("TFIOGQJNVT", err)
		}
	}

	if field, ok := in.Map()["unique-items"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return kerr.Wrap("RAQIXLAFDY", err)
		}
		v.UniqueItems = ob
	}

	return nil

}
func (v *MapRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("IIVXAMEJOL", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("XEIFXOGUMH", err)
	}

	if field, ok := in.Map()["items"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return kerr.Wrap("TQUGWPSOXQ", err)
		}
		v.Items = ob
	}

	if field, ok := in.Map()["keys"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return kerr.Wrap("OCLRYFYHYQ", err)
		}
		v.Keys = ob
	}

	if field, ok := in.Map()["max-items"]; ok && field.Type() != packer.J_NULL {
		if v.MaxItems == nil {
			v.MaxItems = New_Int(ctx).(*Int)
		}
		if err := v.MaxItems.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("TRSCHOGQAD", err)
		}
	}

	if field, ok := in.Map()["min-items"]; ok && field.Type() != packer.J_NULL {
		if v.MinItems == nil {
			v.MinItems = New_Int(ctx).(*Int)
		}
		if err := v.MinItems.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("ABWCWEXHAL", err)
		}
	}

	return nil

}
func (v *StringRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("LAUWEPOSIY", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("YBREONMQCU", err)
	}

	if field, ok := in.Map()["default"]; ok && field.Type() != packer.J_NULL {
		if v.Default == nil {
			v.Default = New_String(ctx).(*String)
		}
		if err := v.Default.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("LUNFAPKKON", err)
		}
	}

	if field, ok := in.Map()["enum"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_ARRAY {
			return kerr.New("BOANVNIBJI", "StringRule: enum field must be an array. Found %s", field.Type())
		}
		for _, value := range field.Array() {
			ob, err := UnpackString(ctx, value)
			if err != nil {
				return kerr.Wrap("QNPYCXSSAK", err)
			}
			v.Enum = append(v.Enum, ob)
		}
	}

	if field, ok := in.Map()["equal"]; ok && field.Type() != packer.J_NULL {
		if v.Equal == nil {
			v.Equal = New_String(ctx).(*String)
		}
		if err := v.Equal.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("CIAKTABAYB", err)
		}
	}

	if field, ok := in.Map()["format"]; ok && field.Type() != packer.J_NULL {
		if v.Format == nil {
			v.Format = New_String(ctx).(*String)
		}
		if err := v.Format.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("QOOGWLNAAJ", err)
		}
	}

	if field, ok := in.Map()["long"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return kerr.Wrap("RQBUBSJYHN", err)
		}
		v.Long = ob
	}

	if field, ok := in.Map()["max-length"]; ok && field.Type() != packer.J_NULL {
		if v.MaxLength == nil {
			v.MaxLength = New_Int(ctx).(*Int)
		}
		if err := v.MaxLength.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("AFEWIKYHSJ", err)
		}
	}

	if field, ok := in.Map()["min-length"]; ok && field.Type() != packer.J_NULL {
		if v.MinLength == nil {
			v.MinLength = New_Int(ctx).(*Int)
		}
		if err := v.MinLength.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("EXXDUXGCLU", err)
		}
	}

	if field, ok := in.Map()["pattern"]; ok && field.Type() != packer.J_NULL {
		if v.Pattern == nil {
			v.Pattern = New_String(ctx).(*String)
		}
		if err := v.Pattern.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("GJGNOKJMBH", err)
		}
	}

	if field, ok := in.Map()["pattern-not"]; ok && field.Type() != packer.J_NULL {
		if v.PatternNot == nil {
			v.PatternNot = New_String(ctx).(*String)
		}
		if err := v.PatternNot.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("VSLVEJQQEC", err)
		}
	}

	return nil

}
func (v *IntRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("HTVEIXVFDS", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("VBUBGQYJET", err)
	}

	if field, ok := in.Map()["default"]; ok && field.Type() != packer.J_NULL {
		if v.Default == nil {
			v.Default = New_Int(ctx).(*Int)
		}
		if err := v.Default.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("DKOIHHYGYH", err)
		}
	}

	if field, ok := in.Map()["maximum"]; ok && field.Type() != packer.J_NULL {
		if v.Maximum == nil {
			v.Maximum = New_Int(ctx).(*Int)
		}
		if err := v.Maximum.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("OSVDVJMQQB", err)
		}
	}

	if field, ok := in.Map()["minimum"]; ok && field.Type() != packer.J_NULL {
		if v.Minimum == nil {
			v.Minimum = New_Int(ctx).(*Int)
		}
		if err := v.Minimum.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("QFNCATUYYT", err)
		}
	}

	if field, ok := in.Map()["multiple-of"]; ok && field.Type() != packer.J_NULL {
		if v.MultipleOf == nil {
			v.MultipleOf = New_Int(ctx).(*Int)
		}
		if err := v.MultipleOf.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("BXPUOFWYYW", err)
		}
	}

	return nil

}
func (v *BoolRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("OSAXPFCHUW", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("WUUAPISXEE", err)
	}

	if field, ok := in.Map()["default"]; ok && field.Type() != packer.J_NULL {
		if v.Default == nil {
			v.Default = New_Bool(ctx).(*Bool)
		}
		if err := v.Default.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("IPMYFWJEWK", err)
		}
	}

	return nil

}
func (v *NumberRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("NQHFFYTGTN", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("KFNRXEFXTS", err)
	}

	if field, ok := in.Map()["default"]; ok && field.Type() != packer.J_NULL {
		if v.Default == nil {
			v.Default = New_Number(ctx).(*Number)
		}
		if err := v.Default.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("MPASFHUMXW", err)
		}
	}

	if field, ok := in.Map()["exclusive-maximum"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return kerr.Wrap("VKISBSJFFD", err)
		}
		v.ExclusiveMaximum = ob
	}

	if field, ok := in.Map()["exclusive-minimum"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return kerr.Wrap("QIDUBWWJAD", err)
		}
		v.ExclusiveMinimum = ob
	}

	if field, ok := in.Map()["maximum"]; ok && field.Type() != packer.J_NULL {
		if v.Maximum == nil {
			v.Maximum = New_Number(ctx).(*Number)
		}
		if err := v.Maximum.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("UXPKORGOFA", err)
		}
	}

	if field, ok := in.Map()["minimum"]; ok && field.Type() != packer.J_NULL {
		if v.Minimum == nil {
			v.Minimum = New_Number(ctx).(*Number)
		}
		if err := v.Minimum.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("AYQBRLKPVF", err)
		}
	}

	if field, ok := in.Map()["multiple-of"]; ok && field.Type() != packer.J_NULL {
		if v.MultipleOf == nil {
			v.MultipleOf = New_Number(ctx).(*Number)
		}
		if err := v.MultipleOf.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("NKQKFFFJUX", err)
		}
	}

	return nil

}
func (v *ReferenceRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("SLLCQCARUM", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("WNACMHDQXN", err)
	}

	if field, ok := in.Map()["default"]; ok && field.Type() != packer.J_NULL {
		if v.Default == nil {
			v.Default = New_Reference(ctx).(*Reference)
		}
		if err := v.Default.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("EBHOIXSQHK", err)
		}
	}

	if field, ok := in.Map()["pattern"]; ok && field.Type() != packer.J_NULL {
		if v.Pattern == nil {
			v.Pattern = New_String(ctx).(*String)
		}
		if err := v.Pattern.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("GJNSTJMVXW", err)
		}
	}

	if field, ok := in.Map()["pattern-not"]; ok && field.Type() != packer.J_NULL {
		if v.PatternNot == nil {
			v.PatternNot = New_String(ctx).(*String)
		}
		if err := v.PatternNot.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("KAKUFMGIQW", err)
		}
	}

	return nil

}

func (v *PackageRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("BEIWFULUTD", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("KRTUNQFXHA", err)
	}

	return nil

}

func (v *RuleRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("KOXFJJVRLY", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("MEOERNKOAJ", err)
	}

	return nil

}

func (v *ObjectRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("RLHVYNBHJW", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("HSENNIBDLG", err)
	}

	return nil

}

func (v *TagsRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("MDEXNPEJWF", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("QYUBYTNSIS", err)
	}

	return nil

}
func (v *TypeRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("YOPEOVEEXS", err)
	}

	if v.Rule == nil {
		v.Rule = New_Rule(ctx).(*Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("JCFNRNXDOM", err)
	}

	return nil

}

func (v *Tags) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if in.Type() == packer.J_MAP {
		in = in.Map()["value"]
	}

	if in.Type() != packer.J_ARRAY {
		return kerr.New("JCPPNLLVRA", "Tags.Unpack: %s must by J_ARRAY", in.Type())
	}

	for _, value := range in.Array() {
		ob, err := UnpackString(ctx, value)
		if err != nil {
			return kerr.Wrap("HOLNALWYBA", err)
		}
		*v = append(*v, ob)
	}

	return nil

}
func (v *Number) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if in.Type() == packer.J_MAP {
		in = in.Map()["value"]
	}

	if in.Type() != packer.J_NUMBER {
		return kerr.New("QCFGULWGNV", "Number.Unpack: %s must by J_NUMBER", in.Type())
	}

	*v = Number(in.Number())

	return nil
}

func (v *Bool) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if in.Type() == packer.J_MAP {
		in = in.Map()["value"]
	}

	if in.Type() != packer.J_BOOL {
		return kerr.New("JKDSOEAHLG", "Bool.Unpack: %s must by J_BOOL", in.Type())
	}

	*v = Bool(in.Bool())

	return nil
}

func (v *String) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if in.Type() == packer.J_MAP {
		in = in.Map()["value"]
	}

	if in.Type() != packer.J_STRING {
		return kerr.New("LXNBXXRGLS", "String.Unpack: %s must by J_STRING", in.Type())
	}

	*v = String(in.String())

	return nil
}

func (v *Rule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if field, ok := in.Map()["interface"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Interface = ob
	}

	if field, ok := in.Map()["optional"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Optional = ob
	}

	if field, ok := in.Map()["selector"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackString(ctx, field)
		if err != nil {
			return err
		}
		v.Selector = ob
	}

	return nil
}

func (v *Object) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if field, ok := in.Map()["description"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackString(ctx, field)
		if err != nil {
			return err
		}
		v.Description = ob
	}

	if field, ok := in.Map()["id"]; ok && field.Type() != packer.J_NULL {
		if v.Id == nil {
			v.Id = New_Reference(ctx).(*Reference)
		}
		if err := v.Id.Unpack(ctx, field, false); err != nil {
			return err
		}
	}

	if field, ok := in.Map()["rules"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_ARRAY {
			return kerr.New("LOIJEOTTWE", "Object: rules field must be an array. Found %s", field.Type())
		}
		for _, value := range field.Array() {
			ob, err := UnpackRuleInterface(ctx, value)
			if err != nil {
				return kerr.Wrap("LIATJWMGXM", err)
			}
			v.Rules = append(v.Rules, ob)
		}
	}

	if field, ok := in.Map()["tags"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_ARRAY {
			return kerr.New("SRGRWWKETJ", "Object: tags field must be an array. Found %s", field.Type())
		}
		for _, value := range field.Array() {
			ob, err := UnpackString(ctx, value)
			if err != nil {
				return kerr.Wrap("STFMYAQUUS", err)
			}
			v.Tags = append(v.Tags, ob)
		}
	}

	if field, ok := in.Map()["tags-new"]; ok && field.Type() != packer.J_NULL {
		if v.TagsNew == nil {
			v.TagsNew = New_Tags(ctx).(Tags)
		}
		if err := v.TagsNew.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("QCYFMPVOSS", err)
		}
	}

	if field, ok := in.Map()["type"]; ok && field.Type() != packer.J_NULL {
		if v.Type == nil {
			v.Type = New_Reference(ctx).(*Reference)
		}
		if err := v.Type.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("MXDNBBSBRJ", err)
		}
	}

	return nil

}

func (v *Package) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("JILILQONWW", err)
	}

	if field, ok := in.Map()["aliases"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_MAP {
			return kerr.New("DTJLNRDEDD", "Type: aliases field must be a map. Found %s", field.Type())
		}
		for key, value := range field.Map() {
			ob, err := UnpackString(ctx, value)
			if err != nil {
				return kerr.Wrap("CUGRUCHKTM", err)
			}
			if v.Aliases == nil {
				v.Aliases = make(map[string]string)
			}
			v.Aliases[key] = ob
		}
	}

	if field, ok := in.Map()["recursive"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return kerr.Wrap("JTPLHHWICO", err)
		}
		v.Recursive = ob
	}

	return nil

}

func (v *Type) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = New_Object(ctx).(*Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("MOPEPGLSEJ", err)
	}

	if field, ok := in.Map()["alias"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackRuleInterface(ctx, field)
		if err != nil {
			return kerr.Wrap("MJTTFDTSBC", err)
		}
		v.Alias = ob
	}

	if field, ok := in.Map()["basic"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return kerr.Wrap("MJTTFDTSBC", err)
		}
		v.Basic = ob
	}

	if field, ok := in.Map()["custom"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return kerr.Wrap("AAQKLHIWPC", err)
		}
		v.Custom = ob
	}

	if field, ok := in.Map()["custom-kind"]; ok && field.Type() != packer.J_NULL {
		if v.CustomKind == nil {
			v.CustomKind = New_String(ctx).(*String)
		}
		if err := v.CustomKind.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("FGDQPCAUYG", err)
		}
	}

	if field, ok := in.Map()["embed"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_ARRAY {
			return kerr.New("BOANVNIBJI", "Type: embed field must be an array. Found %s", field.Type())
		}
		for _, value := range field.Array() {
			ob := New_Reference(ctx).(*Reference)
			if err := ob.Unpack(ctx, value, false); err != nil {
				return kerr.Wrap("FVKBSKEPJV", err)
			}
			v.Embed = append(v.Embed, ob)
		}
	}

	if field, ok := in.Map()["fields"]; ok && field.Type() != packer.J_NULL {
		if field.Type() != packer.J_MAP {
			return kerr.New("DTJLNRDEDD", "Type: fields field must be a map. Found %s", field.Type())
		}
		for key, value := range field.Map() {
			ob, err := UnpackRuleInterface(ctx, value)
			if err != nil {
				return kerr.Wrap("FVKBSKEPJV", err)
			}
			if v.Fields == nil {
				v.Fields = make(map[string]RuleInterface)
			}
			v.Fields[key] = ob
		}
	}

	if field, ok := in.Map()["interface"]; ok && field.Type() != packer.J_NULL {
		ob, err := UnpackBool(ctx, field)
		if err != nil {
			return kerr.Wrap("YKHGPVBDPM", err)
		}
		v.Interface = ob
	}

	if field, ok := in.Map()["native"]; ok && field.Type() != packer.J_NULL {
		if v.Native == nil {
			v.Native = New_String(ctx).(*String)
		}
		if err := v.Native.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("CGWPEEMIHY", err)
		}
	} else {
		value := packer.Pack("object")
		if v.Native == nil {
			v.Native = New_String(ctx).(*String)
		}
		if err := v.Native.Unpack(ctx, value, false); err != nil {
			return kerr.Wrap("CGWPEEMIHY", err)
		}
	}

	if field, ok := in.Map()["rule"]; ok && field.Type() != packer.J_NULL {
		if v.Rule == nil {
			v.Rule = New_Type(ctx).(*Type)
		}
		if err := v.Rule.Unpack(ctx, field, false); err != nil {
			return kerr.Wrap("JLBATRRELG", err)
		}
	}

	return nil

}

func New_ArrayRule(ctx context.Context) interface{} {
	v := new(ArrayRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_MapRule(ctx context.Context) interface{} {
	v := new(MapRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_StringRule(ctx context.Context) interface{} {
	v := new(StringRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_IntRule(ctx context.Context) interface{} {
	v := new(IntRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_BoolRule(ctx context.Context) interface{} {
	v := new(BoolRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_NumberRule(ctx context.Context) interface{} {
	v := new(NumberRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_ReferenceRule(ctx context.Context) interface{} {
	v := new(ReferenceRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_RuleRule(ctx context.Context) interface{} {
	v := new(RuleRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_ObjectRule(ctx context.Context) interface{} {
	v := new(ObjectRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_PackageRule(ctx context.Context) interface{} {
	v := new(PackageRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_TagsRule(ctx context.Context) interface{} {
	v := new(TagsRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}

func New_TypeRule(ctx context.Context) interface{} {
	v := new(TypeRule)
	v.Object = New_Object(ctx).(*Object)
	v.Rule = New_Rule(ctx).(*Rule)
	return v
}
