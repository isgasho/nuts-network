// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1_create_table_document.down.sql
// 1_create_table_document.up.sql
// bindata.go
package migrations

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

var __1_create_table_documentDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x09\xf2\x74\x77\x77\x0d\x52\x28\x2d\x48\x49\x2c\x49\x8d\x2f\x2e\x49\x2c\x29\x8e\x4f\xce\xcf\x2b\x49\xcd\x2b\x29\x8e\x2f\xce\xac\x4a\xb5\xe6\x82\x28\x74\x74\xf2\x71\x55\x00\xcb\x43\x45\x3c\xfd\x5c\x5c\x23\x14\x3c\xdd\x14\x5c\x23\x3c\x83\x43\x82\x15\x32\x53\x2a\x40\x3a\x8b\x33\x8b\x4b\x52\xf3\x92\x2b\xe3\x33\x12\x8b\x33\xf0\x28\x2d\xc9\xcc\x4d\x2d\x2e\x49\xcc\x2d\x40\x56\x08\xb1\x25\x25\x3f\xb9\x34\x37\x35\xaf\xc4\x1a\x10\x00\x00\xff\xff\xf3\xed\x99\xe6\xa3\x00\x00\x00")

func _1_create_table_documentDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_create_table_documentDownSql,
		"1_create_table_document.down.sql",
	)
}

func _1_create_table_documentDownSql() (*asset, error) {
	bytes, err := _1_create_table_documentDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_create_table_document.down.sql", size: 163, mode: os.FileMode(436), modTime: time.Unix(1593158817, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_create_table_documentUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x41\x6f\x9b\x30\x1c\xc5\xef\xfe\x14\x6f\xa7\x82\xd6\x4c\x99\xb4\xd3\x50\x0f\x90\xb8\xc4\x1a\x33\x95\x71\xda\xf5\x84\x50\xb0\x06\x5a\x62\xa3\xda\x6c\xcd\x3e\xfd\x64\x02\x49\x49\x97\x43\x24\xec\x9f\xdf\xff\xff\xde\x5b\x09\x1a\x4b\x0a\x19\x27\x19\x45\x6d\x76\xfd\x41\x69\x47\x02\x02\x00\x4d\x65\x1b\xbc\xfd\xad\x36\xb1\x08\xbe\x2c\x43\x3c\x08\xf6\x3d\x16\xcf\xf8\x46\x9f\x6f\x07\x74\x67\xb4\x6d\xad\x53\x7a\x77\x2c\x87\x67\x13\x7a\xba\x76\xc7\x4e\xcd\x94\x1e\x63\x31\x10\x9f\x97\xcb\x10\x3c\x97\xe0\xdb\x2c\x1b\xd9\xf6\xa0\xac\xab\x0e\xdd\xc4\x26\x2c\x65\x5c\x8e\x1f\x73\x76\x67\xb4\x53\xda\xd9\xb3\x6e\x92\xe5\x09\x09\x23\x42\x16\x0b\x3c\x29\x58\xf3\xe2\x60\xf4\x45\xf3\xf6\x64\xca\x1a\xb4\x1a\xae\x51\xe6\xe5\x88\x4a\xa3\xd5\xb5\x7a\x85\x6d\x4c\xbf\xaf\x61\x3b\xa5\x6a\xf4\x1d\x5c\xd3\xea\x9f\xd6\x4b\xc9\x7c\x9d\x7f\x85\x54\xd6\xf9\x43\xfb\x81\x8c\xb1\x31\xbe\xa6\x3f\xd0\xd6\xaf\xe5\x79\xc2\xc9\x7e\xce\xcf\x61\x22\xb8\x9a\x1e\x46\xef\x9f\xbf\xcb\x6f\x26\x70\x7d\xeb\x0d\xce\x8a\xb3\xae\x72\x76\x6c\xed\x97\x3a\x9e\xe3\xfd\x6f\x57\xbf\xab\x7d\xaf\x66\x0d\x90\xcb\x4a\x52\xb0\x34\xa5\x02\x7d\x57\x57\x4e\x95\x83\x70\x39\xc5\x5c\xda\xf6\xaf\xc2\xf6\x61\xed\xc9\xfc\xfe\x12\xff\x9b\x6d\x49\x42\x53\xc6\x87\x41\x23\x38\x68\xa0\xa0\x72\x9c\x7c\x87\x55\x5c\xc8\x20\x18\xfe\x4f\x47\x71\x01\xc6\x25\x4d\xa9\x08\xf1\x11\xec\xde\x37\x1c\x64\x94\xa7\x72\x13\x68\xf5\xe7\xd3\x34\x28\xbc\xc5\x32\xc4\xe2\x8a\x30\xfb\x7a\x4e\x84\x5e\x70\x34\x18\xe2\x69\x43\x05\x1d\x62\xb9\xc3\xcd\xc4\x2d\xbc\x95\x9b\x88\x50\xbe\x8e\x08\xe3\x05\x15\xd2\xaf\x90\x8f\xdb\x3e\xc6\xd9\x96\x16\x08\xae\x78\x2f\x1e\xfd\x0b\x00\x00\xff\xff\x39\xbf\x75\x3f\x33\x03\x00\x00")

