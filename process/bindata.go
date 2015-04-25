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

var _base_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8e\xb1\x8e\x83\x30\x10\x44\x6b\xf3\x15\x2b\x8b\x02\xae\xf0\xf5\x27\x5d\x93\x2e\x55\x8a\x7c\x01\x31\x03\x71\xc0\x5e\xcb\x36\x05\x42\xfc\x7b\x0c\x08\xd1\x8d\xe6\xbd\xd5\xce\xb2\x50\xe9\x1b\x3d\x34\x3d\xe8\xef\x9f\xd4\x99\xd7\xb5\xd8\x90\xb1\x9e\x43\x8a\x3b\x3a\xf3\x81\x4c\x47\x2a\xcd\x1e\xea\xf9\xe6\x69\x6c\xef\xd6\x8f\xb0\x70\xe9\xd6\x44\xa3\x37\x47\x74\x93\xd3\x54\x31\xfd\x64\x5b\xf5\x9c\xbb\x9a\x32\x45\x55\xef\x55\x40\x87\x00\xa7\x41\xf2\xf1\xfa\x40\x27\x49\x72\x40\xcf\xca\xf0\x6f\x9c\x63\x82\x95\xd7\xff\xf2\xda\x45\x4b\x21\x44\x40\x9a\x82\x23\x56\xc7\x69\x21\xf6\x4d\x70\x6d\x16\xbe\x01\x00\x00\xff\xff\x3e\x86\xa2\x1c\xd3\x00\x00\x00")

func base_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_base_tmpl,
		"base.tmpl",
	)
}

