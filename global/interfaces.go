package global

type Packable interface {
	Unpack(context DataContext, in interface{}) (null bool, err error)
	Repack(context DataContext) (value interface{}, dict bool, null bool, err error)
}

type Package interface {
	Path() string
	Unpack(context DataContext, in interface{}, name string) (value interface{}, null bool, err error)
	Repack(context DataContext, in interface{}, name string) (value interface{}, dict bool, null bool, err error)
	GetType(name string) string
	GetData(name string) string
	GetImportedPackages(packages map[string]Package)
}

// PackageContext describes the package, the package path, all the packages that are directly imported, and the
// flattened package import tree
type PackageContext interface {
	Path() string
	Register(packages ...Package)
	Get(path string) Package
}

// RootContext describes a single json file that includes a "_import" field which maps the aliases to package paths.
type RootContext interface {
	Package() PackageContext
	RegisterImport(path string) string
	ParseImports(v interface{}) error
	Imports() map[string]string
}

type ValidationContext interface {
	Package() PackageContext
	Location() Location
}

type DataContext interface {
	Package() PackageContext
	Root() RootContext
	Location() Location
	New(location Location) DataContext
}

type Location interface {
	String() string
	Child(item stackItem) Location
}

type Loader interface {
	Load(pkg Package, filename string) interface{}
}
