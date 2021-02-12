package utils

import (
	"crypto/rand"
	"encoding/base64"
)

const DefaultBytes = 32

func DefaultStringGenerator() string {
	return UniqueStringGenerator(DefaultBytes)
}

// UniqueStringGenerator - Generates crypto secure base64 encoded string
func UniqueStringGenerator(length int) string {
	return base64.RawURLEncoding.EncodeToString(UniqueBytesGenerator(length))
}

// UniqueBytesGenerator generates crypto random bytes
func UniqueBytesGenerator(length int) []byte {
	bytes := make([]byte, length)
	n, err := rand.Read(bytes)

	if err != nil || n != length {
		return nil
	}

	return bytes
}
