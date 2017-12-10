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

	"frizz.io/generate/jast"
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
		env:     env,
		path:    path,
		pcache:  patsy.NewCache(env),
		dir:     dir,
		types:   map[string]*scanner.TypeDef{},
		stubs:   map[string]*stub{},
		imports: map[string]bool{},
	}

	if err := scanData(prog); err != nil {
		return false, err
	}

	if err := scanSource(prog); err != nil {
		return false, err
	}

	// now we have scanned the source, se can create the jast
	prog.jast = jast.New(nil, prog.info.Uses)

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

	s := scanner.New(prog.path, prog.env)
	if err := s.Scan(); err != nil {
		return err
	}

	prog.types = s.Types
	prog.info = s.Info
	for _, i := range s.Imports {
		prog.imports[i] = true
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

	var sorted []*scanner.TypeDef
	for _, t := range prog.types {
		sorted = append(sorted, t)
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Name < sorted[j].Name })

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
					g.Case(Lit(t.Name)).Block(
						Return(Id("p").Dot("Unpack"+t.Name).Call(
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
		f.Add(prog.unpacker(t.Spec, t.Name, true, t.Packable))
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
					g.Case(Lit(t.Name)).Block(
						Return(Id("p").Dot("Repack"+t.Name).Call(
							Id("context"),
							Id("in").Assert(Id(t.Name)),
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
		f.Add(prog.repacker(t.Spec, t.Name, true, t.Packable))
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
		if td.Annotation == "" {
			continue
		}
		sorted = append(sorted, def{name, td.Annotation})
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
