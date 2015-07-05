package selectors

import (
	system "kego.io/system"
)

var ptr8729021824 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}}
var ptr8729021952 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Id: system.Reference{}}
var ptr8728698688 = &Photo{Base: ptr8729021952, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}, Path: system.String{Value: "/me.jpg", Exists: true}}
var ptr8729022080 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Id: system.Reference{}}
var ptr8729371808 = &Diagram{Base: ptr8729022080, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr8729527504 = &Gallery{Base: ptr8729021824, Images: map[string]Image{"dave": ptr8728698688, "foo": ptr8729371808}}
var Faces = ptr8729527504
