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

	"sort"

	"frizz.io/utils"
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
		fset:      token.NewFileSet(),
		path:      path,
		pcache:    patsy.NewCache(env),
		dir:       dir,
		types:     map[string]*typeDef{},
		typeFiles: map[string][]byte{},
		stubs:     map[string]*stub{},
		imports:   map[string]bool{},
	}

	if err := scanData(prog); err != nil {
		return false, err
	}

	if err := scanSource(prog); err != nil {
		return false, err
	}

	if len(prog.types) == 0 && len(prog.typeFiles) == 0 && len(prog.imports) == 0 {
		return false, nil
	}

	generatePackage(prog, f)
	generatePackers(prog, f)
	generateTypeFiles(prog, f)
	generateImports(prog, f)

	if err := f.Render(writer); err != nil {
		return false, errors.Wrap(err, "error saving generated file")
	}
	return true, nil
}

func scanData(prog *progDef) error {

	// scan all .frizz.json files for stubs

	files, err := filepath.Glob(filepath.Join(prog.dir, "*.frizz.json"))
	if err != nil {
		return errors.Wrapf(err, "finding files in %s", prog.dir)
	}

	for _, fpath := range files {
		s, err := decodeStub(fpath)
		if err != nil {
			return errors.Wrapf(err, "decoding %s", fpath)
		}
		prog.stubs[fpath] = s
	}

	for _, stub := range prog.stubs {
		if stub.Import == nil {
			continue
		}
		for _, pkg := range stub.Import {
			prog.imports[pkg] = true
		}
	}

	return nil
}

func scanSource(prog *progDef) error {

	// scan all .go files for packers and type files referenced by annotations

	pkgs, err := parser.ParseDir(prog.fset, prog.dir, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "parsing Go code in package %s", prog.path)
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
		return nil
	}
	if len(filtered) > 1 {
		// notest
		return errors.Errorf("found more than one package in %s", prog.path)
	}
	pkg := filtered[0]

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
					if len(n.Type.Params.List) != 2 {
						return true
					}
					param0 := file.getTypeInfo(n.Type.Params.List[0].Type)
					if param0 != (typeInfo{false, "frizz.io/global", "DataContext"}) {
						return true
					}
					if ift, ok := n.Type.Params.List[1].Type.(*ast.InterfaceType); !ok || len(ift.Methods.List) != 0 {
						return true
					}
					// check return is bool, error
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
					if len(n.Type.Params.List) != 1 {
						return true
					}
					param0 := file.getTypeInfo(n.Type.Params.List[0].Type)
					if param0 != (typeInfo{false, "frizz.io/global", "DataContext"}) {
						return true
					}
					// check return is error
					if len(n.Type.Results.List) != 4 {
						return true
					}
					if ift, ok := n.Type.Results.List[0].Type.(*ast.InterfaceType); !ok || len(ift.Methods.List) != 0 {
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
				prog.types[ts.Name.Name] = &typeDef{
					fileDef: file,
					name:    ts.Name.Name,
					spec:    ts.Type,
				}
				if a, ok := annotation.(string); ok {

					// for now, only support annotations as strings

					fpath := filepath.Join(prog.dir, a)

					s, ok := prog.stubs[fpath]
					if !ok {
						outer = errors.Errorf("can't find type file %s in annotation", s)
						return false
					}

					tpath, tname, err := utils.ParseReference(s.Type, prog.path, s.Import)
					if err != nil {
						outer = err
						return false
					}

					if tpath != "frizz.io/system" || tname != "Type" {
						outer = errors.Errorf("expected frizz.Type, got: %s", s.Type)
					}

					prog.typeFiles[ts.Name.Name] = s.bytes

				}
			}

			return true
		})
		if outer != nil {
			return errors.WithMessage(outer, "error parsing Go source")
		}
		for _, td := range prog.types {
			_, unpacker := unpackers[td.name]
			_, repacker := repackers[td.name]
			if unpacker && repacker {
				td.packable = true
			}
		}
	}

	return nil
}

func generatePackage(prog *progDef, f *File) {

	/*
		const Package packageType = 0

		type packageType int

		func (p packageType) Path() string {
			return "<path>"
		}
	*/
	f.Const().Id("Package").Id("packageType").Op("=").Lit(0)
	f.Type().Id("packageType").Int()
	f.Func().Params(Id("p").Id("packageType")).Id("Path").Params().String().Block(
		Return(Lit(prog.path)),
	)

}

