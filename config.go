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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\xdd\x6e\x1b\x37\x13\xbd\x8e\x9f\x82\xd0\x03\x24\xc8\x07\xe3\xbb\x10\x0a\x03\x49\x2d\xb4\x46\x53\xc7\x90\xd5\xa2\xbd\x12\x66\xb9\xe3\x15\x23\x2e\x67\x33\xe4\xaa\x56\x9f\xbe\x58\xd9\x92\xf7\x47\xb2\x38\xe1\x36\xd5\x8d\x00\xee\x9c\x3d\x33\xc3\xe1\xe1\x2c\xe9\x8c\x0b\xe8\x82\x57\x3f\x5c\x28\xf5\x85\xb2\xdd\xbf\x52\x0e\x4a\x9c\xaa\x49\x56\x1b\x9b\x7b\xe4\x0d\xf2\x64\x37\x5e\xd0\xb2\x82\xb0\x9a\xaa\x49\x61\xc2\xaa\xce\xde\x6a\x2a\xdf\x65\x4c\x61\x85\x6c\xa9\x30\xfa\xdd\x00\x52\x01\x07\x03\x76\x99\x11\x05\x1f\x18\xaa\xa9\x0a\x5c\xe3\xee\x59\xc6\x08\x6b\xaa\x43\x6b\x88\xf1\x6b\x6d\x18\xcb\x83\x53\xcd\x4f\x43\xc0\x82\x78\x3b\x55\xd7\x37\xf7\xbf\x3c\x0f\x56\x4c\x15\x72\x30\xe8\xa7\x6a\xe2\x35\x43\xd0\xab\x27\xca\xab\x98\x37\xfd\xfc\xf9\x7e\xb1\x5c\xfc\x79\x37\x3b\xf6\xba\x39\xf8\x2a\x43\xe6\xad\xba\x33\xea\x52\xfd\x4a\x39\x5a\xf5\x51\xcd\x71\xa3\xde\xbf\xbd\x6c\x58\x1a\x0e\xc6\xbc\x76\x39\x38\xbd\x9d\xaa\xf9\xec\xfa\xb7\xdb\xeb\x0f\xb7\x8b\xce\xf8\xff\x2f\xdb\x4f\xae\x2e\x4e\xa4\xfb\x7c\x5a\x1f\x8c\x45\x4d\x95\xd9\x67\xf5\x79\x82\xfa\xc3\xa2\x64\xf7\x43\xf8\xe9\xd3\xe7\x8f\x1f\x3e\x5d\x28\xe5\x68\x59\x82\x0f\xc8\xcf\xa6\x27\xfd\x7e\xf6\x62\x83\xec\x0d\xb9\xc0\xa0\xd7\xf1\x95\x72\x0c\xf5\xdf\xf8\xbf\xc6\xad\x0f\xc4\x18\xeb\x79\xd7\x5e\x5e\xc4\xf7\xb3\xf9\xef\xb3\xf9\xb1\xba\xd3\xb6\xf6\xe1\x7f\x2f\x55\x7c\x22\xc4\x73\x01\x3d\x58\xaa\xaa\x6d\x6c\x38\x6d\xeb\x31\x83\x69\xbf\x37\x21\x98\x1c\x02\x88\xa6\xa7\x07\xf8\x46\x45\xe9\xbd\xe5\x74\x00\x6f\xce\x47\x20\x73\xbf\x65\xfd\x7d\xc4\x73\x8c\x50\xcf\xc7\xf5\x34\xae\x81\xf3\x8e\x84\xf5\x86\x77\xda\x7d\x5c\x58\xcf\xa6\x99\xa9\xca\xe8\xd1\x6f\x9d\x8e\xce\x75\x17\x72\xa5\xd4\x49\x59\x3f\xc7\xbe\x42\xb0\x61\xa5\x57\x28\xd1\xc0\x01\x28\xc5\x03\xc8\x37\xcd\x43\xd1\x76\xdd\xc7\xa4\xf0\x3f\x30\x94\x28\xdc\x04\xfa\x98\x14\xfe\x8a\xe9\x31\x5a\xf4\x5a\xc6\xf2\x85\x34\xfb\x63\x31\x9b\xdf\xee\xf6\x9a\xc1\x62\xc2\xc7\x80\xec\xc0\x2e\x19\x21\xef\xa9\x9f\xa6\xda\x85\xa9\x7a\x7f\x36\x14\x4b\x45\x61\x5c\x11\x1b\x4c\xc7\xfc\x5f\x6c\xaa\x5e\xd9\x71\xc3\x92\xdc\xf2\xa9\x8e\xa6\x6a\xc2\xb9\xf1\x95\x85\xed\xe4\xbc\x38\x66\x88\x2c\x6c\x31\x3b\x88\x94\x8a\xc1\x47\xd4\x75\xa0\x68\xe6\xae\xfd\xb0\x21\x19\xb8\x12\xb9\xc5\x55\x4c\x1b\xd3\x34\x43\xf1\x49\x18\x40\x8e\x4c\xfb\x6b\xbb\x87\xaf\x73\x1a\xc3\xf1\xe6\xb1\xc0\xe9\x96\xb9\xbc\x4e\xe7\xb3\x1f\x67\x37\x77\x8b\xe5\xdd\xfc\xe6\x76\xf1\xdc\x73\x5c\x25\x96\xc0\xbe\x50\x63\x37\x8c\xb6\x79\x44\x01\xc4\x3b\xd2\x54\xb5\x50\x3c\x07\x90\xb3\x29\x7d\x23\x68\xdf\xfe\x46\xa6\x76\x2b\xba\x8b\x6d\xa7\x60\x2a\x46\xc2\x18\x35\x71\xae\xc9\x5a\xd4\xc1\x90\x8b\x0d\xea\x18\x2e\x25\xaf\x4f\x3c\x8c\x1a\x8d\x40\x63\x86\xa8\xa4\x9d\xa9\xb6\xb6\x99\x08\xf4\x92\xc5\xd2\x07\xa5\x78\xc0\x58\x1a\x97\x23\xfb\xf8\x69\x68\x01\xd2\x98\x9b\xf9\x2c\x30\x08\x42\xef\x63\xd2\xf9\x3d\x71\x01\xce\x78\xa9\x0f\x5d\x5c\xd2\xfa\x06\x6f\xf4\x17\xca\xa2\x17\x77\xcb\x3e\x3d\x7e\xc8\x73\x69\xe8\x07\x48\x3a\x7b\x56\x37\x93\x29\xa3\x7f\xc1\x8c\x10\xbd\x6d\x54\x2d\xbe\x9f\x1a\xa2\xd2\x7d\x28\x9b\x56\x4a\x3a\x07\x2d\xd0\x08\x1e\x90\x40\x00\x7b\x90\x74\x76\x61\xab\x30\x00\x8d\xe1\x01\x69\xf4\x02\x05\xec\x81\x46\x50\x21\xb0\x28\xe4\x3f\x40\x46\x60\x0f\x10\xa4\xec\x7b\x48\x0a\xbb\xde\x27\x31\xbe\xd1\x1e\x40\xbe\xf1\x0b\x86\xc1\xe4\xa7\x3f\x5f\xa4\x09\xfc\x0b\x9c\x34\x81\x07\x48\x7a\x0b\x13\xc0\xaf\xa3\xd9\x7b\x90\xb4\x0e\xb9\x30\x01\x6c\x13\x89\x35\x3e\x5a\xc5\x8f\xc0\x46\x28\x61\x4d\x2c\x5e\x41\x07\x4c\x0a\x3f\x71\x21\x54\xb0\x2e\x22\x85\xbb\x79\x5f\x74\xff\x7c\x30\x4e\x61\x0c\x58\x56\xb1\x84\x7b\xdb\x14\xbe\x35\x6e\x4b\xa8\xaa\xf8\xe4\x76\x00\x29\xcc\xbb\x05\x2a\xa9\xeb\x0e\x20\xad\x21\xd4\xeb\x3a\x3a\xcb\x6d\xeb\x71\xc4\x30\x79\x31\x6e\xc0\x9a\x1c\x04\xc7\x27\x47\x60\x69\xe7\x7d\xc6\x85\x8c\x4d\x5e\x44\x9f\xad\xf7\x20\x29\xec\x3e\x90\x8c\xbb\x03\x48\x93\x03\xbd\x8e\x57\x83\x27\xdb\x14\xbe\xaf\x35\xd6\xd1\x51\x1e\x8c\xd3\xeb\xeb\x01\x1c\xd5\xc2\xaf\x86\x17\x4c\x3a\xbf\xb6\x08\x82\x13\xb9\x01\x28\xe9\x64\xb2\xb6\xf1\xcc\x07\xe3\x71\xbe\x93\xe5\x5f\xc8\xc9\x2d\x06\xf8\x55\x46\x87\x9b\xa1\x98\x2b\xbe\x0e\x60\xcc\x5b\xcb\x0a\xd9\x93\x03\xfb\x9a\x52\x46\x9e\x8e\x7a\xd4\x35\xa3\xe8\x5a\x62\x00\xf9\x4e\x97\x13\x49\xfb\x98\xcb\x35\x94\x95\xf0\x10\x7f\x80\x4a\xf2\xa1\x36\x56\x70\xc0\xd1\x32\x7f\x95\xf5\x9f\x00\x00\x00\xff\xff\x7d\xd5\xfe\xf0\x1e\x23\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 8990, mode: os.FileMode(436), modTime: time.Unix(1657387656, 0)}
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
