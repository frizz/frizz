package json

import (
	"bytes"
	"encoding/base64"
	"reflect"
	"strconv"

	"fmt"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/editor/client/console"
	"kego.io/kerr"
)

type unpackStruct struct {
	unknownType    string // have we encountered an unknown type?
	unknownPackage string // have we encountered an unknown package?
}

func Unpack(ctx context.Context, in Packed, out *interface{}) error {

	if in.Type() != J_MAP {
		return kerr.New("XOOUKLGORQ", nil, "Type %s should be J_MAP", in.Type())
	}

	us := &unpackStruct{}
	typ, err := us.getTypeFromField(ctx, in, reflect.Value{})
	if err != nil {
		return kerr.New("NUHCPRKRXT", err, "getTypeFromField (global)")
	}
	//if typ == nil {
	//return kerr.New("GREMVFEUMH", nil, "Unknown global type %#v", in.Map()["type"])
	//}

	// we don't wrap the error in kerr because in can return special
	// error types
	return us.unpackFragment(ctx, in, out, typ, false)
}

func UnpackFragment(ctx context.Context, in Packed, out *interface{}, typ reflect.Type, debug bool) error {

	us := &unpackStruct{}
	// we don't wrap the error in kerr because in can return special
	// error types
	return us.unpackFragment(ctx, in, out, typ, debug)
}

func (us *unpackStruct) unpackFragment(ctx context.Context, in Packed, out *interface{}, typ reflect.Type, debug bool) error {

	var p reflect.Value
	if typ != nil {
		p = getEmptyValue(typ)
	}

	if debug {
		console.Log("typ.String()", typ.String())
		console.Log("in.String()", in.String())
	}

	if debug {
		console.Log("p", p.Type().String(), p.Interface())
		console.Log("in", in.Type(), in.String())
	}

	err := us.unpack(ctx, in, p)

	if debug {
		console.Log("p", p.Type().String(), p.Interface())
	}

	if err == nil || us.unknownPackage != "" || us.unknownType != "" {
		// Sometimes we want to tolerate UnknownPackageError, so we should still set v
		v := reflect.ValueOf(out)
		v.Elem().Set(p)

	}

	if us.unknownPackage != "" {
		return UnknownPackageError{us.unknownPackage}
	}
	if us.unknownType != "" {
		return UnknownTypeError{us.unknownType}
	}
	if err != nil {
		return kerr.New("BCVBRIKFJX", err, "unpack (fragment)")
	}

	if debug {
		console.Log("unpacked successfully...", fmt.Sprintf("%T %s", out, *out))
	}

	return nil
}

func (us *unpackStruct) unpack(ctx context.Context, in Packed, v reflect.Value) error {

	if !v.IsValid() {
		return nil
	}

	typ := J_NULL
	if in != nil {
		typ = in.Type()
	}

	switch typ {
	case J_MAP:
		if err := us.unpackObject(ctx, in, v); err != nil {
			return kerr.New("LMLUICBTBA", err, "unpackObject")
		}
	case J_ARRAY:
		if err := us.unpackArray(ctx, in, v); err != nil {
			return kerr.New("ITJMJWULKO", err, "unpackArray")
		}
	default:
		if err := us.unpackLiteral(ctx, in, v); err != nil {
			return kerr.New("BSTNWUKLYO", err, "unpackLiteral")
		}
	}
	return nil
}

