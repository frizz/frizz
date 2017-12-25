package bootstrap

import (
	"go/ast"
	"go/parser"
	"go/token"

	"fmt"

	"bytes"

	"strings"

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
					fmt.Println("Found in cache:", path)
					archive, err := compiler.ReadArchive(path+".a", path, bytes.NewBuffer(resp), b.Packages)
					if err != nil {
						return nil, err
					}
					b.Archives[path] = archive
					return archive, nil
				}

				fmt.Println("Compiling:", path)
				files, err := parseFiles(fset, source.Source)
				if err != nil {
					return nil, err
				}
				archive, err := compiler.Compile(path, files, fset, importContext, false)
				if err != nil {
					return nil, err
				}

				buf := &bytes.Buffer{}
				if err := compiler.WriteArchive(archive, buf); err != nil {
					return nil, err
				}
				if err := b.Cache.Put(name, buf.Bytes()); err != nil {
					return nil, err
				}

				b.Archives[path] = archive
				return archive, nil
			}
			return nil, fmt.Errorf("can't find %s", path)
		},
	}

	fmt.Println("Compiling main")
	files, err := parseFiles(fset, map[string][]byte{"prog.go": src})
	if err != nil {
		return err
	}

	archive, err := compiler.Compile("main", files, fset, importContext, false)
	if err != nil {
		if strings.Contains(err.Error(), "Package not declared by") {
			fmt.Println("Not a frizz package")
			return nil
		}
		return err
	}
	b.Archives["main"] = archive

	archives, err := compiler.ImportDependencies(archive, importContext.Import)
	if err != nil {
		return err
	}

	fmt.Println("Running editor...")
	buf := bytes.NewBuffer(nil)
	buf.WriteString("try{\n")
	compiler.WriteProgramCode(archives, &compiler.SourceMapFilter{Writer: buf})
	buf.WriteString("} catch (err) {\ngoPanicHandler(err.message);\n}\n")

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
