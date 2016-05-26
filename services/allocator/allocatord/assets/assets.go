// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package assets contains template strings and other assets for the allocatord web interface.
//
// This package is auto-generated by "jiri go generate v.io/x/ref/services/allocator/allocatord"
// which in-turn uses https://github.com/jteeuwen/go-bindata/
// Code generated by go-bindata.
// sources:
// bad-request.tmpl.html
// dash.js
// dashboard.tmpl.html
// error.tmpl.html
// head.tmpl.html
// header.tmpl.html
// home.tmpl.html
// root.tmpl.html
// style.css
// DO NOT EDIT!

package assets

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

var _badRequestTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\x41\x0a\x02\x31\x0c\x00\xef\xfb\x8a\xf8\x00\x5b\xbc\x4a\xe8\x41\xf0\x2e\xfe\xa0\x9a\x40\x17\x56\x5b\x6b\x44\x96\xd2\xbf\xdb\x36\x7a\xf0\x14\x3a\x9d\x21\xc1\x0d\xc5\xab\xac\x89\x21\xc8\x6d\x71\x13\xfe\x06\x7b\x72\x13\x00\xca\x2c\x0b\xbb\x83\x27\x38\xf3\xe3\xc5\x4f\x41\xab\x68\x42\xab\x12\x5e\x22\xad\xc3\x0d\xbb\x7f\xb1\xbd\x3b\x4e\x99\xfb\x04\x28\xc5\x9c\xbc\x84\x5a\x3b\xb5\x5f\x5c\xca\x7b\x96\x00\xe6\x98\x73\xcc\xb0\xd5\xcf\xa4\xc1\x60\xfb\xde\xd5\x6a\x34\xd2\x84\xef\xd4\x44\xb4\xba\xba\x6d\x1a\x57\x7f\x02\x00\x00\xff\xff\xe0\x6f\x86\xe4\xcd\x00\x00\x00")

func badRequestTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_badRequestTmplHtml,
		"bad-request.tmpl.html",
	)
}

