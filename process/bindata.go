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

var _cmd_main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x50\xcd\x4e\xfb\x30\x0c\x3f\x27\x4f\xe1\x7f\xf5\x3f\x24\x52\xd4\x89\x2b\x68\x07\x0e\x0c\xed\xc0\xb4\x37\x98\xa2\xce\x1b\xd6\xda\xa4\x24\x19\x62\x9a\xfa\xee\x38\x25\x2b\x88\x0f\x09\x91\x43\xeb\xd8\xce\xef\xeb\x7c\x86\xff\xbd\x4d\x8f\x70\x3d\x87\x7a\x9d\x8b\x61\x90\xb9\x49\x5d\xef\x43\x8a\x63\x7f\x59\x6a\x1e\xf5\xb6\x39\xd8\x3d\x42\x67\xc9\x49\xf9\xb6\x04\x4a\x8a\xca\xc7\x8a\xbf\xbb\x2e\xe5\xdf\x01\xf7\xbe\x26\x3f\xeb\x83\x6f\x30\xf2\x44\x8a\x0d\x4c\xdd\x78\x8a\x09\xbb\x4a\x02\x1f\xa6\xa2\x1d\x38\xcf\x20\xf8\x54\xa4\x7c\x5e\xd4\x99\x58\x7c\x45\x98\xa5\x53\x8f\x99\x96\x41\xd0\x6d\xf3\xd6\x58\x07\xeb\x58\x61\x71\xb0\xb2\x1d\x9a\xcb\x65\x5d\x9c\x4e\xee\x2e\xc0\xef\x8e\x4b\x06\xd5\x0f\xfd\x6f\x38\xb5\x94\xbb\xa3\x6b\xc6\x48\x94\x86\xb3\x14\x5b\x0a\x06\x12\xc6\x64\x20\x60\x73\x0c\x91\x9e\x59\x43\xf6\x66\xa0\x30\x1b\xc0\x10\xb2\x94\x12\x51\xbd\x74\x94\xc8\xb6\x14\x51\x69\x29\x38\x93\x3c\xff\x37\x07\x47\x6d\x86\x14\x9c\x6c\xbd\x0e\xe4\x52\xeb\x14\x8f\xf4\x98\x5e\x3e\x3e\xd6\x77\x2f\x94\xd4\x15\x3f\x1b\xa6\x97\x1f\x90\xef\xd1\x61\xb0\x09\x17\xd4\x62\x54\x97\xee\x62\xf3\x70\xbb\x5c\x19\xf8\x85\x56\x7d\xf3\x47\x31\xc3\x6b\x00\x00\x00\xff\xff\x9c\x2e\x29\x64\x5e\x02\x00\x00")

func cmd_main_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_cmd_main_tmpl,
		"cmd_main.tmpl",
	)
}

