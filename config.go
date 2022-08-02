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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\xdb\x6e\xe3\x36\x10\x7d\xde\x7c\x05\xe1\x0f\xe8\xa2\xc5\xa2\x0f\x46\x11\x60\xdb\x18\x6d\xd0\x45\x36\x70\xdc\xa2\x7d\x32\x28\x6a\x22\x73\x4d\x71\xb4\x43\xd2\x8d\xfb\xf5\x85\x7c\x8b\x2e\x76\x44\x66\x94\x6e\x80\xc0\x00\x35\x87\x67\x66\x38\x3c\x1c\x51\x56\x5b\x0f\xd6\x3b\xf1\xd3\x95\x10\x5f\x30\xdb\xfd\x0a\x61\x65\x09\x53\x31\xc9\x82\x36\xb9\x03\xda\x00\x4d\x76\xe3\x05\x2e\x2b\xe9\x57\x53\x31\x29\xb4\x5f\x85\xec\x3b\x85\xe5\xfb\x8c\xd0\xaf\x80\x0c\x16\x5a\xbd\xef\x41\x2a\x49\x5e\x4b\xb3\xcc\x10\xbd\xf3\x24\xab\xa9\xf0\x14\x60\xf7\x2c\x23\x90\x6b\x0c\xbe\x31\x44\xf0\x35\x68\x82\xf2\xe4\x54\xfd\xa7\xa4\x87\x02\x69\x3b\x15\x37\xb7\x0f\xbf\x1f\x06\x2b\xc2\x0a\xc8\x6b\x70\x53\x31\x71\x8a\xa4\x57\xab\x3d\xe5\x75\xcc\x4c\xbf\x7d\x7e\x58\x2c\x17\x7f\xdf\xcf\xce\x4d\x77\xaf\xc5\x87\xe7\xb9\xea\x7f\x82\x3c\xd8\x5c\x5a\xb5\x9d\x8a\xf9\xec\xe6\x8f\xbb\x9b\x8f\x77\x8b\xd6\xf8\x8f\x1f\x9a\x4f\xae\xaf\x2e\xa4\x76\x38\x85\x8f\xda\x80\xc2\x4a\x1f\x33\x78\x58\x8c\xee\x70\x52\x62\xbb\x21\xfc\xfa\xe9\xf3\xcf\x1f\x3f\x5d\x09\x61\x71\x59\x4a\xe7\x81\x0e\xa6\x17\xfd\x3e\x78\xb1\x01\x72\x1a\xad\x27\xa9\xd6\xf1\x55\x71\x0e\xf5\x6d\xfc\x5f\xc3\xd6\x79\x24\x88\xf5\xbc\x6d\x9f\x5e\xb0\x0f\xb3\xf9\x9f\xb3\xf9\xb9\x1a\x53\x26\x38\xff\xc3\xe5\x2a\x3b\x84\x38\x14\xd0\xa3\xc1\xaa\xda\xc6\x86\xd3\xb4\x7e\xb5\xe7\xcd\x49\x18\x9e\xe7\xd2\xcb\xa4\xb5\xe8\x00\x5e\x29\x15\x9d\x59\x2e\x07\xf0\x6e\x38\x82\x34\xf7\x1b\xd6\xff\x8f\x2a\x8e\x11\xea\x70\x5c\xfb\x71\x25\x29\x6f\xe9\x55\x67\x78\x27\xca\xe7\x55\x74\x30\xcd\x84\x55\x86\x4f\x6e\x6b\x55\x74\xae\xdb\x90\x6b\x21\x2e\x6a\xf8\x10\xfb\x0a\xa4\xf1\x2b\xb5\x82\x14\xc1\xeb\x81\x38\x1e\xc8\x7c\x53\x3f\x4c\x3a\x87\xbb\x18\x0e\xff\x23\xc9\x12\x12\x15\xbf\x8b\xe1\xf0\x57\x84\x4f\xd1\x0a\xd7\x30\x1e\xd8\x35\xb3\xbf\x16\xb3\xf9\xdd\xee\x14\xe9\xed\x1c\x78\xf2\x40\x56\x9a\x25\x81\xcc\x3b\x52\xa7\x30\x58\x3f\x15\xdf\x0f\xfa\x6d\xb0\x28\xb4\x2d\x62\x3d\x6f\x99\xbf\x61\x6b\xf4\xc2\x59\xea\x97\x68\x97\xfb\xa2\x99\x8a\x09\xe5\xda\x55\x46\x6e\x27\xc3\x4a\x98\x01\x50\x62\xa3\xd8\x42\x70\xca\x03\x9e\x40\x05\x8f\xd1\xcc\x6d\xfb\x7e\xab\xd1\x73\x25\xf2\x3c\xab\x08\x37\xba\x6e\x73\xe2\x93\xd0\x83\x9c\x59\xf6\x97\x8e\x0a\x17\x72\x1c\xc3\xf1\xfa\x71\x82\xd3\x0d\xf3\x81\xa2\x9c\xcf\x7e\x99\xdd\xde\x2f\x96\xf7\xf3\xdb\xbb\xc5\xa1\x9b\xb8\x66\xae\xf7\xb1\x2a\x63\x8f\x82\x67\x73\x0e\x6b\x5d\xaf\x89\x1a\xd8\x83\xb4\x93\xf5\x2e\xa1\xe5\xfa\x17\x08\x9b\xbd\xe2\x2e\x90\x9d\x10\x89\x18\x25\x22\x50\x48\xb9\x42\x63\x40\x79\x8d\x36\x36\x82\x73\x38\x4e\x12\xf7\x3c\x04\x0a\x74\x82\x54\xf4\x51\xac\xd3\x24\x18\x53\x2f\x04\xb8\x94\x9a\xef\x82\x38\x1e\x10\x94\xda\xe6\x40\x2e\x7e\x19\x1a\x00\x1e\x73\xbd\x9e\x05\xf8\x84\xd0\xbb\x18\x3e\xbf\x43\x2a\xa4\xd5\x2e\xd5\x87\x36\x8e\xb5\x99\xa5\xd3\xea\x0b\x66\xd1\x3b\xb9\x61\xcf\x8f\x5f\xe6\x79\x6a\xe8\x27\x08\x9f\x3d\x0b\xf5\x62\xa6\xd1\x3f\x63\x46\x88\xde\xd4\xaa\x16\xdf\x16\xf5\x51\x7c\x1f\xca\xba\x23\x4a\x5d\x83\x06\x68\x04\x0f\x30\x41\x00\x3b\x10\x3e\x7b\xe2\x89\xdf\x03\x8d\xe1\x01\x2a\x70\x09\x0a\xd8\x01\x8d\xa0\x42\xd2\x40\x22\xff\x09\x32\x02\xbb\x97\x3e\x95\xfd\x08\xe1\xb0\xab\x63\x12\xe3\xfb\xe5\x1e\xe4\x95\x2f\x22\x24\x75\x1e\x73\xa9\x1a\x97\xc0\x7f\xa4\x4d\x4d\xe0\x09\xc2\x6f\x61\xbc\x74\xeb\x68\xf6\x0e\x84\xd7\xfb\x16\xda\x4b\x53\x47\x62\xb4\x8b\x56\xf1\x33\xb0\x11\x4a\x58\x21\x25\xef\xa0\x13\x86\xc3\x8f\x54\x24\x2a\x58\x1b\xc1\xe1\xae\xe7\x8b\xee\x9f\x4f\xc6\x1c\x46\x0f\x65\x15\x4b\x78\xb4\xe5\xf0\xad\x61\x5b\xca\xaa\x8a\x4f\x6e\x0b\xc0\x61\xde\x6d\xd0\x94\xba\x6e\x01\x78\x0d\xa1\x5a\x87\xe8\x2c\x37\xad\xc7\x11\x43\xf6\x66\xdc\x48\xa3\x73\x99\x70\x0b\x72\x06\xc6\xbb\xa3\xd3\xd6\x67\xa4\xf3\x22\xfa\x3e\xbc\x03\xe1\xb0\x3b\x8f\x69\xdc\x2d\x00\x4f\x0e\xd4\x3a\x5e\x0d\xf6\xb6\x1c\xbe\xaf\x01\x42\x74\x94\x27\x63\x7e\x7d\x3d\x4a\x8b\x21\xf1\xad\xe1\x19\xc3\xe7\x57\x06\x64\xc2\xc5\x5a\x0f\xc4\xba\x60\x0c\x26\x9e\xf9\x64\x3c\xce\x7b\x72\xfa\x1b\x32\xbb\xc5\x90\x6e\x95\xe1\xe9\x6b\x4e\xcc\x67\xb9\x16\xe0\xd5\x9f\x15\x2b\x20\x87\x56\x9a\x97\x64\x31\xf2\x46\xd3\x81\x0a\x04\x49\xdf\x0d\x7a\x90\xb7\xf8\x7a\xc0\x3a\xa1\x6c\xae\x64\x59\x25\xde\xb2\xf7\x50\x2c\x1f\x82\x36\x09\x57\x17\x0d\xf3\x17\x59\xff\x0b\x00\x00\xff\xff\x4e\x8c\x2e\x22\x85\x22\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 8837, mode: os.FileMode(436), modTime: time.Unix(1659451982, 0)}
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
