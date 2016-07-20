package editors

import (
	"fmt"

	"strconv"

	"time"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/editor/client/views"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
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
	input  *vecty.Element
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

func (v *NumberEditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*NumberEditorView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *NumberEditorView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
	if notif.Type == stores.NodeFocus {
		v.Focus()
	}
}

func (v *NumberEditorView) Focus() {
	v.input.Node().Call("focus")
}

func (v *NumberEditorView) Render() vecty.Component {
	id := randomId()

	v.input = elem.Input(
		prop.Type(prop.TypeNumber),
		prop.Value(fmt.Sprintf("%v", v.model.Node.ValueNumber)),
		prop.Class("form-control"),
		prop.ID(id),
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
				<-time.After(time.Millisecond * 50)
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

	group := elem.Div(
		vecty.ClassMap{
			"form-group": true,
			"has-error":  v.node.Invalid,
		},
		elem.Label(
			prop.For(id),
			vecty.Text(v.model.Node.Label(v.Ctx)),
		),
		v.input,
	)

	if v.format == editable.Inline {
		return group
	}

	helpBlock(v.Ctx, v.model.Node).Apply(group)
	return group
}
