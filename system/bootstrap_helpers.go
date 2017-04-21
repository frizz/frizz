package system

import (
	"context"
	"encoding/json"

	"reflect"

	"frizz.io/context/envctx"
	"frizz.io/context/jsonctx"
	"github.com/dave/kerr"
)

func NewContext(ctx context.Context, path string, aliases map[string]string) context.Context {
	ctx = envctx.NewContext(ctx, &envctx.Env{Path: path, Aliases: aliases})
	ctx = jsonctx.AutoContext(ctx)
	return ctx
}

func Marshal(ctx context.Context, data interface{}) ([]byte, error) {

	p, err := Repack(ctx, data)
	if err != nil {
		return nil, kerr.Wrap("OMFSDKIBCY", err)
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, kerr.Wrap("YCEBRGOCGO", err)
	}

	return b, nil
}

func Unmarshal(ctx context.Context, in []byte, out interface{}) error {

	var i interface{}
	if err := json.Unmarshal(in, &i); err != nil {
		return kerr.Wrap("PDTPGAYXRX", err)
	}

	p := Pack(i)

	if err := Unpack(ctx, p, out); err != nil {
		return kerr.Wrap("BCTGSARAXH", err)
	}

	return nil
}

func UnpackRefelctValue(ctx context.Context, p Packed, val reflect.Value) error {

	if up, ok := val.Interface().(Unpacker); ok {
		if err := up.Unpack(ctx, p, false); err != nil {
			return kerr.Wrap("FGSFQHHMNE", err)
		}
		return nil
	}

	if val.Type() == reflect.TypeOf("") {
		unpacked, err := UnpackString(ctx, p)
		if err != nil {
			return kerr.Wrap("XDGRKOVCHD", err)
		}
		val.Set(reflect.ValueOf(unpacked))
		return nil
	}

	if val.Type() == reflect.TypeOf(0.0) {
		unpacked, err := UnpackNumber(ctx, p)
		if err != nil {
			return kerr.Wrap("IPYJYKLNBE", err)
		}
		val.Set(reflect.ValueOf(unpacked))
		return nil
	}

	if val.Type() == reflect.TypeOf(false) {
		unpacked, err := UnpackBool(ctx, p)
		if err != nil {
			return kerr.Wrap("RBILDURHWV", err)
		}
		val.Set(reflect.ValueOf(unpacked))
		return nil
	}

	if val.Type().Kind() == reflect.Ptr && val.Type().Elem().Kind() == reflect.Interface {
		i, err := UnpackUnknownType(ctx, p, true, "", "")
		if err != nil {
			return kerr.Wrap("TDEFTWSBFD", err)
		}
		val.Set(reflect.ValueOf(i))
	} else {
		return kerr.New("IFCTNULWUA", "If unmarshaling into an unknown type, must pass in a *interface{}. Found: %v", val.Type())
	}

	return nil
}

func Repack(ctx context.Context, data interface{}) (interface{}, error) {

	if re, ok := data.(Repacker); ok {
		p, pkg, name, jt, err := re.Repack(ctx)
		if err != nil {
			return nil, kerr.Wrap("ISAOPMMEYI", err)
		}
		if jt == J_OBJECT {
			return p, nil
		}
		typRef := NewReference(pkg, name)
		typeVal, err := typRef.ValueContext(ctx)
		if err != nil {
			return nil, kerr.Wrap("IWSCKOWIMW", err)
		}
		out := map[string]interface{}{}
		out["type"] = typeVal
		out["value"] = p
		return out, nil
	}

	switch data.(type) {
	case string, float64, bool:
		return data, nil
	}

	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Slice:
		out := []interface{}{}
		for i := 0; i < val.Len(); i++ {
			c, err := Repack(ctx, val.Index(i).Interface())
			if err != nil {
				return nil, kerr.Wrap("KSHOKVIBRS", err)
			}
			out = append(out, c)
		}
		return out, nil
	case reflect.Map:
		out := map[string]interface{}{}
		for _, key := range val.MapKeys() {
			c, err := Repack(ctx, val.MapIndex(key).Interface())
			if err != nil {
				return nil, kerr.Wrap("SOLGDVLWSK", err)
			}
			out[key.Interface().(string)] = c
		}
		return out, nil
	}

	return nil, kerr.New("POKFWVHEDY", "%T does can't be repacked", data)
}

func Unpack(ctx context.Context, p Packed, out interface{}) error {

	if up, ok := out.(Unpacker); ok {
		if err := up.Unpack(ctx, p, false); err != nil {
			return kerr.Wrap("YLILPEEMNJ", err)
		}
		return nil
	}

	ptri, ok := out.(*interface{})
	if !ok {
		return kerr.New("VFLOLSLCUQ", "If unmarshaling into an unknown type, must pass in a *interface{}. Found: %T", out)
	}
	i, err := UnpackUnknownType(ctx, p, true, "", "")
	if err != nil {
		return kerr.Wrap("RIOAIQLPIG", err)
	}
	*ptri = i

	return nil
}