func badRequestTmplHtml() (*asset, error) {
	bytes, err := badRequestTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bad-request.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dashJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x57\x6d\x53\xdb\x3e\x12\x7f\xcf\xa7\xd8\xb6\xdc\xd8\xb9\x0b\x26\xa1\x0c\xed\x84\xe9\x74\x20\x70\x25\x73\xe5\xa1\x24\xf4\x5e\x30\x4c\x47\x58\x4a\xac\x56\x96\x52\x49\x06\xd2\x9b\x7c\xf7\x5b\xf9\x51\x26\x86\xff\xfc\x79\x61\xe2\xdd\xfd\xed\xae\xf6\x49\xeb\xdd\x5d\x18\xab\xe5\x4a\xf3\x45\x62\x61\x6f\x30\x3c\x80\x59\xc2\xe0\x3b\x91\x84\xf2\x2c\x85\xa3\xcc\x26\x4a\x9b\x08\x8e\x84\x80\x5c\xc8\x80\x66\x86\xe9\x07\x46\xa3\x2d\x04\xdf\x18\x06\x6a\x0e\x36\xe1\x06\x8c\xca\x74\xcc\x20\x56\x94\x01\xbe\x2e\xd4\x03\xd3\x92\x51\xb8\x5f\x01\x81\xe3\xe9\xc9\x8e\xb1\x2b\xc1\x1c\x4a\xf0\x98\x49\x44\xda\x84\x58\x88\x89\x84\x7b\x06\x73\x95\x49\x0a\x5c\x22\x91\xc1\xd7\xc9\xf8\xf4\x62\x7a\x0a\x73\x2e\x58\xb4\xb5\xf5\x40\x34\x50\x62\x12\xf8\x04\xf3\x4c\xc6\x96\x2b\x19\xf6\xe0\x7f\x5b\x00\xa8\xec\x84\x58\x02\x99\x41\x43\x56\xc1\x82\x49\xa6\x89\x45\x2f\x12\xa2\xad\x89\x50\xc4\x81\xc7\x67\x47\xd7\xb3\x29\xc2\x6f\x91\x00\x39\xd2\xfd\x59\x6e\x05\x1b\x41\xf0\x15\x11\x32\x5e\x41\x98\x9a\x5e\xd0\x2f\x99\x14\xf5\xfe\x87\xad\x1a\x76\xcd\xe1\x14\x89\xa2\x24\xe6\xb4\x75\xbf\x53\xef\xb7\xab\x29\x84\x9a\xfd\xde\xed\x54\xfb\x6d\x69\xda\x2a\x7f\x23\xe1\x35\x75\xe3\xab\x1b\x0c\x38\x59\x30\x08\xff\xd1\xa5\x70\xba\x32\x28\x92\x4b\x5c\xc5\xb6\xad\x3b\x5e\x66\x3b\x99\xe3\xec\x2c\x91\xf5\x9a\x95\x73\x96\x2a\xbd\xaa\x0c\x7d\x39\x7e\xc1\x12\x8a\xe5\x22\xc7\x2b\xcb\x9e\x9d\x23\x65\x69\x69\xeb\xbe\xc5\x34\x31\x71\x16\x86\xbb\xc3\xc1\xde\x7e\xf3\x78\xcd\x99\x13\x6e\x7e\xfd\xa5\x2b\x4e\xe8\x25\x5f\x28\xf2\xfe\xbe\x33\xf8\xbc\x3b\xdc\x2a\x8b\xe7\xe4\xe6\xfa\x68\x36\xb9\xbc\x98\xfe\x98\x5d\xfe\x98\x9e\x8e\x2f\x2f\x4e\x5c\x29\x15\xae\x06\xc3\x24\x18\xc1\xfb\x83\xc1\xa0\x50\x1c\xec\x55\xef\xf0\x4f\xd8\x2b\x69\xfb\x1e\x6d\xbf\xa4\x1d\x78\xb4\x83\x92\x36\xf4\xc1\xc3\x0a\x3d\xa4\x9e\xc6\x0a\xfe\xa1\x45\xc4\xc7\x07\xa4\xaf\x6b\x9f\xf3\x82\xff\xf1\xdf\xc9\xc9\xec\x0c\x5d\xfd\x30\x18\x1c\xb6\x18\x67\xa7\x93\x2f\x67\x33\xe4\xec\xbd\x7f\xc6\x19\x5f\x7e\xbd\xbc\x46\x46\xf0\x6e\x30\xf8\xf8\xfe\xe3\xbf\x83\x5a\x25\xcd\xb0\xb1\xb0\xf3\x26\x72\xca\x62\x25\xa9\x41\x29\xe7\x40\x85\x4f\xb1\x81\x2d\xa3\x17\x24\x65\x0e\x1f\x54\xf4\xd8\xe8\x79\x49\x40\x4a\xd5\xbf\xc0\x9e\xac\x26\xb1\xbd\x22\x1a\x01\x96\x69\x53\x76\x74\x01\xda\xce\xb4\x40\xd0\x76\x84\xff\xc3\xde\x61\x4e\x6f\x1b\x70\x12\xd1\xd2\xa1\xc3\x40\x06\xa5\x48\x69\xcb\xe7\x39\x52\xc5\xd6\xcc\x66\x5a\xb6\x14\xbd\xf9\xf4\x09\x70\xf0\xb0\x39\xc7\x51\xe5\xa4\xd6\x5b\xc5\x58\x99\x32\x9c\x75\xd9\x12\x12\x22\xa9\x40\xf7\x70\x40\x69\x37\x55\xe4\x82\xcb\x45\x1d\x8c\xc8\x3f\x92\x41\xf5\xcb\x93\x92\x33\x56\xd2\x6a\x25\x10\x5a\x1f\x6c\x3b\x0c\xde\x55\x40\x13\xf4\xa2\x38\xe1\x82\x6a\x86\xb3\x2c\x8a\x71\x24\xfe\x0a\x9f\x0d\x37\xf7\x87\x9e\x60\x88\x70\x56\x56\x40\x10\xe4\x9e\x09\x40\xaf\x9c\xbd\xcd\xac\x44\x55\x87\x74\xa4\xab\xab\x92\x6f\xb7\x43\x37\xbd\x7b\x91\xc5\x8c\x84\xbd\xbb\xc3\x12\x9f\x2d\xb1\xc7\x98\x8b\x7d\xe3\xc9\x4d\x4e\x83\x9b\x49\x65\x04\x4f\x14\x55\x86\x76\xb8\x65\x69\x64\x98\x60\x31\x06\x17\x8f\x67\xd5\x62\x21\xd8\x58\x10\x63\xc2\xa0\xa1\x1f\xd6\xd8\xd2\xee\x6b\x62\xeb\x9e\x9f\x93\x2f\x2e\x27\xae\xf5\xf3\xe3\x17\x1e\x1a\x6f\xd4\xd7\x99\xa8\x9c\xf7\x02\x2f\x14\x5e\x69\x72\xb1\x93\x87\x0f\xbd\x33\x89\x7a\xac\x2a\xcb\x55\x5c\x5e\x2e\xa6\xee\x6b\x00\x39\xf2\x0b\xa5\x1e\x3c\xa3\xcd\xc8\x56\x3c\x57\x6a\xa3\xfc\x59\xf8\x7e\x58\x1b\x67\x5a\x2b\xbd\x93\x9a\x05\x1a\x4e\x38\x65\x95\xe1\xed\x88\xfc\x24\x4f\x78\x6c\x4b\xac\xf9\x1c\xc0\xbf\x90\x52\xd4\x6d\xe1\x4e\xaf\x57\xaa\x06\x88\xa8\x92\xac\xa9\x10\x17\x85\xa6\x4a\x9a\x84\x8d\xf3\x60\x14\xec\xc3\x9a\xbb\xf6\xf4\xcc\x09\x17\x5e\xa5\xf9\x2a\x9e\xbb\xea\xc7\xe8\xb9\x9a\x58\xa5\x4b\x81\x8d\xdb\x55\xb4\x2f\xc5\xdc\x3f\xfa\x66\x76\x6f\xca\x7c\x12\xdc\x35\x5e\xcc\xa9\x7f\xbe\xd2\x5e\x71\xb9\x47\xd8\x9f\xa7\x24\x4e\x1a\x7f\x72\x1d\x8d\x53\xf9\x2c\xc2\xfc\x4a\xf6\xe8\xb9\xb9\x50\x0a\xab\x2f\x7a\xe0\x26\x23\x82\xff\x29\x9a\xfa\x48\x33\x92\x1b\x0a\xa9\x8a\xb3\x94\x49\x1b\x2d\x98\x3d\x15\xcc\xfd\x3c\x5e\x4d\x68\xa1\x3b\xe2\xb4\x57\x9f\xc6\xa9\x57\xcb\xbc\xb5\xbd\x22\xaa\x6f\xb2\x02\x90\xbf\xf4\xdb\xbc\x19\xf6\xdd\xd4\x6d\x46\xa3\x56\xfc\xe6\x38\x3f\xa6\xfc\x8f\xbb\xa2\xf6\x1b\xc4\xba\xf9\xf9\xc8\xa9\x4d\x46\xfe\xa4\x6f\x78\x09\x73\xbb\xda\xa8\x35\xed\x1b\xae\x60\xb8\x26\xd1\xb6\xb9\xa5\x32\xdc\x39\x8f\x57\xa6\xc4\x42\x0b\xba\x2c\xe2\xde\xc7\x99\x69\xe3\x06\xed\x57\x6c\x02\x25\x94\x1e\xf9\xb7\x89\xc7\x5e\x77\xa9\x7d\x38\x7a\xe2\x4e\x6b\xca\xe5\x77\x22\x32\x3c\xf1\xc0\xe3\x26\x25\xd7\xd3\xd2\x08\x62\x2e\xdd\xfe\xc7\xf2\x7a\x88\xce\xb9\x9c\x71\x9c\xe9\x78\x77\x0e\x06\x83\x5e\xdf\x87\x90\xa7\x6e\x08\x79\x7a\x09\xb2\xd0\x9c\x0a\xbc\x15\xcc\xe6\x09\x71\x2a\x8c\x60\x67\xd8\x6f\x91\x33\xc9\xed\x86\xac\xdb\x54\x56\x8e\x8a\xd5\x99\x12\x44\xdd\x06\xe7\xe7\xe7\x40\x69\x70\xb7\xee\x3f\x93\x4c\x70\x7f\x6e\x8b\x26\xa3\x34\x05\x12\xf4\x21\x48\xf0\xdf\x26\x22\xe5\x42\x70\x53\x0c\xa1\x2e\xe0\xdd\xba\x05\x68\xe1\xbd\x54\x94\xbf\xd6\x55\x29\xc7\x11\xd5\xe4\x31\xc4\x22\x71\xbb\xf5\x8c\xdc\x8b\x22\x5a\xb7\x45\x11\x97\xbb\xd7\x5d\xbf\x2c\xea\x7c\x93\xea\xf5\xab\xea\x7f\x61\x72\x17\x7b\xb9\x81\x5a\x25\xa8\xfb\x9f\x38\xeb\x61\xae\x55\x9a\x2f\xfd\x0b\xfe\xc0\x24\x16\x22\x97\xcf\x3a\xbf\xe5\x48\xc1\xef\x17\xfb\x9b\xbf\x32\x50\x0b\x45\x77\x77\xf7\x74\xa3\xa1\x74\x90\xda\x88\x50\x3a\x56\x22\x4b\x65\x18\xb8\xd9\x62\xb1\x12\x5c\xb0\x83\x4e\x09\x99\xa5\xf7\x4c\xfb\x7c\x3e\x87\xd2\x9d\x66\xca\x14\x98\x6b\xf5\x68\x4a\x56\x94\x92\x65\x33\x93\x96\xd6\x9f\x92\xe5\x3e\x72\xdb\xca\x52\x5d\xa0\x4b\x1b\xb9\xda\xc4\xbb\x21\x5d\x76\x15\x68\xb9\xc3\xc2\x67\x40\xc9\xbc\xba\x51\xaa\x20\x8d\x6a\x52\x2d\x5f\x5f\xec\xeb\x6a\x66\xad\x81\x09\x5c\x2c\x36\x3c\x0f\x6f\x6b\x17\x30\xab\x83\xbb\x4a\xdc\x5f\xa1\xa8\xad\xb3\x5b\xa7\x89\x63\x03\xd4\x97\x80\x8b\xcd\x9b\x8e\x1d\xaf\x39\x3e\xfa\x89\x03\x36\x78\x2b\xdf\xba\xaf\x45\xfc\x4a\xca\xb8\xc6\x8f\x38\xb7\x66\x71\x89\x87\x96\xf8\x2d\x59\x5e\xc0\x20\x11\xdf\x6c\x0e\x85\x0f\x9e\x53\xdb\x5d\x97\xd0\x8b\xdb\xd8\xe6\x86\x53\xbc\x37\x0b\x0e\xc3\x2f\xd7\x95\x1b\x33\x99\x65\x51\xa3\x6e\x82\xbe\xe8\x07\x22\xc2\x02\xd9\x07\xdc\x7e\x31\x27\x1b\xf5\x5e\xc6\xa8\x0c\x04\x46\x65\x94\x3f\x8b\x05\x7d\xed\x0c\xfe\x3f\x00\x00\xff\xff\xf4\xd9\x7c\x20\x77\x0f\x00\x00")

func dashJsBytes() ([]byte, error) {
	return bindataRead(
		_dashJs,
		"dash.js",
	)
}

func dashJs() (*asset, error) {
	bytes, err := dashJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dash.js", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dashboardTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x4d\x4f\x1b\x3d\x10\xbe\xf3\x2b\xfc\xee\x65\x79\x45\xd6\x4e\x80\xb6\x52\xba\x49\x55\x01\xad\x90\x10\xa0\xd2\x4b\x55\xf5\xe0\xd8\x93\xac\x13\xef\x7a\xb1\x67\x03\x2b\xca\x7f\xaf\xf7\x23\x21\x81\x10\x11\xa9\x87\x2a\x87\x8c\xc7\xcf\x3c\xf3\xed\x24\xfe\xef\xf4\xea\xe4\xfb\x8f\xeb\x33\x92\x60\xaa\x87\x7b\x71\xf3\x45\x48\x9c\x00\x97\x95\xe0\x45\x54\xa8\x61\xf8\xf0\x40\x6f\xc0\xce\xc1\x5e\xf2\x14\x1e\x1f\xc9\x29\x77\xc9\xc8\x70\x2b\x63\xd6\x00\xf6\x1a\xb4\x56\xd9\x8c\x24\x16\xc6\x83\x90\xb1\xb1\xc9\xd0\xd1\x89\x31\x13\x0d\x3c\x57\x8e\x0a\x93\x32\xe1\xdc\xa7\x31\x4f\x95\x2e\x07\x37\xa6\xb0\x02\x0e\x4e\x8c\x84\x83\x6b\x6b\xfa\xc7\xdd\x6e\xe7\x5d\xb7\xfb\xfb\x9b\x19\x19\x34\x7d\x2f\x76\xbc\x4a\x21\xd7\x4a\x74\x8e\x9a\xcb\xa7\x53\x2b\x79\x44\x58\xfb\x26\xc4\x82\x1e\x84\x0e\x4b\x0d\x2e\x01\xc0\x85\x1a\xcb\x1c\x06\x21\xc2\x3d\x56\xce\xc3\x36\x2f\x27\xac\xca\x91\x38\x2b\x06\x41\x82\x98\xbb\x3e\x63\x7c\xca\xef\x9f\xc7\x5b\xe9\x98\x56\x23\xc7\xa6\xb7\x05\xd8\x92\x1d\x52\xff\x69\x0f\x34\x55\x19\x9d\xba\x60\x18\xb3\x86\x6f\x0b\xb9\x90\xd9\xd4\x53\x6a\x53\xc8\xb1\xe6\x16\x36\xb2\x47\x85\xd5\x51\xce\xad\x03\xeb\x1d\x1d\xd1\x1e\xcb\xbd\xe6\xef\xbb\x11\xc6\xcc\x14\xb0\x1e\x3d\xf6\x2e\xda\x5c\x1a\xdd\x6b\xbe\x9e\x3a\x1b\x38\xe4\xa8\x04\xab\x2b\x4d\x7d\x49\x83\xba\xf2\xc1\x53\xe5\x83\x0d\xf1\xb5\x46\xd2\x0f\xce\x96\x54\xea\x5e\x05\x75\xaf\xa6\x7c\xce\x1b\x6d\xb0\x9e\xe1\xdd\xdd\x1d\x9d\x34\x74\xcd\x44\x25\xdc\xa2\x63\xda\x70\x09\x76\x67\xee\x61\x3b\x25\x4d\xdb\x69\x43\x46\x2b\xb2\xfd\x50\x14\xd6\x42\x86\x61\x87\x3c\xe4\x5c\xcc\xf8\x04\x5c\x9f\xfc\x0c\x85\xb1\x50\xe3\xc2\x5f\x8f\xff\x7f\xdc\x68\xef\x00\xaf\xb2\x0b\x4f\x72\xc2\xb5\x1e\x79\xdb\xfd\x3a\x71\x95\x29\x6c\x2d\x56\x63\x8c\xd9\x62\xe1\xe2\x91\x91\x65\x1b\xb6\x54\x73\xa2\xe4\x20\x10\x7e\x8d\xb8\xca\xc0\x2e\x63\x5d\x5e\x55\x66\x2b\xfa\x95\x9b\x7a\x2b\x83\x6d\x7b\xeb\x91\xeb\x76\x15\xf8\x3c\xf3\x85\xcd\x84\x87\x6e\x00\xd4\xc4\x90\x72\xa5\x6b\xe2\xb3\x4a\x7a\x06\x5c\x3f\x2c\x33\xa8\x8b\xf2\x3c\x4c\xa1\xb9\x73\xed\x65\x50\xe3\x34\x47\xc8\x44\x59\x35\xf0\xa5\xf3\x97\xf0\xdb\xdc\xbd\x15\x2a\xf2\x22\x2a\x9c\x6f\x60\x94\x0b\x7c\xab\x51\x0a\x69\x6b\x34\x2a\x11\xde\xec\x4b\x2a\x37\xdb\xcd\x8e\x8c\x95\xd6\x55\x1f\xff\x15\xe0\xca\x61\x55\x5c\xf4\x53\x16\xd6\x6f\x9f\xc9\x5c\xb4\x65\x36\xab\x15\x52\xd9\x24\xd2\x7c\x04\x7e\x60\x2e\xae\x3e\x9f\x9e\x5f\x7e\xa5\x94\x6e\x9e\x91\x25\xe7\x2b\x63\xb2\xb8\x8f\x14\x42\x4a\x1c\x68\x10\x08\x32\x18\xf6\x92\x2d\x29\xae\x19\x05\xc3\xc3\x1d\xb0\xc7\x3b\x60\xdf\xef\x80\xed\xed\x12\x44\x6f\xd3\x9e\xbe\x82\xfd\x20\x77\x69\x21\x58\x6b\x6c\x94\xba\xc9\xb2\xdc\x5f\xfc\x36\x83\x24\x68\xfc\x6b\x8e\x56\xc1\x1c\x88\xe4\xc8\x29\xb9\xf6\xbf\x86\x0e\x08\xda\x92\xf0\x89\x6f\x36\xa9\xd6\xd4\xd2\x75\xea\x98\x35\x2f\x97\x7f\xca\xea\x3f\x11\x7f\x02\x00\x00\xff\xff\x2f\xfc\x99\xc7\x5c\x08\x00\x00")

func dashboardTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dashboardTmplHtml,
		"dashboard.tmpl.html",
	)
}

