package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/services"
)

type Contact struct {
	Service services.ContactService
}

func (c Contact) Message(ctx *fiber.Ctx) error {
	contactDto := dto.Contact{}
	context, cancel := context.WithCancel(ctx.UserContext())
	defer cancel()

	if err := ctx.BodyParser(&contactDto); err != nil {
		return err
	}

	contact, err := c.Service.AddMessage(context, contactDto)

	if err != nil {
		return err
	}

	redirect := ctx.Query("redirect", "")

	if redirect == "" {
		return ctx.Status(fiber.StatusCreated).JSON(contact)
	}

	redirect, err = Redirect(redirect)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: err.Error(),
		})

	}

	return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
}
