package system

import "kego.io/json"

// RulesApplyToObjects returns true if we should apply the rules collection to this object.
// Rules on types apply to instances of that type, not to the actual type object.
// Rules on rules apply to objects defined by that rule, not to the actual rule object.
// Rules on other objects apply to that object.
func RulesApplyToObjects(object interface{}) bool {
	_, isRule := object.(RuleInterface)
	_, isType := object.(*Type)
	_, isObject := object.(ObjectInterface)
	return !isRule && !isType && isObject
}

func (b *Object) InitializeType(path string, name string) error {
	if b.Type != nil {
		// We should return an error if we're trying to set the type to a different type
		if path != b.Type.Package || name != b.Type.Name {
			return json.InitializableTypeError{
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
