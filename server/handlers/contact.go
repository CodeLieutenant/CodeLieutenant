package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/malusev998/malusev998/dto"
	"github.com/malusev998/malusev998/services"
	"github.com/malusev998/malusev998/utils"
)

type Contact struct {
	Service services.ContactService
}

func (c Contact) Index(ctx *fiber.Ctx) error {
	return ctx.Render("contact", struct {
		Title string
	}{
		Title: "Contact Page",
	}, "layouts/main")
}

func (c Contact) Message(ctx *fiber.Ctx) error {
	contactDto := dto.Contact{}

	if err := ctx.BodyParser(&contactDto); err != nil {
		return err
	}

	contact, err := c.Service.AddMessage(ctx.Context(), contactDto)

	if err != nil {
		return err
	}

	if ctx.Accepts(fiber.MIMEApplicationJSON) != "" {
		return ctx.Status(fiber.StatusCreated).JSON(contact)
	}

	redirect := ctx.Context().Referer()

	if len(redirect) == 0 {
		redirect = []byte("/contact")
	}

	return ctx.Redirect(utils.UnsafeString(redirect))
}
