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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x93\x51\x6b\xea\x40\x10\x85\x9f\x6f\x7e\xc5\x90\x1f\x70\x7d\x0f\x17\xc1\x8b\x72\xb9\x54\x2c\x68\xdb\x57\x99\x6c\xa6\xc9\xd6\x64\x27\x9d\x9d\x95\xe6\xdf\x17\x35\x0a\x49\x4d\xd5\xf6\xa1\x81\x10\x38\xd9\x49\xbe\x73\xf6\xac\xb3\x4e\xc9\xa9\x87\x3f\x11\xc0\x0b\xa7\xfb\x27\x80\xc3\x8a\x12\x88\xd3\x60\xcb\xcc\x93\x6c\x49\xe2\xbd\x9e\xf3\xba\x46\x2d\x12\x88\x73\xab\x45\x48\x7f\x1b\xae\x46\xa9\xb0\x16\x24\x25\xe7\xd6\x8c\x3e\x8c\xd4\x28\x6a\xb1\x5c\xa7\xcc\xea\x55\xb0\x4e\x40\x25\xd0\xfe\x9d\xd0\x6b\xb0\x42\xd5\x89\x60\x77\x19\x54\xca\x59\x9a\x04\xa6\xff\x57\x77\xad\x58\x0b\xd7\x24\x6a\xc9\x27\x10\x7b\x23\xa8\xa6\x38\x7c\x7f\x1c\x1d\x6e\xa1\x2c\xb8\x0c\x9d\x69\x12\x58\xce\xa6\x8f\x8b\xe9\x64\xf1\x10\x8d\xa3\x01\x87\x97\x9d\x3c\xdb\x92\x0c\xd7\xf6\x68\xa4\xcd\xa4\x2f\x7f\xe6\x2f\x15\xc2\x0d\x07\x3d\x49\x7d\xd0\x7f\xf3\xfb\xbf\x93\x79\x04\xe0\x78\x5d\xa1\x57\x92\x76\xe9\x20\x77\x4b\xb1\x25\xf1\x96\x9d\x0a\x9a\xcd\xf5\x9b\x73\x6e\xea\x67\xf8\x37\xd4\x78\x65\xa1\x6b\xc9\xbb\xeb\x2f\xf4\x66\x35\x5b\x3e\xcd\x96\xe7\x9a\x23\xc1\xb9\xa3\xf1\x73\xc5\x69\xfd\x5c\xa2\xcf\x50\xf1\x26\xfc\xde\xc0\x17\x7b\xdf\x0d\x61\x98\xff\xd7\x77\x6a\x7f\xd0\x0d\x4a\xd6\xa9\x7d\x4f\xde\xfd\xf8\xd6\x23\x77\x0c\x4f\xb8\x4e\xf9\xcd\x37\xce\x5c\x1d\x5f\x77\x64\x0c\x30\x7c\xe0\xdf\x03\x00\x00\xff\xff\xe2\x88\x37\x88\xd2\x04\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 1234, mode: os.FileMode(420), modTime: time.Unix(1593480881, 0)}
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
