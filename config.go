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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\xdb\x6e\xe3\x36\x10\x7d\xde\x7c\x05\xe1\x0f\xe8\xa2\x7d\x34\x8a\x00\xdb\xc6\x28\x82\x2e\xd2\x85\x37\x2d\xfa\x66\x8c\xa8\x89\xc4\x35\xc5\xd1\x0e\x29\x37\xee\xd7\x17\xf2\x45\xab\x8b\x13\x93\x19\xb5\x5d\x60\x11\x80\x9a\xa3\x33\x67\x48\x1e\x8e\x29\x67\x5c\x40\x17\xbc\xfa\xf1\x46\xa9\x2f\x94\x1d\xfe\x2a\xe5\xa0\xc2\xa5\x5a\x64\x8d\xb1\xb9\x47\xde\x21\x2f\x0e\xe3\x05\x6d\x6a\x08\xe5\x52\x2d\x0a\x13\xca\x26\xfb\x4e\x53\xf5\x3e\x63\x0a\x25\xb2\xa5\xc2\xe8\xf7\x13\x48\x0d\x1c\x0c\xd8\x4d\x46\x14\x7c\x60\xa8\x97\x2a\x70\x83\x87\x67\x8c\x5f\x1b\xc3\x58\x75\x19\xb4\xff\x34\x04\x2c\x88\xf7\x4b\x75\x77\xff\xf9\xd7\xd3\x60\xcd\x54\x23\x07\x83\x7e\xa9\x16\x5e\x33\x04\x5d\x1e\xdf\x7f\x7b\x73\xfc\xcf\x98\x37\x2e\x07\xa7\xf7\x4b\xb5\x5e\xdd\xfd\xfe\x70\xf7\xe1\xe1\xf1\xe6\xf6\xe6\x05\x85\xd7\x95\x3c\x19\x8b\x9a\x6a\x73\x16\x72\xaa\xc9\x78\xf8\x35\x7d\x19\x23\x6c\xa9\x09\xdd\xd0\x38\xd1\x5f\x3e\xfe\xf6\xd3\x87\x8f\x37\x4a\x39\xda\x54\xe0\x03\xf2\x29\xf4\xc5\xbc\x4f\x59\xec\x90\xbd\x21\x17\x18\xf4\x36\x7e\x72\x2e\xa1\xfe\x9f\xfc\xb7\xb8\xf7\x81\x18\x63\x33\x1f\xc6\x4f\xf2\xba\xba\x94\x3e\xaf\xd6\x7f\xac\xd6\x97\x16\xd3\x13\x43\x85\xa1\x64\xc4\x97\xd7\xd3\x49\xe6\x35\x51\x4f\x96\xea\x7a\x1f\x2b\xa9\x1f\x3d\xab\xa0\xde\x7b\x05\x62\x72\x08\x90\x34\x45\x23\xc0\x1b\xf7\xf6\xe8\x2d\x2f\x0b\x78\x77\x5d\x41\x5a\xfa\xbd\xe8\x7f\xc1\xb3\xe6\xd0\x75\x5d\xc4\x71\x5c\x03\xe7\x03\xcf\x1a\x0d\xb7\xc4\xa9\x7e\x79\xae\x29\x53\x9d\xd1\xb3\xdf\x3b\x1d\x5d\xd8\x21\xe4\x56\xa9\x74\xb7\x3e\xb1\x97\x08\x36\x94\xba\xc4\x14\xd3\x9b\x80\x24\x19\x40\xbe\x6b\x1f\x26\x1d\x89\x63\x8c\x84\xff\xe8\x56\x69\xae\x3f\xc6\x48\xf8\x6b\xa6\xe7\x68\x87\xeb\x05\xa7\x1b\xdc\xea\xcf\xc7\xd5\xfa\xe1\x70\xb8\x4c\x36\x13\x3e\x07\x64\x07\x76\xc3\x08\xf9\xc8\xea\x34\x35\x2e\x2c\xd5\xf7\x57\xa5\x58\x2a\x0a\xe3\x8a\x58\x31\x83\xf0\x74\x39\x6f\xee\x65\x7a\x47\x6c\xd8\x90\xdb\x1c\xd7\xd1\x52\x2d\x38\x37\xbe\xb6\xb0\x5f\x5c\x77\xc2\x0c\x91\x13\xdb\xb8\x01\x42\xb2\x62\xf0\x19\x75\x13\x28\x9a\x79\x18\x3f\xed\x40\x26\xa9\x44\x9e\x67\x35\xd3\xce\xb4\xdd\x4f\x7c\x11\x26\x90\x0b\xd3\xee\x9b\x9c\xe6\x48\xae\x7d\x9c\x90\x58\x2f\x3c\x7d\x2d\xae\x57\x3f\xaf\xee\x3f\x3d\x6e\x3e\xad\xef\x1f\x1e\x4f\x4d\xc4\xad\x70\x9a\xcf\x8b\x31\xf6\x50\xe8\x87\x47\x4c\x72\x7c\x22\xed\xca\x4d\x34\xc8\x09\xe4\x6a\x49\xdf\x25\xf4\x63\x7f\x23\xd3\x0f\x8b\x51\x91\x0f\x2e\xa5\x62\x6c\x8a\x51\x13\xe7\x9a\xac\x45\x1d\x0c\xb9\x58\x51\x97\x70\x92\xba\x1e\x79\x18\x35\x9a\x04\x1f\x99\xa2\x44\xa7\x4f\x63\x6d\x3b\x11\xe8\x53\x36\xcb\x18\x24\xc9\x80\xb1\x32\x2e\x47\xf6\xf1\xd3\xd0\x03\xc8\x98\xdb\xf9\x2c\x30\x24\x48\x1f\x63\xe4\xfc\x9e\xb8\x00\x67\x7c\x6a\x0e\x43\x9c\x68\x7f\x83\x37\xfa\x0b\x65\xd1\x9b\xbb\x17\x2f\xd7\x0f\x79\x9e\x2a\xbd\x83\xc8\xd9\xb3\xa6\x9d\xcc\x34\xfa\x6f\x98\x19\xd4\xdb\xd6\xd5\xe2\x7b\xa6\x29\x4a\x9e\x43\xd5\xb6\x4b\xa9\x73\xd0\x03\xcd\x90\x01\x25\x18\xe0\x08\x22\x67\x4f\x6c\x15\x26\xa0\x39\x32\x20\x8d\x3e\xc1\x01\x47\xa0\x19\x5c\x08\x2c\x26\xf2\x77\x90\x19\xd8\x03\x84\x54\xf6\x33\x44\xc2\xae\xcf\x45\x8c\x6f\xa6\x27\x90\x37\xfe\x4a\x61\x30\xb9\xe8\xba\x75\x50\xc0\xbf\xc0\xa5\x16\xb0\x83\xc8\x5b\x98\x00\x7e\x1b\xcd\x3e\x82\xc8\x3a\xe4\xc2\x04\xb0\xad\x12\x6b\x7c\xb4\x8b\x5f\x80\xcd\xb0\x84\x35\x71\xf2\x0e\xea\x30\x12\x7e\xe2\x22\xd1\xc1\x86\x08\x09\x77\xfb\xbe\xe8\xfe\xb9\x0b\x96\x30\x06\xac\xea\x58\xc2\x73\xac\x84\x6f\x8b\xfb\x0a\xea\x3a\xbe\xb8\x03\x80\x84\xf9\xb0\x41\x53\xd6\xf5\x00\x20\x6b\x08\xf5\xb6\x89\xae\x72\x3f\x7a\x1e\x33\x14\x6f\xc6\x1d\x58\x93\x43\xc2\x15\xc9\x05\x98\xec\x4e\xcf\xb8\x90\xb1\xc9\x8b\xe8\xcb\xf2\x11\x44\xc2\xee\x03\xa5\x71\x0f\x00\x32\x3b\xd0\xdb\x78\x37\x38\xc6\x4a\xf8\xbe\x36\xd8\x44\xab\xec\x82\xe5\xeb\xeb\x09\x1c\x35\x89\xbf\x1a\xbe\x61\xe4\xfc\xda\x22\x24\xdc\xba\x4d\x40\xa2\xdb\xc7\xc6\xc6\x33\x77\xc1\xa2\x63\x1e\x7c\x99\x51\xf7\x05\x26\xe6\xbb\xd9\x00\x30\xe7\xa7\xc0\x1a\xd9\x93\x03\xfb\x9a\x5b\x45\xde\x50\x7a\xd4\x0d\x63\xd2\xf5\xff\x04\xf2\x1f\x7d\x04\x78\x6d\xf2\xfe\x09\x00\x00\xff\xff\x05\x72\x5d\xec\xf0\x20\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 8432, mode: os.FileMode(436), modTime: time.Unix(1642625342, 0)}
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
