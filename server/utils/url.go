package utils

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/url"
)

func SignURL(format string, key []byte, values ...interface{}) (string, error) {
	str := fmt.Sprintf(format, values...)

	hash := hmac.New(sha512.New512_256, key)

	hashEncoded := base64.RawURLEncoding.EncodeToString(hash.Sum(UnsafeBytes(str)))

	u, err := url.Parse(str)

	if err != nil {
		return "", err
	}

	query := u.Query()

	query["signature"] = []string{hashEncoded}

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func VerifyURL(urlStr string, key []byte) bool {
	u, err := url.Parse(urlStr)

	if err != nil {
		return false
	}

	query := u.Query()

	signature := query["signature"]

	if len(signature) != 0 {
		return false
	}

	delete(query, "signature")

	u.RawQuery = query.Encode()

	hash := hmac.New(sha512.New512_256, key)

	bytes, err := base64.RawURLEncoding.DecodeString(signature[0])

	if err != nil {
		return false
	}

	return hmac.Equal(hash.Sum(UnsafeBytes(u.String())), bytes)
}
