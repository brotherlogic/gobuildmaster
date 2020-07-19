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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdb\x8a\xdb\x30\x10\x7d\x6e\xbe\x42\xe4\x03\xba\xf4\xd5\x94\xc0\x96\x0d\xa5\x74\x49\x21\xbb\x2d\x7d\x0b\xb2\x3c\x75\xb4\x91\x35\xee\x68\x14\xe2\xbf\x2f\xce\xad\xb1\x9d\x8b\xb4\xda\xd2\x40\x08\x8c\xe7\x64\xce\x19\x9d\x19\xd9\x6a\xcb\x60\xd9\x89\x8f\x23\x21\x5e\x30\xdf\xfe\x0a\x61\x65\x05\x99\x18\xe7\x5e\x9b\xc2\x01\xad\x81\xc6\xdb\x78\x89\x8b\x5a\xf2\x32\x13\xe3\x52\xf3\xd2\xe7\xef\x15\x56\x77\x39\x21\x2f\x81\x0c\x96\x5a\xdd\x0d\x20\xb5\x24\xd6\xd2\x2c\x72\x44\x76\x4c\xb2\xce\x04\x93\x87\xed\x33\x82\xdf\x5e\x13\x54\x47\x06\xed\x47\x49\x86\x12\xa9\xc9\xc4\xc3\x97\xa7\xaf\xfb\x60\x4d\x58\x03\xb1\x06\x97\x89\xb1\x53\x24\x59\x2d\x77\xff\x3f\x19\xed\xbe\x04\x85\xb7\x85\xb4\xaa\xc9\xc4\x7c\xfa\xf0\x7d\xf6\x70\x3f\x7b\x1e\x4d\x46\x17\x14\xde\x56\xf2\x4b\x1b\x50\x58\xeb\x83\x90\x7d\x4f\xfa\xe1\x6b\xfa\x72\x02\xb9\x42\xcf\xc7\x50\x9f\xe8\xe7\xc7\x6f\x9f\xee\x1f\x47\x42\x58\x5c\x54\xd2\x31\xd0\x3e\xf5\x22\xef\x3d\x8b\x35\x90\xd3\x68\x99\xa4\x5a\x85\x1f\xce\x39\xd4\xff\xe1\xbf\x82\xc6\x31\x12\x84\x32\xef\xe6\xdf\xf0\xcd\xd3\x74\xfe\x63\x3a\x3f\xe7\x1c\xf2\xd6\x1e\x84\x9f\x33\xce\x5e\xcf\x2d\xf6\x85\x64\x19\x45\xbf\x07\x78\xa5\xef\xbb\x4d\xb8\xcc\xff\x5d\x8a\xed\x77\x71\x25\xa9\xe8\xd8\xbe\x17\x6e\x0b\xc7\x8e\xdc\xa1\x79\x84\x75\x8e\x1b\xd7\x58\x15\xdc\xbe\x2e\x64\x22\x44\xfc\xc0\x1f\xc6\x97\x64\x05\x91\x63\xd3\xc7\xa4\xd4\xaf\x09\x37\x4d\x68\xe1\x93\xe4\xc1\x28\xde\x74\xd1\xf4\xe7\xf3\x74\x3e\xdb\x4e\xe7\xc0\x49\xb0\x61\x20\x2b\xcd\x82\x40\x16\x4d\xd7\x4f\x0a\xbd\xe5\x4c\x7c\xb8\x29\xc5\x60\x59\x6a\x5b\x86\x8a\xe9\xa4\xff\xf3\xcd\x7f\x65\x06\x0e\x97\x1b\x00\x45\xde\x6d\x1d\x44\x8a\x0b\x60\x03\xca\x33\x06\x57\xee\xe6\x0f\xd7\xf2\x80\x4a\xe0\x22\xab\x09\xd7\xba\xbd\x12\xc2\x9b\x30\x80\x9c\x71\xa6\xf3\x05\xbe\x05\xb9\xf6\x71\x04\xb1\x93\xf4\x84\xb7\x8e\xcb\xb7\x47\xa7\xc2\xe4\x28\x2d\x78\x60\x5a\xff\x44\xae\x9e\x1e\xe4\xd5\xbd\x24\x50\x48\x85\x42\x63\x40\xb1\x46\x1b\x5a\xff\x1c\x2e\xc5\xf8\xbb\x3a\x04\x0a\x74\xc4\xe0\x0d\x51\x49\x2b\xd8\x1b\xd3\x7a\x00\x5c\x8c\xbb\xfa\xa0\x14\x06\x04\x95\xb6\x05\x90\x0b\x3f\x86\x13\x40\x5a\xe5\xf6\x3c\x4b\xe0\x08\xe9\x7d\x4c\x7a\x7d\x87\x54\x4a\xab\x5d\x2c\x87\x2e\x2e\x85\x47\x2e\x9d\x56\x2f\x98\x07\xcf\xe1\x49\x7e\xba\x7e\x59\x14\xb1\xd2\x8f\x90\xf4\xea\xb9\x6f\x0f\x33\xae\xfc\x5f\xcc\x1b\xa8\x37\xed\x46\x0d\x7f\x71\x18\xa2\xae\x72\xf8\x13\x00\x00\xff\xff\x99\x4c\x2b\x15\xd3\x0e\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 3795, mode: os.FileMode(436), modTime: time.Unix(1595197863, 0)}
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
