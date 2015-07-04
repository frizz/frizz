package selectors

import (
	system "kego.io/system"
)

var ptr8729121024 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Description: "This is an instance of the gallery type containing two nested images"}
var ptr8729121152 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Id: system.Reference{}}
var ptr8730001296 = &Photo{Base: ptr8729121152, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}, Path: system.String{Value: "/me.jpg", Exists: true}}
var ptr8729121280 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Id: system.Reference{}}
var ptr8730126944 = &Diagram{Base: ptr8729121280, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr8729553312 = &Gallery{Base: ptr8729121024, Images: map[string]Image{"dave": ptr8730001296, "foo": ptr8730126944}}
var Faces = ptr8729553312