func dashboardTmplHtml() (*asset, error) {
	bytes, err := dashboardTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dashboard.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errorTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8f\x41\x0a\xc2\x30\x10\x45\xf7\x3d\xc5\xd8\xbd\x0d\xdd\xca\x18\x70\x21\xb8\x10\x04\xc1\x03\xd4\x66\x24\x85\xb6\x29\x43\x44\x4a\xc8\xdd\x4d\x32\x76\xe1\xea\x93\xff\xdf\x83\x09\xee\x8c\xeb\xfd\xba\x10\x58\x3f\x8d\xba\xc2\x2d\xa8\x33\xba\x02\x40\x3f\xf8\x91\xf4\x99\xd9\x31\x2a\x79\x54\xa8\x64\xc6\xa7\x33\x6b\xa1\x6c\xab\x4f\x33\x14\x0a\x6e\x7d\xff\x66\x26\x93\xa8\xb6\x8c\x0b\x53\x4e\x80\x10\x9a\xc7\xfd\x1a\x63\x2e\xd5\xaf\x0d\xe1\x33\x78\x0b\x8d\xb8\x7b\x19\x17\xe1\x4b\x77\xc8\x5a\x8c\x8d\x48\xa2\xd0\x6c\xfe\x40\xec\xc0\x32\xbd\x8e\x75\x42\x2f\x6e\xa2\x18\x6b\x9d\x13\x55\xa7\x37\x0f\x95\x5c\x9b\xce\x2a\x5f\xfc\x06\x00\x00\xff\xff\x96\xa3\xfc\xdf\xfa\x00\x00\x00")

func errorTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_errorTmplHtml,
		"error.tmpl.html",
	)
}

