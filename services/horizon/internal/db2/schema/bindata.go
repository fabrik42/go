// Code generated by go-bindata.
// sources:
// latest.sql
// migrations/1_initial_schema.sql
// migrations/2_index_participants_by_toid.sql
// migrations/3_use_sequence_in_history_accounts.sql
// migrations/4_add_protocol_version.sql
// migrations/5_create_trades_table.sql
// migrations/6_create_assets_table.sql
// DO NOT EDIT!

package schema

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

var _latestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5b\x5b\x6f\xdb\x38\x16\x7e\xcf\xaf\x20\xe6\xc5\x36\x60\x07\x71\x9b\x5b\x1d\xa4\x80\xeb\x68\xb6\xc6\xb8\xca\x34\x56\xb6\x53\x0c\x06\x04\x2d\xd1\x36\xb7\x92\xa8\x92\x74\x9a\xcc\x62\xff\xfb\x42\x37\xeb\x46\x8a\x92\xad\x74\x1e\x63\x1d\x7e\xfc\x3e\xf2\x90\xe7\xf0\x90\x19\x8d\x4e\x46\x23\xf0\x3b\xe5\x62\xc3\xf0\xf2\xf3\x02\x38\x48\xa0\x15\xe2\x18\x38\x3b\x2f\x38\x19\x8d\x4e\xc2\xef\x77\x3b\x2f\xc0\x0e\x58\x33\xea\x65\x06\x4f\x98\x71\x42\x7d\xf0\xee\xf4\xf2\xf4\x22\x67\xb5\x7a\x01\xc1\x06\x86\xcd\x4b\x26\x27\x4b\xc3\x02\x5c\x20\x81\x3d\xec\x0b\x28\x88\x87\xe9\x4e\x80\x5b\x70\x76\x13\x7d\x72\xa9\xfd\xad\xfa\xab\xed\x92\xd0\x1a\xfb\x36\x75\x88\xbf\x01\xb7\xa0\xf7\x68\xfd\x7a\xdd\xbb\x49\xe1\x7c\x07\x31\x07\xda\xd4\x5f\x53\xe6\x11\x7f\x03\xb9\x60\xc4\xdf\x70\x70\x0b\xa8\x9f\x60\x6c\xb1\xfd\x0d\xae\x77\xbe\x2d\x08\xf5\xe1\x8a\x3a\x04\x87\xdf\xd7\xc8\xe5\xb8\xd0\x8d\x47\x7c\xe8\x61\xce\xd1\x26\x32\xf8\x81\x98\x4f\xfc\xcd\x4d\xc2\x1d\x23\x66\x6f\x61\x80\xc4\x16\xdc\x82\x60\xb7\x72\x89\x3d\x0c\xc5\xda\x48\x20\x97\xa6\x66\x0e\x5e\xa3\x9d\x2b\xa0\x40\x2b\x17\xf3\x00\xd9\x38\x24\xdd\x2b\x7d\xfd\x41\xc4\x16\x52\xe2\xe4\x78\x9c\xc4\xb3\x61\x22\x0f\x4f\xc0\x86\xb2\x00\x7a\x64\xc3\x50\xc8\x99\xdf\x00\xeb\x25\xc0\x13\x60\x4d\x3f\x2c\x8c\x1b\xb0\xb4\xb7\xd8\x43\x93\x84\xc4\x0d\xb8\xff\xe1\x63\x36\x01\xa3\x68\xc6\x66\x0f\xc6\xd4\x32\x62\xd3\x32\x0e\xe8\x9f\x00\x00\x00\x71\x80\xc0\xcf\x02\x98\xf7\x16\x30\x1f\x17\x8b\x61\xf4\x2b\x0a\x02\x97\x60\x07\x22\x01\xc2\x79\xe0\x02\x79\x01\x08\x89\x46\x7f\x82\xbf\xa9\x8f\x4f\x06\x37\x27\x45\xa2\x5b\xc2\x05\x65\x2f\x10\xd9\x36\xdd\xf9\x82\x43\xe2\x40\x8e\xbf\xa7\x84\x97\xc6\xe7\x47\xc3\x9c\x35\xe4\x9c\x5a\xab\x50\x23\x9a\x4b\x6b\xfa\x60\x81\x2f\x73\xeb\x23\x18\x47\x3f\xcc\xcd\xd9\x83\xf1\xc9\x30\x2d\xf0\xe1\x6b\xf2\x93\x79\x0f\x3e\xcd\xcd\x7f\x4f\x17\x8f\xc6\xfe\xef\xe9\x1f\xd9\xdf\xb3\xe9\xec\xa3\x01\xc6\x3a\x31\x07\x0f\x7b\x19\x28\x1b\xf7\x15\xd9\x10\x5f\x80\x3b\xe3\xd7\xe9\xe3\xc2\x02\x3e\x7e\x16\x4f\xc8\xed\xf7\x14\x8a\x7b\x93\x09\xc3\x1b\xdb\x45\x9c\x0f\xca\xd3\xe5\x38\x0c\x73\x0e\xec\x2d\x62\xc8\x16\x98\x81\x27\xc4\x5e\x88\xbf\xe9\x5f\x9e\x0f\x6a\x26\x8a\x73\xdc\x85\xb2\x08\x26\xd3\x45\x7c\x81\x37\x98\x95\x39\x86\x56\x50\xbc\x04\x58\x4e\x53\x6a\x6e\x53\x47\x66\x3e\x7e\x23\x37\x27\x9c\xef\x30\x93\x34\xb8\xb8\xcc\x1a\xe8\xc6\xa3\x63\xb7\xcd\x63\xfe\x34\xa7\xad\x13\x02\xee\xbf\x98\xc6\x1d\xf8\xf0\x55\xa3\x68\xba\xb0\x8c\x07\x8d\xa0\x3d\x56\xe9\xf3\x29\x71\x54\xdc\xf0\x7a\x8d\xed\x0e\xbc\x2e\xc1\x49\xdc\xae\xb4\x66\x60\xb6\xbc\x8a\x7e\x92\xda\xd1\x00\xc7\xfb\xa0\xd2\xf2\x17\xca\x1c\xcc\x7e\x51\x78\x73\xe4\xc7\xf2\x4f\x0e\x16\x88\xb8\x1c\xfc\x87\x53\x7f\xa5\x76\x36\x17\x3b\x1b\xcc\x8e\x1f\x87\x04\x27\x19\x07\x8e\xbf\xef\xb0\x6f\xab\xb8\xc5\xc6\x70\x8b\xf8\xb6\xd1\x2a\x0c\x18\x7e\x22\x74\xc7\xa1\xb6\x61\x32\x2c\x0c\xf9\x1c\xc5\xc1\x35\x9a\x88\x3d\x8f\x74\x97\x3b\x2b\xf5\x90\x4d\x44\x33\x7b\xdb\xa5\x5c\x16\x98\xc2\x54\x61\x1f\x9b\xca\x6d\x18\x46\x42\xdb\x28\xb6\xdd\x05\x4e\x63\xdb\xbd\xeb\x24\x7f\x7a\x01\x65\x02\x33\x98\x66\x3b\x65\x2d\xe3\xb2\x13\x51\x81\x5c\x68\x53\xe2\x73\xb9\x0f\xae\x31\x86\x01\xa5\xae\xfc\x6b\x98\x7c\xc1\x35\x56\xcd\x75\xf4\x99\x61\x8e\xd9\x93\xca\xc4\x43\xcf\x50\x3c\xc3\x70\xeb\xe4\xe4\x6f\x95\x55\xc0\xa8\xa0\x36\x75\x95\xba\xce\x1a\xec\xad\xd9\x3c\x07\x88\x09\x62\x93\x00\x75\x11\x55\xe5\xb0\xba\x58\xd4\x7c\x17\xd0\xef\x2b\x6d\x25\x77\x1b\x5e\x6a\xfb\xf8\x59\xe1\xa6\x95\xd0\x23\xc3\x4f\x6d\x5f\xd5\x70\x24\x37\xaf\x09\x4f\xfb\x06\x1d\xfa\x66\x35\xe7\x2b\xed\x03\xb9\x5d\x53\x65\x13\x65\xe4\x76\x2c\x25\x8a\x4c\x47\x06\xa6\xf8\x27\x4e\x77\xcc\xc6\xa9\x77\x2b\x42\x42\xba\xcc\x7b\xbd\xc9\xa4\x62\xd1\x60\x1d\x08\x86\x1c\x7c\xfc\x70\xc6\x30\xa5\x78\x7f\x6c\x1c\xa7\xeb\x35\x66\xca\xb6\x1c\xbb\x6e\xcd\xe7\xd5\xee\xa5\xae\x31\x75\x1d\xd8\x32\xeb\xcd\xb5\x69\x94\xcb\x56\x5a\x35\x4e\x98\xe3\x36\x5e\x34\xed\x72\x71\x74\xb7\xd9\x8a\xb6\x02\x0a\xad\x5a\x48\x28\xb4\x6b\x2c\x22\x6d\x55\x23\x63\x76\x6f\x2e\xad\x87\xe9\xdc\xb4\x4a\x8e\x04\x0b\x8d\x61\x54\x19\x00\xb3\x8f\xc6\xec\x37\xd0\xef\x17\x81\xdf\x83\xb3\xc1\x40\x07\x97\x1b\xd0\x12\x58\x7e\xa8\x23\xa8\xda\xa5\xb2\xdf\x09\x3a\x8d\x93\x2a\xe0\xa6\x91\xb2\xc9\x16\x75\x4c\xac\x54\xf1\xeb\x36\x5a\x6a\x7a\xf9\x59\xf1\xb2\xa5\xd8\x23\x23\xa6\xa6\xb7\x6a\xcc\x54\x35\xa8\x89\x9a\xb9\x26\x9d\xfa\x6a\xea\x9f\x79\x4a\x8d\x0f\x2f\xc9\x99\x45\x73\x24\x6a\x1a\x58\xeb\x63\xa4\xd4\x36\xeb\x5a\x9d\xdd\x23\xe5\xd2\x53\x9d\x8c\xfe\x91\xb3\x8d\x78\x86\xd8\x7f\xc2\x2e\x0d\xb0\xac\x5e\x28\x9e\xc3\x93\xc6\xce\x15\x8a\x8f\x1e\x16\x48\xf1\x29\x1c\x05\xd5\x67\x4e\x36\x3e\x12\x3b\x86\x65\xa5\xad\x77\x97\x83\x3f\xff\xca\x92\x93\xff\xfe\x4f\x96\x9e\xfc\xf9\x57\xf9\xc8\x83\x3d\xaa\x08\x67\x19\x96\x4f\x7d\x5c\x9b\xec\x64\x58\x55\x98\x44\x19\xf1\x70\x18\x62\x7c\x87\x87\x33\x77\xcd\x90\xbf\xa9\xab\x99\xc6\x35\x34\xe2\xa4\xab\x27\xe1\xd2\x68\xc9\xc7\xcb\xe7\xde\x5c\x94\xcb\x30\x20\xfe\x3e\xbb\x5f\x3c\x7e\x32\xc3\x29\x5d\x1a\x56\x4d\xbd\x31\x5f\xd9\xc9\x57\x1b\xdb\xe5\xfd\xdd\x89\x50\xe0\xb7\x12\x55\x7b\x5e\x68\x22\x52\x19\x39\x3b\x93\xa9\xec\xa1\x95\x50\xcd\x36\x2f\x97\x7a\x87\x04\x02\x6b\xca\x1a\x5c\x35\x80\xbb\xa9\x35\xd5\x48\x9c\x9b\x4b\xe3\xc1\x02\x73\xd3\xba\xaf\x5c\x37\x44\xd1\x71\x09\xfa\xbd\x31\x24\x3e\x11\x04\xb9\x90\x47\x58\xa7\xfc\xbb\xdb\x1b\x82\xde\x9b\xb3\xf1\xd5\x68\x7c\x36\x1a\x5f\x83\xf1\xd9\x64\x7c\x35\x19\x8f\x4f\xcf\xaf\xae\xce\x2f\xc6\xa3\xb3\xab\xde\xe0\xa6\x19\xfa\x1b\x48\x7c\x07\x3f\x17\x87\x60\xf5\x02\x05\x25\x4e\x6d\x4f\xd7\xe3\x8b\x37\xe7\x6d\x7a\x7a\x0b\x77\x1c\xef\xb7\x78\x48\x7c\x58\x2e\xdc\xd7\xf7\x77\xfe\xf6\xa2\x4d\x77\xe7\x10\x39\x0e\x2c\xd7\x62\xea\xbb\x78\x77\x39\xbe\x6c\xd3\xc7\x05\x8c\xc3\x49\x9a\xd3\x46\x17\x57\xb5\x5d\xbc\xbb\xb8\xba\x6e\x35\x3f\x97\x69\x17\xc9\x6e\xd3\xa0\x8b\x77\xd7\x6f\xaf\x92\x2e\x14\x7e\x5b\x7b\x59\xd3\xc4\x71\x0f\xba\xc8\x0a\xd7\xa3\x06\x77\x69\x2c\x8c\x99\x95\xbb\x19\x3c\xe5\xb8\xfe\x92\x67\x08\xc6\xc3\xf8\x1a\xb0\x81\xdc\xea\xfd\xcd\x11\x62\x6b\xef\x0c\x3a\x91\x5a\x88\x2f\x6d\x84\xca\xee\x0c\xda\x28\x55\xc0\xca\x4a\xf0\x1d\xc0\x36\x28\x75\x1e\x3e\x4d\xed\x6a\x6d\x5d\x4c\x5b\x7d\x04\x6d\x33\x8d\x8a\xda\x5a\x07\x43\x2e\x29\x31\x75\x83\xaa\x3f\x8d\x1f\x3e\x95\x6d\x8f\x81\x5d\x4c\xa6\x2e\x4b\x68\x33\x9d\xca\x43\x5f\xfb\x21\x29\x87\x8b\xd2\xdf\x30\xf8\x86\x5f\xd2\x2e\xb2\x12\x4c\xdb\x84\xab\x84\x1a\xe5\xe9\xd3\xbb\xbb\x7c\x51\x47\xd6\x31\xf8\xfd\x61\xfe\x69\xfa\xf0\x15\xfc\x66\x7c\x05\x7d\xe2\xe8\x92\xf8\xd2\x7e\x97\xd5\xb4\x72\xc5\xb4\x42\x85\x0c\x76\xa2\xae\xd8\xad\x4c\xdc\x41\xc4\xc0\xa3\x39\xff\xfc\x68\x80\x7e\x66\x3e\xcc\x5d\xe6\x0f\x0b\x57\xef\x2d\x87\xa6\x9b\x69\x6d\x2d\xbc\xd5\xa4\x2a\x0e\x21\x9a\xdd\xb1\x5b\x65\xf2\x4e\xea\x94\xd6\xd0\x6a\xac\x5c\x79\x2e\xd1\x6e\x26\xdd\xaa\x57\x75\x53\xa7\xbf\x96\x9a\x76\x04\x62\x97\x5e\xbd\x24\x5e\x9d\x4a\x99\x9b\x77\xc6\x1f\xcd\x2a\x5a\x91\x69\x19\x07\xdc\x9b\xe5\x05\xf1\xb8\x9c\x9b\xff\x02\x2b\xc1\x30\x4e\x57\x98\x62\x25\xad\xf6\xb9\xe2\xc1\x74\x32\x88\x3c\x93\x42\xb9\xad\xc8\x27\x36\x1e\x56\xea\x59\x32\x72\x5b\xc4\xb7\xc7\x30\x8b\xca\x7a\x8d\x68\x95\x8b\x81\x32\x36\x71\x6a\x77\x0c\x9f\x18\xa1\x19\xa3\x52\xa5\x71\x58\x2d\x2a\x4a\x17\x19\xc4\xa1\x6f\x44\xdf\x0f\x60\x9a\xec\xcb\x31\xe1\x12\x5c\x9e\x76\xfa\x50\xa6\xc0\x58\x76\x7f\x36\x4c\xef\xca\x54\x64\xb3\x8a\xc7\x91\x34\x89\xd3\x98\x60\x76\x99\x30\x94\x5e\xfa\x69\x48\xd3\x00\x06\x5d\xf1\x4e\xb0\xf2\xd4\x15\xc1\xe1\x20\x25\x72\x01\xe2\xb9\x3b\x01\x09\x96\xc2\xa7\x0f\x94\x50\xbc\x19\xaa\x8a\xa0\x41\xe8\x95\x5b\x7a\x90\x86\x84\x7c\x86\x71\xe8\xe0\xd7\x0f\xf4\xfe\x7d\x53\xb8\x55\x1f\x3f\xd6\x45\xb8\x3c\xe5\xf4\xb1\x56\x81\xa3\x9c\x51\x7e\x5c\xbb\xa2\x55\xc1\x6c\xb6\xbd\xc9\x08\x8a\x78\x4a\xc4\x31\xd3\x9a\x61\x1c\xee\x92\x3a\xf7\x13\xd1\x2c\xc4\xf7\xb9\x47\x30\xcd\xa1\x94\xb8\x3a\xb8\xc4\xac\x72\x71\x3e\xac\xde\x6e\x0f\x65\x17\xe5\x2a\xf2\x9c\xba\xc7\x0c\xf2\x1e\x43\x47\xbc\xf4\x60\x61\x58\x7e\x57\x30\xac\x3e\x4f\x90\x51\x76\xa2\x28\xb4\x5e\x1f\x11\x7e\x0b\x28\x3a\xda\xe9\x13\x0e\x39\x97\xa0\x83\x85\x93\xe0\xe8\x88\xb4\x0b\x4f\x71\x99\xba\x52\x06\xa4\x3e\x4c\x5e\x73\x1f\x4b\x5b\xdb\x41\x21\x1d\x4d\x5f\xa7\x17\x13\xc0\xd8\xb0\x05\xf7\xe3\x47\xbb\x0e\x5b\xcf\x58\xe2\x06\x45\xc0\x24\xd9\x08\xf1\x42\x27\x3f\xd8\x45\x6b\x51\xb5\xd9\x4d\x68\xa4\x21\x9a\x84\x8a\x10\x72\xff\xe6\xb5\x23\xb6\x32\x68\x6d\x94\xda\x5b\x36\xe7\xdd\xb5\x33\x14\xa0\x0f\x09\xab\x6a\xb8\xd2\xd3\xdd\xee\x07\xba\xf2\x38\x58\x4b\xbf\xd4\xa0\xb9\x98\xdc\x5b\xed\x57\x1b\xff\xfc\x7b\x70\x9d\x92\x9c\x6d\x73\x11\xb2\x97\xe7\xaf\xa6\x46\xfa\xcc\x5d\x27\x4b\xd6\xa8\xb9\xbe\xf4\xac\xf8\x6a\x9a\xf6\x6f\x4f\x74\x3a\x94\x87\xfa\x22\x74\x56\xbc\x7f\x8d\xa5\x5d\x46\x97\xe6\xf9\x6d\x17\x78\x11\xb4\x98\x29\x76\xb4\xc2\xeb\xba\x68\xa2\x41\x93\xbe\xd6\x76\xd6\x5d\xf8\xaa\x02\x37\xe2\xae\x0f\x62\xf9\x33\xc5\x6b\xb8\x4d\x15\xff\xe0\x13\x4d\x94\xd1\xed\x03\x79\x5a\x48\x81\x2b\x4a\xbf\x1d\x3c\xca\x35\x98\xda\x14\xa1\xdf\x4f\xdf\x6b\x8f\xde\xbf\x07\xbd\x52\x72\xde\x9b\x4c\x04\x7e\x16\x83\xc1\x10\xa8\x0d\xc3\xa4\xbd\x91\x61\x9c\xcc\xab\x4d\x2b\x47\x9a\x86\xa6\xf5\x04\x24\x47\xa0\xbd\xf1\x00\x7c\xf9\x68\x3c\x18\xb1\x93\x81\x5b\xf0\xf6\x6d\x6e\xc2\x54\xff\xa5\x0c\x6c\xea\x05\x2e\x16\x38\x9a\x89\xff\x07\x00\x00\xff\xff\x7f\xa0\x0a\x87\xd2\x3c\x00\x00")

