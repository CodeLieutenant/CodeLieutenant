package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
	"github.com/pkg/errors"

	"github.com/malusev998/dusanmalusev/container"
	"github.com/malusev998/dusanmalusev/utils"
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
	f.app.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return utils.UniqueStringGenerator(64)
		},
		ContextKey: "request_id",
	}))

	if !f.debug {
		c.Logger.Debug().Msg("Running in production mode, recover and compression middleware are enabled")

		f.app.Use(recover.New())
		f.app.Use(compress.New(compress.Config{
			Next:  nil,
			Level: compress.LevelBestSpeed,
		}))
	} else {
		c.Logger.Debug().Msg("Running in DEBUG mode, PProf and Monitor (GET /monitor) are enabled")
		f.app.Use(pprof.New())
		f.app.Get("/monitor", monitor.New())
		f.app.Use(logger.New())
	}

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
