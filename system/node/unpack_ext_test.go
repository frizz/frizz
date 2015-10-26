package node_test

import (
	"testing"

	_ "kego.io/demo/common/images"
	_ "kego.io/demo/common/images/types"
	_ "kego.io/demo/common/units"
	_ "kego.io/demo/common/units/types"
	_ "kego.io/demo/common/words"
	_ "kego.io/demo/common/words/types"
	_ "kego.io/demo/site"
	_ "kego.io/demo/site/types"
	"kego.io/ke"
	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/scan"
	"kego.io/system"
	"kego.io/system/node"
	_ "kego.io/system/types"
)

func TestUnpack(t *testing.T) {
	testUnpack(t, "kego.io/system")
	testUnpack(t, "kego.io/demo/site")
}
func testUnpack(t *testing.T, path string) {
	set, err := process.InitialiseManually(false, false, false, false, path)
	assert.NoError(t, err)
	err = scan.ScanForSource(set)
	assert.NoError(t, err)

	sha := system.GetAllSourceInPackage(set.Path)

	for _, sh := range sha {
		var n node.Node
		err := ke.UnmarshalNode(sh.Source, &n, set.Path, set.Aliases)
		assert.NoError(t, err, sh.Id.Name)
	}
}
