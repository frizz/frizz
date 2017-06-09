package frizz

func New(path string) *Unpacker {
	return &Unpacker{
		Path:     path,
		Registry: DefaultRegistry,
	}
}

type Unpacker struct {
	Path     string
	Registry *Registry
}

type Root struct {
	*Unpacker
	Imports map[string]string
}
