package middleware

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Timeout(d time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ch := make(chan error, 1)

		cancel := c.Locals(CancelFuncContextKey)

		if cancel != nil {
			c.Locals(CancelWillBeCalledContextKey, true)
			cancelFn := cancel.(context.CancelFunc)
			defer cancelFn()
		}

		go func() {
			err := c.Next()
			if ch != nil {
				ch <- err
			}
		}()

		select {
		case <-time.After(d):
			return fiber.ErrRequestTimeout
		case err := <-ch:
			return err
		}
	}
}
