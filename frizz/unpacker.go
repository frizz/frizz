package frizz

func New(path string, packers ...Packer) *Unpacker {
	m := make(map[string]Packer, len(packers))
	for _, p := range packers {
		m[p.Path()] = p
	}
	return &Unpacker{
		Path:    path,
		Packers: m,
	}
}

type Unpacker struct {
	Path    string
	Packers map[string]Packer
}

func (u *Unpacker) Register(p Packer) {
	u.Packers[p.Path()] = p
}

type Root struct {
	*Unpacker
	Imports map[string]string
}

type Packable interface {
	Unpack(*Root, Stack, interface{}) error
	Repack(*Root, Stack) (interface{}, bool, error)
}

type Packer interface {
	Path() string
	Unpack(*Root, Stack, interface{}, string) (interface{}, error)
	Repack(*Root, Stack, interface{}, string) (interface{}, bool, error)
}
