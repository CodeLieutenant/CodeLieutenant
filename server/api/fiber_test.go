package api

import (
	"bytes"
	"testing"
	"time"

	"github.com/rs/zerolog"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"

	"github.com/malusev998/malusev998/container"
)

func TestFiberRegister(t *testing.T) {
	t.Parallel()
	register := func(*container.Container, *fiber.App) {
	}
	assert := require.New(t)
	payload := []struct {
		debug    bool
		register RegisterRoutesHandler
	}{
		{debug: false, register: nil},
		{debug: false, register: register},
		{debug: true, register: nil},
		{debug: true, register: register},
	}

	for _, item := range payload {
		writer := bytes.NewBuffer([]byte{})
		logger := zerolog.New(writer)

		if item.debug {
			logger.Level(zerolog.DebugLevel)
		}

		c := container.Container{
			Logger: logger,
		}

		f := &Fiber{
			app:            fiber.New(),
			address:        ":9999",
			registerRoutes: item.register,
			debug:          item.debug,
		}

		assert.Nil(f.Register(&c))

	}
}

func TestFiberListen(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	f := &Fiber{
		app:            fiber.New(),
		address:        ":9999",
		registerRoutes: nil,
		debug:          false,
	}

	time.AfterFunc(10*time.Millisecond, func() {
		assert.Nil(f.Close())
	})

	assert.Nil(f.Listen())

}
