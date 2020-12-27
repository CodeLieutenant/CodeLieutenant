package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// UniqueStringGenerator - Generates crypto secure base64 encoded string
func UniqueStringGenerator(len int) string {
	return base64.RawStdEncoding.EncodeToString(UniqueBytesGenerator(len))
}

// UniqueBytesGenerator generates crypto random bytes
func UniqueBytesGenerator(len int) []byte {
	bytes := make([]byte, len)
	n, err := rand.Read(bytes)

	if err != nil || n != len {
		return nil
	}

	return bytes
}
