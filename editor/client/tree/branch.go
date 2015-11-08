package tree

import (
	"time"

	"honnef.co/go/js/dom"
	"kego.io/editor"
	"kego.io/kerr"
)

type Branch struct {
	Tree     *Tree
	Item     Item
	Parent   *Branch
	children []*Branch
	index    int
	siblings []*Branch
	next     *Branch
	prev     *Branch
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
	dirty    map[*Branch]bool // map of dirty descendants
}

func (b *Branch) Level() int {
	return b.level
}

func (b *Branch) IsRoot() bool {
	return b.Parent == nil
}

func NewRootBranch(tree *Tree, element *dom.HTMLDivElement, inner *dom.HTMLDivElement) *Branch {
	return &Branch{
		Tree:    tree,
		open:    true,
		element: element,
		inner:   inner,
	}
}

func NewBranch(item Item, parent *Branch) *Branch {
	b := &Branch{Tree: parent.Tree, Parent: parent, Item: item}
	b.initializeElement()
	b.initializeOpener()
	b.initializeContent()
	b.initializeLabel()
	b.initializeBadge()
	b.initializeInner()
	return b
}

func (c *Branch) initializeElement() {
	element := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	element.SetAttribute("class", "node")

	c.element = element
}

func (c *Branch) initializeOpener() {
	opener := dom.GetWindow().Document().CreateElement("a").(*dom.HTMLAnchorElement)
	opener.SetAttribute("class", "toggle")
	opener.AddEventListener("click", true, func(e dom.Event) {
		if c.CanOpen() {
			c.Toggle()
		} else {
			c.Select(false)
		}
	})

	c.opener = opener
	c.element.AppendChild(opener)
}

func (c *Branch) initializeContent() {
	content := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	content.SetAttribute("class", "content")

	c.content = content
	c.element.AppendChild(content)
}

func (c *Branch) initializeLabel() {
	label := dom.GetWindow().Document().CreateElement("span").(*dom.HTMLSpanElement)
	label.SetAttribute("class", "label")
	label.AddEventListener("click", true, func(e dom.Event) {
		c.Select(false)
	})

	c.label = label
	c.content.AppendChild(label)
}

func (c *Branch) initializeBadge() {
	badge := dom.GetWindow().Document().CreateElement("span").(*dom.HTMLSpanElement)
	badge.SetAttribute("class", "badge")
	badge.SetInnerHTML(`<svg fill="#ff4081" height="12" width="12" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"/><path d="M0 0h24v24H0z" fill="none"/></svg>`)
	badge.Style().Set("display", "none")

	c.badge = badge
	c.content.AppendChild(badge)
}

func (c *Branch) initializeInner() {
	inner := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	inner.SetAttribute("class", "children")
	// children should be hidden by default
	inner.Style().Set("display", "none")

	c.inner = inner
	c.element.AppendChild(inner)
}

// lastVisible returns the last visible descendant in list order. If the branch is closed or has no
// children, we return the branch itself.
func (b *Branch) LastVisible() *Branch {
	i := b
	for i.open && len(i.children) > 0 {
		// if the node is open, test it's last child
		i = i.children[len(i.children)-1]
	}
	// return the first node we find that's closed
	return i
}

// nextVisible returns the next visible branch in list order.
func (b *Branch) NextVisible(includeChildren bool) *Branch {
	if includeChildren && b.open && len(b.children) > 0 {
		// If the branch is open and has children, we return it's first child.
		return b.children[0]
	}
	i := b
	for i.index >= len(i.siblings)-1 {
		// if the node is the last of the siblings, test it's parent
		if i.IsRoot() {
			// if we get to the root node, we're testing the last visible node, so we return nil
			return nil
		}
		i = i.Parent
	}
	// return the next sibling of the first ancestor that has one
	return i.siblings[i.index+1]
}

// prevVisible returns the previous visible branch in list order
func (c *Branch) PrevVisible() *Branch {
	if c.index == 0 {
		// at the start of a list of children, the previous in list order is always the parent
		return c.Parent
	}
	// in the middle of a list of children, the previous in list order is the lastVisible of the
	// previous sibling.
	return c.Parent.children[c.index-1].LastVisible()
}

func (c *Branch) SetLabel(text string) {
	c.label.SetTextContent(text)
}

func (c *Branch) Append(child *Branch) {
	if c.inner != nil {
		c.inner.AppendChild(child.element)
	}
	c.children = append(c.children, child)
	c.Update()
}

