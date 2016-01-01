package shared // import "kego.io/editor/shared"

type Info struct {
	// Package path
	Path string
	// Map of path:alias
	Aliases map[string]string
	// Array of source names for data
	Data []string
	// Array of source names for types
	Types []string
	// Package object
	Package []byte
	// Flattened list of all imports
	Imports map[string]*ImportInfo
}

type ImportInfo struct {
	Path    string
	Aliases map[string]string
	Types   map[string][]byte
}
