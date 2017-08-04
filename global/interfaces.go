package global

type Packable interface {
	Unpack(context Context, root Root, stack Stack, in interface{}) (null bool, err error)
	Repack(context Context, root Root, stack Stack) (value interface{}, dict bool, null bool, err error)
}

type Package interface {
	Path() string
	Unpack(context Context, root Root, stack Stack, in interface{}, name string) (value interface{}, null bool, err error)
	Repack(context Context, root Root, stack Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error)
	Get(name string) string
	Add(packages map[string]Package)
}

type Root interface {
	Parse(value interface{}) error
	Register(path string) string
	Imports() map[string]string
}

type Context interface {
	Path() string
	Register(packages ...Package)
	Package(path string) Package
	Unmarshal(in []byte) (interface{}, error)
}