func (us *unpackStruct) unpackLiteral(ctx context.Context, in Packed, v reflect.Value) error {

	wantptr := in == nil || in.Type() == J_NULL
	_, _, up, pv := indirect(v, wantptr, false, true)
	if up != nil {
		if err := up.Unpack(ctx, in); err != nil {
			return kerr.New("RYSLUEEOAW", err, "Unpack (plain)")
		}
		return nil
	}
	v = pv

	typ := J_NULL
	if in != nil {
		typ = in.Type()
	}

	switch typ {
	case J_NULL:
		switch v.Kind() {
		case reflect.Interface, reflect.Ptr:
			v.Set(reflect.Zero(v.Type()))
			// otherwise, ignore null for primitives/string
		}
	case J_BOOL:
		switch v.Kind() {
		default:
			return &UnmarshalTypeError{"bool", v.Type()}
		case reflect.Bool:
			v.SetBool(in.Bool())
		case reflect.Interface:
			if v.NumMethod() == 0 {
				v.Set(reflect.ValueOf(in.Bool()))
			} else {
				if err := setDefaultNativeValueUnpack(ctx, v, in); err != nil {
					return kerr.New("DLQUIGTSLP", err, "setDefaultNativeValueUnpack (bool)")
				}
			}
		}
	case J_STRING:
		switch v.Kind() {
		default:
			return &UnmarshalTypeError{"string", v.Type()}
		case reflect.Slice:
			if v.Type().Elem().Kind() != reflect.Uint8 {
				return &UnmarshalTypeError{"string", v.Type()}
			}
			b := make([]byte, base64.StdEncoding.DecodedLen(len(in.String())))
			n, err := base64.StdEncoding.Decode(b, []byte(in.String()))
			if err != nil {
				return kerr.New("OKMBMDOFNL", err, "base64.StdEncoding.Decode")
			}
			v.Set(reflect.ValueOf(b[0:n]))
		case reflect.String:
			v.SetString(in.String())
		case reflect.Interface:
			if v.NumMethod() == 0 {
				v.Set(reflect.ValueOf(in.String()))
			} else {
				if err := setDefaultNativeValueUnpack(ctx, v, in); err != nil {
					return kerr.New("PRTYFVSEFT", err, "setDefaultNativeValueUnpack (string)")
				}
			}
		}
	case J_NUMBER:
		s := strconv.FormatFloat(in.Number(), 'f', -1, 64)
		switch v.Kind() {
		default:
			if v.Kind() == reflect.String && v.Type() == numberType {
				v.SetString(s)
				break
			}
			return &UnmarshalTypeError{"number", v.Type()}
		case reflect.Interface:
			if v.NumMethod() != 0 {
				if err := setDefaultNativeValueUnpack(ctx, v, in); err != nil {
					return kerr.New("FYHICWEELI", err, "setDefaultNativeValueUnpack (number)")
				}
				break
			}
			v.Set(reflect.ValueOf(in.Number()))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(s, 10, 64)
			if err != nil || v.OverflowInt(n) {
				return &UnmarshalTypeError{"number " + s, v.Type()}
			}
			v.SetInt(n)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			n, err := strconv.ParseUint(s, 10, 64)
			if err != nil || v.OverflowUint(n) {
				return &UnmarshalTypeError{"number " + s, v.Type()}
			}
			v.SetUint(n)
		case reflect.Float32, reflect.Float64:
			if v.OverflowFloat(in.Number()) {
				return &UnmarshalTypeError{"number " + s, v.Type()}
			}
			v.SetFloat(in.Number())
		}
	}
	return nil
}

func setDefaultNativeValueUnpack(ctx context.Context, v reflect.Value, in Packed) error {
	t, ok := jsonctx.FromContext(ctx).GetTypeByInterface(v.Type())
	if !ok {
		return kerr.New("YSBBTCVOUU", nil, "No type found for %s", v.Type().Name())
	}
	var i interface{}
	err := UnpackFragment(ctx, in, &i, t, false)
	if err != nil {
		return kerr.New("BCVBRIKFJX", err, "unpack (fragment)")
	}
	v.Set(reflect.ValueOf(i))
	return nil
}

