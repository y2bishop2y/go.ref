// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by go-bindata.
// sources:
// assets/alltrace.html
// assets/blessings.html
// assets/chrome.html
// assets/glob.html
// assets/logs.html
// assets/no_syncbase.html
// assets/profiles.html
// assets/resolve.html
// assets/stats.html
// assets/syncbase.html
// assets/vtrace.html
// DO NOT EDIT!

package browseserver

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

var _alltraceHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\x4f\x6f\xdb\x3e\x0c\xbd\xe7\x53\x10\x46\xf1\x73\xda\xfe\x62\xb7\x05\xb6\x43\x62\x7b\x68\xbb\x0e\x28\x50\xf4\xd2\x6e\x97\x61\x07\xc5\xa2\x63\xad\x8a\x14\x48\x72\xb6\xc0\xf3\x77\x1f\x25\x3b\x7f\x77\xd9\x7c\x68\x29\x9a\x7c\x7a\x7c\x8f\x4e\xdb\xa6\x17\x23\xa0\xe7\x5e\xaf\x36\x46\x2c\x6a\x07\x37\x57\xd7\xef\xe1\xb5\x46\xf8\xc2\x14\xe3\xa2\x59\xc2\x6d\xe3\x6a\x6d\x6c\x02\xb7\x52\x42\x28\xb2\x60\xd0\xa2\x59\x23\x4f\x42\xf7\x67\x8b\xa0\x2b\x70\xb5\xb0\x60\x75\x63\x4a\x84\x52\x73\x04\x3a\x2e\xf4\x1a\x8d\x42\x0e\xf3\x0d\x30\xb8\x7b\xf9\x38\xb1\x6e\x23\x31\xb4\x49\x51\xa2\xa2\x56\x57\x33\x07\x25\x53\x30\x47\xa8\x74\xa3\x38\x08\x45\x49\x84\xa7\xc7\xfb\x87\xe7\x97\x07\xa8\x84\xc4\x64\x74\x91\x76\xdd\x68\xd4\xb6\x1c\x2b\xa1\x10\xa2\x52\x2b\x87\xca\x45\x94\xcd\x6c\x69\xc4\xca\x81\x64\x6a\xd1\xb0\x05\xe6\xf1\x77\xb6\x66\x7d\x32\x2e\xc2\x65\x55\xa3\x4a\x27\xb4\x02\x5b\xeb\x1f\x8f\xdc\x7e\xd2\xe6\x89\xcd\x51\x8e\xa5\xff\x7b\x0e\x6d\xa8\xf2\xcf\x9a\x19\x50\x44\xdf\x42\x0e\x5c\x97\xcd\x92\x2e\x49\x16\xe8\x1e\x24\xfa\xd0\xde\x6d\xee\x25\xb3\xf6\x99\x2d\x71\x1c\x87\xee\x89\xa7\xc2\x88\x94\x89\xcf\x67\x3b\x9c\x4a\x1b\x18\x7b\x30\x41\x40\x57\x33\xfa\x97\xf5\xb8\x89\x44\xb5\x70\x35\x65\x2e\x2f\x0f\x2f\x3e\xbc\x9c\x5a\x42\xed\x57\xf1\x6d\x76\x54\x20\x2a\x18\xfb\x37\x09\x67\x8e\x59\x74\x49\x60\x00\x79\x9e\xc3\x1f\x93\x6c\x9f\x50\x1f\x74\x4f\xb8\xb0\x2b\xc9\x36\x04\x1f\xcf\xa5\x2e\xdf\xe2\x63\xf4\x0e\x50\x92\x23\x7f\x0b\xa1\xb4\xc2\x53\x84\xd1\x71\x44\xe6\xa4\xbd\x11\x05\xd9\x84\xbd\x07\xa5\x17\x30\x8f\x86\xe3\x84\xd6\xc0\xa1\x81\x25\x97\x93\x85\x11\x3c\xea\x1d\x6b\xdb\xb3\x79\x53\xbe\x21\xad\xdb\x34\x87\xe4\xae\x8f\xbb\x1e\x36\x5b\x15\x6d\x6b\xc8\x6e\x84\xb3\x5e\x01\x5f\xf3\xa2\x8d\x43\x1e\x7c\xa5\x42\x8f\x20\x78\xe8\x16\x8a\xe3\x4f\xd8\xe1\xf5\x2d\xa1\x82\xf4\x24\x3b\x20\x14\xfe\x02\xe9\xe0\xaa\xeb\x32\x06\xb5\xc1\x2a\x8f\xf6\x5b\x34\x3d\x5d\x9b\x98\xc0\x07\x94\xf8\x3c\x2a\xf6\xa7\x2c\x65\x74\xf2\x32\x76\xdd\x3e\x4b\x19\xc5\xbb\xbd\x38\x87\xe7\x2c\x5d\x15\x5e\xa5\x5e\x8d\x7f\x95\x29\xab\xdf\x15\xaf\x86\x95\x68\xb3\x94\x42\x3f\xb4\xa2\xd5\xec\xf5\xf0\x1f\xa9\xf1\x9b\x1a\x66\xdd\xea\xf5\x3f\x6c\x75\x39\x55\x95\x8b\x35\xf8\xc5\x9a\x04\xda\x79\xe4\x07\xe8\xba\x08\x82\xf3\x79\x34\x58\x3f\xf5\xbe\xcf\x56\xda\x0a\xcf\x6b\x6a\x50\x32\x27\xd6\x18\x6d\x09\x7b\x86\x25\xd2\xaf\xc5\x36\x98\x4c\xae\x6f\xe8\x23\x91\x70\xf2\xb9\x0c\x23\x0c\x6e\xcb\x03\x7d\xb2\x46\x16\x47\x7b\xb5\x63\x2f\xb8\x27\xee\x07\x38\x28\x0f\x2d\x52\x14\x3b\xe7\xd2\xb5\xf3\x9a\x7c\x50\xf9\xa0\x47\xd7\xfd\xe7\xf2\xb0\x10\xc9\x23\x09\x1f\x1c\x1b\x62\xef\x58\x96\x52\xf7\xc9\x85\xc7\x8e\x65\xe9\x96\x52\x96\x92\x4c\xc5\xf6\xfd\x81\x71\x43\xea\x77\x00\x00\x00\xff\xff\xbf\xe1\x74\x56\x56\x05\x00\x00")

func alltraceHtmlBytes() ([]byte, error) {
	return bindataRead(
		_alltraceHtml,
		"alltrace.html",
	)
}

