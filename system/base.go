package system

type Object interface {
	GetBase() *Base
}

func (b *Base) GetBase() *Base {
	if b == nil {
		return &Base{}
	}
	return b
}

// SetContext satisfies the json.Contexter interface, which allows the json unmarshal
// function to store the unmarshal context in every object.
func (o *Base) SetContext(path string, imports map[string]string) {
	o.Context = &Context{Package: path, Imports: imports}
}

// RulesApplyToObjects returns true if we should apply the rules collection to this object.
// Rules on types apply to instances of that type, not to the actual type object.
// Rules on rules apply to objects defined by that rule, not to the actual rule object.
// Rules on other objects apply to that object.
func RulesApplyToObjects(object interface{}) bool {
	_, isRule := object.(Rule)
	_, isType := object.(*Type)
	_, isObject := object.(Object)
	return !isRule && !isType && isObject
}
