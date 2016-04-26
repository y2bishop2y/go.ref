// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by go-bindata.
// sources:
// v.io/v23/vdlroot/math/math.vdl
// v.io/v23/vdlroot/math/vdl.config
// v.io/v23/vdlroot/signature/signature.vdl
// v.io/v23/vdlroot/time/time.vdl
// v.io/v23/vdlroot/time/vdl.config
// v.io/v23/vdlroot/vdltool/config.vdl
// DO NOT EDIT!

package builtin_vdlroot

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

var _mathMathVdl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x8e\xb1\x4e\xc3\x30\x14\x45\xe7\xfa\x2b\xee\x0f\x34\xa5\x69\x15\xb1\x96\xd2\xa1\x12\x62\xa0\xc0\xfe\x92\xbc\x24\x16\x8e\x1d\xd9\xcf\x15\x11\xea\xbf\xe3\x64\x40\x19\x90\x10\xe3\xb1\x7c\xde\x3d\x9b\x0d\x8e\x6e\x18\xbd\x6e\x3b\x41\x7e\xb7\x2d\xf0\xda\x31\xde\xc9\x52\xad\x63\x8f\x43\x94\xce\xf9\x90\xe1\x60\x0c\xe6\x4f\x01\x9e\x03\xfb\x2b\xd7\x99\x4a\xf2\x5b\x60\xb8\x06\xd2\xe9\x80\xe0\xa2\xaf\x18\x95\xab\x19\x09\x5b\x77\x65\x6f\xb9\x46\x39\x82\xf0\x70\x79\x5c\x07\x19\x0d\x4f\x96\xd1\x15\xdb\x64\x4a\x47\x82\x8a\x2c\x4a\x46\xe3\xa2\xad\xa1\x6d\x7a\x64\x3c\x9d\x8f\xa7\xe7\xcb\x09\x8d\x36\x9c\x29\x35\x50\xf5\x41\x2d\xa3\x27\xe9\x94\x9a\x9b\xfb\xc1\xf0\x67\xb1\x9f\x86\x28\x4d\xce\x08\x1b\xfb\x92\xfd\x8c\x2e\xa4\xe5\x54\xb6\xcb\xd7\xa5\x96\x14\x4d\x06\x34\x0d\xf4\xd4\x6a\x4b\x7e\xc4\x40\x5e\x42\xa6\x64\x1c\x78\x71\x2f\x88\x8f\x95\xe0\x4b\xad\x5e\x26\xa5\x31\x8e\x64\x97\xab\xd5\x39\x79\x3f\x74\x5b\x46\x6c\xf3\xfb\x3f\x2b\x8a\xfd\x3f\x2a\xa6\x83\xbf\x65\x14\xfb\x65\x46\xa2\x9b\xfa\x0e\x00\x00\xff\xff\xce\xca\xd1\x3a\xbf\x01\x00\x00")

func mathMathVdlBytes() ([]byte, error) {
	return bindataRead(
		_mathMathVdl,
		"math/math.vdl",
	)
}

