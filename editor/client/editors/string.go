package editors

import (
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

var _ editable.Editable = (*StringEditor)(nil)

type StringEditor struct{}

func (s *StringEditor) Format(rule *system.RuleWrapper) editable.Format {
	if rule != nil {
		if r, ok := rule.Interface.(*system.StringRule); ok && r.Long {
			return editable.Block
		}
	}
	return editable.Inline
}

func (s *StringEditor) EditorView(ctx context.Context, node *node.Node, format editable.Format) vecty.Component {
	return NewStringEditorView(ctx, node, format)
}

type StringEditorView struct {
	*views.View

	model  *models.EditorModel
	node   *models.NodeModel
	input  *vecty.HTML
	format editable.Format
}

func NewStringEditorView(ctx context.Context, node *node.Node, format editable.Format) *StringEditorView {
	v := &StringEditorView{}
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

func (v *StringEditorView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	vecty.Rerender(v)
	if notif.Type == stores.NodeFocus {
		v.Focus()
	}
}

func (v *StringEditorView) Focus() {
	v.input.Node.Call("focus")
}

func (v *StringEditorView) Render() *vecty.HTML {

	contents := vecty.List{
		prop.Value(v.model.Node.ValueString),
		prop.Class("form-control"),
		event.KeyUp(func(e *vecty.Event) {
			getVal := func() interface{} {
				return e.Target.Get("value").String()
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
	}

	if sr, ok := v.model.Node.Rule.Interface.(*system.StringRule); ok && sr.Long {
		v.input = elem.TextArea(
			contents,
		)
	} else {
		v.input = elem.Input(
			prop.Type(prop.TypeText),
			contents,
		)
	}

	return elem.Div(v.input)
}
