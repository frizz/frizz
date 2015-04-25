package system

// SetContext satisfies the json.Contexter interface, which allows the json unmarshal
// function to store the unmarshal context in every object.
func (o *Object) SetContext(path string, imports map[string]string) {
	o.Context = &Context{Package: NewString(path), Imports: NewStringMap(imports)}
}

func (o *Object) GetType() (*Type, bool) {
	return o.Type.GetType()
}
