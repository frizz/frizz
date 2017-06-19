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
	if err := Generate(buf, env, path, dir); err != nil {
		return errors.WithMessage(err, "error generating source")
	}

	if err := ioutil.WriteFile(filepath.Join(dir, "packer.frizz.go"), buf.Bytes(), 0777); err != nil {
		return errors.Wrap(err, "error saving source")
	}
	return nil
}

func Generate(writer io.Writer, env vos.Env, path string, dir string) error {

	f := NewFilePath(path)

	prog := &progDef{
		fset:   token.NewFileSet(),
		path:   path,
		pcache: patsy.NewCache(env),
		dir:    dir,
	}

	pkgs, err := parser.ParseDir(prog.fset, prog.dir, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "error parsing Go code in package %s", path)
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
		// notest
		return errors.Errorf("can't find Go code in package %s", path)
	}
	if len(filtered) > 1 {
		// notest
		return errors.Errorf("found more than one package in %s", path)
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
			if strings.HasSuffix(fname, "packer.frizz.go") {
				// ignore everything in the generated file
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
				// check the signature
				// Unpack: (*frizz.Root, frizz.Stack, interface{}) error
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
					if len(n.Type.Results.List) != 1 {
						return true
					}
					ide, ok := n.Type.Results.List[0].Type.(*ast.Ident)
					if !ok {
						return true
					}
					if ide.Name != "error" {
						return true
					}
					// unpack method is the correct signature.
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
					if len(n.Type.Results.List) != 2 {
						return true
					}
					ift, ok := n.Type.Results.List[0].Type.(*ast.InterfaceType)
					if !ok {
						return true
					}
					if len(ift.Methods.List) != 0 {
						return true
					}
					ide, ok := n.Type.Results.List[1].Type.(*ast.Ident)
					if !ok {
						return true
					}
					if ide.Name != "error" {
						return true
					}
					// repack method is the correct signature.
					repackers[receiverType] = true
				}

			case *ast.GenDecl:
				var found bool
				if n.Doc == nil {
					return true
				}
				for _, c := range n.Doc.List {
					if c.Text == "// frizz" {
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
			}

			return true
		})
		if outer != nil {
			return errors.WithMessage(outer, "error parsing Go source")
		}
		for _, td := range all {
			_, unpacker := unpackers[td.name]
			_, repacker := repackers[td.name]
			if unpacker && repacker {
				td.packable = true
			}
		}
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
		func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, error) {
			switch name {
			<types...>
			case "<name>":
				return p.Unpack<name>(root, stack, in)
			}
			</types>
			return nil, errors.Errorf("%s: type %s not found", stack, "<name>")
		}
	*/
	f.Func().Params(Id("p").Id("packer")).Id("Unpack").Params(
		Id("root").Op("*").Qual("frizz.io/frizz", "Root"),
		Id("stack").Qual("frizz.io/frizz", "Stack"),
		Id("in").Interface(),
		Id("name").String(),
	).Params(
		Interface(),
		Error(),
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

	if err := f.Render(writer); err != nil {
		return errors.Wrap(err, "error saving generated file")
	}
	return nil
}
