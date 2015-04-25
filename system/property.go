package system

import (
	"fmt"
	"reflect"
	"strconv"

	"kego.io/json"
)

func (p *Property) GoName(name string) string {
	return IdToGoName(name)
}

func getPointer(t *Type) string {
	isNative := t.IsNativeValue()
	isInterface := t.Interface.Value
	if !isNative && !isInterface {
		return "*"
	}
	return ""
}

func getRuleTypeReference(r Rule, localPackagePath string, localImports map[string]string) (*Reference, error) {
	var ruleReference Reference
	b, ok := r.(Basic)
	if ok {
		ruleReference = b.Base().Type
	} else {
		// Looks like we unmarshaled a type that isn't registered. We should give out
		// best guess at a type name.
		i, ok := r.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Error in Property.getParentType: unknown type rule is not map[string]interface{}: %T\n", r)
		}
		t, ok := i["type"]
		if !ok {
			return nil, fmt.Errorf("Error in Property.getParentType: unknown type rule has no type\n")
		}
		s, ok := t.(string)
		if !ok {
			return nil, fmt.Errorf("Error in Property.getParentType: unknown type rule type is not string: %T\n", t)
		}
		r := &Reference{}
		err := r.UnmarshalJSON([]byte(strconv.Quote(s)), localPackagePath, localImports)
		if err != nil {
			return nil, fmt.Errorf("Error in Property.getParentType: r.UnmarshalJSON returned an error:\n%v\n", err)
		}
		ruleReference = *r
	}
	return &ruleReference, nil
}

func getTypeName(parentType *Type, localPackagePath string, localImports map[string]string) (string, error) {
	name, err := parentType.GoTypeReference(localImports, localPackagePath)
	if err != nil {
		return "", fmt.Errorf("Error in Property.getTypeName: typeReference.GoReference returned an error:\n%v\n", err)
	}
	return name, nil
}

func getTag(r Rule, ruleType *Type, localPackagePath string, localImports map[string]string) (string, error) {

	name, _, ok := ruleType.Defaulter()
	if !ok {
		//fmt.Printf("Doesn't support defaulters\n")
		// This rule type doesn't support defaulters
		return "", nil
	}

	if i, ok := r.(map[string]interface{}); ok {
		// This rule is an unknown type, so we have to extract the default
		// value manually
		di, ok := i[name]
		if !ok {
			//fmt.Printf("Doesn't have a default (interface)\n")
			// This rule instance doesn't have a default
			return "", nil
		}
		defaultBytes, err := json.Marshal(di)
		if err != nil {
			return "", fmt.Errorf("Error in Property.getTag: json.Marshal (interface) returned an error:\n%v\n", err)
		}
		//fmt.Printf("Default: %s\n", defaultBytes)
		return formatTag(defaultBytes), nil
	}

	val := reflect.ValueOf(r)
	if val.Kind() != reflect.Ptr {
		return "", fmt.Errorf("Error in Property.getTag: val.Kind is not a Ptr: %v\n", val.Kind())
	}
	if val.Elem().Kind() != reflect.Struct {
		return "", fmt.Errorf("Error in Property.getTag: val.Elem().Kind is not a Struct: %v\n", val.Elem().Kind())
	}
	field := val.Elem().FieldByName(IdToGoName(name))
	empty := reflect.Value{}
	if field == empty {
		//fmt.Printf("Doesn't have a default (typed)\n")
		return "", nil
	}

	//fmt.Printf("Not interface: %#v\n", field.Interface())
	if um, ok := field.Addr().Interface().(json.Marshaler); ok {
		defaultBytes, err := um.MarshalJSON()
		if err != nil {
			return "", fmt.Errorf("Error in Property.getTag: um.MarshalJSON returned an error:\n%v\n", err)
		}
		//fmt.Printf("Default (MarshalJSON): %s\n", defaultBytes)
		return formatTag(defaultBytes), nil
	}
	defaultBytes, err := json.Marshal(field.Interface())
	if err != nil {
		return "", fmt.Errorf("Error in Property.getTag: json.Marshal (typed) returned an error:\n%v\n", err)
	}
	//fmt.Printf("Default (json.Marshal): %s\n", defaultBytes)
	return formatTag(defaultBytes), nil
}

