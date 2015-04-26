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

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\x51\x8f\x94\x30\x10\x7e\x86\x5f\x31\x99\x6c\x74\xd9\x5c\xe0\xdd\xc4\x67\x73\x2f\x7a\x51\xa3\xcf\x95\x1d\x38\x3c\xb6\xe5\x4a\x31\x6e\x90\xff\xee\xb4\x3b\x40\xc1\x68\xcc\xf1\xd4\x32\xf3\xcd\xf7\xf5\xeb\xd7\x71\x84\x43\xa7\xca\x27\x55\xd3\x7b\x75\x21\x78\xf3\x16\xf2\x87\x68\x3f\x4d\x69\xd4\xf2\xa0\xdc\x63\xdc\x12\xf6\xd2\xd2\x5c\x3a\x63\x5d\x1f\xca\xf7\xb2\xe6\x92\x20\x61\x4f\xc4\xa5\xf4\x06\x81\x63\x9a\xa0\xa5\xaa\xa5\xd2\x61\xca\xeb\x27\xaa\x4d\xde\x98\xe2\x7b\x6f\xb4\xff\xc1\xd0\xa6\x02\x6d\xb8\x93\x9e\xb7\x5a\x96\xde\xfe\xda\x3b\xba\x60\xe6\xe7\x26\xc9\xfe\x77\x98\x41\xfa\x1c\x58\xfd\xda\x2a\xcd\x9a\x44\xb3\xd7\x73\x37\x6f\xe6\x23\x2e\xe7\x09\x03\xbd\x84\x9b\xd8\x08\xb3\x81\x6c\x64\x79\xcc\xca\x98\xa5\xe9\x4a\xe9\xae\x1d\x09\xa1\x5f\x06\xbb\x3e\xf3\xa2\x5f\xb4\x1d\xb4\x5c\x44\x68\xc8\xef\xcf\xf9\x17\xd5\x0e\x34\x0f\x3d\x54\x43\xdb\x2e\xe5\xd9\xcb\x50\xa9\xcd\x0a\x7b\x67\x16\x9b\x43\xed\x87\xb2\xe1\x6e\x2e\xaa\x03\xf4\x1d\x28\x02\xd0\xd3\xa1\xb0\xa2\x1f\x8e\xc2\x81\xb5\xc1\x30\x14\xe5\x6c\xb8\x33\x5f\x2c\xc2\x8d\x59\x69\x52\x14\xa7\x97\x7f\x02\x87\xc5\x88\x69\x82\xf5\xf7\x8b\xa7\xc6\x29\x12\x5f\xb5\x23\x5b\xa9\x92\x22\x7b\xb9\xa9\x38\xc1\x57\x82\xb3\xd1\xaf\x1d\x68\xa2\x33\xb8\x47\xea\x09\xbe\x35\x7c\xba\xca\x58\x68\x66\x18\xb8\x70\x6b\xa7\x62\x0e\x08\x07\xad\x6b\x95\x63\x13\x7b\x67\x87\xd2\xe5\x8e\xf7\x28\xce\x6f\x13\x11\xa5\xb1\x1a\x74\xc9\x43\x1b\x77\xcc\x60\x8c\xb3\xf9\xcf\xa0\x24\xff\x77\x9c\xc4\x3f\xa1\xfc\x23\xd5\x0d\x3f\x03\xeb\xd1\x47\xef\xeb\x12\x9c\x5f\xf0\x3c\x18\xe7\x9b\xef\x40\x5e\x60\xe0\xf8\x50\x1d\x5f\xcd\x8d\x6b\x92\xc6\x29\xcb\x6e\xcc\x22\x3e\xac\x19\x47\x96\x34\x1b\x82\x31\x11\xfe\xf1\x36\xd7\x94\xec\xde\xca\xdf\x35\xc5\x22\x3e\x5d\xb5\x53\x3f\xb7\xd8\x28\x77\x59\xec\xef\xf4\x3b\x00\x00\xff\xff\x95\x46\xb6\xf8\xd7\x04\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1239, mode: os.FileMode(420), modTime: time.Unix(1430063342, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\x31\x4f\x03\x31\x0c\x85\xe7\xdc\xaf\xf0\x70\x13\x42\xe9\x8e\xc4\x46\x55\x58\x2a\x84\x10\x7b\x74\xb8\x55\x04\x8d\x23\xc7\x45\xad\x4e\xfd\xef\x24\x3e\xd2\x5c\x29\x62\x73\xf2\xbe\x3c\xfb\x39\xe3\x08\x7d\x74\xc3\x87\xdb\x22\xdc\xdd\x83\xad\xf5\xe9\xd4\x15\xc9\xef\x22\xb1\x24\x95\x6a\x3d\x49\x7e\x03\x81\x04\xac\x1c\x23\xda\xa7\xb4\x76\xe2\xbf\xf0\xcd\x7d\xee\xb1\x01\x93\xf8\x80\x69\x60\x1f\xc5\x53\xb0\xcb\x83\x4f\x6a\xb1\x58\x40\x66\xae\x81\xea\x90\x45\x0c\xef\xc5\xaa\x20\xca\x6e\x29\x1f\x21\x09\xef\x07\x81\xb1\xeb\xcc\xbc\xc9\xa3\x4b\xcb\x83\xe4\x27\x3a\x9f\x31\x37\x67\xf7\x9f\x6b\xbb\xa2\x17\xdc\x20\x63\x18\xb0\xe5\xea\x67\x79\x4d\xeb\xa9\x35\xbb\x90\x85\x3e\x32\x45\x64\x39\xae\xdd\x0e\x6f\xdb\x51\x57\xa2\x0d\x9e\xa7\x1b\x8f\xa9\xba\xe4\xa9\xce\xdc\x7f\xf1\xff\x86\xae\x57\x60\x2e\xd8\x15\x95\x51\x2e\x07\x2b\x9b\xf9\x05\xbd\xe6\xd9\xaa\x2d\x71\x4b\xda\xcf\xfe\x71\x16\x59\xff\x6c\xaa\xbf\x03\x00\x00\xff\xff\x24\x90\x7f\x9c\x14\x02\x00\x00")

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

	info := bindata_file_info{name: "struct.tmpl", size: 532, mode: os.FileMode(420), modTime: time.Unix(1430038843, 0)}
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
