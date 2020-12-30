package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/malusev998/dusanmalusev/dto"
	"github.com/malusev998/dusanmalusev/services"
	"github.com/malusev998/dusanmalusev/utils"
)

type Contact struct {
	ContactService services.ContactService
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

	contact, err := c.ContactService.AddMessage(contactDto)

	if err != nil {
		return err
	}

	if ctx.XHR() && ctx.Accepts(fiber.MIMEApplicationJSON) != "" {
		return ctx.Status(fiber.StatusCreated).JSON(contact)
	}

	return ctx.Redirect(utils.UnsafeString(ctx.Context().Referer()))
}