func latestSqlBytes() ([]byte, error) {
	return bindataRead(
		_latestSql,
		"latest.sql",
	)
}

func latestSql() (*asset, error) {
	bytes, err := latestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "latest.sql", size: 15570, mode: os.FileMode(420), modTime: time.Unix(1508347260, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1_initial_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\x5f\x6f\xdb\xc8\x11\x7f\xf7\xa7\x18\xdc\x8b\x6c\xd4\x6a\x2f\xb8\xe2\x70\x95\xe1\x03\x14\x99\x69\x84\xca\x54\x22\x51\x4d\x82\xc3\x61\xb1\x22\x47\xd4\xd6\xe4\x2e\xb3\xbb\x74\xa4\x2b\xfa\xdd\x0b\x52\x24\xc5\xff\xa4\x1c\xc9\xf7\x28\xee\xec\xcc\xfc\x66\x66\x7f\x33\x5c\x6a\x38\x84\xbf\xf8\xcc\x95\x54\x23\xac\x82\xab\xe1\xf0\x6a\x38\x84\x0f\x42\x69\x57\xe2\xf2\xe3\x0c\x1c\xaa\xe9\x9a\x2a\x04\x27\xf4\xe3\xe5\xab\xa5\x61\x81\xd2\x54\xa3\x8f\x5c\x13\xcd\x7c\x14\xa1\x86\x7b\xf8\xf1\x2e\x5e\xf2\x84\xfd\x54\x7d\x6a\x7b\x2c\x92\x46\x6e\x0b\x87\x71\x17\xee\x61\xb0\xb2\xde\xfd\x32\xb8\x4b\xd5\x71\x87\x4a\x87\xd8\x82\x6f\x84\xf4\x19\x77\x89\xd2\x92\x71\x57\xc1\x3d\x08\x9e\xe8\xd8\xa2\xfd\x44\x36\x21\xb7\x35\x13\x9c\xac\x85\xc3\x30\x5a\xdf\x50\x4f\x61\xc1\x8c\xcf\x38\xf1\x51\x29\xea\xc6\x02\xdf\xa8\xe4\x8c\xbb\x77\x57\x09\x3c\x93\xfa\x38\x82\xc0\x0b\x5c\xf5\xd5\xbb\x03\x6b\x1f\xe0\x08\x8c\xcf\x96\x61\x2e\xa7\x73\xf3\x0e\x96\xf6\x16\x7d\x3a\x82\xe1\x1d\xcc\xbf\x71\x94\x23\x18\xc6\xc8\x27\x0b\x63\x6c\x19\x47\x49\x98\xbe\x03\x73\x6e\x81\xf1\x79\xba\xb4\x96\xa9\x42\xf8\x34\xb5\xde\xc3\x72\xf2\xde\x78\x1c\x43\xe0\x12\x9b\x6a\xea\x89\xc8\x7a\xc1\xfc\x51\x4b\xc9\x91\xc9\xfc\xf1\xd1\x30\xad\x16\x37\x0e\x02\x30\x37\xab\x4a\x60\xba\x84\xc1\x87\xd9\xdf\x02\x37\x4a\x5e\x20\x85\x8d\x4e\x28\xa9\x07\x1e\xe5\x6e\x48\x5d\x1c\x94\xfd\xd8\x2a\x2d\x24\x9e\x2f\x0a\x07\x7d\xc5\x20\x84\x6b\x8f\xd9\xcd\x01\x28\xba\xf0\x32\xfc\x89\xd9\x08\x7e\x54\xb2\xa0\xf7\x01\xc2\x46\x48\x88\x9e\x47\x15\xa7\x50\x2b\x10\x1b\xb8\x7e\xc2\xfd\x2d\x3c\x53\x2f\xc4\x1b\x08\x28\x93\x2a\x0e\x49\x5c\x86\x48\xa5\xbd\x25\x01\xd5\x5b\xb8\x4f\xbc\xbe\x2d\xa6\x30\x12\x73\x70\x43\x43\x4f\x13\x4d\xd7\x1e\xaa\x80\xda\x18\x95\xf3\xa0\xb4\xfa\x8d\xe9\x2d\x11\xcc\xc9\x55\x68\x31\xee\x2c\xf2\x6c\x4f\xa8\x6d\x8b\x90\x6b\x95\xc2\xb7\xc6\x6f\x67\xc6\x11\x7c\x12\xbb\x2c\x02\x77\x60\x65\x66\x47\xf9\x7c\xc4\xfb\x2a\x5a\xe1\xfa\x0a\x00\x80\x39\xb0\x66\x2e\xe3\x3a\xce\x94\xb9\x9a\xcd\x6e\xe3\xe7\xd4\x71\x24\x2a\x05\xf6\x96\x4a\x6a\x6b\x94\xf0\x4c\xe5\x9e\x71\xf7\xfa\xe7\xbf\xdf\x5c\xdd\x54\x6a\x25\xd1\x8e\x9b\x0d\xda\xe7\x76\x39\x51\x9a\x78\x5c\x02\x42\x9a\x10\xa4\x72\x22\x40\x49\x63\x5e\x68\x92\xfc\x41\x48\x07\xe5\x0f\xc0\xb8\x46\x17\x65\x69\x35\xae\x97\xfa\x25\x07\x35\x65\x9e\x82\xff\x28\xc1\xd7\xcd\x41\xf1\xd0\x71\x51\x9e\x39\x28\x89\xd2\x24\x28\x0a\xbf\x86\xc8\xed\x26\x47\x0f\xc2\x64\x4b\xd5\xb6\x3e\xa3\x25\xf9\x40\xe2\x33\x13\xa1\x22\x9d\x1b\x93\x18\x49\xca\x15\x3d\xb0\x6f\x9c\x95\xcc\x8f\x07\xe3\xdd\x78\x35\xb3\xe0\xc7\x92\x85\x63\x56\xfa\xc9\xdb\x9e\x50\xe8\x10\xaa\x21\xea\x20\x4a\x53\x3f\x80\xe8\x20\x45\xbd\x24\x7a\x02\x7f\x08\x8e\xe5\x3d\x12\xa9\xee\xdc\x74\x90\x0d\x03\xa7\xb7\x6c\x56\x47\xc9\x4f\x3f\x10\x52\xa3\x24\xcf\x28\x15\x13\xbc\x82\xe5\x4d\xb9\xa2\x84\xa6\x1e\xb1\x05\xe3\xaa\xbe\x20\x37\x88\x24\x10\xc2\xab\x5f\x8d\x9a\x2e\xd9\x60\x53\xae\xe3\x65\x89\x0a\xe5\x73\x93\x88\x4f\x77\x44\xef\x88\x42\x4d\x14\xfb\xa3\x2a\xd5\x5c\xca\xc7\xb4\x05\x54\x6a\x66\xb3\x80\x9e\x9d\xa1\xea\x6d\x1c\xf9\xaa\x1e\x53\xff\xe3\xde\x4d\x20\xa7\xe2\x27\xcc\x21\x0a\xbf\xa6\x61\x58\x1a\x1f\x57\x86\x39\x69\x89\x44\x1e\x7c\x2a\xdd\xcf\x46\x8c\x60\x69\x8d\x17\xd6\xa1\x91\xbe\x89\x1f\x4c\xcd\xc9\xc2\x88\x5b\xdf\xdb\x2f\xc9\x23\x73\x0e\x8f\x53\xf3\xdf\xe3\xd9\xca\xc8\x7e\x8f\x3f\x1f\x7f\x4f\xc6\x93\xf7\x06\xbc\x39\x0b\x50\x98\x7f\x32\x8d\x07\x78\xfb\xa5\x03\xf1\x78\x66\x19\x8b\x13\x01\x67\xba\x3b\xc4\xff\xca\x9c\x4e\x2c\x97\x2a\xd4\xae\x66\x9a\xa7\xc7\xc6\x86\x1b\x04\x1e\xb3\x0f\xb8\xe2\x7e\xf4\x9d\xed\xe8\xf0\x48\x89\x50\xda\x98\x96\x7a\x03\xf7\xa7\x3c\x35\x18\x8c\x46\x15\x89\x1e\x87\x22\x0f\xef\x72\xb4\xd0\x64\x25\x8e\x7d\x03\x2d\xd4\xed\xad\x4f\xc0\xf7\x90\x42\x93\x67\xe7\xa5\x85\x0e\x2b\xaf\x45\x0c\x27\x82\xfd\x4e\x6a\xe8\xb0\x56\x25\x87\xa6\x0d\x2d\xf4\x90\xdb\x72\xb9\x92\x4d\x29\x22\xef\x5f\xef\x71\x2c\x99\xc2\x3a\x86\xbc\xbe\x0c\xd2\x4e\x06\xb5\xb2\x47\xd3\xcd\xf3\x0a\x6d\x6c\xcd\x4d\xb3\xde\x9f\x32\xad\xe9\x1d\x41\xfe\x8c\x9e\x08\x10\x34\xee\x2a\x54\xbd\x8b\x66\xa7\xd0\xd3\x0d\x8b\x3e\x46\xaf\x90\xb5\x4b\x51\x14\x9a\x96\x15\x73\x39\xd5\xa1\xc4\xba\x37\xaa\x7f\xfc\x7c\xf3\xdb\xef\x47\x16\xfe\xef\xff\xea\x78\xf8\xb7\xdf\xcb\x43\x1c\xfa\x82\xc4\xdd\xa0\xca\xd9\x99\x2e\x2e\x38\xb6\xb2\xfa\x51\x57\x55\x4d\x82\x8c\xf9\x48\xd6\x22\xe4\x8e\x8a\x32\xf7\x8b\xa4\xdc\xc5\x98\x0c\xf3\x87\x89\x39\xe9\xd1\x49\x6c\xf7\x3a\xef\x87\xe3\x32\x37\x67\x5d\xdd\x1d\x0e\xf2\x93\xf9\x6c\xf5\x68\x46\x29\x8d\x5e\xa8\x53\x94\x1c\x77\xfa\x99\x7a\xd7\x83\x5e\x03\xc5\x60\x34\x92\xe8\xda\x1e\x55\xaa\xc2\xe8\x67\x43\xd1\xd8\xac\x4e\xc2\xd1\xc1\x7e\x6d\x48\x3a\x42\x11\x3c\xe1\xfe\x78\xad\x62\x2e\xad\xc5\x78\x6a\xb6\xa0\xad\x12\xde\x89\x09\x8c\x4b\x69\xfc\xf0\x90\xb3\xd6\xc7\x47\xf8\xb0\x98\x3e\x8e\x17\x5f\xe0\x5f\xc6\x17\xb8\x66\xce\xe9\x3d\xf8\x82\x48\x9b\x6c\xb6\x61\x6d\xf5\xb3\x13\xed\x3a\x1b\x50\x52\x48\x53\xf3\xc1\xf8\xfc\x82\x46\x15\xef\xcb\xe9\x83\xb9\x59\xdf\xb6\x56\xcb\xa9\xf9\x4f\x58\x6b\x89\x08\xd7\x89\xf0\x6d\xa5\x2f\xd4\x79\x1a\xb5\xb7\xb3\xb9\x19\xf7\xca\x5e\x3e\x96\x3b\x6c\x9d\x6b\x87\x86\x7a\x36\xe7\x0e\xea\xfa\xb9\x57\xea\xe5\xb7\xd5\xb6\x5d\x5b\xe3\x04\xc9\x7a\x7f\x58\xff\x5e\xb7\x57\xe6\xf4\xe3\x2a\xf5\xbe\xa4\x3b\x8f\x21\xbd\x76\x2b\xb8\x5f\xf7\x9a\x7d\x9b\xde\xa0\x35\x79\x7e\xa4\xd5\x73\xfa\xcc\x9c\xde\xde\x1e\xa7\xfa\xdb\xda\x8b\x82\x0e\x04\x22\x20\xc1\x45\x40\x24\x8a\xf3\x38\x1a\xfa\xdf\x8b\x60\x55\xd1\x64\x37\x7a\xeb\xfd\xd9\x01\x15\x75\xe7\x31\xa5\x77\x95\x05\x10\xf5\xee\xe5\x4f\xef\x45\x7c\xac\x18\xe8\x77\x6c\x6b\xbc\x65\xdc\xc1\x1d\x29\xdf\xab\x13\xc1\x49\x72\x79\x7e\x56\xd7\x3b\xad\xe5\x71\x64\x97\xfc\x45\xf6\x3e\x08\x9e\x00\xe4\xcc\xe1\x6f\x33\xd4\xed\x7e\x67\x0a\x12\x0a\x88\xf4\x45\x73\xf1\x79\xe8\xbd\xd5\x44\x27\x01\x45\x42\x1d\x5e\x27\x87\x23\x52\x99\x5d\x72\x5f\xc2\xf5\x3a\x3b\x9d\x87\x34\x93\xec\x0f\xe2\xa2\x35\x53\xb0\xf3\x12\x8a\x69\x56\x57\xba\xc5\xbf\x70\x0a\x2a\x1f\x0d\x3a\xb1\x94\x36\xf4\x47\x96\xfb\x86\xf3\x3a\x99\xc9\x7f\x34\xea\x82\x95\x93\xed\x8f\xa8\xee\xf3\xd4\xeb\x40\xab\xfd\x30\xd6\x85\xb1\x6e\x53\x7f\xb0\xe9\xa4\xf8\x3a\x00\xb3\x8b\x9e\x2e\x50\x8d\x93\x7f\x51\xf5\xf1\x8e\xfc\xe2\xdc\x50\x36\x55\x3b\x55\x9d\xca\x10\x45\xa5\xc5\x7b\xe4\x4b\x50\x44\x9b\xbd\x3e\x80\x8a\x3b\x4e\x03\x77\xa1\x9e\x59\xb5\xd2\x0b\x48\x5d\xe7\x8c\x87\x66\xbd\xbb\xd0\x34\x9e\x28\x6e\x18\x08\x5f\x38\x8f\x57\x13\xd2\x9c\x8f\xfc\xf8\x79\xf1\xe3\x52\x35\xf6\xe2\x49\x58\x4b\xea\x60\x36\x1b\xa5\xef\x92\x64\x2d\xc4\xd3\x79\x0a\xaa\xc5\x40\xe7\x08\x76\x7d\x9d\x7e\x17\x1b\xfe\xfa\x2b\x0c\x94\xf0\x1c\x42\x95\x42\x1d\x97\xe2\x60\x34\xd2\xb8\xd3\x37\x37\xb7\xd0\x2c\x68\x0b\xa7\x9f\x20\x53\x2a\x44\xd9\x2c\xba\x16\xa1\xbb\xd5\xbd\xcc\x17\x44\xdb\x1d\x28\x88\x96\x5c\xb8\x81\x4f\xef\x8d\x85\x71\x38\x4f\x70\x0f\x3f\xfd\x94\xcb\x5e\xd3\xbf\xf9\xc0\x16\x7e\xe0\xa1\xc6\x38\x13\xf9\x3f\x02\x3e\x88\x6f\xfc\xca\x91\x22\x80\xf8\x3f\x4e\xf5\xe5\x62\x53\x65\x53\x07\xef\x3a\x04\x8b\x07\xaa\x6d\x53\x8e\x23\x7a\x89\xf5\xd7\x9c\xb6\xb6\x36\x99\xb4\xaa\xda\x64\xb2\x37\x96\x4c\xe8\xff\x01\x00\x00\xff\xff\x5d\xb2\x1f\x7d\x3f\x29\x00\x00")

func migrations1_initial_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1_initial_schemaSql,
		"migrations/1_initial_schema.sql",
	)
}

