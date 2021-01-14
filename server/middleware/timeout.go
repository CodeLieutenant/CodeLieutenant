package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Timeout(d time.Duration, h ...fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ch := make(chan error, 1)
		defer close(ch)

		go func() {
			if len(h) > 0 {
				ch <- h[0](c)
			} else {
				ch <- c.Next()
			}
		}()

		select {
		case err := <-ch:
			return err
		case <-time.After(d):
			return fiber.ErrRequestTimeout
		}
	}
}
