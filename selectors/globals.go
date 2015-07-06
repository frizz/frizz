package selectors

import (
	system "kego.io/system"
)

var ptr8729021952 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "gallery", Exists: true}, Id: system.Reference{Package: "kego.io/selectors", Name: "faces", Exists: true}, Description: "This is an instance of the gallery type containing two nested images", Rules: []system.Rule(nil)}
var ptr8729022080 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "photo", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8730280736 = &Photo{Base: ptr8729022080, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}, Path: system.String{Value: "/me.jpg", Exists: true}}
var ptr8729022208 = &system.Base{Type: system.Reference{Package: "kego.io/selectors", Name: "diagram", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
var ptr8729060224 = &Diagram{Base: ptr8729022208, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr8729977264 = &Gallery{Base: ptr8729021952, Images: map[string]Image{"dave": ptr8730280736, "foo": ptr8729060224}}
var Faces = ptr8729977264
