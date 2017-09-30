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

	"regexp"

	"strconv"

	"frizz.io/generate/scanner"
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
		fset:    token.NewFileSet(),
		path:    path,
		pcache:  patsy.NewCache(env),
		dir:     dir,
		types:   map[string]*typeDef{},
		stubs:   map[string]*stub{},
		imports: map[string]bool{},
	}

	if err := scanData(prog); err != nil {
		return false, err
	}

	if err := scanSource(prog); err != nil {
		return false, err
	}

	if len(prog.types) == 0 && len(prog.stubs) == 0 && len(prog.imports) == 0 {
		return false, nil
	}

	generatePackage(prog, f)
	generatePackers(prog, f)
	generateStubs(prog, f)
	generateTypes(prog, f)
	generateImports(prog, f)
	generateData(prog, f)

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
		s, err := decodeStub(fpath, prog.path)
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

	prog.scanner = scanner.New(prog.path, vos.Os())
	if err := prog.scanner.Scan(); err != nil {
		return err
	}

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

		for _, cg := range f.Comments {
			for _, c := range cg.List {
				if strings.HasPrefix(c.Text, "// frizz-import: ") {
					s, err := strconv.Unquote(strings.TrimPrefix(c.Text, "// frizz-import: "))
					if err != nil {
						return errors.Wrap(err, "parsing frizz-import annotation")
					}
					prog.imports[s] = true
				}
			}
		}

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

					if s.typePackage != "frizz.io/system" || s.typeName != "Type" {
						outer = errors.Errorf("expected system.Type, got: %s", s.Type)
					}

					prog.types[ts.Name.Name].typeFilename = a

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
			</types>
			}
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
	).BlockFunc(func(g *Group) {
		if len(sorted) > 0 {
			g.Switch(Id("name")).BlockFunc(func(g *Group) {
				for _, t := range sorted {
					g.Case(Lit(t.name)).Block(
						Return(Id("p").Dot("Unpack"+t.name).Call(
							Id("context"),
							Id("in"),
						)),
					)
				}
			})
		}
		g.Return(
			Nil(),
			False(),
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("%s: type %s not found"),
				Id("context").Dot("Location").Call(),
				Id("name"),
			),
		)
	})

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
	).BlockFunc(func(g *Group) {
		if len(sorted) > 0 {
			g.Switch(Id("name")).BlockFunc(func(g *Group) {
				for _, t := range sorted {
					g.Case(Lit(t.name)).Block(
						Return(Id("p").Dot("Repack"+t.name).Call(
							Id("context"),
							Id("in").Assert(Id(t.name)),
						)),
					)
				}
			})
		}
		g.Return(
			Nil(),
			False(),
			False(),
			Qual("github.com/pkg/errors", "Errorf").Call(
				Lit("%s: type %s not found"),
				Id("context").Dot("Location").Call(),
				Id("name"),
			),
		)
	})

	for _, t := range sorted {
		f.Add(t.repacker(t.spec, t.name, true, t.packable))
	}
}

func generateStubs(prog *progDef, f *File) {

	type def struct {
		filename string
		stub     *stub
	}
	var sorted []def
	for fpath, stub := range prog.stubs {
		_, fname := filepath.Split(fpath)
		sorted = append(sorted, def{fname, stub})
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].filename < sorted[j].filename })

	/*
		func (p packageType) GetData(filename string) string {
			switch filename {
			<stubs...>
			case <filename>:
				return <bytes>
			</types>
			}
			return ""
		}
	*/
	f.Func().Params(
		Id("p").Id("packageType"),
	).Id("GetData").Params(
		Id("filename").String(),
	).String().BlockFunc(func(g *Group) {
		if len(sorted) > 0 {
			g.Switch(Id("filename")).BlockFunc(func(g *Group) {
				for _, d := range sorted {
					g.Case(Lit(d.filename)).Block(
						Return(Lit(base64.StdEncoding.EncodeToString([]byte(d.stub.bytes)))),
					)
				}
			})
		}
		g.Return(Lit(""))
	})

}

func generateTypes(prog *progDef, f *File) {

	type def struct {
		name     string
		filename string
	}
	var sorted []def
	for name, td := range prog.types {
		if td.typeFilename == "" {
			continue
		}
		sorted = append(sorted, def{name, td.typeFilename})
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].name < sorted[j].name })

	/*
		func (p packageType) GetType(name string) string {
			switch name {
			<types...>
			case <name>:
				return <filename>
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
						Return(Lit(d.filename)),
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
		g.Id("packages").Index(Lit(prog.path)).Op("=").Id("Package")
		for _, pkg := range sorted {
			g.Qual(pkg, "Package").Dot("GetImportedPackages").Call(Id("packages"))
		}
	})

}

func generateData(prog *progDef, f *File) {

	if len(prog.stubs) == 0 {
		return
	}

	/*
		func (p packageType) Loader(loader pack.Loader) dataType {
			return dataType{loader}
		}

		var Data = Package.Loader(pack.DefaultLoader)

		type dataType struct {
			loader pack.Loader
		}

		<stubs...>
		func (d dataType)<name>() (out <type>) {
			return d.loader.Load(Package, <filename>).(<type>)
		}
		<stubs>
	*/
	f.Func().Params(
		Id("p").Id("packageType"),
	).Id("Loader").Params(
		Id("loader").Qual("frizz.io/global", "Loader"),
	).Id("dataType").Block(
		Return(Id("dataType").Values(Id("loader"))),
	)
	f.Var().Id("Data").Op("=").Id("Package").Dot("Loader").Call(
		Qual("frizz.io/pack", "DefaultLoader"),
	)
	f.Type().Id("dataType").Struct(
		Id("loader").Qual("frizz.io/global", "Loader"),
	)
	type def struct {
		filename string
		stub     *stub
	}
	var sorted []def
	for fpath, stub := range prog.stubs {
		_, fname := filepath.Split(fpath)
		sorted = append(sorted, def{fname, stub})
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].filename < sorted[j].filename })
	for _, d := range sorted {
		name := getName(d.filename)
		f.Func().Params(Id("d").Id("dataType")).Id(name).Params().Qual(d.stub.typePackage, d.stub.typeName).Block(
			Return(Id("d").Dot("loader").Dot("Load").Call(
				Id("Package"),
				Lit(d.filename),
			).Assert(Qual(d.stub.typePackage, d.stub.typeName))),
		)
	}

}

var nameRegex = regexp.MustCompile(`[^a-z0-9\-_ ]`)
var initRegex = regexp.MustCompile(`^[^a-z]+`)
var sepRegex = regexp.MustCompile(`[\-_ ]+`)

func getName(filename string) string {
	name := filename
	name = strings.TrimSuffix(name, ".frizz.json")
	name = strings.ToLower(name)
	name = nameRegex.ReplaceAllString(name, "") // replace all illegal characters, leaving separators intact
	name = initRegex.ReplaceAllString(name, "") // replace all illegal starting characters
	name = sepRegex.ReplaceAllString(name, " ") // replaces separators with spaces
	name = strings.Title(name)                  // upper cases first char of each word
	name = strings.Replace(name, " ", "", -1)   // removes spaces
	return name
}

type stub struct {
	Import      map[string]string `json:"_import"`
	Type        string            `json:"_type"`
	bytes       []byte
	typePackage string
	typeName    string
}

func decodeStub(fpath string, packagePath string) (*stub, error) {
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
	tpath, tname, err := utils.ParseReference(v.Type, packagePath, v.Import)
	if err != nil {
		return nil, err
	}
	v.typePackage = tpath
	v.typeName = tname
	v.bytes = b
	return v, nil
}
