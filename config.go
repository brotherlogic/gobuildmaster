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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x61\x8b\x9b\x40\x10\xfd\x5c\x7f\xc5\x90\x1f\xd0\xa3\x5f\xa5\x04\xae\x44\x4a\xe9\x91\x42\xee\x5a\xfa\x2d\x8c\xeb\xd4\x6c\xa3\x3b\xdb\xd9\xd9\x23\xfe\xfb\x72\x89\x1e\xd5\x4b\xa2\xf6\x28\x15\x44\x78\xce\x73\xdf\x9b\x7d\xe3\x3a\xeb\x94\x9c\x06\x78\x9f\x00\xfc\xe4\xfc\xf8\x04\x70\x58\x53\x0a\x8b\x3c\xda\xaa\x08\x24\x8f\x24\x8b\x23\x5e\xf2\xd6\xa3\xee\x52\x58\x94\x56\x77\x31\x7f\x6b\xb8\xbe\xc9\x85\x75\x47\x52\x71\x69\xcd\xcd\x0b\x8a\x47\x51\x8b\xd5\x36\x67\xd6\xa0\x82\x3e\x05\x95\x48\xc7\x77\x42\xbf\xa2\x15\xaa\x9f\x15\x3c\x5d\x06\x95\x4a\x96\x26\x85\xd5\xa7\xfb\xcf\x2d\xe8\x85\x3d\x89\x5a\x0a\x29\x2c\x82\x11\x54\xb3\x3b\x7d\x7f\x99\x9c\x6e\xa1\x22\xba\x02\x9d\x69\x52\xd8\x64\xab\xaf\xeb\xd5\xed\xfa\x21\x59\x26\x17\x1c\x8e\x3b\xf9\x61\x2b\x32\xec\x6d\x67\xa4\xed\xc9\x10\xbe\xe6\x2f\x17\xc2\x3d\x47\x7d\x86\x86\x42\x3f\xde\x7d\xf9\x70\x7b\x97\x00\x38\xde\xd6\x18\x94\xa4\x2d\xbd\xa8\xbb\x55\xf1\x48\x12\x2c\x3b\x15\x34\xfb\xe9\x9b\x73\x8e\xf5\x7f\xf4\xef\xa9\x09\xca\x42\x53\x95\xf7\xeb\x47\x72\x73\x9f\x6d\xbe\x65\x9b\x73\xc9\x91\xe8\x5c\x67\xfc\x5c\x70\x5a\x3f\x63\xea\x0b\x54\x9c\x25\x7f\x40\xf8\xcb\xdc\xf7\x9b\x70\x59\xff\x9b\xd7\xc4\xfe\x84\x1b\x94\xa2\x17\xfb\x01\xfc\xb4\xf0\xdc\x91\xeb\x9a\x27\xec\x73\x3e\x84\xc6\x99\xc9\xed\xeb\x53\x96\x00\xf3\x07\xbe\x1b\x5f\xc1\x9a\x66\x8e\xcd\x90\xf3\x9a\xf5\xbd\xf0\xa1\x99\xba\xf0\x1f\xc5\x23\x91\xc9\xbe\x3f\x64\x9b\xf5\x71\x14\x5f\xc4\x86\x0e\x4a\xe2\xb0\xda\x0a\x61\xd1\xf4\xc3\x63\x38\x3a\x4d\xe1\xdd\xa8\xee\x8a\xcb\xd2\xba\x72\xaa\xf2\x5e\xf9\x3f\xff\xcd\x5f\x09\x7c\x77\x92\x11\xc9\xcc\x83\xac\xc7\xb8\xbe\xe5\xbf\x03\x00\x00\xff\xff\x15\x55\x17\xb0\x46\x07\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 1862, mode: os.FileMode(420), modTime: time.Unix(1593735155, 0)}
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
