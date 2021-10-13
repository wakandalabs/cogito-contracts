// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../contracts/Cogito.cdc (3.701kB)
// ../../../contracts/NonFungibleToken.cdc (4.832kB)

package assets

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
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

var _cogitoCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x57\x5b\x6f\x22\x37\x14\x7e\xe7\x57\x9c\xee\xc3\x76\x46\xdd\xc0\x56\xaa\xfa\x80\x92\x26\xab\xb0\x48\x3c\x14\xad\x12\xaa\x3e\xac\x56\x5a\x33\x3e\x80\x95\x19\x1b\xd9\x1e\x28\x45\xfc\xf7\xea\xd8\x1e\xcf\x95\x64\x8b\x22\x31\x8c\xcf\xe5\x3b\xdf\xb9\x39\xa2\xd8\x2b\x6d\x61\xa9\xe4\xbc\x94\x5b\xb1\xce\x71\xa5\x5e\x50\xc2\x46\xab\x02\x3e\xfe\xb3\x9c\xaf\x3e\xcd\x66\x4f\x9f\x9f\x9f\x47\xa3\x7d\xb9\x86\x4c\x49\xab\x59\x66\xe1\x51\x6d\x85\x55\xd3\xbe\xe2\x79\x34\x02\x00\x20\xe1\x03\xd3\x60\x95\x65\xf9\x73\xb9\xdf\xe7\xa7\x29\xfc\xb5\x90\xf6\xf7\xdf\x6a\x01\x3c\xa0\x24\x53\xde\xe6\x42\x0a\x2b\x58\x2e\xfe\x45\x9e\xa4\x1d\x99\xbf\x85\xdd\x71\xcd\x8e\x89\xe0\x95\x99\x0f\x0e\xe3\x14\x3e\x71\xae\xd1\x98\xfb\xae\xca\x0c\xf7\xca\x08\xdb\xd2\x20\xc4\xb5\x7c\x54\xd0\x68\x54\xa9\x33\x84\xe5\x7c\xd5\x0f\x69\xbc\x58\xce\x57\x70\x76\xd2\x95\x46\x8e\x16\x6a\xc3\xa3\xd6\x19\x85\x5d\xa0\x65\x9c\x59\x36\x85\xf3\xb3\xd5\x42\x6e\xa7\xe0\xbf\x2f\xb5\xac\x90\x84\x4e\x0a\xbb\x98\x55\x86\xd2\x86\x1b\xfa\x18\xcc\x37\x63\xc1\xe1\x0e\xbc\x5c\xff\xb0\x72\x04\x77\x70\xbe\xc4\x63\xff\x74\x19\x88\xf0\x51\xe5\x39\x66\x56\x28\x39\x10\xe8\x17\xad\x0e\x82\xa3\xfe\xd0\x3f\x7a\xc2\x0c\xc5\x61\xf0\xa8\x36\xf9\xa5\x5c\xe7\x22\x6b\xc4\x30\x99\x00\x17\xee\x88\xe9\x13\xa8\x0d\x11\x4c\x35\xb4\x51\xba\x10\x72\x0b\x96\x0c\x98\xa6\x38\x09\x08\x03\xac\x46\x6c\x4f\x7b\x84\xa3\xb0\x3b\x60\x12\xbe\x7b\x9e\xbe\xc3\x62\x06\x1b\x81\x39\xef\x31\xaf\x8e\x12\xf9\x72\xbe\x32\x53\x78\x38\x7b\xe9\x81\x48\x97\xf3\x55\x27\x13\x90\x0c\x92\x1f\xcd\xc1\xed\x4d\x9b\xe1\x26\xea\x63\x28\x4f\xd0\x58\xa8\x03\x1a\x82\x4a\x91\xb8\x2e\xb2\x3b\x84\x2c\x72\x04\x4c\x72\xf0\x42\xc2\x82\x55\xfe\x98\xe5\x39\xea\x56\x2c\x9b\x52\x46\xb3\x49\xf5\xd0\xa8\x94\x29\x3c\x0c\x45\xd5\x89\x81\xea\xd4\x91\x4c\xf0\xdb\x01\x8d\x3d\xd6\xe4\x05\x4f\x53\xa8\x1d\xa4\x70\x7f\x0f\x7b\x26\x45\x96\xbc\x2b\x84\x31\x94\xa6\xe5\x7c\xf5\x2e\x1d\xb5\x0c\x63\x21\x3a\x5d\xe9\xdc\x8c\x05\xaf\xfa\x32\x7a\xd3\xf7\x63\xe6\x7b\xae\x63\x43\xa3\x2d\x35\x21\x73\xaa\x57\xa8\xe5\xbe\x8d\xc1\xb2\x17\xe2\xd5\xd1\x4a\x14\x32\xce\x5b\x0c\x46\x82\x4d\xa3\xe4\x9a\x86\xa2\x12\x89\x2f\x66\x95\xa2\xe0\xc0\xb4\x66\xa7\x1e\xf9\xc1\x71\xe2\xc0\x5d\x61\xbb\x5b\x32\x2d\xba\xfd\x03\x33\x3f\xc1\x83\x9f\x96\xa4\x31\xea\xc9\xd7\x63\x04\xee\x22\x89\x6d\x31\x42\xcf\xb9\x83\x2b\xf1\x18\x0c\x07\xfc\x8d\xfe\x3a\xee\x44\xb6\x8b\x25\x48\x87\x2a\xe7\xa0\x24\xf6\x7c\xaa\x9c\xaf\x86\xab\xe2\xab\xe0\xdf\x22\xf8\x81\x94\x37\xa7\x2a\xe5\x9a\x26\xea\xdb\x99\xe6\x68\xac\x56\xa7\xe8\xf7\x4a\xae\xb7\x68\x17\x33\x13\xea\xc2\x35\x91\x4b\x0d\x0d\x0e\x9f\x34\x8a\x8a\x59\x60\x1a\x41\xc8\x4e\xde\x7b\x09\xf4\xd6\x92\x74\x0a\x5f\x3d\xbf\xdf\x3a\xd9\x0a\xf5\xd7\x69\x8b\x17\x3c\x99\x2b\xf8\xd6\x4a\x6b\x75\xa4\x0a\xdc\xa2\xf5\x43\x6a\x83\x1a\x25\x4d\x29\x55\xf5\xfc\x75\x60\x93\x09\x18\xe5\x23\xa8\x9b\x1e\x32\x26\x41\x23\xe3\x20\xac\x89\x7b\xc3\x55\x2b\x09\x54\x6f\x77\x8a\x9b\x5e\x84\x11\x4f\x63\xc9\xa5\x53\x78\xff\x03\x83\x21\xc4\xfe\x7e\x20\xfb\xcc\x0c\x5b\x18\x22\x25\x24\xb6\x37\x3b\xab\x84\xb7\xcd\x0f\xef\xa7\xc9\x84\x02\xa2\xc5\xb1\x29\xa5\x9f\x91\x3e\xcb\xf2\xa4\x24\x3a\x7e\x1c\x13\x56\x41\xa6\x91\x59\x04\xe6\xda\x00\x8b\xbd\x3d\x75\x79\xae\xa8\xf1\x92\x9f\x49\xa4\xde\x4f\xc9\xe0\xd4\xac\xcf\x1b\x41\xc4\xd9\x54\xf9\x6c\x5a\xe9\xa0\x7f\x8a\xab\xca\xc3\x06\xc6\x0b\x21\x41\x69\x30\x8a\x52\x47\x23\xd4\x88\x42\xe4\x4c\xc3\x51\x95\xd4\x93\x47\xd7\xbf\x6b\xac\x4c\xb0\x75\xee\x6a\xa8\x10\xd2\xba\xe0\x22\x5d\x93\xc9\xe0\x35\xe5\x4f\x21\x2d\xea\xea\xb6\x15\xac\x90\x36\xa5\x9a\xbe\x4d\x60\x89\x7e\xfb\xed\xe9\x7e\x36\xee\x10\x61\x26\x56\x03\x96\xfe\x7c\xe9\x6a\xcc\xc4\x5e\x20\xd9\x68\xec\xad\xd2\xad\x02\xbb\x43\xa1\x9b\xaf\x63\x0b\xf4\xaa\x33\xa0\x49\xa2\xb9\x29\xbc\x3f\xbf\x79\x7b\xb8\xa4\xcd\xa0\x02\xce\x56\xde\x9b\xa5\x48\x1f\xda\xfa\x12\x5d\x5b\xd6\xe9\x72\x4d\x11\xee\x56\x61\xfa\x36\xae\xa2\x69\xcf\xc3\x2b\x2c\xfc\x6c\x80\x65\x99\x2a\xa5\x6d\x71\xd0\x0f\xdc\xd7\x4d\xd0\x1a\x77\xf6\xc7\xed\x8d\xc7\xd8\x71\xdd\xc7\x06\x77\x43\x2f\x7f\x09\xed\x9d\xfc\x9a\x0e\x77\x91\xbb\x4d\xa6\xed\xbb\x57\x7d\xa3\x76\x11\x39\x7b\x60\x9c\xc1\x28\xe6\x3a\xb4\xed\xfe\x63\xab\xa8\x1e\x2b\xee\x1f\x9b\x49\x0f\x95\x48\x05\x64\xd8\x01\xc3\x26\x36\x56\x69\xb6\xad\x19\xa1\x5d\xd3\xa8\x95\x57\xba\x29\x42\x09\x54\x8f\xc9\x6a\x72\x7b\x53\x6b\xfb\x5d\x33\x09\x2e\x26\x9e\xa3\xda\x4e\xda\x02\x1d\x0b\x26\x4c\x96\x8c\xed\xd9\x5a\xe4\xc2\x9e\x60\xa3\xf4\xb5\x01\xdd\x02\x90\x0b\xf9\x72\xfb\x23\x05\xfb\x47\xd2\x2e\x26\xef\xb2\x07\xf0\x43\x4b\xca\x32\xbd\x45\xfb\x4a\x3c\x51\x3a\x1d\xce\x46\x18\x00\xff\x27\x13\x85\x57\x69\x35\x89\x37\xf3\x46\x12\xbc\xe2\x60\x02\xbc\x7e\x03\xa3\xbb\x25\x5c\xff\x97\xee\x32\xba\x8c\x46\xff\x05\x00\x00\xff\xff\x89\x9b\xdf\xbf\x75\x0e\x00\x00"