func migrations1_initial_schemaSql() (*asset, error) {
	bytes, err := migrations1_initial_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1_initial_schema.sql", size: 10559, mode: os.FileMode(420), modTime: time.Unix(1508345986, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations2_index_participants_by_toidSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\xb1\xca\xc2\x50\x0c\x46\xf7\x3c\x45\xc6\xff\x47\xfa\x04\x9d\xc4\x16\xe9\xd2\x4a\xb5\xe0\x76\x49\xdb\x8b\xcd\xe0\xcd\x25\x37\x20\x7d\x7b\x41\x07\x5b\xbb\xb8\x86\x8f\x73\x72\xb2\x0c\x77\x77\xbe\x29\x99\xc7\x2e\x02\x1c\xda\x72\x7f\x29\xb1\xaa\x8b\xf2\x8a\x93\x44\xd7\xcf\x6e\x12\x1e\xb1\xa9\x71\xe2\x64\xa2\xb3\x93\xe8\x95\x8c\x25\xb8\x48\x6a\x3c\x70\xa4\x60\x09\xbb\x73\x55\x1f\xb1\x37\xf5\x1e\xff\xb6\x5b\x1e\xff\xf3\x2f\xbc\xbd\xf1\xb6\xc6\x9b\x52\x48\x34\xfc\x28\x58\xae\x5f\x0a\x58\x26\x15\xf2\x08\x00\x45\xdb\x9c\xb6\x49\xf9\xea\xfe\xf9\x25\x87\x67\x00\x00\x00\xff\xff\x33\xec\x54\x7a\x15\x01\x00\x00")

