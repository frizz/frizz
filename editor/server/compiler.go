package server

import (
	"bytes"

	"go/parser"
	"go/token"

	"go/ast"

	"context"

	"github.com/davelondon/kerr"
	"github.com/gopherjs/gopherjs/build"
	"github.com/gopherjs/gopherjs/compiler"
	"github.com/neelance/sourcemap"
	"kego.io/context/vosctx"
)

func Compile(ctx context.Context, source []byte, mapping bool) ([]byte, error) {
	vos := vosctx.FromContext(ctx)

	options := &build.Options{
		GOROOT:        vos.Getenv("GOROOT"),
		GOPATH:        vos.Getenv("GOPATH"),
		CreateMapFile: mapping,
	}
	s := build.NewSession(options)

	packages := make(map[string]*compiler.Archive)
	importContext := &compiler.ImportContext{
		Packages: s.Types,
		Import:   s.BuildImportPath,
	}

	fileSet := token.NewFileSet()

	file, err := parser.ParseFile(fileSet, "prog.go", source, parser.ParseComments)
	if err != nil {
		return nil, kerr.Wrap("NCYFEKGCWX", err)
	}

	mainPkg, err := compiler.Compile("main", []*ast.File{file}, fileSet, importContext, false)
	if err != nil {
		return nil, kerr.Wrap("KPHUKOLTBX", err)
	}
	packages["main"] = mainPkg

	bufCode := bytes.NewBuffer(nil)
	filter := &compiler.SourceMapFilter{Writer: bufCode}
	allPkgs, err := compiler.ImportDependencies(mainPkg, importContext.Import)
	if err != nil {
		return nil, kerr.Wrap("TIMQHFQTWL", err)
	}

	if mapping {
		bufMap := bytes.NewBuffer(nil)
		smap := &sourcemap.Map{File: "script.js"}
		filter.MappingCallback = build.NewMappingCallback(smap, options.GOROOT, options.GOPATH, false)
		if err := compiler.WriteProgramCode(allPkgs, filter); err != nil {
			return nil, kerr.Wrap("YKQEKRKBPL", err)
		}
		if err := smap.WriteTo(bufMap); err != nil {
			return nil, kerr.Wrap("VYQGYAAADG", err)
		}
		return bufMap.Bytes(), nil
	}

	if err := compiler.WriteProgramCode(allPkgs, filter); err != nil {
		return nil, kerr.Wrap("DPPVHCOTBQ", err)
	}
	if _, err := bufCode.WriteString("//# sourceMappingURL=script.js.map\n"); err != nil {
		return nil, kerr.Wrap("CXXKWQVGUI", err)
	}
	return bufCode.Bytes(), nil
}
