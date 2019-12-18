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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x61\x8b\x9b\x40\x10\xfd\x5c\x7f\xc5\xe2\x0f\xe8\x71\x5f\xa5\x08\x57\x22\xa5\xf4\x48\x21\x77\x2d\xfd\x16\xd6\x75\x2e\xd9\x46\x77\xec\xec\xec\x11\xff\x7d\x31\x89\x56\x63\x52\x57\xaf\xa1\x81\x10\x98\x9d\xd9\xbc\x79\xf3\xc6\xa7\xd1\x86\xc1\xb0\x15\x1f\x02\x21\x7e\x62\x7a\xf8\x15\xc2\xc8\x02\x22\x11\xa6\x4e\xe7\x99\x05\x7a\x05\x0a\x0f\xf1\x0d\xae\x4b\xc9\xdb\x48\x84\x1b\xcd\x5b\x97\xbe\x57\x58\xdc\xa5\x84\xbc\x05\xca\x71\xa3\xd5\xdd\xa0\xa4\x94\xc4\x5a\xe6\xeb\x14\x91\x2d\x93\x2c\x23\xc1\xe4\xe0\x70\x46\xf0\xcb\x69\x82\xa2\x45\x50\x7f\x94\x64\xd8\x20\x55\x91\x58\x7c\x7e\xfa\x72\x0a\x96\x84\x25\x10\x6b\xb0\x91\x08\xad\x22\xc9\x6a\x7b\xbc\x3f\x0e\x8e\x5f\x82\xcc\x99\x4c\x1a\x55\x45\x62\x95\x2c\xbe\x2d\x17\x0f\xcb\xe7\x20\x0e\xae\x74\x38\xde\xc9\x8b\xce\x41\x61\xa9\x9b\x46\x4e\x9c\x9c\x87\xff\xd6\x5f\x4a\x20\x77\xe8\xb8\x0d\x9d\x03\xfd\xf4\xf8\xf5\xe3\xc3\x63\x20\x84\xc1\x75\x21\x2d\x03\x9d\x52\xaf\xe2\x3e\xa1\x78\x05\xb2\x1a\x0d\x93\x54\x3b\xff\xe1\x5c\xaa\xfa\x3f\xf8\x65\x0e\xc4\xfe\xc0\x3b\xe9\xb1\x10\xd3\xa7\xdd\x67\x6d\x9a\xa2\x2f\x14\x8d\xe8\xf6\x29\x59\x7d\x4f\x56\x97\x94\x4b\xce\x98\xe6\x92\x46\xb8\x0a\x9d\xe1\x48\xdc\x8f\xa2\xdf\x41\x65\x19\x09\x7c\x81\xf7\xf3\xff\x29\xe6\x0b\x1a\x18\x43\x9f\x11\x96\x29\xee\x6d\x65\x94\x6f\x03\x67\x25\x6f\x99\x7c\x49\xb8\xaf\x7c\xff\xb7\x93\x3c\xc2\x5a\xf2\xe3\x39\x59\x2d\x0f\x1b\x30\xe0\x0d\xf6\x0c\x64\x64\xbe\x26\x90\x59\x35\x73\xe6\x47\x88\x04\x0a\xf4\x04\xc9\x0e\xab\xde\xc4\x9d\xcb\xf3\x9a\x07\xb0\x13\x36\x76\x50\x34\x0b\x81\x6f\xab\x4a\x52\x16\x0e\x79\x6b\xc3\x35\xe7\x33\xbb\x2f\xd0\x68\x46\xef\xbe\x7b\xe9\x37\xf7\xb7\x77\xa3\xf0\x61\x0f\xca\x4d\xc0\xdf\xcf\x1f\x1a\xc0\x60\x8a\x9e\xeb\x5f\x7b\x8e\xbf\x7a\xfe\x64\xcf\xb0\xf6\x76\xe5\xeb\x63\x7f\xc1\x76\xd3\x67\x3f\x2c\x7b\xb7\xc4\x2d\x61\xde\xfb\xfe\x42\xb2\x80\x89\xae\x7e\x5e\x73\xa3\x4d\x4b\x01\xa8\x6b\x82\xcd\x3b\x62\x2f\x1c\x07\x33\xa7\xd5\x98\x55\x2a\xd5\xce\x95\x53\x2d\xae\x5b\x35\x73\xe9\x48\xea\xec\xfe\xfa\xdc\x7e\x07\x00\x00\xff\xff\x84\x87\x46\xa5\x29\x0b\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 2857, mode: os.FileMode(436), modTime: time.Unix(1576700116, 0)}
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
