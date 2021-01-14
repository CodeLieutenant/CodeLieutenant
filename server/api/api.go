package api

import (
	"io"

	"github.com/malusev998/malusev998/container"
)

type (
	Interface interface {
		io.Closer
		Register(c *container.Container) error
		Listen() error
	}
)
