package mdl

import "honnef.co/go/js/dom"

type anchor struct {
	*dom.HTMLAnchorElement
}

func Anchor() *anchor {
	a := dom.GetWindow().Document().CreateElement("a").(*dom.HTMLAnchorElement)
	a.Href = ""
	return &anchor{a}
}

func (a *anchor) Text(text string) *anchor {
	a.SetTextContent(text)
	return a
}

func (a *anchor) Url(href string) *anchor {
	a.Href = href
	return a
}

func (a *anchor) Click(action func(dom.Event)) *anchor {
	a.AddEventListener("click", true, action)
	return a
}
