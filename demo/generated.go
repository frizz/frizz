// info:{"Path":"frizz.io/demo","Hash":2458267122473194559}
package demo

import jsonctx "frizz.io/context/jsonctx"

// notest

func init() {
	pkg := jsonctx.InitPackage("frizz.io/demo")
	pkg.SetHash(uint64(0x221d84ec6609543f))
}
