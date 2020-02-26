// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// api.txt
// api_test.txt
package data

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _apiTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x5f\x6f\x9b\x30\x10\x7f\xc6\x9f\xe2\x26\xf5\x21\x54\x15\xbc\x67\xe2\x21\x6b\x37\x69\x0f\xab\xba\x66\x7b\x9a\xa6\xc9\x81\x4b\x82\x66\x8c\x67\x8e\x25\x29\xe2\xbb\x4f\xb6\x81\x90\xa4\xa4\x4d\xd3\x6a\x6f\xf8\xb8\xbb\xdf\x9f\xf3\xc9\x61\x38\x29\x29\x87\x05\x4a\xd4\x9c\x30\x81\x55\x4a\x4b\xf8\x82\xc4\x27\x2a\x85\x25\x91\x2a\xc6\x61\xb8\x48\x69\x59\xce\x82\x38\xcf\x42\x5c\x6f\x1e\x1e\x36\x61\x86\xc4\xb9\x4a\x99\xe2\xf1\x6f\xbe\x40\xa8\x2a\x08\xee\x9a\xef\xba\x66\xcc\x9c\x3f\x67\x2a\xd7\xd4\x1e\x35\x97\x0b\x84\x8b\x54\x26\xb8\xbe\x82\x0b\xe2\x33\x81\x30\x8e\x20\xf8\x66\xbe\x0a\x93\x16\x86\xd7\x1a\x39\x21\xd8\x10\x9b\x97\x32\x06\x17\xb1\x81\xaa\x6a\xca\x82\x6b\xae\x6e\x79\x66\x90\x46\xc9\x0c\x2e\x8b\x3f\x22\xb8\xf9\xe0\xc3\x08\xb5\x06\xd4\x3a\xd7\x3e\x54\xcc\xfb\x75\x65\x0e\x10\x41\x32\x0b\x3e\xae\x31\xde\xd6\xdf\xe8\x5c\xd9\x96\x53\xe2\x84\x19\x4a\xcb\xd2\x4b\xe7\xb6\xe0\x5d\x04\x32\x15\xa6\x83\xa7\x91\x4a\x2d\x99\x57\x1f\xeb\xd6\xa3\xb8\xdb\xaf\xa9\xae\x19\x0b\x43\x03\xd9\xd7\xd5\x51\x78\x73\x55\x7d\x16\x53\xd2\x65\x4c\x8c\x36\xca\x8e\x6c\x0b\x3c\x4d\xe5\xa2\x01\x87\xc2\x26\x41\xc5\xb6\x19\xae\xee\x53\x8a\x22\xb1\x93\xb2\xcd\x9c\x6e\x27\x67\xd4\xcb\xed\xb5\xba\x1c\x02\xf1\x9b\xc1\x0e\xfd\xdf\x75\x40\x63\x51\x0a\x1a\x64\xec\x1c\xd9\x5a\x54\x50\x46\x2e\x36\xb6\x36\xdd\x69\x54\x5c\xe3\xfe\xc4\x4e\x18\x7e\x82\x73\xd4\x60\xfa\x06\xd7\x22\x2f\x70\xe4\x33\xcf\x4d\xc1\xc6\xbe\x96\xa8\x37\xf7\xf9\x6a\x1f\xc1\xc6\x0f\x86\x70\x8f\xa4\x53\xfc\xfb\x62\xe7\xda\xfa\x37\xf1\xae\x49\x8f\x06\x0b\xaa\xba\x95\x9e\xcc\x1e\x11\xde\x92\x1b\xbe\x83\x6d\xc6\x44\x08\xe7\x40\x2f\xf0\xf4\x32\x6c\x33\x5a\xa7\x7e\xfc\x7c\xb6\xb6\x7c\x55\xf4\xef\x85\x65\x7f\x48\x7d\x22\xc4\x09\x57\x63\x9e\x6b\x30\x8d\x83\x5b\x5c\xd3\xc8\x6f\x7e\x5b\x13\xc7\xc7\x5d\x6c\xfb\x46\xb6\xbe\x77\x0d\x62\x2e\x27\x42\x40\x5d\xbf\xdf\xc7\xed\x80\x0d\xb2\x77\xe8\x45\x04\x5c\x29\x94\xc9\xa1\x4d\x57\xe0\x48\xf9\x96\xb4\x25\xdc\xdd\xe4\xfe\x74\xbe\xab\xe4\x8c\xa5\x76\xd5\xff\x67\xa9\x1d\xf6\x79\x4b\xfd\xc4\x56\x3b\x88\xc7\xb7\xfa\x06\x05\xbe\xdc\x38\x57\xfd\x3c\xe3\x4e\xf4\xc5\xb5\x3e\xd7\x97\xee\xd9\xb1\xf1\xbd\x87\xc7\x22\x1c\xb3\xa5\xdb\xf5\xee\x38\x24\xb5\x78\x05\xad\xfd\x05\x7e\x2d\xb5\x3b\x7b\x52\x55\x80\x32\x31\x5a\xff\x05\x00\x00\xff\xff\x74\x45\xe3\x61\x40\x09\x00\x00")

