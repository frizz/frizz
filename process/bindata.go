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

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\xc1\x8e\xd3\x30\x10\x3d\x27\x5f\x31\x1a\x55\xd0\x54\xab\xe4\x8e\xc4\x19\xed\x05\x56\x80\xc4\xd9\xa4\x93\x6c\xd8\xd6\xce\x3a\x53\x44\x15\xf2\xef\xd8\xee\xc4\x71\x82\x60\xa5\xcd\xc9\xce\xcc\x9b\xf7\xfc\xfc\x3c\x8e\xb0\xeb\x55\xfd\xa4\x5a\xfa\xa8\xce\x04\xef\xde\x43\xf9\x90\xec\xa7\x29\x4f\x5a\x1e\x14\x3f\xa6\x2d\x61\x2f\x2d\xdd\xb9\x37\x96\x87\x50\xbe\x97\xb5\x2b\x09\x12\xb6\x44\xae\x94\xdf\x20\xb0\xcf\x33\xb4\xd4\x9c\xa8\x66\xcc\xdd\xfa\x89\x5a\x53\x76\xa6\xfa\x31\x18\xed\x7f\x38\x68\xd7\x80\x36\xae\x93\x9e\xd7\x5a\x62\xef\x70\x1d\x98\xce\x58\xf8\xb9\x59\xb6\xfd\x1d\x66\x90\x3e\x06\x56\xbf\xb6\x4a\x3b\x4d\xa2\xd9\xeb\xb9\x9b\x37\xf3\x11\xe3\x79\xc2\x40\x2f\xe1\x26\x36\xc1\xac\x20\x2b\x59\x1e\xb3\x30\x16\x79\xbe\x50\xf2\xb5\x27\x21\xf4\xcb\x60\xd7\x57\xb7\x18\xa2\xb6\x9d\x96\x8b\x08\x0d\xe5\xfd\x71\x1e\xb7\x6b\x2e\xa7\x53\x2c\xcc\x2e\x86\x4a\x6b\x16\xc0\x07\x13\x0d\x0e\xb5\x9f\xca\x86\x5b\x39\xab\x1e\xd0\x77\xa0\x50\xa3\x27\x42\xe1\x43\x3f\x1c\x85\x03\x5b\x83\x61\x28\xca\xa9\x70\x63\xbb\x98\x83\x2b\x9b\xf2\xac\xaa\x0e\xaf\xff\x04\x0e\xd1\x82\x69\x82\xe5\xf7\xab\xa7\xa6\xf9\x11\x47\x35\x93\x6d\x54\x4d\xf3\xdd\x56\x07\xf8\x46\x70\x34\xfa\x2d\x83\x26\x3a\x02\x3f\xd2\x40\xf0\xbd\x73\xe7\x6a\x8c\x85\x2e\x02\x38\xdc\xd4\xa1\x9a\x43\xe1\xc2\xd5\x9f\x14\x3b\xfb\x06\xb6\x97\x9a\x4b\x76\x7b\x14\xcf\xd7\x29\x48\x12\xd8\x5c\x74\xed\x86\x76\xbc\x2f\x60\x4c\xf3\xf8\xdf\x70\x64\x2f\x1d\x24\xf3\x0f\xa6\xfc\x4c\x6d\xe7\x42\x6f\x3d\x6e\xef\xbd\x8c\x61\xf9\x0d\xcf\x17\xc3\xbe\xf9\x0e\xe4\xbd\x85\xe9\x9f\x9a\xfd\x9b\xb9\x71\x49\xcf\x38\x15\xc5\x8d\x53\x64\x87\xb5\xc3\x91\x25\xed\x28\x31\x25\xc2\xbf\x5e\xe2\x92\x8c\xcd\xcb\xf8\xb7\xa6\x54\xc4\x97\xab\x66\xf5\x6b\x8d\x4d\xb2\x56\xa4\xce\x4e\x7f\x02\x00\x00\xff\xff\x40\x75\x1a\x40\xc5\x04\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1221, mode: os.FileMode(420), modTime: time.Unix(1430656311, 0)}
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
