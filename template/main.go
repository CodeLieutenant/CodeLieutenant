package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/malusev998/template/jet"
)

func main() {
	engine := jet.NewFileSystem(http.Dir("./views"), ".jet")
	engine.AddFunc("now", time.Now)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public/img", "./img/")
	app.Static("/public", "./dist/")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Dusan Malusev - Index Page",
		})
	})

	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.Render("contact", fiber.Map{
			"Title": "Dusan Malusev - Contact Page",
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"Title": "Dusan Malusev - About me",
		})
	})

	log.Fatalf("Error while starting the fiber application: %v", app.Listen(":3000"))
}