func UnpackUnknownType(ctx context.Context, in Packed, iface bool, ifacePackage, ifaceName string) (interface{}, error) {
	if in.Type() != J_MAP {
		return nil, kerr.New("YVFHXVKPSG", "Unpacking an unknown type, so input must be a map. Found %s", in.Type())
	}
	nf, df, err := GetNewFromTypeField(ctx, in, iface, ifacePackage, ifaceName)
	if err != nil {
		return nil, kerr.Wrap("SJFYDULNOS", err)
	}
	i := nf()
	un, ok := i.(Unpacker)
	if !ok {
		return nil, kerr.New("FOLMDMHTVL", "%T does not implement Unpacker", i)
	}
	if err := un.Unpack(ctx, in, iface); err != nil {
		return nil, kerr.Wrap("BNHVLQPFKC", err)
	}
	if df != nil {
		i = df(i)
	}
	return i, nil
}

func GetNewFromTypeField(ctx context.Context, in Packed, iface bool, ifacePackage, ifaceName string) (newFunc func() interface{}, derefFunc func(interface{}) interface{}, err error) {
	t, ok := in.Map()["type"]
	if !ok {
		return nil, nil, kerr.New("RXEPCCGFKV", "Type field not found.")
	}
	tr := new(Reference)
	if err := tr.Unpack(ctx, t, false); err != nil {
		return nil, nil, kerr.Wrap("MVOYONPYER", err)
	}
	jctx := jsonctx.FromContext(ctx)
	nf, df, ok := jctx.GetNewFunc(tr.Package, tr.Name)
	if ok {
		return nf, df, nil
	}
	if iface && ifacePackage != "" && ifaceName != "" {
		r := NewReference(ifacePackage, ifaceName)
		nf, ok := jsonctx.FromContext(ctx).GetDummyFunc(r.Package, r.Name)
		if ok {
			// TODO: fix this
			return nf, nil, nil
		}
	}
	return nil, nil, kerr.New("BQIJMVGJDT", "NewFunc %s not found.", tr.String())
}

func init() {
	jpkg := jsonctx.InitPackage("frizz.io/json")
	jpkg.Init(
		"string",
		func() interface{} { return "" },
		nil,
		func() interface{} { return new(JsonStringRule) },
		nil,
	)
	jpkg.Init(
		"number",
		func() interface{} { return 0.0 },
		nil,
		func() interface{} { return new(JsonNumberRule) },
		nil,
	)
	jpkg.Init(
		"bool",
		func() interface{} { return false },
		nil,
		func() interface{} { return new(JsonBoolRule) },
		nil,
	)
}

func (v *JsonStringRule) Unpack(ctx context.Context, in Packed, iface bool) error {

	if in == nil || in.Type() == J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("BEIWFULUTD", err)
	}

	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("KRTUNQFXHA", err)
	}

	return nil

}

func (v *JsonNumberRule) Unpack(ctx context.Context, in Packed, iface bool) error {

	if in == nil || in.Type() == J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("WXAFWTOKHU", err)
	}

	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("PSSXFMIYLS", err)
	}

	return nil

}

func (v *JsonBoolRule) Unpack(ctx context.Context, in Packed, iface bool) error {

	if in == nil || in.Type() == J_NULL {
		return nil
	}

	if v.Object == nil {
		v.Object = new(Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("UQUKPBNRVJ", err)
	}

	if v.Rule == nil {
		v.Rule = new(Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return kerr.Wrap("RSAVQRLPML", err)
	}

	return nil

}

func UnpackBool(ctx context.Context, in Packed) (bool, error) {
	if in == nil || in.Type() == J_NULL {
		return false, nil
	}
	if in.Type() != J_BOOL {
		return false, kerr.New("ITLWMRUKKD", "UnpackBool: %s must by J_BOOL", in.Type())
	}
	return in.Bool(), nil
}

func UnpackString(ctx context.Context, in Packed) (string, error) {
	if in == nil || in.Type() == J_NULL {
		return "", nil
	}
	if in.Type() != J_STRING {
		return "", kerr.New("CIECONONEF", "UnpackString: %s must by J_STRING", in.Type())
	}
	return in.String(), nil
}

func UnpackNumber(ctx context.Context, in Packed) (float64, error) {
	if in == nil || in.Type() == J_NULL {
		return 0.0, nil
	}
	if in.Type() != J_NUMBER {
		return 0.0, kerr.New("PFHJQLAIFP", "UnpackNumber: %s must by J_NUMBER", in.Type())
	}
	return in.Number(), nil
}

// ShouldUseExplicitTypeNotation determines if a value should be packed
// into an interface using the explicit typed notation:
// {"type":"foo","value":"bar"}
//
// Rules:
// value == map ALWAYS
// value == object NEVER
// value == native/array USUALLY, but not if the type is the native
//          type of the interface. This will only apply to rule interfaces.
//
func ShouldUseExplicitTypeNotation(dataTypePackage string, dataTypeName string, dataJsonType JsonType, interfaceTypePackage string, interfaceTypeName string) bool {
	switch dataJsonType {
	case J_OBJECT:
		return false
	case J_MAP:
		return true
	case J_STRING, J_BOOL, J_NUMBER, J_ARRAY:
		return dataTypePackage != interfaceTypePackage ||
			dataTypeName != interfaceTypeName
	default: // J_NULL
		return true // can we get here?
	}
}