func _1_create_table_documentUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_create_table_documentUpSql,
		"1_create_table_document.up.sql",
	)
}

func _1_create_table_documentUpSql() (*asset, error) {
	bytes, err := _1_create_table_documentUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_create_table_document.up.sql", size: 819, mode: os.FileMode(436), modTime: time.Unix(1593161462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x4d\x8f\xdb\xc8\x11\x3d\x93\xbf\xa2\x23\x60\x17\x62\xa0\x68\xf8\xfd\x21\x60\x2e\x6b\x3b\x80\x0f\xf1\x02\x59\xef\x29\x15\x0c\x9a\x64\xb7\x42\x44\x12\x65\x8a\x5a\xd7\x8c\x31\xff\x3d\x78\x5d\xad\x19\xad\xb3\xf6\x0c\x82\x1c\x28\x91\xcd\xee\xaa\x57\x1f\xef\x15\x6f\x6e\xd4\x9b\xb1\x37\x6a\x6b\x0e\x66\xd2\xb3\xe9\x55\x7b\xaf\xb6\xe3\x5f\xda\xe1\xd0\xeb\x59\xaf\xd5\xdb\x9f\xd5\x87\x9f\x3f\xaa\x77\x6f\xdf\x7f\x5c\x87\x37\x37\xea\x34\x9e\xa7\xce\x9c\x36\xb8\x4f\xee\xba\xc9\xe8\xd9\xdc\xcd\xba\xdd\x99\xbb\x7e\xec\xce\x7b\x73\x98\xd7\xfd\xf8\xf9\xb0\x3e\x7d\xda\x7d\x6f\xcf\xf9\x78\xd9\x71\x71\xb5\x1d\xc3\xa3\xee\xfe\xad\xb7\x46\xed\x87\xed\xa4\xe7\x61\x3c\x9c\xc2\x70\xd8\x1f\xc7\x69\x56\xcb\x30\x58\xb4\xf7\xb3\x39\x2d\xc2\x60\xd1\x8d\xfb\xe3\x64\x4e\xa7\x9b\xed\xc3\x70\xc4\x82\xdd\xcf\xf8\x1b\x46\xf9\xbd\x19\xc6\xf3\x3c\xec\xf0\x30\xba\x03\x47\x3d\xff\xeb\xc6\x0e\x3b\x83\x1b\x2c\x9c\xe6\x69\x38\x6c\xdd\xbb\x79\xd8\x9b\x45\x18\x85\xa1\x3d\x1f\xba\x0b\x9a\xbf\x1b\xdd\x2f\x71\xa3\xfe\xf1\x4f\xb8\x5d\xa9\x83\xde\x1b\x25\xc7\x22\xb5\xbc\xac\x9a\x69\x1a\xa7\x48\x7d\x09\x83\xed\x83\x7b\x52\x9b\x5b\x05\x54\xeb\x0f\xe6\x33\x8c\x98\x69\xe9\x60\xe3\xf9\xa7\xb3\xb5\x66\x72\x66\xa3\x28\x0c\x06\xeb\x0e\xfc\xe9\x56\x1d\x86\x1d\x4c\x04\x93\x99\xcf\xd3\x01\x8f\x2b\x65\xf7\xf3\xfa\x1d\xac\xdb\xe5\x02\x86\xd4\x0f\x9f\x36\xea\x87\xdf\x16\x82\xc4\xf9\x8a\xc2\xe0\x31\x0c\x83\xdf\xf4\xa4\xda\xb3\x55\xe2\x47\x9c\x84\xc1\x9d\xc0\xb9\x55\xc3\xb8\x7e\x33\x1e\xef\x97\x3f\xb6\x67\xbb\x52\xdb\x87\x28\x0c\xba\xdd\xbb\x0b\xd2\xf5\x9b\xdd\x78\x32\xcb\x28\xfc\x7f\xe1\x81\x19\xb1\xff\x0d\x43\x66\x9a\x04\xb7\x5f\x6c\xcf\x76\xfd\x13\xa0\x2f\xa3\x15\x76\x84\x8f\x61\x38\xdf\x1f\x8d\xd2\xa7\x93\x99\x91\xf2\x73\x37\xc3\x8a\x8b\xcf\xd7\x23\x0c\x86\x83\x1d\x95\x1a\x4f\xeb\xbf\x0e\x3b\xf3\xfe\x60\xc7\xa7\x73\xbe\x84\x97\xf5\x2b\x0b\xae\x86\x4a\xf9\x32\x86\xc1\x69\x78\x70\xcf\xc3\x61\x2e\xf3\x30\xd8\x83\x09\xea\xc9\xe8\xdf\xc6\xde\xb8\xc5\x8f\xc3\xde\x28\xb4\xc9\x1a\x77\xf0\xe3\x5a\x65\x69\x87\xaf\x7d\x45\xea\x83\xde\x9b\x65\xe4\x3d\xc0\xa7\x8f\xd2\x0e\x6b\x78\x0f\x1f\xbf\x73\xf6\x97\xe1\x01\x67\x1d\x9a\xdf\x1f\x05\xd0\xef\x1e\x05\xd6\x65\x74\x8d\xfc\xf7\x06\x10\xda\x4b\x06\x10\xdc\x32\x7a\x0e\xf4\xbf\x2c\xf8\xe8\xbf\x6d\xe4\xfd\xe9\xed\x30\x2d\x23\xd5\x8e\xe3\xee\xfa\xb4\xde\x9d\x5e\x88\xfc\xfe\x24\x81\x9b\xc9\xea\xce\x7c\x79\xbc\x3a\xed\x5b\x02\x5d\x7e\x77\xf7\x0d\x39\x79\x3b\x7e\x3e\xfc\xf2\x69\xa7\x6e\x7d\x7b\x2c\x17\xc4\x89\x25\xae\x5b\xe2\xb8\x26\x8e\xe3\x3f\xbe\xac\x25\xae\x52\xe2\xb8\x21\xb6\xf8\xb7\xc4\x45\xec\xcf\xf8\xb5\x2a\x27\xae\x2a\xb9\xe2\x9e\xb8\x48\x89\xd3\x9a\x38\xed\x89\xf3\x9a\x38\x6f\x88\xd3\x4e\xfe\xeb\x9e\x38\xb5\xc4\xa9\x79\x5e\x4f\xb1\x8e\x67\x4b\xdc\x19\xe2\x0e\xef\x5b\x79\xdf\xf5\x72\x7f\xd9\x93\xfa\x3d\x1a\xf6\x34\x71\x5b\x10\x9b\x92\xb8\xf6\x3e\x1d\x96\xdc\xe3\x4a\x88\x8b\x42\xe2\xe8\x60\x2f\x23\xce\x0b\xe2\xac\x23\xb6\xc0\xd9\xc9\x95\x66\xc4\x49\x2e\xeb\x7d\x2f\xf7\x97\x75\xac\xd5\x99\x9c\x85\x8f\x04\xe7\x53\xe2\x22\x23\x4e\x35\x71\x1e\x13\x67\x5a\xf2\x98\x65\xf2\x9f\xb7\x92\x03\x9b\x11\x37\xa9\xe0\x37\x99\xbc\x4f\xd2\xe7\xbd\x36\x7e\xce\x53\x87\x58\x11\x53\x2f\xcf\x97\xfc\x60\xcd\xe5\x31\x26\x2e\x4a\xc9\x7b\x9b\x10\xa7\x85\x5c\x99\x25\x6e\x1b\xe2\x0c\xf8\x2b\xe2\xac\x20\xd6\xc8\x51\x4e\x9c\x68\xe2\xe4\xab\x5a\xba\x2b\x23\x36\x3d\x71\xd3\x48\xee\x74\x76\x5d\xf3\xc5\x45\xe8\x5f\x68\x25\xaf\x48\x7f\xa4\xf4\x17\xdd\xba\x9a\x14\x61\x10\xbc\xd4\x9b\xab\x30\x08\x16\x2f\x4d\xcc\xc5\x2a\x0c\xa2\x27\x81\x79\xc1\x22\xd0\xfd\xd9\x89\xe4\x35\x3a\xa7\x92\x4f\xa3\xe8\x75\x51\xbe\xa4\xfd\x4f\x92\xed\x44\x77\x73\xfb\x35\x81\xbf\x40\xda\x36\xea\x15\xe1\x29\x28\xd9\x46\x25\x65\xb6\x52\xd0\xa4\xcd\xb5\x64\x2d\xf3\xac\x8c\xdc\x3a\x94\x66\x23\x4a\xf4\xeb\x61\xe0\x65\x52\x34\x59\x52\xd4\x75\x52\xad\x54\x1c\x3d\x86\x81\x06\x8a\x1f\x5d\xf0\x5f\x5c\xc4\x1b\xe5\x03\x07\xc4\x8d\xfb\x7d\x7c\x2a\x95\x5e\xbd\x4a\x45\x7e\x3d\xfe\xaf\x1a\x52\x76\xc4\x4d\x42\x9c\x27\xc4\xa5\x25\x6e\xc0\x81\x98\x38\xe9\x88\x3b\x70\x18\x7d\x69\x84\x7b\x78\xaf\x2b\xe1\x5b\x5f\x12\xe7\x9d\xf4\x6a\x9b\x13\xf7\x99\xd7\x1f\xd8\x88\x89\xdb\xfa\xb9\xd7\xc1\xa9\xa6\x10\xde\xf7\x9a\xd8\x16\xc4\x75\x2e\xfb\x5b\x9c\x29\x89\x0b\x4d\x5c\xa6\xd2\xf3\xd8\x03\x5c\xd0\x98\xcc\x88\x26\x94\x39\x71\x9c\x0a\xf7\x70\x35\x95\x70\x3f\xcd\x89\x0d\x70\x58\xe2\xde\x3e\x73\xa9\x37\xc4\x45\x2b\x5a\x08\x0c\xe0\x7e\x0c\xde\x41\xdb\x2a\xf9\x87\xe6\x94\x3d\x71\x55\x8a\x0f\x97\x03\xac\x57\xe2\xcb\x5d\xb1\x70\xbf\x2c\x88\x93\x96\xb8\xf5\x1a\xa5\x81\xad\x14\xce\x3b\xee\x1b\xaf\xa5\x5e\x9b\xb0\x66\x4b\xe2\x0c\xba\x55\x8a\x76\x5a\x68\x6e\x29\x58\x91\xcb\xb8\x12\x4d\x2c\x2b\xc9\x61\xe9\xed\x42\xc3\x2a\x2d\xba\x0d\x9b\x75\x25\x7b\x12\xbf\xde\x6a\xc1\xdc\x21\x07\x46\xf2\xd4\x20\xd7\x86\xb8\x84\xff\x44\xf4\x05\x7e\x90\x27\x68\x2c\x9e\x81\x0b\xcf\x26\x26\xee\x5b\xb1\x8d\x98\xfa\x46\xb0\xea\x58\xf4\x5b\x23\x57\x46\x34\x17\x35\x4c\x4b\xd9\x8b\x1c\x40\x7f\xa1\xf7\xe8\xaf\x2a\x13\x1c\x17\xfc\x85\xaf\x1d\xec\xb5\x19\x71\x69\x44\x67\x4d\x21\x75\xc0\x05\xed\xce\x53\xc9\x49\xdc\x0a\x26\xcc\x90\xa2\xf6\xfa\x97\x12\x97\xe8\xcd\x5c\xea\x83\x35\x60\x43\xfd\x3b\x2d\xb5\x6c\x7d\x5f\x69\xd4\x38\x11\xad\x84\x8f\xba\x96\xd9\xe3\x7a\xa8\x90\x19\x84\x7c\xd5\xbe\xce\xe8\xd7\xd6\x8a\x0e\x97\x09\x71\xd6\x12\x6b\xbc\xd3\xe2\x2f\xf1\xb3\x07\xbd\x6c\xb4\xef\xa9\x52\x7a\x07\x73\xa0\x42\xfe\xd0\x2f\x56\x6c\x16\xb9\xbc\xb7\x8d\xd4\xdd\xb6\xc4\x75\x42\x5c\x77\xd2\x17\xa8\x03\xfa\x42\x97\x32\x0f\xfa\x58\xf6\xc3\x3f\xf0\x56\x98\x23\xa9\xb7\x6d\x9e\x67\x2c\xb0\xa5\xa9\x70\xa8\xd1\x52\x57\xf4\x0d\xb8\x08\x4c\x88\x01\x35\x45\x0f\xa1\x36\x55\x4c\x5c\xf5\xc4\xa6\x95\x39\x0f\x3b\xb5\x96\x1a\x20\x47\xf8\x4e\x70\x75\xea\xfc\x7c\xa9\x64\x36\x36\x46\xe6\x9f\xe3\x17\xbe\x1f\x2a\x9f\x9f\x56\xec\xb9\x3c\x95\x62\xb3\xb9\xcc\x6a\x2d\x33\x14\xbc\xc5\x6c\x43\xfe\xc0\x13\xec\xc7\xf9\xc2\xf7\x24\xf8\x8e\xf9\x0c\x6c\x59\x23\x5a\xe2\x72\xab\x7d\xaf\x59\x89\xdd\xd5\x36\x11\x5c\xc8\x83\xed\xbc\xde\xa4\xc2\x63\xe8\x91\xab\x5d\x23\x7d\x83\x1e\xeb\x4a\xe1\x04\xf8\xea\xbe\x05\x6a\xe2\xb2\xf6\xfd\x1b\x4b\xbc\x4d\x27\xf5\xc2\xbe\xc2\xfb\xee\x30\xd3\xa1\x6d\xb5\xf8\xc0\x77\x0d\xb8\x0e\x5d\x8a\x13\xb1\x8b\xd9\x0d\xbe\xeb\xc6\x73\x18\x5c\x4a\x44\x63\xa0\x29\x98\xe5\xd0\x46\xf4\x23\x38\x07\x3d\x84\x3f\x70\x13\x18\xa0\x6d\xa6\x92\x1e\x42\x7e\xf0\x1d\xd1\x7a\x2d\xc5\xf7\x09\xf8\x84\x1e\x47\x7d\xa0\xb3\xe8\x19\xf4\x28\x72\xe6\x34\xd1\xf8\x9c\xe5\x82\xd3\xf1\xc1\xf7\x57\x5c\x48\x8f\x42\x27\xf1\x6d\xd1\x65\xc2\x7d\xc7\xd1\x5e\xfc\x20\xef\xc8\x19\xb8\xe0\xf4\xd5\x08\x5f\x11\x8b\xb9\xd8\x28\x88\xfb\x54\x6a\x80\xba\xd6\x56\xf4\x00\x5a\x8b\x1c\x80\xbf\x4d\xe9\x39\x5a\xfb\x1e\xaa\xe5\x1b\x2f\xf1\x7a\x0c\xee\x7e\xfd\xfd\x82\x3a\xa3\x7f\x2a\xff\xfd\x03\xdd\x8f\xb3\x57\x7f\xbf\xb8\x21\xf6\xea\xaf\x97\xff\x04\x00\x00\xff\xff\xcd\xe8\xb4\x76\x00\x10\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 8192, mode: os.FileMode(436), modTime: time.Unix(1593161692, 0)}
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
	"1_create_table_document.down.sql": _1_create_table_documentDownSql,
	"1_create_table_document.up.sql": _1_create_table_documentUpSql,
	"bindata.go": bindataGo,
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
	"1_create_table_document.down.sql": &bintree{_1_create_table_documentDownSql, map[string]*bintree{}},
	"1_create_table_document.up.sql": &bintree{_1_create_table_documentUpSql, map[string]*bintree{}},
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
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
