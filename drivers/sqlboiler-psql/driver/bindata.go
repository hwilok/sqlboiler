// Code generated by go-bindata.
// sources:
// override/templates/17_upsert.tpl
// override/templates_test/singleton/psql_main_test.tpl
// override/templates_test/singleton/psql_suites_test.tpl
// override/templates_test/upsert.tpl
// DO NOT EDIT!

package driver

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

var _templates17_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x58\xdf\x6f\xdb\x38\x12\x7e\x96\xfe\x8a\xa9\x71\x40\xa4\x83\xa3\xdc\x73\x0f\x7e\xb8\xa4\x3f\xae\xd8\x6d\xea\x6d\xda\x2d\xb0\x45\x11\xd0\xd2\xc8\x26\x42\x91\x2a\x45\x25\xeb\x55\xf5\xbf\x2f\x66\x28\x59\x52\xec\x6d\xdc\xcd\x16\x7d\x0a\xc4\x21\x67\x3e\x7e\xf3\x71\x66\xe2\xa6\x39\x85\x7f\x39\xb1\x52\x78\x29\x0a\xbc\x92\x7a\x5d\x2b\x61\xe1\xe9\x02\x92\x77\xb4\x9a\xd0\x32\x7c\x81\xaa\xb7\x7c\x01\x27\x9d\xc2\x0b\x51\x21\x9c\xb6\x6d\xc8\x0e\x6e\x85\x3d\xfa\x78\x2a\x0a\x54\xd3\xe3\x55\xba\xc1\x42\xf0\x81\xfd\xa3\xc9\xd5\x60\x6d\xdb\xf0\xec\x0c\xde\x97\x15\x5a\xf7\x12\x84\x73\x58\x94\xae\x02\xa1\x41\x6a\x5a\x9b\x83\xd0\x19\x64\x06\x79\xad\x2e\x33\xe1\x10\x8c\x05\xb9\xd6\xc6\x22\x18\x0d\xa9\xd1\xb9\x92\xa9\x4b\xc2\xbc\xd6\x29\x44\x06\xfe\xdd\x34\xfb\x04\xb4\x6d\xdc\x87\x89\xbc\x9b\x37\xfa\xa2\x3b\x0a\x2b\x63\xd4\x7c\xe7\xe9\xc2\xa8\xba\xd0\x15\x7c\xfc\x54\x39\x2b\xf5\x7a\xde\xc5\xdd\x5f\xbf\xdb\x48\x87\x4a\x56\x0e\x92\x24\xf1\x8b\x31\xa0\xb5\xc6\x42\x13\x06\x16\x5d\x6d\x35\x98\xc4\xc7\x8d\x56\x46\xaa\xe4\x25\xba\x67\xe7\x51\xdc\xfb\x1c\x40\xec\xc5\xbf\x17\x76\x14\x2d\x49\x92\x38\x6c\xc3\x11\x73\xcb\xc7\x51\x07\x4b\xa1\x65\x5a\xd1\x1a\xa3\x3f\x9e\xcb\xe5\x77\x26\xb3\x09\x03\x99\x13\x28\x92\xd1\xf7\x61\xf2\xbf\xec\xfe\xc9\x02\xb4\x54\x14\x2f\x28\x89\x0c\x1f\xe3\x83\x15\xe5\x73\x6b\x23\xb4\x36\x8e\xc3\xa0\x9d\xb0\x7e\x88\x74\xa8\xe9\x61\xd0\x37\xfe\x8e\x69\xed\x8c\xfd\x16\x05\x8f\x5c\x97\x7f\x33\x23\xcb\x88\x02\x03\xa3\x7f\xbe\x83\xf0\x23\x92\x44\x38\x7e\x5c\x72\xfe\xf1\xdc\x1c\xc9\xff\x0f\xa0\x7f\x57\x70\x64\x0e\x06\x16\x03\x57\x5d\x01\x62\x7b\x95\x5c\xe2\x5d\x34\x6b\x9a\x64\x79\xb3\x26\xe8\x6d\xfb\x14\xb4\x81\xa6\x19\x15\xe7\xb6\x85\xd2\x9a\x5b\x99\x61\x06\xb9\xb1\x50\xf3\x95\x66\x4c\x6e\x18\x50\x65\x27\x4e\x15\xb1\x34\x73\xb2\xc0\xca\x89\xa2\xbc\xf6\xbb\xae\x37\xa8\x4a\xb4\x33\x48\xa0\xf5\xbb\x65\x0e\xda\x38\x48\x2e\xcd\xff\x8d\xb9\xa9\xb8\x3b\x4c\x84\x92\x99\x73\xcc\x8d\x45\xcf\x1c\x6f\x62\xfa\xf6\xb3\x3e\xdc\x84\xa0\x30\x12\xd4\x19\xc7\xd1\x7f\x3c\xc3\x5c\xd4\xca\x55\xe4\xf4\x73\x8d\x56\x62\x95\x5c\x1a\xfd\x1b\x5a\xd3\x99\xae\xd0\x45\x4d\x73\xbf\xa9\xb5\x6d\x47\xf0\x07\xe9\x36\xdd\xce\x39\x98\x38\x0c\x83\xb3\x33\x38\xaf\xa5\xca\x20\x15\xe9\x06\xe1\x06\xb7\x20\xf5\xa9\x92\x1a\xa1\x5e\x2b\xa9\xb6\x70\x0a\xc5\xb6\xfa\xac\xe0\xb6\x82\x92\xfe\x96\xd6\xac\x14\x16\x55\x18\xac\xea\x9c\x90\x54\xce\x16\x42\xaf\x15\x52\xa1\x3a\xaf\xf3\x1c\x6d\x14\xf3\xfd\xf7\xd4\x40\x37\x5c\xd5\x79\xf2\xc1\x4a\x87\xe7\x5b\x87\xd1\x89\x3b\x21\xd2\x01\x55\x85\x87\xcc\x39\x9b\xc3\xfb\xcb\x09\x2d\x53\xe2\xae\xe7\x90\x12\x08\x2b\xf4\x1a\xf7\x74\x36\x71\x78\xc5\x2a\x8a\xd2\x6f\x71\x38\x95\xe7\xa3\xdd\x0d\xaa\x7e\xb4\xab\x91\x1a\xbe\xe2\x8b\xf2\xf9\x74\x01\x64\xed\x0c\x71\x18\x0c\x09\x5b\xd6\x7d\xc2\x56\x75\x1e\xb3\x96\xf7\xb5\xe3\x45\x7b\x41\xfa\x78\x5d\xbb\xe4\xed\xcf\x26\xbd\x21\x37\xac\x98\xb9\x17\x4e\x46\x51\x1e\x38\xfc\xf1\x06\xb7\x9f\x8e\x0b\xf1\x5e\x2b\x1f\x24\x0c\x6e\x85\xe5\x37\xc2\x6f\x3b\x64\x59\x3d\xe9\x42\xd2\xbd\xfb\x09\xc0\xa2\x9b\x6a\xf1\x15\x1b\x7c\xe6\xe8\x55\x84\x41\x70\x30\x74\x5f\x8d\x1f\xb0\x8f\x5f\xce\x11\x5b\x4d\xed\xc6\xbb\x87\x64\xf1\xe7\x4e\x06\xf4\x15\x87\x41\xd0\x15\xe4\xc9\x05\xde\x8f\xa4\xf7\x98\x0b\x2c\xad\x2c\x84\xdd\xfe\x84\xdb\xf1\xce\x69\x2b\x62\x18\xc4\x66\x0e\x0a\x75\x37\xe9\xc4\x54\x61\xff\xc3\x2c\x3f\x5c\x60\x6b\xcd\x23\xb0\x33\x5d\x29\xbd\x5f\x6e\xa9\x07\xd4\x2a\xe3\x3a\xb9\xe2\x82\xd3\xdd\x39\x65\x08\x40\x6c\x50\xf9\xe5\xfa\x1b\xf4\xef\x98\x18\xb9\xf7\xa6\x07\x94\xbd\x61\x8c\x73\x77\x70\x01\x85\xb8\xc1\x68\x68\x28\x74\xe2\x28\x7a\xa8\xc3\x92\xa3\x72\xbb\x8b\x30\x3f\x28\xed\xfd\x93\x0c\x3f\xf0\x0f\x23\xa1\x02\xbd\x85\xa1\x50\x73\x99\xf5\x42\xff\x85\x4c\x4b\x53\xb9\xb5\xc5\x2a\xca\xa4\x50\x48\x41\x66\x4d\x33\xfe\x7f\xa2\x6d\x67\x87\xa6\x0a\x8b\xae\x5f\x1e\x1a\xeb\xbc\xeb\xfc\x9c\x45\x1f\xff\x56\xa8\x1a\x5f\x8b\xb2\xe4\xeb\xd3\x13\x1a\x61\x91\x3a\xeb\x4c\x07\x49\x79\xb7\x2d\xf1\xf0\xa5\x77\x0e\xfb\x78\x41\xdf\xe8\x46\x4d\x6c\xd2\xc5\x98\x92\x2e\x65\x16\x5d\x4c\x1b\xfb\x6c\x31\x50\x8b\xee\xfb\xc1\xa4\x88\x14\xea\x00\xc8\x29\x4a\x86\xd9\xfa\xfe\xcf\xd4\x71\xa9\xc5\x9c\x52\x93\xbc\xd2\x99\xb4\x98\xba\xa8\x5f\xf8\x95\x76\xbc\xc9\x23\x43\x6a\xb9\x15\x6a\xd2\x92\xd9\x58\xbd\xb0\xa6\xe8\xc1\xb3\xc3\xae\x54\x4e\x12\x13\xfb\x02\xe7\x91\xd0\xfc\x23\xb5\x43\x9b\x8b\x14\x1b\x3f\x42\xb0\xd0\xef\xd1\x34\xa2\xb0\x3f\x38\x04\x5f\x3a\xfb\xd7\xa1\x47\x3e\xfc\x4d\x65\xee\x27\xb8\x67\xb8\xaa\xd7\xaf\x4d\xe6\x7b\x70\x5e\xb8\xe4\x45\x69\xa5\x76\x4a\x47\x83\x9d\xfb\x8b\xed\x7d\xb1\xbe\xe3\x87\x77\x13\x3b\x43\xb4\x07\xee\xe3\xf3\x4f\xa3\x51\xc2\x8f\xe4\xad\xb9\x8b\x46\xe1\xbc\x37\x1a\x97\x93\xab\x54\xb0\x9e\xe8\xfa\xfc\x6f\xe2\x2e\xc7\x8b\x05\x54\x9f\x55\xf2\xdc\xda\x4b\xf3\xd6\xdc\xf9\xfe\xd8\x79\xa6\xe4\x9f\x9d\x41\xff\xf6\x78\x22\xd6\x27\xae\x4b\x00\x08\xbd\x75\x1b\x1a\x9d\xef\x36\xa8\xc1\x6d\xd0\xe2\x49\x45\xc3\xa3\x7f\x6f\x9d\x42\x86\x61\xe5\xba\x57\x2c\x23\xa6\x31\xf8\x30\x5a\xee\xc5\xfb\x12\x9c\x56\x55\x9a\xf2\x69\xc4\xe7\x42\xf0\x2d\xb5\x75\x36\xd0\x3b\x6e\x8d\x47\x34\xda\xbe\x95\x3f\xb4\x97\x5b\x37\x2c\x7c\xe6\x8f\x73\xbd\x6b\xe1\xc1\x57\x66\xe4\xdd\x0f\x07\x99\xf9\x5f\xee\xd0\xee\xcd\xc7\xdd\x04\xac\xba\x9f\x5c\xfa\x03\x5a\xaa\xf1\x6c\xdc\x86\x7f\x06\x00\x00\xff\xff\x2a\x72\xec\xd3\x0c\x12\x00\x00")