func mathMathVdl() (*asset, error) {
	bytes, err := mathMathVdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "math/math.vdl", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mathVdlConfig = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x8d\x41\x4b\xc4\x30\x10\x85\xcf\xed\xaf\x78\xf4\x2c\x5d\x14\x11\x09\x78\xda\x83\x07\x71\x4f\x15\xc1\x5b\x4c\xa6\xed\x60\x9a\xa9\xd9\x74\x51\x64\xff\xbb\x49\xa3\x7b\xee\x21\x13\xf8\xde\x9b\x6f\x76\x3b\x74\x23\xe1\x64\x5d\x6b\xc4\xf7\x3c\xa0\x67\x47\xe8\x25\x20\x26\x3e\xe9\x38\x62\xd6\xe6\x43\x0f\x84\x52\x58\x02\x1d\xc1\x96\x25\x65\x6c\xe0\xd3\x3c\x11\xe2\xf7\x4c\xc7\x3a\xd9\xd8\x83\xb4\x19\x31\x90\xa7\xa0\x23\x59\x38\xed\x87\x25\xed\xb7\xf5\xdf\x85\x87\x7c\x2e\x8a\xb8\x76\xbf\x82\x9f\xba\x7a\x14\x85\xf4\x55\xaf\x1c\xa8\x93\xc3\xea\xec\xb2\xb2\xe0\xaa\xd9\xcb\x34\x3b\xfa\xba\xbb\x6d\x14\x0a\xaa\x9e\xd8\x5b\x85\xc3\x32\xbd\x53\xb8\x5a\x49\xde\x50\x68\xcc\xa5\x5b\xf0\x1b\x85\xac\x7f\x16\x9b\xd2\x17\xcf\x9f\x0b\x9d\xd7\xa4\xcc\x7f\xf7\xf5\xcd\x7d\xa3\x36\xba\x73\x77\x8b\x3c\x8f\xf4\xce\xf5\x6f\x00\x00\x00\xff\xff\x03\xf0\x4f\xb5\x68\x01\x00\x00")

func mathVdlConfigBytes() ([]byte, error) {
	return bindataRead(
		_mathVdlConfig,
		"math/vdl.config",
	)
}

func mathVdlConfig() (*asset, error) {
	bytes, err := mathVdlConfigBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "math/vdl.config", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _signatureSignatureVdl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\x41\x8f\xd3\x3c\x14\x3c\xd7\xbf\x62\x8e\xdf\x77\x20\x0b\x48\x9c\x51\xd9\xed\xa1\x12\x74\x2b\xb5\x70\x59\xed\xc1\x89\x5f\x12\xb3\x89\x1d\xd9\xce\x4a\x11\xda\xff\xce\xb3\x63\xd2\xd2\xa2\x15\xa2\x27\xfb\x79\x66\xde\xcc\xb4\xbd\xb9\xc1\xad\x1d\x26\xa7\x9b\x36\xe0\xfd\xdb\x77\x1f\x70\x6c\x09\xdf\xa4\x91\x4a\x8f\x3d\xd6\x63\x68\xad\xf3\x05\xd6\x5d\x87\x04\xf2\x70\xe4\xc9\x3d\x93\x2a\x04\x93\xbf\x7a\x82\xad\x11\x5a\xed\xe1\xed\xe8\x2a\x42\x65\x15\x81\xaf\x8d\x7d\x26\x67\x48\xa1\x9c\x20\xf1\xe9\x70\xf7\xc6\x87\xa9\xa3\xc8\xea\x74\x45\x86\x99\xa1\x95\x01\x95\x34\x28\x09\xb5\x1d\x8d\x82\x36\x3c\x24\x7c\xde\xde\x6e\x76\x87\x0d\x6a\xdd\x51\x21\x22\x65\x2f\xab\x27\xd9\x10\xbc\x6e\x8c\x0c\xa3\x23\x28\xaa\xb5\x21\x8f\x30\x0d\x14\x5d\x0d\xd1\x98\x09\xda\x34\xac\x12\xc8\xd5\x92\xcd\x48\xd6\xec\x89\x43\xa8\x13\xd3\x17\x62\xb8\x54\x4b\x3b\xb6\x0b\x4d\x91\xaf\x9c\x2e\xa3\x7a\x7b\xbe\x93\xa3\xb2\xdb\x45\xbe\x10\x71\xf9\x19\xcf\x07\x37\x56\x01\x3f\xc4\x6a\x27\x7b\x02\x7f\x78\xc2\x8e\xc4\x6a\xff\xd4\xec\x65\x68\x97\xfb\x9d\xad\x70\xfe\xbe\xe9\x4b\x52\x1e\x78\x78\x4c\x27\x80\xfd\xec\x2c\xfc\x40\x95\x96\x1d\xac\x53\x14\x81\x85\x58\x7d\x49\x71\x3c\x23\xe7\x53\x44\xde\xc7\xe7\xb9\xea\x9c\xd6\xf0\xfe\x42\xbc\xa4\x5c\xb3\xe2\xeb\x99\x28\x62\x14\xa9\xab\x70\x33\xf9\x9f\x83\xcd\x0e\xb2\xd3\xbf\xac\x35\x67\xc8\x06\x32\xf7\xda\x01\xae\x76\x9e\x26\x5b\xb3\x76\x8d\x8f\x93\x87\x47\x3e\x21\x7d\xbb\xc3\x18\x20\x5d\x33\xf6\xfc\x33\xf1\x62\x75\x3f\x86\x0c\x5a\x30\x3c\xba\x00\x6d\xcd\x21\x38\x92\x3d\xf0\x31\x62\x4e\x42\x7e\x1e\xff\x67\x87\xa0\xad\x91\xdd\xff\x49\x30\x83\x17\x6c\x16\xfc\x03\xf8\x28\x67\x7f\x71\xbb\x34\x13\x4e\x2d\x05\x7e\xc9\xbd\x45\x99\xd7\x4a\xe3\xab\x69\x3a\x5a\x0c\xe7\xca\x22\xed\xa2\xaf\xdf\xaa\xfa\x75\x39\x46\x70\x64\xd8\xf2\x3b\x31\x98\x37\xa6\x51\xfa\x47\x9f\xab\xbe\x88\x9f\x01\x00\x00\xff\xff\x7a\xf5\x7c\x0e\x29\x04\x00\x00")

func signatureSignatureVdlBytes() ([]byte, error) {
	return bindataRead(
		_signatureSignatureVdl,
		"signature/signature.vdl",
	)
}

func signatureSignatureVdl() (*asset, error) {
	bytes, err := signatureSignatureVdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "signature/signature.vdl", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _timeTimeVdl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x58\x5d\x73\xdb\xba\x11\x7d\xb6\x7f\x05\xc6\x7d\xb8\xc9\x54\x96\x68\xf7\x36\xad\x93\x27\xd7\xf9\x68\x3a\x69\x92\xa9\x9d\x76\xa6\x2f\x1a\x88\x84\x44\x8c\x49\x80\x05\x40\xc9\x4a\x27\xff\xbd\x67\x17\x00\x49\x29\xae\xeb\x66\x7c\x1f\x24\x61\x17\xfb\x71\xce\xd9\xc5\x5d\x2c\xc4\x8d\xed\xf6\x4e\x6f\xea\x20\x2e\x8b\x8b\xdf\x8b\xbb\x5a\x89\xbf\x4b\x23\x2b\xdd\xb7\xe2\xba\x0f\xb5\x75\x7e\x2e\xae\x9b\x46\xf0\x21\x2f\x9c\xf2\xca\x6d\x55\x35\x3f\x85\xf1\x37\xaf\x84\x5d\x8b\x50\x6b\x2f\xbc\xed\x5d\xa9\x44\x69\x2b\x25\xf0\x71\x63\xb7\xca\x19\x55\x89\xd5\x5e\x48\xf1\xa7\xdb\xb7\xe7\x3e\xec\x1b\x45\x56\x8d\x2e\x95\x81\x65\xa8\x65\x10\xa5\x34\x62\xa5\xc4\xda\xf6\xa6\x12\xda\xe0\x4b\x25\x3e\x7d\xbc\x79\xf7\xf9\xf6\x9d\x58\xeb\x46\xcd\x4f\xc9\xe4\xab\x2c\xef\xe5\x06\x26\xba\x55\xa2\x52\x6b\x6d\x14\x6e\x0c\xd2\x54\xd2\x55\x88\xa9\xa3\xb0\x4c\x90\x41\x5b\xe3\x29\x24\xb9\xf2\xb6\xe9\x83\x12\x38\x82\xdf\x1b\xfc\xb2\x8d\xe6\x9e\x22\x27\x9f\x94\xea\xb1\x65\xa5\x7c\xe9\xf4\x8a\xc2\x56\x8d\xdd\x09\xe9\xe8\xcc\xbf\x7a\xed\xf0\x55\xb0\xa2\x73\x76\xab\x91\xe0\x0e\x5f\x90\x8f\xd2\xb6\x1d\x4c\x57\xba\xd1\x61\x0f\x9b\xb0\x53\xca\x88\x4a\xaf\xd7\xca\xc1\x2b\x9d\xdf\x38\xd9\xb6\xda\x6c\x84\x32\x5b\xed\xac\x69\xf1\x3d\x4a\x2a\x3e\x28\xa3\x9c\x0c\xf0\xcb\x25\x5b\x5b\x47\x0e\x47\xd3\xe9\x71\x11\xf6\x9d\x2e\x65\xd3\xec\x87\x08\x64\x1f\x6c\x8b\xab\x4b\x98\x1b\x94\xda\x73\xfc\xda\x20\x48\xc3\xc9\x92\xb7\xa3\xf4\x66\x74\x8b\xf0\xba\xed\x1a\xe5\x84\xae\x74\xf2\xd0\x7b\xd4\x76\x7e\xda\x4d\x8a\xcc\x55\xff\x5b\x36\xa7\xf0\x2b\xc4\xea\xb9\x9c\x5c\x45\xea\xb1\xe4\xfc\xd1\x4e\x4e\x23\x58\xc4\x88\xc4\x6e\x6d\x9b\x50\x41\x40\xf0\xbd\xf2\xaf\xc9\xd9\xc5\x4b\xae\xb8\x92\x2e\xd4\xbf\x00\x47\xb0\xa1\x53\xce\xc6\xe0\xc8\x9f\xb1\x01\x21\x6f\xfa\x46\x3a\x86\x17\xfe\xd5\x21\x74\xaf\x17\x0b\x65\xe6\x3b\x7d\xaf\x3b\x55\x69\x39\xb7\x6e\xb3\xa0\x4f\x8b\x4f\x4a\x76\x4b\xaf\x50\x80\xea\x39\xc7\x6f\xac\x75\x95\x36\x14\xec\xf2\x9b\xd1\x54\x34\xd9\x2c\xef\x28\xdb\x67\x58\x7f\x34\x01\x80\xe6\x58\x61\x75\x8d\xea\xeb\x72\x30\xbe\xe4\xe4\x00\x16\x02\xcc\xd8\xc3\x8c\x50\xcf\x85\x67\xdc\x52\x29\xd1\x49\x45\xdf\xfb\x67\x65\x59\x2c\x5f\xec\x51\xb5\x97\xcf\x39\xfb\x97\xbe\xd1\xd2\x2c\xf3\x05\xcf\x31\xf9\x80\x8a\x5b\xf7\xff\x5a\x7d\x75\xb6\x51\x1d\xc0\xb3\x7c\xd4\x3e\xf3\xab\xb5\x3e\x08\xe0\xcd\x3a\x54\x02\x64\xaf\x2d\xb8\x4f\xad\xde\xd5\x0a\xf8\x70\xc4\xa9\x01\xa4\x91\xdd\x32\xe2\xaa\xc7\x67\x02\x91\x2e\xef\x3d\x79\xf3\xda\x94\xc4\x66\xa1\x3a\x5b\xd6\xe2\x05\xfd\xc0\x06\x2f\x67\x02\xb5\x65\x2b\x82\xe8\x82\x9d\x50\xb9\xa5\xe8\x80\x35\x5d\x12\x9c\xc4\x34\xb7\x17\xf9\x43\xb4\x07\x64\xef\xb2\x33\xd1\xca\xce\x13\x81\x7a\xc7\x6c\x43\x78\x35\xfa\xb7\xa3\xb6\x96\x8d\x45\x2c\x33\x8a\xbe\xac\x81\x3c\xfc\x78\xc0\xb1\x19\x93\x03\x66\x76\xe7\x41\x29\x6a\x74\xa4\x9a\xf0\xfd\x2a\x38\x59\x32\xcc\x29\x32\x62\x4d\x1f\x49\x85\x7b\x98\x9a\x8c\x84\x9b\x69\x5c\x74\xcf\x48\xfb\xd6\x22\x02\xaf\x5a\x54\x31\x7f\xa3\x24\x81\x69\xdd\x37\x48\xe0\x66\xa2\x01\x49\x85\xc8\xe1\x50\x24\x0e\xed\x20\x6d\x46\xea\x94\xbe\xd0\x69\x62\xac\x5c\x41\xb8\x13\x6f\xb3\x52\x3e\xcd\x8b\xdb\x2f\xcb\x3f\xbe\x2a\x2e\xfe\xe7\xc1\xdb\xbd\x0f\xaa\x5d\x0e\xea\x42\xdc\x19\xab\x47\xd2\x32\xaa\x76\x67\xa1\x64\x3c\x0e\xe8\xd0\x4e\x87\x5a\xf4\x54\x70\xb4\xc6\xd8\xc8\x78\x28\xa1\x2a\x35\xa5\x3c\x08\x7a\xaa\xda\xe0\x13\x59\x01\x17\x94\x55\xae\x33\x8a\xb3\xa6\x4a\x12\x3c\xd6\xe0\x33\x40\xb2\xd6\x0f\x38\xc7\xa8\x22\x24\xe0\xf0\x77\xe5\xec\x63\x01\x92\xa3\x88\xbe\xa2\x28\x2e\xce\xf9\xef\xae\x28\x5e\xf3\xdf\xbc\xc8\xff\xfe\xc9\x6e\x34\x61\x40\xb1\x11\xf9\xea\x32\x5f\xc4\xc0\x97\xa1\x1f\x6f\xd8\xf3\xd0\x1d\xd7\xd3\x08\x33\x8c\xf5\x07\xa0\x46\xfc\x5a\x14\x82\x14\x40\x94\xfb\x92\x06\x22\xdc\x91\xf2\x89\x58\x07\xcf\x9d\x3c\xf3\xc0\x03\x86\xd4\xd9\x0c\x83\xc3\xf7\x8e\xc0\xc5\xf3\xd5\x58\xd1\x8c\x87\x45\x90\xab\x86\x8b\x64\x54\xa9\xbc\x97\x6e\x4f\xee\x08\x94\x9a\xf4\x0d\xb9\x46\x39\x1e\x87\xa4\x66\xa9\x07\x92\x35\xf1\x08\x2d\xf8\x60\xb9\x2b\x73\x2a\xcf\x4c\xac\xfa\x00\xcc\xef\xe4\xde\xe7\xe9\xfd\xed\xee\x46\x80\x2a\x83\x9b\x8c\x8a\x8d\x6d\xa4\xd9\x30\x1e\xba\xfb\x0d\xf3\x74\xf1\x9b\x24\xa0\xff\xe5\x2a\x1e\x6d\x26\x4a\xae\x38\xeb\x8d\x7e\xe0\xab\xcf\xe2\xb5\x0c\x8b\xb1\x29\x71\x2d\x40\xcb\x03\x79\xe3\x82\x5d\xd0\x98\x61\x99\xa9\x51\xcd\xf8\xd5\xd5\x1f\x8a\xdc\xa0\xc4\xd5\x90\x7a\x1e\x1b\x8e\x6b\xb1\x8f\xa0\x77\x0c\x1e\x49\xbe\x92\x1e\x08\x82\x81\xdb\xcf\x92\x16\x69\x00\xd4\x9b\x5f\x90\xbe\xd8\xca\x46\xc7\xb9\xc8\xa5\x04\x4b\x69\x50\x13\xe3\xd1\x55\x21\xbb\x48\xb2\x81\xe7\x4f\xd2\x04\x93\xe9\x21\x92\x04\xec\x57\x31\xa6\xb7\x09\xbe\xcc\x9b\xfc\xe1\x27\x68\x36\xd0\x2e\x44\x3d\xc1\x7a\x5c\x47\xc2\xce\x46\x42\xf9\xcc\xa8\x19\xd7\x8e\x9c\x3d\xc5\x2a\xbe\x7f\xb8\xcd\x07\xd7\x03\x8c\xff\x3e\x3d\x81\xd9\x6d\x82\xde\x51\x08\x19\x91\x09\x09\x39\x92\xc4\x2b\x87\xf6\x47\x82\xda\x7e\x53\x37\x7b\xf6\xf4\xdb\xc5\xf9\xe5\x55\x21\xb0\x41\x35\x74\x0b\xf5\x08\x1a\x8b\xfe\x6f\x72\xdb\x38\x37\x8f\xb0\x59\xa6\x68\x41\x49\xcb\x45\x1f\x67\x38\x18\x71\x72\x3b\x5c\x1c\x5e\xfd\xca\x7e\x3f\x53\x4a\xc7\xf1\xad\x93\x0a\xc7\x15\x31\x73\x82\x58\x32\x16\x00\xe7\x49\x83\x62\xd8\x7f\xed\x7d\x60\x77\x80\x44\x4a\x0a\xbd\x6f\xa0\xee\xdb\x9c\x4f\xae\x32\x12\xb9\xba\xba\x9a\xa5\xff\xe6\x64\xc5\x96\x1f\x0d\x38\xe8\x5a\x20\xe4\x3b\xc2\x07\x3c\xda\xd9\x28\xfb\xe0\xa6\xf7\x31\x4d\x6b\xd4\x10\x90\x53\x07\x22\xc6\x40\x2f\x0e\x0a\x4f\x7a\x8e\x1b\x39\x4b\xc4\xf9\x9e\xd6\x8b\xc1\xe9\xc4\x15\xa1\x11\x82\x37\x8b\xed\xd1\x1b\x43\x99\xb3\x15\xbb\x6b\x91\x1f\xe0\x1a\x40\x9f\xe4\x99\xa7\x29\xb2\x2d\x90\x41\xac\x21\x6a\xfa\xbb\xcb\xd3\x1f\x8c\xbd\x7f\x60\xef\x7d\xab\x64\xd5\x80\x69\xc7\xc5\xad\xf2\xf7\x3c\x7c\x71\x51\xa7\x62\x40\x33\x9a\xf6\x4e\xf1\xa1\xe1\x4b\xe0\x80\x1c\xaa\x87\x4e\x95\x71\x7b\x14\xb4\x1d\xf9\x3a\x4b\xf4\xd4\x65\x02\x10\x49\x94\xa9\x70\x98\x17\x56\x42\x52\x9c\xf3\x65\xa3\x11\x05\x0f\x3c\x8b\xc4\x89\x91\x43\x30\x51\x48\x27\xb1\x78\xb9\xe7\x02\x61\x27\xa7\x49\xb3\x76\xb6\x15\x67\xc6\xee\xce\xe2\xfc\x46\x3e\x55\x16\xed\xc1\x09\xdc\x12\x5a\xf0\xee\x71\x29\x94\xf8\x81\x42\x78\x24\x83\xc3\x44\x63\x3e\x1c\xdd\x34\xa5\xa4\x7a\x5f\x0c\xb9\x86\x5c\x37\xf4\xdb\x16\xeb\xd1\x2c\x2a\x4b\x5a\x1e\xd0\x9f\x83\x1d\x69\x92\x9a\x3c\x98\x99\x71\xde\xb7\xea\x0d\x55\x3b\x42\x95\x05\x09\xa9\xc1\x5c\x62\x65\x39\x6c\xd3\xec\xb0\x69\xb5\xe4\x54\xe8\xa1\xc3\x1a\xf5\x67\xbb\x53\x94\x21\x3b\xa3\xa2\xd0\x44\x99\x5c\x3e\xbc\x75\x38\x66\x1f\x31\xda\xd9\x40\x72\x0d\xcd\xe3\x65\x49\xf8\x7b\xb5\xa3\x74\xc8\xa1\x26\xd2\xe1\xe7\xb8\xca\x38\xbb\xea\x8f\x53\x3b\x88\x27\x2e\x74\x59\x7c\xc6\x2e\xc5\xfa\xf3\x46\x48\x58\xc8\xcb\x94\x3a\x9a\xf3\x6c\xa0\x03\xbf\x29\xe8\x33\x50\x9f\x9a\xbc\xab\xf1\xaa\xe4\xc3\x50\x3a\x45\x0a\x92\x1f\x72\x83\x97\x61\x0d\x19\x21\x9d\xdd\xd9\x9d\x39\x76\x79\x30\x2a\x0f\x5f\x5c\xa2\xb2\x2a\x8e\x88\x32\x6e\xb4\x84\x57\xac\x02\x9d\xdc\xa4\xdf\x21\xda\xfb\x38\xcd\xe8\x68\x9c\xdd\x09\x2b\xf2\xa8\x22\x3c\x92\xd2\x96\x96\x43\x8f\xaf\x26\x7a\x54\x81\x02\x13\xbd\xcc\x46\x5c\xf8\xcc\x9a\x64\x1b\xd9\x82\x3a\x42\x99\x78\x52\x21\xf7\x96\x04\x1e\xaf\x72\xe9\xb4\x8d\x35\xf8\x29\xca\x44\x58\x0f\x25\x6b\xe2\xbb\x37\xbd\x89\xf9\xf4\xa3\x8d\x6f\xe5\x7d\xdc\x43\xf8\xad\x98\x96\x61\xe4\xca\xc4\xc7\x1b\xa9\xe4\x47\x39\xc6\xa3\xb3\xc0\x67\xae\xe3\x81\xc6\x8c\x1b\x70\xcd\x88\x88\x6f\xdb\xe3\x32\xd3\x09\x04\x33\x5a\x51\xf1\xe8\x52\xde\x82\x87\x76\xc6\x8d\x05\x00\x80\xdc\x0d\xef\xe7\xf8\xaa\x56\x8e\x94\x39\x4f\x31\x06\x56\xfc\x3f\x07\x53\x88\x8c\x3b\xf6\x04\x8f\x71\x4c\x1e\x04\x7d\x30\x2a\xdf\xe3\xe8\x67\xe6\xe0\xe3\x6a\xf9\x04\xd0\xaf\x69\xff\x20\x1f\x1e\x12\x83\xe2\x9e\x97\xd2\x27\xde\x16\xa3\x89\x06\x37\x4b\x19\xb1\x1b\x31\xe3\x54\xc2\x45\xbe\xe4\x8d\xd0\x73\x88\x28\xc9\xda\xc9\x54\xd7\x70\xea\x4c\x1b\x52\xad\xa0\x90\xc9\x49\x8e\x75\xd8\x37\x7e\x9c\xfe\x27\x00\x00\xff\xff\xec\x04\xfd\xa5\x23\x12\x00\x00")

func timeTimeVdlBytes() ([]byte, error) {
	return bindataRead(
		_timeTimeVdl,
		"time/time.vdl",
	)
}

func timeTimeVdl() (*asset, error) {
	bytes, err := timeTimeVdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "time/time.vdl", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _timeVdlConfig = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\x4d\x8b\xdb\x30\x10\x3d\xc7\xbf\x62\xf0\x69\x17\x82\x0c\xed\x2d\xd0\x43\xd9\xc0\x92\x7e\x78\x0b\x49\x5b\xe8\x6d\x6a\x4d\x1c\xb5\xb6\xc6\x95\x65\x2f\x25\xe4\xbf\x57\x23\x25\x21\xd9\x26\x85\xa5\x1b\x88\x62\xcf\xc7\x7b\xf3\x66\x26\x2a\x0a\x58\x6d\x08\x46\xdd\xa8\x8a\xed\xda\xd4\xb0\x36\x0d\xc1\x9a\x1d\xf8\x60\xef\x3d\x5a\x8d\x4e\x83\x37\x2d\x41\x87\xd5\x4f\xac\x09\x52\xe4\xe0\xa8\x07\xa3\x0d\xb7\xe8\x4d\x05\x36\x9c\x23\x65\x01\xd0\xff\xee\xc4\x63\x81\xb0\xda\x40\x4d\x96\x1c\x7a\xd2\xd0\xa0\xad\x87\x90\xaf\xb2\x3d\xd5\x1b\xe1\xf5\xcc\x8d\xba\x8b\x86\x6d\x36\xb9\xe7\x19\x84\x9f\xc9\x57\xe3\x68\xc5\x65\xc4\x5c\x09\x5e\x32\x4f\x02\xfc\x97\xf9\x87\x58\x8e\x9a\x0f\x01\xd7\xb0\x05\xd3\x83\xa3\x2e\x94\x43\x56\x78\xb0\x87\x7b\x3e\x0f\x51\x92\x9b\x1f\xde\xf2\x3d\xd8\xe4\xbd\xb1\x7a\x06\xe1\x53\x0e\xed\x77\x72\xd3\x68\x14\xba\x68\xcc\xcf\x20\xf2\xe4\x5d\xb4\x1d\x3b\x2f\xe5\x6c\x3f\xa1\xdf\xcc\x52\x54\x3e\x85\x12\x5b\x3a\xbc\xed\x76\x29\xf8\x1b\x39\x8e\x50\xdb\x8f\xac\x83\xf7\xb3\x35\xbf\x06\x4a\xce\x74\x9e\xea\x59\x49\x8f\xaf\x6b\x11\x77\xd2\x21\x4f\x7f\x6b\x58\x7a\x37\x54\xfe\xa2\x86\x98\xf0\x02\xf5\xdf\xa1\x65\x6b\x2a\x6c\xa6\xb0\xe8\x93\x33\x57\xe9\xe9\xe6\x36\xbf\xa6\x4b\x66\x39\x27\xd4\x8d\xb1\x97\xf4\x49\x5c\x5a\x9e\xfd\xcc\xf6\xa1\x49\xeb\x69\xf2\xb3\x34\x1f\x93\xae\xe9\x1e\x95\xe1\x62\x7c\xf5\xba\x08\x4b\xe8\x98\x7d\xf1\xbf\x8d\x88\x03\xba\xd4\x0d\x39\xe4\xfb\x0e\x47\xfc\xf7\x7a\x9f\xae\x68\xce\xae\x56\x3f\x58\xa3\xba\xb0\x87\x87\x1d\x78\x1a\x14\xfe\x68\xc7\x61\x1f\x68\x97\x8f\x66\xed\x9f\xc1\x5b\x2e\x05\x62\x11\xe6\xe3\x46\x6c\x9e\xf0\x95\x4b\xe1\x38\xc3\x97\x3b\xe4\x61\xfe\x70\xe3\x59\xeb\xc7\xdb\x19\xbc\xd5\xfa\x38\xd0\xc0\x01\xfd\xd0\x49\xef\xe3\xa5\x22\x2d\xe8\x2b\x67\x3a\xaf\xb2\x5d\xf6\x27\x00\x00\xff\xff\x46\x95\x86\x66\x7e\x04\x00\x00")

func timeVdlConfigBytes() ([]byte, error) {
	return bindataRead(
		_timeVdlConfig,
		"time/vdl.config",
	)
}

func timeVdlConfig() (*asset, error) {
	bytes, err := timeVdlConfigBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "time/vdl.config", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vdltoolConfigVdl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x5a\x5b\x6f\x1b\x37\x16\x7e\xb6\x7e\x05\x61\x2c\xb0\x4e\xa1\x4a\xbb\xdd\xcb\x83\x17\xc5\xc2\xb5\x1d\xc3\xdd\xd4\xc9\x46\x76\x13\x6c\xd0\x07\xce\x0c\x25\x31\x9e\x19\xaa\x24\x47\xb2\x1a\xe4\xbf\xef\xb9\x90\x1c\x8e\x2c\xdb\x4d\xdb\xa7\x24\x0f\x8e\x34\x43\x1e\x9e\xcb\x77\xbe\x73\x48\x6a\x3a\x15\xa7\x66\xb5\xb5\x7a\xb1\xf4\xe2\x9b\xbf\xfc\xf5\x1f\xe2\x7a\xa9\xc4\x8f\xb2\x95\x95\xee\x1a\x71\xd2\xf9\xa5\xb1\x6e\x22\x4e\xea\x5a\xd0\x20\x27\xac\x72\xca\xae\x55\x35\x19\xc1\xe4\x1b\xa7\x84\x99\x0b\xbf\xd4\x4e\x38\xd3\xd9\x52\x89\xd2\x54\x4a\xc0\xd7\x85\x59\x2b\xdb\xaa\x4a\x14\x5b\x21\xc5\x77\xb3\xb3\xaf\x9d\xdf\xd6\x0a\x67\xd5\xba\x54\x2d\xcc\xf4\x4b\xe9\x45\x29\x5b\x51\x28\x31\x37\x5d\x5b\x09\xdd\xc2\x43\x25\x5e\x5c\x9e\x9e\x5f\xcd\xce\xc5\x5c\xd7\x6a\x32\xc2\x29\xaf\x64\x79\x2b\x17\x4a\xac\xab\xda\x1b\x53\x8b\x4a\xcd\x75\xab\x9c\xf0\xdb\x15\xfc\xed\x1c\xaf\x83\x73\x61\x84\xa0\x21\xda\x3b\x55\xcf\xc7\x20\xb3\xac\xbb\x4a\xb7\x0b\x7c\x8d\xb2\xe6\xc6\x36\xb0\x30\xe8\x0d\x63\x27\xa5\x69\xe7\x7a\x41\x4b\xb9\xc9\x68\x35\x5c\x87\xd6\x3e\xe5\x11\x6e\xa5\x4a\x3d\xd7\xb8\xe8\x12\xcd\xc4\x87\x9d\x95\x5e\x9b\x16\x45\x0e\x16\x9f\x08\x70\x24\x38\x41\x93\x86\xba\x94\x75\xbd\x45\x51\x56\xad\xd0\x7f\xad\x57\x64\xab\x59\xe1\x6c\x59\x8b\xc3\x5e\x93\x43\x56\x05\x5f\x2b\x59\x2e\x49\x64\x70\x6d\x50\x0e\x84\x9f\xc3\x1b\x94\xb7\x63\x80\xd0\xcd\x0a\x9c\xab\x7d\xbd\xc5\x8f\xc6\x7a\xc7\xb1\xc9\x66\x4e\x16\x13\xb1\x35\x9d\x68\xe4\x16\xd4\x99\x2b\x50\xdc\x04\x51\xa4\x79\xb0\x36\x44\xe2\x9e\x62\x62\xa3\x01\x13\x9d\x17\xea\x6e\x67\x29\xf4\x70\x14\x32\xc2\xb8\x24\xc7\x79\xdb\x95\x5e\x7c\x18\x1d\xc0\x32\x17\xaa\x7d\x21\xdb\x45\x07\xda\x10\x96\xbc\xd5\xa5\x67\x97\x3a\x45\x41\x21\x00\x2d\x54\xab\x82\x6f\xeb\x38\x1c\xb4\xbf\x9c\xa7\x91\xda\x91\x3c\xd5\xac\xfc\x76\x2c\xc0\xc1\xc2\x75\x2b\xd4\x03\x5c\x9b\xa6\x08\x69\x15\xbe\x33\x1b\x78\xea\x0d\x02\x2d\x48\x46\x00\x1f\x0c\x94\x01\xa1\xef\xb2\x07\x3f\x8d\x48\x7e\xfc\xfa\x75\x08\x7f\x39\x8c\xbc\x43\x29\x46\xc4\x7f\x17\x86\x6d\x1e\x1d\x7c\x2f\xd7\x32\x3c\xc4\x8f\xf9\x63\x57\x5a\xbd\xf2\xa2\xff\x18\x5f\xce\x36\x7a\xee\x79\x0e\x7d\x0c\xcf\x3f\x12\x0a\x33\xd5\x84\x6a\xbb\x86\x8c\x60\xc7\xdd\xb6\x66\xd3\x3e\xe2\x37\x8e\xc6\xae\x00\x0c\xc8\x85\x61\x95\x72\xc5\x82\x1e\x71\x59\x73\x0f\xfe\x0b\x23\xf6\x3b\x23\x2e\x64\xf6\x05\xfe\x8d\xb6\xea\xda\x5c\xc1\xc0\xb5\xba\xa6\xac\x1d\x26\x54\x23\x57\x2b\x84\xd0\xdc\x9a\x06\x08\xe3\xc7\xb3\x17\x00\x35\x88\x1e\x89\x84\xd0\x41\x3a\x0b\x54\x17\x44\xb5\x24\x84\xdf\xa4\x8c\x62\x0d\xfa\xbc\xb3\x10\x7a\xc0\x26\x51\x83\x04\x6a\x51\xd2\x6d\x51\x4e\xd4\x58\x91\x28\x20\x07\x63\xad\x2a\x01\xc6\xff\x82\xb1\xe8\x1b\x98\x5b\x76\xf0\xac\x45\x68\x47\x84\x32\x7c\x42\x1a\x39\x66\x2e\xc4\x56\x9f\x05\x24\x6e\xb3\xd4\x5e\xd5\xda\x31\xbc\xa6\xd3\x07\x2c\xc7\x99\x8d\x92\xad\x27\xe2\x70\xc0\x86\xd2\x6a\xe3\x60\xba\x42\x8b\xe9\x2f\xa8\x01\xcc\xa8\x2b\x6d\x80\xac\xc0\xd3\xc1\x76\x32\x9a\x8c\x82\x0c\x85\x3c\xb6\x14\xf7\xb1\x28\x20\x25\x31\xad\x5b\x85\xf6\x82\xeb\xc1\x66\x69\x2b\xf2\xe3\xd0\x47\xb4\x26\xba\x96\xe4\x95\xa6\x59\xc1\xe3\x42\xd7\xda\x6f\x23\x3f\x60\x3c\x70\xa2\xd7\x4d\x62\x9d\x44\xb9\x67\x91\xf4\xd0\xab\xd7\x38\x02\x04\xb2\xf5\x18\xaf\x81\x40\xd6\x6b\x03\xac\x81\xc6\xa2\xd8\x94\x7e\x8c\x57\x70\x6a\x47\x95\x40\xf5\x2a\x47\x4b\xb3\xc5\x7b\x67\x62\x7d\xba\x55\x5b\xae\x3a\x84\x1a\x22\x59\xf8\xd8\xca\x46\xc5\xc7\xa4\x3c\x7a\xea\x48\xde\x4a\x0e\x00\x7c\x7b\x36\xc6\x00\x01\xa7\x36\x9d\xf3\x24\xad\x88\x66\xa5\xd2\x83\x74\x1b\x2d\x96\xce\x99\x52\x93\xb2\xc8\x7a\xf1\x7d\xce\xb8\x43\xc5\xf2\x14\x04\x27\xc3\x7c\xc8\x53\xd6\x4e\xdd\x21\x2a\xda\x92\x54\x94\xb0\x84\xb6\x4c\x76\x2d\x94\x49\x47\x61\xe9\xda\x92\x28\x25\xc4\x05\x5f\x10\xa9\x16\xca\x6f\x94\x62\xf5\xc8\xc5\xe8\xf8\x2c\x07\xdc\x98\x9e\x6c\x34\x70\xa0\xec\x3c\xe3\x05\x2b\x8e\xc0\xbf\x24\x0d\xb5\x7e\xad\x16\xa8\x82\x65\x20\x26\x8b\xc0\xfb\x71\x65\x72\xa1\xeb\x4d\x3a\x41\xfd\x43\xe1\x1c\xb8\x37\x3a\x14\x5d\xff\xdc\x98\x63\x1a\x2d\x48\x0e\x80\xc1\x44\xb0\x1f\xdd\xe1\xdb\xb1\x68\xc5\x57\xfc\xe0\x99\x50\xd6\x06\xb0\xf4\xc3\x9f\x43\xc6\xa7\x09\x5f\x85\x19\x3b\x13\xee\xe7\x10\x44\xfe\x1d\xe6\x66\xbb\xf8\xe9\xc2\xe0\xa3\x44\x58\xa4\x59\xa5\x90\xcd\x8a\xe0\x7d\xa0\x68\x82\x83\x6e\xb9\xf8\x13\x7c\x77\xa2\x2b\x13\x6a\xa8\xb5\x99\x29\x08\xf6\x44\x9b\xe9\xdd\x14\xca\xe4\xb4\xd6\xc5\x14\xbc\x38\x05\xce\xf5\x95\xf4\x72\x1a\x02\x80\xc9\xa4\xee\x24\xd4\xdd\x9e\x68\x59\x83\x01\xfb\xfd\x47\x63\x77\x13\xe8\x1a\x3f\x83\x27\xaf\x69\xa9\x03\x7a\x75\x61\xf0\x3f\x06\x52\x70\x6c\xae\x77\x48\x13\xc0\xe8\x30\x7f\xb0\xbb\x81\x98\xca\x6a\x17\xf8\x7d\x10\x43\xc9\x8c\x92\xac\xfa\xb9\x03\x5f\x02\xb9\x54\x95\x0e\xdd\x47\x68\x15\xc6\x81\x91\xb9\x8f\x0a\x44\x13\xfb\x26\x14\x95\x32\xb4\x36\x80\xac\x94\x26\x84\x0c\xe4\x2c\xc6\xa1\xac\x9d\x19\x48\x62\xf1\x69\xb8\x6e\x59\x2f\x5e\x34\x51\xce\x70\x4e\x4e\xf1\x48\x04\x93\xc8\x3a\x11\x6b\xe8\xae\x63\xac\x94\x57\x5d\x53\xa8\x08\x29\xf4\x1d\x3d\x3d\x1c\x4c\x3a\x1c\x87\xf7\x61\xd1\x63\xf1\xe1\xc3\x2b\xe9\x97\xc7\x3c\xee\x70\x0c\x70\x6b\x54\xfc\xf6\xf1\x63\x1c\xfe\x3f\x65\x0d\x89\xfb\xf0\x03\x38\xfb\x58\xdc\xb4\xfa\xe7\x4e\x7d\x1c\x1d\xc4\x00\x83\x77\x72\x63\x88\xd8\x83\xb7\x63\xff\x15\x1c\x9e\xfa\x53\x9c\x3a\x4e\x0c\x3e\x0c\x27\x44\x2d\x4a\x7a\x07\xa8\xe6\xcf\xb4\x00\x6a\x92\x15\x4c\xfa\x04\x21\x28\xd4\x52\xae\x35\x40\x10\x61\xf8\x0b\x8e\x59\xcb\xba\xa3\x06\x89\xa8\x03\xe3\x51\x2e\x55\x09\x90\x5b\x80\x70\x92\x72\x61\xf0\xbf\x94\x2d\x04\xbf\x61\xb6\x44\x80\x06\xcc\x24\x58\xd3\xd0\xd8\x3a\xcc\x08\xde\xa3\x83\xef\xb0\x4d\x3e\x88\x51\x98\x05\x9f\x9c\x58\x2b\xa1\x24\xce\xb0\xd5\x1f\x1d\xfc\x20\x57\xa3\x83\x57\x46\x43\xdf\xcb\x91\xba\x7e\x79\xf6\xf2\xc8\x9b\xaa\xda\x3c\x3b\x16\x33\xee\xdb\x62\xcc\x69\xd4\x5c\x96\xca\xfd\x3b\x29\xc9\x9e\xc8\xd4\x4c\x0e\xce\x53\x3a\x29\x1a\x86\x0f\x32\x10\xe3\x1d\xd3\x2a\x42\x71\x25\x89\x03\xa1\x96\x77\x14\x59\xec\x65\x2b\xa8\x93\x7d\x57\xc2\x8b\x40\x78\xfa\x42\x44\x82\xb2\xd8\x23\x74\xf6\xd5\xa0\x84\xf7\x28\x91\x20\x80\xb3\x01\xf3\x67\x1d\x25\xf5\x85\x61\xa2\x6f\x89\xf8\x19\x76\xda\x0f\xb6\x0d\xe2\x3d\x14\x2b\x12\x58\x48\xa8\xe1\x61\x01\x14\xc3\xc5\x15\xfb\xf8\x82\x5b\x62\xcc\x21\xd3\x28\xe8\xf8\x83\x6a\x95\x9e\x43\x83\x0f\xd2\x85\x9e\xef\x35\x47\x54\x46\xb9\xf6\xcf\xd8\x87\x60\x93\xbc\xa3\x0f\x18\x4b\xc6\x05\x63\x63\x30\x08\x45\x83\x50\x64\xc8\x8b\x88\x44\x0e\xc5\xe1\xaf\xcf\xff\x7b\x73\xf9\xfa\xfc\xec\x58\x9c\x6b\x6c\x6c\x04\x66\x92\xf8\xf6\xdb\x90\x4b\x02\x70\x7b\xe9\x48\xa2\xa6\x16\x1c\x7c\xf3\x46\x71\x2d\x6b\x0d\x90\x06\x40\x01\xc2\x3b\x4a\x5d\xdf\x34\x34\x19\xb1\x6c\x3a\xea\xe9\x10\xe2\x99\x16\x3d\x13\x73\xde\x24\x1c\xd0\xe2\xfc\x14\x3f\x72\xee\xba\x9d\xe4\x92\x50\xd9\x55\x5d\x8d\x05\xba\xd2\x54\xa8\x62\xaa\x90\x04\x16\xab\x7c\x67\x71\x65\x0b\x06\x80\x8b\x13\x41\x06\x04\xb3\x27\x52\xd7\xe5\x12\x33\xf7\x1a\x0e\xf8\x39\x6a\xe0\x25\x66\x7e\x28\x46\x15\x58\x7f\x34\x81\x96\x85\xf1\xc0\xbd\x44\xdc\xc4\x04\x15\x51\x35\x56\x72\x92\x7a\x10\x7e\x41\xe6\xa7\xb7\x41\x63\x7e\x8a\xcd\x0f\xca\x28\x78\x9f\x7a\x31\x6c\xc8\xd0\xf3\x24\x0b\xda\xf1\x9a\x19\x2b\xa9\x67\x69\x8b\xa7\x6c\x01\x66\x36\xa8\x09\x74\xbf\x60\x22\x87\x21\xa0\x3e\x23\xee\x7d\x36\x46\xbc\x91\xad\x99\xa9\x8f\x59\x1a\x7d\xbf\x59\x1a\xc7\x7e\xfe\x34\x73\xa8\x0d\x22\xe5\x7a\x51\xb4\x70\xb0\xd4\x94\xdc\xef\x03\xe3\xa0\x15\x6f\xdf\xbe\xc5\xd8\xd5\xc0\x40\xb1\x33\x68\x33\x4b\x1f\x35\x74\xc0\x6b\xd7\xf9\x8a\x88\x62\x9c\x9a\xc0\x8e\x5e\xc0\xed\xba\x6a\xe8\x38\x20\xab\x9d\xd8\xc2\x73\x13\x28\x21\x05\x30\x67\x42\x4c\x8d\x58\xc2\xa0\x3a\xa6\xb1\x8b\xb5\x30\x08\xdd\x57\x3a\x06\xb1\xbb\xaf\xe2\x49\x85\x9b\x85\x99\xf2\x34\x2a\xc0\x9e\x7b\x64\x88\x44\xad\x6f\x63\xf4\x99\x6b\x22\xf0\x25\x97\xf0\xcc\x29\x94\x17\x90\xbe\x8c\xf5\x0c\xe7\xa0\x9c\xcf\xcc\x82\xd5\x7a\xf5\x92\x5d\x24\x6d\x68\xdb\x64\x40\x37\x94\xb7\xc3\x22\x05\xbb\x3b\xde\x81\x2f\xf5\x6a\xd0\x22\x0f\xf9\x68\x10\xad\x51\xd8\x40\xb1\xb3\xf7\x65\x65\x46\x1b\xb4\x66\xac\x76\x78\xd6\xd5\xf2\x3e\x3b\xdf\xb6\x4a\xbf\x67\x49\xc4\x38\x91\x57\x62\x80\xc1\x52\xdc\x90\x07\xe7\xa4\x53\xa2\x25\xed\x44\x64\x57\xfb\x2c\x6d\x5a\x3e\xbd\x69\x08\xc9\xfd\xde\x49\x38\xe4\xea\xb8\x7f\x0a\x14\xf9\x20\x3f\xc6\xdd\x59\xec\x09\x72\x9a\x04\x05\xf6\xb0\x27\xae\x15\xda\x96\xde\x46\xc6\x01\x49\x7a\x98\x1e\x03\x87\x8f\x0e\x82\xb3\x68\xf8\x29\xc0\xb8\xc5\x62\xf6\xb4\xeb\xee\xe7\x56\x70\x0e\xef\x89\x92\xa0\x9d\x1d\x6d\xd6\xfc\xf6\xc2\x18\xb2\x9a\x43\x81\xaf\x4d\x0b\x84\x36\x8c\xc0\xfd\xe3\x83\xb4\xff\x1d\x6a\x15\x35\xc2\xae\x92\xf6\xbc\xbf\x52\xad\x5f\x52\xee\x70\xc6\x0f\xf4\x10\xd0\x5d\x97\xbc\x92\xa3\x1c\x09\x59\x05\xdb\x34\x6a\xe8\xfa\xf2\x8d\x0d\x77\x2c\xc9\x4f\x81\x63\x2e\x61\x3a\x81\x83\x0e\x17\x9e\x8e\x7a\xef\x8e\xc7\x22\xff\x6b\xa2\x9e\x22\x1d\x12\x86\x2a\xfc\x6f\x0f\x3a\x79\x2a\x10\xfe\xd3\xf1\x7e\x38\x78\x83\xbd\x00\xc9\x3b\x62\x8e\x5b\x4a\x47\x5d\x55\xdc\x1c\xfa\x7f\xfe\xfd\x59\xbe\xf6\x03\x01\xfd\xe4\x20\x14\x06\x5b\xc6\x3d\x3d\xf9\xae\xff\x29\x73\xd0\x6b\x81\xfd\xfa\xa3\xc3\xcc\x8d\xef\xf1\x68\xf1\xd1\x13\xb8\x7c\xda\x1f\x7d\x06\xc7\xc7\x85\x5f\xea\x29\x5c\xb2\xfe\xf3\x3d\x87\x33\x76\x31\x79\x6f\x2a\x39\xc1\x55\x49\xde\x67\x72\x08\x17\xc8\x67\xdf\x91\x1b\xf7\x89\x05\x02\x99\x0f\xc9\x54\x15\x80\xc9\xd9\x96\x29\x34\x79\xe2\x30\x2a\xb4\x5a\x3d\xe6\x60\xc0\x6b\xda\xb7\x7d\x52\xb6\x91\x23\x39\xe5\x48\x14\xdd\x27\xe4\x29\x87\x03\x76\xd0\x9d\xad\xf4\x24\xb6\xfb\xe5\x70\x42\x3c\x72\x9c\xd7\x74\x1b\x03\x66\xe3\x46\x12\xd0\xb2\xed\x9d\x9e\x6b\x40\x8b\x8c\x85\x02\x90\x8e\x83\xea\x97\xd0\xc5\x2e\xc2\xf6\x9e\x14\x2c\x4d\x07\xcd\x5d\xc1\x83\x21\x01\xc2\x00\x52\xa8\xac\x41\x1d\xc0\x79\xaf\xff\x52\xe1\x0d\x46\x53\x10\x0c\x68\xcd\x7b\x5e\x1e\x33\xd7\xce\x95\x84\x2e\x34\x5c\xeb\xe0\xdd\x1f\xba\x49\x7a\x8f\x57\x62\xd4\x9d\xa4\xa0\xc2\xf3\x88\x41\x2c\x7a\xcf\xfb\xa3\xba\x31\xe6\x02\xdf\x80\x91\xa3\x4f\x4a\x68\xfe\xdd\x0b\x2c\xbc\x71\x46\x88\x3f\xc8\x40\x45\xb2\x01\xa1\x75\x6c\x09\xf6\x3b\xef\x60\x34\x52\xc1\x6e\xb0\xf8\x34\xb5\x1f\xc6\x09\x10\x13\xb9\x57\xd8\xb4\x29\x34\x78\xc6\xf8\x5b\x12\x2e\x21\x28\x4f\x33\x88\x02\x9f\x07\xfc\xde\x34\xdb\x85\xd9\x7d\xdc\xf7\x75\x2b\xbf\xdb\xda\xa9\x5e\xe1\x06\xec\xc9\x1a\x36\x14\x11\x2b\x19\x2f\x91\xdd\x90\xe5\x07\x63\x74\x85\xb6\x5f\x70\xb8\xa5\xcd\xae\x43\x51\xce\x95\xf1\xa1\xd7\x81\xfd\xc5\x4a\x7b\xce\x8e\x5c\x3a\xdf\xd2\x52\x23\x0d\x58\x5f\x40\x50\xdb\xe4\xb6\xd0\x48\xb8\x65\x00\x3b\xdd\xa4\x2a\xbb\xa5\xc2\x87\x7e\x5e\x59\x09\xc1\x2d\x89\xf6\xf0\x24\x85\x0e\x36\x14\x17\x32\x6c\x91\xef\x05\x91\x2d\x80\x86\xa2\xe3\xcb\xf2\xc0\x5b\xbd\xd6\xbc\xf3\x51\x04\x6f\xac\x86\xd6\xc0\x0e\x00\x6b\x6c\x5b\x01\xad\x61\xbd\x89\x47\x63\x05\xde\xa9\x43\xa2\xc1\xd8\xb0\x41\x42\x81\x30\xe1\x3d\x94\x5a\x37\x9d\x5b\x58\x76\x63\xec\xad\x9b\xf2\x72\x2e\x1e\xda\xbc\xc1\x4d\x0c\x2f\x5a\x6c\x19\x35\xd4\xb9\xf0\x0d\x30\x6e\xa7\xc1\x80\x43\xd2\x94\x67\x1e\x86\xbb\x7c\x03\xb5\x4d\x83\xab\xbe\xbf\x99\x5d\xc7\xeb\xf6\xdc\xb8\x59\x66\x9c\x08\x7c\x6c\x8d\xa1\x7b\x5f\x2a\xa2\x08\xe1\x58\xdf\x27\x83\x8c\x45\x51\x9a\x46\x41\xdb\xb6\xee\x39\x8c\x6f\xcb\x31\x36\x24\xe0\x2d\xd1\x7f\xb0\x71\xea\xa5\x5d\x28\x1f\x35\xbe\x29\x94\x05\x91\xa7\x12\x38\x15\x77\xd7\x5e\x4c\x6f\x1c\xf8\x6c\x2a\xa5\x35\xed\xb4\x2b\x68\x27\x05\x9d\xb6\xc3\xa3\xfe\x31\x67\x79\x69\x81\x6f\xc8\x90\xc7\x06\x4f\x81\xbb\x26\xf9\xb3\xcc\x39\xc4\x18\xa4\xb3\xf6\x7c\x9a\x17\xfc\x84\x42\x0f\x33\xa5\x0e\x27\xe8\x79\x5a\x95\xea\x91\xc7\xa5\x61\xf6\x36\x33\x34\x4c\x8d\x49\x8c\xb0\xd5\xd8\x39\x19\xc0\x1c\xac\x43\x76\xc1\x04\xd7\x15\xf1\x39\x1d\x69\x21\x7a\xac\x8f\x61\xc8\xd6\x1c\x84\x64\x2c\x60\xff\xa9\xa1\xd3\xc1\x9e\xad\xe6\x7e\x95\x31\xc8\xfd\x04\x68\x60\x1f\xba\x6b\x22\xc0\x28\xfe\x95\x88\xa0\x23\x48\x0c\x17\x9d\x5f\xb8\x08\x48\x8a\x50\x9e\x7c\x97\x1c\xd0\xca\xe0\xde\x28\x42\x8e\x1c\xd0\x9a\xc2\x54\xdb\x54\x97\x65\x51\x73\x39\xcc\x04\x25\x6e\x47\x57\x90\x19\x63\x36\xbf\xe2\x1f\x38\x90\x26\x0a\xbe\x75\x2b\xe6\x34\x75\xe7\xad\x84\xbd\x14\xa8\x06\x05\xc8\xb9\xe9\xea\x76\xc1\xa5\x4c\x1c\x9d\x9a\x26\x0b\xdf\x4c\xd9\x35\x24\xad\x7b\x65\xcd\x5c\x7b\x82\x5e\x7f\xcf\x32\x7c\x49\xd0\xfb\x13\xe8\xf2\xfa\xe5\xcb\xeb\xfb\x30\x08\x63\xa7\x2b\x1a\xfc\x2c\x66\x58\x30\x1c\x49\x82\x00\x86\xee\x6a\xd0\xf9\x00\xf5\x41\x48\x86\xb4\x13\xd7\x61\x1f\x11\x96\xe8\xa0\x0a\x45\xe6\x88\x63\xac\x84\x04\xab\xcd\x82\x76\xa7\x3d\x1d\x84\xcc\xea\x4b\x61\x74\x3d\xae\x45\x97\x5e\xeb\x6f\xfe\x46\x81\xa7\x33\x38\x68\xaf\xf9\xd4\xd8\x73\x8c\xe3\xef\x83\x4e\x0d\xa8\x9f\x58\x84\x7b\x4c\x85\xec\xc7\xe1\x27\x42\x4a\xe2\x92\x2b\xf8\x17\x40\x6b\x0e\x5b\x2e\x2e\xfa\xb5\x17\x19\x4a\xc1\x80\xea\xff\xe8\xfd\x0c\xff\xd8\xe1\x8b\xdd\xd0\xf4\xe6\x7f\x9e\x3b\x9a\xab\xd9\x19\xee\xc7\x41\xb1\xab\x19\x8a\xc4\x2e\xd4\xae\xf9\x84\xe2\x8b\xdb\xcb\x30\xb1\x7c\xfa\x66\xe6\xe3\xe8\xff\x01\x00\x00\xff\xff\xda\x50\x5e\xb2\x29\x28\x00\x00")

func vdltoolConfigVdlBytes() ([]byte, error) {
	return bindataRead(
		_vdltoolConfigVdl,
		"vdltool/config.vdl",
	)
}

func vdltoolConfigVdl() (*asset, error) {
	bytes, err := vdltoolConfigVdlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vdltool/config.vdl", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
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
	"math/math.vdl":           mathMathVdl,
	"math/vdl.config":         mathVdlConfig,
	"signature/signature.vdl": signatureSignatureVdl,
	"time/time.vdl":           timeTimeVdl,
	"time/vdl.config":         timeVdlConfig,
	"vdltool/config.vdl":      vdltoolConfigVdl,
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
	"math": &bintree{nil, map[string]*bintree{
		"math.vdl":   &bintree{mathMathVdl, map[string]*bintree{}},
		"vdl.config": &bintree{mathVdlConfig, map[string]*bintree{}},
	}},
	"signature": &bintree{nil, map[string]*bintree{
		"signature.vdl": &bintree{signatureSignatureVdl, map[string]*bintree{}},
	}},
	"time": &bintree{nil, map[string]*bintree{
		"time.vdl":   &bintree{timeTimeVdl, map[string]*bintree{}},
		"vdl.config": &bintree{timeVdlConfig, map[string]*bintree{}},
	}},
	"vdltool": &bintree{nil, map[string]*bintree{
		"config.vdl": &bintree{vdltoolConfigVdl, map[string]*bintree{}},
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
