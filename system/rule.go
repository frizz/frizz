package system

import (
	"fmt"
	"reflect"
	"strconv"
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
		return nil, fmt.Errorf("Error in NewRuleHolder: ruleTypes returned an error:\n%v\n", err)
	}
	return &RuleHolder{rule: r, ruleType: rt, parentType: pt, path: path, imports: imports}, nil
}

func ruleTypes(r Rule, path string, imports map[string]string) (ruleType *Type, parentType *Type, err error) {
	ruleReference, err := ruleTypeReference(r, path, imports)
	if err != nil {
		return nil, nil, fmt.Errorf("Error in RuleHolder.Types: r.TypeReference returned an error:\n%v\n", err)
	}
	rt, ok := ruleReference.GetType()
	if !ok {
		return nil, nil, fmt.Errorf("Error in RuleHolder.Types: type %v not found\n", ruleReference.Value)
	}
	typeReference, err := ruleReference.RuleToParentType()
	if err != nil {
		return nil, nil, fmt.Errorf("Error in RuleHolder.Types: ruleReference.RuleToParentType returned an error:\n%v\n", err)
	}
	pt, ok := typeReference.GetType()
	if !ok {
		return nil, nil, fmt.Errorf("Error in RuleHolder.Types: type %v not found\n", typeReference.Value)
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
			return nil, fmt.Errorf("Error in RuleHolder.TypeReference: unknown type rule has no type\n")
		}
		s, ok := t.(string)
		if !ok {
			return nil, fmt.Errorf("Error in RuleHolder.TypeReference: unknown type rule type is not string: %T\n", t)
		}
		ref := &Reference{}
		err := ref.UnmarshalJSON([]byte(strconv.Quote(s)), path, imports)
		if err != nil {
			return nil, fmt.Errorf("Error in RuleHolder.TypeReference: ref.UnmarshalJSON returned an error:\n%v\n", err)
		}
		return ref, nil
	}

	// Looks like we have a standard rule. We can get the type by reflection "Type" field.
	ti, _, ok, err := ruleFieldByReflection(r, "Type")
	if err != nil {
		return nil, fmt.Errorf("Error in RuleHolder.TypeReference: ruleFieldByReflection returned an error:\n%v\n", err)
	}
	if !ok {
		return nil, fmt.Errorf("Error in RuleHolder.TypeReference: ruleFieldByReflection did not find Type field.\n")
	}
	tr, ok := ti.(Reference)
	if !ok {
		return nil, fmt.Errorf("Error in RuleHolder.TypeReference: Type field is not a reference.\n")
	}
	return &tr, nil
}

// ItemsRule returns Items rule for a collection Rule.
func (r *RuleHolder) ItemsRule() (*RuleHolder, error) {
	if !r.parentType.IsNativeCollection() {
		return nil, fmt.Errorf("Error in RuleHolder.ItemsRule: parentType %s is not a collection\n", r.parentType.Id)
	}
	items, _, ok, err := ruleFieldByReflection(r.rule, "Items")
	if err != nil {
		return nil, fmt.Errorf("Error in RuleHolder.ItemsRule: ruleFieldByReflection returned an error:\n%v\n", err)
	}
	if !ok {
		return nil, fmt.Errorf("Error in RuleHolder.ItemsRule: ruleFieldByReflection could not find Items field.\n")
	}
	rule, ok := items.(Rule)
	if !ok {
		return nil, fmt.Errorf("Error in RuleHolder.ItemsRule: items is not a rule.\n")
	}
	rh, err := NewRuleHolder(rule, r.path, r.imports)
	if err != nil {
		return nil, fmt.Errorf("Error in RuleHolder.ItemsRule: NewRuleHolder returned an error:\n%v\n", err)
	}
	return rh, nil
}

func ruleFieldByReflection(object interface{}, name string) (value interface{}, pointer interface{}, ok bool, err error) {
	val := reflect.ValueOf(object)
	if val.Kind() != reflect.Ptr {
		return nil, nil, false, fmt.Errorf("Error in FieldByReflection(%s): val.Kind is not a Ptr: %v\n", name, val.Kind())
	}
	if val.Elem().Kind() != reflect.Struct {
		return nil, nil, false, fmt.Errorf("Error in FieldByReflection(%s): val.Elem().Kind is not a Struct: %v\n", name, val.Elem().Kind())
	}
	field := val.Elem().FieldByName(name)
	empty := reflect.Value{}
	if field == empty {
		// Doesn't have an items
		return nil, nil, false, nil
		//return nil, fmt.Errorf("Error in FieldByReflection(%s): val.Elem() does not have field.\n", name)
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