func apiTxtBytes() ([]byte, error) {
	return bindataRead(
		_apiTxt,
		"api.txt",
	)
}

func apiTxt() (*asset, error) {
	bytes, err := apiTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.txt", size: 2368, mode: os.FileMode(420), modTime: time.Unix(1581361184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _api_testTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x6d\x73\xe3\xb6\x11\xfe\x2c\xfe\x8a\x0d\x27\xce\x90\x39\x1d\x95\xbb\xb6\x5f\x94\x2a\x33\x8e\xad\x9b\x5c\x1b\xbf\xd4\xd2\x25\x9d\xba\x9e\x1b\x88\x5c\xc9\xa8\x29\x80\x06\x40\xcb\x3e\x8d\xfe\x7b\x67\x01\x92\xa2\x2c\xea\xc5\xf2\x75\x3a\xc9\x87\x58\xc4\xcb\xe2\xd9\xc5\xb3\x0f\x16\x48\x3a\x9d\xe3\xdc\x48\x98\xa0\x40\xc5\x0c\x26\x30\xe3\xe6\x16\xce\xd0\xb0\xe3\x8c\xc3\xad\x31\x99\xee\x76\x3a\x13\x6e\x6e\xf3\x51\x14\xcb\x69\x07\x1f\x9f\xbe\x7c\x79\xea\x4c\xd1\x30\x96\x71\x2f\x63\xf1\x1d\x9b\x20\xcc\xe7\x10\x5d\x16\xbf\x17\x0b\xcf\xe3\xd3\x4c\x2a\x03\x81\xd7\xf2\x13\x66\xd8\x88\x69\xec\xe8\xfb\xd4\xf7\x5a\x3e\x8a\x58\x26\x5c\x4c\x3a\xff\xd1\x52\x50\xc3\x78\x6a\xe8\x4f\x2a\x27\xf4\x47\x6a\xfa\xb7\xc2\x71\x8a\xb1\x6d\xd7\x46\x71\x31\xb1\xad\x06\xb5\xe1\xc2\x0e\x33\x7c\x8a\xbe\x17\x7a\xde\x03\x53\x40\xed\xa7\x23\xf8\x5e\xdf\xa7\xd1\xe9\xcf\xb6\x29\x96\x62\xcc\x27\xc9\x08\xa6\x2c\xbb\x76\x26\x6e\xb8\x30\xa8\xc6\x2c\xc6\xf9\xc2\x8b\xa5\xd0\xa6\x98\x78\xce\xa6\x08\x3d\xb0\xe6\xc9\x91\x0f\x3c\xc5\x4b\x85\x63\xfe\x08\x8b\x85\xef\x79\x9d\x0e\xf4\xdc\x3f\x70\x8b\x69\x86\x4a\x53\x1b\xd3\x3a\x9f\xa2\x06\x56\xac\x95\xca\x98\xa5\xc9\x28\x22\xb7\x60\xcc\x53\x04\xa6\xbb\x5e\xa7\x33\xa7\xf9\x00\xe0\xff\x22\xb5\xf1\xbb\xe0\xdb\x81\xb7\xf4\xd1\x2e\xbb\x2e\xa5\xb2\x5d\x7f\xf9\xf3\x9f\xde\x2f\x5b\x3f\x69\x54\xd4\x9a\x8c\x04\x9b\x62\x6d\x34\xd3\xba\xa9\x9d\xfc\x68\x6a\x1f\x0c\x7e\x3d\x93\x89\xeb\xe2\x9a\x8d\x52\xf4\xbd\x4e\x67\xe1\x8d\x73\x11\x43\x2a\x59\x72\x62\x1d\x08\x42\x98\x7b\xad\xf1\xd4\x44\x97\x8a\x0b\x93\x8a\xc0\x87\x5a\xb7\x1f\x7a\x2d\x72\xab\x0d\xa8\x14\x74\x7b\x20\x75\x74\x91\xa1\x08\xfc\x75\xff\x69\x2c\x1f\xdb\x81\xdf\xf4\x40\xf0\x94\x2c\xb7\x52\x39\x89\x3e\x30\xc3\x52\x32\x7d\xc2\x84\x90\x06\x64\x86\x62\x35\x80\x36\x76\xbe\x5d\x25\xf4\x5a\x0b\xaf\x95\x60\x2c\x13\xb4\x4b\x92\xed\xe8\x1c\x67\xa7\xae\x29\xa0\xa1\xa1\xd7\xa2\x75\x7a\x50\x8c\x8b\x5c\x67\xf0\x5d\x49\x81\xbd\xb1\x4c\xd0\x80\x05\x51\x00\xca\x15\x33\x5c\x0a\x82\xa4\xe4\xf4\x39\xae\x85\xe7\x02\x18\x2b\x64\x06\x4f\x47\x41\x52\x31\xb0\x0d\x89\x23\x95\x23\x5e\x1b\xe4\x4c\xa0\x2a\xbe\x42\x08\x08\x0b\x2a\x25\x95\x8d\xb8\xd6\xe4\x1a\xc5\x7d\x90\x51\xe0\xc7\x81\x7f\x72\xd5\x3f\x1e\xf6\xe1\xf4\x78\x78\xfc\xf3\xf1\xa0\x0f\x47\x1a\x2e\x7e\x3f\xef\x5f\xc1\x91\xf6\x4b\xe3\x85\xd5\x70\x6d\xcb\x7c\x78\x03\x5a\x87\x5e\xeb\xb3\xdb\xaa\x1e\x24\xa3\xa8\xff\x88\x71\x60\x5b\x15\x9a\x5c\x89\x0a\xbe\x46\x33\xfc\xb2\x82\x7e\x1f\x80\x83\xfe\x10\x86\x1f\xcf\xfa\xf0\xaf\x8b\xf3\x3e\x7c\x1a\x9e\xf8\xaf\xc5\x91\x28\x99\x6d\x0b\xe2\x3e\xa8\x4e\xaf\x2e\x2e\xeb\x41\xab\x62\xf5\x5a\x70\x4a\xce\xfa\x8f\x5c\x1b\xbd\x82\xef\x3e\x47\xf5\x54\xed\x31\x53\x13\x0d\x51\x14\xd5\x54\x86\x30\xdb\x69\x30\x92\x32\x75\x0b\x2d\x1d\x70\xd3\xd7\x02\xfb\x6b\xff\x64\x08\xfd\x7f\x7e\x1c\x0c\x07\x10\x1c\xe9\xd0\x2f\x16\x6a\xf6\xa1\xec\xaa\x7c\xf8\x07\x35\x5c\xc9\x59\x60\x7b\x1c\xac\x28\x8a\xc2\x68\x10\x33\x11\x7c\xe7\xf0\xac\x7b\x68\x48\x16\x1a\x7c\xb4\xed\xcb\x2d\x78\x60\x29\x4f\x9a\xbc\xf1\x5a\xb6\xab\x0c\x66\x3d\x60\x6d\x28\xbd\x7a\x07\x1f\xae\x2e\xce\x20\x9b\x7c\xb6\x66\x35\xfc\xfe\x4b\xff\xaa\xef\xd6\x10\x4e\x84\xbf\x7d\xe7\x17\x8b\xae\x63\xe4\x82\x9b\xa1\x55\xec\x60\x8d\x0d\x75\x19\xf3\x5a\x99\xbe\x4f\x3f\x8a\xb1\x5c\x63\x08\xc9\x6e\xef\x48\x03\x9d\x4e\xf4\x37\xd7\xa8\xec\x37\xd3\x7a\x26\x55\x42\xbf\xfd\x37\x5e\xab\xe5\x6b\x9d\x4e\x65\x82\x3d\xcb\xa2\x52\x4a\xae\x9d\x88\xdf\xd4\x5b\xac\x76\xaf\xb4\x58\xdd\x5e\x1d\x43\x8a\xbd\xd2\x52\x6a\xf2\x4d\xe8\xb5\xdc\x29\x54\x86\x8e\x22\xef\x74\x35\x93\xda\x4c\x14\x12\x80\xd2\xa1\xf5\xa0\x50\x40\xce\x18\x17\xc1\x14\xbe\x2f\xce\xc7\xe8\xcc\x86\xa4\xd3\xa1\x6f\x4a\xf0\x3c\x73\x04\xe9\xf6\x56\x62\xb8\x49\x17\x2f\x99\xe0\xf1\x38\xf0\x63\x27\x8b\xcb\x29\x96\x71\xa8\x54\xd4\xa7\xb8\x07\xa1\xd5\xc1\x92\x7b\x95\x0e\x96\xee\x2c\x0f\xd7\xf5\xe0\x44\x41\x41\xa9\x7d\x41\x9c\x14\xd6\xb7\x42\x28\xb5\xcc\xad\xbc\xaf\xe9\x62\xd6\x06\xcb\x9d\x8e\xca\x85\xf5\x45\x7b\x2d\x7c\xe4\xe6\x37\x96\x52\x20\xa7\xd1\x55\x2e\x82\xd0\xab\xe2\x6c\x90\xa9\x44\xce\x44\x95\x8b\x4e\xd0\xd6\xa3\xb1\x2f\xae\x53\x6b\xa0\x11\x56\x4b\xea\xa8\xff\xc8\x4d\x50\x00\x0a\x89\x0e\xe6\x29\x43\x88\xe5\x34\x63\x0a\x87\xf4\x9b\xf8\x11\xd4\xe4\xa8\x0d\x2b\xda\x44\x59\x5c\x90\x48\xc8\x13\x37\x2f\x50\xa8\xf3\xd4\xb4\x01\x1f\x33\x8c\xcd\xfa\x84\x95\x12\x61\x1c\xf8\xd5\xcc\x2e\x1c\x3d\xb4\xe1\xe8\x01\xde\x02\x1c\x0d\xdb\x70\x34\x84\x7f\x0b\xbf\x0d\xab\x06\x9f\x7f\x57\x7c\x86\xc0\xa8\x1c\xc3\xe5\x71\x80\x63\x96\xa7\xe6\x50\x58\xab\xd3\x5f\x8d\xcd\x75\x40\xaf\x57\x75\x95\x38\xa9\x22\x39\x14\x64\x6d\xee\xa1\x08\x89\x7d\x45\xc1\x89\x69\xa2\x21\x66\x02\x46\x08\x4c\x3c\x81\x54\x54\x34\xb1\xb1\x41\x05\xc9\x08\x9c\x2b\x6d\xd0\x12\x14\xb2\x84\xd0\xc9\xc6\xe2\x18\x98\x48\x20\x95\xf2\x0e\x48\x36\xdc\x82\x67\x2c\xb3\x9c\x67\x77\x18\x34\x17\xd4\x61\xa5\x30\xb6\x46\xfb\x24\xa6\x4c\xe9\x5b\x96\x06\xd7\x37\xa3\x27\x53\xc6\x66\x99\xf6\x6d\xf8\xae\x32\xbd\x3d\x21\x82\xaa\x10\x74\x4e\xef\x8d\x65\x13\x14\x67\x66\x05\x4a\x65\x79\x4f\x28\x5e\x6b\x2c\x15\xdc\xb5\xe1\x81\xa0\x28\x26\x26\x08\x4b\x74\x34\x87\x8f\xe1\x81\x6c\x54\x4e\x5e\xdf\xdd\xd8\x8e\x15\x06\xfc\x1d\x9f\xdc\xce\x5f\xd9\x61\xee\x77\xdf\x1a\xa2\xdf\x7e\x9b\xd6\xa8\x9b\x68\xc3\x43\x48\x46\x0a\x62\x8e\x59\xaa\xd1\x6b\x51\x70\x16\x15\x5b\x29\x91\x96\xe5\x9d\x75\xf2\x5c\x9a\xfe\x7d\xce\xd2\xe7\x14\x75\x11\xa9\xd8\xb9\x24\x7c\x71\xe9\x8a\x86\x8a\x4f\x07\x19\x8b\xb1\x68\x29\x0c\x84\x21\xf9\xb6\x71\x50\xc1\xce\x30\xf4\x9e\xe1\x78\x41\xa6\x94\xd4\xfb\x9b\x96\xa2\xba\x72\xd8\x0d\x3d\x2b\xb6\xb3\x80\xf2\x22\xf2\x6c\xb6\x56\xe5\xfc\x5e\xd6\xea\x9b\xb8\xe2\xda\xa6\x44\x5e\x09\x1f\xa1\x08\xab\xb6\x25\xb2\x70\x8b\x04\xed\xd8\x11\x3b\x3d\x24\x85\xda\xb1\x2b\x6e\xa0\x15\xb0\x4e\x87\x1b\x7b\xed\x07\x73\xab\x64\x3e\xb9\x05\x64\xf1\xad\x93\x12\x90\x63\xb2\x94\xc7\xc6\xea\x01\xcb\xb2\xf4\x09\xcc\x6d\x75\xc0\xd8\xc3\x85\x6e\x47\x60\x64\x7d\x1a\xdd\xf4\x13\x90\x62\xe5\x20\x9a\xb2\xcc\xb1\x00\x89\x84\x1f\x68\xe0\x16\x0a\xb4\xed\x5c\x4a\xa4\x5a\x7a\xd7\xcc\x85\xae\xec\xb3\x1c\xc9\x6d\xfe\xb9\xc7\x82\xe8\x37\x96\xe6\x78\x31\x5e\x6e\xe5\x43\x53\x6f\x45\x1b\x3a\x2f\x2f\xc6\x03\x1a\x93\x47\x64\xd8\x9e\xe6\x94\xd8\x9c\xda\x7e\xf8\x11\x38\xfc\x15\xf2\xe8\x3c\x9f\x3a\xc8\xe1\x8f\xc0\xdf\xbc\xb1\xeb\x12\x4b\xbe\x09\x0a\x9c\xd7\x85\xa5\xc8\x0d\xe3\x61\x44\xc7\xfc\x4d\x18\x3c\x2c\x5b\x3e\x96\xfe\x05\x61\x1b\xf2\xc6\xf6\xd0\xe9\x43\x99\xda\x53\xe3\x8e\xfc\x71\xe0\xdb\xd1\x5d\x38\xd2\x5b\x84\xa2\x11\x43\x1b\x5e\x08\xe1\xb9\x96\x08\x9e\x12\x53\xbc\x4e\xa7\xb8\x08\x64\x18\xf3\x31\x8f\xc1\xf3\xe6\xf3\x42\xf8\xbe\xe5\x22\xc1\xc7\x36\x7c\xeb\x86\x74\x7b\x10\x0d\x5d\x79\xbf\x28\xdf\x59\xe6\xf3\xa2\xd7\xa2\x82\xc5\xc2\x7e\x94\xef\x2e\x6b\xbd\xf6\xcd\xc5\x16\x55\x09\x33\x0c\xde\x82\x90\x06\xbb\x90\xc8\x9c\x16\x18\x29\x16\xdf\xa1\xd1\xc0\x5d\x5d\xe6\x06\x09\xc4\x04\x34\xf1\x1d\x46\x68\x66\x88\x02\xa4\xb9\x45\x35\xe3\x1a\x81\x08\x6b\x19\x96\x29\x34\x98\x00\xd3\x60\x70\x9a\xa5\xc4\x7d\x66\x69\x5c\xbd\x23\x2d\xc1\x9c\xb0\x6c\x40\xda\xe9\x30\x41\x0f\xae\xdf\xdf\x6c\xea\x9d\xd7\x5c\xb4\xb5\x32\x41\x7a\x07\x8b\x45\xbb\xa9\xe3\x3d\xd9\x5b\xb8\xb7\xab\x3c\x4b\x98\xc1\x2d\xab\x6e\x5c\xb2\x79\x41\x0a\xdd\xf3\x1c\xd5\xc5\x93\x98\x6d\x5d\xb1\xb7\x5c\xa6\x39\xd7\xe6\x5e\x6d\xb8\x6b\x3e\x63\xd9\x07\x57\x6c\x2c\x16\x4e\x45\xaa\xd7\x31\x5b\x25\x77\x37\x22\x2e\x2f\xd3\xf8\x80\x4a\x37\xe2\x08\xd6\xa8\x00\xd7\x1b\x63\x1e\x56\x75\xd9\xb6\x31\xf3\x7a\x4e\xa7\x28\xd6\x97\x08\xe1\x2d\xbc\xa3\x64\xff\xc9\x25\xfd\xdb\xb7\x36\x11\xcb\x9a\x8f\xb4\x0f\xc5\x52\xb0\xd6\xe6\x5f\xf3\x9b\xb0\x96\x33\x2b\x37\x33\x77\x6d\xb1\xe9\xd0\xe8\xaf\x59\x5e\xda\x86\xeb\xcf\x6f\xbd\xde\x8e\xf9\x7e\xe8\x55\x75\xd7\xce\xa5\x36\x5f\x8a\x96\x4a\xb3\x72\xdb\xda\x68\xab\xe9\x4e\x02\x98\x6a\xb4\xd6\x9e\xbd\x52\x9c\x4a\x81\xdd\x5d\x16\xfd\xf2\x90\xe6\xda\xe8\xea\x80\xae\x3f\x4a\x94\xf7\xa8\xcd\x52\xb2\x9f\x6b\x35\x9b\xcd\x57\x2b\x52\xf6\xe2\xd5\x66\xd5\x40\x1d\x4d\x03\x89\xdc\xee\x63\x02\x33\x25\xc5\x04\xb4\x61\x26\xd7\x10\xcb\x04\xbb\x30\x91\x86\x6a\x82\x19\x13\xc6\xa9\x75\xe9\xa8\xbb\xf4\xec\x8c\x5e\x6d\x69\x7f\xe5\xf1\x71\xc9\xb1\x4d\x19\xb0\x27\xc5\x36\x4d\xb7\x0c\xab\xce\xea\x62\x5b\xb6\xc8\xe4\xf5\x0f\x37\xd1\x2e\x40\x2f\x26\xe2\x46\x75\x3c\x94\x87\x5b\xbc\xad\x3d\x27\x34\x54\x2b\x3b\x5c\x6f\x6f\x13\xd9\x1d\x1e\xb3\x8c\x2f\x39\x94\x0b\x57\xc4\x60\x52\x54\x84\x51\x23\x59\xeb\x34\xb8\x42\xa3\x38\x3e\x1c\x4e\x84\x5d\x06\x0e\xa1\xc2\x6e\x50\x2f\x20\xc3\x2e\x63\x07\xd0\x61\x0f\x9f\xad\x2e\xfd\x81\xf9\x70\x9c\xa6\x07\x1d\x3d\x3b\xe6\x5b\x36\x7c\xde\x8f\x08\xef\xfe\xf8\x9a\xf0\x9c\xf8\x3b\xa3\x7b\x00\xaf\x37\xd8\x7a\x05\xad\x37\xef\x5d\xf9\xb6\xe9\x6a\x31\x18\x61\xcc\x72\x2a\x94\x33\x0e\x5c\xc3\x69\x7f\x70\xd2\x86\xeb\xee\x0d\x7d\xe8\x94\xc7\x48\xf7\x41\x96\xa6\xc0\x94\x62\x4f\x80\x29\x4e\x51\xb8\x17\x51\x7b\x7d\xb3\xf7\xac\xcd\x65\xdd\x56\x76\x74\xa9\x7c\xb2\x25\x5a\x1b\x3e\x3f\x7f\x4f\xb1\x3e\x6e\x48\xbf\x6b\x7e\x53\xde\x1f\xed\xcf\xed\xb9\xd6\xb0\x13\xaf\xca\x36\x1b\xc1\x85\x7d\xe1\x58\x26\xdd\xa7\xad\xd5\xfc\xce\x9c\xdb\x3e\xbd\x49\x80\xb7\x5f\x1f\xa2\x5d\x78\x5e\xc0\xd2\xed\xa6\x0e\x20\xe9\x4e\x67\xb7\x2a\xef\x76\xc7\xff\x9f\xc2\x7b\x8a\x29\xbe\x82\x03\xdb\xa7\xd7\x2b\xfe\x5d\x87\xef\x2e\x20\x2f\xd8\xfc\xed\xa6\x0e\xd8\xfc\x9d\x5e\xd2\xe6\x57\xff\x6d\xf6\xab\x97\x19\xbd\x97\x6f\xbc\xab\xe2\xaf\xe4\xcc\x95\xf1\xe7\xf2\x4a\xce\xfc\x55\x4f\x97\xc6\xf5\x7d\x4a\x66\xed\x20\xfd\xec\xdd\xf7\x25\xf5\x87\xfb\xff\x71\x84\x2c\x30\xf8\x56\x77\x96\x0b\x7e\x95\x1a\xc9\xbd\xf8\xac\x33\xf8\xd0\xc2\x61\xeb\xec\x3a\x7f\x77\x2c\xf3\x62\x7a\x7e\xbd\xf3\x73\x97\x0b\xaf\x28\x08\xfe\x57\x15\x81\xb5\x9b\xa2\x28\xdf\x37\xe1\x27\xf8\xe1\x6b\xf3\xfb\x90\x5a\xa3\x81\xc1\x44\xb6\xf9\x1c\x50\x24\xb0\x58\xfc\x37\x00\x00\xff\xff\xbf\xb8\x2e\x79\x82\x26\x00\x00")

func api_testTxtBytes() ([]byte, error) {
	return bindataRead(
		_api_testTxt,
		"api_test.txt",
	)
}

func api_testTxt() (*asset, error) {
	bytes, err := api_testTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api_test.txt", size: 9858, mode: os.FileMode(420), modTime: time.Unix(1582324886, 0)}
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
	"api.txt":      apiTxt,
	"api_test.txt": api_testTxt,
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
	"api.txt":      &bintree{apiTxt, map[string]*bintree{}},
	"api_test.txt": &bintree{api_testTxt, map[string]*bintree{}},
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