func migrations2_index_participants_by_toidSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations2_index_participants_by_toidSql,
		"migrations/2_index_participants_by_toid.sql",
	)
}

func migrations2_index_participants_by_toidSql() (*asset, error) {
	bytes, err := migrations2_index_participants_by_toidSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/2_index_participants_by_toid.sql", size: 277, mode: os.FileMode(420), modTime: time.Unix(1508345986, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations3_use_sequence_in_history_accountsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x4d\x6b\xb3\x40\x14\x85\xf7\xf3\x2b\xce\x2e\xca\xfb\x66\x91\x6d\x5c\x4d\xc6\x1b\x22\x8c\x63\x3b\x5e\xdb\x64\x25\xa2\x43\x3a\x90\x6a\xeb\xd8\xaf\x7f\x5f\x48\xd3\x0f\x08\x6d\xa1\xcb\x73\x78\xe0\x39\xdc\x3b\x9f\xe3\xdf\xad\xdf\x8f\xcd\xe4\x50\xdd\x09\x65\x49\x32\xa1\xa4\xcb\x8a\x8c\x22\xdc\xf8\x30\x0d\xe3\x4b\xdd\xb4\xed\xf0\xd0\x4f\xa1\xf6\x5d\x1d\xdc\xbd\x00\x80\x92\xa5\x65\x5c\x67\xbc\xc1\xe2\x58\x64\x46\x59\xca\xc9\x30\x56\xbb\x53\x65\x0a\xe4\x99\xb9\x92\xba\xa2\x8f\x2c\xb7\x9f\x59\x49\xb5\x21\x2c\x12\x51\x92\x26\xc5\x08\x6e\x7a\x6c\x0e\xd1\xec\x1b\xef\xec\x3f\xa2\x13\x99\xcb\x6d\xe4\xbb\x18\x6b\x5b\xe4\x67\x33\xe3\x38\x11\x52\x33\x59\xb0\x5c\x69\x42\x61\xf4\xee\x0c\xc2\x1b\xa1\x0a\x5d\xe5\x06\xbe\x43\x49\x8c\x94\xd6\xb2\xd2\x8c\xde\x3d\xff\xbc\x64\xb9\x1c\xdd\xbe\x3d\x34\x21\xc4\x89\x10\x5f\xcf\x98\x0e\x4f\xfd\x1f\xec\xa9\x2d\x2e\xde\xf5\x89\x38\xa6\xdf\xde\x90\x88\xd7\x00\x00\x00\xff\xff\x55\xe2\xdd\x2c\xbf\x01\x00\x00")

