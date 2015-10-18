package json

import (
	"bytes"
	"encoding/base64"
	"reflect"
	"strconv"

	"kego.io/kerr"
)

type unpackStruct struct {
	unknownType    string // have we encountered an unknown type?
	unknownPackage string // have we encountered an unknown package?
}

func Unpack(in Unpackable, out *interface{}, path string, aliases map[string]string) error {

	if in.UpType() != J_MAP {
		return kerr.New("XOOUKLGORQ", nil, "Type %s should be J_MAP", in.UpType())
	}

	us := &unpackStruct{}
	typ, err := us.getTypeFromField(in, reflect.Value{}, path, aliases)
	if err != nil {
		return kerr.New("NUHCPRKRXT", err, "getTypeFromField (global)")
	}
	if typ == nil {
		return kerr.New("GREMVFEUMH", nil, "Unknown global type")
	}

	// we don't wrap the error in kerr because in can return special
	// error types
	return us.unpackFragment(in, out, typ, path, aliases)
}

func UnpackFragment(in Unpackable, out *interface{}, typ reflect.Type, path string, aliases map[string]string) error {

	us := &unpackStruct{}
	// we don't wrap the error in kerr because in can return special
	// error types
	return us.unpackFragment(in, out, typ, path, aliases)
}

func (us *unpackStruct) unpackFragment(in Unpackable, out *interface{}, typ reflect.Type, path string, aliases map[string]string) error {

	p := getEmptyValue(typ)

	err := us.unpack(in, p, path, aliases)

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
	return nil
}

func (us *unpackStruct) unpack(in Unpackable, v reflect.Value, path string, aliases map[string]string) error {

	if !v.IsValid() {
		return nil
	}

	typ := J_NULL
	if in != nil {
		typ = in.UpType()
	}

	switch typ {
	case J_MAP:
		if err := us.unpackObject(in, v, path, aliases); err != nil {
			return kerr.New("LMLUICBTBA", err, "unpackObject")
		}
	case J_ARRAY:
		if err := us.unpackArray(in, v, path, aliases); err != nil {
			return kerr.New("ITJMJWULKO", err, "unpackArray")
		}
	default:
		if err := us.unpackLiteral(in, v, path, aliases); err != nil {
			return kerr.New("BSTNWUKLYO", err, "unpackLiteral")
		}
	}
	return nil
}

