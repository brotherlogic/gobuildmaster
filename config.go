package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _config_pb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x6d\x6f\xe3\x36\x0c\xfe\x7c\xf9\x15\x42\x7e\xc0\x0e\xdb\xc7\x60\x28\x70\x5b\x83\xa1\xd8\xa1\x3b\xe4\xba\x61\xdf\x02\x5a\x66\x1d\x9d\x65\xd1\x47\xc9\x5d\xb3\x5f\x3f\x38\x2f\x6e\x6c\x27\x8d\x58\x79\xdb\x15\x28\x02\x28\x7c\xfc\xf0\xa1\x48\x4a\x66\x9c\x71\x01\x5d\xf0\xea\xc7\x99\x52\x5f\x28\xdb\x7d\x2a\xe5\xa0\xc2\x85\x9a\x67\x8d\xb1\xb9\x47\x7e\x42\x9e\xef\xd6\x0b\x5a\xd7\x10\x36\x0b\x35\x2f\x4c\xd8\x34\xd9\x77\x9a\xaa\xf7\x19\x53\xd8\x20\x5b\x2a\x8c\x7e\x3f\x82\xd4\xc0\xc1\x80\x5d\x67\x44\xc1\x07\x86\x7a\xa1\x02\x37\xb8\xfb\x8e\xf1\x6b\x63\x18\xab\xce\x83\xf6\x4f\x43\xc0\x82\x78\xbb\x50\xb7\x77\x9f\x7f\x3d\x2c\xd6\x4c\x35\x72\x30\xe8\x17\x6a\xee\x35\x43\xd0\x9b\xfd\xf3\x6f\x66\xfb\x7f\xc6\xbc\x71\x39\x38\xbd\x5d\xa8\xd5\xf2\xf6\xf7\xfb\xdb\x0f\xf7\x0f\xb3\x9b\xd9\x05\x85\xd7\x95\x3c\x1a\x8b\x9a\x6a\x73\x14\x72\x88\xc9\x70\xf9\x35\x7d\x19\x23\x94\xd4\x84\x6e\x69\xe8\xe8\x2f\x1f\x7f\xfb\xe9\xc3\xc7\x99\x52\x8e\xd6\x15\xf8\x80\x7c\x30\xbd\xe8\xf7\xc1\x8b\x27\x64\x6f\xc8\x05\x06\x5d\xc6\x6f\xce\x39\xd4\xff\xe3\x7f\x89\x5b\x1f\x88\x31\xd6\xf3\xbe\xfd\xc8\xaf\xab\xa9\xf4\x79\xb9\xfa\x63\xb9\x3a\x97\x4c\x8f\x0c\x15\x86\x0d\x23\x5e\xce\xa7\x83\xcc\x6b\xa2\x72\x08\x20\x52\x35\x00\xbc\xb1\x1c\x06\x4f\xb9\x2c\xe0\xdd\x75\x05\x32\xf7\xbf\x0d\xdf\xaf\x3b\xba\x5f\xd7\xc0\x79\xaf\x94\x07\xcb\x2d\xb1\xb4\x8d\x1c\xe3\xc6\x54\x67\xf4\xec\xb7\x4e\x47\x07\xaf\x0f\xb9\x51\x4a\xde\xc4\x8e\x2d\x69\x97\xc2\xb2\x56\x30\xc4\xa4\xf0\xd7\x4c\xcf\xdb\x58\xe2\x13\x63\x79\x19\x2f\xff\x7c\x58\xae\xee\x77\x1d\x67\x94\x4a\xf8\x1c\x90\x1d\xd8\x35\x23\xe4\xdb\x7e\x3e\x69\x6a\x5c\x58\xa8\xef\xaf\x4a\xb1\x54\x14\xc6\x15\xb1\x62\x7a\xe6\xff\xd6\x69\x16\x5b\xbc\x19\x22\x0b\x0f\xeb\x1e\x22\x25\x05\xf0\x19\x75\x13\x28\x9a\xb9\x6f\x3f\x3e\x67\x46\xae\x44\xb6\xe0\x9a\xe9\xc9\xb4\x67\x5c\x7c\x10\x46\x90\x33\x69\xe9\x9b\x9c\xa6\x70\xae\xfd\x5a\xe0\xd8\x89\xb9\xbc\x56\x56\xcb\x9f\x97\x77\x9f\x1e\xd6\x9f\x56\x77\xf7\x0f\x87\xb3\xef\x26\x71\x9b\xdb\x84\x11\x36\x9a\x11\xe4\xaa\x92\x77\x82\xd3\xfb\x6f\x64\xfa\x61\x3e\xd0\xb6\xab\x76\x15\x53\xee\x8c\x9a\x38\xd7\x64\x2d\xea\x60\xc8\xc5\x8a\x3a\x87\x4b\x89\xeb\x9e\x87\x51\xa3\x11\x94\xef\x18\x95\xd4\xc5\x1b\x6b\xdb\x8d\x40\x2f\xc9\xd1\x21\x28\xc5\x03\xc6\xca\xb8\x1c\xd9\xc7\x6f\xc3\x09\x20\x8d\xb9\xdd\xcf\x02\x83\x40\xfa\x10\x93\xce\xef\x89\x0b\x70\xc6\x4b\x7d\xe8\xe3\x92\xea\x1b\xbc\xd1\x5f\x28\x8b\x2e\xee\x13\xfb\x74\xfd\x90\xe7\x52\xe9\x1d\x24\x9d\x3d\x6b\xda\xcd\x94\xd1\xbf\x60\x26\x50\x6f\xdb\xae\x16\x7f\xf7\x18\xa3\xd2\x7d\xa8\xda\x9b\x88\x74\x0f\x4e\x40\x13\x78\x40\x82\x06\x38\x80\xa4\xb3\x0b\x4f\xe8\x11\x68\x0a\x0f\x48\xa3\x17\x74\xc0\x01\x68\x82\x2e\x04\x16\x85\xfc\x1d\x64\x02\xf6\x00\x41\xca\x7e\x84\xa4\xb0\xeb\x63\x10\xe3\xef\xb0\x23\xc8\x1b\x5f\x00\x18\x4c\x9e\x34\xcb\xea\x05\xf0\x2f\x70\xd2\x00\x76\x90\xa4\x37\x80\xa0\x4b\x63\x6d\x7c\xf1\x0c\x00\xd3\xbd\x03\xec\x09\x03\xf8\x32\x3a\x0e\x03\x48\x4a\x1c\x72\x53\x98\x00\xb6\x8d\xa9\x35\x3e\xfa\x3c\x39\x03\x9b\xa0\x98\x34\xb1\xb8\x96\x3b\x4c\x0a\x3f\x71\x21\xec\xa5\x7d\x44\x0a\x77\xfb\xbc\xe8\x9b\x7c\x67\x9c\xc2\x18\xb0\xaa\x63\x09\x8f\xb6\x29\x7c\x25\x6e\x2b\xa8\xeb\xf8\xe0\xf6\x00\x29\xcc\xbb\x56\x21\xc9\xeb\x1e\x20\xed\x6a\xaa\xcb\x26\x3a\xca\x2f\xd6\xe9\x55\xf4\x04\xd6\xe4\x20\x18\x6e\x9c\x81\xa5\x8d\xd7\x8c\x0b\x19\x9b\xbc\x88\x9e\xcc\x0e\x20\x29\xec\x3e\x90\x8c\xbb\x07\x48\xab\x63\x5d\xc6\x97\xf1\xde\x36\x85\xef\x6b\x83\x4d\xb4\xca\xce\x38\x3d\xbf\x1e\xc1\x51\x23\x7c\xf1\x78\xc1\xa4\xf3\x6b\x8b\x20\x98\x97\x8d\x40\x49\xa7\x25\xf8\x4d\x46\xdd\x60\x3e\xe6\x27\x93\x1e\x60\xca\x5f\x82\x6a\x64\x4f\x0e\xec\x70\x9c\xf4\x86\x2b\x88\x47\xdd\x30\x8a\xe6\xe2\x23\xc8\x7f\x34\x1d\x7f\x6d\xf3\xfe\x09\x00\x00\xff\xff\xab\x5f\x88\x0b\x1e\x1e\x00\x00")

func config_pb() ([]byte, error) {
	return bindata_read(
		_config_pb,
		"config.pb",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"config.pb": config_pb,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"config.pb": &_bintree_t{config_pb, map[string]*_bintree_t{
	}},
}}
