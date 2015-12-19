package node_test

import (
	"testing"

	"kego.io/context/envctx"
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
	"kego.io/process/tests"
	"kego.io/system"
	"kego.io/system/node"
	_ "kego.io/system/types"
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

	env, ok := envctx.FromContext(ctx)
	assert.True(t, ok)

	err = scan.ScanForSource(ctx)
	assert.NoError(t, err)

	sha := system.GetAllSourceInPackage(env.Path)

	for _, sh := range sha {
		var n node.Node
		err := ke.UnmarshalNode(ctx, sh.Source, &n)
		assert.NoError(t, err, sh.Id.Name)
	}
}

func TestNodeUnpack(t *testing.T) {
	j := `{"type":"system:package","aliases":{"kego.io/demo/common/images":"images","kego.io/demo/common/units":"units","kego.io/demo/common/words":"words"}}`

	packageNode := &node.Node{}
	err := ke.UnmarshalNode(tests.PathCtx("kego.io/demo/site"), []byte(j), packageNode)
	assert.NoError(t, err)

}
