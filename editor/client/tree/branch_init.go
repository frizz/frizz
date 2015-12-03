package tree

import "honnef.co/go/js/dom"

func (b *Branch) init() {
	b.initElement()
	b.initOpener()
	b.initContent()
	b.initLabel()
	b.initBadge()
	b.initInner()
}

func (b *Branch) initElement() {
	element := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	element.SetAttribute("class", "node")

	b.element = element
}

func (b *Branch) initOpener() {
	opener := dom.GetWindow().Document().CreateElement("a").(*dom.HTMLAnchorElement)
	opener.SetAttribute("class", "toggle")
	opener.AddEventListener("click", true, func(e dom.Event) {
		if b.canOpen() {
			b.toggle()
		} else {
			b.Select(false)
		}
	})

	b.opener = opener
	b.element.AppendChild(opener)
}

func (b *Branch) initContent() {
	content := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	content.SetAttribute("class", "content")

	b.content = content
	b.element.AppendChild(content)
}

func (b *Branch) initLabel() {
	label := dom.GetWindow().Document().CreateElement("span").(*dom.HTMLSpanElement)
	label.SetAttribute("class", "label")
	label.AddEventListener("click", true, func(e dom.Event) {
		b.Select(false)
	})

	b.label = label
	b.content.AppendChild(label)
}

func (b *Branch) initBadge() {
	badge := dom.GetWindow().Document().CreateElement("span").(*dom.HTMLSpanElement)
	badge.SetAttribute("class", "badge")
	badge.SetInnerHTML(`<svg fill="#ff4081" height="12" width="12" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M0 0h24v24H0z" fill="none"/><path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"/><path d="M0 0h24v24H0z" fill="none"/></svg>`)
	badge.Style().Set("display", "none")

	b.badge = badge
	b.content.AppendChild(badge)
}

func (b *Branch) initInner() {
	inner := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	inner.SetAttribute("class", "children")
	// children should be hidden by default
	inner.Style().Set("display", "none")

	b.inner = inner
	b.element.AppendChild(inner)
}