func alltraceHtml() (*asset, error) {
	bytes, err := alltraceHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "alltrace.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _blessingsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x4f\x6f\xeb\x36\x0c\xbf\xe7\x53\x10\xd9\x0e\x5d\x31\xdb\x6b\x81\xf5\x50\xb8\x06\xda\x34\x87\x62\xc3\x36\xa4\xdd\xbb\x16\x8a\xc5\xc4\x7a\x50\xa4\x42\x92\xd3\xe6\x19\xfe\xee\x8f\xb2\x1c\xe7\xcf\x4b\x52\x24\x2f\xb9\x44\x12\xc9\x1f\x49\xfd\x48\xca\x55\x95\x5c\xf6\x80\x7e\x03\xfd\xb6\x30\x62\x5a\x38\xb8\xfe\xe3\xea\x06\x5e\x0a\x84\x2f\x4c\x31\x2e\xca\x19\xdc\x97\xae\xd0\xc6\xc6\x70\x2f\x25\x34\x4a\x16\x0c\x5a\x34\x73\xe4\x71\x63\xfd\xbf\x45\xd0\x13\x70\x85\xb0\x60\x75\x69\x72\x84\x5c\x73\x04\xda\x4e\xf5\x1c\x8d\x42\x0e\xe3\x05\x30\x78\x78\x7e\x8c\xac\x5b\x48\x6c\xcc\xa4\xc8\x51\x91\xa9\x2b\x98\x83\x9c\x29\x18\x23\x4c\x74\xa9\x38\x08\x45\x87\x08\x7f\x3f\x0d\x86\xff\x3c\x0f\x61\x22\x24\xc6\xbd\xcb\xa4\xae\x7b\xbd\xaa\xe2\x38\x11\x0a\xa1\x9f\x6b\xe5\x50\xb9\x3e\x9d\xa6\x16\x73\x27\xb4\x82\x5c\x32\x6b\xef\xfa\xed\x36\x8a\xc8\x81\x43\x03\x33\x2e\xa3\xa9\x11\xbc\x9f\x35\x8e\xd3\xe2\xcf\x6c\x20\x99\x98\x21\x4f\x13\x5a\x87\x43\x2e\xe6\x4b\x7b\xaf\x9f\x23\xa5\xbb\x5c\x44\xd1\xd5\x75\x94\x6b\xd9\x02\xf8\x5f\x55\xc5\x0f\x12\xad\x15\x6a\x6a\x29\x84\x06\x22\x21\x8c\xac\x97\x26\xad\xff\xcc\x87\xfb\x2e\x5c\x01\xf1\x08\x73\x3d\x55\xe2\x1b\xf2\xd3\xc2\x5d\xd9\x9f\x1e\x71\x5a\xca\xd5\x26\xa4\x60\x98\x9a\x22\xc4\x6d\xfc\x9d\xa2\x14\x19\xa5\x57\xd7\x69\x42\xab\x2d\x13\x54\x7c\x4d\x3d\x4d\x96\xa0\x3f\x26\xbf\xd4\x3d\x25\xdf\xff\xca\x31\x95\xc7\x5f\xb8\x38\x32\x5d\xaa\x95\x0f\xe4\xd1\xbb\xe0\xae\xd8\x43\x56\xdc\x61\x7f\x4e\xdb\x00\x8d\x13\x13\x91\x33\x87\x83\x82\x09\x65\x4f\x2c\xb6\x15\x0c\x04\x1c\xb8\x78\xd1\x8e\xc9\x5b\x8a\x4c\xa2\xf2\x0c\xfc\xf6\x33\xa5\x18\x78\xfc\x35\xf7\xd8\x82\x7f\xfc\xde\x2e\xe1\xf6\x6e\x83\xdc\x23\x43\xef\xcc\x8a\x9b\xac\x09\x1b\x7e\xa9\xaa\xce\x89\x2f\x0f\x12\x6c\x6a\x1e\x19\x79\x67\xe7\xd8\x58\xe2\xba\x25\x67\x8e\x45\xe1\xd4\x6f\xbf\xda\xed\x93\xd5\x36\x8a\x2c\x4a\xca\xa4\x13\xd9\x82\x71\xfd\x1e\x45\xd7\xfc\x6d\x87\xaf\xe0\xaf\x40\xc6\x77\xcb\x82\xdc\xec\x17\x06\x05\xbe\x3b\xda\xd7\xd7\x90\xaa\xa2\x6b\x55\xe5\x0c\x8d\xc8\xfb\xeb\xfc\xa7\x89\x3b\xe0\xf7\x78\xe8\xe1\x07\x0d\x41\x4b\x24\x9e\x1b\xb8\xe9\x17\x1a\xdc\xa1\x5b\xa0\x69\xc5\xf3\x7a\x18\xb0\x39\x32\x67\x0f\xc3\x92\x74\x0f\x15\x24\xd9\x4f\x62\xea\xc6\x9a\x2f\xf6\xc3\xae\x5a\xc6\x53\x13\x3a\x86\x56\xbe\x61\x42\x81\x6f\x8d\xc4\x4d\xec\xcf\x8b\x23\xf3\x8d\x12\xa0\x7d\x9f\x9c\xf7\xe2\x5a\xec\xb8\xe3\xfe\xdc\x2e\x36\xa7\x68\x55\x95\x6a\xc6\x0c\x75\x95\xec\x46\x67\xb8\xae\xf5\x51\x7a\xfe\x24\xfd\x60\x0c\x6e\xda\x4a\xa9\xeb\x35\xde\xd8\xbc\xa5\x8d\xcd\x03\x6b\x1b\x8a\x07\x03\x69\x82\x19\x9b\x24\x6b\xc6\x59\x03\x54\xd7\x7e\x14\xfb\x4d\xe3\x64\xf3\x91\xdb\x69\x7f\x62\xd5\xfa\xdf\x21\x7c\xb2\xdc\x5d\xb9\x24\xf0\x77\xb5\x35\x6f\xc3\xd3\xb5\xda\x76\x4f\xd8\x2e\x5f\xfb\x9f\xe8\xe5\x7b\x37\x34\x46\x9b\xd3\xde\xb8\x54\x74\xe4\xd2\x98\x33\x82\xc9\x48\xd0\x27\x9a\xed\x67\x42\x4d\x74\x9a\x88\x6c\x38\x1a\xfd\x3b\xba\xa8\x2a\xfa\x22\x24\x2f\x4f\x8f\xa7\xbc\x7a\xdb\x95\xd9\x7c\xa4\xec\xcf\xaa\xf9\xff\x1e\x00\x00\xff\xff\x37\x19\x1f\xac\xe6\x0a\x00\x00")

func blessingsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_blessingsHtml,
		"blessings.html",
	)
}

