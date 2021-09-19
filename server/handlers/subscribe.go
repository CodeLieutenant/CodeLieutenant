package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/malusev998/malusev998/server/dto"
	"github.com/malusev998/malusev998/server/services"
	"github.com/malusev998/malusev998/server/utils"
)

type Subscribe struct {
	FrontendRedirectURL string
	BaseURL             string
	SubscriptionService services.SubscribeService
	Signer              utils.URLSigner
}

func (s Subscribe) Unsubscribe(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(c.UserContext())
	defer cancel()

	err := s.Signer.Verify(fmt.Sprintf("%s%s", s.BaseURL, c.OriginalURL()))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("Invalid URL")
	}

	idStr := c.Query("id")

	// Error should never happen here, URL signature has been checked above
	// Nobody could have tempered with it at this point
	id, _ := strconv.ParseUint(idStr, 10, 64)

	err = s.SubscriptionService.Unsubscribe(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString("An error has occurred")
	}

	redirect, _ := Redirect(s.FrontendRedirectURL, "Successfully unsubscribed")

	return c.
		Redirect(redirect, fiber.StatusTemporaryRedirect)
}

func (s Subscribe) Subscribe(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(c.UserContext())
	defer cancel()

	var subDto dto.Subscription

	if err := c.BodyParser(&subDto); err != nil {
		return err
	}

	subscription, err := s.SubscriptionService.Subscribe(ctx, subDto)
	if err != nil {
		return err
	}

	redirect := c.Query("redirect", "")

	if redirect == "" {
		return c.Status(fiber.StatusCreated).JSON(subscription)
	}

	redirect, err = Redirect(redirect)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Redirect(redirect, fiber.StatusTemporaryRedirect)
}
