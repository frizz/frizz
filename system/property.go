package system

import (
	"fmt"
	"strconv"

	"kego.io/json"
)

func (p *Property) GoName(name string) string {
	return IdToGoName(name)
}

func getPointer(t *Type) string {
	isNative := t.IsNativeValue()
	isInterface := t.Interface
	if !isNative && !isInterface {
		return "*"
	}
	return ""
}

func formatTag(defaultBytes []byte, r *RuleHolder) (string, error) {

	if string(defaultBytes) == "null" {
		return "", nil
	}

	defaultRaw := json.RawMessage(defaultBytes)
	t := r.parentType.FullName()
	var tag json.KegoTag
	if t == "kego.io/system:string" || t == "kego.io/system:number" || t == "kego.io/system:bool" {
		// If our default is one of the basic system native types, we know we can unmarshal it without
		// the extra context, so we omit type, path and imports. This makes the generated code easier to
		// understand.
		tag = json.KegoTag{
			Default: &json.KegoDefault{
				Value: &defaultRaw,
			},
		}
	} else {
		tag = json.KegoTag{
			Default: &json.KegoDefault{
				Value:   &defaultRaw,
				Path:    r.path,
				Imports: r.imports,
				Type:    t,
			},
		}
	}

	jsonBytes, err := json.Marshal(tag)
	if err != nil {
		return "", fmt.Errorf("Error in property.getTag: json.Marshal(tag) returned an error:\n%v\n", err)
	}

	return fmt.Sprintf("`kego:%s`", strconv.Quote(string(jsonBytes))), nil
}

func getTag(r *RuleHolder) (string, error) {

	name, _, ok := r.ruleType.Defaulter()
	if !ok {
		// This rule type doesn't support defaulters
		return "", nil
	}

	if i, ok := r.rule.(map[string]interface{}); ok {
		// This rule is an unknown type, so we have to extract the default
		// value manually
		di, ok := i[name]
		if !ok {
			// This rule instance doesn't have a default
			return "", nil
		}
		defaultBytes, err := json.Marshal(di)
		if err != nil {
			return "", fmt.Errorf("Error in Property.getTag: json.Marshal (interface) returned an error:\n%v\n", err)
		}
		return formatTag(defaultBytes, r)
	}

	value, pointer, ok, err := ruleFieldByReflection(r.rule, IdToGoName(name))
	if !ok {
		// Doesn't have a default field
		return "", nil
	}

	// If we have a marshaler, we have to call it manually
	if m, ok := pointer.(json.Marshaler); ok {
		defaultBytes, err := m.MarshalJSON()
		if err != nil {
			return "", fmt.Errorf("Error in Property.getTag: um.MarshalJSON returned an error:\n%v\n", err)
		}
		return formatTag(defaultBytes, r)
	}

	defaultBytes, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("Error in Property.getTag: json.Marshal (typed) returned an error:\n%v\n", err)
	}

	return formatTag(defaultBytes, r)
}

// GoTypeDescriptor returns the Go source for the definition of the type of this property
// [collection prefix][optional pointer][type name]
func (p *Property) GoTypeDescriptor(path string, imports map[string]string) (string, error) {

	outer, err := NewRuleHolder(p.Item, path, imports)
	if err != nil {
		return "", Err(err, "Property.GoTypeDescriptor", "NewRuleHolder")
	}

	prefix, inner, err := collectionPrefixInnerRule("", outer)
	if err != nil {
		return "", Err(err, "Property.GoTypeDescriptor", "collectionPrefixInnerRule")
	}

	pointer := getPointer(inner.parentType)

	name, err := inner.parentType.GoTypeReference(path, imports)
	if err != nil {
		return "", Err(err, "Property.GoTypeDescriptor", "inner.parentType.GoTypeReference")
	}

	tag, err := getTag(inner)
	if err != nil {
		return "", Err(err, "Property.GoTypeDescriptor", "getTag")
	}
	if tag != "" {
		tag = " " + tag
	}

	return fmt.Sprintf("%s%s%s%s", prefix, pointer, name, tag), nil
}

// collectionPrefix recursively digs down through collection rules, recursively
// calling itself as long as it finds a collection rule (map or array). It returns
// the full collection prefix (e.g. any number of appended [] and map[string]'s)
// and the inner (non collection) rule.
func collectionPrefixInnerRule(prefix string, outer *RuleHolder) (fullPrefix string, inner *RuleHolder, err error) {
	p := outer.parentType
	if p.IsNativeCollection() {
		if p.Native.Value == "array" {
			prefix += "[]"
		} else if p.Native.Value == "map" {
			prefix += "map[string]"
		}
		items, err := outer.ItemsRule()
		if err != nil {
			return "", nil, fmt.Errorf("Error in property.collectionPrefixInnerRule: outer.ItemsRule returned an error: \n%v\n", err)
		}
		return collectionPrefixInnerRule(prefix, items)
	} else {
		return prefix, outer, nil
	}
}
