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

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\xc1\x6e\x9c\x30\x10\x3d\xe3\xaf\x18\x8d\x56\xed\xb2\x8a\xe0\x5e\xa9\xe7\x2a\x97\x36\x6a\xa3\xf6\xec\xb0\x03\xa1\x01\x9b\xc2\x50\x75\x45\xf9\xf7\xda\xce\x00\x26\x91\xaa\x68\xf7\x64\x33\xef\xcd\xcc\x7b\x7e\x3b\x4d\x70\xe8\x74\xf1\xa4\x2b\xfa\xac\x5b\x82\x0f\x1f\x21\xbb\x8b\xee\xf3\xac\x22\xc8\x9d\xe6\xc7\x18\x12\xee\x02\xa9\xdb\xce\xf6\x3c\x84\xf2\xad\x9c\x5d\x49\x98\xf0\x72\x90\x2b\xa9\x67\x0a\x1c\x55\x82\x3d\x95\x0d\x15\x8c\x4a\x25\x0e\xd9\x6b\xe3\x28\xd2\xd2\xc3\x6f\x96\xcb\xb2\xc1\x3a\xce\xf5\x49\x3c\x45\x7a\x45\x9c\x1d\x65\xa7\xc0\x73\x1c\x85\xcc\xd9\x1f\x53\xa5\xb6\x91\x7c\xe9\x48\x06\xfa\x63\x50\x73\xef\x0e\x61\x50\x60\x1d\x8c\xf8\x14\x00\xd9\xed\x39\xfb\xae\x9b\x91\x96\xa6\x87\x72\x6c\x9a\xb5\xbc\x48\x0d\x95\xca\x6e\xb4\x4f\x76\x75\x21\xd4\x7e\xeb\x3e\x58\xd7\xea\x0e\xd0\x23\x50\x16\x40\x3f\x0e\x65\x2a\xfa\xe6\x28\x33\xb0\xb2\x18\x9a\xa2\x68\xc3\xbd\x4a\x14\x8b\x70\x67\x96\x4a\xf2\xfc\x74\xfd\x4f\xe8\xb0\x1a\x31\xcf\xb0\x7d\xbe\xba\x6b\x30\xa1\x2e\xc1\x58\x5e\x7c\x35\x4c\x7d\xa9\x0b\x8a\xec\x75\xa0\xfc\x04\x3f\x08\xce\xd6\xbc\x67\x30\x44\x67\xe0\x47\x1a\x08\x1e\x6a\xa7\xae\xb4\x3d\xd4\x0b\x0d\x38\xbc\xda\x29\x5f\x02\xc2\xd4\x76\x8d\x66\x67\xe2\xc0\xfd\x58\x70\xc6\xee\x8e\xe2\xfc\x6b\xd0\x83\x1e\xe8\x15\x64\x0b\xcd\x76\x52\xe5\x68\x0a\x37\xb7\xe6\x63\x0a\x53\x1c\xdf\xff\x66\x29\x79\x9b\xe2\xd0\x8e\x4a\xea\xc9\x38\x4d\xf8\x95\xaa\x7a\x70\xb0\xfb\x90\x10\x7c\xa2\xca\x66\xb5\xcd\x7f\x0e\xd6\x44\xcf\xfc\x22\xec\x47\xff\x58\x6b\x1a\xff\xc2\xaf\xd1\xb2\x6f\x7f\x03\xf2\xaf\x0b\x5b\x7d\x29\x8f\xef\x16\xe0\x16\xcf\x69\x4e\xd3\xe7\x5d\x45\xee\x1b\x37\x1a\x2e\xee\x63\x7b\xcd\x4e\xf1\x12\xdf\x2e\x86\xf5\x9f\x3d\x37\x0a\x73\x1a\xbf\xc8\xfc\x2f\x00\x00\xff\xff\xf9\x81\x89\x71\xcb\x04\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1227, mode: os.FileMode(420), modTime: time.Unix(1429983099, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x31\x4f\x03\x31\x0c\x85\xe7\xdc\xaf\xf0\x70\x13\x42\xe9\x8e\xc4\x46\x55\x58\x2a\x84\x10\x7b\x74\xb8\x55\x04\x8d\x23\xc7\x45\xad\x4e\xfd\xef\x24\x3e\xd2\xb4\x14\xd1\xcd\xf1\xfb\xee\xf9\xd9\x37\x8e\xd0\x47\x37\x7c\xb8\x35\xc2\xdd\x3d\xd8\x5a\x1f\x0e\x5d\x91\xfc\x26\x12\x4b\x52\xa9\xd6\x93\xe4\x57\x10\x48\xc0\xca\x3e\xa2\x7d\x4a\x4b\x27\xfe\x0b\xdf\xdc\xe7\x16\x1b\x30\x89\x0f\x98\x06\xf6\x51\x3c\x05\x3b\xdf\xf9\xa4\x16\xb3\x19\x64\xe6\x12\xa8\x0e\x59\xc4\xf0\x5e\xac\x0a\xa2\xec\x9a\xf2\x13\x92\xf0\x76\x10\x18\xbb\xce\x9c\x0e\x79\x74\x69\xbe\x93\xfc\x89\xe6\x33\xe6\xe6\xe8\xfe\xd3\xb6\x0b\x7a\xc1\x15\x32\x86\x01\xdb\x5e\xfd\xc9\xbe\xa6\xcd\xd4\x9a\x5d\xc8\x42\x1f\x99\x22\xb2\xec\x97\x6e\x83\xb7\xed\xa9\x27\xd1\x01\xcf\x53\xc7\x63\xaa\x2e\x39\xd5\x91\xfb\x6f\xfd\xbf\xa1\xcb\x13\x98\x33\x76\x41\x25\xca\x79\xb0\x72\x99\x5f\xd0\x6b\xce\x56\x6d\x89\xaf\xae\xac\xff\x6c\xaa\xbf\x03\x00\x00\xff\xff\x54\xc7\x5c\x0e\x14\x02\x00\x00")

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

	info := bindata_file_info{name: "struct.tmpl", size: 532, mode: os.FileMode(420), modTime: time.Unix(1429994275, 0)}
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
