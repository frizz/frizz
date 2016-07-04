package views

import (
	"net/rpc"

	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/editor/shared"
	"kego.io/flux"
	"kego.io/system/node"
)

type BranchView struct {
	*View

	model    *models.BranchModel
	children vecty.List
}

func NewBranchView(ctx context.Context, model *models.BranchModel) *BranchView {
	v := &BranchView{}
	v.View = New(ctx, v)
	v.model = model
	v.Watch(v.reaction, v.model,
		stores.BranchOpening,
		stores.BranchOpened,
		stores.BranchClose,
		stores.BranchLoaded,
		stores.BranchSelecting,
		stores.BranchChildAdded,
		stores.BranchChildDeleted,
		stores.BranchChildrenReordered,
	)
	return v
}

// ke: {"func": {"notest": true}}
func (v *BranchView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *BranchView) reaction(notif flux.NotifPayload) {
	wait := &flux.Waiter{}
	defer wait.Go(notif.Done)

	switch notif.Type {
	case stores.BranchOpening:
		loaded := loadBranch(v.Ctx, v.App, v.model, wait)
		wait.Add(v.App.Dispatch(&actions.BranchOpened{Branch: v.model, Loaded: loaded}))
	case stores.BranchSelecting:
		loaded := loadBranch(v.Ctx, v.App, v.model, wait)
		data := notif.Data.(*stores.BranchSelectOperationData)
		if loaded && data.Op == models.BranchOpClickToggle {
			wait.Add(v.App.Dispatch(&actions.BranchOpening{Branch: v.model}))
		}
		wait.Add(v.App.Dispatch(&actions.BranchSelected{Branch: v.model, Loaded: loaded}))
	case stores.BranchOpened,
		stores.BranchClose,
		stores.BranchLoaded,
		stores.BranchChildAdded,
		stores.BranchChildDeleted,
		stores.BranchChildrenReordered:
		v.ReconcileBody()
		if ds, ok := notif.Data.(*stores.BranchDescendantSelectData); ok {
			wait.Add(v.App.Dispatch(&actions.BranchSelecting{Branch: ds.Branch, Op: ds.Op}))
		}
	}
}

func loadBranch(ctx context.Context, app *stores.App, b *models.BranchModel, wait *flux.Waiter) bool {
	c, ok := b.Contents.(*models.SourceContents)
	if !ok {
		return false
	}
	loaded := false
	c.Once.Do(func() {
		request := &shared.DataRequest{
			File:    c.Filename,
			Name:    c.Name,
			Package: envctx.FromContext(ctx).Path,
		}
		data := shared.DataResponse{}
		done := make(chan *rpc.Call, 1)
		call := app.Conn.Go("Server.Data", request, &data, done, app.Fail)
		wait.Add(app.Dispatcher.Dispatch(&actions.LoadSourceSent{Branch: b}))

		<-call.Done

		gr, ok := call.Reply.(*shared.DataResponse)
		if !ok {
			app.Fail <- kerr.New("OCVFGLPIQG", "%T is not a *shared.DataResponse", call.Reply)
			return
		}

		n, err := node.Unmarshal(ctx, gr.Data)
		if err != nil {
			app.Fail <- kerr.Wrap("IOOQWKIEGC", err)
			return
		}

		c.Node = n

		wait.Add(app.Dispatcher.Dispatch(&actions.LoadSourceSuccess{Branch: b}))

		loaded = true
	})

	if !loaded {
		wait.Add(app.Dispatcher.Dispatch(&actions.LoadSourceCancelled{Branch: b}))
	}
	return loaded

}

func (v *BranchView) Render() vecty.Component {
	if v.model == nil {
		return elem.Div()
	}

	v.children = vecty.List{}
	if v.model.Open {
		for _, c := range v.model.Children {
			v.children = append(v.children, NewBranchView(v.Ctx, c))
		}
	}

	return elem.Div(
		prop.Class("node"),
		NewBranchControlView(v.Ctx, v.model),
		elem.Div(
			prop.Class("children"),
			v.children,
		),
	)
}
