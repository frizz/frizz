package system

import (
	"reflect"
	"strconv"

	"kego.io/kerr"
)

type Rule interface{}

type RuleHolder struct {
	Rule       Rule
	RuleType   *Type
	ParentType *Type
	Path       string
	Imports    map[string]string
}

func NewRuleHolder(r Rule, path string, imports map[string]string) (*RuleHolder, error) {
	rt, pt, err := ruleTypes(r, path, imports)
	if err != nil {
		return nil, kerr.New("VRCWUGOTMA", err, "NewRuleHolder", "ruleTypes")
	}
	return &RuleHolder{Rule: r, RuleType: rt, ParentType: pt, Path: path, Imports: imports}, nil
}

func ruleTypes(r Rule, path string, imports map[string]string) (ruleType *Type, parentType *Type, err error) {
	ruleReference, err := ruleTypeReference(r, path, imports)
	if err != nil {
		return nil, nil, kerr.New("BNEKIFYDDL", err, "ruleTypes", "ruleTypeReference")
	}
	rt, ok := ruleReference.GetType()
	if !ok {
		return nil, nil, kerr.New("PFGWISOHRR", nil, "ruleTypes", "ruleReference.GetType: type %v not found", ruleReference.Value)
	}
	typeReference, err := ruleReference.RuleToParentType()
	if err != nil {
		return nil, nil, kerr.New("NXRCPQMUIE", err, "ruleTypes", "ruleReference.RuleToParentType")
	}
	pt, ok := typeReference.GetType()
	if !ok {
		return nil, nil, kerr.New("KYCTDXKFYR", nil, "ruleTypes", "typeReference.GetType: type %v not found", typeReference.Value)
	}
	return rt, pt, nil
}

func ruleTypeReference(r Rule, path string, imports map[string]string) (*Reference, error) {

	m, ok := r.(map[string]interface{})
	if ok {
		// Looks like we unmarshaled a type that isn't registered. We
		// can get the type information from the "type" field.
		i, ok := m["type"]
		if !ok {
			return nil, kerr.New("OLHOVKXEXN", nil, "ruleTypeReference", "map[string]interface{} rule has no type attribute")
		}
		s, ok := i.(string)
		if !ok {
			return nil, kerr.New("IILEXGQDXL", nil, "ruleTypeReference", "map[string]interface{} rule type attribute is not string: %T", i)
		}

		i, ok = m["_path"]
		if !ok {
			return nil, kerr.New("FUMLBULJCP", nil, "ruleTypeReference", "map[string]interface{} rule has no _path attribute")
		}
		innerPath, ok := i.(string)
		if !ok {
			return nil, kerr.New("ERIHIYYQYL", nil, "ruleTypeReference", "map[string]interface{} rule _path attribute is not string: %T", i)
		}

		i, ok = m["_imports"]
		if !ok {
			return nil, kerr.New("PJJYHJMPUI", nil, "ruleTypeReference", "map[string]interface{} rule has no _imports attribute")
		}
		innerImports, ok := i.(map[string]string)
		if !ok {
			return nil, kerr.New("KSVRLDNLTL", nil, "ruleTypeReference", "map[string]interface{} rule _imports attribute is not map[string]string: %T", i)
		}

		ref := &Reference{}
		err := ref.UnmarshalJSON([]byte(strconv.Quote(s)), innerPath, innerImports)
		if err != nil {
			return nil, kerr.New("QBTHPRVBWN", err, "ruleTypeReference", "ref.UnmarshalJSON")
		}
		return ref, nil
	}

	// Looks like we have a standard rule. We can get the type by
	// reflection "Type" field.
	ti, _, ok, err := ruleFieldByReflection(r, "Type")
	if err != nil {
		return nil, kerr.New("QJQAIGPYXC", err, "ruleTypeReference", "ruleFieldByReflection")
	}
	if !ok {
		return nil, kerr.New("NXYRAJITEV", nil, "ruleTypeReference", "ruleFieldByReflection did not find Type field")
	}
	tr, ok := ti.(Reference)
	if !ok {
		return nil, kerr.New("FHUPSRTRFE", nil, "ruleTypeReference", "Type field is not a reference")
	}
	return &tr, nil
}

// ItemsRule returns Items rule for a collection Rule.
func (r *RuleHolder) ItemsRule() (*RuleHolder, error) {
	if !r.ParentType.IsNativeCollection() {
		return nil, kerr.New("VPAGXSTQHM", nil, "RuleHolder.ItemsRule", "parentType %s is not a collection", r.ParentType.Id)
	}
	items, _, ok, err := ruleFieldByReflection(r.Rule, "Items")
	if err != nil {
		return nil, kerr.New("LIDXIQYGJD", err, "RuleHolder.ItemsRule", "ruleFieldByReflection")
	}
	if !ok {
		return nil, kerr.New("VYTHGJTSNJ", nil, "RuleHolder.ItemsRule", "ruleFieldByReflection could not find Items field")
	}
	rule, _ := items.(Rule)
	// Rule is interface{} so everything is a rule... Maybe this will change in future?
	//if !ok {
	//	return nil, kerr.New("DIFVRMVWMC", nil, "RuleHolder.ItemsRule", "items is not a rule")
	//}
	rh, err := NewRuleHolder(rule, r.Path, r.Imports)
	if err != nil {
		return nil, kerr.New("FGYMQPNBQJ", err, "RuleHolder.ItemsRule", "NewRuleHolder")
	}
	return rh, nil
}

func ruleFieldByReflection(object interface{}, name string) (value interface{}, pointer interface{}, ok bool, err error) {
	v := reflect.ValueOf(object)
	if v.Kind() != reflect.Ptr {
		return nil, nil, false, kerr.New("QOYMWPXWUO", nil, "ruleFieldByReflection", "val.Kind (%s) is not a Ptr: %v", name, v.Kind())
	}
	if v.Elem().Kind() != reflect.Struct {
		return nil, nil, false, kerr.New("IGOUOBGXAN", nil, "ruleFieldByReflection", "val.Elem().Kind (%s) is not a Struct: %v", name, v.Elem().Kind())
	}
	value, pointer, found, zero, err := GetField(v.Elem(), name)

	// zero => !ok
	return value, pointer, found && !zero, err
}
func GetField(v reflect.Value, name string) (value interface{}, pointer interface{}, found bool, zero bool, err error) {
	field := v.FieldByName(name)
	empty := reflect.Value{}
	if field == empty {
		// Field does not exist
		return
	}
	if field.Kind() == reflect.Ptr {
		// If it's a pointer we should only return not found if
		// it's nil:
		if field.IsNil() {
			zero = true
			return
		}
	} else if field.Kind() == reflect.Map || field.Kind() == reflect.Slice {
		if field.Len() == 0 {
			zero = true
		}
	} else {
		// If it's not a pointer, we return not found if it's an
		// zero value
		nilValue := reflect.Zero(field.Type())
		if field.Interface() == nilValue.Interface() {
			zero = true
		}
	}
	found = true
	value = field.Interface()
	// This prevents **Foo being returned for pointer when value is already *Foo
	if field.Kind() == reflect.Ptr {
		pointer = field.Interface()
	} else {
		pointer = field.Addr().Interface()
	}
	return
}
