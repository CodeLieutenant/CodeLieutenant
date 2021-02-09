package handlers

import "github.com/gofiber/fiber/v2"

type Home struct{}

func (h Home) Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{"Title": "Dusan Malusev - Home"})
}

func (h Home) About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title": "Dusan Malusev - About me",
	})
}
