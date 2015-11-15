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
	dirty    map[editor.Editor]bool // descendant editors that have changes
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

func (b *Branch) initializeElement() {
	element := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	element.SetAttribute("class", "node")

	b.element = element
}

func (b *Branch) initializeOpener() {
	opener := dom.GetWindow().Document().CreateElement("a").(*dom.HTMLAnchorElement)
	opener.SetAttribute("class", "toggle")
	opener.AddEventListener("click", true, func(e dom.Event) {
		if b.CanOpen() {
			b.Toggle()
		} else {
			b.Select(false)
		}
	})

	b.opener = opener
	b.element.AppendChild(opener)
}

func (b *Branch) initializeContent() {
	content := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	content.SetAttribute("class", "content")

	b.content = content
	b.element.AppendChild(content)
}

func (b *Branch) initializeLabel() {
	label := dom.GetWindow().Document().CreateElement("span").(*dom.HTMLSpanElement)
	label.SetAttribute("class", "label")
	label.AddEventListener("click", true, func(e dom.Event) {
		b.Select(false)
	})

	b.label = label
	b.content.AppendChild(label)
}

func (b *Branch) initializeBadge() {
	badge := dom.GetWindow().Document().CreateElement("span").(*dom.HTMLSpanElement)
	badge.SetAttribute("class", "badge")
	badge.SetInnerHTML(`<svg fill="#ff4081" height="12" width="12" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"/><path d="M0 0h24v24H0z" fill="none"/></svg>`)
	badge.Style().Set("display", "none")

	b.badge = badge
	b.content.AppendChild(badge)
}

