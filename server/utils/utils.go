package utils

import (
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"unsafe"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/blake2b"
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

func LimiterKeyGenerator(key []byte) func(c *fiber.Ctx) string {
	blakeHash, err := blake2b.New256(key)
	if err != nil {
		panic(err.Error())
	}

	return func(c *fiber.Ctx) string {
		userAgent := c.Context().UserAgent()
		ip := UnsafeBytes(c.IP())

		data := make([]byte, 0, len(userAgent)+len(ip))

		copy(data, userAgent)
		copy(data, ip)

		return base64.RawStdEncoding.EncodeToString(blakeHash.Sum(data))
	}
}
