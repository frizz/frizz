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

	v.notifs = v.app.Branches.Watch(v.model,
		stores.BranchOpen,
		stores.BranchOpenPostLoad,
		stores.BranchClose,
		stores.BranchLoaded,
		stores.BranchSelect,
		stores.BranchChildAdded,
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
	case stores.BranchOpen:
		loaded := loadBranch(v.ctx, v.app, v.model, wait)
		wait.Add(v.app.Dispatch(&actions.BranchOpenPostLoad{Branch: v.model, Loaded: loaded}))
	case stores.BranchSelect:
		loaded := loadBranch(v.ctx, v.app, v.model, wait)
		if v.model.LastOp == models.BranchOpClickToggle && loaded {
			wait.Add(v.app.Dispatch(&actions.BranchOpen{Branch: v.model}))
		}
		wait.Add(v.app.Dispatch(&actions.BranchSelectPostLoad{Branch: v.model, Loaded: loaded}))
	case stores.BranchOpenPostLoad,
		stores.BranchClose,
		stores.BranchLoaded,
		stores.BranchChildAdded:
		v.ReconcileBody()
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
		v.app.Branches.Delete(v.notifs)
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
