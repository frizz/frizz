package process

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\xc1\x8e\xd3\x30\x10\x3d\x27\x5f\x31\x1a\x55\xd0\x54\xab\xe4\x8e\xc4\x19\xed\x05\x56\x80\xe0\x6c\xd2\x49\x36\x6c\x6a\x67\x9d\x09\xa2\x2a\xf9\x77\x6c\x77\x9c\x38\x45\x20\xb4\x39\xd9\x99\x79\xf3\x9e\x9f\x9f\x2f\x17\xd8\x0d\xaa\x7e\x52\x2d\xbd\x57\x27\x82\x37\x6f\xa1\x7c\x48\xf6\xf3\x9c\x27\x2d\x0f\x8a\x1f\xd3\x96\xb0\x97\x96\xee\x34\x18\xcb\x63\x28\xdf\xcb\xda\x95\x04\x09\xb7\x44\xae\x94\x5f\x21\xb0\xcf\x33\xb4\xd4\xf4\x54\x33\xe6\x6e\xfd\x44\xad\x29\x3b\x53\x7d\x1f\x8d\xf6\x3f\x1c\xb4\x6b\x40\x1b\xd7\x49\xcf\x5b\x2d\x4b\xef\x78\x1e\x99\x4e\x58\xf8\xb9\x59\x76\xfb\x3b\xcc\x20\x7d\x0c\xac\x7e\x6d\x95\x76\x9a\x44\xb3\xd7\x73\x17\x37\xf1\x88\xcb\x79\xc2\x40\x2f\xe1\x2a\x36\xc1\x6c\x20\x1b\x59\x1e\xb3\x32\x16\x79\xbe\x52\xf2\x79\x20\x21\xf4\xcb\x60\xd7\x67\xb7\x18\x17\x6d\x3b\x2d\x17\x11\x1a\xca\xfb\x63\x1c\xb7\x6b\xa6\xbe\x5f\x0a\xd1\xc5\x50\x69\xcd\x0a\x78\x67\x16\x83\x43\xed\x87\xb2\xe1\x56\x4e\x6a\x00\xf4\x1d\x28\xd4\xe8\x89\x50\xf8\xd0\x0f\x47\xe1\xc0\xd6\x60\x18\x8a\x72\x2a\xbc\xb1\x5d\xcc\xc1\x8d\x4d\x79\x56\x55\x87\x97\x7f\x02\x87\xc5\x82\x79\x86\xf5\xf7\x8b\xa7\xa6\xf9\x11\x47\x35\x93\x6d\x54\x4d\xe5\x17\xd5\x4f\x14\x6f\xb8\x3a\xc0\x57\x82\xa3\xd1\xaf\x19\x34\xd1\x11\xf8\x91\x46\x82\x6f\x9d\x3b\x5d\x63\x2c\x74\x11\x06\x1c\xee\xeb\x50\xc5\x68\xb8\x88\x0d\xbd\x62\x67\xe2\xc8\x76\xaa\xb9\x64\xb7\x47\x71\x7e\x9b\x85\x24\x87\xcd\xa4\x6b\x37\xb4\xe3\x7d\x01\x97\x34\x95\xff\x8c\x48\xf6\x7f\xc7\xc9\xfc\xe3\x29\x3f\x52\xdb\xb9\x07\x60\x3d\x7a\xef\x7d\x5d\x82\xf3\x0b\x9e\x27\xc3\xbe\xf9\x0e\xe4\xed\x05\x8e\x0f\xcd\xfe\x55\x6c\x5c\x93\x74\x99\x8b\xe2\xca\x2c\xe2\xc3\xda\xe1\xc8\x92\x76\x86\x60\x4a\x84\x7f\xbc\xca\x35\x25\x37\xaf\xe4\xef\x9a\x52\x11\x9f\xce\x9a\xd5\xcf\x2d\x36\xc9\x5d\x91\xfa\x3b\xff\x0e\x00\x00\xff\xff\xaa\x87\x6b\xe8\xd1\x04\x00\x00")

func main_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_main_tmpl,
		"main.tmpl",
	)
}

func main_tmpl() (*asset, error) {
	bytes, err := main_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "main.tmpl", size: 1233, mode: os.FileMode(420), modTime: time.Unix(1430655257, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x91\xcf\x4e\xc3\x30\x0c\xc6\xcf\xe9\x53\x58\x53\x0f\x80\x50\x77\x47\xe2\x06\x1a\x5c\x26\x84\x10\xf7\xa8\x78\x53\x04\x8b\x83\xe3\x21\xa6\xaa\xef\x4e\xfe\xac\x4b\x0a\xec\xe6\xf8\xfb\xc5\xf6\x67\x0f\x03\xb4\x4e\xf7\xef\x7a\x8b\x70\x73\x0b\xdd\x14\x8f\x63\x13\x25\xb3\x73\xc4\xe2\x93\x34\xc5\x59\x32\x1b\xb0\x24\xd0\xc9\xc1\x61\xf7\xe8\xd7\x5a\xcc\x17\xbe\xea\x8f\x3d\xce\x81\x0b\xfc\x3c\x42\x77\xe8\x7b\x36\x4e\x0c\x59\x58\x2c\x2e\x03\xb6\x5c\x42\x00\xff\xaa\xe3\x18\xd2\x68\xdf\x62\xa5\x28\x26\x6a\x4b\xe1\x09\x5e\x78\xdf\x0b\x0c\x4d\xa3\x72\x8f\xfc\xfb\x41\xfb\xfb\x6f\x09\x5f\xd2\x78\x4a\x5d\x9d\xea\x1e\xd3\xdd\x8a\x9e\x71\x83\x8c\xb6\xc7\x62\xab\xad\xec\xaa\xd2\x33\xc5\xac\x6d\x10\x5a\xc7\xe4\x90\xe5\xb0\xd6\x3b\xbc\x2e\xcf\xb4\x91\xd4\xe0\x29\x67\x0c\xfa\xa9\x4a\xe5\xfc\xc4\x9f\x75\xff\x3f\x51\x6f\x40\xcd\xa8\x15\xc5\x49\xe6\x73\xc5\xc5\xfc\x82\x5e\xc2\x68\x53\x41\xe2\x62\xb4\xad\xae\x58\x39\x4e\x17\xcb\xf1\x4f\x00\x00\x00\xff\xff\x4f\x66\x1a\x25\x12\x02\x00\x00")

func struct_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_struct_tmpl,
		"struct.tmpl",
	)
}

func struct_tmpl() (*asset, error) {
	bytes, err := struct_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "struct.tmpl", size: 530, mode: os.FileMode(420), modTime: time.Unix(1430655327, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"main.tmpl":   main_tmpl,
	"struct.tmpl": struct_tmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"main.tmpl":   &_bintree_t{main_tmpl, map[string]*_bintree_t{}},
	"struct.tmpl": &_bintree_t{struct_tmpl, map[string]*_bintree_t{}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
