// Code generated by go-bindata.
// sources:
// templates/enum.go.tpl
// templates/model.go.tpl
// templates/pogo.go.tpl
// DO NOT EDIT!

package templates

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

var _templatesEnumGoTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templatesEnumGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesEnumGoTpl,
		"templates/enum.go.tpl",
	)
}

func templatesEnumGoTpl() (*asset, error) {
	bytes, err := templatesEnumGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/enum.go.tpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1501948970, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGoTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x5d\x73\xe3\xb6\x15\x7d\x26\x7f\xc5\x2d\xc7\xc9\x4a\x8e\x4a\x35\xaf\xee\xe8\x61\x6d\x6b\x5b\x77\x5c\xef\xd6\x96\xd3\x87\x4d\xc6\x84\xc9\x2b\x09\x31\x09\xd0\x00\x24\x59\xe5\xf2\xbf\x77\x00\x7e\x4b\xa4\xbc\xde\x9a\x4e\x3a\x93\x9d\xd9\x19\x19\xb8\x00\x2e\x2e\xce\x39\x3c\x40\x92\x8c\x8f\x5f\xeb\xdf\x38\x4d\x6d\x3d\x1f\xfc\x44\x04\x25\xf7\x21\x4a\x28\xdb\x5e\x73\x0d\x3b\x49\xe0\x48\xc1\xc9\x04\x94\x5e\x85\x91\x08\xe1\xc8\xbd\xf1\x97\x18\x11\x70\x67\xba\x0d\xcc\xb2\x70\xa4\x98\x0e\xcb\xda\xdc\x2b\x1d\x98\x77\xf8\xba\x5d\x77\x7f\x01\x9f\xc4\x54\x91\x90\xfe\xa7\xea\x5c\xb7\xf6\x7e\x81\x90\x6f\x50\x14\x51\x51\x15\x24\x29\x5b\xac\x42\x52\x75\xad\x5b\xfa\x76\x87\xaf\x17\x3a\x28\x16\x94\x29\x70\xbe\x77\xcc\x28\xc7\x75\x8a\xfe\xf9\xf3\x53\xc4\xf9\x04\x11\x11\xdb\x62\x93\x67\x3c\x5c\x45\x4c\x96\x21\xa6\x4e\x3e\x47\xe1\xd7\x8a\x74\x14\xbb\xe7\x44\x91\xd9\x36\xae\xf6\xcc\xb3\xc0\x50\x97\x53\x76\xcc\x46\x83\x27\xa9\xc3\x68\xf0\xd4\x08\xbb\x60\x01\x3e\xa1\xac\xe6\x32\xc9\x47\x24\x36\xf3\x0e\x22\xb3\xc9\x39\x38\x3f\x3b\xdf\xc9\x9f\x1d\x67\x08\x5f\xe0\x57\x4e\x19\x38\x23\x70\xaa\x41\x8b\xfa\xa0\xa8\x56\x78\x3d\x01\xce\xe9\x93\x29\xda\xde\xe0\x57\xc7\x97\xc1\xf0\xc7\x95\x80\x4f\xc4\x7f\x20\x0b\xec\x09\xc5\x71\x3e\x7b\x92\x80\x7b\x83\x4a\x51\xb6\x90\x6e\xb1\x64\x6f\xfb\x9a\x2d\x31\xa3\x0d\x6c\xf0\x5d\x18\x02\x51\x8a\xf8\x4b\xe0\x2b\x01\x67\xd7\xb7\xe7\x10\xa1\x5a\xf2\x40\x02\x67\x8a\xf7\xb4\xf1\xf1\x18\x32\x0a\xa6\x29\xf8\x21\x91\xd2\x56\x1a\x88\x65\x9b\x54\x62\xe5\x2b\x48\x6c\x80\xe0\x1e\x8e\xcf\x4f\xed\x3e\x8b\x11\xf1\x00\x43\x50\x4b\xa2\xc0\xe7\x4c\x11\xca\x24\x90\x30\x34\x15\x09\x88\x22\xf7\x44\x22\xcc\x29\x86\x41\x5f\x72\x96\xd7\x23\xd2\x7b\x37\xd9\x54\xf5\x88\x76\xea\x91\x24\x20\x08\x5b\xe0\x3e\x3f\x4b\x4d\xdc\xe5\x7a\x83\xe9\x66\x86\x4c\x08\x77\x84\x0f\xb2\x09\xd2\x14\xbc\x5f\x25\x67\x27\x4e\x19\x98\xa6\x23\x1e\x51\x85\x51\xac\xb6\x8e\xa7\xe3\xe8\x1c\xdc\x33\x1e\x45\xc8\xf4\x80\x2c\xfd\x5a\x43\x92\x00\xb2\xa0\xfe\xa3\xb7\x03\xfc\x24\xe8\x9a\x28\xcc\x60\xa4\xcf\x2f\xab\x15\x17\x23\x20\xbe\x8f\x52\x62\x00\x6b\x4a\x20\xe6\x0b\xee\x1e\xcd\xde\x9f\x5e\x4e\xfb\x3d\xc4\xb9\x01\x75\x95\x87\x3d\x5f\x31\xbf\xec\x19\x68\x3c\xc7\x8b\x27\xf7\x8c\x33\x36\x84\xe3\x12\xf2\xfa\x6c\x05\xaa\x95\x60\xf0\x7d\xd1\x98\x04\xf7\xfd\x17\x6e\x89\x61\x8c\x02\x14\x87\x05\x2a\x83\x7b\xb5\x44\x60\x9c\xfd\x99\xd1\xb0\x80\x3d\x67\x86\x0d\x19\x53\x7a\x2b\x5f\x77\x02\x59\x11\xb3\xdf\x83\xfc\xfb\x9a\xa6\x59\xf9\x34\x43\x86\xfa\xab\xf1\x59\x2a\x41\xd9\xe2\x17\xca\x14\x8a\x39\xf1\x31\xc9\xca\x9a\xef\xc1\x7c\x5a\x1e\x70\xd0\x1e\x39\x7c\x96\x5b\x66\x1a\x63\x26\xda\xd8\x63\x83\x26\x45\x99\x9a\x5b\x8d\x48\x53\xf8\xd3\x04\xf4\x5e\x74\x32\x45\x3a\x9f\xeb\xf4\x72\x7e\x81\x49\xc7\x58\x1b\xa0\xc6\x22\x80\x0a\x26\x79\x65\xfa\xc2\x47\x8d\x30\xee\x07\xca\x82\xc1\xf0\x04\xe6\x94\x05\xc0\x19\x82\xe0\x1b\xb8\xdf\x02\x55\xef\x64\x69\x3f\x1e\x70\xdb\x9f\xdb\xa3\x73\x6d\x76\xd2\x54\xc3\x44\x67\x03\xa4\xaa\x97\xce\xc4\x31\x56\xa7\xac\x67\x86\x97\x41\x6e\xe5\x0a\xa4\xf8\x06\x29\x66\x33\x8d\x70\xc8\x7d\x92\xee\x6d\x43\xd7\x08\x50\x08\xfd\x9f\x8b\x21\x24\xb6\x35\x1e\x83\x7c\x0c\x41\x62\x88\xbe\x82\xc7\x15\x8a\xed\xa8\x51\x86\x58\xf0\x35\x0d\x30\xd0\x99\x49\x7c\x5c\x21\xf3\xd1\xb6\xe4\x63\x28\x95\xd0\x08\xf2\x0c\x10\x6e\xa6\x97\xd3\xb3\x19\x14\x7e\xc9\x1c\x2f\xc0\x87\xeb\x8f\xff\x2c\xa5\xd8\xb4\xfc\xfb\xef\xd3\xeb\xe9\xee\x0e\x61\x02\x47\x3f\x7a\xb6\x6d\x5d\xf2\xc5\x20\x9b\x79\x04\x8d\x90\xa1\x6d\xe9\x63\x3a\xc9\x90\x65\xca\xe0\x06\xf7\xee\xbf\x74\xba\xd7\x7c\xd3\x35\x08\xcc\x66\x27\xfa\x88\xdd\x1b\x9f\xb0\x41\x61\xcd\xcc\x8c\x74\x6e\xba\x4b\x40\x5b\x45\xcb\x64\x02\x5a\xd6\xa6\x42\x5c\xf1\x6b\xbe\x91\xa6\xcf\xca\xa1\x5a\xd6\x74\x04\x30\x15\xa2\xa8\xeb\x15\x57\x1f\xf8\x8a\x05\xb6\x65\xa5\x76\x5b\x30\x0a\x61\x5b\xa9\x6d\x5b\x75\x65\xcc\xfb\x18\x0d\x6d\xe3\x1b\x73\x6a\xbc\x0d\x0b\x4e\xb7\xae\xeb\x96\x54\x20\x15\x11\x24\xac\x18\x7d\x5c\x65\xc2\x55\x60\x81\xe6\xae\xb8\x37\x5a\x64\x72\xa5\x5d\x79\xed\x96\xb3\xe7\xc5\xc3\xc2\xb2\x97\xe6\xde\x8c\xa8\x2c\x7d\x66\xf9\x6a\xe6\x3b\x94\x0d\xfb\x5d\xf8\xed\xf7\x2c\x70\x6a\xc3\x62\x22\x48\x54\xcc\x9d\xff\x51\xdd\x30\xea\x4b\xe8\x8c\xd6\x44\xc8\xbd\x25\x22\x6c\x2c\x90\x1b\xfa\x9c\xe4\xa7\xdb\x66\x7e\x69\x5a\x94\xbd\xc1\xfd\x72\x93\xc5\x2c\x1e\x10\x16\x80\xe7\xc0\x17\x28\xee\x1d\xde\x77\xd2\x73\xd2\xf4\x39\x5d\xd8\x5f\x71\xd0\xdc\xeb\x8e\x44\xf4\xa8\x10\x96\xb5\xaf\x0f\x96\xd5\x54\x07\xcb\xca\xb4\x41\x0b\x64\xf0\xb4\x59\xa2\xc0\xa2\xee\x6d\xd2\x50\x1d\xc3\xcb\xd4\x61\x67\xdc\xef\x57\x1f\x7e\x53\x79\xb8\x60\x12\x85\xd2\xda\x40\xcd\x2f\x20\xc0\x70\x63\x14\x82\xea\xab\x94\x2a\xef\x5d\xbd\x19\xa8\x8b\x62\x61\xaf\xac\x84\x57\x2d\xee\x15\xb8\xf1\xb2\x3c\xdc\x43\x6c\xc8\x77\xb3\x0f\xf5\x21\x0c\x9a\x1f\xc6\x0a\xf2\x07\x0c\xa4\x26\x64\x2c\x30\x26\x02\x75\x77\x04\x73\x2e\x4c\x9c\xe1\x86\x6d\xdd\xf9\x23\xb8\xa3\x23\xb8\x33\xef\x26\xa6\xac\x32\xa4\x3e\x0e\x76\x3d\xdf\x70\x04\x7f\x19\xda\x25\xc3\xf2\x52\x7f\x1b\xc3\x2e\xae\x6e\xa6\xd7\x33\xb8\xb8\x9a\x7d\xac\x2e\x3f\x03\x0f\x7e\x80\xcc\x1f\x4a\xf7\x1f\x9c\xb2\x81\xce\x4d\x2b\xd3\x10\x7e\x00\x6f\x68\x5b\x3f\xbd\xbf\xbc\x9d\xde\xb4\x05\xd2\x66\xe0\xf5\x74\x76\x7b\x7d\x75\x71\xf5\xb7\x1a\x87\x77\x79\x79\xb7\x76\x5d\xf7\x6b\xc9\x58\x04\x6b\x4e\xe9\xe8\x36\x0e\xfe\x15\xb0\xe9\x37\x73\x6e\x30\x1a\x8e\x00\xb5\x9d\x3c\xf4\x3d\xed\x9f\x26\xb7\x71\x40\x14\x6a\x9a\xac\xcc\x2f\x20\x0c\xf0\x89\x4a\x45\xd9\x22\x27\xcb\x5b\x50\xe5\x36\x5f\xbc\xf9\x29\xd1\xdf\x71\xaf\x61\x88\xbc\x43\x24\xc9\xf7\xd2\xf6\x3d\xe8\x36\x98\xad\xec\xc9\x60\x8e\xe6\xfe\xbe\x07\xf9\x0c\xee\x73\x2a\xa4\x02\x7f\x89\xfe\x83\xb6\xc3\x1b\x84\x25\x59\x1b\x36\xd5\x71\x6f\xc0\xd1\x5c\x7c\x52\x89\x71\x03\x0b\x7a\x71\xe9\x5e\xe1\x66\xe0\xd5\x89\xb3\x6b\x34\xa3\x95\x54\x70\x5f\x32\x5a\x03\x3b\xcd\x32\x0a\x38\x7b\xa7\x8a\x63\xdc\xcb\x23\xc0\x10\x55\x41\x60\x54\xa3\xdd\x89\xf3\x6d\xd5\x54\x01\x0c\xe5\xe5\xb3\x6a\xa0\x27\xfb\xb1\x26\x02\xb9\x86\xd4\xb8\x7d\xfb\xe9\xfc\xfd\x6c\x5a\x91\xfa\x66\x3a\x33\x7c\xb5\x2d\xeb\x00\xb5\x61\xd2\x1e\xb4\x43\x6b\xab\xdb\x90\xdb\x56\x27\xe9\xc7\x63\x10\x2b\x56\xe4\xba\x26\xe1\x0a\x8d\x1d\x22\x71\x8c\x2c\x18\x7c\xae\x5f\x46\x93\x9d\x13\x4c\x2b\xea\xd7\xb5\x23\x9b\xc4\xb4\x7f\xa5\x80\xd4\x46\xbc\x9a\x88\xf4\xab\x21\xc7\x87\x75\xa4\xf4\xe3\x85\x98\xfc\x56\x8e\xfc\xf8\xff\xcb\x95\x97\x0a\xf8\x8c\xc7\x86\xc3\x9e\xb9\x6b\x8e\x4e\x49\xdc\x71\xd2\xaf\xac\x85\x85\xf9\x78\xc0\xad\xd4\x8d\x0c\x31\xb0\xad\xc6\x89\xd4\x1e\x75\x32\xa1\xac\x5e\x73\xf2\xbb\xc8\xd7\x29\x66\xc7\xc8\x56\xb1\xb4\x6a\x36\xb4\x55\x37\x75\xbe\xdd\x79\xb6\x0a\x69\xcb\xe2\xce\x70\x6f\xa5\x6f\x14\xd7\x24\x81\x10\x59\x0e\xb9\xb2\xe8\x6f\x2e\xb5\x90\x45\x78\x9d\x97\x9c\x22\xe0\xa5\x9a\xdb\x14\xdb\xb4\x56\xfa\x5a\xd9\xf3\xf0\x52\xa1\xb3\xbf\x47\x1d\x98\x69\xd4\xde\xea\x1a\xfb\x72\x15\x3f\x3f\x3d\xac\xe2\x2f\xbb\x8c\xfd\x7e\xae\x50\xc7\x5d\xba\x7e\x9e\xa1\x7d\x78\x02\x19\xee\x73\x49\x5f\x49\xed\x0f\xb5\xaa\x37\xa4\xbc\x47\x21\x1f\x8f\xe1\xbc\xc8\xa0\x7e\x9f\x9a\x0b\x1e\xb5\xde\xa7\x0e\x09\x65\xbe\xab\x4e\x57\x68\xa4\xa5\xf6\x72\xb0\xcf\xb3\xf3\xe9\xe5\x74\x36\x6d\x3e\x0f\x3e\xf3\x34\xd8\x84\xff\xa1\x87\x42\x3a\x87\xbb\x51\x66\x07\x1a\x06\x62\xfa\x84\x7e\xc7\xa0\x3d\x8b\x60\xfc\x44\xcb\x35\x1f\xa0\x66\x20\xda\x6e\xf7\xba\x3b\xad\xbb\x8c\x1d\x83\xf1\x96\xae\x22\x3b\xa7\xd2\x55\x34\x20\xf8\x87\xab\xe8\xe9\xad\xaf\x28\xfa\xff\xe2\x44\xba\xe6\x68\x79\xbf\xdb\xe5\x5a\x7e\xc8\x2f\xa2\xdc\xb7\xbd\xb8\xe5\x34\x13\x62\x4f\xe3\xf7\x88\xd6\x18\xf9\xd6\x5c\xab\x34\xff\xbf\x01\x00\x00\xff\xff\x5a\x1a\xd9\xa9\x1e\x24\x00\x00")

func templatesModelGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModelGoTpl,
		"templates/model.go.tpl",
	)
}

