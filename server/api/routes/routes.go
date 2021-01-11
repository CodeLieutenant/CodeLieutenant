package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/malusev998/dusanmalusev/container"
	"github.com/malusev998/dusanmalusev/handlers"
)

func RegisterRouter(c *container.Container, app *fiber.App) {
	app.Static("/public", "./public", fiber.Static{
		Browse:    false,
		Compress:  true,
		ByteRange: true,
	})

	globalGroup := app.Group("")

	registerHomeRoutes(c, globalGroup)
	registerSubscribeRoutes(c, globalGroup)
	registerContactRoutes(c, app.Group("/contact"))
}

func registerHomeRoutes(c *container.Container, app fiber.Router) {
	home := handlers.Home{}

	app.Get("/", home.Home)
}

func registerSubscribeRoutes(c *container.Container, app fiber.Router) {
	sub := handlers.Subscribe{
		Service: c.GetSubscriptionService(),
	}

	app.Post("/subscribe", sub.Subscribe)
	app.Get("/unsubscribe", sub.Unsubscribe)
}

func registerContactRoutes(c *container.Container, router fiber.Router) {
	contact := handlers.Contact{
		Service: c.GetContactService(),
	}

	router.Get("/", contact.Index)
	router.Post("/", contact.Message)
}
