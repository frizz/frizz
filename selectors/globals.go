package selectors

import (
	system "kego.io/system"
)

var ptr8728923520 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Description: "This is an instance of the gallery type containing two nested images"}
var ptr8728923648 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Id: system.Reference{}}
var ptr8729040512 = &Photo{Base: ptr8728923648, Server: system.String{Value: "www.google.com", Exists: true}, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}}
var ptr8728923776 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Id: system.Reference{}}
var ptr8729405664 = &Diagram{Base: ptr8728923776, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr8729413648 = &Gallery{Base: ptr8728923520, Images: map[string]Image{"dave": ptr8729040512, "foo": ptr8729405664}}
var Faces = ptr8729413648
