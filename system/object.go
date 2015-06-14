package system

// SetContext satisfies the json.Contexter interface, which allows the json unmarshal
// function to store the unmarshal context in every object.
func (o *Object) SetContext(path string, imports map[string]string) {
	o.Context = &Context{Package: path, Imports: imports}
}

func (o *Object) GetType() (*Type, bool) {
	return o.Type.GetType()
}
func (o *Object) GetTypeReference() Reference {
	return o.Type
}

type Typer interface {
	GetType() (*Type, bool)
	GetTypeReference() Reference
}

type Ruler interface {
	GetRules() []Rule
	RulesApply() rulesApplication
}

func (o *Object) GetRules() []Rule {
	return []Rule{}
}

func (o *Object) RulesApply() rulesApplication {
	if o.Type.Value == "kego.io/system:type" || o.Type.Value == "kego.io/system:property" {
		return RULES_APPLY_TO_TYPES
	} else if o.Type.Type[0:0] == "@" {
		return RULES_APPLY_TO_TYPES
	}
	return RULES_APPLY_TO_OBJECTS
}

type rulesApplication string

const (
	RULES_APPLY_TO_TYPES rulesApplication = "types"
	RULES_APPLY_TO_OBJECTS 				  = "objects"
)