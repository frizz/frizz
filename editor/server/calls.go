package server

import (
	"path/filepath"

	"io/ioutil"

	"os"

	"context"

	"frizz.io/context/envctx"
	"frizz.io/editor/server/auther"
	"frizz.io/editor/server/pkghelp"
	"frizz.io/editor/shared"
	"frizz.io/process/parser"
	"frizz.io/process/scanner"
	"frizz.io/system"
	"github.com/dave/kerr"
	"github.com/ghodss/yaml"
)

type Server struct {
	ctx  context.Context
	auth auther.Auther
}

func (s *Server) Save(request *shared.SaveRequest, response *shared.SaveResponse) error {
	if !s.auth.Auth([]byte(request.Path), request.Hash) {
		return kerr.New("GIONMMGOWA", "Auth failed")
	}
	pkg, err := pkghelp.Scan(s.ctx, request.Path)
	if err != nil {
		return kerr.Wrap("YKYVDSDGNV", err)
	}
	localContext := envctx.NewContext(s.ctx, pkg.Env)
	for _, info := range request.Files {

		// Check we only have yml, yaml or json extension.
		ext := filepath.Ext(info.File)
		if ext != ".json" && ext != ".yml" && ext != ".yaml" {
			return kerr.New("NDTPTCDOET", "Unsupported extension %s in %s", ext, info.File)
		}

		// Check the bytes are well formed json...
		o := &system.ObjectStub{}
		if err := system.Unmarshal(localContext, info.Bytes, o); err != nil {
			return kerr.Wrap("QISVPOXTCJ", err)
		}
		// Check type field exists
		if o.Type == nil {
			return kerr.New("PHINYFTGEC", "%s has no type", info.File)
		}
		// Check id field exists apart from system:package type.
		if o.Id == nil && *o.Type != *system.NewReference("frizz.io/system", "package") {
			return kerr.New("NNOEQPRQXS", "%s has no id", info.File)
		}
		// Convert output to YAML if needed.
		output := info.Bytes
		if ext == ".yaml" || ext == ".yml" {
			var err error
			if output, err = yaml.JSONToYAML(output); err != nil {
				return kerr.Wrap("EAMEWSCAGB", err)
			}
		}

		var mode os.FileMode
		var full string

		file := pkg.File(info.File)
		if file != nil {
			// The file already exists, so we should use the existing filemode
			full = file.AbsoluteFilepath
			fs, err := os.Stat(full)
			if err != nil {
				return kerr.Wrap("VLIJSSVSXU", err)
			}
			mode = fs.Mode()
		} else {
			if full, err = pkghelp.Check(pkg.Env.Dir, info.File, pkg.Env.Recursive); err != nil {
				return kerr.Wrap("YSQEHPFIVF", err)
			}
			mode = 0644
			if _, err := os.Stat(full); err == nil || !os.IsNotExist(err) {
				return kerr.New("XOEPAUNCXB", "Can't overwrite %s - existing file is not a valid frizz data file", info.File)
			}
		}

		if err := ioutil.WriteFile(full, output, mode); err != nil {
			return kerr.Wrap("KPDYGCYOYQ", err)
		}

		response.Files = append(response.Files, shared.SaveResponseFile{
			File: info.File,
			Hash: info.Hash,
		})
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

	full, err := pkghelp.Check(env.Dir, request.File, env.Recursive)
	if err != nil {
		return kerr.Wrap("JEYTFWKMYF", err)
	}

	bytes, err := scanner.ProcessFile(full)
	if err != nil {
		return kerr.Wrap("HQXMIMWXFY", err)
	}
	if bytes == nil {
		return kerr.New("HIUINHIAPY", "Error reading %s", request.File)
	}

	localContext := envctx.NewContext(s.ctx, env)
	o := &system.ObjectStub{}
	if err := system.Unmarshal(localContext, bytes, o); err != nil {
		return kerr.Wrap("SVINFEMKBG", err)
	}
	if o.Id.Name != request.Name {
		return kerr.New("TNJSLMPMLB", "Id %s in %s does not match request %s", o.Id.Name, request.File, request.Name)
	}

	response.Package = request.Package
	response.Name = request.Name
	response.Found = true
	response.Data = bytes
	return nil
}