func (us *unpackStruct) unpackArray(ctx context.Context, in Packed, v reflect.Value) error {

	if in.Type() != J_ARRAY {
		return kerr.New("PVJMVPULMY", nil, "Type %s should be J_ARRAY", in.Type())
	}

	// Check for unpackers.
	_, _, up, pv := indirect(v, false, false, true)
	if up != nil {
		if err := up.Unpack(ctx, in); err != nil {
			return kerr.New("PQRNFAYAYQ", err, "Unpack (plain)")
		}
		return nil
	}
	v = pv

	// Check type of target.
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return kerr.New("AODAOUPIED", nil, "Array must be Array or Slice. This is %s", v.Kind())
	}

	i := 0

	for _, in := range in.Array() {

		// Get element of array, growing if necessary.
		if v.Kind() == reflect.Slice {
			// Grow slice if necessary
			if i >= v.Cap() {
				newcap := v.Cap() + v.Cap()/2
				if newcap < 4 {
					newcap = 4
				}
				newv := reflect.MakeSlice(v.Type(), v.Len(), newcap)
				reflect.Copy(newv, v)
				v.Set(newv)
			}
			if i >= v.Len() {
				v.SetLen(i + 1)
			}
		}

		if i < v.Len() {
			// Decode into element.
			if err := us.unpack(ctx, in, v.Index(i)); err != nil {
				return kerr.New("PKGCUIXVWS", err, "unpack")
			}
		}
		i++
	}

	if i < v.Len() {
		if v.Kind() == reflect.Array {
			// Array.  Zero the rest.
			z := reflect.Zero(v.Type().Elem())
			for ; i < v.Len(); i++ {
				v.Index(i).Set(z)
			}
		} else {
			v.SetLen(i)
		}
	}
	if i == 0 && v.Kind() == reflect.Slice {
		v.Set(reflect.MakeSlice(v.Type(), 0, 0))
	}
	return nil
}

func (us *unpackStruct) unpackObject(ctx context.Context, in Packed, v reflect.Value) error {

	if in.Type() != J_MAP {
		return kerr.New("PBAXKEKVTA", nil, "Type %s should be J_MAP", in.Type())
	}

	hasConcreteType := false
	concreteTypePath := ""
	concreteTypeName := ""

	_, _, _, val := indirect(v, false, false, false)

	// If the type we're unmarshaling into is an interface, we should scan for a "type"
	// attribute and initialise the correct type.
	switch val.Kind() {
	case reflect.Interface:
		// This sets the value of v to the correct type based on the "type" attribute.
		typ, err := us.getTypeFromField(ctx, in, v)
		if err != nil {
			return kerr.New("BGJEIXFQHL", err, "getTypeFromField (interface) %s", v.String())
		}
		if typ != nil {
			if err := setType(v, typ); err != nil {
				return kerr.New("KBWJCMHWYF", err, "setType")
			}
		}
	case reflect.Struct:
		// If we're unmarshaling into a concrete type, we want to be able to omit the "type"
		// attribute, so we should add it back in if it's missing so the system:base object is
		// correct.
		path, name, ok := jsonctx.FromContext(ctx).GetTypeByReflectType(val.Type())
		if ok {
			hasConcreteType = true
			concreteTypePath = path
			concreteTypeName = name
		}
	}

	_, _, up, rv := indirect(v, false, false, true)
	if up != nil {
		if err := up.Unpack(ctx, in); err != nil {
			return kerr.New("NPDUYUXVVK", err, "Unpack")
		}
		return nil
	}
	v = rv

	// Decoding into nil interface? Just use the input value
	if v.Kind() == reflect.Interface && v.NumMethod() == 0 {
		v.Set(reflect.ValueOf(in.Interface()))
		return nil
	}

	// Check type of target: struct or map[string]T
	switch v.Kind() {
	case reflect.Map:
		// map must have string kind
		t := v.Type()
		if t.Key().Kind() != reflect.String {
			return kerr.New("TXNQGFVHOT", nil, "Map must have string keys. This has %s", t.Key().Kind())
		}
		if v.IsNil() {
			v.Set(reflect.MakeMap(t))
		}
	case reflect.Struct:
		// This is ok.
	default:
		return kerr.New("AMDJPDYCGI", nil, "unpackObject only unpacks maps and structs. This is %s %s", v.Type().Name(), v.Kind())
	}

	var mapElem reflect.Value

	foundFields := make([]field, 10)
	for key, val := range in.Map() {

		// Figure out field corresponding to key.
		var subv reflect.Value

		if v.Kind() == reflect.Map {
			elemType := v.Type().Elem()
			if !mapElem.IsValid() {
				mapElem = reflect.New(elemType).Elem()
			} else {
				mapElem.Set(reflect.Zero(elemType))
			}
			subv = mapElem
		} else {
			var f *field
			fields := cachedTypeFields(v.Type())
			for i := range fields {
				ff := &fields[i]
				if bytes.Equal(ff.nameBytes, []byte(key)) {
					f = ff
					break
				}
				if f == nil && ff.equalFold(ff.nameBytes, []byte(key)) {
					f = ff
				}
			}
			if f != nil {
				subv = v
				if f.quoted {
					return kerr.New("SRULCNWOWM", nil, "Quoted json not supported by Unpack")
				}
				for _, i := range f.index {
					if subv.Kind() == reflect.Ptr {
						if subv.IsNil() {
							subv.Set(reflect.New(subv.Type().Elem()))
						}
						subv = subv.Elem()
					}
					subv = subv.Field(i)
				}
				foundFields = append(foundFields, *f)
			}
		}

		if err := us.unpack(ctx, val, subv); err != nil {
			return kerr.New("SIJHJHWXYF", err, "unpack")
		}

		// Write value back to map;
		// if using struct, subv points into struct already.
		if v.Kind() == reflect.Map {
			kv := reflect.ValueOf(key).Convert(v.Type().Key())
			v.SetMapIndex(kv, subv)
		}
	}

	if err := initialiseUnmarshaledObject(ctx, v, foundFields, true, hasConcreteType, concreteTypePath, concreteTypeName); err != nil {
		return kerr.New("XWHQSWVNLF", err, "initialiseUnmarshaledObject")
	}
	return nil
}