func errorTmplHtml() (*asset, error) {
	bytes, err := errorTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "error.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _headTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xdf\x8e\xa2\x3c\x14\xbf\xff\x9e\xc2\xf4\xbb\xf0\x62\x44\x10\x15\xc7\xc9\x32\x9b\xcd\x5c\xed\xdd\x64\x77\x5f\xa0\x94\xa3\x36\x03\x94\xd0\xa2\x75\x59\xdf\x7d\x4f\x5b\x23\xa3\xce\xcc\x0e\x44\x13\x02\x9c\x9e\xdf\xbf\x42\x38\x36\x4d\x0a\x2b\x5e\xc0\x80\x6c\x80\xa6\xe4\x70\xf8\x6f\x30\xf8\x92\x83\xa2\x78\xc6\x5f\x41\x73\x88\xc9\x96\xc3\xae\x14\x95\x22\xae\xc8\x44\xa1\xa0\x50\x31\xd9\xf1\x54\x6d\xe2\x14\xb6\x9c\x81\x67\x6f\x46\xae\xe3\xf4\xe3\x05\x57\x9c\x66\x9e\x64\x34\x83\x78\x72\xb9\x9c\x53\xcd\xf3\x3a\x7f\x6f\xb9\x96\x50\xd9\x35\x9a\xe0\x72\x21\xae\xe0\x48\x9f\x23\x7b\xcd\xc9\xe3\x5b\xbe\x69\x59\x66\xe0\xe5\x22\xe1\x78\xda\x41\xe2\x61\xc1\x63\xb4\x34\x74\x97\x59\xf6\x20\x91\xe4\xd3\x2c\x52\x51\x55\x4b\x2f\xa1\x68\x50\xed\xaf\xe9\x92\x8c\xb2\x17\x43\x68\x18\x33\x5e\xbc\x0c\x36\x15\xac\xe2\xa1\xef\xaf\xb0\x45\x8e\xd7\x42\xac\x33\xa0\x25\x97\x63\x26\x72\x9f\x49\xf9\x75\x45\x73\x9e\xed\xe3\x9f\xa2\xae\x18\xdc\x3d\x89\x14\xee\x9e\x2b\xf1\x30\x0b\x82\xd1\x3c\x08\xfe\xfc\x10\x89\x50\xe2\x01\x2f\x47\x58\xe2\x8a\x66\x9c\x8d\xa6\x6e\xb1\xbd\x3b\x5e\x61\xc7\xd0\x3a\xaa\x20\x8b\x87\xd6\xa1\xdc\x00\x28\x57\x54\xfb\x12\xe2\xa1\x02\xad\x8c\xf0\xf0\xb1\x35\x69\xda\x49\xdb\x4e\x9c\x6b\xd2\x34\xe3\x6f\x52\x82\x92\xcf\x78\xcb\xf5\xe1\xe0\xf3\x14\x63\x72\xb5\x1f\x23\x81\xdb\xfc\x36\x24\x31\x9b\xc3\x99\x6f\x79\x6c\xc3\x15\xef\xa5\xa4\xdb\x63\x25\x6a\xb6\xf1\x38\x6e\x22\x19\x48\xfe\x1b\x64\x4c\xe6\x0b\x3d\x5f\xbc\x6f\x63\x45\xb7\xa6\x5d\xfa\x97\x04\x9e\x05\x8e\xcb\x62\xfd\xca\xdd\xc7\x52\x93\xc9\x4c\xe3\xd1\x47\xec\x08\xed\x24\xb7\x08\xf5\x22\xec\x23\x66\x81\xdd\x92\xcd\xd0\xde\xac\x5f\x32\x07\xed\x24\x17\x05\x3a\x0a\xfa\x88\x59\x60\xb7\x64\x61\xa0\xf1\xe8\x95\xcc\x41\xbb\x3d\xb3\x48\x2f\xa2\x5e\xcf\xcc\x00\xbb\x25\x9b\x87\x1a\x8f\x5e\xc9\x1c\xb4\x9b\xdc\x3d\xee\xc6\x7d\xbf\x8d\x74\xd0\xb7\xe4\x9c\x84\xfd\xdc\x10\xfc\x52\xaf\xc1\x37\x4d\xff\xd6\x38\x5e\x78\x93\x25\x26\x59\xba\x24\x27\xa7\xae\x76\x73\x29\x7c\xf9\x26\xc7\xb7\xef\x24\xe5\x6a\xb7\x96\x5a\x46\x7a\x19\x9d\x09\xd9\xca\xed\x13\xa1\xf9\x8b\x3c\x58\xb9\xb5\xcc\x34\xd4\xd3\xf3\x27\x64\x2b\xed\x44\x3e\x8e\xd1\x5c\x9a\x37\x87\x33\x1c\x0d\x88\xfa\x85\xc3\xf4\x49\x64\xa2\x22\xed\xcc\xfc\x3f\xa5\xf3\x69\xc8\x3e\x81\xfc\x6e\x7c\xbe\x42\x7e\x60\x33\x97\xca\xcc\xed\xf3\xef\x58\xd3\x40\x91\xe2\x7f\x9d\xbf\x01\x00\x00\xff\xff\x52\x29\x6b\xf4\xfe\x08\x00\x00")

func headTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_headTmplHtml,
		"head.tmpl.html",
	)
}

func headTmplHtml() (*asset, error) {
	bytes, err := headTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "head.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _headerTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x31\x6e\xc6\x20\x0c\x46\xf7\x9e\xc2\x72\xe7\x86\x0b\x90\x48\x1d\xba\x76\xa9\xd4\xdd\x0a\x26\x41\x22\x50\x01\xcd\x62\x71\xf7\xd2\x26\xa4\xbf\x94\x0d\x7f\xcf\x96\x1f\x16\x31\x6c\x5d\x60\xc0\x95\xc9\x70\xc2\x5a\x9f\x00\xf4\x51\x4c\xed\xd9\x8a\x40\x3b\xcc\x9e\x72\x1e\xd1\xb3\x2d\x78\xc4\x0d\x10\xac\x89\xed\x88\xcf\x78\xf1\xb8\x44\x9c\x3e\x29\x90\x71\xdf\x9b\x56\x74\xf5\xe6\x2f\x0a\xbd\x2b\x73\xda\xdd\xcc\x2f\x81\x36\xc6\x49\xa4\xb8\xe2\x19\x86\x8f\x16\x73\x7a\x6f\x61\xad\xf0\xea\x7d\x9c\xa9\xc4\xa4\xd5\xef\xe8\xa9\xa2\x9a\xcb\xdd\x2a\xb9\x65\xfd\xd7\x12\x71\x16\x86\xb7\x8d\x9c\xff\xfb\xcc\x4d\xb6\x6d\xec\xf8\xc1\x50\x84\x83\x39\x07\xae\x3d\x5a\xf5\x4b\x74\xfc\x13\x00\x00\xff\xff\xe2\x89\x41\x0f\x31\x01\x00\x00")

func headerTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_headerTmplHtml,
		"header.tmpl.html",
	)
}

