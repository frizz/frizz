package node_test

import (
	"testing"

	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	_ "kego.io/demo/common/images"
	_ "kego.io/demo/common/units"
	_ "kego.io/demo/common/words"
	_ "kego.io/demo/site"
	"kego.io/ke"
	"kego.io/kerr/assert"
	"kego.io/process"
	"kego.io/process/scanutils"
	"kego.io/system/node"
)

func TestUnpack(t *testing.T) {
	testUnpack(t, "kego.io/system")
	testUnpack(t, "kego.io/demo/site")
}
func testUnpack(t *testing.T, path string) {
	ctx, _, err := process.Initialise(&process.FromDefaults{
		Path: path,
	})
	assert.NoError(t, err)

	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)

	files := scanutils.ScanDirToFiles(ctx, cmd.Dir, env.Recursive)
	bytes := scanutils.ScanFilesToBytes(ctx, files)
	for b := range bytes {
		n := node.NewNode()
		err := ke.UnmarshalUntyped(ctx, b.Bytes, n)
		assert.NoError(t, err, b.File)
	}
}

func TestNodeUnpack(t *testing.T) {

	ctx, _, err := process.Initialise(&process.FromDefaults{
		Path: "kego.io/demo/site",
	})
	assert.NoError(t, err)

	j := `{"type":"system:package","aliases":{"kego.io/demo/common/images":"images","kego.io/demo/common/units":"units","kego.io/demo/common/words":"words"}}`

	packageNode := node.NewNode()
	err = ke.UnmarshalUntyped(ctx, []byte(j), packageNode)
	assert.NoError(t, err)

}
