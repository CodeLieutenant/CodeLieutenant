package container

import (
	"context"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/blake2b"

	"github.com/malusev998/dusanmalusev/config"
	"github.com/malusev998/dusanmalusev/database"
	"github.com/malusev998/dusanmalusev/services"
	"github.com/malusev998/dusanmalusev/services/email"
	"github.com/malusev998/dusanmalusev/services/subscribe"
	"github.com/malusev998/dusanmalusev/utils"
	"github.com/malusev998/dusanmalusev/validators"
)

type Container struct {
	Ctx    context.Context
	Logger zerolog.Logger
	DB     *pgxpool.Pool
	Config *config.Config

	contactService      services.ContactService
	validator           *validator.Validate
	translator          ut.Translator
	subscriptionService subscribe.Service
	session             *session.Store
	urlSigner           utils.URLSigner
}

func (c *Container) GetDatabasePool() *pgxpool.Pool {
	if c.DB == nil {
		var err error
		c.DB, err = database.ConnectDB(c.Ctx, database.Config{
			URL:                   c.Config.Database.URI,
			Host:                  c.Config.Database.Host,
			User:                  c.Config.Database.User,
			Password:              c.Config.Database.Password,
			DbName:                c.Config.Database.DBName,
			TimeZone:              c.Config.Database.TimeZone,
			MaxConnectionLifetime: c.Config.Database.MaxConnectionIdleTime,
			MaxConnectionIdleTime: c.Config.Database.MaxConnectionIdleTime,
			HealthCheck:           c.Config.Database.HealthCheck,
			MaxConns:              c.Config.Database.MaxConns,
			MinConns:              c.Config.Database.MinConns,
			Port:                  uint16(c.Config.Database.Port),
			SslMode:               c.Config.Database.SSLMode,
			Lazy:                  c.Config.Database.Lazy,
		}, c.Logger)

		if err != nil {
			log.Fatal().Err(err).Msg("Error while connecting to database")
			c.Logger.Fatal().Err(err).Msg("Error while connecting to database")
		}
	}

	return c.DB
}

func (c *Container) GetEmailService() email.Interface {
	service, err := email.NewEmailService(email.Config{
		Addr: "",
		From: "",
		// Auth:     smtp.PlainAuth(),
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
		c.contactService = services.NewContactService(c.DB, c.GetValidator())
	}

	return c.contactService
}

func (c *Container) GetSubscriptionService() subscribe.Service {
	if c.contactService == nil {
		if c.Config.Subscription.SendEmail {
			c.subscriptionService = subscribe.NewSubscriptionWithEmail(c.GetEmailService(), c.GetDatabasePool(), c.GetValidator())
		} else {
			c.subscriptionService = subscribe.NewSubscriptionService(c.GetDatabasePool(), c.GetValidator())
		}
	}

	return c.subscriptionService
}

func (c *Container) GetURLSigner() utils.URLSigner {
	if c.urlSigner == nil {
		h, err := blake2b.New512(c.Config.Key)
		if err != nil {
			c.Logger.Fatal().Err(err).Msg("Cannot create blake2b algorithm")
		}

		c.urlSigner = utils.NewURLSigner(h)
	}

	return c.urlSigner
}

func (c *Container) GetValidator() *validator.Validate {
	if c.validator == nil && c.translator == nil {
		english := en.New()
		uni := ut.New(english, english)

		c.translator, _ = uni.GetTranslator(c.Config.Locale)
		c.validator = validator.New()

		if err := en_translations.RegisterDefaultTranslations(c.validator, c.translator); err != nil {
			c.Logger.Fatal().Err(err).Msg("Error while registering english translations")
		}

		if err := validators.Register(c.validator, c.translator); err != nil {
			c.Logger.Fatal().Err(err).Msg("Error while registering custom validators")
		}
	}

	return c.validator
}

func (c *Container) GetTranslator() ut.Translator {
	if c.translator == nil {
		c.GetValidator()
	}

	return c.translator
}

func (c *Container) GetStorage(database int) fiber.Storage {
	return redis.New(redis.Config{
		Host:     c.Config.Redis.Host,
		Port:     c.Config.Redis.Port,
		Username: c.Config.Redis.Username,
		Password: c.Config.Redis.Password,
		Database: database,
	})
}

func (c *Container) GetSession() *session.Store {
	if c.session == nil {
		c.session = session.New(session.Config{
			Storage:        c.GetStorage(0),
			CookieHTTPOnly: true,
			Expiration:     c.Config.Session.Expiration,
			CookieName:     c.Config.Session.CookieName,
			CookieDomain:   c.Config.Session.CookieDomain,
			CookiePath:     c.Config.Session.CookiePath,
			CookieSecure:   c.Config.Session.Secure,
			CookieSameSite: "Lax",
			KeyGenerator:   utils.DefaultStringGenerator,
		})
	}

	return c.session
}

func (c *Container) Close() error {
	c.DB.Close()

	return nil
}
