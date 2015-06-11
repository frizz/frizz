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

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\xcf\x6f\x9d\x30\x0c\x3e\xc3\x5f\x61\x59\x4f\x1b\x3c\x55\x70\x9f\xb4\xf3\xd4\x4b\x57\x4d\x95\x76\xce\xc0\xa1\xac\x90\xd0\x10\x26\x55\x94\xff\x7d\x71\x08\x01\x76\x58\xa5\xc7\xc9\xf8\xc7\xf7\xd9\x9f\x9d\x79\x86\xcb\x20\xec\x33\x7c\xf9\x0a\xc5\x23\x1b\xcb\x92\xb2\xb3\xed\x07\x6d\xec\xe8\xfd\xf7\xc1\x76\xa1\x41\x54\x2f\xa2\x21\x88\x75\xef\xa0\x44\x4f\x1c\x4a\xd7\x12\xc8\xd2\x04\x0d\xc9\x8e\x2a\x8b\xa9\xb3\x5f\xa8\xd1\x45\xab\xcb\xdf\xa3\x56\xec\x70\xa5\xad\x04\xa5\x5d\x26\xbd\x06\x94\x98\x34\xbe\x8d\x96\x7a\xcc\x19\x30\x49\xfe\x75\xfb\x62\x52\xb5\xa7\x63\xdb\x08\xe5\x9a\x09\xcd\x3e\xb8\x46\xee\xb6\x9f\xc7\x30\x54\x1c\xc4\x03\x32\xf7\xda\xe5\xa1\xe6\x54\xb2\xf6\xc3\xc9\x3b\x55\x9e\xa6\x3b\x97\x7d\x1b\x28\x30\xb1\xe9\x05\x7a\x72\xc6\x18\x9b\xba\x78\x45\x98\x9b\x13\x8a\xfb\x7a\x83\xbb\xc8\xa9\xeb\x62\xe0\x21\xe8\xe6\x23\x8d\xde\x0b\xbe\xe9\x2d\xb4\xc6\xfe\x08\xe3\xf7\xd0\x8b\x01\x90\x33\x30\x50\x23\x13\x61\xe0\x43\x06\xc7\xc0\x81\x8d\x46\x0f\x8a\x3c\x0e\x6e\x2a\x07\x2d\xf0\xa4\x4a\x9a\x94\xe5\xf5\xf6\x2f\x94\x43\x1c\x7c\x59\x60\x77\xdf\x8c\x7a\xbc\x93\xa0\xa3\xb2\x64\xa4\xa8\x68\x5b\x65\x79\x85\x9f\x04\xb5\x56\x9f\x2d\x28\xa2\x1a\xec\x33\x8d\x04\xbf\x5a\x37\x97\xd4\x06\xda\x58\x60\xfd\x7e\xae\xe5\x76\x03\xee\x96\x86\x4e\x58\x27\xda\x68\xcd\x54\xd9\xc2\xba\x7f\x0c\x4a\x9f\x77\x7f\x38\x38\x39\xa9\xca\x81\xb6\x36\xcb\x61\x3e\x9e\xdf\x7f\x4f\x22\xf9\x68\x10\x0f\x44\x92\x0c\x29\xe7\x3a\x3f\x17\xc0\x1f\xd4\xb4\xee\xf6\xcd\xd3\xba\x77\xbf\xc7\xc3\xf6\x32\x96\x3d\x5e\xd3\x3b\xbc\x4e\xda\x32\xee\x1d\x84\x27\xe8\x1b\xf9\x2e\xb3\x4f\x5b\xe2\x7e\x5e\xf3\x92\xe7\x6b\x7b\x61\xc2\x83\xb9\xfc\x0d\x00\x00\xff\xff\x30\x4f\xdd\x51\x19\x04\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1049, mode: os.FileMode(420), modTime: time.Unix(1432913174, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x51\x4d\x6f\xc2\x30\x0c\x3d\xa7\xbf\xc2\x42\x1c\xd8\x34\x95\xfb\xa4\xdd\x86\xd8\x2e\x68\x9a\xa6\xdd\xa3\x62\x58\x24\x48\xb2\xc4\x9d\x40\x55\xff\xfb\x1c\xa7\x94\xb0\x0f\x6e\xce\x7b\xcf\xcf\x7e\x4e\xd7\xc1\xd4\x6b\xfa\x80\xfb\x07\xa8\xa5\xe8\xfb\x2a\x81\x66\xef\x5d\xa0\x28\xf8\xa9\xce\x94\xd9\x80\xb6\x6b\x98\x59\x47\x50\xd3\xd1\x63\xfd\x1c\x57\x9a\xcc\x17\xbe\xeb\x5d\x8b\x37\x25\xb3\x38\x34\xbb\x76\xcd\xd8\xd8\x9a\xb8\x19\x7e\x0e\xfc\x23\xc6\x26\x18\x4f\xc6\x59\x98\x4c\x92\x6c\x3e\x07\x16\xfe\x66\xfb\x9e\x61\xe4\xb9\xec\x94\x48\x51\x6d\x1d\x3f\x21\x52\x68\x1b\x82\xae\xaa\x54\x9e\x91\xbb\x9f\x74\x5c\x1c\x88\x5b\x64\x71\xa5\x6e\x47\xdf\x01\xae\x97\xee\x15\x37\x18\xd0\x36\x38\x5c\x61\x5a\x44\x55\xe7\x81\x52\x07\x6d\xb7\x49\x17\x9c\xc7\x40\xc7\x95\xde\xe3\xdd\xf9\x29\x87\x12\xf7\x97\x8c\x18\x1c\x5d\x8a\xd8\xa3\xfe\xdf\xe8\x7f\x2b\xca\xf8\xea\x42\xb5\x74\x69\x93\xcb\xbd\xd2\x55\x7e\x88\xde\x78\xb5\x93\xa1\x0b\xd7\xe3\xca\x5f\xe5\xfa\x3b\x00\x00\xff\xff\x9f\xbd\xf5\x7e\x20\x02\x00\x00")

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

	info := bindata_file_info{name: "struct.tmpl", size: 544, mode: os.FileMode(420), modTime: time.Unix(1432631021, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _types_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x52\xcb\x6a\xc3\x30\x10\x3c\x5b\x5f\xb1\x98\x1e\x62\x08\xc9\xbd\xd0\x0f\xe8\xa5\x84\xd2\x7b\x11\xee\xda\x55\x83\x25\x45\x92\x0f\xc1\xd5\xbf\x57\xab\x87\x23\xa7\x14\x7a\x5b\x6b\x76\x66\x67\xd6\xbb\x2c\xf0\xa0\xb9\xfb\x84\xc7\x27\x38\x9c\xa8\xf0\x9e\xd1\xa3\x98\xb4\x32\xce\xc6\xf7\xe7\x5c\x07\x48\xf3\xfe\xcc\x47\x84\x95\xf7\x0d\x92\x4f\x48\x10\x4b\x14\xd8\xb1\xa6\x3d\xe3\xa8\x0e\x42\x1d\xbf\xac\x92\x2d\x63\x4d\x68\x17\x03\x48\x15\x50\xbc\x64\xe6\xda\x64\xaf\xd6\xe1\xd4\x76\x24\xd2\x34\xf7\xcf\x91\x8c\xf2\x23\x8e\xf8\x8f\xd0\xd1\x5d\x35\xda\x22\xf7\xfe\x07\xfc\x4b\xd6\x70\x19\x72\xe5\xdc\x2f\x21\xd3\xbe\x7c\x9c\xf2\x7e\xd6\x9d\x44\x61\x72\x92\x02\x57\x9c\x0d\x25\xb9\x2b\xcd\x71\xec\x06\x27\xe4\x66\xa2\x63\x6c\x98\x65\x0f\x42\x0a\xb7\xeb\x60\xa9\x4d\x69\x25\xa4\x43\x93\x7e\x52\xaa\x6f\x2e\x56\x74\xab\x57\xf3\x69\x76\x8e\x64\xd5\x6c\x7a\x8c\x4a\x6f\xd1\x51\x91\x31\x38\xa0\x41\x19\xb0\xfb\x3f\x00\xed\x2b\x8e\x22\x94\x86\x18\x6d\xce\x55\x6d\x63\x47\x2e\xca\x8c\x70\x11\x97\x59\x39\x3a\x89\x7d\x3c\x93\x3c\xd1\xfb\xae\xb6\xe7\x7f\x02\x00\x00\xff\xff\x9e\x11\x97\xd0\x7a\x02\x00\x00")

func types_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_types_tmpl,
		"types.tmpl",
	)
}

func types_tmpl() (*asset, error) {
	bytes, err := types_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "types.tmpl", size: 634, mode: os.FileMode(420), modTime: time.Unix(1434061445, 0)}
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
	"types.tmpl":  types_tmpl,
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
	"types.tmpl":  &_bintree_t{types_tmpl, map[string]*_bintree_t{}},
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
