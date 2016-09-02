package client // import "kego.io/editor/client"

// ke: {"package": {"jstest": true}}

import (
	"bytes"
	"net/rpc"

	"fmt"

	"encoding/gob"

	"encoding/base64"

	"context"

	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/websocket"
	"honnef.co/go/js/dom"
	"kego.io/context/envctx"
	"kego.io/context/jsonctx"
	"kego.io/context/sysctx"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/clientctx"
	"kego.io/editor/client/connection"
	"kego.io/editor/client/editors"
	"kego.io/editor/client/stores"
	"kego.io/editor/client/views"
	"kego.io/editor/shared"
	"kego.io/process/parser"
	"kego.io/system"
)

func Start() error {

	loc := dom.GetWindow().Location()
	addr := fmt.Sprintf("ws://%s:%s/_rpc", loc.Hostname, loc.Port)
	ws, err := websocket.Dial(addr)
	if err != nil {
		return kerr.Wrap("HNQFLPFAJD", err)
	}

	app := &stores.App{
		Conn: connection.New(rpc.NewClient(ws)),
		Fail: make(chan error),
	}

	// We parse the info attribute from the body tag
	info, err := getInfo(getRawInfo())
	if err != nil {
		return kerr.Wrap("MGLVIQIDDY", err)
	}

	var ctx context.Context
	ctx = context.Background()
	ctx = envctx.NewContext(ctx, &envctx.Env{Path: info.Path, Aliases: info.Aliases})
	ctx = sysctx.NewContext(ctx)
	ctx = jsonctx.AutoContext(ctx)
	ctx = stores.NewContext(ctx, app)
	ctx = clientctx.NewContext(ctx)

	app.Init(ctx)

	// Don't do this. Implement the Editable interface instead. We can't do
	// this for system types so we use this method instead.
	editors.Register(ctx)

	if _, err := registerTypes(ctx, info.Path, info.Imports); err != nil {
		return kerr.Wrap("MMJDDOBAUK", err)
	}

	p := views.NewPage(ctx)
	vecty.RenderAsBody(p)

	// TODO: work out why I can't seem to call this without using eval
	js.Global.Get("window").Call("eval", "Split(['#tree', '#main'], {sizes:[25, 75]});")

	app.Dispatch(&actions.InitialState{
		Info: info,
	})

	go func() {
		err, open := <-app.Fail
		if !open {
			// Channel has been closed, so app should gracefully exit.
			fmt.Println("Server disconnected")
		} else {
			// Error received, so app should display error.
			fmt.Println(err.Error())
		}
	}()

	return nil
}

func registerTypes(ctx context.Context, path string, imports map[string]shared.ImportInfo) (*sysctx.SysPackageInfo, error) {
	system.RegisterJsonTypes(ctx)
	scache := sysctx.FromContext(ctx)
	var current *sysctx.SysPackageInfo
	for _, info := range imports {
		env := &envctx.Env{Path: info.Path, Aliases: info.Aliases}
		pcache := scache.SetEnv(env)
		for _, ti := range info.Types {
			if err := parser.ProcessTypeFileBytes(ctx, env, ti.File, ti.Bytes, pcache, nil); err != nil {
				return nil, kerr.Wrap("UJLXYWCVUC", err)
			}
		}
		if path == info.Path {
			current = pcache
		}
	}
	return current, nil
}

func getInfo(infoBase64 string) (info *shared.Info, err error) {
	info = &shared.Info{}
	infoBytes, err := base64.StdEncoding.DecodeString(infoBase64)
	if err != nil {
		return nil, kerr.Wrap("UTKDDLYKKH", err)
	}
	buf := bytes.NewBuffer(infoBytes)
	if err := gob.NewDecoder(buf).Decode(info); err != nil {
		return nil, kerr.Wrap("AAFXLQRUEW", err)
	}
	return info, nil
}

func getRawInfo() string {
	return dom.
		GetWindow().
		Document().(dom.HTMLDocument).
		GetElementByID("body").(*dom.HTMLBodyElement).
		GetAttribute("info")
}