func migrations3_use_sequence_in_history_accountsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations3_use_sequence_in_history_accountsSql,
		"migrations/3_use_sequence_in_history_accounts.sql",
	)
}

func migrations3_use_sequence_in_history_accountsSql() (*asset, error) {
	bytes, err := migrations3_use_sequence_in_history_accountsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/3_use_sequence_in_history_accounts.sql", size: 447, mode: os.FileMode(420), modTime: time.Unix(1508345986, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations4_add_protocol_versionSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcd\xb1\x0a\xc2\x30\x10\x06\xe0\x3d\x4f\xf1\xef\x52\x70\xef\x14\x4d\x9d\xce\x44\x4a\x32\x38\x15\xd1\xa3\x06\x6a\xae\x5c\x82\xe2\xdb\xbb\xba\x88\x4f\xf0\x75\x1d\x36\x8f\x3c\xeb\xa5\x31\xd2\x6a\x2c\xc5\x61\x44\xb4\x3b\x1a\x10\x3c\x9d\x71\xcf\xb5\x89\xbe\xa7\x85\x6f\x33\x6b\x85\x01\xac\x73\xd8\x07\x4a\x47\x8f\x55\xa5\xc9\x55\x96\xe9\xc9\x5a\xb3\x14\xe4\xd2\x78\x66\x85\x1b\x0e\x36\x51\xc4\x16\x3e\x44\xf8\x44\xd4\x1b\xf3\x6d\x39\x79\x95\xff\x9a\x1b\xc3\xe9\x97\xd5\x9b\x4f\x00\x00\x00\xff\xff\x83\xbb\x30\x2e\xbc\x00\x00\x00")

