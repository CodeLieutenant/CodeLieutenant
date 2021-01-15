package handlers

import "github.com/gofiber/fiber/v2"

type Home struct{}

func (h Home) Home(c *fiber.Ctx) error {
	return c.Render("index", struct {
		Title string
	}{Title: "Dusan Malusev - Home"}, "layouts/main")
}