func (c *Branch) Select(fromKeyboard bool) {

	if c.IsRoot() {
		// we never select the root node
		return
	}

	if c.Tree.Selected != nil {
		c.Tree.Selected.Unselect()
	}
	c.content.Class().Add("selected")
	c.Tree.Selected = c
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

func (b *Branch) showEditPanel(fromKeyboard bool) {

	if b.IsRoot() {
		return
	}

	if b.editor != nil && b.Tree.Editor != nil && b.editor == b.Tree.Editor {
		return
	}

	if b.Tree.Editor != nil {
		b.Tree.Editor.Hide()
		b.Tree.Editor = nil
	}

	done, ok := b.ensureContentLoaded()

	if !ok {
		// if the operation should be cancelled, we should return immediately
		return
	}

	success := func() {

		if !fromKeyboard && b.CanOpen() && !b.open {
			// if we clicked on an item, and it's not open, we should open it
			b.Open()
		}

		if b.editor == nil {
			ed, ok := b.Item.(Editable)
			if !ok {
				return
			}
			b.editor = ed.Editor()
			if err := b.editor.Initialize(b, editor.Page, b.Tree.Path, b.Tree.Aliases); err != nil {
				b.Tree.Fail <- kerr.New("KKOBKWJDBI", err, "Initialize")
				return
			}
			b.Tree.Content.AppendChild(b.editor)
		}
		if b.editor == nil {
			return
		}
		if b.Tree.Editor != nil {
			b.Tree.Editor.Hide()
		}
		b.editor.Show()
		b.Tree.Editor = b.editor

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

func (c *Branch) Unselect() {
	// un-select
	c.content.Class().Remove("selected")
	c.Tree.Selected = nil
	c.selected = false
}

// ensureContentLoaded ensures the content is loaded. For asynchronous branches, it initializes the
// load event and returns a done chanel. For synchronous branches (or asynchronous branches that are
// already loaded), it returns nil so that the calling code can react synchronously.
func (c *Branch) ensureContentLoaded() (done chan bool, success bool) {

	if async, ok := c.Item.(Async); ok && !async.ContentLoaded() {

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
func (c *Branch) Open() {

	if c.IsRoot() {
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
func (c *Branch) Close() {
	if c.IsRoot() {
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
func (c *Branch) Toggle() {
	if c.open {
		c.Close()
	} else {
		c.Open()
	}
}

// afterStateChange is fired every time a branch is opened or closed.
func (c *Branch) afterStateChange() {
	c.Update()
	if next := c.NextVisible(false); next != nil {
		// we must also update the next in the list to ensure it's prev property is set correctly
		next.Update()
	}
	if c.Tree.Selected != nil && !c.Tree.Selected.IsVisible() {
		// if the selected branch is now invisible, we should un-select it.
		c.Tree.Selected.Unselect()
	}
	c.SetDirtyIconState()
	if c.Parent != nil {
		c.Parent.SetDirtyIconState()
	}
}

// update assumes parent and index are sources of truth, and updates siblings, prev and next, and
// updates the children
func (c *Branch) Update() {
	if c.Parent == nil {
		// special case for the root node
		c.siblings = []*Branch{c}
		c.level = 0
	} else {
		c.siblings = c.Parent.children
		c.level = c.Parent.level + 1
	}
	c.prev = c.PrevVisible()
	c.next = c.NextVisible(true)

	if c.open && len(c.children) > 0 {
		for index, child := range c.children {
			child.Parent = c
			child.index = index
			child.Update()
		}
	}
	if c.opener != nil {

		plus := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M13 7h-2v4H7v2h4v4h2v-4h4v-2h-4V7zm-1-5C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/></svg>`
		minus := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M7 11v2h10v-2H7zm5-9C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/></svg>`
		point := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8z"/><path d="M0 0h24v24H0z" fill="none"/></svg>`
		notLoaded := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M10 9c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1-.45-1-1-1zm0 4c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1-.45-1-1-1zM7 9.5c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zm3 7c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zm-3-3c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zm3-6c.28 0 .5-.22.5-.5s-.22-.5-.5-.5-.5.22-.5.5.22.5.5.5zM14 9c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1-.45-1-1-1zm0-1.5c.28 0 .5-.22.5-.5s-.22-.5-.5-.5-.5.22-.5.5.22.5.5.5zm3 6c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zm0-4c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zM12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8zm2-3.5c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zm0-3.5c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1-.45-1-1-1z"/></svg>`

		if c.isAsyncAndNotLoaded() {
			c.opener.SetInnerHTML(notLoaded)
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
func (c *Branch) isAsyncAndNotLoaded() bool {
	async, isAsync := c.Item.(Async)
	if isAsync && !async.ContentLoaded() {
		return true
	}
	return false
}

// canOpen gives us an indication of whether the branch can be opened. We use this to work out if we
// should display the plus icon, or the empty icon. If a branch is async but the children aren't
// loaded, we don't know if it has children or not. In that case we assume it can be opened.
func (c *Branch) CanOpen() bool {
	if c.isAsyncAndNotLoaded() {
		return true
	}
	if len(c.children) == 0 {
		return false
	}
	return true
}

// isDescendantOf tells us if this branch is a direct descendant of the specified branch
func (c *Branch) isDescendantOf(ancestor *Branch) bool {
	test := c.Parent
	for test != nil {
		if test == ancestor {
			return true
		}
		test = test.Parent
	}
	return false
}

// isAncestorOf tells us if this branch is a direct ancestor of the specified branch
func (c *Branch) isAncestorOf(child *Branch) bool {
	test := child.Parent
	for test != nil {
		if test == c {
			return true
		}
		test = test.Parent
	}
	return false
}

// isVisible tells us if the branch is currently visible. For a branch to be visible, all it's
// ancestors must be open.
func (c *Branch) IsVisible() bool {
	test := c.Parent
	for test != nil {
		if !test.open {
			return false
		}
		test = test.Parent
	}
	return true
}

func (c *Branch) DirtyDescendant(state bool, descendant *Branch) {
	if state {
		if c.dirty == nil {
			c.dirty = map[*Branch]bool{}
		}
		c.dirty[descendant] = true
	} else {
		if c.dirty != nil {
			delete(c.dirty, descendant)
		}
	}
	c.SetDirtyIconState()
}

func (c *Branch) SetDirtyIconState() {
	if c.badge == nil {
		return
	}

	if c.dirtyIconState() {
		c.badge.Style().Set("display", "inline")
		return
	}

	c.badge.Style().Set("display", "none")

}

func (c *Branch) dirtyIconState() bool {

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

func (c *Branch) MarkDirty(state bool) {
	ancestor := c
	for ancestor != nil {
		ancestor.DirtyDescendant(state, c)
		ancestor = ancestor.Parent
	}
}
