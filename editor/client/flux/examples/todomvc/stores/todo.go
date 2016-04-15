package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/flux/examples/todomvc/actions"
	"kego.io/editor/client/flux/examples/todomvc/model"
)

type TodoStore struct {
	*flux.Store
	ctx context.Context
	app *App

	items  []*model.Item
	filter model.FilterState
}

type todoKey string

const TodoChange todoKey = "TodoChange"

func NewTodoStore(ctx context.Context) *TodoStore {
	return &TodoStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
}

func (t *TodoStore) Filter() model.FilterState {
	return t.filter
}

func (t *TodoStore) Items() []*model.Item {
	return t.items
}

func (t *TodoStore) ActiveItemCount() int {
	return t.count(false)
}

func (t *TodoStore) CompletedItemCount() int {
	return t.count(true)
}

func (t *TodoStore) count(completed bool) int {
	count := 0
	for _, item := range t.items {
		if item.Complete == completed {
			count++
		}
	}
	return count
}

func (t *TodoStore) Handle(payload *flux.Payload) (finished bool) {
	switch a := payload.Action.(type) {
	case *actions.ReplaceItems:
		t.items = a.Items

	case *actions.Create:
		t.items = append(t.items, &model.Item{Title: a.Text, Complete: false})

	case *actions.Destroy:
		t.items = append(t.items[:a.Index], t.items[a.Index+1:]...)

	case *actions.UpdateText:
		t.items[a.Index].Title = a.Text

	case *actions.SetComplete:
		t.items[a.Index].Complete = a.Complete

	case *actions.SetCompleteAll:
		for _, item := range t.items {
			item.Complete = a.Complete
		}

	case *actions.DestroyCompleted:
		var activeItems []*model.Item
		for _, item := range t.items {
			if !item.Complete {
				activeItems = append(activeItems, item)
			}
		}
		t.items = activeItems

	case *actions.SetFilter:
		t.filter = a.Filter

	default:
		return true // don't notify
	}

	t.Notify(TodoChange)
	return true
}
