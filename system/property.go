package system

import (
	"fmt"
	"strconv"

	"kego.io/json"
	"kego.io/uerr"
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
	t := r.ParentType.FullName()
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
				Path:    r.Path,
				Imports: r.Imports,
				Type:    t,
			},
		}
	}

	jsonBytes, err := json.Marshal(tag)
	if err != nil {
		return "", uerr.New("LKBWJTMJCF", err, "property.formatTag", "json.Marshal(tag)")
	}

	return fmt.Sprintf("`kego:%s`", strconv.Quote(string(jsonBytes))), nil
}

func getTag(r *RuleHolder) (string, error) {

	name, _, ok := r.RuleType.Defaulter()
	if !ok {
		// This rule type doesn't support defaulters
		return "", nil
	}

	if i, ok := r.Rule.(map[string]interface{}); ok {
		// This rule is an unknown type, so we have to extract the default
		// value manually
		di, ok := i[name]
		if !ok {
			// This rule instance doesn't have a default
			return "", nil
		}
		defaultBytes, err := json.Marshal(di)
		if err != nil {
			return "", uerr.New("FYMGUTAOCR", err, "Property.getTag", "json.Marshal (interface)")
		}
		return formatTag(defaultBytes, r)
	}

	value, pointer, ok, err := ruleFieldByReflection(r.Rule, IdToGoName(name))
	if !ok {
		// Doesn't have a default field
		return "", nil
	}

	// If we have a marshaler, we have to call it manually
	if m, ok := pointer.(json.Marshaler); ok {
		defaultBytes, err := m.MarshalJSON()
		if err != nil {
			return "", uerr.New("YIEMHYFVCD", err, "Property.getTag", "m.MarshalJSON")
		}
		return formatTag(defaultBytes, r)
	}

	defaultBytes, err := json.Marshal(value)
	if err != nil {
		return "", uerr.New("QQDOLAJKLU", err, "Property.getTag", "json.Marshal (typed)")
	}

	return formatTag(defaultBytes, r)
}

// GoTypeDescriptor returns the Go source for the definition of the type of this property
// [collection prefix][optional pointer][type name]
func (p *Property) GoTypeDescriptor(path string, imports map[string]string) (string, error) {

	outer, err := NewRuleHolder(p.Item, path, imports)
	if err != nil {
		return "", uerr.New("QKCXSRPMOQ", err, "Property.GoTypeDescriptor", "NewRuleHolder")
	}

	prefix, inner, err := collectionPrefixInnerRule("", outer)
	if err != nil {
		return "", uerr.New("SNATGPVLAS", err, "Property.GoTypeDescriptor", "collectionPrefixInnerRule")
	}

	pointer := getPointer(inner.ParentType)

	name, err := inner.ParentType.GoTypeReference(path, imports)
	if err != nil {
		return "", uerr.New("CGNYBDFUNP", err, "Property.GoTypeDescriptor", "inner.parentType.GoTypeReference")
	}

	tag, err := getTag(inner)
	if err != nil {
		return "", uerr.New("LKIXCKRCYG", err, "Property.GoTypeDescriptor", "getTag")
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
	p := outer.ParentType
	if p.IsNativeCollection() {
		if p.Native.Value == "array" {
			prefix += "[]"
		} else if p.Native.Value == "map" {
			prefix += "map[string]"
		}
		items, err := outer.ItemsRule()
		if err != nil {
			return "", nil, uerr.New("SUTYJEGBKW", err, "property.collectionPrefixInnerRule", "outer.ItemsRule")
		}
		return collectionPrefixInnerRule(prefix, items)
	} else {
		return prefix, outer, nil
	}
}
