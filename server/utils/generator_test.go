package utils_test

import (
	"encoding/base64"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/server/utils"
)

func TestGenerateDefaultLengthString(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	str := utils.DefaultStringGenerator()

	bytes, err := base64.RawURLEncoding.DecodeString(str)

	assert.NoError(err)
	assert.NotEmpty(bytes)
	assert.Len(bytes, utils.DefaultBytes)
}

func TestGenerateRandomBytes(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	tests := []struct {
		length int64
	}{
		{length: 10},
		{length: 32},
		{length: 64},
	}

	for _, test := range tests {
		t.Run("Test_Length_"+strconv.FormatInt(test.length, 10), func(t *testing.T) {
			bytes := utils.UniqueBytesGenerator(int(test.length))
			assert.NotEmpty(bytes)
			assert.Len(bytes, int(test.length))
		})
	}
}