func blessingsHtml() (*asset, error) {
	bytes, err := blessingsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "blessings.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _chromeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x57\x4d\x8f\xdb\x36\x10\xbd\xfb\x57\x4c\x04\xf4\x92\x86\x52\x36\x45\x02\x24\x95\x1d\xec\x17\xda\x00\x9b\x6d\x80\xdd\xa6\xe8\xc9\xa0\xa5\xb1\xc4\x84\x22\x0d\x92\xb6\xd6\x30\xfc\xdf\x3b\x14\x65\xc7\x2b\xcb\x87\x40\xb5\x2e\x96\xc8\xe1\x9b\x79\xa3\x79\xa3\xf1\x66\x93\xbc\x1c\x01\x5d\xd7\x7a\xb1\x36\xa2\x28\x1d\xbc\x79\x7d\xf1\x0e\x1e\x4b\x84\xaf\x5c\xf1\x5c\x2c\x2b\xb8\x5c\xba\x52\x1b\x1b\xc3\xa5\x94\xd0\x18\x59\x30\x68\xd1\xac\x30\x8f\x9b\xd3\x7f\x5b\x04\x3d\x07\x57\x0a\x0b\x56\x2f\x4d\x86\x90\xe9\x1c\x81\x1e\x0b\xbd\x42\xa3\x30\x87\xd9\x1a\x38\x5c\x3d\xdc\x30\xeb\xd6\x12\x9b\x63\x52\x64\xa8\xe8\xa8\x2b\xb9\x83\x8c\x2b\x98\x21\xcc\xf5\x52\xe5\x20\x14\x2d\x22\xdc\x7d\xba\xbe\xbd\x7f\xb8\x85\xb9\x90\x18\x8f\x5e\x26\xdb\xed\x68\x94\xbe\xb8\xf9\xeb\xfa\xf1\xdf\x2f\xb7\x50\xba\x4a\x4e\x46\xa9\xff\x81\xa7\x4a\x2a\x3b\x8e\x4a\xe7\x16\x1f\x92\xa4\xae\xeb\xb8\xfe\x2d\xd6\xa6\x48\x2e\xde\xbf\x7f\x9f\x3c\x79\x9b\x68\x42\x87\x4b\xe4\xf9\xa4\x71\x9e\x3a\xe1\x24\x4e\x6e\x70\xb6\x2c\x0a\xa1\x0a\xd8\x6c\xe2\x07\xcf\xc9\xdc\xf3\x0a\xb7\xdb\x34\x09\x06\xc1\xb8\x42\xc7\x41\xd1\xc6\x38\x5a\x09\xac\x17\xda\xb8\x88\x38\x2a\x87\xca\x8d\xa3\x5a\xe4\xae\x1c\xe7\xb8\x22\x42\xac\x79\x78\x45\x14\x84\x13\x5c\x32\x9b\x71\x89\xe3\x8b\xf8\x75\xd4\x42\x49\xa1\xbe\x53\xfe\xe4\x38\x6a\x32\x61\x4b\x44\xc2\x2a\x0d\xce\x43\xfc\x96\x08\x58\xa7\x0d\x2f\x30\x2e\xb4\x2e\x24\xf2\x85\xb0\x71\xa6\xab\xc4\x27\x35\x2e\xd0\x55\xb9\x8c\x85\x4e\x08\x34\x7e\x97\x54\xdc\xa1\x21\x4f\xb1\x43\x72\x37\x93\x4b\x8c\x2b\xa1\xe2\xcc\xda\x9d\x47\x9b\x19\xb1\x70\x60\x4d\x36\xd0\x83\xc7\xfd\x46\xb0\x69\x12\x20\x7f\x82\xd1\x9c\x72\x65\xbb\xde\x04\x65\xf0\xe3\x9c\x57\x42\xae\xc7\x9f\x5b\x27\xbf\x7e\xa2\xc5\x1f\xa1\x7b\xc0\x70\xef\xaf\x78\x2e\x9e\x30\x0f\x29\x86\xcd\x7e\xd9\x5f\xde\x01\x0b\x58\x1f\xa0\xd2\x4a\xdb\x05\xcf\xf0\xf7\xbd\xcd\x36\x00\x26\x2d\x62\x9a\x84\x52\x18\xa5\x33\x9d\xaf\x5b\x6f\x2f\x18\xa3\x22\xaf\xf9\x9a\xca\xb8\xd4\xb5\xa5\x8a\xf5\x56\x68\x5e\x01\xae\x50\xf9\xb2\xb4\x15\x97\x12\x0d\x50\x02\x90\x8a\x37\x06\xc6\xda\xc3\xb9\x58\x41\x26\xb9\xa5\x32\xa4\xec\x31\xc9\xd7\x7a\xe9\xc0\xdf\x7e\xb3\x87\x4f\xe1\x96\xb1\x40\x25\xe0\x47\x3f\x28\xa6\x61\xe5\x18\x6a\x3a\x3d\xb2\x3d\xed\x77\x67\xcc\x8c\xae\x3b\x07\xf6\x4c\x1f\x7d\x7d\xef\xe3\x7f\xb6\x4d\xb9\x53\xc7\xa0\xac\x51\x44\x14\x34\x43\x99\x24\xa3\x13\xd0\x97\x79\x0e\x4d\xfe\x29\x73\x4e\x03\x97\xa2\x50\x24\x9f\x95\x28\xb8\x13\x5a\xf9\x35\xaf\xef\xd0\x71\x7a\x23\xe8\x65\xc5\x02\xa6\x2f\x40\xda\x3f\xe1\xfb\x7e\xef\x26\x86\x7f\x10\x4a\xe1\x1b\x91\xdb\xbf\xbb\xe3\x37\xf7\x0c\x81\x82\x3c\xf4\x7b\x10\xf3\xe1\xbb\x93\xdc\x14\xc8\x02\x12\xd3\x4a\xae\x7b\x72\xdc\xc0\xf1\x7e\xb0\xe9\xd4\x6b\x66\xa7\x91\xe4\xa3\x1a\x77\xba\x4f\x34\xf1\xbf\x69\xc2\x07\xe1\xce\x48\x8c\x96\x7a\x9b\xed\x73\x70\xb5\xdb\x1c\xea\xc5\x3a\xee\x7a\x3d\x3c\xf8\x8d\xa1\xe8\x52\xf7\x87\x7f\xa7\x87\x47\x5e\x48\x3d\xeb\xc3\xfe\x83\xd6\x87\x62\x2f\x8c\xf6\x9f\xad\xde\xd8\xbf\xb4\x7b\x43\x7d\xac\x9c\x21\x3d\xf4\xba\x78\x6c\x76\x06\xbf\xda\xb5\xca\x66\xdc\x62\xef\xdb\xa5\xbd\x2b\xda\xeb\xf5\x91\x26\x84\xd9\x69\x54\xcf\x35\x1b\x1a\x30\x9a\x83\x95\x13\x9d\x2c\x37\xbc\x3e\x6e\x7b\x27\x85\xda\xd7\xee\xce\x20\xc3\xf3\x4b\xf0\x7c\xf2\x3b\x97\xf4\xce\x25\xbb\x73\x4b\xee\xac\x72\x3b\xa7\xd4\x3a\x32\xeb\x4a\xac\xe2\xa2\xe7\x33\x3e\x9d\xb6\x83\x6b\x57\x52\x16\xb3\xe6\x4b\xd7\x9e\x68\x1f\x19\xa3\x29\x9d\x66\xb3\xe6\x03\x58\x18\x91\xf7\x09\xac\x7c\xbb\x3b\x75\x30\xa1\x45\x93\xa3\x99\xba\x7c\xdb\x25\xd0\x7a\x79\xbe\xbc\xd9\xcc\xa4\xce\xbe\x43\xb4\x8b\x14\xe2\xed\xf6\xb3\x67\x73\x1d\x16\xe0\x4f\x34\xb8\xd9\xa0\xca\xb7\xdb\xe7\x80\xa5\x49\x26\x74\xbe\x16\x34\x21\xc6\xd7\xba\xaa\xb8\xca\xef\x84\xc2\xae\xdd\x20\xb2\x93\x03\xe0\x63\x52\x8d\x51\xa7\x99\x65\x48\xa3\xc7\xee\x86\xb1\x8b\x37\x2c\xd3\x12\xba\xc9\xf2\x29\x3a\x9a\x6c\x4e\xe6\xa8\x61\xbf\xe3\xba\x30\x42\xb9\x39\x44\xbf\xac\x28\x59\x5f\x9b\x4a\xfd\x7f\x29\x37\x25\x3e\x80\x2c\x4d\x6d\x0b\x83\x3b\x96\xfe\xf6\xa7\xb9\x1e\x14\xba\x2f\xed\x76\xf0\x0e\x28\x69\x12\x06\x79\x8a\xb0\xf9\x5f\xf8\x5f\x00\x00\x00\xff\xff\x9c\xc5\x61\xad\xdb\x0e\x00\x00")

func chromeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_chromeHtml,
		"chrome.html",
	)
}

