package branch

import (
	"time"

	"honnef.co/go/js/dom"
	"kego.io/editor"
	"kego.io/editor/client/tree"
	"kego.io/kerr"
)

type Common struct {
	this     tree.Branch
	tree     *tree.Tree
	root     bool
	parent   tree.Branch
	children []tree.Branch
	index    int
	siblings []tree.Branch
	next     tree.Branch
	prev     tree.Branch
	open     bool
	loading  bool
	level    int
	element  *dom.HTMLDivElement
	opener   *dom.HTMLAnchorElement
	inner    *dom.HTMLDivElement
	content  *dom.HTMLDivElement
	label    *dom.HTMLSpanElement
	badge    *dom.HTMLSpanElement
	selected bool
	editor   editor.Editor
	dirty    map[tree.Branch]bool // map of dirty descendants
}

func (c *Common) Initialise(parent tree.Branch) {
	c.initializeElement()
	c.initializeOpener()
	c.initializeContent()
	c.initializeLabel()
	c.initializeBadge()
	c.initializeInner()
	c.parent = parent
}

func (c *Common) initializeElement() {
	element := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	element.SetAttribute("class", "node")

	c.element = element
}

func (c *Common) initializeOpener() {
	opener := dom.GetWindow().Document().CreateElement("a").(*dom.HTMLAnchorElement)
	opener.SetAttribute("class", "toggle")
	opener.AddEventListener("click", true, func(e dom.Event) {
		if c.CanOpen() {
			c.toggle()
		} else {
			c.Select(false)
		}
	})

	c.opener = opener
	c.element.AppendChild(opener)
}

func (c *Common) initializeContent() {
	content := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	content.SetAttribute("class", "content")

	c.content = content
	c.element.AppendChild(content)
}

func (c *Common) initializeLabel() {
	label := dom.GetWindow().Document().CreateElement("span").(*dom.HTMLSpanElement)
	label.SetAttribute("class", "label")
	label.AddEventListener("click", true, func(e dom.Event) {
		c.Select(false)
	})

	c.label = label
	c.content.AppendChild(label)
}

func (c *Common) initializeBadge() {
	badge := dom.GetWindow().Document().CreateElement("span").(*dom.HTMLSpanElement)
	badge.SetAttribute("class", "badge")
	badge.SetInnerHTML(`<svg fill="#ff4081" height="12" width="12" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"/><path d="M0 0h24v24H0z" fill="none"/></svg>`)
	badge.Style().Set("display", "none")

	c.badge = badge
	c.content.AppendChild(badge)
}

func (c *Common) initializeInner() {
	inner := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	inner.SetAttribute("class", "children")
	// children should be hidden by default
	inner.Style().Set("display", "none")

	c.inner = inner
	c.element.AppendChild(inner)
}

// lastVisible returns the last visible descendant in list order. If the branch is closed or has no
// children, we return the branch itself.
func (c *Common) LastVisible() tree.Branch {
	i := c.this
	for i.IsOpen() && len(i.Children()) > 0 {
		// if the node is open, test it's last child
		i = i.Child(len(i.Children()) - 1)
	}
	// return the first node we find that's closed
	return i
}

// nextVisible returns the next visible branch in list order.
func (c *Common) NextVisible(includeChildren bool) tree.Branch {
	if includeChildren && c.open && len(c.children) > 0 {
		// If the branch is open and has children, we return it's first child.
		return c.children[0]
	}
	i := c.this
	for i.Index() >= len(i.Siblings())-1 {
		// if the node is the last of the siblings, test it's parent
		if i.Root() {
			// if we get to the root node, we're testing the last visible node, so we return nil
			return nil
		}
		i = i.Parent()
	}
	// return the next sibling of the first ancestor that has one
	return i.Sibling(i.Index() + 1)
}

// prevVisible returns the previous visible branch in list order
func (c *Common) PrevVisible() tree.Branch {
	if c.index == 0 {
		// at the start of a list of children, the previous in list order is always the parent
		return c.parent
	}
	// in the middle of a list of children, the previous in list order is the lastVisible of the
	// previous sibling.
	return c.parent.Child(c.index - 1).LastVisible()
}

func (c *Common) SetLabel(text string) {
	c.label.SetTextContent(text)
}

func (c *Common) Element() *dom.HTMLDivElement {
	return c.element
}

func (c *Common) Tree() *tree.Tree {
	return c.tree
}

func (c *Common) SetTree(t *tree.Tree) {
	c.tree = t
}

func (c *Common) Parent() tree.Branch {
	return c.parent
}

func (c *Common) Root() bool {
	return c.root
}

func (c *Common) IsOpen() bool {
	return c.open
}

func (c *Common) Children() []tree.Branch {
	return c.children
}

