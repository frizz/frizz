package selectors

import (
	system "kego.io/system"
)

var ptr8729071872 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Description: "This is an instance of the gallery type containing two nested images", Rules: []system.Rule(nil)}
var ptr8729072000 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8730402624 = &Photo{Base: ptr8729072000, Server: system.String{Value: "www.google.com", Exists: true}, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}}
var ptr8729072128 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8730234368 = &Diagram{Base: ptr8729072128, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr8730071040 = &Gallery{Base: ptr8729071872, Images: map[string]Image{"dave": ptr8730402624, "foo": ptr8730234368}}
var Faces = ptr8730071040
