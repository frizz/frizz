package builder

import (
	"fmt"
	"strconv"

	"context"

	"github.com/davelondon/kerr"
	"kego.io/system"
)

// AliasTypeDefinition returns the Go source for the definition of the type
// of this alias type [collection prefix][optional pointer][type name]
func AliasTypeDefinition(ctx context.Context, alias system.RuleInterface, path string, getAlias func(string) string) (string, error) {

	outer := system.WrapRule(ctx, alias)

	// if the rule is a complex collection, with possibly several maps and
	// arrays, this iterates over the rule and returns the go collection prefix
	// - e.g. []map[string] for an array of maps. It also returns the inner
	// rule.
	prefix, inner, err := collectionPrefixInnerRule(ctx, "", outer)
	if err != nil {
		return "", kerr.Wrap("BSWTXBHVTH", err)
	}

	pointer := ""
	// if we have a prefix we should also work out the innerPointer
	if prefix != "" && inner.Pointer(ctx) {
		pointer = "*"
	}

	name := Reference(inner.Parent.Id.Package, system.GoName(inner.Parent.Id.Name), path, getAlias)

	return fmt.Sprint(prefix, pointer, name), nil
}

// FieldTypeDefinition returns the Go source for the definition of the type
// of this field [optional pointer][collection prefix][optional pointer][type
// name]
func FieldTypeDefinition(ctx context.Context, fieldName string, field system.RuleInterface, path string, getAlias func(string) string) (string, error) {
	outer := system.WrapRule(ctx, field)

	outerPointer := ""
	if outer.Pointer(ctx) {
		outerPointer = "*"
	}

	// if the rule is a complex collection, with possibly several maps and
	// arrays, this iterates over the rule and returns the go collection prefix
	// - e.g. []map[string] for an array of maps. It also returns the inner rule.
	prefix, inner, err := collectionPrefixInnerRule(ctx, "", outer)
	if err != nil {
		return "", kerr.Wrap("SOGEFOPJHB", err)
	}

	innerPointer := ""
	// if we have a prefix we should also work out the innerPointer
	if prefix != "" && inner.Pointer(ctx) {
		innerPointer = "*"
	}

	var n string
	if inner.Struct.Interface {
		n = system.GoInterfaceName(inner.Parent.Id.Name)
	} else {
		n = system.GoName(inner.Parent.Id.Name)
	}
	name := Reference(inner.Parent.Id.Package, n, path, getAlias)

	tag := getTag(fieldName)
	if tag != "" {
		tag = " " + tag
	}

	return fmt.Sprint(outerPointer, prefix, innerPointer, name, tag), nil
}

// collectionPrefix recursively digs down through collection rules, recursively
// calling itself as long as it finds a collection rule (map or array). It returns
// the full collection prefix (e.g. any number of appended [] and map[string]'s)
// and the inner (non collection) rule.
func collectionPrefixInnerRule(ctx context.Context, prefix string, outer *system.RuleWrapper) (fullPrefix string, inner *system.RuleWrapper, err error) {

	kind, alias := outer.Kind(ctx)
	if alias {
		return prefix, outer, nil
	}
	switch kind {
	case system.KindValue, system.KindStruct, system.KindInterface:
		return prefix, outer, nil
	case system.KindArray:
		prefix += "[]"
	case system.KindMap:
		prefix += "map[string]"
	default:
		panic("unknown kind")
	}

	items, err := outer.ItemsRule()
	if err != nil {
		return "", nil, kerr.Wrap("SUTYJEGBKW", err)
	}
	return collectionPrefixInnerRule(ctx, prefix, items)
}

func formatTag(fieldName string) string {

	tag := ""
	tag = addSubTag(tag, "json", fieldName)

	if tag == "" {
		return ""
	}
	if strconv.CanBackquote(tag) {
		return "`" + tag + "`"
	}
	return strconv.Quote(tag)
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

func getTag(fieldName string) string {

	return formatTag(fieldName)

	/*
		dr, ok := r.Interface.(system.DefaultRule)
		if !ok {
			// Doesn't have a default field
			return formatTag(ctx, fieldName)
		}

		d := dr.GetDefault()
		if d == nil {
			return formatTag(ctx, fieldName)
		}

		// If we have a marshaler, we have to call it manually
		if rp, ok := d.(packer.Repacker); ok {
			i, err := rp.Repack(ctx)
			if err != nil {
				return "", kerr.Wrap("YIEMHYFVCD", err)
			}
			defaultBytes, err := json.Marshal(i)
			if err != nil {
				return "", kerr.Wrap("OFWRBEMHAL", err)
			}
			return formatTag(ctx, fieldName, defaultBytes, r)
		}

		defaultBytes, err := json.Marshal(d)
		if err != nil {
			return "", kerr.Wrap("QQDOLAJKLU", err)
		}

		return formatTag(ctx, fieldName, defaultBytes, r)
	*/
}
