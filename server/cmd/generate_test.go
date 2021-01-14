package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestGenerate_Success(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	g := genkey{length: 32}

	buf := bytes.NewBufferString("")
	command := &cobra.Command{}
	command.SetOut(buf)
	assert.NoError(g.execute(command, nil))

	assert.Regexp("^Your application key: [a-zA-Z0-9-_]{32}", buf.String())
}

func TestGenerate_KeyLengthError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	g := genkey{length: 16}

	buf := bytes.NewBufferString("")
	command := &cobra.Command{}
	command.SetOut(buf)
	err := g.execute(command, nil)
	assert.Error(err)
	assert.Equal(fmt.Sprintf("Length must be between %d and %d", MinLength, MaxLength), err.Error())
}
