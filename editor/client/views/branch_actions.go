package views

import (
	"net/rpc"
	"time"

	"code.google.com/p/go.net/context"
	"kego.io/context/envctx"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/editor/shared"
	"kego.io/kerr"
	"kego.io/system/node"
)

func LoadBranchDebounced(ctx context.Context, app *stores.App, b *models.BranchModel) {
	if async, ok := b.Contents.(models.AsyncInterface); !ok || async.Loaded() {
		return
	}
	go func() {
		<-time.After(time.Millisecond * 50)
		if app.Branches.Selected() == b {
			LoadBranch(ctx, app, b, false)
		}
	}()
}

func LoadBranch(ctx context.Context, app *stores.App, b *models.BranchModel, openBranch bool) {
	c, ok := b.Contents.(*models.SourceContents)
	if !ok {
		return
	}
	c.Once.Do(func() {
		request := &shared.DataRequest{
			File:    c.Filename,
			Name:    c.Name,
			Package: envctx.FromContext(ctx).Path,
		}
		data := shared.DataResponse{}
		done := make(chan *rpc.Call, 1)
		call := app.Conn.Go("Server.Data", request, &data, done, app.Fail)

		app.Dispatcher.Dispatch(&actions.LoadSourceSent{Branch: b})

		<-call.Done

		gr, ok := call.Reply.(*shared.DataResponse)
		if !ok {
			app.Fail <- kerr.New("OCVFGLPIQG", "%T is not a *shared.DataResponse", call.Reply)
		}

		n, err := node.Unmarshal(ctx, gr.Data)
		if err != nil {
			app.Fail <- kerr.Wrap("IOOQWKIEGC", err)
		}

		c.Node = n

		app.Dispatcher.Dispatch(&actions.LoadSourceSuccess{Branch: b})
		if openBranch {
			app.Dispatcher.Dispatch(&actions.OpenBranch{Branch: b})
		}
		/**


		ed := s.Node.Editor()

		if err := ed.Initialize(s.tree.ctx, s, editor.Page, s.tree.Fail); err != nil {
			return kerr.Wrap("WXAHRIGKHI", err)
		}

		s.ListenForEditorChanges(ed.Listen().Ch)

		if err := addEntryChildren(s.Node, s, ed); err != nil {
			return kerr.Wrap("IQMWLPORUF", err)
		}

		return nil
		*/

	})

}
