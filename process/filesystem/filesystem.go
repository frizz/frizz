package filesystem // import "kego.io/process/filesystem"

import (
	"os"

	"github.com/blang/vfs"
	"github.com/blang/vfs/memfs"
)

func Real() vfs.Filesystem {
	return vfs.OS()
}

func Mock(files map[string]string) vfs.Filesystem {
	fs := memfs.Create()
	for name, file := range files {
		f, _ := fs.OpenFile(name, os.O_RDWR, 0)
		f.Write([]byte(file))
		f.Close()
	}
	return fs
}
