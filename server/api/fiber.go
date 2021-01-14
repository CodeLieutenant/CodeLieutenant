package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/pkg/errors"

	"github.com/malusev998/malusev998/container"
)

type RegisterRoutesHandler func(*container.Container, *fiber.App)

type Fiber struct {
	app            *fiber.App
	address        string
	registerRoutes RegisterRoutesHandler
	debug          bool
}

func NewFiberAPI(
	address string,
	prefork,
	debug bool,
	errorHandler fiber.ErrorHandler,
	register RegisterRoutesHandler,
) Interface {
	engine := html.New("views", ".html")
	engine.AddFunc("now", time.Now)

	return Fiber{
		app: fiber.New(fiber.Config{
			Prefork:      prefork,
			ErrorHandler: errorHandler,
			Views:        engine,
		}),
		debug:          debug,
		address:        address,
		registerRoutes: register,
	}
}

func (f Fiber) Register(c *container.Container) error {
	if fiber.IsChild() {
		c.Logger.Debug().Msg("Starting the preforked process")
	} else {
		c.Logger.Debug().Msg("Starting the main application")
	}

	if f.registerRoutes != nil {
		c.Logger.Debug().Msg("Loading the routes")
		f.registerRoutes(c, f.app)
	}

	return nil
}

func (f Fiber) Listen() error {
	if err := f.app.Listen(f.address); err != nil {
		return errors.Wrap(err, "Error while starting application")
	}
	return nil
}

func (f Fiber) Close() error {
	return f.app.Shutdown()
}
