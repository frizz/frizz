package site

import (
	"reflect"

	"kego.io/demo/common/images"
	"kego.io/demo/common/words"
	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for gallery1
type Gallery1Rule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for gallery1a
type Gallery1aRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for gallery2
type Gallery2Rule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for gallery2a
type Gallery2aRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for gallery2b
type Gallery2bRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for gallery3
type Gallery3Rule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for gallery3a
type Gallery3aRule struct {
	*system.Object
	*system.Rule
}

// This represents a gallery - it's just a list of images
type Gallery1 struct {
	*system.Object
	Title system.String `json:"title"`
}
type Gallery1Interface interface {
	GetGallery1() *Gallery1
}

func (o *Gallery1) GetGallery1() *Gallery1 {
	return o
}

// This represents a gallery - it's just a list of images
type Gallery1a struct {
	*system.Object
	Title system.String `json:"title"`
}
type Gallery1aInterface interface {
	GetGallery1a() *Gallery1a
}

func (o *Gallery1a) GetGallery1a() *Gallery1a {
	return o
}

// This represents a gallery - it's just a list of images
type Gallery2 struct {
	*system.Object
	Images map[string]*images.Photo `json:"images"`
	Title  system.String            `json:"title"`
}
type Gallery2Interface interface {
	GetGallery2() *Gallery2
}

func (o *Gallery2) GetGallery2() *Gallery2 {
	return o
}

// This represents a gallery - it's just a list of images
type Gallery2a struct {
	*system.Object
	Images map[string]*images.Photo `json:"images"`
	Title  system.String            `json:"title"`
}
type Gallery2aInterface interface {
	GetGallery2a() *Gallery2a
}

func (o *Gallery2a) GetGallery2a() *Gallery2a {
	return o
}

// This represents a gallery - it's just a list of images
type Gallery2b struct {
	*system.Object
	Images map[string]*images.Photo `json:"images"`
	Title  system.String            `json:"title"`
}
type Gallery2bInterface interface {
	GetGallery2b() *Gallery2b
}

func (o *Gallery2b) GetGallery2b() *Gallery2b {
	return o
}

// This represents a gallery - it's just a list of images
type Gallery3 struct {
	*system.Object
	Images map[string]images.Image `json:"images"`
	Title  words.Localizer         `json:"title"`
}
type Gallery3Interface interface {
	GetGallery3() *Gallery3
}

func (o *Gallery3) GetGallery3() *Gallery3 {
	return o
}

// This represents a gallery - it's just a list of images.
type Gallery3a struct {
	*system.Object
	Images map[string]images.Image `json:"images"`
	Title  words.Localizer         `json:"title"`
}
type Gallery3aInterface interface {
	GetGallery3a() *Gallery3a
}

func (o *Gallery3a) GetGallery3a() *Gallery3a {
	return o
}
func init() {
	json.Register("kego.io/demo/site", "@gallery1", reflect.TypeOf(&Gallery1Rule{}), nil, 9914653089594241493)
	json.Register("kego.io/demo/site", "@gallery1a", reflect.TypeOf(&Gallery1aRule{}), nil, 17736143005787605048)
	json.Register("kego.io/demo/site", "@gallery2", reflect.TypeOf(&Gallery2Rule{}), nil, 4609417854725294272)
	json.Register("kego.io/demo/site", "@gallery2a", reflect.TypeOf(&Gallery2aRule{}), nil, 5335559691650145607)
	json.Register("kego.io/demo/site", "@gallery2b", reflect.TypeOf(&Gallery2bRule{}), nil, 10214881650404650557)
	json.Register("kego.io/demo/site", "@gallery3", reflect.TypeOf(&Gallery3Rule{}), nil, 3637886575604579657)
	json.Register("kego.io/demo/site", "@gallery3a", reflect.TypeOf(&Gallery3aRule{}), nil, 7797969851483200773)
	json.Register("kego.io/demo/site", "gallery1", reflect.TypeOf(&Gallery1{}), reflect.TypeOf((*Gallery1Interface)(nil)).Elem(), 9914653089594241493)
	json.Register("kego.io/demo/site", "gallery1a", reflect.TypeOf(&Gallery1a{}), reflect.TypeOf((*Gallery1aInterface)(nil)).Elem(), 17736143005787605048)
	json.Register("kego.io/demo/site", "gallery2", reflect.TypeOf(&Gallery2{}), reflect.TypeOf((*Gallery2Interface)(nil)).Elem(), 4609417854725294272)
	json.Register("kego.io/demo/site", "gallery2a", reflect.TypeOf(&Gallery2a{}), reflect.TypeOf((*Gallery2aInterface)(nil)).Elem(), 5335559691650145607)
	json.Register("kego.io/demo/site", "gallery2b", reflect.TypeOf(&Gallery2b{}), reflect.TypeOf((*Gallery2bInterface)(nil)).Elem(), 10214881650404650557)
	json.Register("kego.io/demo/site", "gallery3", reflect.TypeOf(&Gallery3{}), reflect.TypeOf((*Gallery3Interface)(nil)).Elem(), 3637886575604579657)
	json.Register("kego.io/demo/site", "gallery3a", reflect.TypeOf(&Gallery3a{}), reflect.TypeOf((*Gallery3aInterface)(nil)).Elem(), 7797969851483200773)
}
