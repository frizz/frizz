package system

import (
	"context"
	"encoding/json"

	"github.com/davelondon/kerr"
	"kego.io/context/jsonctx"
	"kego.io/packer"
)

func Unmarshal(ctx context.Context, in []byte, out interface{}) error {

	var i interface{}
	if err := json.Unmarshal(in, &i); err != nil {
		return kerr.Wrap("PDTPGAYXRX", err)
	}

	p := packer.Pack(i)

	if err := Unpack(ctx, p, out); err != nil {
		return kerr.Wrap("BCTGSARAXH", err)
	}

	return nil
}

func Unpack(ctx context.Context, p packer.Packed, out interface{}) error {

	if up, ok := out.(packer.Unpacker); ok {
		if err := up.Unpack(ctx, p, false); err != nil {
			return kerr.Wrap("YLILPEEMNJ", err)
		}
		return nil
	}

	ptri, ok := out.(*interface{})
	if !ok {
		return kerr.New("VFLOLSLCUQ", "If unmarshaling into an unknown type, must pass in a *interface{}. Found: %T", out)
	}
	i, err := UnpackUnknownType(ctx, p, true)
	if err != nil {
		return kerr.Wrap("RIOAIQLPIG", err)
	}
	*ptri = i

	return nil
}

func UnpackUnknownType(ctx context.Context, in packer.Packed, iface bool) (interface{}, error) {
	if in.Type() != packer.J_MAP {
		return nil, kerr.New("YVFHXVKPSG", "Unpacking an unknown type, so input must be a map. Found %s", in.Type())
	}
	i, err := GetNewFromTypeField(ctx, in)
	if err != nil {
		return nil, kerr.Wrap("SJFYDULNOS", err)
	}
	un, ok := i.(packer.Unpacker)
	if !ok {
		return nil, kerr.New("FOLMDMHTVL", "%T does not implement packer.Unpacker", i)
	}
	if err := un.Unpack(ctx, in, iface); err != nil {
		return nil, kerr.Wrap("BNHVLQPFKC", err)
	}
	return i, nil
}

func GetNewFromTypeField(ctx context.Context, in packer.Packed) (interface{}, error) {
	t, ok := in.Map()["type"]
	if !ok {
		return nil, kerr.New("RXEPCCGFKV", "Type field not found.")
	}
	tr := new(Reference)
	if err := tr.Unpack(ctx, t, false); err != nil {
		return nil, kerr.Wrap("MVOYONPYER", err)
	}
	jctx := jsonctx.FromContext(ctx)
	nf, ok := jctx.GetNewFunc(tr.Package, tr.Name)
	if !ok {
		return nil, kerr.New("BQIJMVGJDT", "NewFunc %s not found.", tr.String())
	}
	return nf(), nil
}

func init() {
	jpkg := jsonctx.InitPackage("kego.io/json", 0)
	jpkg.InitNew("string", nil, func() interface{} { return new(JsonStringRule) })
	jpkg.InitNew("number", nil, func() interface{} { return new(JsonNumberRule) })
	jpkg.InitNew("bool", nil, func() interface{} { return new(JsonBoolRule) })
}

func (v *JsonStringRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
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

func (v *JsonNumberRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
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

func (v *JsonBoolRule) Unpack(ctx context.Context, in packer.Packed, iface bool) error {

	if in == nil || in.Type() == packer.J_NULL {
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

func UnpackNumber(ctx context.Context, in packer.Packed) (float64, error) {
	if in == nil || in.Type() == packer.J_NULL {
		return 0.0, nil
	}
	if in.Type() != packer.J_NUMBER {
		return 0.0, kerr.New("PFHJQLAIFP", "UnpackNumber: %s must by J_NUMBER", in.Type())
	}
	return in.Number(), nil
}