func migrations4_add_protocol_versionSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations4_add_protocol_versionSql,
		"migrations/4_add_protocol_version.sql",
	)
}

func migrations4_add_protocol_versionSql() (*asset, error) {
	bytes, err := migrations4_add_protocol_versionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/4_add_protocol_version.sql", size: 188, mode: os.FileMode(420), modTime: time.Unix(1508345986, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations5_create_trades_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x94\x51\x6f\xaa\x40\x10\x85\xdf\xf9\x15\x13\x9f\x30\x17\x93\x7b\x6f\x5a\x5f\x4c\x9a\x58\x25\xad\xa9\xc1\xd6\x4a\xd2\x37\xb2\xb0\x23\x6c\xa2\x2c\x99\x1d\xda\xf0\xef\x1b\x68\x69\x10\x57\xad\xaf\x9c\x39\x67\x38\xbb\x5f\x76\x34\x82\x3f\x7b\x95\x92\x60\x84\xb0\x70\x66\x6b\x7f\xba\xf1\x61\x33\xbd\x5f\xfa\x90\x29\xc3\x9a\xaa\x88\x49\x48\x34\xe0\x3a\x00\xf0\xf3\x51\x17\x48\x82\x95\xce\x23\x25\x21\x56\xa9\xca\x19\x82\xd5\x06\x82\x70\xb9\xf4\x9a\xc9\x81\x26\x89\x34\x00\x95\x33\xa6\x48\x1d\xb5\x91\xf5\x76\x8b\x64\x35\x37\xb2\xc1\xdd\xee\x84\x5e\xcb\x71\x59\x9d\x75\xeb\x9d\x8c\x84\x31\xc8\x11\x57\x05\x42\x92\x09\x12\x09\x23\xc1\xbb\xa0\x4a\xe5\xa9\x3b\xbe\x19\xf6\x22\x3b\x1e\x65\x4c\x89\x64\x71\xdd\x8e\xcf\xb8\x12\x2d\x6d\x9b\xfe\xfd\xb7\x7b\xf6\xba\xcc\xb9\xff\xff\x30\x7b\xf4\x67\x4f\xe0\x76\x47\xee\xe0\xef\xf0\xbb\x57\xac\xcb\x34\xe3\x6b\x9b\x1d\xb8\xae\xe8\x76\xe0\xfb\x75\xbb\xd6\x75\xb6\xdf\xe1\x50\xdd\xd0\x19\x4e\x9c\x96\xbf\x30\x58\xbc\x84\x3e\x2c\x82\xb9\xff\x06\x19\x93\x8c\x0a\x25\x61\x15\xf4\x91\x0c\x5f\x17\xc1\x03\xc4\x4c\x88\xe0\xda\xc8\xf4\x5a\x0a\x3b\xe1\x9d\xd4\xb8\x8a\x1a\x0c\x2f\x45\xb7\xac\xda\x52\xea\x90\xfa\xb6\x2e\x65\xf4\x90\xf4\xfa\xe4\x78\xc7\x00\x9e\x5a\xf7\x75\x78\x97\x16\x1e\xb1\xe2\x1d\x5f\xa8\x67\x63\xa3\x5e\xdb\x7d\x17\xe6\xfa\x23\x77\xe6\xeb\xd5\xb3\xfd\x5d\x48\x84\x49\x84\xc4\x89\xf3\x19\x00\x00\xff\xff\x79\x87\x24\x6b\x4c\x04\x00\x00")

