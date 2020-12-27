package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSignUrl(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	key := []byte("Test Key")

	signedUrl, err := SignURL("http://localhost:3000/subscribe?email=%s", key, "test@gmail.com")

	assert.Nil(err)
	assert.Regexp("^http://localhost:3000/subscribe\\?email=test%40gmail\\.com&signature=[a-zA-Z0-9-_]+$", signedUrl)
}

func TestVerify(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	key := []byte("Test Key")

	signedUrl := "http://localhost:3000/subscribe?email=test%40gmail.com&signature=aHR0cDovL2xvY2FsaG9zdDozMDAwL3N1YnNjcmliZT9lbWFpbD10ZXN0QGdtYWlsLmNvbQB5azgGygndjdUFS67Qb51rYaQT6Jk3huTeOYBMIiLE"

	assert.False(VerifyURL(signedUrl, key))
}
