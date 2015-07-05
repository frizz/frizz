package selectors

import (
	system "kego.io/system"
)

var ptr8728973568 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Description: "This is an instance of the gallery type containing two nested images", Rules: []system.Rule(nil)}
var ptr8728973696 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8728991760 = &Photo{Base: ptr8728973696, Server: system.String{Value: "www.google.com", Exists: true}, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}}
var ptr8728973824 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8729355264 = &Diagram{Base: ptr8728973824, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr8729527376 = &Gallery{Base: ptr8728973568, Images: map[string]Image{"dave": ptr8728991760, "foo": ptr8729355264}}
var Faces = ptr8729527376
