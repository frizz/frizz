// info:{"Path":"frizz.io/demo","Hash":4074124933468102755}
package demo

import jsonctx "frizz.io/context/jsonctx"

// notest

func init() {
	pkg := jsonctx.InitPackage("frizz.io/demo")
	pkg.SetHash(uint64(0x388a330a95eda063))
}
