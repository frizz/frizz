package frizz

func New(path string, packers ...Packer) *Context {
	m := make(map[string]Packer, len(packers))
	for _, p := range packers {
		m[p.Path()] = p
	}
	return &Context{
		Path:    path,
		Packers: m,
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
