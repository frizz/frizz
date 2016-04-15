package views

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"github.com/davelondon/vecty/style"
	"kego.io/editor/client/flux/examples/todomvc/actions"
	"kego.io/editor/client/flux/examples/todomvc/model"
	"kego.io/editor/client/flux/examples/todomvc/stores"
)

type PageView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Items        []*model.Item
	newItemTitle string
}

func NewPageView(ctx context.Context) *PageView {
	p := &PageView{
		ctx:   ctx,
		app:   stores.FromContext(ctx),
		Items: []*model.Item{},
	}

	go func() {
		for _ = range p.app.Todos.Watch(stores.TodoChange) {
			p.Items = p.app.Todos.Items()
			p.ReconcileBody()
		}
	}()

	return p
}

// Apply implements the vecty.Markup interface.
func (p *PageView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *PageView) Reconcile(old vecty.Component) {
	if old, ok := old.(*PageView); ok {
		p.Body = old.Body
		p.Items = old.Items
		p.newItemTitle = old.newItemTitle
	}
	p.RenderFunc = p.render
	p.ReconcileBody()
}

func (p *PageView) onNewItemTitleInput(event *vecty.Event) {
	p.newItemTitle = event.Target.Get("value").String()
	p.ReconcileBody()
}

func (p *PageView) onAdd(event *vecty.Event) {
	go func() {
		p.app.Dispatcher.Dispatch(&actions.Create{Text: p.newItemTitle})
		p.newItemTitle = ""
		p.ReconcileBody()
	}()
}

func (p *PageView) onClearCompleted(event *vecty.Event) {
	go func() {
		p.app.Dispatcher.Dispatch(&actions.DestroyCompleted{})
	}()
}

func (p *PageView) onToggleAllCompleted(event *vecty.Event) {
	go func() {
		p.app.Dispatcher.Dispatch(&actions.SetCompleteAll{Complete: event.Target.Get("checked").Bool()})
	}()
}

func (p *PageView) render() vecty.Component {
	return elem.Div(
		elem.Section(
			prop.Class("todoapp"),

			p.renderHeader(),
			vecty.If(len(p.app.Todos.Items()) > 0,
				p.renderItemList(),
				p.renderFooter(),
			),
		),

		p.renderInfo(),
	)
}

func (p *PageView) renderHeader() vecty.Markup {
	return elem.Header(
		prop.Class("header"),

		elem.Header1(
			vecty.Text("todos"),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			event.Submit(p.onAdd).PreventDefault(),

			elem.Input(
				prop.Class("new-todo"),
				prop.Placeholder("What needs to be done?"),
				prop.Autofocus(true),
				prop.Value(p.newItemTitle),
				event.Input(p.onNewItemTitleInput),
			),
		),
	)
}

func (p *PageView) renderFooter() vecty.Component {
	count := p.app.Todos.ActiveItemCount()
	var itemsLeftText = " items left"
	if count == 1 {
		itemsLeftText = " item left"
	}

	return elem.Footer(
		prop.Class("footer"),

		elem.Span(
			prop.Class("todo-count"),

			elem.Strong(
				vecty.Text(fmt.Sprintf("%d", count)),
			),
			vecty.Text(itemsLeftText),
		),

		elem.UnorderedList(
			prop.Class("filters"),
			NewFilterButton(p.ctx, "All", model.All),
			vecty.Text(" "),
			NewFilterButton(p.ctx, "Active", model.Active),
			vecty.Text(" "),
			NewFilterButton(p.ctx, "Completed", model.Completed),
		),

		vecty.If(p.app.Todos.CompletedItemCount() > 0,
			elem.Button(
				prop.Class("clear-completed"),
				vecty.Text(fmt.Sprintf("Clear completed (%d)", p.app.Todos.CompletedItemCount())),
				event.Click(p.onClearCompleted),
			),
		),
	)
}

func (p *PageView) renderInfo() vecty.Component {
	return elem.Footer(
		prop.Class("info"),

		elem.Paragraph(
			vecty.Text("Double-click to edit a todo"),
		),
		elem.Paragraph(
			vecty.Text("Created by "),
			elem.Anchor(
				prop.Href("http://github.com/neelance"),
				vecty.Text("Richard Musiol"),
			),
			vecty.Text(", "),
			elem.Anchor(
				prop.Href("http://github.com/davelondon"),
				vecty.Text("Dave Brophy"),
			),
		),
		elem.Paragraph(
			vecty.Text("Part of "),
			elem.Anchor(
				prop.Href("http://todomvc.com"),
				vecty.Text("TodoMVC"),
			),
		),
	)
}

func (p *PageView) renderItemList() vecty.Component {
	var items vecty.List
	for i, item := range p.app.Todos.Items() {
		if (p.app.Todos.Filter() == model.Active && item.Complete) || (p.app.Todos.Filter() == model.Completed && !item.Complete) {
			continue
		}
		items = append(items, NewItemView(p.ctx, i, item))
	}

	return elem.Section(
		prop.Class("main"),

		elem.Input(
			prop.ID("toggle-all"),
			prop.Class("toggle-all"),
			prop.Type(prop.TypeCheckbox),
			prop.Checked(p.app.Todos.CompletedItemCount() == len(p.app.Todos.Items())),
			event.Change(p.onToggleAllCompleted),
		),
		elem.Label(
			prop.For("toggle-all"),
			vecty.Text("Mark all as complete"),
		),

		elem.UnorderedList(
			prop.Class("todo-list"),
			items,
		),
	)
}
