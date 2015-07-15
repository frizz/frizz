package tests

import (
	"kego.io/system"
)

var ptr0 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "kego.io/selectors/tests", Name: "faces", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "gallery", Exists: true}}
var ptr1 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "photo", Exists: true}}
var ptr2 = &Photo{Base: ptr1, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "http", Exists: true}, Server: system.String{Value: "www.google.com", Exists: true}, Size: (*Rectangle)(nil)}
var ptr3 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/selectors/tests", Name: "diagram", Exists: true}}
var ptr4 = &Diagram{Base: ptr3, Url: system.String{Value: "http://www.other.com/foo.jpg", Exists: true}}
var ptr5 = &Gallery{Base: ptr0, Images: map[string]Image{"dave": ptr2, "foo": ptr4}}
var Faces = ptr5
