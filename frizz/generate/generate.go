package generate

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"path/filepath"

	"strings"

	"encoding/json"

	"encoding/base64"

	"os"

	"frizz.io/frizz"
	. "github.com/dave/jennifer/jen"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/pkg/errors"
)

func Save(path string) error {
	// notest
	env := vos.Os()

	dir, err := patsy.Dir(env, path)
	if err != nil {
		return errors.Wrapf(err, "can't find dir for package %s", path)
	}

	buf := &bytes.Buffer{}
	hasContent, err := Generate(buf, env, path, dir)
	if err != nil {
		return errors.WithMessage(err, "error generating source")
	}

	fname := filepath.Join(dir, "generated.frizz.go")
	if hasContent {
		if err := ioutil.WriteFile(fname, buf.Bytes(), 0777); err != nil {
			return errors.Wrap(err, "saving source")
		}
	} else {
		if _, err := os.Stat(fname); err == nil {
			if err := os.Remove(fname); err != nil {
				return errors.Wrap(err, "removing source")
			}
		}
	}

	return nil
}

func Generate(writer io.Writer, env vos.Env, path string, dir string) (bool, error) {

	f := NewFilePath(path)

	prog := &progDef{
		fset:        token.NewFileSet(),
		path:        path,
		pcache:      patsy.NewCache(env),
		dir:         dir,
		annotations: map[string]string{},
	}

	var err error
	var packers, types bool

	if packers, err = generatePackers(prog, f); err != nil {
		return false, err
	}

	if types, err = generateTypes(prog, f); err != nil {
		return false, err
	}

	if !packers && !types {
		// no source
		return false, nil
	}

	if err := f.Render(writer); err != nil {
		return false, errors.Wrap(err, "error saving generated file")
	}
	return true, nil
}

func generateTypes(prog *progDef, f *File) (bool, error) {
	types := map[string][]byte{}
	for name, annotation := range prog.annotations {
		fpath := filepath.Join(prog.dir, annotation)

		// sanity check to make sure the filename in the annotation hasn't escaped from the package dir with ..'s
		fdir, fname := filepath.Split(fpath)
		// TODO: needed to join to "." to get the paths to be exactly the same... Better way?
		if filepath.Join(fdir, ".") != filepath.Join(prog.dir, ".") {
			return false, errors.Errorf("type filename %s points to a location outside the package dir %s %s", annotation)
		}

		if !strings.HasSuffix(fname, ".frizz.json") {
			return false, errors.Errorf("type filename %s should end .frizz.json", annotation)
		}

		s, b, err := decodeStub(fpath)
		if err != nil {
			return false, errors.Wrapf(err, "decoding %s", annotation)
		}

		tpath, tname, err := frizz.ParseReference(s.Type, prog.path, s.Import)

		if tpath == "frizz.io/system" && tname == "Type" {
			types[name] = b
		}
	}

	if len(types) == 0 {
		return false, nil
	}

	/*
		const Types types = 0

		type types int

		func (t types) Path() string {
			return "<path>"
		}
	*/
	f.Const().Id("Types").Id("types").Op("=").Lit(0)
	f.Type().Id("types").Int()
	f.Func().Params(Id("t").Id("types")).Id("Path").Params().String().Block(
		Return(Lit(prog.path)),
	)
	/*
		func (t types) Get(name string) string {
			switch name {
			<types...>
			case <name>:
				return <bytes>
			</types>
			}
			return nil
		}
	*/
	f.Func().Params(
		Id("t").Id("types"),
	).Id("Get").Params(
		Id("name").String(),
	).String().BlockFunc(func(g *Group) {
		if len(types) > 0 {
			g.Switch(Id("name")).BlockFunc(func(g *Group) {
				for name, data := range types {
					g.Case(Lit(name)).Block(
						Return(Lit(base64.StdEncoding.EncodeToString([]byte(data)))),
					)
				}
			})
		}
		g.Return(Nil())
	})

	return true, nil

}

type stub struct {
	Import map[string]string `json:"_import"`
	Type   string            `json:"_type"`
}

func decodeStub(fpath string) (stub, []byte, error) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return stub{}, nil, err
	}
	var v stub
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	if err := d.Decode(&v); err != nil {
		return stub{}, nil, err
	}
	return v, b, nil
}

