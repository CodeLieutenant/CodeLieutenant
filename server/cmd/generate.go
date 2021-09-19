package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/malusev998/malusev998/server/utils"
)

const (
	MinLength int8 = 32
	MaxLength int8 = 64
)

type genkey struct {
	length int8
}

func newGenerateKeyCommand() *cobra.Command {
	g := genkey{}
	command := &cobra.Command{
		Use:   "generate",
		Short: "Generates application key used in encryption and digital signatures",
		RunE:  g.execute,
	}

	command.Flags().Int8VarP(
		&g.length,
		"length",
		"l",
		MaxLength,
		fmt.Sprintf("Length of the generated key (in bytes) (between %d and %d)", MinLength, MaxLength),
	)

	return command
}

func (g *genkey) execute(cmd *cobra.Command, args []string) error {
	if g.length < MinLength || g.length > MaxLength {
		return fmt.Errorf("length must be between %d and %d", MinLength, MaxLength)
	}

	str := utils.UniqueStringGenerator(int(g.length))
	cmd.Printf("Your application key: %s\n", str)

	return nil
}
