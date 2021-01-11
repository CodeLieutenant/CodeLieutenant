package container

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/malusev998/dusanmalusev/config"
	"github.com/malusev998/dusanmalusev/services"
	"github.com/malusev998/dusanmalusev/services/email"
	"github.com/malusev998/dusanmalusev/services/subscribe"
	"github.com/rs/zerolog"
)

type Container struct {
	Ctx       context.Context
	Logger    *zerolog.Logger
	DB        *pgxpool.Pool
	Validator *validator.Validate
	Config    *config.Config

	contactService      services.ContactService
	subscriptionService subscribe.Service
}

func (c *Container) GetEmailService() email.Interface {
	service, err := email.NewEmailService(email.Config{
		Addr:     "",
		From:     "",
		//Auth:     smtp.PlainAuth(),
		TLS:      nil,
		Logger:   c.Logger,
		PoolSize: 0,
		Senders:  0,
	})

	if err != nil {
		panic(err.Error())
	}

	return service
}

func (c *Container) GetContactService() services.ContactService {
	if c.contactService == nil {
		c.contactService = services.NewContactService(c.DB, c.Validator)
	}

	return c.contactService
}

func (c *Container) GetSubscriptionService() subscribe.Service {
	if c.contactService == nil {
		if c.Config.Subscription.SendEmail {
			c.subscriptionService = subscribe.NewSubscriptionWithEmail(c.GetEmailService(), c.DB, c.Validator)
		} else {
			c.subscriptionService = subscribe.NewSubscriptionService(c.DB, c.Validator)
		}
	}

	return c.subscriptionService
}
