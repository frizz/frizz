package generate

import (
	"bytes"
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

	"fmt"

	"frizz.io/utils"
	. "github.com/dave/jennifer/jen"
	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
	"github.com/pkg/errors"
)

func Save(out io.Writer, paths ...string) error {
	// notest
	env := vos.Os()

	def, err := Scan(env, paths...)
	if err != nil {
		return errors.WithMessage(err, "scanning")
	}

	for _, def := range def.Packages {
		generated := &bytes.Buffer{}
		fname := filepath.Join(def.Dir, "generated.frizz.go")

		hasContent, err := def.Generate(generated)
		if err != nil {
			return errors.WithMessage(err, "generating source")
		}

		var existing []byte
		if _, err := os.Stat(fname); err == nil {
			existing, err = ioutil.ReadFile(fname)
			if err != nil {
				return errors.WithMessage(err, "reading existing source")
			}
		}

		if hasContent {
			if bytes.Equal(existing, generated.Bytes()) {
				fmt.Fprintf(out, "unchanged: %s\n", def.Path)
			} else {
				if err := ioutil.WriteFile(fname, generated.Bytes(), 0777); err != nil {
					return errors.Wrap(err, "saving source")
				}
				fmt.Fprintf(out, "updated: %s\n", def.Path)
			}
		} else {
			if _, err := os.Stat(fname); err == nil {
				if err := os.Remove(fname); err != nil {
					return errors.Wrap(err, "removing source")
				}
				fmt.Fprintf(out, "removed: %s\n", def.Path)
			}
		}
	}
	return nil
}

func Scan(env vos.Env, paths ...string) (*scanDef, error) {
	def := &scanDef{
		fset:     token.NewFileSet(),
		env:      env,
		input:    []struct{ path, dir string }{},
		pcache:   patsy.NewCache(env),
		Packages: map[string]*packageDef{},
	}

	for _, p := range paths {
		d, err := patsy.Dir(env, p)
		if err != nil {
			return nil, err
		}
		def.input = append(def.input, struct{ path, dir string }{path: p, dir: d})
	}

	// first scan the data recoding stubs and recursively scanning imported packages
	for _, p := range def.input {
		if err := scanData(p.path, def); err != nil {
			return nil, err
		}
	}

	// then scan all source files
	if err := scanSource(def); err != nil {
		return nil, err
	}

	return def, nil
}

func (p *packageDef) Generate(writer io.Writer) (bool, error) {

	if len(p.types) == 0 && len(p.stubs) == 0 && len(p.imports) == 0 {
		return false, nil
	}

	f := NewFilePath(p.Path)
	generatePackage(p, f)
	generatePackers(p, f)
	generateStubs(p, f)
	generateTypes(p, f)
	generateImports(p, f)
	generateData(p, f)

	if err := f.Render(writer); err != nil {
		return false, errors.Wrap(err, "error saving generated file")
	}
	return true, nil
}

func scanData(path string, prog *scanDef) error {

	if _, ok := prog.Packages[path]; ok {
		// already scanned
		return nil
	}

	dir, err := patsy.Dir(prog.env, path)
	if err != nil {
		return err
	}

	pkg := &packageDef{
		scan:    prog,
		Path:    path,
		Dir:     dir,
		types:   map[string]*typeDef{},
		stubs:   map[string]*stub{},
		imports: map[string]bool{},
	}

	prog.Packages[path] = pkg

	// scan all .frizz.json files for stubs
	files, err := filepath.Glob(filepath.Join(dir, "*.frizz.json"))
	if err != nil {
		return errors.Wrapf(err, "finding files in %s", dir)
	}

	for _, fpath := range files {
		s, err := decodeStub(fpath, pkg.Path)
		if err != nil {
			return errors.Wrapf(err, "decoding %s", fpath)
		}
		pkg.stubs[fpath] = s
		if s.Import != nil {
			for _, importPath := range s.Import {
				pkg.imports[importPath] = true
				if err := scanData(importPath, prog); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func generatePackage(p *packageDef, f *File) {

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
		Return(Lit(p.Path)),
	)

}

func generatePackers(p *packageDef, f *File) {

	var sorted []*typeDef
	for _, t := range p.types {
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
		f.Add(p.unpacker(t.spec, t.name, true, t.packable))
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
		f.Add(p.repacker(t.spec, t.name, true, t.packable))
	}
}

func generateStubs(p *packageDef, f *File) {

	type def struct {
		filename string
		stub     *stub
	}
	var sorted []def
	for fpath, stub := range p.stubs {
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

func generateTypes(p *packageDef, f *File) {

	type def struct {
		name     string
		filename string
	}
	var sorted []def
	for name, td := range p.types {
		if td.annotation == "" {
			continue
		}
		sorted = append(sorted, def{name, td.annotation})
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

func generateImports(p *packageDef, f *File) {

	var sorted []string
	for pkg := range p.imports {
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
		g.Id("packages").Index(Lit(p.Path)).Op("=").Id("Package")
		for _, pkg := range sorted {
			g.Qual(pkg, "Package").Dot("GetImportedPackages").Call(Id("packages"))
		}
	})

}

func generateData(p *packageDef, f *File) {

	if len(p.stubs) == 0 {
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
	for fpath, stub := range p.stubs {
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
