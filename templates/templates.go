// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/enum.gotmpl
// templates/many.gotmpl
// templates/model.gotmpl
// templates/pogo.gotmpl

package templates

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _templatesEnumGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x61\x6b\xdb\x30\x14\xfc\x3c\xfd\x8a\x23\x04\x6a\x87\x4c\xfe\x3e\xe8\x87\x76\xf1\x4a\x61\xc4\x63\x09\x85\x31\xc6\x90\x9d\xe7\x54\xcc\x96\x3d\xc9\x0e\x04\x4d\xff\x7d\xd8\x72\x16\x1b\x02\x65\xd0\x16\x9a\x4f\x41\xba\x77\xef\xde\xdd\x93\xad\x8d\x16\x78\x10\x5a\x8a\xb4\x20\x83\x45\xe4\x1c\xb3\x16\xf3\x4f\xba\x95\x0d\x3e\x5c\x83\xc7\xaa\x2d\xf9\x5a\x94\x84\x3f\xa8\x85\xc9\x44\x81\x01\x93\x5f\xc4\x64\xa2\xa4\x1e\xc2\x6a\x91\xfd\x12\x7b\x02\xa9\xb6\x64\x2c\x8a\x70\x17\xaf\xe3\xaf\x37\xdb\x78\x85\xdb\x6f\xf8\x92\xdc\x25\x1c\xab\x04\xeb\x64\x8b\x78\x75\xbf\xe5\x8c\x75\x6a\x9e\xeb\xe7\x27\x89\x16\x48\x5a\xdd\x4b\x40\x73\xac\x09\xff\xce\x9f\xb3\x4f\x37\x9c\xb5\xde\x34\xe7\x20\x0d\x9a\x47\xc2\xcc\xda\xf9\xd9\x1a\xe7\x66\x23\x19\xb9\xae\x4a\x0f\xd8\x64\x8f\x54\x8a\x13\x84\xb3\xfe\x7a\x44\x66\x1a\x2d\xd5\xfe\x0d\x7b\x93\x55\xca\x34\x08\x18\x60\xed\x7b\x68\xa1\xf6\x84\xf9\x41\x14\xdd\xe6\x0c\xfe\x3c\x88\xa2\x25\xd3\x2d\x0d\xd0\xc1\x30\xbf\xa9\xeb\x82\x7a\xc4\x41\x14\xfc\xb3\x48\xa9\x98\xec\xdf\x09\x27\x2e\xe0\x86\xeb\x49\x26\xd6\x7a\xca\x73\x3a\x57\xd6\xfa\x6a\xe7\xae\x46\x40\x3e\x50\x5f\x28\xbc\x1e\x1d\x07\xb3\x73\xf9\x2c\x1c\x66\x23\xb5\xeb\x9a\x87\x2f\x14\xd6\x46\x34\xd2\xe4\xc7\x5e\xbd\xf9\x5d\x44\x3b\x2d\x0f\xa4\xbd\x79\x1a\x52\x35\xa4\x73\x91\xbd\xe0\x8a\xf7\x9d\x60\x7a\x19\x92\xcc\x93\x42\xf2\x4a\x4f\x9c\xcd\x5b\x95\x21\xb0\xd6\x7f\x39\x9c\x1b\x5d\x86\x9e\x3c\x08\x11\x8c\xd9\x96\x20\xad\x2b\x1d\xc2\x32\x40\x53\xd3\x6a\x35\x3c\x88\x11\x4d\xb8\x84\x92\x05\x73\xaf\x60\x7b\xbd\xef\x1e\x8a\xe1\xb1\xca\xaa\x1d\xdd\x4a\x25\xf4\xf1\x35\x9c\x9f\xf4\xdb\x4c\x02\x78\x42\xd2\x7f\x65\x30\xe6\x08\x7e\x62\xe1\xb9\xf9\xc7\x4a\xa9\x7b\x95\x57\x4b\xa4\xf8\xfe\x23\x3d\x36\x14\x22\xf0\x7f\x2e\x05\x24\xea\x9a\xd4\x2e\x48\x97\x03\x78\x1c\x15\xe7\xfc\x14\xd7\xbb\xbf\x01\x00\x00\xff\xff\x3b\x56\x37\x75\x7c\x06\x00\x00")

func templatesEnumGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesEnumGotmpl,
		"templates/enum.gotmpl",
	)
}

func templatesEnumGotmpl() (*asset, error) {
	bytes, err := templatesEnumGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/enum.gotmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1, 0x80, 0x5c, 0xe4, 0x7a, 0x54, 0x28, 0x97, 0x98, 0x42, 0xcd, 0x72, 0xfd, 0x46, 0x3c, 0xf2, 0xe6, 0x53, 0xd4, 0x46, 0x73, 0x27, 0xa8, 0xec, 0xb0, 0xf, 0x14, 0x9c, 0xdb, 0xa, 0xca, 0xa2}}
	return a, nil
}

