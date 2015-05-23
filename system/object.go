package system

// SetContext satisfies the json.Contexter interface, which allows the json unmarshal
// function to store the unmarshal context in every object.
func (o *Object) SetContext(path string, imports map[string]string) {
	o.Context = &Context{Package: path, Imports: imports}
}

func (o *Object) GetType() (*Type, bool) {
	return o.Type.GetType()
}

type Typer interface {
	GetType() (*Type, bool)
}
