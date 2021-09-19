package container

import (
	"golang.org/x/crypto/blake2b"

	"github.com/malusev998/malusev998/server/utils"
)

func (c *Container) GetURLSigner() utils.URLSigner {
	if c.urlSigner == nil {
		h, err := blake2b.New512(c.Config.Key)
		if err != nil {
			c.Logger.Fatal().Err(err).Msg("Cannot create blake2b algorithm")
		}

		c.urlSigner = utils.NewURLSigner(h)
	}

	return c.urlSigner
}