func (us *unpackStruct) unpackLiteral(in Unpackable, v reflect.Value, path string, aliases map[string]string) error {

	wantptr := in == nil
	_, _, up, cup, pv := indirect(v, wantptr, false, true)
	if up != nil {
		if err := up.Unpack(in); err != nil {
			return kerr.New("RYSLUEEOAW", err, "Unpack (plain)")
		}
		return nil
	}
	if cup != nil {
		if err := cup.Unpack(in, path, aliases); err != nil {
			return kerr.New("MOKLHNANQB", err, "Unpack (context)")
		}
		return nil
	}
	v = pv

	typ := J_NULL
	if in != nil {
		typ = in.UpType()
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
			v.SetBool(in.UpBool())
		case reflect.Interface:
			if v.NumMethod() == 0 {
				v.Set(reflect.ValueOf(in.UpBool()))
			} else {
				if err := setDefaultNativeValueUnpack(v, in, path, aliases); err != nil {
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
			b := make([]byte, base64.StdEncoding.DecodedLen(len(in.UpString())))
			n, err := base64.StdEncoding.Decode(b, []byte(in.UpString()))
			if err != nil {
				return kerr.New("OKMBMDOFNL", err, "base64.StdEncoding.Decode")
			}
			v.Set(reflect.ValueOf(b[0:n]))
		case reflect.String:
			v.SetString(in.UpString())
		case reflect.Interface:
			if v.NumMethod() == 0 {
				v.Set(reflect.ValueOf(in.UpString()))
			} else {
				if err := setDefaultNativeValueUnpack(v, in, path, aliases); err != nil {
					return kerr.New("PRTYFVSEFT", err, "setDefaultNativeValueUnpack (string)")
				}
			}
		}
	case J_NUMBER:
		s := strconv.FormatFloat(in.UpNumber(), 'f', -1, 64)
		switch v.Kind() {
		default:
			if v.Kind() == reflect.String && v.Type() == numberType {
				v.SetString(s)
				break
			}
			return &UnmarshalTypeError{"number", v.Type()}
		case reflect.Interface:
			if v.NumMethod() != 0 {
				if err := setDefaultNativeValueUnpack(v, in, path, aliases); err != nil {
					return kerr.New("FYHICWEELI", err, "setDefaultNativeValueUnpack (number)")
				}
				break
			}
			v.Set(reflect.ValueOf(in.UpNumber()))
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
			if v.OverflowFloat(in.UpNumber()) {
				return &UnmarshalTypeError{"number " + s, v.Type()}
			}
			v.SetFloat(in.UpNumber())
		}
	}
	return nil
}

func setDefaultNativeValueUnpack(v reflect.Value, in Unpackable, path string, aliases map[string]string) error {
	t, ok := GetTypeByInterface(v.Type())
	if !ok {
		return kerr.New("YSBBTCVOUU", nil, "No type found for %s", v.Type().Name())
	}
	var i interface{}
	err := UnpackFragment(in, &i, t, path, aliases)
	if err != nil {
		return kerr.New("BCVBRIKFJX", err, "unpack (fragment)")
	}
	v.Set(reflect.ValueOf(i))
	return nil
}

func (us *unpackStruct) unpackArray(in Unpackable, v reflect.Value, path string, aliases map[string]string) error {

	if in.UpType() != J_ARRAY {
		return kerr.New("PVJMVPULMY", nil, "Type %s should be J_ARRAY", in.UpType())
	}

	// Check for unpackers.
	_, _, up, cup, pv := indirect(v, false, false, true)
	if up != nil {
		if err := up.Unpack(in); err != nil {
			return kerr.New("PQRNFAYAYQ", err, "Unpack (plain)")
		}
		return nil
	}
	if cup != nil {
		if err := cup.Unpack(in, path, aliases); err != nil {
			return kerr.New("HPUKPAPBDC", err, "Unpack (context)")
		}
		return nil
	}
	v = pv

	// Check type of target.
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return kerr.New("AODAOUPIED", nil, "Array must be Array or Slice. This is %s", v.Kind())
	}

	i := 0

	for _, in := range in.UpArray() {

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
			if err := us.unpack(in, v.Index(i), path, aliases); err != nil {
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

func (us *unpackStruct) unpackObject(in Unpackable, v reflect.Value, path string, aliases map[string]string) error {

	if in.UpType() != J_MAP {
		return kerr.New("PBAXKEKVTA", nil, "Type %s should be J_MAP", in.UpType())
	}

	hasConcreteType := false
	concreteTypePath := ""
	concreteTypeName := ""

	_, _, _, _, val := indirect(v, false, false, false)

	// If the type we're unmarshaling into is an interface, we should scan for a "type"
	// attribute and initialise the correct type.
	switch val.Kind() {
	case reflect.Interface:
		// This sets the value of v to the correct type based on the "type" attribute.
		typ, err := us.getTypeFromField(in, v, path, aliases)
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
		path, name, ok := GetTypeByReflectType(val.Type())
		if ok {
			hasConcreteType = true
			concreteTypePath = path
			concreteTypeName = name
		}
	}

	_, _, u, cu, rv := indirect(v, false, false, true)
	if u != nil {
		if err := u.Unpack(in); err != nil {
			return kerr.New("NPDUYUXVVK", err, "Unpack (plain)")
		}
		return nil
	}
	if cu != nil {
		if err := cu.Unpack(in, path, aliases); err != nil {
			return kerr.New("TXWXHXCDWW", err, "Unpack (context)")
		}
		return nil
	}
	v = rv

	// Decoding into nil interface? Just use the input value
	if v.Kind() == reflect.Interface && v.NumMethod() == 0 {
		v.Set(reflect.ValueOf(in.UpInterface()))
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
	for key, val := range in.UpMap() {

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

		if err := us.unpack(val, subv, path, aliases); err != nil {
			return kerr.New("SIJHJHWXYF", err, "unpack")
		}

		// Write value back to map;
		// if using struct, subv points into struct already.
		if v.Kind() == reflect.Map {
			kv := reflect.ValueOf(key).Convert(v.Type().Key())
			v.SetMapIndex(kv, subv)
		}
	}

	if err := initialiseUnmarshaledObject(v, foundFields, true, hasConcreteType, concreteTypePath, concreteTypeName); err != nil {
		return kerr.New("XWHQSWVNLF", err, "initialiseUnmarshaledObject")
	}
	return nil
}

func (us *unpackStruct) getTypeFromField(in Unpackable, iface reflect.Value, path string, aliases map[string]string) (reflect.Type, error) {

	if in.UpType() != J_MAP {
		return nil, kerr.New("LCJRIHJXFU", nil, "Type %s should be J_MAP", in.UpType())
	}

	m := in.UpMap()
	t, ok := m["type"]
	if !ok {
		return nil, kerr.New("RMMVQNVHTU", nil, "Input missing type field")
	}
	if t.UpType() != J_STRING {
		return nil, kerr.New("RPBSKPRLJQ", nil, "Type field %s is not string", t.UpType())
	}
	typePath, typeName, err := GetReferencePartsFromTypeString(t.UpString(), path, aliases)
	if err != nil {
		if unk, ok := err.(UnknownPackageError); ok {
			// We don't want to throw an error here, because when we're scanning for
			// aliases we need to tolerate unknown packages
			us.unknownPackage = unk.UnknownPackage
		} else {
			return nil, kerr.New("KXBNXCCRYH", err, "GetReferencePartsFromTypeString")
		}
	}
	return us.getType(typePath, typeName, iface), nil
}

func (us *unpackStruct) getType(typePath string, typeName string, iface reflect.Value) reflect.Type {

	typ, _, ok := GetType(typePath, typeName)
	if !ok && iface.Kind() == reflect.Interface {

		// If we can't find the type in the resolver, and
		// we're unmarshaling into an interface, then look
		// the interface in the dummy interface resolver.
		typ, _ = GetInterface(iface.Type())
	}
	if typ == nil {
		us.unknownType = typePath + ":" + typeName
	}
	return typ
}
