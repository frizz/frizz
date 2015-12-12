package site

import (
	"reflect"

	"kego.io/demo/common/images"
	"kego.io/demo/common/units"
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

// This represents a gallery - it just has a title
type Gallery1 struct {
	*system.Object
	Bar   system.RuleInterface   `json:"bar"`
	Foo   system.IntInterface    `json:"foo"`
	Title system.StringInterface `json:"title"`
}
type Gallery1Interface interface {
	GetGallery1() *Gallery1
}

func (o *Gallery1) GetGallery1() *Gallery1 {
	return o
}

// This represents a gallery - it has a title and an image
type Gallery1a struct {
	*system.Object
	Image *images.Icon     `json:"image"`
	Size  *units.Rectangle `json:"size"`
	Title *system.String   `json:"title"`
}
type Gallery1aInterface interface {
	GetGallery1a() *Gallery1a
}

func (o *Gallery1a) GetGallery1a() *Gallery1a {
	return o
}

// This represents a gallery - it has a title and a map of photos
type Gallery2 struct {
	*system.Object
	Images map[string]*images.Photo `json:"images"`
	Title  *system.String           `json:"title"`
}
type Gallery2Interface interface {
	GetGallery2() *Gallery2
}

func (o *Gallery2) GetGallery2() *Gallery2 {
	return o
}

// This represents a gallery - it has a title and a map of images with a restriction rule
type Gallery2a struct {
	*system.Object
	Images map[string]*images.Photo `json:"images"`
	Title  *system.String           `json:"title"`
}
type Gallery2aInterface interface {
	GetGallery2a() *Gallery2a
}

func (o *Gallery2a) GetGallery2a() *Gallery2a {
	return o
}

// This represents a gallery - it has a title and a list of images with a restriction rule
type Gallery2b struct {
	*system.Object
	Images map[string]*images.Photo `json:"images"`
	Title  *system.String           `json:"title"`
}
type Gallery2bInterface interface {
	GetGallery2b() *Gallery2b
}

func (o *Gallery2b) GetGallery2b() *Gallery2b {
	return o
}

// This represents a gallery - it has a localizer title and a map of images
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

// This represents a gallery - it has a localizer title and a map of images with a custom rule
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
	json.Register("kego.io/demo/site", "@gallery1", reflect.TypeOf((*Gallery1Rule)(nil)), nil, 18124994774018985003)
	json.Register("kego.io/demo/site", "@gallery1a", reflect.TypeOf((*Gallery1aRule)(nil)), nil, 7234649160741892262)
	json.Register("kego.io/demo/site", "@gallery2", reflect.TypeOf((*Gallery2Rule)(nil)), nil, 1938792963787640012)
	json.Register("kego.io/demo/site", "@gallery2a", reflect.TypeOf((*Gallery2aRule)(nil)), nil, 3956662793233262470)
	json.Register("kego.io/demo/site", "@gallery2b", reflect.TypeOf((*Gallery2bRule)(nil)), nil, 13459768753991397545)
	json.Register("kego.io/demo/site", "@gallery3", reflect.TypeOf((*Gallery3Rule)(nil)), nil, 7443957450467844980)
	json.Register("kego.io/demo/site", "@gallery3a", reflect.TypeOf((*Gallery3aRule)(nil)), nil, 14112747451711469678)
	json.Register("kego.io/demo/site", "gallery1", reflect.TypeOf((*Gallery1)(nil)), reflect.TypeOf((*Gallery1Interface)(nil)).Elem(), 18124994774018985003)
	json.Register("kego.io/demo/site", "gallery1a", reflect.TypeOf((*Gallery1a)(nil)), reflect.TypeOf((*Gallery1aInterface)(nil)).Elem(), 7234649160741892262)
	json.Register("kego.io/demo/site", "gallery2", reflect.TypeOf((*Gallery2)(nil)), reflect.TypeOf((*Gallery2Interface)(nil)).Elem(), 1938792963787640012)
	json.Register("kego.io/demo/site", "gallery2a", reflect.TypeOf((*Gallery2a)(nil)), reflect.TypeOf((*Gallery2aInterface)(nil)).Elem(), 3956662793233262470)
	json.Register("kego.io/demo/site", "gallery2b", reflect.TypeOf((*Gallery2b)(nil)), reflect.TypeOf((*Gallery2bInterface)(nil)).Elem(), 13459768753991397545)
	json.Register("kego.io/demo/site", "gallery3", reflect.TypeOf((*Gallery3)(nil)), reflect.TypeOf((*Gallery3Interface)(nil)).Elem(), 7443957450467844980)
	json.Register("kego.io/demo/site", "gallery3a", reflect.TypeOf((*Gallery3a)(nil)), reflect.TypeOf((*Gallery3aInterface)(nil)).Elem(), 14112747451711469678)
}