func migrations5_create_trades_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations5_create_trades_tableSql,
		"migrations/5_create_trades_table.sql",
	)
}

func migrations5_create_trades_tableSql() (*asset, error) {
	bytes, err := migrations5_create_trades_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/5_create_trades_table.sql", size: 1100, mode: os.FileMode(420), modTime: time.Unix(1508345986, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations6_create_assets_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x3d\x4f\xc3\x30\x18\x84\x77\xff\x8a\x1b\x1d\x91\x0e\x20\xe8\x92\xc9\x34\x16\x58\x18\xa7\xb8\x31\xa2\x53\xe5\x26\x16\x78\x80\x54\xb6\x11\xca\xbf\x47\xaa\x28\xf9\x50\xe6\x7b\xf4\xbc\xef\xdd\x6a\x85\xab\x4f\xff\x1e\x6c\x72\x30\x27\xb2\xd1\x9c\xd5\x1c\x35\xbb\x97\x1c\x1f\x3e\xa6\x2e\xf4\x07\x1b\xa3\x4b\x11\x94\x00\x80\x6f\xb1\xe3\x5a\x30\x89\xad\x16\xcf\x4c\xef\xf1\xc4\xf7\xc8\xcf\xd9\x19\x3c\xa4\xfe\xe4\xf0\xca\xf4\xe6\x91\x69\xba\xbe\xcd\xa0\xaa\x1a\xca\x48\x39\x86\x9a\xae\x1d\xa0\xeb\x9b\x65\xc8\xc7\xf8\xed\xc2\x3f\x76\xb7\x9e\x63\x46\x89\x17\xc3\xe9\xa0\xcc\x47\x3f\xe4\x13\x4b\x46\xb2\x82\x5c\xfa\x09\x55\xf2\xb7\xbf\xf8\xd8\x5f\xee\x54\x6a\x5e\xd9\xec\x84\x7a\xc0\x31\x05\xe7\x40\x27\xb6\x82\x90\xf1\x74\x65\xf7\xf3\x45\x4a\x5d\x6d\x97\xa7\x6b\x6c\x6c\x6c\xeb\x8a\xdf\x00\x00\x00\xff\xff\xfb\x53\x3e\x81\x6e\x01\x00\x00")