func chromeHtml() (*asset, error) {
	bytes, err := chromeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "chrome.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _globHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x91\xdf\x8b\xd4\x30\x10\xc7\xdf\xfb\x57\x0c\x45\x44\x0f\xda\x78\x07\xfa\x20\xd9\xca\x79\xd7\x87\x03\x39\x65\x57\x7d\xcf\x26\xd3\x26\x90\x4d\xce\x49\xba\xba\x84\xfc\xef\xa6\xb5\xcb\x2d\xa2\xf3\xd2\xce\xaf\xcf\xcc\x7c\x93\x12\xbb\xaa\xa0\xd8\x9d\x7f\x3a\x91\x19\x75\x84\x9b\x37\xd7\xef\xe0\xab\x46\xf8\x2e\x9c\x50\x66\x3a\xc0\xed\x14\xb5\xa7\xd0\xc2\xad\xb5\xb0\x14\x05\x20\x0c\x48\x47\x54\xed\xd2\xfd\x2d\x20\xf8\x01\xa2\x36\x01\x82\x9f\x48\x22\x48\xaf\x10\x8a\x3b\xfa\x23\x92\x43\x05\xfb\x13\x08\xf8\xb8\xbb\x6f\x42\x3c\x59\x5c\xda\xac\x91\xe8\x4a\x6b\xd4\x22\x82\x14\x0e\xf6\x08\x83\x9f\x9c\x02\xe3\x4a\x10\xe1\xd3\xc3\x5d\xff\xb8\xeb\x61\x30\x16\xdb\xea\x8a\xe5\x5c\x55\x29\x29\x1c\x8c\x43\xa8\xa5\x77\x11\x5d\xac\x4b\x94\x07\x94\xd1\x78\x07\xd2\x8a\x10\x36\xf5\xea\x36\x4d\x19\x10\x91\xe0\xa0\x6c\x33\x92\x51\x75\xb7\x0c\xe6\xfa\x6d\x97\x52\xfb\x45\xc4\x92\x74\x39\x73\x56\x02\x7f\x32\xca\x1c\xc1\xa8\x4d\xfd\x24\x68\x66\x9f\x81\x33\x40\x62\xb9\xff\xfc\xd3\x34\xd7\x37\x8d\xf4\x76\x25\x2e\xbd\xde\x3e\x3b\xb3\xa5\x44\xc2\x8d\x08\x6d\xef\x22\x19\x0c\x39\x97\xd0\x4f\x13\x35\xb4\xbb\x69\x18\xcc\xaf\xb2\xf9\x65\x3d\xb7\xa6\xe3\x02\x34\xe1\xb0\xa9\xd9\x68\xfd\xfe\x83\xdb\xa4\x34\x91\xfd\x31\x21\x9d\xe0\x45\xbb\x9b\x45\xa7\x47\x71\xc0\x9c\x5f\x86\xcb\x5c\x9b\x73\x3d\xdf\x34\x1f\x23\x3a\xce\x0a\xea\xaf\x5d\xd0\xa9\x8b\x05\x7a\x22\x4f\xff\x98\xdf\x6f\xb7\x9f\xb7\xaf\x52\x2a\x63\x4a\xc1\xc3\xfd\x0c\x7e\xfd\x1e\x56\xf2\xff\xa9\xcb\xf7\x59\x0a\x76\xd6\x82\xb3\xa2\x68\x57\x71\xb6\x3e\x49\x57\xad\xb5\xbf\x03\x00\x00\xff\xff\x8e\x08\x47\x71\x7b\x02\x00\x00")

func globHtmlBytes() ([]byte, error) {
	return bindataRead(
		_globHtml,
		"glob.html",
	)
}

func globHtml() (*asset, error) {
	bytes, err := globHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "glob.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _logsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x53\xc1\x8e\xda\x30\x10\xbd\xf3\x15\xa3\x9c\x60\xa5\x24\xdd\x95\xda\x4b\x81\x6a\x97\x22\x75\x25\xba\x17\xb6\xbd\xae\xbc\xf6\x90\xb8\x72\x6c\x3a\x76\xa8\x50\x94\x7f\xef\xc4\x24\x5d\x60\xa1\x39\xc5\xe3\x79\x6f\x9e\x9f\x9f\x9b\x26\xbf\x19\x01\x7f\x0b\xb7\xdd\x93\x2e\xca\x00\x77\x1f\x6e\x3f\xc1\x73\x89\xf0\x53\x58\xa1\x74\x5d\xc1\x7d\x1d\x4a\x47\x3e\x83\x7b\x63\x20\x36\x79\x20\xf4\x48\x3b\x54\x59\x44\xff\xf0\x08\x6e\x03\xa1\xd4\x1e\xbc\xab\x49\x22\x48\xa7\x10\x78\x59\xb8\x1d\x92\x45\x05\xaf\x7b\x10\xf0\xb0\xfe\x9a\xfa\xb0\x37\x18\x61\x46\x4b\xb4\x0c\x0d\xa5\x08\x20\x85\x85\x57\x84\x8d\xab\xad\x02\x6d\xb9\x88\xb0\x7a\x5c\x2c\x9f\xd6\x4b\xd8\x68\x83\xd9\xe8\x26\x6f\xdb\xd1\xa8\x69\x14\x6e\xb4\x45\x48\xa4\xb3\x01\x6d\x48\xb8\x3a\xf5\x28\x83\x76\x16\xa4\x11\xde\xcf\x92\x7e\x99\xa6\x3c\x20\x20\x41\xa5\x4c\x5a\x90\x56\xc9\x3c\x0e\x9e\x96\x1f\xe7\x2b\xed\x43\x27\xda\xb8\x22\xf2\xfb\x69\xce\xd5\xc3\xb6\xd2\x3b\xd0\x6a\x96\x6c\x05\x75\x03\x06\xd6\x8e\x45\x22\x9b\x30\xfc\xa4\xe9\xed\x5d\x2a\x9d\xe9\x69\x23\xd6\x4b\xd2\xdb\xf0\x56\xe8\xbe\x9d\xa0\xc1\x97\x19\x58\xfc\x03\xcb\x1d\xf3\xae\x63\x65\x9c\xe4\x2c\xc1\x7f\xb1\xb3\xa6\xa9\xc9\xfc\xae\x91\xf6\x90\xad\x3b\x77\xe9\x49\x54\xd8\xb6\xc9\xe4\xf3\x09\xdb\x81\x29\x73\xb6\x42\xef\x45\xd1\x71\x6e\x6a\x1b\x0f\x3c\xc6\x8e\x78\x02\xcd\x09\x60\x90\xe0\x0c\xb7\x2a\x27\xeb\x8a\x9b\xb2\x02\xc3\xd2\x60\xf7\xfb\xb0\x7f\x54\xe3\x84\x55\x44\x1f\xce\xc7\x0d\x68\xa3\x8f\xd1\x92\x50\x04\xec\x09\x18\xac\x2f\xc1\x8c\xce\xb4\xb5\x48\xdf\x9e\xbf\xaf\x18\x1c\xc5\x65\x4a\x04\xf1\xbe\xd5\x99\x4c\x6c\xb7\x68\xd5\xa2\xd4\x46\x8d\x8d\x3e\xa3\x6b\xaf\x78\x80\x44\x8e\x8e\x1d\xb8\x74\xf8\xbe\x5b\x1a\xe7\x71\x7c\x41\xe7\x55\x4f\xfa\x00\x4c\x32\xc2\x8a\x83\x7c\xd0\xf6\x0e\xfe\x7f\x0a\x72\x05\xbf\x16\xb6\xf5\xfa\x89\xa6\xf9\x79\x6a\xde\x22\x38\xc0\x8f\x43\x38\x14\x63\x10\x7f\xf9\xd3\xf5\xb0\x78\x79\xd1\x56\x21\xa7\xbf\xd2\x96\xaf\x2a\x99\x4f\x73\x26\x3d\x1a\xc1\x71\xe8\x26\xfc\xbb\xf7\xa3\xad\xdc\x99\xfe\x29\x1c\x30\x2c\xf0\xf0\xa2\xe6\xfc\xfe\xf8\x92\xda\xf6\x6f\x00\x00\x00\xff\xff\xb0\xbe\x01\x8d\x3a\x04\x00\x00")

func logsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_logsHtml,
		"logs.html",
	)
}

