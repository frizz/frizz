package views

import (
	"testing"

	"reflect"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/ctests"
	"kego.io/editor/client/models"
	"kego.io/kerr/assert"
)

func TestBranchRender(t *testing.T) {

	cb := ctests.New().App()

	b := NewBranchView(cb.Ctx(), nil)
	expected := elem.Div()
	equal(t, expected, b.render().(*vecty.Element))

	b = NewBranchView(cb.Ctx(), &models.BranchModel{
		Contents: &models.RootContents{Name: "a"},
	})
	expected = elem.Div(
		prop.Class("node"),
		&BranchControlView{},
		elem.Div(
			prop.Class("children"),
		),
	)
	equal(t, expected, b.render().(*vecty.Element))

	b.model.Children = append(b.model.Children, &models.BranchModel{
		Contents: &models.RootContents{Name: "b"},
	})
	// Extra child but HTML doesn't change because branch is closed.
	equal(t, expected, b.render().(*vecty.Element))

	b.model.Open = true
	expected = elem.Div(
		prop.Class("node"),
		&BranchControlView{},
		elem.Div(
			prop.Class("children"),
			&BranchView{},
		),
	)
	equal(t, expected, b.render().(*vecty.Element))
}

func equal(t *testing.T, expected, actual *vecty.Element) {
	assert.Equal(t, expected.TagName, actual.TagName)

	assert.Equal(t, len(expected.Properties), len(actual.Properties))
	for i, v := range expected.Properties {
		assert.Equal(t, v, actual.Properties[i])
	}

	assert.Equal(t, len(expected.Style), len(actual.Style))
	for i, v := range expected.Style {
		assert.Equal(t, v, actual.Style[i])
	}

	assert.Equal(t, len(expected.Dataset), len(actual.Dataset))
	for i, v := range expected.Dataset {
		assert.Equal(t, v, actual.Dataset[i])
	}

	assert.Equal(t, len(expected.EventListeners), len(actual.EventListeners))
	for i, e := range expected.EventListeners {
		a := actual.EventListeners[i]
		assert.Equal(t, e.Name, a.Name)
		assert.Equal(t, e.Listener, a.Listener)
	}

	assert.Equal(t, len(expected.Children), len(actual.Children))
	for i, v := range expected.Children {
		if e, ok := v.(*vecty.Element); ok {
			a, ok := actual.Children[i].(*vecty.Element)
			if !ok {
				assert.Fail(t, "Should be *vecty.Element")
			}
			equal(t, e, a)
		} else {
			expectedType := reflect.TypeOf(v)
			actualType := reflect.TypeOf(actual.Children[i])
			assert.Equal(t, expectedType, actualType)
		}
		// TODO: should we be comparing non *vecty.Element?
	}
}
