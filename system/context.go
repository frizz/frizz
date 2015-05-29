package system

// Unmarshal context.
type Context struct {

	// A list of imports.
	Imports map[string]string

	// The path of the local package.
	Package string
}
