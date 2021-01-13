package utils

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"unsafe"
)

// #nosec G103
// GetBytes returns a byte pointer without allocation
func UnsafeBytes(s string) (bs []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return
}

// #nosec G103
// GetString returns a string pointer without allocation
func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}


func CreateFile(path string) (file io.WriteCloser, err error) {
	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			return nil, err
		}
	}

	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, 0744); err != nil {
		return nil, err
	}

	file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)

	if err != nil {
		return nil, err
	}

	return
}
