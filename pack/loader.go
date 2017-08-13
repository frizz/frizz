package pack

import (
	"encoding/base64"

	"fmt"

	"frizz.io/global"
)

var DefaultLoader = &defaultLoader{
	cache: map[string]*packageDef{},
}

type defaultLoader struct {
	cache map[string]*packageDef
}

type packageDef struct {
	path    string
	context global.PackageContext
	files   map[string]interface{}
}

func (d defaultLoader) Load(pkg global.Package, filename string) interface{} {

	var def *packageDef
	var ok bool
	if def, ok = d.cache[pkg.Path()]; !ok {
		all := map[string]global.Package{}
		pkg.GetImportedPackages(all)
		context := NewPackageContext(pkg.Path(), all)
		def = &packageDef{
			path:    pkg.Path(),
			context: context,
			files:   map[string]interface{}{},
		}
		d.cache[pkg.Path()] = def
	}

	if i, ok := def.files[filename]; ok {
		return i
	}

	database64 := pkg.GetData(filename)
	if database64 == "" {
		panic(fmt.Sprintf("data not found for %s", filename))
	}

	databytes, err := base64.StdEncoding.DecodeString(database64)
	if err != nil {
		panic(fmt.Sprintf("error decoding base54 data for %s: %s", filename, err))
	}

	dataiface, err := Unmarshal(def.context, databytes)
	if err != nil {
		panic(fmt.Sprintf("error unmarshaling %s: %s", filename, err))
	}

	def.files[filename] = dataiface

	return dataiface

}
