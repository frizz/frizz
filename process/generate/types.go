package generate

/*
func Types(ctx context.Context) (source []byte, err error) {

	env := envctx.FromContext(ctx)

	types := system.GetAllGlobalsInPackage(env.Path, system.NewReference("kego.io/system", "type"))
	typesPath := fmt.Sprintf("%s/types", env.Path)
	g := generator.New(typesPath)

	if len(types) == 0 {
		b, err := g.Build()
		if err != nil {
			return nil, kerr.New("JMHUIFWPWV", err, "Build")
		}
		return b, nil
	}
	if typesPath != "kego.io/system/types" {
		g.Imports.Anonymous("kego.io/system/types")
	}
	for p, _ := range env.Aliases {
		g.Imports.Anonymous(fmt.Sprint(p, "/types"))
	}
	g.Println("func init() {")
	{
		registers := []string{}
		pointers := []literal.Pointer{}
		systemRegisterType := generator.Reference("kego.io/system", "Register", typesPath, g.Imports.Add)
		for _, hashed := range types {
			t := hashed.Object.(*system.Type)
			pkg := strconv.Quote(t.Id.Package)
			name := strconv.Quote(t.Id.Name)
			pointer := literal.Build(t, &pointers, typesPath, g.Imports.Add)
			// e.g.
			// system.Register("kego.io/gallery/data", "gallery", ptr8728815248)
			// system.Register("kego.io/gallery/data", "@gallery", ptr8728815360)
			line := fmt.Sprintf("%s(%s, %s, %s, %#v)", systemRegisterType, pkg, name, pointer.Name, hashed.Hash)
			registers = append(registers, line)
		}
		for _, p := range pointers {
			g.Println(p.Name, " := ", p.Source)
		}
		for _, s := range registers {
			g.Println(s)
		}
	}
	g.Println("}")

	b, err := g.Build()
	if err != nil {
		return nil, kerr.New("BATUJJFGHT", err, "Build")
	}
	return b, nil
}
*/
