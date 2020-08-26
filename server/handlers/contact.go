package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/leebenson/conform"
	"github.com/malusev998/dusanmalusev/services"
	"github.com/malusev998/dusanmalusev/validator"
)

func contact(cs services.ContactService) func(ctx *fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		var c services.Contact

		if err := ctx.BodyParser(&c); err != nil {
			ctx.Next(fiber.NewError(400, "No payload sent"))
			return
		}

		if err := conform.Strings(&c); err != nil {
			ctx.Next(err)
			return
		}

		if err := validator.Validator.Struct(c); err != nil {
			ctx.Next(err)
			return
		}

		contactMessage, err := cs.AddMessage(c)

		if err != nil {
			ctx.Next(err)
			return
		}

		ctx.Status(201).JSON(contactMessage)
	}
}

func AddContactRoutes(g fiber.Router) {
	contactService := services.NewContactService()
	g.Post("/contact", contact(contactService))
}
