package views

import (
	"testing"

	"reflect"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/ctests"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/editor/shared"
	"kego.io/ke"
	"kego.io/process/parser"
	"kego.io/system"
)

func TestBranchRender1(t *testing.T) {
	cb := ctests.New().App()
	b := NewBranchView(cb.Ctx(), nil)
	expected := elem.Div()
	equal(t, expected, b.render().(*vecty.Element))
}

func TestBranchRender2(t *testing.T) {
	cb := ctests.New().App()
	b := NewBranchView(cb.Ctx(), &models.BranchModel{
		Contents: &models.RootContents{Name: "a"},
	})
	expected := elem.Div(
		prop.Class("node"),
		&BranchControlView{},
		elem.Div(
			prop.Class("children"),
		),
	)
	equal(t, expected, b.render().(*vecty.Element))

	b.model.Open = false
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

func TestBranchNotify1(t *testing.T) {
	cb := ctests.New().App()

	app := stores.FromContext(cb.Ctx())

	m := &models.BranchModel{
		Contents: &models.RootContents{Name: "a"},
	}
	b := NewBranchView(cb.Ctx(), m)

	app.Branches.Notify(b.model, stores.BranchOpen)

	assert.Equal(t, 0, len(cb.Conn().Log))
	shouldDispatch(t, cb, new(actions.BranchOpenPostLoad))
	a := cb.GetDispatcher().Log[0].(*actions.BranchOpenPostLoad)
	assert.Equal(t, m, a.Branch)
	assert.Equal(t, false, a.Loaded)
}
func TestBranchNotify2(t *testing.T) {

	cb := ctests.New().App()

	// We have to unmarshal the
	cb.Base.Path("").Jauto().Sauto(parser.Parse)

	app := stores.FromContext(cb.Ctx())

	m := &models.BranchModel{
		Contents: &models.SourceContents{Name: "a", Filename: "b"},
	}

	b := NewBranchView(cb.Ctx(), m)

	var bytes []byte
	bytes, err := ke.MarshalContext(cb.Ctx(), &system.Object{Type: system.NewReference("system", "object")})
	assert.NoError(t, err)
	cb.ConnReply(&shared.DataResponse{Data: bytes})

	_, done := app.Branches.Notify(b.model, stores.BranchOpen)
	<-done

	assert.Equal(t, 1, len(cb.Conn().Log))
	shouldDispatch(t, cb,
		new(actions.LoadSourceSent),
		new(actions.LoadSourceSuccess),
		new(actions.BranchOpenPostLoad),
	)

}

func TestBranchNotify3(t *testing.T) {

	cb := ctests.New().App()
	cb.Base.Path("")

	app := stores.FromContext(cb.Ctx())

	m := &models.BranchModel{
		Contents: &models.SourceContents{Name: "a", Filename: "b"},
	}

	b := NewBranchView(cb.Ctx(), m)

	cb.ConnReply("a")

	_, done := app.Branches.Notify(b.model, stores.BranchOpen)
	<-done

	err := readErrorOrNil(app.Fail)
	assert.IsError(t, err, "OCVFGLPIQG")

	shouldDispatch(t, cb,
		new(actions.LoadSourceSent),
		new(actions.LoadSourceCancelled),
		new(actions.BranchOpenPostLoad),
	)

}

func shouldDispatch(t *testing.T, cb *ctests.ClientContextBuilder, expected ...interface{}) {
	assert.Equal(t, len(expected), len(cb.GetDispatcher().Log))
	for i, action := range expected {
		assert.IsType(t, action, cb.GetDispatcher().Log[i])
	}
}

func readErrorOrNil(c chan error) error {
	select {
	case err := <-c:
		return err
	default:
		return nil
	}
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
