package api

import (
	"io"

	"github.com/malusev998/dusanmalusev/container"
)

type (
	Interface interface {
		io.Closer
		Register(c *container.Container) error
		Listen() error
	}
)
