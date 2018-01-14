package bootstrap

import (
	"go/ast"
	"go/parser"
	"go/token"

	"fmt"

	"bytes"

	"frizz.io/edit/common"
	"frizz.io/edit/hasher"
	"github.com/gopherjs/gopherjs/compiler"
)

func (b *Bootstrap) Compile() error {

	src, err := b.GenStub()
	if err != nil {
		return err
	}

	var compiled, cached int

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
					cached++
					archive, err := compiler.ReadArchive(path+".a", path, bytes.NewBuffer(resp), b.Packages)
					if err != nil {
						return nil, err
					}
					b.Archives[path] = archive
					return archive, nil
				}

				compiled++
				files, err := parseFiles(fset, source.Files)
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

	b.Log.Println("Compiling...")

	files := map[string][]byte{"main.go": src}
	hash, err := hasher.Hash(files)
	if err != nil {
		return err
	}
	b.Source["main"] = common.Bundle{Hash: hash, Files: files}
	archive, err := importContext.Import("main")
	if err != nil {
		return err
	}

	archives, err := compiler.ImportDependencies(archive, importContext.Import)
	if err != nil {
		return err
	}

	b.Log.Printf("Finished compiling: %d compiled, %d found in cache.\n", compiled, cached)

	b.Log.Println("Running editor...")
	buf := bytes.NewBuffer(nil)
	buf.WriteString("try{\n")
	compiler.WriteProgramCode(archives, &compiler.SourceMapFilter{Writer: buf})
	buf.WriteString("} catch (err) {\ngoPanicHandler(err.message);\n}\n")

	b.Code = buf.String()

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
