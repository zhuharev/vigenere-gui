// Code generated by go-bindata.
// sources:
// templates/hello.tmpl
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _templatesHelloTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\x5b\x6f\x1b\x45\x14\x7e\xce\xfe\x8a\x61\x54\xc8\x5a\x75\x7c\x59\x1b\xa9\x72\x6d\x47\x82\x06\x82\xd4\x0a\x84\x8c\x04\x4f\xd5\x5e\x26\xf1\x28\xeb\x19\x33\x3b\xeb\xc4\x44\x96\xd2\x08\x09\x10\x37\xa9\x2f\x3c\x21\x78\x42\x3c\x5a\x15\x86\x34\xa5\xe9\x5f\x98\xfd\x47\x9c\xd9\x19\xdb\x6b\x70\x6b\x4b\xc0\x53\x66\xbe\x3d\x97\xef\x7c\xe7\xcc\x89\xdb\xaf\xdd\x7b\xff\xed\xde\x27\x1f\x1c\xa0\xbe\x1c\xc4\x5d\xa7\xad\xff\xa0\xd8\x67\xc7\x1d\x4c\x18\xd6\x00\xf1\xa3\xae\xb3\xd3\x1e\x10\xe9\xa3\xb0\xef\x8b\x84\xc8\x0e\xfe\xa8\xf7\xce\xde\x1d\xac\x71\x49\x65\x4c\xba\xf7\x78\x98\x0e\x08\x93\xed\xaa\xb9\xc3\x87\x44\x8e\xf5\x21\xe0\xd1\x18\x9d\x3b\x3b\x03\x5f\x1c\x53\xb6\x17\x70\x29\xf9\xa0\x85\xde\xac\x0d\xcf\xee\x3a\x13\xa7\x12\x72\x26\x7d\xca\x88\x30\x46\x67\x7b\xa7\x34\x92\xfd\x16\xba\x53\xcb\x2d\xac\x5f\x0b\xd5\x90\x9f\x4a\xae\x5d\x24\x39\x93\xbe\x20\x7e\x99\xb2\x61\x2a\xc1\x0d\xa1\x88\x26\xc3\xd8\x1f\xb7\x50\x10\xf3\xf0\xe4\x2e\x20\x36\x4a\xbd\x56\x7b\x5d\x5f\x79\x2a\x63\x48\xd2\x42\x8c\x33\xa2\x01\x41\x12\xfa\xd9\xe2\x3e\x71\xda\x55\xcb\xb7\x5d\x35\x15\xb7\x35\x71\xf8\x13\xd1\x11\x0a\x63\x3f\x49\x3a\x78\x41\x35\xd7\xc5\xeb\xaa\x5f\xd5\x55\xf6\x79\x76\x81\xd4\x2f\x6a\xa6\xfe\x50\xd3\xec\x22\xfb\x1e\xfc\x3d\xfd\xb9\xd1\x55\x3f\x02\xf2\x28\xbb\xcc\x2e\xd4\x8d\x7a\xaa\xae\xd5\x15\x7c\x6b\x68\x69\x0c\x71\x39\x1e\x92\x0e\xd6\xd5\x60\xc4\xfc\x01\x9c\xf9\xd1\x11\xa8\x8b\x11\x8d\x96\x67\xa8\x2b\x24\x7d\x1e\x47\x44\x74\xb0\x7a\xac\x9e\x40\xaa\xdf\x20\xef\xa5\x9a\xa1\xec\x91\xfa\x53\xcd\xb2\xaf\x00\x7a\xae\xae\xd4\x0c\xa3\x91\x1f\xa7\x10\xa8\x81\x2d\x85\x1f\x80\xc2\x97\x86\x25\x90\x78\x02\x24\xa6\xab\x24\x74\xaa\xfc\x54\xc7\x2b\x84\x6c\x20\xf5\x13\x38\x5e\xe9\xa4\xd9\x65\x19\x69\x77\x75\x8d\x20\xff\x4c\x3d\x53\xd3\x7d\x74\x48\xcb\xa8\xcf\x4f\x11\xb4\x03\x8d\x79\xba\x9f\x4f\x44\x90\x42\x8b\x59\x1e\xd9\x1c\xeb\xf8\x9f\x44\xa6\xa0\xcb\x37\xed\xaa\x31\xc8\xc7\x65\xe8\x1b\x1f\xe8\x4c\x1a\x03\x9d\x2e\xb4\x04\xb0\x3c\xa2\xb0\xe5\xfc\x9c\x2b\xba\x5d\x41\xde\xff\x57\x90\x87\xd7\x51\xd9\x54\x92\xb7\x2c\xe9\x6f\xe3\xf3\x18\x18\xfd\x9e\x37\x71\x06\xb1\xa6\xdb\x8f\xd0\xab\x86\x03\xac\x9f\x65\xdf\x65\x5f\xe0\x35\x83\x36\xa2\xc9\xc3\x13\x32\x36\x93\xb6\xb8\x58\x89\x60\xaa\x66\xea\x1a\x52\x82\x44\xea\x79\xf6\xb5\x7a\xba\x8c\xb5\x79\xac\xe6\xaf\xd3\x66\x3a\xe2\xe2\x21\x61\xa1\xc9\xb4\xb8\x84\x3c\x86\xf7\xd4\xa8\x61\x24\xf8\x29\x9c\xea\x35\x50\xf4\xc5\x2b\x3b\x93\x7d\x0b\x6c\x6e\xd0\x2a\x39\x48\x0e\x85\xbe\x50\x37\x00\x83\x95\x79\x05\xb0\x83\x2c\x87\xb5\xcd\x6b\x6c\x33\x8d\x7a\xe2\xe6\x22\xc3\x72\xf1\x83\x98\x44\x85\x56\x36\xf0\xf6\x63\xb9\x46\x90\x88\x14\x04\xc9\x2f\xeb\x04\xd9\x50\x47\x73\xbb\x21\x7c\x69\x0d\x4d\xdd\xcc\x2a\xac\x37\x3d\x8e\x49\x28\xe8\x50\x76\x1d\xf7\x28\x65\xa1\xa4\x9c\xb9\xa5\x73\x67\x7e\x46\xb7\x5c\x1a\x95\xce\x05\x91\xa9\x60\x28\xb2\xab\xbe\x72\x4c\xe4\x41\x4c\xf4\xf1\xad\xf1\x7b\x91\x36\x99\x2c\x5d\xe0\xa3\x9b\x8a\xb8\x1c\x06\x25\xbd\xd8\x47\xbe\x40\x83\xf1\x87\xe4\xd3\x94\x24\x12\x75\x10\x23\xa7\xe8\xe3\x07\xf7\x0f\xa5\x1c\x5a\xd0\x2d\xe9\x65\x3f\x37\xa9\x70\x16\x73\x3f\x02\xcb\x02\xa1\x30\x70\x97\x06\x50\xc4\x90\xb3\x84\xf4\x40\x23\x48\x5c\x74\x1d\x12\xe6\xee\xbe\x7b\xd0\xdb\x2d\x23\x4d\x01\x49\x91\x92\xd5\xe8\x09\x61\x91\xcb\xd2\x38\x2e\xe9\xdd\xaf\xc9\xd1\x3a\xe4\xba\xe5\xee\x9a\x5d\xb8\x5b\xca\xc1\xc0\x82\x76\x8d\x59\x54\x58\xd4\x2e\x2a\x8b\x52\xaf\x10\xc0\x9b\x07\xf0\x8a\x01\xe6\xa8\xf0\x8a\x01\x00\x2d\x3b\x66\xdb\x1b\xd8\x9c\x35\x1a\x34\x8a\xee\x8d\x1c\x6a\x16\xa1\x26\x44\x74\x82\x3a\xa8\x15\xc6\x34\x3c\x59\x91\x4b\xcb\xae\xbb\x80\xab\xf0\xe4\xf6\x61\xa6\x6e\x17\xca\xab\xe4\x6f\x1d\xdd\x46\xf8\x0d\x93\xae\x83\xe1\x52\x48\x6e\x0c\xca\xcb\x78\x22\x0f\xb8\x53\xac\xbb\x42\x19\xfc\x33\x3c\xec\x3d\xb8\x0f\x89\x85\xb3\x33\x29\x81\x98\x81\xb7\x81\x0d\xcc\xfb\x0a\x1b\xef\xbf\x60\xe3\xbd\x84\x0d\x48\xb8\x81\xce\x88\x16\xe4\xb1\x0b\x6a\x85\x11\xac\xc6\x39\x1d\xbb\x29\xb7\xe0\xd3\x58\x84\x28\x72\x69\x6e\xe4\x52\x10\xc7\xee\x86\x7f\xcf\xa5\xb9\x8e\xcb\xa4\xe4\x96\xf4\x8f\x1e\xfb\xee\x61\x63\x98\x9f\x3b\xd5\xfc\x77\xe0\x5f\x01\x00\x00\xff\xff\xa6\x23\x2c\x06\x17\x0a\x00\x00")

func templatesHelloTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHelloTmpl,
		"templates/hello.tmpl",
	)
}

func templatesHelloTmpl() (*asset, error) {
	bytes, err := templatesHelloTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/hello.tmpl", size: 2583, mode: os.FileMode(436), modTime: time.Unix(1443507178, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"templates/hello.tmpl": templatesHelloTmpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"hello.tmpl": &bintree{templatesHelloTmpl, map[string]*bintree{
		}},
	}},
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

