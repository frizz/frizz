package views

import (
	"net/rpc"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/editor/shared"
	"kego.io/kerr"
	"kego.io/system/node"
)

func LoadBranch(ctx context.Context, app *stores.App, b *models.BranchModel) bool {
	c, ok := b.Contents.(*models.SourceContents)
	if !ok {
		return false
	}
	did := false
	c.Once.Do(func() {
		did = true
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

	})

	if !did {
		app.Dispatcher.Dispatch(&actions.LoadSourceCancelled{Branch: b})
	}

	return did

}