func (b *Branch) initializeInner() {
	inner := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	inner.SetAttribute("class", "children")
	// children should be hidden by default
	inner.Style().Set("display", "none")

	b.inner = inner
	b.element.AppendChild(inner)
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
func (b *Branch) PrevVisible() *Branch {
	if b.index == 0 {
		// at the start of a list of children, the previous in list order is always the parent
		return b.Parent
	}
	// in the middle of a list of children, the previous in list order is the lastVisible of the
	// previous sibling.
	return b.Parent.children[b.index-1].LastVisible()
}

func (b *Branch) SetLabel(text string) {
	b.label.SetTextContent(text)
}

func (c *Branch) Append(child *Branch) {
	if c.inner != nil {
		c.inner.AppendChild(child.element)
	}
	c.children = append(c.children, child)
	c.Update()
}

func (b *Branch) Select(fromKeyboard bool) {

	if b.IsRoot() {
		// we never select the root node
		return
	}

	if !b.IsVisible() {
		// if the branch we're selecting isn't visible, we should open all it's closed ancestors.
		ancestor := b.Parent
		for ancestor != nil {
			if !ancestor.open && ancestor.CanOpen() {
				ancestor.Open()
			}
			ancestor = ancestor.Parent
		}
	}

	if b.Tree.Selected != nil {
		b.Tree.Selected.Unselect()
	}
	b.content.Class().Add("selected")
	b.Tree.Selected = b
	b.selected = true

	if fromKeyboard && b.isAsyncAndNotLoaded() {
		// wait 50ms before showing the edit panel so we don't generate
		// lots of content load requests if we scroll quickly. We only
		// do this for keyboard events.
		go func() {
			time.Sleep(time.Millisecond * 50)
			if b.selected {
				b.showEditPanel(fromKeyboard)
			}
		}()
	} else {
		b.showEditPanel(fromKeyboard)
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
			b.Listen(b.editor.Listen().Ch)
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

func (b *Branch) Unselect() {
	// un-select
	b.content.Class().Remove("selected")
	b.Tree.Selected = nil
	b.selected = false
}

// ensureContentLoaded ensures the content is loaded. For asynchronous branches, it initializes the
// load event and returns a done chanel. For synchronous branches (or asynchronous branches that are
// already loaded), it returns nil so that the calling code can react synchronously.
func (b *Branch) ensureContentLoaded() (done chan bool, success bool) {

	if async, ok := b.Item.(Async); ok && !async.ContentLoaded() {

		done = make(chan bool, 1)

		if b.loading {
			// if we're already in the process of loading the contents, we should cancel the
			// operation
			return nil, false
		}

		// load content asynchronously
		b.loading = true
		responseChannel := async.LoadContent()

		go func() {
			<-responseChannel
			b.loading = false
			done <- true
		}()

		return done, true

	} else {
		// if item is not async or content is already loaded, just open the node.
		return nil, true
	}
}

// Open opens the branch. For asynchronously loaded branches it initialises the load.
func (b *Branch) Open() {

	if b.IsRoot() {
		return
	}

	done, ok := b.ensureContentLoaded()

	if !ok {
		// if the operation should be cancelled, we should return immediately
		return
	}

	success := func(fromAsync bool) {

		if b.inner != nil {
			b.inner.Style().Set("display", "block")
		}

		if fromAsync && !b.CanOpen() {
			// if after loading the content, we can't we can't open the branch, this means we just
			// there were no children. In this case we want to select the branch.
			b.Select(false)
		} else {
			b.open = true
			b.afterStateChange()
		}
	}

	if done == nil {
		// if the done chanel is nil, the operation was synchronous, so we should call success
		// synchronously
		success(false)
	} else {
		go func() {
			// block and wait until the response arrives
			<-done
			success(true)
		}()
	}
	return
}

// Close closes a branch and hides the children
func (b *Branch) Close() {
	b.close()
	b.afterStateChange()
	return
}

func (b *Branch) close() {
	if b.IsRoot() {
		return
	}
	if b.inner != nil {
		b.inner.Style().Set("display", "none")
	}
	b.open = false
	b.loading = false
	for _, child := range b.children {
		if child.open {
			child.close()
		}
	}
}

// toggle inverts the open/closed state of a branch
func (b *Branch) Toggle() {
	if b.open {
		b.Close()
	} else {
		b.Open()
	}
}

// afterStateChange is fired every time a branch is opened or closed.
func (b *Branch) afterStateChange() {
	b.Update()
	if next := b.NextVisible(false); next != nil {
		// we must also update the next in the list to ensure it's prev property is set correctly
		next.Update()
	}
	if b.Tree.Selected != nil && !b.Tree.Selected.IsVisible() {
		// if the selected branch is now invisible, we should un-select it.
		b.Tree.Selected.Unselect()
	}
	b.setDirtyIconState()
	if b.Parent != nil {
		b.Parent.setDirtyIconState()
	}
}

// update assumes parent and index are sources of truth, and updates siblings, prev and next, and
// updates the children
func (b *Branch) Update() {
	if b.Parent == nil {
		// special case for the root node
		b.siblings = []*Branch{b}
		b.level = 0
	} else {
		b.siblings = b.Parent.children
		b.level = b.Parent.level + 1
	}
	b.prev = b.PrevVisible()
	b.next = b.NextVisible(true)

	if b.open && len(b.children) > 0 {
		for index, child := range b.children {
			child.Parent = b
			child.index = index
			child.Update()
		}
	}
	if b.opener != nil {

		plus := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M13 7h-2v4H7v2h4v4h2v-4h4v-2h-4V7zm-1-5C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/></svg>`
		minus := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M7 11v2h10v-2H7zm5-9C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/></svg>`
		point := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8z"/><path d="M0 0h24v24H0z" fill="none"/></svg>`
		notLoaded := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M10 9c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1-.45-1-1-1zm0 4c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1-.45-1-1-1zM7 9.5c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zm3 7c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zm-3-3c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zm3-6c.28 0 .5-.22.5-.5s-.22-.5-.5-.5-.5.22-.5.5.22.5.5.5zM14 9c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1-.45-1-1-1zm0-1.5c.28 0 .5-.22.5-.5s-.22-.5-.5-.5-.5.22-.5.5.22.5.5.5zm3 6c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zm0-4c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zM12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8zm2-3.5c-.28 0-.5.22-.5.5s.22.5.5.5.5-.22.5-.5-.22-.5-.5-.5zm0-3.5c-.55 0-1 .45-1 1s.45 1 1 1 1-.45 1-1-.45-1-1-1z"/></svg>`

		if b.isAsyncAndNotLoaded() {
			b.opener.SetInnerHTML(notLoaded)
		} else if b.children == nil || len(b.children) == 0 {
			b.opener.SetInnerHTML(point)
		} else if b.open {
			b.opener.SetInnerHTML(minus)
		} else {
			b.opener.SetInnerHTML(plus)
		}
	}
}

// isAsyncAndNotLoaded is true if the branch loads it's contents asynchronously, and the contents
// are not yet loaded
func (b *Branch) isAsyncAndNotLoaded() bool {
	async, isAsync := b.Item.(Async)
	if isAsync && !async.ContentLoaded() {
		return true
	}
	return false
}

// canOpen gives us an indication of whether the branch can be opened. We use this to work out if we
// should display the plus icon, or the empty icon. If a branch is async but the children aren't
// loaded, we don't know if it has children or not. In that case we assume it can be opened.
func (b *Branch) CanOpen() bool {
	if b.isAsyncAndNotLoaded() {
		return true
	}
	if len(b.children) == 0 {
		return false
	}
	return true
}

// isDescendantOf tells us if this branch is a direct descendant of the specified branch
func (b *Branch) isDescendantOf(ancestor *Branch) bool {
	test := b.Parent
	for test != nil {
		if test == ancestor {
			return true
		}
		test = test.Parent
	}
	return false
}

// isAncestorOf tells us if this branch is a direct ancestor of the specified branch
func (b *Branch) isAncestorOf(child *Branch) bool {
	test := child.Parent
	for test != nil {
		if test == b {
			return true
		}
		test = test.Parent
	}
	return false
}

// isVisible tells us if the branch is currently visible. For a branch to be visible, all it's
// ancestors must be open.
func (b *Branch) IsVisible() bool {
	test := b.Parent
	for test != nil {
		if !test.open {
			return false
		}
		test = test.Parent
	}
	return true
}

func (b *Branch) markEditorDirtyState(e editor.Editor, state bool) {
	if state {
		if b.dirty == nil {
			b.dirty = map[editor.Editor]bool{}
		}
		b.dirty[e] = true
	} else {
		if b.dirty != nil {
			delete(b.dirty, e)
		}
	}
	b.setDirtyIconState()
}

func (b *Branch) setDirtyIconState() {
	if b.badge == nil {
		return
	}

	if b.dirtyIconState() {
		b.badge.Style().Set("display", "inline")
		return
	}

	b.badge.Style().Set("display", "none")

}

func (b *Branch) dirtyIconState() bool {

	if b.dirty == nil || len(b.dirty) == 0 {
		return false
	}

	if !b.open {
		// if descendants are dirty, only show the icon if this branch is closed
		return true
	}

	return false
}

func (b *Branch) notify(editor editor.Editor) {
	ancestor := b
	for ancestor != nil {
		ancestor.markEditorDirtyState(editor, editor.Dirty())
		ancestor = ancestor.Parent
	}
}

func (b *Branch) Listen(changes <-chan interface{}) {
	go func() {
		for e := range changes {
			b.notify(e.(editor.Editor))
		}
	}()
}
