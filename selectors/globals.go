package selectors

import (
	system "kego.io/system"
)

var ptr8729227520 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Description: "This is an instance of the gallery type containing two nested images", Rules: []system.Rule(nil)}
var ptr8729227648 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8728829120 = &Photo{Base: ptr8729227648, Server: system.String{Value: "www.google.com", Exists: true}, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}}
var ptr8729227776 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8729585920 = &Diagram{Base: ptr8729227776, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr8729471040 = &Gallery{Base: ptr8729227520, Images: map[string]Image{"dave": ptr8728829120, "foo": ptr8729585920}}
var Faces = ptr8729471040
