package views

import (
	"testing"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/ctests"
	"kego.io/editor/client/models"
)

func TestBranchControlView_Render_Nil(t *testing.T) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	b := NewBranchControlView(cb.Ctx(), nil)
	expected := elem.Div()
	equal(t, expected, b.Render().(*vecty.Element))

	cb.AssertAppSuccess()
}

func TestBranchControlView_Render(t *testing.T) {

	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	b := NewBranchControlView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))

	expected := elem.Div(
		elem.Anchor(
			vecty.ClassMap{
				"toggle": true,
				"empty":  true,
			},
			event.Click(nil),
		),
		elem.Div(
			vecty.ClassMap{
				"node-content": true,
			},
			elem.Span(
				prop.Class("node-label"),
				event.Click(nil),
				vecty.Text("a"),
			),
			elem.Span(
				prop.Class("badge"),
				vecty.Style("display", "none"),
			),
		),
	)
	equal(t, expected, b.Render().(*vecty.Element))

	cb.AssertAppSuccess()

}
