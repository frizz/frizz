package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
)

type BranchView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	model    *models.BranchModel
	c        chan struct{}
	children vecty.List
}

func (b *BranchView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchView); ok {
		b.Body = old.Body
		b.model = old.model
	}
	b.RenderFunc = b.render
	b.ReconcileBody()
}

func NewBranchView(ctx context.Context, model *models.BranchModel) *BranchView {
	if model == nil {
		return nil
	}
	app := stores.FromContext(ctx)
	b := &BranchView{
		ctx:   ctx,
		app:   app,
		model: model,
	}
	b.Mount()
	return b
}

func (b *BranchView) Mount() {
	if b.c != nil {
		panic("mounting a mounted branch")
	}
	b.c = b.app.Branches.Changed(b.model)
	go func() {
		for range b.c {
			b.ReconcileBody()
		}
	}()
}

func (b *BranchView) Unmount() {
	if b.c == nil {
		return
	}
	b.app.Branches.Delete(b.c)
	close(b.c)
	b.c = nil
	b.Body.Unmount()
}

// Apply implements the vecty.Markup interface.
func (b *BranchView) Apply(element *vecty.Element) {
	element.AddChild(b)
}

func (b *BranchView) toggleClick(*vecty.Event) {
	go func() {
		if _, success := <-b.ensureLoaded(); !success {
			return
		}
		if b.model.CanOpen() {
			b.app.Dispatch(&actions.ToggleBranch{Branch: b.model})
		} else {
			b.app.Dispatch(&actions.SelectBranch{Branch: b.model, Keyboard: false})
		}
	}()
}

func (b *BranchView) selectClick(*vecty.Event) {
	go func() {
		if _, success := <-b.ensureLoaded(); !success {
			return
		}
		b.app.Dispatch(&actions.SelectBranch{Branch: b.model, Keyboard: false})
	}()
}

func (b *BranchView) ensureLoaded() chan struct{} {

	signal := make(chan struct{}, 1)

	if !b.model.Contents.Async() || b.model.Contents.Loaded() {
		// If item is not async or content is already loaded, send to the signal before returning.
		signal <- struct{}{}
		return signal
	}

	// Load content asynchronously
	action := &actions.LoadSource{
		Contents: b.model.Contents,
		Signal:   signal,
	}

	go func() {
		b.app.Dispatch(action)
	}()

	return signal
}

func (b *BranchView) render() vecty.Component {
	if b.model == nil {
		return elem.Div()
	}

	selected := b.app.Branches.Selected() == b.model

	b.children = vecty.List{}
	if b.model.Open {
		for _, c := range b.model.Children {
			b.children = append(b.children, NewBranchView(b.ctx, c))
		}
	}

	icon := b.model.Icon()

	return elem.Div(
		prop.Class("node"),
		elem.Anchor(
			vecty.ClassMap{
				"toggle":   true,
				"selected": selected,
				"plus":     icon == "plus",
				"minus":    icon == "minus",
				"unknown":  icon == "unknown",
				"empty":    icon == "empty",
			},
			event.Click(b.toggleClick),
		),
		elem.Div(
			vecty.ClassMap{
				"node-content": true,
				"selected":     selected,
			},
			elem.Span(
				prop.Class("node-label"),
				event.Click(b.selectClick),
				vecty.Text(b.model.Contents.Label()),
			),
			elem.Span(
				prop.Class("badge"),
				vecty.Style("display", "none"),
			),
		),
		elem.Div(
			prop.Class("children"),
			b.children,
		),
	)
}