func (c *Common) Child(i int) tree.Branch {
	return c.children[i]
}

func (c *Common) Siblings() []tree.Branch {
	return c.siblings
}

func (c *Common) Sibling(i int) tree.Branch {
	return c.siblings[i]
}

func (c *Common) Index() int {
	return c.index
}

func (c *Common) Level() int {
	return c.level
}

func (c *Common) SetParent(parent tree.Branch) {
	c.parent = parent
}

func (c *Common) SetIndex(index int) {
	c.index = index
}

func (c *Common) Append(child tree.Branch) {
	if c.inner != nil {
		c.inner.AppendChild(child.Element())
	}
	c.children = append(c.children, child)
	c.Update()
}

func (c *Common) Select(fromKeyboard bool) {

	if c.root {
		// we never select the root node
		return
	}

	if c.tree.Selected != nil {
		c.tree.Selected.Unselect()
	}
	c.content.Class().Add("selected")
	c.tree.Selected = c
	c.selected = true

	if fromKeyboard && c.isAsyncAndNotLoaded() {
		// wait 50ms before showing the edit panel so we don't generate
		// lots of content load requests if we scroll quickly. We only
		// do this for keyboard events.
		go func() {
			time.Sleep(time.Millisecond * 50)
			if c.selected {
				c.showEditPanel(fromKeyboard)
			}
		}()
	} else {
		c.showEditPanel(fromKeyboard)
	}

	return
}

func (c *Common) showEditPanel(fromKeyboard bool) {

	if c.root {
		return
	}

	if c.editor != nil && c.tree.Editor != nil && c.editor == c.tree.Editor {
		return
	}

	if c.tree.Editor != nil {
		c.tree.Editor.Hide()
		c.tree.Editor = nil
	}

	done, ok := c.ensureContentLoaded()

	if !ok {
		// if the operation should be cancelled, we should return immediately
		return
	}

	success := func() {

		if !fromKeyboard && c.CanOpen() && !c.open {
			// if we clicked on an item, and it's not open, we should open it
			c.Open()
		}

		if c.editor == nil {
			ed, ok := c.this.(tree.Editable)
			if !ok {
				return
			}
			c.editor = ed.Editor()
			if err := c.editor.Initialize(c.tree.Content, c, c.tree.Path, c.tree.Aliases); err != nil {
				c.tree.Fail <- kerr.New("KKOBKWJDBI", err, "Initialize")
				return
			}
		}
		if c.editor == nil {
			return
		}
		if c.tree.Editor != nil {
			c.tree.Editor.Hide()
		}
		c.editor.Show()
		c.tree.Editor = c.editor

	}

	if done == nil {
		// if the done chanel is nil, the operation was synchronous, so we should call success synchronously
		success()
	} else {
		go func() {
			// block and wait until the response arrives
			<-done
			success()
		}()
	}
	return
}

func (c *Common) Unselect() {
	// un-select
	c.content.Class().Remove("selected")
	c.tree.Selected = nil
	c.selected = false
}

// ensureContentLoaded ensures the content is loaded. For asynchronous branches, it initializes the
// load event and returns a done chanel. For synchronous branches (or asynchronous branches that are
// already loaded), it returns nil so that the calling code can react synchronously.
func (c *Common) ensureContentLoaded() (done chan bool, success bool) {

	if async, ok := c.this.(tree.Async); ok && !async.ContentLoaded() {

		done = make(chan bool, 1)

		if c.loading {
			// if we're already in the process of loading the contents, we should cancel the
			// operation
			return nil, false
		}

		// load content asynchronously
		c.loading = true
		responseChannel := async.LoadContent()

		go func() {
			<-responseChannel
			c.loading = false
			done <- true
		}()

		return done, true

	} else {
		// if item is not async or content is already loaded, just open the node.
		return nil, true
	}
}

// Open opens the branch. For asynchronously loaded branches it initialises the load.
func (c *Common) Open() {

	if c.root {
		return
	}

	done, ok := c.ensureContentLoaded()

	if !ok {
		// if the operation should be cancelled, we should return immediately
		return
	}

	success := func() {
		if c.inner != nil {
			c.inner.Style().Set("display", "block")
		}
		c.open = true
		c.afterStateChange()
	}

	if done == nil {
		// if the done chanel is nil, the operation was synchronous, so we should call success
		// synchronously
		success()
	} else {
		go func() {
			// block and wait until the response arrives
			<-done
			success()
		}()
	}
	return
}

// Close closes a branch and hides the children
func (c *Common) Close() {
	if c.root {
		return
	}
	if c.inner != nil {
		c.inner.Style().Set("display", "none")
	}
	c.open = false
	c.loading = false
	c.afterStateChange()
	return
}