func cmd_main_tmpl() (*asset, error) {
	bytes, err := cmd_main_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "cmd_main.tmpl", size: 606, mode: os.FileMode(420), modTime: time.Unix(1434841461, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmd_types_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x91\xb1\x4e\xc3\x30\x10\x86\x67\xfb\x29\x8e\x8a\x21\x91\xa2\x54\xac\xa0\x0e\x95\x68\x51\x17\x14\x89\x2e\x4c\x95\x95\x5e\xcb\xa9\x89\x1d\x6c\x17\x51\x55\x79\x77\xce\xc1\x8d\x02\xed\x00\x12\x78\x48\xce\x77\xe7\xff\x3e\xff\x3e\x1e\xe1\xba\x51\xfe\x05\x6e\x27\x90\x17\x21\x68\x5b\x19\x92\x54\x37\xc6\x7a\xd7\xe5\x17\x31\xe6\x52\xa3\xca\x9d\xda\x22\xd4\x8a\xb4\x94\x9f\x4d\x90\x48\x31\x32\x6e\xc4\xdf\x4d\xed\xc3\x6f\x87\x5b\x93\x93\x19\x37\xd6\x94\xe8\xb8\x22\xc5\x0a\xfa\xac\x3b\x38\x8f\xf5\x48\x02\x2f\x1e\x45\x1b\xd0\x86\x45\xf0\x35\xa2\x7c\x6f\x4c\xc3\x60\x71\xae\x30\xf6\x87\x06\xc3\x58\x16\x41\xbd\x0e\x5d\x5d\xfc\x0b\xc1\xfe\xfa\x6d\x7b\xae\x63\x95\xe6\x9b\x46\x27\x1e\x55\x8d\xd9\x69\x53\x44\xc7\x7a\x97\x86\x7a\x83\x96\x4e\xf5\x62\xfe\x02\x7b\x2a\xe5\x66\xaf\xcb\xce\xda\x24\x85\xa3\x14\x6b\xb2\x19\x78\x74\x3e\x03\x8b\xe5\xde\x3a\x7a\x63\x86\xc0\x9b\x41\x9c\x9c\x01\x5a\x1b\x50\xa2\xd5\xf9\x42\x93\x27\x55\x91\xc3\x24\x95\x82\xad\x08\xf5\xab\x09\x68\xaa\x82\xa4\xe0\x17\xca\x0b\x4b\xda\x57\x3a\xe1\x52\xda\xbd\x42\x58\xc6\xe5\xb3\x77\xf2\xc9\x0d\x1f\x6b\xfb\x93\x03\xe5\x07\xd4\x68\x95\xc7\x39\x55\xe8\x92\x53\x76\xbe\x5a\x3e\x17\xb3\xa7\x0c\x7e\x00\x9b\xde\xfd\x3f\xcd\xfd\x74\x39\xfd\x2b\x18\x21\xbe\x62\xb4\x1f\x01\x00\x00\xff\xff\xe6\x3c\x1a\xd6\x2e\x03\x00\x00")

func cmd_types_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_cmd_types_tmpl,
		"cmd_types.tmpl",
	)
}

func cmd_types_tmpl() (*asset, error) {
	bytes, err := cmd_types_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "cmd_types.tmpl", size: 814, mode: os.FileMode(420), modTime: time.Unix(1434737862, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cmd_validate_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x50\xb1\x4e\xc3\x30\x10\x9d\xed\xaf\x38\x22\x06\x5b\x8a\x52\xb1\x82\x3a\x32\x74\x41\x9d\x58\x2b\x2b\xbd\x96\x53\x1b\x3b\xd8\x2e\xa2\xaa\xfc\xef\x9c\x83\x1b\x25\x50\x44\x86\xe4\xf2\xee\xee\xbd\x77\xef\x72\x81\xfb\xde\xc4\x37\x78\x5c\x42\xb3\xce\x45\x4a\x32\x83\xd4\xf5\xce\xc7\x30\xe0\xab\x52\x73\xab\x37\xed\xc1\xec\x11\x3a\x43\x56\xca\xef\x21\x50\x52\x54\x2e\x54\xfc\xde\x75\x31\x7f\x0e\xb8\x77\x0d\xb9\x45\xef\x5d\x8b\x81\x3b\x52\x6c\x60\x44\xc3\x39\x44\xec\xaa\x1b\xd8\x22\x9e\x7b\x1c\xc6\xd9\x02\xed\xc0\x3a\x26\xc7\xf7\x62\xf1\x27\x81\xce\x86\x44\x66\x19\xaf\x48\xa9\xfa\x85\x5c\x49\x33\x27\xda\x6d\x5e\x1a\x6a\x6f\x2c\x1f\x52\x0e\x7d\x31\x1d\xd6\xd7\x9f\x75\x09\x64\x0c\x61\xaa\x33\x19\x99\xa9\xcd\xf0\x1b\x9a\x5a\xca\xdd\xc9\xb6\x43\x72\x4a\xc3\x45\x8a\x2d\xf9\x1a\x36\x35\x78\x6c\x4f\x3e\xd0\x07\x1b\xc8\x96\x6b\x28\xb2\x35\xa0\xf7\xd9\x47\x89\xb1\x59\x59\x8a\x64\x8e\x14\x50\x69\x29\x38\x9f\xdc\xbf\x5b\x82\xa5\x63\xe6\x13\x9c\x7e\xb3\xf6\x64\xe3\xd1\x2a\x6e\x69\x09\xe5\x71\xa1\x79\xfe\xa4\xa8\x1e\x78\x2d\x8d\x9b\x13\xe6\x57\x66\xdd\x9a\x88\x6a\xf0\xf4\x97\x21\xfd\xf4\xaf\xa2\x10\x73\xad\xf4\x15\x00\x00\xff\xff\x46\x99\x1e\xd6\x61\x02\x00\x00")

func cmd_validate_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_cmd_validate_tmpl,
		"cmd_validate.tmpl",
	)
}

