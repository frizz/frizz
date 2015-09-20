package tree

import (
	"honnef.co/go/js/dom"
	"kego.io/kerr"
)

type Node struct {
	root     bool
	item     Item
	parent   *Node
	children []*Node
	index    int
	siblings []*Node
	next     *Node
	prev     *Node
	open     bool
	opening  bool
	level    int
	element  *dom.HTMLDivElement
	opener   *dom.HTMLAnchorElement
	inner    *dom.HTMLDivElement
	content  *dom.HTMLDivElement
}

func New(parent dom.Element) *Node {
	node := &Node{root: true, open: true, item: &root{}}
	// We must tolerate passing in a nil dom element in order to run tests in pure go
	if parent != nil {
		nodeDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
		nodeDiv.SetAttribute("class", "node root")
		parent.AppendChild(nodeDiv)

		contentDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
		contentDiv.SetAttribute("class", "content")
		nodeDiv.AppendChild(contentDiv)

		innerDiv := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
		innerDiv.SetAttribute("class", "children")
		nodeDiv.AppendChild(innerDiv)

		node.element = nodeDiv
		node.inner = innerDiv
		node.content = contentDiv
	}
	node.Initialise()
	return node
}
func (n *Node) Item() Item {
	return n.item
}

func (n *Node) Initialise() {
	// We must tolerate having a nil dom element in order to run tests in pure go
	if n.element != nil {
		n.item.Initialise(n.content)
	}
}

func NewNode(item Item) *Node {
	return &Node{item: item}
}

func (n *Node) AppendItem(item Item) *Node {
	return n.AppendNodes(NewNode(item))
}

func (n *Node) AppendNodes(nodes ...*Node) *Node {
	for _, child := range nodes {
		child.parent = n
		// We must tolerate passing in a nil dom element in order to run tests in pure go
		if n.element != nil {

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

			n.inner.AppendChild(nodeDiv)
		}
		n.children = append(n.children, child)
		child.Initialise()
	}
	n.update()
	return nodes[0]
}

func (n *Node) Each(f func(*Node) error) error {
	c := n
	for c != nil {
		if err := f(c); err != nil {
			return kerr.New("MPQEUJHVPN", err, "f")
		}
		c = c.next
	}
	return nil
}

func (n *Node) Open() *Node {
	if n.root {
		return n
	}

	if n.opening {
		// if we're already in the process of opening the node, we should cancel the open operation
		return n
	}

	opened := func() {
		if n.inner != nil {
			n.inner.Style().Set("display", "block")
		}
		n.open = true
		n.opening = false
		n.afterStateChange()
	}

	if async, ok := n.item.(AsyncItem); ok && !async.ContentLoaded() {

		// load content asynchronously
		n.opening = true
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

	return n
}

func (n *Node) Close() *Node {
	if n.root {
		return n
	}
	if n.inner != nil {
		n.inner.Style().Set("display", "none")
	}
	n.open = false
	n.opening = false
	n.afterStateChange()
	return n
}

func (n *Node) Toggle() *Node {
	if n.open {
		return n.Close()
	} else {
		return n.Open()
	}
}

func (n *Node) afterStateChange() {
	n.update()
	if next := n.nextVisible(); next != nil {
		// we must also update the next in the list to ensure it's prev
		// property is set correctly
		next.update()
	}
}

// update assumes parent and index are sources of truth, and updates
// siblings, prev and next, and updates the children
func (n *Node) update() {
	if n.parent == nil {
		// special case for the root node
		n.siblings = []*Node{n}
		n.level = 0
	} else {
		n.siblings = n.parent.children
		n.level = n.parent.level + 1
	}
	if n.index == 0 {
		// at the start of a list of children, the previous in list
		// order is always the parent
		n.prev = n.parent
	} else {
		// in the middle of a list of children, the previous in list
		// order is the lastVisible of the previous sibling.
		n.prev = n.parent.children[n.index-1].lastVisible()
	}
	if n.open && len(n.children) > 0 {
		// if the node is open, the next in list order will be the
		// first child
		n.next = n.children[0]
	} else {
		// if it's closed, we use the nextVisible method to get the
		// next sibling, or if we're the last sibling, the next sibling
		// of the first ancestor that has one.
		n.next = n.nextVisible()
	}
	if n.open && len(n.children) > 0 {
		for index, child := range n.children {
			child.parent = n
			child.index = index
			child.update()
		}
	}
	if n.opener != nil {

		plus := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M13 7h-2v4H7v2h4v4h2v-4h4v-2h-4V7zm-1-5C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/></svg>`
		minus := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M7 11v2h10v-2H7zm5-9C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z"/></svg>`
		point := `<svg fill="#000000" height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8z"/><path d="M0 0h24v24H0z" fill="none"/></svg>`

		async, isAsync := n.item.(AsyncItem)
		if isAsync && !async.ContentLoaded() {
			n.opener.SetInnerHTML(plus)
		} else if n.children == nil || len(n.children) == 0 {
			n.opener.SetInnerHTML(point)
		} else if n.open {
			n.opener.SetInnerHTML(minus)
		} else {
			n.opener.SetInnerHTML(plus)
		}
	}
}

// lastVisible returns the last visible descendant in list order
func (n *Node) lastVisible() *Node {
	i := n
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
func (n *Node) nextVisible() *Node {
	i := n
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
