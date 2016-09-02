package views

import (
	"testing"

	"reflect"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/mock_vecty"
	"github.com/davelondon/vecty/prop"
	"github.com/go-errors/errors"
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

var branchViewWatchNotifs = []interface{}{
	stores.BranchOpening,
	stores.BranchOpened,
	stores.BranchClose,
	stores.BranchLoaded,
	stores.BranchSelecting,
	stores.BranchChildAdded,
	stores.BranchChildDeleted,
	stores.BranchChildrenReordered,
}

func TestBranchRender1(t *testing.T) {

	cb := ctests.New(t).SetApp(true, true, false)
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), nil)
	expected := elem.Div()
	equal(t, expected, b.Render().(*vecty.Element))

	cb.AssertAppSuccess()

}

func TestBranchRender2(t *testing.T) {

	cb := ctests.New(t).SetApp(true, true, false)
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))
	expected := elem.Div(
		prop.Class("node"),
		NewBranchControlView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), nil)),
		elem.Div(
			prop.Class("children"),
		),
	)
	equal(t, expected, b.Render().(*vecty.Element))

}

func TestBranchNotifyOpen(t *testing.T) {
	cb := ctests.New(t).SetApp(true, false, false)
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))

	cb.GetConnection().EXPECT().Go(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	cb.ExpectDispatched(&actions.BranchOpened{Branch: b.model, Loaded: false})

	done := cb.GetApp().Notifier.Notify(b.model, stores.BranchOpening)
	<-done

	cb.AssertAppSuccess()
}

func TestBranchNotifyOpenFile(t *testing.T) {
	cb := ctests.New(t).SetApp(true, false, false)
	defer cb.Finish()

	setupForSuccessfulFileLoad(t, cb)

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.AsyncContents{NodeContents: models.NodeContents{Name: "a"}, Filename: "b"}))

	cb.ExpectDispatched(
		&actions.LoadFileSent{Branch: b.model},
		&actions.LoadFileSuccess{Branch: b.model},
		&actions.BranchOpened{Branch: b.model, Loaded: true},
	)

	done := cb.GetApp().Notifier.Notify(b.model, stores.BranchOpening)
	<-done

	cb.AssertAppSuccess()

}

func TestBranchNotifyOpenFileError(t *testing.T) {
	cb := ctests.New(t).SetApp(true, false, false)
	defer cb.Finish()

	setupForFailedFileLoad(cb)

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.AsyncContents{NodeContents: models.NodeContents{Name: "a"}, Filename: "b"}))

	cb.ExpectDispatched(
		&actions.LoadFileSent{Branch: b.model},
		&actions.LoadFileCancelled{Branch: b.model},
		&actions.BranchOpened{Branch: b.model, Loaded: false},
	)

	done := cb.GetApp().Notifier.Notify(b.model, stores.BranchOpening)
	<-done

	cb.AssertAppFail("XQQEKDDQSL")

}

func TestBranchNotifySelect(t *testing.T) {
	cb := ctests.New(t).SetApp(true, false, false)
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))

	cb.ExpectDispatched(
		&actions.BranchSelected{Branch: b.model, Loaded: false},
	)

	done := cb.GetApp().Notifier.NotifyWithData(
		b.model,
		stores.BranchSelecting,
		&stores.BranchSelectOperationData{Op: models.BranchOpClickLabel},
	)
	<-done

	cb.AssertAppSuccess()

}

func TestBranchNotifySelectFile(t *testing.T) {
	cb := ctests.New(t).SetApp(true, false, false)
	defer cb.Finish()

	setupForSuccessfulFileLoad(t, cb)

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.AsyncContents{NodeContents: models.NodeContents{Name: "a"}, Filename: "b"}))

	cb.ExpectDispatched(
		&actions.LoadFileSent{Branch: b.model},
		&actions.LoadFileSuccess{Branch: b.model},
		&actions.BranchSelected{Branch: b.model, Loaded: true},
	)

	done := cb.GetApp().Notifier.NotifyWithData(
		b.model,
		stores.BranchSelecting,
		&stores.BranchSelectOperationData{Op: models.BranchOpClickLabel},
	)
	<-done

	cb.AssertAppSuccess()

}

func TestBranchNotifySelectFileClick(t *testing.T) {
	cb := ctests.New(t).SetApp(true, false, false)
	defer cb.Finish()

	setupForSuccessfulFileLoad(t, cb)

	m := models.NewBranchModel(cb.Ctx(), &models.AsyncContents{NodeContents: models.NodeContents{Name: "a"}, Filename: "b"})
	b := NewBranchView(cb.Ctx(), m)

	cb.ExpectDispatched(
		&actions.LoadFileSent{Branch: b.model},
		&actions.LoadFileSuccess{Branch: b.model},
		&actions.BranchOpening{Branch: b.model},
		&actions.BranchSelected{Branch: b.model, Loaded: true},
	)

	done := cb.GetApp().Notifier.NotifyWithData(
		b.model,
		stores.BranchSelecting,
		&stores.BranchSelectOperationData{Op: models.BranchOpClickToggle},
	)
	<-done

	cb.AssertAppSuccess()

}

func TestBranchReconcile(t *testing.T) {
	testBranchReconcile(t, stores.BranchClose)
	testBranchReconcile(t, stores.BranchOpened)
	testBranchReconcile(t, stores.BranchLoaded)
}

func testBranchReconcile(t *testing.T, notif flux.Notif) {
	cb := ctests.New(t).SetApp(true, false, false)
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))

	cb.ExpectReconcile(&b.Composite)
	cb.ExpectNoneDispatched()

	done := cb.GetApp().Notifier.Notify(b.model, notif)
	<-done

	cb.AssertAppSuccess()

}

func TestBranchView_Unmount(t *testing.T) {
	cb := ctests.New(t).SetApp(true, true, false)
	defer cb.Finish()

	b := NewBranchView(cb.Ctx(), models.NewBranchModel(cb.Ctx(), &models.RootContents{Name: "a"}))

	c := mock_vecty.NewMockComponent(cb.GetMock())
	c.EXPECT().Unmount()
	b.Body = c
	b.Unmount()
	assert.Nil(t, b.Notifs)
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

func setupForFailedFileLoad(cb *ctests.ClientContextBuilder) {

	// We have to unmarshal the received object, so we'll have to load some types.
	cb.Base.Path("")

	done := make(chan error, 1)
	done <- errors.New("error")

	cb.GetConnection().EXPECT().Go(shared.Data, gomock.Any(), gomock.Any(), gomock.Any()).Return(done)

}

func setupForSuccessfulFileLoad(t *testing.T, cb *ctests.ClientContextBuilder) {

	// We have to unmarshal the received object, so we'll have to load some types.
	cb.Base.Path("").Jauto().Sauto(parser.Parse)

	done := make(chan error, 1)

	do := func(method shared.Method, args interface{}, reply interface{}, fail chan error) chan error {
		// Create a simple ke object and marshal it to a []byte
		var bytes []byte
		bytes, err := ke.MarshalContext(cb.Ctx(), &system.Object{Type: system.NewReference("system", "object")})
		assert.NoError(t, err)
		reply.(*shared.DataResponse).Found = true
		reply.(*shared.DataResponse).Data = bytes
		close(done)
		return done
	}

	cb.GetConnection().EXPECT().Go(shared.Data, gomock.Any(), gomock.Any(), gomock.Any()).Do(do).Return(done)
}
