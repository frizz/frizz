package selectors

import (
	system "kego.io/system"
)

var ptr8729178368 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Description: "This is an instance of the gallery type containing two nested images", Rules: []system.Rule(nil)}
var ptr8729178496 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8728700320 = &Photo{Base: ptr8729178496, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}, Path: system.String{Value: "/me.jpg", Exists: true}}
var ptr8729178624 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8729322976 = &Diagram{Base: ptr8729178624, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr8729560640 = &Gallery{Base: ptr8729178368, Images: map[string]Image{"dave": ptr8728700320, "foo": ptr8729322976}}
var Faces = ptr8729560640