func headerTmplHtml() (*asset, error) {
	bytes, err := headerTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "header.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _homeTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x57\xdd\x6e\xdb\x38\x13\xbd\xf7\x53\x4c\xf5\x05\xb0\xfc\xc5\xa2\x52\xf4\x2e\x91\x0d\x34\xdb\x02\x5b\xa0\xce\x16\x6d\xba\x40\xaf\x16\xb4\x44\x5b\x6c\x24\xd2\x4b\x52\x76\x05\x43\xef\xbe\x33\xfa\xb3\x95\xd8\xd9\x06\xdd\x02\xb5\x4d\xf2\xf0\xf0\xcc\x0f\x87\x93\xe8\x55\xa2\x63\x57\x6e\x04\xa4\x2e\xcf\xe6\xa3\xa8\xfb\x12\x3c\x99\x8f\x00\x22\x27\x5d\x26\xe6\x7f\x72\xc5\x13\x59\xe4\xb0\xdf\xd7\x13\xc0\xbe\x08\xb3\x15\xe6\x8e\xe7\xa2\xaa\xe0\x6d\x96\xe9\x98\x3b\x6d\xa2\xb0\xc1\xe3\x4e\x44\x8a\x7c\x93\x71\x27\xc0\x23\x36\x0f\x58\x55\x8d\xa2\xb0\x61\x8e\x5e\x05\x01\xdc\xfd\x71\xff\xfe\x1a\x76\x02\xc6\x4b\x6d\x8c\xde\x8d\x81\x43\xa6\x1d\xe8\x15\xb8\x54\x80\x75\x65\x26\xd5\x9a\x86\x32\x11\x0a\x99\xcb\x64\x0a\xa9\x50\xb1\xa0\xf5\x91\x15\x71\x61\x70\x32\x30\x82\x8e\x49\xc0\x09\x93\x4b\xa5\x33\xbd\x2e\x41\xaa\x9e\x43\x80\xc8\x44\x8e\x04\x76\x0a\x99\x7c\x10\x3d\x1b\x70\x95\x8c\x96\x99\xb0\x16\x8f\x61\x10\x04\xa8\x6b\xa9\x93\x12\xe2\x8c\x5b\x3b\xf3\x3a\xdc\xc6\xe8\x2d\xfe\x36\x41\xc6\x4b\x5d\x38\xe0\x9d\xb9\x89\x77\xca\x52\x61\x1a\x5b\xd1\x7b\x36\x36\x72\xe3\x08\x04\xb0\x2a\x54\xec\xa4\x56\x10\xa7\x5c\xad\xc5\xad\x53\xbe\x44\x7b\x94\xd8\xfd\x7e\xbf\xf8\x38\x81\x7d\x8d\x02\xc0\x80\x14\xa4\x96\xad\x85\x7b\xdf\x08\xbf\x2d\x3f\x24\x08\x9e\x30\xa9\x94\x30\x04\x87\x19\x78\xd1\x4a\x2b\x07\x31\xda\x6b\x66\xe3\xb5\xe1\xe5\x78\xee\x5d\xb6\x74\x97\x5e\x14\xd2\xf2\xdc\xbb\xa9\x69\x6b\x35\xe1\x41\x4e\x94\xc8\x6d\x67\x66\xce\xa5\xf2\x1a\x8d\x51\xfa\xba\x9b\xdd\xf0\xb5\x08\xea\xc0\xcd\x7f\x33\x82\x8c\x43\x6f\xc1\x02\x13\x61\x2d\xe0\x83\xb2\x8e\x63\x1c\x2c\x05\xe7\x64\x4e\x60\xa0\x5f\xb7\x9c\x47\x47\x75\xce\xb6\x41\x26\xad\x6b\x0f\x3d\x07\x69\x7d\xd9\x81\x6a\x75\xf3\x6f\xba\x30\x18\xdc\xf6\xf8\xc3\x29\x64\x1e\xb2\x34\x83\xfd\xde\x90\x8b\xe1\x42\xaa\x44\xfc\x98\xc2\x45\x9b\x00\x70\x3d\x03\xd6\x6b\xaf\x43\x74\xfe\x78\x89\x51\x3d\x3e\xfc\x04\x28\x48\x84\xe3\x32\xb3\x47\x30\x52\xf9\x66\xbe\xdf\x3f\x14\x4b\x41\x9e\x00\xd6\xfb\xe3\xcd\x00\xb5\x79\x42\x16\xf3\x2d\xfa\x79\x48\x46\x39\xb4\xe1\xaa\x09\x01\x25\xcf\xbd\xcc\x05\x06\x92\xe6\xa2\xa5\x09\x4f\x60\x3b\xde\x42\xc9\x1f\x0e\xd1\x1e\x24\xdc\xf1\xa0\x1b\xce\xf6\x7b\xd6\xb1\x11\x19\xfb\x8a\x0b\x55\x35\x7f\x3c\xfd\xc5\x19\x14\x45\xc2\xeb\xc3\x8e\xa5\x87\x9b\x5f\xb0\x64\xa1\x0b\x0c\x04\x39\xe5\xbc\x19\x92\xd4\x10\xe4\xb3\xd6\xae\xaa\xc2\x28\x94\xf3\x68\x47\x38\x8c\x2c\x5d\x76\xb9\x1d\x38\xf7\x3f\x13\x77\xdb\x82\xe0\x13\x77\x58\x4d\x94\x3d\xab\xb1\x4b\x31\xd6\x6d\xe9\x76\x0c\xd4\x34\x48\x2c\x06\xa7\x08\x84\x4a\xce\x4b\x3f\xca\xe6\x7a\x78\x9c\x7d\x4e\xd9\x00\xaf\xfd\x30\xeb\x86\x80\xc7\xd6\x71\x48\x8d\x58\xcd\x3c\x14\xf3\x8e\xdb\x74\xa9\xb9\x49\xbe\x7e\xfe\x58\x55\x5e\xbf\xa9\x70\x4e\xab\x60\x83\x03\xf4\xae\x07\x8e\x1b\x2c\x40\x33\xef\xaf\x65\xc6\xd5\x83\x37\xef\xb7\x45\x21\x1f\x66\xc3\x40\xe8\x8b\x94\x88\x65\xb1\x7e\xa1\x0a\xda\xf2\x42\x05\x80\xff\xf1\xa6\x5a\x67\x74\xe9\x61\xf1\x9f\x79\xed\x00\x4b\xf0\x7e\xdf\xd4\x08\x94\xf0\x9c\xce\x1a\xfe\xbc\x52\xad\xe2\x4c\xc6\x0f\x33\xef\x50\xde\xc7\xa7\xce\x19\x4f\x61\xdc\x12\x52\xa2\xf9\x8e\x3f\x60\x19\xe5\xb0\x12\x3b\xc0\x17\x4d\xab\xc4\x4e\x80\x31\x36\x9e\x90\xb5\x35\xee\x79\x7b\x07\xc3\x41\x11\x14\x99\x3d\x5c\x8e\x3b\xad\x04\xac\xf0\xf2\x25\x6c\xf4\x38\xff\x6a\x87\x91\x63\xe2\xba\xd2\xdf\x52\xd4\x8e\x1d\xd0\x3c\x00\x2f\xb7\xbf\xa7\x23\xa3\x9b\xea\xf2\x6f\x26\xb7\x6f\xcd\x9d\xd8\x91\xd5\x03\x73\x76\xd2\xa5\xc0\x16\x78\xdd\xe8\x09\x0a\x1e\x6b\xcf\x9b\x85\x3e\x8e\xf5\xbd\x1b\x3d\x75\xca\xc1\xea\xe6\x85\x1c\x1d\x01\xda\x07\x1b\xac\x89\x6b\xc3\xdf\x5a\x2b\x9c\xfd\x84\x7e\xa0\x1a\x19\x76\x0d\x41\x98\xeb\xfa\x85\xfe\x8e\x55\x64\xf0\xaa\xfe\xe4\xf6\xef\x7f\x17\xc2\x94\xe7\xb6\xd3\xcf\xbe\x59\x40\x02\x2a\xc7\xf7\xe2\x87\xf3\xe9\x15\xeb\xfa\x84\x2d\x37\x40\xe5\x1c\xdf\xb2\x7c\x83\xdd\x00\xad\x31\x2a\xf4\xfe\xa1\xf0\x4f\x6e\x7a\x68\x8e\x90\x46\xb4\xdf\xef\xfa\xff\xeb\xab\xab\x2b\x76\x75\x84\x6a\xba\xa5\x01\x59\x3d\xd5\x31\xc9\x15\xf8\x2d\x66\x86\x0d\x08\x5f\x5a\x9d\x15\x0e\x97\xfb\xe6\xa5\xde\x49\x3d\xa4\xef\x75\x09\x34\xfe\xdf\xb8\x4f\x8e\xb1\x11\xae\x30\x0a\x56\x1c\x53\xf3\x06\xfb\x15\xb8\x84\x9c\xad\xb4\xc9\xb9\xf3\xbd\xc5\x62\x01\xef\xde\x4d\xe1\x1b\xfe\x83\xf4\x3a\xcf\xaf\x2d\xa6\x09\xb2\x5f\x62\xb3\x83\xd9\xd0\xc9\x80\xa7\x02\xa7\xe0\xad\x8c\xce\xef\xf4\xae\x03\x55\x40\xe9\xff\x8b\xc2\x1a\x4a\xff\x27\x15\x1c\xfc\x71\x68\xba\xaa\x11\x7e\x5c\xf8\x5d\x53\x37\x61\x98\xdf\x49\xe9\x77\xf1\xf5\x3b\xd7\x5d\xf8\x1e\x3b\x04\x8e\x09\x1e\xa7\x4f\x41\x00\x61\x08\xb5\x60\xba\x47\xd4\xdf\x1e\x52\xc0\xa6\xba\xc8\xb0\x07\xd6\xeb\x75\x56\xf7\xc7\x90\x48\x8b\x5d\x69\x09\x8d\x7b\x59\xcb\x70\xe1\xbb\x54\xda\x09\xab\x69\x06\x47\x0c\x72\xad\x85\x4d\x6e\xa0\xea\x6d\x3e\xb9\xde\x58\x5a\x7f\x37\x9f\x87\x8c\x8e\x42\xea\xa6\xe9\xbb\xf9\xab\xe2\x9f\x00\x00\x00\xff\xff\x57\x95\x4e\xc2\x6d\x0c\x00\x00")

func homeTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_homeTmplHtml,
		"home.tmpl.html",
	)
}