func templates17_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertTpl,
		"templates/17_upsert.tpl",
	)
}

func templates17_upsertTpl() (*asset, error) {
	bytes, err := templates17_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.tpl", size: 4620, mode: os.FileMode(420), modTime: time.Unix(1527186642, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_main_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x6d\x6f\xe3\xb8\x11\xfe\x2c\xfd\x8a\xa9\x81\x1c\xa4\xad\xc2\x1c\xfa\xf2\x25\x07\x63\x91\x38\x4e\xba\xb8\xac\xe3\xb3\xdd\x1e\x8a\x6e\xdb\xa3\xad\x91\x42\x44\x22\x19\x92\x8a\xd7\x3d\xe4\xbf\x17\x43\x4a\xb6\x9c\xb5\x93\x6d\x8b\x05\xfa\xcd\x26\x9f\x19\xce\x33\xef\x7a\xe2\x06\x4c\xf9\x79\x7a\x73\xfd\x80\x1b\x18\x82\xc1\x12\x3f\x6b\xf6\xb1\xb1\x6e\xa4\x6a\x2d\x2a\x4c\x7e\x49\xde\xd7\xe9\x3f\x2e\x6e\x17\xe3\x19\x2c\x2e\x2e\x6f\xc7\x70\x37\xb9\xfd\x2b\xb0\x77\x9f\xe4\x27\xfb\xdb\x8b\xab\x2b\x18\xdd\x4d\xe6\x8b\xd9\xc5\x87\xc9\x02\xd8\xbb\xf7\x70\x7d\x37\x1b\x7f\xb8\x99\xc0\x8f\x63\x42\xbd\xff\xe1\x93\xfc\x25\x8d\x63\xb7\xd1\x08\xba\x5c\xa0\x75\x68\xc0\x3a\xd3\xac\x1c\xfc\x1a\x47\xf9\x72\xa4\xa4\x84\x77\xf6\xb1\x62\x57\x97\x31\x1d\x4c\x78\x8d\x40\x10\x21\xcb\x38\xba\x57\xd6\x01\xec\xfe\x37\x16\x4d\xff\xbf\xe6\xd6\xf6\xff\x5b\x5b\xd5\x2a\xc7\xdd\xbd\x32\x5e\x5e\x48\x17\xc7\x91\x2e\xa7\xdc\xda\x6b\x51\x6d\x01\x71\xe4\xd0\xba\xab\x4b\xff\x6a\x7b\xf6\x1c\xc7\x45\x23\x57\x20\xa4\x70\x49\x1a\xcc\xfc\xc8\x85\x84\x21\x7c\xd7\x71\xf8\xf5\x99\x60\x67\x67\x60\xd1\x35\x1a\xf2\xa6\xd6\x16\xdc\x3d\x42\xce\x1d\x5f\x72\x8b\x60\x57\xf7\x58\x73\xe0\x32\x07\x51\x93\x19\x16\x84\x23\x3b\x14\x70\x70\x48\x47\xdc\x6c\xc0\x70\x99\xab\xba\xda\x90\xae\x12\x25\x1a\xee\x30\x07\x32\xaa\xa7\x4a\x81\xbb\xe7\xce\x9f\x5a\x58\x71\x09\x4b\x04\xd3\x48\xe0\x25\x17\xd2\x3a\x52\xdc\x58\x21\x4b\xb2\x60\x5f\x91\x7d\xac\x96\x4a\x54\x68\xe0\x6e\xf6\x11\x34\x5f\x3d\xf0\x12\x59\xe0\x97\x68\x78\xd7\xf1\x49\x03\x91\x24\x05\x34\x46\x19\x22\x4d\xd9\x81\xc6\x84\x83\x38\x8e\x9e\x84\x46\xc3\xe6\xe8\xae\xb0\xe0\x4d\xe5\x92\x81\xa6\xb0\x05\x9e\x83\x0c\x06\xba\x59\x56\x62\x35\x48\x8f\x42\xc9\x0b\x83\x0c\xfe\xf8\x87\xdf\xff\xee\x38\xa8\x8d\x20\x29\x34\xf8\xd8\x08\x83\x83\x94\x42\xc7\xda\xd4\x18\x42\x10\xbc\x41\x37\xf7\xf1\x6a\xe5\xf2\xa5\xe4\x35\x61\x23\xcd\x7c\xd6\x1c\x03\xd2\x65\x80\xf9\x64\x3a\x06\xa3\xcb\x00\xf3\x39\x76\x0c\x46\x97\x2d\x8c\x52\xad\x07\xfb\x20\xf7\x78\x7b\x4c\x97\x9e\xc7\xb4\x75\xe4\x89\x31\xf9\x7e\x08\x4f\xbc\xe2\xec\x12\x4b\x21\xff\xc2\x2b\x91\x73\x27\x94\x4c\x52\xd6\xfe\xc1\x24\x8e\x22\x0f\x09\x6a\x26\xca\x8d\x6b\xed\x36\x49\x20\x47\x41\xd9\x71\xc9\x8e\x62\xc9\x25\x1d\x36\xb8\x67\x8b\x9d\x28\x97\xf8\x1f\xe3\xc7\x86\x57\x36\x09\x3c\x33\xf8\xbe\xc3\x07\x72\xaf\x28\x0f\x71\xeb\xe0\x5d\x98\x8e\xe3\x5b\x1f\x74\x02\x5b\x97\x64\x71\x94\xb2\xd1\x3d\xae\x1e\x12\x72\x8f\x28\x7c\x76\xfe\x66\x08\x52\x54\x94\xaf\x91\x41\xd7\x18\x49\xa7\x71\xf4\x1c\xc7\xd1\xd9\x19\x8c\x0c\x72\x87\xc0\xdb\x32\x13\xff\xc2\x1c\xf2\x25\x90\x09\x8c\xe2\xd1\x2b\xfe\xe1\x0e\xc3\xe6\x8e\x2f\x2b\x0c\x17\x5b\x06\xbd\x47\x87\xa0\x59\xcd\x1f\x70\x7a\xd3\xf5\x93\x24\xfd\xe1\x2d\x73\x7a\xb2\xb9\x51\x7a\xe1\x9f\x7e\x53\xae\x2f\xb6\xf2\x6c\xbe\x52\x30\x8e\xa8\x29\x8d\xea\x1c\xce\x87\x80\x9f\x71\xc5\x46\xaa\xae\xb9\xcc\x93\x81\x2e\xff\x49\x77\x54\x62\xa7\xa7\xa1\x7e\x4f\x95\xac\x36\x83\x0c\x76\x64\x3b\x71\x36\x96\x4f\x30\x04\xae\x35\xca\x3c\x51\x96\xfe\x0b\x43\x49\x48\x68\x5d\x8e\xe5\x53\x92\x32\xc6\xd2\x38\x0a\xf6\x1d\x7e\xd2\x3e\x56\x5e\xfd\xce\xe3\x7d\x81\xaf\x7f\x24\x8e\x4c\x06\x6b\x7a\x40\x28\x36\x15\x1a\x93\x9e\xa9\x73\x97\xab\x86\x8a\x70\xdd\xd7\x3d\x77\xb9\x6f\xde\x12\xd7\xd7\x3f\xe2\xe6\x0a\xad\x33\x6a\x83\x26\xd9\xce\xbe\x0c\xcc\x5e\x74\x77\xfa\xb8\x71\xaf\x7a\x5a\x19\xcb\x7e\x36\x5c\x27\x68\xa8\xda\x0a\x2e\x2a\x6a\xdf\x0a\x2c\x89\x42\xeb\x69\x58\x05\x3f\x50\x13\xe8\x87\xb4\x6f\xe3\xff\xfa\x92\x7d\xac\xf6\x9f\x39\xc0\xe7\x67\x2e\x0e\x3d\x52\xd4\x8e\x4d\x8d\x90\xae\x92\xa4\x3d\xfd\xba\x77\xd7\x5c\x38\x28\x94\x39\x4c\x32\x8e\xd6\x6c\x54\x29\x8b\x49\x0a\x67\x67\x70\x51\xd0\xe0\xef\x32\x52\x58\xc8\x95\xc4\x0c\x56\x84\xf0\x73\x73\x6d\x84\x43\x40\x99\x83\x2a\xfc\x81\x16\x1a\xe3\x83\xbe\xfa\x46\x2c\x0e\x38\xb0\x95\x97\xa2\xda\x2e\x05\xfb\x43\xd3\x34\x72\x54\xe7\x89\xa5\x0c\xcb\x3a\xe9\x76\x8f\xc8\x80\x9b\xd2\x02\x63\x2c\xfc\xef\x8d\xd6\xd5\x81\x12\x69\x85\x83\x54\x5b\x4f\xff\x59\x61\x88\x02\x2a\x94\xc1\x98\x94\x3c\xf3\xbd\xf7\xcb\xaa\x57\x02\xc1\x12\xcb\x26\xb8\x9e\x21\xcf\xd1\xb4\xe8\x40\xd7\x86\xf2\x39\x1f\xc2\x77\xcb\x8d\x43\xcb\x2e\x9b\xa2\xf0\xbb\x0e\x5d\x91\xbb\x0f\x5d\xad\xfa\x85\x17\x54\x6c\x0f\x43\xe8\x82\xf0\x36\x96\xe7\x43\xa0\xeb\x59\x23\xdf\x88\x62\x17\x26\xd3\x48\x29\x64\x79\x3e\xd8\xba\x38\x78\x29\x7d\x81\x0f\x8f\xb7\x13\x25\x49\x0f\x5c\xa3\x31\x7b\xd7\x2f\x5b\xe6\x9b\x01\x6f\x3d\x0e\x7f\xfb\x7b\x70\x25\xd9\xdc\x0a\x75\x47\x1d\x8b\xb9\xa6\x77\x8b\x64\x30\xbd\xf9\xd3\xdd\x7c\x31\x3c\xb1\xbe\x01\xd2\x7c\xf5\xd3\xef\x05\x66\x7a\x37\x5b\x0c\x4f\x72\x8f\xa1\x99\x7a\x08\xf3\xe7\xf9\x78\xd6\xe9\xa1\x99\x7e\x50\xcf\xc5\x7c\x7e\xfd\xe1\x76\xdc\xe1\x76\x3b\x2f\xa1\x9f\x8f\xf0\x7a\x39\xcd\x76\xb9\xea\x6a\x9d\x75\x61\x13\xaa\x71\xa2\x62\x0b\xac\xb5\x87\x0d\xfc\xda\x57\x76\x3b\xd0\x6b\x23\xf9\x68\x01\x86\xba\x06\xa5\x69\xb3\x81\x42\x54\xd8\x55\x1f\x11\xbb\x6e\x89\x79\x2b\x06\x27\xf6\xfc\x24\x3f\xd7\xca\xba\xd2\xa0\x3d\xef\x79\xb4\xf3\xda\xd6\x33\xdb\x72\x08\xfb\x5b\xaf\x1e\xbe\x54\xdb\x29\xf2\x40\xdf\xa1\x77\x98\x4a\x12\x28\x7d\xc5\x9c\x93\xa3\x86\x74\x9b\xcf\xff\x91\x49\xbb\xf1\xfb\x0d\xcd\xea\x27\x1d\x0c\xc1\xd5\x9a\xf9\x4d\x2a\xdd\xd6\x0a\x1d\xb5\xd3\xe1\x48\x42\xee\xef\x3a\xbb\x74\x6c\x15\x68\xd6\xb6\x5e\x9f\x82\x01\x9c\x2f\xbf\xd8\x30\x0e\xeb\xee\xaf\x5f\x6f\x68\x26\xa8\xd7\x3b\x38\x3d\x15\xc5\x29\x7e\x16\xd6\xd9\x43\xcf\x9c\x9d\x81\x43\x6e\x72\xb5\x96\xbe\xaf\x37\x0e\x2d\xac\x2a\xe4\xb2\xd1\xe0\xb8\x7d\xb0\xb0\xbe\x47\xe9\x47\x5b\xf8\x8e\x2b\x84\x14\xf6\xbe\x6b\x6e\x87\xec\xec\x14\x1e\xff\x2a\xdb\x5b\x2a\xfd\xb7\x74\xe7\xd6\xb7\xd6\xca\x0e\x0f\x1e\xf1\x5f\xaf\xa7\x5b\xb7\x29\xcb\x66\x58\xab\x27\xda\x97\x7b\x2d\xe7\x58\x74\x95\x24\x56\x49\xfb\xe1\x9f\x05\x3a\xfe\x5b\x5b\x14\x5b\x2e\x07\x9e\xed\xae\x32\x6f\xb5\x37\xe0\x85\x47\x76\x88\x76\xf8\x3c\x56\xec\x4e\xa3\x4c\x06\x5d\xdf\x18\x64\x90\x1b\xf1\x84\x86\x4d\xe7\x3f\xdd\x5e\x36\xa2\xca\x7f\x6a\xd0\x6c\xda\xc1\xd0\x7d\x3a\x85\x2c\xff\xb2\x68\x5e\x96\x54\xfb\x81\x92\xbe\xd6\x00\xa5\xa8\xb2\x2f\x5c\xb6\xcf\xe5\x39\xfe\x77\x00\x00\x00\xff\xff\x00\xef\xd0\xbf\x8f\x11\x00\x00")

