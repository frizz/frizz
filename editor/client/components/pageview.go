package components // import "kego.io/editor/client/components"

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"kego.io/context/envctx"
	"kego.io/editor"
)

type PageView struct {
	vecty.Composite

	Environment *envctx.Env
	Package     *editor.Node
	Heading     string
}

// Apply implements the vecty.Markup interface.
func (p *PageView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *PageView) Reconcile(oldComp vecty.Component) {
	if oldComp, ok := oldComp.(*PageView); ok {
		p.Body = oldComp.Body
		p.Heading = oldComp.Heading
	}
	p.RenderFunc = p.render
	p.ReconcileBody()
}

func (p *PageView) render() vecty.Component {
	return elem.Div(
		prop.ID("wrapper"),
		elem.Navigation(
			prop.Class("navbar navbar-inverse navbar-fixed-top"),
			elem.Div(
				prop.Class("container-fluid"),
				elem.Div(
					prop.Class("navbar-header"),
					elem.Button(
						prop.Type("button"),
						prop.Class("navbar-toggle collapsed"),
						vecty.Data("toggle", "collapse"),
						vecty.Data("target", "#navbar"),
						elem.Span(
							prop.Class("sr-only"),
							vecty.Text("Toggle navigation"),
						),
						elem.Span(
							prop.Class("icon-bar"),
						),
						elem.Span(
							prop.Class("icon-bar"),
						),
						elem.Span(
							prop.Class("icon-bar"),
						),
					),
					elem.Anchor(
						prop.Class("navbar-brand"),
						prop.Href("#"),
						vecty.Text(p.Environment.Path),
					),
				),
				elem.Div(
					prop.ID("navbar"),
					prop.Class("navbar-collapse collapse"),
					elem.UnorderedList(
						prop.Class("nav navbar-nav"),
						elem.ListItem(
							prop.Class("active"),
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Home"),
							),
						),
						elem.ListItem(
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("About"),
							),
						),
						elem.ListItem(
							prop.Class("dropdown"),
							elem.Anchor(
								prop.Href("#"),
								prop.Class("dropdown-toggle"),
								vecty.Data("toggle", "dropdown"),
								vecty.Text("Dropdown"),
								elem.Span(
									prop.Class("caret"),
								),
							),
							elem.UnorderedList(
								prop.Class("dropdown-menu"),
								elem.ListItem(
									elem.Anchor(
										prop.Href("#"),
										vecty.Text("Action"),
									),
								),
								elem.ListItem(
									elem.Anchor(
										prop.Href("#"),
										vecty.Text("Another action"),
									),
								),
								elem.ListItem(
									prop.Class("divider"),
								),
								elem.ListItem(
									prop.Class("dropdown-header"),
									vecty.Text("Nav header"),
								),
								elem.ListItem(
									elem.Anchor(
										prop.Href("#"),
										vecty.Text("Separated"),
									),
								),
							),
						),
					),
					elem.Form(
						prop.Class("navbar-form navbar-right"),
						elem.Input(
							prop.Type("text"),
							prop.Class("form-control"),
							prop.Placeholder("Search..."),
						),
					),
				),
			),
			/*
				<form class="navbar-form navbar-right">
				            <input type="text" class="form-control" placeholder="Search...">
				          </form>
							         <ul class="nav navbar-nav navbar-right">
							           <li><a href="../navbar/">Default</a></li>
							           <li><a href="../navbar-static-top/">Static top</a></li>
							           <li class="active"><a href="./">Fixed top <span class="sr-only">(current)</span></a></li>
							         </ul>
							       </div><!--/.nav-collapse -->
							     </div>
							   </nav>
			*/
		),
		elem.Div(
			prop.Class("wrapper"),
			elem.Div(
				prop.ID("tree"),
				prop.Class("split split-horizontal"),
				elem.Div(
					prop.Class("content"),
					vecty.Text("Sidebar"),
				),
			),
			elem.Div(
				prop.ID("main"),
				prop.Class("split split-horizontal"),
				elem.Div(
					prop.Class("content"),
					vecty.Text("Content"),
				),
			),
		),
	)
}
