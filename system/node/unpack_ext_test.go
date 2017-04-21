package node_test

import (
	"testing"

	"github.com/dave/ktest/require"

	"context"

	"frizz.io/context/envctx"
	_ "frizz.io/demo/common/images"
	_ "frizz.io/demo/common/units"
	_ "frizz.io/demo/common/words"
	_ "frizz.io/demo/site"
	"frizz.io/process"
	"frizz.io/process/scanner"
	"frizz.io/system/node"
)

func TestUnpack(t *testing.T) {
	testUnpack(t, "frizz.io/system")
	testUnpack(t, "frizz.io/demo/site")
}
func testUnpack(t *testing.T, path string) {
	ctx, _, err := process.Initialise(context.Background(), &process.Options{
		Path: path,
	})
	require.NoError(t, err)

	env := envctx.FromContext(ctx)

	files := scanner.ScanDirToFiles(ctx, env.Dir, env.Recursive)
	bytes := scanner.ScanFilesToBytes(ctx, files)
	for b := range bytes {
		_, err := node.Unmarshal(ctx, b.Bytes)
		require.NoError(t, err, b.File)
	}
}

func TestNodeUnpack(t *testing.T) {

	ctx, _, err := process.Initialise(context.Background(), &process.Options{
		Path: "frizz.io/demo/site",
	})
	require.NoError(t, err)

	j := `{"type":"system:package","aliases":{"images":"frizz.io/demo/common/images","units":"frizz.io/demo/common/units","words":"frizz.io/demo/common/words"}}`

	_, err = node.Unmarshal(ctx, []byte(j))
	require.NoError(t, err)

}
