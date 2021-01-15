package cmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/config"
)

func TestExecute(t *testing.T) {
	t.Parallel()

	assert := require.New(t)

	buf := bytes.NewBufferString("")
	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"--version"})
	assert.NoError(Execute(context.Background(), "v1.0.0", &config.Config{}, log.Logger))

	assert.Contains(buf.String(), "v1.0.0")
}
