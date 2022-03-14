// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// config.pb
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\x6d\x6f\xe3\x36\x0c\xfe\x7c\xfd\x15\x42\x7e\xc0\x0e\x1b\x86\x7d\x08\x86\x02\xb7\x35\x18\x8a\x1d\xba\x43\xae\x1b\xf6\x2d\xa0\x65\xd6\xd1\x45\x11\x7d\x94\x9c\x35\xfb\xf5\x83\xf3\xe2\xf3\x4b\x52\x8b\x95\xd1\x15\x28\x02\xc8\x7c\xfc\x90\x14\xf5\x88\x96\x9c\x71\x01\x5d\xf0\xea\xe7\x1b\xa5\xbe\x50\x76\xf8\x55\xca\xc1\x16\xe7\x6a\x96\x55\xc6\xe6\x1e\x79\x87\x3c\x3b\x8c\x17\xb4\x2a\x21\xac\xe7\x6a\x56\x98\xb0\xae\xb2\xef\x34\x6d\xdf\x67\x4c\x61\x8d\x6c\xa9\x30\xfa\xfd\x00\x52\x02\x07\x03\x76\x95\x11\x05\x1f\x18\xca\xb9\x0a\x5c\xe1\xe1\x59\xc6\x08\x1b\xaa\x42\x6b\x88\xf1\x6b\x65\x18\xb7\x8d\x53\xf5\x9f\x86\x80\x05\xf1\x7e\xae\xee\xee\x3f\xff\x7e\x1a\x2c\x99\x4a\xe4\x60\xd0\xcf\xd5\xcc\x6b\x86\xa0\xd7\x47\xca\xdb\x9b\xe3\x3f\x63\x5e\xb9\x1c\x9c\xde\xcf\xd5\x72\x71\xf7\xe7\xc3\xdd\x87\x87\xc7\xce\xf8\x4f\x3f\xb6\x9f\xdc\xde\x5c\x49\xc7\x78\xd8\x4f\xc6\xa2\xa6\xd2\x9c\xa3\x3e\x25\xb0\x3f\x2c\x4a\x46\x3f\x84\xdf\x3e\xfe\xf1\xcb\x87\x8f\x37\x4a\x39\x5a\x6d\xc1\x07\xe4\x93\xe9\x55\xbf\x4f\x5e\xec\x90\xbd\x21\x17\x18\xf4\x26\x7e\x26\x2f\xa1\xfe\x1f\xff\x37\xb8\xf7\x81\x18\x63\x3d\xef\xda\xcb\x8b\xec\xf3\x62\xf9\xd7\x62\x79\xa9\xcc\x9e\x18\xb6\x18\xd6\x8c\x78\xbd\xd2\x4e\x61\x8e\x05\xf5\x64\xa9\x2c\xf7\xb1\x21\xb5\xad\x27\x0d\xa8\xf5\xde\x84\x60\x72\x08\x20\x9a\xa2\x1e\xe0\x95\xab\xbe\xf7\x96\xeb\x01\xbc\x1b\x8f\x40\xe6\x7e\xcb\xfa\x6d\x04\x6e\x8a\x50\xc7\xe3\x3a\x8e\x6b\xe0\xbc\x23\x63\xbd\xe1\x9a\xf8\x8a\xb8\x8e\xa6\x99\xa9\xcc\xe8\xd9\xef\x9d\x8e\xce\x75\x17\x72\xab\xd4\x55\x69\x1f\x63\x5f\x23\xd8\xb0\xd6\x6b\x94\xe8\xe0\x00\x94\xe2\x01\xe4\xbb\xfa\xa1\x68\x4b\xed\x63\x52\xf8\x8f\x02\x26\xdb\x08\xfa\x98\x14\xfe\x92\xe9\x39\x5a\xf4\x5a\xc6\xf2\x85\xb4\xf8\xfb\x71\xb1\x7c\x38\xec\x37\x83\xc5\x84\xcf\x01\xd9\x81\x5d\x31\x42\xde\x53\x3f\x4d\x95\x0b\x73\xf5\xfd\x68\x28\x96\x8a\xc2\xb8\x22\x36\x98\x8e\xf9\x1b\x36\x3e\xad\x5d\x37\xac\xc8\xad\x8e\x75\x34\x57\x33\xce\x8d\x2f\x2d\xec\x67\xe3\xe2\x98\x21\xb2\xb0\x0d\xec\x20\x52\x2a\x06\x9f\x51\x57\x81\xa2\x99\xbb\xf6\xc3\xa6\x64\xe0\x4a\xe4\x16\x57\x32\xed\x4c\xdd\x10\xc5\x27\x61\x00\xb9\x30\xed\x2f\xed\x1e\xbe\xca\x69\x0a\xc7\xeb\xc7\x02\xa7\x5b\xe6\xf2\x3a\x5d\x2e\x7e\x5d\xdc\x7f\x7a\x5c\x7d\x5a\xde\x3f\x3c\x9e\x7a\x8e\xdb\xc4\x12\x38\x17\x6a\xec\x86\xd1\x36\x8f\x28\x80\x78\x47\xea\xaa\x16\x8a\xe7\x00\x32\x9a\xd2\x77\x82\xf6\xed\x5f\x64\xfa\x61\xd6\x4b\xf2\x41\xc1\x54\x8c\x84\x31\x6a\xe2\x5c\x93\xb5\xa8\x83\x21\x17\x1b\xd4\x25\x5c\x4a\x5e\x8f\x3c\x8c\x1a\x8d\x40\x63\x86\xa8\xa4\x9d\xa9\xb2\xb6\x9e\x08\xf4\x92\xc5\xd2\x07\xa5\x78\xc0\xb8\x35\x2e\x47\xf6\xf1\xd3\xd0\x02\xa4\x31\xd7\xf3\x59\x60\x10\x84\xde\xc7\xa4\xf3\x7b\xe2\x02\x9c\xf1\x52\x1f\xba\xb8\xa4\xf5\x0d\xde\xe8\x2f\x94\x45\x2f\xee\x96\x7d\x7a\xfc\x90\xe7\xd2\xd0\x1b\x48\x3a\x7b\x56\xd5\x93\x29\xa3\xff\x86\x99\x20\x7a\x5b\xab\x5a\x7c\x3f\x35\x44\xa5\xfb\xb0\xad\x5b\x29\xe9\x1c\xb4\x40\x13\x78\x40\x02\x01\xec\x41\xd2\xd9\x85\xad\xc2\x00\x34\x85\x07\xa4\xd1\x0b\x14\xb0\x07\x9a\x40\x85\xc0\xa2\x90\xbf\x81\x4c\xc0\x1e\x20\x48\xd9\xcf\x90\x14\x76\x7d\x4e\x62\x7c\xa3\x3d\x80\xbc\xf2\x0b\x86\xc1\xe4\x31\xe7\xb6\x71\x09\xfc\x07\x9c\x34\x81\x0d\x24\xbd\x85\x09\xe0\x37\xd1\xec\x3d\x48\x5a\x87\x5c\x98\x00\xb6\x8e\xc4\x1a\x1f\xad\xe2\x17\x60\x13\x94\xb0\x26\x16\xaf\xa0\x06\x93\xc2\x4f\x5c\x08\x15\xac\x8b\x48\xe1\xae\xdf\x17\xdd\x3f\x37\xc6\x29\x8c\x01\xb7\x65\x2c\xe1\xd9\x36\x85\x6f\x83\xfb\x2d\x94\x65\x7c\x72\x3b\x80\x14\xe6\xc3\x02\x95\xd4\x75\x07\x90\xd6\x10\xea\x4d\x15\x9d\xe5\xb6\xf5\x34\x62\x98\xbc\x18\x77\x60\x4d\x0e\x82\xe3\x93\x0b\xb0\xb4\xf3\x3e\xe3\x42\xc6\x26\x2f\xa2\xcf\xd6\x7b\x90\x14\x76\x1f\x48\xc6\xdd\x01\xa4\xc9\x81\xde\xc4\xab\xc1\xd1\x36\x85\xef\x6b\x85\x55\x74\x94\x8d\x71\x7a\x7d\x3d\x81\xa3\x4a\xf8\xd5\xf0\x0d\x93\xce\xaf\x2d\x82\xe0\x44\x6e\x00\x4a\x3a\x99\xac\x6c\x3c\x73\x63\x3c\xcd\x77\xb2\xfc\x0b\x39\xb9\xc5\x00\xbf\xce\xa8\xb9\x19\x8a\xb9\xe2\xeb\x00\xa6\xbc\xb5\x2c\x91\x3d\x39\xb0\x2f\x29\x65\xe4\xe9\xa8\x47\x5d\x31\x8a\xae\x25\x06\x90\x37\xba\x9c\x78\x69\xf2\xfe\x0b\x00\x00\xff\xff\xc0\x9e\x21\xaf\xc8\x21\x00\x00")

func configPbBytes() ([]byte, error) {
	return bindataRead(
		_configPb,
		"config.pb",
	)
}

func configPb() (*asset, error) {
	bytes, err := configPbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.pb", size: 8648, mode: os.FileMode(436), modTime: time.Unix(1647219272, 0)}
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
	"config.pb": configPb,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"config.pb": &bintree{configPb, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