func base_tmpl() (*asset, error) {
	bytes, err := base_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "base.tmpl", size: 211, mode: os.FileMode(420), modTime: time.Unix(1429906606, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\xc1\x6e\x9c\x30\x10\x3d\xe3\xaf\x18\x8d\x56\xed\xb2\x8a\xe0\x5e\xa9\xe7\x2a\x97\x36\x6a\xa3\xf6\xec\xb0\x03\xa1\x01\x9b\xc2\x50\x75\x45\xf9\xf7\xda\xce\x00\x26\x91\xaa\x68\xf7\x34\x66\xde\x9b\xe7\x79\x7e\x3b\x4d\x70\xe8\x74\xf1\xa4\x2b\xfa\xac\x5b\x82\x0f\x1f\x21\xbb\x8b\xce\xf3\xac\x22\xc8\x9d\xe6\xc7\x18\x12\xce\x02\xa9\xdb\xce\xf6\x3c\x84\xf6\xad\xd4\xae\x25\x4c\x78\x29\xe4\x5a\xea\x99\x02\x47\x95\x60\x4f\x65\x43\x05\xa3\x52\x89\x43\xf6\xda\x38\x8a\x8c\xf4\xf0\x9b\xe5\xb0\xdc\x60\x95\x73\x73\x12\x4f\x91\x59\x11\x67\x47\xd9\x6d\xe0\x39\x8e\x42\xe6\xec\xcb\x54\xa9\x4d\x92\x2f\x1d\x89\xa0\x2f\xc3\x36\xf7\xae\x08\x42\x81\x75\x30\xe2\x53\x00\x64\xb7\xe7\xec\xbb\x6e\x46\x5a\x86\x1e\xca\xb1\x69\xd6\xf6\xb2\x6a\xe8\x54\x76\xa3\x7d\xb2\xab\x0b\xa1\xf7\x5b\xf7\xc1\xba\x56\x77\x80\x1e\x81\x72\x01\xf4\x72\x28\xaa\xe8\x87\xa3\x68\x60\x65\x31\x0c\x45\xd9\x0d\xf7\x5b\xa2\x58\x84\x3b\xb3\x54\x92\xe7\xa7\xeb\x7f\x42\x87\xd5\x88\x79\x86\xed\xf3\xd5\x53\x83\x09\x75\x09\xc6\xf2\xe2\xab\x61\xea\x4b\x5d\x50\x64\xaf\x03\xe5\x27\xf8\x41\x70\xb6\xe6\x3d\x83\x21\x3a\x03\x3f\xd2\x40\xf0\x50\xbb\xed\x4a\xdb\x43\xbd\xd0\x80\xc3\xab\x9d\xf2\x25\x20\x4c\x6d\xd7\x68\x76\x26\x0e\xdc\x8f\x05\x67\xec\xce\x28\xce\xbf\x06\x3d\xe8\x81\x5e\x41\xb6\xd0\x6c\x95\x2a\x47\x53\x38\xdd\x9a\x8f\x29\x4c\x71\x7c\xff\x9b\xa5\xe4\x6d\x1b\x87\x71\x54\x52\x4f\xc6\xed\x84\x5f\xa9\xaa\x07\x07\xbb\x0f\x09\xc1\x27\xaa\x6c\x56\xdb\xfc\xe7\x60\x4d\xf4\xcc\x2f\xc2\x7e\xf4\x8f\xb5\xa6\xf1\x2f\xfc\x1a\x2d\xfb\xf1\x37\x20\xff\xba\x70\xab\x2f\xe5\xf1\xdd\x02\xdc\xe2\x39\xcd\x69\xfa\xe6\x6b\x0c\x17\xf7\xb1\xbd\xe6\x22\xb1\xf2\xb7\x8b\x61\xfd\x67\xcf\x8d\x12\x9c\x3e\x5b\x27\xee\x47\xe5\xfc\x2f\x00\x00\xff\xff\x70\x30\xd6\x43\xcc\x04\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1228, mode: os.FileMode(420), modTime: time.Unix(1429909442, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\x31\x4f\x03\x31\x0c\x85\xe7\xdc\xaf\xc8\x70\x13\x42\xe9\x8e\xc4\x46\x55\x58\x2a\x84\x10\xbb\x75\xb8\xa7\x08\x1a\x47\x8e\x8b\x5a\x9d\xfa\xdf\x49\x7c\xa4\xd7\xf6\x2a\x36\x27\xef\xcb\xf3\xb3\x33\x0c\xb6\x8d\xd0\x7d\x41\x8f\xf6\xe1\xd1\xba\x5a\x1f\x8f\x4d\x91\xfc\x36\x12\x4b\x52\xa9\xd6\xa3\xe4\x37\x36\x90\x58\x27\x87\x88\xee\x25\xad\x41\xfc\x0f\x7e\xc0\xf7\x0e\x27\x60\x14\x9f\x30\x75\xec\xa3\x78\x0a\x6e\xb9\xf7\x49\x2d\x16\x0b\x9b\x99\x39\x50\x1d\xb2\x88\xe1\xb3\x58\x15\x44\xd9\x9e\xf2\xd1\x26\xe1\x5d\x27\x76\x68\x1a\x73\xde\xe4\x19\xd2\x72\x2f\xf9\x89\xe6\x33\xe6\xee\xe4\xfe\x77\xed\x56\xf4\x86\x1b\x64\x0c\x1d\x4e\x73\xb5\x67\xf3\x9a\xa9\xa7\xd6\x0c\x21\x0b\x6d\x64\x8a\xc8\x72\x58\xc3\x16\xef\xa7\xa3\xae\x44\x1b\xbc\x8e\x37\x1e\x53\x75\xc9\xa9\x4e\xdc\x7f\xe3\xdf\x86\xe6\x2b\x30\x17\xec\x8a\x4a\x94\xcb\x60\x65\x33\x57\xd0\x7b\xce\x56\x6d\x89\x6f\x8e\x3c\x7b\x02\xfd\xd5\x22\xf4\x27\xc7\xfa\x37\x00\x00\xff\xff\x97\x1a\xc1\x2f\x2a\x02\x00\x00")

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

	info := bindata_file_info{name: "struct.tmpl", size: 554, mode: os.FileMode(420), modTime: time.Unix(1429888909, 0)}
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
	"base.tmpl":   base_tmpl,
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
	"base.tmpl":   &_bintree_t{base_tmpl, map[string]*_bintree_t{}},
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