func cogitoCdcBytes() ([]byte, error) {
	return bindataRead(
		_cogitoCdc,
		"Cogito.cdc",
	)
}

func cogitoCdc() (*asset, error) {
	bytes, err := cogitoCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Cogito.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0x92, 0xe5, 0x7c, 0x9d, 0xa1, 0x8d, 0xd0, 0x79, 0xe3, 0x79, 0x96, 0x19, 0x7d, 0xd7, 0x4d, 0x90, 0xef, 0xd4, 0x8e, 0xf, 0x4b, 0x4d, 0x6a, 0xee, 0x68, 0x64, 0x88, 0x14, 0x76, 0xdd, 0x59}}
	return a, nil
}

var _nonfungibletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x41\x8f\xdb\xba\x11\xbe\xeb\x57\xcc\xcb\x03\x9a\xdd\xc0\x6b\xf7\x50\xf4\x60\x20\x68\xda\xb7\x6f\x01\x5f\xb6\x0f\x5b\x17\x3d\x04\x01\x4c\x8b\x23\x9b\x08\x45\x2a\x24\x65\xc7\x0d\xf6\xbf\x17\x33\x24\x25\xca\xf6\x26\x9b\x5b\x73\x89\x57\x22\xbf\x99\xf9\xe6\x9b\x8f\xd4\xe2\xdd\xbb\xaa\xfa\xf5\x57\x58\xef\x11\x1e\xb4\x3d\xc2\xa3\x35\x77\x0f\xbd\xd9\xa9\xad\x46\x58\xdb\xcf\x68\xc0\x07\x61\xa4\x70\x92\x17\x6e\x1e\xad\xc9\xef\xf9\xf5\x06\x6a\x6b\x82\x13\x75\x00\x65\x02\xba\x46\xd4\x58\x55\x84\x37\xfc\x09\x61\x2f\x02\x08\xad\xc1\x58\x73\xd7\x64\xf4\xc0\xe8\x79\xb7\x87\xda\xf6\x5a\xd2\xdf\x8d\x75\x2d\x04\x3b\xaf\x56\x0d\x08\xe8\x3d\x3a\x38\x0a\x13\x3c\x04\x0b\x12\x3b\x6d\x4f\x20\xc0\xe0\x11\x4c\x13\x86\xfd\x33\x08\x7b\x54\x6e\xcc\xe6\xc8\x70\x06\x51\x56\xc1\x82\x6a\x3b\x8d\x2d\x9a\x40\xcb\xe0\xbc\x88\x31\xd7\x39\xe7\x7e\x89\xb3\x17\x07\xca\x18\x1a\xab\x89\x26\x2a\x86\x80\x5c\xaf\xd1\x83\x30\x12\x8c\x68\x95\xd9\x55\x5c\x6a\x98\x54\xef\x3b\xac\x55\xa3\xd0\xcf\x13\x83\x0f\xeb\x0d\x38\xf4\xb6\x77\x99\xaa\xda\x3a\x1c\x1e\x41\x38\x75\x89\x33\x87\x9d\x43\x8f\x54\xbb\x30\xf0\xf8\xb0\x06\x65\x18\xdd\xb7\xc2\x8d\xb5\x27\xe0\xdf\xac\xd6\x58\x07\x65\xcd\x06\x9e\x26\xf8\x23\x34\xa1\xfa\x60\x1d\x65\xcd\xd4\xbe\xf5\x8c\x5b\x0f\x7b\xe7\xd5\x8a\x5a\x59\xeb\x5e\xf2\xa2\x06\x8f\xd0\xf4\x86\xdf\x71\x0b\x04\x33\x40\x59\xd8\xa3\x41\x47\x8f\x50\x78\xa5\x4f\x55\x6b\x0f\xa9\xad\x9e\x12\x25\x5a\x6c\x1f\xc0\x36\xbc\xba\x0c\xc1\xf9\xfe\xe1\xec\x41\x49\x74\x1b\x5e\xb9\x79\xc2\x1a\xd5\x81\xfe\x1c\xd2\x1d\x48\xf4\x5c\x87\x2f\x9f\x80\xc4\x5a\x0b\x87\x45\x72\x47\x15\xf6\xe0\x6d\x8b\xd0\x39\x64\xd0\xce\x7a\xa6\x49\x2a\x5e\x51\x25\x56\xbf\xf4\xca\x21\x27\x35\x72\x56\x74\xb7\x46\x17\x84\x32\xa9\xa7\x0c\xb4\xc5\xbd\x38\x28\xeb\x86\x69\xf0\x51\x29\x27\xa0\x14\x3c\x76\xc2\x89\x80\xb0\xc5\x5a\xf4\x94\x66\x80\x9d\x3a\xa0\xe7\x18\xac\x60\xfa\x21\xb6\x4a\xab\x70\xa2\x48\x7e\x4f\xfb\x04\x38\x6c\xd0\xa1\xa9\x91\x44\x1a\x15\x5c\xa6\x44\xe9\x5a\xa3\x4f\x80\x5f\x3b\xeb\x13\x5e\xa3\x50\xcb\xa8\xba\xb1\x76\x65\xc0\x1a\x04\xeb\xa0\xb5\x0e\xab\xc4\xf9\x48\xd7\x1c\x56\x34\x83\xde\xa6\xc4\x28\x29\x7f\x9e\x55\x2b\x3e\x23\xd4\xbd\x0f\xb6\x1d\x9a\x90\x48\x9b\x0c\xd0\xb4\x11\x34\x96\x16\x0e\xc2\x29\xdb\x13\xa4\x32\xbb\xd4\x0b\x82\x8f\x7a\x98\x57\xd5\x3f\x4e\xd0\x7b\xe2\x73\x40\xe6\x12\x46\xa0\x59\x4a\xca\x36\x2c\xc9\xa9\xc6\x3d\xd4\xc2\x80\x47\x23\x2b\xda\xe5\xa2\x58\xb2\xda\x3a\x44\x77\x17\xec\x1d\xfd\x3f\xe3\xd8\x24\x3c\x6a\x99\xd9\x51\x7e\x1c\x84\xa7\x99\xd2\x12\x50\x23\xa1\x6a\xd0\x28\x77\xe8\xaa\x8b\x71\x5a\x5b\x0e\x95\xa7\x8e\x54\x6f\x6c\xd8\xa3\xe3\x14\x67\x83\x2d\xb1\x37\x78\xe2\xe6\xc4\xd0\xd2\x89\x38\x1a\x8f\x0f\xeb\xaa\x71\xb6\xbd\xe8\x29\xfb\x94\x81\x3a\x3b\x88\xc4\xce\x7a\x15\x86\x4e\x82\x35\x93\x58\x6f\x7d\x35\xd5\x68\x6d\xa9\x13\x21\xca\x37\x38\x61\x7c\x83\x6e\x5e\x55\xef\x16\x55\xb5\x58\xb0\x93\xb7\x24\xde\x38\xd5\xe7\xd6\x3c\x87\x7f\x32\x74\xf9\x96\x9a\xa5\x35\x6d\x56\x6d\x67\x5d\x88\x6d\x29\xfa\xad\x7c\xe1\xed\x8b\x45\xd5\xf5\xdb\x2b\xd0\x97\xae\xfa\xad\xaa\x00\x00\x52\x56\xc1\x06\xa1\xc1\xf4\xed\x16\x1d\x7b\x42\x6c\x1d\x2b\x55\xf9\xe8\x7a\xca\x00\x7e\x55\x3e\xf0\x44\xd0\x5e\x0a\x75\x10\x2e\x6e\xfe\x57\xdf\x75\xfa\xb4\x84\x7f\xaf\x4c\xf8\xeb\x5f\x06\xf0\xdf\x0f\x31\x4d\x11\x00\x5b\x15\x02\x4a\x38\x12\xc7\xa9\x0f\x45\xaa\x54\x87\x0a\x4a\x68\xf5\x5f\x94\x69\xfb\x10\x06\x19\xe6\xb7\xb4\x78\x35\x2e\xbc\xb9\xbd\x16\x4a\xf9\x69\x34\x91\x0e\x34\xe5\x07\x25\x98\x59\xde\xa7\x8c\x54\xb5\x08\xac\xc6\xc1\x38\x2f\x7c\x31\x01\x07\x38\x8a\x02\x04\x48\x47\xf3\x32\xdb\xc5\x02\x56\x17\x7b\x95\x07\x63\x43\xf4\x5d\x10\x75\x6d\x7b\x13\xde\x7a\x36\x7b\xb1\xc3\x19\x6c\x08\x66\xc3\xad\x86\x2d\xc2\xc6\x28\xbd\x99\x5f\xe7\xe0\x3f\x29\xf4\x8d\x92\x99\xec\x19\x67\xb1\x84\xbf\x4b\xe9\xd0\xfb\xbf\x5d\xa5\xe4\x25\x3e\x92\xc6\x51\xf2\x20\x4d\x0e\x82\xb3\xaa\x42\x66\x2a\x59\xdd\x6b\x88\x2a\xd1\x5f\x28\xe8\x3e\x2e\x99\xd4\x13\xec\xb5\x6a\x56\xd3\x4b\x4b\x92\x90\x1f\xce\xff\xf1\x7a\x72\x1e\xe9\xf2\xd0\x82\x15\xa9\xef\x1b\xaf\x28\xe6\xa0\x37\xea\x4b\x8f\xb0\xba\x4f\xa4\x89\x7a\xcf\x32\xdd\x0b\x3f\x2c\x25\x40\x8d\x01\xc6\x84\xf9\xd5\xf3\x90\xe7\x53\x3c\xc3\xda\x81\x7b\xf2\x93\x94\x1c\xa9\xec\x9a\x81\x52\x0d\x79\x3f\x5f\xa5\x1a\x65\xe2\x19\x94\x32\x27\x53\x42\x19\x1d\x8f\x30\x13\x1e\x3b\x3c\xd5\x72\x59\xeb\xe3\xc3\x7a\x79\x5e\xe6\x0f\x73\x2f\x38\xb6\xd0\xa2\x54\x74\x72\x66\xb9\x7b\xc8\xb6\x59\x98\xe6\x2b\xb8\xce\x97\x89\x29\xdf\x83\x27\x3b\xa4\xcb\xc9\x70\x8d\x1a\x62\x14\x9a\x22\xd7\x8b\x8b\x54\x80\x78\x1a\x47\x46\xdc\xa4\xb4\xa6\x37\x03\xec\x4d\xfe\xb1\xba\xcf\xb5\xde\x2e\xe1\xc3\x94\x0f\xde\x48\xf7\x90\xe9\x23\xfa\xe7\xd0\xf7\x3a\xcc\x95\x84\xf7\xef\xa1\xc4\x7a\x43\x42\x59\xdd\x67\xe5\x8f\x5e\x10\x67\xaa\xed\x7d\xa0\x21\xe6\xab\xa0\x68\x11\x44\x1c\x17\xba\xd9\xa0\xa7\x51\x58\xdd\xbf\x99\x44\x7b\xae\xa6\xbf\x7e\xd0\x8d\x34\x53\x3e\xf3\xf0\x53\xad\xc8\x17\xb9\xec\xff\x29\x50\x3e\xe9\x82\xf8\x3c\x36\x42\xf0\x2f\xe1\x76\x3d\x4b\x99\x7a\x20\xa4\x2c\x5b\x70\x16\xba\x08\x5f\x76\x24\x81\xdf\x30\x3f\xb1\x05\xb7\x2f\x17\xca\x03\x33\xb8\x64\x3a\xc6\x6b\xdb\xb6\x7c\xd7\xca\x1b\xba\x7e\xab\x95\xdf\x43\x63\xdd\xf0\x71\x31\xc9\xe5\x85\xfa\xc7\x8c\xff\x20\x84\xfa\x6c\x36\xbe\x9b\x6e\xb9\x68\x87\x61\x75\xef\x6f\x6e\x97\xf0\x31\x6a\xeb\xd3\xc5\x92\xad\x75\xce\x1e\x1f\x1f\xd6\x85\xb5\xdd\x2e\xe1\x4f\x79\x58\xaf\x1b\x46\x2a\x28\x0d\x80\xa9\x1d\x5d\x27\x26\x9f\x1f\x85\x4d\x6c\x31\xdf\xb4\x65\xfe\xfa\x18\xee\x06\xe4\x34\xd9\x5f\x5e\x14\xc6\x48\xc7\x72\x98\xd2\xd9\x20\x92\xd9\x35\xba\x4a\xd9\xdc\x2b\x7e\x27\x1c\xdf\x50\xf7\x56\xcb\xd1\x95\x53\x3e\x57\x24\x92\xef\x0d\x74\x80\x48\x5a\xbb\x84\x0f\xdf\x22\x3f\x4b\xda\xfb\x5c\xfd\x5f\xd8\xc4\xf7\x06\x24\xce\xc7\xe5\x40\x8c\xb9\x78\x90\x03\x39\x25\xd0\xb0\x29\x44\x17\x49\x1b\x95\x04\xe1\x9c\x38\xbd\x4e\x8d\x25\x60\x54\x22\x38\x0c\xbd\x33\x69\x62\x9d\x38\x65\x7b\xa2\x77\x71\xa6\x1c\xe6\x9e\xd4\xd7\x7b\xf2\x82\xae\xcb\x60\x4f\x39\x4a\x52\x37\xca\xf1\x2b\x29\xde\xc4\xcb\x2f\xe1\x2b\x71\x16\x0b\xf0\x76\x3c\xbf\x63\x73\xf8\xf3\xc1\xa1\x90\x20\x45\x10\x4c\x11\xdf\xc1\x5b\x0c\x7b\x2b\xd3\xa9\xa3\xc2\xcf\x4c\xd8\xb9\xc7\x3b\xbc\x62\xf1\x1e\x75\x33\x1f\x54\xf8\x51\xc9\x4f\xf0\xcb\x7b\x30\x4a\x2f\xe1\x0d\x61\x48\x8b\xf1\xe2\xc6\xf7\xde\xcb\xaa\x7e\x79\xad\x8f\xd7\x0e\x45\xc0\xdf\xdb\x2e\x9c\x8a\x0f\x86\xf8\x94\x5b\x86\xf4\xea\xd2\xc9\x21\x7e\x4e\x45\xce\xcf\x25\x5d\x12\x79\x62\x0a\xed\x91\xe9\xf7\x55\x49\xd2\xd5\xd8\xd4\xe0\x0f\x45\x2a\x85\x0b\x5e\x9e\x86\xe9\x24\xcc\xd2\x98\x6b\x34\xbb\xb0\xa7\x63\xf1\xcf\xe9\x34\x8c\x31\x64\x39\x8a\xf9\x18\xe4\xca\x0a\xa2\x32\x35\xcf\xd5\xff\x02\x00\x00\xff\xff\x33\x4d\x81\x27\xe0\x12\x00\x00"

func nonfungibletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_nonfungibletokenCdc,
		"NonFungibleToken.cdc",
	)
}

func nonfungibletokenCdc() (*asset, error) {
	bytes, err := nonfungibletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NonFungibleToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0x61, 0xca, 0x9d, 0xaa, 0x66, 0x36, 0xdf, 0xbc, 0x51, 0xdb, 0x7b, 0x51, 0xd8, 0x3d, 0x6f, 0x4e, 0x9c, 0x8e, 0x50, 0x28, 0x7c, 0x18, 0x1d, 0x2, 0xb2, 0xc2, 0x2b, 0x26, 0xa1, 0xfe, 0x2d}}
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
	"Cogito.cdc":           cogitoCdc,
	"NonFungibleToken.cdc": nonfungibletokenCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"Cogito.cdc":           {cogitoCdc, map[string]*bintree{}},
	"NonFungibleToken.cdc": {nonfungibletokenCdc, map[string]*bintree{}},
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