func logsHtml() (*asset, error) {
	bytes, err := logsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logs.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _no_syncbaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x34\x90\x31\x6b\xc3\x30\x14\x84\x77\xff\x8a\xc3\x63\x40\x71\x13\x68\x27\xd7\x90\xa6\x19\x0a\xa5\x4b\xda\xee\x8a\xf4\x1c\x09\x94\xa7\x22\xc9\x06\x63\xf2\xdf\x2b\x3b\xb6\x27\xbf\x27\xdd\x9d\xbe\x1b\xc7\x6a\x53\x20\x7f\x47\xff\x37\x04\x7b\x35\x09\xfb\xa7\xdd\x0b\xbe\x0d\xe1\x57\xb2\xd4\xb6\xbb\xe1\xd0\x25\xe3\x43\xdc\xe2\xe0\x1c\xe6\x4b\x11\x81\x22\x85\x9e\xf4\x76\x56\xff\x44\x82\x6f\x91\x8c\x8d\x88\xbe\x0b\x8a\xa0\xbc\x26\xe4\xf1\xea\x7b\x0a\x4c\x1a\x97\x01\x12\x6f\xe7\x77\x11\xd3\xe0\x68\x96\x39\xab\x88\xb3\x34\x19\x99\xa0\x24\xe3\x42\x68\x7d\xc7\x1a\x96\xf3\x92\xf0\xf9\x71\x3c\x7d\x9d\x4f\x68\xad\xa3\x6d\xb1\xa9\xee\xf7\xa2\x18\x47\x4d\xad\x65\x42\xa9\x3c\x27\xe2\x54\x4e\xdb\x3a\x92\x4a\xd6\x33\x94\x93\x31\xbe\x96\xcb\x28\x44\x4e\x48\x14\x70\xd3\x4e\x5c\x83\xd5\x65\x33\x27\xd7\xe6\xb9\x39\x0f\xac\x2e\x32\x52\x5d\xe5\xe1\xb1\xd5\xb6\x5f\x0d\x26\x81\xa2\x0c\xbc\xfe\x08\xb1\xdb\x0b\xe5\xdd\xe2\x80\xdc\xd1\x44\x3b\xb5\x10\xa0\x3d\x45\xb0\x4f\x30\xb2\xa7\xcc\xb9\x7a\xcf\xe7\x99\xf2\x51\x53\x5d\xe5\x80\xa6\xa8\xab\xe5\x75\xcd\x44\x43\xac\x33\xc0\x7f\x00\x00\x00\xff\xff\xfe\x8f\x17\xb9\x89\x01\x00\x00")

func no_syncbaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_no_syncbaseHtml,
		"no_syncbase.html",
	)
}

func no_syncbaseHtml() (*asset, error) {
	bytes, err := no_syncbaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "no_syncbase.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _profilesHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\xcd\x6e\xdb\x3c\x10\xbc\xfb\x29\x16\x3a\x7c\x5f\x1b\x40\x52\x13\xa0\x3d\x14\xb2\x8a\x24\x0d\xd2\x02\x41\x10\xd4\x49\xef\x34\xb9\x92\x88\x50\x5c\x95\xa4\x92\x18\x82\xde\xbd\x2b\x55\xce\x6f\x9b\x18\xb0\x7d\x30\xc4\xc5\xce\xcc\xce\x52\x1e\x77\x5d\xba\x37\x03\xfe\x1c\x53\xb3\x72\xba\xac\x02\x1c\x7c\xd8\xff\x04\x97\x15\xc2\x4f\x61\x85\xd2\x6d\x0d\x87\x6d\xa8\xc8\xf9\x04\x0e\x8d\x81\xb1\xc9\x83\x43\x8f\xee\x06\x55\x32\xa2\xaf\x3c\x02\x15\x10\x2a\xed\xc1\x53\xeb\x24\x82\x24\x85\xc0\xc7\x92\x6e\xd0\x59\x54\xb0\x5c\x81\x80\xa3\xc5\xd7\xd8\x87\x95\xc1\x11\x66\xb4\x44\xcb\xd0\x50\x89\x00\x52\x58\x58\x22\x14\xd4\x5a\x05\xda\x72\x11\xe1\xec\xfb\xf1\xc9\xf9\xe2\x04\x0a\x6d\x30\x99\xed\xa5\x7d\x3f\x9b\x75\x9d\xc2\x42\x5b\x84\x48\x92\x0d\x68\x43\xc4\xd5\xcc\xa3\x0c\x9a\x2c\x48\x23\xbc\x9f\x47\xd3\x31\x8e\x59\x20\xa0\x83\x5a\x99\xb8\x74\x5a\x45\xf9\x28\x9c\x55\x1f\xf3\x0b\x47\x4c\xab\x6d\x99\xa5\x7c\xfa\x53\x56\xfa\x06\xb4\x9a\x47\x8d\x70\x03\xf1\x9a\x6d\x40\x4b\x64\xf3\xeb\x87\x38\xde\x3f\x88\x25\x99\x89\x6e\xc4\xb6\xe6\xe1\x30\x16\x8c\xce\x8f\x2f\xae\x9e\xd4\xee\x45\x26\xe2\x42\xdf\xa1\x8a\x6f\xb5\x0a\x55\x94\x97\x04\x81\xc8\x40\xd3\xf0\x64\xd0\x75\xc9\xd5\x8f\xb3\x0b\xc7\x5e\xef\xfa\x3e\x6d\xc6\x69\xf1\x8b\x9d\x77\x5d\xeb\xcc\xaf\x16\xdd\x0a\x92\xc5\x70\x07\xee\x5c\xd4\xd8\xf7\x59\xca\xc4\xcf\x46\x48\x79\x86\x17\x43\x65\x02\x2a\xa6\x9d\x47\xcf\x24\x2a\x14\xcd\x2b\xfc\xff\x29\x5c\xb6\xe5\x7c\x3f\xca\xbf\x71\x63\x96\x8a\x7c\x47\xd6\xde\xd0\xdd\xda\xd7\xd2\x90\xbc\xde\xc8\xd8\xd1\xd0\xb9\x43\x67\x6f\x29\x6f\x6d\x2d\x70\x55\x28\xc9\x5f\xe1\xb5\x57\xe3\xc1\xe1\xe5\x23\xc0\x0e\x8d\x6e\x38\xc7\xc6\x7e\x4f\xc9\x51\x1b\xf8\x47\xee\x3f\xbf\x9c\xf0\x1f\xbb\x28\xd7\x98\x8d\x16\xf1\x4e\x52\xdd\x08\x19\xde\xff\x7d\x0b\xdb\x6b\x1c\xb0\x46\xd1\x1a\xf3\x52\xe0\xa9\xe5\x2c\x7d\x1c\x1c\xdb\x24\xd0\x88\xd7\xf7\x08\xbe\x0a\xa7\x85\x89\x35\xc7\xa4\x8f\x72\x6d\x0b\xca\x52\x9d\x0f\xc9\xce\xde\x6b\x61\x95\x07\xb1\xe4\x74\x86\x5a\xac\xc0\x52\x80\x5b\x72\xd7\xa0\x8b\x31\x79\x1d\xd6\x14\x10\xf8\xb2\x25\x7a\xcf\x41\x6e\xff\xe7\x06\xa7\x03\x27\xee\x90\xce\xa7\x94\x3c\x11\x5e\xb4\x4d\x43\x2e\x70\x7c\x3b\x68\xd6\xc9\x3a\xfd\x09\x58\x20\xa6\x74\x60\x84\x2d\x5b\x51\xe2\x40\xb7\x4e\xf8\x5b\xed\x2b\xa3\x7d\x48\x1e\x2d\xe4\xfe\x25\x99\x1e\xb3\x74\x8a\xf3\x9c\xc3\x1f\xad\xea\xfb\xdf\x01\x00\x00\xff\xff\xd3\x06\xf2\xb8\xb7\x06\x00\x00")

func profilesHtmlBytes() ([]byte, error) {
	return bindataRead(
		_profilesHtml,
		"profiles.html",
	)
}

