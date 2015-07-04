package selectors

import (
	system "kego.io/system"
)

var ptr8729309440 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}}
var ptr8729309568 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Id: system.Reference{}}
var ptr8730416048 = &Photo{Base: ptr8729309568, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}, Path: system.String{Value: "/me.jpg", Exists: true}}
var ptr8729309696 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Id: system.Reference{}}
var ptr8728686976 = &Diagram{Base: ptr8729309696, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr8730167392 = &Gallery{Base: ptr8729309440, Images: map[string]Image{"dave": ptr8730416048, "foo": ptr8728686976}}
var Faces = ptr8730167392
