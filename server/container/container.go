package container

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/go-playground/validator/v10"
	"github.com/malusev998/dusanmalusev/services"
	"github.com/rs/zerolog"
)

type Container struct {
	Ctx                 context.Context
	Logger              *zerolog.Logger
	DB                  *pgxpool.Pool
	Validator           *validator.Validate
	contactService      services.ContactService
	subscriptionService services.SubscriptionService
}

func (c *Container) GetContactService() services.ContactService {
	if c.contactService == nil {
		c.contactService = services.NewContactService(c.DB, c.Validator)
	}

	return c.contactService
}

func (c *Container) GetSubscriptionService() services.SubscriptionService {
	if c.contactService == nil {
		c.subscriptionService = services.NewSubscriptionService(c.DB, c.Validator)
	}

	return c.subscriptionService
}
