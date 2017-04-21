package scanner

import (
	"testing"

	"path/filepath"

	"os"

	"io/ioutil"

	"frizz.io/tests"
	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
)

func TestScanFilesToBytes(t *testing.T) {
	cb := tests.New().TempGopath(false)
	defer cb.Cleanup()
	_, dir := cb.TempPackage("a", map[string]string{
		"a.json": "a",
		"b.json": "b",
		"c.txt":  "c",
		"d.go":   "d",
		"e.yml":  "e",
		"f.yaml": "f",
	})

	in := make(chan File)
	ch := ScanFilesToBytes(cb.Ctx(), in)
	go func() {
		in <- File{filepath.Join(dir, "b.json"), nil}
		close(in)
	}()
	out := []Content{}
	for c := range ch {
		require.NoError(t, c.Err)
		out = append(out, c)
	}
	assert.Equal(t, 1, len(out))
	assert.Equal(t, "b", string(out[0].Bytes))

	// check only json, yml and yaml files are included, and check yaml files
	// are converted
	in = make(chan File)
	ch = ScanFilesToBytes(cb.Ctx(), in)
	go func() {
		in <- File{filepath.Join(dir, "b.json"), nil}
		in <- File{filepath.Join(dir, "c.txt"), nil}
		in <- File{filepath.Join(dir, "d.go"), nil}
		in <- File{filepath.Join(dir, "e.yml"), nil}
		in <- File{filepath.Join(dir, "f.yaml"), nil}
		close(in)
	}()
	out = []Content{}
	for c := range ch {
		require.NoError(t, c.Err)
		out = append(out, c)
	}
	assert.Equal(t, 3, len(out))
	assert.Equal(t, "b", string(out[0].Bytes))
	assert.Equal(t, `"e"`, string(out[1].Bytes))
	assert.Equal(t, `"f"`, string(out[2].Bytes))

	// TODO: This test doesn't run reliably. Fix it!
	/*
		// check cb.Done will interrupt
		in = make(chan File)
		ch = ScanFilesToBytes(cb.Ctx(), in)
		go func() {
			cb.Cancel()
			in <- File{filepath.Join(dir, "a.json"), nil}
			in <- File{filepath.Join(dir, "b.json"), nil}
			close(in)
		}()
		cancelled := false
		for c := range ch {
			if c.Err != nil {
				assert.IsError(t, c.Err, "AFBJCTFOKX")
				cancelled = true
				break
			}
		}
		assert.True(t, cancelled)
	*/
}

func TestScanDirToFiles(t *testing.T) {
	cb := tests.New().TempGopath(false)
	defer cb.Cleanup()
	_, dir := cb.TempPackage("a", map[string]string{
		"b.json": "c",
	})
	// make a dir so we hit the IsRegular block
	require.NoError(t, os.Mkdir(filepath.Join(dir, "d"), 0777))

	// put another file in it
	require.NoError(t, ioutil.WriteFile(filepath.Join(dir, "d", "e.json"), []byte("f"), 0777))

	ch := ScanDirToFiles(cb.Ctx(), dir, false)
	out := []File{}
	for f := range ch {
		require.NoError(t, f.Err)
		out = append(out, f)
	}
	assert.Equal(t, 1, len(out))
	assert.Equal(t, filepath.Join(dir, "b.json"), out[0].File)

	ch = ScanDirToFiles(cb.Ctx(), dir, true)
	out = []File{}
	for f := range ch {
		require.NoError(t, f.Err)
		out = append(out, f)
	}
	assert.Equal(t, 2, len(out))
	assert.Equal(t, filepath.Join(dir, "b.json"), out[0].File)
	assert.Equal(t, filepath.Join(dir, "d", "e.json"), out[1].File)

}