func templatesModelGoTpl() (*asset, error) {
	bytes, err := templatesModelGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/model.go.tpl", size: 9246, mode: os.FileMode(420), modTime: time.Unix(1501952883, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPogoGoTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x53\x4d\x6f\xab\x46\x14\x5d\x33\xbf\xe2\x08\x55\x0d\xc4\x08\xd4\xed\x93\xbc\x78\x2e\x56\x1a\xa9\xb2\xdd\x94\x2c\xaa\x34\x8b\xf1\x70\x81\x69\x60\x86\x37\x33\x60\x47\x16\xff\xbd\x1a\x6c\xe7\x25\xad\xf2\x56\x5c\xee\x9c\x73\xee\x77\xcf\xc5\x0b\xaf\x09\xa7\x13\xd2\xdd\xc5\x9e\x26\xc6\x64\xd7\x6b\xe3\x10\xb1\x20\xb4\xce\x08\xad\xc6\x90\xb1\x20\xac\xa5\x6b\x86\x7d\x2a\x74\x97\xfd\xc3\xc5\x8b\xc8\xfa\xfa\x18\xb2\x98\xb1\x2c\xc3\xdd\x7a\xb3\x7e\xf8\x5a\xac\x73\xac\xfe\xc2\x6e\x7b\xb7\x4d\x91\x6f\xb1\xd9\x16\x58\xe7\xf7\x45\x3a\x63\xf2\x15\xa4\x85\x6b\x08\x42\x77\x9d\x56\x90\xca\x91\xa9\xb8\x20\x54\xda\xa0\xe4\x8e\xef\xb9\x25\xe8\x9e\x0c\x77\x52\x2b\x0f\xe6\x0e\x82\x2b\xec\x09\x83\xa5\x12\x07\xe9\x1a\xaf\xe5\x5e\x7b\xb2\xa8\x8c\xee\x60\x45\x43\x1d\xc7\x8d\xaf\xe2\xcf\xb3\x3d\x4d\x37\x29\xcb\x32\x0f\x2c\x1a\x69\x61\x1b\x3d\xb4\x25\x0e\xda\xbc\xcc\x0a\x6f\xb1\x32\xfb\xad\x4d\xf3\x15\xb8\x2a\x3f\xfa\x8a\x63\xca\x7c\x8c\x39\xe9\xb7\x34\x4f\x2c\x58\x1f\x49\x44\xd6\x19\xa9\xea\x04\x69\x9a\xbe\x3d\x9e\xa6\x18\x51\x5f\x1f\xd3\x5f\x75\xd7\x71\x55\x16\xbc\x4e\x40\xc6\x68\x13\xb3\xe0\x8f\x81\xcc\xeb\xe7\xb4\x5b\xcf\x7b\xd0\x07\xfb\x1f\xc6\x83\x3e\x7c\x4a\xba\x72\xd8\xc4\xce\x99\xee\x74\xad\x61\x9d\x19\x84\xc3\x89\x01\x05\xf1\xee\x0b\x6e\xfd\x27\xf1\xa0\x6a\x50\x02\x1b\x3a\x44\xe5\x1e\xf9\x2a\xc6\xed\x4c\xf0\x48\x43\x6e\x30\x0a\x3f\x7b\x87\xff\xbf\x72\x1d\xf1\x2e\x2a\xf7\x71\xc2\x80\xc9\x4b\x64\x19\x7e\xd7\x35\x7a\xa3\x47\x59\xd2\x79\x96\xad\xae\x31\x4b\xcf\xf3\xd9\xbf\xa2\x26\xe5\xe7\x47\x25\xbe\x0d\x64\x24\xd9\x94\x8d\xdc\xcc\xc4\xe5\x8c\xfc\xb4\xa4\xd3\xc4\x98\xd0\xca\xfa\xd5\x03\xb2\x0c\x8f\xbd\x25\xe3\x72\xbd\xd1\xae\x91\xaa\x46\xae\xa1\x2e\xa6\xac\x7c\x74\x43\x37\x16\x1c\x42\xab\xaa\x95\xc2\x31\xfc\x8f\xb2\x44\x78\x5e\xc4\xdf\xee\x37\x77\xe1\x47\xd9\xc7\xbe\xe4\x8e\xb0\x23\x53\x69\xd3\x81\x2b\x0c\x67\xcf\xa1\x21\xf5\x63\xf9\x0b\xf5\xac\xfe\xb8\xcb\xbf\x16\xeb\xf9\x14\xe6\x4e\xd8\x56\x0a\x8a\x2a\x49\x6d\x69\xd1\xf1\xfe\xe9\x5c\xf0\xf3\xbb\x62\x13\xe8\xaa\xb2\xe4\xfc\x72\xc5\x88\x04\x9e\x9e\xaf\x5d\x91\xef\xec\x11\x4f\xcf\x1f\x5b\xc4\x02\x85\x2f\xcb\x2b\x7b\x81\x5f\x58\xe0\x8f\x47\xe8\x36\xc1\xc8\x5b\xff\x66\xb8\xaa\x09\x97\xf0\x27\x16\x04\x02\x4b\xf0\xbe\x27\x55\x46\x22\x41\xf8\x77\x18\x2e\x84\x6e\x17\xde\x88\x59\x10\xc8\xef\xcf\x32\x41\xf8\x53\xb8\xb8\x5c\x7d\x7a\xef\x34\x8f\x54\xec\x41\xe3\x77\xd0\x38\x47\xf2\x4e\xb5\x58\xb0\x60\x62\xc1\x65\x81\x44\x02\x99\x60\x64\x13\xfb\x37\x00\x00\xff\xff\x15\x6a\x43\x63\x5a\x04\x00\x00")

func templatesPogoGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPogoGoTpl,
		"templates/pogo.go.tpl",
	)
}

func templatesPogoGoTpl() (*asset, error) {
	bytes, err := templatesPogoGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/pogo.go.tpl", size: 1114, mode: os.FileMode(420), modTime: time.Unix(1501953138, 0)}
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
	"templates/enum.go.tpl": templatesEnumGoTpl,
	"templates/model.go.tpl": templatesModelGoTpl,
	"templates/pogo.go.tpl": templatesPogoGoTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"enum.go.tpl": &bintree{templatesEnumGoTpl, map[string]*bintree{}},
		"model.go.tpl": &bintree{templatesModelGoTpl, map[string]*bintree{}},
		"pogo.go.tpl": &bintree{templatesPogoGoTpl, map[string]*bintree{}},
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

