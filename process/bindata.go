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

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\xc1\x6e\x9c\x40\x0c\x3d\x33\x5f\x61\x59\xab\x76\x59\x45\x70\xaf\xd4\x73\x95\x4b\x1b\xb5\x51\x7b\x9e\xb2\x86\xd0\xc0\x0c\x05\x53\x75\x45\xf9\xf7\xce\x4c\x0c\x0c\x39\x44\xab\xdd\x93\x07\xbf\xf7\x6c\x3f\x7b\xa7\x09\x0e\x9d\x2e\x9e\x75\x45\x9f\x75\x4b\xf0\xe1\x23\x64\x0f\xd1\x7b\x9e\x55\x04\x79\xd0\xfc\x14\x43\xc2\x5b\x20\x75\xdb\xd9\x9e\x87\x90\xbe\x97\xd8\xa5\x84\x09\xaf\x0b\xb9\x94\x7a\xa1\xc0\x51\x25\xd8\x53\xd9\x50\xc1\xa8\x54\xe2\x90\xbd\x36\x8e\x22\x92\x1e\x7e\xb7\x3c\x96\x0e\xd6\x72\x4e\x27\xf1\x14\xd1\x8a\x38\x3b\xca\x6e\x02\xcf\x71\x14\x32\x67\x1f\xa6\x4a\x6d\x25\xf9\xd2\x91\x14\xf4\x61\x98\xe6\xd1\x05\xa1\x50\x60\x1d\x8c\xf8\x14\x00\xd9\xfd\x39\xfb\xae\x9b\x91\x16\xd1\x43\x39\x36\xcd\x9a\x5e\x46\x0d\x99\xca\x6e\xb4\x4f\x76\x75\x21\xe4\xfe\xe8\x3e\x58\xd7\xea\x0e\xd0\x23\x50\x1a\x40\x5f\x0e\xa5\x2a\x7a\x71\x94\x1a\x58\x59\x0c\xa2\x28\xb3\xe1\x7e\x4a\x14\x8b\x70\x67\x96\x4a\xf2\xfc\x74\xfb\x4f\xe8\xb0\x1a\x31\xcf\xb0\x7d\xbe\x59\x35\x98\x50\x97\x60\x2c\x2f\xbe\x1a\xa6\xbe\xd4\x05\x45\xf6\x3a\x50\x7e\x82\x1f\x04\x67\x6b\xde\x33\x18\xa2\x33\xf0\x13\x0d\x04\x3f\x6b\x37\x5d\x69\x7b\xa8\x17\x1a\x70\xd8\xda\x29\x5f\x0e\x84\xa9\xed\x1a\xcd\xce\xc4\x81\xfb\xb1\xe0\x8c\xdd\x1b\xc5\xf9\xfd\x45\x6c\x91\x2a\x47\x53\x38\xd1\x9a\x8f\x29\x4c\xf1\x6d\xbe\x79\x28\xc9\x75\xe3\x04\x39\x2a\xa9\x27\xe3\x1a\xc6\xaf\x54\xd5\x83\x83\x3d\x86\xf5\xe3\x33\x55\x36\xab\x6d\xfe\x6b\xb0\x26\xda\xe1\xab\x4b\x3e\xfa\x4d\xac\xa7\xf6\x0f\x7e\x8f\x96\xbd\xfc\x1d\xc8\x5f\x2a\x74\xf5\xa5\x3c\xbe\x5b\x80\xdb\xed\x4d\x73\x9a\xbe\xf4\x2a\xe3\x5e\xd9\xd1\x70\x71\x1f\xdb\x5b\x7a\x8a\x9b\xf8\x76\x31\xac\xff\xee\xb9\xd1\xa5\xa6\xf1\x46\xe6\xff\x01\x00\x00\xff\xff\x89\xa0\x45\x4d\xa8\x04\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 1192, mode: os.FileMode(420), modTime: time.Unix(1430038669, 0)}
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
