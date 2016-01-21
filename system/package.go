package system

func EmptyPackage() *Package {
	return &Package{
		Object: &Object{
			Type: NewReference("kego.io/system", "package"),
		},
		Aliases:   map[string]string{},
		Recursive: false,
	}
}