func cmd_validate_tmpl() (*asset, error) {
	bytes, err := cmd_validate_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "cmd_validate.tmpl", size: 609, mode: os.FileMode(420), modTime: time.Unix(1434737862, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _global_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x51\xcd\x4a\x03\x31\x10\x3e\x6f\x9e\x62\x28\x1e\x2c\xc8\x7a\x17\x7a\x16\x2f\x4b\xdf\x40\x62\x9d\x6e\x43\xbb\x99\x35\x89\x82\xb4\x7d\x77\x33\x33\xd9\xcd\x2a\x08\xbd\x4d\xe6\x9b\xef\x87\x2f\xe7\x33\xdc\x8d\x36\x1d\xe0\x69\x03\xed\x96\x87\xeb\xd5\xf0\xd2\x0d\x23\x85\x14\x65\xff\x52\xe6\x0c\x8d\x76\x77\xb4\x3d\xc2\xcc\xbb\x80\xb7\x03\x32\x64\x94\x02\xf7\xa6\xc9\xa8\xdb\x83\xa7\xfc\xc0\x8f\x72\xb8\x3a\x62\x4f\xad\xa3\xc7\xf8\x1d\x13\x0e\xab\x35\x73\x9a\xe6\xef\x5a\xc8\xe8\xdf\x45\x91\xe7\x60\x7d\xf6\x2b\x79\xba\xec\xf5\x30\x3d\xb6\x25\xf7\x9c\x55\x04\xd9\x5b\x83\x2c\x38\xbf\x28\x9a\x87\x8f\xab\xd5\xda\x98\xea\x35\x92\xf3\x09\x83\x76\xa2\x73\x9c\xee\x67\x50\x7b\x9a\x92\x56\x72\xc0\x7d\x4e\x18\xe9\x33\xec\x50\x14\x9e\x4f\xf4\x66\x4f\x2a\xf0\x65\x83\x54\x97\x8f\x5a\xc9\x75\x81\x9e\x4a\x81\xb0\x11\xa8\x30\xff\x93\xbf\xb5\x07\x76\x7a\x85\xfa\x93\x9d\x9a\xb4\x1d\xa5\x83\xf3\xfd\x52\xfc\x27\x00\x00\xff\xff\x6c\xd3\xae\x22\x04\x02\x00\x00")

func global_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_global_tmpl,
		"global.tmpl",
	)
}

