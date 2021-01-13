package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	engine.AddFunc("now", time.Now)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public/img", "./img/")
	app.Static("/public", "./dist/")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Dusan Malusev - Index Page",
		}, "layouts/main")
	})

	log.Fatalf("Error while starting the fiber application: %v", app.Listen(":3000"))
}
