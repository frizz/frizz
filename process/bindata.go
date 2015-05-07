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

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\xc1\x8e\xd3\x30\x10\x3d\x27\x5f\x31\x1a\x55\xd0\x54\xab\xe4\x8e\xc4\x19\xed\x05\x56\x80\xc4\xd9\xa4\x93\x6c\xd8\xd6\xce\x3a\x13\x44\x15\xf2\xef\xd8\xee\xc4\x71\x8a\x16\xa4\xcd\xc9\xce\xcc\x9b\xf7\xfc\xfc\x3c\x4d\xb0\xeb\x55\xfd\xa4\x5a\xfa\xa8\xce\x04\xef\xde\x43\xf9\x90\xec\xe7\x39\x4f\x5a\x1e\x14\x3f\xa6\x2d\x61\x2f\x2d\xdd\xb9\x37\x96\x87\x50\xbe\x97\xb5\x2b\x09\x12\x6e\x89\x5c\x29\xbf\x42\x60\x9f\x67\x68\xa9\x39\x51\xcd\x98\xbb\xf5\x13\xb5\xa6\xec\x4c\xf5\x63\x30\xda\xff\x70\xd0\xae\x01\x6d\x5c\x27\x3d\x6f\xb5\xc4\xde\xe1\x32\x30\x9d\xb1\xf0\x73\xb3\xec\xf6\x77\x98\x41\xfa\x18\x58\xfd\xda\x2a\xed\x34\x89\x66\xaf\xe7\x6e\xd9\x2c\x47\x8c\xe7\x09\x03\xbd\x84\xab\xd8\x04\xb3\x81\x6c\x64\x79\xcc\xca\x58\xe4\xf9\x4a\xc9\x97\x9e\x84\xd0\x2f\x83\x5d\x5f\xdd\x62\x88\xda\x76\x5a\x2e\x22\x34\x94\xf7\xc7\x65\xdc\xae\x19\x4f\xa7\x58\x58\x5c\x0c\x95\xd6\xac\x80\x0f\x26\x1a\x1c\x6a\x3f\x95\x0d\xb7\x72\x56\x3d\xa0\xef\x40\xa1\x46\x4f\x84\xc2\x87\x7e\x38\x0a\x07\xb6\x06\xc3\x50\x94\x53\xe1\x8d\xed\x62\x0e\x6e\x6c\xca\xb3\xaa\x3a\xbc\xfe\x13\x38\x44\x0b\xe6\x19\xd6\xdf\xaf\x9e\x9a\xe6\x47\x1c\xd5\x4c\xb6\x51\x35\x2d\x77\x5b\x1d\xe0\x1b\xc1\xd1\xe8\xb7\x0c\x9a\xe8\x08\xfc\x48\x03\xc1\xf7\xce\x9d\xab\x31\x16\xba\x08\xe0\x70\x53\x87\x6a\x09\x85\x0b\x57\x7f\x52\xec\xec\x1b\xd8\x8e\x35\x97\xec\xf6\x28\x9e\x6f\x53\x90\x24\xb0\x19\x75\xed\x86\x76\xbc\x2f\x60\x4a\xf3\xf8\xcf\x70\x64\xff\x3b\x48\xe6\x1f\x4c\xf9\x99\xda\xce\x85\xde\x7a\xdc\xde\x7b\x19\xc3\xf2\x1b\x9e\x47\xc3\xbe\xf9\x0e\xe4\xbd\x85\xe9\x9f\x9a\xfd\x9b\xa5\x71\x4d\xcf\x34\x17\xc5\x95\x53\x64\x87\xb5\xc3\x91\x25\xed\x28\x31\x25\xc2\xbf\x5e\xe2\x36\x30\x49\x4c\x5e\xd6\x94\x8a\xf8\x72\xd1\xac\x7e\xbd\x38\xa4\x48\x9d\x9d\xff\x04\x00\x00\xff\xff\xd5\x39\x3f\x77\xc5\x04\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1221, mode: os.FileMode(420), modTime: time.Unix(1431010732, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\xcb\x4a\x43\x31\x10\x86\xd7\x39\x4f\x31\x94\x2e\x54\xe4\x74\x2f\xb8\x53\xaa\x9b\x22\x22\xee\x43\x9c\x96\xa0\xcd\xc4\xc9\x54\x2c\x87\xbc\xbb\xb9\xf4\xdc\xbc\xe0\x6e\x32\xff\x37\x97\x7f\xd2\x75\xb0\xf4\xda\xbc\xea\x1d\xc2\xd5\x35\xb4\x7d\x1c\x63\x93\x25\xbb\xf7\xc4\x12\x8a\xd4\xc7\x55\xb2\x5b\x70\x24\xd0\xca\xd1\x63\x7b\x1f\x36\x5a\xec\x07\x3e\xeb\xb7\x03\xce\x81\x33\x7c\x3f\x41\x37\x18\x0c\x5b\x2f\x96\x1c\x2c\x16\xe7\x09\x5b\xad\x20\x81\x3f\xd5\x18\x53\x1a\xdd\x4b\xee\x94\xc5\x42\xed\x28\x3d\x21\x08\x1f\x8c\x40\xd7\x34\xaa\xce\xa8\xd5\x77\x3a\xdc\x7e\x4a\x2a\x29\xeb\x29\x75\x31\xf4\x3d\xa5\xdb\x35\x3d\xe2\x16\x19\x9d\xc1\xd1\xf1\x72\xe2\x49\x8d\x33\x4b\xcc\xda\x65\xc2\x33\x79\x64\x39\x6e\xf4\x1e\x2f\xc7\x67\xb9\x48\x19\xf0\x50\x33\x16\x87\x2e\x13\xe7\x03\xff\xa7\xfb\xdf\x89\xe9\x05\xd4\x8c\x5a\x53\xde\x64\xbe\x57\x3e\xcc\x37\xe8\x29\xad\xd6\x37\x24\xfe\xd7\x71\xf9\xb1\x1a\x7f\x05\x00\x00\xff\xff\x9b\x89\x45\x23\x12\x02\x00\x00")

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

	info := bindata_file_info{name: "struct.tmpl", size: 530, mode: os.FileMode(420), modTime: time.Unix(1431010555, 0)}
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
