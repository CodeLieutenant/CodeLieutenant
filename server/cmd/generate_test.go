package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate_Success(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	command := newGenerateKeyCommand()
	buf := bytes.NewBufferString("")
	command.SetOut(buf)
	command.Flag("length").Value.Set("32")
	assert.NoError(command.Execute())

	assert.Regexp("^Your application key: [a-zA-Z0-9-_]{32}", buf.String())
}

func TestGenerate_KeyLengthError(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	command := newGenerateKeyCommand()
	buf := bytes.NewBufferString("")
	command.SetOut(buf)
	command.Flag("length").Value.Set("16")
	command.SetOut(buf)
	err := command.Execute()
	assert.Error(err)
	assert.Equal(fmt.Sprintf("Length must be between %d and %d", MinLength, MaxLength), err.Error())
}
