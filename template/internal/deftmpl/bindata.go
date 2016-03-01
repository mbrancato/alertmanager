// Code generated by go-bindata.
// sources:
// template/default.tmpl
// DO NOT EDIT!

package deftmpl

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

var _templateDefaultTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1b\x69\x6f\xdb\x38\xf6\xbb\x7e\xc5\x1b\x15\x8b\x69\x00\xcb\x4a\xd2\x03\x93\xc3\x59\xb8\x8e\xd2\x18\xeb\xd8\x81\xed\xb4\x53\x0c\x06\x05\x2d\xd1\x36\x5b\x49\xd4\x88\x54\x9c\xcc\xec\xfc\xf7\x7d\x8f\x92\x0f\xc5\x76\xe2\x16\xb3\x89\x77\x27\x49\x0f\xf3\x89\xef\xe4\xbb\x48\xd1\x7f\xfc\x01\x01\x1f\x8a\x98\x83\xfd\xf9\x33\x0b\x79\xaa\x23\x16\xb3\x11\x4f\x6d\xf8\xf3\xcf\x3a\x8d\x2f\xf2\x31\x4e\xe4\x71\x80\x40\x6b\x2d\xca\x55\xb7\x45\x58\xf8\xbc\xea\xdd\x68\x9e\xc6\x2c\x44\x10\x42\xdc\x17\xae\x99\xa7\xfe\x99\x72\x9f\x8b\x6b\x9e\xd6\x68\x52\xb7\x18\xe4\x38\x05\xf5\x32\x79\x95\x0d\xbe\x70\x5f\x13\xd9\x5f\x08\xa5\xa7\x99\xce\x14\xfc\x1b\xb4\xbc\x4a\x92\x29\xaa\x18\x02\xff\x6d\xf6\xd0\x1e\x8a\x54\xc4\x23\xc2\x39\x24\x1c\xa3\x85\xaa\x9e\x19\x28\xa2\x86\x3c\x5e\xe4\xf8\x2b\xd0\xa4\xf7\xa9\xcc\x92\x16\x1b\xf0\x50\x55\x7b\x32\xd5\x3c\xb8\x64\x22\x55\xd5\x0f\x2c\xcc\x38\x31\xfc\x22\x45\x0c\x36\x10\x55\xc8\x59\x8e\x34\xbc\x24\x5a\xd5\x86\x8c\x22\x19\xe7\xc8\x3b\x05\x6c\x81\xde\x0e\xa2\xbc\x44\x94\x89\xd0\xe3\xf2\x64\xb4\x40\x24\xaf\x79\x99\x7b\x9b\x45\xc8\x30\x37\xe3\x2a\xee\x33\xc1\x77\x66\x9f\xd6\xac\x4d\xc0\x95\x9f\x8a\x44\x0b\x19\xdb\xf7\xd8\x58\xf3\x1b\x9d\xaf\xe3\xe7\x50\x28\x5d\x4c\x4d\x59\x3c\x42\xc9\x70\x90\xcb\x75\x68\xcd\x81\xcb\x76\x22\xab\x38\xc6\x90\x24\x3e\x8d\x6a\x30\x53\xa0\x10\x2c\x67\x5e\x8f\x63\x89\xeb\x84\x32\x95\x48\x2e\x80\xbf\x8f\x6e\x4f\x66\xa9\xcf\x0f\xf3\xc5\xe4\x31\x4f\x99\x96\x69\xee\x7e\xd6\x0a\x43\x95\x6c\xa0\x42\xe6\x7f\xad\xe2\x88\x65\xa1\xae\x6a\xa1\x43\x5e\x58\x41\xf3\x28\x09\x99\x2e\xfb\x62\x75\x9d\xc9\xcb\x74\x32\x45\x21\x10\xad\x22\x55\x0e\xb4\x0d\xe9\x0d\x59\x18\x0e\x10\xb0\x44\x6f\xa5\xf8\x44\x14\x1d\xe7\xa1\x89\xa1\x88\xbf\x6e\x2c\x41\x92\x72\x72\x16\x7b\xb3\xd9\x0b\xf4\xef\x35\x80\x49\x1b\x1b\x4a\x20\x7c\x19\x63\xcc\x7c\x11\x9b\xca\xb0\x24\x6e\x69\xe1\xc7\x22\xf1\xc7\x4c\xcf\x4d\x9c\xca\xe8\xfb\x97\xeb\x2e\x35\x8c\x63\x85\x28\x9b\xbb\x52\x49\xb6\x84\xb8\x05\x99\xbe\x9d\xd1\x5b\x8e\xe7\x6f\x73\xcf\x65\x8a\x7e\x28\x78\xac\xbf\x5f\xe3\x75\x14\xe7\x95\xe0\xfb\x16\x7d\x99\xae\x88\x95\x66\xb1\xcf\xd5\x0a\xba\x4b\x09\xec\x1e\xab\xca\x44\x8d\x78\x2c\xf8\x5f\x66\xd4\x25\x82\xca\x24\xa2\x6f\x57\xbf\x24\x26\x8f\x98\x08\xe7\x24\xe7\x55\xf0\x9b\xe5\x2b\x53\x1a\xeb\x28\x24\x32\xd6\xf1\x0f\xa7\x9d\x46\xff\xd3\xa5\x07\x04\x82\xcb\xab\x77\xad\x66\x03\x6c\xc7\x75\x3f\xbe\x6a\xb8\xee\x69\xff\x14\x7e\x3e\xef\x5f\xb4\x60\xaf\xba\x0b\x7d\x4c\xd2\x4a\x90\x89\x58\xe8\xba\x5e\x1b\x0b\xd1\x58\xeb\xe4\xd0\x75\x27\x93\x49\x75\xf2\xaa\x2a\xd3\x91\xdb\xef\xba\x37\x44\x6b\x8f\x90\x8b\x8f\x8e\x5e\xc0\xac\x06\x3a\xb0\x4f\x90\xb3\xe3\x58\x3d\x7d\x1b\x72\x60\x28\xad\x61\x12\xf0\x14\x1b\x81\x00\x28\xfe\x80\x48\x2b\xa4\x3d\xc2\x7a\x99\x0d\xaa\xbe\x8c\x5c\xd2\x61\x94\xc5\xae\x21\xc7\xfc\x9c\x9e\x63\x54\x73\xa6\xe6\x50\x68\xc1\xfe\x98\xc3\x45\xb3\x0f\x2d\xe1\xf3\x58\x71\x78\x89\x83\x1d\xcb\x6a\xc8\xe4\x36\x15\xa3\x31\x16\x6c\x7f\x07\xf6\x77\xf7\x5e\xc3\x45\x4e\xd1\xb2\x2e\x79\x1a\x09\xa5\x90\x22\x08\x05\x63\x9e\xf2\xc1\x2d\x8c\x90\x0f\x96\xa0\x0a\x0a\xc4\x39\xc8\x21\x60\x5c\xa7\x23\x5e\xc1\xb6\x03\x85\xbe\x05\xec\x3c\x14\x22\xc8\x81\x66\x22\xa6\xae\x82\x81\x8f\x3c\x2c\x9c\xa9\xc7\x48\x46\xc9\xa1\x9e\xb0\x34\xd7\x90\x29\x25\x7d\x81\x12\x06\x10\x48\x3f\x8b\x30\x3e\x4c\xa1\x83\xa1\x08\xb1\xb8\xbf\xd4\x28\xb4\xdd\x2b\x30\xec\x1d\xc3\x24\xe0\x2c\xb4\xb0\xe4\xd3\xb3\xe9\x23\xd3\x40\xc8\x4c\x43\xca\x95\x4e\x85\xb1\x42\x05\x44\xec\x87\x59\x40\x32\x4c\x1f\x87\x22\x12\x05\x07\x42\x37\x8a\x2b\x0b\x89\x62\x41\xaa\x18\x39\x2b\x10\xc9\x40\x0c\xe9\x7f\x6e\xd4\x4a\xb2\x01\x46\xce\xb8\x02\x81\x20\xd2\x83\x4c\x23\x50\x11\xd0\xd8\xb1\x42\x7a\xb8\x32\x05\xc5\xc3\xd0\x42\x0a\x02\xe5\x36\xba\xce\xa5\x33\x73\x48\xf4\x84\x0c\xaa\x0b\x13\x29\x82\x4c\xc6\xb8\xaa\x25\x4d\x84\xb2\x86\x59\x1a\x23\x4b\x6e\x70\x02\x89\x26\x33\x1c\xc9\x9b\x09\x42\xd3\x87\x32\x0c\xe5\x84\x54\xc3\xac\x1f\x88\xa2\x67\x30\x8b\xcc\x06\xd4\x37\xf9\xb3\x75\xc5\xe6\x01\x45\xcd\x45\xa0\x05\x48\xe6\xab\x5a\x3c\x52\x63\x2c\x9f\x30\xe0\x85\xc1\x90\x2f\x9a\x97\x2d\xa8\x93\x12\x7b\x4a\x32\x5a\xb0\x10\x12\xec\x41\x88\xdf\x5d\x35\xab\xc8\xff\xdc\x83\x5e\xe7\xac\xff\xb1\xde\xf5\xa0\xd9\x83\xcb\x6e\xe7\x43\xf3\xd4\x3b\x05\xbb\xde\xc3\xb1\x5d\x81\x8f\xcd\xfe\x79\xe7\xaa\x0f\x38\xa3\x5b\x6f\xf7\x3f\x41\xe7\x0c\xea\xed\x4f\xf0\xaf\x66\xfb\xb4\x02\xde\xcf\x97\x5d\xaf\xd7\x83\x4e\xd7\x6a\x5e\x5c\xb6\x9a\x1e\xc2\x9a\xed\x46\xeb\xea\xb4\xd9\x7e\x0f\xef\x10\xaf\xdd\x41\x17\x6e\xa2\xef\x22\xd1\x7e\x07\x88\x61\x41\xaa\xe9\xf5\x88\xd8\x85\xd7\x6d\x9c\xe3\xb0\xfe\xae\xd9\x6a\xf6\x3f\x55\xac\xb3\x66\xbf\x4d\x34\xcf\x3a\x5d\xa8\xc3\x65\xbd\xdb\x6f\x36\xae\x5a\xf5\x2e\x06\x76\xf7\xb2\xd3\xf3\x90\xfd\x29\x92\x6d\x37\xdb\x67\x5d\xe4\xe2\x5d\x78\xed\x7e\x15\xb9\x22\x0c\xbc\x0f\x38\x80\xde\x79\xbd\xd5\x22\x56\x56\xfd\x0a\xa5\xef\x92\x7c\xd0\xe8\x5c\x7e\xea\x36\xdf\x9f\xf7\xe1\xbc\xd3\x3a\xf5\x10\xf8\xce\x43\xc9\xea\xef\x5a\x5e\xce\x0a\x95\x6a\xb4\xea\xcd\x8b\x0a\x9c\xd6\x2f\xea\xef\x3d\x83\xd5\x41\x2a\x5d\x8b\xa6\xe5\xd2\xc1\xc7\x73\x8f\x40\xc4\xaf\x8e\x7f\x1a\xfd\x66\xa7\x4d\x6a\x34\x3a\xed\x7e\x17\x87\x15\xd4\xb2\xdb\x9f\xa1\x7e\x6c\xf6\xbc\x0a\xd4\xbb\xcd\x1e\x19\xe4\xac\xdb\xb9\xa8\x58\x64\x4e\xc4\xe8\x18\x22\x88\xd7\xf6\x72\x2a\x64\x6a\x28\xad\x08\x4e\xa1\xf1\x55\xcf\x9b\x11\x84\x53\xaf\xde\x42\x5a\x3d\x42\x26\x15\xa7\x93\xab\x96\xe3\x60\x46\x32\x29\xf0\x26\x0a\x63\x55\x5b\x91\xd8\xf6\x0e\x0e\x0e\xf2\x7c\x66\x6f\x36\x49\x51\x72\xab\xd9\x43\x19\x6b\x67\xc8\x22\x11\xde\x1e\xc2\x8f\xe7\x3c\xbc\xe6\xe8\x89\x0c\xda\x3c\xe3\x3f\x56\x60\x06\x40\x55\x53\x74\x39\x74\x7f\x4c\x6e\x0e\x36\x8d\x62\x78\x04\x03\x79\xe3\x28\xf1\x3b\x3a\xff\x21\x7e\x4e\x31\x41\x3a\x08\x3a\x02\x43\x14\x1f\x60\xa7\xbb\xf7\x3a\x41\x40\x84\x89\x49\xc4\x87\xb0\x7b\x44\xb9\x75\xcc\x59\xf0\x94\xfc\x23\xae\x19\x50\xd3\x5b\xb3\xaf\x05\x9f\x50\x14\xd9\x14\xbd\x1a\x93\x5e\xcd\x9e\x88\x40\x8f\x6b\x01\xbf\xc6\x80\x74\xcc\xe0\xe9\x8c\x05\xee\x54\x5c\x5a\x4c\x87\xff\x96\x89\xeb\x9a\xdd\xc8\x45\x75\xfa\xb7\x09\x5f\x10\x9c\x3a\x0c\x97\x16\xf7\xc8\x54\x02\xc5\x75\xed\xaa\x7f\xe6\xfc\xf4\xc4\xe2\x9b\x0e\xfb\xe9\x96\xfb\xbe\x5e\xe4\xd8\x35\xc2\x9d\x58\xd6\xb1\x4b\x4e\x49\x1f\x06\x32\xb8\x05\x81\x28\x0a\x73\x2e\x4a\x6c\x9b\x81\xbe\xa5\xcf\x45\x44\x29\x7f\x8c\x55\xdd\x44\x94\x47\xd5\xfd\x62\xda\x42\x3f\xaa\x92\xce\x84\x0f\xbe\x0a\x64\x64\x1e\x44\x52\x62\x4d\x21\xa4\xbc\x36\x08\xa6\x78\x30\x9f\x44\xbe\x61\xb0\x1d\x16\x7c\xc9\x94\x3e\xc4\x8a\x13\xf3\x23\x6c\x25\xa8\x32\x21\xc9\xdd\xdd\x7f\x1c\x61\x51\x8e\xb9\x33\x03\x55\xdf\xf2\xe8\x08\x4c\x04\xe4\x13\xe0\x07\x11\x51\xb0\x20\x07\x94\x13\x77\x30\xa3\x54\x66\x71\xe0\xf8\x32\x94\xe9\x21\xbc\x18\xbe\xa5\xdf\x45\xf3\x43\xc2\x82\xc0\x48\x45\xde\x30\x18\x99\x99\x35\xbb\x98\x69\x93\xbd\x35\x1b\x3c\xb6\x7b\x2c\xa8\xb4\xa1\x1e\x2b\x65\x07\x38\xd6\xe9\x13\xe6\x31\x00\x92\xe0\x91\x33\xe9\x35\x6e\x0f\x90\x48\xe8\xa0\x8b\x8d\x50\x12\x2d\x93\xb2\xa1\xae\xcd\x03\xcc\x46\x32\xb1\x4f\x30\xc0\x82\xb9\xa0\x79\x66\xb5\xdf\xee\xee\x3e\x72\xa8\xac\x14\x1a\xbb\x48\xcc\x0a\xc8\x76\x10\x4a\xff\x6b\xc9\xb7\x23\x76\xe3\x14\x4e\x82\xc2\x26\x37\xa5\x87\x7e\xc8\x59\x4a\x0c\xf5\xb8\x04\x5f\x17\x28\x33\xe3\x00\xcb\xb4\xbc\x13\x12\x25\x6b\x19\x43\xa1\xa9\x02\x71\xfd\xd8\x6e\x55\xd6\xf7\xae\x71\xee\x57\x62\x2a\x37\x2d\xb2\x09\xe6\x62\x9d\xc9\x12\x58\x9e\xb0\x1b\x2f\x66\xd7\xec\xdd\x7c\xac\x12\xe6\x4f\xc7\x8f\xaa\x68\xf1\x30\x65\x81\xc8\xd4\x21\xbc\x32\xb0\x15\x09\x60\x38\x2c\x65\xb1\x1c\x0d\x89\xa0\x2b\x28\x19\x8a\x00\x5e\xf0\x03\xfa\x2d\x27\x86\xe1\x70\xc1\x16\xdb\x90\x1d\xe6\x92\x3c\x5e\x96\x78\xbb\x36\xe0\x4a\xd6\x35\x28\x93\xa2\xd4\xbc\xd9\x45\x23\x9b\x12\x55\xcc\xc7\x0d\x9d\xe6\xe9\xaa\xf5\x32\x7f\x77\xcd\xa2\x2c\xaf\x9b\xf7\xf6\xcd\xfe\x7e\x63\x75\x01\xda\x27\xbf\xb6\xa1\x88\xb7\x9c\xc1\xe2\xea\xe5\xb8\xab\x23\x72\xfa\x33\x3f\xa8\x9f\x9d\xd0\x83\x39\x30\xb9\x73\xd6\x9e\xcf\xd9\x81\x3d\x9c\xa0\x66\x07\x1e\xa8\x73\x0a\xf3\xc3\xe4\x35\x87\xf9\x74\xee\x01\xb0\xcc\xb7\x38\x5a\xae\x2d\x1e\x2c\xc3\xb2\x7c\xc5\xd9\x4a\x69\xf5\x67\x49\x78\x36\x4e\x9f\xfd\x74\x93\x6a\x36\xf7\x9e\xbd\xdc\x7b\xee\x73\x8e\xad\x4f\x7e\x6b\xcd\xbe\x5d\x4e\xb0\xed\xae\x80\xc9\x67\x9a\x4c\xee\x73\x87\x42\x0d\xdc\xb9\xa5\x7c\x58\xb3\x37\x39\x66\x7d\x64\x7f\x98\x66\xcd\xb3\xb3\xb3\x22\xfb\x06\xdc\x97\xa9\x39\x94\x9b\xee\x0f\x4a\x3b\x82\x7d\xda\x0f\x94\x12\xf7\x40\x86\xc1\xea\xcc\xed\x67\xa9\x22\xea\x89\x14\x39\x60\xd6\x51\x88\xd8\x10\x2d\x1a\x8b\x3b\x19\xfe\x0d\x09\x66\xe8\x99\x53\x54\xcc\x98\x11\xd2\x64\x89\xd0\x48\xff\x77\xbe\x32\xeb\xbf\x7a\xfd\x13\x0f\xd8\x8a\x82\xbd\x34\xa3\x00\x1b\x2b\x1f\xe6\x95\x7c\x06\x9c\xb5\x6f\x58\x5f\xf2\xe5\x3d\xf9\x20\xf8\x84\x0e\xe0\x1e\x7c\xed\x70\xec\xb2\x95\x3e\x7c\x27\xf1\xae\x4e\xbf\xf9\xcf\xca\x0a\x52\xbc\x0e\xde\x41\x8f\x5b\x51\x15\x9e\x43\xf6\xbf\x13\xb2\x4a\xa7\x32\x1e\x3d\x9d\x69\x7f\x59\x7f\x1f\xe0\x57\xc8\x01\xc7\x6e\x2e\xe4\x5f\xe0\x75\x2b\x1a\x86\xe2\xc9\xf4\xa5\x77\x49\x92\x67\x3f\xfc\xdb\xf8\x61\xde\x9b\xce\x5c\xed\x78\xf0\x74\xcb\x4c\x07\x89\xab\x6c\xf4\xc0\x6d\x8f\xf5\x57\x32\x9e\x58\x99\xf5\x71\x57\x68\x55\xaa\x05\xf3\x5b\x27\x79\x25\x78\x72\xcf\x58\x90\x68\x5b\xdc\xe3\x41\x8b\x3e\x78\x85\xe7\x7f\xd4\x59\x16\x3b\xcc\xbb\x77\x8a\x9e\xa8\xa1\x9c\xb6\x5b\x4b\x3d\x25\x76\x6d\x3c\xa5\xee\xaf\xec\x4e\xf9\xad\x28\x6a\xa2\xb6\x2f\xc7\x7c\x5f\x35\xdd\xb0\xbd\xeb\x72\xec\x41\xaf\x79\xb0\xa6\xc1\x7b\xee\x0a\xb7\xa8\x1a\x6f\x9d\x67\xa2\x4c\xe3\x2d\x94\x69\xeb\xec\xf4\x2d\x11\x7c\x5f\x47\xfc\x1c\x58\xff\x9f\x6d\xee\xe2\x76\x6b\x9a\x90\x17\x36\x5c\x53\xd0\x13\x6c\xb9\x66\xd2\x3c\x7b\xe3\xdf\xc6\x1b\x9f\x37\x5d\xcf\x9b\xae\xe7\x4d\xd7\xb6\x3b\xcb\xf3\xa6\x6b\x6b\x5a\xb6\x75\x0b\x85\xb3\xe9\x7d\xdc\xc9\x37\xbc\x0a\x9d\xa1\xcc\x21\x8f\x7e\x15\xa3\x74\x37\x69\xe1\xaa\xc9\x7c\xa1\x0f\x0e\x0e\xee\x7b\xc3\x5d\x7e\xb3\xbb\xfc\x4a\x72\x3b\x9a\x86\x6d\x6a\x5f\x1e\xb3\x75\xd9\x5f\xdb\xba\xac\x7c\x89\xf6\xd0\x92\x2f\xf4\x36\x77\x2e\x36\x94\xaf\x61\x2d\xa6\xab\xf2\xb7\x1e\x1f\xcf\x21\xf6\x17\xb3\x95\xd1\x68\xe3\x54\x85\x3a\xc1\xe0\x76\xb3\xf7\x70\xcb\xb9\x63\xe9\xbe\xc3\xdd\xcc\x70\xec\x62\x98\x9f\xe4\xff\x5a\xe5\x34\xb1\x6d\x6d\xed\x9a\xfb\x75\xb9\x8a\xf3\xfc\x75\xec\xd2\x35\x56\x82\xd0\x7d\xe0\x13\x6b\xfe\xe5\x42\xeb\x3f\x01\x00\x00\xff\xff\xe7\x17\xfa\x7f\x4c\x3b\x00\x00")

func templateDefaultTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDefaultTmpl,
		"template/default.tmpl",
	)
}

func templateDefaultTmpl() (*asset, error) {
	bytes, err := templateDefaultTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/default.tmpl", size: 15180, mode: os.FileMode(420), modTime: time.Unix(1455637291, 0)}
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
	"template/default.tmpl": templateDefaultTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"default.tmpl": &bintree{templateDefaultTmpl, map[string]*bintree{}},
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
