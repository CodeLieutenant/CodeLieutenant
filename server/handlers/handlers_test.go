package handlers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedirect(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	url, err := Redirect("/")

	assert.NoError(err)
	assert.Equal("/?message=Message+sent&status=success", url)
}
