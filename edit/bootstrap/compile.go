package bootstrap

import (
	"go/ast"
	"go/parser"
	"go/token"

	"fmt"

	"bytes"

	"github.com/gopherjs/gopherjs/compiler"
	"github.com/gopherjs/gopherjs/js"
)

func (b *Bootstrap) Compile() error {

	src, err := b.GenStub()
	if err != nil {
		return err
	}

	var importContext *compiler.ImportContext
	fset := token.NewFileSet()

	importContext = &compiler.ImportContext{
		Packages: b.Packages,
		Import: func(path string) (*compiler.Archive, error) {
			if archive, ok := b.Archives[path]; ok {
				return archive, nil
			}
			if source, ok := b.Source[path]; ok {

				name := fmt.Sprintf("v4/%s?%d", path, source.Hash)

				found, resp, err := b.Cache.Get(name)
				if err != nil {
					return nil, err
				}
				if found {
					fmt.Println("Found in cache", path)
					archive, err := compiler.ReadArchive(path+".a", path, bytes.NewBuffer(resp), b.Packages)
					if err != nil {
						return nil, err
					}
					b.Archives[path] = archive
					return archive, nil
				}

				fmt.Println("Parsing", path)
				files, err := parseFiles(fset, source.Source)
				if err != nil {
					return nil, err
				}
				fmt.Println("Compiling", path)
				archive, err := compiler.Compile(path, files, fset, importContext, false)
				if err != nil {
					return nil, err
				}
				fmt.Println("Finished compiling", path)

				buf := &bytes.Buffer{}
				if err := compiler.WriteArchive(archive, buf); err != nil {
					return nil, err
				}

				if err := b.Cache.Put(name, buf.Bytes()); err != nil {
					return nil, err
				}
				fmt.Println("Stored as", name)

				b.Archives[path] = archive
				return archive, nil
			}
			panic(fmt.Sprintf("can't find %s", path))
			return &compiler.Archive{}, nil
		},
	}

	fmt.Println("Parsing main")
	files, err := parseFiles(fset, map[string][]byte{"prog.go": src})
	if err != nil {
		return err
	}

	fmt.Println("Compiling main")
	mainArchive, err := compiler.Compile("main", files, fset, importContext, false)
	if err != nil {
		return err
	}
	b.Archives["main"] = mainArchive

	fmt.Println("Importing dependancies")
	allArchives, err := compiler.ImportDependencies(mainArchive, importContext.Import)
	if err != nil {
		return err
	}

	fmt.Println("Writing program code")
	buf := bytes.NewBuffer(nil)
	buf.WriteString("try{\n")
	compiler.WriteProgramCode(allArchives, &compiler.SourceMapFilter{Writer: buf})
	buf.WriteString("} catch (err) {\ngoPanicHandler(err.message);\n}\n")

	fmt.Println("Running js, length", buf.Len())
	js.Global.Set("$checkForDeadlock", true)
	js.Global.Call("eval", js.InternalObject(buf.String()))

	return nil
}

func parseFiles(fset *token.FileSet, files map[string][]byte) ([]*ast.File, error) {
	var out []*ast.File
	for name, file := range files {
		parsed, err := parser.ParseFile(fset, name, file, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		out = append(out, parsed)
	}
	return out, nil
}
