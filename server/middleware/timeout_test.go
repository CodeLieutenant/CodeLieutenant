// +build !race

package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"

	"github.com/malusev998/malusev998/server/middleware"
)

type TimeoutTest struct {
	suite.Suite
	duration, sleep time.Duration
	httpStatus      int
	app             *fiber.App
}

func (t *TimeoutTest) SetupTest() {
	t.app = fiber.New()
	t.app.Get("/wrapper", middleware.Timeout(t.duration, func(c *fiber.Ctx) error {
		time.Sleep(t.sleep)
		c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Everything ok"})
		return nil
	}))
	t.app.Use(middleware.Timeout(t.duration))
	t.app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(t.sleep)
		c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Everything ok"})
		return nil
	})

}

func (t *TimeoutTest) TestMiddleware() {
	t.T().Parallel()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, err := t.app.Test(req)

	t.Nil(err)
	t.NotNil(res)
	t.Equal(t.httpStatus, res.StatusCode)
}

func (t *TimeoutTest) TestHandlerWrapper() {
	t.T().Parallel()
	req := httptest.NewRequest(http.MethodGet, "/wrapper", nil)
	res, err := t.app.Test(req)

	t.Nil(err)
	t.NotNil(res)
	t.Equal(t.httpStatus, res.StatusCode)
}

func TestTimeoutSuccess(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TimeoutTest{
		duration:   20 * time.Millisecond,
		sleep:      10 * time.Millisecond,
		httpStatus: http.StatusOK,
	})
}

func TestTimeoutExpired(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TimeoutTest{
		duration:   20 * time.Millisecond,
		sleep:      50 * time.Millisecond,
		httpStatus: http.StatusRequestTimeout,
	})
}
