// info:{"Path":"kego.io/selectors/tests","Hash":995749587514777472}
package tests

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for basic
type BasicRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for c
type CRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for collision
type CollisionRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for diagram
type DiagramRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for expr
type ExprRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for gallery
type GalleryRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for image
type ImageRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for instance
type InstanceRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for instanceItem
type InstanceItemRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for kid
type KidRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for photo
type PhotoRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for polykids
type PolykidsRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for rectangle
type RectangleRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for rightscale
type RightscaleRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for rightscaleLink
type RightscaleLinkRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for rightscaleList
type RightscaleListRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for sibling
type SiblingRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for simple
type SimpleRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for simpleItem
type SimpleItemRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for typed
type TypedRule struct {
	*system.Object
	*system.Rule
}
type Basic struct {
	*system.Object
	DrinkPreference   []*system.String            `json:"drinkPreference"`
	FavoriteColor     *system.String              `json:"favoriteColor"`
	LanguagesSpoken   []map[string]*system.String `json:"languagesSpoken"`
	Name              map[string]*system.String   `json:"name"`
	SeatingPreference []*system.String            `json:"seatingPreference"`
	Weight            *system.Number              `json:"weight"`
}
type BasicInterface interface {
	GetBasic(ctx context.Context) *Basic
}

func (o *Basic) GetBasic(ctx context.Context) *Basic {
	return o
}

type C struct {
	*system.Object
	A *system.Number            `json:"a"`
	B *system.Number            `json:"b"`
	C map[string]*system.Number `json:"c"`
}
type CInterface interface {
	GetC(ctx context.Context) *C
}

func (o *C) GetC(ctx context.Context) *C {
	return o
}

type Collision struct {
	*system.Object
	Number map[string]*system.String `json:"number"`
}
type CollisionInterface interface {
	GetCollision(ctx context.Context) *Collision
}

func (o *Collision) GetCollision(ctx context.Context) *Collision {
	return o
}

type Diagram struct {
	*system.Object
	Url *system.String `json:"url"`
}
type DiagramInterface interface {
	GetDiagram(ctx context.Context) *Diagram
}

func (o *Diagram) GetDiagram(ctx context.Context) *Diagram {
	return o
}

type Expr struct {
	*system.Object
	False   *system.Bool   `json:"false"`
	Float   *system.Number `json:"float"`
	Int     *system.Number `json:"int"`
	Null    *system.String `json:"null"`
	String  *system.String `json:"string"`
	String2 *system.String `json:"string2"`
	True    *system.Bool   `json:"true"`
}
type ExprInterface interface {
	GetExpr(ctx context.Context) *Expr
}

func (o *Expr) GetExpr(ctx context.Context) *Expr {
	return o
}

// This represents a gallery - it's just a list of images
type Gallery struct {
	*system.Object
	Images map[string]Image `json:"images"`
}
type GalleryInterface interface {
	GetGallery(ctx context.Context) *Gallery
}

func (o *Gallery) GetGallery(ctx context.Context) *Gallery {
	return o
}

type Instance struct {
	*system.Object
	Child        map[string]*InstanceItem `json:"child"`
	Cloud_type   *system.String           `json:"cloud_type"`
	Display_name *system.String           `json:"display_name"`
	Links        []*RightscaleLink        `json:"links"`
	Name         *system.String           `json:"name"`
}
type InstanceInterface interface {
	GetInstance(ctx context.Context) *Instance
}

func (o *Instance) GetInstance(ctx context.Context) *Instance {
	return o
}

type InstanceItem struct {
	*system.Object
	Name *system.String `json:"name"`
}
type InstanceItemInterface interface {
	GetInstanceItem(ctx context.Context) *InstanceItem
}

func (o *InstanceItem) GetInstanceItem(ctx context.Context) *InstanceItem {
	return o
}

type Kid struct {
	*system.Object
	Language  *system.String `json:"language"`
	Level     *system.String `json:"level"`
	Preferred *system.Bool   `json:"preferred"`
}
type KidInterface interface {
	GetKid(ctx context.Context) *Kid
}