func (us *unpackStruct) getTypeFromField(ctx context.Context, in Packed, iface reflect.Value) (reflect.Type, error) {

	if in.Type() != J_MAP {
		return nil, kerr.New("LCJRIHJXFU", nil, "Type %s should be J_MAP", in.Type())
	}

	m := in.Map()
	t, ok := m["type"]
	if !ok {
		return nil, kerr.New("RMMVQNVHTU", nil, "Input missing type field")
	}
	if t.Type() != J_STRING {
		return nil, kerr.New("RPBSKPRLJQ", nil, "Type field %s is not string", t.Type())
	}
	typePath, typeName, err := GetReferencePartsFromTypeString(ctx, t.String())
	if err != nil {
		if unk, ok := err.(UnknownPackageError); ok {
			// We don't want to throw an error here, because when we're scanning for
			// aliases we need to tolerate unknown packages
			us.unknownPackage = unk.UnknownPackage
		} else {
			return nil, kerr.New("KXBNXCCRYH", err, "GetReferencePartsFromTypeString")
		}
	}
	return us.getType(ctx, typePath, typeName, iface), nil
}

func (us *unpackStruct) getType(ctx context.Context, typePath string, typeName string, iface reflect.Value) reflect.Type {
	jcache := jsonctx.FromContext(ctx)
	typ, ok := jcache.GetType(typePath, typeName)
	if !ok && iface.Kind() == reflect.Interface {

		// If we can't find the type in the resolver, and
		// we're unmarshaling into an interface, then look
		// the interface in the dummy interface resolver.
		typ, _ = jcache.Dummies.Get(iface.Type())
	}
	if typ == nil {
		us.unknownType = typePath + ":" + typeName
	}
	return typ
}
