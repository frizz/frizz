package process

import (
	"flag"

	"kego.io/process/parser"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/sysctx"
	"kego.io/context/vosctx"
	"kego.io/context/wgctx"
	"kego.io/kerr"
	"kego.io/process/packages"
)

type optionsSpec interface {
	getOptions() FromDefaults
}

type FromFlags struct {
	Edit     *bool
	Validate *bool
	Update   *bool
	Log      *bool
	Path     *string
	Debug    *bool
}

type FromDefaults struct {
	Edit     bool
	Validate bool
	Update   bool
	Log      bool
	Path     string
	Debug    bool
}

func (f FromDefaults) getOptions() FromDefaults {
	return f
}

func (f FromFlags) getOptions() FromDefaults {

	var edit, update, log, debug, validate *bool
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
	if f.Log == nil {
		log = flag.Bool("l", false, "Log")
	} else {
		log = f.Log
	}
	if f.Validate == nil {
		validate = flag.Bool("v", false, "Validate")
	} else {
		validate = f.Validate
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
		Edit:     *edit,
		Update:   *update,
		Log:      *log,
		Path:     *path,
		Debug:    *debug,
		Validate: *validate,
	}
}

func Initialise(ctx context.Context, overrides optionsSpec) (context.Context, context.CancelFunc, error) {
	if overrides == nil {
		overrides = FromFlags{}
	}
	options := overrides.getOptions()

	ctx, cancel := context.WithCancel(ctx)
	ctx = jsonctx.AutoContext(ctx)
	ctx = wgctx.NewContext(ctx)
	ctx = sysctx.NewContext(ctx)

	cmd := &cmdctx.Cmd{}

	vos := vosctx.FromContext(ctx)

	path := ""
	cmd.Edit = options.Edit
	cmd.Validate = options.Validate
	cmd.Update = options.Update
	cmd.Log = options.Log
	cmd.Debug = options.Debug
	if options.Path == "" {
		dir, err := vos.Getwd()
		if err != nil {
			return nil, nil, kerr.New("OKOLXAMBSJ", err, "vos.Getwd")
		}
		p, err := packages.GetPackageFromDir(ctx, dir)
		if err != nil {
			return nil, nil, kerr.New("ADNJKTLAWY", err, "packages.GetPackageFromDir")
		}
		path = p
	} else {
		path = options.Path
	}

	ctx = cmdctx.NewContext(ctx, cmd)

	pcache, err := parser.Parse(ctx, path)
	if err != nil {
		return nil, nil, kerr.New("EBMBIBIKUF", err, "parser.Parse")
	}

	ctx = envctx.NewContext(ctx, pcache.Environment)

	return ctx, cancel, nil
}
