// info:{"Path":"kego.io/demo/common/images","Hash":16772887215174173915}
package images

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/demo/common/units"
	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for icon
type IconRule struct {
	*system.Object
	*system.Rule
}

// Restriction rules for images
type ImageRule struct {
	*system.Object
	*system.Rule
	Secure *system.Bool `json:"secure"`
}

// Automatically created basic rule for photo
type PhotoRule struct {
	*system.Object
	*system.Rule
}

// This is a type of image, which just contains the url of the image
type Icon struct {
	*system.Object
	Url *system.String `json:"url"`
}
type IconInterface interface {
	GetIcon(ctx context.Context) *Icon
}

func (o *Icon) GetIcon(ctx context.Context) *Icon {
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
	Server *system.String   `json:"server"`
	Size   *units.Rectangle `json:"size"`
}
type PhotoInterface interface {
	GetPhoto(ctx context.Context) *Photo
}

func (o *Photo) GetPhoto(ctx context.Context) *Photo {
	return o
}
func init() {
	json.RegisterPackage("kego.io/demo/common/images", 16772887215174173915)
	json.RegisterType("kego.io/demo/common/images", "icon", reflect.TypeOf((*Icon)(nil)), reflect.TypeOf((*IconRule)(nil)), reflect.TypeOf((*IconInterface)(nil)).Elem())
	json.RegisterType("kego.io/demo/common/images", "image", reflect.TypeOf((*Image)(nil)).Elem(), reflect.TypeOf((*ImageRule)(nil)), nil)
	json.RegisterType("kego.io/demo/common/images", "photo", reflect.TypeOf((*Photo)(nil)), reflect.TypeOf((*PhotoRule)(nil)), reflect.TypeOf((*PhotoInterface)(nil)).Elem())
}
