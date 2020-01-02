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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xed\x6a\xdb\x40\x10\xfc\x5d\x3d\xc5\xa1\x07\x68\xc8\x5f\x51\x0c\x29\x31\xa5\x34\xb8\xe0\xa4\xa5\xff\xcc\xe9\xb4\x91\xaf\x96\x6e\xd5\xbd\x55\x6a\xbd\x7d\xb1\x65\xab\xfa\x4a\x74\x3a\xd7\x34\x10\x0c\x7b\x3b\xa7\x99\xd9\x5d\xad\x8c\x36\x0c\x86\xad\xf8\x10\x08\xf1\x13\xe3\xe3\xaf\x10\x46\xe6\x10\x89\x30\x2e\x75\x96\x58\xa0\x17\xa0\xf0\x18\x4f\x71\x53\x48\xde\x46\x22\x4c\x35\x6f\xcb\xf8\xbd\xc2\xfc\x26\x26\xe4\x2d\x50\x86\xa9\x56\x37\x03\x48\x21\x89\xb5\xcc\x36\x31\x22\x5b\x26\x59\x44\x82\xa9\x84\xe3\x19\xc1\xaf\x52\x13\xe4\x0d\x83\xc3\x9f\x92\x0c\x29\x52\x15\x89\xfb\xcf\x8f\x5f\x4e\xc1\x82\xb0\x00\x62\x0d\x36\x12\xa1\x55\x24\x59\x6d\xeb\xfb\x17\x41\xfd\x4f\x90\x94\x26\x91\x46\x55\x91\x58\x2f\xef\xbf\xad\xee\xef\x56\x4f\xc1\x22\x78\x45\xe1\xb4\x92\x67\x9d\x81\xc2\x42\x9f\x85\x9c\x3c\xe9\x87\xdf\xd2\x17\x13\xc8\x1d\x96\xdc\x84\xfa\x44\x3f\x3d\x7c\xfd\x78\xf7\x10\x08\x61\x70\x93\x4b\xcb\x40\xa7\xd4\x57\x79\x9f\x58\xbc\x00\x59\x8d\x86\x49\xaa\x9d\x7b\x71\xc6\x50\xff\x87\xbf\xcc\x80\xd8\x9d\x78\x2b\x7d\x21\xc4\xfc\x6a\x77\x5d\x9b\xd7\xd1\x23\xa0\x89\xbe\x7d\x5c\xae\xbf\x2f\xd7\x63\x9d\x4b\xa5\x31\xe7\x4b\xce\x8d\xab\xb0\x34\x1c\x89\xdb\x49\xf6\x3b\xa8\x2c\x23\x81\x2b\xf1\x6e\xfe\x3f\xe5\x3c\xd2\x03\x53\xec\x13\xc2\x22\xc6\xbd\xad\x8c\x72\x15\xd0\x83\x5c\x52\xf9\x82\x70\x5f\xb9\x3e\xb7\x95\x3c\xe1\xda\xf2\xc7\xd3\x72\xbd\x3a\x4e\xc0\xc0\x37\xd8\x33\x90\x91\xd9\x86\x40\x26\x95\x67\xcd\x6b\x8a\x04\x0a\xf4\x8c\x96\x1d\xa2\x2e\xf2\xae\xcc\xb2\x83\x0f\x60\x67\x4c\xec\x00\xe4\xc5\xc0\x55\xaa\x92\x94\x84\x43\xdf\x9a\xf0\xc1\x73\x4f\xf5\x39\x1a\xcd\xe8\xac\xbb\x93\x7e\xf5\xfd\xf6\x6e\x92\x3e\xec\x41\x95\x33\xf8\x77\xf3\x87\x0b\x60\x50\x45\xc7\xf1\x3f\xec\x1c\xf7\xee\xf9\x9b\xed\xb1\xda\x9b\x91\x3f\x1c\xbb\x37\x6c\x3b\xdd\xfb\x65\xd9\xb9\x65\xd1\x18\xe6\x3c\xef\xcf\x24\x73\x98\xb9\xd5\xfb\x98\x2b\x4d\x5a\x0c\x40\xed\x25\x78\xfe\x46\xec\x84\x17\x81\x67\xb5\xce\xcb\x2a\x96\x6a\x57\x16\x73\x57\x5c\x1b\xe5\x39\x74\x24\x75\x72\xeb\x5d\x37\x02\x85\x94\x28\xcc\x32\x50\xac\xd1\xb8\x0a\x18\xc3\xf9\xb7\x7c\x7d\x9b\x4c\x12\xf7\xe6\xe9\x41\x2e\xd9\x13\xf5\x55\x29\xf0\x8c\xa9\xeb\x63\xae\xf2\x25\x5f\x3f\xe4\xb7\x34\x6c\xc3\x21\xe1\x26\xee\xef\xfb\xf1\x8a\x4c\x5b\x76\x95\xdd\x01\x5c\x69\x5e\x09\x72\x6d\x12\xa0\xbe\xe6\x76\xd4\x67\x2f\xba\xfa\x6d\x91\x52\x69\xb4\xed\xbd\x2f\x46\x0f\xdf\xe2\xf1\x27\x00\x00\xff\xff\xf1\xc0\xa9\xa5\xa3\x0e\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 3747, mode: os.FileMode(436), modTime: time.Unix(1577995468, 0)}
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
