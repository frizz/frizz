// info:{"Path":"kego.io/demo/site","Hash":10775463831198305239}
package site

import (
	"reflect"

	"golang.org/x/net/context"
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

// This represents a gallery - it just has a title.
type Gallery1 struct {
	*system.Object
	Bar   system.RuleInterface   `json:"bar"`
	Foo   system.IntInterface    `json:"foo"`
	Title system.StringInterface `json:"title"`
}
type Gallery1Interface interface {
	GetGallery1(ctx context.Context) *Gallery1
}

func (o *Gallery1) GetGallery1(ctx context.Context) *Gallery1 {
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
	GetGallery1a(ctx context.Context) *Gallery1a
}

func (o *Gallery1a) GetGallery1a(ctx context.Context) *Gallery1a {
	return o
}

// This represents a gallery - it has a title and a map of photos
type Gallery2 struct {
	*system.Object
	Images map[string]*images.Photo `json:"images"`
	Title  *system.String           `json:"title"`
}
type Gallery2Interface interface {
	GetGallery2(ctx context.Context) *Gallery2
}

func (o *Gallery2) GetGallery2(ctx context.Context) *Gallery2 {
	return o
}

// This represents a gallery - it has a title and a map of images with a restriction rule
type Gallery2a struct {
	*system.Object
	Images map[string]*images.Photo `json:"images"`
	Title  *system.String           `json:"title"`
}
type Gallery2aInterface interface {
	GetGallery2a(ctx context.Context) *Gallery2a
}

func (o *Gallery2a) GetGallery2a(ctx context.Context) *Gallery2a {
	return o
}

// This represents a gallery - it has a title and a list of images with a restriction rule
type Gallery2b struct {
	*system.Object
	Images map[string]*images.Photo `json:"images"`
	Title  *system.String           `json:"title"`
}
type Gallery2bInterface interface {
	GetGallery2b(ctx context.Context) *Gallery2b
}

func (o *Gallery2b) GetGallery2b(ctx context.Context) *Gallery2b {
	return o
}

// This represents a gallery - it has a localizer title and a map of images
type Gallery3 struct {
	*system.Object
	Images map[string]images.Image `json:"images"`
	Title  words.Localizer         `json:"title"`
}
type Gallery3Interface interface {
	GetGallery3(ctx context.Context) *Gallery3
}

func (o *Gallery3) GetGallery3(ctx context.Context) *Gallery3 {
	return o
}

// This represents a gallery - it has a localizer title and a map of images with a custom rule
type Gallery3a struct {
	*system.Object
	Images map[string]images.Image `json:"images"`
	Title  words.Localizer         `json:"title"`
}
type Gallery3aInterface interface {
	GetGallery3a(ctx context.Context) *Gallery3a
}

func (o *Gallery3a) GetGallery3a(ctx context.Context) *Gallery3a {
	return o
}
func init() {
	json.RegisterPackage("kego.io/demo/site", 10775463831198305239)
	json.RegisterType("kego.io/demo/site", "gallery1", reflect.TypeOf((*Gallery1)(nil)), reflect.TypeOf((*Gallery1Rule)(nil)), reflect.TypeOf((*Gallery1Interface)(nil)).Elem())
	json.RegisterType("kego.io/demo/site", "gallery1a", reflect.TypeOf((*Gallery1a)(nil)), reflect.TypeOf((*Gallery1aRule)(nil)), reflect.TypeOf((*Gallery1aInterface)(nil)).Elem())
	json.RegisterType("kego.io/demo/site", "gallery2", reflect.TypeOf((*Gallery2)(nil)), reflect.TypeOf((*Gallery2Rule)(nil)), reflect.TypeOf((*Gallery2Interface)(nil)).Elem())
	json.RegisterType("kego.io/demo/site", "gallery2a", reflect.TypeOf((*Gallery2a)(nil)), reflect.TypeOf((*Gallery2aRule)(nil)), reflect.TypeOf((*Gallery2aInterface)(nil)).Elem())
	json.RegisterType("kego.io/demo/site", "gallery2b", reflect.TypeOf((*Gallery2b)(nil)), reflect.TypeOf((*Gallery2bRule)(nil)), reflect.TypeOf((*Gallery2bInterface)(nil)).Elem())
	json.RegisterType("kego.io/demo/site", "gallery3", reflect.TypeOf((*Gallery3)(nil)), reflect.TypeOf((*Gallery3Rule)(nil)), reflect.TypeOf((*Gallery3Interface)(nil)).Elem())
	json.RegisterType("kego.io/demo/site", "gallery3a", reflect.TypeOf((*Gallery3a)(nil)), reflect.TypeOf((*Gallery3aRule)(nil)), reflect.TypeOf((*Gallery3aInterface)(nil)).Elem())
}
