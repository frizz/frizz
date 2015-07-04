package selectors

import (
	system "kego.io/system"
)

var ptr8728875264 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}}
var ptr8728875392 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Id: system.Reference{}}
var ptr8728665440 = &Photo{Base: ptr8728875392, Server: system.String{Value: "www.google.com", Exists: true}, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}}
var ptr8728875520 = &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Id: system.Reference{}}
var ptr8730002592 = &Diagram{Base: ptr8728875520, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr8729052560 = &Gallery{Base: ptr8728875264, Images: map[string]Image{"dave": ptr8728665440, "foo": ptr8730002592}}
var Faces = ptr8729052560
