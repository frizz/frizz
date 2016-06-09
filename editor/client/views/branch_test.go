package views

import (
	"net/rpc"
	"testing"

	"reflect"

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

	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), nil)
	expected := elem.Div()
	equal(t, expected, b.render().(*vecty.Element))

	cb.AssertAppSuccess()

}

func TestBranchRender2(t *testing.T) {

	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))
	expected := elem.Div(
		prop.Class("node"),
		&BranchControlView{},
		elem.Div(
			prop.Class("children"),
		),
	)
	equal(t, expected, b.render().(*vecty.Element))

	b.model.Open = false
	b.model.Children = append(b.model.Children, models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "b"}))
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

	cb.AssertAppSuccess()

}

func TestBranchNotifyOpen(t *testing.T) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))

	cb.GetConnection().EXPECT().Go(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	cb.ExpectDispatched(&actions.BranchOpened{Branch: b.model, Loaded: false})

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchOpening)
	<-done

	cb.AssertAppSuccess()
}

func TestBranchNotifyOpenSource(t *testing.T) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	setupForSuccessfulSourceLoad(t, cb)

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.SourceContents{Name: "a", Filename: "b"}))

	cb.ExpectDispatched(
		&actions.LoadSourceSent{Branch: b.model},
		&actions.LoadSourceSuccess{Branch: b.model},
		&actions.BranchOpened{Branch: b.model, Loaded: true},
	)

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchOpening)
	<-done

	cb.AssertAppSuccess()

}

func TestBranchNotifyOpenSourceError(t *testing.T) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	setupForFailedSourceLoad(cb)

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.SourceContents{Name: "a", Filename: "b"}))

	cb.ExpectDispatched(
		&actions.LoadSourceSent{Branch: b.model},
		&actions.LoadSourceCancelled{Branch: b.model},
		&actions.BranchOpened{Branch: b.model, Loaded: false},
	)

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchOpening)
	<-done

	cb.AssertAppFail("OCVFGLPIQG")

}

func TestBranchNotifySelect(t *testing.T) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))

	cb.ExpectDispatched(
		&actions.BranchSelected{Branch: b.model, Loaded: false},
	)

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchSelecting)
	<-done

	cb.AssertAppSuccess()

}

func TestBranchNotifySelectSource(t *testing.T) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	setupForSuccessfulSourceLoad(t, cb)

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.SourceContents{Name: "a", Filename: "b"}))

	cb.ExpectDispatched(
		&actions.LoadSourceSent{Branch: b.model},
		&actions.LoadSourceSuccess{Branch: b.model},
		&actions.BranchSelected{Branch: b.model, Loaded: true},
	)

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchSelecting)
	<-done

	cb.AssertAppSuccess()

}

func TestBranchNotifySelectSourceClick(t *testing.T) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	setupForSuccessfulSourceLoad(t, cb)

	m := models.NewBranchModel(cb.Ctx(), &models.SourceContents{Name: "a", Filename: "b"})
	m.LastOp = models.BranchOpClickToggle
	b := NewBranchView(cb.Ctx(), m)

	cb.ExpectDispatched(
		&actions.LoadSourceSent{Branch: b.model},
		&actions.LoadSourceSuccess{Branch: b.model},
		&actions.BranchOpening{Branch: b.model},
		&actions.BranchSelected{Branch: b.model, Loaded: true},
	)

	done := cb.GetApp().Branches.Notify(b.model, stores.BranchSelecting)
	<-done

	cb.AssertAppSuccess()

}

func TestBranchReconcile(t *testing.T) {
	testBranchReconcile(t, stores.BranchClose)
	testBranchReconcile(t, stores.BranchOpened)
	testBranchReconcile(t, stores.BranchLoaded)
}

func testBranchReconcile(t *testing.T, notif flux.Notif) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))

	cb.ExpectReconcile(&b.Composite)
	cb.ExpectNoneDispatched()

	done := cb.GetApp().Branches.Notify(b.model, notif)
	<-done

	cb.AssertAppSuccess()

}

func TestBranchView_Unmount(t *testing.T) {
	cb := ctests.New(t).SetApp()
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))

	c := mock_vecty.NewMockComponent(cb.GetMock())
	c.EXPECT().Unmount()
	b.Body = c
	b.Unmount()
	assert.Nil(t, b.notifs)
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
		// TODO: Compare events?
		//assert.Equal(t, e.Listener, a.Listener)
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

func setupForFailedSourceLoad(cb *ctests.ClientContextBuilder) {

	// We have to unmarshal the received object, so we'll have to load some types.
	cb.Base.Path("")

	done := make(chan *rpc.Call, 1)

	reply := &rpc.Call{
		ServiceMethod: "Server.Data",
		Args:          nil,
		Reply:         "a",
		Error:         nil,
		Done:          done,
	}

	done <- reply

	cb.GetConnection().EXPECT().Go("Server.Data", gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(reply)

}

func setupForSuccessfulSourceLoad(t *testing.T, cb *ctests.ClientContextBuilder) {

	// We have to unmarshal the received object, so we'll have to load some types.
	cb.Base.Path("").Jauto().Sauto(parser.Parse)

	// Create a simple ke object and marshal it to a []byte
	var bytes []byte
	bytes, err := ke.MarshalContext(cb.Ctx(), &system.Object{Type: system.NewReference("system", "object")})
	assert.NoError(t, err)

	done := make(chan *rpc.Call, 1)

	reply := &rpc.Call{
		ServiceMethod: "Server.Data",
		Args:          nil,
		Reply:         &shared.DataResponse{Data: bytes},
		Error:         nil,
		Done:          done,
	}

	done <- reply

	cb.GetConnection().EXPECT().Go("Server.Data", gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(reply)

}
