package builder

import (
	"fmt"
	"strconv"

	"context"

	"github.com/davelondon/kerr"
	"kego.io/context/envctx"
	"kego.io/json"
	"kego.io/system"
)

// Type returns the Go source for the definition of the type of this property
// [collection prefix][optional pointer][type name]
func Type(ctx context.Context, fieldName string, field system.RuleInterface, path string, getAlias func(string) string) (string, error) {

	outer, err := system.WrapRule(ctx, field)
	if err != nil {
		return "", kerr.Wrap("TFXFBIRXHN", err)
	}

	// if the rule is a complex collection, with possibly several maps and
	// arrays, this iterates over the rule and returns the go collection prefix
	// - e.g. []map[string] for an array of maps. It also returns the inner rule.
	prefix, inner, err := collectionPrefixInnerRule("", outer)
	if err != nil {
		return "", kerr.Wrap("SOGEFOPJHB", err)
	}

	var name, pointer string
	if inner.Struct.Interface {
		// if this is an interface rule, we print the interface name of the
		// inner type, which never has a pointer asterisk.
		pointer = ""
		name = Reference(inner.Parent.Id.Package, system.GoInterfaceName(inner.Parent.Id.Name), path, getAlias)
	} else {
		// this returns a "*" if the type should be prefixed by it. Native and
		// interface types don't have a *.
		pointer = getPointer(inner.Parent)
		name = Reference(inner.Parent.Id.Package, system.GoName(inner.Parent.Id.Name), path, getAlias)
	}

	// TODO: Why aren't we giving getTag the correct path and aliases?!?
	tag, err := getTag(envctx.Empty, fieldName, inner)
	if err != nil {
		return "", kerr.Wrap("CSJHNCMHRU", err)
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

	if _, ok := outer.Interface.(system.CollectionRule); !ok {
		return prefix, outer, nil
	}
	if !outer.IsCollection() {
		// DummyRule is a system.CollectionRule but may not actually be a
		// collection
		return prefix, outer, nil
	}

	switch outer.Parent.Native.Value() {
	case "array":
		prefix += "[]"
	case "map":
		prefix += "map[string]"
	}
	items, err := outer.ItemsRule()
	if err != nil {
		return "", nil, kerr.Wrap("SUTYJEGBKW", err)
	}
	return collectionPrefixInnerRule(prefix, items)
}

func getPointer(t *system.Type) string {
	isJson := t.IsJsonValue()
	isInterface := t.Interface
	if !isJson && !isInterface {
		return "*"
	}
	return ""
}

func formatTag(ctx context.Context, fieldName string, defaultBytes []byte, r *system.RuleWrapper) (string, error) {

	env := envctx.FromContext(ctx)

	kegoTag := ""
	if defaultBytes != nil && string(defaultBytes) != "null" {
		defaultRaw := json.RawMessage(defaultBytes)
		t := r.Parent.Id.Value()
		var tag json.KegoTag
		if t == "kego.io/system:string" || t == "kego.io/system:number" || t == "kego.io/system:bool" {
			// If our default is one of the basic system native types, we know we can unmarshal it
			// without the extra context, so we omit type, path and aliases. This makes the
			// generated code easier to understand.
			tag = json.KegoTag{
				Default: &json.KegoDefault{
					Value: &defaultRaw,
				},
			}
		} else {
			tag = json.KegoTag{
				Default: &json.KegoDefault{
					Value:   &defaultRaw,
					Path:    env.Path,
					Aliases: env.Aliases,
					Type:    t,
				},
			}
		}

		jsonBytes, err := json.MarshalPlain(tag)
		if err != nil {
			return "", kerr.Wrap("LKBWJTMJCF", err)
		}
		kegoTag = string(jsonBytes)
	}

	tag := ""
	tag = addSubTag(tag, "kego", kegoTag)
	tag = addSubTag(tag, "json", fieldName)

	if tag == "" {
		return "", nil
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

func getTag(ctx context.Context, fieldName string, r *system.RuleWrapper) (string, error) {

	dr, ok := r.Interface.(system.DefaultRule)
	if !ok {
		// Doesn't have a default field
		return formatTag(ctx, fieldName, nil, r)
	}

	d := dr.GetDefault()
	if d == nil {
		return formatTag(ctx, fieldName, nil, r)
	}

	// If we have a marshaler, we have to call it manually
	if m, ok := d.(json.Marshaler); ok {
		defaultBytes, err := m.MarshalJSON(ctx)
		if err != nil {
			return "", kerr.Wrap("YIEMHYFVCD", err)
		}
		return formatTag(ctx, fieldName, defaultBytes, r)
	}

	defaultBytes, err := json.MarshalPlain(d)
	if err != nil {
		return "", kerr.Wrap("QQDOLAJKLU", err)
	}

	return formatTag(ctx, fieldName, defaultBytes, r)
}
