package system

type Rule interface{}

type CollectionType interface {
	CollectionTypePrefix() string
	CollectionItemsRule() Rule
}

func (m *Map_rule) CollectionTypePrefix() string {
	return "map[string]"
}

func (m *Map_rule) CollectionItemsRule() Rule {
	return m.Items
}

func (a *Array_rule) CollectionTypePrefix() string {
	return "[]"
}

func (a *Array_rule) CollectionItemsRule() Rule {
	return a.Items
}

type NativeType interface {
	NativeType() string
}

func (s *String_rule) NativeType() string {
	return "string"
}
func (n *Number_rule) NativeType() string {
	return "number"
}
func (b *Bool_rule) NativeType() string {
	return "bool"
}
func (r *Reference_rule) NativeType() string {
	return "string"
}

type InterfaceType interface {
	IsInterface() bool
}

func (r *Rule_rule) IsInterface() bool {
	return true
}
