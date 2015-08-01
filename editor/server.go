package editor

import (
	"bytes"
	"go/scanner"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"syscall"
	"text/template"
	"time"

	"golang.org/x/tools/go/types"

	"github.com/gopherjs/gopherjs/compiler"
	"github.com/neelance/sourcemap"

	"fmt"

	gbuild "github.com/gopherjs/gopherjs/build"
)

type serveCommandFileSystem struct {
	options    *gbuild.Options
	dirs       []string
	sourceMaps map[string][]byte
}

func (fs serveCommandFileSystem) Open(name string) (http.File, error) {
	dir, file := path.Split(name)
	base := path.Base(dir) // base is parent folder name, which becomes the output file name.

	isPkg := file == base+".js"
	isMap := file == base+".js.map"
	isIndex := file == "index.html"

	if isPkg || isMap || isIndex {
		// If we're going to be serving our special files, make sure there's a Go command in this folder.
		s := gbuild.NewSession(fs.options)
		buildPkg, err := gbuild.Import(path.Dir(name[1:]), 0, s.InstallSuffix(), fs.options.BuildTags)
		if err != nil || buildPkg.Name != "main" {
			isPkg = false
			isMap = false
			isIndex = false
		}

		switch {
		case isPkg:
			buf := bytes.NewBuffer(nil)
			browserErrors := bytes.NewBuffer(nil)
			exitCode := handleError(func() error {
				pkg := &gbuild.PackageData{Package: buildPkg}
				if err := s.BuildPackage(pkg); err != nil {
					return err
				}

				sourceMapFilter := &compiler.SourceMapFilter{Writer: buf}
				m := &sourcemap.Map{File: base + ".js"}
				sourceMapFilter.MappingCallback = gbuild.NewMappingCallback(m, fs.options.GOROOT, fs.options.GOPATH)

				deps, err := compiler.ImportDependencies(pkg.Archive, s.ImportContext.Import)
				if err != nil {
					return err
				}
				if err := compiler.WriteProgramCode(deps, sourceMapFilter); err != nil {
					return err
				}

				mapBuf := bytes.NewBuffer(nil)
				m.WriteTo(mapBuf)
				buf.WriteString("//# sourceMappingURL=" + base + ".js.map\n")
				fs.sourceMaps[name+".map"] = mapBuf.Bytes()

				return nil
			}, fs.options, browserErrors)
			if exitCode != 0 {
				buf = browserErrors
			}
			return newFakeFile(base+".js", buf.Bytes()), nil

		case isMap:
			if content, ok := fs.sourceMaps[name]; ok {
				return newFakeFile(base+".js.map", content), nil
			}
		}
	}

	for _, d := range fs.dirs {
		f, err := http.Dir(filepath.Join(d, "src")).Open(name)
		if err == nil {
			return f, nil
		}
	}

	if isIndex {
		// If there was no index.html file in any dirs, supply our own.
		return newFakeFile("index.html", []byte(`
			<html>
				<head>
					<meta charset="utf-8">
					<link rel="stylesheet" href="https://storage.googleapis.com/code.getmdl.io/1.0.0/material.indigo-pink.min.css">
					<script src="https://storage.googleapis.com/code.getmdl.io/1.0.0/material.min.js"></script>
					<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
					<script src="`+base+`.js"></script>
				</head>

			</html>`)), nil
	}

	return nil, os.ErrNotExist
}

type fakeFile struct {
	name string
	size int
	io.ReadSeeker
}

func newFakeFile(name string, content []byte) *fakeFile {
	return &fakeFile{name: name, size: len(content), ReadSeeker: bytes.NewReader(content)}
}

func (f *fakeFile) Close() error {
	return nil
}

func (f *fakeFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, os.ErrInvalid
}

func (f *fakeFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *fakeFile) Name() string {
	return f.name
}

func (f *fakeFile) Size() int64 {
	return int64(f.size)
}

func (f *fakeFile) Mode() os.FileMode {
	return 0
}

func (f *fakeFile) ModTime() time.Time {
	return time.Time{}
}

func (f *fakeFile) IsDir() bool {
	return false
}

func (f *fakeFile) Sys() interface{} {
	return nil
}

var currentDirectory string

func init() {
	var err error
	currentDirectory, err = os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	currentDirectory, err = filepath.EvalSymlinks(currentDirectory)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// If browserErrors is non-nil, errors are written for presentation in browser.
func handleError(f func() error, options *gbuild.Options, browserErrors *bytes.Buffer) int {
	switch err := f().(type) {
	case nil:
		return 0
	case compiler.ErrorList:
		for _, entry := range err {
			printError(entry, options, browserErrors)
		}
		return 1
	case *exec.ExitError:
		return err.Sys().(syscall.WaitStatus).ExitStatus()
	default:
		printError(err, options, browserErrors)
		return 1
	}
}

// sprintError returns an annotated error string without trailing newline.
func sprintError(err error) string {
	makeRel := func(name string) string {
		if relname, err := filepath.Rel(currentDirectory, name); err == nil {
			return relname
		}
		return name
	}

	switch e := err.(type) {
	case *scanner.Error:
		return fmt.Sprintf("%s:%d:%d: %s", makeRel(e.Pos.Filename), e.Pos.Line, e.Pos.Column, e.Msg)
	case types.Error:
		pos := e.Fset.Position(e.Pos)
		return fmt.Sprintf("%s:%d:%d: %s", makeRel(pos.Filename), pos.Line, pos.Column, e.Msg)
	default:
		return fmt.Sprintf("%s", e)
	}
}

// printError prints err to Stderr with options. If browserErrors is non-nil, errors are also written for presentation in browser.
func printError(err error, options *gbuild.Options, browserErrors *bytes.Buffer) {
	e := sprintError(err)
	options.PrintError("%s\n", e)
	if browserErrors != nil {
		fmt.Fprintln(browserErrors, `console.error("`+template.JSEscapeString(e)+`");`)
	}
}
