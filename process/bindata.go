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

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x4d\x8f\xd3\x30\x10\x3d\xc7\xbf\x62\x64\x55\xd0\x54\x55\x73\x5f\x89\x33\xda\xcb\x52\xc1\x4a\x1c\x10\x07\x93\x4e\xb2\x66\x5b\x3b\xeb\x38\x88\x2a\x9b\xff\x8e\xc7\x1f\x49\xba\x6d\x05\xb4\x27\xd7\xf3\xf1\xe6\xbd\x79\x71\xdf\xc3\xa2\x11\xf6\x09\xee\x3e\xc0\x66\x4b\x87\x61\x60\x74\x29\x0f\x8d\x36\xb6\xf5\xf7\xf7\xf1\xec\x42\x8d\x28\x9f\x45\x8d\x30\xd6\xbd\x82\x12\x07\xa4\x10\x0b\x25\xb0\x64\x19\x37\x58\xed\xb1\xb4\x9c\xb9\xf3\x33\xd6\x7a\x23\x75\xf1\xb3\xd5\x8a\x2e\x5c\xa9\xac\x40\x69\x97\x89\x2f\xb1\xcb\x98\xd4\x1e\x5b\x8b\x07\x9e\x53\xc3\x2c\x7b\x7b\xed\x8b\x51\xed\x3c\x1c\x9d\x8d\x50\x6e\x98\x38\xec\x83\x1b\x64\x9d\xfe\x6c\x23\xa9\x91\x88\x6f\x48\xd8\x61\xca\x59\xcd\x49\x49\x98\x87\x92\x27\xa8\x9c\xb1\x09\xcb\x1e\x1b\x8c\x48\x74\xf4\x02\x3d\xba\x43\x3b\x0e\xb5\xf0\x8a\x10\x36\x25\x6c\xee\x77\xa9\xdd\xa2\xea\xf6\xfb\x31\xf0\x10\x75\xf3\x91\x5a\x4f\x05\x1f\x75\x0a\x85\xd8\x2f\x61\xfc\x1e\x0e\xa2\x01\x4e\x19\x3c\x42\x73\x02\xe2\x11\x8f\x53\x73\x1e\x31\x78\xad\xb9\x6f\xca\x89\x0e\x4f\x2a\x47\x2d\xf8\x89\x2a\x2c\x2b\x8a\xd5\xed\xbf\x58\x0e\x23\xf1\x61\x80\xe9\xfa\xe6\xae\x73\x9f\x44\x1d\x95\x45\x53\x89\x12\xd3\x2a\x8b\x15\x7c\x45\xd8\x69\xf5\xde\x82\x42\xdc\x81\x7d\xc2\x16\xe1\x87\x74\xbc\x2a\x6d\x40\x8e\x05\xd6\xef\x67\x55\x24\x0f\x38\x2f\x35\x7b\x61\x9d\x68\xad\x35\x5d\x69\x37\xd6\xfd\xe7\x51\xe9\xd3\xdd\xcf\x0c\x57\x75\xaa\x04\x72\xb1\x5f\xf7\x32\xa7\x85\x7c\x73\x0d\xa4\xaa\xbf\x47\xc3\x7b\x27\x40\xcf\x32\x83\xb6\x33\xea\x5a\x46\x1f\xc6\xf8\x17\x47\x65\xd9\xdf\x84\xf0\x09\x93\xa3\x5e\xe1\xa5\xd3\x96\x62\x77\x30\xc7\xfc\x54\x2d\xdf\xa5\xc4\xc9\x62\xfd\x90\xaf\x23\x46\xa4\x79\x72\x1e\x58\xe2\x1d\xbe\xc0\x0b\xcc\x57\x44\x04\x2b\x34\xa8\xdc\x44\x6f\xbf\x58\xe0\x8f\xc1\xb0\xde\x80\x33\xdb\x5d\x16\xe9\xc6\x66\xff\xad\xe7\x65\xb9\xe6\xfa\x7c\x39\x2a\x2b\x7e\x9f\x43\xad\xaf\x09\x24\x95\xb4\x4e\x99\x3e\xbc\x4b\xe7\x1c\xfc\xf3\x07\xfc\x33\xd6\xd2\x91\x31\xdb\xf0\x90\x5e\x20\xb3\x9c\xbd\xad\x69\xb6\xf5\xe4\xbb\xfc\x1a\xc0\x28\xd2\xad\x10\xb3\x15\xe7\x8e\xd6\x9f\x00\x00\x00\xff\xff\x3a\x7a\x98\x1a\x1b\x06\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1563, mode: os.FileMode(420), modTime: time.Unix(1431188535, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\xcd\x4e\x03\x31\x0c\x84\xcf\xd9\xa7\xb0\xaa\x1e\x00\xa1\xed\x1d\x89\x1b\xa8\x70\xa9\x10\x42\xdc\xa3\xc5\x2d\x91\x68\x1c\x1c\x17\x51\xad\xf6\xdd\x49\x9c\xed\xfe\xf0\x77\x73\xc6\x5f\x26\x1e\xa7\x6d\x61\x19\xac\xbc\xc2\xd5\x35\xd4\x5a\x74\x5d\x95\x45\xb7\x0f\xc4\x12\x55\x3f\xd5\xa5\xe5\xb6\xe0\x49\xa0\x96\x63\xc0\xfa\x3e\x6e\xac\xb8\x0f\x7c\xb6\x6f\x07\x9c\x03\x67\xf8\xde\x43\x37\x18\x1b\x76\x41\x1c\x79\x58\x2c\xce\x13\xb6\x5a\x41\x02\x7f\x76\xbb\x2e\xc9\xe8\x5f\xb2\x53\x6e\x2a\xb5\xa3\x74\x84\x28\x7c\x68\x04\xda\xaa\x32\xe5\x8d\x72\xfb\xce\xc6\xdb\x4f\x49\x57\x74\x3c\x63\x2e\x06\xdf\x5e\xae\xd7\xf4\x88\x5b\x64\xf4\x0d\xf6\x59\x97\x93\x40\x66\x7c\x50\x6b\xb6\x7e\x97\x39\xa6\x80\x2c\xc7\x8d\xdd\xe3\xe5\x78\xd4\x75\xa8\xfb\x43\x51\x1c\x0e\x2e\x93\xd8\x03\xff\x67\xf4\xdf\x89\x69\x7c\x33\xa3\xd6\x94\x27\x99\xcf\x95\xb7\xf2\x0d\x7a\x4a\xa3\x9d\x0c\x89\xff\x8f\xab\x7f\x55\xea\xaf\x00\x00\x00\xff\xff\xd6\x83\x8a\x1c\x06\x02\x00\x00")

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

	info := bindata_file_info{name: "struct.tmpl", size: 518, mode: os.FileMode(420), modTime: time.Unix(1431170447, 0)}
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
