//go:build !race
// +build !race

package middleware

import (
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestTimeout_Excited(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	app.Use(Timeout(2 * time.Millisecond))

	app.Get("/", func(ctx *fiber.Ctx) error {
		time.Sleep(5 * time.Millisecond)
		return nil
	})

	req, _ := http.NewRequest(fiber.MethodGet, "/", nil)

	res, err := app.Test(req)

	assert.NoError(err)
	assert.Equal(fiber.StatusRequestTimeout, res.StatusCode)
}

func TestTimeout_Expire(t *testing.T) {
	t.Parallel()

	assert := require.New(t)

	app := fiber.New()

	app.Use(Timeout(2 * time.Millisecond))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	req, _ := http.NewRequest(fiber.MethodGet, "/", nil)

	res, err := app.Test(req)

	assert.NoError(err)
	assert.Equal(fiber.StatusOK, res.StatusCode)
}

func TestTimeout_Cancel(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	app := fiber.New()

	app.Use(Context)
	app.Use(Timeout(2 * time.Millisecond))

	app.Get("/", func(ctx *fiber.Ctx) error {
		<-ctx.UserContext().Done()
		assert.True(true)

		<-time.After(5 * time.Millisecond)

		return nil
	})

	req, _ := http.NewRequest(fiber.MethodGet, "/", nil)

	res, err := app.Test(req)

	assert.NoError(err)
	assert.Equal(fiber.StatusRequestTimeout, res.StatusCode)
}
