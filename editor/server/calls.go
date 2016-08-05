package server

import (
	"path/filepath"

	"strings"

	"os"

	"io/ioutil"

	"github.com/davelondon/kerr"
	"github.com/ghodss/yaml"
	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/editor/server/auther"
	"kego.io/editor/shared"
	"kego.io/json"
	"kego.io/ke"
	"kego.io/process/packages"
	"kego.io/process/parser"
	"kego.io/process/scanner"
	"kego.io/system"
)

type Server struct {
	ctx  context.Context
	auth auther.Auther
}

func (s *Server) Save(request *shared.SaveRequest, response *shared.SaveResponse) error {
	if !s.auth.Auth([]byte(request.Path), request.Hash) {
		return kerr.New("GIONMMGOWA", "Auth failed")
	}
	root, err := packages.GetDirFromPackage(s.ctx, request.Path)
	if err != nil {
		return kerr.Wrap("YKYVDSDGNV", err)
	}
	root = filepath.Clean(root)
	for _, file := range request.Files {
		// Check the filename is ok
		f := filepath.Clean(file.File)
		if err := checkFilename(root, f); err != nil {
			return kerr.Wrap("VVYOJAVPWC", err)
		}
		fullfilepath := filepath.Join(root, f)
		fs, err := os.Stat(fullfilepath)
		if err != nil {
			return kerr.New("KJKEVETWEY", "File %s not found", f)
		}
		if fs.IsDir() {
			return kerr.New("BALIAIMCMO", "File %s is a dir", f)
		}
		// Check the bytes are well formed json...
		var data interface{}
		if err := json.UnmarshalPlain(file.Bytes, &data); err != nil {
			return kerr.Wrap("IHFFPTCQPR", err)
		}
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			return kerr.New("AMJFJHFWVP", "Data in %s is not a json map", f)
		}
		if _, hasType := dataMap["type"]; !hasType {
			return kerr.New("HXYNWQQMFS", "Data in %s has no type field", f)
		}
		var output []byte
		if strings.HasSuffix(fullfilepath, ".json") {
			output = file.Bytes
		} else if strings.HasSuffix(fullfilepath, ".yml") || strings.HasSuffix(fullfilepath, ".yaml") {
			var err error
			if output, err = yaml.JSONToYAML(file.Bytes); err != nil {
				return kerr.Wrap("EAMEWSCAGB", err)
			}
		} else {
			return kerr.New("SHGMNTGCNG", "File %s has invalid extension", f)
		}
		if err := ioutil.WriteFile(fullfilepath, output, fs.Mode()); err != nil {
			return kerr.Wrap("KPDYGCYOYQ", err)
		}

		response.Files = append(response.Files, shared.SaveResponseFile{
			File: file.File,
			Hash: file.Hash,
		})
	}
	return nil
}

func checkFilename(root string, file string) error {
	root = filepath.Clean(root)
	file = filepath.Clean(file)
	full := filepath.Join(root, file)
	rel, err := filepath.Rel(root, full)
	if err != nil || strings.HasPrefix(rel, "..") {
		return kerr.New("CNMFTLYQFG", "Error in filename %s", file)
	}
	return nil
}

func (s *Server) Data(request *shared.DataRequest, response *shared.DataResponse) error {
	if !s.auth.Auth([]byte(request.Path), request.Hash) {
		return kerr.New("SYEKLIUMVY", "Auth failed")
	}
	env, err := parser.ScanForEnv(s.ctx, request.Package)
	if err != nil {
		return kerr.Wrap("PNAGGKHDYL", err)
	}

	file := filepath.Join(env.Dir, request.File)

	bytes, err := scanner.ProcessFile(file)

	localContext := envctx.NewContext(s.ctx, env)
	o := &system.Object{}
	if err := ke.UnmarshalUntyped(localContext, bytes, o); err != nil {
		return kerr.Wrap("SVINFEMKBG", err)
	}
	if o.Id.Name != request.Name {
		return kerr.New("TNJSLMPMLB", "Id does not match")
	}

	response.Package = request.Package
	response.Name = request.Name
	response.Found = true
	response.Data = bytes
	return nil
}
