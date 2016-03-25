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

	model *models.BranchModel
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
	return b
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

	selected := b.app.Selected.Get() == b.model

	var childern vecty.List
	for _, c := range b.model.Children {
		childern = append(childern, NewBranchView(b.ctx, c))
	}

	if b.model.Root {
		return elem.Div(
			prop.Class("node root"),
			elem.Div(
				prop.Class("children"),
				childern,
			),
		)
	}

	icon := b.model.Icon()

	return elem.Div(
		prop.Class("node"),
		elem.Anchor(
			vecty.ClassMap{
				"toggle":  true,
				"plus":    icon == "plus",
				"minus":   icon == "minus",
				"unknown": icon == "unknown",
				"empty":   icon == "empty",
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
			vecty.If(!b.model.Open, vecty.Style("display", "none")),
			childern,
		),
	)
}
