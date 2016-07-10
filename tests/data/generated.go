// info:{"Path":"kego.io/tests/data","Hash":1099647111195420677}
package data

// ke: {"file": {"notest": true}}

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for face
type FaceRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for facea
type FaceaRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for faceb
type FacebRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for multi
type MultiRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for simple
type SimpleRule struct {
	*system.Object
	*system.Rule
}
type Facea struct {
	*system.Object
	A *system.String `json:"a"`
}
type FaceaInterface interface {
	GetFacea(ctx context.Context) *Facea
}

func (o *Facea) GetFacea(ctx context.Context) *Facea {
	return o
}

type Faceb struct {
	*system.Object
	B *system.String `json:"b"`
}
type FacebInterface interface {
	GetFaceb(ctx context.Context) *Faceb
}

func (o *Faceb) GetFaceb(ctx context.Context) *Faceb {
	return o
}

type Multi struct {
	*system.Object
	Ai   []Face                            `json:"ai"`
	Ajb  []bool                            `json:"ajb"`
	Ajn  []float64                         `json:"ajn"`
	Ajs  []string                          `json:"ajs"`
	Am   []*Multi                          `json:"am"`
	Anri []system.StringInterface          `json:"anri"`
	Ari  []system.RuleInterface            `json:"ari"`
	Asb  []*system.Bool                    `json:"asb"`
	Asi  []*system.Int                     `json:"asi"`
	Asn  []*system.Number                  `json:"asn"`
	Asp  []*system.Package                 `json:"asp"`
	Asr  []*system.Reference               `json:"asr"`
	Ass  []*system.String                  `json:"ass"`
	I    Face                              `json:"i"`
	Jb   bool                              `json:"jb"`
	Jn   float64                           `json:"jn"`
	Js   string                            `json:"js"`
	M    *Multi                            `json:"m"`
	Mi   map[string]Face                   `json:"mi"`
	Mjb  map[string]bool                   `json:"mjb"`
	Mjn  map[string]float64                `json:"mjn"`
	Mjs  map[string]string                 `json:"mjs"`
	Mm   map[string]*Multi                 `json:"mm"`
	Mnri map[string]system.StringInterface `json:"mnri"`
	Mri  map[string]system.RuleInterface   `json:"mri"`
	Msb  map[string]*system.Bool           `json:"msb"`
	Msi  map[string]*system.Int            `json:"msi"`
	Msn  map[string]*system.Number         `json:"msn"`
	Msp  map[string]*system.Package        `json:"msp"`
	Msr  map[string]*system.Reference      `json:"msr"`
	Mss  map[string]*system.String         `json:"mss"`
	Nri  system.StringInterface            `json:"nri"`
	Ri   system.RuleInterface              `json:"ri"`
	Sb   *system.Bool                      `json:"sb"`
	Si   *system.Int                       `json:"si"`
	Sn   *system.Number                    `json:"sn"`
	Sp   *system.Package                   `json:"sp"`
	Sr   *system.Reference                 `json:"sr"`
	Ss   *system.String                    `json:"ss"`
}
type MultiInterface interface {
	GetMulti(ctx context.Context) *Multi
}

func (o *Multi) GetMulti(ctx context.Context) *Multi {
	return o
}

type Simple struct {
	*system.Object
	Js string `json:"js"`
}
type SimpleInterface interface {
	GetSimple(ctx context.Context) *Simple
}

func (o *Simple) GetSimple(ctx context.Context) *Simple {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/tests/data", 1099647111195420677)
	pkg.InitType("face", reflect.TypeOf((*Face)(nil)).Elem(), reflect.TypeOf((*FaceRule)(nil)), nil)
	pkg.InitType("facea", reflect.TypeOf((*Facea)(nil)), reflect.TypeOf((*FaceaRule)(nil)), reflect.TypeOf((*FaceaInterface)(nil)).Elem())
	pkg.InitType("faceb", reflect.TypeOf((*Faceb)(nil)), reflect.TypeOf((*FacebRule)(nil)), reflect.TypeOf((*FacebInterface)(nil)).Elem())
	pkg.InitType("multi", reflect.TypeOf((*Multi)(nil)), reflect.TypeOf((*MultiRule)(nil)), reflect.TypeOf((*MultiInterface)(nil)).Elem())
	pkg.InitType("simple", reflect.TypeOf((*Simple)(nil)), reflect.TypeOf((*SimpleRule)(nil)), reflect.TypeOf((*SimpleInterface)(nil)).Elem())
}