func profilesHtml() (*asset, error) {
	bytes, err := profilesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "profiles.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resolveHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x51\x6f\xe2\x38\x10\x7e\xe7\x57\x8c\xa2\x7b\xb8\xab\x14\x72\xad\x74\xf7\x50\x85\x9c\xda\xc2\x03\x52\xaf\x2b\x41\x77\x5f\x2b\x13\x4f\x88\x77\x8d\xcd\xda\x4e\x29\x4a\xf9\xef\x3b\x4e\x02\x8d\x28\xac\x44\x37\xdd\xdd\xbc\x80\x67\x3e\x7f\x33\xf3\x8d\x13\x4f\x59\x46\x67\x3d\xa0\xe7\x46\x2f\xd7\x46\xcc\x73\x07\x17\x7f\x9f\xff\x0b\xf7\x39\xc2\x27\xa6\x18\x17\xc5\x02\xae\x0a\x97\x6b\x63\xfb\x70\x25\x25\x54\x20\x0b\x06\x2d\x9a\x47\xe4\xfd\x6a\xf7\x47\x8b\xa0\x33\x70\xb9\xb0\x60\x75\x61\x52\x84\x54\x73\x04\x5a\xce\xf5\x23\x1a\x85\x1c\x66\x6b\x60\x70\x3d\x1d\x86\xd6\xad\x25\x56\xdb\xa4\x48\x51\xd1\x56\x97\x33\x07\x29\x53\x30\x43\xc8\x74\xa1\x38\x08\x45\x46\x84\xdb\xf1\xcd\xe8\x6e\x3a\x82\x4c\x48\xec\xf7\xce\xa2\xcd\xa6\xd7\x2b\x4b\x8e\x99\x50\x08\x41\xaa\x95\x43\xe5\x02\xb2\xc6\x16\x53\x27\xb4\x82\x54\x32\x6b\x07\x41\xb3\x0c\x43\x0a\xe0\xd0\xc0\x82\xcb\x70\x6e\x04\x0f\x92\x2a\x70\x9c\xff\x93\xdc\xb1\x05\xfa\x32\xb4\x2c\x3c\x34\x8e\xc8\x56\x3b\xb9\x78\xdc\xf2\xf8\x7d\x29\x52\xd9\xdb\x3f\x61\x78\x7e\x11\xa6\x5a\x36\x44\xfe\x29\xcb\x95\x70\x39\xf4\xff\xa7\xcc\xdd\x48\x39\xb3\xa6\x84\xb6\xce\xb8\x90\x2f\xc8\x36\xda\x87\x6f\xe1\x2a\xac\x14\xc9\xb4\xc8\x32\xf1\x74\x49\xb0\xfe\x66\x13\x47\x64\x29\x4b\x54\x7c\xb3\xd9\x6d\x9c\x7a\xdd\x6d\x15\xec\x9e\xcd\xe4\x21\x92\xfb\xaa\x0f\x1e\x67\x7c\x0b\x18\x2c\x3c\xda\x79\xf4\x41\xca\xb1\xbd\x45\x96\x1d\x23\xaa\x18\x24\x01\x1a\xca\x36\xc5\x4b\x9d\x51\xbb\xd0\xb2\x34\x4c\xcd\xb1\x49\xd6\xd8\x36\xf0\x44\x75\xab\x3d\x55\xea\xed\x5d\x9c\x39\x16\xd6\x56\xbf\xfc\x6c\xf7\x2d\x2f\xcb\x30\xb4\x28\xe9\x38\xec\x5c\x36\x67\x5c\xaf\xc2\xf0\x82\x2f\xf7\xe2\xd4\xb1\x66\x9a\xaf\x5f\xdb\x6b\x9f\x39\xec\xa8\x9d\xfc\x70\x86\x0f\x0f\x75\x69\x8a\xce\xa3\x2a\x16\x68\x44\x1a\x24\x23\xc5\x97\x5a\x28\x17\x47\x8e\x77\xc5\x19\x33\xc8\x0d\x66\x83\x20\xfa\x4f\x0d\xaa\x06\x55\x21\xb6\x5d\x80\x67\xd8\x9a\xaa\xb3\xff\x0c\x85\x91\x5f\x0b\xf4\xe7\x35\xa0\x86\x36\x30\x7f\xec\x58\x72\x3c\x2f\xf2\x1c\x51\xa1\x4b\x79\x9e\x96\x82\xde\xcd\x2e\xd5\xa1\x02\x87\xc8\xb8\xa4\x4f\x87\x2f\xf1\xf4\xf2\x9a\xb7\xe5\x0f\x5c\xc2\xe5\x00\xf6\xd5\xdd\x7b\x7b\xde\x41\x93\x3b\x74\x2b\x6d\xbe\x74\xac\xc9\x15\xe7\xa6\xdf\x50\xbf\x4d\x97\x0e\x4b\xf4\xc9\xa0\xed\xba\xed\x9e\xf5\x97\x97\x36\xd1\x74\xcd\xa8\xf9\x78\xd8\x71\x71\x3b\xde\x1f\x3a\xd4\xfd\x6b\x49\xba\x13\x8f\xff\x32\xd8\xf7\x3f\xcb\xdb\x70\x16\xfe\x2c\x4b\x89\x0a\xe8\xb6\xfb\xab\x63\x65\xde\x2a\x48\x7d\x37\x7e\x57\x2d\x2f\xfa\xcf\x90\xa9\x8e\xf3\xbb\x6a\x74\x92\x8f\xd8\x5e\x5f\xac\x64\xf4\xd9\x25\xad\x31\x82\xc6\x83\xf6\x1c\x81\xd2\xb6\x27\x9c\xbd\xa1\x8d\x86\x46\x5a\x17\x4b\xc0\xc5\xd2\xad\x7b\xc7\x72\x68\xaf\x9b\x08\x71\xd4\x8c\x88\x89\x9f\x28\xeb\xae\x8e\x8c\xd1\xe6\x6d\xc3\x64\x2c\x76\xba\x33\x82\x08\x26\x43\x41\xe3\xa9\x0d\x12\xa1\x32\x1d\x47\x22\x19\x4d\x26\x1f\x26\xd4\x48\xba\x2a\x28\xca\x78\xd8\x74\xf3\xa4\xa1\x93\xe6\xe0\x27\xe4\xe1\x4a\x70\x97\xef\xda\xb7\x5f\xce\xb6\xd8\xe6\xf7\x5b\x00\x00\x00\xff\xff\xbc\x42\xa9\x8f\xe2\x0b\x00\x00")

func resolveHtmlBytes() ([]byte, error) {
	return bindataRead(
		_resolveHtml,
		"resolve.html",
	)
}

