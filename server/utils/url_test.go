package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUrl_InvalidUrl(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	key := []byte("Test Key")
	h := hmac.New(sha256.New, key)

	signedUrl, err := NewURLSigner(h).Sign("h😊tp:/")
	assert.Error(err)
	assert.Empty(signedUrl)
}

func TestSignUrl(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	key := []byte("Test Key")
	h := hmac.New(sha256.New, key)

	signedUrl, err := NewURLSigner(h).Sign("http://localhost:3000/subscribe?email=%s", "test@gmail.com")

	assert.Nil(err)
	fmt.Println(signedUrl)
	assert.Regexp("^http://localhost:3000/subscribe\\?email=test%40gmail\\.com&signature=[a-zA-Z0-9-_]+$", signedUrl)
}

func TestVerify(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)
	key := []byte("Test Key")
	h := hmac.New(sha512.New512_256, key)
	tests := []struct {
		name  string
		url   string
		valid bool
	}{
		{
			name:  "Valid",
			url:   "http://localhost:3000/subscribe?email=test%40gmail.com&signature=aHR0cDovL2xvY2FsaG9zdDozMDAwL3N1YnNjcmliZT9lbWFpbD10ZXN0QGdtYWlsLmNvbQB5azgGygndjdUFS67Qb51rYaQT6Jk3huTeOYBMIiLE",
			valid: true,
		},
		{
			name:  "Invalid",
			url:   "http://localhost:3000/subscribe?email=test%40gmail.com&signature=aHR1cDovL2xvY2FsaG9zdDozMDAwL3N1YnNjcmliZT9lbWFpbD10ZXN0QGdtYWlsLmNvbQB5azgGygndjdUFS67Qb51rYaQT6Jk3huTeOYBMIiLE",
			valid: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.valid {
				ass.NoError(NewURLSigner(h).Verify(test.url))
			} else {
				err := NewURLSigner(h).Verify(test.url)
				ass.Error(err)
				ass.Contains(err.Error(), "signature is invalid")
			}
		})
	}
}

func TestVerify_InvalidBase64Signature(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	key := []byte("Test Key")
	h := hmac.New(sha256.New, key)
	signedUrl := "http://localhost:3000/subscribe?email=test%40gmail.com&signature=aHR0cDovL2xv Y2FsaG9zdDozMDAwL3N1YnNjcmliZT9lbWFpbD10ZXN0QGdtYWlsLmNvbQB5azgGygndjdUFS67Qb51rYaQT6Jk3huTeOYBMIiLE"

	err := NewURLSigner(h).Verify(signedUrl)
	assert.Error(err)
	assert.Contains(err.Error(), "illegal base64 data")
}

func TestVerify_NoSignatureFound(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	key := []byte("Test Key")
	h := hmac.New(sha256.New, key)
	tests := []struct {
		name string
		url  string
	}{
		{
			name: "EmptySignature",
			url:  "http://localhost:3000/subscribe?email=test%40gmail.com&signature=",
		},
		{
			name: "NoSignatureKey",
			url:  "http://localhost:3000/subscribe?email=test%40gmail.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := NewURLSigner(h).Verify(test.url)
			assert.Error(err)
			assert.Contains(err.Error(), "signature not found")
		})
	}
}
