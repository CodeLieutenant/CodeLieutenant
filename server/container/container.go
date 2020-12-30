package container

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/malusev998/dusanmalusev/services"
)

type Container struct {
	Ctx                 context.Context
	Logger              *zerolog.Logger
	DB                  *gorm.DB
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
