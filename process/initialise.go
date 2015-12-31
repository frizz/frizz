package process

import (
	"flag"
	"os"

	"golang.org/x/net/context"
	"kego.io/context/cachectx"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/wgctx"
	"kego.io/kerr"
	"kego.io/parse"
	"kego.io/process/pkgutils"
)

type optionsSpec interface {
	getOptions() FromDefaults
}

type FromFlags struct {
	Edit    *bool
	Update  *bool
	Verbose *bool
	Path    *string
	Debug   *bool
}

type FromDefaults struct {
	Edit    bool
	Update  bool
	Verbose bool
	Path    string
	Debug   bool
}

func (f FromDefaults) getOptions() FromDefaults {
	return f
}

func (f FromFlags) getOptions() FromDefaults {

	var edit, update, verbose, debug *bool
	var path *string
	if f.Edit == nil {
		edit = flag.Bool("e", false, "Edit: open the editor")
	} else {
		edit = f.Edit
	}
	if f.Update == nil {
		update = flag.Bool("u", false, "Update: update all import packages e.g. go get -u")
	} else {
		update = f.Update
	}
	if f.Verbose == nil {
		verbose = flag.Bool("v", false, "Verbose")
	} else {
		verbose = f.Verbose
	}
	if f.Debug == nil {
		debug = flag.Bool("d", false, "Debug: don't close the server when the connection is closed")
	} else {
		debug = f.Debug
	}
	if !flag.Parsed() {
		flag.Parse()
	}
	if f.Path == nil {
		p := flag.Arg(0)
		path = &p
	} else {
		path = f.Path
	}

	return FromDefaults{
		Edit:    *edit,
		Update:  *update,
		Verbose: *verbose,
		Path:    *path,
		Debug:   *debug,
	}
}

func Initialise(overrides optionsSpec) (context.Context, context.CancelFunc, error) {
	if overrides == nil {
		overrides = FromFlags{}
	}
	options := overrides.getOptions()

	ctx, cancel := context.WithCancel(context.Background())
	ctx = wgctx.NewContext(ctx)
	ctx = cachectx.NewContext(ctx)

	cmd := &cmdctx.Cmd{}

	path := ""
	cmd.Edit = options.Edit
	cmd.Update = options.Update
	cmd.Verbose = options.Verbose
	cmd.Debug = options.Debug
	if options.Path == "" {

		dir, err := os.Getwd()
		if err != nil {
			return nil, nil, kerr.New("OKOLXAMBSJ", err, "os.Getwd")
		}
		cmd.Dir = dir

		p, err := pkgutils.GetPackageFromDir(cmd.Dir)
		if err != nil {
			return nil, nil, kerr.New("ADNJKTLAWY", err, "pkgutils.GetPackageFromDir")
		}
		path = p

	} else {

		path = options.Path
		dir, err := pkgutils.GetDirFromPackage(options.Path)
		if err != nil {
			return nil, nil, kerr.New("QKPSIYSKUN", err, "pkgutils.GetDirFromPackage")
		}
		cmd.Dir = dir
	}

	env, err := parse.Parse(ctx, path, []string{})
	if err != nil {
		return nil, nil, kerr.New("EBMBIBIKUF", err, "parse.Parse")
	}

	ctx = envctx.NewContext(ctx, env)
	ctx = cmdctx.NewContext(ctx, cmd)

	return ctx, cancel, nil
}
