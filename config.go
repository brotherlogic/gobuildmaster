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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\xd1\x8a\xdb\x3a\x10\x7d\xbe\xf9\x0a\x91\x0f\xb8\x4b\x5f\x4d\x09\x6c\xd9\x50\x4a\x97\x2d\x64\xb7\xa5\x6f\x41\x96\xa7\x8e\x76\x65\x8d\x3b\x1a\x6d\x93\xbf\x2f\x4e\xe2\x34\xb6\x93\x8d\x14\xb9\x34\x10\x02\xf2\x9c\x9c\x33\xa3\x33\x23\x61\xab\x2d\x83\x65\x27\xde\x4f\x84\x78\xc6\x7c\xfb\x2b\x84\x95\x15\x64\x62\x9a\x7b\x6d\x0a\x07\xf4\x0a\x34\xdd\xae\x97\xb8\xac\x25\xaf\x32\x31\x2d\x35\xaf\x7c\xfe\xbf\xc2\xea\x26\x27\xe4\x15\x90\xc1\x52\xab\x9b\x01\xa4\x96\xc4\x5a\x9a\x65\x8e\xc8\x8e\x49\xd6\x99\x60\xf2\xb0\x7d\x46\xf0\xd3\x6b\x82\xea\xa0\xa0\xf9\x28\xc9\x50\x22\x6d\x32\x71\xf7\xe9\xf1\xf3\x7e\xb1\x26\xac\x81\x58\x83\xcb\xc4\xd4\x29\x92\xac\x56\xbb\xff\x9f\x4d\x76\x5f\x82\xc2\xdb\x42\x5a\xb5\xc9\xc4\x62\x7e\xf7\xf5\xe1\xee\xf6\xe1\x69\x32\x9b\x9c\xc9\xf0\x72\x26\x3f\xb4\x01\x85\xb5\x6e\x13\xd9\xd7\xa4\xbf\xfc\x56\x7e\x39\x81\x7c\x41\xcf\x87\xa5\xbe\xd0\x8f\xf7\x5f\x3e\xdc\xde\x4f\x84\xb0\xb8\xac\xa4\x63\xa0\x7d\xe8\x59\xdd\x7b\x15\xaf\x40\x4e\xa3\x65\x92\xea\x25\x7c\x73\x4e\xa1\xfe\x8d\xfe\x17\xd8\x38\x46\x82\x50\xe5\xdd\xf8\x0b\xbe\x79\x9c\x2f\xbe\xcd\x17\xa7\x9c\x43\xde\xda\x36\xf1\x53\xc6\xd9\xe7\x73\x49\x7d\x21\x59\x46\xc9\xef\x01\xae\xf4\x7d\xb7\x08\xe7\xf5\xff\x97\x62\xfb\xdd\xba\x92\x54\x74\x6c\xdf\x5b\x6e\x88\x63\x5b\xae\x2d\x1e\x61\x9d\xe3\xda\x6d\xac\x0a\x2e\x5f\x17\x32\x13\x22\xbe\xe1\xdb\xf6\x25\x59\x41\x64\xdb\xf4\x31\x29\xfc\x35\xe1\x7a\x13\x4a\x7c\x14\x3c\x68\xc5\x8b\x2e\x9a\x7f\x7f\x9a\x2f\x1e\xb6\xdd\x39\x70\x12\xac\x19\xc8\x4a\xb3\x24\x90\xc5\xa6\xeb\x27\x85\xde\x72\x26\xde\x5d\x4c\xc5\x60\x59\x6a\x5b\x86\x26\xd3\x09\xff\xeb\x93\xff\x8d\x1e\x68\x0f\x37\x00\x8a\x3c\xdb\x3a\x88\x14\x17\xc0\x1a\x94\x67\x0c\x66\xee\xc6\x0f\xc7\xf2\x40\x4a\xe0\x20\xab\x09\x5f\x75\x73\x24\x84\x17\x61\x00\x39\xe1\x4c\xe7\x0b\x1c\x43\x5c\xf3\x38\x42\xd8\x51\x78\xc2\xad\xe3\xfc\xe9\xd1\x61\x98\x1d\x52\x0b\x6e\x98\xc6\x3f\x91\xa3\xa7\x07\xb9\xba\x96\x04\x0a\xa9\x50\x68\x0c\x28\xd6\x68\x43\xf9\x4f\xe1\x52\x8c\xbf\xe3\x21\x50\xa0\x23\x1a\x6f\x88\x4a\x1a\xc1\xde\x98\xc6\x03\xe0\x62\xdc\xd5\x07\xa5\x28\x20\xa8\xb4\x2d\x80\x5c\xf8\x36\x1c\x01\xd2\x98\x9b\xfd\x2c\x81\x23\x52\xef\x63\xd2\xf9\x1d\x52\x29\xad\x76\xb1\x1a\xba\xb8\x14\x1d\xb9\x74\x5a\x3d\x63\x1e\xdc\x87\x47\xf1\xe9\xf9\xcb\xa2\x88\x4d\xfd\x00\x49\x67\xcf\x7d\xb3\x99\x71\xf4\x7f\x30\x23\x64\x6f\x9a\x89\x1a\x7e\x71\x18\xa2\xd2\x35\x54\xcd\x35\x22\x76\x0f\x8e\x40\x23\x28\xc0\x88\x01\xd8\x83\xa4\xb3\x47\x9e\xad\x03\xd0\x18\x0a\x50\x81\x8b\x98\x80\x3d\xd0\x08\x53\x48\x1a\x88\xe4\x3f\x40\x46\x60\x67\xc9\xb1\xec\x2d\x24\x85\x5d\xb5\x45\x0c\xbf\x7d\x0e\x20\x57\xde\xde\x2b\xef\xb4\x4a\x7a\x6b\xd3\xa9\xe0\x2f\x69\x63\x2b\x78\x80\xbc\x5d\xc1\xdf\x01\x00\x00\xff\xff\x3a\x09\xc3\xe6\x16\x13\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 4886, mode: os.FileMode(436), modTime: time.Unix(1597100173, 0)}
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
