package frizz

func New(imports Importer) *Context {
	packers := map[string]Packer{}
	imports.Add(packers, nil)
	return &Context{
		Path:    imports.Path(),
		Packers: packers,
	}
}

type Context struct {
	Path    string
	Packers map[string]Packer
}

func (u *Context) Register(p Packer) {
	u.Packers[p.Path()] = p
}

type Root struct {
	*Context
	Imports map[string]string
}

type Packable interface {
	Unpack(root *Root, stack Stack, in interface{}) (null bool, err error)
	Repack(root *Root, stack Stack) (value interface{}, dict bool, null bool, err error)
}

type Packer interface {
	Path() string
	Unpack(root *Root, stack Stack, in interface{}, name string) (value interface{}, null bool, err error)
	Repack(root *Root, stack Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error)
}

type Typer interface {
	Path() string
	Get(name string) string
}

type Importer interface {
	Path() string
	Add(map[string]Packer, map[string]Typer)
}
