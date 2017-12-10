package scanner

import (
	"encoding/json"
	"go/ast"
	"go/build"
	"io/ioutil"

	"os"

	"strings"

	"go/types"

	"go/parser"

	"go/token"

	"path/filepath"

	"github.com/dave/patsy/vos"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

type Scanner struct {
	env     vos.Env
	pkg     string
	dir     string
	prog    *loader.Program
	Types   map[string]*TypeDef
	Info    types.Info
	Imports []string // frizz packages imported by Go source
}

func New(pkg string, env vos.Env) *Scanner {
	return &Scanner{
		pkg:   pkg,
		env:   env,
		Types: map[string]*TypeDef{},
	}
}

type TypeDef struct {
	Packable   bool
	Spec       ast.Expr
	Name       string
	Annotation string
}

func (s *Scanner) Scan() error {
	conf := loader.Config{}
	conf.Import(s.pkg)
	conf.Import("frizz.io/global") // contains UnpackInterface and RepackInterface
	conf.Build = &build.Default
	conf.Build.GOPATH = s.env.Getenv("GOPATH")

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

	prog, err := conf.Load()
	if err != nil {
		return errors.WithStack(err)
	}
	s.prog = prog

	global := prog.Package("frizz.io/global")
	var packable *types.Interface
	for i, v := range global.Defs {
		if i.Name == "Packable" {
			packable = v.Type().Underlying().(*types.Interface)
		}
	}
	if packable == nil {
		return errors.New("can't find frizz.io/global.Packable")
	}

	pkg := prog.Package(s.pkg)

	s.Info = pkg.Info

	packables := map[string]bool{}
	for i, v := range pkg.Defs {
		if v, ok := v.(*types.TypeName); ok {
			if types.Implements(types.NewPointer(v.Type()), packable) {
				packables[i.Name] = true
			}
		}
	}

	for _, f := range pkg.Files {
		for _, d := range f.Decls {
			found, spec, annotation, err := isFrizzDecl(prog.Fset, d)
			if err != nil {
				return err
			}
			if !found {
				continue
			}
			td := &TypeDef{
				Name:       spec.Name.Name,
				Spec:       spec.Type,
				Annotation: annotation,
				Packable:   packables[spec.Name.Name],
			}
			s.Types[spec.Name.Name] = td
		}
	}

Package:
	for _, p := range pkg.Pkg.Imports() {
		if !strings.Contains(p.Path(), "/") {
			// package paths not containing a slash are standard library packages and can be ignored
			// TODO: better way of doing this?
			continue
		}
		if p.Path() == s.pkg {
			// ignore the local package
			continue
		}
		pi := prog.Package(p.Path())

		// first search for a generated.frizz.go file
		for _, f := range pi.Files {
			if _, fname := filepath.Split(prog.Fset.File(f.Pos()).Name()); fname == "generated.frizz.go" {
				s.Imports = append(s.Imports, p.Path())
				continue Package
			}
		}

		// if we can't find the generated file, perhaps it's not been created yet... if so, search
		// the source for frizz annotated types, but exit as soon as we find one.
		// TODO: Do we need this?
		for _, f := range pi.Files {
			for _, d := range f.Decls {
				found, _, _, _ := isFrizzDecl(prog.Fset, d) // ignore errors in other packages
				if found {
					s.Imports = append(s.Imports, p.Path())
					continue Package
				}
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