// toggle inverts the open/closed state of a branch
func (c *Common) toggle() {
	if c.open {
		c.Close()
	} else {
		c.Open()
	}
}

// afterStateChange is fired every time a branch is opened or closed.
func (c *Common) afterStateChange() {
	c.Update()
	if next := c.NextVisible(false); next != nil {
		// we must also update the next in the list to ensure it's prev property is set correctly
		next.Update()
	}
	if c.tree.Selected != nil && !c.tree.Selected.IsVisible() {
		// if the selected branch is now invisible, we should un-select it.
		c.tree.Selected.Unselect()
	}
	c.SetDirtyIconState()
	if c.parent != nil {
		c.parent.SetDirtyIconState()
	}
}

// update assumes parent and index are sources of truth, and updates siblings, prev and next, and
// updates the children
func (c *Common) Update() {
	if c.parent == nil {
		// special case for the root node
		c.siblings = []tree.Branch{c.this}
		c.level = 0
	} else {
		c.siblings = c.parent.Children()
		c.level = c.parent.Level() + 1
	}
	c.prev = c.PrevVisible()
	c.next = c.NextVisible(true)

	if c.open && len(c.children) > 0 {
		for index, child := range c.children {
			child.SetParent(c.this)
			child.SetIndex(index)
			child.Update()
		}
	}
	if c.opener != nil {

		plus := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M13 7h-2v4H7v2h4v4h2v-4h4v-2h-4V7zm-1-5C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/></svg>`
		minus := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M7 11v2h10v-2H7zm5-9C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/></svg>`
		point := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8z"/><path d="M0 0h24v24H0z" fill="none"/></svg>`

		if c.isAsyncAndNotLoaded() {
			c.opener.SetInnerHTML(plus)
		} else if c.children == nil || len(c.children) == 0 {
			c.opener.SetInnerHTML(point)
		} else if c.open {
			c.opener.SetInnerHTML(minus)
		} else {
			c.opener.SetInnerHTML(plus)
		}
	}
}

// isAsyncAndNotLoaded is true if the branch loads it's contents asynchronously, and the contents
// are not yet loaded
func (c *Common) isAsyncAndNotLoaded() bool {
	async, isAsync := c.this.(tree.Async)
	if isAsync && !async.ContentLoaded() {
		return true
	}
	return false
}

// canOpen gives us an indication of whether the branch can be opened. We use this to work out if we
// should display the plus icon, or the empty icon. If a branch is async but the children aren't
// loaded, we don't know if it has children or not. In that case we assume it can be opened.
func (c *Common) CanOpen() bool {
	if c.isAsyncAndNotLoaded() {
		return true
	}
	if len(c.children) == 0 {
		return false
	}
	return true
}

// isDescendantOf tells us if this branch is a direct descendant of the specified branch
func (c *Common) isDescendantOf(ancestor tree.Branch) bool {
	test := c.parent
	for test != nil {
		if test == ancestor {
			return true
		}
		test = test.Parent()
	}
	return false
}

// isAncestorOf tells us if this branch is a direct ancestor of the specified branch
func (c *Common) isAncestorOf(child tree.Branch) bool {
	test := child.Parent()
	for test != nil {
		if test == c {
			return true
		}
		test = test.Parent()
	}
	return false
}

// isVisible tells us if the branch is currently visible. For a branch to be visible, all it's
// ancestors must be open.
func (c *Common) IsVisible() bool {
	test := c.parent
	for test != nil {
		if !test.IsOpen() {
			return false
		}
		test = test.Parent()
	}
	return true
}

func (c *Common) DirtyDescendant(state bool, descendant tree.Branch) {
	if state {
		if c.dirty == nil {
			c.dirty = map[tree.Branch]bool{}
		}
		c.dirty[descendant] = true
	} else {
		if c.dirty != nil {
			delete(c.dirty, descendant)
		}
	}
	c.SetDirtyIconState()
}

func (c *Common) SetDirtyIconState() {
	if c.badge == nil {
		return
	}

	if c.dirtyIconState() {
		c.badge.Style().Set("display", "inline")
		return
	}

	c.badge.Style().Set("display", "none")

}

func (c *Common) dirtyIconState() bool {

	if c.dirty == nil || len(c.dirty) == 0 {
		return false
	}

	if c.dirty[c] == true {
		// if this branch is dirty, show the icon
		return true
	}

	if !c.open {
		// if descendants are dirty, only show the icon if this branch is closed
		return true
	}

	return false
}

func (c *Common) MarkDirty(state bool) {
	ancestor := c.this
	for ancestor != nil {
		ancestor.DirtyDescendant(state, c)
		ancestor = ancestor.Parent()
	}
}