var _templatesManyGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x57\xdf\x53\xdb\x3e\x12\x7f\xae\xff\x8a\x3d\x4f\xe7\xce\xa1\xa9\x73\xcf\xdc\xe4\x81\x12\x97\xa3\x43\x03\x4d\x42\x3b\x37\x4c\x07\x14\x7b\x13\x54\x6c\xc9\x95\xe4\x06\xc6\xf8\x7f\xbf\x91\x64\x3b\x72\x08\x90\x42\xbf\x4f\xdf\xa7\x38\xde\xdd\x8f\xf6\xc7\x67\x77\xe5\xb2\x1c\xec\xc1\x57\x22\x28\x99\xa7\x28\x61\x6f\x50\x55\x5e\x59\xc2\xdb\x73\x89\x02\xf6\x87\x10\xce\xb4\x20\xfc\xcc\x13\x4c\xa1\x96\x15\xdb\x64\xf7\x20\x29\x5b\x16\x29\x11\x70\x0f\x31\xc9\xba\xea\xd2\xd1\x1f\x93\x0c\xe1\x1e\xf2\xb4\x10\x24\x75\x95\xbd\x9c\xc4\x37\x64\x89\x50\x96\xb5\xd1\x3d\xa4\x7c\x85\x42\xcb\x06\x03\x38\x8a\xc6\xd1\xe4\x60\x16\x8d\xe0\xc3\xff\xe0\xec\xf4\xe8\x34\x84\xd1\x29\x8c\x4f\x67\x10\x8d\x8e\x67\xa1\xe7\xd1\x2c\xe7\x42\x41\xe0\x01\xf8\x31\x67\x0a\x6f\x95\xaf\x9f\x91\xc5\x3c\xa1\x6c\x39\xf8\x21\x39\xf3\xbd\x37\x3e\x0a\xc1\x85\xd4\x4f\x52\x09\xca\x96\xe6\x51\xd1\x0c\x7d\xcf\x7b\xe3\x2f\xa9\xba\x2e\xe6\x61\xcc\xb3\xc1\x0f\x12\xdf\xc4\x83\x7c\x79\xab\x71\x74\x30\xe1\x59\xed\x63\x55\x81\x5f\x96\x50\x1f\x59\x55\xbe\xd7\x33\x4e\x46\x42\x94\xa5\x49\x5f\x55\x8d\xb9\xfa\xc8\x0b\x96\x80\x40\x55\x08\x86\x09\xd0\x05\xa8\x6b\x1d\xa0\x93\x8b\xaa\x02\x2a\x81\x71\x05\x0b\xad\xec\xfd\x22\x62\x3b\xca\x10\xac\xdf\xe1\x18\x57\x81\xbf\x89\xd1\x02\xf8\xd6\x91\xd6\xfe\x98\xe5\x85\x82\xcc\x54\x69\xc1\x85\x76\xfb\x6d\x38\x8d\xaf\x31\x23\xb5\xad\x1f\x9a\x77\x2e\x9c\xef\xa9\xbb\x1c\x37\x41\xa4\x12\x45\xac\xa0\x34\xc9\x78\x0f\x82\xb0\x25\xc2\xdb\x98\xa7\x4e\x79\x0f\x79\x5a\x64\x4c\xea\x82\xda\x94\xc5\x3c\x6d\x6a\xde\x94\x1a\xf6\x1a\xc1\x11\x9f\xe9\x73\x6a\xe5\xf7\x80\x2c\xd1\x7f\xaa\x6e\x08\xaf\xf5\xfe\x25\x8e\xbb\x7e\xe7\x44\xc6\x24\xad\x2a\x1d\x0f\x5d\xd4\x21\x71\x35\x2e\x52\x1d\x4d\xad\x6b\x43\xd1\x7f\x01\x53\xa9\x63\xda\xdb\x22\xb1\xf1\x6d\x09\x76\x8c\xab\x1d\x83\x83\x83\xb3\x63\x6f\x51\xb0\x58\xdb\x04\x3d\x9d\xcc\x4e\x95\x74\x94\x96\x71\xf0\xcf\xae\xa8\x34\x87\x95\xe5\xb3\x09\xd0\xe5\x89\x32\x42\xad\xb8\x93\x85\x46\x8c\x9b\xe2\xb6\x8f\x6d\xe5\x8c\xb9\x4e\x3d\x2a\x69\x48\xaf\x03\x41\xfb\xd2\xb7\xee\x07\x75\xa3\x5b\x42\xb8\x8e\xf6\x1c\x84\x60\x6d\x07\x1b\x09\xdd\x1e\x7b\x8b\x1a\x3a\x96\x43\x93\x8b\xfa\xdf\x3a\x41\xad\xae\xa7\xa3\x6a\x4b\xa2\x63\xf8\x4c\x84\xbc\x26\xe9\xa7\xe9\xe9\x18\x32\xfb\xbc\x0e\xc4\x1a\xf9\x40\x99\xe2\xa0\x55\x9e\x8f\xc8\xc1\x0b\x7a\x10\x5c\x7c\x9f\xdf\x29\xec\xdb\x9e\xee\xb9\x45\xd3\x53\x2a\xac\xb5\xd7\x80\xbd\x9a\x27\xe7\x2c\x73\x1c\x2b\x58\xeb\x9a\x36\x03\xc5\x81\xb8\x0e\x3e\xef\x56\x07\x2f\x48\x88\x22\x60\x5d\xeb\x59\xd7\x1e\x78\xd6\x1a\x18\xe5\x3e\x74\x3d\x7c\xf6\xbc\xa9\x99\xb8\x41\x0f\xec\xe8\x75\xe1\x5d\xbf\x1b\xa8\x25\xaa\x9a\x97\x4f\x81\x66\x24\xbf\xb0\x78\xdf\x29\x53\x28\x16\x24\xc6\xb2\x32\xd8\x71\xcd\xea\xfd\x21\x64\xe4\x06\x83\xed\xaa\x3d\xcf\x4e\xab\xdd\x46\x5a\xcb\xfd\x6d\xc3\xcd\x03\x3d\x23\xb6\x92\xf0\x1f\x43\x60\x34\x35\x6e\xb5\x8e\x5d\xf8\xce\xac\xa9\x2a\xff\x3b\x0c\x4d\x7c\x0f\xac\x3d\x80\xfa\xfc\x86\xa4\x6d\xe2\x6a\xa8\x9a\x21\xdf\xae\x51\xe0\x61\x4a\x0a\x89\x7a\xb3\x90\x66\xfa\x29\x0e\xd7\x84\x25\x29\xc2\x4a\x6b\x40\x6c\x54\xa4\x9d\x95\xae\x91\x33\x2d\x63\xce\x12\xaa\x28\x67\x75\xb9\x3c\x80\x9c\x08\x92\x49\x1d\xc1\x85\x9b\x41\xf7\x70\x90\x39\xc6\x74\x41\xd1\x36\x4c\x0b\x22\x6d\x4d\x8d\x4e\xb0\x09\xdd\x6f\x90\xc3\x30\x74\x2b\x03\x7b\xae\x6f\xee\x70\x73\xde\x37\x29\xad\x21\xf7\xd7\x8f\x7d\x23\xb1\xd0\xfb\xeb\xc7\xbe\x49\xa7\xf5\xf9\x98\x49\x14\xaa\xdb\x38\xb6\xb3\x77\xdc\x36\x26\x2a\x8b\x12\xc4\xea\x16\xea\xcb\x47\x78\x68\x7f\xfb\x90\xcc\x35\x21\x9a\xcb\x43\x55\x85\xa3\x0f\x4e\xe3\xd8\xd1\xf5\x80\xd4\xc1\xfa\x4d\x67\x4c\x0c\x06\xba\x2b\x80\xa4\xa9\xc9\x2e\xe3\xec\xbd\x66\xd5\x82\x62\x9a\x48\x20\x2c\x81\x5c\x60\x4e\x04\x6a\x71\x66\xb6\xa6\xd6\xfb\x59\xa0\xb8\xf3\x00\x2e\xe3\x3e\x5c\xd2\x3e\x5c\xfe\xd2\x0c\xee\xba\x35\x4d\x69\x8c\xc1\xb6\x9e\xb3\x3e\xf5\xe1\xdf\xa6\x53\x06\x03\x90\x3f\x53\xa0\x36\x71\x06\xb9\x0f\xb9\xa0\x19\x11\x77\x70\x83\x77\x90\x0b\xfe\x8b\x26\x98\xc0\xfc\x0e\x24\xfe\x2c\x90\xc5\xe8\x81\x36\x92\xca\x5c\x1e\xaf\x3c\x80\xe3\xf1\x34\x9a\xcc\xe0\x78\x3c\x3b\xdd\x75\xf5\x05\x57\xf0\xae\xe6\x8b\x0c\x3f\x71\xca\x02\x1d\x8e\xdf\x07\xbf\x07\xef\xe0\xaa\xe7\x01\x7c\x3d\x38\x39\x8f\xa6\xdb\x34\xe9\x86\xe6\x24\x9a\x9d\x4f\xc6\xc7\xe3\x23\x58\x1f\x34\xfd\x72\x32\x31\xf4\x32\x0d\x77\x65\xd7\x89\x93\xa1\x13\xbe\x0c\x6c\x14\x3a\x83\x61\x18\x36\xf9\x10\x05\xeb\xa4\x59\x5f\xe3\xd6\x13\xab\xad\xa4\xa6\x2f\x5f\xe9\x0c\x24\xf3\xf0\x8b\xd6\x9d\xf0\xd5\x26\xa2\x99\x20\xa8\x95\x04\x5f\x85\xd3\x98\xb0\x60\xed\xe1\x11\xd7\x2f\xaa\xaa\xf7\x1f\xc0\xee\x44\xa9\xfb\x82\xd1\xb4\x0f\x68\xf8\xdd\xbd\x09\x14\x35\x97\x18\x4d\xeb\x7b\x00\x5d\xac\x67\x5c\x96\x73\x49\x15\x36\xfb\x3d\xce\x65\x67\x04\x3a\xe2\xc1\x00\x3e\x52\x96\xd4\x0d\x73\x5e\x37\xcc\xfc\x0e\xa8\x92\x66\x45\xe7\x32\x1c\xa1\x8c\x05\xcd\x75\x07\x56\x95\x6d\x10\x6d\xf3\x7b\xed\xa1\x81\x8e\xf8\x99\x69\x58\xbd\xed\x9f\xe8\x08\x4d\x47\x89\x29\xc6\x2f\xa7\x23\xc0\x34\x3a\x89\x0e\x67\x5d\x36\x4c\x0d\xa8\x29\x1c\xc0\xc7\xc9\xe9\xe7\x2d\x54\x0d\x1f\x8e\x04\xad\xfd\xed\xbf\xd1\x24\x6a\xe2\x98\x7e\x39\x31\xf3\x6a\x07\x5e\xb5\x91\x7f\x25\x42\xc7\xfd\x27\x19\xf6\x00\xfb\x85\x5c\x33\x46\xc3\x21\xe4\xcb\xdb\x30\x12\x62\xcc\x27\x7c\x25\x6b\x59\x97\x89\xdb\xbe\x65\x8c\x5a\xf5\x02\xd2\x3a\xdb\xef\x35\xf4\x3d\xcf\x13\xa2\xf0\x77\x09\x6c\xad\x5e\x43\xe1\x97\x0f\xfd\x7a\xba\xef\x0f\xb7\xde\x88\xac\x71\xcd\x93\x84\xb3\x7f\x29\x28\x6c\x88\x9a\x30\x0b\x2e\x90\x2e\x99\xee\x04\xb9\xf1\x05\x94\xcb\xe6\x76\x63\x88\x93\x60\x8a\x0a\x03\x7b\x58\x5f\x67\xa7\x21\x74\xaf\xfb\xcd\x62\x4f\x72\xf6\x0c\x48\xbd\x3a\xe4\x4e\xfb\xa5\x81\xaf\xd3\x33\x2e\xb2\xd6\x07\x77\xbb\x34\x3c\x77\xba\xf4\xfc\x6c\x74\x30\x8b\x76\x5d\x16\xd3\x68\x66\xd6\x80\xe1\xd9\x13\x4b\x03\x86\x8f\xa9\x6d\x6e\x8c\xa7\x7a\xfa\xd9\x7d\x72\xd5\xc4\x86\xaa\xc8\x37\x3a\x39\x2d\xd0\x14\xb7\x73\xa5\x2a\xab\xa7\xab\x55\x9b\x0d\x81\xe4\x39\xb2\x24\xb0\xff\x75\x5e\x3b\x77\x52\xdb\xe5\x4e\xf1\x1e\xb7\x6c\xf7\xcf\xe3\xc3\xc9\xaa\xfe\xe9\xc5\xe7\xa0\xfe\x2d\x07\xd2\xc8\x34\xde\xee\x03\x29\xb4\x13\xc9\x9a\xbd\x6e\xa9\xae\xbf\xf2\x36\xfa\x2e\xe6\x4c\xaa\xa6\xfb\x9a\x15\x39\x8a\x4e\xa2\x59\xf4\xd8\x22\xdc\x7a\x3b\xfe\x8b\x56\x61\xe3\x26\x5d\xc0\x65\xdf\x12\x26\x99\x87\xd1\x2d\xc6\x8f\x5a\xbf\x8c\x34\xbb\xf1\x65\x93\x2a\x1b\xf4\xf8\x7f\x00\x00\x00\xff\xff\xc2\x6e\xb8\xec\x69\x15\x00\x00")

func templatesManyGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesManyGotmpl,
		"templates/many.gotmpl",
	)
}

func templatesManyGotmpl() (*asset, error) {
	bytes, err := templatesManyGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/many.gotmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9a, 0xcb, 0x48, 0x19, 0xe9, 0x27, 0xff, 0x7f, 0x26, 0xce, 0xea, 0x66, 0xac, 0xe5, 0xff, 0xb7, 0x78, 0xa7, 0x8, 0x4e, 0xd6, 0x66, 0x91, 0xee, 0xe7, 0x1, 0x96, 0x9e, 0x20, 0xc1, 0x81, 0x59}}
	return a, nil
}

var _templatesModelGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x73\x9b\x48\x12\x7f\x0e\x9f\xa2\x8f\x72\x25\xc8\xab\xa0\xdb\xd7\x5c\xe9\xc1\xb1\x88\x4f\x7b\x8e\xe4\x58\x72\xf6\xae\x5c\xa9\x68\x0c\x23\x99\x35\x9a\x21\x33\x60\x5b\xa5\xf0\xdd\xaf\xe6\x0f\x30\x48\xe8\xbf\xbd\xbb\xd9\xf5\x93\x10\xd3\xd3\x74\xf7\xfc\xfa\x37\x3d\x0d\xf3\x79\xeb\x18\x3e\x23\x16\xa2\x9b\x08\x73\x38\x6e\x65\x99\x35\x9f\xc3\xd1\x15\xc7\x0c\xde\xb5\xc1\x1d\x8a\x01\xb7\x87\xa6\x18\xbe\x03\x0f\xc9\x24\x8d\x10\x83\xef\x10\x23\xee\xa3\x08\xb4\x78\xba\x49\xdc\x47\x53\x5c\x91\xe6\xcb\xe2\x71\x94\x32\x14\x99\xc2\x56\x8c\xfc\x3b\x34\xc1\x30\x9f\xeb\x49\xdf\x21\xa2\x0f\x98\x89\xb1\x56\x0b\xce\xbc\x9e\x77\x79\x32\xf4\x3a\xf0\xfe\x7f\x70\xd1\x3f\xeb\xbb\xd0\xe9\x43\xaf\x3f\x04\xaf\xd3\x1d\xba\x96\x15\x4e\x63\xca\x12\x70\x2c\x00\xdb\xa7\x24\xc1\x8f\x89\x2d\xae\x31\xf1\x69\x10\x92\x49\xeb\x37\x4e\x89\x6d\xbd\xb2\x31\x63\x94\x71\x71\xc5\x13\x16\x92\x89\xbc\x4c\xc2\x29\xb6\x2d\xeb\x95\x3d\x09\x93\xdb\xf4\xc6\xf5\xe9\xb4\xf5\x1b\xf2\xef\xfc\x56\x3c\x79\x14\x7a\x84\x33\xee\x85\xb6\x31\xcb\xc0\x9e\xcf\x41\x3f\x32\xcb\x6c\xab\x21\x8d\xf4\x18\x9b\xcf\x65\x40\xb3\xac\x47\x93\x0f\x34\x25\x01\x30\x9c\xa4\x8c\xe0\x00\xc2\x31\x24\xb7\xc2\x41\x23\x16\x59\x06\x21\x07\x42\x13\x18\x0b\x61\xeb\x1e\xb1\x7a\x2d\x6d\x50\x76\xbb\x3d\xfc\xe0\xd8\x8b\x3a\x0a\x05\xb6\x32\xa4\x98\xdf\x25\x71\x9a\xc0\x94\x06\x38\x82\x31\x65\xc2\xec\x23\x77\xe0\xdf\xe2\x29\xd2\x73\x6d\x57\xde\x33\xd5\xd9\x56\x32\x8b\xf1\xa2\x12\x9e\xb0\xd4\x4f\x60\x2e\x83\xf1\x16\x18\x22\x13\x0c\x47\x3e\x8d\x8c\xe5\x3d\xa5\x51\x3a\x25\x5c\x2c\xa8\x0a\x99\x4f\xa3\x7c\xcd\xf3\xa5\x86\xe3\x7c\xe0\x8c\x0e\xc5\x73\xb4\xf0\x5b\xc0\x24\x10\x7f\xb2\xaa\x0b\x87\x5a\xbf\x8f\xe1\xa6\xdd\x0a\xfe\x59\x26\xfc\x09\xc7\xda\x25\x9a\xf4\xd2\x48\x78\xa3\x65\x95\x2b\xe2\x2f\xe0\x88\x0b\x9f\x8e\x6b\x46\x94\x7f\x35\xce\xf6\xf0\xc3\x96\xce\xc1\xc9\x45\xd7\x1a\xa7\xc4\x17\x73\x9c\x86\x08\x66\x65\x95\x84\x97\x0a\x71\xf0\xba\x3a\x34\x97\x0f\x9b\xcf\x37\x06\x40\x2c\x8f\x37\x45\xa1\x1a\xae\x44\x21\x1f\xc6\x8b\xc3\x45\x1e\xab\x95\x93\xd3\x45\xe8\x71\xc2\x25\xe8\x85\x23\x58\xdd\xb4\x95\xf9\x8e\x4e\x74\x05\x08\xd3\xd0\x86\xa1\xc1\x29\xe7\xc1\x42\x40\xeb\x7d\x2f\xb4\xba\xc6\xcc\xb6\x8c\x85\xfe\x57\x06\xa8\x90\xb5\x84\x57\xc5\x92\x08\x1f\x3e\x22\xc6\x6f\x51\xf4\xcb\xa0\xdf\x83\xa9\xba\x2e\x1d\x51\x93\x6c\x08\x49\x42\x41\x88\x6c\xf6\xc8\xd0\xe7\x34\xc0\xb9\xfe\x72\x33\x4b\x70\x53\xe5\x74\xc3\x5c\x34\xc1\x52\xae\x96\x2e\x15\x36\x34\x4e\xae\xc8\xd4\x30\x2c\x25\x85\x69\x62\x1a\x24\x14\x90\x69\xe0\x66\xb3\x2a\xfa\x9c\x00\x25\x08\x94\x69\x0d\x65\xda\x92\x65\xc5\x04\x29\xdc\x84\xaa\x85\x1b\x9f\x37\x90\x8c\xeb\x34\x40\x51\xaf\xa9\xde\xb4\x3b\x57\x35\xc1\x89\xc6\xe5\x3a\xa5\x53\x14\x5f\x2b\x7d\x5f\x42\x92\x60\x36\x46\x3e\x9e\x67\x52\xb7\xaf\x51\xfd\xae\x0d\x53\x74\x87\x9d\x7a\xd1\x86\xa5\xd8\x6a\x3b\x4a\x2b\xb0\x5f\x47\x6e\x16\x08\x8e\xa8\x05\xe1\x3f\xda\x40\xc2\x48\x9a\x55\x18\x76\x6d\x1b\x5c\x93\x65\xf6\x17\x68\x4b\xff\x96\x66\x5b\x00\xfa\xf9\x39\x48\x8b\xc0\x69\x55\x1a\x21\xbf\xde\x62\x86\x4f\x23\x94\x72\x2c\x76\x16\x94\xb3\x5f\x42\xe1\x16\x91\x20\xc2\xf0\x20\x24\xc0\x97\x22\x5c\x71\xa5\x39\xc9\x60\x4b\x9f\x92\x20\x4c\x42\x4a\xf4\x72\x59\x00\x31\x62\x68\xca\x85\x07\xd7\x66\x04\xcd\x87\x03\x8f\xb1\x1f\x8e\x43\xac\x12\xa6\x50\xc2\xd5\x9a\x4a\x19\x67\x51\x75\x33\xd7\xec\xba\xae\xb9\x32\x70\x6c\xda\x66\x92\x9b\x71\x3f\x0f\xa9\x56\xf9\xae\xbc\x6c\xca\x11\xa5\xfa\x5d\x79\xd9\x94\xe1\x54\x36\x77\x09\xc7\x2c\xa9\x26\x8e\xca\xec\x3c\xdd\xb7\xd8\x71\xa4\x67\x4a\x93\xe3\x27\x8f\xa0\x0b\x10\xf7\x54\xfd\x36\x21\xb8\x11\xa0\xc8\x0b\x88\x2c\x73\x3b\xef\x8d\xe4\x51\xf4\xb5\x04\x6c\xa7\xbc\x53\xa1\x8a\x56\x4b\x64\x06\xa0\x28\x92\x36\x12\x4a\xde\x0a\x64\xe5\x60\x47\x24\x80\x98\xe1\x18\x31\x2c\xc6\xa7\x72\xeb\x14\x82\xdf\x52\xcc\x66\x16\xc0\x57\xbf\x09\x5f\xc3\x26\x7c\xbd\x17\x30\xae\xda\x35\x88\x42\x1f\x3b\x75\x89\xa7\x8c\x6a\xc2\x3f\x65\xba\xb4\x5a\xc0\xbf\x45\x10\xaa\xe8\x49\xcd\x4d\x88\x59\x38\x45\x6c\x06\x77\x78\x06\x31\xa3\xf7\x61\x80\x03\xb8\x99\x01\xc7\xdf\x52\x4c\x7c\x6c\x81\x98\xc4\x13\x59\x40\x8e\xe4\xda\x74\x7b\x03\xef\x72\x08\xdd\xde\xb0\xbf\xed\x0e\xe8\x8c\xe0\x27\x0d\x1b\xee\xfe\x42\x43\xe2\x08\x87\xec\x26\xd8\x0d\xf8\x09\x46\x0d\xa9\xf7\xf3\xc9\xf9\x95\x37\xa8\x93\x0d\x97\x64\x2f\xbd\xe1\xd5\x65\xaf\xdb\x3b\x83\xf2\x61\x83\x4f\xe7\x97\x12\x69\x32\xf7\x46\x6a\x67\x31\xe2\x74\x4e\x27\x8e\xf2\x45\xc4\xd1\x75\x5d\x19\x15\x51\xbe\x95\x4c\x55\xac\x9e\x80\x2d\x7d\x10\x4e\x07\x37\xee\x27\x11\xab\x4b\xfa\xb0\x38\x5d\x32\x07\x16\x42\x8c\x3e\xb8\x03\x1f\x11\xa7\x34\xe7\x8c\x8a\x1b\x59\xd6\xf8\x17\xe0\x2a\x93\xe8\x7c\x20\x61\xd4\x04\x2c\x71\x5d\xad\x00\x52\x8d\x1f\x12\x46\x7a\xff\x0f\xc7\x39\xb7\x5d\xa8\xf5\xfa\x0f\x9e\xe5\x1b\x7b\x7c\x67\x50\xdf\xf2\x70\xc9\x7d\xf1\x5d\xdd\xd6\xff\x21\x24\x01\x20\x18\x15\x9e\x8f\xc4\xf2\x87\xc9\x1b\x5e\xc1\x86\x18\xd7\x0a\xb2\x6c\xa4\xb2\x47\x4c\xdd\x2d\x77\x8c\xba\x20\xbe\x33\xca\x82\x35\x69\x23\x20\xcb\x71\x84\xfd\x43\x20\x3b\xf0\xce\xbd\xd3\x61\x59\xc7\x0f\x3e\x9d\x0f\xa4\x4e\xb9\xd0\x00\x1f\x2e\xfb\x1f\xb7\x25\x0e\x21\xff\xeb\xbf\xbd\x4b\x4f\x4e\x28\x62\x62\x43\x1b\x8e\x7e\xde\x08\x3c\xa3\xaa\xd9\x1f\x7e\x15\x25\x7b\x82\x50\x4e\x6a\xb7\x21\x9e\x3c\xba\x1e\x63\x3d\x7a\x49\x1f\xb8\x1e\xab\x42\xb4\xee\x70\x23\xc5\xb2\x3d\xd0\x6c\x6c\x87\xe5\xf6\x1d\x06\x8f\x06\x86\xbb\x24\xc0\x8f\x58\xd7\xb5\x6f\x85\xa1\x8e\xe0\x47\x21\xe5\x76\xf9\x15\x09\xbf\xa5\x18\x1c\x71\x7e\xd2\xb7\x34\xe8\x1b\x0d\x03\xd2\xef\x67\xf3\xb9\x1c\x3e\xa3\x1f\x71\x72\x4b\x83\x2c\x83\xb1\x82\x7a\x19\xee\x9b\x19\x68\xa9\x0e\xe6\x3e\x0b\x63\xb1\x01\x65\x59\x89\xee\x1a\x2d\xbb\x01\x5e\xcd\xbd\x90\x5b\xd8\xef\x0a\xf4\xa3\x27\x45\xba\xf6\x64\xf0\xe9\x5c\xee\xe1\x5b\x10\x6c\xe1\xfb\x67\xc4\xf8\x21\x58\x5f\x52\x04\x22\x6e\x9b\xe0\xae\xb3\x82\xb1\x65\xd0\x33\xb6\x1d\xec\x77\xc2\x3d\x63\x5b\x20\xbf\x38\xae\x18\x49\xa0\xe1\xda\x27\x58\xc1\x93\x12\x5c\x05\x28\x2a\x4b\xa3\x12\x97\x7d\x82\x77\xc0\xa1\x2a\x1f\xcd\xb2\xec\xb9\x81\xb8\x9e\x6f\x77\xc3\xa0\x42\xa0\x28\x0a\xa4\x1f\x6e\x19\x8e\x75\xf8\x53\xb2\xaa\x76\x3c\x68\xab\x5f\x52\xf4\xec\x7c\xfb\xa4\x84\xfb\xaa\x80\x98\xaa\xf4\x25\xc8\xa6\x88\xcc\xcc\xe2\xf9\x0d\x57\x48\x9b\x84\xf7\x98\xd4\xe1\x4d\x9f\x00\x0e\x43\xdc\xf5\x97\x15\x98\xcb\xfb\x87\x59\x26\xa2\x6a\x8a\xcd\x33\xeb\x2f\x8e\x49\x46\x1f\x78\x33\x67\xb4\x1c\x7f\x9b\xc0\xb7\x48\x6b\xd5\xe6\x08\xd7\xe1\xd5\x07\xd0\x00\x8f\x31\x93\xcf\x71\x4f\x23\xca\xb1\x23\x9f\x2b\x0e\x18\xf2\x5e\x0f\x3f\x26\x4e\x43\x2b\x5a\x9d\x21\x15\xd4\xf3\xed\x61\xbf\x09\xf8\xb5\xc6\xaf\x4c\x80\x3c\x05\xd6\xb8\x9c\x8b\x18\xa0\x6a\x03\x8a\x63\x4c\x02\xc7\x94\x7e\x6d\xf4\x41\xd4\x9c\x70\xac\x7c\xf3\x18\x73\x1a\x1b\x03\x5c\x8a\x2e\xa4\xa0\x29\xa4\x2b\x78\x9d\x81\x27\x51\xa4\x6b\x90\x28\xaa\xa6\x5f\x99\x69\x27\x51\xb4\x75\x9e\xfd\x85\x72\x6a\x5d\x31\xb1\x31\x4f\x5e\xd2\xe2\xc7\x4c\x8b\xe7\x3d\xd8\x5e\xc5\x01\x4a\xf0\x52\x97\x68\xdb\x96\x85\x3c\x04\xf3\x85\x33\x9e\xca\x53\xa5\xf9\x29\x4e\xbe\xfb\x37\x94\xc6\x21\x8e\x02\xd9\x25\x5d\xdd\xf8\xd1\x79\x1e\x50\xf2\x26\x81\x54\x85\x23\xb9\xc5\x66\x9a\xcb\x44\x88\x70\x82\x1d\xa5\xb0\xb9\xe0\x70\xae\xc3\x68\x4e\x01\x8f\x42\x1f\xf3\xad\x9a\x52\xb9\xd2\x9f\xcd\x06\x54\xde\xd3\x32\xa8\xe4\xea\xa2\x73\x32\xf4\xb6\x5d\x9a\x81\x37\x94\x5d\x22\x8d\xf9\x95\x7d\x25\x39\x3e\x6a\x40\x7b\xb5\x74\xb8\x24\xbd\xf1\x7c\xbf\xb1\xf9\x34\xca\x7d\xc5\x49\x1a\x17\xde\xde\xa3\x28\xc5\x72\xc1\x74\xd6\x55\xba\xb1\x73\xe3\x48\x9f\x19\xfd\xa5\xd5\xc5\x84\xd2\x97\x97\x11\xad\x16\xb0\x94\x54\x5a\x86\x7b\x55\xbb\x86\xd6\x1f\xab\xce\xdd\xa6\xb1\xa0\x5a\x07\xbc\xc2\x0f\xeb\xfb\x04\x3a\x64\x46\xd6\x3f\x71\x47\xe0\x8f\x21\x80\x3b\x3c\xe3\xd5\x57\xab\xc2\x2a\xad\x24\xcb\x6a\x49\xa1\x64\x84\xca\xcb\xd0\x27\xa1\x07\x1d\x96\x5e\x3a\x2d\x6c\x78\x3e\xc2\x58\xd3\x86\x36\x99\x62\x4d\x07\x1a\x72\x99\xd1\x8a\xee\x48\x29\xb0\x13\x55\x54\xb3\x37\xa7\x8b\x2a\x4f\x64\xeb\xd7\x4d\x4f\x2b\x48\x46\xfd\x17\x11\xae\x6c\x91\x6a\x77\x37\x96\x71\xf5\xcc\xbf\x05\x15\x3d\x17\x13\x15\x34\xa3\x0e\xdf\xa2\x38\xda\xad\x04\x59\x71\x26\x37\x94\x1e\x74\x2a\xdf\x82\x7e\x0e\x39\x61\xec\x49\x0b\xeb\x5e\x65\x45\x98\x38\xe6\xa9\xb8\xf1\x23\x11\x45\x4d\xc3\x40\xc8\x3d\x33\x5d\xac\xca\xec\x9a\xe6\xc2\x53\x93\x40\x6e\xde\xfa\x2e\xc7\x52\xa6\xbf\x1c\xe4\x5e\x0e\x72\xe5\x41\xae\x23\x6b\x91\x85\x83\xdc\x98\xd1\xe9\x0e\xaf\xfb\x21\x11\xff\x14\x79\x2a\x7d\x4f\xf3\xe2\xb2\xfc\xe4\x66\x35\x05\x75\xbc\x73\x6f\xe8\xed\xd2\x0d\x59\x7d\xfe\xd8\xf8\xd2\xc5\x7c\xb7\xb8\x90\x84\xe1\x18\xbe\x36\x15\xc6\x83\x1b\xd7\x7b\xc4\x7e\xdd\xc4\xfd\x76\xd1\xed\x36\xd0\xc5\xbd\x73\xcf\xca\x5d\x2d\x60\x5d\xe5\xae\xca\x56\x5e\xf7\x4d\xd5\xca\x49\x87\xbd\xd0\x5b\x02\x80\x32\xe1\x49\x71\xb0\x54\x5f\xee\xfe\xea\x6d\xfd\xda\x57\xc4\xff\x50\x00\x14\xab\xab\x0a\x26\x1d\xcc\xfa\xf7\x15\x22\xfb\x6b\xab\x23\x43\xc3\x81\xef\x2c\x96\x56\xf7\x90\x7e\xe9\xde\xeb\xff\x14\xef\x19\xd6\x00\x60\x51\x7c\xd5\xf7\x2a\xf5\x8b\x77\x00\xf9\xcb\xea\x58\x7e\x8c\xb4\xf8\x99\x66\x6d\xf3\xcd\xcd\x6b\xdf\xbf\xdf\x37\x5b\x35\x6c\xf2\x0c\x5f\x61\xe9\x3a\x74\xb4\xe5\x87\x58\xb9\x78\xbf\x07\xa7\xfd\xde\x87\xf3\xee\xe9\x10\x9c\x85\xf6\x61\x2e\xd3\xe9\x83\xae\xc4\xf3\x12\x7b\x53\x79\x0d\xde\x7f\x4f\xcf\xaf\x3a\x5e\xc7\x5d\x25\x5c\x08\x2c\xd8\xb3\xb1\x88\xde\xee\xc3\xb0\x85\xdd\xf3\x4f\xff\x9d\xd8\x1e\x0d\x30\x91\x48\x6b\x1a\x60\x76\xe1\xa5\x6d\xa6\xde\xc1\x7b\xe8\x4b\x36\xfe\x5e\xd9\x58\x56\x0e\x65\x63\xed\x25\x25\xf7\x4a\xc9\xd7\xaf\xd5\x75\x5d\x15\xb4\x67\xa2\xfe\x3f\x00\x00\xff\xff\x6f\x31\x34\x27\xdf\x35\x00\x00")

func templatesModelGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModelGotmpl,
		"templates/model.gotmpl",
	)
}

func templatesModelGotmpl() (*asset, error) {
	bytes, err := templatesModelGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/model.gotmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3c, 0x73, 0x7b, 0xaf, 0x27, 0xa2, 0x22, 0xe3, 0x93, 0x6e, 0x9c, 0xdb, 0xf0, 0xeb, 0x15, 0xb4, 0x4, 0xc4, 0x7, 0xfb, 0x8a, 0x50, 0x85, 0x4, 0x37, 0xf5, 0xc4, 0x97, 0xae, 0xcc, 0xa4, 0x13}}
	return a, nil
}

var _templatesPogoGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4d\x6f\xe3\x36\x10\x3d\x8b\xbf\xe2\xc1\x28\x5a\x7b\x6d\x48\xe8\x75\x81\x1c\x36\x95\x91\x06\x58\xd8\xde\xc4\x29\x50\x04\x01\x4c\x53\x23\x99\x88\xc4\x51\x48\x4a\x4e\x6a\xf8\xbf\x17\x94\x3f\x36\xd9\x22\xdd\x4b\xa4\x13\x31\x9c\x37\x7c\xf3\xf8\x86\xda\xed\x92\x4f\x1f\xf5\x25\xfb\xbd\x08\xf5\xf0\x97\xb4\x5a\xae\x4b\x72\x38\xc7\x3e\xf2\x0c\x21\x3e\xbc\x64\x47\x7b\xde\x58\x2c\xa4\x7a\x94\x05\xf5\x44\xbc\x3e\x56\xdf\xed\x10\x9f\x4e\x0a\xf1\x5e\xda\x59\x70\xc1\xa8\xa4\x7d\x24\xdb\x53\x3b\x49\x82\xab\xe9\x6c\x7a\xf3\x65\x39\x4d\x71\xf9\x37\x16\xf3\xab\x79\x8c\x74\x8e\xd9\x7c\x89\x69\x7a\xbd\x8c\x7b\x6a\x2d\x95\x5e\xae\xa5\x23\x68\xe3\xc9\xe6\x52\x11\xdc\x86\x9b\x32\xc3\x96\xed\x23\xb6\xda\x6f\x90\x1d\x73\x12\xf7\x54\xc6\xe9\x25\x7e\x7d\x1b\x59\x3e\xf7\x27\x4a\x7a\x09\xed\xe0\x37\x04\xc5\x55\xc5\xe6\x15\xcd\x9c\xed\x99\x07\xb8\x26\x2b\xbd\x66\x13\x92\xa5\x87\x92\x06\x6b\x42\xe3\x28\xeb\x7a\x08\xb5\xfc\x4b\x4d\x0e\xb9\xe5\x0a\x4e\x6d\xa8\x92\x58\x05\xf7\xdc\x76\xeb\x78\x26\xab\xe0\xa0\x55\x2c\x92\x24\x64\x2f\x37\xda\xfd\x5c\x0a\x69\xb2\x1f\xc5\x88\x45\x38\xa8\x63\x7e\xe6\xba\x13\xd1\xf4\x99\xd4\xd0\x79\xab\x4d\x31\x41\x1c\xc7\xe7\xcd\xdd\x7e\x84\x61\x5d\x3c\xc7\x7f\x70\x55\x49\x93\x2d\x65\x31\x01\x59\xcb\x76\x24\xa2\x6f\x0d\xd9\x97\xf7\x61\x9f\x02\xee\x86\xb7\xee\x07\xc4\x0d\x6f\xdf\x05\x9d\x30\xa2\xb7\x69\x69\xd6\xa5\x56\x50\x8d\xf3\x5c\xe9\x7f\xc2\xf3\x85\x92\x8b\x42\x9b\xe2\x95\x24\xbd\x59\xe6\x2b\x17\xa8\x2d\xb7\x3a\xa3\x83\x73\x4a\x2e\x90\x37\x46\x1d\xdc\xb0\x7e\x41\x41\x26\xb8\x85\x32\x3c\x35\x64\x35\xb9\x58\xb4\xd2\x76\xc0\x8b\x2e\xf3\x5d\xed\x76\x3d\x6b\xc6\xf5\xc1\xc3\xc1\xdb\x4d\xed\xc8\x7a\x87\xe1\x72\x9e\xce\x3f\xa3\x92\x8f\x04\x1f\x4c\x29\x0d\xc8\x34\xd5\xa8\x27\x09\x15\x1b\xe7\x31\x14\x40\x92\xe0\xae\x23\x91\xf2\x8c\xfd\x26\xdc\x5f\xca\x30\xc7\xa5\xce\x83\xba\x96\x7e\x73\x90\x50\x6c\xf2\x52\x2b\x2f\xf0\x1f\xc8\x05\x06\x87\x77\xec\xcf\xeb\xd9\xd5\xe0\x6d\xd9\xbb\x3a\x93\x9e\xb0\x20\x9b\xb3\xad\x42\x67\xcd\x21\xb2\xdd\x90\xf9\xff\xf2\x47\xe8\xa1\xfa\xdd\x22\xfd\xb2\x9c\x0e\xc4\xa8\xdf\xeb\xd9\x50\x59\x93\xed\x2c\x12\xee\x09\x9e\xe1\x4a\xad\x08\xdc\x58\xe4\x9a\xca\xcc\x05\x87\x33\x6e\xbf\x7d\x45\x6e\x35\x99\xac\x7c\x81\x36\x75\xe3\xfb\xfa\x81\x27\x09\x6e\x3b\x0a\x8a\x4d\xdb\xf9\x25\x70\x51\x5c\x36\x95\x39\x92\x71\x5c\xd1\xe1\x2a\xc2\x34\xb8\xa7\x12\x99\xd5\x2d\xd9\xee\x91\x6c\x4c\x46\xd6\x79\x69\x32\xd1\xcd\x48\x57\x6c\x78\xc2\x57\xb2\xbe\x3f\xcc\xc2\xc3\xab\x39\x98\x80\xf3\xdc\x91\x0f\xe5\x47\x18\x2a\xdc\x3f\x9c\x06\x46\xbf\x5a\xb7\xb8\x7f\x78\x3b\x3d\x22\x32\xf8\x7c\x71\x42\x8f\xf1\xbb\x88\x82\xd3\x15\x97\x13\xb4\xb2\x0c\x7b\x56\x9a\x82\xce\xfc\x77\x22\x8a\x14\x2e\x20\xeb\x9a\x4c\x36\x54\x13\xac\x06\xab\xb1\xe2\x72\xbc\x1a\xac\x46\x22\x8a\xf4\xf7\x4d\x3d\xc1\xe0\x97\xc1\xd8\x79\x1b\xa4\x88\xaf\x3d\xcb\xa1\x19\x85\xa4\xf6\x7b\x52\xdb\x1d\x14\x82\x66\x3c\x16\xd1\x5e\x44\x96\x7c\x63\x0d\xd4\x04\x7a\x82\x56\xec\xc5\xbf\x01\x00\x00\xff\xff\xa5\xd7\xa5\x26\xc8\x09\x00\x00")

func templatesPogoGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesPogoGotmpl,
		"templates/pogo.gotmpl",
	)
}

func templatesPogoGotmpl() (*asset, error) {
	bytes, err := templatesPogoGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/pogo.gotmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5d, 0xeb, 0xc6, 0xd9, 0xb1, 0xeb, 0x52, 0x4d, 0x2f, 0x29, 0x7e, 0x2b, 0x49, 0x32, 0x8e, 0x1a, 0x84, 0x29, 0x28, 0x34, 0x5c, 0xd7, 0xd6, 0x27, 0x4b, 0x11, 0x1d, 0x5e, 0xa4, 0xee, 0xc6, 0x47}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"templates/enum.gotmpl": templatesEnumGotmpl,

	"templates/many.gotmpl": templatesManyGotmpl,

	"templates/model.gotmpl": templatesModelGotmpl,

	"templates/pogo.gotmpl": templatesPogoGotmpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
		"enum.gotmpl":  &bintree{templatesEnumGotmpl, map[string]*bintree{}},
		"many.gotmpl":  &bintree{templatesManyGotmpl, map[string]*bintree{}},
		"model.gotmpl": &bintree{templatesModelGotmpl, map[string]*bintree{}},
		"pogo.gotmpl":  &bintree{templatesPogoGotmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