func generatePackers(prog *progDef, f *File) {

	var sorted []*typeDef
	for _, t := range prog.types {
		sorted = append(sorted, t)
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].name < sorted[j].name })

	/*
		func (p packageType) Unpack(context global.DataContext, in interface{}, name string) (value interface{}, null bool, err error) {
			switch name {
			<types...>
			case "<name>":
				return p.Unpack<name>(context, in)
			}
			</types>
			return nil, false, errors.Errorf("%s: type %s not found", context.Location(), "<name>")
		}
	*/
	f.Func().Params(Id("p").Id("packageType")).Id("Unpack").Params(
		Id("context").Qual("frizz.io/global", "DataContext"),
		Id("in").Interface(),
		Id("name").String(),
	).Params(
		Id("value").Interface(),
		Id("null").Bool(),
		Err().Error(),
	).Block(
		Switch(Id("name")).BlockFunc(func(g *Group) {
			for _, t := range sorted {
				g.Case(Lit(t.name)).Block(
					Return(Id("p").Dot("Unpack"+t.name).Call(
						Id("context"),
						Id("in"),
					)),
				)
			}
		}),
		Return(
			Nil(),
			False(),
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("%s: type %s not found"),
				Id("context").Dot("Location").Call(),
				Id("name"),
			),
		),
	)

	for _, t := range sorted {
		f.Add(t.unpacker(t.spec, t.name, true, t.packable))
	}

	/*
		func (p packageType) Repack(context global.DataContext, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
			switch name {
			<types...>
			case "<name>":
				return p.Repack<name>(context, in.(<name>))
			}
			</types>
			return nil, false, false, errors.Errorf("%s: type %s not found", context.Location(), "<name>")
		}
	*/
	f.Func().Params(Id("p").Id("packageType")).Id("Repack").Params(
		Id("context").Qual("frizz.io/global", "DataContext"),
		Id("in").Interface(),
		Id("name").String(),
	).Params(
		Id("value").Interface(),
		Id("dict").Bool(),
		Id("null").Bool(),
		Err().Error(),
	).Block(
		Switch(Id("name")).BlockFunc(func(g *Group) {
			for _, t := range sorted {
				g.Case(Lit(t.name)).Block(
					Return(Id("p").Dot("Repack"+t.name).Call(
						Id("context"),
						Id("in").Assert(Id(t.name)),
					)),
				)
			}
		}),
		Return(
			Nil(),
			False(),
			False(),
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("%s: type %s not found"),
				Id("context").Dot("Location").Call(),
				Id("name"),
			),
		),
	)

	for _, t := range sorted {
		f.Add(t.repacker(t.spec, t.name, true, t.packable))
	}
}

func generateTypeFiles(prog *progDef, f *File) {

	type def struct {
		name string
		data []byte
	}
	var sorted []def
	for name, data := range prog.typeFiles {
		sorted = append(sorted, def{name, data})
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].name < sorted[j].name })

	/*
		func (p packageType) GetType(name string) string {
			switch name {
			<types...>
			case <name>:
				return <bytes>
			</types>
			}
			return ""
		}
	*/
	f.Func().Params(
		Id("p").Id("packageType"),
	).Id("GetType").Params(
		Id("name").String(),
	).String().BlockFunc(func(g *Group) {
		if len(sorted) > 0 {
			g.Switch(Id("name")).BlockFunc(func(g *Group) {
				for _, d := range sorted {
					g.Case(Lit(d.name)).Block(
						Return(Lit(base64.StdEncoding.EncodeToString([]byte(d.data)))),
					)
				}
			})
		}
		g.Return(Lit(""))
	})

}

func generateImports(prog *progDef, f *File) {

	var sorted []string
	for pkg := range prog.imports {
		sorted = append(sorted, pkg)
	}
	sort.Strings(sorted)

	/*
		func (p packageType) GetImportedPackages(packages map[string]global.Package) {
			packages["<path>"] = Package
			<imports...>
			<pkg>.Package.Add(packages)
			</imports>
		}
	*/
	f.Func().Params(Id("p").Id("packageType")).Id("GetImportedPackages").Params(
		Id("packages").Map(String()).Qual("frizz.io/global", "Package"),
	).BlockFunc(func(g *Group) {
		if len(prog.types) > 0 || len(prog.typeFiles) > 0 {
			g.Id("packages").Index(Lit(prog.path)).Op("=").Id("Package")
		}
		for _, pkg := range sorted {
			g.Qual(pkg, "Package").Dot("GetImportedPackages").Call(Id("packages"))
		}
	})

}

type stub struct {
	Import map[string]string `json:"_import"`
	Type   string            `json:"_type"`
	bytes  []byte
}

func decodeStub(fpath string) (*stub, error) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	v := &stub{}
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	if err := d.Decode(v); err != nil {
		return nil, err
	}
	v.bytes = b
	return v, nil
}
