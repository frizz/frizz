package tree // import "kego.io/editor/client/tree"

import (
	"honnef.co/go/js/dom"
	"kego.io/kerr"
)

type Branch struct {
	Tree     *Tree
	root     bool
	item     Item
	parent   *Branch
	children []*Branch
	index    int
	siblings []*Branch
	next     *Branch
	prev     *Branch
	open     bool
	opening  bool
	level    int
	element  *dom.HTMLDivElement
	opener   *dom.HTMLAnchorElement
	inner    *dom.HTMLDivElement
	content  *dom.HTMLDivElement
}

func (b *Branch) Initialise() {
	// We must tolerate having a nil dom element in order to run tests in pure go
	if b.element != nil {
		b.item.Initialise(b.content)
	}
}

func (b *Branch) Append(child *Branch) *Branch {
	child.parent = b
	// We must tolerate passing in a nil dom element in order to run tests in pure go
	if b.element != nil {

		opener := dom.GetWindow().Document().CreateElement("a").(*dom.HTMLAnchorElement)
		opener.SetAttribute("class", "toggle")
		opener.AddEventListener("click", true, func(e dom.Event) {
			child.Toggle()
		})

		contentDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
		contentDiv.SetAttribute("class", "content")

		innerDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
		innerDiv.SetAttribute("class", "children")

		nodeDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
		nodeDiv.SetAttribute("class", "node")

		nodeDiv.AppendChild(opener)
		nodeDiv.AppendChild(contentDiv)
		nodeDiv.AppendChild(innerDiv)

		child.opener = opener
		child.element = nodeDiv
		child.inner = innerDiv
		child.content = contentDiv

		b.inner.AppendChild(nodeDiv)
	}
	b.children = append(b.children, child)
	child.Initialise()
	b.update()
	return child
}

func (b *Branch) Each(f func(*Branch) error) error {
	c := b
	for c != nil {
		if err := f(c); err != nil {
			return kerr.New("MPQEUJHVPN", err, "f")
		}
		c = c.next
	}
	return nil
}

func (b *Branch) Open() *Branch {
	if b.root {
		return b
	}

	if b.opening {
		// if we're already in the process of opening the node, we should cancel the open operation
		return b
	}

	opened := func() {
		if b.inner != nil {
			b.inner.Style().Set("display", "block")
		}
		b.open = true
		b.opening = false
		b.afterStateChange()
	}

	if async, ok := b.item.(AsyncItem); ok && !async.ContentLoaded() {

		// load content asynchronously
		b.opening = true
		successChannel := async.LoadContent()
		go func() {
			<-successChannel
			opened()
		}()

	} else {

		// if item is not async or content is already loaded, just
		// open the node.
		opened()
	}

	return b
}

func (b *Branch) Close() *Branch {
	if b.root {
		return b
	}
	if b.inner != nil {
		b.inner.Style().Set("display", "none")
	}
	b.open = false
	b.opening = false
	b.afterStateChange()
	return b
}

func (b *Branch) Toggle() *Branch {
	if b.open {
		return b.Close()
	} else {
		return b.Open()
	}
}

func (b *Branch) afterStateChange() {
	b.update()
	if next := b.nextVisible(); next != nil {
		// we must also update the next in the list to ensure it's prev
		// property is set correctly
		next.update()
	}
}

// update assumes parent and index are sources of truth, and updates
// siblings, prev and next, and updates the children
func (b *Branch) update() {
	if b.parent == nil {
		// special case for the root node
		b.siblings = []*Branch{b}
		b.level = 0
	} else {
		b.siblings = b.parent.children
		b.level = b.parent.level + 1
	}
	if b.index == 0 {
		// at the start of a list of children, the previous in list
		// order is always the parent
		b.prev = b.parent
	} else {
		// in the middle of a list of children, the previous in list
		// order is the lastVisible of the previous sibling.
		b.prev = b.parent.children[b.index-1].lastVisible()
	}
	if b.open && len(b.children) > 0 {
		// if the node is open, the next in list order will be the
		// first child
		b.next = b.children[0]
	} else {
		// if it's closed, we use the nextVisible method to get the
		// next sibling, or if we're the last sibling, the next sibling
		// of the first ancestor that has one.
		b.next = b.nextVisible()
	}
	if b.open && len(b.children) > 0 {
		for index, child := range b.children {
			child.parent = b
			child.index = index
			child.update()
		}
	}
	if b.opener != nil {

		plus := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M13 7h-2v4H7v2h4v4h2v-4h4v-2h-4V7zm-1-5C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/></svg>`
		minus := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M7 11v2h10v-2H7zm5-9C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/></svg>`
		point := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8z"/><path d="M0 0h24v24H0z" fill="none"/></svg>`

		async, isAsync := b.item.(AsyncItem)
		if isAsync && !async.ContentLoaded() {
			b.opener.SetInnerHTML(plus)
		} else if b.children == nil || len(b.children) == 0 {
			b.opener.SetInnerHTML(point)
		} else if b.open {
			b.opener.SetInnerHTML(minus)
		} else {
			b.opener.SetInnerHTML(plus)
		}
	}
}

// lastVisible returns the last visible descendant in list order
func (b *Branch) lastVisible() *Branch {
	i := b
	for i.open && len(i.children) > 0 {
		// if the node is open, test it's last child
		i = i.children[len(i.children)-1]
	}
	// return the first node we find that's closed
	return i
}

// nextVisible returns the next sibling. If we're the last sibling,
// we find the nearest ancestor that has a next sibling, or nil if
// we're at the end of the tree.
func (b *Branch) nextVisible() *Branch {
	i := b
	for i.index >= len(i.siblings)-1 {
		// if the node is the last of the siblings,
		// test it's parent
		if i.root {
			// if we get to the root node, we're testing the last visible
			// node, so we return nil
			return nil
		}
		i = i.parent
	}
	// return the next sibling of the first ancestor that has one
	return i.siblings[i.index+1]
}
