package system

import "context"

type Labelled interface {
	Label(ctx context.Context) string
}

// RulesApplyToObjects returns true if we should apply the rules collection to this object. It
// returns false if the rules collection should not apply to this object (e.g. it should apply to
// instances of this type or objects defined by this rule).
// The rules collection on a type apply to instances of that type, not to the actual type object.
// The rules collection on a rule apply to objects defined by that rule, not to the actual rule
// object.
// The rules on any other object applies to the object itself.
func RulesApplyToObjects(object interface{}) bool {
	_, isRule := object.(RuleInterface)
	_, isType := object.(*Type)
	_, isObject := object.(ObjectInterface)
	return !isRule && !isType && isObject
}

var _ InitializableType = (*Object)(nil)

// InitializeType implements the InitializableType interface. If we are
// unpacking an object into a concrete type defined by the schema, we should
// set the type using this. This enables us to allow the type to be omitted
// from the system.
func (b *Object) InitializeType(path string, name string) error {
	if b.Type != nil {
		// We should return an error if we're trying to set the type to a different type
		if path != b.Type.Package || name != b.Type.Name {
			return InitializableTypeError{
				UnmarshalledPath: b.Type.Package,
				UnmarshalledName: b.Type.Name,
				IntoPath:         path,
				IntoName:         name,
			}
		}
	}
	b.Type = NewReference(path, name)
	return nil
}

type ObjectStub struct {
	Id     *Reference `json:"id"`
	Type   *Reference `json:"type"`
	Export bool       `json:"export"`
}

func (v *ObjectStub) Unpack(ctx context.Context, in Packed, iface bool) error {

	if in == nil || in.Type() == J_NULL {
		return nil
	}

	if field, ok := in.Map()["id"]; ok && field.Type() != J_NULL {
		ob0 := new(Reference)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Id = ob0
	}

	if field, ok := in.Map()["type"]; ok && field.Type() != J_NULL {
		ob0 := new(Reference)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Type = ob0
	}

	if field, ok := in.Map()["export"]; ok && field.Type() != J_NULL {
		ob0, err := UnpackBool(ctx, field)
		if err != nil {
			return err
		}
		v.Export = ob0
	}
	return nil
}
