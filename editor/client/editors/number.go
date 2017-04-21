package editors

import (
	"fmt"

	"strconv"

	"time"

	"context"

	"frizz.io/editor/client/actions"
	"frizz.io/editor/client/common"
	"frizz.io/editor/client/editable"
	"frizz.io/editor/client/models"
	"frizz.io/editor/client/stores"
	"frizz.io/editor/client/views"
	"frizz.io/flux"
	"frizz.io/system"
	"frizz.io/system/node"
	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/event"
	"github.com/dave/vecty/prop"
)

var _ editable.Editable = (*NumberEditor)(nil)

type NumberEditor struct{}

func (s *NumberEditor) Format(rule *system.RuleWrapper) editable.Format {
	return editable.Inline
}

func (s *NumberEditor) EditorView(ctx context.Context, node *node.Node, format editable.Format) vecty.Component {
	return NewNumberEditorView(ctx, node, format)
}

type NumberEditorView struct {
	*views.View

	model  *models.EditorModel
	node   *models.NodeModel
	input  *vecty.HTML
	format editable.Format
}

func NewNumberEditorView(ctx context.Context, node *node.Node, format editable.Format) *NumberEditorView {
	v := &NumberEditorView{}
	v.View = views.New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.node = v.App.Nodes.Get(node)
	v.format = format
	v.Watch(v.model.Node,
		stores.NodeFocus,
		stores.NodeValueChanged,
		stores.NodeErrorsChanged,
	)
	return v
}

func (v *NumberEditorView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	vecty.Rerender(v)
	if notif.Type == stores.NodeFocus {
		v.Focus()
	}
}

func (v *NumberEditorView) Focus() {
	v.input.Node.Call("focus")
}

func (v *NumberEditorView) Render() *vecty.HTML {

	v.input = elem.Input(
		prop.Type(prop.TypeNumber),
		prop.Value(fmt.Sprintf("%v", v.model.Node.ValueNumber)),
		prop.Class("form-control"),
		event.KeyUp(func(e *vecty.Event) {
			getVal := func() interface{} {
				val, err := strconv.ParseFloat(e.Target.Get("value").String(), 64)
				if err != nil {
					// if there's an error converting to a float, ignore it
					return nil
				}
				return val
			}
			val := getVal()
			changed := func() bool {
				return val != getVal()
			}
			go func() {
				<-time.After(common.EditorKeyboardDebounceShort)
				if changed() {
					return
				}
				v.App.Dispatch(&actions.Modify{
					Undoer:  &actions.Undoer{},
					Editor:  v.model,
					Before:  v.model.Node.NativeValue(),
					After:   val,
					Changed: changed,
				})
			}()
		}),
	)

	return elem.Div(v.input)
}