func templates_testSingletonPsql_main_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_main_testTpl,
		"templates_test/singleton/psql_main_test.tpl",
	)
}

func templates_testSingletonPsql_main_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_main_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_main_test.tpl", size: 4495, mode: os.FileMode(420), modTime: time.Unix(1527187113, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonPsql_suites_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\xb1\x8e\x83\x30\x10\x44\x7b\xbe\x62\x84\x28\xe0\x04\xfe\x80\x93\xae\xba\xea\xae\x48\x11\x91\x0f\x70\xc2\x82\x2c\x39\x1b\x84\x17\x29\x92\xf1\xbf\x47\x18\x8b\x90\xce\xe3\x99\xb7\x3b\xdb\xcf\x7c\x43\x4b\x4e\x2e\xa3\xa3\x49\x4a\xc1\x97\x90\x13\xc3\x83\x6a\x2b\xf8\x0c\xf0\xbe\xc1\xa4\x79\x20\x14\x86\x3b\x7a\xd6\x28\x44\x5f\x2d\xe1\xfb\x07\xaa\x5d\x5f\x2e\x84\x94\x33\x7d\x32\xd5\x9f\xfb\x7f\x18\x8e\x36\x9a\xdd\x27\xeb\x8e\x72\xcb\x9e\xf4\x3d\x0e\x4b\x64\x94\x0b\x46\x3b\x4f\xda\x62\x81\x18\xb1\xf4\xab\x77\x50\xd4\x79\xe6\x32\xf7\xfe\x4d\x87\x90\xd7\x58\x6b\x7f\x7e\x6e\x27\x55\x71\x19\x71\x77\xec\x91\x54\xc8\x5e\x01\x00\x00\xff\xff\x2f\xea\xf2\xb5\x00\x01\x00\x00")

