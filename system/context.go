package system

// Unmarshal context.
type Context struct {

	// A list of imports.
	Imports map[string]string

	// The path of the local package.
	Package string
}

func (c *Context) Clone() *Context {
	new := &Context{}
	new.Package = c.Package
	new.Imports = map[string]string{}
	for k, v := range c.Imports {
		new.Imports[k] = v
	}
	return new
}

/*
type Node struct {
	Value    interface{}
	Object   *Object
	Parent   *Node
	Key      string
	Idx      int
	Siblings int
	// Do we need the json?
	//json   *json.Json
}
*/
