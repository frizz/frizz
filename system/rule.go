package system

import (
	"reflect"
	"strconv"

	"kego.io/uerr"
)

type Rule interface{}

type RuleHolder struct {
	rule       Rule
	ruleType   *Type
	parentType *Type
	path       string
	imports    map[string]string
}

func NewRuleHolder(r Rule, path string, imports map[string]string) (*RuleHolder, error) {
	rt, pt, err := ruleTypes(r, path, imports)
	if err != nil {
		return nil, uerr.New("VRCWUGOTMA", err, "NewRuleHolder", "ruleTypes")
	}
	return &RuleHolder{rule: r, ruleType: rt, parentType: pt, path: path, imports: imports}, nil
}

func ruleTypes(r Rule, path string, imports map[string]string) (ruleType *Type, parentType *Type, err error) {
	ruleReference, err := ruleTypeReference(r, path, imports)
	if err != nil {
		return nil, nil, uerr.New("BNEKIFYDDL", err, "ruleTypes", "ruleTypeReference")
	}
	rt, ok := ruleReference.GetType()
	if !ok {
		return nil, nil, uerr.New("PFGWISOHRR", nil, "ruleTypes", "ruleReference.GetType: type %v not found", ruleReference.Value)
	}
	typeReference, err := ruleReference.RuleToParentType()
	if err != nil {
		return nil, nil, uerr.New("NXRCPQMUIE", err, "ruleTypes", "ruleReference.RuleToParentType")
	}
	pt, ok := typeReference.GetType()
	if !ok {
		return nil, nil, uerr.New("KYCTDXKFYR", nil, "ruleTypes", "typeReference.GetType: type %v not found", typeReference.Value)
	}
	return rt, pt, nil
}

func ruleTypeReference(r Rule, path string, imports map[string]string) (*Reference, error) {

	i, ok := r.(map[string]interface{})

	if ok {
		// Looks like we unmarshaled a type that isn't registered. We can get the type information
		// from the "type" field.
		t, ok := i["type"]
		if !ok {
			return nil, uerr.New("OLHOVKXEXN", nil, "ruleTypeReference", "map[string]interface{} rule has no type attribute")
		}
		s, ok := t.(string)
		if !ok {
			return nil, uerr.New("IILEXGQDXL", nil, "ruleTypeReference", "map[string]interface{} rule type attribute is not string: %T", t)
		}
		ref := &Reference{}
		err := ref.UnmarshalJSON([]byte(strconv.Quote(s)), path, imports)
		if err != nil {
			return nil, uerr.New("QBTHPRVBWN", err, "ruleTypeReference", "ref.UnmarshalJSON")
		}
		return ref, nil
	}

	// Looks like we have a standard rule. We can get the type by reflection "Type" field.
	ti, _, ok, err := ruleFieldByReflection(r, "Type")
	if err != nil {
		return nil, uerr.New("QJQAIGPYXC", err, "ruleTypeReference", "ruleFieldByReflection")
	}
	if !ok {
		return nil, uerr.New("NXYRAJITEV", nil, "ruleTypeReference", "ruleFieldByReflection did not find Type field")
	}
	tr, ok := ti.(Reference)
	if !ok {
		return nil, uerr.New("FHUPSRTRFE", nil, "ruleTypeReference", "Type field is not a reference")
	}
	return &tr, nil
}

// ItemsRule returns Items rule for a collection Rule.
func (r *RuleHolder) ItemsRule() (*RuleHolder, error) {
	if !r.parentType.IsNativeCollection() {
		return nil, uerr.New("VPAGXSTQHM", nil, "RuleHolder.ItemsRule", "parentType %s is not a collection", r.parentType.Id)
	}
	items, _, ok, err := ruleFieldByReflection(r.rule, "Items")
	if err != nil {
		return nil, uerr.New("LIDXIQYGJD", err, "RuleHolder.ItemsRule", "ruleFieldByReflection")
	}
	if !ok {
		return nil, uerr.New("VYTHGJTSNJ", nil, "RuleHolder.ItemsRule", "ruleFieldByReflection could not find Items field")
	}
	rule, ok := items.(Rule)
	if !ok {
		return nil, uerr.New("DIFVRMVWMC", nil, "RuleHolder.ItemsRule", "items is not a rule")
	}
	rh, err := NewRuleHolder(rule, r.path, r.imports)
	if err != nil {
		return nil, uerr.New("FGYMQPNBQJ", err, "RuleHolder.ItemsRule", "NewRuleHolder")
	}
	return rh, nil
}

func ruleFieldByReflection(object interface{}, name string) (value interface{}, pointer interface{}, ok bool, err error) {
	val := reflect.ValueOf(object)
	if val.Kind() != reflect.Ptr {
		return nil, nil, false, uerr.New("QOYMWPXWUO", nil, "ruleFieldByReflection", "val.Kind (%s) is not a Ptr: %v", name, val.Kind())
	}
	if val.Elem().Kind() != reflect.Struct {
		return nil, nil, false, uerr.New("IGOUOBGXAN", nil, "ruleFieldByReflection", "val.Elem().Kind (%s) is not a Struct: %v", name, val.Elem().Kind())
	}
	field := val.Elem().FieldByName(name)
	empty := reflect.Value{}
	if field == empty {
		// Doesn't have an items
		return nil, nil, false, nil
	}
	if field.Kind() == reflect.Ptr {
		// If it's a pointer we should only return not found if
		// it's nil:
		if field.IsNil() {
			return nil, nil, false, nil
		}
	} else {
		// If it's not a pointer, we return not found if it's an
		// zero value
		nilValue := reflect.Zero(field.Type())
		if field.Interface() == nilValue.Interface() {
			return nil, nil, false, nil
		}
	}
	// This prevents **Foo being returned for pointer when value is already *Foo
	if field.Kind() == reflect.Ptr {
		pointer = field.Interface()
	} else {
		pointer = field.Addr().Interface()
	}
	return field.Interface(), pointer, true, nil
}
