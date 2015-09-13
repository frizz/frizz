package shared

type Info struct {
	// Package path
	Path string
	// Map of path:alias
	Aliases map[string]string
	// Array of global names
	Globals []string
}
