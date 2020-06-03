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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xed\x6a\xdb\x3c\x14\xfe\xfd\xe6\x2a\x44\x2e\xe0\x2d\xfd\x1b\x46\xa0\xa3\x61\x8c\x95\x0e\xd2\x6e\xec\x5f\x90\xe5\x53\x47\x8d\xad\xe3\x1d\x49\x5d\x72\xf7\xc3\x49\xed\xf8\xab\xb1\x24\xd7\x74\x85\x12\x90\xce\x23\x3d\xe7\xfb\xc8\x4a\x2a\x03\xca\x68\xf6\x69\xc6\xd8\x33\x46\xc7\x5f\xc6\x14\xcf\x60\xc1\xe6\xb1\xd4\x02\x5f\x80\x0e\xf3\xe3\x6a\x82\x9b\x9c\x9b\xed\x82\xcd\x13\x69\xb6\x36\xfa\x5f\x60\x76\x15\x11\x9a\x2d\x50\x8a\x89\x14\x57\x2d\x40\x44\xc0\x77\x68\xcd\x82\x19\xb2\x30\x63\x6c\xc9\xd8\x8c\x31\x82\xd8\xaa\x98\x2b\x71\x58\xb0\x2f\x77\xdf\x3f\xdf\xdc\xcd\x96\xb3\xcb\x4c\x22\x2b\xd3\x58\x03\xbd\x00\xb9\x72\xe9\x40\x72\x4e\x46\xf2\x74\x13\x21\x1a\x6d\x88\xe7\x15\xad\x82\xd2\x6f\x2b\x09\xb2\x8a\x41\xf1\x27\xb8\x81\x04\xe9\xb0\x60\xb7\x5f\x1f\xbe\xbd\x2e\xe6\x84\x39\x90\x91\xa0\x17\x6c\xae\x05\x71\x23\xb6\xa7\xf3\x97\xb3\xd3\x7f\x5d\xbd\xf5\xea\xf6\xc7\xfd\xed\xcd\xfd\xe3\xdb\x1a\x0e\x6b\xf2\x24\x53\x10\x98\xcb\x52\x91\x57\x9b\xb4\x97\x2f\xe9\xd7\xf5\x44\xbf\x1f\x18\x53\xb8\xc9\xb8\x36\x40\xaf\xa2\x43\x9e\x79\x01\xd2\x12\x95\x21\x2e\x76\xee\xce\xe9\x43\x7d\x0c\x7f\x9e\x02\x19\x77\xe2\x35\xf1\x6e\x30\x3b\x78\xbb\x69\x35\xbf\x88\xee\x01\x0d\xc4\xed\xc3\x6a\xfd\x73\xb5\xee\x8b\x5c\xb2\x4a\x95\x87\x94\x81\x2b\xd0\x2a\xb3\x60\xd7\x83\xec\x77\x70\xd0\x06\x09\x5c\x89\x37\xe5\xdf\x95\x73\x40\x2d\x89\x09\xf3\x08\xf7\xfa\xa0\x84\x73\x5d\x6b\x42\xc6\x78\x3e\x27\xdc\x3b\xd7\xd3\x9a\xf0\x80\xd5\x56\xbf\x1e\x57\xeb\xfb\x63\x06\x74\xec\x06\x7b\x03\xa4\x78\xba\x21\xe0\xf1\x21\xd0\xe7\x27\x8a\x04\x02\xa4\x47\xc8\x76\x51\xa3\x6c\x67\xd3\xb4\xb0\x03\x68\x8f\x8c\xed\x80\x82\x18\xb8\xaa\x2a\x38\xc5\xf3\xae\xdd\xaa\xe5\xc2\xe6\x81\xda\x67\xa8\xa4\x41\x67\xbd\x1b\xe2\x93\xf7\xb7\xff\x86\xd3\x8e\x1b\xee\x55\x35\x5a\x80\x40\x15\x9a\xb5\xe7\xed\xb2\x51\x28\x30\xa0\x41\x8a\x49\x22\x55\xe2\xca\xbf\x21\xfe\x0f\x38\x00\xf6\x20\xac\x47\x00\x35\xe5\xdf\x6f\x96\x2b\x9a\xbe\x7b\xfa\x9e\xa5\x03\x66\xab\xaa\xe6\x16\xdb\xee\x15\xa3\x2e\x1e\xdc\xad\x1a\xa7\x2c\x2b\x83\x39\x17\xdc\x27\xe2\x19\x78\x8e\x55\x6d\xcc\x44\xa5\x2e\x02\xa0\xfa\x14\x52\x0e\xe9\x8d\xe5\xe5\x2c\xd0\x5b\x65\xc6\x46\x5c\xec\x6c\xee\x3b\x63\xd4\x51\x81\x49\x47\x5c\xc6\xd7\xc1\x7e\x23\x10\x48\xb1\xc0\x34\x05\x61\x24\x2a\x57\x05\xfa\x70\xe1\x21\x7f\x3a\x8d\xc7\xb1\x7b\xf0\xb4\x20\x63\x1a\xf5\xe9\xa8\x04\x8c\x47\xd6\xb5\x31\x93\x3c\xa5\x4e\x97\xfc\xe1\xca\xe8\x79\x97\x70\xb5\x1e\x6e\xf7\xe3\x11\xa9\xd4\xc6\x55\xed\x06\x60\xa2\x7c\x25\xc8\xa4\x8a\x81\xda\x3a\xd7\x57\x43\x06\x13\x57\x7b\x6b\xa4\x84\x2b\xa9\x5b\xf5\xa2\x77\x73\x4a\x1e\x39\xa1\x00\xdd\xe7\xf9\xfa\xce\xf2\x54\x13\x26\xe2\x90\x61\xbb\x6c\xb6\xd7\xdf\xb6\xc0\x85\x1e\xef\x7c\x7d\x31\x4e\xf4\x13\xa8\xed\x4c\x4a\xe1\xf8\x84\xae\x06\xa3\x66\xbd\xaa\x6d\x4d\x98\xff\x8d\xd6\xdc\x8a\x83\x6a\x67\x6c\xed\xd5\x3c\x05\xed\x57\xfc\x2a\xc8\x98\xda\x2b\xca\x60\x76\x9f\xf4\x3a\x90\xc0\xc6\x99\x59\x2d\xc5\xa8\x8f\x61\x8d\x47\x93\xe1\x7a\xe7\x6c\xc1\x16\x64\x7c\xf7\x8a\x6c\xd1\x8c\xfc\x1c\x78\xc6\x8c\xb9\xdf\x6f\xf0\xf9\xf0\x81\xc7\xeb\x51\x37\xfe\x41\x37\xf8\x24\x72\x7c\x86\x14\x7c\x9c\x87\xb3\x4a\x38\x3c\xa6\xc1\xa6\xee\xc3\x50\x25\x1c\x7e\x9f\x06\x61\x09\xbc\x3e\x37\x75\x20\x53\x7c\x74\x1a\x91\x18\xb1\x24\xaf\xaf\x76\x67\xf1\xc0\x18\x29\x5e\x34\x9e\x4f\xb0\x16\x24\xf4\x62\xae\xa5\x78\xc6\xc8\xbd\x08\x9c\xe5\x47\x37\x2e\xc3\x8d\x6f\xe3\x2a\x21\xe1\x77\x6f\xd1\x6a\x78\x3a\x0e\xa4\xae\x77\xb7\x20\x17\xee\xfe\x1b\x00\x00\xff\xff\x81\x09\x87\x80\xe7\x1a\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 6887, mode: os.FileMode(436), modTime: time.Unix(1591131103, 0)}
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