func migrations6_create_assets_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations6_create_assets_tableSql,
		"migrations/6_create_assets_table.sql",
	)
}

func migrations6_create_assets_tableSql() (*asset, error) {
	bytes, err := migrations6_create_assets_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/6_create_assets_table.sql", size: 366, mode: os.FileMode(420), modTime: time.Unix(1508347260, 0)}
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
	"latest.sql": latestSql,
	"migrations/1_initial_schema.sql": migrations1_initial_schemaSql,
	"migrations/2_index_participants_by_toid.sql": migrations2_index_participants_by_toidSql,
	"migrations/3_use_sequence_in_history_accounts.sql": migrations3_use_sequence_in_history_accountsSql,
	"migrations/4_add_protocol_version.sql": migrations4_add_protocol_versionSql,
	"migrations/5_create_trades_table.sql": migrations5_create_trades_tableSql,
	"migrations/6_create_assets_table.sql": migrations6_create_assets_tableSql,
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
	"latest.sql": &bintree{latestSql, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"1_initial_schema.sql": &bintree{migrations1_initial_schemaSql, map[string]*bintree{}},
		"2_index_participants_by_toid.sql": &bintree{migrations2_index_participants_by_toidSql, map[string]*bintree{}},
		"3_use_sequence_in_history_accounts.sql": &bintree{migrations3_use_sequence_in_history_accountsSql, map[string]*bintree{}},
		"4_add_protocol_version.sql": &bintree{migrations4_add_protocol_versionSql, map[string]*bintree{}},
		"5_create_trades_table.sql": &bintree{migrations5_create_trades_tableSql, map[string]*bintree{}},
		"6_create_assets_table.sql": &bintree{migrations6_create_assets_tableSql, map[string]*bintree{}},
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