func homeTmplHtml() (*asset, error) {
	bytes, err := homeTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "home.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rootTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\x41\x6b\x84\x30\x10\x85\xef\xfe\x8a\x69\xee\x1a\xf6\x9e\x15\x4a\x2f\x2d\xb4\xbd\x14\x7a\x9f\x35\xb3\x1a\x88\x49\x88\xa3\x20\xe2\x7f\xdf\x51\xd7\xdb\x9e\x26\x19\xde\xfb\x78\xf3\xcc\x9b\x8d\x0d\xcf\x89\xa0\xe3\xde\xd7\x85\x39\x07\xa1\xad\x0b\x00\xc3\x8e\x3d\xd5\xff\x18\xd0\xba\xb1\x87\x65\xd9\x17\x50\xfd\x51\x9e\x28\xff\x62\x4f\xeb\x0a\xef\xde\xc7\x06\x39\x66\xa3\x0f\xbd\x38\x45\x49\x7d\xf2\xc8\x04\x6a\xa3\x29\xa8\xd6\xb5\x30\xfa\x20\x9b\x5b\xb4\x33\x34\x1e\x87\xe1\xaa\x9c\xa5\x20\xbe\x39\xe5\x38\xc9\x3b\x97\x1e\xe7\x38\x32\xe0\x89\xb5\xea\x15\x91\xf2\xc1\x94\x94\xd6\x4d\x27\xac\x47\x17\x76\xb9\xac\xbb\xcb\xb9\x4d\xd8\x52\xb9\xc7\xa8\x3f\x32\x6d\x08\x0c\x16\x7e\xe4\xac\x96\xe0\x2b\x0c\x8c\xa1\xa1\x01\xe2\xfd\xf5\x85\x12\xfb\xf2\x64\xa6\x63\xca\x0b\x4f\xf8\x6d\x64\x8e\xa1\x4c\xf2\x71\x13\x29\xe8\x32\xdd\xaf\x6a\x59\xaa\xcf\xb8\x99\x55\xfd\x1d\x5b\x70\xc1\x68\x7c\x32\xf4\x0e\x31\x5a\x62\x4b\x15\x7a\xeb\x62\x9b\x47\xf7\x8f\x00\x00\x00\xff\xff\xe3\x95\x3a\xa7\x93\x01\x00\x00")

func rootTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_rootTmplHtml,
		"root.tmpl.html",
	)
}

