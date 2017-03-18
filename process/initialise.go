package process

import (
	"flag"

	"kego.io/process/parser"

	"context"

	"github.com/dave/kerr"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/sysctx"
	"kego.io/context/vosctx"
	"kego.io/context/wgctx"
	"kego.io/process/packages"
)

type OptionsInterface interface {
	getOptions() Options
}

type Flags struct {
	Edit     *bool
	Validate *bool
	Update   *bool
	Log      *bool
	Path     *string
	Debug    *bool
	Port     *int
}

type Options struct {
	Edit     bool
	Validate bool
	Update   bool
	Log      bool
	Path     string
	Debug    bool
	Port     int
}

func (f Options) getOptions() Options {
	return f
}

func (f Flags) getOptions() Options {

	var edit, update, log, debug, validate *bool
	var port *int
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
	if f.Port == nil {
		port = flag.Int("p", 0, "Port: use a specific port for the editor server. Default: use a random port")
	} else {
		port = f.Port
	}
	if f.Debug == nil {
		debug = flag.Bool("d", false, "Debug: don't close the server when the connection is closed")
	} else {
		debug = f.Debug
	}
	if !flag.Parsed() {
		// notest
		flag.Parse()
	}
	if f.Path == nil {
		p := flag.Arg(0)
		path = &p
	} else {
		path = f.Path
	}

	return Options{
		Edit:     *edit,
		Update:   *update,
		Log:      *log,
		Path:     *path,
		Debug:    *debug,
		Validate: *validate,
		Port:     *port,
	}
}

func Initialise(ctx context.Context, overrides OptionsInterface) (context.Context, context.CancelFunc, error) {
	if overrides == nil {
		overrides = Flags{}
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
	cmd.Port = options.Port
	if options.Path == "" {
		dir, err := vos.Getwd()
		if err != nil {
			return nil, nil, kerr.Wrap("OKOLXAMBSJ", err)
		}
		p, err := packages.GetPackageFromDir(ctx, dir)
		if err != nil {
			return nil, nil, kerr.Wrap("ADNJKTLAWY", err)
		}
		path = p
	} else {
		path = options.Path
	}

	ctx = cmdctx.NewContext(ctx, cmd)

	pcache, err := parser.Parse(ctx, path)
	if err != nil {
		return nil, nil, kerr.Wrap("EBMBIBIKUF", err)
	}

	ctx = envctx.NewContext(ctx, pcache.Env)

	return ctx, cancel, nil
}
