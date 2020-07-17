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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdb\x8a\xdb\x30\x10\x7d\x6e\xbe\x42\xe4\x03\xba\xf4\xd5\x94\xc0\x96\x84\x52\xba\xa4\x90\xdd\x96\xbe\x05\x59\x9a\x3a\xda\xd8\x1a\x77\x34\x0a\xf1\xdf\x17\xe7\x46\x6c\xe7\x22\xad\x5a\x1a\x08\x81\xf1\x9c\xcc\x39\xa3\x33\x23\x5b\x63\x19\x2c\x3b\xf1\x71\x24\xc4\x2b\xe6\xbb\x5f\x21\xac\xac\x20\x13\xe3\xdc\x9b\x52\x3b\xa0\x0d\xd0\x78\x17\x2f\x70\x59\x4b\x5e\x65\x62\x5c\x18\x5e\xf9\xfc\xbd\xc2\xea\x21\x27\xe4\x15\x50\x89\x85\x51\x0f\x03\x48\x2d\x89\x8d\x2c\x97\x39\x22\x3b\x26\x59\x67\x82\xc9\xc3\xee\x19\xc1\x6f\x6f\x08\xaa\x13\x83\xf6\xa3\x24\x43\x81\xd4\x64\x62\xfa\xe5\xf9\xeb\x21\x58\x13\xd6\x40\x6c\xc0\x65\x62\xec\x14\x49\x56\xab\xfd\xff\x4f\x46\xfb\x2f\x81\xf6\x56\x4b\xab\x9a\x4c\x2c\x66\xd3\xef\xf3\xe9\xe3\xfc\x65\x34\x19\x5d\x51\x78\x5f\xc9\x2f\x53\x82\xc2\xda\x1c\x85\x1c\x7a\xd2\x0f\xdf\xd2\x97\x13\xc8\x35\x7a\x3e\x85\xfa\x44\x3f\x3f\x7d\xfb\xf4\xf8\x34\x12\xc2\xe2\xb2\x92\x8e\x81\x0e\xa9\x57\x79\x1f\x58\x6c\x80\x9c\x41\xcb\x24\xd5\x3a\xfc\x70\x2e\xa1\xfe\x0f\xff\x35\x34\x8e\x91\x20\x94\x79\x37\xff\x8e\x6f\x9e\x67\x8b\x1f\xb3\xc5\x25\xe7\x90\xb7\xf6\x28\xfc\x92\x71\x0e\x7a\xee\xb1\xd7\x92\x65\x14\xfd\x1e\xe0\x8d\xbe\xef\x36\xe1\x3a\xff\x77\x29\xb6\xdf\xc7\x95\x24\xdd\xb1\x7d\x2f\xdc\x16\x8e\x1d\xb9\x63\xf3\x08\xeb\x1c\xb7\xae\xb1\x2a\xb8\x7d\x5d\xc8\x44\x88\xf8\x81\x3f\x8e\x2f\xc9\x0a\x22\xc7\xa6\x8f\x49\xa9\x5f\x13\x6e\x9b\xd0\xc2\x67\xc9\x77\x2c\x33\xfb\xf9\x32\x5b\xcc\x77\xa3\x38\xb0\x0d\x6c\x19\xc8\xca\x72\x49\x20\x75\xd3\x35\x8f\x42\x6f\x39\x13\x1f\xee\xf2\x2e\xb1\x28\x8c\x2d\x42\x99\x77\xd2\xff\xf9\x9a\xbf\x61\xf8\xe3\x4d\x06\x40\x91\x17\x59\x07\x91\x72\xe4\xb0\x05\xe5\x19\x83\x2b\x77\xf3\x87\x3b\x78\x40\x25\x70\x6b\xd5\x84\x1b\xd3\xee\xff\xf0\x26\x0c\x20\x03\x36\x42\x38\xaf\xf1\x6f\x90\x6b\x1f\x47\x10\x3b\x4b\x4f\x78\xc5\xb8\x7e\x55\x74\x2a\x4c\x4e\xd2\x82\x07\xa6\xf5\x4f\xe4\x9e\xe9\x41\xde\xdc\x4b\x02\x85\xa4\x15\x96\x25\x28\x36\x68\x43\xeb\x5f\xc2\xa5\x18\x7f\x5f\x87\x40\x81\x89\x18\xbc\x21\x2a\x69\xdf\xfa\xb2\x6c\x3d\x00\x2e\xc6\x5d\x7d\x50\x0a\x03\x82\xca\x58\x0d\xe4\xc2\x8f\xe1\x0c\x90\x56\xb9\x3d\xcf\x02\x38\x42\x7a\x1f\x93\x5e\xdf\x21\x15\xd2\x1a\x17\xcb\xa1\x8b\x4b\xe1\x91\x4b\x67\xd4\x2b\xe6\xc1\x73\x78\x96\x9f\xae\x5f\x6a\x1d\x2b\xfd\x04\xb9\x59\xfd\x4f\x00\x00\x00\xff\xff\x80\xec\x86\xfa\xbc\x0d\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 3516, mode: os.FileMode(436), modTime: time.Unix(1594958863, 0)}
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