func (o *Kid) GetKid(ctx context.Context) *Kid {
	return o
}

// This represents an image, and contains path, server and protocol separately
type Photo struct {
	*system.Object
	// The path for the url - e.g. /foo/bar.jpg
	Path *system.String `json:"path"`
	// The protocol for the url - e.g. http or https
	Protocol *system.String `kego:"{\"default\":{\"value\":\"http\"}}" json:"protocol"`
	// The server for the url - e.g. www.google.com
	Server *system.String `json:"server"`
	Size   *Rectangle     `json:"size"`
}
type PhotoInterface interface {
	GetPhoto(ctx context.Context) *Photo
}

func (o *Photo) GetPhoto(ctx context.Context) *Photo {
	return o
}

type Polykids struct {
	*system.Object
	A []*Kid `json:"a"`
}
type PolykidsInterface interface {
	GetPolykids(ctx context.Context) *Polykids
}

func (o *Polykids) GetPolykids(ctx context.Context) *Polykids {
	return o
}

type Rectangle struct {
	*system.Object
	Height *system.Int `json:"height"`
	Width  *system.Int `json:"width"`
}
type RectangleInterface interface {
	GetRectangle(ctx context.Context) *Rectangle
}

func (o *Rectangle) GetRectangle(ctx context.Context) *Rectangle {
	return o
}

type Rightscale struct {
	*system.Object
	Child        map[string]*InstanceItem `json:"child"`
	Cloud_type   *system.String           `json:"cloud_type"`
	Display_name *system.String           `json:"display_name"`
	Links        []*RightscaleLink        `json:"links"`
	Name         *system.String           `json:"name"`
}
type RightscaleInterface interface {
	GetRightscale(ctx context.Context) *Rightscale
}

func (o *Rightscale) GetRightscale(ctx context.Context) *Rightscale {
	return o
}

type RightscaleLink struct {
	*system.Object
	Href *system.String `json:"href"`
	Rel  *system.String `json:"rel"`
}
type RightscaleLinkInterface interface {
	GetRightscaleLink(ctx context.Context) *RightscaleLink
}

func (o *RightscaleLink) GetRightscaleLink(ctx context.Context) *RightscaleLink {
	return o
}

type RightscaleList struct {
	*system.Object
	Foo []*Rightscale `json:"foo"`
}
type RightscaleListInterface interface {
	GetRightscaleList(ctx context.Context) *RightscaleList
}

func (o *RightscaleList) GetRightscaleList(ctx context.Context) *RightscaleList {
	return o
}

type Sibling struct {
	*system.Object
	A *system.Number            `json:"a"`
	B *system.Number            `json:"b"`
	C *C                        `json:"c"`
	D map[string]*system.Number `json:"d"`
	E map[string]*system.Number `json:"e"`
}
type SiblingInterface interface {
	GetSibling(ctx context.Context) *Sibling
}

func (o *Sibling) GetSibling(ctx context.Context) *Sibling {
	return o
}

type Simple struct {
	*system.Object
	A *SimpleItem `json:"a"`
}
type SimpleInterface interface {
	GetSimple(ctx context.Context) *Simple
}

func (o *Simple) GetSimple(ctx context.Context) *Simple {
	return o
}

type SimpleItem struct {
	*system.Object
	B *system.String `json:"b"`
}
type SimpleItemInterface interface {
	GetSimpleItem(ctx context.Context) *SimpleItem
}

func (o *SimpleItem) GetSimpleItem(ctx context.Context) *SimpleItem {
	return o
}

type Typed struct {
	*system.Object
	Avatar          Image                     `json:"avatar"`
	DrinkPreference []*system.String          `json:"drinkPreference"`
	FavoriteColor   *system.String            `json:"favoriteColor"`
	Kids            map[string]*Kid           `json:"kids"`
	Name            map[string]*system.String `json:"name"`
	Weight          *system.Number            `json:"weight"`
}
type TypedInterface interface {
	GetTyped(ctx context.Context) *Typed
}