func generatePackers(prog *progDef, f *File) (bool, error) {
	pkgs, err := parser.ParseDir(prog.fset, prog.dir, nil, parser.ParseComments)
	if err != nil {
		return false, errors.Wrapf(err, "error parsing Go code in package %s", prog.path)
	}
	// ignore any test package
	var filtered []*ast.Package
	for name, pkg := range pkgs {
		if strings.HasSuffix(name, "_test") {
			// notest
			continue
		}
		filtered = append(filtered, pkg)
	}
	if len(filtered) == 0 {
		// data only package?
		return false, nil
	}
	if len(filtered) > 1 {
		// notest
		return false, errors.Errorf("found more than one package in %s", prog.path)
	}
	pkg := filtered[0]

	var all []*typeDef
	var unpackers = map[string]bool{}
	var repackers = map[string]bool{}
	for fname, f := range pkg.Files {
		var file *fileDef // lazy init because not all files need action
		var outer error
		ast.Inspect(f, func(n ast.Node) bool {
			if err != nil {
				// as soon as error is detected, stop processing
				return false
			}
			if strings.HasSuffix(fname, ".frizz.go") {
				// ignore everything in the generated files
				return false
			}
			switch n := n.(type) {
			case nil:
				return true
			case *ast.FuncDecl:
				if n.Name.Name != "Unpack" && n.Name.Name != "Repack" {
					return true
				}
				if n.Recv == nil {
					// function not method
					return true
				}
				if len(n.Recv.List) != 1 {
					return true // impossible?
				}
				s, ok := n.Recv.List[0].Type.(*ast.StarExpr)
				if !ok {
					return true
				}
				id, ok := s.X.(*ast.Ident)
				if !ok {
					return true
				}
				receiverType := id.Name

				if file == nil {
					file, err = prog.newFileDef(f)
					if err != nil {
						outer = err
						return false
					}
				}
				// We check the signature of "Unpack" and "Repack" methods to work out if the receiver implements the
				// frizz.Packable interface.
				switch n.Name.Name {
				case "Unpack":
					if len(n.Type.Params.List) != 3 {
						return true
					}
					param0 := file.getTypeInfo(n.Type.Params.List[0].Type)
					param1 := file.getTypeInfo(n.Type.Params.List[1].Type)
					if param0 != (typeInfo{true, "frizz.io/frizz", "Root"}) {
						return true
					}
					if param1 != (typeInfo{false, "frizz.io/frizz", "Stack"}) {
						return true
					}
					ift, ok := n.Type.Params.List[2].Type.(*ast.InterfaceType)
					if !ok {
						return true
					}
					if len(ift.Methods.List) != 0 {
						return true
					}

					// check return is error
					if len(n.Type.Results.List) != 2 {
						return true
					}
					if ide, ok := n.Type.Results.List[0].Type.(*ast.Ident); !ok || ide.Name != "bool" {
						return true
					}
					if ide, ok := n.Type.Results.List[1].Type.(*ast.Ident); !ok || ide.Name != "error" {
						return true
					}
					// Unpack method is the correct signature.
					unpackers[receiverType] = true
				case "Repack":
					if len(n.Type.Params.List) != 2 {
						return true
					}
					param0 := file.getTypeInfo(n.Type.Params.List[0].Type)
					param1 := file.getTypeInfo(n.Type.Params.List[1].Type)
					if param0 != (typeInfo{true, "frizz.io/frizz", "Root"}) {
						return true
					}
					if param1 != (typeInfo{false, "frizz.io/frizz", "Stack"}) {
						return true
					}

					// check return is error
					if len(n.Type.Results.List) != 4 {
						return true
					}
					ift, ok := n.Type.Results.List[0].Type.(*ast.InterfaceType)
					if !ok {
						return true
					}
					if len(ift.Methods.List) != 0 {
						return true
					}
					if ide, ok := n.Type.Results.List[1].Type.(*ast.Ident); !ok || ide.Name != "bool" {
						return true
					}
					if ide, ok := n.Type.Results.List[2].Type.(*ast.Ident); !ok || ide.Name != "bool" {
						return true
					}
					if ide, ok := n.Type.Results.List[3].Type.(*ast.Ident); !ok || ide.Name != "error" {
						return true
					}
					// Repack method is the correct signature.
					repackers[receiverType] = true
				}

			case *ast.GenDecl:
				var found bool
				var annotation interface{}
				if n.Doc == nil {
					return true
				}
				for _, c := range n.Doc.List {
					if c.Text == "// frizz" {
						found = true
						break
					}
					if strings.HasPrefix(c.Text, "// frizz: ") {
						s := strings.TrimPrefix(c.Text, "// frizz: ")
						if err := json.Unmarshal([]byte(s), &annotation); err != nil {
							outer = errors.Wrapf(err, "annotation %s must be well formed json", s)
							return false
						}
						found = true
						break
					}
				}
				if !found {
					return true
				}
				if n.Tok != token.TYPE {
					outer = errors.Errorf("unsupported token tagged with frizz comment at %s", prog.fset.Position(n.TokPos))
					return false
				}
				if len(n.Specs) != 1 {
					outer = errors.Errorf("must be a single spec in type definition at %s", prog.fset.Position(n.TokPos))
					return false
				}
				ts, ok := n.Specs[0].(*ast.TypeSpec)
				if !ok {
					outer = errors.Errorf("must be type spec at %s", prog.fset.Position(n.TokPos))
					return false
				}
				if file == nil {
					file, err = prog.newFileDef(f)
					if err != nil {
						outer = err
						return false
					}
				}
				all = append(all, &typeDef{
					fileDef: file,
					name:    ts.Name.Name,
					spec:    ts.Type,
				})
				if a, ok := annotation.(string); ok {
					// for now, only support annotations as strings
					prog.annotations[ts.Name.Name] = a
				}
			}

			return true
		})
		if outer != nil {
			return false, errors.WithMessage(outer, "error parsing Go source")
		}
		for _, td := range all {
			_, unpacker := unpackers[td.name]
			_, repacker := repackers[td.name]
			if unpacker && repacker {
				td.packable = true
			}
		}
	}

	if len(all) == 0 {
		return false, nil
	}

	/*
		const Packer packer = 0

		type packer int

		func (p packer) Path() string {
			return "<path>"
		}
	*/
	f.Const().Id("Packer").Id("packer").Op("=").Lit(0)
	f.Type().Id("packer").Int()
	f.Func().Params(Id("p").Id("packer")).Id("Path").Params().String().Block(
		Return(Lit(prog.path)),
	)
	/*
		func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
			switch name {
			<types...>
			case "<name>":
				return p.Unpack<name>(root, stack, in)
			}
			</types>
			return nil, false, errors.Errorf("%s: type %s not found", stack, "<name>")
		}
	*/
	f.Func().Params(Id("p").Id("packer")).Id("Unpack").Params(
		Id("root").Op("*").Qual("frizz.io/frizz", "Root"),
		Id("stack").Qual("frizz.io/frizz", "Stack"),
		Id("in").Interface(),
		Id("name").String(),
	).Params(
		Id("value").Interface(),
		Id("null").Bool(),
		Err().Error(),
	).Block(
		Switch(Id("name")).BlockFunc(func(g *Group) {
			for _, t := range all {
				g.Case(Lit(t.name)).Block(
					Return(Id("p").Dot("Unpack"+t.name).Call(Id("root"), Id("stack"), Id("in"))),
				)
			}
		}),
		Return(
			Nil(),
			False(),
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("%s: type %s not found"),
				Id("stack"),
				Id("name"),
			),
		),
	)

	for _, t := range all {
		f.Add(t.unpacker(t.spec, t.name, true, t.packable))
	}

	/*
		func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
			switch name {
			<types...>
			case "<name>":
				return p.Repack<name>(root, stack, in.(<name>))
			}
			</types>
			return nil, false, false, errors.Errorf("%s: type %s not found", stack, "<name>")
		}
	*/
	f.Func().Params(Id("p").Id("packer")).Id("Repack").Params(
		Id("root").Op("*").Qual("frizz.io/frizz", "Root"),
		Id("stack").Qual("frizz.io/frizz", "Stack"),
		Id("in").Interface(),
		Id("name").String(),
	).Params(
		Id("value").Interface(),
		Id("dict").Bool(),
		Id("null").Bool(),
		Err().Error(),
	).Block(
		Switch(Id("name")).BlockFunc(func(g *Group) {
			for _, t := range all {
				g.Case(Lit(t.name)).Block(
					Return(Id("p").Dot("Repack"+t.name).Call(Id("root"), Id("stack"), Id("in").Assert(Id(t.name)))),
				)
			}
		}),
		Return(
			Nil(),
			False(),
			False(),
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("%s: type %s not found"),
				Id("stack"),
				Id("name"),
			),
		),
	)

	for _, t := range all {
		f.Add(t.repacker(t.spec, t.name, true, t.packable))
	}
	return true, nil
}
