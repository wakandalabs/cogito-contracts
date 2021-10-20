// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../contracts/Cogito.cdc (5.916kB)
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

var _cogitoCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x58\x5f\x6f\xe3\xb8\x11\x7f\x3e\x7f\x8a\xd9\x3c\xec\xca\x68\xd6\xba\x02\x45\x1f\x8c\xb8\xd9\x34\x5e\x1f\x8c\xa2\xc6\x62\xe3\xe2\x1e\x0e\x87\x1e\x2d\x8e\x2c\x22\x32\x69\x90\x54\x7c\x6e\x2e\xdf\xbd\x18\x52\x94\xa8\x3f\x4e\xd2\xe6\x61\x57\x26\x39\xc3\xe1\xfc\xf9\xcd\x8f\x14\x87\xa3\xd2\x16\x36\x4a\xae\x2a\xb9\x17\xbb\x12\xb7\xea\x11\x25\xe4\x5a\x1d\xe0\xc7\xdf\x37\xab\xed\xdd\x72\xf9\xfd\xeb\xc3\xc3\x64\x92\xa6\x93\x34\x85\xcd\x6a\x0b\xb9\xd2\x70\xaf\xf6\xc2\x2a\x40\xbd\x57\x60\xaa\xc3\x07\x9a\x3e\x56\x3b\xc8\x94\xb4\x9a\x65\xb6\x5e\x30\x1f\xaa\x7e\x9e\x4c\x00\x00\xd2\x14\xbe\x3e\xa1\xb4\xa6\xfe\xe5\xfe\x23\x0d\x48\xa3\x70\x5f\xeb\x59\x4b\x61\x05\x2b\xc5\x7f\x90\x27\xd3\xde\x9a\x9f\x85\x2d\xb8\x66\xa7\x44\xf0\x39\xfc\x6b\x2d\xed\x5f\xff\x72\xed\x2c\x9f\xc3\x1d\xe7\x1a\x8d\xb9\xed\x8b\x2c\xf1\xa8\x8c\xb0\x1d\x09\xb2\xf2\xd2\xfa\x7f\x0a\x69\x91\x77\x96\x1f\xd0\x32\xce\x2c\x9b\xc3\x83\xd5\x42\xee\xfb\x22\x7f\xaf\xb4\xec\x88\x4c\x9b\x03\x6f\xd8\x01\x39\x7c\x63\xb6\x18\x9c\xba\x44\x3a\x73\x59\x62\x66\x85\x92\x0f\x56\x69\xb6\x47\x5a\x49\xdb\x34\x3f\x2e\x2c\xff\x56\xed\x4a\x91\xf9\xd5\xed\x77\xb3\xad\x55\x96\x95\x0f\xd5\xf1\x58\x9e\xc3\xd0\xb6\x40\x3f\x0c\xb2\x3a\xec\x50\x83\xca\x43\x48\x6d\xc1\x2c\x14\xec\x09\x61\x87\x28\xe1\xe0\x3c\xd0\x37\xf7\x89\xe9\x58\x6d\x38\x6a\x7b\xd2\xd5\x36\x7c\xde\xc1\x3f\x84\xb5\x67\x58\x5b\x3c\x00\x33\xc0\x64\x34\xdb\x28\xd4\x68\x54\xa5\x33\xa4\xb9\x61\xd2\xcc\xd6\x94\x77\x75\xe6\xc4\x3e\x68\xbd\x3c\x9c\xeb\x07\xaa\x5d\x91\xa6\x20\x9a\xbc\xd2\xd1\x70\xf3\x49\xd3\xaf\x47\x1d\x9e\x9b\xc5\xf4\x67\xb0\xcc\x67\x82\xc3\x02\x04\x1f\x4e\x04\x69\x58\x34\x8a\x9a\x45\x2f\x13\xff\x6f\x1b\x1a\x61\x40\x18\xb0\x05\x02\xf9\x5e\xe7\x2c\x43\x1f\x95\xca\xa0\x36\x90\x31\x09\x19\x33\x96\x56\x88\xa6\x12\xdb\x6c\x00\x66\xda\xc8\x03\x2b\x4b\x75\x02\x65\x0b\x12\xb5\x0a\xb8\xaf\x80\x20\x26\xa4\x0b\xb9\x57\x14\x34\xcc\x60\x6d\x81\x95\xa6\x96\x36\xae\xe2\x35\x32\x4e\x5e\x0c\xaa\x0b\x04\x8e\x96\x89\xd2\x44\xc9\x23\xa4\x9b\x88\x54\x0d\x23\xdc\x1e\xca\x0b\xf5\xf3\x38\xf2\x2c\xc9\xe5\x95\x84\x9d\xd2\x5a\x9d\xfc\xf2\xb8\xb2\xe6\xf0\xd1\x0f\xce\x36\xab\xed\x6d\x2f\x24\x69\x0a\xeb\xdc\x99\xa3\xd1\x54\xa5\x05\x61\xe4\x27\x0b\x52\x94\xd7\xde\xb9\x9c\x0c\xf7\xf3\xd6\x15\x2d\x68\xcc\x51\xa3\xcc\xb0\xaf\xc8\x14\xaa\x2a\x39\xec\xd0\xad\x37\xec\x80\x94\xca\xf4\xcd\xf4\xbe\x3a\x50\xdd\x7b\x3f\x92\xb9\xee\x2c\x1d\x0d\x47\x65\x6c\xcf\x3a\xfa\x4b\x6a\xc3\x16\x0b\xb2\x6a\x0a\x7f\xfc\x11\x86\x6e\x5d\x2e\x51\x32\x4d\xe7\x03\x31\xfa\xbb\xba\x67\x52\x2a\x5b\x7b\x26\xb8\xbf\xb1\x7f\xee\x2a\x7c\xbd\xbc\x7c\x42\x4a\x31\x21\x33\xa5\x35\x66\xf6\xaa\xb3\xc7\xcb\x64\xf8\x15\x22\x41\x70\xb0\x59\x6d\x93\x71\x14\xec\xc4\xac\xd2\x92\x56\xc6\x40\xd8\xcd\xf4\x36\xf0\x2d\x56\x64\x6d\x1a\xb7\x59\xb5\x59\x6d\x0d\xa8\x13\xd9\xbf\x3b\x13\x80\xb0\x2c\x53\x95\xb4\x17\x41\xa4\xd5\x3c\xbf\x90\x64\xd7\x43\x8c\xf9\xa6\xd5\x93\xe0\xa8\x47\xa6\xbe\x63\x86\xe2\x69\x74\xea\x95\xec\x4d\x53\xe0\xc2\x4d\x31\x7d\xa6\xe3\x10\x88\x65\x4a\xe6\x4a\x1f\x84\xdc\x83\x25\x05\x26\x5e\x4e\x0b\x84\x01\xd6\x9e\xc4\x9e\x8f\x08\x27\x61\x0b\x3a\xf6\x6f\xde\x8f\xbf\x51\x60\x73\x81\x25\x1f\x83\xae\x00\xcf\xce\x5f\xe4\xb9\x39\x7c\x79\xf6\x82\x23\xc0\xba\x59\x6d\x5f\x3a\xc0\x78\xaa\x9b\x6a\x3c\xf6\x1d\x0f\xea\x09\x03\x74\x7b\x72\x40\x59\x15\x05\x8b\x49\x0e\x7e\x91\x68\x4a\x21\x63\x65\x39\x0e\xaf\x21\x47\xc2\x66\x49\xf8\x58\x2f\xa3\xda\xfe\x32\x66\x6d\xaf\x8c\x08\xe8\x9d\x1f\xe1\xe6\xb3\xc7\xda\xe6\xdc\x33\xed\xcc\x4e\x1e\xf1\x3c\x87\x76\x83\x29\xdc\xde\xc2\x91\x49\x91\x25\x57\x07\x61\x0c\x45\x62\xb3\xda\x5e\x4d\x27\x1d\xc5\x78\x10\x3d\x86\xe1\xb6\x99\x09\x1e\x38\x46\xb3\x9b\xbe\x9d\x31\xcf\x1f\x7a\x3a\x7c\xd5\xc1\xcd\x67\x27\x1a\x55\x55\x27\x45\x3c\x20\xc7\x43\x5b\xf6\x48\xde\x76\xce\x26\xc7\x32\xce\x3b\x7e\x6d\xdc\x6e\xa2\x04\x8b\x15\x34\x42\xd6\xa3\x40\x2d\x28\x38\x30\xad\xd9\xf9\xb5\x90\xd4\xe6\x24\xce\xe4\x0b\x31\xe8\x37\xbf\x4e\x10\xfc\x07\x33\x1f\xe0\x4b\x8b\xcd\x93\xc1\xfa\x16\x16\x60\xd1\xb8\x76\xd2\xc7\x5d\xc6\xb9\xb3\x5c\xe2\xa9\x56\x5c\x1f\x25\x2a\xac\x53\x21\xb2\x02\x74\x9d\xa3\x34\xa9\x4a\x0e\x4a\xe2\x60\x4f\x55\xf2\xed\x78\xae\xfc\x22\xf8\xaf\x8d\xf1\x23\x89\x10\xf3\x46\xca\x00\xe2\x8c\x6f\xc7\x9f\xa3\xb1\x5a\x9d\x9b\x7d\x2f\x64\xc0\x1e\xed\x7a\x69\xba\x05\x47\x89\xe3\x0a\xce\x05\x2c\xc0\xf8\x7a\x69\x3c\x13\x60\x1a\x43\xaf\xcd\xba\x28\x7a\x21\xac\x7e\x8f\x64\x3a\x87\x5f\xbc\xd7\x7f\xed\xc5\xb0\xce\xd5\x5e\x09\x3d\xe2\xd9\x5c\xb0\xda\x77\x9e\x40\xe5\xea\xc1\x9f\xd0\x7a\xf8\x0a\x5d\x86\x18\x88\x47\x8d\xd7\xcc\x05\x53\xf3\xce\x16\x36\x1c\xcd\x21\xce\x01\xc2\x9a\x86\x35\xb9\xcc\xa6\x05\x61\xb4\x50\xdc\xbc\x76\xee\xc6\xca\x1e\x6d\x78\x07\xb4\xd4\x1e\xf9\x38\x92\x29\xcc\x8c\x6b\x78\xd5\x55\xbe\x1a\xfe\x5f\x6f\x39\xea\x5c\x37\xb3\xeb\x58\x09\xfe\x4e\xa9\x29\xf7\xc4\xd5\x28\x4f\xc8\x31\xae\x39\x18\x48\x84\xcc\xca\x8a\xbb\x5e\x53\xf8\x5e\xb2\x5e\x4e\x67\x1d\xa8\xa9\xd9\xa6\x61\x79\xa0\x34\x1a\x5d\x7a\x49\xd5\x70\x99\x3a\xe9\x28\x22\xbb\x3a\x3e\x54\x60\x35\xd5\x73\x35\xfe\x76\x0c\xfe\x47\xf6\x26\xf2\xb1\x12\xfd\xe0\x98\xd2\x08\x95\xa2\xf2\xd6\x98\xc3\xe2\x52\xbc\x58\x65\x8b\x37\x82\xd6\x0b\x3c\xa9\x23\x18\x8b\x8c\xec\xd2\x24\xc0\xd2\xe0\x88\x2d\xb5\xbc\x14\xe5\x45\x5a\xf5\x7e\x5e\xd5\xd3\xef\xd0\x28\xba\x95\xd6\xb6\x45\x77\xb1\xf8\xae\x12\xbe\x7a\xb8\xe4\x3c\xd4\x03\xfa\x9b\xcf\x99\x46\x66\x9b\x80\x86\x92\x79\xe7\x06\xbd\x1d\x86\x52\xb0\x18\x1b\xfc\x13\x24\x7f\xa6\xe8\xc4\x04\xf1\x82\x87\x46\xf8\xe4\x98\x6f\xa2\xeb\xb7\xe0\xd3\x61\xdf\xe9\x92\x85\x51\x02\x32\x90\x0b\x38\xfe\x7a\x1b\x37\x56\x57\x99\x55\x2d\xe1\xa9\xc5\x92\xbe\x9d\x41\x5d\x37\x51\x2f\xe8\x7d\xc7\x45\x15\x06\x3b\x74\x35\xd3\x69\x9f\xfb\xa4\xbe\xa5\xe1\x3e\xee\x5f\x0f\x47\x7b\x1e\x32\xf2\xa3\x27\xb5\x01\x09\xea\xee\x23\xcf\x4a\x62\x7d\x11\x2d\x4b\x82\xad\x3a\x79\x98\xeb\xd4\x48\xba\xfa\x48\x1f\xf1\x74\x0a\xe7\xe8\xae\xc9\x28\xeb\x8b\xae\xb6\xed\x31\x1b\x6e\x05\x4d\xde\xb6\x5a\x7a\x27\xcc\xd1\x66\x45\xf8\xf1\x13\xda\x01\xe6\x86\x8b\x86\x63\xb6\xed\xfd\xe2\x93\x89\xb4\x5e\x13\x24\xb1\x27\x26\x4a\xb6\x2b\x71\x16\xd4\xad\xf3\x48\x00\xb8\x42\x03\x74\x2f\x73\x0f\x28\x41\xef\x2c\xd6\xe2\xa8\x67\x2c\x2e\x68\x35\x81\x7b\x04\xf6\xbb\x2a\xd2\x95\x29\x69\x59\xdd\x10\x88\x7c\xb4\xf8\xf2\x96\x1a\xea\x96\x1e\xba\xdb\xb1\x5a\x9b\xe9\xab\xeb\xf9\x84\xc4\x66\x63\x81\x73\xce\x4c\xfe\xdd\x7d\x68\xbb\x86\xf7\xe0\x3a\xd5\x5f\x64\xc9\x82\x58\xc9\x9d\xf7\x5c\x42\xea\xa6\x9d\x24\x9e\xed\xd1\xde\xb3\x23\xdb\x89\x52\xd8\x73\x32\x70\x65\xfb\xd2\xd5\x93\xf3\xfd\xe6\xe6\xe3\x40\xe2\xb9\x19\x19\xbb\x12\xbe\xfc\x2d\xe9\x2a\x6a\xef\x09\xf7\x74\xfb\x97\x9f\x2c\x19\x1c\x9d\xe0\x6a\x1a\x57\xea\xcf\x08\x56\x57\xc6\x0e\x83\x3e\x8b\x3b\x20\x39\x97\xd4\xf8\xf6\xee\xae\xdf\xf1\x9b\x51\x9a\x42\x22\x2c\x64\x05\x66\x8f\x8e\xf4\xef\x30\x57\x3a\xdc\xe0\xa9\x95\x0b\x1b\xb5\xf0\x3a\x78\xd9\xf8\x5e\x0d\x0c\x76\xeb\xa1\x8f\x29\x69\x3a\xf9\xc1\xbd\x78\x4d\xbb\xb7\xd7\x07\xe2\xcc\x95\x06\xe9\xde\x2d\x8f\xcd\xbb\x65\x03\x31\xa3\x2f\x96\xb0\x80\xd4\xf8\x9f\x69\xd6\xf3\xf4\x25\xf1\x36\x96\x24\xed\x41\x67\x28\x1c\xdb\xd6\x3e\x0c\x7b\x6e\xe3\x1e\x33\x4d\xfb\xc8\xd9\x6c\xd2\xed\x41\x3f\xb6\x4a\x5c\xbb\x18\x7f\x67\xfe\xe1\x65\xf2\x32\xf9\x6f\x00\x00\x00\xff\xff\xad\xdb\x35\x24\x1c\x17\x00\x00"

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x70, 0xe6, 0x4b, 0xc5, 0xad, 0xb0, 0xf0, 0x98, 0xa2, 0x91, 0x8e, 0x74, 0xe9, 0xc, 0xc3, 0x71, 0x98, 0x10, 0xea, 0x84, 0xf1, 0xf6, 0xa5, 0xb6, 0xa9, 0x6e, 0xf7, 0xfd, 0x16, 0xcb, 0x1f, 0xa6}}
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
