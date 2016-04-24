package mdl

import "honnef.co/go/js/dom"

type AnchorStruct struct {
	*dom.HTMLAnchorElement
}

func Anchor() *AnchorStruct {
	a := get("a").(*dom.HTMLAnchorElement)
	a.Href = ""
	return &AnchorStruct{a}
}

func (a *AnchorStruct) Text(text string) *AnchorStruct {
	a.SetTextContent(text)
	return a
}

func (a *AnchorStruct) Url(href string) *AnchorStruct {
	a.Href = href
	return a
}

func (a *AnchorStruct) Click(action func(dom.Event)) *AnchorStruct {
	a.AddEventListener("click", true, action)
	return a
}