func rootTmplHtml() (*asset, error) {
	bytes, err := rootTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "root.tmpl.html", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _styleCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x56\x4d\x6f\xe3\x36\x13\xbe\xfb\x57\x10\x08\x5e\x20\x59\x44\x5e\xc9\x71\x62\xaf\x73\xca\xe6\xdd\x00\x05\x8a\x1e\xba\x6d\xef\x94\x38\xb6\xd8\xa5\x48\x83\xa4\x62\x3b\x41\xff\x7b\x87\x1f\xfa\xb4\x9c\x3d\xd4\x08\x1c\x6b\x34\xdf\xf3\xcc\x43\x7e\xfe\x44\x9e\xd5\xfe\xa4\xf9\xae\xb4\x64\x91\x66\x0f\xe4\x8f\x12\xc8\x5f\x54\x52\xc6\xeb\x8a\x3c\xd5\xb6\x54\xda\xcc\xc9\x93\x10\xc4\x2b\x19\xa2\xc1\x80\x7e\x05\x36\x27\x9f\x3e\xcf\xd0\xfe\x4f\x03\x44\x6d\x89\x2d\xb9\x21\x46\xd5\xba\x00\x52\x28\x06\x04\x1f\x77\xea\x15\xb4\x04\x46\xf2\x13\xa1\xe4\xeb\xf7\xff\x27\xc6\x9e\x04\x44\x43\xc1\x0b\x90\x68\x6c\x4b\x6a\x49\x41\x25\xc9\x81\x6c\x55\x2d\x19\xe1\x12\x85\x40\x7e\xfd\xe5\xf9\xdb\x6f\xdf\xbf\x91\x2d\x17\xe0\xa3\xcd\x72\xc5\x4e\xe4\x7d\x46\x50\x4f\xda\x64\x4b\x2b\x2e\x4e\x1b\xf2\xbb\xca\x95\x55\x8f\x8d\xf8\x00\x2e\xd1\x0d\x59\xa6\x69\x2b\xf3\x71\x37\x44\x2a\x5d\x51\xd1\x4a\x0b\x25\x94\xde\x90\xab\xbb\xd5\x72\xb5\x7c\x71\x62\x34\xce\x7f\x70\x74\xed\x8d\x2a\xa5\xb0\x2c\xb9\xdb\x10\x2a\x2d\xa7\x82\x53\x03\xac\xf3\xc9\xdf\xd0\x65\xb6\xd8\x1f\x9d\x48\x70\x09\x49\x19\x43\x2f\xd2\x20\xdc\x53\xc6\xbc\x7d\x7c\xae\xa8\xde\x71\xd9\x3e\xba\xf6\x6c\x85\x3a\x6c\x48\xc9\x19\x03\xf9\x38\xfb\x67\x36\x9b\x53\x21\x54\x41\xad\xd2\xcc\x97\xda\x29\xd1\xda\x55\x39\x52\x99\x57\x14\xdb\xf5\xde\x73\xbe\x5e\xee\x8f\x64\xb1\x84\x8a\x2c\xd6\xee\xd7\x7a\x94\x8b\x97\xde\x3f\x34\x19\x1d\x93\x03\x67\xb6\xdc\x90\x2f\xe9\x20\xcb\x44\x87\x5a\x32\x2f\x1d\x47\xcd\xad\x34\xae\x7f\x3e\x72\x74\x9d\x58\xb5\x47\xfd\x87\x41\xbc\xd6\x4d\x14\xc7\x60\xd9\x5d\x0c\xb6\x15\x70\xc4\x86\x78\x0b\x65\xb8\xe5\x0a\x4b\xd0\x20\xa8\xe5\xaf\xf0\x41\xdc\x79\x5e\x5b\xab\x64\xb2\xa7\xc6\xa0\xa6\xcf\xa3\xf1\x9d\x45\xdf\x16\x8e\x36\xc1\xb1\xed\xd0\x25\x62\xcd\x82\xee\x95\xe7\x93\x5d\x5e\xa8\x2d\x61\x60\xac\x56\x0e\xb6\xef\x83\xcc\x68\x6e\x94\xa8\x2d\x38\x47\x88\x3a\xab\xaa\x6e\xda\x0d\x9c\x5e\x5e\x96\x2f\xd9\x53\xd0\xd0\x0c\x50\x94\x61\xc7\xd1\x8e\xb3\xfe\xcb\x06\x6b\x56\x53\xd9\xb8\xcf\x69\xf1\x63\xa7\xdd\x12\x04\x70\x92\x74\xbe\x58\x1b\x52\xd4\x39\x2f\x92\x1c\xde\x38\xe8\xeb\x74\xbe\xbc\x45\x79\xea\xbe\x16\xb7\x24\xbb\xf1\xa5\xfe\x77\x27\x1f\xb6\x61\x53\x3a\x1c\xde\x92\x0f\x34\x68\x61\x9b\x41\x74\x9d\x70\x1f\xdf\x89\x51\x4e\xbd\x36\x9d\xc5\x15\x80\x23\x95\x3b\x93\x70\x8b\x20\x6e\x9f\x31\x94\xa5\x5c\x98\x3e\xe4\x02\xae\x3c\x46\xbb\x5f\x1d\xac\xb2\xcb\xde\x93\x82\xbe\x02\xb5\xa6\xa3\x93\xb8\xcf\x77\x13\xfb\x9c\xdd\x0f\x31\xed\xb1\xb3\x9e\xc4\xce\xd8\xbd\xd9\x53\xd9\xc5\x68\xb8\xe9\xde\x71\xd3\xc8\xf6\xaa\xd0\x68\x01\x5f\x6d\x7f\x9d\x43\xa8\xc5\x70\x33\x87\xc0\x1b\xbb\xa9\x30\x03\xba\x83\xf3\xcd\x8c\x4c\xd5\x2f\x36\x0b\xa2\x80\xd3\xb8\xbf\x1d\x56\x9f\x9f\x9f\xbd\xfb\xab\x02\x4d\x90\x66\x40\x8f\xb6\x61\xcb\x8f\x81\x10\x19\x37\x7b\x41\x91\x8b\x5d\xdf\x9b\xfe\x27\x8c\x6b\x28\x82\x26\x0e\xbd\xae\xa4\x7b\xd3\xf6\x34\x4d\xff\xd7\x27\x04\xff\xe8\x82\x95\x40\x59\x8c\x74\xe6\xf6\xef\xda\x58\xbe\x3d\x25\x2e\x21\x5c\xe8\x8d\x6b\x6f\x01\x08\x6a\x7b\x00\x90\x17\x80\x96\xa6\xeb\xbb\x75\x40\xa1\x3a\x26\xa6\xa4\xcc\xb1\xa9\xde\xe5\xf4\xda\xe1\x3f\xfc\xcd\x17\xe9\x8d\xe3\x65\xd7\x25\x47\x0b\x0d\x47\x0f\xd2\xeb\xb2\x5f\xae\x27\x60\xd2\x08\x63\xe0\x43\xc9\x03\x51\xf4\x3b\xfe\xd0\x1b\xc2\xe8\xa4\x7a\x4b\xb8\x64\x0e\xb6\x5f\xf0\xf3\x78\x91\x76\xfc\x94\x62\x76\x02\xb6\x36\x3e\xf4\x5a\xc7\xf8\x6b\x1f\x42\x41\x29\xfb\x98\xdd\x1b\xdb\x2b\xcb\xad\x18\xd0\xe9\x5d\x7a\xa6\x03\x78\xe8\x88\x09\x9d\x21\xe5\xfa\x10\x01\x9f\xd5\x2e\x39\x25\xf4\x88\x77\x02\xa7\x80\x34\x82\x82\x63\x27\x38\x5b\xc2\xac\x01\x36\xea\xf5\xf0\x70\xd6\xc8\xa8\x12\xa8\x27\x61\xd4\xd2\xbd\xe2\xf2\xdc\xdf\xa2\xa9\xa0\x28\xa9\x8e\x4b\x3f\x8d\xd9\x83\xa6\xd8\x5e\xf7\xdd\x3f\xa1\x93\x53\x73\xfc\x12\xe2\xab\xeb\x10\xe8\xad\x8c\x45\xb7\x3f\x19\x59\x83\x8e\x66\x7b\x63\xc7\xce\x20\xdd\x9c\x54\xae\x36\x9f\xee\x05\x3a\x68\xdf\xcf\xf1\x82\x24\x62\x87\xe2\x3c\x56\xdd\xcc\x58\xad\xa9\xcb\xc8\x24\x97\xb6\xb8\x9f\x69\x0f\x50\x67\x99\x9e\xaf\x56\x8b\x70\x85\x6b\xc8\x2d\x36\x29\x9d\xaf\x43\x58\xa1\xa8\x67\x1e\x41\x73\x18\xde\x12\x42\x88\x87\xe9\x33\x73\x90\xf1\xf4\x98\xc6\x97\xac\x0b\xe4\x15\x77\x28\xbe\x6d\xe0\x3e\x7a\xef\x7a\xd8\x84\x0b\xc7\xcd\x24\x72\xda\xbe\x2e\x96\xfb\xc9\x1c\x26\x6f\x7f\xd3\xd7\x8f\xa2\xd6\xc6\x95\xec\x81\x1a\x45\xb1\x09\xab\xd5\x6a\x22\xa5\xb9\x01\x81\x34\x0a\xe1\x3a\xd8\x98\x33\xd8\xd2\x5a\xd8\x9f\x52\xde\x70\x52\x67\xce\x37\x52\xd9\xeb\x36\xc2\x4d\x38\xe6\x7d\xa0\x09\xaf\xed\x79\x00\x5a\x2b\x9d\x54\x66\x77\xe9\x3c\x88\xed\x9e\xc6\x51\x77\x13\x8d\x54\x1b\xff\x5f\xa8\xe5\xe9\x29\x4d\x03\x41\x8e\x51\xd7\x42\x43\x2a\xe9\x8b\xfb\x37\x00\x00\xff\xff\x6b\xe4\x4a\xcf\xcf\x0c\x00\x00")

func styleCssBytes() ([]byte, error) {
	return bindataRead(
		_styleCss,
		"style.css",
	)
}

func styleCss() (*asset, error) {
	bytes, err := styleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "style.css", size: 0, mode: os.FileMode(420), modTime: time.Unix(0, 0)}
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
	"bad-request.tmpl.html": badRequestTmplHtml,
	"dash.js":               dashJs,
	"dashboard.tmpl.html":   dashboardTmplHtml,
	"error.tmpl.html":       errorTmplHtml,
	"head.tmpl.html":        headTmplHtml,
	"header.tmpl.html":      headerTmplHtml,
	"home.tmpl.html":        homeTmplHtml,
	"root.tmpl.html":        rootTmplHtml,
	"style.css":             styleCss,
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
	"bad-request.tmpl.html": &bintree{badRequestTmplHtml, map[string]*bintree{}},
	"dash.js":               &bintree{dashJs, map[string]*bintree{}},
	"dashboard.tmpl.html":   &bintree{dashboardTmplHtml, map[string]*bintree{}},
	"error.tmpl.html":       &bintree{errorTmplHtml, map[string]*bintree{}},
	"head.tmpl.html":        &bintree{headTmplHtml, map[string]*bintree{}},
	"header.tmpl.html":      &bintree{headerTmplHtml, map[string]*bintree{}},
	"home.tmpl.html":        &bintree{homeTmplHtml, map[string]*bintree{}},
	"root.tmpl.html":        &bintree{rootTmplHtml, map[string]*bintree{}},
	"style.css":             &bintree{styleCss, map[string]*bintree{}},
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
