package mdl

import "honnef.co/go/js/dom"

type Image struct {
	*dom.HTMLImageElement
}

func NewImage(url string) *Image {
	i := &Image{}
	i.HTMLImageElement = dom.GetWindow().Document().CreateElement("img").(*dom.HTMLImageElement)
	i.Src = url
	return i
}
