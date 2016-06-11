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
	vecty.Composite
	ctx context.Context
	app *stores.App

	model    *models.BranchModel
	notifs   chan flux.NotifPayload
	children vecty.List
}

func NewBranchView(ctx context.Context, model *models.BranchModel) *BranchView {
	v := &BranchView{
		ctx:   ctx,
		app:   stores.FromContext(ctx),
		model: model,
	}
	v.Mount()
	return v
}

// ke: {"func": {"notest": true}}
func (v *BranchView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *BranchView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *BranchView) Mount() {

	v.notifs = v.app.Watch(v.model,
		stores.BranchOpening,
		stores.BranchOpened,
		stores.BranchClose,
		stores.BranchLoaded,
		stores.BranchSelecting,
		stores.BranchChildAdded,
		stores.BranchChildDeleted,
		stores.BranchChildrenReordered,
	)

	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *BranchView) reaction(notif flux.NotifPayload) {
	wait := &flux.Waiter{}
	defer wait.Go(notif.Done)

	switch notif.Type {
	case stores.BranchOpening:
		loaded := loadBranch(v.ctx, v.app, v.model, wait)
		wait.Add(v.app.Dispatch(&actions.BranchOpened{Branch: v.model, Loaded: loaded}))
	case stores.BranchSelecting:
		loaded := loadBranch(v.ctx, v.app, v.model, wait)
		data := notif.Data.(*stores.BranchSelectOperationData)
		if loaded && data.Op == models.BranchOpClickToggle {
			wait.Add(v.app.Dispatch(&actions.BranchOpening{Branch: v.model}))
		}
		wait.Add(v.app.Dispatch(&actions.BranchSelected{Branch: v.model, Loaded: loaded}))
	case stores.BranchOpened,
		stores.BranchClose,
		stores.BranchLoaded,
		stores.BranchChildAdded,
		stores.BranchChildDeleted,
		stores.BranchChildrenReordered:
		v.ReconcileBody()
		if ds, ok := notif.Data.(*stores.BranchDescendantSelectData); ok {
			wait.Add(v.app.Dispatch(&actions.BranchSelecting{Branch: ds.Branch, Op: ds.Op}))
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

func (v *BranchView) Unmount() {
	if v.notifs != nil {
		v.app.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *BranchView) render() vecty.Component {
	if v.model == nil {
		return elem.Div()
	}

	v.children = vecty.List{}
	if v.model.Open {
		for _, c := range v.model.Children {
			v.children = append(v.children, NewBranchView(v.ctx, c))
		}
	}

	return elem.Div(
		prop.Class("node"),
		NewBranchControlView(v.ctx, v.model),
		elem.Div(
			prop.Class("children"),
			v.children,
		),
	)
}
