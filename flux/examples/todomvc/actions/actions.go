package actions

import "kego.io/flux/examples/todomvc/model"

type ReplaceItems struct {
	Items []*model.Item
}

type Create struct {
	Text string
}

type UpdateText struct {
	Index int
	Text  string
}

type SetComplete struct {
	Index    int
	Complete bool
}

type SetCompleteAll struct {
	Complete bool
}

type Destroy struct {
	Index int
}

type DestroyCompleted struct{}

type SetFilter struct {
	Filter model.FilterState
}