func resolveHtml() (*asset, error) {
	bytes, err := resolveHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resolve.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _statsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x55\xdd\x4f\xdb\x30\x10\x7f\xcf\x5f\x71\x8a\xd8\x04\x48\x4e\x06\xd2\xf6\x80\xdc\x4c\x0c\x3a\x84\xc4\x40\x02\xd6\x57\xe4\xc6\x97\xc6\x93\x6b\x77\xb6\x03\xab\xac\xfc\xef\x73\x3e\x0a\x05\x8a\xf8\x50\xdb\x97\xfa\xce\xf7\xf5\xfb\xdd\xe5\xec\x7d\xba\x1b\x41\xf8\x1d\xe9\xd9\xdc\x88\x49\xe9\x60\xff\xcb\xde\x37\xb8\x2e\x11\x46\x4c\x31\x2e\xaa\x29\x1c\x56\xae\xd4\xc6\x26\x70\x28\x25\xb4\x46\x16\x0c\x5a\x34\xb7\xc8\x93\xd6\xfb\xb7\x45\xd0\x05\xb8\x52\x58\xb0\xba\x32\x39\x42\xae\x39\x42\x10\x27\xfa\x16\x8d\x42\x0e\xe3\x39\x30\xf8\x71\x75\x4c\xac\x9b\x4b\x6c\xdd\xa4\xc8\x51\x05\x57\x57\x32\x07\x39\x53\x30\x46\x28\x74\xa5\x38\x08\x15\x94\x08\x67\xa7\x47\xc3\xf3\xab\x21\x14\x42\x62\x12\xed\xa6\x75\x1d\x45\xde\x73\x2c\x84\x42\x88\x73\xad\x1c\x2a\x17\xd7\x35\x78\x2f\x0a\x48\x46\x4c\x56\x18\x6c\xa8\xc5\xdc\x09\xad\x20\x97\xcc\xda\x41\xdc\x8b\x84\x84\x74\x0e\x0d\x4c\xb9\x24\x13\x23\x78\x9c\xb5\x65\x50\x2e\x6e\x17\xa6\xcd\x55\x8e\x01\xe7\xe2\x40\xc8\xde\x3e\xc9\xb5\xec\x6d\x5b\x7b\xc7\xc6\x12\x97\x3d\x38\x73\x8c\x74\xda\x46\xfc\x63\x9f\x6a\x1e\x44\x42\x2c\xca\x50\xcf\xfd\x95\x2d\x19\xd7\x77\x84\xec\xf3\xd9\x52\x8e\x2e\xcf\x58\xf3\xf9\x63\x5d\xf3\xf3\xfe\x4e\xb8\x12\xb6\x26\x3a\x00\x86\x83\x01\xb4\x87\x0a\x7f\x1a\x3d\x1d\x5d\xfc\x7a\xe0\xe1\xa9\x23\x75\xe6\x79\xb4\xee\x82\xaf\x86\x73\x73\xd3\x71\xa0\x02\x7b\xaa\x9a\xa2\x11\x79\x9c\xb5\xe1\x61\xfb\x44\xef\xd0\xd4\xf1\x75\x44\xa4\x33\x83\x99\xf7\x1d\xa2\xba\xa6\x69\x23\xaf\x0e\x1e\xb4\x2b\x40\xac\x09\xd9\xf5\x7c\xb6\x4e\x60\x61\x6e\xff\x21\x27\x77\x82\xbb\x32\x0e\xf8\x66\x46\x28\x57\x40\xfc\xe9\x3a\x86\x07\xac\x6f\x47\xe9\x3d\x4a\xbb\xf9\xc6\x8e\x8e\xcf\x36\x44\xc0\x62\x32\xdf\x09\x5a\xf1\xcd\x61\xee\x5a\xbe\x69\xc8\x49\x93\xe6\xed\xb8\x83\xe6\xf1\xa7\x1f\x14\x4d\xbe\x7e\x5f\xa5\x61\x61\x65\x11\x4d\xfb\xbd\x96\x45\x0b\x8e\xba\x35\x78\x22\xf5\x78\x8c\xfc\x03\x8b\xb0\xfc\x9a\x35\xce\x34\x0d\x87\x8f\xad\xc6\x4a\x3e\x46\xe2\xbd\x61\x6a\x82\x90\x1c\x95\x42\x72\x83\xea\x49\x27\xa9\x14\x19\x65\x50\x1a\x2c\x06\x71\x6a\x1d\x73\xf6\xbb\x1a\x78\x5f\x19\xf9\xb7\x42\x33\x87\xad\xe4\xaa\x79\x68\xcc\x39\x9b\x06\x02\x3f\xdb\xe5\xbb\xa4\xae\x5b\x8a\x1b\x62\x59\xd8\x17\x21\x56\xf4\xda\xe8\x3c\x2b\x68\x68\x4c\x78\xdc\x56\x94\x35\xbc\xbc\xbc\xb8\xdc\xf6\x3e\x24\x0f\x16\xa7\xc7\x4d\xba\x9d\x03\xe8\xf3\xbd\x96\x8b\xa6\x0b\x2a\x56\xb6\xab\xfb\x8e\xfb\x4d\x9e\xb4\x35\x7c\xac\x5f\x54\xdc\xf7\x87\x05\x13\xc1\x24\x11\xe1\x61\xb4\x71\x26\x54\xa1\x69\xfa\x02\x8e\x77\xb6\xf8\xd9\x54\x37\x14\xbc\x30\x85\xb0\x34\x8e\xed\xff\xff\x00\x00\x00\xff\xff\x30\x43\x30\x2b\x65\x08\x00\x00")

func statsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_statsHtml,
		"stats.html",
	)
}

func statsHtml() (*asset, error) {
	bytes, err := statsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stats.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _syncbaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\x4d\x8b\xdb\x30\x10\xbd\xfb\x57\x3c\x7c\x69\xbb\x60\xbb\xbb\xd0\x9e\x5c\x43\x36\xbb\x85\x85\x92\x4b\xb6\xbd\x2b\xd2\xc4\x16\x28\x52\xd0\xc8\x81\x60\xf2\xdf\x2b\x3b\xb6\x93\xb4\xa7\x8d\x4f\x1a\xcd\x9b\x79\x1f\x72\xd7\x15\x0f\x09\xe2\xb7\x74\xfb\xa3\xd7\x75\x13\xf0\xf4\xf5\xf1\x3b\xde\x1b\xc2\x1f\x61\x85\xd2\xed\x0e\x8b\x36\x34\xce\x73\x8e\x85\x31\x18\x40\x0c\x4f\x4c\xfe\x40\x2a\x1f\xa6\x7f\x33\xc1\x6d\x11\x1a\xcd\x60\xd7\x7a\x49\x90\x4e\x11\x62\x59\xbb\x03\x79\x4b\x0a\x9b\x23\x04\x9e\xd7\x2f\x19\x87\xa3\xa1\x61\xcc\x68\x49\x36\x8e\x86\x46\x04\x48\x61\xb1\x21\x6c\x5d\x6b\x15\xb4\x8d\x97\x84\x5f\x6f\xcb\xd7\xd5\xfa\x15\x5b\x6d\x28\x4f\x1e\x8a\xd3\x29\x49\xba\x4e\xd1\x56\x5b\x42\x2a\x9d\x0d\x64\x43\xda\xdf\x96\x4c\x32\x68\x67\x21\x8d\x60\xfe\x91\x8e\x65\x96\x45\x86\x40\x1e\x3b\x65\xb2\xda\x6b\x95\x56\x03\x73\xd9\x7c\xab\xd6\x47\x2b\x37\x82\xa9\x2c\x62\x71\xbe\x55\xfa\x30\x2d\xe8\x07\x24\x45\xc3\xd3\x21\xcb\x1e\x9f\x32\xe9\xcc\xb8\x01\x58\xc7\x00\xa2\x01\xa4\x5d\x97\x8f\xe7\xfc\x67\x6b\xcc\x4a\xec\xe8\x74\x4a\xcf\x1b\x8b\xb8\xb2\x4a\xca\x62\xd4\x53\xf5\xfa\xbd\xb0\x35\x21\x7f\xf7\x14\x71\x11\x76\x87\xf6\x17\x11\x44\xaf\x7d\x20\x9f\x8a\xfc\x4d\xe5\x23\xf9\xc5\xd3\x47\x5d\x01\x8b\xfd\x1e\xcf\x86\x98\xb5\xad\xff\xdb\x3f\x35\x26\x83\xb3\xc5\xfb\xb8\x4a\x65\x2e\xc5\x9c\xcc\xd2\x19\x73\xce\x80\x87\x80\x2e\xe8\x50\x5d\x7a\x83\xb6\x2b\xcb\xb1\x79\x83\x55\xd5\xd2\x93\x08\xce\x7f\xe2\xfe\x07\xf5\xb7\xa6\x6e\xbd\x44\x13\xea\x5a\x08\x59\x35\x33\xc7\xde\x28\x72\xb6\x7a\xf5\x9e\x11\x6b\xf8\xbe\x77\xfc\x60\x56\x9f\x57\x0e\x6a\x7c\x0a\xfe\xf2\x6f\xfa\xe5\xb5\xa2\x41\x7d\x32\x9f\xfe\x06\x00\x00\xff\xff\x0c\xa0\x9a\x98\xe6\x03\x00\x00")

func syncbaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_syncbaseHtml,
		"syncbase.html",
	)
}

