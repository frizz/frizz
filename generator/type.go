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
func Type(fieldName string, field system.RuleInterface, path string, getAlias func(string) string) (string, error) {

	outer, err := system.WrapRule(field)
	if err != nil {
		return "", kerr.New("TFXFBIRXHN", err, "NewRuleHolder")
	}

	// if the rule is a complex collection, with possibly several maps and arrays, this
	// iterates over the rule and returns the go collection prefix - e.g. []map[string]
	// for an array of maps. It also returns the inner rule.
	prefix, inner, err := collectionPrefixInnerRule("", outer)
	if err != nil {
		return "", kerr.New("SOGEFOPJHB", err, "collectionPrefixInnerRule")
	}

	var name, pointer string
	if inner.Struct.Interface {
		// if this is an interface rule, we print the interface name of the inner type,
		// which never has a pointer asterisk.
		pointer = ""
		name = Reference(inner.Parent.Id.Package, system.GoInterfaceName(inner.Parent.Id.Name), path, getAlias)
	} else {
		// this returns a "*" if the type should be prefixed by it. Native and interface types
		// don't have a *.
		pointer = getPointer(inner.Parent)
		name = Reference(inner.Parent.Id.Package, system.GoName(inner.Parent.Id.Name), path, getAlias)
	}

	// TODO: Why aren't we giving getTag the correct path and aliases?!?
	tag, err := getTag(fieldName, field.GetRule().Exclude, inner, "", map[string]string{})
	if err != nil {
		return "", kerr.New("CSJHNCMHRU", err, "getTag")
	}
	if tag != "" {
		tag = " " + tag
	}

	return fmt.Sprint(prefix, pointer, name, tag), nil
}

// collectionPrefix recursively digs down through collection rules, recursively
// calling itself as long as it finds a collection rule (map or array). It returns
// the full collection prefix (e.g. any number of appended [] and map[string]'s)
// and the inner (non collection) rule.
func collectionPrefixInnerRule(prefix string, outer *system.RuleWrapper) (fullPrefix string, inner *system.RuleWrapper, err error) {
	p := outer.Parent
	if p.IsNativeCollection() {
		if p.Native.Value == "array" {
			prefix += "[]"
		} else if p.Native.Value == "map" {
			prefix += "map[string]"
		}
		items, err := outer.ItemsRule()
		if err != nil {
			return "", nil, kerr.New("SUTYJEGBKW", err, "outer.ItemsRule")
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

func formatTag(fieldName string, exclude bool, defaultBytes []byte, r *system.RuleWrapper, path string, aliases map[string]string) (string, error) {

	kegoTag := ""
	if defaultBytes != nil && string(defaultBytes) != "null" {
		defaultRaw := json.RawMessage(defaultBytes)
		t := r.Parent.FullName()
		var tag json.KegoTag
		if t == "kego.io/system:string" || t == "kego.io/system:number" || t == "kego.io/system:bool" {
			// If our default is one of the basic system native types, we know we can unmarshal it without
			// the extra context, so we omit type, path and aliases. This makes the generated code easier to
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
					Path:    path,
					Aliases: aliases,
					Type:    t,
				},
			}
		}

		jsonBytes, err := json.MarshalPlain(tag)
		if err != nil {
			return "", kerr.New("LKBWJTMJCF", err, "json.Marshal(tag)")
		}
		kegoTag = string(jsonBytes)
	}

	tag := ""
	tag = addSubTag(tag, "kego", kegoTag)
	if exclude {
		tag = addSubTag(tag, "json", "-")
	} else {
		tag = addSubTag(tag, "json", fieldName)
	}

	if strconv.CanBackquote(tag) {
		return "`" + tag + "`", nil
	}
	return strconv.Quote(tag), nil
}

func addSubTag(tag string, name string, content string) string {
	if content == "" {
		return tag
	}
	if tag != "" {
		tag += " "
	}
	return fmt.Sprintf("%s%s:%s", tag, name, strconv.Quote(content))
}

func getTag(fieldName string, exclude bool, r *system.RuleWrapper, path string, aliases map[string]string) (string, error) {

	dr, ok := r.Interface.(system.DefaultRule)
	if !ok {
		// Doesn't have a default field
		return formatTag(fieldName, exclude, nil, r, path, aliases)
	}

	d := dr.GetDefault()
	if d == nil {
		return formatTag(fieldName, exclude, nil, r, path, aliases)
	}

	// If we have a marshaler, we have to call it manually
	if m, ok := d.(json.Marshaler); ok {
		defaultBytes, err := m.MarshalJSON()
		if err != nil {
			return "", kerr.New("YIEMHYFVCD", err, "m.MarshalJSON")
		}
		return formatTag(fieldName, exclude, defaultBytes, r, path, aliases)
	}

	defaultBytes, err := json.MarshalPlain(d)
	if err != nil {
		return "", kerr.New("QQDOLAJKLU", err, "json.Marshal (typed)")
	}

	return formatTag(fieldName, exclude, defaultBytes, r, path, aliases)
}