func templates_testSingletonPsql_suites_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_suites_testTpl,
		"templates_test/singleton/psql_suites_test.tpl",
	)
}

func templates_testSingletonPsql_suites_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_suites_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_suites_test.tpl", size: 256, mode: os.FileMode(420), modTime: time.Unix(1527187008, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x93\xcf\x6e\xda\x40\x10\xc6\xcf\xde\xa7\x98\x5a\x6d\x65\x57\xce\x46\xbd\xa6\xe2\x90\x10\x0e\x55\x55\x84\x82\x79\x80\xc5\x1e\x93\x15\xeb\x5d\x6b\x77\x1c\xa0\x8e\xdf\xbd\x5a\x03\x09\x7f\x8c\xc4\x25\x55\x0f\xec\x61\x76\xbe\xd9\xdf\x7c\x7c\x6e\x9a\x1b\xf8\x4c\x62\xae\x70\x2c\x4a\x9c\x4a\xbd\xa8\x95\xb0\x70\x37\x00\x9e\xfa\x2a\xf7\x65\x78\x05\xb7\xbf\x79\x05\x92\xa4\x70\x28\x1c\xc2\x4d\xdb\xb2\xe3\x01\x13\x55\x5b\xa1\xce\xe5\xd5\xb6\xde\x2b\x7e\x11\xf6\x2a\x69\x26\x4a\x54\xbd\xd2\xab\xb0\x8f\xe5\x45\xad\x33\x20\x74\xd4\x34\xa7\xf4\x6d\x3b\xab\x1c\x5a\x8a\x08\xbe\xf9\x0e\xa9\x17\x3c\x8d\xa1\x61\x01\xf1\x89\xb0\x42\x29\x54\x51\xcc\x58\x20\x0b\x50\xa8\xa3\xa6\x39\xe5\x68\xdb\xa1\x51\x75\xa9\x5d\x0c\x83\xc1\xc5\x9e\x89\x95\xa5\xb0\x9b\x5f\xb8\x79\xeb\x6e\x58\x10\x10\x9f\x2e\x65\x15\x85\xfe\xac\xa4\x5e\x40\x87\x07\x2b\x49\xcf\x60\xb4\xda\x40\xb5\xd5\xc1\x12\x37\x90\x6d\x95\x61\xcc\x82\x96\xb1\xc0\x21\xe6\xde\x04\x2b\x74\x6e\x4a\xf9\x07\xf9\x18\x57\x53\xc4\x3c\x8a\x59\xf0\x22\x2c\xa0\xed\x7e\xc6\xb2\xe0\xf6\x16\xee\x89\xb0\xac\x08\xe8\x19\xe1\xe7\x78\x3a\x7a\x4a\xc1\xc9\x1c\xc1\x14\x20\x34\xcc\x26\xbe\xc2\x82\x3e\x76\xff\xc8\xa1\x73\xef\x37\x4d\xdb\x19\xe3\x1f\x3a\xe4\x98\x92\xad\x33\x8a\x3c\x60\x02\x5f\xfb\x46\x26\xd0\x57\x7d\x7c\x48\x37\x15\xba\x04\xc8\xd6\x18\xff\xe8\xe6\x7e\x1a\x80\x96\x6a\x67\xd6\xc8\x6f\x53\x44\xe1\x4c\x77\x36\x91\x79\x7f\xf4\x02\x21\xb8\x8e\xe5\x0e\xbe\xb8\x30\xf1\x03\x77\xe6\xd1\xda\x6f\xf5\xbb\x76\x94\xae\xa3\xb9\x91\x8a\x3f\xe0\x42\xea\x28\x8e\x59\x90\x63\x81\x16\x68\xcd\x9f\x8c\x52\x73\x91\x2d\xbd\xa1\x6f\x6b\xf6\x81\xf3\x7d\x88\xd6\x09\x14\x42\x39\x4c\x3c\x74\x77\x5c\xb7\x46\xdd\xe9\x2f\xec\x70\x06\x9f\x99\x5a\x53\x57\x38\xfd\x6b\xf6\xa1\x8e\x68\x1d\xf3\xa1\x6f\x3b\x60\x3f\x67\x88\xf6\x33\x7d\x4b\x37\xd5\x37\x7d\x3f\x6a\x09\x57\x42\x13\x18\x8d\x60\x31\x33\x36\x4f\x60\x61\xe8\x2e\x4c\xb6\xfd\x3b\xa2\x93\x7c\xcd\x26\x8f\xf7\xe9\xa8\x2f\x5f\x1f\x91\x96\x9d\xe3\x57\x7d\x76\x9c\xf3\x8f\x0d\xd6\xd5\x39\xf1\x19\xff\x67\x31\xf9\x1f\x52\xd2\xb2\xbf\x01\x00\x00\xff\xff\x06\x55\xfd\x54\x80\x06\x00\x00")

func templates_testUpsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertTpl,
		"templates_test/upsert.tpl",
	)
}

func templates_testUpsertTpl() (*asset, error) {
	bytes, err := templates_testUpsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.tpl", size: 1664, mode: os.FileMode(420), modTime: time.Unix(1527186259, 0)}
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
	"templates/17_upsert.tpl": templates17_upsertTpl,
	"templates_test/singleton/psql_main_test.tpl": templates_testSingletonPsql_main_testTpl,
	"templates_test/singleton/psql_suites_test.tpl": templates_testSingletonPsql_suites_testTpl,
	"templates_test/upsert.tpl": templates_testUpsertTpl,
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
		"17_upsert.tpl": &bintree{templates17_upsertTpl, map[string]*bintree{}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_main_test.tpl": &bintree{templates_testSingletonPsql_main_testTpl, map[string]*bintree{}},
			"psql_suites_test.tpl": &bintree{templates_testSingletonPsql_suites_testTpl, map[string]*bintree{}},
		}},
		"upsert.tpl": &bintree{templates_testUpsertTpl, map[string]*bintree{}},
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