func syncbaseHtml() (*asset, error) {
	bytes, err := syncbaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syncbase.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vtraceHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x94\xdf\x6f\xdb\x36\x10\xc7\xdf\xf3\x57\xdc\xb4\x06\x4e\xba\x51\x76\x02\x64\x18\xe4\xd8\x40\x9b\xe6\x21\x40\xd1\x87\xa5\xeb\x9e\x69\xf2\x24\x71\xa1\x49\x81\xa4\xdd\x1a\x9a\xfe\xf7\x1d\x29\xc9\x3f\x82\x06\x89\x1f\x0c\xf2\xc4\x2f\xbf\x77\x1f\x1e\xd9\xb6\xd3\xf7\x67\x40\xbf\x3b\xdb\xec\x9c\xaa\xea\x00\xd7\xb3\xab\x3f\xe0\x6b\x8d\xf0\x8d\x1b\x2e\xd5\x66\x0d\x1f\x36\xa1\xb6\xce\xe7\xf0\x41\x6b\x48\x8b\x3c\x38\xf4\xe8\xb6\x28\xf3\xa4\xfe\xdb\x23\xd8\x12\x42\xad\x3c\x78\xbb\x71\x02\x41\x58\x89\x40\xd3\xca\x6e\xd1\x19\x94\xb0\xda\x01\x87\x8f\x8f\x9f\x98\x0f\x3b\x8d\x49\xa6\x95\x40\x43\xd2\x50\xf3\x00\x82\x1b\x58\x21\x94\x76\x63\x24\x28\x43\x41\x84\xcf\x0f\x77\xf7\x5f\x1e\xef\xa1\x54\x1a\xf3\xb3\xf7\xd3\xae\x3b\x6b\x5b\x89\xa5\x32\x08\x59\xee\x1b\x6e\x32\x0a\xdd\x4a\xb5\x85\xb4\xeb\x22\x6b\xac\x57\x41\x59\x53\x38\xd4\x3c\xa8\x2d\xce\x35\x96\xa1\x68\xdb\xfc\x31\x70\x17\xba\xee\x7c\xfe\x5d\xc9\x50\xc7\xc8\x3f\x71\x10\x23\x6b\xee\x2a\x65\x8a\x59\xf3\x63\xde\x70\x29\x95\xa9\x58\xb0\x4d\x71\x4d\xf3\x0c\x94\x5c\x64\x64\xc0\x48\xf0\x20\xbb\x2e\x5b\xa6\xd4\x6f\x7f\x61\x0c\xfe\xb2\x36\x40\xcc\x02\x18\x1b\xc2\x31\x95\xa8\x70\xf4\x29\x83\xa0\x42\x4c\x8a\xa4\x5f\xf8\x1a\x49\xfc\x72\x9a\x7d\x56\x57\xb3\xd9\xf9\x7c\xc5\xc5\x53\xe5\x22\x07\xca\xd2\xe0\x8f\x70\x67\xb5\x75\x5d\x37\xaf\x31\xc2\x2f\xae\x6e\x28\x31\xa9\x7c\xa3\xf9\xae\x58\x69\x2b\x9e\x7e\x52\x41\x1c\x67\xcb\xdb\x29\x25\xd4\xa7\xd6\xb6\x8e\x9b\x0a\xe1\x9d\xfa\x1d\xde\x89\x5a\x69\x09\xc5\x02\xf2\xbb\x38\x72\x68\xba\x8e\x56\x04\x5c\xd3\xa6\x61\x0f\x77\x58\x98\xbe\xa1\x91\x91\x75\xbf\xe1\x38\x3d\x3e\x0e\x61\xb5\xe6\x8d\x47\x66\xf8\x76\x7f\x2c\x91\x45\x70\x88\x7b\x7c\xaf\x1c\xd4\x4d\xcc\xfa\x04\xe5\x24\xa2\x9c\xbc\x2c\x3b\x86\x52\x5a\x43\xa3\x59\x43\xe8\xad\x11\xd4\x5e\x4f\x8b\xc9\xbf\x7c\xcb\xbd\x70\xaa\x09\x45\xb0\x55\xa5\xf1\xab\xe3\x02\x2f\xb2\x31\xa1\xcb\xc9\xb2\x6d\x55\x09\x1a\xcd\x81\x06\xfc\x07\x3a\xc0\xac\xeb\xda\xf6\x24\x1e\x03\xa9\xf2\x9f\x80\x7d\x11\xe5\x09\x18\xc8\xdf\x44\x53\x50\x25\x68\x42\xe2\x98\x4a\x87\xb0\x6b\xa8\xfe\x40\xed\x30\x15\xde\x0f\x90\xf2\x5a\x49\x64\x62\xcc\x7a\x59\x18\x1b\x2e\x7e\x8d\xc4\x2e\xa1\x4d\x2b\xe2\x6f\xec\x15\x30\xd6\xe0\x3c\x85\xa3\x77\xda\x77\x49\xfb\x27\x3a\xa0\xa9\x8a\x0d\xaf\xc8\xe4\x80\x6c\xb0\x29\x37\x46\x44\xea\x70\x4c\x50\xc9\x63\x8f\x2d\x77\x10\x0f\x3a\xdd\x89\x05\x48\x2b\x36\x6b\x2a\x20\xaf\x30\xdc\x6b\x8c\xc3\x8f\xbb\x07\x79\xd1\x37\x43\x06\xbf\xd1\xd1\x5e\xce\xf7\xea\x51\x99\x0b\xcd\xbd\xff\xac\x7c\xc8\x7b\xab\x8b\xc9\x49\x85\x93\x23\x4d\x74\x24\x7c\xaf\x18\x4e\xe2\xdd\x9d\x3c\xf7\x1b\x74\x6f\xb5\x4b\xb4\x12\x90\x88\x0b\x7b\x16\x49\xbb\xc8\x86\x29\x63\xf4\x96\x05\x74\xb0\x96\x9a\x55\x4e\xc9\xb1\x8d\xeb\x9b\xe5\xb7\x10\x81\xd1\xd3\xe6\x60\xe8\xba\xdb\x29\x85\xfb\xef\x8d\x43\x6a\xc0\xfc\x13\xae\x36\x55\x02\x1b\xbf\xc6\xe0\xe1\x16\x0c\x4e\x71\x67\x81\xf4\x06\x8f\x03\xc6\xae\xae\x19\x75\xd7\x60\xb5\x5f\x3f\xdc\x95\xf1\xd8\x4b\x8d\x74\x35\xe8\x8f\x49\xe5\xfa\x6c\x0b\x67\xbf\x1f\xa9\x9e\x2b\xd7\xca\xb0\xf1\x41\x3a\x7f\xb6\xae\x6f\xf9\x97\xdb\x3b\x72\xa5\xae\x3d\xd9\xfa\x70\x59\x4e\xdc\xe2\xe3\xd0\x70\x17\xfb\x7c\x74\xee\x5d\xff\x7c\xdd\xb5\x7f\x9f\xde\xe0\x76\x34\x1d\x86\x74\x96\x3d\x85\xc3\xcd\xfb\x3f\x00\x00\xff\xff\x06\x9b\x5a\x17\x08\x07\x00\x00")

func vtraceHtmlBytes() ([]byte, error) {
	return bindataRead(
		_vtraceHtml,
		"vtrace.html",
	)
}

func vtraceHtml() (*asset, error) {
	bytes, err := vtraceHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vtrace.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
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
	"alltrace.html":    alltraceHtml,
	"blessings.html":   blessingsHtml,
	"chrome.html":      chromeHtml,
	"glob.html":        globHtml,
	"logs.html":        logsHtml,
	"no_syncbase.html": no_syncbaseHtml,
	"profiles.html":    profilesHtml,
	"resolve.html":     resolveHtml,
	"stats.html":       statsHtml,
	"syncbase.html":    syncbaseHtml,
	"vtrace.html":      vtraceHtml,
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
	"alltrace.html":    &bintree{alltraceHtml, map[string]*bintree{}},
	"blessings.html":   &bintree{blessingsHtml, map[string]*bintree{}},
	"chrome.html":      &bintree{chromeHtml, map[string]*bintree{}},
	"glob.html":        &bintree{globHtml, map[string]*bintree{}},
	"logs.html":        &bintree{logsHtml, map[string]*bintree{}},
	"no_syncbase.html": &bintree{no_syncbaseHtml, map[string]*bintree{}},
	"profiles.html":    &bintree{profilesHtml, map[string]*bintree{}},
	"resolve.html":     &bintree{resolveHtml, map[string]*bintree{}},
	"stats.html":       &bintree{statsHtml, map[string]*bintree{}},
	"syncbase.html":    &bintree{syncbaseHtml, map[string]*bintree{}},
	"vtrace.html":      &bintree{vtraceHtml, map[string]*bintree{}},
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
