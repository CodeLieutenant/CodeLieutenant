package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// engine := html.New("./views", ".html")
	// engine.AddFunc("now", time.Now)

	app := fiber.New(fiber.Config{
		// Views: engine,
	})

	// app.Static("/public/img", "./img/")
	// app.Static("/public", "./dist/")

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Render("index", fiber.Map{
	// 		"Title": "Dusan Malusev - Index Page",
	// 	}, "layouts/main")
	// })

	// app.Get("/contact", func(c *fiber.Ctx) error {
	// 	return c.Render("contact", fiber.Map{
	// 		"Title": "Dusan Malusev - Contact Page",
	// 	}, "layouts/main")
	// })

	// app.Get("/about", func(c *fiber.Ctx) error {
	// 	return c.Render("about", fiber.Map{
	// 		"Title": "Dusan Malusev - About me",
	// 	}, "layouts/main")
	// })

	log.Fatalf("Error while starting the fiber application: %v", app.Listen(":3000"))
}
