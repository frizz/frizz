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

var _base_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x56\xc8\x4c\x53\xd0\x2b\xa9\x2c\x48\xd5\x0b\xce\xc8\x2f\xcd\x49\xf1\xcc\x2d\xc8\x49\xcd\x4d\xcd\x2b\x71\x4a\x2c\xce\x4c\x56\xa8\xad\xe5\xe2\x4c\x2b\xcd\x4b\x56\xd0\xc8\x57\xd0\x02\xaa\xd6\x4b\xcf\x07\x8a\x69\x2a\x00\x65\x53\x35\x34\x15\xb4\xfc\x93\xb2\x52\x93\x4b\x14\xaa\xb9\x38\x39\x8b\x52\x4b\x4a\x8b\xf2\x14\xf2\xf5\x20\x62\x5c\x9c\xb5\x5c\x40\x0d\xa9\x79\x29\x40\x0d\x80\x00\x00\x00\xff\xff\xd0\x94\x10\x88\x68\x00\x00\x00")

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

	info := bindata_file_info{name: "base.tmpl", size: 104, mode: os.FileMode(420), modTime: time.Unix(1429888857, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _init_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x14\xca\x41\x0a\x85\x20\x10\x00\xd0\xf5\xf7\x14\xb3\xfa\x28\x84\x57\x09\xa2\x1b\xc8\x8c\x18\xa2\xa5\xe3\x22\xa6\xb9\x7b\xb9\x7e\x8f\x46\x09\x90\x4a\x62\xeb\x40\xcc\xef\xe8\xb5\xf8\x0d\x63\xea\x8c\x6d\xbf\x4f\xb4\x22\xe0\x69\xe4\x0c\x0f\x5c\xa3\x32\x82\xea\x02\x0d\x29\x63\x60\x3f\xc7\x4a\xf6\x3f\x53\xac\x1f\x89\x3a\x67\xd4\xbc\x01\x00\x00\xff\xff\x6e\xd8\x13\x08\x56\x00\x00\x00")

func init_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_init_tmpl,
		"init.tmpl",
	)
}

func init_tmpl() (*asset, error) {
	bytes, err := init_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "init.tmpl", size: 86, mode: os.FileMode(420), modTime: time.Unix(1429886416, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\x4d\x8f\xd4\x30\x0c\x3d\x37\xbf\xc2\xb2\x46\x30\x33\x5a\x75\xee\x48\x5c\xb8\xa0\xbd\xc0\x0a\x10\x9c\xb3\x1d\xb7\x5b\x68\x93\x92\xa6\x48\xa3\xd2\xff\x8e\x93\x3a\xfd\x60\x25\x84\x36\x87\x91\x3d\x7e\xcf\x1f\xcf\xee\x38\xc2\xa1\xbf\xf5\x9e\xda\x77\xd6\x36\xf0\xe6\x2d\xe4\x9f\xa3\x0b\xd3\xa4\x42\xb0\xd3\xc5\x0f\x5d\xd1\x07\xdd\x52\x8c\x3e\x6c\xfc\x3d\xe4\x41\xfb\xa7\x2d\x24\xfa\x02\xa9\xdb\xce\x3a\xdf\xc7\xf0\xbd\xd8\x1c\x12\x26\xfc\x5d\x88\x43\x6a\xa6\xc0\x51\x65\xe8\xa8\x6c\xa8\xf0\xa8\x54\xc6\x48\xa7\x0d\x53\x24\x65\x80\xdf\x25\x27\x75\xb0\x94\xe3\x3c\x59\xa0\x48\xae\x0d\x67\x47\xd9\x4d\x10\x38\x4c\x21\x73\x0d\xe6\x49\xa9\xb5\xa4\xbf\x75\x24\x05\x83\x19\xa7\xf9\xc2\x46\x2c\x14\x59\x07\x23\x3a\x45\x40\x7e\x7f\xcd\xbf\xea\x66\xa0\x94\xf4\x50\x0e\x4d\xb3\x84\xd3\xa8\x31\x52\xd9\x95\xf6\xde\xee\x42\xf3\x7e\x42\xd8\x93\x33\xda\xdd\x76\x2b\x43\x04\x9c\xdd\x1c\xd7\x3e\x7e\x69\x17\xd5\x6e\x75\x07\x18\x92\xa2\xf4\x8c\xa1\x43\x94\x46\x31\xf4\x83\xd2\x16\x56\x16\x63\x1f\x92\x0e\x97\xca\x28\xfa\xe0\x5e\x29\x14\x99\x71\x27\xb8\xca\x2e\xe7\x97\x3f\x95\xf1\x0f\x2c\x4a\x4e\x13\xa4\x3f\x5f\xfa\x2e\xb3\x22\x75\x09\xc6\xfa\xb4\x17\xc3\x4a\x96\xba\xa0\xcd\x7a\x18\x74\x39\xc3\x37\x82\xab\x35\xaf\x3d\x18\xa2\x2b\xf8\x27\xea\x09\x1e\x6b\x9e\xac\xb4\x0e\xea\x44\x03\x1f\xb7\x7e\xbe\xa4\x03\x63\x95\xba\x46\x7b\x56\xb4\xf7\x6e\x28\x7c\xee\xd9\x47\x59\xc3\x73\xd0\xa3\xee\xe9\x19\x64\x3d\xba\xd5\x52\xe5\x60\x0a\xae\x5b\xfb\xe3\x09\xc6\xed\xf9\xff\xf3\x16\xb3\xff\x9b\x38\xfb\xde\x5b\x93\x7f\xa2\xaa\xe6\x35\xbb\xc0\x3e\x06\xe5\x97\xdb\xfc\x0d\x3f\x07\xeb\x03\xf8\x0e\xe4\x1b\x8c\x35\x3e\x96\xc7\x57\x09\xb8\x1e\xeb\x38\x9d\x4e\x73\x65\x69\x7e\x63\x4e\x7f\x02\x00\x00\xff\xff\xca\xf1\x5f\x29\x68\x04\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1128, mode: os.FileMode(420), modTime: time.Unix(1429897117, 0)}
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
	"init.tmpl":   init_tmpl,
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
	"init.tmpl":   &_bintree_t{init_tmpl, map[string]*_bintree_t{}},
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