func (o *Typed) GetTyped(ctx context.Context) *Typed {
	return o
}
func init() {
	json.RegisterPackage("kego.io/selectors/tests", 995749587514777472)
	json.RegisterType("kego.io/selectors/tests", "basic", reflect.TypeOf((*Basic)(nil)), reflect.TypeOf((*BasicRule)(nil)), reflect.TypeOf((*BasicInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "c", reflect.TypeOf((*C)(nil)), reflect.TypeOf((*CRule)(nil)), reflect.TypeOf((*CInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "collision", reflect.TypeOf((*Collision)(nil)), reflect.TypeOf((*CollisionRule)(nil)), reflect.TypeOf((*CollisionInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "diagram", reflect.TypeOf((*Diagram)(nil)), reflect.TypeOf((*DiagramRule)(nil)), reflect.TypeOf((*DiagramInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "expr", reflect.TypeOf((*Expr)(nil)), reflect.TypeOf((*ExprRule)(nil)), reflect.TypeOf((*ExprInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "gallery", reflect.TypeOf((*Gallery)(nil)), reflect.TypeOf((*GalleryRule)(nil)), reflect.TypeOf((*GalleryInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "image", reflect.TypeOf((*Image)(nil)).Elem(), reflect.TypeOf((*ImageRule)(nil)), nil)
	json.RegisterType("kego.io/selectors/tests", "instance", reflect.TypeOf((*Instance)(nil)), reflect.TypeOf((*InstanceRule)(nil)), reflect.TypeOf((*InstanceInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "instanceItem", reflect.TypeOf((*InstanceItem)(nil)), reflect.TypeOf((*InstanceItemRule)(nil)), reflect.TypeOf((*InstanceItemInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "kid", reflect.TypeOf((*Kid)(nil)), reflect.TypeOf((*KidRule)(nil)), reflect.TypeOf((*KidInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "photo", reflect.TypeOf((*Photo)(nil)), reflect.TypeOf((*PhotoRule)(nil)), reflect.TypeOf((*PhotoInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "polykids", reflect.TypeOf((*Polykids)(nil)), reflect.TypeOf((*PolykidsRule)(nil)), reflect.TypeOf((*PolykidsInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "rectangle", reflect.TypeOf((*Rectangle)(nil)), reflect.TypeOf((*RectangleRule)(nil)), reflect.TypeOf((*RectangleInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "rightscale", reflect.TypeOf((*Rightscale)(nil)), reflect.TypeOf((*RightscaleRule)(nil)), reflect.TypeOf((*RightscaleInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "rightscaleLink", reflect.TypeOf((*RightscaleLink)(nil)), reflect.TypeOf((*RightscaleLinkRule)(nil)), reflect.TypeOf((*RightscaleLinkInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "rightscaleList", reflect.TypeOf((*RightscaleList)(nil)), reflect.TypeOf((*RightscaleListRule)(nil)), reflect.TypeOf((*RightscaleListInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "sibling", reflect.TypeOf((*Sibling)(nil)), reflect.TypeOf((*SiblingRule)(nil)), reflect.TypeOf((*SiblingInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "simple", reflect.TypeOf((*Simple)(nil)), reflect.TypeOf((*SimpleRule)(nil)), reflect.TypeOf((*SimpleInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "simpleItem", reflect.TypeOf((*SimpleItem)(nil)), reflect.TypeOf((*SimpleItemRule)(nil)), reflect.TypeOf((*SimpleItemInterface)(nil)).Elem())
	json.RegisterType("kego.io/selectors/tests", "typed", reflect.TypeOf((*Typed)(nil)), reflect.TypeOf((*TypedRule)(nil)), reflect.TypeOf((*TypedInterface)(nil)).Elem())
}
