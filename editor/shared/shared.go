package shared // import "frizz.io/editor/shared"

// frizz: {"package": {"jstest": true}}

type Info struct {
	Path            string                // Package path
	Aliases         map[string]string     // Map of alias:path
	Data            map[string]string     // Map of data names and relative file names
	Package         []byte                // Package object
	PackageFilename string                // Filename of the package object
	Imports         map[string]ImportInfo // Flattened list of all imports
	Hash            []byte                // Auth hash that is transmitted on every request
}

type ImportInfo struct {
	Path    string
	Aliases map[string]string
	Types   map[string]TypeInfo
}

type TypeInfo struct {
	File  string
	Bytes []byte
}

type Method string

// Request should be embeddedin all requests, and contains some common data and
// very simple auth that prevents the server being abused.
type Request struct {
	Path string // Package path
	Hash []byte // Hash round-tripped to the server to authenticate requests
}

const Save Method = "Server.Save"

type SaveRequest struct {
	Request
	Files []SaveRequestFile
}

type SaveRequestFile struct {
	File  string
	Bytes []byte
	Hash  uint64
}

type SaveResponse struct {
	Files []SaveResponseFile
}

type SaveResponseFile struct {
	File string
	Hash uint64
}

const Data Method = "Server.Data"

type DataRequest struct {
	Request
	File    string
	Name    string
	Package string
}

type DataResponse struct {
	Data    []byte
	Found   bool
	Name    string
	Package string
}
