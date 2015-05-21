// +build bin

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
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

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1d\x6b\x73\xdb\x36\xf2\x7b\x7e\x05\x4f\x4d\x6f\x9a\x4e\x25\x5d\x72\xed\xb5\x93\xda\x9e\x51\x2c\x39\xd1\xd4\x0f\x8d\x25\xbb\xd7\x4f\x1d\x4a\x84\x24\xc4\x14\xc9\x92\x90\x1c\x57\xe3\xff\x7e\xbb\x78\xf0\x25\x40\x22\xa9\x87\x9d\x5c\x32\x93\x31\x09\x60\xb1\xd8\xf7\x62\x49\x50\x47\xff\x68\x5f\x9d\x0e\xfe\xe8\x75\xac\x29\x9b\xb9\x2f\x4e\x8e\xf8\x1f\xcb\x82\x0b\x62\x3b\x70\x81\x97\x33\xc2\x6c\x6b\x34\xb5\xc3\x88\xb0\xe3\xda\x9c\x8d\xeb\xbf\xd4\xd2\x5d\x53\xc6\x82\x3a\xf9\x6b\x4e\x17\xc7\xb5\xff\xd6\x6f\x5a\xf5\x53\x7f\x16\xd8\x8c\x0e\x5d\x52\xb3\x46\xbe\xc7\x88\x07\x70\xdd\xce\x31\x71\x26\x24\x03\xe9\xd9\x33\x72\x5c\x5b\x50\x72\x1f\xf8\x21\x4b\x0d\xbe\xa7\x0e\x9b\x1e\x3b\x64\x41\x47\xa4\xce\x6f\x7e\xb0\xa8\x47\x19\xb5\xdd\x7a\x34\xb2\x5d\x72\xfc\x5a\x4d\xc4\x28\x73\x09\xbf\x86\xbb\xe5\xb2\xd1\xb6\x99\xdd\xf8\xe0\x47\x0c\x27\x7f\x7c\xb4\xe0\x0a\xa6\x3c\x6a\x26\xe3\x4e\x8e\x5c\xea\xdd\x59\x21\x71\x8f\x6b\x14\x50\xd6\x2c\xf6\x10\xc0\x3a\xe8\xcc\x9e\x90\x66\xe0\x4d\x6a\xd6\x34\x24\xe3\xe3\x5a\x73\x6c\x2f\x70\x40\x03\xdb\x56\x40\x23\xf6\xe0\x92\x68\x4a\x08\x53\x13\x30\xf2\x89\x35\x47\x51\x14\xc3\xc3\x75\x93\x7a\x0e\xf9\xd4\xc0\x56\x39\x43\x34\x0a\x69\xc0\xd2\x20\x1f\xed\x85\x2d\x5a\x6b\x79\x3e\x5b\x51\x38\x82\x89\x3e\x46\xcd\x10\x39\x1c\x12\xb8\x7a\xd3\x78\xdd\x78\xfd\xb3\x6a\x68\xcc\xa8\xd7\xf8\x08\x38\x1d\x20\xbc\x3e\xb3\xa9\x27\xc6\x2f\x97\x74\x6c\x35\x06\xad\xf7\xef\x3b\xed\x21\xf5\x1e\x1f\x61\x9c\x5c\x8c\x80\x58\x2e\x89\x1b\x01\x87\x00\x43\x73\x46\xdd\x3b\xd9\xc9\x3b\x3c\xe7\xf1\xb1\xa6\x98\x7a\xd4\x14\x8b\x93\xeb\x6f\x4a\xd5\x38\x39\x1a\xfa\xce\x83\x6c\xf4\xec\x85\x35\x72\xed\x28\x3a\xae\xc1\xe5\xd0\x0e\x2d\xf1\xa7\xee\x90\xb1\x3d\x77\x99\xba\x1d\xd3\x4f\xc4\xa9\x33\x3f\xa8\x59\xa1\x0f\x62\xc4\xd1\x74\x02\xca\x02\x62\x88\xf1\x39\x34\x9e\x0c\x35\x02\x48\x22\x00\xe9\xce\xa9\xa3\xc6\xe4\x46\xc9\xc9\x71\x61\x24\x4c\xc6\xe8\x47\x0d\x43\xdb\x73\xd2\x83\x70\x98\x6d\x51\xe7\xb8\x36\x95\x6a\x13\x0b\x50\x72\x35\x70\xed\x11\x99\x71\xcd\x1c\xfa\x8c\xf9\x33\xd9\xce\xfc\xc9\x04\x89\x08\xfc\xc0\x5f\x00\x66\xd9\x1a\xd2\xc9\x84\x84\x38\x1d\x34\x5a\x63\x7f\x34\x57\xe2\x89\xc9\x39\xae\x7d\xe3\xd0\x08\x46\x0e\xeb\x59\xe0\xd8\x04\xc2\xb9\xe7\x51\x6f\x62\x29\x95\x6e\x77\xfb\x83\xeb\xee\xbb\x44\x2e\xf1\xe2\x57\x95\xfe\xa8\x69\xe7\xe8\x6b\x02\x1f\x32\x7c\x89\x02\xdb\xe3\x34\xe7\x97\x91\x07\xc4\x81\x2f\x8c\x53\x09\x06\xe3\x3c\x92\xbb\x23\xdf\x75\xed\x20\x42\xcb\x57\x12\x14\x0d\x56\x7e\xc0\x8a\x98\x70\x96\x79\xc0\xe8\x8c\xd4\x03\x3b\x04\x26\xac\xcc\x91\x5b\x9c\x18\xfc\x4f\x6f\x18\x05\xbf\x26\x04\x89\x56\x23\x97\x6e\x78\x37\xf2\x28\x47\x99\x96\x4d\x73\x37\xa5\x3d\x8a\x04\xf8\x93\x67\x93\x4b\x39\x6e\x1a\xd4\x72\xda\x86\x06\x6e\x5c\x4b\xb7\x87\xeb\x70\xa9\x7e\x32\xd7\x2e\x35\xd9\x79\x4b\x3b\x59\x73\xee\x16\xa2\x48\x5d\x82\xf2\x4e\x99\x86\xbc\x02\x2b\x49\xcd\xed\x82\x56\xd5\xa9\x07\xde\x92\x18\xb9\xa6\xa6\x5e\x69\x14\x06\x29\x4c\xf0\x1b\xf4\x15\x9a\x11\x30\x06\xbb\xf2\x9a\x1e\x53\xad\x99\x76\x33\xae\x20\x32\xa1\x0a\xa2\xdd\x62\x5a\x4c\x4c\x98\x16\xf6\x04\x1c\x14\x2b\x87\x2e\x2b\x64\xfd\xc8\xb4\x14\xed\x11\xa3\x8b\x55\x1b\x89\xd7\x87\x21\x3d\x7a\xdb\x6c\xde\xdf\xdf\x37\xc0\xab\x84\xf0\xbf\x31\xf2\x67\x4d\x11\x4c\x21\xf4\xb8\xc4\x8e\x48\xd4\x74\x6d\x46\x22\x8c\x7f\x18\x5c\x21\x92\x93\x30\x02\x4f\x1e\xfb\xac\xdb\xce\x75\xbf\x7b\x75\xb9\xea\xb3\x00\xd7\x55\x7f\xd0\xb9\x1c\xac\x92\xa9\x5b\x78\x5e\x87\x73\x2e\x28\x75\x0b\x37\xa0\x67\x32\x28\x29\xa7\xc2\xc3\x8d\x3e\xa0\x58\x81\xed\x38\xe0\x66\x31\x24\x45\x64\x01\xc4\x3d\x44\x53\x6d\x24\x0a\xfd\x7b\x43\xf4\x19\xba\xfe\xe8\x0e\x72\x17\xb7\x3e\x73\xea\x6f\xd4\x85\x3f\x1e\x43\x1c\xaf\xbf\x36\x86\x23\xf0\x3e\xc4\x7d\xbd\x62\x69\xf6\x90\xb8\x79\xbf\x57\x17\xad\x1c\x82\x47\x39\x0c\x0c\x43\xe6\xd5\x39\xea\x1a\x84\x18\x08\x28\xce\x7c\x36\x7b\x98\x91\x19\x50\x38\xa6\x93\xd1\x94\x8c\xee\x86\xfe\x27\xde\x59\x8f\x30\x61\xf9\x26\xee\xd4\x09\x5e\x11\x03\xb3\x8a\x38\x2a\x26\xe7\x69\x04\xf9\xcb\xaa\x8d\x6d\x17\xbd\xfa\x77\x1f\x23\x10\xb0\x10\xef\xa9\x4b\x81\x61\x8d\x0f\xd4\x21\x62\xde\x8b\xce\xc5\x2b\x48\xb9\x84\x76\xa9\x34\x42\xa9\x3c\x60\xd7\xa8\xc1\x05\x99\xf9\xe1\x83\x4e\x0d\xa8\x17\xcc\x99\x08\x51\x06\xca\xf2\x5c\x4a\x7a\x44\x72\x95\xdc\xdb\x21\xb5\xeb\x53\xea\x38\x04\xb2\x23\x16\xce\x31\xc2\xf3\x3b\x8d\xf2\x21\xab\x5f\xe8\x96\xb2\x6b\x74\x5c\x1b\x56\x16\xa0\x94\x36\x11\x96\x42\x1c\x8b\x42\xcc\x58\x40\x12\xf1\x52\x05\xf6\x5c\x62\x97\xc6\x0a\x4a\x32\x53\x78\xa6\x7e\x48\xff\x46\x33\x71\xeb\xbc\x79\xe8\x87\x5c\x1d\x78\xe6\xc3\x9b\x74\xbe\x2b\xa3\xdb\x38\xa8\x3e\x09\xfd\x79\x50\x47\x03\x23\x8e\xc1\xdb\x65\xcc\x08\x34\x8f\x83\x58\xf1\x55\x3d\xca\x67\x5a\xc3\x39\x2c\xc2\x33\xb8\xe9\x15\xfb\x81\x89\xf8\x64\x2a\x09\xcd\xda\x53\xce\x34\x4c\x73\x26\x0a\x90\x13\xb3\x79\xfc\x07\xce\x6e\x9d\x2e\x15\x5d\x2a\x97\x35\xc4\xce\x0d\x66\xd7\xff\xbd\xd5\x33\x5b\x1c\x48\x9d\x44\xf7\xb6\x21\x7e\x56\xa5\x8c\x58\x38\xe7\x26\xe2\x9a\xab\xba\xbd\xa1\xab\x92\x02\x2d\x97\xf7\x94\x4d\xb3\x8c\xb9\x26\xe3\x10\x36\x66\x60\x02\x8f\x8f\xcb\x25\x23\xb3\x00\xc3\x94\x55\x03\xd6\x82\xe3\x8f\xde\xbe\x0d\xc5\x80\x9a\xd5\xc0\x01\x9c\x67\xe6\x45\x35\x71\x29\xba\x90\xab\x33\xde\x95\xc6\x8d\x36\x5e\xd2\xba\xcb\xda\x75\x9c\x15\x03\xaa\x3a\xb3\x71\x4b\xae\xa3\x52\xcb\x26\x71\xf1\x27\x40\x0a\x40\x64\x97\x26\x6b\x2e\xc7\x8e\x95\x0c\x5b\xb3\x9d\xd0\xc5\xd5\x7f\x1f\x36\x8e\xd2\xb1\x39\x8c\xaa\xbe\x7d\x45\xd1\xee\x99\xd9\xa4\xe9\x58\x23\x65\x95\x6e\xc9\x99\x06\xf6\xb0\x7b\xd6\x18\x60\x4a\xb6\xba\x15\xd4\x44\xd6\x55\x5a\xbf\xe8\xc0\x1a\x8b\xaf\x4a\x5c\xe5\xb2\xd9\x49\x58\xfd\x5c\x63\x28\xd0\x4c\x3d\xc3\xee\x25\x2d\xd4\x43\x07\xcb\xf5\xd6\xd5\xf9\x04\x46\xef\xa0\x1b\x13\xd6\xe5\xd0\x08\x6f\x9c\x22\xf6\x55\x99\xb2\x9c\x5d\x8a\x35\xe0\xce\xb9\x7b\xc6\xb7\xea\xcf\x3c\x80\xe2\x2a\xf7\x17\x3f\x57\xea\x10\x18\x9e\xb4\x4a\x99\xde\xb8\xd2\x71\x3d\x82\x05\x8f\xa6\x4a\xa7\xed\x21\xda\x74\x5c\x50\x39\x1b\xb4\xde\xf5\x1b\xf4\xac\xd7\x3a\xfd\xad\x33\xe8\x37\x6e\xa8\xc7\x74\xa6\x29\xe6\xb5\x13\xb1\x07\xf6\xe8\x8e\x30\xb3\x51\xf4\x44\xbf\x7e\x5b\x6e\xda\x98\x57\x5b\x7b\xe7\xfa\xfa\xea\xba\xc4\xd2\x49\x18\xfa\xa1\x79\xe5\x1d\xde\xbd\x93\x85\xcb\xa0\xb4\x61\xfd\xef\xfe\x18\x74\x4a\x2c\x7f\xf8\xc0\x88\x79\xf5\xef\xb0\xb7\xfc\xe2\x57\x0b\x22\xbb\xca\xd9\x12\x55\x49\xb1\x07\x58\xa1\x32\xf6\xb4\xf3\xd1\xaa\xe3\x2b\x9d\x7f\x12\x31\x9b\x77\x83\x77\x32\x04\x18\xc9\x76\x51\x6b\x31\xaa\xbc\x0c\xfc\x5b\x1a\x48\x2a\x75\x8c\x29\xae\x98\x40\xc6\xf0\x7f\x56\xcd\x23\x13\xde\x4b\x5d\x2f\xc1\xfa\x94\x35\xed\x8b\xf3\x12\x45\x31\xc6\x6f\xb0\xee\x0c\xdf\x05\xb5\x95\xd9\x2e\xc0\x77\xc0\x75\x61\xa2\x25\x98\x9e\xb8\x80\x7d\xf1\x5c\x60\x28\xc6\xf2\xf5\x0e\x29\xc3\x71\x4e\x69\x65\x86\x73\xe8\xca\xfc\xde\xe1\x6e\xe9\xcd\x61\x77\x4b\xa3\x60\x6e\xde\x2e\xc5\x9d\xfb\xda\x2f\x9d\xf6\x6e\xcc\x1b\x26\xc0\xae\x11\x39\x80\x14\xd8\x18\x69\xc8\xfa\xa2\x77\x46\x89\xa4\xaa\x6c\x8d\x84\x18\xbe\x96\x1c\x8b\x96\x1c\xb5\x9a\x99\x57\x80\xe7\xba\x8b\x92\x36\x67\xda\x46\xed\x96\x36\xe3\x3e\x0a\x56\xf1\x39\x6c\xa4\xf8\x32\x9f\x77\x25\x12\x05\x56\xce\xe8\xcb\x9a\x7b\x1c\x64\x01\x55\xc5\x00\x0b\x90\xcf\x20\xb6\x1e\xb8\x12\xe9\xac\xa9\x44\x3a\x7b\xae\x44\xb6\xd7\x54\x22\x9d\x62\x95\xc8\x76\x99\x4a\xe4\x2a\xad\x5f\x74\xbc\x75\xb6\xa9\x44\xb6\xbf\x56\x22\x3f\xf3\x4a\x64\x7b\x6d\x25\x52\x67\x5f\x95\x29\x33\x46\xd0\xf6\x67\x51\x89\x6c\x3f\xc3\x4a\xa4\xb3\xbe\x9a\xd7\x16\x9b\x4f\xe7\xac\x7b\x79\xd5\x2e\x5e\x0e\x73\x80\x30\xdf\x59\x53\x0f\xeb\xf2\xee\x6d\xaa\x79\x4e\xa1\x6a\x5e\xbc\xfe\x52\xd5\x3c\xe7\x73\xab\xe6\xc5\xfc\x4e\x71\xc7\x50\xdc\xd0\x49\xd4\x54\xdd\x68\x57\xab\x6e\xe4\x51\xac\x2d\x6f\x14\xd5\xb0\xf4\xbb\x9e\x92\xda\x8a\xf9\x97\x02\xdf\xba\xa2\xe4\xe4\x2a\x4a\x05\x98\xbe\xb9\xa2\xb4\x2d\xcf\x0b\x54\x94\x0a\x1a\x45\x86\xe3\xdb\x54\x94\x9c\xc3\x57\x94\xb2\xef\xd2\xa9\x9b\x22\xef\x64\x57\x79\x5b\xee\x97\xdc\xdb\x72\x07\xae\x5b\x05\x91\x39\xb7\x56\x7d\xfb\xca\xad\x7b\x7d\x73\x6e\xad\x7b\x0f\xf5\xa4\x17\xfa\x23\x12\x45\x3a\xff\xb9\x92\x48\xaf\x12\xf6\x45\x27\xd2\xb1\xac\xaa\x24\xd2\x5c\x10\xd5\x12\x69\xf1\x82\xb3\xa9\x64\x65\x61\x82\xa5\x7b\x9f\x5a\xa3\xde\x3b\x48\x92\x7a\xfd\x5d\x24\x49\xcf\x3f\xed\x37\x57\xd4\x4c\xaf\x6f\x8b\x49\xf7\xbf\x19\x70\xc1\x3c\x2d\xd3\x8e\xc0\xa8\x89\xbd\xbe\xe7\xb3\x36\x19\x85\xc4\xe6\x3b\x81\x75\x9b\x82\x20\x42\x1c\x3b\xa4\xb2\x5e\x8a\xc0\x99\x1f\x92\x8a\x04\x26\xdb\x9e\xf5\xf4\x21\x8a\xfd\x6d\x7c\x7a\xfd\xc0\x9d\x47\x03\xb0\xcb\xfd\x6c\x7b\xf6\x59\xa2\x0b\xa2\x92\xee\xad\xa4\x63\x8b\x73\x96\xa0\x6a\xba\x12\x44\x5f\x13\x95\xfd\x24\x2a\x8b\x89\x39\x51\x51\x7d\xfb\x4a\x54\x6e\xdf\x9b\x13\x15\xdd\x31\x96\x93\x5b\x71\x82\xc5\x9a\xb8\xfe\x10\xcf\x6b\x32\x9b\xcd\x8b\x24\x2d\xab\x44\x7e\xd1\x49\x4b\x2c\xb7\x2a\x49\x0b\x17\xca\x17\x90\xb4\xdc\xbe\xff\x7f\x4f\x5a\x4c\x27\xc1\xc4\xa4\x7b\x49\x5a\x9e\x5b\x64\x5b\x94\xb5\x81\x92\xda\x1f\x47\xb6\xc5\xa4\x62\x64\x5b\x4c\x9e\x41\x64\x5b\x7b\x70\x5c\x01\x2c\xec\xd0\x42\x76\x59\xc7\xea\x14\xde\xe3\xe3\xaf\xf9\x13\xdc\xf2\xe0\x36\x9e\xe5\xe6\xc7\xfe\x97\xcb\xe6\xf7\x2f\xbe\x6f\xa2\xa5\x09\x82\x35\xcf\xde\x5c\x1a\xd4\x70\xc0\xcb\x53\xeb\xed\xb1\xe5\x87\x56\xe3\xf4\xbc\xd5\xef\x5f\xb6\x2e\x3a\x56\x8d\x4b\x0f\xba\xa5\xbe\x02\xe6\x6e\xbb\xc5\x58\x08\x81\xc3\x42\x90\xc7\xc7\x2d\xdd\xea\x89\x52\x0c\x40\x2f\xb0\x5d\xe2\x39\x6d\x94\x06\x47\x89\x03\x84\x93\x10\x66\x08\x2b\x38\xf3\x43\xb9\x84\x78\x05\x48\x03\x76\x5d\xfc\xde\xc7\xc3\xfb\xa2\xfb\x85\x48\x13\xb3\x49\xa1\x9a\xae\x08\x67\x82\x39\x24\x11\x11\x67\x8e\xf0\x73\x57\xb7\x9d\xeb\xeb\x6e\xbb\x83\x2d\xc2\x7d\xf1\x83\xf6\x1b\x78\xb7\x5c\x42\xd4\x9c\x10\xeb\x25\xfd\xc1\x7a\x39\xc2\x4c\x1b\xc6\x4a\xd5\xef\xdd\x34\xce\x69\x84\xeb\x63\xe1\x72\x79\x47\x1e\x04\x1f\x82\x79\x1d\x50\x0f\x1f\xea\x97\x35\x01\xd2\xb8\x14\xf4\x1c\x31\x47\x59\x54\xe2\xcf\x2d\xcf\xbf\x0f\xd5\x69\x20\x20\x30\x86\x38\x6a\x32\xc7\x08\x55\x7b\x91\x32\x22\xc9\x49\x05\x7b\x13\x91\xf0\x14\x01\x38\xd6\xcc\xb4\xd8\x95\xb1\x95\x93\xca\x58\xfa\x0f\x91\x44\x62\xe5\xb1\x40\xd7\x8e\x90\x74\x1d\x97\x18\x48\xc1\x2e\x03\x96\x26\x0b\x53\xaa\x92\xfd\x53\x44\x71\x84\x57\x41\xb1\xe2\x45\xbc\x62\xbc\x79\x6d\xe1\xb9\xfc\x80\x24\x1e\x0b\x88\x4a\x7d\x95\x03\x30\x2b\x1f\xc1\xa6\x69\xef\x11\xdf\x61\xc7\x06\x2d\xe0\xc3\x50\x54\x82\x27\x72\xf0\xdc\xa3\x99\xe0\x7f\xf2\x6d\xce\xe7\x55\xc1\x02\xa2\xda\x3f\x12\x14\xd5\x36\x58\x9a\x92\xab\xbc\x55\x7d\xe6\x82\x25\xdf\xb9\xd8\xf0\x76\x02\xf7\x03\x22\x46\xb0\xc4\xc7\x72\x01\x96\xf1\x28\xaa\xec\x5b\xd8\xad\x9c\x29\xb7\xf2\x61\x70\x71\x7e\x76\x75\x6d\x61\x1a\xb4\xd9\x5b\xc7\x1e\x07\xb6\xe3\x77\x89\xb7\x69\x9f\x71\xf4\x3a\x8f\xa3\xca\xd9\xc2\xeb\x38\x34\x14\x5f\xcb\xe0\x13\x34\xda\x34\xbc\xe4\x9f\x9f\xc8\x19\x60\xce\xed\xac\xe1\x20\xf7\xcf\xdf\xe1\x1f\x0b\x68\x02\x84\xaf\xdf\x20\xce\xba\x40\xea\x90\x45\x06\x1d\x59\x20\xba\x57\x5a\xff\xb5\x05\xce\xff\xa4\x51\x6a\x29\xd4\xa3\x5c\xf1\x33\xc0\x7d\x0e\xd5\x5a\xd8\xd4\x2d\x07\x02\x26\x09\xaa\xa2\xbe\x6c\x31\x0f\x12\x97\x64\xa6\x64\x1e\x91\x80\x84\x23\xfe\xdd\x8c\xef\x92\x1b\x24\x29\x9e\xb4\x27\xda\x38\x01\x6a\xda\x42\x3e\x53\x2d\x6c\xe0\x33\x3b\x4d\xcb\x56\x0e\x30\xfb\x78\x63\x73\x72\xb1\x8d\x93\x94\x46\xf4\xf2\x3c\xd1\xf3\x73\xea\xdd\x61\xf8\x48\x5c\x4b\xfe\x03\x32\x27\x89\xe3\xb0\xd4\x43\x4e\x58\xe4\x39\x98\xc8\x59\xbf\xf1\x01\x6e\x31\xc1\xb1\x52\xb1\x44\x76\xa9\x78\x65\x59\x69\xdf\xd3\xe6\xdf\x33\x12\xfe\x49\x03\x63\x87\x84\x69\x01\x39\xa3\x57\xbf\xc3\x62\x6f\x74\x95\x25\xe8\xb9\xe8\x19\xe9\x81\x2e\x3d\x3d\x17\xfe\xdc\x63\xc4\xc9\xc7\x53\x05\xb3\x2f\x7a\xb4\x5b\xd4\x93\x1c\x3d\xad\xdb\x56\xf7\x5c\x91\x94\x5d\x9b\xe8\xca\x06\x7b\x3e\x05\x37\x54\x29\x1e\x3d\x4c\x8a\xa4\x03\xd3\x73\xd3\xef\xb4\x63\x09\x65\xd7\xc6\xbb\x72\x19\x12\x9f\x02\xbd\x88\x22\x47\x0b\x93\x96\xd0\x81\xe9\x19\x5c\x0d\x5a\x06\xf9\x88\x2e\x8d\x7c\xb8\xf3\xd1\xcb\x47\xc2\xec\x42\x3e\x71\xb5\x61\x8b\x5c\x20\x13\xc1\x77\x91\x10\xc8\x07\xef\x4f\x96\x11\x08\xfc\xfa\x94\x40\xbe\x53\x70\xe0\x9c\x40\x62\x3d\x68\x52\x60\xc0\x59\x29\x2b\xe8\x8e\x43\x42\x4a\x82\xcc\xf7\x91\x16\xe0\xac\x5b\xe7\x05\xe2\xb5\xa4\xdd\x25\x06\x99\x17\x4d\xb6\xdf\x1f\x19\x03\xa2\x0c\xc9\x05\xc3\xa7\x0a\x78\xe5\xbd\x9f\x08\x2d\xe5\xe1\xb8\x0b\x2f\x0f\x26\x3c\xe5\xae\x76\x36\x59\xf7\xb3\xbd\x3b\xa3\x87\xde\xdf\xd0\x71\xe2\xcb\xba\xe6\xdd\x0d\xcd\xec\x6e\xa4\x89\xd3\x71\x23\xe3\xc1\xaa\xba\x2c\xaa\xf6\x14\xd9\x79\x8b\x7b\x0d\x00\x68\x13\x97\xd9\x5d\xaf\x34\xc8\xd5\x9c\x95\x81\x29\x87\x21\x3b\xf9\x56\x76\x4f\x73\x1b\x82\xdd\x95\x45\xba\x60\xb7\xe1\xd8\x36\x9a\xfa\x6a\x51\x41\x7d\xd3\xec\x5d\x77\xd0\xb7\xc0\x69\x5a\x11\x3e\x7a\x49\x7f\x81\xb2\xeb\xad\x2f\x35\x1c\xe5\x3e\x63\x36\x3c\x6a\xa6\x5b\x4e\xf0\x9b\x72\x15\x2b\x1e\x9b\x17\x07\x52\x79\xf2\xd5\x31\xf4\x42\x16\x7f\x57\xce\x9a\xf9\xce\xdc\xf5\xad\x1f\xdf\x6f\xc1\xc0\x77\xb9\x25\x7e\xfb\xe3\xfb\xbd\xaf\xb1\x34\x1f\xcb\x2e\x72\x1b\xc7\x4c\x77\x9b\x66\xc6\x27\x86\x9f\xc8\x31\x0b\xec\x7a\xcf\x2c\x8f\x42\xef\xcb\x35\xcb\xe9\xbf\xfa\xe6\xcd\xca\xf1\x9c\x9c\xb3\xc1\x2f\xa7\x3f\x09\x6b\xb2\xdb\x1d\xf8\x37\x93\xe3\x3d\x10\x7a\xbd\x4b\x2d\x82\x7c\x17\x8e\xd3\xe0\x2c\x77\x81\x7e\x3b\x97\x98\xf6\x61\xbb\xf0\x89\xea\xe3\x15\x4f\xe4\x14\x25\x7a\xbd\x57\x54\x1f\xe6\xd8\x97\x5b\x54\xf3\x7f\xf5\x8b\x05\x14\xe4\xab\x63\x7c\x26\x8e\x51\xe4\x74\x4f\xeb\x1e\xcd\x6b\x78\x16\x4e\x32\xe3\xd4\xb6\xf7\x92\x33\x32\xdb\xfd\x3b\x10\x30\x69\xe2\x08\x2f\x3a\x17\x3a\x17\x88\x9f\xb5\x14\xce\xef\x8e\x82\xb6\x71\xa0\xc6\x6f\x14\x97\x9d\xf3\x7e\xa9\x9e\x82\x7e\x02\x01\xce\xca\x14\xe9\x10\x60\xd7\x4f\xee\xe4\x9c\xdb\x14\xe8\x70\x8a\x9d\x3e\xb7\x8b\xbf\x08\x7a\xd0\x17\x17\x52\x26\x84\x52\x29\x33\x7e\x4d\x15\xed\x00\xe5\x33\x65\x1c\xdb\x9b\x59\x10\x1d\x2c\x09\x09\x42\x7f\x94\x58\x5f\xaf\xcf\xd7\xaa\xb3\xc0\x40\x65\x1f\x01\x45\xfb\x43\xb8\x46\xaf\xdb\x46\x1d\xf0\x26\x6b\xde\x46\x92\xfa\x89\xcf\x71\x14\x4c\x31\x43\x8b\x21\x6e\x72\x10\xb2\x3b\xf5\xaf\x7c\xce\x13\x44\x60\x7f\xa1\x96\x16\x79\x8f\x6f\xcb\x14\x4c\x80\x12\xda\x42\xea\x87\x94\x3d\x94\x04\xbb\xa4\xa3\x82\xce\x27\x06\xe9\xd3\xbf\xcb\x82\x5c\x93\x88\x3a\xc4\x5b\x93\x6d\xa1\xe7\x51\x75\xf0\x13\x05\x36\x90\xbf\xa2\xb1\xfe\x29\x4b\x05\x31\xfc\xc8\xc5\x80\x19\xe7\x3a\x31\xe4\xf3\xd0\xad\x1c\x9a\x3a\xb8\x50\xe5\x15\x84\x37\xfb\x78\x05\x41\x5f\xd5\xcf\x3e\xb3\xef\xf5\x81\x27\xda\x47\xc2\xa2\x47\xf7\x44\x18\xda\xe3\x07\xc2\x3a\x90\x43\x3e\x10\xce\x53\x73\x63\xa4\xe6\xc6\x40\xcd\x8d\x99\x9a\x9b\xfd\x51\xa3\x7f\x9d\x62\x85\x9a\x7e\xe7\x5a\xf7\x78\x5b\xf6\x68\x9e\x6e\x63\xbb\xf6\x69\xbd\x02\x39\xe0\xcb\x07\x2b\x9a\x76\xdd\x35\x69\x1a\xf4\x68\x35\x2d\xa1\x45\x0b\xf2\x94\x9a\x76\xd9\x3d\xed\xe8\x65\xc3\x7b\x34\xb2\xb9\xec\x1a\xa9\x11\x20\x4f\x28\x9b\xdb\xee\xf5\x40\x4f\x0d\xef\xd1\x50\x83\xed\x06\x4d\x13\x20\x4f\x48\xcd\x75\xa7\x6f\xd0\x34\xec\xd1\x69\x1a\xb4\x9b\xbc\x00\x07\xd9\xa7\xa6\xa5\x23\xa3\x18\x97\xa5\x66\xd0\xbd\x30\x68\x1a\xef\xd1\xbd\xe3\x02\xed\x06\xd9\x08\x90\xbd\xc8\xa6\x98\x4f\xc3\x58\x68\xb0\x9b\x96\x9e\x9a\xd3\xab\x8b\x8b\xd6\x65\x7b\xf5\x0d\x31\x05\xf2\x5c\x5e\xd8\x91\x99\xf5\xf6\x29\x7a\x6a\x2f\x07\x53\x49\xba\x05\x6b\x32\xc7\x0d\x6e\x6d\x77\x0e\xf9\x93\x7a\x1d\xb9\x04\x86\xc5\x84\x57\x38\x79\x36\x12\x5a\x90\x86\xe3\xb9\x16\xd1\x14\x67\x62\xbe\x8b\xb3\x1e\xd7\x7e\x8a\x93\x36\xf9\x83\x4c\xe2\x68\x1e\xff\x24\xb4\x31\x79\x2a\xb0\x82\x78\x1b\xb2\xf5\x3e\x63\x66\x8f\xa6\x49\x3e\x24\xd7\x77\x01\x8d\x88\x51\xb7\xdf\x90\xbf\x8f\x25\x37\x1d\xf3\x39\x4f\x11\x71\x96\xc6\x0d\xdf\x10\xe8\x36\x03\xe5\x12\xd0\x9f\x13\x24\xf9\xe9\x5f\x99\x37\x1c\xb8\xd1\xc6\x61\x97\x76\x2e\x2f\xd6\x8e\xea\x85\xfe\x02\xb2\xee\x70\xdd\xc8\x32\x6b\xfe\x25\x59\x33\x9e\xb7\x24\xe9\x45\xcb\xcb\x3e\xb6\xaf\x25\xa0\x0c\xc2\x9f\xfe\x95\x60\x1c\x53\x97\x04\x36\x9b\x6a\x90\x4a\x81\xe2\x88\x3f\x71\xc8\xc6\x94\xbd\x80\xf2\xc9\xcf\x2e\x15\xd8\x05\x1b\xcf\x92\xa5\xed\xc0\x79\x65\x38\x82\x98\xb3\x37\x2b\x33\xf5\xba\xf1\xb1\x2b\xa9\xba\x1f\x51\xc7\xcd\xe2\x0d\x87\x38\xbc\x26\xcf\xa0\xed\xa0\xd8\x42\xf3\xe5\x90\xa4\x0b\x37\x5e\xc6\xce\x40\xaa\xad\x71\x00\x57\x3e\x63\xaf\x43\x43\x32\x62\xfc\xf7\xbe\x76\x54\x5b\x89\xb5\x61\x0b\xcf\xad\x8e\x9a\xc2\x04\xa9\x43\xa1\xfc\x58\x99\x3c\x16\x9a\xba\xae\x47\x33\x4b\x02\x88\x06\xc1\xf3\x74\x29\x37\x3d\xda\x76\x1c\xf5\xcb\xa6\x27\xf2\xc8\x6b\x1c\xd9\x72\xc7\x8e\xf9\xc9\x55\x3c\xb2\x1e\xfa\x6e\x8c\x82\x8f\x48\xff\xbe\x6c\xcd\xe2\x3f\x46\x3a\xf5\x5d\x07\x7f\x45\x14\x1d\xbb\x38\x62\x2a\xf5\x3d\xf5\xfd\x24\xd1\x1e\x7f\x2e\xa9\x47\x42\xea\xa3\xb2\xd7\xac\x05\x06\x1f\x0e\x2c\x1a\x63\xdd\x14\xc4\xf0\x33\x87\x8a\x73\xff\x0b\x00\x00\xff\xff\xee\x88\x3f\xad\x23\x78\x00\x00")

func index_html_bytes() ([]byte, error) {
	return bindata_read(
		_index_html,
		"index.html",
	)
}

func index_html() (*asset, error) {
	bytes, err := index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "index.html", size: 30755, mode: os.FileMode(384), modTime: time.Unix(1400000000, 0)}
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
	"index.html": index_html,
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

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"index.html": &_bintree_t{index_html, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

