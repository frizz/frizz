package generate

import (
	"encoding/json"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"frizz.io/generate/jast"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

func scanSource(prog *scanDef) error {

	conf := loader.Config{}
	conf.Import("frizz.io/global") // contains Packable interface
	for _, input := range prog.input {
		// Import all the input packages
		conf.Import(input.path)
	}
	for path := range prog.Packages {
		// Manually import all the packages that have been found while scanning data
		conf.Import(path)
	}
	conf.Build = &build.Default
	conf.Build.GOPATH = prog.env.Getenv("GOPATH")

	// We must exclude all generated files from the scan. This is because when a user makes changes
	// to the signature of a type, the generated files would become invalid and go/types would
	// throw an error when scanning. Additionally, performance is increased by preventing the
	// scanning of the complex generated code. The results of scanning this code will not be needed
	// during the generation process.
	conf.Build.ReadDir = func(dir string) ([]os.FileInfo, error) {
		raw, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		var list []os.FileInfo
		for _, fi := range raw {
			if !strings.HasSuffix(fi.Name(), "generated.frizz.go") {
				list = append(list, fi)
			}
		}
		return list, nil
	}

	// Allow errors to ensure any references to the excluded generated file don't cause errors.
	conf.AllowErrors = true
	conf.ParserMode = parser.ParseComments
	conf.TypeChecker.Error = func(e error) {}

	loaded, err := conf.Load()
	if err != nil {
		return errors.WithStack(err)
	}
	prog.loaded = loaded

	global := loaded.Package("frizz.io/global")
	var packable *types.Interface
	for i, v := range global.Defs {
		if i.Name == "Packable" {
			packable = v.Type().Underlying().(*types.Interface)
			break
		}
	}
	if packable == nil {
		return errors.New("can't find frizz.io/global.Packable")
	}

	for pa, pi := range loaded.AllPackages {
		if len(pi.Files) == 0 {
			continue // virtual stdlib package
		}
		if strings.HasPrefix(loaded.Fset.File(pi.Files[0].Pos()).Name(), conf.Build.GOROOT) {
			continue // stdlib package
		}

		packables := map[string]bool{}
		for i, v := range pi.Defs {
			if v, ok := v.(*types.TypeName); ok {
				if types.Implements(types.NewPointer(v.Type()), packable) {
					packables[i.Name] = true
				}
			}
		}

		typeDefs := map[string]*typeDef{}

		for _, f := range pi.Files {
			for _, d := range f.Decls {
				found, spec, annotation, err := isFrizzDecl(loaded.Fset, d)
				if err != nil {
					return err
				}
				if !found {
					continue
				}
				td := &typeDef{
					name:       spec.Name.Name,
					spec:       spec.Type,
					annotation: annotation,
					packable:   packables[spec.Name.Name],
				}
				typeDefs[spec.Name.Name] = td
			}
		}

		if len(typeDefs) > 0 {
			def, ok := prog.Packages[pa.Path()]
			if !ok {
				dir, _ := filepath.Split(loaded.Fset.File(pi.Files[0].Pos()).Name())
				def = &packageDef{
					scan:    prog,
					Path:    pa.Path(),
					Dir:     dir,
					stubs:   map[string]*stub{},
					imports: map[string]bool{},
				}
				prog.Packages[pa.Path()] = def
			}
			def.types = typeDefs
			def.jast = jast.New(nil, pi.Uses)
			def.info = pi
		}
	}

	for _, def := range prog.Packages {
		if def.info == nil {
			continue
		}
		for _, importPkg := range def.info.Pkg.Imports() {
			// only add to imports if the imported package is a frizz package
			if _, ok := prog.Packages[importPkg.Path()]; ok {
				def.imports[importPkg.Path()] = true
			}
		}
	}

	return nil
}

func isFrizzDecl(fset *token.FileSet, d ast.Decl) (found bool, spec *ast.TypeSpec, annotation string, err error) {
	g, ok := d.(*ast.GenDecl)
	if !ok {
		return false, nil, "", nil
	}
	if g.Doc == nil {
		return false, nil, "", nil
	}
	for _, c := range g.Doc.List {
		if c.Text == "// frizz" {
			found = true
			break
		}
		if strings.HasPrefix(c.Text, "// frizz: ") {
			s := strings.TrimPrefix(c.Text, "// frizz: ")
			var ai interface{}
			if err := json.Unmarshal([]byte(s), &ai); err != nil {
				return false, nil, "", errors.Wrapf(err, "annotation %s must be well formed json", s)
			}
			annotation, _ = ai.(string)
			found = true
			break
		}
	}
	if !found {
		return false, nil, "", nil
	}
	if g.Tok != token.TYPE {
		return false, nil, "", errors.Errorf("unsupported token tagged with frizz comment at %s", fset.Position(g.TokPos))
	}
	if len(g.Specs) != 1 {
		return false, nil, "", errors.Errorf("must be a single spec in type definition at %s", fset.Position(g.TokPos))
	}
	ts, ok := g.Specs[0].(*ast.TypeSpec)
	if !ok {
		return false, nil, "", errors.Errorf("must be type spec at %s", fset.Position(g.TokPos))
	}
	return true, ts, annotation, nil
}
