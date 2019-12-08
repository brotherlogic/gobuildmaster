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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x61\x6b\xdb\x30\x10\xfd\x3c\xff\x0a\xe1\x1f\xb0\xb2\xaf\x66\x18\x3a\x62\xc6\x58\xc9\x20\xed\xc6\xbe\x05\x59\xbe\x26\x5a\x64\x9d\x77\x3a\x95\xf8\xdf\x0f\xbb\x89\x1b\xc7\xc9\x2c\xbb\x1d\x0b\x84\xc0\xf9\xdd\xe9\xe9\xdd\x5d\x9e\xad\xb6\x0c\x96\x9d\xf8\x18\x09\xf1\x0b\xf3\xf6\x57\x08\x2b\x4b\x48\x44\x9c\x7b\x6d\x0a\x07\xf4\x04\x14\xb7\xf1\x0d\xae\x2b\xc9\xdb\x44\xc4\x1b\xcd\x5b\x9f\xbf\x57\x58\xde\xe4\x84\xbc\x05\x32\xb8\xd1\xea\x66\x90\x52\x49\x62\x2d\xcd\x3a\x47\x64\xc7\x24\xab\x44\x30\x79\x68\x9f\x11\xfc\xf6\x9a\xa0\xec\x18\x34\x1f\x25\x19\x36\x48\x75\x22\x16\x5f\xee\xbf\x1e\x82\x15\x61\x05\xc4\x1a\x5c\x22\x62\xa7\x48\xb2\xda\x3e\xd7\x4f\x43\x2a\xdd\x67\xab\x1f\xd9\xea\x52\x2d\x96\x6e\x67\xb4\xe3\x97\x62\xcd\x97\xa0\xf0\xb6\x90\x56\xd5\x89\x58\x65\x8b\xef\xcb\xc5\xed\xf2\x21\x4a\xa3\x2b\x72\x8d\xcb\xf2\xa8\x0d\x28\xac\xf4\x51\x95\x83\xc0\xe7\xe1\xbf\x89\x95\x13\xc8\x1d\x7a\xee\x42\xe7\x44\x3f\xdf\x7d\xfb\x74\x7b\x17\x09\x61\x71\x5d\x4a\xc7\x40\x07\xe8\x55\xde\x07\x16\x4f\x40\x4e\xa3\x65\x92\x6a\x17\xde\xe9\x4b\x59\xff\x87\xbf\x34\x40\x1c\x4e\xfc\x04\x9e\x0a\x31\xbd\xdb\x7d\xd5\xa6\xad\xc7\x85\xa4\xd9\xa3\x4b\xde\xda\x63\x91\xe3\xe0\x2a\xf4\x96\x13\xf1\x61\x94\xfd\x0e\x6a\xc7\x48\x10\x4a\xbc\x8f\x7f\x53\xce\x17\x66\x60\x8c\x7d\x41\x58\xe5\xb8\x77\xb5\x55\xa1\x17\x38\x4b\x79\x4d\xe7\x2b\xc2\x7d\x1d\x7a\xee\x09\x78\x44\xb5\xec\xe7\x43\xb6\x5a\xb6\x1b\x30\xd0\x0d\xf6\x0c\x64\xa5\x59\x13\xc8\xa2\x9e\xd9\xf3\x67\x8a\x04\x0a\xf4\x84\x91\x1d\x66\xbd\x4a\x3b\x6f\x4c\xa3\x03\xb8\x09\x1b\x3b\x48\x9a\xc5\x20\xf4\xaa\x4a\x52\x11\x0f\x75\xeb\xc2\xad\xe3\xcc\xbb\x7d\x89\x56\x33\x06\xdf\xbb\x07\x7f\x2b\xb3\xbc\xea\x6f\xef\x46\xe9\xc3\x1e\x94\x9f\xc0\xbf\x8f\x1f\x1a\xc0\xa0\x8b\x81\xeb\xdf\x78\x4e\xf8\xf4\xbc\xa0\x67\x58\x7b\xb7\xf2\xcd\xe3\xf0\x81\x3d\x85\xcf\xfe\xb3\xec\x55\x49\x3b\xc1\x82\xf7\xfd\x91\x64\x09\x13\x5d\xfd\x3c\xe7\x1f\x6d\x5a\x0e\x40\xa7\x26\x78\x7c\xe1\xec\x85\xd3\xe8\xda\xc1\x7f\x02\x00\x00\xff\xff\x3f\x0e\x2c\x9e\xb7\x0a\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 2743, mode: os.FileMode(436), modTime: time.Unix(1575813591, 0)}
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
