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

var _config_pb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\xed\x6e\xe3\x36\x10\xfc\x7d\x7e\x0a\xc2\x0f\xd0\x43\xfb\xd3\x28\x02\x5c\x1b\xa3\x08\x7a\x48\x0f\xbe\xb4\xe8\x3f\x63\x45\x6d\x64\x9e\x28\xae\x6e\x49\xa5\x71\x9f\xbe\x90\x3f\x14\x4b\xb2\x63\x32\xab\xb6\x17\x20\x30\x40\xef\x68\x76\x96\x9c\x25\x45\x3b\xe3\x02\xba\xe0\xd5\x8f\x33\xa5\xbe\x50\xb6\xfb\x54\xca\x41\x85\x0b\x35\xcf\x1a\x63\x73\x8f\xfc\x84\x3c\xdf\x8d\x17\xb4\xae\x21\x6c\x16\x6a\x5e\x98\xb0\x69\xb2\xef\x34\x55\xef\x33\xa6\xb0\x41\xb6\x54\x18\xfd\x7e\x04\xa9\x81\x83\x01\xbb\xce\x88\x82\x0f\x0c\xf5\x42\x05\x6e\x70\xf7\x1d\xe3\xd7\xc6\x30\x56\x5d\x06\xed\x9f\x86\x80\x05\xf1\x76\xa1\x6e\xef\x3e\xff\x7a\x18\xac\x99\x6a\xe4\x60\xd0\x2f\xd4\xdc\x6b\x86\xa0\x37\xfb\xe7\xdf\xcc\xf6\xff\x8c\x79\xe3\x72\x70\x7a\xbb\x50\xab\xe5\xed\xef\xf7\xb7\x1f\xee\x1f\x66\x37\xb3\x0b\x0a\xaf\x2b\x79\x34\x16\x35\xd5\xe6\x28\xe4\x50\x93\xe1\xf0\x6b\xfa\x32\x46\x28\xa9\x09\xdd\xd0\x30\xd1\x5f\x3e\xfe\xf6\xd3\x87\x8f\x33\xa5\x1c\xad\x2b\xf0\x01\xf9\x10\x7a\x31\xef\x43\x16\x4f\xc8\xde\x90\x0b\x0c\xba\x8c\x9f\x9c\x73\xa8\xff\x27\xff\x12\xb7\x3e\x10\x63\x6c\xe6\xfd\xf8\x51\x5e\x57\x97\xd2\xe7\xe5\xea\x8f\xe5\xea\xdc\x62\x7a\x64\xa8\x30\x6c\x18\xf1\xf2\x7a\x3a\xc8\xbc\x26\xea\xd1\x52\x5d\x6f\x63\x25\x9d\x46\x4f\x2a\xe8\xe4\xb9\x02\x31\x39\x04\x48\x9a\xa2\x01\xe0\x8d\xde\x1e\x3c\xe5\xb2\x80\x77\xd7\x15\xa4\xa5\xff\x6d\xe4\x7e\x3d\xd1\xfd\xb8\x06\xce\x7b\x7d\x69\x30\xdc\x12\xa7\xf6\xc4\x63\xdd\x98\xea\x8c\x9e\xfd\xd6\xe9\xe8\xe2\xf5\x21\x37\x4a\xa5\x77\xe4\xa3\x89\x76\x7e\x4c\xeb\x6b\x43\x8c\x84\xbf\x66\x7a\x8e\xf6\xf0\x49\x70\xba\x85\x97\x7f\x3e\x2c\x57\xf7\xbb\xf6\x39\x5a\x4a\xf8\x1c\x90\x1d\xd8\x35\x23\xe4\x03\x33\x6b\x6a\x5c\x58\xa8\xef\xaf\x4a\xb1\x54\x14\xc6\x15\xb1\x62\x7a\xe1\xff\xd6\xd6\x1c\x6b\xde\x0c\x91\x13\x4f\x1e\x3d\x84\x64\x09\xe0\x33\xea\x26\x50\x34\x73\x3f\x7e\xbc\x69\x8e\x52\x89\x6c\xc1\x35\xd3\x93\x69\x37\xec\xf8\x22\x8c\x20\x67\x96\xa5\x6f\x72\x9a\x22\xb9\xf6\xeb\x84\xc4\x4e\xc2\xd3\xbd\xb2\x5a\xfe\xbc\xbc\xfb\xf4\xb0\xfe\xb4\xba\xbb\x7f\x38\xec\x7b\x37\xc2\x69\x6e\x17\x4c\x62\xa3\x19\x41\xae\x2a\x79\x97\xb0\x73\xff\x8d\x4c\x3f\xcc\x07\xda\x76\x6e\x57\x31\x76\x67\xd4\xc4\xb9\x26\x6b\x51\x07\x43\x2e\x56\xd4\x39\x9c\xa4\xae\x7b\x1e\x46\x8d\x26\xc1\xbe\x63\x94\xa8\x8b\x37\xd6\xb6\x13\x81\x3e\x65\x8d\x0e\x41\x92\x0c\x18\x2b\xe3\x72\x64\x1f\x3f\x0d\x27\x00\x19\x73\x3b\x9f\x05\x86\x04\xe9\x43\x8c\x9c\xdf\x13\x17\xe0\x8c\x4f\xcd\xa1\x8f\x13\xf9\x1b\xbc\xd1\x5f\x28\x8b\x36\xf7\x49\xbc\x5c\x3f\xe4\x79\xaa\xf4\x0e\x22\x67\xcf\x9a\x76\x32\xd3\xe8\x5f\x30\x13\xa8\xb7\x6d\x57\x8b\x3f\x7b\x8c\x51\xf2\x1c\xaa\xf6\x24\x92\x3a\x07\x27\xa0\x09\x32\xa0\x84\x06\x38\x80\xc8\xd9\x13\x77\xe8\x11\x68\x8a\x0c\x48\xa3\x4f\xe8\x80\x03\xd0\x04\x5d\x08\x2c\x26\xf2\x77\x90\x09\xd8\x03\x84\x54\xf6\x23\x44\xc2\xae\x8f\x45\x8c\x3f\xc3\x8e\x20\x6f\x7c\x01\x60\x30\xb9\xe8\x62\xae\x57\xc0\xbf\xc0\xa5\x16\xb0\x83\x88\xde\x00\x82\x2e\x8d\xb5\xf1\xe6\x19\x00\xa6\x7b\x07\xd8\x13\x06\xf0\x65\x74\x1d\x06\x10\x49\x1d\x72\x53\x98\x00\xb6\xad\xa9\x35\x3e\x7a\x3f\x39\x03\x9b\xc0\x4c\x9a\x38\xd9\xcb\x1d\x46\xc2\x4f\x5c\x24\xf6\xd2\x3e\x42\xc2\xdd\x3e\x2f\xfa\x24\xdf\x05\x4b\x18\x03\x56\x75\x2c\xe1\x31\x56\xc2\x57\xe2\xb6\x82\xba\x8e\x2f\x6e\x0f\x20\x61\xde\xb5\x8a\x94\x75\xdd\x03\xc8\x8e\xa6\xba\x6c\xa2\xab\xfc\x12\x2d\x77\xd1\x13\x58\x93\x43\xc2\xe5\xc6\x19\x98\xec\x7a\xcd\xb8\x90\xb1\xc9\x8b\xe8\x9b\xd9\x01\x44\xc2\xee\x03\xa5\x71\xf7\x00\x32\x1f\xeb\x32\xde\xc6\xfb\x58\x09\xdf\xd7\x06\x9b\x68\x95\x5d\xb0\x7c\x7d\x3d\x82\xa3\x26\xf1\xc5\xe3\x05\x23\xe7\xd7\x16\x21\xe1\xbe\x6c\x04\x12\xed\x96\xe0\x37\x19\x75\x17\xf3\x31\x3f\x99\xf4\x00\x53\xfe\x0a\x54\x23\x7b\x72\x60\x87\xd7\x49\x6f\x38\x82\x78\xd4\x0d\x63\xd2\xbd\xf8\x08\xf2\x1f\xdd\x8e\xbf\x36\x79\xff\x04\x00\x00\xff\xff\xab\x54\x79\x5d\xeb\x1e\x00\x00")

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