func formatTag(defaultBytes []byte) string {
	if string(defaultBytes) == "null" {
		return ""
	}
	jsonString := fmt.Sprintf("{\"default\": %s}", defaultBytes)
	return fmt.Sprintf("`kego:%s`", strconv.Quote(jsonString))
}

func getTypes(rule Rule, path string, imports map[string]string) (ruleType *Type, parentType *Type, err error) {
	ruleReference, err := getRuleTypeReference(rule, path, imports)
	if err != nil {
		return nil, nil, fmt.Errorf("Error in Property.getTypes: getRuleTypeReference returned an error:\n%v\n", err)
	}
	rt, ok := ruleReference.GetType()
	if !ok {
		return nil, nil, fmt.Errorf("Error in Property.getTypes: type %v not found\n", ruleReference.Value)
	}
	typeReference, err := ruleReference.RuleToParentType()
	if err != nil {
		return nil, nil, fmt.Errorf("Error in Property.getTypes: ruleReference.RuleToParentType returned an error:\n%v\n", err)
	}
	pt, ok := typeReference.GetType()
	if !ok {
		return nil, nil, fmt.Errorf("Error in Property.getTypes: type %v not found\n", typeReference.Value)
	}
	return rt, pt, nil
}

// GoTypeDescriptor returns the Go source for the definition of the type of this property
// [collection prefix][optional pointer][type name]
func (p *Property) GoTypeDescriptor(localImports map[string]string, localPackagePath string) (string, error) {

	prefix, rule, err := collectionPrefix("", p.Item, localPackagePath, localImports)
	if err != nil {
		return "", fmt.Errorf("Error in Property.GoTypeDescriptor: collectionPrefix returned an error:\n%v\n", err)
	}

	ruleType, parentType, err := getTypes(rule, localPackagePath, localImports)

	pointer := getPointer(parentType)

	name, err := getTypeName(parentType, localPackagePath, localImports)
	if err != nil {
		return "", fmt.Errorf("Error in Property.GoTypeDescriptor: getTypeName returned an error:\n%v\n", err)
	}

	tag, err := getTag(rule, ruleType, localPackagePath, localImports)
	if err != nil {
		return "", fmt.Errorf("Error in Property.GoTypeDescriptor: getTag returned an error:\n%v\n", err)
	}
	if tag != "" {
		tag = " " + tag
	}

	return fmt.Sprintf("%s%s%s%s", prefix, pointer, name, tag), nil
}

// collectionPrefix recursively digs down through collection rules, recursively
// calling itself as long as it finds a collection rule (map or array).
func collectionPrefix(prefix string, rule Rule, path string, imports map[string]string) (string, Rule, error) {

	ruleType, parentType, err := getTypes(rule, path, imports)
	if err != nil {
		return "", nil, fmt.Errorf("Error in property.collectionPrefix: getTypes returned an error: \n%v\n", err)
	}

	if parentType.IsNativeCollection() {
		if parentType.Native.Value == "array" {
			prefix += "[]"
		} else if parentType.Native.Value == "map" {
			prefix += "map[string]"
		}
		items, err := getItemsRuleFromCollectionRule(rule, ruleType)
		if err != nil {
			return "", nil, fmt.Errorf("Error in property.collectionPrefix: getItemsRuleFromCollectionRule returned an error: \n%v\n", err)
		}
		return collectionPrefix(prefix, items, path, imports)
	} else {
		return prefix, rule, nil
	}
}

func getItemsRuleFromCollectionRule(rule Rule, ruleType *Type) (Rule, error) {
	val := reflect.ValueOf(rule)
	if val.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("Error in getItemsRuleFromCollectionRule: val.Kind is not a Ptr: %v\n", val.Kind())
	}
	if val.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("Error in getItemsRuleFromCollectionRule: val.Elem().Kind is not a Struct: %v\n", val.Elem().Kind())
	}
	field := val.Elem().FieldByName("Items")
	empty := reflect.Value{}
	if field == empty {
		// Doesn't have a default
		return nil, nil
	}
	r, ok := field.Interface().(Rule)
	if !ok {
		return nil, fmt.Errorf("Error in getItemsRuleFromCollectionRule: items is not a rule.\n")
	}
	return r, nil
}
