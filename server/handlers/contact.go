package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/services"
	"github.com/malusev998/malusev998/server/utils"
)

type Contact struct {
	Service services.ContactService
}

func (c Contact) Index(ctx *fiber.Ctx) error {
	return ctx.Render("contact", fiber.Map{"Title": "Dusan Malusev - Contact"})
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
