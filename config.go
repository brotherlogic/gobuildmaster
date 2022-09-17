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

var _configPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\xdb\x6e\xe3\x36\x10\x7d\xde\x7c\x05\xe1\x0f\xe8\xa2\xc5\xa2\x0f\x46\x11\x60\xdb\x18\x6d\xd0\x45\x36\x70\xdc\xa2\x7d\x32\x46\xd4\x44\xe6\x9a\xe2\x68\x87\xa4\x1b\xf7\xeb\x0b\xf9\x16\x5d\xec\x88\x0c\x95\x6e\x80\x20\x00\x35\x87\x67\x66\x38\x3c\x1c\x51\x31\xca\x38\x34\xce\x8a\x9f\xae\x84\xf8\x42\xd9\xee\xaf\x10\x06\x4a\x9c\x8a\x49\xe6\x95\xce\x2d\xf2\x06\x79\xb2\x1b\x2f\x68\x59\x81\x5b\x4d\xc5\xa4\x50\x6e\xe5\xb3\xef\x24\x95\xef\x33\x26\xb7\x42\xd6\x54\x28\xf9\xbe\x07\xa9\x80\x9d\x02\xbd\xcc\x88\x9c\x75\x0c\xd5\x54\x38\xf6\xb8\x7b\x96\x31\xc2\x9a\xbc\x6b\x0c\x31\x7e\xf5\x8a\xb1\x3c\x39\x55\xff\x48\x70\x58\x10\x6f\xa7\xe2\xe6\xf6\xe1\xf7\xc3\x60\xc5\x54\x21\x3b\x85\x76\x2a\x26\x56\x32\x38\xb9\xda\x53\x5e\x87\xcc\xf4\xdb\xe7\x87\xc5\x72\xf1\xf7\xfd\xec\xdc\x74\xf7\x4a\x7c\x78\x9e\xab\xfe\x65\xcc\xbd\xc9\xc1\xc8\xed\x54\xcc\x67\x37\x7f\xdc\xdd\x7c\xbc\x5b\xb4\xc6\x7f\xfc\xd0\x7c\x72\x7d\x75\x21\xb5\xc3\x29\x7c\x54\x1a\x25\x55\xea\x98\xc1\xc3\x62\x74\x87\xa3\x12\xdb\x0d\xe1\xd7\x4f\x9f\x7f\xfe\xf8\xe9\x4a\x08\x43\xcb\x12\xac\x43\x3e\x98\x5e\xf4\xfb\xe0\xc5\x06\xd9\x2a\x32\x8e\x41\xae\xc3\xab\xe2\x1c\xea\xdb\xf8\xbf\xc6\xad\x75\xc4\x18\xea\x79\xdb\x3e\xbe\x60\x1f\x66\xf3\x3f\x67\xf3\x73\x35\x06\x5c\x90\xb9\x5c\x64\x87\x08\x87\xe2\x79\xd4\x54\x55\xdb\xd0\x68\x9a\xd6\xaf\x76\xbc\x39\x49\x82\xe7\x39\x38\x88\x5a\x8a\x0e\xe0\x95\x4a\xd1\x99\xe5\x72\x00\xef\x86\x23\x88\x73\xbf\x61\xfd\xff\x88\xe2\x18\xa1\x0e\xc7\xb5\x1f\x97\xc0\x79\x4b\xae\x3a\xc3\x3b\x4d\x3e\x2f\xa2\x83\x69\x66\xaa\x32\x7a\xb2\x5b\x23\x83\x73\xdd\x86\x5c\x0b\x71\x51\xc2\x87\xd8\x57\x08\xda\xad\xe4\x0a\x63\xf4\xae\x07\x4a\xf1\x00\xf2\x4d\xfd\x30\xea\x18\xee\x62\x52\xf8\x1f\x19\x4a\x8c\x14\xfc\x2e\x26\x85\xbf\x62\x7a\x0a\x56\xb8\x86\xf1\xc0\xae\x99\xfd\xb5\x98\xcd\xef\x76\x87\x48\x6f\xe7\xe0\x93\x43\x36\xa0\x97\x8c\x90\x77\xa4\x4e\x92\x37\x6e\x2a\xbe\x1f\xf4\x5b\x53\x51\x28\x53\x84\x7a\xde\x32\x7f\xc3\xce\xe8\x85\xa3\xd4\x2d\xc9\x2c\xf7\x45\x33\x15\x13\xce\x95\xad\x34\x6c\x27\xc3\x4a\x88\x4f\x28\xbd\xa3\xe0\xf2\x68\xdb\xf7\x4f\xfb\x5e\xb5\x04\x9e\x29\x15\xd3\x46\xd5\x9d\x46\x78\xa5\xf6\x20\x67\x52\xff\x92\x5c\x5b\x9f\xd3\x18\x8e\xd7\x8f\x23\x9c\x6e\x98\xc7\xd7\xca\x7c\xf6\xcb\xec\xf6\x7e\xb1\xbc\x9f\xdf\xde\x2d\x0e\x87\xfc\x75\xe2\x2e\x3d\x16\x4b\xa8\x42\x3f\x9b\xa7\xb0\x66\x88\x1c\x29\x4d\x3d\x48\x3b\x59\xef\x22\x3a\xa1\x7f\x91\xe9\x87\x49\x27\x7d\x3b\x7d\x10\x21\x02\xc1\x28\x89\x73\x49\x5a\xa3\x74\xea\xd8\x0c\x0e\x47\x70\x0e\x97\x92\xc4\x3d\x0f\xa3\x44\x15\x71\xc4\xf4\x51\x49\x22\xef\xb5\xae\x17\x02\x6d\xcc\x36\xe8\x82\x52\x3c\x60\x2c\x95\xc9\x91\x6d\xf8\x32\x34\x00\x69\xcc\xf5\x7a\x16\xe8\x22\x42\xef\x62\xd2\xf9\x2d\x71\x01\x46\xd9\x58\x1f\xda\xb8\xa4\xcd\x0c\x56\xc9\x2f\x94\x05\xef\xe4\x86\x7d\x7a\xfc\x90\xe7\xb1\xa1\x9f\x20\xe9\xec\x99\xaf\x17\x33\x8e\xfe\x19\x33\x42\xf4\xba\x56\xb5\xf0\x6e\xa5\x8f\x4a\xf7\xa1\xac\x1b\x95\xd8\x35\x68\x80\x46\xf0\x80\x22\x04\xb0\x03\x49\x67\x8f\x6c\x02\x7a\xa0\x31\x3c\x20\x89\x36\x42\x01\x3b\xa0\x11\x54\x08\x34\x46\xf2\x9f\x20\x23\xb0\x3b\x70\xb1\xec\x47\x48\x0a\xbb\x3c\x26\x31\xbc\x85\xee\x41\x5e\xf9\x7e\xc0\xa0\xf2\x90\xab\xce\xb0\x04\xfe\x03\x26\x36\x81\x27\x48\x7a\x0b\xe3\xc0\xae\x83\xd9\x3b\x90\xb4\xde\xb7\x50\x0e\x74\x1d\x89\x56\x36\x58\xc5\xcf\xc0\x46\x28\x61\x49\x1c\xbd\x83\x4e\x98\x14\x7e\xe2\x22\x52\xc1\xda\x88\x14\xee\x7a\xbe\xe0\xfe\xf9\x64\x9c\xc2\xe8\xb0\xac\x42\x09\x8f\xb6\x29\x7c\x6b\xdc\x96\x50\x55\xe1\xc9\x6d\x01\x52\x98\x77\x1b\x34\xa6\xae\x5b\x80\xb4\x86\x50\xae\x7d\x70\x96\x9b\xd6\xe3\x88\x61\xf2\x66\xdc\x80\x56\x39\x44\x5c\x8c\x9c\x81\xa5\x5d\x9d\x29\xe3\x32\x56\x79\x11\x7c\x4d\xdd\x81\xa4\xb0\x5b\x47\x71\xdc\x2d\x40\x9a\x1c\xc8\x75\xb8\x1a\xec\x6d\x53\xf8\xbe\x7a\xf4\xc1\x51\x9e\x8c\xd3\xeb\xeb\x11\x0c\xf9\xc8\xb7\x86\x67\x4c\x3a\xbf\xd4\x08\x11\x77\x6d\x3d\x50\x8a\x07\xe8\x75\x38\xf3\xc9\x78\x9c\xf7\xe4\xf8\x37\xe4\xe4\x16\x03\xec\x2a\xa3\xd3\x47\x96\x90\xaf\x65\x2d\xc0\xab\xbf\xf6\x55\xc8\x96\x0c\xe8\x97\x64\x31\xf0\x92\xd3\xa2\xf4\x8c\x51\xd7\xf9\x3d\xc8\x5b\x5c\xea\x27\x9d\x50\x26\x97\x50\x56\x91\xff\x24\xd1\x43\x25\xf9\xe0\x95\x8e\xb8\xba\x68\x98\xbf\xc8\xfa\x5f\x00\x00\x00\xff\xff\xbf\x45\xc7\x74\x1b\x22\x00\x00")

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

	info := bindataFileInfo{name: "config.pb", size: 8731, mode: os.FileMode(436), modTime: time.Unix(1663377105, 0)}
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
