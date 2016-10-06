// info:{"Path":"kego.io/tests/data","Hash":6167257871666940584}
package data

// ke: {"file": {"notest": true}}

import (
	"context"
	"reflect"

	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for alajs
type AlajsRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for alas
type AlasRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for alass
type AlassRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for aljb
type AljbRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for aljn
type AljnRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for aljs
type AljsRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for almjs
type AlmjsRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for alms
type AlmsRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for almss
type AlmssRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for als
type AlsRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for alss
type AlssRule struct {
	*system.Object
	*system.Rule
}

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
type Alajs []string
type AlajsInterface interface {
	GetAlajs(ctx context.Context) *Alajs
}

func (o Alajs) GetAlajs(ctx context.Context) Alajs {
	return o
}

type Alas []*Simple
type AlasInterface interface {
	GetAlas(ctx context.Context) *Alas
}

func (o Alas) GetAlas(ctx context.Context) Alas {
	return o
}

type Alass []*system.String
type AlassInterface interface {
	GetAlass(ctx context.Context) *Alass
}

func (o Alass) GetAlass(ctx context.Context) Alass {
	return o
}

type Aljb bool
type AljbInterface interface {
	GetAljb(ctx context.Context) *Aljb
}

func (o *Aljb) GetAljb(ctx context.Context) *Aljb {
	return o
}

type Aljn float64
type AljnInterface interface {
	GetAljn(ctx context.Context) *Aljn
}

func (o *Aljn) GetAljn(ctx context.Context) *Aljn {
	return o
}

type Aljs string
type AljsInterface interface {
	GetAljs(ctx context.Context) *Aljs
}

func (o *Aljs) GetAljs(ctx context.Context) *Aljs {
	return o
}

type Almjs map[string]string
type AlmjsInterface interface {
	GetAlmjs(ctx context.Context) *Almjs
}

func (o Almjs) GetAlmjs(ctx context.Context) Almjs {
	return o
}

type Alms map[string]*Simple
type AlmsInterface interface {
	GetAlms(ctx context.Context) *Alms
}

func (o Alms) GetAlms(ctx context.Context) Alms {
	return o
}

type Almss map[string]*system.String
type AlmssInterface interface {
	GetAlmss(ctx context.Context) *Almss
}

func (o Almss) GetAlmss(ctx context.Context) Almss {
	return o
}

type Als Simple
type AlsInterface interface {
	GetAls(ctx context.Context) *Als
}

func (o *Als) GetAls(ctx context.Context) *Als {
	return o
}

type Alss system.String
type AlssInterface interface {
	GetAlss(ctx context.Context) *Alss
}

func (o *Alss) GetAlss(ctx context.Context) *Alss {
	return o
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

// v2
type Multi struct {
	*system.Object
	Aalajs []Alajs                           `json:"aalajs"`
	Aalas  []Alas                            `json:"aalas"`
	Aalass []Alass                           `json:"aalass"`
	Aaljb  []*Aljb                           `json:"aaljb"`
	Aaljn  []*Aljn                           `json:"aaljn"`
	Aaljs  []*Aljs                           `json:"aaljs"`
	Aalmjs []Almjs                           `json:"aalmjs"`
	Aalms  []Alms                            `json:"aalms"`
	Aalmss []Almss                           `json:"aalmss"`
	Aals   []*Als                            `json:"aals"`
	Aalss  []*Alss                           `json:"aalss"`
	Ai     []Face                            `json:"ai"`
	Ajb    []bool                            `json:"ajb"`
	Ajn    []float64                         `json:"ajn"`
	Ajs    []string                          `json:"ajs"`
	Alajs  Alajs                             `json:"alajs"`
	Alas   Alas                              `json:"alas"`
	Alass  Alass                             `json:"alass"`
	Aljb   *Aljb                             `json:"aljb"`
	Aljbi  AljbInterface                     `json:"aljbi"`
	Aljn   *Aljn                             `json:"aljn"`
	Aljni  AljnInterface                     `json:"aljni"`
	Aljs   *Aljs                             `json:"aljs"`
	Aljsi  AljsInterface                     `json:"aljsi"`
	Almjs  Almjs                             `json:"almjs"`
	Alms   Alms                              `json:"alms"`
	Almss  Almss                             `json:"almss"`
	Als    *Als                              `json:"als"`
	Alss   *Alss                             `json:"alss"`
	Am     []*Multi                          `json:"am"`
	Anri   []system.StringInterface          `json:"anri"`
	Ari    []system.RuleInterface            `json:"ari"`
	Asb    []*system.Bool                    `json:"asb"`
	Asi    []*system.Int                     `json:"asi"`
	Asn    []*system.Number                  `json:"asn"`
	Asp    []*system.Package                 `json:"asp"`
	Asr    []*system.Reference               `json:"asr"`
	Ass    []*system.String                  `json:"ass"`
	Bri    system.BoolInterface              `json:"bri"`
	I      Face                              `json:"i"`
	Jb     bool                              `json:"jb"`
	Jn     float64                           `json:"jn"`
	Js     string                            `json:"js"`
	M      *Multi                            `json:"m"`
	Malajs map[string]Alajs                  `json:"malajs"`
	Malas  map[string]Alas                   `json:"malas"`
	Malass map[string]Alass                  `json:"malass"`
	Maljb  map[string]*Aljb                  `json:"maljb"`
	Maljn  map[string]*Aljn                  `json:"maljn"`
	Maljs  map[string]*Aljs                  `json:"maljs"`
	Malmjs map[string]Almjs                  `json:"malmjs"`
	Malms  map[string]Alms                   `json:"malms"`
	Malmss map[string]Almss                  `json:"malmss"`
	Mals   map[string]*Als                   `json:"mals"`
	Malss  map[string]*Alss                  `json:"malss"`
	Mi     map[string]Face                   `json:"mi"`
	Mjb    map[string]bool                   `json:"mjb"`
	Mjn    map[string]float64                `json:"mjn"`
	Mjs    map[string]string                 `json:"mjs"`
	Mm     map[string]*Multi                 `json:"mm"`
	Mnri   map[string]system.StringInterface `json:"mnri"`
	Mri    map[string]system.RuleInterface   `json:"mri"`
	Msb    map[string]*system.Bool           `json:"msb"`
	Msi    map[string]*system.Int            `json:"msi"`
	Msn    map[string]*system.Number         `json:"msn"`
	Msp    map[string]*system.Package        `json:"msp"`
	Msr    map[string]*system.Reference      `json:"msr"`
	Mss    map[string]*system.String         `json:"mss"`
	Nri    system.NumberInterface            `json:"nri"`
	Ri     system.RuleInterface              `json:"ri"`
	Sb     *system.Bool                      `json:"sb"`
	Si     *system.Int                       `json:"si"`
	Sn     *system.Number                    `json:"sn"`
	Sp     *system.Package                   `json:"sp"`
	Sr     *system.Reference                 `json:"sr"`
	Sri    system.StringInterface            `json:"sri"`
	Ss     *system.String                    `json:"ss"`
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
	pkg := jsonctx.InitPackage("kego.io/tests/data", 6167257871666940584)
	pkg.InitType("alajs", reflect.TypeOf((Alajs)(nil)), reflect.TypeOf((*AlajsRule)(nil)), reflect.TypeOf((*AlajsInterface)(nil)).Elem())
	pkg.InitType("alas", reflect.TypeOf((Alas)(nil)), reflect.TypeOf((*AlasRule)(nil)), reflect.TypeOf((*AlasInterface)(nil)).Elem())
	pkg.InitType("alass", reflect.TypeOf((Alass)(nil)), reflect.TypeOf((*AlassRule)(nil)), reflect.TypeOf((*AlassInterface)(nil)).Elem())
	pkg.InitType("aljb", reflect.TypeOf((*Aljb)(nil)), reflect.TypeOf((*AljbRule)(nil)), reflect.TypeOf((*AljbInterface)(nil)).Elem())
	pkg.InitType("aljn", reflect.TypeOf((*Aljn)(nil)), reflect.TypeOf((*AljnRule)(nil)), reflect.TypeOf((*AljnInterface)(nil)).Elem())
	pkg.InitType("aljs", reflect.TypeOf((*Aljs)(nil)), reflect.TypeOf((*AljsRule)(nil)), reflect.TypeOf((*AljsInterface)(nil)).Elem())
	pkg.InitType("almjs", reflect.TypeOf((Almjs)(nil)), reflect.TypeOf((*AlmjsRule)(nil)), reflect.TypeOf((*AlmjsInterface)(nil)).Elem())
	pkg.InitType("alms", reflect.TypeOf((Alms)(nil)), reflect.TypeOf((*AlmsRule)(nil)), reflect.TypeOf((*AlmsInterface)(nil)).Elem())
	pkg.InitType("almss", reflect.TypeOf((Almss)(nil)), reflect.TypeOf((*AlmssRule)(nil)), reflect.TypeOf((*AlmssInterface)(nil)).Elem())
	pkg.InitType("als", reflect.TypeOf((*Als)(nil)), reflect.TypeOf((*AlsRule)(nil)), reflect.TypeOf((*AlsInterface)(nil)).Elem())
	pkg.InitType("alss", reflect.TypeOf((*Alss)(nil)), reflect.TypeOf((*AlssRule)(nil)), reflect.TypeOf((*AlssInterface)(nil)).Elem())
	pkg.InitType("face", reflect.TypeOf((*Face)(nil)).Elem(), reflect.TypeOf((*FaceRule)(nil)), nil)
	pkg.InitType("facea", reflect.TypeOf((*Facea)(nil)), reflect.TypeOf((*FaceaRule)(nil)), reflect.TypeOf((*FaceaInterface)(nil)).Elem())
	pkg.InitType("faceb", reflect.TypeOf((*Faceb)(nil)), reflect.TypeOf((*FacebRule)(nil)), reflect.TypeOf((*FacebInterface)(nil)).Elem())
	pkg.InitType("multi", reflect.TypeOf((*Multi)(nil)), reflect.TypeOf((*MultiRule)(nil)), reflect.TypeOf((*MultiInterface)(nil)).Elem())
	pkg.InitType("simple", reflect.TypeOf((*Simple)(nil)), reflect.TypeOf((*SimpleRule)(nil)), reflect.TypeOf((*SimpleInterface)(nil)).Elem())
}
