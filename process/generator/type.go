package generator

import (
	"fmt"
	"strconv"

	"kego.io/json"
	"kego.io/kerr"
	"kego.io/system"
)

// Type returns the Go source for the definition of the type of this property
// [collection prefix][optional pointer][type name]
func Type(field system.Rule, path string, getAlias func(string) string) (string, error) {

	outer, err := system.NewRuleHolder(field, "", map[string]string{})
	if err != nil {
		return "", kerr.New("TFXFBIRXHN", err, "generator.Type", "NewRuleHolder")
	}

	// if the rule is a complex collection, with possibly several maps and arrays, this
	// iterates over the rule and returns the go collection prefix - e.g. []map[string]
	// for an array of maps. It also returns the inner rule.
	prefix, inner, err := collectionPrefixInnerRule("", outer)
	if err != nil {
		return "", kerr.New("SOGEFOPJHB", err, "generator.Type", "collectionPrefixInnerRule")
	}

	// this returns a "*" if the type should be prefixed by it. Native and interface types
	// don't have a *.
	pointer := getPointer(inner.ParentType)

	name := Reference(inner.ParentType.Id.Package, system.GoName(inner.ParentType.Id.Name), path, getAlias)

	tag, err := getTag(inner)
	if err != nil {
		return "", kerr.New("CSJHNCMHRU", err, "generator.Type", "getTag")
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
func collectionPrefixInnerRule(prefix string, outer *system.RuleHolder) (fullPrefix string, inner *system.RuleHolder, err error) {
	p := outer.ParentType
	if p.IsNativeCollection() {
		if p.Native.Value == "array" {
			prefix += "[]"
		} else if p.Native.Value == "map" {
			prefix += "map[string]"
		}
		items, err := outer.ItemsRule()
		if err != nil {
			return "", nil, kerr.New("SUTYJEGBKW", err, "property.collectionPrefixInnerRule", "outer.ItemsRule")
		}
		return collectionPrefixInnerRule(prefix, items)
	} else {
		return prefix, outer, nil
	}
}

func getPointer(t *system.Type) string {
	isNative := t.IsNativeValue()
	isInterface := t.Interface
	if !isNative && !isInterface {
		return "*"
	}
	return ""
}

func formatTag(defaultBytes []byte, r *system.RuleHolder) (string, error) {

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
		return "", kerr.New("LKBWJTMJCF", err, "property.formatTag", "json.Marshal(tag)")
	}

	return fmt.Sprintf("`kego:%s`", strconv.Quote(string(jsonBytes))), nil
}

func getTag(r *system.RuleHolder) (string, error) {

	value, pointer, ok, err := system.RuleFieldByReflection(r.Rule, "Default")
	if !ok {
		// Doesn't have a default field
		return "", nil
	}

	// If we have a marshaler, we have to call it manually
	if m, ok := pointer.(json.Marshaler); ok {
		defaultBytes, err := m.MarshalJSON()
		if err != nil {
			return "", kerr.New("YIEMHYFVCD", err, "Property.getTag", "m.MarshalJSON")
		}
		return formatTag(defaultBytes, r)
	}

	defaultBytes, err := json.Marshal(value)
	if err != nil {
		return "", kerr.New("QQDOLAJKLU", err, "Property.getTag", "json.Marshal (typed)")
	}

	return formatTag(defaultBytes, r)
}
