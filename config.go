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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xed\x8a\xdb\x30\x10\xfc\x5d\x3f\x85\xf0\x03\xf4\xe8\x5f\x53\x0c\x57\x62\x4a\xe9\x91\x42\xee\x5a\xfa\x2f\xc8\xf2\x5e\xa2\xc6\xd6\xaa\xab\xd5\x11\xbf\x7d\x71\x3e\xdc\x38\x4e\x6a\xd9\xd7\x70\x81\x10\x58\xef\x4a\x33\xb3\xb3\x59\x1b\x6d\x18\x0c\x3b\xf1\x31\x12\xe2\x17\xe6\xbb\x5f\x21\x8c\xac\x20\x11\x71\xee\x75\x59\x38\xa0\x17\xa0\x78\x17\x5f\xe1\xd2\x4a\x5e\x27\x22\x5e\x69\x5e\xfb\xfc\xbd\xc2\xea\x2e\x27\xe4\x35\x50\x89\x2b\xad\xee\x7a\x25\x56\x12\x6b\x59\x2e\x73\x44\x76\x4c\xd2\x26\x82\xc9\xc3\xee\x19\xc1\x6f\xaf\x09\xaa\x16\x41\xf3\x51\x92\x61\x85\x54\x27\x62\xf6\xe5\xf1\xeb\x21\x68\x09\x2d\x10\x6b\x70\x89\x88\x9d\x22\xc9\x6a\xbd\x3f\x3f\x8d\xf6\x5f\x82\xc2\x9b\x42\x1a\x55\x27\x62\x91\xcd\xbe\xcf\x67\xf7\xf3\xa7\x28\x8d\xae\x30\x1c\x66\xf2\xac\x4b\x50\x68\xf5\x91\xc8\x41\x93\xf3\xf0\xbf\xf8\xe5\x04\x72\x83\x9e\xdb\xd0\x39\xd0\xcf\x0f\xdf\x3e\xdd\x3f\x44\x42\x18\x5c\x56\xd2\x31\xd0\x21\xf5\x2a\xee\x03\x8a\x17\x20\xa7\xd1\x30\x49\xb5\x09\x6f\xce\xa5\xaa\xb7\xc1\x2f\x4b\x20\x0e\x07\x7e\x92\x9e\x0a\x31\xbe\xdb\x5d\xd5\xc6\x39\xfa\x42\xd1\x80\x6f\x1f\xb3\xc5\x8f\x6c\x71\xc9\xb9\xe4\x8d\x39\x1e\x72\x34\xae\x42\x6f\x38\x11\x1f\x06\xd1\x6f\xa0\x76\x8c\x04\xa1\xc0\xbb\xf9\xff\x15\xf3\x05\x0f\x0c\xa1\x2f\x08\x6d\x8e\x5b\x57\x1b\x15\x4a\xe0\xac\xe4\x35\x9d\xb7\x84\xdb\x3a\xf4\xde\x93\xe4\x01\xd5\xb2\x9f\x4f\xd9\x62\xbe\x9b\x80\x9e\x6e\xb0\x65\x20\x23\xcb\x25\x81\x2c\xea\x89\x3d\xdf\x43\x24\x50\xa0\x47\x58\xb6\x5f\xf5\x2a\xed\x7c\x59\x36\x3a\x80\x1b\x31\xb1\xbd\xa2\x49\x08\x42\xa9\x2a\x49\x45\xdc\xd7\xad\x0d\x37\x9a\x4f\x64\x5f\xa1\xd1\x8c\xc1\xbc\x3b\xe9\x37\xdf\x6f\xef\x06\xe1\xc3\x16\x94\x1f\x81\xbf\x9b\xdf\x5f\x00\xbd\x2e\x06\x8e\x7f\xb3\x73\xc2\xdd\xf3\x37\x7b\xc2\x6a\x6f\x47\xbe\x79\x1c\x6e\xd8\xd3\xf4\xc9\x7f\x96\x9d\x53\xd2\x56\xb0\xe0\x79\x7f\x26\x59\xc1\xc8\xad\x7e\x5e\x73\xa3\x49\xcb\x01\xe8\x74\x09\x1e\xdf\x11\x3b\xe1\x34\xba\x76\xf1\x9f\x00\x00\x00\xff\xff\xa9\xf2\xa8\xdb\x6a\x0a\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 2666, mode: os.FileMode(436), modTime: time.Unix(1576269881, 0)}
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
