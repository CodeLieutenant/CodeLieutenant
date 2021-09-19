package container

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"

	"github.com/malusev998/malusev998/server/utils"
)

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
			KeyLookup:      c.Config.Session.CookieName,
			CookieDomain:   c.Config.Session.CookieDomain,
			CookiePath:     c.Config.Session.CookiePath,
			CookieSecure:   c.Config.Session.Secure,
			CookieSameSite: "Lax",
			KeyGenerator:   utils.DefaultStringGenerator,
		})
	}

	return c.session
}
