package views

import (
	"testing"

	"reflect"

	"fmt"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/mock_vecty"
	"github.com/davelondon/vecty/prop"
	"github.com/golang/mock/gomock"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/ctests"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/editor/shared"
	"kego.io/flux"
	"kego.io/ke"
	"kego.io/process/parser"
	"kego.io/system"
)

func TestBranchRender1(t *testing.T) {
	cb := ctests.New().SetApp()
	b := NewBranchView(cb.Ctx(), nil)
	expected := elem.Div()
	equal(t, expected, b.render().(*vecty.Element))
}

func TestBranchRender2(t *testing.T) {
	cb := ctests.New().SetApp()
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

func TestBranchNotifyOpen(t *testing.T) {
	cb := ctests.New().SetApp()
	b := NewBranchView(cb.Ctx(), &models.BranchModel{
		Contents: &models.RootContents{Name: "a"},
	})
	done := cb.GetApp().Branches.Notify(b.model, stores.BranchOpen)
	<-done
	assert.Equal(t, 0, len(cb.GetConn().Log))
	shouldDispatch(t, cb, new(actions.BranchOpenPostLoad))
	a := cb.GetDispatcher().Log[0].(*actions.BranchOpenPostLoad)
	assert.Equal(t, b.model, a.Branch)
	assert.Equal(t, false, a.Loaded)
}

func TestBranchNotifyOpenSource(t *testing.T) {
	cb := ctests.New().SetApp()

	setupForSuccessfulSourceLoad(t, cb)

	b := NewBranchView(cb.Ctx(), &models.BranchModel{
		Contents: &models.SourceContents{Name: "a", Filename: "b"},
	})

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchOpen)
	<-done

	assert.Equal(t, 1, len(cb.GetConn().Log))
	shouldDispatch(t, cb,
		new(actions.LoadSourceSent),
		new(actions.LoadSourceSuccess),
		new(actions.BranchOpenPostLoad),
	)

}

func setupForSuccessfulSourceLoad(t *testing.T, cb *ctests.ClientContextBuilder) {

	// We have to unmarshal the received object, so we'll have to load some types.
	cb.Base.Path("").Jauto().Sauto(parser.Parse)

	// Create a simple ke object and marshal it to a []byte
	var bytes []byte
	bytes, err := ke.MarshalContext(cb.Ctx(), &system.Object{Type: system.NewReference("system", "object")})
	assert.NoError(t, err)
	cb.SetConnReply(&shared.DataResponse{Data: bytes})

}

func TestBranchNotifyOpenSourceError(t *testing.T) {

	cb := ctests.New().SetApp()
	cb.Base.Path("")

	b := NewBranchView(cb.Ctx(), &models.BranchModel{
		Contents: &models.SourceContents{Name: "a", Filename: "b"},
	})

	cb.SetConnReply("a")

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchOpen)
	<-done

	assert.IsError(t, readErrorOrNil(cb.GetApp().Fail), "OCVFGLPIQG")

	shouldDispatch(t, cb,
		new(actions.LoadSourceSent),
		new(actions.LoadSourceCancelled),
		new(actions.BranchOpenPostLoad),
	)

}

func TestBranchNotifySelect(t *testing.T) {
	cb := ctests.New().SetApp()

	b := NewBranchView(cb.Ctx(), &models.BranchModel{
		Contents: &models.RootContents{Name: "a"},
	})

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchSelect)
	<-done

	assert.NoError(t, readErrorOrNil(cb.GetApp().Fail))

	shouldDispatch(t, cb,
		new(actions.BranchSelectPostLoad),
	)
}

func TestBranchNotifySelectSource(t *testing.T) {
	cb := ctests.New().SetApp()

	setupForSuccessfulSourceLoad(t, cb)

	b := NewBranchView(cb.Ctx(), &models.BranchModel{
		Contents: &models.SourceContents{Name: "a", Filename: "b"},
	})

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchSelect)
	<-done

	assert.NoError(t, readErrorOrNil(cb.GetApp().Fail))

	shouldDispatch(t, cb,
		new(actions.LoadSourceSent),
		new(actions.LoadSourceSuccess),
		new(actions.BranchSelectPostLoad),
	)

}

func TestBranchNotifySelectSourceClick(t *testing.T) {
	cb := ctests.New().SetApp()

	setupForSuccessfulSourceLoad(t, cb)

	b := NewBranchView(cb.Ctx(), &models.BranchModel{
		LastOp:   models.BranchOpClickToggle,
		Contents: &models.SourceContents{Name: "a", Filename: "b"},
	})

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchSelect)
	<-done

	assert.NoError(t, readErrorOrNil(cb.GetApp().Fail))

	shouldDispatch(t, cb,
		new(actions.LoadSourceSent),
		new(actions.LoadSourceSuccess),
		new(actions.BranchOpen),
		new(actions.BranchSelectPostLoad),
	)

}

func TestBranchReconcile(t *testing.T) {
	testBranchReconcile(t, stores.BranchClose)
	testBranchReconcile(t, stores.BranchOpenPostLoad)
	testBranchReconcile(t, stores.BranchLoaded)
}
func testBranchReconcile(t *testing.T, notif flux.Notif) {
	cb := ctests.New().SetApp()
	mc := gomock.NewController(t)
	defer mc.Finish()

	b := NewBranchView(cb.Ctx(), &models.BranchModel{
		Contents: &models.RootContents{Name: "a"},
	})

	c := mock_vecty.NewMockComponent(mc)
	c.EXPECT().Reconcile(nil)

	b.RenderFunc = func() vecty.Component { return c }

	done := cb.GetApp().Branches.Notify(b.model, notif)
	<-done

	assert.NoError(t, readErrorOrNil(cb.GetApp().Fail))
	shouldDispatch(t, cb)

}

func TestBranchView_Unmount(t *testing.T) {
	cb := ctests.New().SetApp()
	mc := gomock.NewController(t)
	defer mc.Finish()

	b := NewBranchView(cb.Ctx(), &models.BranchModel{
		Contents: &models.RootContents{Name: "a"},
	})
	c := mock_vecty.NewMockComponent(mc)
	c.EXPECT().Unmount()
	b.Body = c
	b.Unmount()
	assert.Nil(t, b.notifs)
}

func shouldDispatch(t *testing.T, cb *ctests.ClientContextBuilder, expected ...interface{}) {
	if assert.Equal(t, len(expected), len(cb.GetDispatcher().Log)) {
		for i, action := range expected {
			assert.IsType(t, action, cb.GetDispatcher().Log[i])
		}
	} else {
		for _, v := range cb.GetDispatcher().Log {
			fmt.Printf("%T\n", v)
		}
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
