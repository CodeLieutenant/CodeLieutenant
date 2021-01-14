package utils

import (
	"crypto/hmac"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"net/url"
)

type URLSigner interface {
	Sign(string, ...interface{}) (string, error)
	Verify(string) error
}

type signer struct {
	h hash.Hash
}

func NewURLSigner(h hash.Hash) signer {
	return signer{
		h: h,
	}
}

func (s signer) Sign(format string, values ...interface{}) (string, error) {
	str := fmt.Sprintf(format, values...)

	hashEncoded := base64.RawURLEncoding.EncodeToString(s.h.Sum(UnsafeBytes(str)))

	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}

	query := u.Query()

	query["signature"] = []string{hashEncoded}

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func (s signer) Verify(urlStr string) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	query := u.Query()

	signature, ok := query["signature"]

	if !ok || (len(signature) > 0 && signature[0] == "") {
		return errors.New("signature not found")
	}

	delete(query, "signature")

	u.RawQuery = query.Encode()

	bytes, err := base64.RawURLEncoding.DecodeString(signature[0])
	if err != nil {
		return err
	}

	str, _ := url.QueryUnescape(u.String())
	calculated := s.h.Sum(UnsafeBytes(str))

	if !hmac.Equal(calculated, bytes) {
		return errors.New("signature is invalid")
	}

	return nil
}
