package builder

import (
	"context"

	"frizz.io/system"
	"github.com/dave/jennifer/jen"
	"github.com/dave/kerr"
)

// AliasTypeDefinition returns the Go source for the definition of the type
// of this alias type [collection prefix][optional pointer][type name]
func AliasTypeDefinition(ctx context.Context, alias system.RuleInterface) (*jen.Statement, error) {

	s := &jen.Statement{}

	outer := system.WrapRule(ctx, alias)

	// if the rule is a complex collection, with possibly several maps and
	// arrays, this iterates over the rule and returns the go collection prefix
	// - e.g. []map[string] for an array of maps. It also returns the inner
	// rule.
	prefix, inner, err := collectionPrefixInnerRule(ctx, nil, outer)
	if err != nil {
		return nil, kerr.Wrap("BSWTXBHVTH", err)
	}
	s.Add(prefix)

	// if we have a prefix we should also work out the innerPointer
	if prefix != nil && inner.PassedAsPointer(ctx) {
		s.Op("*")
	}

	s.Add(Reference(inner.Parent.Id))

	return s, nil
}

// FieldTypeDefinition returns the Go source for the definition of the type
// of this field [optional pointer][collection prefix][optional pointer][type
// name]
func FieldTypeDefinition(ctx context.Context, fieldName string, field system.RuleInterface) (*jen.Statement, error) {

	s := &jen.Statement{}

	typeDef, err := TypeDefinition(ctx, field)
	if err != nil {
		return nil, kerr.Wrap("BAJJNKKUWI", err)
	}
	s.Add(typeDef)

	if fieldName != "" {
		s.Tag(map[string]string{"json": fieldName})
	}

	return s, nil
}

// TypeDefinition returns the Go source for the definition of the type
// [optional pointer][collection prefix][optional pointer][type name]
func TypeDefinition(ctx context.Context, field system.RuleInterface) (*jen.Statement, error) {

	s := &jen.Statement{}

	outer := system.WrapRule(ctx, field)

	if outer.PassedAsPointer(ctx) {
		s.Op("*")
	}

	// if the rule is a complex collection, with possibly several maps and
	// arrays, this iterates over the rule and returns the go collection prefix
	// - e.g. []map[string] for an array of maps. It also returns the inner rule.
	prefix, inner, err := collectionPrefixInnerRule(ctx, nil, outer)
	if err != nil {
		return nil, kerr.Wrap("SOGEFOPJHB", err)
	}
	s.Add(prefix)

	// if we have a prefix we should also work out the innerPointer
	if prefix != nil && inner.PassedAsPointer(ctx) {
		s.Op("*")
	}

	if inner.Struct.Interface {
		s.Add(InterfaceReference(inner.Parent.Id))
	} else {
		s.Add(Reference(inner.Parent.Id))
	}

	return s, nil
}

// collectionPrefix recursively digs down through collection rules, recursively
// calling itself as long as it finds a collection rule (map or array). It returns
// the full collection prefix (e.g. any number of appended [] and map[string]'s)
// and the inner (non collection) rule.
func collectionPrefixInnerRule(ctx context.Context, prefix *jen.Statement, outer *system.RuleWrapper) (fullPrefix *jen.Statement, inner *system.RuleWrapper, err error) {

	kind, alias := outer.Kind(ctx)
	if alias {
		return prefix, outer, nil
	}
	switch kind {
	case system.KindValue, system.KindStruct, system.KindInterface:
		return prefix, outer, nil
	case system.KindArray:
		if prefix == nil {
			prefix = &jen.Statement{}
		}
		prefix.Index()
	case system.KindMap:
		if prefix == nil {
			prefix = &jen.Statement{}
		}
		prefix.Map(jen.String())
	default:
		panic("unknown kind")
	}

	items, err := outer.ItemsRule()
	if err != nil {
		return nil, nil, kerr.Wrap("SUTYJEGBKW", err)
	}
	return collectionPrefixInnerRule(ctx, prefix, items)
}