func global_tmpl() (*asset, error) {
	bytes, err := global_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "global.tmpl", size: 516, mode: os.FileMode(420), modTime: time.Unix(1434820454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _main_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x53\x5d\x6b\xeb\x30\x0c\x7d\xae\x7f\x85\x30\xe5\x92\x40\xc9\x7d\xbf\x70\x9f\x47\x5f\xba\x32\xfa\x3e\x4c\xa6\xa4\x59\x1b\x3b\x75\xdc\x41\x49\xf3\xdf\x27\xcb\xce\xd7\xd8\xd8\xf6\xa6\xc8\xd2\xd1\x39\x47\x4a\xd7\xc1\xba\x51\xee\x08\xff\xfe\x43\xb6\xf7\x41\xdf\x0b\x9f\xac\xea\xc6\x58\xd7\x72\x7e\x1b\x63\x7a\x6a\x54\x7e\x52\x25\xc2\xd8\x77\x07\xad\x6a\xf4\x4f\x22\xb4\x40\x22\x56\xd2\x62\x71\xc6\xdc\x49\x41\xf1\x09\x4b\x93\x55\xe6\xef\x6b\x6b\xb4\x4f\x50\x6b\x55\x80\x36\x54\x89\x97\x88\x32\x16\xb5\xb7\xd6\x61\x2d\x53\x0f\xb8\x5a\x7d\x4c\x73\x33\xea\x17\x1e\xe7\x63\xab\x34\x91\x89\x64\x77\x44\x64\x33\x7c\xec\xa3\xa8\x51\x08\x03\xfa\xd9\x81\xe5\xac\x67\xd1\x12\xf8\xf8\xe2\x69\x54\x2a\xc4\x34\x8b\xa4\xd1\x10\x77\x6b\x90\xbd\x39\x50\xd0\x0e\xf5\xeb\x37\x65\xd9\xb2\x5a\x35\x20\x7d\x8d\x8c\xa5\xd2\xc3\xca\x41\x6d\xe4\x24\x97\xec\x26\x63\xb8\x27\xdb\x6a\x87\xb6\x50\x39\x8e\xdc\xc9\x83\xe6\xac\x1c\xc1\xb5\xce\x5e\x73\x97\x39\xfa\x96\x71\xec\x92\xf3\xcc\xa8\xe2\xaa\x73\xa8\x74\xe5\x92\x14\xba\xb9\x6d\x5f\x4a\xf9\x8e\x0a\x63\x60\x81\x16\x35\xa5\x96\x1b\x06\xf9\x84\x65\x45\xeb\xb2\x87\xa0\x9f\x25\xcf\x84\x26\xde\x28\xea\xa6\x73\x0b\xb7\x74\x87\xcb\xd5\x38\x0f\xbd\x81\xe1\x8d\xf7\x32\x7f\x88\x17\xc5\x24\x1f\x8b\xe4\x8f\x2f\x64\x6a\x0f\x66\x17\x0e\xb0\xeb\xd3\x34\x50\x8f\xc2\x67\x61\xcf\x1b\xfc\xe1\xd9\x91\x9b\xf0\x0c\x21\x95\xed\x8c\x3b\x56\xba\x14\x0b\x67\x7f\x79\x76\x01\x70\xfa\xab\x22\xe1\x4f\xb0\x05\x2f\x23\xe6\x21\x2c\xb9\xeb\xdf\x03\x00\x00\xff\xff\x0f\x20\x3f\x97\xa5\x03\x00\x00")

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

	info := bindata_file_info{name: "main.tmpl", size: 933, mode: os.FileMode(420), modTime: time.Unix(1434819787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _struct_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\x41\x4b\xc4\x30\x10\x85\xcf\xed\xaf\x18\x4a\x0f\x2a\x12\xef\x82\x17\x41\xc5\xcb\x1e\x3c\x78\x95\x6c\x32\xad\xc1\x36\x29\xc9\x54\x5c\xba\xfb\xdf\x4d\x26\xa9\x5b\x10\xf6\xf6\xfa\x3e\xde\x4c\xde\x74\x59\xa0\x9d\x24\x7d\xc2\xfd\x03\x08\x16\xa7\x53\x9d\x4c\x33\x4e\xce\x53\x60\x7f\xd5\x19\x99\x0e\xa4\xd5\x70\x65\x1d\x81\xa0\xc3\x84\xe2\x35\xec\x24\x99\x6f\x7c\x97\xc3\x8c\xd7\x5b\xf2\xf4\xa3\x86\x59\x47\x2f\x47\xd9\x84\x23\x68\x0c\xca\x9b\x89\x8c\xb3\x89\xb0\xbb\x62\xf1\xe2\x76\x72\xc4\xe8\x43\x20\x3f\x2b\x82\xa5\xae\xab\xbc\xf7\x3c\xf8\x51\x06\xa3\x52\xb6\xaa\x6e\x22\xf3\xd8\xa1\x47\xab\xf0\xc3\x68\x68\xbe\xb0\x77\xc2\xb8\xbb\x70\x08\x84\x63\x03\xcd\x5e\x06\x6c\x4a\xd1\x76\xd3\x26\x8d\xc5\xd8\xa5\x48\x2f\x6d\x8f\xd0\xe2\xb8\x47\xcd\xc5\x73\x07\xfe\xfc\x5b\x95\x71\x7c\xe5\xdb\xba\xf3\xf2\xe0\xed\xe4\xce\xe0\xa0\x53\xbb\xdb\xa2\xcf\x5b\x9e\xd3\x67\xce\xa6\x40\xc1\xff\x4e\xb5\x81\x7c\xa5\x23\xf4\xce\x96\x73\x45\xd2\x3b\xbe\x65\x49\x5f\x7c\x17\xff\x90\xac\x7f\x03\x00\x00\xff\xff\x3b\xb8\x66\x9e\x05\x02\x00\x00")

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

	info := bindata_file_info{name: "struct.tmpl", size: 517, mode: os.FileMode(420), modTime: time.Unix(1434819787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _types_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x52\x4d\x6b\x83\x40\x10\x3d\xbb\xbf\x62\x90\x1e\x14\x42\x72\x2f\xf4\x07\xf4\x22\xa1\xf4\x1e\x16\x3b\x9a\x25\xb8\x6b\xd6\xb5\x10\x8c\xff\xbd\xb3\x5f\xba\xb1\x04\xda\xdb\x38\x6f\xdf\xf3\xcd\xcc\x9b\x26\x78\xe9\xb9\x39\xc3\xeb\x1b\xec\x8f\xb6\x98\x67\x66\x9b\xa2\xeb\x95\x36\x83\xeb\xbf\x87\x3a\x40\x52\xc9\xcf\x5b\x8f\xc3\x31\xf2\xaa\xb4\x41\x8f\x7a\x5e\x5f\x78\x8b\xb0\x88\xdf\x41\xf2\x0e\x2d\xc4\xbc\x2e\x14\x2c\x23\x54\x34\x20\x15\x7d\xe0\x35\x3c\xcc\x2f\xd8\xaa\xbd\x50\x87\xe1\x36\x18\xec\xf2\xd2\x72\xb2\x6c\xdb\x76\x64\x94\x5f\x4e\xf1\x2f\x42\x07\x63\xfd\x45\xb9\xd3\x13\xf8\x97\xac\xe6\x92\xc6\x08\xbb\xa8\x68\x84\x5d\xfc\x88\xb3\x2f\x7b\x72\xc2\xd6\x89\x9f\x2f\xe1\x3c\x50\xbc\xbb\xf8\xd8\xfd\xf6\x09\xbe\x5a\x29\x19\x6b\x46\x59\x83\x90\xc2\x14\x25\x4c\xa9\xb5\x5e\x09\x69\x50\xfb\xf3\xf9\x7a\xf5\xb2\xa0\x8f\x7a\x29\x5f\x63\x43\x33\x0d\x6a\xd4\x35\x3a\x11\x77\xc8\x45\x81\x60\xd4\x28\x09\xdb\x9e\x00\xf2\x0f\x6c\x05\x95\xda\x32\xf2\x60\x3c\x59\x47\x61\x0d\x10\x9f\x52\xe5\xd3\x70\x87\xeb\xa8\x8c\x4d\xc1\x0e\x22\xe6\xf6\xb3\x05\x82\x9b\x79\x2e\x53\xd7\xe4\xfb\xff\x17\xf9\xe6\x1a\x4e\xb0\xe6\xb9\xf2\x31\xa4\xc4\x9a\xb3\x90\x2d\x5b\xe4\x7f\x02\x00\x00\xff\xff\x76\x5a\x7d\xc8\x08\x03\x00\x00")

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

	info := bindata_file_info{name: "types.tmpl", size: 776, mode: os.FileMode(420), modTime: time.Unix(1434819787, 0)}
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
	"cmd_main.tmpl":     cmd_main_tmpl,
	"cmd_types.tmpl":    cmd_types_tmpl,
	"cmd_validate.tmpl": cmd_validate_tmpl,
	"global.tmpl":       global_tmpl,
	"main.tmpl":         main_tmpl,
	"struct.tmpl":       struct_tmpl,
	"types.tmpl":        types_tmpl,
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
	"cmd_main.tmpl":     &_bintree_t{cmd_main_tmpl, map[string]*_bintree_t{}},
	"cmd_types.tmpl":    &_bintree_t{cmd_types_tmpl, map[string]*_bintree_t{}},
	"cmd_validate.tmpl": &_bintree_t{cmd_validate_tmpl, map[string]*_bintree_t{}},
	"global.tmpl":       &_bintree_t{global_tmpl, map[string]*_bintree_t{}},
	"main.tmpl":         &_bintree_t{main_tmpl, map[string]*_bintree_t{}},
	"struct.tmpl":       &_bintree_t{struct_tmpl, map[string]*_bintree_t{}},
	"types.tmpl":        &_bintree_t{types_tmpl, map[string]*_bintree_t{}},
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
