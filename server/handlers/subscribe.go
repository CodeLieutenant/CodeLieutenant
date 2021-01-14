package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/malusev998/malusev998/dto"
	"github.com/malusev998/malusev998/services/subscribe"
	"github.com/malusev998/malusev998/utils"
)

type Subscribe struct {
	Service subscribe.Service
}

func (s Subscribe) Unsubscribe(c *fiber.Ctx) error {
	return nil
}

func (s Subscribe) Subscribe(c *fiber.Ctx) error {
	fmt.Println("Hello World")
	var subDto dto.Subscription

	if err := c.BodyParser(&subDto); err != nil {
		return err
	}

	subscription, err := s.Service.Subscribe(c.Context(), subDto)
	if err != nil {
		return err
	}

	if c.Accepts(fiber.MIMEApplicationJSON) != "" {
		return c.Status(fiber.StatusCreated).JSON(subscription)
	}

	redirect := c.Context().Referer()

	if len(redirect) == 0 {
		redirect = []byte("/contact")
	}

	return c.Redirect(utils.UnsafeString(redirect))
}
