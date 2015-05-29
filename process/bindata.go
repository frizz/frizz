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

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x53\x4d\x6f\xdb\x30\x0c\x3d\xdb\xbf\x82\x20\x82\x2d\x0e\x8a\xf8\x3e\x60\xe7\xa1\x97\xae\xd8\x0a\xec\xac\x39\xb4\xeb\x35\x91\x5c\x99\x19\x16\xa4\xfe\xef\x13\x65\x5a\x76\xf6\x55\x20\xf3\x89\xe6\xd7\x7b\x7c\xa4\xce\x67\x58\x75\x86\x1f\xe1\xdd\x7b\xd8\xde\x8b\x31\x0c\xb9\x38\xdb\x43\xe7\x3c\xf7\xd1\x7f\xab\x76\x08\x75\xa6\x7a\x32\x0d\x41\xaa\x7b\x01\x6b\x0e\x24\xa1\x7c\x2c\x81\x75\x9e\xa1\xa7\x7a\x4f\x15\x63\x1e\xec\x27\x6a\xdc\xb6\x75\xe5\xb7\xde\x59\x71\x84\xd2\xb6\x06\xeb\x42\x26\x3d\x6b\x97\x94\xd4\x9f\x7a\xa6\x03\x16\xd2\x30\xcb\x7e\x75\xc7\x62\xb2\xbb\x08\x27\xb6\x37\x36\x90\x51\xb2\x77\x81\xc8\xcd\xf4\x73\xaf\x43\xa5\x41\x62\x43\xc1\x1e\x59\x2e\x6a\x2e\x4a\x46\x3e\x92\x3c\x43\x15\x79\x3e\x63\xf1\xa9\x23\x45\x12\x33\x0a\xf4\x10\x8c\x3e\x91\x5a\x45\x45\x04\x5b\x12\xb6\xb7\xbb\xa9\xdd\xaa\x3e\xee\xf7\x29\x70\xa7\xba\xc5\x48\xe3\xe6\x82\x0f\x6e\x0a\x8d\xb1\xef\xc6\xc7\x3d\x1c\x4c\x07\x28\x19\xa8\xd0\x28\x40\xa8\x78\x28\xcd\x51\x31\xb0\x71\x18\x9b\xa2\x8c\x83\x93\xca\xaa\x05\x5e\xa8\x92\x67\x65\xb9\xb9\xfe\xd3\x72\x48\x83\x0f\x03\xcc\xee\xab\xbb\x2e\xef\x44\x75\xb4\x4c\xbe\x36\x15\x4d\xab\x2c\x37\xf0\x85\x60\xe7\xec\x5b\x06\x4b\xb4\x03\x7e\xa4\x9e\xe0\x6b\x1b\xe6\xaa\x9d\x87\x36\x15\x70\xdc\xcf\xa6\x9c\x6e\x20\xdc\x52\xb7\x37\x1c\x44\xeb\xd9\x1f\x2b\xde\x72\xf8\x47\x55\xfa\x72\xf7\x8b\x83\xab\x8f\xb6\x0a\x4d\x5b\x5e\x17\x70\x5e\x9e\xdf\x3f\x4f\x22\x7b\x6d\x90\xd8\x88\x6a\xf2\x64\x83\xeb\xf2\xb9\x00\x7e\xa2\xa6\x0d\xb7\xef\x1f\xc6\xbd\xc7\x3d\x2e\xb6\xb7\x16\xd9\xd3\x35\xbd\xc0\xf3\xd1\xb1\xf4\xbd\x01\x7d\x82\x91\xc8\xc7\x7a\xfd\x66\x4a\x9c\xcf\xeb\x3c\x14\xc5\x48\x4f\x27\xfc\xcb\xeb\x7a\x7d\xbc\x3f\xd0\xd7\x17\xfb\x1f\x03\x2c\x19\x7f\x3e\x59\x36\x3f\x7e\xaf\x2e\x96\x9c\x87\x9f\x01\x00\x00\xff\xff\x43\x1b\xcd\x01\xcd\x04\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1229, mode: os.FileMode(420), modTime: time.Unix(1432893715, 0)}
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
